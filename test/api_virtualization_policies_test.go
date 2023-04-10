/*
Delphix DCT API

Testing VirtualizationPoliciesApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package delphix_dct_api

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/delphix/dct-sdk-go"
)

func Test_delphix_dct_api_VirtualizationPoliciesApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test VirtualizationPoliciesApiService GetVirtualizationPolicyById", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var policyId string

		resp, httpRes, err := apiClient.VirtualizationPoliciesApi.GetVirtualizationPolicyById(context.Background(), policyId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VirtualizationPoliciesApiService ListVirtualizationPolicies", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.VirtualizationPoliciesApi.ListVirtualizationPolicies(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VirtualizationPoliciesApiService SearchVirtualizationPolicies", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.VirtualizationPoliciesApi.SearchVirtualizationPolicies(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}