package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

//	func split_if(data string) []string {
//		count := strings.Count(data, " if ") + strings.Count(data, " if(")
//		result := make([]string, count+1)
//		for i := 0; i < count; i++ {
//			now_ind := min(strings.Index(data, " if "),
//				strings.Index(data, " if("))
//			result[i] = data[:now_ind]
//			data = data[now_ind+4:]
//		}
//		result[count] = data
//		return result
//	}
func double_split(data, first_delim, second_delim string) []string {
	count := strings.Count(data, first_delim) + strings.Count(data, second_delim)
	result := make([]string, count+1)
	now_len := 0
	for i := 0; i < count; i++ {
		now_ind := min(strings.Index(data, first_delim),
			strings.Index(data, second_delim))
		if strings.Index(data, first_delim) < strings.Index(data, second_delim) {
			now_len = len(first_delim)
		} else {
			now_len = len(second_delim)
		}
		result[i] = data[:now_ind]
		data = data[now_ind+now_len:]
	}
	result[count] = data
	return result
}
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
	result := 0
	max_deep := 0
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
		new_data := double_split(to_analyze[i], " if ", " if(")
		now_deep := 0
		max_local_deep := 0
		for g := 0; g < len(new_data); g++ {
			if strings.Contains(new_data[g], "else ") || strings.Contains(new_data[g], "else(") {
				now_deep--
			} else {
				now_deep++
			}
			if now_deep > max_local_deep {
				max_local_deep = now_deep
			}
		}
		if max_local_deep < now_deep {
			max_local_deep = now_deep
		}
		if max_local_deep > max_deep {
			result = i
			max_deep = max_local_deep
		}
	}
	fmt.Println(result, max_deep)
}
