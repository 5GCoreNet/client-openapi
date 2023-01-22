# ScscfRestorationInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserName** | Pointer to **string** | IMS Private Identity of the UE | [optional] 
**RestorationInfo** | Pointer to [**[]RestorationInfo**](RestorationInfo.md) |  | [optional] 
**RegistrationTimeOut** | Pointer to **time.Time** | string with format &#39;date-time&#39; as defined in OpenAPI. | [optional] 
**SipAuthenticationScheme** | Pointer to [**SipAuthenticationScheme**](SipAuthenticationScheme.md) |  | [optional] 

## Methods

### NewScscfRestorationInfo

`func NewScscfRestorationInfo() *ScscfRestorationInfo`

NewScscfRestorationInfo instantiates a new ScscfRestorationInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScscfRestorationInfoWithDefaults

`func NewScscfRestorationInfoWithDefaults() *ScscfRestorationInfo`

NewScscfRestorationInfoWithDefaults instantiates a new ScscfRestorationInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserName

`func (o *ScscfRestorationInfo) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *ScscfRestorationInfo) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *ScscfRestorationInfo) SetUserName(v string)`

SetUserName sets UserName field to given value.

### HasUserName

`func (o *ScscfRestorationInfo) HasUserName() bool`

HasUserName returns a boolean if a field has been set.

### GetRestorationInfo

`func (o *ScscfRestorationInfo) GetRestorationInfo() []RestorationInfo`

GetRestorationInfo returns the RestorationInfo field if non-nil, zero value otherwise.

### GetRestorationInfoOk

`func (o *ScscfRestorationInfo) GetRestorationInfoOk() (*[]RestorationInfo, bool)`

GetRestorationInfoOk returns a tuple with the RestorationInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestorationInfo

`func (o *ScscfRestorationInfo) SetRestorationInfo(v []RestorationInfo)`

SetRestorationInfo sets RestorationInfo field to given value.

### HasRestorationInfo

`func (o *ScscfRestorationInfo) HasRestorationInfo() bool`

HasRestorationInfo returns a boolean if a field has been set.

### GetRegistrationTimeOut

`func (o *ScscfRestorationInfo) GetRegistrationTimeOut() time.Time`

GetRegistrationTimeOut returns the RegistrationTimeOut field if non-nil, zero value otherwise.

### GetRegistrationTimeOutOk

`func (o *ScscfRestorationInfo) GetRegistrationTimeOutOk() (*time.Time, bool)`

GetRegistrationTimeOutOk returns a tuple with the RegistrationTimeOut field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistrationTimeOut

`func (o *ScscfRestorationInfo) SetRegistrationTimeOut(v time.Time)`

SetRegistrationTimeOut sets RegistrationTimeOut field to given value.

### HasRegistrationTimeOut

`func (o *ScscfRestorationInfo) HasRegistrationTimeOut() bool`

HasRegistrationTimeOut returns a boolean if a field has been set.

### GetSipAuthenticationScheme

`func (o *ScscfRestorationInfo) GetSipAuthenticationScheme() SipAuthenticationScheme`

GetSipAuthenticationScheme returns the SipAuthenticationScheme field if non-nil, zero value otherwise.

### GetSipAuthenticationSchemeOk

`func (o *ScscfRestorationInfo) GetSipAuthenticationSchemeOk() (*SipAuthenticationScheme, bool)`

GetSipAuthenticationSchemeOk returns a tuple with the SipAuthenticationScheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSipAuthenticationScheme

`func (o *ScscfRestorationInfo) SetSipAuthenticationScheme(v SipAuthenticationScheme)`

SetSipAuthenticationScheme sets SipAuthenticationScheme field to given value.

### HasSipAuthenticationScheme

`func (o *ScscfRestorationInfo) HasSipAuthenticationScheme() bool`

HasSipAuthenticationScheme returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


