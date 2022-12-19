package helpers

import (
	"bufio"
	"fmt"
	"os"
)

func Println(items ...interface{}) {
	printStr := ""
	for _, i := range items {
		printStr += fmt.Sprintf("%+v", i)
	}

	fmt.Println(printStr)
}

func WaitForInput() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
}
