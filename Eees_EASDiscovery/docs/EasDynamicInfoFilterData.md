# EasDynamicInfoFilterData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EecId** | **string** | Represents a unique identifier of the EEC. | 
**EasStatus** | Pointer to **bool** | Notify if EAS status changed. | [optional] 
**EasAcIds** | Pointer to **bool** | Notify if list of AC identifiers changed. | [optional] 
**EasDesc** | Pointer to **bool** | Notify if EAS description changed. | [optional] 
**EasPt** | Pointer to **bool** | Notify if EAS endpoint changed. | [optional] 
**EasFeature** | Pointer to **bool** | NotiNotify if EAS feature changed. | [optional] 
**EasSchedule** | Pointer to **bool** | Notify if EAS schedule changed. | [optional] 
**SvcArea** | Pointer to **bool** | Notify if EAS service area changed. | [optional] 
**SvcKpi** | Pointer to **bool** | Notify if EAS KPIs changed. | [optional] 
**SvcCont** | Pointer to **bool** | Notify if EAS supported ACR changed. | [optional] 

## Methods

### NewEasDynamicInfoFilterData

`func NewEasDynamicInfoFilterData(eecId string, ) *EasDynamicInfoFilterData`

NewEasDynamicInfoFilterData instantiates a new EasDynamicInfoFilterData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEasDynamicInfoFilterDataWithDefaults

`func NewEasDynamicInfoFilterDataWithDefaults() *EasDynamicInfoFilterData`

NewEasDynamicInfoFilterDataWithDefaults instantiates a new EasDynamicInfoFilterData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEecId

`func (o *EasDynamicInfoFilterData) GetEecId() string`

GetEecId returns the EecId field if non-nil, zero value otherwise.

### GetEecIdOk

`func (o *EasDynamicInfoFilterData) GetEecIdOk() (*string, bool)`

GetEecIdOk returns a tuple with the EecId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEecId

`func (o *EasDynamicInfoFilterData) SetEecId(v string)`

SetEecId sets EecId field to given value.


### GetEasStatus

`func (o *EasDynamicInfoFilterData) GetEasStatus() bool`

GetEasStatus returns the EasStatus field if non-nil, zero value otherwise.

### GetEasStatusOk

`func (o *EasDynamicInfoFilterData) GetEasStatusOk() (*bool, bool)`

GetEasStatusOk returns a tuple with the EasStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEasStatus

`func (o *EasDynamicInfoFilterData) SetEasStatus(v bool)`

SetEasStatus sets EasStatus field to given value.

### HasEasStatus

`func (o *EasDynamicInfoFilterData) HasEasStatus() bool`

HasEasStatus returns a boolean if a field has been set.

### GetEasAcIds

`func (o *EasDynamicInfoFilterData) GetEasAcIds() bool`

GetEasAcIds returns the EasAcIds field if non-nil, zero value otherwise.

### GetEasAcIdsOk

`func (o *EasDynamicInfoFilterData) GetEasAcIdsOk() (*bool, bool)`

GetEasAcIdsOk returns a tuple with the EasAcIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEasAcIds

`func (o *EasDynamicInfoFilterData) SetEasAcIds(v bool)`

SetEasAcIds sets EasAcIds field to given value.

### HasEasAcIds

`func (o *EasDynamicInfoFilterData) HasEasAcIds() bool`

HasEasAcIds returns a boolean if a field has been set.

### GetEasDesc

`func (o *EasDynamicInfoFilterData) GetEasDesc() bool`

GetEasDesc returns the EasDesc field if non-nil, zero value otherwise.

### GetEasDescOk

`func (o *EasDynamicInfoFilterData) GetEasDescOk() (*bool, bool)`

GetEasDescOk returns a tuple with the EasDesc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEasDesc

`func (o *EasDynamicInfoFilterData) SetEasDesc(v bool)`

SetEasDesc sets EasDesc field to given value.

### HasEasDesc

`func (o *EasDynamicInfoFilterData) HasEasDesc() bool`

HasEasDesc returns a boolean if a field has been set.

