# \NetworkSliceInformationDocumentApi

All URIs are relative to *https://example.com/nnssf-nsselection/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**NSSelectionGet**](NetworkSliceInformationDocumentApi.md#NSSelectionGet) | **Get** /network-slice-information | Retrieve the Network Slice Selection Information



## NSSelectionGet

> AuthorizedNetworkSliceInfo NSSelectionGet(ctx).NfType(nfType).NfId(nfId).SliceInfoRequestForRegistration(sliceInfoRequestForRegistration).SliceInfoRequestForPduSession(sliceInfoRequestForPduSession).SliceInfoRequestForUeCu(sliceInfoRequestForUeCu).HomePlmnId(homePlmnId).Tai(tai).SupportedFeatures(supportedFeatures).Execute()

Retrieve the Network Slice Selection Information

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/Nnssf_NSSelection"
)

func main() {
    nfType := *openapiclient.NewNFType() // NFType | NF type of the NF service consumer
    nfId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | NF Instance ID of the NF service consumer
    sliceInfoRequestForRegistration := map[string][]openapiclient.SliceInfoForRegistration{ ... } // SliceInfoForRegistration | Requested network slice information during Registration procedure (optional)
    sliceInfoRequestForPduSession := map[string][]openapiclient.SliceInfoForPDUSession{ ... } // SliceInfoForPDUSession | Requested network slice information during PDU session establishment procedure (optional)
    sliceInfoRequestForUeCu := map[string][]openapiclient.SliceInfoForUEConfigurationUpdate{ ... } // SliceInfoForUEConfigurationUpdate | Requested network slice information during UE confuguration update procedure (optional)
    homePlmnId := map[string][]openapiclient.PlmnId{ ... } // PlmnId | PLMN ID of the HPLMN (optional)
    tai := map[string][]openapiclient.Tai{ ... } // Tai | TAI of the UE (optional)
    supportedFeatures := "supportedFeatures_example" // string | Features required to be supported by the NFs in the target slice instance (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworkSliceInformationDocumentApi.NSSelectionGet(context.Background()).NfType(nfType).NfId(nfId).SliceInfoRequestForRegistration(sliceInfoRequestForRegistration).SliceInfoRequestForPduSession(sliceInfoRequestForPduSession).SliceInfoRequestForUeCu(sliceInfoRequestForUeCu).HomePlmnId(homePlmnId).Tai(tai).SupportedFeatures(supportedFeatures).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkSliceInformationDocumentApi.NSSelectionGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `NSSelectionGet`: AuthorizedNetworkSliceInfo
    fmt.Fprintf(os.Stdout, "Response from `NetworkSliceInformationDocumentApi.NSSelectionGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiNSSelectionGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **nfType** | [**NFType**](NFType.md) | NF type of the NF service consumer | 
 **nfId** | **string** | NF Instance ID of the NF service consumer | 
 **sliceInfoRequestForRegistration** | [**SliceInfoForRegistration**](SliceInfoForRegistration.md) | Requested network slice information during Registration procedure | 
 **sliceInfoRequestForPduSession** | [**SliceInfoForPDUSession**](SliceInfoForPDUSession.md) | Requested network slice information during PDU session establishment procedure | 
 **sliceInfoRequestForUeCu** | [**SliceInfoForUEConfigurationUpdate**](SliceInfoForUEConfigurationUpdate.md) | Requested network slice information during UE confuguration update procedure | 
 **homePlmnId** | [**PlmnId**](PlmnId.md) | PLMN ID of the HPLMN | 
 **tai** | [**Tai**](Tai.md) | TAI of the UE | 
 **supportedFeatures** | **string** | Features required to be supported by the NFs in the target slice instance | 

### Return type

[**AuthorizedNetworkSliceInfo**](AuthorizedNetworkSliceInfo.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

