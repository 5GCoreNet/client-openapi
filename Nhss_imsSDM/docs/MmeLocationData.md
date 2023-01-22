# MmeLocationData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MmeAddress** | **string** | Fully Qualified Domain Name | 
**PlmnId** | [**PlmnId**](PlmnId.md) |  | 
**MmeLocation** | Pointer to [**EutraLocation**](EutraLocation.md) |  | [optional] 
**CsgInformation** | Pointer to [**CsgInformation**](CsgInformation.md) |  | [optional] 
**TimeZone** | Pointer to **string** | String with format \&quot;time-numoffset\&quot; optionally appended by \&quot;daylightSavingTime\&quot;, where  - \&quot;time-numoffset\&quot; shall represent the time zone adjusted for daylight saving time and be    encoded as time-numoffset as defined in clause 5.6 of IETF RFC 3339;  - \&quot;daylightSavingTime\&quot; shall represent the adjustment that has been made and shall be    encoded as \&quot;+1\&quot; or \&quot;+2\&quot; for a +1 or +2 hours adjustment.   The example is for 8 hours behind UTC, +1 hour adjustment for Daylight Saving Time.  | [optional] 
**RatType** | Pointer to [**RatType**](RatType.md) |  | [optional] 

## Methods

### NewMmeLocationData

`func NewMmeLocationData(mmeAddress string, plmnId PlmnId, ) *MmeLocationData`

NewMmeLocationData instantiates a new MmeLocationData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMmeLocationDataWithDefaults

`func NewMmeLocationDataWithDefaults() *MmeLocationData`

NewMmeLocationDataWithDefaults instantiates a new MmeLocationData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMmeAddress

`func (o *MmeLocationData) GetMmeAddress() string`

GetMmeAddress returns the MmeAddress field if non-nil, zero value otherwise.

### GetMmeAddressOk

`func (o *MmeLocationData) GetMmeAddressOk() (*string, bool)`

GetMmeAddressOk returns a tuple with the MmeAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMmeAddress

`func (o *MmeLocationData) SetMmeAddress(v string)`

SetMmeAddress sets MmeAddress field to given value.


### GetPlmnId

`func (o *MmeLocationData) GetPlmnId() PlmnId`

GetPlmnId returns the PlmnId field if non-nil, zero value otherwise.

### GetPlmnIdOk

`func (o *MmeLocationData) GetPlmnIdOk() (*PlmnId, bool)`

GetPlmnIdOk returns a tuple with the PlmnId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlmnId

`func (o *MmeLocationData) SetPlmnId(v PlmnId)`

SetPlmnId sets PlmnId field to given value.


### GetMmeLocation

`func (o *MmeLocationData) GetMmeLocation() EutraLocation`

GetMmeLocation returns the MmeLocation field if non-nil, zero value otherwise.

### GetMmeLocationOk

`func (o *MmeLocationData) GetMmeLocationOk() (*EutraLocation, bool)`

GetMmeLocationOk returns a tuple with the MmeLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMmeLocation

`func (o *MmeLocationData) SetMmeLocation(v EutraLocation)`

SetMmeLocation sets MmeLocation field to given value.

### HasMmeLocation

`func (o *MmeLocationData) HasMmeLocation() bool`

HasMmeLocation returns a boolean if a field has been set.

### GetCsgInformation

`func (o *MmeLocationData) GetCsgInformation() CsgInformation`

GetCsgInformation returns the CsgInformation field if non-nil, zero value otherwise.

### GetCsgInformationOk

`func (o *MmeLocationData) GetCsgInformationOk() (*CsgInformation, bool)`

GetCsgInformationOk returns a tuple with the CsgInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCsgInformation

`func (o *MmeLocationData) SetCsgInformation(v CsgInformation)`

SetCsgInformation sets CsgInformation field to given value.

### HasCsgInformation

`func (o *MmeLocationData) HasCsgInformation() bool`

HasCsgInformation returns a boolean if a field has been set.

### GetTimeZone

`func (o *MmeLocationData) GetTimeZone() string`

GetTimeZone returns the TimeZone field if non-nil, zero value otherwise.

### GetTimeZoneOk

`func (o *MmeLocationData) GetTimeZoneOk() (*string, bool)`

GetTimeZoneOk returns a tuple with the TimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeZone

`func (o *MmeLocationData) SetTimeZone(v string)`

SetTimeZone sets TimeZone field to given value.

### HasTimeZone

`func (o *MmeLocationData) HasTimeZone() bool`

HasTimeZone returns a boolean if a field has been set.

### GetRatType

`func (o *MmeLocationData) GetRatType() RatType`

GetRatType returns the RatType field if non-nil, zero value otherwise.

### GetRatTypeOk

`func (o *MmeLocationData) GetRatTypeOk() (*RatType, bool)`

GetRatTypeOk returns a tuple with the RatType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRatType

`func (o *MmeLocationData) SetRatType(v RatType)`

SetRatType sets RatType field to given value.

### HasRatType

`func (o *MmeLocationData) HasRatType() bool`

HasRatType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


