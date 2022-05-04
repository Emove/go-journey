package main

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"log"
	"os"
	"yaml/conf"
)

func main() {
	root, err := os.Getwd()
	if err != nil {
		log.Fatal(err.Error())
	}
	filePath := fmt.Sprintf("%s%s", root, "/conf/service.yaml")
	fd, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer func(fd *os.File) {
		_ = fd.Close()
	}(fd)

	service := &conf.ServiceConfig{}
	err = yaml.NewDecoder(fd).Decode(service)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Printf("%+v", service)
}
