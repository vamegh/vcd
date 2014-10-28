package vcd

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
  "os"
)

type Client struct {
	Token      string
	URL        string
	httpClient *http.Client
	creds      string
}

func NewClient(url string, creds string) (*Client, error) {

	client := Client{
		Token:      "FAIL",
		URL:        url,
		httpClient: http.DefaultClient,
		creds:      creds,
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

	//fmt.Println("Body: ", re.Body)

	v.Token = re.Header["X-Vcloud-Authorization"][0]
	return nil
}

func (v *Client) ShowToken() {
	if v.Token == "FAIL" {
		fmt.Println("Auth Has Failed")
		os.Exit(1)
	} else {
		fmt.Println("VCD Token: ", v.Token)
	}
}
