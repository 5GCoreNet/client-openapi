# ServiceParameterDataPatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ParamOverPc5** | Pointer to **string** | Represents configuration parameters for V2X communications over PC5 reference point.  | [optional] 
**ParamOverUu** | Pointer to **string** | Represents configuration parameters for V2X communications over Uu reference point.  | [optional] 
**ParamForProSeDd** | Pointer to **string** | Represents the service parameters for 5G ProSe direct discovery. | [optional] 
**ParamForProSeDc** | Pointer to **string** | Represents the service parameters for 5G ProSe direct communications. | [optional] 
**ParamForProSeU2NRelUe** | Pointer to **string** | Represents the service parameters for 5G ProSe UE-to-network relay UE. | [optional] 
**ParamForProSeRemUe** | Pointer to **string** | Represents the service parameters for 5G ProSe Remate UE. | [optional] 
**UrspInfluence** | Pointer to [**[]UrspRuleRequest**](UrspRuleRequest.md) | Contains the service parameter used to influence the URSP. | [optional] 
**DeliveryEvents** | Pointer to [**[]Event**](Event.md) | Contains the outcome of the UE Policy Delivery. | [optional] 
**PolicDelivNotifUri** | Pointer to **string** | String providing an URI formatted according to RFC 3986. | [optional] 

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

### GetUrspInfluence

`func (o *ServiceParameterDataPatch) GetUrspInfluence() []UrspRuleRequest`

GetUrspInfluence returns the UrspInfluence field if non-nil, zero value otherwise.

### GetUrspInfluenceOk

`func (o *ServiceParameterDataPatch) GetUrspInfluenceOk() (*[]UrspRuleRequest, bool)`

GetUrspInfluenceOk returns a tuple with the UrspInfluence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrspInfluence

`func (o *ServiceParameterDataPatch) SetUrspInfluence(v []UrspRuleRequest)`

SetUrspInfluence sets UrspInfluence field to given value.

### HasUrspInfluence

`func (o *ServiceParameterDataPatch) HasUrspInfluence() bool`

HasUrspInfluence returns a boolean if a field has been set.

### GetDeliveryEvents

`func (o *ServiceParameterDataPatch) GetDeliveryEvents() []Event`

GetDeliveryEvents returns the DeliveryEvents field if non-nil, zero value otherwise.

### GetDeliveryEventsOk

`func (o *ServiceParameterDataPatch) GetDeliveryEventsOk() (*[]Event, bool)`

GetDeliveryEventsOk returns a tuple with the DeliveryEvents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryEvents

`func (o *ServiceParameterDataPatch) SetDeliveryEvents(v []Event)`

SetDeliveryEvents sets DeliveryEvents field to given value.

### HasDeliveryEvents

`func (o *ServiceParameterDataPatch) HasDeliveryEvents() bool`

HasDeliveryEvents returns a boolean if a field has been set.

### GetPolicDelivNotifUri

`func (o *ServiceParameterDataPatch) GetPolicDelivNotifUri() string`

GetPolicDelivNotifUri returns the PolicDelivNotifUri field if non-nil, zero value otherwise.

### GetPolicDelivNotifUriOk

`func (o *ServiceParameterDataPatch) GetPolicDelivNotifUriOk() (*string, bool)`

GetPolicDelivNotifUriOk returns a tuple with the PolicDelivNotifUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicDelivNotifUri

`func (o *ServiceParameterDataPatch) SetPolicDelivNotifUri(v string)`

SetPolicDelivNotifUri sets PolicDelivNotifUri field to given value.

### HasPolicDelivNotifUri

`func (o *ServiceParameterDataPatch) HasPolicDelivNotifUri() bool`

HasPolicDelivNotifUri returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


