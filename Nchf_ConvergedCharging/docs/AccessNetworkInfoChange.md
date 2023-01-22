# AccessNetworkInfoChange

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessNetworkInformation** | Pointer to **[]string** |  | [optional] 
**CellularNetworkInformation** | Pointer to **string** |  | [optional] 
**ChangeTime** | Pointer to **time.Time** | string with format &#39;date-time&#39; as defined in OpenAPI. | [optional] 

## Methods

### NewAccessNetworkInfoChange

`func NewAccessNetworkInfoChange() *AccessNetworkInfoChange`

NewAccessNetworkInfoChange instantiates a new AccessNetworkInfoChange object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccessNetworkInfoChangeWithDefaults

`func NewAccessNetworkInfoChangeWithDefaults() *AccessNetworkInfoChange`

NewAccessNetworkInfoChangeWithDefaults instantiates a new AccessNetworkInfoChange object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessNetworkInformation

`func (o *AccessNetworkInfoChange) GetAccessNetworkInformation() []string`

GetAccessNetworkInformation returns the AccessNetworkInformation field if non-nil, zero value otherwise.

### GetAccessNetworkInformationOk

`func (o *AccessNetworkInfoChange) GetAccessNetworkInformationOk() (*[]string, bool)`

GetAccessNetworkInformationOk returns a tuple with the AccessNetworkInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessNetworkInformation

`func (o *AccessNetworkInfoChange) SetAccessNetworkInformation(v []string)`

SetAccessNetworkInformation sets AccessNetworkInformation field to given value.

### HasAccessNetworkInformation

`func (o *AccessNetworkInfoChange) HasAccessNetworkInformation() bool`

HasAccessNetworkInformation returns a boolean if a field has been set.

### GetCellularNetworkInformation

`func (o *AccessNetworkInfoChange) GetCellularNetworkInformation() string`

GetCellularNetworkInformation returns the CellularNetworkInformation field if non-nil, zero value otherwise.

### GetCellularNetworkInformationOk

`func (o *AccessNetworkInfoChange) GetCellularNetworkInformationOk() (*string, bool)`

GetCellularNetworkInformationOk returns a tuple with the CellularNetworkInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCellularNetworkInformation

`func (o *AccessNetworkInfoChange) SetCellularNetworkInformation(v string)`

SetCellularNetworkInformation sets CellularNetworkInformation field to given value.

### HasCellularNetworkInformation

`func (o *AccessNetworkInfoChange) HasCellularNetworkInformation() bool`

HasCellularNetworkInformation returns a boolean if a field has been set.

### GetChangeTime

`func (o *AccessNetworkInfoChange) GetChangeTime() time.Time`

GetChangeTime returns the ChangeTime field if non-nil, zero value otherwise.

### GetChangeTimeOk

`func (o *AccessNetworkInfoChange) GetChangeTimeOk() (*time.Time, bool)`

GetChangeTimeOk returns a tuple with the ChangeTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeTime

`func (o *AccessNetworkInfoChange) SetChangeTime(v time.Time)`

SetChangeTime sets ChangeTime field to given value.

### HasChangeTime

`func (o *AccessNetworkInfoChange) HasChangeTime() bool`

HasChangeTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


