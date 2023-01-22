# SecurityNotification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiInvokerId** | **string** | String identifying the API invoker assigned by the CAPIF core function. | 
**AefId** | Pointer to **string** | String identifying the AEF. | [optional] 
**ApiIds** | **[]string** | Identifier of the service API | 
**Cause** | [**Cause**](Cause.md) |  | 

## Methods

### NewSecurityNotification

`func NewSecurityNotification(apiInvokerId string, apiIds []string, cause Cause, ) *SecurityNotification`

NewSecurityNotification instantiates a new SecurityNotification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecurityNotificationWithDefaults

`func NewSecurityNotificationWithDefaults() *SecurityNotification`

NewSecurityNotificationWithDefaults instantiates a new SecurityNotification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiInvokerId

`func (o *SecurityNotification) GetApiInvokerId() string`

GetApiInvokerId returns the ApiInvokerId field if non-nil, zero value otherwise.

### GetApiInvokerIdOk

`func (o *SecurityNotification) GetApiInvokerIdOk() (*string, bool)`

GetApiInvokerIdOk returns a tuple with the ApiInvokerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiInvokerId

`func (o *SecurityNotification) SetApiInvokerId(v string)`

SetApiInvokerId sets ApiInvokerId field to given value.


### GetAefId

`func (o *SecurityNotification) GetAefId() string`

GetAefId returns the AefId field if non-nil, zero value otherwise.

### GetAefIdOk

`func (o *SecurityNotification) GetAefIdOk() (*string, bool)`

GetAefIdOk returns a tuple with the AefId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAefId

`func (o *SecurityNotification) SetAefId(v string)`

SetAefId sets AefId field to given value.

### HasAefId

`func (o *SecurityNotification) HasAefId() bool`

HasAefId returns a boolean if a field has been set.

### GetApiIds

`func (o *SecurityNotification) GetApiIds() []string`

GetApiIds returns the ApiIds field if non-nil, zero value otherwise.

### GetApiIdsOk

`func (o *SecurityNotification) GetApiIdsOk() (*[]string, bool)`

GetApiIdsOk returns a tuple with the ApiIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiIds

`func (o *SecurityNotification) SetApiIds(v []string)`

SetApiIds sets ApiIds field to given value.


### GetCause

`func (o *SecurityNotification) GetCause() Cause`

GetCause returns the Cause field if non-nil, zero value otherwise.

### GetCauseOk

`func (o *SecurityNotification) GetCauseOk() (*Cause, bool)`

GetCauseOk returns a tuple with the Cause field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCause

`func (o *SecurityNotification) SetCause(v Cause)`

SetCause sets Cause field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


