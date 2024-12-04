package main

import (
	"fmt"

	"github.com/codedwells/avcs/cli"
)

func init(){
    cli.GenerateCommand()
    fmt.Println("CLI initialized")
}

func main() {
    cli.Execute()
}
