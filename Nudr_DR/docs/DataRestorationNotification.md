# DataRestorationNotification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SupiRanges** | Pointer to [**[]SupiRange**](SupiRange.md) |  | [optional] 
**GpsiRanges** | Pointer to [**[]IdentityRange**](IdentityRange.md) |  | [optional] 
**ResetIds** | Pointer to **[]string** |  | [optional] 
**SNssaiList** | Pointer to [**[]Snssai**](Snssai.md) |  | [optional] 
**DnnList** | Pointer to **[]string** |  | [optional] 
**UdrGroupId** | Pointer to **string** | Identifier of a group of NFs. | [optional] 

## Methods

### NewDataRestorationNotification

`func NewDataRestorationNotification() *DataRestorationNotification`

NewDataRestorationNotification instantiates a new DataRestorationNotification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataRestorationNotificationWithDefaults

`func NewDataRestorationNotificationWithDefaults() *DataRestorationNotification`

NewDataRestorationNotificationWithDefaults instantiates a new DataRestorationNotification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSupiRanges

`func (o *DataRestorationNotification) GetSupiRanges() []SupiRange`

GetSupiRanges returns the SupiRanges field if non-nil, zero value otherwise.

### GetSupiRangesOk

`func (o *DataRestorationNotification) GetSupiRangesOk() (*[]SupiRange, bool)`

GetSupiRangesOk returns a tuple with the SupiRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupiRanges

`func (o *DataRestorationNotification) SetSupiRanges(v []SupiRange)`

SetSupiRanges sets SupiRanges field to given value.

### HasSupiRanges

`func (o *DataRestorationNotification) HasSupiRanges() bool`

HasSupiRanges returns a boolean if a field has been set.

### GetGpsiRanges

`func (o *DataRestorationNotification) GetGpsiRanges() []IdentityRange`

GetGpsiRanges returns the GpsiRanges field if non-nil, zero value otherwise.

### GetGpsiRangesOk

`func (o *DataRestorationNotification) GetGpsiRangesOk() (*[]IdentityRange, bool)`

GetGpsiRangesOk returns a tuple with the GpsiRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpsiRanges

`func (o *DataRestorationNotification) SetGpsiRanges(v []IdentityRange)`

SetGpsiRanges sets GpsiRanges field to given value.

### HasGpsiRanges

`func (o *DataRestorationNotification) HasGpsiRanges() bool`

HasGpsiRanges returns a boolean if a field has been set.

### GetResetIds

`func (o *DataRestorationNotification) GetResetIds() []string`

GetResetIds returns the ResetIds field if non-nil, zero value otherwise.

### GetResetIdsOk

`func (o *DataRestorationNotification) GetResetIdsOk() (*[]string, bool)`

GetResetIdsOk returns a tuple with the ResetIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResetIds

`func (o *DataRestorationNotification) SetResetIds(v []string)`

SetResetIds sets ResetIds field to given value.

### HasResetIds

`func (o *DataRestorationNotification) HasResetIds() bool`

HasResetIds returns a boolean if a field has been set.

### GetSNssaiList

`func (o *DataRestorationNotification) GetSNssaiList() []Snssai`

GetSNssaiList returns the SNssaiList field if non-nil, zero value otherwise.

### GetSNssaiListOk

`func (o *DataRestorationNotification) GetSNssaiListOk() (*[]Snssai, bool)`

GetSNssaiListOk returns a tuple with the SNssaiList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSNssaiList

`func (o *DataRestorationNotification) SetSNssaiList(v []Snssai)`

SetSNssaiList sets SNssaiList field to given value.

### HasSNssaiList

`func (o *DataRestorationNotification) HasSNssaiList() bool`

HasSNssaiList returns a boolean if a field has been set.

### GetDnnList

`func (o *DataRestorationNotification) GetDnnList() []string`

GetDnnList returns the DnnList field if non-nil, zero value otherwise.

### GetDnnListOk

`func (o *DataRestorationNotification) GetDnnListOk() (*[]string, bool)`

GetDnnListOk returns a tuple with the DnnList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnnList

`func (o *DataRestorationNotification) SetDnnList(v []string)`

SetDnnList sets DnnList field to given value.

### HasDnnList

`func (o *DataRestorationNotification) HasDnnList() bool`

HasDnnList returns a boolean if a field has been set.

### GetUdrGroupId

`func (o *DataRestorationNotification) GetUdrGroupId() string`

GetUdrGroupId returns the UdrGroupId field if non-nil, zero value otherwise.

### GetUdrGroupIdOk

`func (o *DataRestorationNotification) GetUdrGroupIdOk() (*string, bool)`

GetUdrGroupIdOk returns a tuple with the UdrGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUdrGroupId

`func (o *DataRestorationNotification) SetUdrGroupId(v string)`

SetUdrGroupId sets UdrGroupId field to given value.

### HasUdrGroupId

`func (o *DataRestorationNotification) HasUdrGroupId() bool`

HasUdrGroupId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


