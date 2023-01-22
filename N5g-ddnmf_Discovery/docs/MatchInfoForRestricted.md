# MatchInfoForRestricted

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Supi** | **string** | String identifying a Supi that shall contain either an IMSI, a network specific identifier, a Global Cable Identifier (GCI) or a Global Line Identifier (GLI) as specified in clause  2.2A of 3GPP TS 23.003. It shall be formatted as follows  - for an IMSI \&quot;imsi-&lt;imsi&gt;\&quot;, where &lt;imsi&gt; shall be formatted according to clause 2.2    of 3GPP TS 23.003 that describes an IMSI.  - for a network specific identifier \&quot;nai-&lt;nai&gt;, where &lt;nai&gt; shall be formatted    according to clause 28.7.2 of 3GPP TS 23.003 that describes an NAI.  - for a GCI \&quot;gci-&lt;gci&gt;\&quot;, where &lt;gci&gt; shall be formatted according to clause 28.15.2    of 3GPP TS 23.003.  - for a GLI \&quot;gli-&lt;gli&gt;\&quot;, where &lt;gli&gt; shall be formatted according to clause 28.16.2 of    3GPP TS 23.003.To enable that the value is used as part of an URI, the string shall    only contain characters allowed according to the \&quot;lower-with-hyphen\&quot; naming convention    defined in 3GPP TS 29.501.  | 
**Rpauid** | **string** | Contains the RPAUID. | 
**TargetRpauid** | **string** | Contains the RPAUID. | 
**ProseRestrictedCode** | **string** | Contains the ProSe Restricted Code. | 

## Methods

### NewMatchInfoForRestricted

`func NewMatchInfoForRestricted(supi string, rpauid string, targetRpauid string, proseRestrictedCode string, ) *MatchInfoForRestricted`

NewMatchInfoForRestricted instantiates a new MatchInfoForRestricted object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMatchInfoForRestrictedWithDefaults

`func NewMatchInfoForRestrictedWithDefaults() *MatchInfoForRestricted`

NewMatchInfoForRestrictedWithDefaults instantiates a new MatchInfoForRestricted object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSupi

`func (o *MatchInfoForRestricted) GetSupi() string`

GetSupi returns the Supi field if non-nil, zero value otherwise.

### GetSupiOk

`func (o *MatchInfoForRestricted) GetSupiOk() (*string, bool)`

GetSupiOk returns a tuple with the Supi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupi

`func (o *MatchInfoForRestricted) SetSupi(v string)`

SetSupi sets Supi field to given value.


### GetRpauid

`func (o *MatchInfoForRestricted) GetRpauid() string`

GetRpauid returns the Rpauid field if non-nil, zero value otherwise.

### GetRpauidOk

`func (o *MatchInfoForRestricted) GetRpauidOk() (*string, bool)`

GetRpauidOk returns a tuple with the Rpauid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRpauid

`func (o *MatchInfoForRestricted) SetRpauid(v string)`

SetRpauid sets Rpauid field to given value.


### GetTargetRpauid

`func (o *MatchInfoForRestricted) GetTargetRpauid() string`

GetTargetRpauid returns the TargetRpauid field if non-nil, zero value otherwise.

### GetTargetRpauidOk

`func (o *MatchInfoForRestricted) GetTargetRpauidOk() (*string, bool)`

GetTargetRpauidOk returns a tuple with the TargetRpauid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetRpauid

`func (o *MatchInfoForRestricted) SetTargetRpauid(v string)`

SetTargetRpauid sets TargetRpauid field to given value.


### GetProseRestrictedCode

`func (o *MatchInfoForRestricted) GetProseRestrictedCode() string`

GetProseRestrictedCode returns the ProseRestrictedCode field if non-nil, zero value otherwise.

### GetProseRestrictedCodeOk

`func (o *MatchInfoForRestricted) GetProseRestrictedCodeOk() (*string, bool)`

GetProseRestrictedCodeOk returns a tuple with the ProseRestrictedCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProseRestrictedCode

`func (o *MatchInfoForRestricted) SetProseRestrictedCode(v string)`

SetProseRestrictedCode sets ProseRestrictedCode field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


