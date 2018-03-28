package main

import "fmt"
import "os"

func main() {
	for i, ele := range os.Args[1:] {
		fmt.Println(i, ele)
	}
}
