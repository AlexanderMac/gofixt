package internal

import (
	"fmt"
	"slices"
	"strings"
	"time"
)

type Logger struct {
	printMode      PrintMode
	start          time.Time
	totalCnt       int
	ignoredCnt     int
	fixRequiredCnt int
	fixedCnt       int
	errorCnt       int
}

func NewLogger(printMode PrintMode) *Logger {
	return &Logger{printMode, time.Now(), 0, 0, 0, 0, 0}
}

func (log *Logger) PrintIntro(dirPath string, exts string, workMode WorkMode) {
	if !log.shouldPrint(PM_ALL, PM_IMPORTANT, PM_REPORT) {
		return
	}
	msg := fmt.Sprintf("Running gofixt in %s mode on %s...\n", workMode, dirPath)
	msg += fmt.Sprintf("- extension white list: %s\n", exts)
	msg += fmt.Sprintf("- print mode: %s\n", log.printMode)
	fmt.Printf("%s\n", msg)
}

func (log *Logger) PrintTableHeader() {
	if !log.shouldPrint(PM_ALL, PM_IMPORTANT) {
		return
	}

	fmt.Printf("| File%50s | Mime%12s | Orig ext. | Real ext. | Notes%10s |\n", "", "", "")
	fmt.Printf("| %s |\n", strings.Repeat("-", 115))
}

func (log *Logger) PrintTableRow(fi *_FileInfo) {
	log.totalCnt++
	var notes string
	switch fi.status {
	case FS_FIX_REQUIRED:
		log.fixRequiredCnt++
		notes = "Fix required"
	case FS_IGNORED:
		log.ignoredCnt++
		notes = "Ignored"
	case FS_FIXED:
		log.fixedCnt++
		notes = "Fixed"
	case FS_ERROR:
		log.errorCnt++
		notes = fi.notes
	}

	shouldPrint := false
	switch fi.status {
	case FS_NONE:
		shouldPrint = log.shouldPrint(PM_ALL)
	case FS_FIX_REQUIRED, FS_IGNORED, FS_FIXED, FS_ERROR:
		shouldPrint = log.shouldPrint(PM_ALL, PM_IMPORTANT)
	}
	if !shouldPrint {
		return
	}

	fmt.Printf(
		"| %-54s | %-16s | %-9s | %-9s | %-15s |\n",
		trimString(fi.fileTrimPath, 54, true),
		trimString(fi.mime, 16, true),
		trimString(fi.origExt, 9, false),
		trimString(fi.realExt, 9, false),
		trimString(notes, 15, false),
	)
}

func (log *Logger) PrintResult() {
	if !log.shouldPrint(PM_ALL, PM_IMPORTANT, PM_REPORT) {
		return
	}
	if log.shouldPrint(PM_ALL, PM_IMPORTANT) {
		fmt.Println("")
	}

	duration := time.Since(log.start)
	fmt.Printf(
		`Process has been completed in %v.
- %d file(s) processed
- %d file(s) ignored
- %d file(s) require fix
- %d file(s) fixed
- %d error(s)
`,
		duration,
		log.totalCnt,
		log.ignoredCnt,
		log.fixRequiredCnt,
		log.fixedCnt,
		log.errorCnt,
	)
}

func (log *Logger) shouldPrint(allowedModes ...PrintMode) bool {
	return slices.Contains(allowedModes, log.printMode)
}

func trimString(str string, n int, trimLet bool) string {
	if n < 0 || n >= len(str) {
		return str
	}

	n = n - 3
	if trimLet {
		return "..." + str[len(str)-n:]
	} else {
		return str[:n] + "..."
	}
}
