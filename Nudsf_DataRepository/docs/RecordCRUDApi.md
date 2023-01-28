# \RecordCRUDApi

All URIs are relative to *https://example.com/nudsf-dr/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BulkDeleteRecords**](RecordCRUDApi.md#BulkDeleteRecords) | **Delete** /{realmId}/{storageId}/records | Bulk Deletion of Records
[**CreateOrModifyRecord**](RecordCRUDApi.md#CreateOrModifyRecord) | **Put** /{realmId}/{storageId}/records/{recordId} | Create/Modify Record
[**DeleteRecord**](RecordCRUDApi.md#DeleteRecord) | **Delete** /{realmId}/{storageId}/records/{recordId} | Delete a Record with an user provided RecordId
[**GetMeta**](RecordCRUDApi.md#GetMeta) | **Get** /{realmId}/{storageId}/records/{recordId}/meta | Record&#39;s meta access
[**GetRecord**](RecordCRUDApi.md#GetRecord) | **Get** /{realmId}/{storageId}/records/{recordId} | Record access
[**SearchRecord**](RecordCRUDApi.md#SearchRecord) | **Get** /{realmId}/{storageId}/records | Records search with get
[**UpdateMeta**](RecordCRUDApi.md#UpdateMeta) | **Patch** /{realmId}/{storageId}/records/{recordId}/meta | Record&#39;s meta update



## BulkDeleteRecords

> RecordIdList BulkDeleteRecords(ctx, realmId, storageId).Filter(filter).SupportedFeatures(supportedFeatures).Execute()

Bulk Deletion of Records



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/Nudsf_DataRepository"
)

func main() {
    realmId := "Realm01" // string | Identifier of the Realm
    storageId := "Storage01" // string | Identifier of the Storage
    filter := map[string][]openapiclient.SearchExpression{ ... } // SearchExpression | 
    supportedFeatures := "supportedFeatures_example" // string | Features required to be supported by the target NF (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RecordCRUDApi.BulkDeleteRecords(context.Background(), realmId, storageId).Filter(filter).SupportedFeatures(supportedFeatures).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RecordCRUDApi.BulkDeleteRecords``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BulkDeleteRecords`: RecordIdList
    fmt.Fprintf(os.Stdout, "Response from `RecordCRUDApi.BulkDeleteRecords`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realmId** | **string** | Identifier of the Realm | 
**storageId** | **string** | Identifier of the Storage | 

### Other Parameters

Other parameters are passed through a pointer to a apiBulkDeleteRecordsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **filter** | [**SearchExpression**](SearchExpression.md) |  | 
 **supportedFeatures** | **string** | Features required to be supported by the target NF | 

### Return type

[**RecordIdList**](RecordIdList.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateOrModifyRecord

> Record CreateOrModifyRecord(ctx, realmId, storageId, recordId).Record(record).IfNoneMatch(ifNoneMatch).IfMatch(ifMatch).GetPrevious(getPrevious).SupportedFeatures(supportedFeatures).Execute()

Create/Modify Record



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/Nudsf_DataRepository"
)

func main() {
    realmId := "Realm01" // string | Identifier(name) of the Realm
    storageId := "Storage01" // string | Identifier of the Storage
    recordId := "UserRecordValue000000001" // string | Identifier of the Record
    record := *openapiclient.NewRecord(*openapiclient.NewRecordMeta()) // Record | The record multipart request body. The meta part shall be the first part and is mandatory but can be empty and zero or more block parts may follow the meta part.
    ifNoneMatch := "ifNoneMatch_example" // string | Validator for conditional requests, as described in RFC 7232, 3.2 (optional)
    ifMatch := "ifMatch_example" // string | Record validator for conditional requests, as described in RFC 7232, 3.2 (optional)
    getPrevious := true // bool | Retrieve the Record before update (optional) (default to false)
    supportedFeatures := "supportedFeatures_example" // string | Features required to be supported by the target NF (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RecordCRUDApi.CreateOrModifyRecord(context.Background(), realmId, storageId, recordId).Record(record).IfNoneMatch(ifNoneMatch).IfMatch(ifMatch).GetPrevious(getPrevious).SupportedFeatures(supportedFeatures).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RecordCRUDApi.CreateOrModifyRecord``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOrModifyRecord`: Record
    fmt.Fprintf(os.Stdout, "Response from `RecordCRUDApi.CreateOrModifyRecord`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realmId** | **string** | Identifier(name) of the Realm | 
**storageId** | **string** | Identifier of the Storage | 
**recordId** | **string** | Identifier of the Record | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrModifyRecordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **record** | [**Record**](Record.md) | The record multipart request body. The meta part shall be the first part and is mandatory but can be empty and zero or more block parts may follow the meta part. | 
 **ifNoneMatch** | **string** | Validator for conditional requests, as described in RFC 7232, 3.2 | 
 **ifMatch** | **string** | Record validator for conditional requests, as described in RFC 7232, 3.2 | 
 **getPrevious** | **bool** | Retrieve the Record before update | [default to false]
 **supportedFeatures** | **string** | Features required to be supported by the target NF | 

### Return type

[**Record**](Record.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: multipart/mixed
- **Accept**: multipart/mixed, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteRecord

> Record DeleteRecord(ctx, realmId, storageId, recordId).IfMatch(ifMatch).GetPrevious(getPrevious).SupportedFeatures(supportedFeatures).Execute()

Delete a Record with an user provided RecordId

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/Nudsf_DataRepository"
)

func main() {
    realmId := "Realm01" // string | Identifier(name) of the Realm
    storageId := "Storage01" // string | Identifier of the Storage
    recordId := "UserRecordValue000000001" // string | Identifier of the Record
    ifMatch := "ifMatch_example" // string | Record validator for conditional requests, as described in RFC 7232, 3.2 (optional)
    getPrevious := true // bool | Retrieve the Record before delete (optional) (default to false)
    supportedFeatures := "supportedFeatures_example" // string | Features required to be supported by the target NF (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RecordCRUDApi.DeleteRecord(context.Background(), realmId, storageId, recordId).IfMatch(ifMatch).GetPrevious(getPrevious).SupportedFeatures(supportedFeatures).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RecordCRUDApi.DeleteRecord``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteRecord`: Record
    fmt.Fprintf(os.Stdout, "Response from `RecordCRUDApi.DeleteRecord`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realmId** | **string** | Identifier(name) of the Realm | 
**storageId** | **string** | Identifier of the Storage | 
**recordId** | **string** | Identifier of the Record | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRecordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **ifMatch** | **string** | Record validator for conditional requests, as described in RFC 7232, 3.2 | 
 **getPrevious** | **bool** | Retrieve the Record before delete | [default to false]
 **supportedFeatures** | **string** | Features required to be supported by the target NF | 

### Return type

[**Record**](Record.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: multipart/mixed, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMeta

> RecordMeta GetMeta(ctx, realmId, storageId, recordId).IfNoneMatch(ifNoneMatch).IfModifiedSince(ifModifiedSince).SupportedFeatures(supportedFeatures).Execute()

Record's meta access



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/Nudsf_DataRepository"
)

func main() {
    realmId := "Realm01" // string | Identifier of the Realm
    storageId := "Storage01" // string | Identifier of the Storage
    recordId := "UserRecordValue000000001" // string | Identifier of the Record
    ifNoneMatch := "ifNoneMatch_example" // string | Validator for conditional requests, as described in RFC 7232, 3.2 (optional)
    ifModifiedSince := "ifModifiedSince_example" // string | Validator for conditional requests, as described in RFC 7232, 3.3 (optional)
    supportedFeatures := "supportedFeatures_example" // string | Features required to be supported by the target NF (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RecordCRUDApi.GetMeta(context.Background(), realmId, storageId, recordId).IfNoneMatch(ifNoneMatch).IfModifiedSince(ifModifiedSince).SupportedFeatures(supportedFeatures).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RecordCRUDApi.GetMeta``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMeta`: RecordMeta
    fmt.Fprintf(os.Stdout, "Response from `RecordCRUDApi.GetMeta`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realmId** | **string** | Identifier of the Realm | 
**storageId** | **string** | Identifier of the Storage | 
**recordId** | **string** | Identifier of the Record | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMetaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **ifNoneMatch** | **string** | Validator for conditional requests, as described in RFC 7232, 3.2 | 
 **ifModifiedSince** | **string** | Validator for conditional requests, as described in RFC 7232, 3.3 | 
 **supportedFeatures** | **string** | Features required to be supported by the target NF | 

### Return type

[**RecordMeta**](RecordMeta.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRecord

> Record GetRecord(ctx, realmId, storageId, recordId).IfNoneMatch(ifNoneMatch).IfModifiedSince(ifModifiedSince).SupportedFeatures(supportedFeatures).Execute()

Record access



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/Nudsf_DataRepository"
)

func main() {
    realmId := "Realm01" // string | Identifier of the Realm
    storageId := "Storage01" // string | Identifier of the Storage
    recordId := "UserRecordValue000000001" // string | Identifier of the Record
    ifNoneMatch := "ifNoneMatch_example" // string | Validator for conditional requests, as described in RFC 7232, 3.2 (optional)
    ifModifiedSince := "ifModifiedSince_example" // string | Validator for conditional requests, as described in RFC 7232, 3.3 (optional)
    supportedFeatures := "supportedFeatures_example" // string | Features required to be supported by the target NF (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RecordCRUDApi.GetRecord(context.Background(), realmId, storageId, recordId).IfNoneMatch(ifNoneMatch).IfModifiedSince(ifModifiedSince).SupportedFeatures(supportedFeatures).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RecordCRUDApi.GetRecord``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRecord`: Record
    fmt.Fprintf(os.Stdout, "Response from `RecordCRUDApi.GetRecord`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realmId** | **string** | Identifier of the Realm | 
**storageId** | **string** | Identifier of the Storage | 
**recordId** | **string** | Identifier of the Record | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRecordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **ifNoneMatch** | **string** | Validator for conditional requests, as described in RFC 7232, 3.2 | 
 **ifModifiedSince** | **string** | Validator for conditional requests, as described in RFC 7232, 3.3 | 
 **supportedFeatures** | **string** | Features required to be supported by the target NF | 

### Return type

[**Record**](Record.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: multipart/mixed, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchRecord

> RecordSearchResult SearchRecord(ctx, realmId, storageId).LimitRange(limitRange).Filter(filter).CountIndicator(countIndicator).SupportedFeatures(supportedFeatures).RetrieveRecords(retrieveRecords).MaxPayloadSize(maxPayloadSize).Execute()

Records search with get



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/Nudsf_DataRepository"
)

func main() {
    realmId := "Realm01" // string | Identifier of the Realm
    storageId := "Storage01" // string | Identifier of the Storage
    limitRange := int32(56) // int32 | The most number of record references to fetch (optional)
    filter := map[string][]openapiclient.SearchExpression{ ... } // SearchExpression | Query filter using conditions on tags (optional)
    countIndicator := true // bool | Indicates whether the number of records that matched the criteria shall be returned. (optional) (default to false)
    supportedFeatures := "supportedFeatures_example" // string | Features required to be supported by the target NF (optional)
    retrieveRecords := *openapiclient.NewRetrieveRecords() // RetrieveRecords | Indicates whether the UDSF is requested to include matching records within the response. (optional)
    maxPayloadSize := int32(56) // int32 | Indicates the number of kilo octets the consumer is prepared to receive (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RecordCRUDApi.SearchRecord(context.Background(), realmId, storageId).LimitRange(limitRange).Filter(filter).CountIndicator(countIndicator).SupportedFeatures(supportedFeatures).RetrieveRecords(retrieveRecords).MaxPayloadSize(maxPayloadSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RecordCRUDApi.SearchRecord``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchRecord`: RecordSearchResult
    fmt.Fprintf(os.Stdout, "Response from `RecordCRUDApi.SearchRecord`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realmId** | **string** | Identifier of the Realm | 
**storageId** | **string** | Identifier of the Storage | 

### Other Parameters

Other parameters are passed through a pointer to a apiSearchRecordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **limitRange** | **int32** | The most number of record references to fetch | 
 **filter** | [**SearchExpression**](SearchExpression.md) | Query filter using conditions on tags | 
 **countIndicator** | **bool** | Indicates whether the number of records that matched the criteria shall be returned. | [default to false]
 **supportedFeatures** | **string** | Features required to be supported by the target NF | 
 **retrieveRecords** | [**RetrieveRecords**](RetrieveRecords.md) | Indicates whether the UDSF is requested to include matching records within the response. | 
 **maxPayloadSize** | **int32** | Indicates the number of kilo octets the consumer is prepared to receive | 

### Return type

[**RecordSearchResult**](RecordSearchResult.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateMeta

> PatchResult UpdateMeta(ctx, realmId, storageId, recordId).PatchItem(patchItem).IfMatch(ifMatch).SupportedFeatures(supportedFeatures).Execute()

Record's meta update



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/Nudsf_DataRepository"
)

func main() {
    realmId := "Realm01" // string | Identifier of the Realm
    storageId := "Storage01" // string | Identifier of the Storage
    recordId := "UserRecordValue000000001" // string | Identifier of the Record
    patchItem := []openapiclient.PatchItem{*openapiclient.NewPatchItem(*openapiclient.NewPatchOperation(), "Path_example")} // []PatchItem | Meta data to patch
    ifMatch := "ifMatch_example" // string | Record validator for conditional requests, as described in RFC 7232, 3.2 (optional)
    supportedFeatures := "supportedFeatures_example" // string | Features required to be supported by the target NF (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RecordCRUDApi.UpdateMeta(context.Background(), realmId, storageId, recordId).PatchItem(patchItem).IfMatch(ifMatch).SupportedFeatures(supportedFeatures).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RecordCRUDApi.UpdateMeta``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateMeta`: PatchResult
    fmt.Fprintf(os.Stdout, "Response from `RecordCRUDApi.UpdateMeta`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realmId** | **string** | Identifier of the Realm | 
**storageId** | **string** | Identifier of the Storage | 
**recordId** | **string** | Identifier of the Record | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateMetaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **patchItem** | [**[]PatchItem**](PatchItem.md) | Meta data to patch | 
 **ifMatch** | **string** | Record validator for conditional requests, as described in RFC 7232, 3.2 | 
 **supportedFeatures** | **string** | Features required to be supported by the target NF | 

### Return type

[**PatchResult**](PatchResult.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json-patch+json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

