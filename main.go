package main

import (
	"fmt"
	fmt2 "github.com/ArtisanCloud/go-libs/fmt"
	"github.com/ArtisanCloud/go-wechat/src/work"
)

func main() {

	fmt.Printf("hello Wechat!")

	config := map[string]string{
		"corp_id" : "ww454dfb9d6f6d432a",
		//'agent_id' : 100020,
		"secret"   : "",
	}

	app := work.NewWork(config)
	fmt2.Dump(app)



}