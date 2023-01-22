# RecipientAddress

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RecipientAddressInfo** | Pointer to [**SMAddressInfo**](SMAddressInfo.md) |  | [optional] 
**SMaddresseeType** | Pointer to [**SMAddresseeType**](SMAddresseeType.md) |  | [optional] 

## Methods

### NewRecipientAddress

`func NewRecipientAddress() *RecipientAddress`

NewRecipientAddress instantiates a new RecipientAddress object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecipientAddressWithDefaults

`func NewRecipientAddressWithDefaults() *RecipientAddress`

NewRecipientAddressWithDefaults instantiates a new RecipientAddress object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRecipientAddressInfo

`func (o *RecipientAddress) GetRecipientAddressInfo() SMAddressInfo`

GetRecipientAddressInfo returns the RecipientAddressInfo field if non-nil, zero value otherwise.

### GetRecipientAddressInfoOk

`func (o *RecipientAddress) GetRecipientAddressInfoOk() (*SMAddressInfo, bool)`

GetRecipientAddressInfoOk returns a tuple with the RecipientAddressInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipientAddressInfo

`func (o *RecipientAddress) SetRecipientAddressInfo(v SMAddressInfo)`

SetRecipientAddressInfo sets RecipientAddressInfo field to given value.

### HasRecipientAddressInfo

`func (o *RecipientAddress) HasRecipientAddressInfo() bool`

HasRecipientAddressInfo returns a boolean if a field has been set.

### GetSMaddresseeType

`func (o *RecipientAddress) GetSMaddresseeType() SMAddresseeType`

GetSMaddresseeType returns the SMaddresseeType field if non-nil, zero value otherwise.

### GetSMaddresseeTypeOk

`func (o *RecipientAddress) GetSMaddresseeTypeOk() (*SMAddresseeType, bool)`

GetSMaddresseeTypeOk returns a tuple with the SMaddresseeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSMaddresseeType

`func (o *RecipientAddress) SetSMaddresseeType(v SMAddresseeType)`

SetSMaddresseeType sets SMaddresseeType field to given value.

### HasSMaddresseeType

`func (o *RecipientAddress) HasSMaddresseeType() bool`

HasSMaddresseeType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


