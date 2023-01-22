# M1QoSSpecification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**QosReference** | Pointer to **string** |  | [optional] 
**MaxBtrUl** | Pointer to **string** | String representing a bit rate; the prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \&quot;K\&quot; is used to represent the standard symbol \&quot;k\&quot;.  | [optional] 
**MaxBtrDl** | Pointer to **string** | String representing a bit rate; the prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \&quot;K\&quot; is used to represent the standard symbol \&quot;k\&quot;.  | [optional] 
**MaxAuthBtrUl** | Pointer to **string** | String representing a bit rate; the prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \&quot;K\&quot; is used to represent the standard symbol \&quot;k\&quot;.  | [optional] 
**MaxAuthBtrDl** | Pointer to **string** | String representing a bit rate; the prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \&quot;K\&quot; is used to represent the standard symbol \&quot;k\&quot;.  | [optional] 
**DefPacketLossRateDl** | Pointer to **int32** |  | [optional] 
**DefPacketLossRateUl** | Pointer to **int32** |  | [optional] 

## Methods

### NewM1QoSSpecification

`func NewM1QoSSpecification() *M1QoSSpecification`

NewM1QoSSpecification instantiates a new M1QoSSpecification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewM1QoSSpecificationWithDefaults

`func NewM1QoSSpecificationWithDefaults() *M1QoSSpecification`

NewM1QoSSpecificationWithDefaults instantiates a new M1QoSSpecification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQosReference

`func (o *M1QoSSpecification) GetQosReference() string`

GetQosReference returns the QosReference field if non-nil, zero value otherwise.

### GetQosReferenceOk

`func (o *M1QoSSpecification) GetQosReferenceOk() (*string, bool)`

GetQosReferenceOk returns a tuple with the QosReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQosReference

`func (o *M1QoSSpecification) SetQosReference(v string)`

SetQosReference sets QosReference field to given value.

### HasQosReference

`func (o *M1QoSSpecification) HasQosReference() bool`

HasQosReference returns a boolean if a field has been set.

### GetMaxBtrUl

`func (o *M1QoSSpecification) GetMaxBtrUl() string`

GetMaxBtrUl returns the MaxBtrUl field if non-nil, zero value otherwise.

### GetMaxBtrUlOk

`func (o *M1QoSSpecification) GetMaxBtrUlOk() (*string, bool)`

GetMaxBtrUlOk returns a tuple with the MaxBtrUl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxBtrUl

`func (o *M1QoSSpecification) SetMaxBtrUl(v string)`

SetMaxBtrUl sets MaxBtrUl field to given value.

### HasMaxBtrUl

`func (o *M1QoSSpecification) HasMaxBtrUl() bool`

HasMaxBtrUl returns a boolean if a field has been set.

### GetMaxBtrDl

`func (o *M1QoSSpecification) GetMaxBtrDl() string`

GetMaxBtrDl returns the MaxBtrDl field if non-nil, zero value otherwise.

### GetMaxBtrDlOk

`func (o *M1QoSSpecification) GetMaxBtrDlOk() (*string, bool)`

GetMaxBtrDlOk returns a tuple with the MaxBtrDl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxBtrDl

`func (o *M1QoSSpecification) SetMaxBtrDl(v string)`

SetMaxBtrDl sets MaxBtrDl field to given value.

### HasMaxBtrDl

`func (o *M1QoSSpecification) HasMaxBtrDl() bool`

HasMaxBtrDl returns a boolean if a field has been set.

### GetMaxAuthBtrUl

`func (o *M1QoSSpecification) GetMaxAuthBtrUl() string`

GetMaxAuthBtrUl returns the MaxAuthBtrUl field if non-nil, zero value otherwise.

### GetMaxAuthBtrUlOk

`func (o *M1QoSSpecification) GetMaxAuthBtrUlOk() (*string, bool)`

GetMaxAuthBtrUlOk returns a tuple with the MaxAuthBtrUl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxAuthBtrUl

`func (o *M1QoSSpecification) SetMaxAuthBtrUl(v string)`

SetMaxAuthBtrUl sets MaxAuthBtrUl field to given value.

### HasMaxAuthBtrUl

`func (o *M1QoSSpecification) HasMaxAuthBtrUl() bool`

HasMaxAuthBtrUl returns a boolean if a field has been set.

### GetMaxAuthBtrDl

`func (o *M1QoSSpecification) GetMaxAuthBtrDl() string`

GetMaxAuthBtrDl returns the MaxAuthBtrDl field if non-nil, zero value otherwise.

### GetMaxAuthBtrDlOk

`func (o *M1QoSSpecification) GetMaxAuthBtrDlOk() (*string, bool)`

GetMaxAuthBtrDlOk returns a tuple with the MaxAuthBtrDl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxAuthBtrDl

`func (o *M1QoSSpecification) SetMaxAuthBtrDl(v string)`

SetMaxAuthBtrDl sets MaxAuthBtrDl field to given value.

### HasMaxAuthBtrDl

`func (o *M1QoSSpecification) HasMaxAuthBtrDl() bool`

HasMaxAuthBtrDl returns a boolean if a field has been set.

### GetDefPacketLossRateDl

`func (o *M1QoSSpecification) GetDefPacketLossRateDl() int32`

GetDefPacketLossRateDl returns the DefPacketLossRateDl field if non-nil, zero value otherwise.

### GetDefPacketLossRateDlOk

`func (o *M1QoSSpecification) GetDefPacketLossRateDlOk() (*int32, bool)`

GetDefPacketLossRateDlOk returns a tuple with the DefPacketLossRateDl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefPacketLossRateDl

`func (o *M1QoSSpecification) SetDefPacketLossRateDl(v int32)`

SetDefPacketLossRateDl sets DefPacketLossRateDl field to given value.

### HasDefPacketLossRateDl

`func (o *M1QoSSpecification) HasDefPacketLossRateDl() bool`

HasDefPacketLossRateDl returns a boolean if a field has been set.

### GetDefPacketLossRateUl

`func (o *M1QoSSpecification) GetDefPacketLossRateUl() int32`

GetDefPacketLossRateUl returns the DefPacketLossRateUl field if non-nil, zero value otherwise.

### GetDefPacketLossRateUlOk

`func (o *M1QoSSpecification) GetDefPacketLossRateUlOk() (*int32, bool)`

GetDefPacketLossRateUlOk returns a tuple with the DefPacketLossRateUl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefPacketLossRateUl

`func (o *M1QoSSpecification) SetDefPacketLossRateUl(v int32)`

SetDefPacketLossRateUl sets DefPacketLossRateUl field to given value.

### HasDefPacketLossRateUl

`func (o *M1QoSSpecification) HasDefPacketLossRateUl() bool`

HasDefPacketLossRateUl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


