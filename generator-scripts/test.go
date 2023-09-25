package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"net/http"
	"os"

	openapi "github.com/delphix/dct-sdk-go/v2"
)

func main() {

	var KEY string
	var HOST string

	fmt.Println("Enter Your KEY (excluding 'apk'): ")
	fmt.Scanln(&KEY)

	fmt.Println("Enter Your HOST ('hostname.domain'): ")
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
	req := client.ManagementApi.GetRegisteredEngines(ctx)
	engines, _, err := client.ManagementApi.GetRegisteredEnginesExecute(req)
	if err != nil {
		fmt.Print("Error while retrieving Delphix Engines.")
		fmt.Print(err)
		os.Exit(1)
	}
	fmt.Printf("Retrieved list of Delphix Engines: \n")
	for _, engine := range engines.GetItems() {
		fmt.Printf("%s \n", engine.Hostname)
	}

	req_vdb := client.VDBsApi.GetVdbs(ctx)

	vdbs, _, err := client.VDBsApi.GetVdbsExecute(req_vdb)

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
