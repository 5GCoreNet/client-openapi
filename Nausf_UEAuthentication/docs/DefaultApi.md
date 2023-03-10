# \DefaultApi

All URIs are relative to *https://example.com/nausf-auth/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EapAuthMethod**](DefaultApi.md#EapAuthMethod) | **Post** /ue-authentications/{authCtxId}/eap-session | 
[**ProseAuth**](DefaultApi.md#ProseAuth) | **Post** /prose-authentications/{authCtxId}/prose-auth | 
[**ProseAuthenticationsPost**](DefaultApi.md#ProseAuthenticationsPost) | **Post** /prose-authentications | 
[**RgAuthenticationsPost**](DefaultApi.md#RgAuthenticationsPost) | **Post** /rg-authentications | 
[**UeAuthenticationsAuthCtxId5gAkaConfirmationPut**](DefaultApi.md#UeAuthenticationsAuthCtxId5gAkaConfirmationPut) | **Put** /ue-authentications/{authCtxId}/5g-aka-confirmation | 
[**UeAuthenticationsDeregisterPost**](DefaultApi.md#UeAuthenticationsDeregisterPost) | **Post** /ue-authentications/deregister | 
[**UeAuthenticationsPost**](DefaultApi.md#UeAuthenticationsPost) | **Post** /ue-authentications | 



## EapAuthMethod

> EapSession EapAuthMethod(ctx, authCtxId).EapSession(eapSession).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/Nausf_UEAuthentication"
)

func main() {
    authCtxId := "authCtxId_example" // string | 
    eapSession := *openapiclient.NewEapSession(NullableString(123)) // EapSession |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.EapAuthMethod(context.Background(), authCtxId).EapSession(eapSession).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.EapAuthMethod``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EapAuthMethod`: EapSession
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.EapAuthMethod`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authCtxId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEapAuthMethodRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **eapSession** | [**EapSession**](EapSession.md) |  | 

### Return type

[**EapSession**](EapSession.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/3gppHal+json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProseAuth

> ProSeEapSession ProseAuth(ctx, authCtxId).ProSeEapSession(proSeEapSession).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/Nausf_UEAuthentication"
)

func main() {
    authCtxId := "authCtxId_example" // string | 
    proSeEapSession := *openapiclient.NewProSeEapSession(NullableString(123)) // ProSeEapSession |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ProseAuth(context.Background(), authCtxId).ProSeEapSession(proSeEapSession).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ProseAuth``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProseAuth`: ProSeEapSession
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ProseAuth`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authCtxId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiProseAuthRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **proSeEapSession** | [**ProSeEapSession**](ProSeEapSession.md) |  | 

### Return type

[**ProSeEapSession**](ProSeEapSession.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/3gppHal+json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProseAuthenticationsPost

> ProSeAuthenticationResult ProseAuthenticationsPost(ctx).ProSeAuthenticationInfo(proSeAuthenticationInfo).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/Nausf_UEAuthentication"
)

func main() {
    proSeAuthenticationInfo := *openapiclient.NewProSeAuthenticationInfo(int32(123), NullableString(123)) // ProSeAuthenticationInfo | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ProseAuthenticationsPost(context.Background()).ProSeAuthenticationInfo(proSeAuthenticationInfo).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ProseAuthenticationsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProseAuthenticationsPost`: ProSeAuthenticationResult
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ProseAuthenticationsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProseAuthenticationsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **proSeAuthenticationInfo** | [**ProSeAuthenticationInfo**](ProSeAuthenticationInfo.md) |  | 

### Return type

[**ProSeAuthenticationResult**](ProSeAuthenticationResult.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/3gppHal+json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RgAuthenticationsPost

> RgAuthCtx RgAuthenticationsPost(ctx).RgAuthenticationInfo(rgAuthenticationInfo).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/Nausf_UEAuthentication"
)

func main() {
    rgAuthenticationInfo := *openapiclient.NewRgAuthenticationInfo("Suci_example", false) // RgAuthenticationInfo | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.RgAuthenticationsPost(context.Background()).RgAuthenticationInfo(rgAuthenticationInfo).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.RgAuthenticationsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RgAuthenticationsPost`: RgAuthCtx
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.RgAuthenticationsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRgAuthenticationsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **rgAuthenticationInfo** | [**RgAuthenticationInfo**](RgAuthenticationInfo.md) |  | 

### Return type

[**RgAuthCtx**](RgAuthCtx.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UeAuthenticationsAuthCtxId5gAkaConfirmationPut

> ConfirmationDataResponse UeAuthenticationsAuthCtxId5gAkaConfirmationPut(ctx, authCtxId).ConfirmationData(confirmationData).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/Nausf_UEAuthentication"
)

func main() {
    authCtxId := "authCtxId_example" // string | 
    confirmationData := *openapiclient.NewConfirmationData("ResStar_example") // ConfirmationData |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.UeAuthenticationsAuthCtxId5gAkaConfirmationPut(context.Background(), authCtxId).ConfirmationData(confirmationData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UeAuthenticationsAuthCtxId5gAkaConfirmationPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UeAuthenticationsAuthCtxId5gAkaConfirmationPut`: ConfirmationDataResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UeAuthenticationsAuthCtxId5gAkaConfirmationPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authCtxId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUeAuthenticationsAuthCtxId5gAkaConfirmationPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **confirmationData** | [**ConfirmationData**](ConfirmationData.md) |  | 

### Return type

[**ConfirmationDataResponse**](ConfirmationDataResponse.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UeAuthenticationsDeregisterPost

> UeAuthenticationsDeregisterPost(ctx).DeregistrationInfo(deregistrationInfo).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/Nausf_UEAuthentication"
)

func main() {
    deregistrationInfo := *openapiclient.NewDeregistrationInfo("Supi_example") // DeregistrationInfo | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.UeAuthenticationsDeregisterPost(context.Background()).DeregistrationInfo(deregistrationInfo).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UeAuthenticationsDeregisterPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUeAuthenticationsDeregisterPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deregistrationInfo** | [**DeregistrationInfo**](DeregistrationInfo.md) |  | 

### Return type

 (empty response body)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UeAuthenticationsPost

> UEAuthenticationCtx UeAuthenticationsPost(ctx).AuthenticationInfo(authenticationInfo).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/Nausf_UEAuthentication"
)

func main() {
    authenticationInfo := *openapiclient.NewAuthenticationInfo("SupiOrSuci_example", "ServingNetworkName_example") // AuthenticationInfo | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.UeAuthenticationsPost(context.Background()).AuthenticationInfo(authenticationInfo).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UeAuthenticationsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UeAuthenticationsPost`: UEAuthenticationCtx
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UeAuthenticationsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUeAuthenticationsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authenticationInfo** | [**AuthenticationInfo**](AuthenticationInfo.md) |  | 

### Return type

[**UEAuthenticationCtx**](UEAuthenticationCtx.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/3gppHal+json, application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

