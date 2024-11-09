package gofit

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"text/tabwriter"

	"github.com/h2non/filetype"
)

type _FileInfo struct {
	filePath    string
	filePathCut string
	mime        string
	origExt     string
	realExt     string
	fixed       bool
}

func Scan(dir string) error {
	fis, err := walkDir(dir)
	if err != nil {
		return err
	}

	w := tabwriter.NewWriter(os.Stdout, 1, 1, 2, ' ', 0)
	fmt.Fprintln(w, "| File\t| Mime\t| Orig ext.\t| Real ext.\t| ")
	fmt.Fprintln(w, "| ----\t| ----\t| ---------\t| ---------\t| ")
	for _, fi := range fis {
		fmt.Fprintf(w, "| %s\t| %s\t| %s\t| %s\t|\n", fi.filePathCut, fi.mime, fi.origExt, fi.realExt)
	}
	w.Flush()

	return nil
}

func Fix(dirPath string) error {
	fis, err := walkDir(dirPath)
	if err != nil {
		return err
	}

	w := tabwriter.NewWriter(os.Stdout, 1, 1, 2, ' ', 0)
	fmt.Fprintln(w, "| File\t| Mime\t| Orig ext.\t| Real ext.\t| Fixed\t| ")
	fmt.Fprintln(w, "| ----\t| ----\t| ---------\t| ---------\t| -----\t| ")
	for _, fi := range fis {
		err = fixFileExt(&fi)
		if err != nil {
			return err
		}
		fmt.Fprintf(w, "| %s\t| %s\t| %s\t| %s\t| %s\t|\n", fi.filePathCut, fi.mime, fi.origExt, fi.realExt, If(fi.fixed, "yes", ""))
	}
	w.Flush()

	return nil
}

func walkDir(dirPath string) ([]_FileInfo, error) {
	var result []_FileInfo
	err := filepath.WalkDir(dirPath, func(f string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() {
			fi, err := getFileInfo(dirPath, f)
			if err != nil {
				return err
			}
			result = append(result, fi)
		}
		return nil
	})

	return result, err
}

func getFileInfo(dirPath string, filePath string) (_FileInfo, error) {
	var result _FileInfo

	file, err := os.Open(filePath)
	if err != nil {
		return result, err
	}
	defer file.Close()

	head := make([]byte, 1024)
	_, err = file.Read(head)
	if err != nil {
		return result, err
	}

	ft, err := filetype.Get(head)
	if err != nil {
		return result, err
	}

	result.filePath = filePath
	result.filePathCut = strings.Replace(filePath, dirPath, "<dir>/", 1)
	result.mime = ft.MIME.Value
	result.origExt = filepath.Ext(filePath)
	result.realExt = "." + ft.Extension

	return result, nil
}

func fixFileExt(fi *_FileInfo) error {
	if fi.mime != "" && fi.origExt != fi.realExt {
		oldPath := fi.filePath
		newPath := strings.TrimSuffix(fi.filePath, fi.origExt) + fi.realExt
		if err := os.Rename(oldPath, newPath); err != nil {
			return err
		}
		fi.fixed = true
	}

	return nil
}
