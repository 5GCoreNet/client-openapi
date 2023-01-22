# \DefaultApi

All URIs are relative to *https://example.com/eecs-targeteesdiscovery/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EesProfilesGet**](DefaultApi.md#EesProfilesGet) | **Get** /ees-profiles | 



## EesProfilesGet

> ECSServProvResp EesProfilesGet(ctx).EesId(eesId).EasId(easId).TargetDnai(targetDnai).UeId(ueId).UeLocation(ueLocation).Execute()





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
    eesId := "eesId_example" // string | Unique identifier of the S-EES.
    easId := "easId_example" // string | Unique identifier of the S-EAS.
    targetDnai := "targetDnai_example" // string | The DNAI information associated with the potential T-EES(s) and/or T-EAS(s). (optional)
    ueId := "ueId_example" // string | Identifier of the UE. (optional)
    ueLocation := map[string][]openapiclient.LocationArea5G{ ... } // LocationArea5G | The location information of the UE. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.EesProfilesGet(context.Background()).EesId(eesId).EasId(easId).TargetDnai(targetDnai).UeId(ueId).UeLocation(ueLocation).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.EesProfilesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EesProfilesGet`: ECSServProvResp
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.EesProfilesGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEesProfilesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **eesId** | **string** | Unique identifier of the S-EES. | 
 **easId** | **string** | Unique identifier of the S-EAS. | 
 **targetDnai** | **string** | The DNAI information associated with the potential T-EES(s) and/or T-EAS(s). | 
 **ueId** | **string** | Identifier of the UE. | 
 **ueLocation** | [**LocationArea5G**](LocationArea5G.md) | The location information of the UE. | 

### Return type

[**ECSServProvResp**](ECSServProvResp.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

