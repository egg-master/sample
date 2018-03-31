package main

import (
	"fmt"
	"os"
)

func printEnv() {
	fmt.Println("========================================================")
	envList := os.Environ()
	for _, env := range envList {
		fmt.Println(env)
	}
	fmt.Println("========================================================")
}

func main() {
	dummyString := "TEST__________"

	os.Setenv(dummyString, "env test")
	fmt.Println(os.Getenv(dummyString))
	os.Unsetenv(dummyString)
	printEnv()
}
