package gofixt

import (
	"context"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/h2non/filetype"
	"golang.org/x/sync/errgroup"
)

type _FileInfo struct {
	filePath    string
	filePathCut string
	mime        string
	oExt        string
	realExt     string
	fixRequired bool
	fixed       bool
	err         string
}

func Scan(dirPath string, silent bool) error {
	return worker(dirPath, silent, false)
}

func Fix(dirPath string, silent bool) error {
	return worker(dirPath, silent, true)
}

func worker(dirPath string, silent bool, needFix bool) error {
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

	var mutex sync.Mutex
	ctx := context.TODO()
	errGrp, _ := errgroup.WithContext(ctx)
	start := time.Now()
	tableWriter := NewTableWriter()
	if !silent {
		tableWriter.AddHeader()
	}

	var procFileCnt int
	var fixedFileCnt int
	err = filepath.WalkDir(dirPathAbs, func(filePath string, dirEntry fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if dirEntry.IsDir() || filePath == "" {
			return nil
		}
		errGrp.Go(func() error {
			fileInfo, err := getFileInfo(dirPathAbs, filePath)
			if err != nil {
				return err
			}
			if fileInfo.err == "" {
				fileInfo.fixRequired = fileInfo.mime != "" && fileInfo.oExt != fileInfo.realExt
				if needFix {
					err := fixFileExt(&fileInfo)
					if err != nil {
						return err
					}
					if fileInfo.fixed {
						fixedFileCnt++
					}
				}
			}
			if !silent {
				mutex.Lock()
				defer mutex.Unlock()
				tableWriter.AddRow(&fileInfo)
			}
			procFileCnt++
			return nil
		})
		return nil
	})
	if err != nil {
		return err
	}

	err = errGrp.Wait()
	if err != nil {
		return err
	}

	if !silent {
		err = tableWriter.Finish()
		if err != nil {
			return err
		}
	}

	duration := time.Since(start)
	fmt.Printf("\n%d file(s) processed and %d file(s) fixed in %v\n", procFileCnt, fixedFileCnt, duration)
	return nil
}

func getFileInfo(dirPath string, filePath string) (_FileInfo, error) {
	result := _FileInfo{
		filePath:    filePath,
		filePathCut: strings.Replace(filePath, dirPath, "<dir>", 1),
		mime:        "unknown",
		oExt:        filepath.Ext(filePath),
	}

	file, err := os.Open(filePath)
	if err != nil {
		return _FileInfo{}, err
	}
	defer file.Close()

	head := make([]byte, 512)
	_, err = file.Read(head)
	if err != nil {
		if errors.Is(err, io.EOF) {
			result.err = "File is empty"
			return result, nil
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
		result.mime = fileType.MIME.Value
		result.realExt = "." + fileType.Extension
	}

	return result, nil
}

func fixFileExt(fileInfo *_FileInfo) error {
	if fileInfo.fixRequired {
		oldPath := fileInfo.filePath
		newPath := strings.TrimSuffix(fileInfo.filePath, fileInfo.oExt) + fileInfo.realExt
		fileStat, err := os.Stat(newPath)
		if err != nil && !os.IsNotExist(err) {
			return err
		}
		if fileStat != nil {
			fileInfo.err = "File with the same name is already exists"
			return nil
		}
		if err = os.Rename(oldPath, newPath); err != nil {
			return err
		}
		fileInfo.fixed = true
	}
	return nil
}
