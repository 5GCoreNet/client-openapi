# \IndividualMBSParametersProvisioningApi

All URIs are relative to *https://example.com/3gpp-mbs-session/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteIndMBSParamsProvisioning**](IndividualMBSParametersProvisioningApi.md#DeleteIndMBSParamsProvisioning) | **Delete** /mbs-pp/{mbsPpId} | Request the deletion of an existing Individual MBS Parameters Provisioning resource.
[**GetIndMBSParamsProvisioning**](IndividualMBSParametersProvisioningApi.md#GetIndMBSParamsProvisioning) | **Get** /mbs-pp/{mbsPpId} | Request to retrieve an existing Individual MBS Parameters Provisioning resource.
[**ModifyIndMBSParamsProvisioning**](IndividualMBSParametersProvisioningApi.md#ModifyIndMBSParamsProvisioning) | **Patch** /mbs-pp/{mbsPpId} | Request the modification of an existing Individual MBS Parameters Provisioning resource.
[**UpdateIndMBSParamsProvisioning**](IndividualMBSParametersProvisioningApi.md#UpdateIndMBSParamsProvisioning) | **Put** /mbs-pp/{mbsPpId} | Request the update of an existing Individual MBS Parameters Provisioning resource.



## DeleteIndMBSParamsProvisioning

> DeleteIndMBSParamsProvisioning(ctx, mbsPpId).Execute()

Request the deletion of an existing Individual MBS Parameters Provisioning resource.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/MBSSession"
)

func main() {
    mbsPpId := "mbsPpId_example" // string | Represents the identifier of the Individual MBS Parameters Provisioning resource. 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualMBSParametersProvisioningApi.DeleteIndMBSParamsProvisioning(context.Background(), mbsPpId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualMBSParametersProvisioningApi.DeleteIndMBSParamsProvisioning``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mbsPpId** | **string** | Represents the identifier of the Individual MBS Parameters Provisioning resource.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteIndMBSParamsProvisioningRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIndMBSParamsProvisioning

> MbsPpData GetIndMBSParamsProvisioning(ctx, mbsPpId).Execute()

Request to retrieve an existing Individual MBS Parameters Provisioning resource.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/MBSSession"
)

func main() {
    mbsPpId := "mbsPpId_example" // string | Represents the identifier of the Individual MBS Parameters Provisioning resource. 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualMBSParametersProvisioningApi.GetIndMBSParamsProvisioning(context.Background(), mbsPpId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualMBSParametersProvisioningApi.GetIndMBSParamsProvisioning``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetIndMBSParamsProvisioning`: MbsPpData
    fmt.Fprintf(os.Stdout, "Response from `IndividualMBSParametersProvisioningApi.GetIndMBSParamsProvisioning`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mbsPpId** | **string** | Represents the identifier of the Individual MBS Parameters Provisioning resource.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIndMBSParamsProvisioningRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MbsPpData**](MbsPpData.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyIndMBSParamsProvisioning

> MbsPpData ModifyIndMBSParamsProvisioning(ctx, mbsPpId).MbsPpDataPatch(mbsPpDataPatch).Execute()

Request the modification of an existing Individual MBS Parameters Provisioning resource.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/MBSSession"
)

func main() {
    mbsPpId := "mbsPpId_example" // string | Represents the identifier of the Individual MBS Parameters Provisioning resource. 
    mbsPpDataPatch := *openapiclient.NewMbsPpDataPatch() // MbsPpDataPatch | Contains the parameters to request the modification of the Individual Parameters Provisioning resource. 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualMBSParametersProvisioningApi.ModifyIndMBSParamsProvisioning(context.Background(), mbsPpId).MbsPpDataPatch(mbsPpDataPatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualMBSParametersProvisioningApi.ModifyIndMBSParamsProvisioning``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ModifyIndMBSParamsProvisioning`: MbsPpData
    fmt.Fprintf(os.Stdout, "Response from `IndividualMBSParametersProvisioningApi.ModifyIndMBSParamsProvisioning`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mbsPpId** | **string** | Represents the identifier of the Individual MBS Parameters Provisioning resource.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyIndMBSParamsProvisioningRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **mbsPpDataPatch** | [**MbsPpDataPatch**](MbsPpDataPatch.md) | Contains the parameters to request the modification of the Individual Parameters Provisioning resource.  | 

### Return type

[**MbsPpData**](MbsPpData.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/merge-patch+json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateIndMBSParamsProvisioning

> MbsPpData UpdateIndMBSParamsProvisioning(ctx, mbsPpId).MbsPpData(mbsPpData).Execute()

Request the update of an existing Individual MBS Parameters Provisioning resource.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/MBSSession"
)

func main() {
    mbsPpId := "mbsPpId_example" // string | Represents the identifier of the Individual MBS Parameters Provisioning resource. 
    mbsPpData := *openapiclient.NewMbsPpData("AfId_example") // MbsPpData | Represents the updated Individual MBS Parameters Provisioning resource representation. 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualMBSParametersProvisioningApi.UpdateIndMBSParamsProvisioning(context.Background(), mbsPpId).MbsPpData(mbsPpData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualMBSParametersProvisioningApi.UpdateIndMBSParamsProvisioning``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateIndMBSParamsProvisioning`: MbsPpData
    fmt.Fprintf(os.Stdout, "Response from `IndividualMBSParametersProvisioningApi.UpdateIndMBSParamsProvisioning`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mbsPpId** | **string** | Represents the identifier of the Individual MBS Parameters Provisioning resource.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateIndMBSParamsProvisioningRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **mbsPpData** | [**MbsPpData**](MbsPpData.md) | Represents the updated Individual MBS Parameters Provisioning resource representation.  | 

### Return type

[**MbsPpData**](MbsPpData.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

