package gofit

import (
	"fmt"
	"os"
	"text/tabwriter"
)

type TableWriter struct {
	writer *tabwriter.Writer
}

func NewTableWriter() *TableWriter {
	return &TableWriter{
		writer: tabwriter.NewWriter(os.Stdout, 1, 1, 1, ' ', 0),
	}
}

func (tw *TableWriter) AddHeader() {
	fmt.Fprintln(tw.writer, "| File\t| Mime\t| Orig ext.\t| Real ext.\t| Fixed\t| Error\t|")
	fmt.Fprintln(tw.writer, "| ----\t| ----\t| ---------\t| ---------\t| -----\t| -----\t|")
}

func (tw *TableWriter) AddRow(fi *_FileInfo) {
	fmt.Fprintf(tw.writer, "| %s\t| %s\t| %s\t| %s\t| %s\t| %s\t|\n", fi.filePathCut, fi.mime, fi.oExt, fi.realExt, If(fi.fixed, "Yes", ""), fi.err)
}

func (tw *TableWriter) Finish() error {
	return tw.writer.Flush()
}
