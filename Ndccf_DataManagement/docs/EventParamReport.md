# EventParamReport

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the reported parameter. | 
**Values** | **[]interface{}** | The list of values of the reported parameter. | 
**Supi** | Pointer to **string** | String identifying a Supi that shall contain either an IMSI, a network specific identifier, a Global Cable Identifier (GCI) or a Global Line Identifier (GLI) as specified in clause  2.2A of 3GPP TS 23.003. It shall be formatted as follows  - for an IMSI \&quot;imsi-&lt;imsi&gt;\&quot;, where &lt;imsi&gt; shall be formatted according to clause 2.2    of 3GPP TS 23.003 that describes an IMSI.  - for a network specific identifier \&quot;nai-&lt;nai&gt;, where &lt;nai&gt; shall be formatted    according to clause 28.7.2 of 3GPP TS 23.003 that describes an NAI.  - for a GCI \&quot;gci-&lt;gci&gt;\&quot;, where &lt;gci&gt; shall be formatted according to clause 28.15.2    of 3GPP TS 23.003.  - for a GLI \&quot;gli-&lt;gli&gt;\&quot;, where &lt;gli&gt; shall be formatted according to clause 28.16.2 of    3GPP TS 23.003.To enable that the value is used as part of an URI, the string shall    only contain characters allowed according to the \&quot;lower-with-hyphen\&quot; naming convention    defined in 3GPP TS 29.501.  | [optional] 
**Area** | Pointer to [**NetworkAreaInfo**](NetworkAreaInfo.md) |  | [optional] 
**Spacing** | Pointer to [**NumberAverage**](NumberAverage.md) |  | [optional] 
**Duration** | Pointer to [**NumberAverage**](NumberAverage.md) |  | [optional] 
**AvgAndVar** | Pointer to [**NumberAverage**](NumberAverage.md) |  | [optional] 
**MostFreqVal** | Pointer to **interface{}** |  | [optional] 
**LeastFreqVal** | Pointer to **interface{}** |  | [optional] 
**Count** | Pointer to **int32** | Unsigned Integer, i.e. only value 0 and integers above 0 are permissible. | [optional] 
**MinValue** | Pointer to **string** | The minimum value of the parameter. | [optional] 
**MaxValue** | Pointer to **string** | The maximum value of the parameter. | [optional] 

## Methods

### NewEventParamReport

`func NewEventParamReport(name string, values []interface{}, ) *EventParamReport`

NewEventParamReport instantiates a new EventParamReport object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventParamReportWithDefaults

`func NewEventParamReportWithDefaults() *EventParamReport`

NewEventParamReportWithDefaults instantiates a new EventParamReport object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *EventParamReport) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EventParamReport) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EventParamReport) SetName(v string)`

SetName sets Name field to given value.


### GetValues

`func (o *EventParamReport) GetValues() []interface{}`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *EventParamReport) GetValuesOk() (*[]interface{}, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *EventParamReport) SetValues(v []interface{})`

SetValues sets Values field to given value.


### GetSupi

`func (o *EventParamReport) GetSupi() string`

GetSupi returns the Supi field if non-nil, zero value otherwise.

### GetSupiOk

`func (o *EventParamReport) GetSupiOk() (*string, bool)`

GetSupiOk returns a tuple with the Supi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupi

`func (o *EventParamReport) SetSupi(v string)`

SetSupi sets Supi field to given value.

### HasSupi

`func (o *EventParamReport) HasSupi() bool`

HasSupi returns a boolean if a field has been set.

### GetArea

`func (o *EventParamReport) GetArea() NetworkAreaInfo`

GetArea returns the Area field if non-nil, zero value otherwise.

### GetAreaOk

`func (o *EventParamReport) GetAreaOk() (*NetworkAreaInfo, bool)`

GetAreaOk returns a tuple with the Area field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArea

`func (o *EventParamReport) SetArea(v NetworkAreaInfo)`

SetArea sets Area field to given value.

### HasArea

`func (o *EventParamReport) HasArea() bool`

HasArea returns a boolean if a field has been set.

### GetSpacing

`func (o *EventParamReport) GetSpacing() NumberAverage`

GetSpacing returns the Spacing field if non-nil, zero value otherwise.

### GetSpacingOk

`func (o *EventParamReport) GetSpacingOk() (*NumberAverage, bool)`

GetSpacingOk returns a tuple with the Spacing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpacing

`func (o *EventParamReport) SetSpacing(v NumberAverage)`

