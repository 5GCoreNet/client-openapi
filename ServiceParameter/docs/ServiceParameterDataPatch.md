# ServiceParameterDataPatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ParamOverPc5** | Pointer to **NullableString** | Represents the same as the ParameterOverPc5 data type but with the nullable:true property.  | [optional] 
**ParamOverUu** | Pointer to **NullableString** | Represents the same as the ParameterOverUu data type but with the nullable:true property.  | [optional] 
**ParamForProSeDd** | Pointer to **NullableString** | This data type is defined in the same way as the ParamForProSeDd data type, but with the OpenAPI nullable property set to true.  | [optional] 
**ParamForProSeDc** | Pointer to **NullableString** | This data type is defined in the same way as the ParamForProSeDc data type, but with the OpenAPI nullable property set to true.  | [optional] 
**ParamForProSeU2NRelUe** | Pointer to **NullableString** | This data type is defined in the same way as the ParamForProSeU2NRelay data type, but with the OpenAPI nullable property set to true.  | [optional] 
**ParamForProSeRemUe** | Pointer to **NullableString** | This data type is defined in the same way as the ParamForProSeRemUe data type, but with the OpenAPI nullable property set to true.  | [optional] 
**UrspGuidance** | Pointer to [**[]UrspRuleRequest**](UrspRuleRequest.md) | Contains the service parameter used to guide the URSP. | [optional] 
**SubNotifEvents** | Pointer to [**[]Event**](Event.md) |  | [optional] 
**NotificationDestination** | Pointer to **string** | String providing an URI formatted according to RFC 3986. | [optional] 

## Methods

### NewServiceParameterDataPatch

`func NewServiceParameterDataPatch() *ServiceParameterDataPatch`

NewServiceParameterDataPatch instantiates a new ServiceParameterDataPatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceParameterDataPatchWithDefaults

`func NewServiceParameterDataPatchWithDefaults() *ServiceParameterDataPatch`

NewServiceParameterDataPatchWithDefaults instantiates a new ServiceParameterDataPatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetParamOverPc5

`func (o *ServiceParameterDataPatch) GetParamOverPc5() string`

GetParamOverPc5 returns the ParamOverPc5 field if non-nil, zero value otherwise.

### GetParamOverPc5Ok

`func (o *ServiceParameterDataPatch) GetParamOverPc5Ok() (*string, bool)`

GetParamOverPc5Ok returns a tuple with the ParamOverPc5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParamOverPc5

`func (o *ServiceParameterDataPatch) SetParamOverPc5(v string)`

SetParamOverPc5 sets ParamOverPc5 field to given value.

### HasParamOverPc5

`func (o *ServiceParameterDataPatch) HasParamOverPc5() bool`

HasParamOverPc5 returns a boolean if a field has been set.

### SetParamOverPc5Nil

`func (o *ServiceParameterDataPatch) SetParamOverPc5Nil(b bool)`

 SetParamOverPc5Nil sets the value for ParamOverPc5 to be an explicit nil

### UnsetParamOverPc5
`func (o *ServiceParameterDataPatch) UnsetParamOverPc5()`

UnsetParamOverPc5 ensures that no value is present for ParamOverPc5, not even an explicit nil
### GetParamOverUu

`func (o *ServiceParameterDataPatch) GetParamOverUu() string`

GetParamOverUu returns the ParamOverUu field if non-nil, zero value otherwise.

### GetParamOverUuOk

`func (o *ServiceParameterDataPatch) GetParamOverUuOk() (*string, bool)`

GetParamOverUuOk returns a tuple with the ParamOverUu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParamOverUu

`func (o *ServiceParameterDataPatch) SetParamOverUu(v string)`

SetParamOverUu sets ParamOverUu field to given value.

### HasParamOverUu

`func (o *ServiceParameterDataPatch) HasParamOverUu() bool`

HasParamOverUu returns a boolean if a field has been set.

### SetParamOverUuNil

`func (o *ServiceParameterDataPatch) SetParamOverUuNil(b bool)`

 SetParamOverUuNil sets the value for ParamOverUu to be an explicit nil

