package internal

import (
	"errors"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"slices"
	"strings"

	"github.com/h2non/filetype"
	"golang.org/x/sync/errgroup"
)

type WorkMode string
type PrintMode string
type FileStatus int

const (
	WM_SCAN WorkMode = "scan"
	WM_FIX  WorkMode = "fix"
)

const (
	PM_NONE      PrintMode = "none"
	PM_ALL       PrintMode = "all"
	PM_IMPORTANT PrintMode = "important"
	PM_REPORT    PrintMode = "report"
)

const (
	FS_NONE         FileStatus = iota
	FS_FIX_REQUIRED FileStatus = iota
	FS_IGNORED      FileStatus = iota
	FS_FIXED        FileStatus = iota
	FS_ERROR        FileStatus = iota
)

const FILE_PROCESSING_LIMIT = 100

type _FileInfo struct {
	filePath     string
	fileTrimPath string
	mime         string
	origExt      string
	realExt      string
	status       FileStatus
	notes        string
}

func Scan(dirPath string, exts string, printMode PrintMode) error {
	return worker(dirPath, exts, WM_SCAN, printMode)
}

func Fix(dirPath string, exts string, printMode PrintMode) error {
	return worker(dirPath, exts, WM_FIX, printMode)
}

func worker(dirPath string, exts string, workMode WorkMode, printMode PrintMode) error {
	tb := NewLogger(printMode)
	tb.PrintIntro(dirPath, exts, workMode)
	tb.PrintTableHeader()

	dirPathAbs, err := filepath.Abs(dirPath)
	if err != nil {
		return err
	}
	if _, err := os.Stat(dirPathAbs); err != nil {
		if os.IsNotExist(err) {
			return fmt.Errorf("%s doesn't exist", dirPathAbs)
		}
		return err
	}

	var whiteExtList []string
	if len(exts) > 0 {
		whiteExtList = strings.Split(exts, ",")
	}

	var eg errgroup.Group
	eg.SetLimit(FILE_PROCESSING_LIMIT)
	if err := filepath.WalkDir(dirPathAbs, func(filePath string, dirEntry fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if dirEntry.IsDir() || filePath == "" {
			return nil
		}

		eg.Go(func() error {
			fileInfo, err := getFileInfo(filePath)
			if err != nil {
				return err
			}
			fileInfo.fileTrimPath = strings.Replace(filePath, dirPathAbs, ".", 1)

			if fileInfo.mime != "" && fileInfo.origExt != fileInfo.realExt {
				if len(whiteExtList) == 0 || slices.ContainsFunc(whiteExtList, func(s string) bool { return "."+s == fileInfo.realExt }) {
					fileInfo.status = FS_FIX_REQUIRED
				} else {
					fileInfo.status = FS_IGNORED
				}
			}
			if fileInfo.status == FS_FIX_REQUIRED && workMode == WM_FIX {
				if err := fixFileExt(&fileInfo); err != nil {
					return err
				}
			}

			tb.PrintTableRow(&fileInfo)
			return nil
		})
		return nil
	}); err != nil {
		return err
	}

	if err = eg.Wait(); err != nil {
		return err
	}

	tb.PrintResult()

	return nil
}

func getFileInfo(filePath string) (_FileInfo, error) {
	fileInfo := _FileInfo{
		filePath: filePath,
		origExt:  filepath.Ext(filePath),
	}

	file, err := os.Open(filePath)
	if err != nil {
		return _FileInfo{}, err
	}
	defer file.Close()

	head := make([]byte, 261)
	_, err = file.Read(head)
	if err != nil {
		if errors.Is(err, io.EOF) {
			fileInfo.status = FS_ERROR
			fileInfo.notes = "File is empty"
			return fileInfo, nil
		}
		/* TODO: process more expected errors:
		- Access is denied.
		- The process cannot access the file because it is being used by another process.
		*/
		return _FileInfo{}, err
	}

	fileType, err := filetype.Get(head)
	if err != nil {
		return _FileInfo{}, err
	}

	if fileType.MIME.Value != "" {
		fileInfo.mime = fileType.MIME.Value
		fileInfo.realExt = "." + fileType.Extension
	}

	return fileInfo, nil
}

func fixFileExt(fileInfo *_FileInfo) error {
	oldPath := fileInfo.filePath
	newPath := strings.TrimSuffix(fileInfo.filePath, fileInfo.origExt) + fileInfo.realExt

	fileStat, err := os.Stat(newPath)
	if err != nil && !os.IsNotExist(err) {
		return err
	}
	if fileStat != nil {
		fileInfo.status = FS_ERROR
		fileInfo.notes = "File with the same name already exists"
		return nil
	}
	if err = os.Rename(oldPath, newPath); err != nil {
		return err
	}
	fileInfo.status = FS_FIXED
	fileInfo.notes = ""

	return nil
}
