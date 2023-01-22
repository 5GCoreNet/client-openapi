# AmfLocationData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AmfAddress** | **string** | String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.   | 
**PlmnId** | [**PlmnId**](PlmnId.md) |  | 
**AmfLocation** | Pointer to [**NrLocation**](NrLocation.md) |  | [optional] 
**SmsfAddress** | Pointer to **string** | String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.   | [optional] 
**TimeZone** | Pointer to **string** | String with format \&quot;time-numoffset\&quot; optionally appended by \&quot;daylightSavingTime\&quot;, where  - \&quot;time-numoffset\&quot; shall represent the time zone adjusted for daylight saving time and be    encoded as time-numoffset as defined in clause 5.6 of IETF RFC 3339;  - \&quot;daylightSavingTime\&quot; shall represent the adjustment that has been made and shall be    encoded as \&quot;+1\&quot; or \&quot;+2\&quot; for a +1 or +2 hours adjustment.   The example is for 8 hours behind UTC, +1 hour adjustment for Daylight Saving Time.  | [optional] 
**RatType** | Pointer to [**RatType**](RatType.md) |  | [optional] 

## Methods

### NewAmfLocationData

`func NewAmfLocationData(amfAddress string, plmnId PlmnId, ) *AmfLocationData`

NewAmfLocationData instantiates a new AmfLocationData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAmfLocationDataWithDefaults

`func NewAmfLocationDataWithDefaults() *AmfLocationData`

NewAmfLocationDataWithDefaults instantiates a new AmfLocationData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmfAddress

`func (o *AmfLocationData) GetAmfAddress() string`

GetAmfAddress returns the AmfAddress field if non-nil, zero value otherwise.

### GetAmfAddressOk

`func (o *AmfLocationData) GetAmfAddressOk() (*string, bool)`

GetAmfAddressOk returns a tuple with the AmfAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmfAddress

`func (o *AmfLocationData) SetAmfAddress(v string)`

SetAmfAddress sets AmfAddress field to given value.


### GetPlmnId

`func (o *AmfLocationData) GetPlmnId() PlmnId`

GetPlmnId returns the PlmnId field if non-nil, zero value otherwise.

### GetPlmnIdOk

`func (o *AmfLocationData) GetPlmnIdOk() (*PlmnId, bool)`

GetPlmnIdOk returns a tuple with the PlmnId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlmnId

`func (o *AmfLocationData) SetPlmnId(v PlmnId)`

SetPlmnId sets PlmnId field to given value.


### GetAmfLocation

`func (o *AmfLocationData) GetAmfLocation() NrLocation`

GetAmfLocation returns the AmfLocation field if non-nil, zero value otherwise.

### GetAmfLocationOk

`func (o *AmfLocationData) GetAmfLocationOk() (*NrLocation, bool)`

GetAmfLocationOk returns a tuple with the AmfLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmfLocation

`func (o *AmfLocationData) SetAmfLocation(v NrLocation)`

SetAmfLocation sets AmfLocation field to given value.

### HasAmfLocation

`func (o *AmfLocationData) HasAmfLocation() bool`

HasAmfLocation returns a boolean if a field has been set.

### GetSmsfAddress

`func (o *AmfLocationData) GetSmsfAddress() string`

GetSmsfAddress returns the SmsfAddress field if non-nil, zero value otherwise.

### GetSmsfAddressOk

`func (o *AmfLocationData) GetSmsfAddressOk() (*string, bool)`

GetSmsfAddressOk returns a tuple with the SmsfAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmsfAddress

`func (o *AmfLocationData) SetSmsfAddress(v string)`

SetSmsfAddress sets SmsfAddress field to given value.

### HasSmsfAddress

`func (o *AmfLocationData) HasSmsfAddress() bool`

HasSmsfAddress returns a boolean if a field has been set.

### GetTimeZone

`func (o *AmfLocationData) GetTimeZone() string`

GetTimeZone returns the TimeZone field if non-nil, zero value otherwise.

### GetTimeZoneOk

`func (o *AmfLocationData) GetTimeZoneOk() (*string, bool)`

GetTimeZoneOk returns a tuple with the TimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeZone

`func (o *AmfLocationData) SetTimeZone(v string)`

SetTimeZone sets TimeZone field to given value.

### HasTimeZone

`func (o *AmfLocationData) HasTimeZone() bool`

HasTimeZone returns a boolean if a field has been set.

### GetRatType

`func (o *AmfLocationData) GetRatType() RatType`

GetRatType returns the RatType field if non-nil, zero value otherwise.

### GetRatTypeOk

`func (o *AmfLocationData) GetRatTypeOk() (*RatType, bool)`

GetRatTypeOk returns a tuple with the RatType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRatType

`func (o *AmfLocationData) SetRatType(v RatType)`

SetRatType sets RatType field to given value.

### HasRatType

`func (o *AmfLocationData) HasRatType() bool`

HasRatType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


