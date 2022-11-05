package main

import (
	"fmt"
	"log"
)

func init() {
	log.SetPrefix("TRACE: ")
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
}

func main() {
	fmt.Println("Hello World...")
	log.Println("message")
	//log.Fatalln("fatal message")
	log.Panicln("panic message")
}
