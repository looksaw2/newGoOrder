package main

import (
	"log"
	"github.com/looksaw/go-orderv2/common/config"
	"github.com/spf13/viper"
)

//初始化init
//读取配置
func init() {
	if err := config.NewViperConfig(); err != nil {
		log.Fatal(err)
	}
}

//
func main(){
	log.Printf("%+v",viper.Get("order"))
}