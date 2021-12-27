package main

import (
	"fmt"
	"github.com/kevin/boibot/bot"
	"github.com/kevin/boibot/config"
)

func main() {
	err := config.ReadConfig();

	if err != nil {
		fmt.Println(err.Error());
		return
	}

	bot.Start()
	
	<-make(chan struct{})
	return
}
