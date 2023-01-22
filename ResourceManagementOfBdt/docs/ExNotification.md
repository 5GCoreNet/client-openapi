# ExNotification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BdtRefId** | **string** | string identifying a BDT Reference ID as defined in clause 5.3.3 of 3GPP TS 29.154. | 
**LocationArea5G** | Pointer to [**LocationArea5G**](LocationArea5G.md) |  | [optional] 
**TimeWindow** | Pointer to [**TimeWindow**](TimeWindow.md) |  | [optional] 
**CandPolicies** | Pointer to [**[]TransferPolicy**](TransferPolicy.md) | This IE indicates a list of the candidate transfer policies from which the AF may select a new transfer policy due to network performance degradation. | [optional] 

## Methods

### NewExNotification

`func NewExNotification(bdtRefId string, ) *ExNotification`

NewExNotification instantiates a new ExNotification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExNotificationWithDefaults

`func NewExNotificationWithDefaults() *ExNotification`

NewExNotificationWithDefaults instantiates a new ExNotification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBdtRefId

`func (o *ExNotification) GetBdtRefId() string`

GetBdtRefId returns the BdtRefId field if non-nil, zero value otherwise.

### GetBdtRefIdOk

`func (o *ExNotification) GetBdtRefIdOk() (*string, bool)`

GetBdtRefIdOk returns a tuple with the BdtRefId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBdtRefId

`func (o *ExNotification) SetBdtRefId(v string)`

SetBdtRefId sets BdtRefId field to given value.


### GetLocationArea5G

`func (o *ExNotification) GetLocationArea5G() LocationArea5G`

GetLocationArea5G returns the LocationArea5G field if non-nil, zero value otherwise.

### GetLocationArea5GOk

`func (o *ExNotification) GetLocationArea5GOk() (*LocationArea5G, bool)`

GetLocationArea5GOk returns a tuple with the LocationArea5G field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationArea5G

`func (o *ExNotification) SetLocationArea5G(v LocationArea5G)`

SetLocationArea5G sets LocationArea5G field to given value.

### HasLocationArea5G

`func (o *ExNotification) HasLocationArea5G() bool`

HasLocationArea5G returns a boolean if a field has been set.

### GetTimeWindow

`func (o *ExNotification) GetTimeWindow() TimeWindow`

GetTimeWindow returns the TimeWindow field if non-nil, zero value otherwise.

### GetTimeWindowOk

`func (o *ExNotification) GetTimeWindowOk() (*TimeWindow, bool)`

GetTimeWindowOk returns a tuple with the TimeWindow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeWindow

`func (o *ExNotification) SetTimeWindow(v TimeWindow)`

SetTimeWindow sets TimeWindow field to given value.

### HasTimeWindow

`func (o *ExNotification) HasTimeWindow() bool`

HasTimeWindow returns a boolean if a field has been set.

### GetCandPolicies

`func (o *ExNotification) GetCandPolicies() []TransferPolicy`

GetCandPolicies returns the CandPolicies field if non-nil, zero value otherwise.

### GetCandPoliciesOk

`func (o *ExNotification) GetCandPoliciesOk() (*[]TransferPolicy, bool)`

GetCandPoliciesOk returns a tuple with the CandPolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCandPolicies

`func (o *ExNotification) SetCandPolicies(v []TransferPolicy)`

SetCandPolicies sets CandPolicies field to given value.

### HasCandPolicies

`func (o *ExNotification) HasCandPolicies() bool`

HasCandPolicies returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


