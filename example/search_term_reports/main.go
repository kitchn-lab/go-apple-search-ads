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
	start := time.Now()
	cID := int64(262773151)
	// aID := int64(262825521)
	pemdat, _ := ioutil.ReadFile("../cert.pem")
	keydat, _ := ioutil.ReadFile("../cert.key")
	client, err := searchads.NewClient(nil, pemdat, keydat, nil)
	if err != nil {
		log.Fatalf("Client error: %s", err)
		panic(err)
	}
	filter := searchads.ReportFilter{
		StartTime:   "2019-04-01",
		EndTime:     "2019-04-29",
		Granularity: searchads.DAILY,
		TimeZone:    searchads.UTC,
		Selector: searchads.Selector{
			OrderBy: []searchads.OrderBySelector{
				searchads.OrderBySelector{
					Field:     searchads.OrderByImpressions,
					SortOrder: searchads.ASCENDING,
				},
			},
			Pagination: searchads.PaginationSelector{
				Offset: 0,
				Limit:  1000,
			},
		},
		GroupBy:                    []searchads.GroupBy{},
		ReturnRecordsWithNoMetrics: false,
		ReturnRowTotals:            false,
		ReturnGrandTotals:          false,
	}
	report, rs, err := client.Report.SearchTerms(context.Background(), cID, &filter)
	if err != nil {
		log.Fatalf("Campaign Reports error: %s", err)
		panic(err)
	}
	for _, data := range report.ReportingDataResponse.Row {
		res, _ := json.Marshal(&data)
		fmt.Println(string(res))
		fmt.Println("----------------")
	}
	fmt.Println(rs.Pagination.ItemsPerPage)
	fmt.Println(rs.Pagination.StartIndex)
	fmt.Println(rs.Pagination.TotalResults)
	duration := time.Since(start)
	fmt.Println("duration")
	fmt.Println(duration)
}
