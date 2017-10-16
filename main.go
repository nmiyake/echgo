package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/nmiyake/echgo/echo"
)

var version = "none"

func main() {
	versionVar := flag.Bool("version", false, "print version")
	flag.Parse()
	if *versionVar {
		fmt.Println("echgo version:", version)
		return
	}
	echoer := echo.NewEchoer()
	fmt.Println(echoer.Echo(strings.Join(flag.Args(), " ")))
}
