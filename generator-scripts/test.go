package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"net/http"
	"os"

	openapi "local/dct-sdk-go"
)

func main() {

	var KEY string
	var HOST string

	fmt.Println("Enter Your DCT KEY (excluding 'apk'): ")
	fmt.Scanln(&KEY)

	fmt.Println("Enter Your DCT HOST ('hostname.domain'): ")
	fmt.Scanln(&HOST)

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
	req := client.ManagementAPI.GetRegisteredEngines(ctx)
	engines, _, err := client.ManagementAPI.GetRegisteredEnginesExecute(req)
	if err != nil {
		fmt.Print("Error while retrieving Delphix Engines.")
		fmt.Print(err)
		os.Exit(1)
	}
	fmt.Printf("Retrieved list of Delphix Engines: \n")
	for _, engine := range engines.GetItems() {
		fmt.Printf("%s \n", *engine.Hostname)
	}

	reporting, _, err := client.ReportingAPI.GetProductInfo(ctx).Execute()
	if err != nil {
		fmt.Print("Error while retrieving Delphix Engines.")
		fmt.Print(err)
		os.Exit(1)
	}
	fmt.Printf("Product version: %s \n", *reporting.ApiVersion)

	req_vdb := client.VDBsAPI.GetVdbs(ctx)

	vdbs, _, err := client.VDBsAPI.GetVdbsExecute(req_vdb)

	if err != nil {
		fmt.Print("Error while retrieving Delphix Vdb. \n")
		fmt.Print(err)
		os.Exit(1)
	}

	fmt.Printf("Retrieved list of VDBs: \n")
	for _, vdbs := range vdbs.GetItems() {
		fmt.Printf("%s \n", vdbs.GetName())
	}
}
