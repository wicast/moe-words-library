package filter

import (
	"regexp"
)

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
