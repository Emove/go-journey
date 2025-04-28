package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
)

type Data struct {
}

func Write(num int) {
	m := make(map[int]*Data)
	for i := 0; i < num; i++ {
		m[i] = new(Data)
	}

}

func main() {
	go func() {
		log.Println(http.ListenAndServe(":8080", nil))
	}()
	Write(200000)
	select {}
}