SetSpacing sets Spacing field to given value.

### HasSpacing

`func (o *EventParamReport) HasSpacing() bool`

HasSpacing returns a boolean if a field has been set.

### GetDuration

`func (o *EventParamReport) GetDuration() NumberAverage`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *EventParamReport) GetDurationOk() (*NumberAverage, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *EventParamReport) SetDuration(v NumberAverage)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *EventParamReport) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetAvgAndVar

`func (o *EventParamReport) GetAvgAndVar() NumberAverage`

GetAvgAndVar returns the AvgAndVar field if non-nil, zero value otherwise.

### GetAvgAndVarOk

`func (o *EventParamReport) GetAvgAndVarOk() (*NumberAverage, bool)`

GetAvgAndVarOk returns a tuple with the AvgAndVar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvgAndVar

`func (o *EventParamReport) SetAvgAndVar(v NumberAverage)`

SetAvgAndVar sets AvgAndVar field to given value.

### HasAvgAndVar

`func (o *EventParamReport) HasAvgAndVar() bool`

HasAvgAndVar returns a boolean if a field has been set.

### GetMostFreqVal

`func (o *EventParamReport) GetMostFreqVal() interface{}`

GetMostFreqVal returns the MostFreqVal field if non-nil, zero value otherwise.

### GetMostFreqValOk

`func (o *EventParamReport) GetMostFreqValOk() (*interface{}, bool)`

GetMostFreqValOk returns a tuple with the MostFreqVal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMostFreqVal

`func (o *EventParamReport) SetMostFreqVal(v interface{})`

SetMostFreqVal sets MostFreqVal field to given value.

### HasMostFreqVal

`func (o *EventParamReport) HasMostFreqVal() bool`

HasMostFreqVal returns a boolean if a field has been set.

### SetMostFreqValNil

`func (o *EventParamReport) SetMostFreqValNil(b bool)`

 SetMostFreqValNil sets the value for MostFreqVal to be an explicit nil

### UnsetMostFreqVal
`func (o *EventParamReport) UnsetMostFreqVal()`

UnsetMostFreqVal ensures that no value is present for MostFreqVal, not even an explicit nil
### GetLeastFreqVal

`func (o *EventParamReport) GetLeastFreqVal() interface{}`

GetLeastFreqVal returns the LeastFreqVal field if non-nil, zero value otherwise.

### GetLeastFreqValOk

`func (o *EventParamReport) GetLeastFreqValOk() (*interface{}, bool)`

GetLeastFreqValOk returns a tuple with the LeastFreqVal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeastFreqVal

`func (o *EventParamReport) SetLeastFreqVal(v interface{})`

SetLeastFreqVal sets LeastFreqVal field to given value.

### HasLeastFreqVal

`func (o *EventParamReport) HasLeastFreqVal() bool`

HasLeastFreqVal returns a boolean if a field has been set.

### SetLeastFreqValNil

`func (o *EventParamReport) SetLeastFreqValNil(b bool)`

 SetLeastFreqValNil sets the value for LeastFreqVal to be an explicit nil

### UnsetLeastFreqVal
`func (o *EventParamReport) UnsetLeastFreqVal()`

UnsetLeastFreqVal ensures that no value is present for LeastFreqVal, not even an explicit nil
### GetCount

`func (o *EventParamReport) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *EventParamReport) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *EventParamReport) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *EventParamReport) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetMinValue

`func (o *EventParamReport) GetMinValue() string`

GetMinValue returns the MinValue field if non-nil, zero value otherwise.

### GetMinValueOk

`func (o *EventParamReport) GetMinValueOk() (*string, bool)`

GetMinValueOk returns a tuple with the MinValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinValue

`func (o *EventParamReport) SetMinValue(v string)`

SetMinValue sets MinValue field to given value.

### HasMinValue

`func (o *EventParamReport) HasMinValue() bool`

HasMinValue returns a boolean if a field has been set.

### GetMaxValue

`func (o *EventParamReport) GetMaxValue() string`

GetMaxValue returns the MaxValue field if non-nil, zero value otherwise.

### GetMaxValueOk

`func (o *EventParamReport) GetMaxValueOk() (*string, bool)`

GetMaxValueOk returns a tuple with the MaxValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxValue

`func (o *EventParamReport) SetMaxValue(v string)`

SetMaxValue sets MaxValue field to given value.

### HasMaxValue

`func (o *EventParamReport) HasMaxValue() bool`

HasMaxValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


