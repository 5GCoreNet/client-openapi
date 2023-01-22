# LocationEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UeId** | **string** | String identifying a Gpsi shall contain either an External Id or an MSISDN.  It shall be formatted as follows -External Identifier&#x3D; \&quot;extid-&#39;extid&#39;, where &#39;extid&#39;  shall be formatted according to clause 19.7.2 of 3GPP TS 23.003 that describes an  External Identifier.   | 
**LocInf** | Pointer to [**LocationInfo**](LocationInfo.md) |  | [optional] 
**LocInfPred** | Pointer to [**UeMobilityExposure**](UeMobilityExposure.md) |  | [optional] 

## Methods

### NewLocationEvent

`func NewLocationEvent(ueId string, ) *LocationEvent`

NewLocationEvent instantiates a new LocationEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLocationEventWithDefaults

`func NewLocationEventWithDefaults() *LocationEvent`

NewLocationEventWithDefaults instantiates a new LocationEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUeId

`func (o *LocationEvent) GetUeId() string`

GetUeId returns the UeId field if non-nil, zero value otherwise.

### GetUeIdOk

`func (o *LocationEvent) GetUeIdOk() (*string, bool)`

GetUeIdOk returns a tuple with the UeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUeId

`func (o *LocationEvent) SetUeId(v string)`

SetUeId sets UeId field to given value.


### GetLocInf

`func (o *LocationEvent) GetLocInf() LocationInfo`

GetLocInf returns the LocInf field if non-nil, zero value otherwise.

### GetLocInfOk

`func (o *LocationEvent) GetLocInfOk() (*LocationInfo, bool)`

GetLocInfOk returns a tuple with the LocInf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocInf

`func (o *LocationEvent) SetLocInf(v LocationInfo)`

SetLocInf sets LocInf field to given value.

### HasLocInf

`func (o *LocationEvent) HasLocInf() bool`

HasLocInf returns a boolean if a field has been set.

### GetLocInfPred

`func (o *LocationEvent) GetLocInfPred() UeMobilityExposure`

GetLocInfPred returns the LocInfPred field if non-nil, zero value otherwise.

### GetLocInfPredOk

`func (o *LocationEvent) GetLocInfPredOk() (*UeMobilityExposure, bool)`

GetLocInfPredOk returns a tuple with the LocInfPred field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocInfPred

`func (o *LocationEvent) SetLocInfPred(v UeMobilityExposure)`

SetLocInfPred sets LocInfPred field to given value.

### HasLocInfPred

`func (o *LocationEvent) HasLocInfPred() bool`

HasLocInfPred returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


