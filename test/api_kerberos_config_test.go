/*
Delphix DCT API

Testing KerberosConfigApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package delphix_dct_api

import (
	"context"
	"testing"

	openapiclient "github.com/delphix/dct-sdk-go/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_delphix_dct_api_KerberosConfigApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test KerberosConfigApiService GetKerberosConfigById", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var kerberosConfigId string

		resp, httpRes, err := apiClient.KerberosConfigApi.GetKerberosConfigById(context.Background(), kerberosConfigId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test KerberosConfigApiService ListKerberosConfigs", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.KerberosConfigApi.ListKerberosConfigs(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test KerberosConfigApiService SearchKerberosConfigs", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.KerberosConfigApi.SearchKerberosConfigs(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}