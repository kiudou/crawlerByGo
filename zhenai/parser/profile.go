package parser

import (
	"regexp"
	"learngo/crawler/engine"
	"learngo/crawler/model"
	"strconv"
)

var re = regexp.MustCompile(`<div class="m-btn purple" data-v-bff6f798>([^<]+)</div><div class="m-btn purple" data-v-bff6f798>([^<]+)</div><div class="m-btn purple" data-v-bff6f798>([^<]+)</div><div class="m-btn purple" data-v-bff6f798>([^<]+)</div><div class="m-btn purple" data-v-bff6f798>([^<]+)</div><div class="m-btn purple" data-v-bff6f798>([^<]+)</div><div class="m-btn purple" data-v-bff6f798>([^<]+)</div>`)
var idUrlRe = regexp.MustCompile(`http://album.zhenai.com/u/([\d]+)`)

func ParseProfile(contents []byte, name string, url string) engine.ParseResult {
	profile := model.Profile{}
	matches := re.FindSubmatch(contents)
	//log.Printf("matches :%s\n",matches)
	if matches != nil {
		profile.Name = string(name)
		profile.Marriage = string(matches[1])
		profile.Age, _ = strconv.Atoi(string(matches[2]))
		profile.Xinzuo = string(matches[3])
		profile.Height, _ = strconv.Atoi(string(matches[4]))
		profile.Weight, _ = strconv.Atoi(string(matches[5]))
		profile.WorkingGround = string(matches[6])
		profile.Income = string(matches[7])
	}


	result := engine.ParseResult{
		Items: []engine.Item{
			{
				Url: url,
				Type: "zhenai",
				Id: extractString([]byte(url), idUrlRe),
				Payload: profile,
			},

		},
	}
	return result
}
func extractString(contents []byte, re *regexp.Regexp) string {
	match := re.FindSubmatch(contents)

	if len(match) >= 2 {
		return string(match[1])
	} else {
		return ""
	}
}
