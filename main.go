// ship-test-go is a throwaway command that exists to be released.
package main

import (
	"fmt"
	"os"
)

var version = "dev"

func main() {
	if len(os.Args) > 1 && os.Args[1] == "--version" {
		fmt.Println("ship-test-go", version)
		return
	}
	fmt.Println("hello from ship-test-go")
}
