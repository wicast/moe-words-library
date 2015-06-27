package main

import (
	"./src/crawler"
	"./src/filter"
	"./src/go-pinyin"
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
					dict[data] = pinyin.Slug(data, pa)
				}
			}
		}
	}

	Moe_dict_basic, err := os.Create("./dicts/Moe_dict_basic.org")
	defer Moe_dict_basic.Close()
	if err != nil {
		panic(err)
	}
	Moe_dict_unsure, err := os.Create("./dicts/Moe_dict_unsure.org")
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
