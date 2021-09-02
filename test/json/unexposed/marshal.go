package unexposed

import (
	"encoding/json"
	"fmt"
	"log"
)

type node struct {
	NodeID string `json:"node_id"`
}

func Marshal() {
	n := node{
		NodeID: "123213123",
	}
	marshal, err := json.Marshal(n)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(marshal))
}
