# \MBSDistributionSessionsCollectionApi

All URIs are relative to *https://example.com/nmbstf-distsession/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Create**](MBSDistributionSessionsCollectionApi.md#Create) | **Post** /dist-sessions | Create



## Create

> CreateRspData Create(ctx).CreateReqData(createReqData).Execute()

Create

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/Nmbstf_DistSession"
)

func main() {
    createReqData := *openapiclient.NewCreateReqData(openapiclient.DistSession{Interface{}: new(interface{})}) // CreateReqData | Representation of the MBS distribution session to be created in the MBSTF Creates an individual MBS distribution session resource in the MBSTF. 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MBSDistributionSessionsCollectionApi.Create(context.Background()).CreateReqData(createReqData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MBSDistributionSessionsCollectionApi.Create``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Create`: CreateRspData
    fmt.Fprintf(os.Stdout, "Response from `MBSDistributionSessionsCollectionApi.Create`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createReqData** | [**CreateReqData**](CreateReqData.md) | Representation of the MBS distribution session to be created in the MBSTF Creates an individual MBS distribution session resource in the MBSTF.  | 

### Return type

[**CreateRspData**](CreateRspData.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

