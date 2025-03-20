package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

var file_path string

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

func index_preprocessor(ind int) int {
	if ind == -1 {
		return 1e+9
	} else {
		return ind
	}
}

func forth_split(data, first_delim, second_delim, third_delim, forth_delim string) []string {
	count := strings.Count(data, first_delim) + strings.Count(data, second_delim) +
		strings.Count(data, third_delim) + strings.Count(data, forth_delim)
	result := make([]string, count+1)
	now_len := 0
	for i := 0; i < count; i++ {
		now_ind := min(index_preprocessor(strings.Index(data, first_delim)),
			index_preprocessor(strings.Index(data, second_delim)),
			index_preprocessor(strings.Index(data, third_delim)),
			index_preprocessor(strings.Index(data, forth_delim)))
		if now_ind == strings.Index(data, first_delim) {
			now_len = len(first_delim)
		} else if now_ind == strings.Index(data, second_delim) {
			now_len = len(second_delim)
		} else if now_ind == strings.Index(data, third_delim) {
			now_len = len(third_delim)
		} else {
			now_len = len(forth_delim)
		}
		// fmt.Println(data, strings.Contains(data, first_delim), strings.Contains(data, second_delim),
		// strings.Contains(data, third_delim), strings.Contains(data, forth_delim))
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
	fmt.Println("Enter file name with data type")
	fmt.Scan(&file_path)
	file, err := os.Open(file_path)
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
	// fmt.Println(to_analyze, len(to_analyze))
	// fmt.Println(strings.Count(text, "package") + 1)
	for i := 0; i < len(to_analyze); i++ {
		// new_data := double_split(to_analyze[i], " if ", " if(")
		new_data := forth_split(to_analyze[i], " if ", " if(", "\tif ", "\tif(")
		now_deep := 0
		max_local_deep := 0
		for g := 0; g < len(new_data); g++ {
			if strings.Contains(new_data[g], " else ") || strings.Contains(new_data[g], " else(") ||
				strings.Contains(new_data[g], "\telse(") || strings.Contains(new_data[g], "\telse ") {
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
