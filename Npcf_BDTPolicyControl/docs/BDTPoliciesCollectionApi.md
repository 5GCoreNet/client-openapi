# \BDTPoliciesCollectionApi

All URIs are relative to *https://example.com/npcf-bdtpolicycontrol/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateBDTPolicy**](BDTPoliciesCollectionApi.md#CreateBDTPolicy) | **Post** /bdtpolicies | Create a new Individual BDT policy



## CreateBDTPolicy

> BdtPolicy CreateBDTPolicy(ctx).BdtReqData(bdtReqData).Execute()

Create a new Individual BDT policy

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    bdtReqData := *openapiclient.NewBdtReqData("AspId_example", *openapiclient.NewTimeWindow(time.Now(), time.Now()), int32(123), *openapiclient.NewUsageThreshold()) // BdtReqData | Contains information for the creation of a new Individual BDT policy resource. 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BDTPoliciesCollectionApi.CreateBDTPolicy(context.Background()).BdtReqData(bdtReqData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BDTPoliciesCollectionApi.CreateBDTPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateBDTPolicy`: BdtPolicy
    fmt.Fprintf(os.Stdout, "Response from `BDTPoliciesCollectionApi.CreateBDTPolicy`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateBDTPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bdtReqData** | [**BdtReqData**](BdtReqData.md) | Contains information for the creation of a new Individual BDT policy resource.  | 

### Return type

[**BdtPolicy**](BdtPolicy.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

