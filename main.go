package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func split(data, delimeter string) []string {
	result := make([]string, strings.Count(data, delimeter)+1)
	count := strings.Count(data, delimeter)
	for i := 0; i < count; i++ {
		result[i] = data[:strings.Index(data, delimeter)+len(delimeter)]
		data = data[strings.Index(data, delimeter)+1+len(delimeter):]
	}
	fmt.Println(count)
	result[count] = data
	return result
}
func new_split(data, delimeter string) []string {
	result := make([]string, strings.Count(data, delimeter)+1)
	return result
}
func main() {
	file, err := os.Open("hello.txt")
	if err != nil {
		panic(err)
	}
	b_text, err_f := io.ReadAll(file)
	if err_f != nil {
		panic(err_f)
	}
	text := string(b_text)
	fmt.Println(split(text, "package"))
}
