# TransferPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BdtPolicyId** | **int32** | Identifier for the transfer policy | 
**MaxUplinkBandwidth** | Pointer to **int32** | integer indicating a bandwidth in bits per second. | [optional] 
**MaxDownlinkBandwidth** | Pointer to **int32** | integer indicating a bandwidth in bits per second. | [optional] 
**RatingGroup** | **int32** | Indicates the rating group during the time window. | 
**TimeWindow** | [**TimeWindow**](TimeWindow.md) |  | 

## Methods

### NewTransferPolicy

`func NewTransferPolicy(bdtPolicyId int32, ratingGroup int32, timeWindow TimeWindow, ) *TransferPolicy`

NewTransferPolicy instantiates a new TransferPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransferPolicyWithDefaults

`func NewTransferPolicyWithDefaults() *TransferPolicy`

NewTransferPolicyWithDefaults instantiates a new TransferPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBdtPolicyId

`func (o *TransferPolicy) GetBdtPolicyId() int32`

GetBdtPolicyId returns the BdtPolicyId field if non-nil, zero value otherwise.

### GetBdtPolicyIdOk

`func (o *TransferPolicy) GetBdtPolicyIdOk() (*int32, bool)`

GetBdtPolicyIdOk returns a tuple with the BdtPolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBdtPolicyId

`func (o *TransferPolicy) SetBdtPolicyId(v int32)`

SetBdtPolicyId sets BdtPolicyId field to given value.


### GetMaxUplinkBandwidth

`func (o *TransferPolicy) GetMaxUplinkBandwidth() int32`

GetMaxUplinkBandwidth returns the MaxUplinkBandwidth field if non-nil, zero value otherwise.

### GetMaxUplinkBandwidthOk

`func (o *TransferPolicy) GetMaxUplinkBandwidthOk() (*int32, bool)`

GetMaxUplinkBandwidthOk returns a tuple with the MaxUplinkBandwidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxUplinkBandwidth

`func (o *TransferPolicy) SetMaxUplinkBandwidth(v int32)`

SetMaxUplinkBandwidth sets MaxUplinkBandwidth field to given value.

### HasMaxUplinkBandwidth

`func (o *TransferPolicy) HasMaxUplinkBandwidth() bool`

HasMaxUplinkBandwidth returns a boolean if a field has been set.

### GetMaxDownlinkBandwidth

`func (o *TransferPolicy) GetMaxDownlinkBandwidth() int32`

GetMaxDownlinkBandwidth returns the MaxDownlinkBandwidth field if non-nil, zero value otherwise.

### GetMaxDownlinkBandwidthOk

`func (o *TransferPolicy) GetMaxDownlinkBandwidthOk() (*int32, bool)`

GetMaxDownlinkBandwidthOk returns a tuple with the MaxDownlinkBandwidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxDownlinkBandwidth

`func (o *TransferPolicy) SetMaxDownlinkBandwidth(v int32)`

SetMaxDownlinkBandwidth sets MaxDownlinkBandwidth field to given value.

### HasMaxDownlinkBandwidth

`func (o *TransferPolicy) HasMaxDownlinkBandwidth() bool`

HasMaxDownlinkBandwidth returns a boolean if a field has been set.

### GetRatingGroup

`func (o *TransferPolicy) GetRatingGroup() int32`

GetRatingGroup returns the RatingGroup field if non-nil, zero value otherwise.

### GetRatingGroupOk

`func (o *TransferPolicy) GetRatingGroupOk() (*int32, bool)`

GetRatingGroupOk returns a tuple with the RatingGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRatingGroup

`func (o *TransferPolicy) SetRatingGroup(v int32)`

SetRatingGroup sets RatingGroup field to given value.


### GetTimeWindow

`func (o *TransferPolicy) GetTimeWindow() TimeWindow`

GetTimeWindow returns the TimeWindow field if non-nil, zero value otherwise.

### GetTimeWindowOk

`func (o *TransferPolicy) GetTimeWindowOk() (*TimeWindow, bool)`

GetTimeWindowOk returns a tuple with the TimeWindow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeWindow

`func (o *TransferPolicy) SetTimeWindow(v TimeWindow)`

SetTimeWindow sets TimeWindow field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


