# \LocationInformationCollectionApi

All URIs are relative to *https://example.com/ss-lair/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RetrieveUeLocInfo**](LocationInformationCollectionApi.md#RetrieveUeLocInfo) | **Get** /location-retrievals | 



## RetrieveUeLocInfo

> []LMInformation RetrieveUeLocInfo(ctx).LocationInfo(locationInfo).Range_(range_).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/SS_LocationAreaInfoRetrieval"
)

func main() {
    locationInfo := map[string][]openapiclient.LocationInfo{ ... } // LocationInfo | Location information around which the UE(s) information is requested.
    range_ := float32(3.4) // float32 | The range information over which the UE(s) information is required, expressed in meters. 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LocationInformationCollectionApi.RetrieveUeLocInfo(context.Background()).LocationInfo(locationInfo).Range_(range_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LocationInformationCollectionApi.RetrieveUeLocInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveUeLocInfo`: []LMInformation
    fmt.Fprintf(os.Stdout, "Response from `LocationInformationCollectionApi.RetrieveUeLocInfo`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveUeLocInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **locationInfo** | [**LocationInfo**](LocationInfo.md) | Location information around which the UE(s) information is requested. | 
 **range_** | **float32** | The range information over which the UE(s) information is required, expressed in meters.  | 

### Return type

[**[]LMInformation**](LMInformation.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

