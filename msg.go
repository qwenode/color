package color

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

var enableTimePrefix = false
var prefixAddition = ""
var enableWriteToFile = false
var fileHandle *os.File

func SetOutputFile(logFile string, flag int) error {
	var err error
	fileHandle, err = os.OpenFile(logFile, os.O_CREATE|os.O_APPEND, os.ModePerm)
	if err != nil {
		return err
	}
	enableWriteToFile = true
	log.SetOutput(fileHandle)
	log.SetFlags(flag)
	return nil
}
func EnableTimePrefix(additionPrefix ...string) {
	if len(additionPrefix) > 0 {
		prefixAddition = strings.Join(additionPrefix, "")
	}
	enableTimePrefix = true
}
func SuccessMessage(format string, a ...interface{}) {
	sprintf := getTimePrefix() + fmt.Sprintf(format, a...)
	if enableWriteToFile {
		log.Output(2, sprintf)
	}
	colorPrint(sprintf, FgHiGreen)
}

func ErrorMessage(format string, a ...interface{}) {
	sprintf := getTimePrefix() + fmt.Sprintf(format, a...)
	if enableWriteToFile {
		log.Output(2, sprintf)
	}
	colorPrint(sprintf, FgHiRed)
}
func WarningMessage(format string, a ...interface{}) {
	sprintf := getTimePrefix() + fmt.Sprintf(format, a...)
	if enableWriteToFile {
		log.Output(2, sprintf)
	}
	colorPrint(sprintf, FgHiYellow)
}
func InfoMessage(format string, a ...interface{}) {
	sprintf := getTimePrefix() + fmt.Sprintf(format, a...)
	if enableWriteToFile {
		log.Output(2, sprintf)
	}
	colorPrint(sprintf, FgWhite)
}
func InfoBlackMessage(format string, a ...interface{}) {
	sprintf := getTimePrefix() + fmt.Sprintf(format, a...)
	if enableWriteToFile {
		log.Output(2, sprintf)
	}
	colorPrint(sprintf, FgBlack)
}
func getTimePrefix() string {
	if !enableTimePrefix {
		return ""
	}
	s := time.Now().Format("[2006/01/02 15:04:05] ")
	return s + prefixAddition
}
