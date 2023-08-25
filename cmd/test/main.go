package main

import (
	"github.com/qwenode/color"
	"log"
	"strings"
)

func main() {
	color.EnableTimePrefix()
	color.SetOutputFile("a.log", log.Lshortfile)
	color.SuccessMessage("哟西,%s", "bbb")
	aa := "wowifw%s"
	color.SuccessMessage(strings.ReplaceAll(aa, "%s", "[%s]"), "xx")
}
