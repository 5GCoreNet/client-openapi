# SgsnLocationData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SgsnNumber** | **string** |  | 
**PlmnId** | [**PlmnId**](PlmnId.md) |  | 
**SgsnLocation** | Pointer to [**UtraLocation**](UtraLocation.md) |  | [optional] 
**CsgInformation** | Pointer to [**CsgInformation**](CsgInformation.md) |  | [optional] 
**TimeZone** | Pointer to **string** | String with format \&quot;time-numoffset\&quot; optionally appended by \&quot;daylightSavingTime\&quot;, where  - \&quot;time-numoffset\&quot; shall represent the time zone adjusted for daylight saving time and be    encoded as time-numoffset as defined in clause 5.6 of IETF RFC 3339;  - \&quot;daylightSavingTime\&quot; shall represent the adjustment that has been made and shall be    encoded as \&quot;+1\&quot; or \&quot;+2\&quot; for a +1 or +2 hours adjustment.   The example is for 8 hours behind UTC, +1 hour adjustment for Daylight Saving Time.  | [optional] 
**RatType** | Pointer to [**RatType**](RatType.md) |  | [optional] 

## Methods

### NewSgsnLocationData

`func NewSgsnLocationData(sgsnNumber string, plmnId PlmnId, ) *SgsnLocationData`

NewSgsnLocationData instantiates a new SgsnLocationData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSgsnLocationDataWithDefaults

`func NewSgsnLocationDataWithDefaults() *SgsnLocationData`

NewSgsnLocationDataWithDefaults instantiates a new SgsnLocationData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSgsnNumber

`func (o *SgsnLocationData) GetSgsnNumber() string`

GetSgsnNumber returns the SgsnNumber field if non-nil, zero value otherwise.

### GetSgsnNumberOk

`func (o *SgsnLocationData) GetSgsnNumberOk() (*string, bool)`

GetSgsnNumberOk returns a tuple with the SgsnNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSgsnNumber

`func (o *SgsnLocationData) SetSgsnNumber(v string)`

SetSgsnNumber sets SgsnNumber field to given value.


### GetPlmnId

`func (o *SgsnLocationData) GetPlmnId() PlmnId`

GetPlmnId returns the PlmnId field if non-nil, zero value otherwise.

### GetPlmnIdOk

`func (o *SgsnLocationData) GetPlmnIdOk() (*PlmnId, bool)`

GetPlmnIdOk returns a tuple with the PlmnId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlmnId

`func (o *SgsnLocationData) SetPlmnId(v PlmnId)`

SetPlmnId sets PlmnId field to given value.


### GetSgsnLocation

`func (o *SgsnLocationData) GetSgsnLocation() UtraLocation`

GetSgsnLocation returns the SgsnLocation field if non-nil, zero value otherwise.

### GetSgsnLocationOk

`func (o *SgsnLocationData) GetSgsnLocationOk() (*UtraLocation, bool)`

GetSgsnLocationOk returns a tuple with the SgsnLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSgsnLocation

`func (o *SgsnLocationData) SetSgsnLocation(v UtraLocation)`

SetSgsnLocation sets SgsnLocation field to given value.

### HasSgsnLocation

`func (o *SgsnLocationData) HasSgsnLocation() bool`

HasSgsnLocation returns a boolean if a field has been set.

### GetCsgInformation

`func (o *SgsnLocationData) GetCsgInformation() CsgInformation`

GetCsgInformation returns the CsgInformation field if non-nil, zero value otherwise.

### GetCsgInformationOk

`func (o *SgsnLocationData) GetCsgInformationOk() (*CsgInformation, bool)`

GetCsgInformationOk returns a tuple with the CsgInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCsgInformation

`func (o *SgsnLocationData) SetCsgInformation(v CsgInformation)`

SetCsgInformation sets CsgInformation field to given value.

### HasCsgInformation

`func (o *SgsnLocationData) HasCsgInformation() bool`

HasCsgInformation returns a boolean if a field has been set.

### GetTimeZone

`func (o *SgsnLocationData) GetTimeZone() string`

GetTimeZone returns the TimeZone field if non-nil, zero value otherwise.

### GetTimeZoneOk

`func (o *SgsnLocationData) GetTimeZoneOk() (*string, bool)`

GetTimeZoneOk returns a tuple with the TimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeZone

`func (o *SgsnLocationData) SetTimeZone(v string)`

SetTimeZone sets TimeZone field to given value.

### HasTimeZone

`func (o *SgsnLocationData) HasTimeZone() bool`

HasTimeZone returns a boolean if a field has been set.

### GetRatType

`func (o *SgsnLocationData) GetRatType() RatType`

GetRatType returns the RatType field if non-nil, zero value otherwise.

### GetRatTypeOk

`func (o *SgsnLocationData) GetRatTypeOk() (*RatType, bool)`

GetRatTypeOk returns a tuple with the RatType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRatType

`func (o *SgsnLocationData) SetRatType(v RatType)`

SetRatType sets RatType field to given value.

### HasRatType

`func (o *SgsnLocationData) HasRatType() bool`

HasRatType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


