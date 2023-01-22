# \IndividualChargeablePartyResourceOperationApi

All URIs are relative to *https://example.com/3gpp-chargeable-party/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteChargeablePartyTransaction**](IndividualChargeablePartyResourceOperationApi.md#DeleteChargeablePartyTransaction) | **Delete** /{scsAsId}/transactions/{transactionId} | Deletes a chargeable party resource for a given SCS/AS and a transcation Id.
[**FetchIndChargeablePartyTransaction**](IndividualChargeablePartyResourceOperationApi.md#FetchIndChargeablePartyTransaction) | **Get** /{scsAsId}/transactions/{transactionId} | Read a chargeable party resource for a given SCS/AS and a transaction Id.
[**UpdateChargeablePartyTransaction**](IndividualChargeablePartyResourceOperationApi.md#UpdateChargeablePartyTransaction) | **Patch** /{scsAsId}/transactions/{transactionId} | Updates a existing chargeable party resource for a given SCS/AS and transaction Id.



## DeleteChargeablePartyTransaction

> NotificationData DeleteChargeablePartyTransaction(ctx, scsAsId, transactionId).Execute()

Deletes a chargeable party resource for a given SCS/AS and a transcation Id.

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
    scsAsId := "scsAsId_example" // string | Identifier of SCS/AS
    transactionId := "transactionId_example" // string | Identifier of transaction

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualChargeablePartyResourceOperationApi.DeleteChargeablePartyTransaction(context.Background(), scsAsId, transactionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualChargeablePartyResourceOperationApi.DeleteChargeablePartyTransaction``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteChargeablePartyTransaction`: NotificationData
    fmt.Fprintf(os.Stdout, "Response from `IndividualChargeablePartyResourceOperationApi.DeleteChargeablePartyTransaction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**scsAsId** | **string** | Identifier of SCS/AS | 
**transactionId** | **string** | Identifier of transaction | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteChargeablePartyTransactionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**NotificationData**](NotificationData.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchIndChargeablePartyTransaction

> ChargeableParty FetchIndChargeablePartyTransaction(ctx, scsAsId, transactionId).Execute()

Read a chargeable party resource for a given SCS/AS and a transaction Id.

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
    scsAsId := "scsAsId_example" // string | Identifier of SCS/AS
    transactionId := "transactionId_example" // string | Identifier of transaction

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualChargeablePartyResourceOperationApi.FetchIndChargeablePartyTransaction(context.Background(), scsAsId, transactionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualChargeablePartyResourceOperationApi.FetchIndChargeablePartyTransaction``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchIndChargeablePartyTransaction`: ChargeableParty
    fmt.Fprintf(os.Stdout, "Response from `IndividualChargeablePartyResourceOperationApi.FetchIndChargeablePartyTransaction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**scsAsId** | **string** | Identifier of SCS/AS | 
**transactionId** | **string** | Identifier of transaction | 

### Other Parameters

Other parameters are passed through a pointer to a apiFetchIndChargeablePartyTransactionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ChargeableParty**](ChargeableParty.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateChargeablePartyTransaction

> ChargeableParty UpdateChargeablePartyTransaction(ctx, scsAsId, transactionId).ChargeablePartyPatch(chargeablePartyPatch).Execute()

Updates a existing chargeable party resource for a given SCS/AS and transaction Id.

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
    scsAsId := "scsAsId_example" // string | Identifier of SCS/AS
    transactionId := "transactionId_example" // string | Identifier of transaction
    chargeablePartyPatch := *openapiclient.NewChargeablePartyPatch() // ChargeablePartyPatch | representation of the chargeable party resource to be udpated in the SCEF

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualChargeablePartyResourceOperationApi.UpdateChargeablePartyTransaction(context.Background(), scsAsId, transactionId).ChargeablePartyPatch(chargeablePartyPatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualChargeablePartyResourceOperationApi.UpdateChargeablePartyTransaction``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateChargeablePartyTransaction`: ChargeableParty
    fmt.Fprintf(os.Stdout, "Response from `IndividualChargeablePartyResourceOperationApi.UpdateChargeablePartyTransaction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**scsAsId** | **string** | Identifier of SCS/AS | 
**transactionId** | **string** | Identifier of transaction | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateChargeablePartyTransactionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **chargeablePartyPatch** | [**ChargeablePartyPatch**](ChargeablePartyPatch.md) | representation of the chargeable party resource to be udpated in the SCEF | 

### Return type

[**ChargeableParty**](ChargeableParty.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/merge-patch+json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

