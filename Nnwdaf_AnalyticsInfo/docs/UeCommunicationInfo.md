# UeCommunicationInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Supi** | Pointer to **string** | String identifying a Supi that shall contain either an IMSI, a network specific identifier, a Global Cable Identifier (GCI) or a Global Line Identifier (GLI) as specified in clause  2.2A of 3GPP TS 23.003. It shall be formatted as follows  - for an IMSI \&quot;imsi-&lt;imsi&gt;\&quot;, where &lt;imsi&gt; shall be formatted according to clause 2.2    of 3GPP TS 23.003 that describes an IMSI.  - for a network specific identifier \&quot;nai-&lt;nai&gt;, where &lt;nai&gt; shall be formatted    according to clause 28.7.2 of 3GPP TS 23.003 that describes an NAI.  - for a GCI \&quot;gci-&lt;gci&gt;\&quot;, where &lt;gci&gt; shall be formatted according to clause 28.15.2    of 3GPP TS 23.003.  - for a GLI \&quot;gli-&lt;gli&gt;\&quot;, where &lt;gli&gt; shall be formatted according to clause 28.16.2 of    3GPP TS 23.003.To enable that the value is used as part of an URI, the string shall    only contain characters allowed according to the \&quot;lower-with-hyphen\&quot; naming convention    defined in 3GPP TS 29.501.  | [optional] 
**InterGroupId** | Pointer to **string** | String identifying a group of devices network internal globally unique ID which identifies a set of IMSIs, as specified in clause 19.9 of 3GPP TS 23.003.   | [optional] 
**AppId** | Pointer to **string** | String providing an application identifier. | [optional] 
**Comms** | [**[]CommunicationCollection**](CommunicationCollection.md) |  | 

## Methods

### NewUeCommunicationInfo

`func NewUeCommunicationInfo(comms []CommunicationCollection, ) *UeCommunicationInfo`

NewUeCommunicationInfo instantiates a new UeCommunicationInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUeCommunicationInfoWithDefaults

`func NewUeCommunicationInfoWithDefaults() *UeCommunicationInfo`

NewUeCommunicationInfoWithDefaults instantiates a new UeCommunicationInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSupi

`func (o *UeCommunicationInfo) GetSupi() string`

GetSupi returns the Supi field if non-nil, zero value otherwise.

### GetSupiOk

`func (o *UeCommunicationInfo) GetSupiOk() (*string, bool)`

GetSupiOk returns a tuple with the Supi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupi

`func (o *UeCommunicationInfo) SetSupi(v string)`

SetSupi sets Supi field to given value.

### HasSupi

`func (o *UeCommunicationInfo) HasSupi() bool`

HasSupi returns a boolean if a field has been set.

### GetInterGroupId

`func (o *UeCommunicationInfo) GetInterGroupId() string`

GetInterGroupId returns the InterGroupId field if non-nil, zero value otherwise.

### GetInterGroupIdOk

`func (o *UeCommunicationInfo) GetInterGroupIdOk() (*string, bool)`

GetInterGroupIdOk returns a tuple with the InterGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterGroupId

`func (o *UeCommunicationInfo) SetInterGroupId(v string)`

SetInterGroupId sets InterGroupId field to given value.

### HasInterGroupId

`func (o *UeCommunicationInfo) HasInterGroupId() bool`

HasInterGroupId returns a boolean if a field has been set.

### GetAppId

`func (o *UeCommunicationInfo) GetAppId() string`

GetAppId returns the AppId field if non-nil, zero value otherwise.

### GetAppIdOk

`func (o *UeCommunicationInfo) GetAppIdOk() (*string, bool)`

GetAppIdOk returns a tuple with the AppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppId

`func (o *UeCommunicationInfo) SetAppId(v string)`

SetAppId sets AppId field to given value.

### HasAppId

`func (o *UeCommunicationInfo) HasAppId() bool`

HasAppId returns a boolean if a field has been set.

### GetComms

`func (o *UeCommunicationInfo) GetComms() []CommunicationCollection`

GetComms returns the Comms field if non-nil, zero value otherwise.

### GetCommsOk

`func (o *UeCommunicationInfo) GetCommsOk() (*[]CommunicationCollection, bool)`

GetCommsOk returns a tuple with the Comms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComms

`func (o *UeCommunicationInfo) SetComms(v []CommunicationCollection)`

SetComms sets Comms field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


