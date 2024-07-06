package main

import (
	"bytes"
	"fmt"
	"os"
)

func Greet(writer *bytes.Buffer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}
func main() {
	Greet(os.Stdout, "Elodie")
	//such is needed when http response is coming from the server
	//this helps us to test std out
	//http code is explained in the greeting.go file and is commented
}
