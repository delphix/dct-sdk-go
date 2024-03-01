/*
Delphix DCT API

Testing SamlLoginApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package delphix_dct_api

import (
	"context"
	"testing"

	openapiclient "github.com/delphix/dct-sdk-go/v14"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_delphix_dct_api_SamlLoginApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test SamlLoginApiService CheckSaml", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		httpRes, err := apiClient.SamlLoginApi.CheckSaml(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
