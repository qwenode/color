package color

import "time"

var enableTimePrefix = false

func EnableTimePrefix() {
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
	return time.Now().Format("[2006/01/02 15:04:05] ")
}
