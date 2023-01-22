# UeCommunicationCollection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Gpsi** | Pointer to **string** | String identifying a Gpsi shall contain either an External Id or an MSISDN.  It shall be formatted as follows -External Identifier&#x3D; \&quot;extid-&#39;extid&#39;, where &#39;extid&#39;  shall be formatted according to clause 19.7.2 of 3GPP TS 23.003 that describes an  External Identifier.   | [optional] 
**Supi** | Pointer to **string** | String identifying a Supi that shall contain either an IMSI, a network specific identifier, a Global Cable Identifier (GCI) or a Global Line Identifier (GLI) as specified in clause  2.2A of 3GPP TS 23.003. It shall be formatted as follows  - for an IMSI \&quot;imsi-&lt;imsi&gt;\&quot;, where &lt;imsi&gt; shall be formatted according to clause 2.2    of 3GPP TS 23.003 that describes an IMSI.  - for a network specific identifier \&quot;nai-&lt;nai&gt;, where &lt;nai&gt; shall be formatted    according to clause 28.7.2 of 3GPP TS 23.003 that describes an NAI.  - for a GCI \&quot;gci-&lt;gci&gt;\&quot;, where &lt;gci&gt; shall be formatted according to clause 28.15.2    of 3GPP TS 23.003.  - for a GLI \&quot;gli-&lt;gli&gt;\&quot;, where &lt;gli&gt; shall be formatted according to clause 28.16.2 of    3GPP TS 23.003.To enable that the value is used as part of an URI, the string shall    only contain characters allowed according to the \&quot;lower-with-hyphen\&quot; naming convention    defined in 3GPP TS 29.501.  | [optional] 
**ExterGroupId** | Pointer to **string** |  | [optional] 
**InterGroupId** | Pointer to **string** | String identifying a group of devices network internal globally unique ID which identifies a set of IMSIs, as specified in clause 19.9 of 3GPP TS 23.003.   | [optional] 
**AppId** | **string** | String providing an application identifier. | 
**Comms** | [**[]CommunicationCollection**](CommunicationCollection.md) |  | 

## Methods

### NewUeCommunicationCollection

`func NewUeCommunicationCollection(appId string, comms []CommunicationCollection, ) *UeCommunicationCollection`

NewUeCommunicationCollection instantiates a new UeCommunicationCollection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUeCommunicationCollectionWithDefaults

`func NewUeCommunicationCollectionWithDefaults() *UeCommunicationCollection`

NewUeCommunicationCollectionWithDefaults instantiates a new UeCommunicationCollection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGpsi

`func (o *UeCommunicationCollection) GetGpsi() string`

GetGpsi returns the Gpsi field if non-nil, zero value otherwise.

### GetGpsiOk

`func (o *UeCommunicationCollection) GetGpsiOk() (*string, bool)`

GetGpsiOk returns a tuple with the Gpsi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpsi

`func (o *UeCommunicationCollection) SetGpsi(v string)`

SetGpsi sets Gpsi field to given value.

### HasGpsi

`func (o *UeCommunicationCollection) HasGpsi() bool`

HasGpsi returns a boolean if a field has been set.

### GetSupi

`func (o *UeCommunicationCollection) GetSupi() string`

GetSupi returns the Supi field if non-nil, zero value otherwise.

### GetSupiOk

`func (o *UeCommunicationCollection) GetSupiOk() (*string, bool)`

GetSupiOk returns a tuple with the Supi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupi

`func (o *UeCommunicationCollection) SetSupi(v string)`

SetSupi sets Supi field to given value.

### HasSupi

`func (o *UeCommunicationCollection) HasSupi() bool`

HasSupi returns a boolean if a field has been set.

### GetExterGroupId

`func (o *UeCommunicationCollection) GetExterGroupId() string`

GetExterGroupId returns the ExterGroupId field if non-nil, zero value otherwise.

### GetExterGroupIdOk

`func (o *UeCommunicationCollection) GetExterGroupIdOk() (*string, bool)`

GetExterGroupIdOk returns a tuple with the ExterGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExterGroupId

`func (o *UeCommunicationCollection) SetExterGroupId(v string)`

SetExterGroupId sets ExterGroupId field to given value.

### HasExterGroupId

`func (o *UeCommunicationCollection) HasExterGroupId() bool`

HasExterGroupId returns a boolean if a field has been set.

### GetInterGroupId

`func (o *UeCommunicationCollection) GetInterGroupId() string`

GetInterGroupId returns the InterGroupId field if non-nil, zero value otherwise.

### GetInterGroupIdOk

`func (o *UeCommunicationCollection) GetInterGroupIdOk() (*string, bool)`

GetInterGroupIdOk returns a tuple with the InterGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterGroupId

`func (o *UeCommunicationCollection) SetInterGroupId(v string)`

SetInterGroupId sets InterGroupId field to given value.

### HasInterGroupId

`func (o *UeCommunicationCollection) HasInterGroupId() bool`

HasInterGroupId returns a boolean if a field has been set.

### GetAppId

`func (o *UeCommunicationCollection) GetAppId() string`

GetAppId returns the AppId field if non-nil, zero value otherwise.

### GetAppIdOk

`func (o *UeCommunicationCollection) GetAppIdOk() (*string, bool)`

GetAppIdOk returns a tuple with the AppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppId

`func (o *UeCommunicationCollection) SetAppId(v string)`

SetAppId sets AppId field to given value.


### GetComms

`func (o *UeCommunicationCollection) GetComms() []CommunicationCollection`

GetComms returns the Comms field if non-nil, zero value otherwise.

### GetCommsOk

`func (o *UeCommunicationCollection) GetCommsOk() (*[]CommunicationCollection, bool)`

GetCommsOk returns a tuple with the Comms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComms

`func (o *UeCommunicationCollection) SetComms(v []CommunicationCollection)`

SetComms sets Comms field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


