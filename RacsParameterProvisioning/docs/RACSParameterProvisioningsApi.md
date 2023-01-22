# \RACSParameterProvisioningsApi

All URIs are relative to *https://example.com/3gpp-racs-pp/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateRACSParameterProvisioning**](RACSParameterProvisioningsApi.md#CreateRACSParameterProvisioning) | **Post** /{scsAsId}/provisionings | Create a new RACS parameter provisioning.
[**FetchAllRACSParameterProvisionings**](RACSParameterProvisioningsApi.md#FetchAllRACSParameterProvisionings) | **Get** /{scsAsId}/provisionings | Read all RACS parameter provisionings for a given AF.



## CreateRACSParameterProvisioning

> RacsProvisioningData CreateRACSParameterProvisioning(ctx, scsAsId).RacsProvisioningData(racsProvisioningData).Execute()

Create a new RACS parameter provisioning.

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
    scsAsId := "scsAsId_example" // string | Identifier of the SCS/AS as defined in clause 5.2.4 of 3GPP TS 29.122.
    racsProvisioningData := *openapiclient.NewRacsProvisioningData(map[string]RacsConfiguration{"key": *openapiclient.NewRacsConfiguration("RacsId_example", []string{"ImeiTacs_example"})}) // RacsProvisioningData | create new provisionings for a given SCS/AS.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RACSParameterProvisioningsApi.CreateRACSParameterProvisioning(context.Background(), scsAsId).RacsProvisioningData(racsProvisioningData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RACSParameterProvisioningsApi.CreateRACSParameterProvisioning``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateRACSParameterProvisioning`: RacsProvisioningData
    fmt.Fprintf(os.Stdout, "Response from `RACSParameterProvisioningsApi.CreateRACSParameterProvisioning`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**scsAsId** | **string** | Identifier of the SCS/AS as defined in clause 5.2.4 of 3GPP TS 29.122. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateRACSParameterProvisioningRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **racsProvisioningData** | [**RacsProvisioningData**](RacsProvisioningData.md) | create new provisionings for a given SCS/AS. | 

### Return type

[**RacsProvisioningData**](RacsProvisioningData.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchAllRACSParameterProvisionings

> []RacsProvisioningData FetchAllRACSParameterProvisionings(ctx, scsAsId).Execute()

Read all RACS parameter provisionings for a given AF.

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
    scsAsId := "scsAsId_example" // string | Identifier of the SCS/AS as defined in clause 5.2.4 of 3GPP TS 29.122.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RACSParameterProvisioningsApi.FetchAllRACSParameterProvisionings(context.Background(), scsAsId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RACSParameterProvisioningsApi.FetchAllRACSParameterProvisionings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchAllRACSParameterProvisionings`: []RacsProvisioningData
    fmt.Fprintf(os.Stdout, "Response from `RACSParameterProvisioningsApi.FetchAllRACSParameterProvisionings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**scsAsId** | **string** | Identifier of the SCS/AS as defined in clause 5.2.4 of 3GPP TS 29.122. | 

### Other Parameters

Other parameters are passed through a pointer to a apiFetchAllRACSParameterProvisioningsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]RacsProvisioningData**](RacsProvisioningData.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

