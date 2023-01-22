# TwanLocationData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TwanSsid** | **string** |  | 
**PlmnId** | [**PlmnId**](PlmnId.md) |  | 
**TwanBssid** | Pointer to **string** |  | [optional] 
**CivicAddress** | Pointer to **string** |  | [optional] 
**TwanOperatorName** | Pointer to **string** |  | [optional] 
**TimeZone** | Pointer to **string** | String with format \&quot;time-numoffset\&quot; optionally appended by \&quot;daylightSavingTime\&quot;, where  - \&quot;time-numoffset\&quot; shall represent the time zone adjusted for daylight saving time and be    encoded as time-numoffset as defined in clause 5.6 of IETF RFC 3339;  - \&quot;daylightSavingTime\&quot; shall represent the adjustment that has been made and shall be    encoded as \&quot;+1\&quot; or \&quot;+2\&quot; for a +1 or +2 hours adjustment.   The example is for 8 hours behind UTC, +1 hour adjustment for Daylight Saving Time.  | [optional] 
**LogicalAccessId** | Pointer to **string** |  | [optional] 

## Methods

### NewTwanLocationData

`func NewTwanLocationData(twanSsid string, plmnId PlmnId, ) *TwanLocationData`

NewTwanLocationData instantiates a new TwanLocationData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTwanLocationDataWithDefaults

`func NewTwanLocationDataWithDefaults() *TwanLocationData`

NewTwanLocationDataWithDefaults instantiates a new TwanLocationData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTwanSsid

`func (o *TwanLocationData) GetTwanSsid() string`

GetTwanSsid returns the TwanSsid field if non-nil, zero value otherwise.

### GetTwanSsidOk

`func (o *TwanLocationData) GetTwanSsidOk() (*string, bool)`

GetTwanSsidOk returns a tuple with the TwanSsid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwanSsid

`func (o *TwanLocationData) SetTwanSsid(v string)`

SetTwanSsid sets TwanSsid field to given value.


### GetPlmnId

`func (o *TwanLocationData) GetPlmnId() PlmnId`

GetPlmnId returns the PlmnId field if non-nil, zero value otherwise.

### GetPlmnIdOk

`func (o *TwanLocationData) GetPlmnIdOk() (*PlmnId, bool)`

GetPlmnIdOk returns a tuple with the PlmnId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlmnId

`func (o *TwanLocationData) SetPlmnId(v PlmnId)`

SetPlmnId sets PlmnId field to given value.


### GetTwanBssid

`func (o *TwanLocationData) GetTwanBssid() string`

GetTwanBssid returns the TwanBssid field if non-nil, zero value otherwise.

### GetTwanBssidOk

`func (o *TwanLocationData) GetTwanBssidOk() (*string, bool)`

GetTwanBssidOk returns a tuple with the TwanBssid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwanBssid

`func (o *TwanLocationData) SetTwanBssid(v string)`

SetTwanBssid sets TwanBssid field to given value.

### HasTwanBssid

`func (o *TwanLocationData) HasTwanBssid() bool`

HasTwanBssid returns a boolean if a field has been set.

### GetCivicAddress

`func (o *TwanLocationData) GetCivicAddress() string`

GetCivicAddress returns the CivicAddress field if non-nil, zero value otherwise.

### GetCivicAddressOk

`func (o *TwanLocationData) GetCivicAddressOk() (*string, bool)`

GetCivicAddressOk returns a tuple with the CivicAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCivicAddress

`func (o *TwanLocationData) SetCivicAddress(v string)`

SetCivicAddress sets CivicAddress field to given value.

### HasCivicAddress

`func (o *TwanLocationData) HasCivicAddress() bool`

HasCivicAddress returns a boolean if a field has been set.

### GetTwanOperatorName

`func (o *TwanLocationData) GetTwanOperatorName() string`

GetTwanOperatorName returns the TwanOperatorName field if non-nil, zero value otherwise.

### GetTwanOperatorNameOk

`func (o *TwanLocationData) GetTwanOperatorNameOk() (*string, bool)`

GetTwanOperatorNameOk returns a tuple with the TwanOperatorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwanOperatorName

`func (o *TwanLocationData) SetTwanOperatorName(v string)`

SetTwanOperatorName sets TwanOperatorName field to given value.

### HasTwanOperatorName

`func (o *TwanLocationData) HasTwanOperatorName() bool`

HasTwanOperatorName returns a boolean if a field has been set.

### GetTimeZone

`func (o *TwanLocationData) GetTimeZone() string`

GetTimeZone returns the TimeZone field if non-nil, zero value otherwise.

### GetTimeZoneOk

`func (o *TwanLocationData) GetTimeZoneOk() (*string, bool)`

GetTimeZoneOk returns a tuple with the TimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeZone

`func (o *TwanLocationData) SetTimeZone(v string)`

SetTimeZone sets TimeZone field to given value.

### HasTimeZone

`func (o *TwanLocationData) HasTimeZone() bool`

HasTimeZone returns a boolean if a field has been set.

### GetLogicalAccessId

`func (o *TwanLocationData) GetLogicalAccessId() string`

GetLogicalAccessId returns the LogicalAccessId field if non-nil, zero value otherwise.

### GetLogicalAccessIdOk

`func (o *TwanLocationData) GetLogicalAccessIdOk() (*string, bool)`

GetLogicalAccessIdOk returns a tuple with the LogicalAccessId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogicalAccessId

`func (o *TwanLocationData) SetLogicalAccessId(v string)`

SetLogicalAccessId sets LogicalAccessId field to given value.

### HasLogicalAccessId

`func (o *TwanLocationData) HasLogicalAccessId() bool`

HasLogicalAccessId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


