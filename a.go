package main

import (
	"context"
	"github.com/cloudflare/cloudflare-go"
	"log"
	"os"
)

func main() {
	api, err := cloudflare.New(os.Getenv("CLOUDFLARE_API_KEY"), os.Getenv("CLOUDFLARE_API_EMAIL"))
	if err != nil {
		log.Fatal(err)
	}

	zoneID, err := api.ZoneIDByName("newzhxu.com")
	if err != nil {
		log.Fatal(err)
	}

	// Fetch all records for a zone
	recs, _, err := api.ListDNSRecords(context.Background(), cloudflare.ZoneIdentifier(zoneID), cloudflare.ListDNSRecordsParams{})
	if err != nil {
		log.Fatal(err)
	}
	var dnsRecord cloudflare.DNSRecord

	s := "hhhh.newzhxu.com"
	for _, r := range recs {
		if r.Name == s {
			dnsRecord = r
			log.Println(r.ID)
		}
	}

	log.Println(dnsRecord.Content, "33333333333")
	if dnsRecord.Content == "" {
		a, err := api.CreateDNSRecord(context.Background(), cloudflare.ZoneIdentifier(zoneID), cloudflare.CreateDNSRecordParams{
			Type:    "A",
			Name:    s,
			Content: "1.1.1.1",
		})
		if err != nil {
			return
		}
		log.Println(a, "eqweqweqweqwe")
	} else {
		record, err := api.UpdateDNSRecord(context.Background(), cloudflare.ZoneIdentifier(zoneID), cloudflare.UpdateDNSRecordParams{
			Type:    "A",
			Name:    s,
			Content: "2.2.2.2",
			ID:      dnsRecord.ID,
			Proxied: nil,
		})
		if err != nil {
			log.Fatal(err)
			return
		}
		log.Println(record.Content, "fffffffffffffffffff")
	}

	//var	record  cloudflare.DNSRecord
	//for _, r := range recs {
	//	if r.Name == "test.newzhxu.com" {
	//		record = r
	//		//record =r
	//	}
	//	fmt.Printf("%s: %s\n", r.Name, r.Content)
	//}
	//api.UpdateDNSRecord(context.Background(), cloudflare.ZoneIdentifier(zoneID), cloudflare.UpdateDNSRecordParams{
	//	Type:     "",
	//	Name:     "",
	//	Content:  "",
	//	Data:     nil,
	//	ID:       "",
	//	Priority: nil,
	//	TTL:      0,
	//	Proxied:  nil,
	//	Comment:  "",
	//	Tags:     nil,
	//})
	api.CreateDNSRecord(context.Background(), cloudflare.ZoneIdentifier(zoneID), cloudflare.CreateDNSRecordParams{})

}
