package engine

type Request struct {
	Url        string
	ParserFunc func([]byte) ParseResult //对于相对应的Url，要执行的具体函数
}

type ParseResult struct {
	Requests []Request
	Items    []Item
}

type Item struct {
	Url     string
	Id      string
	Type    string
	Payload interface{}
}

func NilParser([]byte) ParseResult {
	return ParseResult{}
}
