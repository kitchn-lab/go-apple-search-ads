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
	campaignID := int64(284979913)
	adGroupID := int64(285227483)
	targetingKeywordID := int64(306701962)
	/*
		campaignID := int64(262931515)
		adGroupID := int64(262987552)
		targetingKeywordID := int64(263022253)
	*/
	pemdat, _ := ioutil.ReadFile("../cert.pem")
	keydat, _ := ioutil.ReadFile("../cert.key")
	client, err := searchads.NewClient(nil, pemdat, keydat, nil)
	if err != nil {
		log.Fatalf("Client error: %s", err)
		panic(err)
	}
	input := []*searchads.TargetingKeyword{
		&searchads.TargetingKeyword{
			ID:        targetingKeywordID,
			AdGroupID: adGroupID,
			Status:    searchads.KEYWORD_PAUSED,
			BidAmount: searchads.Amount{
				Amount:   fmt.Sprintf("%f", 3.0),
				Currency: "EUR",
			},
		},
	}
	updatedKeywords, rs, err := client.AdGroupTargetingKeyword.UpdateBulk(context.Background(), campaignID, adGroupID, input)
	if err != nil {
		log.Fatalf("TargetingKeyword error: %s", err)
		panic(err)
	}
	res, _ := json.Marshal(&updatedKeywords)
	fmt.Println(string(res))
	fmt.Println("----------------")
	fmt.Println(rs.Pagination.ItemsPerPage)
	fmt.Println(rs.Pagination.StartIndex)
	fmt.Println(rs.Pagination.TotalResults)
}
