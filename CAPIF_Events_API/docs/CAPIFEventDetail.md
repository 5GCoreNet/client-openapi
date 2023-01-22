# CAPIFEventDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServiceAPIDescriptions** | Pointer to [**[]ServiceAPIDescription**](ServiceAPIDescription.md) | Description of the service API as published by the APF. | [optional] 
**ApiIds** | Pointer to **[]string** | Identifier of the service API | [optional] 
**ApiInvokerIds** | Pointer to **[]string** | Identity of the API invoker | [optional] 
**AccCtrlPolList** | Pointer to [**AccessControlPolicyListExt**](AccessControlPolicyListExt.md) |  | [optional] 
**InvocationLogs** | Pointer to [**[]InvocationLog**](InvocationLog.md) | Invocation logs. | [optional] 
**ApiTopoHide** | Pointer to [**TopologyHiding**](TopologyHiding.md) |  | [optional] 

## Methods

### NewCAPIFEventDetail

`func NewCAPIFEventDetail() *CAPIFEventDetail`

NewCAPIFEventDetail instantiates a new CAPIFEventDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCAPIFEventDetailWithDefaults

`func NewCAPIFEventDetailWithDefaults() *CAPIFEventDetail`

NewCAPIFEventDetailWithDefaults instantiates a new CAPIFEventDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServiceAPIDescriptions

`func (o *CAPIFEventDetail) GetServiceAPIDescriptions() []ServiceAPIDescription`

GetServiceAPIDescriptions returns the ServiceAPIDescriptions field if non-nil, zero value otherwise.

### GetServiceAPIDescriptionsOk

`func (o *CAPIFEventDetail) GetServiceAPIDescriptionsOk() (*[]ServiceAPIDescription, bool)`

GetServiceAPIDescriptionsOk returns a tuple with the ServiceAPIDescriptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAPIDescriptions

`func (o *CAPIFEventDetail) SetServiceAPIDescriptions(v []ServiceAPIDescription)`

SetServiceAPIDescriptions sets ServiceAPIDescriptions field to given value.

### HasServiceAPIDescriptions

`func (o *CAPIFEventDetail) HasServiceAPIDescriptions() bool`

HasServiceAPIDescriptions returns a boolean if a field has been set.

### GetApiIds

`func (o *CAPIFEventDetail) GetApiIds() []string`

GetApiIds returns the ApiIds field if non-nil, zero value otherwise.

### GetApiIdsOk

`func (o *CAPIFEventDetail) GetApiIdsOk() (*[]string, bool)`

GetApiIdsOk returns a tuple with the ApiIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiIds

`func (o *CAPIFEventDetail) SetApiIds(v []string)`

SetApiIds sets ApiIds field to given value.

### HasApiIds

`func (o *CAPIFEventDetail) HasApiIds() bool`

HasApiIds returns a boolean if a field has been set.

### GetApiInvokerIds

`func (o *CAPIFEventDetail) GetApiInvokerIds() []string`

GetApiInvokerIds returns the ApiInvokerIds field if non-nil, zero value otherwise.

### GetApiInvokerIdsOk

`func (o *CAPIFEventDetail) GetApiInvokerIdsOk() (*[]string, bool)`

GetApiInvokerIdsOk returns a tuple with the ApiInvokerIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiInvokerIds

`func (o *CAPIFEventDetail) SetApiInvokerIds(v []string)`

SetApiInvokerIds sets ApiInvokerIds field to given value.

### HasApiInvokerIds

`func (o *CAPIFEventDetail) HasApiInvokerIds() bool`

HasApiInvokerIds returns a boolean if a field has been set.

### GetAccCtrlPolList

`func (o *CAPIFEventDetail) GetAccCtrlPolList() AccessControlPolicyListExt`

GetAccCtrlPolList returns the AccCtrlPolList field if non-nil, zero value otherwise.

### GetAccCtrlPolListOk

`func (o *CAPIFEventDetail) GetAccCtrlPolListOk() (*AccessControlPolicyListExt, bool)`

GetAccCtrlPolListOk returns a tuple with the AccCtrlPolList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccCtrlPolList

`func (o *CAPIFEventDetail) SetAccCtrlPolList(v AccessControlPolicyListExt)`

SetAccCtrlPolList sets AccCtrlPolList field to given value.

### HasAccCtrlPolList

`func (o *CAPIFEventDetail) HasAccCtrlPolList() bool`

HasAccCtrlPolList returns a boolean if a field has been set.

### GetInvocationLogs

`func (o *CAPIFEventDetail) GetInvocationLogs() []InvocationLog`

GetInvocationLogs returns the InvocationLogs field if non-nil, zero value otherwise.

### GetInvocationLogsOk

`func (o *CAPIFEventDetail) GetInvocationLogsOk() (*[]InvocationLog, bool)`

GetInvocationLogsOk returns a tuple with the InvocationLogs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvocationLogs

`func (o *CAPIFEventDetail) SetInvocationLogs(v []InvocationLog)`

SetInvocationLogs sets InvocationLogs field to given value.

### HasInvocationLogs

`func (o *CAPIFEventDetail) HasInvocationLogs() bool`

HasInvocationLogs returns a boolean if a field has been set.

### GetApiTopoHide

`func (o *CAPIFEventDetail) GetApiTopoHide() TopologyHiding`

GetApiTopoHide returns the ApiTopoHide field if non-nil, zero value otherwise.

### GetApiTopoHideOk

`func (o *CAPIFEventDetail) GetApiTopoHideOk() (*TopologyHiding, bool)`

GetApiTopoHideOk returns a tuple with the ApiTopoHide field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiTopoHide

`func (o *CAPIFEventDetail) SetApiTopoHide(v TopologyHiding)`

SetApiTopoHide sets ApiTopoHide field to given value.

### HasApiTopoHide

`func (o *CAPIFEventDetail) HasApiTopoHide() bool`

HasApiTopoHide returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


