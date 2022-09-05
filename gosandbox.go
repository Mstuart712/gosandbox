package main

import (
	"fmt"
	"github.com/Mstuart712/gosandbox/workers"
)

func main() {
	fmt.Println("Hey world!")
	workers.StartWorker()
}
