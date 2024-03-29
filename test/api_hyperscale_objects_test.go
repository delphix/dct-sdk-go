/*
Delphix DCT API

Testing HyperscaleObjectsApiService

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

func Test_delphix_dct_api_HyperscaleObjectsApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test HyperscaleObjectsApiService GetHyperscaleConnectorById", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var hyperscaleConnectorId string

		resp, httpRes, err := apiClient.HyperscaleObjectsApi.GetHyperscaleConnectorById(context.Background(), hyperscaleConnectorId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test HyperscaleObjectsApiService GetHyperscaleConnectors", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.HyperscaleObjectsApi.GetHyperscaleConnectors(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test HyperscaleObjectsApiService GetHyperscaleDatasetById", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var hyperscaleDatasetId string

		resp, httpRes, err := apiClient.HyperscaleObjectsApi.GetHyperscaleDatasetById(context.Background(), hyperscaleDatasetId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test HyperscaleObjectsApiService GetHyperscaleDatasetTablesOrFiles", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var hyperscaleDatasetId string

		resp, httpRes, err := apiClient.HyperscaleObjectsApi.GetHyperscaleDatasetTablesOrFiles(context.Background(), hyperscaleDatasetId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test HyperscaleObjectsApiService GetHyperscaleDatasets", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.HyperscaleObjectsApi.GetHyperscaleDatasets(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test HyperscaleObjectsApiService GetHyperscaleMountPointById", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var hyperscaleMountPointId string

		resp, httpRes, err := apiClient.HyperscaleObjectsApi.GetHyperscaleMountPointById(context.Background(), hyperscaleMountPointId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test HyperscaleObjectsApiService GetHyperscaleMountPoints", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.HyperscaleObjectsApi.GetHyperscaleMountPoints(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test HyperscaleObjectsApiService SearchHyperscaleConnectors", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.HyperscaleObjectsApi.SearchHyperscaleConnectors(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test HyperscaleObjectsApiService SearchHyperscaleDatasetTablesOrFiles", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var hyperscaleDatasetId string

		resp, httpRes, err := apiClient.HyperscaleObjectsApi.SearchHyperscaleDatasetTablesOrFiles(context.Background(), hyperscaleDatasetId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test HyperscaleObjectsApiService SearchHyperscaleDatasets", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.HyperscaleObjectsApi.SearchHyperscaleDatasets(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test HyperscaleObjectsApiService SearchHyperscaleMountPoints", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.HyperscaleObjectsApi.SearchHyperscaleMountPoints(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
