package main

import (
	"context"
	"log"
	"os"

	"github.com/cloudflare/cloudflare-go"
)

func main() {
	// Construct a new API object using a global API key
	api, err := cloudflare.New(os.Getenv("CLOUDFLARE_API_KEY"), os.Getenv("CLOUDFLARE_API_EMAIL"))
	// alternatively, you can use a scoped API token
	// api, err := cloudflare.NewWithAPIToken(os.Getenv("CLOUDFLARE_API_TOKEN"))
	if err != nil {
		log.Fatal(err)
	}

	// Most API calls require a Context
	ctx := context.Background()

	// Fetch user details on the account
	u, err := api.UserDetails(ctx)
	if err != nil {
		log.Fatal(err)
	}
	// Print user details
	log.Println(u)
	records, r, err := api.ListDNSRecords(ctx, cloudflare.ZoneIdentifier(u.ID), cloudflare.ListDNSRecordsParams{})
	if err != nil {
		return
	}

	log.Println(records, r)
}
