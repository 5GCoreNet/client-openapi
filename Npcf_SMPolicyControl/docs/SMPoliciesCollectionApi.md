# \SMPoliciesCollectionApi

All URIs are relative to *https://example.com/npcf-smpolicycontrol/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSMPolicy**](SMPoliciesCollectionApi.md#CreateSMPolicy) | **Post** /sm-policies | Create a new Individual SM Policy



## CreateSMPolicy

> SmPolicyDecision CreateSMPolicy(ctx).SmPolicyContextData(smPolicyContextData).Execute()

Create a new Individual SM Policy

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/Npcf_SMPolicyControl"
)

func main() {
    smPolicyContextData := *openapiclient.NewSmPolicyContextData("Supi_example", int32(123), *openapiclient.NewPduSessionType(), "Dnn_example", "NotificationUri_example", *openapiclient.NewSnssai(int32(123))) // SmPolicyContextData | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SMPoliciesCollectionApi.CreateSMPolicy(context.Background()).SmPolicyContextData(smPolicyContextData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SMPoliciesCollectionApi.CreateSMPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSMPolicy`: SmPolicyDecision
    fmt.Fprintf(os.Stdout, "Response from `SMPoliciesCollectionApi.CreateSMPolicy`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSMPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **smPolicyContextData** | [**SmPolicyContextData**](SmPolicyContextData.md) |  | 

### Return type

[**SmPolicyDecision**](SmPolicyDecision.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

