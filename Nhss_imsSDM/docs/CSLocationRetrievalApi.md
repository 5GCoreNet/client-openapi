# \CSLocationRetrievalApi

All URIs are relative to *https://example.com/nhss-ims-sdm/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetLocCsDomain**](CSLocationRetrievalApi.md#GetLocCsDomain) | **Get** /{imsUeId}/access-data/cs-domain/location-data | Retrieve the location data in CS domain



## GetLocCsDomain

> CsLocation GetLocCsDomain(ctx, imsUeId).ServingNode(servingNode).LocalTime(localTime).CurrentLocation(currentLocation).SupportedFeatures(supportedFeatures).PrivateIdentity(privateIdentity).Execute()

Retrieve the location data in CS domain

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
    imsUeId := "imsUeId_example" // string | IMS Public Identity
    servingNode := true // bool | Indicates that only the stored NF id/address of the serving node(s) is required  (optional)
    localTime := true // bool | Indicates that only the Local Time Zone information of the location in the visited network where the UE is attached is requested  (optional)
    currentLocation := true // bool | Indicates whether an active location retrieval has to be initiated by the requested node  (optional)
    supportedFeatures := "supportedFeatures_example" // string | Supported Features (optional)
    privateIdentity := "privateIdentity_example" // string | IMS Private Identity (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CSLocationRetrievalApi.GetLocCsDomain(context.Background(), imsUeId).ServingNode(servingNode).LocalTime(localTime).CurrentLocation(currentLocation).SupportedFeatures(supportedFeatures).PrivateIdentity(privateIdentity).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CSLocationRetrievalApi.GetLocCsDomain``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLocCsDomain`: CsLocation
    fmt.Fprintf(os.Stdout, "Response from `CSLocationRetrievalApi.GetLocCsDomain`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**imsUeId** | **string** | IMS Public Identity | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLocCsDomainRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **servingNode** | **bool** | Indicates that only the stored NF id/address of the serving node(s) is required  | 
 **localTime** | **bool** | Indicates that only the Local Time Zone information of the location in the visited network where the UE is attached is requested  | 
 **currentLocation** | **bool** | Indicates whether an active location retrieval has to be initiated by the requested node  | 
 **supportedFeatures** | **string** | Supported Features | 
 **privateIdentity** | **string** | IMS Private Identity | 

### Return type

[**CsLocation**](CsLocation.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

