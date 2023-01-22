# \ChargeablePartyTransactionOperationApi

All URIs are relative to *https://example.com/3gpp-chargeable-party/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateChargeablePartyTransaction**](ChargeablePartyTransactionOperationApi.md#CreateChargeablePartyTransaction) | **Post** /{scsAsId}/transactions | Create a new chargeable party transaction resource.
[**FetchAllChargeablePartyTransactions**](ChargeablePartyTransactionOperationApi.md#FetchAllChargeablePartyTransactions) | **Get** /{scsAsId}/transactions | Read all or queried chargeable party transaction resources for a given SCS/AS.



## CreateChargeablePartyTransaction

> ChargeableParty CreateChargeablePartyTransaction(ctx, scsAsId).ChargeableParty(chargeableParty).Execute()

Create a new chargeable party transaction resource.

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
    chargeableParty := *openapiclient.NewChargeableParty("NotificationDestination_example", *openapiclient.NewSponsorInformation("SponsorId_example", "AspId_example"), false) // ChargeableParty | representation of the Chargeable Party resource to be Created in the SCEF

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ChargeablePartyTransactionOperationApi.CreateChargeablePartyTransaction(context.Background(), scsAsId).ChargeableParty(chargeableParty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ChargeablePartyTransactionOperationApi.CreateChargeablePartyTransaction``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateChargeablePartyTransaction`: ChargeableParty
    fmt.Fprintf(os.Stdout, "Response from `ChargeablePartyTransactionOperationApi.CreateChargeablePartyTransaction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**scsAsId** | **string** | Identifier of SCS/AS | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateChargeablePartyTransactionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **chargeableParty** | [**ChargeableParty**](ChargeableParty.md) | representation of the Chargeable Party resource to be Created in the SCEF | 

### Return type

[**ChargeableParty**](ChargeableParty.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchAllChargeablePartyTransactions

> []ChargeableParty FetchAllChargeablePartyTransactions(ctx, scsAsId).IpAddrs(ipAddrs).IpDomain(ipDomain).MacAddrs(macAddrs).Execute()

Read all or queried chargeable party transaction resources for a given SCS/AS.

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
    ipAddrs := []openapiclient.IpAddr{openapiclient.IpAddr{Interface{}: new(interface{})}} // []IpAddr | The IP address(es) of the requested UE(s). (optional)
    ipDomain := "ipDomain_example" // string | The IPv4 address domain identifier. The attribute may only be provided if IPv4 address is included in the ip-addrs query parameter. (optional)
    macAddrs := []string{"Inner_example"} // []string | The MAC address(es) of the requested UE(s). (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ChargeablePartyTransactionOperationApi.FetchAllChargeablePartyTransactions(context.Background(), scsAsId).IpAddrs(ipAddrs).IpDomain(ipDomain).MacAddrs(macAddrs).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ChargeablePartyTransactionOperationApi.FetchAllChargeablePartyTransactions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchAllChargeablePartyTransactions`: []ChargeableParty
    fmt.Fprintf(os.Stdout, "Response from `ChargeablePartyTransactionOperationApi.FetchAllChargeablePartyTransactions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**scsAsId** | **string** | Identifier of SCS/AS | 

### Other Parameters

Other parameters are passed through a pointer to a apiFetchAllChargeablePartyTransactionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ipAddrs** | [**[]IpAddr**](IpAddr.md) | The IP address(es) of the requested UE(s). | 
 **ipDomain** | **string** | The IPv4 address domain identifier. The attribute may only be provided if IPv4 address is included in the ip-addrs query parameter. | 
 **macAddrs** | **[]string** | The MAC address(es) of the requested UE(s). | 

### Return type

[**[]ChargeableParty**](ChargeableParty.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

