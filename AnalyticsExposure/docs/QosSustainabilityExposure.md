# QosSustainabilityExposure

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LocArea** | [**LocationArea5G**](LocationArea5G.md) |  | 
**StartTs** | **time.Time** | string with format \&quot;date-time\&quot; as defined in OpenAPI. | 
**EndTs** | **time.Time** | string with format \&quot;date-time\&quot; as defined in OpenAPI. | 
**QosFlowRetThd** | Pointer to [**RetainabilityThreshold**](RetainabilityThreshold.md) |  | [optional] 
**RanUeThrouThd** | Pointer to **string** | String representing a bit rate; the prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \&quot;K\&quot; is used to represent the standard symbol \&quot;k\&quot;.  | [optional] 
**Confidence** | Pointer to **int32** | Unsigned Integer, i.e. only value 0 and integers above 0 are permissible. | [optional] 

## Methods

### NewQosSustainabilityExposure

`func NewQosSustainabilityExposure(locArea LocationArea5G, startTs time.Time, endTs time.Time, ) *QosSustainabilityExposure`

NewQosSustainabilityExposure instantiates a new QosSustainabilityExposure object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQosSustainabilityExposureWithDefaults

`func NewQosSustainabilityExposureWithDefaults() *QosSustainabilityExposure`

NewQosSustainabilityExposureWithDefaults instantiates a new QosSustainabilityExposure object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocArea

`func (o *QosSustainabilityExposure) GetLocArea() LocationArea5G`

GetLocArea returns the LocArea field if non-nil, zero value otherwise.

### GetLocAreaOk

`func (o *QosSustainabilityExposure) GetLocAreaOk() (*LocationArea5G, bool)`

GetLocAreaOk returns a tuple with the LocArea field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocArea

`func (o *QosSustainabilityExposure) SetLocArea(v LocationArea5G)`

SetLocArea sets LocArea field to given value.


### GetStartTs

`func (o *QosSustainabilityExposure) GetStartTs() time.Time`

GetStartTs returns the StartTs field if non-nil, zero value otherwise.

### GetStartTsOk

`func (o *QosSustainabilityExposure) GetStartTsOk() (*time.Time, bool)`

GetStartTsOk returns a tuple with the StartTs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTs

`func (o *QosSustainabilityExposure) SetStartTs(v time.Time)`

SetStartTs sets StartTs field to given value.


### GetEndTs

`func (o *QosSustainabilityExposure) GetEndTs() time.Time`

GetEndTs returns the EndTs field if non-nil, zero value otherwise.

### GetEndTsOk

`func (o *QosSustainabilityExposure) GetEndTsOk() (*time.Time, bool)`

GetEndTsOk returns a tuple with the EndTs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTs

`func (o *QosSustainabilityExposure) SetEndTs(v time.Time)`

SetEndTs sets EndTs field to given value.


### GetQosFlowRetThd

`func (o *QosSustainabilityExposure) GetQosFlowRetThd() RetainabilityThreshold`

GetQosFlowRetThd returns the QosFlowRetThd field if non-nil, zero value otherwise.

### GetQosFlowRetThdOk

`func (o *QosSustainabilityExposure) GetQosFlowRetThdOk() (*RetainabilityThreshold, bool)`

GetQosFlowRetThdOk returns a tuple with the QosFlowRetThd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQosFlowRetThd

`func (o *QosSustainabilityExposure) SetQosFlowRetThd(v RetainabilityThreshold)`

SetQosFlowRetThd sets QosFlowRetThd field to given value.

### HasQosFlowRetThd

`func (o *QosSustainabilityExposure) HasQosFlowRetThd() bool`

HasQosFlowRetThd returns a boolean if a field has been set.

### GetRanUeThrouThd

`func (o *QosSustainabilityExposure) GetRanUeThrouThd() string`

GetRanUeThrouThd returns the RanUeThrouThd field if non-nil, zero value otherwise.

### GetRanUeThrouThdOk

`func (o *QosSustainabilityExposure) GetRanUeThrouThdOk() (*string, bool)`

GetRanUeThrouThdOk returns a tuple with the RanUeThrouThd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRanUeThrouThd

`func (o *QosSustainabilityExposure) SetRanUeThrouThd(v string)`

SetRanUeThrouThd sets RanUeThrouThd field to given value.

### HasRanUeThrouThd

`func (o *QosSustainabilityExposure) HasRanUeThrouThd() bool`

HasRanUeThrouThd returns a boolean if a field has been set.

### GetConfidence

`func (o *QosSustainabilityExposure) GetConfidence() int32`

GetConfidence returns the Confidence field if non-nil, zero value otherwise.

### GetConfidenceOk

`func (o *QosSustainabilityExposure) GetConfidenceOk() (*int32, bool)`

GetConfidenceOk returns a tuple with the Confidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfidence

`func (o *QosSustainabilityExposure) SetConfidence(v int32)`

SetConfidence sets Confidence field to given value.

### HasConfidence

`func (o *QosSustainabilityExposure) HasConfidence() bool`

HasConfidence returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


