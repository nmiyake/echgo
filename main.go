package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/nmiyake/echgo/echo"
)

func main() {
	fmt.Println(echo.Echo(strings.Join(os.Args[1:], " ")))
}
