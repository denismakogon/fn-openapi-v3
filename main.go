package main

import (
	"fmt"
	"github.com/denismakogon/fn-openapi/models"
	"io/ioutil"
	"os"
)

func main() {

	yamlFile, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	var fn models.Fn

	err = fn.Unmarshal(yamlFile, os.Stdout)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	var oai models.OpenAPISpec
	err = oai.FromFnSpec("http://localhost:8080", &fn)
	err = oai.Marshal(os.Stdout)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
