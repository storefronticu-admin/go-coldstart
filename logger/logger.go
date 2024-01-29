package logger

import (
	"fmt"
	"log"
	"os"

	"github.com/fatih/color"
)

// Info logs informational messages with formatted string.
func Info(format string, args ...interface{}) {
	infoLog := log.New(os.Stdout, color.GreenString("INFO: "), log.Ldate|log.Ltime)
	infoLog.Println(fmt.Sprintf(format, args...))
}

// Warning logs warning messages with formatted string.
func Warning(format string, args ...interface{}) {
	warningLog := log.New(os.Stdout, color.YellowString("WARNING: "), log.Ldate|log.Ltime|log.Lshortfile)
	warningLog.Println(fmt.Sprintf(format, args...))
}

// Error logs error messages with formatted string.
func Error(format string, args ...interface{}) {
	errorLog := log.New(os.Stderr, color.RedString("ERROR: "), log.Ldate|log.Ltime|log.Lshortfile)
	errorLog.Println(fmt.Sprintf(format, args...))
}

// Success logs success messages with formatted string.
func Success(format string, args ...interface{}) {
	successLog := log.New(os.Stdout, color.MagentaString("SUCCESS: "), log.Ldate|log.Ltime)
	successLog.Println(fmt.Sprintf(format, args...))
}

// Debug logs debug messages with formatted string.
func Debug(format string, args ...interface{}) {
	debugLog := log.New(os.Stdout, color.BlueString("DEBUG: "), log.Ldate|log.Ltime)
	debugLog.Println(fmt.Sprintf(format, args...))
}
