package main

import (
	"fmt"
	"time"
)

func main() {

	aa := "2022-01-06T05:39:27.000Z"

	bb := time.Now().Format("2006-01-02T15:04:05.000Z")

	if aa < bb {
		fmt.Println("111")
	}

}
