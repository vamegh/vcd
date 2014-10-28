package main

import (
	"encoding/base64"
	"github.com/houst0n/vcd"
	"os"
)

func main() {
	creds := os.Getenv("VCD_CREDS")
	url := os.Getenv("VCD_URL")
	creds64 := base64.StdEncoding.EncodeToString([]byte(creds))

	v, _ := vcd.NewClient(url, creds64)
	v.ShowToken()
}
