package main

import (
	//	"fmt"
	"github.com/wicast/moe-words-library/src/crawler"
	"github.com/wicast/moe-words-library/src/filter"
	"github.com/wicast/moe-words-library/src/go-pinyin"
	"os"
)

func main() {
	result := crawler.QueryAll()

	pa := pinyin.NewArgs()
	pa.Heteronym = true
	pa.Separator = "'"

	dict := make(map[string]string)
	for _, R := range result {
		for _, v := range R.Pages {
			for _, data := range filter.Split(v.Title) {
				if len(data) > 3 {
					//fmt.Println(data)
					dict[data] = pinyin.Slug(data, pa)
				}
			}
		}
	}

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
}
