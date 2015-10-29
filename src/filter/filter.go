package filter

import (
	// "fmt"
	"github.com/wicast/moe-words-library/src/crawler"
	"github.com/wicast/moe-words-library/src/go-pinyin"
	"regexp"
)

func ResultSet2Dict(result []crawler.ResultSet) map[string]string {
	dict := make(map[string]string)
	pa := pinyin.NewArgs()
	pa.Heteronym = true
	pa.Separator = "'"
	for _, R := range result {
		for _, v := range R.Pages {
			// fmt.Println(v.Title)
			for _, data := range Split(v.Title) {
				if len(data) > 3 {
					//fmt.Println(data)
					dict[data] = pinyin.Slug(data, pa)
				}
			}
		}
	}

	return dict
}

func Split(origin_s string) []string {
	hanzi := regexp.MustCompile(`[\p{Han}]+`)
	jpn := regexp.MustCompile(`[\p{Hiragana}]+|[\p{Katakana}]+`)

	if jpn.FindAllString(origin_s, -1) != nil {
		return nil
	}
	result := hanzi.FindAllString(origin_s, -1)
	return result
}

//需要人工判断的多音字
func Unsure(s string) bool {
	DuoYinzi := regexp.MustCompile(`长|朝|重|都|角|乐|传|藏|血`)
	if DuoYinzi.FindAllString(s, -1) != nil {
		return true
	} else {
		return false
	}
}