### GetEasPt

`func (o *EasDynamicInfoFilterData) GetEasPt() bool`

GetEasPt returns the EasPt field if non-nil, zero value otherwise.

### GetEasPtOk

`func (o *EasDynamicInfoFilterData) GetEasPtOk() (*bool, bool)`

GetEasPtOk returns a tuple with the EasPt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEasPt

`func (o *EasDynamicInfoFilterData) SetEasPt(v bool)`

SetEasPt sets EasPt field to given value.

### HasEasPt

`func (o *EasDynamicInfoFilterData) HasEasPt() bool`

HasEasPt returns a boolean if a field has been set.

### GetEasFeature

`func (o *EasDynamicInfoFilterData) GetEasFeature() bool`

GetEasFeature returns the EasFeature field if non-nil, zero value otherwise.

### GetEasFeatureOk

`func (o *EasDynamicInfoFilterData) GetEasFeatureOk() (*bool, bool)`

GetEasFeatureOk returns a tuple with the EasFeature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEasFeature

`func (o *EasDynamicInfoFilterData) SetEasFeature(v bool)`

SetEasFeature sets EasFeature field to given value.

### HasEasFeature

`func (o *EasDynamicInfoFilterData) HasEasFeature() bool`

HasEasFeature returns a boolean if a field has been set.

### GetEasSchedule

`func (o *EasDynamicInfoFilterData) GetEasSchedule() bool`

GetEasSchedule returns the EasSchedule field if non-nil, zero value otherwise.

### GetEasScheduleOk

`func (o *EasDynamicInfoFilterData) GetEasScheduleOk() (*bool, bool)`

GetEasScheduleOk returns a tuple with the EasSchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEasSchedule

`func (o *EasDynamicInfoFilterData) SetEasSchedule(v bool)`

SetEasSchedule sets EasSchedule field to given value.

### HasEasSchedule

`func (o *EasDynamicInfoFilterData) HasEasSchedule() bool`

HasEasSchedule returns a boolean if a field has been set.

### GetSvcArea

`func (o *EasDynamicInfoFilterData) GetSvcArea() bool`

GetSvcArea returns the SvcArea field if non-nil, zero value otherwise.

### GetSvcAreaOk

`func (o *EasDynamicInfoFilterData) GetSvcAreaOk() (*bool, bool)`

GetSvcAreaOk returns a tuple with the SvcArea field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSvcArea

`func (o *EasDynamicInfoFilterData) SetSvcArea(v bool)`

SetSvcArea sets SvcArea field to given value.

### HasSvcArea

`func (o *EasDynamicInfoFilterData) HasSvcArea() bool`

HasSvcArea returns a boolean if a field has been set.

### GetSvcKpi

`func (o *EasDynamicInfoFilterData) GetSvcKpi() bool`

GetSvcKpi returns the SvcKpi field if non-nil, zero value otherwise.

### GetSvcKpiOk

`func (o *EasDynamicInfoFilterData) GetSvcKpiOk() (*bool, bool)`

GetSvcKpiOk returns a tuple with the SvcKpi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSvcKpi

`func (o *EasDynamicInfoFilterData) SetSvcKpi(v bool)`

SetSvcKpi sets SvcKpi field to given value.

### HasSvcKpi

`func (o *EasDynamicInfoFilterData) HasSvcKpi() bool`

HasSvcKpi returns a boolean if a field has been set.

### GetSvcCont

`func (o *EasDynamicInfoFilterData) GetSvcCont() bool`

GetSvcCont returns the SvcCont field if non-nil, zero value otherwise.

### GetSvcContOk

`func (o *EasDynamicInfoFilterData) GetSvcContOk() (*bool, bool)`

GetSvcContOk returns a tuple with the SvcCont field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSvcCont

`func (o *EasDynamicInfoFilterData) SetSvcCont(v bool)`

SetSvcCont sets SvcCont field to given value.

### HasSvcCont

`func (o *EasDynamicInfoFilterData) HasSvcCont() bool`

HasSvcCont returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


