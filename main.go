package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

var file_path string

func n_count(data string, delims []string) int {
	to_return := 0
	for i := 0; i < len(delims); i++ {
		to_return += strings.Count(data, delims[i])
	}
	return to_return
}

func ind_qualifier(data string, delims []string) [2]int {
	var to_return [2]int
	to_return[0] = 1e+9
	for i := 0; i < len(delims); i++ {
		if to_return[0] > index_preprocessor(strings.Index(data, delims[i])) {
			to_return[0] = index_preprocessor(strings.Index(data, delims[i]))
			to_return[1] = len(delims[i])
		}
	}
	return to_return
}

func n_split(data string, delims []string) []string {
	count := n_count(data, delims)
	result := make([]string, count+1)
	for i := 0; i < count; i++ {
		now_data := ind_qualifier(data, delims)
		result[i] = data[:now_data[0]]
		data = data[now_data[0]+now_data[1]+1:]
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
		new_data := n_split(to_analyze[i], []string{" if ", " if(", "\tif ", "\tif("})
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
