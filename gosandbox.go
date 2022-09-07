package main

import "fmt"
import "github.com/Mstuart712/gosandbox/workers"

func main() {
	fmt.Println("Hey world!")
	workers.StartWorkerGroup()
}
