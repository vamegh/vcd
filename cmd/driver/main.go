package main

import (
	"encoding/base64"
	"flag"
	"github.com/houst0n/vcd"
	"os"
)

func main() {
	creds := os.Getenv("VCD_CREDS")
	url := os.Getenv("VCD_URL")

	flag_url := flag.String("url", "", "Please provide a valid URL")
	flag_creds := flag.String("cred", "", "Please provide valid Credentials")
	flag.Parse()

	if url == "" {
		if *flag_url == "" {
			fmt.Println("URL Environment not set please set using export VCD_URL or by passing via command line argument --url=",
				"\nUrl ENV Currently : ", url, "\nURL Command Line Currently : ", *flag_url)
			os.Exit(1)
		} else {
			url = *flag_url
		}
	}
	if creds == "" {
		if *flag_creds == "" {
			fmt.Println("Credentials Environment not set please set using export VCD_CREDS or by passing via command line argument --cred=",
				"\nVCD_CREDS Currently : ", creds, "\nCREDS Command Line Currently : ", *flag_creds)
			os.Exit(1)
		} else {
			creds = *flag_creds
		}
	}

	creds64 := base64.StdEncoding.EncodeToString([]byte(creds))
	v, _ := vcd.NewClient(url, creds64)
	v.ShowToken()
}
