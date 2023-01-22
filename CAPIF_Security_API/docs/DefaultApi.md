# \DefaultApi

All URIs are relative to *https://example.com/capif-security/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SecuritiesSecurityIdTokenPost**](DefaultApi.md#SecuritiesSecurityIdTokenPost) | **Post** /securities/{securityId}/token | 
[**TrustedInvokersApiInvokerIdDelete**](DefaultApi.md#TrustedInvokersApiInvokerIdDelete) | **Delete** /trustedInvokers/{apiInvokerId} | 
[**TrustedInvokersApiInvokerIdDeletePost**](DefaultApi.md#TrustedInvokersApiInvokerIdDeletePost) | **Post** /trustedInvokers/{apiInvokerId}/delete | 
[**TrustedInvokersApiInvokerIdGet**](DefaultApi.md#TrustedInvokersApiInvokerIdGet) | **Get** /trustedInvokers/{apiInvokerId} | 
[**TrustedInvokersApiInvokerIdPut**](DefaultApi.md#TrustedInvokersApiInvokerIdPut) | **Put** /trustedInvokers/{apiInvokerId} | 
[**TrustedInvokersApiInvokerIdUpdatePost**](DefaultApi.md#TrustedInvokersApiInvokerIdUpdatePost) | **Post** /trustedInvokers/{apiInvokerId}/update | 



## SecuritiesSecurityIdTokenPost

> AccessTokenRsp SecuritiesSecurityIdTokenPost(ctx, securityId).GrantType(grantType).ClientId(clientId).ClientSecret(clientSecret).Scope(scope).Execute()



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
    securityId := "securityId_example" // string | Identifier of an individual API invoker
    grantType := "grantType_example" // string | 
    clientId := "clientId_example" // string | 
    clientSecret := "clientSecret_example" // string |  (optional)
    scope := "scope_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.SecuritiesSecurityIdTokenPost(context.Background(), securityId).GrantType(grantType).ClientId(clientId).ClientSecret(clientSecret).Scope(scope).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.SecuritiesSecurityIdTokenPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SecuritiesSecurityIdTokenPost`: AccessTokenRsp
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.SecuritiesSecurityIdTokenPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**securityId** | **string** | Identifier of an individual API invoker | 

### Other Parameters

Other parameters are passed through a pointer to a apiSecuritiesSecurityIdTokenPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **grantType** | **string** |  | 
 **clientId** | **string** |  | 
 **clientSecret** | **string** |  | 
 **scope** | **string** |  | 

### Return type

[**AccessTokenRsp**](AccessTokenRsp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TrustedInvokersApiInvokerIdDelete

> TrustedInvokersApiInvokerIdDelete(ctx, apiInvokerId).Execute()



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
    apiInvokerId := "apiInvokerId_example" // string | Identifier of an individual API invoker

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.TrustedInvokersApiInvokerIdDelete(context.Background(), apiInvokerId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.TrustedInvokersApiInvokerIdDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiInvokerId** | **string** | Identifier of an individual API invoker | 

### Other Parameters

Other parameters are passed through a pointer to a apiTrustedInvokersApiInvokerIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TrustedInvokersApiInvokerIdDeletePost

> TrustedInvokersApiInvokerIdDeletePost(ctx, apiInvokerId).SecurityNotification(securityNotification).Execute()



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
    apiInvokerId := "apiInvokerId_example" // string | Identifier of an individual API invoker
    securityNotification := *openapiclient.NewSecurityNotification("ApiInvokerId_example", []string{"ApiIds_example"}, *openapiclient.NewCause()) // SecurityNotification | Revoke the authorization of the API invoker for APIs.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.TrustedInvokersApiInvokerIdDeletePost(context.Background(), apiInvokerId).SecurityNotification(securityNotification).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.TrustedInvokersApiInvokerIdDeletePost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiInvokerId** | **string** | Identifier of an individual API invoker | 

### Other Parameters

Other parameters are passed through a pointer to a apiTrustedInvokersApiInvokerIdDeletePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **securityNotification** | [**SecurityNotification**](SecurityNotification.md) | Revoke the authorization of the API invoker for APIs. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TrustedInvokersApiInvokerIdGet

> ServiceSecurity TrustedInvokersApiInvokerIdGet(ctx, apiInvokerId).AuthenticationInfo(authenticationInfo).AuthorizationInfo(authorizationInfo).Execute()



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
    apiInvokerId := "apiInvokerId_example" // string | Identifier of an individual API invoker
    authenticationInfo := true // bool | When set to 'true', it indicates the CAPIF core function to send the authentication information of the API invoker. Set to false or omitted otherwise.  (optional)
    authorizationInfo := true // bool | When set to 'true', it indicates the CAPIF core function to send the authorization information of the API invoker. Set to false or omitted otherwise.  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.TrustedInvokersApiInvokerIdGet(context.Background(), apiInvokerId).AuthenticationInfo(authenticationInfo).AuthorizationInfo(authorizationInfo).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.TrustedInvokersApiInvokerIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TrustedInvokersApiInvokerIdGet`: ServiceSecurity
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.TrustedInvokersApiInvokerIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiInvokerId** | **string** | Identifier of an individual API invoker | 

### Other Parameters

Other parameters are passed through a pointer to a apiTrustedInvokersApiInvokerIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authenticationInfo** | **bool** | When set to &#39;true&#39;, it indicates the CAPIF core function to send the authentication information of the API invoker. Set to false or omitted otherwise.  | 
 **authorizationInfo** | **bool** | When set to &#39;true&#39;, it indicates the CAPIF core function to send the authorization information of the API invoker. Set to false or omitted otherwise.  | 

### Return type

[**ServiceSecurity**](ServiceSecurity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TrustedInvokersApiInvokerIdPut

> ServiceSecurity TrustedInvokersApiInvokerIdPut(ctx, apiInvokerId).ServiceSecurity(serviceSecurity).Execute()



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
    apiInvokerId := "apiInvokerId_example" // string | Identifier of an individual API invoker
    serviceSecurity := *openapiclient.NewServiceSecurity([]openapiclient.SecurityInformation{openapiclient.SecurityInformation{Interface{}: new(interface{})}}, "NotificationDestination_example") // ServiceSecurity | create a security context for an API invoker

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.TrustedInvokersApiInvokerIdPut(context.Background(), apiInvokerId).ServiceSecurity(serviceSecurity).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.TrustedInvokersApiInvokerIdPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TrustedInvokersApiInvokerIdPut`: ServiceSecurity
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.TrustedInvokersApiInvokerIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiInvokerId** | **string** | Identifier of an individual API invoker | 

### Other Parameters

Other parameters are passed through a pointer to a apiTrustedInvokersApiInvokerIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **serviceSecurity** | [**ServiceSecurity**](ServiceSecurity.md) | create a security context for an API invoker | 

### Return type

[**ServiceSecurity**](ServiceSecurity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TrustedInvokersApiInvokerIdUpdatePost

> ServiceSecurity TrustedInvokersApiInvokerIdUpdatePost(ctx, apiInvokerId).ServiceSecurity(serviceSecurity).Execute()



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
    apiInvokerId := "apiInvokerId_example" // string | Identifier of an individual API invoker
    serviceSecurity := *openapiclient.NewServiceSecurity([]openapiclient.SecurityInformation{openapiclient.SecurityInformation{Interface{}: new(interface{})}}, "NotificationDestination_example") // ServiceSecurity | Update the security context (e.g. re-negotiate the security methods).

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.TrustedInvokersApiInvokerIdUpdatePost(context.Background(), apiInvokerId).ServiceSecurity(serviceSecurity).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.TrustedInvokersApiInvokerIdUpdatePost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TrustedInvokersApiInvokerIdUpdatePost`: ServiceSecurity
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.TrustedInvokersApiInvokerIdUpdatePost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiInvokerId** | **string** | Identifier of an individual API invoker | 

### Other Parameters

Other parameters are passed through a pointer to a apiTrustedInvokersApiInvokerIdUpdatePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **serviceSecurity** | [**ServiceSecurity**](ServiceSecurity.md) | Update the security context (e.g. re-negotiate the security methods). | 

### Return type

[**ServiceSecurity**](ServiceSecurity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

