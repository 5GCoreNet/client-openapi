# UeMobilityCollection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Gpsi** | Pointer to **string** | String identifying a Gpsi shall contain either an External Id or an MSISDN.  It shall be formatted as follows -External Identifier&#x3D; \&quot;extid-&#39;extid&#39;, where &#39;extid&#39;  shall be formatted according to clause 19.7.2 of 3GPP TS 23.003 that describes an  External Identifier.   | [optional] 
**Supi** | Pointer to **string** | String identifying a Supi that shall contain either an IMSI, a network specific identifier, a Global Cable Identifier (GCI) or a Global Line Identifier (GLI) as specified in clause  2.2A of 3GPP TS 23.003. It shall be formatted as follows  - for an IMSI \&quot;imsi-&lt;imsi&gt;\&quot;, where &lt;imsi&gt; shall be formatted according to clause 2.2    of 3GPP TS 23.003 that describes an IMSI.  - for a network specific identifier \&quot;nai-&lt;nai&gt;, where &lt;nai&gt; shall be formatted    according to clause 28.7.2 of 3GPP TS 23.003 that describes an NAI.  - for a GCI \&quot;gci-&lt;gci&gt;\&quot;, where &lt;gci&gt; shall be formatted according to clause 28.15.2    of 3GPP TS 23.003.  - for a GLI \&quot;gli-&lt;gli&gt;\&quot;, where &lt;gli&gt; shall be formatted according to clause 28.16.2 of    3GPP TS 23.003.To enable that the value is used as part of an URI, the string shall    only contain characters allowed according to the \&quot;lower-with-hyphen\&quot; naming convention    defined in 3GPP TS 29.501.  | [optional] 
**AppId** | **string** | String providing an application identifier. | 
**UeTrajs** | [**[]UeTrajectoryCollection**](UeTrajectoryCollection.md) |  | 

## Methods

### NewUeMobilityCollection

`func NewUeMobilityCollection(appId string, ueTrajs []UeTrajectoryCollection, ) *UeMobilityCollection`

NewUeMobilityCollection instantiates a new UeMobilityCollection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUeMobilityCollectionWithDefaults

`func NewUeMobilityCollectionWithDefaults() *UeMobilityCollection`

NewUeMobilityCollectionWithDefaults instantiates a new UeMobilityCollection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGpsi

`func (o *UeMobilityCollection) GetGpsi() string`

GetGpsi returns the Gpsi field if non-nil, zero value otherwise.

### GetGpsiOk

`func (o *UeMobilityCollection) GetGpsiOk() (*string, bool)`

GetGpsiOk returns a tuple with the Gpsi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpsi

`func (o *UeMobilityCollection) SetGpsi(v string)`

SetGpsi sets Gpsi field to given value.

### HasGpsi

`func (o *UeMobilityCollection) HasGpsi() bool`

HasGpsi returns a boolean if a field has been set.

### GetSupi

`func (o *UeMobilityCollection) GetSupi() string`

GetSupi returns the Supi field if non-nil, zero value otherwise.

### GetSupiOk

`func (o *UeMobilityCollection) GetSupiOk() (*string, bool)`

GetSupiOk returns a tuple with the Supi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupi

`func (o *UeMobilityCollection) SetSupi(v string)`

SetSupi sets Supi field to given value.

### HasSupi

`func (o *UeMobilityCollection) HasSupi() bool`

HasSupi returns a boolean if a field has been set.

### GetAppId

`func (o *UeMobilityCollection) GetAppId() string`

GetAppId returns the AppId field if non-nil, zero value otherwise.

### GetAppIdOk

`func (o *UeMobilityCollection) GetAppIdOk() (*string, bool)`

GetAppIdOk returns a tuple with the AppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppId

`func (o *UeMobilityCollection) SetAppId(v string)`

SetAppId sets AppId field to given value.


### GetUeTrajs

`func (o *UeMobilityCollection) GetUeTrajs() []UeTrajectoryCollection`

GetUeTrajs returns the UeTrajs field if non-nil, zero value otherwise.

### GetUeTrajsOk

`func (o *UeMobilityCollection) GetUeTrajsOk() (*[]UeTrajectoryCollection, bool)`

GetUeTrajsOk returns a tuple with the UeTrajs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUeTrajs

`func (o *UeMobilityCollection) SetUeTrajs(v []UeTrajectoryCollection)`

SetUeTrajs sets UeTrajs field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


