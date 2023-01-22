# InternalGroupIdRange

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Start** | Pointer to **string** | String identifying a group of devices network internal globally unique ID which identifies a set of IMSIs, as specified in clause 19.9 of 3GPP TS 23.003.   | [optional] 
**End** | Pointer to **string** | String identifying a group of devices network internal globally unique ID which identifies a set of IMSIs, as specified in clause 19.9 of 3GPP TS 23.003.   | [optional] 
**Pattern** | Pointer to **string** |  | [optional] 

## Methods

### NewInternalGroupIdRange

`func NewInternalGroupIdRange() *InternalGroupIdRange`

NewInternalGroupIdRange instantiates a new InternalGroupIdRange object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInternalGroupIdRangeWithDefaults

`func NewInternalGroupIdRangeWithDefaults() *InternalGroupIdRange`

NewInternalGroupIdRangeWithDefaults instantiates a new InternalGroupIdRange object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStart

`func (o *InternalGroupIdRange) GetStart() string`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *InternalGroupIdRange) GetStartOk() (*string, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *InternalGroupIdRange) SetStart(v string)`

SetStart sets Start field to given value.

### HasStart

`func (o *InternalGroupIdRange) HasStart() bool`

HasStart returns a boolean if a field has been set.

### GetEnd

`func (o *InternalGroupIdRange) GetEnd() string`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *InternalGroupIdRange) GetEndOk() (*string, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *InternalGroupIdRange) SetEnd(v string)`

SetEnd sets End field to given value.

### HasEnd

`func (o *InternalGroupIdRange) HasEnd() bool`

HasEnd returns a boolean if a field has been set.

### GetPattern

`func (o *InternalGroupIdRange) GetPattern() string`

GetPattern returns the Pattern field if non-nil, zero value otherwise.

### GetPatternOk

`func (o *InternalGroupIdRange) GetPatternOk() (*string, bool)`

GetPatternOk returns a tuple with the Pattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPattern

`func (o *InternalGroupIdRange) SetPattern(v string)`

SetPattern sets Pattern field to given value.

### HasPattern

`func (o *InternalGroupIdRange) HasPattern() bool`

HasPattern returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


