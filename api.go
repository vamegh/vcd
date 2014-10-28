package vcd

import (
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
  "flag"
)

type Client struct {
	Token      string
	URL        string
	httpClient *http.Client
	creds      string
}

func NewClient() (*Client, error) {
	creds := os.Getenv("VCD_CREDS")
	url := os.Getenv("VCD_URL")

  flag_url := flag.String("url", "", "Please provide a valid URL")
  flag_creds := flag.String("cred", "", "Please provide valid Credentials")
  flag.Parse()

  if url == "" {
    if *flag_url == "" {
      fmt.Println("URL Environment not set please set using export VCD_URL or by passing via command line argument --url=",
                  "\nUrl ENV Currently : ",url, "\nURL Command Line Currently : ",*flag_url)
      os.Exit(1)
    } else {
      url = *flag_url
    }
  }
  if creds == "" {
    if *flag_creds == "" {
      fmt.Println("Credentials Environment not set please set using export VCD_CREDS or by passing via command line argument --cred=",
                  "\nVCD_CREDS Currently : ",creds,"\nCREDS Command Line Currently : ",*flag_creds )
      os.Exit(1)
    } else {
      creds = *flag_creds
    }
  }

	creds64 := base64.StdEncoding.EncodeToString([]byte(creds))

	client := Client{
		Token:      "FAIL",
		URL:        url,
		httpClient: http.DefaultClient,
		creds:      creds64,
	}

	client.login()
	return &client, nil
}

func (v *Client) login() error {
	endpoint, _ := url.Parse(v.URL)
	endpoint.Path = "/api/sessions"
	var body io.ReadWriter

	hReq, _ := http.NewRequest("POST", endpoint.String(), body)

	hReq.Header.Set("Accept", "application/*+xml;version=5.5")
	hReq.Header.Set("Authorization", fmt.Sprintf("Basic %s", v.creds))

	re, err := v.httpClient.Do(hReq)
	if err != nil {
		return err
	}
	defer re.Body.Close()

	v.Token = re.Header["X-Vcloud-Authorization"][0]
	return nil
}

func (v *Client) ShowToken() {
  if (v.Token == "FAIL") {
    fmt.Println("Auth Has Failed")
    os.Exit(1)
  } else {
    fmt.Println("VCD Token: ", v.Token)
  }
}

