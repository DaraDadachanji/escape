package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		line, readErr := reader.ReadString('\n')
		if len(line) > 0 {
			str := line[:len(line)-1]
			output := strconv.Quote(str)
			fmt.Print(output[1 : len(output)-1])
		}
		if readErr == io.EOF {
			return
		}
		fmt.Println()
	}
}
