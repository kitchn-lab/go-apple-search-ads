package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/kitchn-lab/go-apple-search-ads/searchads"
)

func main() {
	pemdat, _ := ioutil.ReadFile("../api-keys.pem")
	keydat, _ := ioutil.ReadFile("../api-keys.key")
	client, err := searchads.NewClient(nil, pemdat, keydat, nil)
	if err != nil {
		log.Fatalf("Client error: %s", err)
		panic(err)
	}
	opt := searchads.ListOptions{Limit: 1000, Offset: 0}
	list, rs, err := client.Campaign.List(context.Background(), &opt)
	if err != nil {
		log.Fatalf("Campaign List error: %s", err)
		panic(err)
	}
	res, _ := json.Marshal(&list)
	fmt.Println(string(res))
	fmt.Println("----------------")
	fmt.Println(len(list))
	fmt.Println(rs.Pagination.ItemsPerPage)
	fmt.Println(rs.Pagination.StartIndex)
	fmt.Println(rs.Pagination.TotalResults)
}