### UnsetParamOverUu
`func (o *ServiceParameterDataPatch) UnsetParamOverUu()`

UnsetParamOverUu ensures that no value is present for ParamOverUu, not even an explicit nil
### GetParamForProSeDd

`func (o *ServiceParameterDataPatch) GetParamForProSeDd() string`

GetParamForProSeDd returns the ParamForProSeDd field if non-nil, zero value otherwise.

### GetParamForProSeDdOk

`func (o *ServiceParameterDataPatch) GetParamForProSeDdOk() (*string, bool)`

GetParamForProSeDdOk returns a tuple with the ParamForProSeDd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParamForProSeDd

`func (o *ServiceParameterDataPatch) SetParamForProSeDd(v string)`

SetParamForProSeDd sets ParamForProSeDd field to given value.

### HasParamForProSeDd

`func (o *ServiceParameterDataPatch) HasParamForProSeDd() bool`

HasParamForProSeDd returns a boolean if a field has been set.

### SetParamForProSeDdNil

`func (o *ServiceParameterDataPatch) SetParamForProSeDdNil(b bool)`

 SetParamForProSeDdNil sets the value for ParamForProSeDd to be an explicit nil

### UnsetParamForProSeDd
`func (o *ServiceParameterDataPatch) UnsetParamForProSeDd()`

UnsetParamForProSeDd ensures that no value is present for ParamForProSeDd, not even an explicit nil
### GetParamForProSeDc

`func (o *ServiceParameterDataPatch) GetParamForProSeDc() string`

GetParamForProSeDc returns the ParamForProSeDc field if non-nil, zero value otherwise.

### GetParamForProSeDcOk

`func (o *ServiceParameterDataPatch) GetParamForProSeDcOk() (*string, bool)`

GetParamForProSeDcOk returns a tuple with the ParamForProSeDc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParamForProSeDc

`func (o *ServiceParameterDataPatch) SetParamForProSeDc(v string)`

SetParamForProSeDc sets ParamForProSeDc field to given value.

### HasParamForProSeDc

`func (o *ServiceParameterDataPatch) HasParamForProSeDc() bool`

HasParamForProSeDc returns a boolean if a field has been set.

### SetParamForProSeDcNil

`func (o *ServiceParameterDataPatch) SetParamForProSeDcNil(b bool)`

 SetParamForProSeDcNil sets the value for ParamForProSeDc to be an explicit nil

### UnsetParamForProSeDc
`func (o *ServiceParameterDataPatch) UnsetParamForProSeDc()`

UnsetParamForProSeDc ensures that no value is present for ParamForProSeDc, not even an explicit nil
### GetParamForProSeU2NRelUe

`func (o *ServiceParameterDataPatch) GetParamForProSeU2NRelUe() string`

GetParamForProSeU2NRelUe returns the ParamForProSeU2NRelUe field if non-nil, zero value otherwise.

### GetParamForProSeU2NRelUeOk

`func (o *ServiceParameterDataPatch) GetParamForProSeU2NRelUeOk() (*string, bool)`

GetParamForProSeU2NRelUeOk returns a tuple with the ParamForProSeU2NRelUe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParamForProSeU2NRelUe

`func (o *ServiceParameterDataPatch) SetParamForProSeU2NRelUe(v string)`

SetParamForProSeU2NRelUe sets ParamForProSeU2NRelUe field to given value.

### HasParamForProSeU2NRelUe

`func (o *ServiceParameterDataPatch) HasParamForProSeU2NRelUe() bool`

HasParamForProSeU2NRelUe returns a boolean if a field has been set.

### SetParamForProSeU2NRelUeNil

`func (o *ServiceParameterDataPatch) SetParamForProSeU2NRelUeNil(b bool)`

 SetParamForProSeU2NRelUeNil sets the value for ParamForProSeU2NRelUe to be an explicit nil

### UnsetParamForProSeU2NRelUe
`func (o *ServiceParameterDataPatch) UnsetParamForProSeU2NRelUe()`

