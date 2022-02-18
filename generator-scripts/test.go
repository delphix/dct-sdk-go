package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"net/http"
	"os"

	openapi "github.com/delphix/dct-sdk-go"
)

var KEY = "1.GCAE1EOOIa3JUQZcKE2umJa81xOo6MLL210m3NdbMfIEG8VTOz02Xt0wbXlsxS8J"
var HOST = "dct-apigw.dlpxdc.co"

func main() {
	apiKeyMap := make(map[string]openapi.APIKey)
	apiKeyMap["ApiKeyAuth"] = openapi.APIKey{
		Key:    KEY,
		Prefix: "apk",
	}

	ctx := context.WithValue(context.Background(), openapi.ContextAPIKeys, apiKeyMap)

	cfg := openapi.NewConfiguration()
	cfg.Host = HOST
	cfg.Scheme = "https"
	cfg.HTTPClient = &http.Client{Transport: &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}}

	client := openapi.NewAPIClient(cfg)

	req := client.EnginesApi.GetEngines(ctx)

	res, _, err := client.EnginesApi.GetEnginesExecute(req)

	if err != nil {
		fmt.Print("\n__________ERR__________\n")
		fmt.Print(err)
		os.Exit(1)
	}
	fmt.Printf("List of engines are: \n")

	for _, j := range res.Items {
		fmt.Printf("%s \n", *j.Hostname)
	}
}
