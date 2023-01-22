# CongestInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LocArea** | [**LocationArea5G**](LocationArea5G.md) |  | 
**CngAnas** | [**[]CongestionAnalytics**](CongestionAnalytics.md) |  | 

## Methods

### NewCongestInfo

`func NewCongestInfo(locArea LocationArea5G, cngAnas []CongestionAnalytics, ) *CongestInfo`

NewCongestInfo instantiates a new CongestInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCongestInfoWithDefaults

`func NewCongestInfoWithDefaults() *CongestInfo`

NewCongestInfoWithDefaults instantiates a new CongestInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocArea

`func (o *CongestInfo) GetLocArea() LocationArea5G`

GetLocArea returns the LocArea field if non-nil, zero value otherwise.

### GetLocAreaOk

`func (o *CongestInfo) GetLocAreaOk() (*LocationArea5G, bool)`

GetLocAreaOk returns a tuple with the LocArea field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocArea

`func (o *CongestInfo) SetLocArea(v LocationArea5G)`

SetLocArea sets LocArea field to given value.


### GetCngAnas

`func (o *CongestInfo) GetCngAnas() []CongestionAnalytics`

GetCngAnas returns the CngAnas field if non-nil, zero value otherwise.

### GetCngAnasOk

`func (o *CongestInfo) GetCngAnasOk() (*[]CongestionAnalytics, bool)`

GetCngAnasOk returns a tuple with the CngAnas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCngAnas

`func (o *CongestInfo) SetCngAnas(v []CongestionAnalytics)`

SetCngAnas sets CngAnas field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


