# ModelInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AnalyticsId** | [**NwdafEvent**](NwdafEvent.md) |  | 
**MlModelInfos** | [**[]MLModelInfo**](MLModelInfo.md) |  | 

## Methods

### NewModelInfo

`func NewModelInfo(analyticsId NwdafEvent, mlModelInfos []MLModelInfo, ) *ModelInfo`

NewModelInfo instantiates a new ModelInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelInfoWithDefaults

`func NewModelInfoWithDefaults() *ModelInfo`

NewModelInfoWithDefaults instantiates a new ModelInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnalyticsId

`func (o *ModelInfo) GetAnalyticsId() NwdafEvent`

GetAnalyticsId returns the AnalyticsId field if non-nil, zero value otherwise.

### GetAnalyticsIdOk

`func (o *ModelInfo) GetAnalyticsIdOk() (*NwdafEvent, bool)`

GetAnalyticsIdOk returns a tuple with the AnalyticsId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnalyticsId

`func (o *ModelInfo) SetAnalyticsId(v NwdafEvent)`

SetAnalyticsId sets AnalyticsId field to given value.


### GetMlModelInfos

`func (o *ModelInfo) GetMlModelInfos() []MLModelInfo`

GetMlModelInfos returns the MlModelInfos field if non-nil, zero value otherwise.

### GetMlModelInfosOk

`func (o *ModelInfo) GetMlModelInfosOk() (*[]MLModelInfo, bool)`

GetMlModelInfosOk returns a tuple with the MlModelInfos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMlModelInfos

`func (o *ModelInfo) SetMlModelInfos(v []MLModelInfo)`

SetMlModelInfos sets MlModelInfos field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


