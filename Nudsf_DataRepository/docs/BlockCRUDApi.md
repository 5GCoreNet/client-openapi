# \BlockCRUDApi

All URIs are relative to *https://example.com/nudsf-dr/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOrModifyBlock**](BlockCRUDApi.md#CreateOrModifyBlock) | **Put** /{realmId}/{storageId}/records/{recordId}/blocks/{blockId} | Create or Update a specific Block in a Record.
[**DeleteBlock**](BlockCRUDApi.md#DeleteBlock) | **Delete** /{realmId}/{storageId}/records/{recordId}/blocks/{blockId} | Delete a specific Block. Then update the Record
[**GetBlock**](BlockCRUDApi.md#GetBlock) | **Get** /{realmId}/{storageId}/records/{recordId}/blocks/{blockId} | Retrieve a specific Block
[**GetBlockList**](BlockCRUDApi.md#GetBlockList) | **Get** /{realmId}/{storageId}/records/{recordId}/blocks | Record&#39;s Blocks access



## CreateOrModifyBlock

> interface{} CreateOrModifyBlock(ctx, realmId, storageId, recordId, blockId).Body(body).GetPrevious(getPrevious).IfNoneMatch(ifNoneMatch).IfMatch(ifMatch).SupportedFeatures(supportedFeatures).Execute()

Create or Update a specific Block in a Record.



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
    blockId := "userDefjson01" // string | Id of the Block
    body := interface{}(987) // interface{} | information on the Block to create
    getPrevious := true // bool | Retrieve the Block before update (optional) (default to false)
    ifNoneMatch := "ifNoneMatch_example" // string | Validator for conditional requests, as described in RFC 7232, 3.2 (optional)
    ifMatch := "ifMatch_example" // string | Record validator for conditional requests, as described in RFC 7232, 3.2 (optional)
    supportedFeatures := "supportedFeatures_example" // string | Features required to be supported by the target NF (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BlockCRUDApi.CreateOrModifyBlock(context.Background(), realmId, storageId, recordId, blockId).Body(body).GetPrevious(getPrevious).IfNoneMatch(ifNoneMatch).IfMatch(ifMatch).SupportedFeatures(supportedFeatures).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BlockCRUDApi.CreateOrModifyBlock``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOrModifyBlock`: interface{}
    fmt.Fprintf(os.Stdout, "Response from `BlockCRUDApi.CreateOrModifyBlock`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realmId** | **string** | Identifier of the Realm | 
**storageId** | **string** | Identifier of the Storage | 
**recordId** | **string** | Identifier of the Record | 
**blockId** | **string** | Id of the Block | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrModifyBlockRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **body** | **interface{}** | information on the Block to create | 
 **getPrevious** | **bool** | Retrieve the Block before update | [default to false]
 **ifNoneMatch** | **string** | Validator for conditional requests, as described in RFC 7232, 3.2 | 
 **ifMatch** | **string** | Record validator for conditional requests, as described in RFC 7232, 3.2 | 
 **supportedFeatures** | **string** | Features required to be supported by the target NF | 

### Return type

**interface{}**

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteBlock

> interface{} DeleteBlock(ctx, realmId, storageId, recordId, blockId).GetPrevious(getPrevious).IfMatch(ifMatch).SupportedFeatures(supportedFeatures).Execute()

Delete a specific Block. Then update the Record



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
    blockId := "userDefjson01" // string | Id of the Block
    getPrevious := true // bool | Retrieve the Block before delete (optional) (default to false)
    ifMatch := "ifMatch_example" // string | Record validator for conditional requests, as described in RFC 7232, 3.2 (optional)
    supportedFeatures := "supportedFeatures_example" // string | Features required to be supported by the target NF (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BlockCRUDApi.DeleteBlock(context.Background(), realmId, storageId, recordId, blockId).GetPrevious(getPrevious).IfMatch(ifMatch).SupportedFeatures(supportedFeatures).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BlockCRUDApi.DeleteBlock``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteBlock`: interface{}
    fmt.Fprintf(os.Stdout, "Response from `BlockCRUDApi.DeleteBlock`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realmId** | **string** | Identifier of the Realm | 
**storageId** | **string** | Identifier of the Storage | 
**recordId** | **string** | Identifier of the Record | 
**blockId** | **string** | Id of the Block | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteBlockRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **getPrevious** | **bool** | Retrieve the Block before delete | [default to false]
 **ifMatch** | **string** | Record validator for conditional requests, as described in RFC 7232, 3.2 | 
 **supportedFeatures** | **string** | Features required to be supported by the target NF | 

### Return type

**interface{}**

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBlock

> interface{} GetBlock(ctx, realmId, storageId, recordId, blockId).IfNoneMatch(ifNoneMatch).IfModifiedSince(ifModifiedSince).SupportedFeatures(supportedFeatures).Execute()

Retrieve a specific Block



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
    blockId := "userDefjson01" // string | Id of the Block
    ifNoneMatch := "ifNoneMatch_example" // string | Validator for conditional requests, as described in RFC 7232, 3.2 (optional)
    ifModifiedSince := "ifModifiedSince_example" // string | Validator for conditional requests, as described in RFC 7232, 3.3 (optional)
    supportedFeatures := "supportedFeatures_example" // string | Features required to be supported by the target NF (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BlockCRUDApi.GetBlock(context.Background(), realmId, storageId, recordId, blockId).IfNoneMatch(ifNoneMatch).IfModifiedSince(ifModifiedSince).SupportedFeatures(supportedFeatures).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BlockCRUDApi.GetBlock``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBlock`: interface{}
    fmt.Fprintf(os.Stdout, "Response from `BlockCRUDApi.GetBlock`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realmId** | **string** | Identifier of the Realm | 
**storageId** | **string** | Identifier of the Storage | 
**recordId** | **string** | Identifier of the Record | 
**blockId** | **string** | Id of the Block | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBlockRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **ifNoneMatch** | **string** | Validator for conditional requests, as described in RFC 7232, 3.2 | 
 **ifModifiedSince** | **string** | Validator for conditional requests, as described in RFC 7232, 3.3 | 
 **supportedFeatures** | **string** | Features required to be supported by the target NF | 

### Return type

**interface{}**

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBlockList

> GetBlockList200Response GetBlockList(ctx, realmId, storageId, recordId).SupportedFeatures(supportedFeatures).Execute()

Record's Blocks access



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
    supportedFeatures := "supportedFeatures_example" // string | Features required to be supported by the target NF (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BlockCRUDApi.GetBlockList(context.Background(), realmId, storageId, recordId).SupportedFeatures(supportedFeatures).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BlockCRUDApi.GetBlockList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBlockList`: GetBlockList200Response
    fmt.Fprintf(os.Stdout, "Response from `BlockCRUDApi.GetBlockList`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiGetBlockListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **supportedFeatures** | **string** | Features required to be supported by the target NF | 

### Return type

[**GetBlockList200Response**](GetBlockList200Response.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: multipart/parallel, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

