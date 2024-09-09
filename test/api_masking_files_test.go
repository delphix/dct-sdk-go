/*
Delphix DCT API

Testing MaskingFilesApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package delphix_dct_api

import (
	"context"
	"testing"

	openapiclient "github.com/delphix/dct-sdk-go/v21"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_delphix_dct_api_MaskingFilesApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test MaskingFilesApiService UploadMaskingFile", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.MaskingFilesApi.UploadMaskingFile(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
