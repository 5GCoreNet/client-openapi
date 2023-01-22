# DataSubscription

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AmfDataSub** | Pointer to [**AmfEventSubscription**](AmfEventSubscription.md) |  | [optional] 
**SmfDataSub** | Pointer to [**NsmfEventExposure**](NsmfEventExposure.md) |  | [optional] 
**UdmDataSub** | Pointer to [**EeSubscription**](EeSubscription.md) |  | [optional] 
**AfDataSub** | Pointer to [**AfEventExposureSubsc**](AfEventExposureSubsc.md) |  | [optional] 
**NefDataSub** | Pointer to [**NefEventExposureSubsc**](NefEventExposureSubsc.md) |  | [optional] 
**NrfDataSub** | Pointer to [**SubscriptionData**](SubscriptionData.md) |  | [optional] 
**NsacfDataSub** | Pointer to [**SACEventSubscription**](SACEventSubscription.md) |  | [optional] 

## Methods

### NewDataSubscription

`func NewDataSubscription() *DataSubscription`

NewDataSubscription instantiates a new DataSubscription object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataSubscriptionWithDefaults

`func NewDataSubscriptionWithDefaults() *DataSubscription`

NewDataSubscriptionWithDefaults instantiates a new DataSubscription object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmfDataSub

`func (o *DataSubscription) GetAmfDataSub() AmfEventSubscription`

GetAmfDataSub returns the AmfDataSub field if non-nil, zero value otherwise.

### GetAmfDataSubOk

`func (o *DataSubscription) GetAmfDataSubOk() (*AmfEventSubscription, bool)`

GetAmfDataSubOk returns a tuple with the AmfDataSub field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmfDataSub

`func (o *DataSubscription) SetAmfDataSub(v AmfEventSubscription)`

SetAmfDataSub sets AmfDataSub field to given value.

### HasAmfDataSub

`func (o *DataSubscription) HasAmfDataSub() bool`

HasAmfDataSub returns a boolean if a field has been set.

### GetSmfDataSub

`func (o *DataSubscription) GetSmfDataSub() NsmfEventExposure`

GetSmfDataSub returns the SmfDataSub field if non-nil, zero value otherwise.

### GetSmfDataSubOk

`func (o *DataSubscription) GetSmfDataSubOk() (*NsmfEventExposure, bool)`

GetSmfDataSubOk returns a tuple with the SmfDataSub field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmfDataSub

`func (o *DataSubscription) SetSmfDataSub(v NsmfEventExposure)`

SetSmfDataSub sets SmfDataSub field to given value.

### HasSmfDataSub

`func (o *DataSubscription) HasSmfDataSub() bool`

HasSmfDataSub returns a boolean if a field has been set.

### GetUdmDataSub

`func (o *DataSubscription) GetUdmDataSub() EeSubscription`

GetUdmDataSub returns the UdmDataSub field if non-nil, zero value otherwise.

### GetUdmDataSubOk

`func (o *DataSubscription) GetUdmDataSubOk() (*EeSubscription, bool)`

GetUdmDataSubOk returns a tuple with the UdmDataSub field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUdmDataSub

`func (o *DataSubscription) SetUdmDataSub(v EeSubscription)`

SetUdmDataSub sets UdmDataSub field to given value.

### HasUdmDataSub

`func (o *DataSubscription) HasUdmDataSub() bool`

HasUdmDataSub returns a boolean if a field has been set.

### GetAfDataSub

`func (o *DataSubscription) GetAfDataSub() AfEventExposureSubsc`

GetAfDataSub returns the AfDataSub field if non-nil, zero value otherwise.

### GetAfDataSubOk

`func (o *DataSubscription) GetAfDataSubOk() (*AfEventExposureSubsc, bool)`

GetAfDataSubOk returns a tuple with the AfDataSub field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAfDataSub

`func (o *DataSubscription) SetAfDataSub(v AfEventExposureSubsc)`

SetAfDataSub sets AfDataSub field to given value.

### HasAfDataSub

`func (o *DataSubscription) HasAfDataSub() bool`

HasAfDataSub returns a boolean if a field has been set.

### GetNefDataSub

`func (o *DataSubscription) GetNefDataSub() NefEventExposureSubsc`

GetNefDataSub returns the NefDataSub field if non-nil, zero value otherwise.

### GetNefDataSubOk

`func (o *DataSubscription) GetNefDataSubOk() (*NefEventExposureSubsc, bool)`

GetNefDataSubOk returns a tuple with the NefDataSub field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNefDataSub

`func (o *DataSubscription) SetNefDataSub(v NefEventExposureSubsc)`

SetNefDataSub sets NefDataSub field to given value.

### HasNefDataSub

`func (o *DataSubscription) HasNefDataSub() bool`

HasNefDataSub returns a boolean if a field has been set.

### GetNrfDataSub

`func (o *DataSubscription) GetNrfDataSub() SubscriptionData`

GetNrfDataSub returns the NrfDataSub field if non-nil, zero value otherwise.

### GetNrfDataSubOk

`func (o *DataSubscription) GetNrfDataSubOk() (*SubscriptionData, bool)`

GetNrfDataSubOk returns a tuple with the NrfDataSub field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNrfDataSub

`func (o *DataSubscription) SetNrfDataSub(v SubscriptionData)`

SetNrfDataSub sets NrfDataSub field to given value.

### HasNrfDataSub

`func (o *DataSubscription) HasNrfDataSub() bool`

HasNrfDataSub returns a boolean if a field has been set.

### GetNsacfDataSub

`func (o *DataSubscription) GetNsacfDataSub() SACEventSubscription`

GetNsacfDataSub returns the NsacfDataSub field if non-nil, zero value otherwise.

### GetNsacfDataSubOk

`func (o *DataSubscription) GetNsacfDataSubOk() (*SACEventSubscription, bool)`

GetNsacfDataSubOk returns a tuple with the NsacfDataSub field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNsacfDataSub

`func (o *DataSubscription) SetNsacfDataSub(v SACEventSubscription)`

SetNsacfDataSub sets NsacfDataSub field to given value.

### HasNsacfDataSub

`func (o *DataSubscription) HasNsacfDataSub() bool`

HasNsacfDataSub returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


