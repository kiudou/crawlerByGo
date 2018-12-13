package main

import (
	"learngo/crawler/engine"
	"learngo/crawler/zhenai/parser"
	"learngo/crawler/scheduler"
)

func main() {
	e := engine.ConcurrentEngine{
		Scheduler: &scheduler.SimpleScheduler{},
		WorkerCount: 100,
	}
	e.Run(engine.Request{
		Url: "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})
}


