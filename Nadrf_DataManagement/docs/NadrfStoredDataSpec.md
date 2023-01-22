# NadrfStoredDataSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DataSpec** | Pointer to [**DataSubscription**](DataSubscription.md) |  | [optional] 
**AnaSpec** | Pointer to [**NnwdafEventsSubscription**](NnwdafEventsSubscription.md) |  | [optional] 
**TimePeriod** | [**TimeWindow**](TimeWindow.md) |  | 

## Methods

### NewNadrfStoredDataSpec

`func NewNadrfStoredDataSpec(timePeriod TimeWindow, ) *NadrfStoredDataSpec`

NewNadrfStoredDataSpec instantiates a new NadrfStoredDataSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNadrfStoredDataSpecWithDefaults

`func NewNadrfStoredDataSpecWithDefaults() *NadrfStoredDataSpec`

NewNadrfStoredDataSpecWithDefaults instantiates a new NadrfStoredDataSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDataSpec

`func (o *NadrfStoredDataSpec) GetDataSpec() DataSubscription`

GetDataSpec returns the DataSpec field if non-nil, zero value otherwise.

### GetDataSpecOk

`func (o *NadrfStoredDataSpec) GetDataSpecOk() (*DataSubscription, bool)`

GetDataSpecOk returns a tuple with the DataSpec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSpec

`func (o *NadrfStoredDataSpec) SetDataSpec(v DataSubscription)`

SetDataSpec sets DataSpec field to given value.

### HasDataSpec

`func (o *NadrfStoredDataSpec) HasDataSpec() bool`

HasDataSpec returns a boolean if a field has been set.

### GetAnaSpec

`func (o *NadrfStoredDataSpec) GetAnaSpec() NnwdafEventsSubscription`

GetAnaSpec returns the AnaSpec field if non-nil, zero value otherwise.

### GetAnaSpecOk

`func (o *NadrfStoredDataSpec) GetAnaSpecOk() (*NnwdafEventsSubscription, bool)`

GetAnaSpecOk returns a tuple with the AnaSpec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnaSpec

`func (o *NadrfStoredDataSpec) SetAnaSpec(v NnwdafEventsSubscription)`

SetAnaSpec sets AnaSpec field to given value.

### HasAnaSpec

`func (o *NadrfStoredDataSpec) HasAnaSpec() bool`

HasAnaSpec returns a boolean if a field has been set.

### GetTimePeriod

`func (o *NadrfStoredDataSpec) GetTimePeriod() TimeWindow`

GetTimePeriod returns the TimePeriod field if non-nil, zero value otherwise.

### GetTimePeriodOk

`func (o *NadrfStoredDataSpec) GetTimePeriodOk() (*TimeWindow, bool)`

GetTimePeriodOk returns a tuple with the TimePeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimePeriod

`func (o *NadrfStoredDataSpec) SetTimePeriod(v TimeWindow)`

SetTimePeriod sets TimePeriod field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


