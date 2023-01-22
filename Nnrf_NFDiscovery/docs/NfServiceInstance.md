# NfServiceInstance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServiceInstanceId** | Pointer to **string** |  | [optional] 
**NfInstanceId** | Pointer to **string** | String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.   | [optional] 
**NfServiceSetId** | Pointer to **string** | NF Service Set Identifier (see clause 28.12 of 3GPP TS 23.003) formatted as the following  string \&quot;set&lt;Set ID&gt;.sn&lt;Service Name&gt;.nfi&lt;NF Instance ID&gt;.5gc.mnc&lt;MNC&gt;.mcc&lt;MCC&gt;\&quot;, or  \&quot;set&lt;SetID&gt;.sn&lt;ServiceName&gt;.nfi&lt;NFInstanceID&gt;.5gc.nid&lt;NID&gt;.mnc&lt;MNC&gt;.mcc&lt;MCC&gt;\&quot; with  &lt;MCC&gt; encoded as defined in clause 5.4.2 (\&quot;Mcc\&quot; data type definition)   &lt;MNC&gt; encoding the Mobile Network Code part of the PLMN, comprising 3 digits.    If there are only 2 significant digits in the MNC, one \&quot;0\&quot; digit shall be inserted    at the left side to fill the 3 digits coding of MNC.  Pattern: &#39;^[0-9]{3}$&#39; &lt;NID&gt; encoded as defined in clause 5.4.2 (\&quot;Nid\&quot; data type definition)  &lt;NFInstanceId&gt; encoded as defined in clause 5.3.2  &lt;ServiceName&gt; encoded as defined in 3GPP TS 29.510  &lt;Set ID&gt; encoded as a string of characters consisting of alphabetic    characters (A-Z and a-z), digits (0-9) and/or the hyphen (-) and that shall end    with either an alphabetic character or a digit.  | [optional] 

## Methods

### NewNfServiceInstance

`func NewNfServiceInstance() *NfServiceInstance`

NewNfServiceInstance instantiates a new NfServiceInstance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNfServiceInstanceWithDefaults

`func NewNfServiceInstanceWithDefaults() *NfServiceInstance`

NewNfServiceInstanceWithDefaults instantiates a new NfServiceInstance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServiceInstanceId

`func (o *NfServiceInstance) GetServiceInstanceId() string`

GetServiceInstanceId returns the ServiceInstanceId field if non-nil, zero value otherwise.

### GetServiceInstanceIdOk

`func (o *NfServiceInstance) GetServiceInstanceIdOk() (*string, bool)`

GetServiceInstanceIdOk returns a tuple with the ServiceInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceInstanceId

`func (o *NfServiceInstance) SetServiceInstanceId(v string)`

SetServiceInstanceId sets ServiceInstanceId field to given value.

### HasServiceInstanceId

`func (o *NfServiceInstance) HasServiceInstanceId() bool`

HasServiceInstanceId returns a boolean if a field has been set.

### GetNfInstanceId

`func (o *NfServiceInstance) GetNfInstanceId() string`

GetNfInstanceId returns the NfInstanceId field if non-nil, zero value otherwise.

### GetNfInstanceIdOk

`func (o *NfServiceInstance) GetNfInstanceIdOk() (*string, bool)`

GetNfInstanceIdOk returns a tuple with the NfInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNfInstanceId

`func (o *NfServiceInstance) SetNfInstanceId(v string)`

SetNfInstanceId sets NfInstanceId field to given value.

### HasNfInstanceId

`func (o *NfServiceInstance) HasNfInstanceId() bool`

HasNfInstanceId returns a boolean if a field has been set.

### GetNfServiceSetId

`func (o *NfServiceInstance) GetNfServiceSetId() string`

GetNfServiceSetId returns the NfServiceSetId field if non-nil, zero value otherwise.

### GetNfServiceSetIdOk

`func (o *NfServiceInstance) GetNfServiceSetIdOk() (*string, bool)`

GetNfServiceSetIdOk returns a tuple with the NfServiceSetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNfServiceSetId

`func (o *NfServiceInstance) SetNfServiceSetId(v string)`

SetNfServiceSetId sets NfServiceSetId field to given value.

### HasNfServiceSetId

`func (o *NfServiceInstance) HasNfServiceSetId() bool`

HasNfServiceSetId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


