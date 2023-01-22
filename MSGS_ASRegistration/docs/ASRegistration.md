# ASRegistration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AsSvcId** | **string** |  | 
**AppId** | Pointer to **string** |  | [optional] 
**TargetUri** | Pointer to **string** | String providing an URI formatted according to RFC 3986. | [optional] 
**AsProf** | Pointer to [**ASProfile**](ASProfile.md) |  | [optional] 

## Methods

### NewASRegistration

`func NewASRegistration(asSvcId string, ) *ASRegistration`

NewASRegistration instantiates a new ASRegistration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewASRegistrationWithDefaults

`func NewASRegistrationWithDefaults() *ASRegistration`

NewASRegistrationWithDefaults instantiates a new ASRegistration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAsSvcId

`func (o *ASRegistration) GetAsSvcId() string`

GetAsSvcId returns the AsSvcId field if non-nil, zero value otherwise.

### GetAsSvcIdOk

`func (o *ASRegistration) GetAsSvcIdOk() (*string, bool)`

GetAsSvcIdOk returns a tuple with the AsSvcId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsSvcId

`func (o *ASRegistration) SetAsSvcId(v string)`

SetAsSvcId sets AsSvcId field to given value.


### GetAppId

`func (o *ASRegistration) GetAppId() string`

GetAppId returns the AppId field if non-nil, zero value otherwise.

### GetAppIdOk

`func (o *ASRegistration) GetAppIdOk() (*string, bool)`

GetAppIdOk returns a tuple with the AppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppId

`func (o *ASRegistration) SetAppId(v string)`

SetAppId sets AppId field to given value.

### HasAppId

`func (o *ASRegistration) HasAppId() bool`

HasAppId returns a boolean if a field has been set.

### GetTargetUri

`func (o *ASRegistration) GetTargetUri() string`

GetTargetUri returns the TargetUri field if non-nil, zero value otherwise.

### GetTargetUriOk

`func (o *ASRegistration) GetTargetUriOk() (*string, bool)`

GetTargetUriOk returns a tuple with the TargetUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetUri

`func (o *ASRegistration) SetTargetUri(v string)`

SetTargetUri sets TargetUri field to given value.

### HasTargetUri

`func (o *ASRegistration) HasTargetUri() bool`

HasTargetUri returns a boolean if a field has been set.

### GetAsProf

`func (o *ASRegistration) GetAsProf() ASProfile`

GetAsProf returns the AsProf field if non-nil, zero value otherwise.

### GetAsProfOk

`func (o *ASRegistration) GetAsProfOk() (*ASProfile, bool)`

GetAsProfOk returns a tuple with the AsProf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsProf

`func (o *ASRegistration) SetAsProf(v ASProfile)`

SetAsProf sets AsProf field to given value.

### HasAsProf

`func (o *ASRegistration) HasAsProf() bool`

HasAsProf returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


