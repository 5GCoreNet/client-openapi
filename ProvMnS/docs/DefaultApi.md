# \DefaultApi

All URIs are relative to *http://example.com/3GPPManagement/ProvMnS/XXX*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ClassNameidDelete**](DefaultApi.md#ClassNameidDelete) | **Delete** /{className}&#x3D;{id} | Deletes one or multiple resources
[**ClassNameidGet**](DefaultApi.md#ClassNameidGet) | **Get** /{className}&#x3D;{id} | Reads one or multiple resources
[**ClassNameidPatch**](DefaultApi.md#ClassNameidPatch) | **Patch** /{className}&#x3D;{id} | Patches one or multiple resources
[**ClassNameidPut**](DefaultApi.md#ClassNameidPut) | **Put** /{className}&#x3D;{id} | Replaces a complete single resource or creates it if it does not exist



## ClassNameidDelete

> ClassNameidDelete(ctx, className, id).Scope(scope).Filter(filter).Execute()

Deletes one or multiple resources



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
    className := "className_example" // string | 
    id := "id_example" // string | 
    scope := map[string][]openapiclient.Scope{ ... } // Scope | This parameter extends the set of targeted resources beyond the base resource identified with the path component of the URI. No scoping mechanism is specified in the present document. (optional)
    filter := "filter_example" // string | This parameter reduces the targeted set of resources by applying a filter to the scoped set of resource representations. Only resources representations for which the filter construct evaluates to \"true\" are returned. No filter language is specified in the present document. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ClassNameidDelete(context.Background(), className, id).Scope(scope).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ClassNameidDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**className** | **string** |  | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiClassNameidDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **scope** | [**Scope**](Scope.md) | This parameter extends the set of targeted resources beyond the base resource identified with the path component of the URI. No scoping mechanism is specified in the present document. | 
 **filter** | **string** | This parameter reduces the targeted set of resources by applying a filter to the scoped set of resource representations. Only resources representations for which the filter construct evaluates to \&quot;true\&quot; are returned. No filter language is specified in the present document. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ClassNameidGet

> Resource ClassNameidGet(ctx, className, id).Attributes(attributes).Scope(scope).Filter(filter).Fields(fields).Execute()

Reads one or multiple resources



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
    className := "className_example" // string | 
    id := "id_example" // string | 
    attributes := []string{"Inner_example"} // []string | This parameter specifies the attributes of the scoped resources that are returned.
    scope := map[string][]openapiclient.Scope{ ... } // Scope | This parameter extends the set of targeted resources beyond the base resource identified with the path component of the URI. No scoping mechanism is specified in the present document. (optional)
    filter := "filter_example" // string | This parameter reduces the targeted set of resources by applying a filter to the scoped set of resource representations. Only resource representations for which the filter construct evaluates to \"true\" are targeted. No filter language is specified in the present document. (optional)
    fields := []string{"Inner_example"} // []string | This parameter specifies the attribute field of the scoped resources that are returned. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ClassNameidGet(context.Background(), className, id).Attributes(attributes).Scope(scope).Filter(filter).Fields(fields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ClassNameidGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ClassNameidGet`: Resource
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ClassNameidGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**className** | **string** |  | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiClassNameidGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **attributes** | **[]string** | This parameter specifies the attributes of the scoped resources that are returned. | 
 **scope** | [**Scope**](Scope.md) | This parameter extends the set of targeted resources beyond the base resource identified with the path component of the URI. No scoping mechanism is specified in the present document. | 
 **filter** | **string** | This parameter reduces the targeted set of resources by applying a filter to the scoped set of resource representations. Only resource representations for which the filter construct evaluates to \&quot;true\&quot; are targeted. No filter language is specified in the present document. | 
 **fields** | **[]string** | This parameter specifies the attribute field of the scoped resources that are returned. | 

### Return type

[**Resource**](Resource.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ClassNameidPatch

> Resource ClassNameidPatch(ctx, className, id).Resource(resource).Execute()

Patches one or multiple resources



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
    className := "className_example" // string | 
    id := "id_example" // string | 
    resource := openapiclient.Resource{AnyOfresourcesGenericNrmresourcesNrNrmresources5gcNrmresourcesSliceNrmresourcesCoslaNrmresourcesIntentNrmresourcesMdaNrmresourcesAiMlNrm: new(AnyOfresourcesGenericNrmresourcesNrNrmresources5gcNrmresourcesSliceNrmresourcesCoslaNrmresourcesIntentNrmresourcesMdaNrmresourcesAiMlNrm)} // Resource | The request body describes changes to be made to the target resources. The following patch media types are available   - \"application/merge-patch+json\" (RFC 7396)   - \"application/3gpp-merge-patch+json\" (TS 32.158)   - \"application/json-patch+json\" (RFC 6902)   - \"application/3gpp-json-patch+json\" (TS 32.158)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ClassNameidPatch(context.Background(), className, id).Resource(resource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ClassNameidPatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ClassNameidPatch`: Resource
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ClassNameidPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**className** | **string** |  | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiClassNameidPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **resource** | [**Resource**](Resource.md) | The request body describes changes to be made to the target resources. The following patch media types are available   - \&quot;application/merge-patch+json\&quot; (RFC 7396)   - \&quot;application/3gpp-merge-patch+json\&quot; (TS 32.158)   - \&quot;application/json-patch+json\&quot; (RFC 6902)   - \&quot;application/3gpp-json-patch+json\&quot; (TS 32.158) | 

### Return type

[**Resource**](Resource.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/merge-patch+json, application/3gpp-merge-patch+json, application/json-patch+json, application/3gpp-json-patch+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ClassNameidPut

> Resource ClassNameidPut(ctx, className, id).Resource(resource).Execute()

Replaces a complete single resource or creates it if it does not exist



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
    className := "className_example" // string | 
    id := "id_example" // string | 
    resource := openapiclient.Resource{AnyOfresourcesGenericNrmresourcesNrNrmresources5gcNrmresourcesSliceNrmresourcesCoslaNrmresourcesIntentNrmresourcesMdaNrmresourcesAiMlNrm: new(AnyOfresourcesGenericNrmresourcesNrNrmresources5gcNrmresourcesSliceNrmresourcesCoslaNrmresourcesIntentNrmresourcesMdaNrmresourcesAiMlNrm)} // Resource | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ClassNameidPut(context.Background(), className, id).Resource(resource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ClassNameidPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ClassNameidPut`: Resource
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ClassNameidPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**className** | **string** |  | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiClassNameidPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **resource** | [**Resource**](Resource.md) |  | 

### Return type

[**Resource**](Resource.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

