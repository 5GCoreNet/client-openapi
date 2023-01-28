# \DSAIRegistrationInformationApi

All URIs are relative to *https://example.com/nhss-ims-sdm/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetDsaiInfo**](DSAIRegistrationInformationApi.md#GetDsaiInfo) | **Get** /{imsUeId}/service-data/dsai | Retrieve the DSAI information associated to an Application Server



## GetDsaiInfo

> DsaiTagInformation GetDsaiInfo(ctx, imsUeId).ApplicationServerName(applicationServerName).DsaiTag(dsaiTag).SupportedFeatures(supportedFeatures).Execute()

Retrieve the DSAI information associated to an Application Server

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
    imsUeId := "imsUeId_example" // string | IMS Identity
    applicationServerName := "applicationServerName_example" // string | SIP URI of the Application Server Name
    dsaiTag := "dsaiTag_example" // string | The requested instance of Dynamic Service Activation Info (optional)
    supportedFeatures := "supportedFeatures_example" // string | Supported Features (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DSAIRegistrationInformationApi.GetDsaiInfo(context.Background(), imsUeId).ApplicationServerName(applicationServerName).DsaiTag(dsaiTag).SupportedFeatures(supportedFeatures).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DSAIRegistrationInformationApi.GetDsaiInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDsaiInfo`: DsaiTagInformation
    fmt.Fprintf(os.Stdout, "Response from `DSAIRegistrationInformationApi.GetDsaiInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**imsUeId** | **string** | IMS Identity | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDsaiInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **applicationServerName** | **string** | SIP URI of the Application Server Name | 
 **dsaiTag** | **string** | The requested instance of Dynamic Service Activation Info | 
 **supportedFeatures** | **string** | Supported Features | 

### Return type

[**DsaiTagInformation**](DsaiTagInformation.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

