# CsLocation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MscNumber** | **string** |  | 
**VlrNumber** | **string** |  | 
**PlmnId** | [**PlmnId**](PlmnId.md) |  | 
**VlrLocation** | Pointer to [**GeraLocation**](GeraLocation.md) |  | [optional] 
**CsgInformation** | Pointer to [**CsgInformation**](CsgInformation.md) |  | [optional] 
**TimeZone** | Pointer to **string** | String with format \&quot;time-numoffset\&quot; optionally appended by \&quot;daylightSavingTime\&quot;, where  - \&quot;time-numoffset\&quot; shall represent the time zone adjusted for daylight saving time and be    encoded as time-numoffset as defined in clause 5.6 of IETF RFC 3339;  - \&quot;daylightSavingTime\&quot; shall represent the adjustment that has been made and shall be    encoded as \&quot;+1\&quot; or \&quot;+2\&quot; for a +1 or +2 hours adjustment.   The example is for 8 hours behind UTC, +1 hour adjustment for Daylight Saving Time.  | [optional] 
**EUtranCgi** | Pointer to [**Ecgi**](Ecgi.md) |  | [optional] 
**Tai** | Pointer to [**Tai**](Tai.md) |  | [optional] 

## Methods

### NewCsLocation

`func NewCsLocation(mscNumber string, vlrNumber string, plmnId PlmnId, ) *CsLocation`

NewCsLocation instantiates a new CsLocation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCsLocationWithDefaults

`func NewCsLocationWithDefaults() *CsLocation`

NewCsLocationWithDefaults instantiates a new CsLocation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMscNumber

`func (o *CsLocation) GetMscNumber() string`

GetMscNumber returns the MscNumber field if non-nil, zero value otherwise.

### GetMscNumberOk

`func (o *CsLocation) GetMscNumberOk() (*string, bool)`

GetMscNumberOk returns a tuple with the MscNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMscNumber

`func (o *CsLocation) SetMscNumber(v string)`

SetMscNumber sets MscNumber field to given value.


### GetVlrNumber

`func (o *CsLocation) GetVlrNumber() string`

GetVlrNumber returns the VlrNumber field if non-nil, zero value otherwise.

### GetVlrNumberOk

`func (o *CsLocation) GetVlrNumberOk() (*string, bool)`

GetVlrNumberOk returns a tuple with the VlrNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlrNumber

`func (o *CsLocation) SetVlrNumber(v string)`

SetVlrNumber sets VlrNumber field to given value.


### GetPlmnId

`func (o *CsLocation) GetPlmnId() PlmnId`

GetPlmnId returns the PlmnId field if non-nil, zero value otherwise.

### GetPlmnIdOk

`func (o *CsLocation) GetPlmnIdOk() (*PlmnId, bool)`

GetPlmnIdOk returns a tuple with the PlmnId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlmnId

`func (o *CsLocation) SetPlmnId(v PlmnId)`

SetPlmnId sets PlmnId field to given value.


### GetVlrLocation

`func (o *CsLocation) GetVlrLocation() GeraLocation`

GetVlrLocation returns the VlrLocation field if non-nil, zero value otherwise.

### GetVlrLocationOk

`func (o *CsLocation) GetVlrLocationOk() (*GeraLocation, bool)`

GetVlrLocationOk returns a tuple with the VlrLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlrLocation

`func (o *CsLocation) SetVlrLocation(v GeraLocation)`

SetVlrLocation sets VlrLocation field to given value.

### HasVlrLocation

`func (o *CsLocation) HasVlrLocation() bool`

HasVlrLocation returns a boolean if a field has been set.

### GetCsgInformation

`func (o *CsLocation) GetCsgInformation() CsgInformation`

GetCsgInformation returns the CsgInformation field if non-nil, zero value otherwise.

### GetCsgInformationOk

`func (o *CsLocation) GetCsgInformationOk() (*CsgInformation, bool)`

GetCsgInformationOk returns a tuple with the CsgInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCsgInformation

`func (o *CsLocation) SetCsgInformation(v CsgInformation)`

SetCsgInformation sets CsgInformation field to given value.

### HasCsgInformation

`func (o *CsLocation) HasCsgInformation() bool`

HasCsgInformation returns a boolean if a field has been set.

### GetTimeZone

`func (o *CsLocation) GetTimeZone() string`

GetTimeZone returns the TimeZone field if non-nil, zero value otherwise.

### GetTimeZoneOk

`func (o *CsLocation) GetTimeZoneOk() (*string, bool)`

GetTimeZoneOk returns a tuple with the TimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeZone

`func (o *CsLocation) SetTimeZone(v string)`

SetTimeZone sets TimeZone field to given value.

### HasTimeZone

`func (o *CsLocation) HasTimeZone() bool`

HasTimeZone returns a boolean if a field has been set.

### GetEUtranCgi

`func (o *CsLocation) GetEUtranCgi() Ecgi`

GetEUtranCgi returns the EUtranCgi field if non-nil, zero value otherwise.

### GetEUtranCgiOk

`func (o *CsLocation) GetEUtranCgiOk() (*Ecgi, bool)`

GetEUtranCgiOk returns a tuple with the EUtranCgi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEUtranCgi

`func (o *CsLocation) SetEUtranCgi(v Ecgi)`

SetEUtranCgi sets EUtranCgi field to given value.

### HasEUtranCgi

`func (o *CsLocation) HasEUtranCgi() bool`

HasEUtranCgi returns a boolean if a field has been set.

### GetTai

`func (o *CsLocation) GetTai() Tai`

GetTai returns the Tai field if non-nil, zero value otherwise.

### GetTaiOk

`func (o *CsLocation) GetTaiOk() (*Tai, bool)`

GetTaiOk returns a tuple with the Tai field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTai

`func (o *CsLocation) SetTai(v Tai)`

SetTai sets Tai field to given value.

### HasTai

`func (o *CsLocation) HasTai() bool`

HasTai returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


