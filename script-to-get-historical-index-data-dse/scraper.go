package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/antchfx/htmlquery"
	"github.com/araddon/dateparse"
)

type MarketSTAT struct {
	Date               string  `json:"date"`
	DSEXIndex          float64 `json:"dsex_index"`
	DSEXIndexChange    float64 `json:"dsex_index_change"`
	DS30Index          float64 `json:"ds30_index"`
	DS30IndexChange    float64 `json:"ds30_index_change"`
	TotalTrade         int64   `json:"total_trade"`
	TotalValaueTaka    float64 `json:"total_value_taka"`
	TotalVolume        int64   `json:"total_volume"`
	TotalMarketCapital float64 `json:"total_market_capital"`
}

func main() {
	fmt.Println()

	url := "https://www.dsebd.org/market_summary.php?startDate=2019-04-08&endDate=2021-04-08&archive=data"

	doc, err := htmlquery.LoadURL(url)
	if err != nil {
		log.Println(err)
	}

	var marketStats []MarketSTAT

	div, _ := htmlquery.QueryAll(doc, "/html/body/div[2]/section/div/div[3]/div[1]/div")
	for _, v := range div {
		table, _ := htmlquery.QueryAll(v, "//table")
		for indexTA := range table {
			tr, err := htmlquery.QueryAll(table[indexTA], "//tr")
			if err != nil {
				log.Println(err)
			}

			var marketStat MarketSTAT
			for indexTR := range tr {
				td, _ := htmlquery.QueryAll(tr[indexTR], "//td")
				if indexTR == 0 {
					s := strings.Replace(htmlquery.InnerText(td[0]), "Market Summary of ", "", -1)
					t, _ := dateparse.ParseAny(s)
					//log.Println(t)
					taa := strings.Replace(t.String(), " 00:00:00 +0000 UTC", "", -1)
					//log.Println(taa)
					marketStat.Date = taa
				}

				if indexTR == 1 {
					if strings.Contains(htmlquery.InnerText(td[0]), "DSEX Index") {
						//log.Println(htmlquery.InnerText(tr[index]))
						f, _ := strconv.ParseFloat(strings.Replace(htmlquery.InnerText(td[1]), ",", "", -1), 64)
						//log.Println(f)
						marketStat.DSEXIndex = f
					}

					if strings.Contains(htmlquery.InnerText(td[2]), "Total Trade") {
						//log.Println(htmlquery.InnerText(tr[index]))
						f, _ := strconv.ParseInt(strings.Replace(htmlquery.InnerText(td[3]), ",", "", -1), 10, 64)
						//log.Println(f)
						marketStat.TotalTrade = f
					}

				}

				if indexTR == 2 {
					if strings.Contains(htmlquery.InnerText(td[0]), "DSEX Index Change") {
						//log.Println(htmlquery.InnerText(tr[index]))
						f, _ := strconv.ParseFloat(strings.Replace(htmlquery.InnerText(td[1]), ",", "", -1), 64)
						//log.Println(f)
						marketStat.DSEXIndexChange = f
					}

					if strings.Contains(htmlquery.InnerText(td[2]), "Total Value Taka(mn)") {
						//log.Println(htmlquery.InnerText(tr[index]))
						f, _ := strconv.ParseFloat(strings.Replace(htmlquery.InnerText(td[3]), ",", "", -1), 64)
						//log.Println(f)
						marketStat.TotalValaueTaka = f
					}

				}

				if indexTR == 3 {
					if strings.Contains(htmlquery.InnerText(td[0]), "DS30 Index") {
						//log.Println(htmlquery.InnerText(tr[index]))
						f, _ := strconv.ParseFloat(strings.Replace(htmlquery.InnerText(td[1]), ",", "", -1), 64)
						//log.Println(f)
						marketStat.DS30Index = f
					}

					if strings.Contains(htmlquery.InnerText(td[2]), "Total Volume") {
						//log.Println(htmlquery.InnerText(tr[index]))
						f, _ := strconv.ParseInt(strings.Replace(htmlquery.InnerText(td[3]), ",", "", -1), 10, 64)
						//log.Println(f)
						marketStat.TotalVolume = f
					}

				}

				if indexTR == 4 {
					if strings.Contains(htmlquery.InnerText(td[0]), "DS30 Index Change") {
						//log.Println(htmlquery.InnerText(tr[index]))
						f, _ := strconv.ParseFloat(strings.Replace(htmlquery.InnerText(td[1]), ",", "", -1), 64)
						//log.Println(f)
						marketStat.DS30IndexChange = f
					}

					if strings.Contains(htmlquery.InnerText(td[2]), " Total Market Cap. Taka(mn) ") {
						//log.Println(htmlquery.InnerText(tr[index]))
						f, _ := strconv.ParseFloat(strings.Replace(htmlquery.InnerText(td[3]), ",", "", -1), 64)
						//log.Println(f)
						marketStat.TotalMarketCapital = f
					}

				}

			}
			marketStats = append(marketStats, marketStat)
		}
	}

	// for index := range marketStats {
	// 	log.Println(marketStats[index].Date)
	// 	log.Println(marketStats[index].DSEXIndex)
	// 	log.Println(marketStats[index].DSEXIndexChange)
	// 	log.Println(marketStats[index].DS30Index)
	// 	log.Println(marketStats[index].DS30IndexChange)
	// 	log.Println(marketStats[index].TotalMarketCapital)
	// 	log.Println(marketStats[index].TotalTrade)
	// 	log.Println(marketStats[index].TotalValaueTaka)
	// 	log.Println(marketStats[index].TotalVolume)
	// 	log.Println("===========================")
	// }

	f, _ := os.Create("test.json")
	defer f.Close()
	for index := range marketStats {
		marketStat := marketStats[index]
		j, _ := json.Marshal(marketStat)
		_, _ = f.WriteString(string(j)+"\n")
	}

}
