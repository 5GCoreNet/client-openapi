# ChargingInformation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PrimaryChfAddress** | **string** | String providing an URI formatted according to RFC 3986. | 
**SecondaryChfAddress** | Pointer to **string** | String providing an URI formatted according to RFC 3986. | [optional] 
**PrimaryChfSetId** | Pointer to **string** | NF Set Identifier (see clause 28.12 of 3GPP TS 23.003), formatted as the following string \&quot;set&lt;Set ID&gt;.&lt;nftype&gt;set.5gc.mnc&lt;MNC&gt;.mcc&lt;MCC&gt;\&quot;, or  \&quot;set&lt;SetID&gt;.&lt;NFType&gt;set.5gc.nid&lt;NID&gt;.mnc&lt;MNC&gt;.mcc&lt;MCC&gt;\&quot; with  &lt;MCC&gt; encoded as defined in clause 5.4.2 (\&quot;Mcc\&quot; data type definition)  &lt;MNC&gt; encoding the Mobile Network Code part of the PLMN, comprising 3 digits.    If there are only 2 significant digits in the MNC, one \&quot;0\&quot; digit shall be inserted    at the left side to fill the 3 digits coding of MNC.  Pattern: &#39;^[0-9]{3}$&#39; &lt;NFType&gt; encoded as a value defined in Table 6.1.6.3.3-1 of 3GPP TS 29.510 but    with lower case characters &lt;Set ID&gt; encoded as a string of characters consisting of    alphabetic characters (A-Z and a-z), digits (0-9) and/or the hyphen (-) and that    shall end with either an alphabetic character or a digit.   | [optional] 
**PrimaryChfInstanceId** | Pointer to **string** | String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.   | [optional] 
**SecondaryChfSetId** | Pointer to **string** | NF Set Identifier (see clause 28.12 of 3GPP TS 23.003), formatted as the following string \&quot;set&lt;Set ID&gt;.&lt;nftype&gt;set.5gc.mnc&lt;MNC&gt;.mcc&lt;MCC&gt;\&quot;, or  \&quot;set&lt;SetID&gt;.&lt;NFType&gt;set.5gc.nid&lt;NID&gt;.mnc&lt;MNC&gt;.mcc&lt;MCC&gt;\&quot; with  &lt;MCC&gt; encoded as defined in clause 5.4.2 (\&quot;Mcc\&quot; data type definition)  &lt;MNC&gt; encoding the Mobile Network Code part of the PLMN, comprising 3 digits.    If there are only 2 significant digits in the MNC, one \&quot;0\&quot; digit shall be inserted    at the left side to fill the 3 digits coding of MNC.  Pattern: &#39;^[0-9]{3}$&#39; &lt;NFType&gt; encoded as a value defined in Table 6.1.6.3.3-1 of 3GPP TS 29.510 but    with lower case characters &lt;Set ID&gt; encoded as a string of characters consisting of    alphabetic characters (A-Z and a-z), digits (0-9) and/or the hyphen (-) and that    shall end with either an alphabetic character or a digit.   | [optional] 
**SecondaryChfInstanceId** | Pointer to **string** | String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.   | [optional] 

## Methods

### NewChargingInformation

`func NewChargingInformation(primaryChfAddress string, ) *ChargingInformation`

NewChargingInformation instantiates a new ChargingInformation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChargingInformationWithDefaults

`func NewChargingInformationWithDefaults() *ChargingInformation`

NewChargingInformationWithDefaults instantiates a new ChargingInformation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrimaryChfAddress

`func (o *ChargingInformation) GetPrimaryChfAddress() string`

GetPrimaryChfAddress returns the PrimaryChfAddress field if non-nil, zero value otherwise.

### GetPrimaryChfAddressOk

`func (o *ChargingInformation) GetPrimaryChfAddressOk() (*string, bool)`

GetPrimaryChfAddressOk returns a tuple with the PrimaryChfAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryChfAddress

`func (o *ChargingInformation) SetPrimaryChfAddress(v string)`

SetPrimaryChfAddress sets PrimaryChfAddress field to given value.


