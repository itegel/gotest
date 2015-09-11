package main

import (
	"fmt"
	"github.com/itegel/gotest/beegotest/config"
)

func main() {
	beegotest.SetAppName("TestApp")
	var appname = beegotest.GetAppName()
	fmt.Println(appname)

}
