package main

import "github.com/qwenode/color"

func main() {
	color.EnableTimePrefix()
	color.SuccessMessage("success")
	color.WarningMessage("warning")
	color.InfoMessage("info")
	color.InfoBlackMessage("info")
	color.ErrorMessage("error")
}
