package main

import "fmt"
import "os"

import "strings"
import "time"

type fn func()

func echo1() {
	fmt.Println("name: echo1")
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}

func echo2() {
	fmt.Println("name: echo2")
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)

}

func echo3() {
	fmt.Println("name: echo3")
	fmt.Println(strings.Join(os.Args[1:], " "))
}

func bench(f fn) {
	start := time.Now()
	f()
	fmt.Printf("%.10fs elapsed\n", time.Since(start).Seconds())
}

func main() {
	flist := []func(){echo1, echo2, echo3}

	for _, f := range flist {
		bench(f)
	}
}
