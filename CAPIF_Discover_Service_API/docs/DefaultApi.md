# \DefaultApi

All URIs are relative to *https://example.com/service-apis/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AllServiceAPIsGet**](DefaultApi.md#AllServiceAPIsGet) | **Get** /allServiceAPIs | 



## AllServiceAPIsGet

> DiscoveredAPIs AllServiceAPIsGet(ctx).ApiInvokerId(apiInvokerId).ApiName(apiName).ApiVersion(apiVersion).CommType(commType).Protocol(protocol).AefId(aefId).DataFormat(dataFormat).ApiCat(apiCat).PreferredAefLoc(preferredAefLoc).SupportedFeatures(supportedFeatures).ApiSupportedFeatures(apiSupportedFeatures).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/CAPIF_Discover_Service_API"
)

func main() {
    apiInvokerId := "apiInvokerId_example" // string | String identifying the API invoker assigned by the CAPIF core function. It also represents the CCF identifier in the CAPIF-6/6e interface. 
    apiName := "apiName_example" // string | API name, it is set as {apiName} part of the URI structure as defined in clause 5.2.4 of 3GPP TS 29.122.  (optional)
    apiVersion := "apiVersion_example" // string | API major version the URI (e.g. v1). (optional)
    commType := *openapiclient.NewCommunicationType() // CommunicationType | Communication type used by the API (e.g. REQUEST_RESPONSE). (optional)
    protocol := *openapiclient.NewProtocol() // Protocol | Protocol used by the API. (optional)
    aefId := "aefId_example" // string | AEF identifer. (optional)
    dataFormat := *openapiclient.NewDataFormat() // DataFormat | Data formats used by the API (e.g. serialization protocol JSON used). (optional)
    apiCat := "apiCat_example" // string | The service API category to which the service API belongs to. (optional)
    preferredAefLoc := map[string][]openapiclient.AefLocation{ ... } // AefLocation | The preferred AEF location. (optional)
    supportedFeatures := "supportedFeatures_example" // string | Features supported by the NF consumer for the CAPIF Discover Service API. (optional)
    apiSupportedFeatures := "apiSupportedFeatures_example" // string | Features supported by the discovered service API indicated by api-name parameter. This may only be present if api-name query parameter is present.  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.AllServiceAPIsGet(context.Background()).ApiInvokerId(apiInvokerId).ApiName(apiName).ApiVersion(apiVersion).CommType(commType).Protocol(protocol).AefId(aefId).DataFormat(dataFormat).ApiCat(apiCat).PreferredAefLoc(preferredAefLoc).SupportedFeatures(supportedFeatures).ApiSupportedFeatures(apiSupportedFeatures).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.AllServiceAPIsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AllServiceAPIsGet`: DiscoveredAPIs
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.AllServiceAPIsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAllServiceAPIsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiInvokerId** | **string** | String identifying the API invoker assigned by the CAPIF core function. It also represents the CCF identifier in the CAPIF-6/6e interface.  | 
 **apiName** | **string** | API name, it is set as {apiName} part of the URI structure as defined in clause 5.2.4 of 3GPP TS 29.122.  | 
 **apiVersion** | **string** | API major version the URI (e.g. v1). | 
 **commType** | [**CommunicationType**](CommunicationType.md) | Communication type used by the API (e.g. REQUEST_RESPONSE). | 
 **protocol** | [**Protocol**](Protocol.md) | Protocol used by the API. | 
 **aefId** | **string** | AEF identifer. | 
 **dataFormat** | [**DataFormat**](DataFormat.md) | Data formats used by the API (e.g. serialization protocol JSON used). | 
 **apiCat** | **string** | The service API category to which the service API belongs to. | 
 **preferredAefLoc** | [**AefLocation**](AefLocation.md) | The preferred AEF location. | 
 **supportedFeatures** | **string** | Features supported by the NF consumer for the CAPIF Discover Service API. | 
 **apiSupportedFeatures** | **string** | Features supported by the discovered service API indicated by api-name parameter. This may only be present if api-name query parameter is present.  | 

### Return type

[**DiscoveredAPIs**](DiscoveredAPIs.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

