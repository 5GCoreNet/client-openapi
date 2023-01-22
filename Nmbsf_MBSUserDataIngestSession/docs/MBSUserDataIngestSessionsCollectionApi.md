# \MBSUserDataIngestSessionsCollectionApi

All URIs are relative to *https://example.com/nmbsf-mbs-ud-ingest/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateMBSUserDataIngSession**](MBSUserDataIngestSessionsCollectionApi.md#CreateMBSUserDataIngSession) | **Post** /sessions | Request the creation of a new MBS User Data Ingest Session.
[**RetrieveMBSUserDataIngSessions**](MBSUserDataIngestSessionsCollectionApi.md#RetrieveMBSUserDataIngSessions) | **Get** /sessions | Retrieve all the active MBS User Data Ingest Sessions managed by the MBSF.



## CreateMBSUserDataIngSession

> MBSUserDataIngSession CreateMBSUserDataIngSession(ctx).MBSUserDataIngSession(mBSUserDataIngSession).Execute()

Request the creation of a new MBS User Data Ingest Session.

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
    mBSUserDataIngSession := *openapiclient.NewMBSUserDataIngSession("MbsUserServId_example", map[string]MBSDistributionSessionInfo{"key": *openapiclient.NewMBSDistributionSessionInfo("MaxContBitRate_example", openapiclient.DistributionMethod{DistributionMethodOneOf: penapiclient.DistributionMethod_oneOf("OBJECT")})}) // MBSUserDataIngSession | Contains the parameters to request the creation of a new MBS User Data Ingest Session  at the MBSF. 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MBSUserDataIngestSessionsCollectionApi.CreateMBSUserDataIngSession(context.Background()).MBSUserDataIngSession(mBSUserDataIngSession).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MBSUserDataIngestSessionsCollectionApi.CreateMBSUserDataIngSession``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateMBSUserDataIngSession`: MBSUserDataIngSession
    fmt.Fprintf(os.Stdout, "Response from `MBSUserDataIngestSessionsCollectionApi.CreateMBSUserDataIngSession`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateMBSUserDataIngSessionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **mBSUserDataIngSession** | [**MBSUserDataIngSession**](MBSUserDataIngSession.md) | Contains the parameters to request the creation of a new MBS User Data Ingest Session  at the MBSF.  | 

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


## RetrieveMBSUserDataIngSessions

> []MBSUserDataIngSession RetrieveMBSUserDataIngSessions(ctx).Execute()

Retrieve all the active MBS User Data Ingest Sessions managed by the MBSF.

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
    resp, r, err := apiClient.MBSUserDataIngestSessionsCollectionApi.RetrieveMBSUserDataIngSessions(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MBSUserDataIngestSessionsCollectionApi.RetrieveMBSUserDataIngSessions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveMBSUserDataIngSessions`: []MBSUserDataIngSession
    fmt.Fprintf(os.Stdout, "Response from `MBSUserDataIngestSessionsCollectionApi.RetrieveMBSUserDataIngSessions`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveMBSUserDataIngSessionsRequest struct via the builder pattern


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

