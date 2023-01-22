# \SMContextsCollectionApi

All URIs are relative to *https://example.com/nsmf-pdusession/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostSmContexts**](SMContextsCollectionApi.md#PostSmContexts) | **Post** /sm-contexts | Create SM Context



## PostSmContexts

> SmContextCreatedData PostSmContexts(ctx).JsonData(jsonData).BinaryDataN1SmMessage(binaryDataN1SmMessage).BinaryDataN2SmInformation(binaryDataN2SmInformation).BinaryDataN2SmInformationExt1(binaryDataN2SmInformationExt1).Execute()

Create SM Context

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    jsonData := *openapiclient.NewSmContextCreateData("ServingNfId_example", *openapiclient.NewPlmnIdNid("Mcc_example", "Mnc_example"), openapiclient.AccessType("3GPP_ACCESS"), "SmContextStatusUri_example") // SmContextCreateData |  (optional)
    binaryDataN1SmMessage := os.NewFile(1234, "some_file") // *os.File |  (optional)
    binaryDataN2SmInformation := os.NewFile(1234, "some_file") // *os.File |  (optional)
    binaryDataN2SmInformationExt1 := os.NewFile(1234, "some_file") // *os.File |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SMContextsCollectionApi.PostSmContexts(context.Background()).JsonData(jsonData).BinaryDataN1SmMessage(binaryDataN1SmMessage).BinaryDataN2SmInformation(binaryDataN2SmInformation).BinaryDataN2SmInformationExt1(binaryDataN2SmInformationExt1).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SMContextsCollectionApi.PostSmContexts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostSmContexts`: SmContextCreatedData
    fmt.Fprintf(os.Stdout, "Response from `SMContextsCollectionApi.PostSmContexts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostSmContextsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **jsonData** | [**SmContextCreateData**](SmContextCreateData.md) |  | 
 **binaryDataN1SmMessage** | ***os.File** |  | 
 **binaryDataN2SmInformation** | ***os.File** |  | 
 **binaryDataN2SmInformationExt1** | ***os.File** |  | 

### Return type

[**SmContextCreatedData**](SmContextCreatedData.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: multipart/related
- **Accept**: application/json, multipart/related, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

