package util

import "fmt"

func Info(msg string) {
	fmt.Println("\033[36m[INFO]\033[0m", msg)
}

func Success(msg string) {
	fmt.Println("\033[32m[SUCCESS]\033[0m", msg)
}

func Error(msg string) {
	fmt.Println("\033[31m[ERROR]\033[0m", msg)
}
