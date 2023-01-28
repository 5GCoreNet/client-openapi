# \DefaultApi

All URIs are relative to *https://example.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**NiddConfigurationTrigger**](DefaultApi.md#NiddConfigurationTrigger) | **Post** / | 



## NiddConfigurationTrigger

> NiddConfigurationTriggerReply NiddConfigurationTrigger(ctx).NiddConfigurationTrigger(niddConfigurationTrigger).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/NIDDConfigurationTrigger"
)

func main() {
    niddConfigurationTrigger := *openapiclient.NewNiddConfigurationTrigger("AfId_example", "NefId_example", "Gpsi_example", "SuppFeat_example") // NiddConfigurationTrigger | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.NiddConfigurationTrigger(context.Background()).NiddConfigurationTrigger(niddConfigurationTrigger).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.NiddConfigurationTrigger``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `NiddConfigurationTrigger`: NiddConfigurationTriggerReply
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.NiddConfigurationTrigger`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiNiddConfigurationTriggerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **niddConfigurationTrigger** | [**NiddConfigurationTrigger**](NiddConfigurationTrigger.md) |  | 

### Return type

[**NiddConfigurationTriggerReply**](NiddConfigurationTriggerReply.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

