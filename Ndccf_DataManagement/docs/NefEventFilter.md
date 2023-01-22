# NefEventFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TgtUe** | [**TargetUeIdentification**](TargetUeIdentification.md) |  | 
**AppIds** | Pointer to **[]string** |  | [optional] 
**LocArea** | Pointer to [**NetworkAreaInfo**](NetworkAreaInfo.md) |  | [optional] 
**CollAttrs** | Pointer to [**[]CollectiveBehaviourFilter**](CollectiveBehaviourFilter.md) |  | [optional] 

## Methods

### NewNefEventFilter

`func NewNefEventFilter(tgtUe TargetUeIdentification, ) *NefEventFilter`

NewNefEventFilter instantiates a new NefEventFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNefEventFilterWithDefaults

`func NewNefEventFilterWithDefaults() *NefEventFilter`

NewNefEventFilterWithDefaults instantiates a new NefEventFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTgtUe

`func (o *NefEventFilter) GetTgtUe() TargetUeIdentification`

GetTgtUe returns the TgtUe field if non-nil, zero value otherwise.

### GetTgtUeOk

`func (o *NefEventFilter) GetTgtUeOk() (*TargetUeIdentification, bool)`

GetTgtUeOk returns a tuple with the TgtUe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTgtUe

`func (o *NefEventFilter) SetTgtUe(v TargetUeIdentification)`

SetTgtUe sets TgtUe field to given value.


### GetAppIds

`func (o *NefEventFilter) GetAppIds() []string`

GetAppIds returns the AppIds field if non-nil, zero value otherwise.

### GetAppIdsOk

`func (o *NefEventFilter) GetAppIdsOk() (*[]string, bool)`

GetAppIdsOk returns a tuple with the AppIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppIds

`func (o *NefEventFilter) SetAppIds(v []string)`

SetAppIds sets AppIds field to given value.

### HasAppIds

`func (o *NefEventFilter) HasAppIds() bool`

HasAppIds returns a boolean if a field has been set.

### GetLocArea

`func (o *NefEventFilter) GetLocArea() NetworkAreaInfo`

GetLocArea returns the LocArea field if non-nil, zero value otherwise.

### GetLocAreaOk

`func (o *NefEventFilter) GetLocAreaOk() (*NetworkAreaInfo, bool)`

GetLocAreaOk returns a tuple with the LocArea field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocArea

`func (o *NefEventFilter) SetLocArea(v NetworkAreaInfo)`

SetLocArea sets LocArea field to given value.

### HasLocArea

`func (o *NefEventFilter) HasLocArea() bool`

HasLocArea returns a boolean if a field has been set.

### GetCollAttrs

`func (o *NefEventFilter) GetCollAttrs() []CollectiveBehaviourFilter`

GetCollAttrs returns the CollAttrs field if non-nil, zero value otherwise.

### GetCollAttrsOk

`func (o *NefEventFilter) GetCollAttrsOk() (*[]CollectiveBehaviourFilter, bool)`

GetCollAttrsOk returns a tuple with the CollAttrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollAttrs

`func (o *NefEventFilter) SetCollAttrs(v []CollectiveBehaviourFilter)`

SetCollAttrs sets CollAttrs field to given value.

### HasCollAttrs

`func (o *NefEventFilter) HasCollAttrs() bool`

HasCollAttrs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


