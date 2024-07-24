package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func getCIDRsFromASN(asn string) []string {
	url := fmt.Sprintf("https://bgp.he.net/%s", asn)
	doc, err := goquery.NewDocument(url)
	if err != nil {
		log.Fatal(err)
	}

	var ret []string
	doc.Find("table[id*=table_prefixes] tr").Each(func(i int, s *goquery.Selection) {
		td := s.Find("td").First()
		if td != nil {
			ret = append(ret, strings.TrimSpace(td.Text()))
		}
	})

	return ret
}

func lookupASNFromOrg(org string) map[string]string {
	url := fmt.Sprintf("https://bgp.he.net/search?search%%5Bsearch%%5D=%s&commit=Search", org)
	doc, err := goquery.NewDocument(url)
	if err != nil {
		log.Fatal(err)
	}

	database := make(map[string]string)
	doc.Find("table tr").Each(func(i int, s *goquery.Selection) {
		td := s.Find("td")
		if td.Length() >= 3 && strings.TrimSpace(td.Eq(1).Text()) == "ASN" {
			asn := strings.TrimSpace(td.Eq(0).Text())
			description := strings.TrimSpace(td.Eq(2).Text())
			database[asn] = description
		}
	})

	return database
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run script.go <organization>")
		os.Exit(1)
	}

	org := os.Args[1]
	results := lookupASNFromOrg(org)
	for k, v := range results {
		if !strings.Contains(strings.ToLower(v), strings.ToLower(org)) {
			continue
		}
		fmt.Fprintf(os.Stderr, "------- %s - %s -------\n", k, v)
		cidrs := getCIDRsFromASN(k)
		for _, cidr := range cidrs {
			fmt.Println(cidr)
		}
	}
}
