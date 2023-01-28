# \ADRFDataStoreRecordsCollectionApi

All URIs are relative to *https://example.com/nadrf-datamanagement/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateADRFDataStoreRecord**](ADRFDataStoreRecordsCollectionApi.md#CreateADRFDataStoreRecord) | **Post** /data-store-records | Creates a new Individual Data Store Record resource.
[**GetAdrfDataStoreRecords**](ADRFDataStoreRecordsCollectionApi.md#GetAdrfDataStoreRecords) | **Get** /data-store-records | Retrieves existing Individual ADRF Data Store Records.



## CreateADRFDataStoreRecord

> NadrfDataStoreRecord CreateADRFDataStoreRecord(ctx).NadrfDataStoreRecord(nadrfDataStoreRecord).Execute()

Creates a new Individual Data Store Record resource.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/Nadrf_DataManagement"
)

func main() {
    nadrfDataStoreRecord := openapiclient.NadrfDataStoreRecord{Interface{}: new(interface{})} // NadrfDataStoreRecord | ADRF data store record to be stored.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ADRFDataStoreRecordsCollectionApi.CreateADRFDataStoreRecord(context.Background()).NadrfDataStoreRecord(nadrfDataStoreRecord).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ADRFDataStoreRecordsCollectionApi.CreateADRFDataStoreRecord``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateADRFDataStoreRecord`: NadrfDataStoreRecord
    fmt.Fprintf(os.Stdout, "Response from `ADRFDataStoreRecordsCollectionApi.CreateADRFDataStoreRecord`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateADRFDataStoreRecordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **nadrfDataStoreRecord** | [**NadrfDataStoreRecord**](NadrfDataStoreRecord.md) | ADRF data store record to be stored. | 

### Return type

[**NadrfDataStoreRecord**](NadrfDataStoreRecord.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAdrfDataStoreRecords

> NadrfDataStoreRecord GetAdrfDataStoreRecords(ctx).StoreTransId(storeTransId).FetchCorrelationIds(fetchCorrelationIds).AnaSub(anaSub).AmfDataSub(amfDataSub).SmfDataSub(smfDataSub).NefDataSub(nefDataSub).UdmDataSub(udmDataSub).AfDataSub(afDataSub).TimePeriod(timePeriod).Execute()

Retrieves existing Individual ADRF Data Store Records.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/Nadrf_DataManagement"
)

func main() {
    storeTransId := "storeTransId_example" // string | A storage transaction identifier of a data store record in ADRF. (optional)
    fetchCorrelationIds := []string{"Inner_example"} // []string | Fetch correlation identifiers received as part of fetch instruction. (optional)
    anaSub := map[string][]openapiclient.NnwdafEventsSubscription{ ... } // NnwdafEventsSubscription | Represents analytics event subscription. (optional)
    amfDataSub := map[string][]openapiclient.AmfEventSubscription{ ... } // AmfEventSubscription | Represents AMF event subscription. (optional)
    smfDataSub := map[string][]openapiclient.NsmfEventExposure{ ... } // NsmfEventExposure | Represents SMF event subscription. (optional)
    nefDataSub := map[string][]openapiclient.NefEventExposureSubsc{ ... } // NefEventExposureSubsc | Represents NEF event subscription. (optional)
    udmDataSub := map[string][]openapiclient.EeSubscription{ ... } // EeSubscription | Represents UDM event subscription. (optional)
    afDataSub := map[string][]openapiclient.AfEventExposureSubsc{ ... } // AfEventExposureSubsc | Represents AF event subscription. (optional)
    timePeriod := map[string][]openapiclient.TimeWindow{ ... } // TimeWindow | Represents a start time and a stop time during which requested data is collected  or to be collected.  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ADRFDataStoreRecordsCollectionApi.GetAdrfDataStoreRecords(context.Background()).StoreTransId(storeTransId).FetchCorrelationIds(fetchCorrelationIds).AnaSub(anaSub).AmfDataSub(amfDataSub).SmfDataSub(smfDataSub).NefDataSub(nefDataSub).UdmDataSub(udmDataSub).AfDataSub(afDataSub).TimePeriod(timePeriod).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ADRFDataStoreRecordsCollectionApi.GetAdrfDataStoreRecords``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAdrfDataStoreRecords`: NadrfDataStoreRecord
    fmt.Fprintf(os.Stdout, "Response from `ADRFDataStoreRecordsCollectionApi.GetAdrfDataStoreRecords`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAdrfDataStoreRecordsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **storeTransId** | **string** | A storage transaction identifier of a data store record in ADRF. | 
 **fetchCorrelationIds** | **[]string** | Fetch correlation identifiers received as part of fetch instruction. | 
 **anaSub** | [**NnwdafEventsSubscription**](NnwdafEventsSubscription.md) | Represents analytics event subscription. | 
 **amfDataSub** | [**AmfEventSubscription**](AmfEventSubscription.md) | Represents AMF event subscription. | 
 **smfDataSub** | [**NsmfEventExposure**](NsmfEventExposure.md) | Represents SMF event subscription. | 
 **nefDataSub** | [**NefEventExposureSubsc**](NefEventExposureSubsc.md) | Represents NEF event subscription. | 
 **udmDataSub** | [**EeSubscription**](EeSubscription.md) | Represents UDM event subscription. | 
 **afDataSub** | [**AfEventExposureSubsc**](AfEventExposureSubsc.md) | Represents AF event subscription. | 
 **timePeriod** | [**TimeWindow**](TimeWindow.md) | Represents a start time and a stop time during which requested data is collected  or to be collected.  | 

### Return type

[**NadrfDataStoreRecord**](NadrfDataStoreRecord.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

