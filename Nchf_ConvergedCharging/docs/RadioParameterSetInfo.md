# RadioParameterSetInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RadioParameterSetValues** | Pointer to **[]string** |  | [optional] 
**ChangeTimestamp** | Pointer to **time.Time** | string with format &#39;date-time&#39; as defined in OpenAPI. | [optional] 

## Methods

### NewRadioParameterSetInfo

`func NewRadioParameterSetInfo() *RadioParameterSetInfo`

NewRadioParameterSetInfo instantiates a new RadioParameterSetInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRadioParameterSetInfoWithDefaults

`func NewRadioParameterSetInfoWithDefaults() *RadioParameterSetInfo`

NewRadioParameterSetInfoWithDefaults instantiates a new RadioParameterSetInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRadioParameterSetValues

`func (o *RadioParameterSetInfo) GetRadioParameterSetValues() []string`

GetRadioParameterSetValues returns the RadioParameterSetValues field if non-nil, zero value otherwise.

### GetRadioParameterSetValuesOk

`func (o *RadioParameterSetInfo) GetRadioParameterSetValuesOk() (*[]string, bool)`

GetRadioParameterSetValuesOk returns a tuple with the RadioParameterSetValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadioParameterSetValues

`func (o *RadioParameterSetInfo) SetRadioParameterSetValues(v []string)`

SetRadioParameterSetValues sets RadioParameterSetValues field to given value.

### HasRadioParameterSetValues

`func (o *RadioParameterSetInfo) HasRadioParameterSetValues() bool`

HasRadioParameterSetValues returns a boolean if a field has been set.

### GetChangeTimestamp

`func (o *RadioParameterSetInfo) GetChangeTimestamp() time.Time`

GetChangeTimestamp returns the ChangeTimestamp field if non-nil, zero value otherwise.

### GetChangeTimestampOk

`func (o *RadioParameterSetInfo) GetChangeTimestampOk() (*time.Time, bool)`

GetChangeTimestampOk returns a tuple with the ChangeTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeTimestamp

`func (o *RadioParameterSetInfo) SetChangeTimestamp(v time.Time)`

SetChangeTimestamp sets ChangeTimestamp field to given value.

### HasChangeTimestamp

`func (o *RadioParameterSetInfo) HasChangeTimestamp() bool`

HasChangeTimestamp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


