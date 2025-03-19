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
		result[i] = data[:strings.Index(data, delimeter)]
		data = data[strings.Index(data, delimeter)+1+len(delimeter):]
	}
	result[count] = data
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
	to_analyze := split(text, "package")[1:]
	// fmt.Println(to_analise, len(to_analise))
	// fmt.Println(strings.Count(text, "package") + 1)
	for i := 0; i < len(to_analyze); i++ {
		fmt.Println(to_analyze[i])
	}
}
