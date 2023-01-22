# Positioning

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServAttrCom** | Pointer to [**ServAttrCom**](ServAttrCom.md) |  | [optional] 
**Availability** | Pointer to **[]string** |  | [optional] 
**Predictionfrequency** | Pointer to [**Predictionfrequency**](Predictionfrequency.md) |  | [optional] 
**Accuracy** | Pointer to **float32** |  | [optional] 

## Methods

### NewPositioning

`func NewPositioning() *Positioning`

NewPositioning instantiates a new Positioning object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPositioningWithDefaults

`func NewPositioningWithDefaults() *Positioning`

NewPositioningWithDefaults instantiates a new Positioning object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServAttrCom

`func (o *Positioning) GetServAttrCom() ServAttrCom`

GetServAttrCom returns the ServAttrCom field if non-nil, zero value otherwise.

### GetServAttrComOk

`func (o *Positioning) GetServAttrComOk() (*ServAttrCom, bool)`

GetServAttrComOk returns a tuple with the ServAttrCom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServAttrCom

`func (o *Positioning) SetServAttrCom(v ServAttrCom)`

SetServAttrCom sets ServAttrCom field to given value.

### HasServAttrCom

`func (o *Positioning) HasServAttrCom() bool`

HasServAttrCom returns a boolean if a field has been set.

### GetAvailability

`func (o *Positioning) GetAvailability() []string`

GetAvailability returns the Availability field if non-nil, zero value otherwise.

### GetAvailabilityOk

`func (o *Positioning) GetAvailabilityOk() (*[]string, bool)`

GetAvailabilityOk returns a tuple with the Availability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailability

`func (o *Positioning) SetAvailability(v []string)`

SetAvailability sets Availability field to given value.

### HasAvailability

`func (o *Positioning) HasAvailability() bool`

HasAvailability returns a boolean if a field has been set.

### GetPredictionfrequency

`func (o *Positioning) GetPredictionfrequency() Predictionfrequency`

GetPredictionfrequency returns the Predictionfrequency field if non-nil, zero value otherwise.

### GetPredictionfrequencyOk

`func (o *Positioning) GetPredictionfrequencyOk() (*Predictionfrequency, bool)`

GetPredictionfrequencyOk returns a tuple with the Predictionfrequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPredictionfrequency

`func (o *Positioning) SetPredictionfrequency(v Predictionfrequency)`

SetPredictionfrequency sets Predictionfrequency field to given value.

### HasPredictionfrequency

`func (o *Positioning) HasPredictionfrequency() bool`

HasPredictionfrequency returns a boolean if a field has been set.

### GetAccuracy

`func (o *Positioning) GetAccuracy() float32`

GetAccuracy returns the Accuracy field if non-nil, zero value otherwise.

### GetAccuracyOk

`func (o *Positioning) GetAccuracyOk() (*float32, bool)`

GetAccuracyOk returns a tuple with the Accuracy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccuracy

`func (o *Positioning) SetAccuracy(v float32)`

SetAccuracy sets Accuracy field to given value.

### HasAccuracy

`func (o *Positioning) HasAccuracy() bool`

HasAccuracy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


