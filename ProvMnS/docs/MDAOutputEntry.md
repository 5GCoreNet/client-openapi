# MDAOutputEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MDAOutputIEName** | Pointer to **string** |  | [optional] 
**MdaOutputIEValue** | Pointer to **interface{}** |  | [optional] 
**AnalyticsWindow** | Pointer to [**TimeWindow**](TimeWindow.md) |  | [optional] 
**ConfidenceDegree** | Pointer to **float32** |  | [optional] 

## Methods

### NewMDAOutputEntry

`func NewMDAOutputEntry() *MDAOutputEntry`

NewMDAOutputEntry instantiates a new MDAOutputEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMDAOutputEntryWithDefaults

`func NewMDAOutputEntryWithDefaults() *MDAOutputEntry`

NewMDAOutputEntryWithDefaults instantiates a new MDAOutputEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMDAOutputIEName

`func (o *MDAOutputEntry) GetMDAOutputIEName() string`

GetMDAOutputIEName returns the MDAOutputIEName field if non-nil, zero value otherwise.

### GetMDAOutputIENameOk

`func (o *MDAOutputEntry) GetMDAOutputIENameOk() (*string, bool)`

GetMDAOutputIENameOk returns a tuple with the MDAOutputIEName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMDAOutputIEName

`func (o *MDAOutputEntry) SetMDAOutputIEName(v string)`

SetMDAOutputIEName sets MDAOutputIEName field to given value.

### HasMDAOutputIEName

`func (o *MDAOutputEntry) HasMDAOutputIEName() bool`

HasMDAOutputIEName returns a boolean if a field has been set.

### GetMdaOutputIEValue

`func (o *MDAOutputEntry) GetMdaOutputIEValue() interface{}`

GetMdaOutputIEValue returns the MdaOutputIEValue field if non-nil, zero value otherwise.

### GetMdaOutputIEValueOk

`func (o *MDAOutputEntry) GetMdaOutputIEValueOk() (*interface{}, bool)`

GetMdaOutputIEValueOk returns a tuple with the MdaOutputIEValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMdaOutputIEValue

`func (o *MDAOutputEntry) SetMdaOutputIEValue(v interface{})`

SetMdaOutputIEValue sets MdaOutputIEValue field to given value.

### HasMdaOutputIEValue

`func (o *MDAOutputEntry) HasMdaOutputIEValue() bool`

HasMdaOutputIEValue returns a boolean if a field has been set.

### SetMdaOutputIEValueNil

`func (o *MDAOutputEntry) SetMdaOutputIEValueNil(b bool)`

 SetMdaOutputIEValueNil sets the value for MdaOutputIEValue to be an explicit nil

### UnsetMdaOutputIEValue
`func (o *MDAOutputEntry) UnsetMdaOutputIEValue()`

UnsetMdaOutputIEValue ensures that no value is present for MdaOutputIEValue, not even an explicit nil
### GetAnalyticsWindow

`func (o *MDAOutputEntry) GetAnalyticsWindow() TimeWindow`

GetAnalyticsWindow returns the AnalyticsWindow field if non-nil, zero value otherwise.

### GetAnalyticsWindowOk

`func (o *MDAOutputEntry) GetAnalyticsWindowOk() (*TimeWindow, bool)`

GetAnalyticsWindowOk returns a tuple with the AnalyticsWindow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnalyticsWindow

`func (o *MDAOutputEntry) SetAnalyticsWindow(v TimeWindow)`

SetAnalyticsWindow sets AnalyticsWindow field to given value.

### HasAnalyticsWindow

`func (o *MDAOutputEntry) HasAnalyticsWindow() bool`

HasAnalyticsWindow returns a boolean if a field has been set.

### GetConfidenceDegree

`func (o *MDAOutputEntry) GetConfidenceDegree() float32`

GetConfidenceDegree returns the ConfidenceDegree field if non-nil, zero value otherwise.

### GetConfidenceDegreeOk

`func (o *MDAOutputEntry) GetConfidenceDegreeOk() (*float32, bool)`

GetConfidenceDegreeOk returns a tuple with the ConfidenceDegree field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfidenceDegree

`func (o *MDAOutputEntry) SetConfidenceDegree(v float32)`

SetConfidenceDegree sets ConfidenceDegree field to given value.

### HasConfidenceDegree

`func (o *MDAOutputEntry) HasConfidenceDegree() bool`

HasConfidenceDegree returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


