# SdRange

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Start** | Pointer to **string** | First value identifying the start of an SD range. This string shall be formatted as specified for the sd attribute of the Snssai data type in clause 5.4.4.2.  | [optional] 
**End** | Pointer to **string** | Last value identifying the end of an SD range. This string shall be formatted as specified for the sd attribute of the Snssai data type in clause 5.4.4.2.  | [optional] 

## Methods

### NewSdRange

`func NewSdRange() *SdRange`

NewSdRange instantiates a new SdRange object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSdRangeWithDefaults

`func NewSdRangeWithDefaults() *SdRange`

NewSdRangeWithDefaults instantiates a new SdRange object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStart

`func (o *SdRange) GetStart() string`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *SdRange) GetStartOk() (*string, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *SdRange) SetStart(v string)`

SetStart sets Start field to given value.

### HasStart

`func (o *SdRange) HasStart() bool`

HasStart returns a boolean if a field has been set.

### GetEnd

`func (o *SdRange) GetEnd() string`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *SdRange) GetEndOk() (*string, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *SdRange) SetEnd(v string)`

SetEnd sets End field to given value.

### HasEnd

`func (o *SdRange) HasEnd() bool`

HasEnd returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


