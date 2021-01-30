package main

import (
	"encoding/json"
	"fmt"
	"github.com/kataras/iris/v12"
	"os"
)

func main() {
	app := iris.New()
	app.Configure(iris.WithConfiguration(iris.YAML("E:/go/irisDemo/iris.yml")))

	file, _ := os.Open("E:/go/irisDemo/config.json")
	defer file.Close()

	decoder := json.NewDecoder(file)
	conf := Coniguration{}
	err := decoder.Decode(&conf)
	if err != nil {
		fmt.Println("Error:",err)
	}
	fmt.Println(conf.Port)

	app.Run(iris.Addr(":8007"))

}
type Coniguration struct {
	AppName string `json:"appname"`
	Port int `json:"port"`
}
