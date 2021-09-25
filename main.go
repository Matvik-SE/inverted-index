package main

import "fmt"

var dataDir = "./data/"

func main() {
	res := getDirFiles(dataDir)

	fmt.Printf("%v", res)
}
