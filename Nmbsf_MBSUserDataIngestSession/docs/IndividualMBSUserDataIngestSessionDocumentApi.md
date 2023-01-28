# \IndividualMBSUserDataIngestSessionDocumentApi

All URIs are relative to *https://example.com/nmbsf-mbs-ud-ingest/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteIndMBSUserDataIngSession**](IndividualMBSUserDataIngestSessionDocumentApi.md#DeleteIndMBSUserDataIngSession) | **Delete** /sessions/{sessionId} | Request the deletion of an existing Individual MBS User Data Ingest Session resource.
[**ModifyIndMBSUserDataIngSession**](IndividualMBSUserDataIngestSessionDocumentApi.md#ModifyIndMBSUserDataIngSession) | **Patch** /sessions/{sessionId} | Request the modification of an existing Individual MBS User Data Ingest Session resource.
[**RetrieveIndMBSUserDataIngSession**](IndividualMBSUserDataIngestSessionDocumentApi.md#RetrieveIndMBSUserDataIngSession) | **Get** /sessions/{sessionId} | Retrieve an existing Individual MBS User Data Ingest Session resource.
[**UpdateIndMBSUserDataIngSession**](IndividualMBSUserDataIngestSessionDocumentApi.md#UpdateIndMBSUserDataIngSession) | **Put** /sessions/{sessionId} | Request the update of an existing Individual MBS User Data Ingest Session resource.



## DeleteIndMBSUserDataIngSession

> DeleteIndMBSUserDataIngSession(ctx, sessionId).Execute()

Request the deletion of an existing Individual MBS User Data Ingest Session resource.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/Nmbsf_MBSUserDataIngestSession"
)

func main() {
    sessionId := "sessionId_example" // string | Identifier of the Individual MBS User Data Ingest Session resource.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualMBSUserDataIngestSessionDocumentApi.DeleteIndMBSUserDataIngSession(context.Background(), sessionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualMBSUserDataIngestSessionDocumentApi.DeleteIndMBSUserDataIngSession``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteIndMBSUserDataIngSessionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyIndMBSUserDataIngSession

> MBSUserDataIngSession ModifyIndMBSUserDataIngSession(ctx, sessionId).MBSUserDataIngSessionPatch(mBSUserDataIngSessionPatch).Execute()

Request the modification of an existing Individual MBS User Data Ingest Session resource.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/Nmbsf_MBSUserDataIngestSession"
)

func main() {
    sessionId := "sessionId_example" // string | Identifier of the Individual MBS User Data Ingest Session resource.
    mBSUserDataIngSessionPatch := *openapiclient.NewMBSUserDataIngSessionPatch() // MBSUserDataIngSessionPatch | Contains the parameters to request the modification of the Individual MBS User Data Ingest  Session resource. 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualMBSUserDataIngestSessionDocumentApi.ModifyIndMBSUserDataIngSession(context.Background(), sessionId).MBSUserDataIngSessionPatch(mBSUserDataIngSessionPatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualMBSUserDataIngestSessionDocumentApi.ModifyIndMBSUserDataIngSession``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ModifyIndMBSUserDataIngSession`: MBSUserDataIngSession
    fmt.Fprintf(os.Stdout, "Response from `IndividualMBSUserDataIngestSessionDocumentApi.ModifyIndMBSUserDataIngSession`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sessionId** | **string** | Identifier of the Individual MBS User Data Ingest Session resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyIndMBSUserDataIngSessionRequest struct via the builder pattern


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


## RetrieveIndMBSUserDataIngSession

> MBSUserDataIngSession RetrieveIndMBSUserDataIngSession(ctx, sessionId).Execute()

Retrieve an existing Individual MBS User Data Ingest Session resource.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/Nmbsf_MBSUserDataIngestSession"
)

func main() {
    sessionId := "sessionId_example" // string | Identifier of the Individual MBS User Data Ingest Session resource.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualMBSUserDataIngestSessionDocumentApi.RetrieveIndMBSUserDataIngSession(context.Background(), sessionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualMBSUserDataIngestSessionDocumentApi.RetrieveIndMBSUserDataIngSession``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveIndMBSUserDataIngSession`: MBSUserDataIngSession
    fmt.Fprintf(os.Stdout, "Response from `IndividualMBSUserDataIngestSessionDocumentApi.RetrieveIndMBSUserDataIngSession`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sessionId** | **string** | Identifier of the Individual MBS User Data Ingest Session resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveIndMBSUserDataIngSessionRequest struct via the builder pattern


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


## UpdateIndMBSUserDataIngSession

> MBSUserDataIngSession UpdateIndMBSUserDataIngSession(ctx, sessionId).MBSUserDataIngSession(mBSUserDataIngSession).Execute()

Request the update of an existing Individual MBS User Data Ingest Session resource.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/Nmbsf_MBSUserDataIngestSession"
)

func main() {
    sessionId := "sessionId_example" // string | Identifier of the Individual MBS User Data Ingest Session resource.
    mBSUserDataIngSession := *openapiclient.NewMBSUserDataIngSession("MbsUserServId_example", map[string]MBSDistributionSessionInfo{"key": *openapiclient.NewMBSDistributionSessionInfo("MaxContBitRate_example", openapiclient.DistributionMethod{DistributionMethodOneOf: penapiclient.DistributionMethod_oneOf("OBJECT")})}) // MBSUserDataIngSession | Contains the updated representation of the Individual MBS User Data Ingest Session  resource. 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IndividualMBSUserDataIngestSessionDocumentApi.UpdateIndMBSUserDataIngSession(context.Background(), sessionId).MBSUserDataIngSession(mBSUserDataIngSession).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IndividualMBSUserDataIngestSessionDocumentApi.UpdateIndMBSUserDataIngSession``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateIndMBSUserDataIngSession`: MBSUserDataIngSession
    fmt.Fprintf(os.Stdout, "Response from `IndividualMBSUserDataIngestSessionDocumentApi.UpdateIndMBSUserDataIngSession`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sessionId** | **string** | Identifier of the Individual MBS User Data Ingest Session resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateIndMBSUserDataIngSessionRequest struct via the builder pattern


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

