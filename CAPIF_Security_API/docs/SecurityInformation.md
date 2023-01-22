# SecurityInformation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InterfaceDetails** | Pointer to [**InterfaceDescription**](InterfaceDescription.md) |  | [optional] 
**AefId** | Pointer to **string** | Identifier of the API exposing function | [optional] 
**ApiId** | Pointer to **string** | API identifier | [optional] 
**PrefSecurityMethods** | [**[]SecurityMethod**](SecurityMethod.md) | Security methods preferred by the API invoker for the API interface. | 
**SelSecurityMethod** | Pointer to [**SecurityMethod**](SecurityMethod.md) |  | [optional] 
**AuthenticationInfo** | Pointer to **string** | Authentication related information | [optional] 
**AuthorizationInfo** | Pointer to **string** | Authorization related information | [optional] 

## Methods

### NewSecurityInformation

`func NewSecurityInformation(prefSecurityMethods []SecurityMethod, ) *SecurityInformation`

NewSecurityInformation instantiates a new SecurityInformation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecurityInformationWithDefaults

`func NewSecurityInformationWithDefaults() *SecurityInformation`

NewSecurityInformationWithDefaults instantiates a new SecurityInformation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInterfaceDetails

`func (o *SecurityInformation) GetInterfaceDetails() InterfaceDescription`

GetInterfaceDetails returns the InterfaceDetails field if non-nil, zero value otherwise.

### GetInterfaceDetailsOk

`func (o *SecurityInformation) GetInterfaceDetailsOk() (*InterfaceDescription, bool)`

GetInterfaceDetailsOk returns a tuple with the InterfaceDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceDetails

`func (o *SecurityInformation) SetInterfaceDetails(v InterfaceDescription)`

SetInterfaceDetails sets InterfaceDetails field to given value.

### HasInterfaceDetails

`func (o *SecurityInformation) HasInterfaceDetails() bool`

HasInterfaceDetails returns a boolean if a field has been set.

### GetAefId

`func (o *SecurityInformation) GetAefId() string`

GetAefId returns the AefId field if non-nil, zero value otherwise.

### GetAefIdOk

`func (o *SecurityInformation) GetAefIdOk() (*string, bool)`

GetAefIdOk returns a tuple with the AefId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAefId

`func (o *SecurityInformation) SetAefId(v string)`

SetAefId sets AefId field to given value.

### HasAefId

`func (o *SecurityInformation) HasAefId() bool`

HasAefId returns a boolean if a field has been set.

### GetApiId

`func (o *SecurityInformation) GetApiId() string`

GetApiId returns the ApiId field if non-nil, zero value otherwise.

### GetApiIdOk

`func (o *SecurityInformation) GetApiIdOk() (*string, bool)`

GetApiIdOk returns a tuple with the ApiId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiId

`func (o *SecurityInformation) SetApiId(v string)`

SetApiId sets ApiId field to given value.

### HasApiId

`func (o *SecurityInformation) HasApiId() bool`

HasApiId returns a boolean if a field has been set.

### GetPrefSecurityMethods

`func (o *SecurityInformation) GetPrefSecurityMethods() []SecurityMethod`

GetPrefSecurityMethods returns the PrefSecurityMethods field if non-nil, zero value otherwise.

### GetPrefSecurityMethodsOk

`func (o *SecurityInformation) GetPrefSecurityMethodsOk() (*[]SecurityMethod, bool)`

GetPrefSecurityMethodsOk returns a tuple with the PrefSecurityMethods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefSecurityMethods

`func (o *SecurityInformation) SetPrefSecurityMethods(v []SecurityMethod)`

SetPrefSecurityMethods sets PrefSecurityMethods field to given value.


### GetSelSecurityMethod

`func (o *SecurityInformation) GetSelSecurityMethod() SecurityMethod`

GetSelSecurityMethod returns the SelSecurityMethod field if non-nil, zero value otherwise.

### GetSelSecurityMethodOk

`func (o *SecurityInformation) GetSelSecurityMethodOk() (*SecurityMethod, bool)`

GetSelSecurityMethodOk returns a tuple with the SelSecurityMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelSecurityMethod

`func (o *SecurityInformation) SetSelSecurityMethod(v SecurityMethod)`

SetSelSecurityMethod sets SelSecurityMethod field to given value.

### HasSelSecurityMethod

`func (o *SecurityInformation) HasSelSecurityMethod() bool`

HasSelSecurityMethod returns a boolean if a field has been set.

### GetAuthenticationInfo

`func (o *SecurityInformation) GetAuthenticationInfo() string`

GetAuthenticationInfo returns the AuthenticationInfo field if non-nil, zero value otherwise.

### GetAuthenticationInfoOk

`func (o *SecurityInformation) GetAuthenticationInfoOk() (*string, bool)`

GetAuthenticationInfoOk returns a tuple with the AuthenticationInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationInfo

`func (o *SecurityInformation) SetAuthenticationInfo(v string)`

SetAuthenticationInfo sets AuthenticationInfo field to given value.

### HasAuthenticationInfo

`func (o *SecurityInformation) HasAuthenticationInfo() bool`

HasAuthenticationInfo returns a boolean if a field has been set.

### GetAuthorizationInfo

`func (o *SecurityInformation) GetAuthorizationInfo() string`

GetAuthorizationInfo returns the AuthorizationInfo field if non-nil, zero value otherwise.

### GetAuthorizationInfoOk

`func (o *SecurityInformation) GetAuthorizationInfoOk() (*string, bool)`

GetAuthorizationInfoOk returns a tuple with the AuthorizationInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationInfo

`func (o *SecurityInformation) SetAuthorizationInfo(v string)`

SetAuthorizationInfo sets AuthorizationInfo field to given value.

### HasAuthorizationInfo

`func (o *SecurityInformation) HasAuthorizationInfo() bool`

HasAuthorizationInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


