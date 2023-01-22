# EDNConnectionInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DNN** | Pointer to **string** |  | [optional] 
**EDNServiceArea** | Pointer to [**ServingLocation**](ServingLocation.md) |  | [optional] 

## Methods

### NewEDNConnectionInfo

`func NewEDNConnectionInfo() *EDNConnectionInfo`

NewEDNConnectionInfo instantiates a new EDNConnectionInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEDNConnectionInfoWithDefaults

`func NewEDNConnectionInfoWithDefaults() *EDNConnectionInfo`

NewEDNConnectionInfoWithDefaults instantiates a new EDNConnectionInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDNN

`func (o *EDNConnectionInfo) GetDNN() string`

GetDNN returns the DNN field if non-nil, zero value otherwise.

### GetDNNOk

`func (o *EDNConnectionInfo) GetDNNOk() (*string, bool)`

GetDNNOk returns a tuple with the DNN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDNN

`func (o *EDNConnectionInfo) SetDNN(v string)`

SetDNN sets DNN field to given value.

### HasDNN

`func (o *EDNConnectionInfo) HasDNN() bool`

HasDNN returns a boolean if a field has been set.

### GetEDNServiceArea

`func (o *EDNConnectionInfo) GetEDNServiceArea() ServingLocation`

GetEDNServiceArea returns the EDNServiceArea field if non-nil, zero value otherwise.

### GetEDNServiceAreaOk

`func (o *EDNConnectionInfo) GetEDNServiceAreaOk() (*ServingLocation, bool)`

GetEDNServiceAreaOk returns a tuple with the EDNServiceArea field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEDNServiceArea

`func (o *EDNConnectionInfo) SetEDNServiceArea(v ServingLocation)`

SetEDNServiceArea sets EDNServiceArea field to given value.

### HasEDNServiceArea

`func (o *EDNConnectionInfo) HasEDNServiceArea() bool`

HasEDNServiceArea returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


