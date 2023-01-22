# ProseContextInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Supi** | **string** | String identifying a Supi that shall contain either an IMSI, a network specific identifier, a Global Cable Identifier (GCI) or a Global Line Identifier (GLI) as specified in clause  2.2A of 3GPP TS 23.003. It shall be formatted as follows  - for an IMSI \&quot;imsi-&lt;imsi&gt;\&quot;, where &lt;imsi&gt; shall be formatted according to clause 2.2    of 3GPP TS 23.003 that describes an IMSI.  - for a network specific identifier \&quot;nai-&lt;nai&gt;, where &lt;nai&gt; shall be formatted    according to clause 28.7.2 of 3GPP TS 23.003 that describes an NAI.  - for a GCI \&quot;gci-&lt;gci&gt;\&quot;, where &lt;gci&gt; shall be formatted according to clause 28.15.2    of 3GPP TS 23.003.  - for a GLI \&quot;gli-&lt;gli&gt;\&quot;, where &lt;gli&gt; shall be formatted according to clause 28.16.2 of    3GPP TS 23.003.To enable that the value is used as part of an URI, the string shall    only contain characters allowed according to the \&quot;lower-with-hyphen\&quot; naming convention    defined in 3GPP TS 29.501.  | 
**Var5gPruk** | **string** | 5G Prose Remote User Key | 
**Var5gPrukId** | **string** | The 5GPRUK ID is string in NAI format as specified in clause 28.7.19 of 3GPP TS 23.003.  | 
**RelayServiceCode** | **int32** | Relay Service Code to identify a connectivity service provided by the UE-to-Network relay.  | 

## Methods

### NewProseContextInfo

`func NewProseContextInfo(supi string, var5gPruk string, var5gPrukId string, relayServiceCode int32, ) *ProseContextInfo`

NewProseContextInfo instantiates a new ProseContextInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProseContextInfoWithDefaults

`func NewProseContextInfoWithDefaults() *ProseContextInfo`

NewProseContextInfoWithDefaults instantiates a new ProseContextInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSupi

`func (o *ProseContextInfo) GetSupi() string`

GetSupi returns the Supi field if non-nil, zero value otherwise.

### GetSupiOk

`func (o *ProseContextInfo) GetSupiOk() (*string, bool)`

GetSupiOk returns a tuple with the Supi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupi

`func (o *ProseContextInfo) SetSupi(v string)`

SetSupi sets Supi field to given value.


### GetVar5gPruk

`func (o *ProseContextInfo) GetVar5gPruk() string`

GetVar5gPruk returns the Var5gPruk field if non-nil, zero value otherwise.

### GetVar5gPrukOk

`func (o *ProseContextInfo) GetVar5gPrukOk() (*string, bool)`

GetVar5gPrukOk returns a tuple with the Var5gPruk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar5gPruk

`func (o *ProseContextInfo) SetVar5gPruk(v string)`

SetVar5gPruk sets Var5gPruk field to given value.


### GetVar5gPrukId

`func (o *ProseContextInfo) GetVar5gPrukId() string`

GetVar5gPrukId returns the Var5gPrukId field if non-nil, zero value otherwise.

### GetVar5gPrukIdOk

`func (o *ProseContextInfo) GetVar5gPrukIdOk() (*string, bool)`

GetVar5gPrukIdOk returns a tuple with the Var5gPrukId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar5gPrukId

`func (o *ProseContextInfo) SetVar5gPrukId(v string)`

SetVar5gPrukId sets Var5gPrukId field to given value.


### GetRelayServiceCode

`func (o *ProseContextInfo) GetRelayServiceCode() int32`

GetRelayServiceCode returns the RelayServiceCode field if non-nil, zero value otherwise.

### GetRelayServiceCodeOk

`func (o *ProseContextInfo) GetRelayServiceCodeOk() (*int32, bool)`

GetRelayServiceCodeOk returns a tuple with the RelayServiceCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelayServiceCode

`func (o *ProseContextInfo) SetRelayServiceCode(v int32)`

SetRelayServiceCode sets RelayServiceCode field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


