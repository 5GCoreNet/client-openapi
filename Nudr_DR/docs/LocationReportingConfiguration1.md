# LocationReportingConfiguration1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrentLocation** | **bool** |  | 
**OneTime** | Pointer to **bool** |  | [optional] 
**Accuracy** | Pointer to [**LocationAccuracy**](LocationAccuracy.md) |  | [optional] 
**N3gppAccuracy** | Pointer to [**LocationAccuracy**](LocationAccuracy.md) |  | [optional] 

## Methods

### NewLocationReportingConfiguration1

`func NewLocationReportingConfiguration1(currentLocation bool, ) *LocationReportingConfiguration1`

NewLocationReportingConfiguration1 instantiates a new LocationReportingConfiguration1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLocationReportingConfiguration1WithDefaults

`func NewLocationReportingConfiguration1WithDefaults() *LocationReportingConfiguration1`

NewLocationReportingConfiguration1WithDefaults instantiates a new LocationReportingConfiguration1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrentLocation

`func (o *LocationReportingConfiguration1) GetCurrentLocation() bool`

GetCurrentLocation returns the CurrentLocation field if non-nil, zero value otherwise.

### GetCurrentLocationOk

`func (o *LocationReportingConfiguration1) GetCurrentLocationOk() (*bool, bool)`

GetCurrentLocationOk returns a tuple with the CurrentLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentLocation

`func (o *LocationReportingConfiguration1) SetCurrentLocation(v bool)`

SetCurrentLocation sets CurrentLocation field to given value.


### GetOneTime

`func (o *LocationReportingConfiguration1) GetOneTime() bool`

GetOneTime returns the OneTime field if non-nil, zero value otherwise.

### GetOneTimeOk

`func (o *LocationReportingConfiguration1) GetOneTimeOk() (*bool, bool)`

GetOneTimeOk returns a tuple with the OneTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOneTime

`func (o *LocationReportingConfiguration1) SetOneTime(v bool)`

SetOneTime sets OneTime field to given value.

### HasOneTime

`func (o *LocationReportingConfiguration1) HasOneTime() bool`

HasOneTime returns a boolean if a field has been set.

### GetAccuracy

`func (o *LocationReportingConfiguration1) GetAccuracy() LocationAccuracy`

GetAccuracy returns the Accuracy field if non-nil, zero value otherwise.

### GetAccuracyOk

`func (o *LocationReportingConfiguration1) GetAccuracyOk() (*LocationAccuracy, bool)`

GetAccuracyOk returns a tuple with the Accuracy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccuracy

`func (o *LocationReportingConfiguration1) SetAccuracy(v LocationAccuracy)`

SetAccuracy sets Accuracy field to given value.

### HasAccuracy

`func (o *LocationReportingConfiguration1) HasAccuracy() bool`

HasAccuracy returns a boolean if a field has been set.

### GetN3gppAccuracy

`func (o *LocationReportingConfiguration1) GetN3gppAccuracy() LocationAccuracy`

GetN3gppAccuracy returns the N3gppAccuracy field if non-nil, zero value otherwise.

### GetN3gppAccuracyOk

`func (o *LocationReportingConfiguration1) GetN3gppAccuracyOk() (*LocationAccuracy, bool)`

GetN3gppAccuracyOk returns a tuple with the N3gppAccuracy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetN3gppAccuracy

`func (o *LocationReportingConfiguration1) SetN3gppAccuracy(v LocationAccuracy)`

SetN3gppAccuracy sets N3gppAccuracy field to given value.

### HasN3gppAccuracy

`func (o *LocationReportingConfiguration1) HasN3gppAccuracy() bool`

HasN3gppAccuracy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


