package persist

import (
	"testing"
	"learngo/crawler/model"
	"github.com/olivere/elastic"
	"context"
	"encoding/json"
	"learngo/crawler/engine"
)

func TestSave(t *testing.T) {
	expected := engine.Item{
		Url: "http://album.zhenai.com/u/1884420087",
		Id: "1884420087",
		Type: "zhenai",
		Payload: model.Profile{
			Name: "心悦",
			WorkingGround: "工作地:阿坝金川",
			Age: 0,
			Height: 0,
			Weight: 0,
			Income: "月收入:3-5千",
			Marriage: "离异",
			Xinzuo: "魔羯座(12.22-01.19)",
		},
	}
	client, err := elastic.NewClient(elastic.SetSniff(false))

	if err != nil {
		panic(err)
	}
	//save
	const index = "dating_test"
	err = save(client, expected,index)
	if err != nil {
		panic(err)
	}



	//fetch save item
	resp, err := client.Get().Index("dating_profile").
		Type(expected.Type).Id(expected.Id).Do(context.Background())

	if err != nil {
		panic(err)
	}

	t.Logf("%s", []byte(*resp.Source))
	var actual engine.Item
	json.Unmarshal(*resp.Source, &actual)
	actualProfile, _ := model.FromJsonObj(actual.Payload)
	actual.Payload = actualProfile
	//verify
	if actual != expected {
		t.Errorf("Got %v; expected %v",actual,expected)
	}
}