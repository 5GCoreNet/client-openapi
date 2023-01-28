# \NetworkSliceAdaptationRequestApi

All URIs are relative to *https://example.com/ss-nsa/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RequestNetworkSliceAdaptation**](NetworkSliceAdaptationRequestApi.md#RequestNetworkSliceAdaptation) | **Post** /request | request the network slice adaptation.



## RequestNetworkSliceAdaptation

> RequestNetworkSliceAdaptation(ctx).NwSliceAdptInfo(nwSliceAdptInfo).Execute()

request the network slice adaptation.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/SS_NetworkSliceAdaptation"
)

func main() {
    nwSliceAdptInfo := *openapiclient.NewNwSliceAdptInfo("ValServiceId_example", []string{"ValTgtUeIds_example"}) // NwSliceAdptInfo | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworkSliceAdaptationRequestApi.RequestNetworkSliceAdaptation(context.Background()).NwSliceAdptInfo(nwSliceAdptInfo).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkSliceAdaptationRequestApi.RequestNetworkSliceAdaptation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRequestNetworkSliceAdaptationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **nwSliceAdptInfo** | [**NwSliceAdptInfo**](NwSliceAdptInfo.md) |  | 

### Return type

 (empty response body)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

