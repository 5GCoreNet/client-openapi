# PtpCapabilitiesPerUe

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Gpsi** | **string** | String identifying a Gpsi shall contain either an External Id or an MSISDN.  It shall be formatted as follows -External Identifier&#x3D; \&quot;extid-&#39;extid&#39;, where &#39;extid&#39;  shall be formatted according to clause 19.7.2 of 3GPP TS 23.003 that describes an  External Identifier.   | 
**PtpCaps** | [**[]EventFilter**](EventFilter.md) |  | 

## Methods

### NewPtpCapabilitiesPerUe

`func NewPtpCapabilitiesPerUe(gpsi string, ptpCaps []EventFilter, ) *PtpCapabilitiesPerUe`

NewPtpCapabilitiesPerUe instantiates a new PtpCapabilitiesPerUe object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPtpCapabilitiesPerUeWithDefaults

`func NewPtpCapabilitiesPerUeWithDefaults() *PtpCapabilitiesPerUe`

NewPtpCapabilitiesPerUeWithDefaults instantiates a new PtpCapabilitiesPerUe object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGpsi

`func (o *PtpCapabilitiesPerUe) GetGpsi() string`

GetGpsi returns the Gpsi field if non-nil, zero value otherwise.

### GetGpsiOk

`func (o *PtpCapabilitiesPerUe) GetGpsiOk() (*string, bool)`

GetGpsiOk returns a tuple with the Gpsi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpsi

`func (o *PtpCapabilitiesPerUe) SetGpsi(v string)`

SetGpsi sets Gpsi field to given value.


### GetPtpCaps

`func (o *PtpCapabilitiesPerUe) GetPtpCaps() []EventFilter`

GetPtpCaps returns the PtpCaps field if non-nil, zero value otherwise.

### GetPtpCapsOk

`func (o *PtpCapabilitiesPerUe) GetPtpCapsOk() (*[]EventFilter, bool)`

GetPtpCapsOk returns a tuple with the PtpCaps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPtpCaps

`func (o *PtpCapabilitiesPerUe) SetPtpCaps(v []EventFilter)`

SetPtpCaps sets PtpCaps field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


