# RedirectServer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RedirectAddressType** | [**RedirectAddressType**](RedirectAddressType.md) |  | 
**RedirectServerAddress** | **string** |  | 

## Methods

### NewRedirectServer

`func NewRedirectServer(redirectAddressType RedirectAddressType, redirectServerAddress string, ) *RedirectServer`

NewRedirectServer instantiates a new RedirectServer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRedirectServerWithDefaults

`func NewRedirectServerWithDefaults() *RedirectServer`

NewRedirectServerWithDefaults instantiates a new RedirectServer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRedirectAddressType

`func (o *RedirectServer) GetRedirectAddressType() RedirectAddressType`

GetRedirectAddressType returns the RedirectAddressType field if non-nil, zero value otherwise.

### GetRedirectAddressTypeOk

`func (o *RedirectServer) GetRedirectAddressTypeOk() (*RedirectAddressType, bool)`

GetRedirectAddressTypeOk returns a tuple with the RedirectAddressType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectAddressType

`func (o *RedirectServer) SetRedirectAddressType(v RedirectAddressType)`

SetRedirectAddressType sets RedirectAddressType field to given value.


### GetRedirectServerAddress

`func (o *RedirectServer) GetRedirectServerAddress() string`

GetRedirectServerAddress returns the RedirectServerAddress field if non-nil, zero value otherwise.

### GetRedirectServerAddressOk

`func (o *RedirectServer) GetRedirectServerAddressOk() (*string, bool)`

GetRedirectServerAddressOk returns a tuple with the RedirectServerAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectServerAddress

`func (o *RedirectServer) SetRedirectServerAddress(v string)`

SetRedirectServerAddress sets RedirectServerAddress field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


