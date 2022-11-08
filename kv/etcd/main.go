package main

import (
	"context"
	"encoding/json"
	"etcd/cluster"
	"fmt"
	clientv3 "go.etcd.io/etcd/client/v3"
	"log"
)

func main() {

	//addr := []string{"http://etcd-0.etcd.com:2379", "http://etcd-1.etcd.com:2379", "http://etcd-2.etcd.com:2379"}
	addr := []string{"192.168.204.132:32379"}
	client, err := cluster.InitClusterClient(addr)
	if err != nil {
		log.Fatal(err)
	}
	defer func(client *clientv3.Client) {
		_ = client.Close()
	}(client)

	list, err := client.MemberList(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < len(list.Members); i++ {
		member, _ := json.Marshal(list.Members[i])
		fmt.Printf("member-%d: %s\n", i, string(member))
	}

	//put, err := client.Put(context.Background(), "foo", "bar")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//fmt.Printf("%v\n", put)

	//get, err := client.Get(context.Background(), "foo")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Printf("%v\n", get)
}
