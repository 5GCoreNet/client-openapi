# \MetaSchemaCRUDApi

All URIs are relative to *https://example.com/nudsf-dr/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOrModifyMetaSchema**](MetaSchemaCRUDApi.md#CreateOrModifyMetaSchema) | **Put** /{realmId}/{storageId}/meta-schemas/{schemaId} | Create/Modify Meta Schema
[**DeleteMetaSchema**](MetaSchemaCRUDApi.md#DeleteMetaSchema) | **Delete** /{realmId}/{storageId}/meta-schemas/{schemaId} | Delete a Meta Schema with an user provided SchemaId
[**GetMetaSchema**](MetaSchemaCRUDApi.md#GetMetaSchema) | **Get** /{realmId}/{storageId}/meta-schemas/{schemaId} | Meta Schema access



## CreateOrModifyMetaSchema

> MetaSchema CreateOrModifyMetaSchema(ctx, realmId, storageId, schemaId).MetaSchema(metaSchema).IfNoneMatch(ifNoneMatch).IfMatch(ifMatch).GetPrevious(getPrevious).SupportedFeatures(supportedFeatures).Execute()

Create/Modify Meta Schema



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
    realmId := "Realm01" // string | Identifier(name) of the Realm
    storageId := "Storage01" // string | Identifier of the Storage
    schemaId := "schemaId_example" // string | Identifier of the Meta Schema
    metaSchema := *openapiclient.NewMetaSchema("SchemaId_example", []openapiclient.TagType{*openapiclient.NewTagType("TagName_example", *openapiclient.NewKeyType())}) // MetaSchema | 
    ifNoneMatch := "ifNoneMatch_example" // string | Validator for conditional requests, as described in RFC 7232, 3.2 (optional)
    ifMatch := "ifMatch_example" // string | Validator for conditional requests, as described in RFC 7232, 3.2 (optional)
    getPrevious := true // bool | Retrieve the Meta Schema before update (optional) (default to false)
    supportedFeatures := "supportedFeatures_example" // string | Features required to be supported by the target NF (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MetaSchemaCRUDApi.CreateOrModifyMetaSchema(context.Background(), realmId, storageId, schemaId).MetaSchema(metaSchema).IfNoneMatch(ifNoneMatch).IfMatch(ifMatch).GetPrevious(getPrevious).SupportedFeatures(supportedFeatures).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetaSchemaCRUDApi.CreateOrModifyMetaSchema``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOrModifyMetaSchema`: MetaSchema
    fmt.Fprintf(os.Stdout, "Response from `MetaSchemaCRUDApi.CreateOrModifyMetaSchema`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realmId** | **string** | Identifier(name) of the Realm | 
**storageId** | **string** | Identifier of the Storage | 
**schemaId** | **string** | Identifier of the Meta Schema | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrModifyMetaSchemaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **metaSchema** | [**MetaSchema**](MetaSchema.md) |  | 
 **ifNoneMatch** | **string** | Validator for conditional requests, as described in RFC 7232, 3.2 | 
 **ifMatch** | **string** | Validator for conditional requests, as described in RFC 7232, 3.2 | 
 **getPrevious** | **bool** | Retrieve the Meta Schema before update | [default to false]
 **supportedFeatures** | **string** | Features required to be supported by the target NF | 

### Return type

[**MetaSchema**](MetaSchema.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteMetaSchema

> MetaSchema DeleteMetaSchema(ctx, realmId, storageId, schemaId).IfMatch(ifMatch).GetPrevious(getPrevious).SupportedFeatures(supportedFeatures).Execute()

Delete a Meta Schema with an user provided SchemaId

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
    realmId := "Realm01" // string | Identifier(name) of the Realm
    storageId := "Storage01" // string | Identifier of the Storage
    schemaId := "schemaId_example" // string | Identifier of the Meta Schema
    ifMatch := "ifMatch_example" // string | Record validator for conditional requests, as described in RFC 7232, 3.2 (optional)
    getPrevious := true // bool | Retrieve the Meta Schema before delete (optional) (default to false)
    supportedFeatures := "supportedFeatures_example" // string | Features required to be supported by the target NF (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MetaSchemaCRUDApi.DeleteMetaSchema(context.Background(), realmId, storageId, schemaId).IfMatch(ifMatch).GetPrevious(getPrevious).SupportedFeatures(supportedFeatures).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetaSchemaCRUDApi.DeleteMetaSchema``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteMetaSchema`: MetaSchema
    fmt.Fprintf(os.Stdout, "Response from `MetaSchemaCRUDApi.DeleteMetaSchema`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realmId** | **string** | Identifier(name) of the Realm | 
**storageId** | **string** | Identifier of the Storage | 
**schemaId** | **string** | Identifier of the Meta Schema | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMetaSchemaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **ifMatch** | **string** | Record validator for conditional requests, as described in RFC 7232, 3.2 | 
 **getPrevious** | **bool** | Retrieve the Meta Schema before delete | [default to false]
 **supportedFeatures** | **string** | Features required to be supported by the target NF | 

### Return type

[**MetaSchema**](MetaSchema.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMetaSchema

> Record GetMetaSchema(ctx, realmId, storageId, schemaId).IfNoneMatch(ifNoneMatch).IfModifiedSince(ifModifiedSince).SupportedFeatures(supportedFeatures).Execute()

Meta Schema access



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
    realmId := "Realm01" // string | Identifier of the Realm
    storageId := "Storage01" // string | Identifier of the Storage
    schemaId := "schemaId_example" // string | Identifier of the Meta Schema
    ifNoneMatch := "ifNoneMatch_example" // string | Validator for conditional requests, as described in RFC 7232, 3.2 (optional)
    ifModifiedSince := "ifModifiedSince_example" // string | Validator for conditional requests, as described in RFC 7232, 3.3 (optional)
    supportedFeatures := "supportedFeatures_example" // string | Features required to be supported by the target NF (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MetaSchemaCRUDApi.GetMetaSchema(context.Background(), realmId, storageId, schemaId).IfNoneMatch(ifNoneMatch).IfModifiedSince(ifModifiedSince).SupportedFeatures(supportedFeatures).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetaSchemaCRUDApi.GetMetaSchema``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMetaSchema`: Record
    fmt.Fprintf(os.Stdout, "Response from `MetaSchemaCRUDApi.GetMetaSchema`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realmId** | **string** | Identifier of the Realm | 
**storageId** | **string** | Identifier of the Storage | 
**schemaId** | **string** | Identifier of the Meta Schema | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMetaSchemaRequest struct via the builder pattern


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

