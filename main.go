package main

import (
	//	"fmt"
	"encoding/json"
	"github.com/wicast/moe-words-library/src/crawler"
	"github.com/wicast/moe-words-library/src/filter"
	"os"
)

func main() {
	result := crawler.QueryAll()

	dict := filter.ResultSet2Dict(result)

	var path string
	if os.IsPathSeparator('\\') {
		path = "\\"
	} else {
		path = "/"
	}

	dir, _ := os.Getwd()
	if _, err := os.Stat(dir + path + "dicts"); os.IsNotExist(err) {
		err_dir := os.Mkdir(dir+path+"dicts", os.ModePerm)
		if err_dir != nil {
			panic(err_dir)
		}
	}

	Moe_dict_basic, err := os.Create("./dicts/Moe_dict_basic.txt")
	defer Moe_dict_basic.Close()
	if err != nil {
		panic(err)
	}
	Moe_dict_unsure, err := os.Create("./dicts/Moe_dict_unsure.txt")
	defer Moe_dict_unsure.Close()
	if err != nil {
		panic(err)
	}

	for k, v := range dict {
		if filter.Unsure(k) == true {
			Moe_dict_unsure.WriteString(k + " " + v + "\n")
		} else {
			Moe_dict_basic.WriteString(k + " " + v + "\n")
		}
	}
	json_file, err := os.Create("./dicts/test.json")
	defer json_file.Close()
	json_byte, err := json.Marshal(dict)
	if err != nil {
		panic(err)
	}
	json_file.Write(json_byte)
}
