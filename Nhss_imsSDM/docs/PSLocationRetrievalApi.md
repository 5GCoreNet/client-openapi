# \PSLocationRetrievalApi

All URIs are relative to *https://example.com/nhss-ims-sdm/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetLocPsDomain**](PSLocationRetrievalApi.md#GetLocPsDomain) | **Get** /{imsUeId}/access-data/ps-domain/location-data | Retrieve the location data in PS domain



## GetLocPsDomain

> PsLocation GetLocPsDomain(ctx, imsUeId).RequestedNodes(requestedNodes).ServingNode(servingNode).LocalTime(localTime).CurrentLocation(currentLocation).RatType(ratType).SupportedFeatures(supportedFeatures).PrivateIdentity(privateIdentity).Execute()

Retrieve the location data in PS domain

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
    imsUeId := "imsUeId_example" // string | IMS Public Identity
    requestedNodes := []openapiclient.RequestedNode{*openapiclient.NewRequestedNode()} // []RequestedNode | Indicates the serving node(s) for which the request is applicable. (optional)
    servingNode := true // bool | Indicates that only the stored NF id/address of the serving node(s) is required  (optional)
    localTime := true // bool | Indicates that only the Local Time Zone information of the location in the visited network where the UE is attached is requested  (optional)
    currentLocation := true // bool | Indicates whether an active location retrieval has to be initiated by the requested node  (optional)
    ratType := true // bool | Indicates whether RAT Type retrieval is requested (optional)
    supportedFeatures := "supportedFeatures_example" // string | Supported Features (optional)
    privateIdentity := "privateIdentity_example" // string | IMS Private Identity (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PSLocationRetrievalApi.GetLocPsDomain(context.Background(), imsUeId).RequestedNodes(requestedNodes).ServingNode(servingNode).LocalTime(localTime).CurrentLocation(currentLocation).RatType(ratType).SupportedFeatures(supportedFeatures).PrivateIdentity(privateIdentity).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PSLocationRetrievalApi.GetLocPsDomain``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLocPsDomain`: PsLocation
    fmt.Fprintf(os.Stdout, "Response from `PSLocationRetrievalApi.GetLocPsDomain`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**imsUeId** | **string** | IMS Public Identity | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLocPsDomainRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **requestedNodes** | [**[]RequestedNode**](RequestedNode.md) | Indicates the serving node(s) for which the request is applicable. | 
 **servingNode** | **bool** | Indicates that only the stored NF id/address of the serving node(s) is required  | 
 **localTime** | **bool** | Indicates that only the Local Time Zone information of the location in the visited network where the UE is attached is requested  | 
 **currentLocation** | **bool** | Indicates whether an active location retrieval has to be initiated by the requested node  | 
 **ratType** | **bool** | Indicates whether RAT Type retrieval is requested | 
 **supportedFeatures** | **string** | Supported Features | 
 **privateIdentity** | **string** | IMS Private Identity | 

### Return type

[**PsLocation**](PsLocation.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

