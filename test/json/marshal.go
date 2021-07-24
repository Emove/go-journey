package json

import (
	"encoding/json"
	"fmt"
)

type PSBStatus struct {
	NowX     int    `json:"nowX"`
	NowY     int    `json:"nowY"`
	UniqueID uint64 `json:"uniqueID"`
	Seq      uint32 `json:"seq"`
	MapId    int    `json:"mapID"`
	Status   int    `json:"status"`
	HaveBox  int    `json:"haveBox"`
	Version  int    `json:"version"`
}

func MarshalTest() {
	status := PSBStatus{
		NowX:     0,
		NowY:     0,
		UniqueID: 355762198263610168,
		Seq:      4,
		MapId:    1,
		Status:   0,
		HaveBox:  0,
		Version:  400,
	}
	marshal, err := json.Marshal(status)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(string(marshal))
}

type Container struct {
	Code string `json:"code"`
}

func UnmarshalTest() {
	val := "{\"code\": \"skdlaskdl\"}"
	var con Container
	err := json.Unmarshal([]byte(val), &con)
	if err != nil {
		fmt.Println(err.Error())
	}
}
