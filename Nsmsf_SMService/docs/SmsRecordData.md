# SmsRecordData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SmsRecordId** | **string** | Represents a record ID, used to identify a message carrying SMS payload. | 
**SmsPayload** | [**RefToBinaryData**](RefToBinaryData.md) |  | 
**AccessType** | Pointer to [**AccessType**](AccessType.md) |  | [optional] 
**Gpsi** | Pointer to **string** | String identifying a Gpsi shall contain either an External Id or an MSISDN.  It shall be formatted as follows -External Identifier&#x3D; \&quot;extid-&#39;extid&#39;, where &#39;extid&#39;  shall be formatted according to clause 19.7.2 of 3GPP TS 23.003 that describes an  External Identifier.   | [optional] 
**Pei** | Pointer to **string** | String representing a Permanent Equipment Identifier that may contain - an IMEI or IMEISV, as  specified in clause 6.2 of 3GPP TS 23.003; a MAC address for a 5G-RG or FN-RG via  wireline  access, with an indication that this address cannot be trusted for regulatory purpose if this  address cannot be used as an Equipment Identifier of the FN-RG, as specified in clause 4.7.7  of 3GPP TS23.316. Examples are imei-012345678901234 or imeisv-0123456789012345.   | [optional] 
**UeLocation** | Pointer to [**UserLocation**](UserLocation.md) |  | [optional] 
**UeTimeZone** | Pointer to **string** | String with format \&quot;time-numoffset\&quot; optionally appended by \&quot;daylightSavingTime\&quot;, where  - \&quot;time-numoffset\&quot; shall represent the time zone adjusted for daylight saving time and be    encoded as time-numoffset as defined in clause 5.6 of IETF RFC 3339;  - \&quot;daylightSavingTime\&quot; shall represent the adjustment that has been made and shall be    encoded as \&quot;+1\&quot; or \&quot;+2\&quot; for a +1 or +2 hours adjustment.   The example is for 8 hours behind UTC, +1 hour adjustment for Daylight Saving Time.  | [optional] 

## Methods

### NewSmsRecordData

`func NewSmsRecordData(smsRecordId string, smsPayload RefToBinaryData, ) *SmsRecordData`

NewSmsRecordData instantiates a new SmsRecordData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSmsRecordDataWithDefaults

`func NewSmsRecordDataWithDefaults() *SmsRecordData`

NewSmsRecordDataWithDefaults instantiates a new SmsRecordData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSmsRecordId

`func (o *SmsRecordData) GetSmsRecordId() string`

GetSmsRecordId returns the SmsRecordId field if non-nil, zero value otherwise.

### GetSmsRecordIdOk

`func (o *SmsRecordData) GetSmsRecordIdOk() (*string, bool)`

GetSmsRecordIdOk returns a tuple with the SmsRecordId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmsRecordId

`func (o *SmsRecordData) SetSmsRecordId(v string)`

SetSmsRecordId sets SmsRecordId field to given value.


### GetSmsPayload

`func (o *SmsRecordData) GetSmsPayload() RefToBinaryData`

GetSmsPayload returns the SmsPayload field if non-nil, zero value otherwise.

### GetSmsPayloadOk

`func (o *SmsRecordData) GetSmsPayloadOk() (*RefToBinaryData, bool)`

GetSmsPayloadOk returns a tuple with the SmsPayload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmsPayload

`func (o *SmsRecordData) SetSmsPayload(v RefToBinaryData)`

SetSmsPayload sets SmsPayload field to given value.


### GetAccessType

`func (o *SmsRecordData) GetAccessType() AccessType`

GetAccessType returns the AccessType field if non-nil, zero value otherwise.

### GetAccessTypeOk

`func (o *SmsRecordData) GetAccessTypeOk() (*AccessType, bool)`

GetAccessTypeOk returns a tuple with the AccessType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessType

`func (o *SmsRecordData) SetAccessType(v AccessType)`

SetAccessType sets AccessType field to given value.

### HasAccessType

`func (o *SmsRecordData) HasAccessType() bool`

HasAccessType returns a boolean if a field has been set.

### GetGpsi

`func (o *SmsRecordData) GetGpsi() string`

GetGpsi returns the Gpsi field if non-nil, zero value otherwise.

### GetGpsiOk

`func (o *SmsRecordData) GetGpsiOk() (*string, bool)`

GetGpsiOk returns a tuple with the Gpsi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpsi

`func (o *SmsRecordData) SetGpsi(v string)`

SetGpsi sets Gpsi field to given value.

### HasGpsi

`func (o *SmsRecordData) HasGpsi() bool`

HasGpsi returns a boolean if a field has been set.

### GetPei

`func (o *SmsRecordData) GetPei() string`

GetPei returns the Pei field if non-nil, zero value otherwise.

### GetPeiOk

`func (o *SmsRecordData) GetPeiOk() (*string, bool)`

GetPeiOk returns a tuple with the Pei field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPei

`func (o *SmsRecordData) SetPei(v string)`

SetPei sets Pei field to given value.

### HasPei

`func (o *SmsRecordData) HasPei() bool`

HasPei returns a boolean if a field has been set.

### GetUeLocation

`func (o *SmsRecordData) GetUeLocation() UserLocation`

GetUeLocation returns the UeLocation field if non-nil, zero value otherwise.

### GetUeLocationOk

`func (o *SmsRecordData) GetUeLocationOk() (*UserLocation, bool)`

GetUeLocationOk returns a tuple with the UeLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUeLocation

`func (o *SmsRecordData) SetUeLocation(v UserLocation)`

SetUeLocation sets UeLocation field to given value.

### HasUeLocation

`func (o *SmsRecordData) HasUeLocation() bool`

HasUeLocation returns a boolean if a field has been set.

### GetUeTimeZone

`func (o *SmsRecordData) GetUeTimeZone() string`

GetUeTimeZone returns the UeTimeZone field if non-nil, zero value otherwise.

### GetUeTimeZoneOk

`func (o *SmsRecordData) GetUeTimeZoneOk() (*string, bool)`

GetUeTimeZoneOk returns a tuple with the UeTimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUeTimeZone

`func (o *SmsRecordData) SetUeTimeZone(v string)`

SetUeTimeZone sets UeTimeZone field to given value.

### HasUeTimeZone

`func (o *SmsRecordData) HasUeTimeZone() bool`

HasUeTimeZone returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