### GetSecondaryChfAddress

`func (o *ChargingInformation) GetSecondaryChfAddress() string`

GetSecondaryChfAddress returns the SecondaryChfAddress field if non-nil, zero value otherwise.

### GetSecondaryChfAddressOk

`func (o *ChargingInformation) GetSecondaryChfAddressOk() (*string, bool)`

GetSecondaryChfAddressOk returns a tuple with the SecondaryChfAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryChfAddress

`func (o *ChargingInformation) SetSecondaryChfAddress(v string)`

SetSecondaryChfAddress sets SecondaryChfAddress field to given value.

### HasSecondaryChfAddress

`func (o *ChargingInformation) HasSecondaryChfAddress() bool`

HasSecondaryChfAddress returns a boolean if a field has been set.

### GetPrimaryChfSetId

`func (o *ChargingInformation) GetPrimaryChfSetId() string`

GetPrimaryChfSetId returns the PrimaryChfSetId field if non-nil, zero value otherwise.

### GetPrimaryChfSetIdOk

`func (o *ChargingInformation) GetPrimaryChfSetIdOk() (*string, bool)`

GetPrimaryChfSetIdOk returns a tuple with the PrimaryChfSetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryChfSetId

`func (o *ChargingInformation) SetPrimaryChfSetId(v string)`

SetPrimaryChfSetId sets PrimaryChfSetId field to given value.

### HasPrimaryChfSetId

`func (o *ChargingInformation) HasPrimaryChfSetId() bool`

HasPrimaryChfSetId returns a boolean if a field has been set.

### GetPrimaryChfInstanceId

`func (o *ChargingInformation) GetPrimaryChfInstanceId() string`

GetPrimaryChfInstanceId returns the PrimaryChfInstanceId field if non-nil, zero value otherwise.

### GetPrimaryChfInstanceIdOk

`func (o *ChargingInformation) GetPrimaryChfInstanceIdOk() (*string, bool)`

GetPrimaryChfInstanceIdOk returns a tuple with the PrimaryChfInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryChfInstanceId

`func (o *ChargingInformation) SetPrimaryChfInstanceId(v string)`

SetPrimaryChfInstanceId sets PrimaryChfInstanceId field to given value.

### HasPrimaryChfInstanceId

`func (o *ChargingInformation) HasPrimaryChfInstanceId() bool`

HasPrimaryChfInstanceId returns a boolean if a field has been set.

### GetSecondaryChfSetId

`func (o *ChargingInformation) GetSecondaryChfSetId() string`

GetSecondaryChfSetId returns the SecondaryChfSetId field if non-nil, zero value otherwise.

### GetSecondaryChfSetIdOk

`func (o *ChargingInformation) GetSecondaryChfSetIdOk() (*string, bool)`

GetSecondaryChfSetIdOk returns a tuple with the SecondaryChfSetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryChfSetId

`func (o *ChargingInformation) SetSecondaryChfSetId(v string)`

SetSecondaryChfSetId sets SecondaryChfSetId field to given value.

### HasSecondaryChfSetId

`func (o *ChargingInformation) HasSecondaryChfSetId() bool`

HasSecondaryChfSetId returns a boolean if a field has been set.

### GetSecondaryChfInstanceId

`func (o *ChargingInformation) GetSecondaryChfInstanceId() string`

GetSecondaryChfInstanceId returns the SecondaryChfInstanceId field if non-nil, zero value otherwise.

### GetSecondaryChfInstanceIdOk

`func (o *ChargingInformation) GetSecondaryChfInstanceIdOk() (*string, bool)`

GetSecondaryChfInstanceIdOk returns a tuple with the SecondaryChfInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryChfInstanceId

`func (o *ChargingInformation) SetSecondaryChfInstanceId(v string)`

SetSecondaryChfInstanceId sets SecondaryChfInstanceId field to given value.

### HasSecondaryChfInstanceId

`func (o *ChargingInformation) HasSecondaryChfInstanceId() bool`

HasSecondaryChfInstanceId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


