# \MaskingFilesApi

All URIs are relative to */v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UploadMaskingFile**](MaskingFilesApi.md#UploadMaskingFile) | **Post** /masking-file-uploads | Upload a file to a masking engine.



## UploadMaskingFile

> MaskingFileUploadResponse UploadMaskingFile(ctx).MaskingFileUploadParameters(maskingFileUploadParameters).Execute()

Upload a file to a masking engine.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/delphix/dct-sdk-go"
)

func main() {
    maskingFileUploadParameters := *openapiclient.NewMaskingFileUploadParameters("1", "TODO") // MaskingFileUploadParameters | The parameters to upload a file to a masking engine. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MaskingFilesApi.UploadMaskingFile(context.Background()).MaskingFileUploadParameters(maskingFileUploadParameters).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MaskingFilesApi.UploadMaskingFile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UploadMaskingFile`: MaskingFileUploadResponse
    fmt.Fprintf(os.Stdout, "Response from `MaskingFilesApi.UploadMaskingFile`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUploadMaskingFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **maskingFileUploadParameters** | [**MaskingFileUploadParameters**](MaskingFileUploadParameters.md) | The parameters to upload a file to a masking engine. | 

### Return type

[**MaskingFileUploadResponse**](MaskingFileUploadResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

