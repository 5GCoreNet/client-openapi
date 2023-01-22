# ChargingSpecification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SponId** | Pointer to **string** |  | [optional] 
**SponStatus** | Pointer to [**SponsoringStatus**](SponsoringStatus.md) |  | [optional] 
**Gpsi** | Pointer to **[]string** |  | [optional] 

## Methods

### NewChargingSpecification

`func NewChargingSpecification() *ChargingSpecification`

NewChargingSpecification instantiates a new ChargingSpecification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChargingSpecificationWithDefaults

`func NewChargingSpecificationWithDefaults() *ChargingSpecification`

NewChargingSpecificationWithDefaults instantiates a new ChargingSpecification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSponId

`func (o *ChargingSpecification) GetSponId() string`

GetSponId returns the SponId field if non-nil, zero value otherwise.

### GetSponIdOk

`func (o *ChargingSpecification) GetSponIdOk() (*string, bool)`

GetSponIdOk returns a tuple with the SponId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSponId

`func (o *ChargingSpecification) SetSponId(v string)`

SetSponId sets SponId field to given value.

### HasSponId

`func (o *ChargingSpecification) HasSponId() bool`

HasSponId returns a boolean if a field has been set.

### GetSponStatus

`func (o *ChargingSpecification) GetSponStatus() SponsoringStatus`

GetSponStatus returns the SponStatus field if non-nil, zero value otherwise.

### GetSponStatusOk

`func (o *ChargingSpecification) GetSponStatusOk() (*SponsoringStatus, bool)`

GetSponStatusOk returns a tuple with the SponStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSponStatus

`func (o *ChargingSpecification) SetSponStatus(v SponsoringStatus)`

SetSponStatus sets SponStatus field to given value.

### HasSponStatus

`func (o *ChargingSpecification) HasSponStatus() bool`

HasSponStatus returns a boolean if a field has been set.

### GetGpsi

`func (o *ChargingSpecification) GetGpsi() []string`

GetGpsi returns the Gpsi field if non-nil, zero value otherwise.

### GetGpsiOk

`func (o *ChargingSpecification) GetGpsiOk() (*[]string, bool)`

GetGpsiOk returns a tuple with the Gpsi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpsi

`func (o *ChargingSpecification) SetGpsi(v []string)`

SetGpsi sets Gpsi field to given value.

### HasGpsi

`func (o *ChargingSpecification) HasGpsi() bool`

HasGpsi returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


