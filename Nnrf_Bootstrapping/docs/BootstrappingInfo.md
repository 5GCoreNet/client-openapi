# BootstrappingInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to [**Status**](Status.md) |  | [optional] 
**Links** | [**map[string]LinksValueSchema**](LinksValueSchema.md) | Map of link objects where the keys are the link relations defined in 3GPP TS 29.510 clause 6.4.6.3.3  | 
**NrfFeatures** | Pointer to **map[string]string** | Map of features supported by the NRF, where the keys are the NRF services as defined in 3GPP TS 29.510 clause 6.1.6.3.11  | [optional] 
**Oauth2Required** | Pointer to **map[string]bool** | Map indicating whether the NRF requires Oauth2-based authorization for accessing its services. The key of the map shall be the name of an NRF service, e.g. \&quot;nnrf-nfm\&quot; or \&quot;nnrf-disc\&quot;  | [optional] 
**NrfSetId** | Pointer to **string** | NF Set Identifier (see clause 28.12 of 3GPP TS 23.003), formatted as the following string \&quot;set&lt;Set ID&gt;.&lt;nftype&gt;set.5gc.mnc&lt;MNC&gt;.mcc&lt;MCC&gt;\&quot;, or  \&quot;set&lt;SetID&gt;.&lt;NFType&gt;set.5gc.nid&lt;NID&gt;.mnc&lt;MNC&gt;.mcc&lt;MCC&gt;\&quot; with  &lt;MCC&gt; encoded as defined in clause 5.4.2 (\&quot;Mcc\&quot; data type definition)  &lt;MNC&gt; encoding the Mobile Network Code part of the PLMN, comprising 3 digits.    If there are only 2 significant digits in the MNC, one \&quot;0\&quot; digit shall be inserted    at the left side to fill the 3 digits coding of MNC.  Pattern: &#39;^[0-9]{3}$&#39; &lt;NFType&gt; encoded as a value defined in Table 6.1.6.3.3-1 of 3GPP TS 29.510 but    with lower case characters &lt;Set ID&gt; encoded as a string of characters consisting of    alphabetic characters (A-Z and a-z), digits (0-9) and/or the hyphen (-) and that    shall end with either an alphabetic character or a digit.   | [optional] 
**NrfInstanceId** | Pointer to **string** | String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.   | [optional] 

## Methods

### NewBootstrappingInfo

`func NewBootstrappingInfo(links map[string]LinksValueSchema, ) *BootstrappingInfo`

NewBootstrappingInfo instantiates a new BootstrappingInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBootstrappingInfoWithDefaults

`func NewBootstrappingInfoWithDefaults() *BootstrappingInfo`

NewBootstrappingInfoWithDefaults instantiates a new BootstrappingInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *BootstrappingInfo) GetStatus() Status`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BootstrappingInfo) GetStatusOk() (*Status, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BootstrappingInfo) SetStatus(v Status)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *BootstrappingInfo) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetLinks

`func (o *BootstrappingInfo) GetLinks() map[string]LinksValueSchema`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *BootstrappingInfo) GetLinksOk() (*map[string]LinksValueSchema, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *BootstrappingInfo) SetLinks(v map[string]LinksValueSchema)`

SetLinks sets Links field to given value.


### GetNrfFeatures

`func (o *BootstrappingInfo) GetNrfFeatures() map[string]string`

GetNrfFeatures returns the NrfFeatures field if non-nil, zero value otherwise.

### GetNrfFeaturesOk

`func (o *BootstrappingInfo) GetNrfFeaturesOk() (*map[string]string, bool)`

GetNrfFeaturesOk returns a tuple with the NrfFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNrfFeatures

`func (o *BootstrappingInfo) SetNrfFeatures(v map[string]string)`

SetNrfFeatures sets NrfFeatures field to given value.

### HasNrfFeatures

`func (o *BootstrappingInfo) HasNrfFeatures() bool`

HasNrfFeatures returns a boolean if a field has been set.

### GetOauth2Required

`func (o *BootstrappingInfo) GetOauth2Required() map[string]bool`

GetOauth2Required returns the Oauth2Required field if non-nil, zero value otherwise.

### GetOauth2RequiredOk

`func (o *BootstrappingInfo) GetOauth2RequiredOk() (*map[string]bool, bool)`

GetOauth2RequiredOk returns a tuple with the Oauth2Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauth2Required

`func (o *BootstrappingInfo) SetOauth2Required(v map[string]bool)`

SetOauth2Required sets Oauth2Required field to given value.

### HasOauth2Required

`func (o *BootstrappingInfo) HasOauth2Required() bool`

HasOauth2Required returns a boolean if a field has been set.

### GetNrfSetId

`func (o *BootstrappingInfo) GetNrfSetId() string`

GetNrfSetId returns the NrfSetId field if non-nil, zero value otherwise.

### GetNrfSetIdOk

`func (o *BootstrappingInfo) GetNrfSetIdOk() (*string, bool)`

GetNrfSetIdOk returns a tuple with the NrfSetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNrfSetId

`func (o *BootstrappingInfo) SetNrfSetId(v string)`

SetNrfSetId sets NrfSetId field to given value.

### HasNrfSetId

`func (o *BootstrappingInfo) HasNrfSetId() bool`

HasNrfSetId returns a boolean if a field has been set.

### GetNrfInstanceId

`func (o *BootstrappingInfo) GetNrfInstanceId() string`

GetNrfInstanceId returns the NrfInstanceId field if non-nil, zero value otherwise.

### GetNrfInstanceIdOk

`func (o *BootstrappingInfo) GetNrfInstanceIdOk() (*string, bool)`

GetNrfInstanceIdOk returns a tuple with the NrfInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNrfInstanceId

`func (o *BootstrappingInfo) SetNrfInstanceId(v string)`

SetNrfInstanceId sets NrfInstanceId field to given value.

### HasNrfInstanceId

`func (o *BootstrappingInfo) HasNrfInstanceId() bool`

HasNrfInstanceId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


