# SubscriptionData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NfInstanceId** | **string** | String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.   | 
**CallbackReference** | **string** | String providing an URI formatted according to RFC 3986. | 
**MonitoredResourceUris** | **[]string** |  | 
**Expires** | Pointer to **time.Time** | string with format &#39;date-time&#39; as defined in OpenAPI. | [optional] 
**ImmediateReport** | Pointer to **bool** |  | [optional] [default to false]
**Report** | Pointer to [**SubscriptionDataSets**](SubscriptionDataSets.md) |  | [optional] 

## Methods

### NewSubscriptionData

`func NewSubscriptionData(nfInstanceId string, callbackReference string, monitoredResourceUris []string, ) *SubscriptionData`

NewSubscriptionData instantiates a new SubscriptionData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriptionDataWithDefaults

`func NewSubscriptionDataWithDefaults() *SubscriptionData`

NewSubscriptionDataWithDefaults instantiates a new SubscriptionData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNfInstanceId

`func (o *SubscriptionData) GetNfInstanceId() string`

GetNfInstanceId returns the NfInstanceId field if non-nil, zero value otherwise.

### GetNfInstanceIdOk

`func (o *SubscriptionData) GetNfInstanceIdOk() (*string, bool)`

GetNfInstanceIdOk returns a tuple with the NfInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNfInstanceId

`func (o *SubscriptionData) SetNfInstanceId(v string)`

SetNfInstanceId sets NfInstanceId field to given value.


### GetCallbackReference

`func (o *SubscriptionData) GetCallbackReference() string`

GetCallbackReference returns the CallbackReference field if non-nil, zero value otherwise.

### GetCallbackReferenceOk

`func (o *SubscriptionData) GetCallbackReferenceOk() (*string, bool)`

GetCallbackReferenceOk returns a tuple with the CallbackReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallbackReference

`func (o *SubscriptionData) SetCallbackReference(v string)`

SetCallbackReference sets CallbackReference field to given value.


### GetMonitoredResourceUris

`func (o *SubscriptionData) GetMonitoredResourceUris() []string`

GetMonitoredResourceUris returns the MonitoredResourceUris field if non-nil, zero value otherwise.

### GetMonitoredResourceUrisOk

`func (o *SubscriptionData) GetMonitoredResourceUrisOk() (*[]string, bool)`

GetMonitoredResourceUrisOk returns a tuple with the MonitoredResourceUris field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitoredResourceUris

`func (o *SubscriptionData) SetMonitoredResourceUris(v []string)`

SetMonitoredResourceUris sets MonitoredResourceUris field to given value.


### GetExpires

`func (o *SubscriptionData) GetExpires() time.Time`

GetExpires returns the Expires field if non-nil, zero value otherwise.

### GetExpiresOk

`func (o *SubscriptionData) GetExpiresOk() (*time.Time, bool)`

GetExpiresOk returns a tuple with the Expires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpires

`func (o *SubscriptionData) SetExpires(v time.Time)`

SetExpires sets Expires field to given value.

### HasExpires

`func (o *SubscriptionData) HasExpires() bool`

HasExpires returns a boolean if a field has been set.

### GetImmediateReport

`func (o *SubscriptionData) GetImmediateReport() bool`

GetImmediateReport returns the ImmediateReport field if non-nil, zero value otherwise.

### GetImmediateReportOk

`func (o *SubscriptionData) GetImmediateReportOk() (*bool, bool)`

GetImmediateReportOk returns a tuple with the ImmediateReport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImmediateReport

`func (o *SubscriptionData) SetImmediateReport(v bool)`

SetImmediateReport sets ImmediateReport field to given value.

### HasImmediateReport

`func (o *SubscriptionData) HasImmediateReport() bool`

HasImmediateReport returns a boolean if a field has been set.

### GetReport

`func (o *SubscriptionData) GetReport() SubscriptionDataSets`

GetReport returns the Report field if non-nil, zero value otherwise.

### GetReportOk

`func (o *SubscriptionData) GetReportOk() (*SubscriptionDataSets, bool)`

GetReportOk returns a tuple with the Report field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReport

`func (o *SubscriptionData) SetReport(v SubscriptionDataSets)`

SetReport sets Report field to given value.

### HasReport

`func (o *SubscriptionData) HasReport() bool`

HasReport returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


