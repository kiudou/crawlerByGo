package parser

import (
	"learngo/crawler/engine"
	"regexp"
)

const cityListRe = `<a href="(http://www.zhenai.com/zhenghun/[a-z0-9]+)"[^>]*>([^<]+)</a>`

//通过正则表达式，从网页中提取所需要的连接和城市名字
func ParseCityList(contents []byte) engine.ParseResult {
	re := regexp.MustCompile(cityListRe)
	matches := re.FindAllSubmatch(contents, -1) //-1表示匹配所有结果
	result := engine.ParseResult{}
	for _, m := range matches {
		//result.Items = append(result.Items,"City " + string(m[2]))
		result.Requests = append(result.Requests, engine.Request{
			Url: string(m[1]),
			ParserFunc: ParseCity,
		})
	}
	return result
}
