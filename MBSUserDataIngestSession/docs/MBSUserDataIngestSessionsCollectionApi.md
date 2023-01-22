# \MBSUserDataIngestSessionsCollectionApi

All URIs are relative to *https://example.com/3gpp-mbs-ud-ingest/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateMBSUserDataIngestSession**](MBSUserDataIngestSessionsCollectionApi.md#CreateMBSUserDataIngestSession) | **Post** /sessions | Request the creation of a new Individual MBS User Data Ingest Session resource.
[**RetrieveMBSUserDataIngestSessions**](MBSUserDataIngestSessionsCollectionApi.md#RetrieveMBSUserDataIngestSessions) | **Get** /sessions | Retrieve all the active MBS User Data Ingest Sessions managed by the NEF.



## CreateMBSUserDataIngestSession

> MBSUserDataIngSession CreateMBSUserDataIngestSession(ctx).MBSUserDataIngSession(mBSUserDataIngSession).Execute()

Request the creation of a new Individual MBS User Data Ingest Session resource.

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
    mBSUserDataIngSession := *openapiclient.NewMBSUserDataIngSession("MbsUserServId_example", map[string]MBSDistributionSessionInfo{"key": *openapiclient.NewMBSDistributionSessionInfo("MaxContBitRate_example", openapiclient.DistributionMethod{DistributionMethodOneOf: penapiclient.DistributionMethod_oneOf("OBJECT")})}) // MBSUserDataIngSession | Contains the parameters to request the creation of a new MBS User Data Ingest Session  at the NEF. 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MBSUserDataIngestSessionsCollectionApi.CreateMBSUserDataIngestSession(context.Background()).MBSUserDataIngSession(mBSUserDataIngSession).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MBSUserDataIngestSessionsCollectionApi.CreateMBSUserDataIngestSession``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateMBSUserDataIngestSession`: MBSUserDataIngSession
    fmt.Fprintf(os.Stdout, "Response from `MBSUserDataIngestSessionsCollectionApi.CreateMBSUserDataIngestSession`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateMBSUserDataIngestSessionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **mBSUserDataIngSession** | [**MBSUserDataIngSession**](MBSUserDataIngSession.md) | Contains the parameters to request the creation of a new MBS User Data Ingest Session  at the NEF.  | 

### Return type

[**MBSUserDataIngSession**](MBSUserDataIngSession.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveMBSUserDataIngestSessions

> []MBSUserDataIngSession RetrieveMBSUserDataIngestSessions(ctx).Execute()

Retrieve all the active MBS User Data Ingest Sessions managed by the NEF.

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MBSUserDataIngestSessionsCollectionApi.RetrieveMBSUserDataIngestSessions(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MBSUserDataIngestSessionsCollectionApi.RetrieveMBSUserDataIngestSessions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveMBSUserDataIngestSessions`: []MBSUserDataIngSession
    fmt.Fprintf(os.Stdout, "Response from `MBSUserDataIngestSessionsCollectionApi.RetrieveMBSUserDataIngestSessions`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveMBSUserDataIngestSessionsRequest struct via the builder pattern


### Return type

[**[]MBSUserDataIngSession**](MBSUserDataIngSession.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

