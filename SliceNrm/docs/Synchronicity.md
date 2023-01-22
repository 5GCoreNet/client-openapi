# Synchronicity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServAttrCom** | Pointer to [**ServAttrCom**](ServAttrCom.md) |  | [optional] 
**Availability** | Pointer to [**SynAvailability**](SynAvailability.md) |  | [optional] 
**Accuracy** | Pointer to **float32** |  | [optional] 

## Methods

### NewSynchronicity

`func NewSynchronicity() *Synchronicity`

NewSynchronicity instantiates a new Synchronicity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSynchronicityWithDefaults

`func NewSynchronicityWithDefaults() *Synchronicity`

NewSynchronicityWithDefaults instantiates a new Synchronicity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServAttrCom

`func (o *Synchronicity) GetServAttrCom() ServAttrCom`

GetServAttrCom returns the ServAttrCom field if non-nil, zero value otherwise.

### GetServAttrComOk

`func (o *Synchronicity) GetServAttrComOk() (*ServAttrCom, bool)`

GetServAttrComOk returns a tuple with the ServAttrCom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServAttrCom

`func (o *Synchronicity) SetServAttrCom(v ServAttrCom)`

SetServAttrCom sets ServAttrCom field to given value.

### HasServAttrCom

`func (o *Synchronicity) HasServAttrCom() bool`

HasServAttrCom returns a boolean if a field has been set.

### GetAvailability

`func (o *Synchronicity) GetAvailability() SynAvailability`

GetAvailability returns the Availability field if non-nil, zero value otherwise.

### GetAvailabilityOk

`func (o *Synchronicity) GetAvailabilityOk() (*SynAvailability, bool)`

GetAvailabilityOk returns a tuple with the Availability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailability

`func (o *Synchronicity) SetAvailability(v SynAvailability)`

SetAvailability sets Availability field to given value.

### HasAvailability

`func (o *Synchronicity) HasAvailability() bool`

HasAvailability returns a boolean if a field has been set.

### GetAccuracy

`func (o *Synchronicity) GetAccuracy() float32`

GetAccuracy returns the Accuracy field if non-nil, zero value otherwise.

### GetAccuracyOk

`func (o *Synchronicity) GetAccuracyOk() (*float32, bool)`

GetAccuracyOk returns a tuple with the Accuracy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccuracy

`func (o *Synchronicity) SetAccuracy(v float32)`

SetAccuracy sets Accuracy field to given value.

### HasAccuracy

`func (o *Synchronicity) HasAccuracy() bool`

HasAccuracy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


