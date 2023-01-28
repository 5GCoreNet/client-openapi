# \IndividualMBSUserDataIngestSessionDocumentApi

All URIs are relative to *https://example.com/3gpp-mbs-ud-ingest/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteIndivMBSUserDataIngestSession**](IndividualMBSUserDataIngestSessionDocumentApi.md#DeleteIndivMBSUserDataIngestSession) | **Delete** /sessions/{sessionId} | Deletes an existing Individual MBS User Data Ingest Session resource.
[**ModifyIndivMBSUserDataIngestSession**](IndividualMBSUserDataIngestSessionDocumentApi.md#ModifyIndivMBSUserDataIngestSession) | **Patch** /sessions/{sessionId} | Request the modification of an existing Individual MBS User Data Ingest Session resource.
[**RetrieveIndivMBSUserDataIngestSession**](IndividualMBSUserDataIngestSessionDocumentApi.md#RetrieveIndivMBSUserDataIngestSession) | **Get** /sessions/{sessionId} | Retrieve an existing Individual MBS User Data Ingest Session resource.
[**UpdateIndivMBSUserDataIngestSession**](IndividualMBSUserDataIngestSessionDocumentApi.md#UpdateIndivMBSUserDataIngestSession) | **Put** /sessions/{sessionId} | Request the update of an existing Individual MBS User Data Ingest Session resource.



## DeleteIndivMBSUserDataIngestSession

> DeleteIndivMBSUserDataIngestSession(ctx, sessionId).Execute()

Deletes an existing Individual MBS User Data Ingest Session resource.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/MBSUserDataIngestSession"
)

func main() {
    sessionId := "sessionId_example" // string | Identifier of the Individual MBS User Data Ingest Session resource.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualMBSUserDataIngestSessionDocumentApi.DeleteIndivMBSUserDataIngestSession(context.Background(), sessionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualMBSUserDataIngestSessionDocumentApi.DeleteIndivMBSUserDataIngestSession``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sessionId** | **string** | Identifier of the Individual MBS User Data Ingest Session resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteIndivMBSUserDataIngestSessionRequest struct via the builder pattern


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


## ModifyIndivMBSUserDataIngestSession

> MBSUserDataIngSession ModifyIndivMBSUserDataIngestSession(ctx, sessionId).MBSUserDataIngSessionPatch(mBSUserDataIngSessionPatch).Execute()

Request the modification of an existing Individual MBS User Data Ingest Session resource.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/MBSUserDataIngestSession"
)

func main() {
    sessionId := "sessionId_example" // string | Identifier of the Individual MBS User Data Ingest Session resource.
    mBSUserDataIngSessionPatch := *openapiclient.NewMBSUserDataIngSessionPatch() // MBSUserDataIngSessionPatch | Contains the parameters to request the modification of the Individual MBS User Data Ingest  Session resource. 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualMBSUserDataIngestSessionDocumentApi.ModifyIndivMBSUserDataIngestSession(context.Background(), sessionId).MBSUserDataIngSessionPatch(mBSUserDataIngSessionPatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualMBSUserDataIngestSessionDocumentApi.ModifyIndivMBSUserDataIngestSession``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ModifyIndivMBSUserDataIngestSession`: MBSUserDataIngSession
    fmt.Fprintf(os.Stdout, "Response from `IndividualMBSUserDataIngestSessionDocumentApi.ModifyIndivMBSUserDataIngestSession`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sessionId** | **string** | Identifier of the Individual MBS User Data Ingest Session resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyIndivMBSUserDataIngestSessionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **mBSUserDataIngSessionPatch** | [**MBSUserDataIngSessionPatch**](MBSUserDataIngSessionPatch.md) | Contains the parameters to request the modification of the Individual MBS User Data Ingest  Session resource.  | 

### Return type

[**MBSUserDataIngSession**](MBSUserDataIngSession.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/merge-patch+json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveIndivMBSUserDataIngestSession

> MBSUserDataIngSession RetrieveIndivMBSUserDataIngestSession(ctx, sessionId).Execute()

Retrieve an existing Individual MBS User Data Ingest Session resource.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/MBSUserDataIngestSession"
)

func main() {
    sessionId := "sessionId_example" // string | Identifier of the Individual MBS User Data Ingest Session resource.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualMBSUserDataIngestSessionDocumentApi.RetrieveIndivMBSUserDataIngestSession(context.Background(), sessionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualMBSUserDataIngestSessionDocumentApi.RetrieveIndivMBSUserDataIngestSession``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveIndivMBSUserDataIngestSession`: MBSUserDataIngSession
    fmt.Fprintf(os.Stdout, "Response from `IndividualMBSUserDataIngestSessionDocumentApi.RetrieveIndivMBSUserDataIngestSession`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sessionId** | **string** | Identifier of the Individual MBS User Data Ingest Session resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveIndivMBSUserDataIngestSessionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MBSUserDataIngSession**](MBSUserDataIngSession.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateIndivMBSUserDataIngestSession

> MBSUserDataIngSession UpdateIndivMBSUserDataIngestSession(ctx, sessionId).MBSUserDataIngSession(mBSUserDataIngSession).Execute()

Request the update of an existing Individual MBS User Data Ingest Session resource.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/MBSUserDataIngestSession"
)

func main() {
    sessionId := "sessionId_example" // string | Identifier of the Individual MBS User Data Ingest Session resource.
    mBSUserDataIngSession := *openapiclient.NewMBSUserDataIngSession("MbsUserServId_example", map[string]MBSDistributionSessionInfo{"key": *openapiclient.NewMBSDistributionSessionInfo("MaxContBitRate_example", openapiclient.DistributionMethod{DistributionMethodOneOf: penapiclient.DistributionMethod_oneOf("OBJECT")})}) // MBSUserDataIngSession | Contains the updated representation of the Individual MBS User Data Ingest Session  resource. 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualMBSUserDataIngestSessionDocumentApi.UpdateIndivMBSUserDataIngestSession(context.Background(), sessionId).MBSUserDataIngSession(mBSUserDataIngSession).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualMBSUserDataIngestSessionDocumentApi.UpdateIndivMBSUserDataIngestSession``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateIndivMBSUserDataIngestSession`: MBSUserDataIngSession
    fmt.Fprintf(os.Stdout, "Response from `IndividualMBSUserDataIngestSessionDocumentApi.UpdateIndivMBSUserDataIngestSession`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sessionId** | **string** | Identifier of the Individual MBS User Data Ingest Session resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateIndivMBSUserDataIngestSessionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **mBSUserDataIngSession** | [**MBSUserDataIngSession**](MBSUserDataIngSession.md) | Contains the updated representation of the Individual MBS User Data Ingest Session  resource.  | 

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

