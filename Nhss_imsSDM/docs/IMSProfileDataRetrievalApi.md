# \IMSProfileDataRetrievalApi

All URIs are relative to *https://example.com/nhss-ims-sdm/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetProfileData**](IMSProfileDataRetrievalApi.md#GetProfileData) | **Get** /{imsUeId}/ims-data/profile-data | Retrieve the complete IMS profile for a given IMS public identity (and public identities in the same IRS) 



## GetProfileData

> ImsProfileData GetProfileData(ctx, imsUeId).DatasetNames(datasetNames).Execute()

Retrieve the complete IMS profile for a given IMS public identity (and public identities in the same IRS) 

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/Nhss_imsSDM"
)

func main() {
    imsUeId := "imsUeId_example" // string | IMS Identity. In this case it shall be an IMS public identity
    datasetNames := []openapiclient.DataSetName{*openapiclient.NewDataSetName()} // []DataSetName | Datasets to be retrieved (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IMSProfileDataRetrievalApi.GetProfileData(context.Background(), imsUeId).DatasetNames(datasetNames).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IMSProfileDataRetrievalApi.GetProfileData``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetProfileData`: ImsProfileData
    fmt.Fprintf(os.Stdout, "Response from `IMSProfileDataRetrievalApi.GetProfileData`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**imsUeId** | **string** | IMS Identity. In this case it shall be an IMS public identity | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProfileDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **datasetNames** | [**[]DataSetName**](DataSetName.md) | Datasets to be retrieved | 

### Return type

[**ImsProfileData**](ImsProfileData.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