UnsetParamForProSeU2NRelUe ensures that no value is present for ParamForProSeU2NRelUe, not even an explicit nil
### GetParamForProSeRemUe

`func (o *ServiceParameterDataPatch) GetParamForProSeRemUe() string`

GetParamForProSeRemUe returns the ParamForProSeRemUe field if non-nil, zero value otherwise.

### GetParamForProSeRemUeOk

`func (o *ServiceParameterDataPatch) GetParamForProSeRemUeOk() (*string, bool)`

GetParamForProSeRemUeOk returns a tuple with the ParamForProSeRemUe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParamForProSeRemUe

`func (o *ServiceParameterDataPatch) SetParamForProSeRemUe(v string)`

SetParamForProSeRemUe sets ParamForProSeRemUe field to given value.

### HasParamForProSeRemUe

`func (o *ServiceParameterDataPatch) HasParamForProSeRemUe() bool`

HasParamForProSeRemUe returns a boolean if a field has been set.

### SetParamForProSeRemUeNil

`func (o *ServiceParameterDataPatch) SetParamForProSeRemUeNil(b bool)`

 SetParamForProSeRemUeNil sets the value for ParamForProSeRemUe to be an explicit nil

### UnsetParamForProSeRemUe
`func (o *ServiceParameterDataPatch) UnsetParamForProSeRemUe()`

UnsetParamForProSeRemUe ensures that no value is present for ParamForProSeRemUe, not even an explicit nil
### GetUrspGuidance

`func (o *ServiceParameterDataPatch) GetUrspGuidance() []UrspRuleRequest`

GetUrspGuidance returns the UrspGuidance field if non-nil, zero value otherwise.

### GetUrspGuidanceOk

`func (o *ServiceParameterDataPatch) GetUrspGuidanceOk() (*[]UrspRuleRequest, bool)`

GetUrspGuidanceOk returns a tuple with the UrspGuidance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrspGuidance

`func (o *ServiceParameterDataPatch) SetUrspGuidance(v []UrspRuleRequest)`

SetUrspGuidance sets UrspGuidance field to given value.

### HasUrspGuidance

`func (o *ServiceParameterDataPatch) HasUrspGuidance() bool`

HasUrspGuidance returns a boolean if a field has been set.

### GetSubNotifEvents

`func (o *ServiceParameterDataPatch) GetSubNotifEvents() []Event`

GetSubNotifEvents returns the SubNotifEvents field if non-nil, zero value otherwise.

### GetSubNotifEventsOk

`func (o *ServiceParameterDataPatch) GetSubNotifEventsOk() (*[]Event, bool)`

GetSubNotifEventsOk returns a tuple with the SubNotifEvents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubNotifEvents

`func (o *ServiceParameterDataPatch) SetSubNotifEvents(v []Event)`

SetSubNotifEvents sets SubNotifEvents field to given value.

### HasSubNotifEvents

`func (o *ServiceParameterDataPatch) HasSubNotifEvents() bool`

HasSubNotifEvents returns a boolean if a field has been set.

### SetSubNotifEventsNil

`func (o *ServiceParameterDataPatch) SetSubNotifEventsNil(b bool)`

 SetSubNotifEventsNil sets the value for SubNotifEvents to be an explicit nil

### UnsetSubNotifEvents
`func (o *ServiceParameterDataPatch) UnsetSubNotifEvents()`

UnsetSubNotifEvents ensures that no value is present for SubNotifEvents, not even an explicit nil
### GetNotificationDestination

`func (o *ServiceParameterDataPatch) GetNotificationDestination() string`

GetNotificationDestination returns the NotificationDestination field if non-nil, zero value otherwise.

### GetNotificationDestinationOk

`func (o *ServiceParameterDataPatch) GetNotificationDestinationOk() (*string, bool)`

GetNotificationDestinationOk returns a tuple with the NotificationDestination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationDestination

`func (o *ServiceParameterDataPatch) SetNotificationDestination(v string)`

SetNotificationDestination sets NotificationDestination field to given value.

### HasNotificationDestination

`func (o *ServiceParameterDataPatch) HasNotificationDestination() bool`

HasNotificationDestination returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


