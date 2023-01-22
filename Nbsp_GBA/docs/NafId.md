# NafId

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NafFqdn** | **string** | Fully Qualified Domain Name | 
**UaSecProtId** | **string** |  | 

## Methods

### NewNafId

`func NewNafId(nafFqdn string, uaSecProtId string, ) *NafId`

NewNafId instantiates a new NafId object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNafIdWithDefaults

`func NewNafIdWithDefaults() *NafId`

NewNafIdWithDefaults instantiates a new NafId object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNafFqdn

`func (o *NafId) GetNafFqdn() string`

GetNafFqdn returns the NafFqdn field if non-nil, zero value otherwise.

### GetNafFqdnOk

`func (o *NafId) GetNafFqdnOk() (*string, bool)`

GetNafFqdnOk returns a tuple with the NafFqdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNafFqdn

`func (o *NafId) SetNafFqdn(v string)`

SetNafFqdn sets NafFqdn field to given value.


### GetUaSecProtId

`func (o *NafId) GetUaSecProtId() string`

GetUaSecProtId returns the UaSecProtId field if non-nil, zero value otherwise.

### GetUaSecProtIdOk

`func (o *NafId) GetUaSecProtIdOk() (*string, bool)`

GetUaSecProtIdOk returns a tuple with the UaSecProtId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUaSecProtId

`func (o *NafId) SetUaSecProtId(v string)`

SetUaSecProtId sets UaSecProtId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


