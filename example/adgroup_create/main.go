package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"time"

	"github.com/kitchn-lab/go-apple-search-ads/searchads"
)

func main() {
	orgID := int64(1405310)
	campaignID := int64(12312323)
	pemdat, _ := ioutil.ReadFile("../cert.pem")
	keydat, _ := ioutil.ReadFile("../cert.key")
	client, err := searchads.NewClient(nil, pemdat, keydat, &orgID)
	if err != nil {
		log.Fatalf("Client error: %s", err)
		panic(err)
	}
	now := time.Now().UTC()
	startTime := fmt.Sprintf("%4d-%02d-%02dT%02d:%02d:%02d.000",
		now.Year(), now.Month(), now.Day(),
		now.Hour(), now.Minute(), now.Second())
	data := searchads.AdGroup{
		CampaignID:             campaignID,
		Name:                   "US_BRAND_EXACT_2",
		StartTime:              startTime,
		AutomatedKeywordsOptIn: false,
		CpaGoal: &searchads.Amount{
			Amount:   "5",
			Currency: "USD",
		},
		DefaultCpcBid: &searchads.Amount{
			Amount:   "3",
			Currency: "USD",
		},
	}

	createdAdGroup, _, err := client.AdGroup.Create(context.Background(), campaignID, &data)
	if err != nil {
		log.Fatalf("Campaign Create error: %s", err)
		panic(err)
	}
	res, _ := json.Marshal(&createdAdGroup)
	fmt.Println(string(res))
}
