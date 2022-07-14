package main

import "fmt"
import "flag"

func main() {
	var configFilePath string
	flag.StringVar(&configFilePath, "c", "", "config file path")
	flag.Parse()
	fmt.Println(configFilePath)
}
