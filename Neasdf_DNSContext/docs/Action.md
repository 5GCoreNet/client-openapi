# Action

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplyAction** | [**ApplyAction**](ApplyAction.md) |  | 
**FwdParas** | Pointer to [**ForwardingParameters**](ForwardingParameters.md) |  | [optional] 
**ReportingOnceInd** | Pointer to **bool** |  | [optional] [default to false]
**ResetReportingOnceInd** | Pointer to **bool** |  | [optional] [default to false]

## Methods

### NewAction

`func NewAction(applyAction ApplyAction, ) *Action`

NewAction instantiates a new Action object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActionWithDefaults

`func NewActionWithDefaults() *Action`

NewActionWithDefaults instantiates a new Action object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplyAction

`func (o *Action) GetApplyAction() ApplyAction`

GetApplyAction returns the ApplyAction field if non-nil, zero value otherwise.

### GetApplyActionOk

`func (o *Action) GetApplyActionOk() (*ApplyAction, bool)`

GetApplyActionOk returns a tuple with the ApplyAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplyAction

`func (o *Action) SetApplyAction(v ApplyAction)`

SetApplyAction sets ApplyAction field to given value.


### GetFwdParas

`func (o *Action) GetFwdParas() ForwardingParameters`

GetFwdParas returns the FwdParas field if non-nil, zero value otherwise.

### GetFwdParasOk

`func (o *Action) GetFwdParasOk() (*ForwardingParameters, bool)`

GetFwdParasOk returns a tuple with the FwdParas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFwdParas

`func (o *Action) SetFwdParas(v ForwardingParameters)`

SetFwdParas sets FwdParas field to given value.

### HasFwdParas

`func (o *Action) HasFwdParas() bool`

HasFwdParas returns a boolean if a field has been set.

### GetReportingOnceInd

`func (o *Action) GetReportingOnceInd() bool`

GetReportingOnceInd returns the ReportingOnceInd field if non-nil, zero value otherwise.

### GetReportingOnceIndOk

`func (o *Action) GetReportingOnceIndOk() (*bool, bool)`

GetReportingOnceIndOk returns a tuple with the ReportingOnceInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportingOnceInd

`func (o *Action) SetReportingOnceInd(v bool)`

SetReportingOnceInd sets ReportingOnceInd field to given value.

### HasReportingOnceInd

`func (o *Action) HasReportingOnceInd() bool`

HasReportingOnceInd returns a boolean if a field has been set.

### GetResetReportingOnceInd

`func (o *Action) GetResetReportingOnceInd() bool`

GetResetReportingOnceInd returns the ResetReportingOnceInd field if non-nil, zero value otherwise.

### GetResetReportingOnceIndOk

`func (o *Action) GetResetReportingOnceIndOk() (*bool, bool)`

GetResetReportingOnceIndOk returns a tuple with the ResetReportingOnceInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResetReportingOnceInd

`func (o *Action) SetResetReportingOnceInd(v bool)`

SetResetReportingOnceInd sets ResetReportingOnceInd field to given value.

### HasResetReportingOnceInd

`func (o *Action) HasResetReportingOnceInd() bool`

HasResetReportingOnceInd returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


