package color

import (
	"strings"
	"time"
)

var enableTimePrefix = false
var prefixAddition = ""
func EnableTimePrefix(additionPrefix... string) {
	if len(additionPrefix) > 0 {
		prefixAddition=strings.Join(additionPrefix,"")
	}
	enableTimePrefix = true
}
func SuccessMessage(format string, a ...interface{}) {
	colorPrint(getTimePrefix()+format, FgHiGreen, a...)
}
func ErrorMessage(format string, a ...interface{}) {
	colorPrint(getTimePrefix()+format, FgHiRed, a...)
}
func WarningMessage(format string, a ...interface{}) {
	colorPrint(getTimePrefix()+format, FgHiYellow, a...)
}
func InfoMessage(format string, a ...interface{}) {
	colorPrint(getTimePrefix()+format, FgWhite, a...)
}
func InfoBlackMessage(format string, a ...interface{}) {
	colorPrint(getTimePrefix()+format, FgBlack, a...)
}
func getTimePrefix() string {
	if !enableTimePrefix {
		return ""
	}
	s:= time.Now().Format("[2006/01/02 15:04:05] ")
	return s+prefixAddition
}
