# RecipientInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RecipientSUPI** | Pointer to **string** | String identifying a Supi that shall contain either an IMSI, a network specific identifier, a Global Cable Identifier (GCI) or a Global Line Identifier (GLI) as specified in clause  2.2A of 3GPP TS 23.003. It shall be formatted as follows  - for an IMSI \&quot;imsi-&lt;imsi&gt;\&quot;, where &lt;imsi&gt; shall be formatted according to clause 2.2    of 3GPP TS 23.003 that describes an IMSI.  - for a network specific identifier \&quot;nai-&lt;nai&gt;, where &lt;nai&gt; shall be formatted    according to clause 28.7.2 of 3GPP TS 23.003 that describes an NAI.  - for a GCI \&quot;gci-&lt;gci&gt;\&quot;, where &lt;gci&gt; shall be formatted according to clause 28.15.2    of 3GPP TS 23.003.  - for a GLI \&quot;gli-&lt;gli&gt;\&quot;, where &lt;gli&gt; shall be formatted according to clause 28.16.2 of    3GPP TS 23.003.To enable that the value is used as part of an URI, the string shall    only contain characters allowed according to the \&quot;lower-with-hyphen\&quot; naming convention    defined in 3GPP TS 29.501.  | [optional] 
**RecipientGPSI** | Pointer to **string** | String identifying a Gpsi shall contain either an External Id or an MSISDN.  It shall be formatted as follows -External Identifier&#x3D; \&quot;extid-&#39;extid&#39;, where &#39;extid&#39;  shall be formatted according to clause 19.7.2 of 3GPP TS 23.003 that describes an  External Identifier.   | [optional] 
**RecipientOtherAddress** | Pointer to [**SMAddressInfo**](SMAddressInfo.md) |  | [optional] 
**RecipientReceivedAddress** | Pointer to [**SMAddressInfo**](SMAddressInfo.md) |  | [optional] 
**RecipientSCCPAddress** | Pointer to **string** |  | [optional] 
**SMDestinationInterface** | Pointer to [**SMInterface**](SMInterface.md) |  | [optional] 
**SMrecipientProtocolId** | Pointer to **string** |  | [optional] 

## Methods

### NewRecipientInfo

`func NewRecipientInfo() *RecipientInfo`

NewRecipientInfo instantiates a new RecipientInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecipientInfoWithDefaults

`func NewRecipientInfoWithDefaults() *RecipientInfo`

NewRecipientInfoWithDefaults instantiates a new RecipientInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRecipientSUPI

`func (o *RecipientInfo) GetRecipientSUPI() string`

GetRecipientSUPI returns the RecipientSUPI field if non-nil, zero value otherwise.

### GetRecipientSUPIOk

`func (o *RecipientInfo) GetRecipientSUPIOk() (*string, bool)`

GetRecipientSUPIOk returns a tuple with the RecipientSUPI field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipientSUPI

`func (o *RecipientInfo) SetRecipientSUPI(v string)`

SetRecipientSUPI sets RecipientSUPI field to given value.

### HasRecipientSUPI

`func (o *RecipientInfo) HasRecipientSUPI() bool`

HasRecipientSUPI returns a boolean if a field has been set.

### GetRecipientGPSI

`func (o *RecipientInfo) GetRecipientGPSI() string`

GetRecipientGPSI returns the RecipientGPSI field if non-nil, zero value otherwise.

### GetRecipientGPSIOk

`func (o *RecipientInfo) GetRecipientGPSIOk() (*string, bool)`

GetRecipientGPSIOk returns a tuple with the RecipientGPSI field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipientGPSI

`func (o *RecipientInfo) SetRecipientGPSI(v string)`

SetRecipientGPSI sets RecipientGPSI field to given value.

### HasRecipientGPSI

`func (o *RecipientInfo) HasRecipientGPSI() bool`

HasRecipientGPSI returns a boolean if a field has been set.

### GetRecipientOtherAddress

`func (o *RecipientInfo) GetRecipientOtherAddress() SMAddressInfo`

GetRecipientOtherAddress returns the RecipientOtherAddress field if non-nil, zero value otherwise.

### GetRecipientOtherAddressOk

`func (o *RecipientInfo) GetRecipientOtherAddressOk() (*SMAddressInfo, bool)`

GetRecipientOtherAddressOk returns a tuple with the RecipientOtherAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipientOtherAddress

`func (o *RecipientInfo) SetRecipientOtherAddress(v SMAddressInfo)`

SetRecipientOtherAddress sets RecipientOtherAddress field to given value.

### HasRecipientOtherAddress

`func (o *RecipientInfo) HasRecipientOtherAddress() bool`

HasRecipientOtherAddress returns a boolean if a field has been set.

### GetRecipientReceivedAddress

`func (o *RecipientInfo) GetRecipientReceivedAddress() SMAddressInfo`

GetRecipientReceivedAddress returns the RecipientReceivedAddress field if non-nil, zero value otherwise.

### GetRecipientReceivedAddressOk

`func (o *RecipientInfo) GetRecipientReceivedAddressOk() (*SMAddressInfo, bool)`

GetRecipientReceivedAddressOk returns a tuple with the RecipientReceivedAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipientReceivedAddress

`func (o *RecipientInfo) SetRecipientReceivedAddress(v SMAddressInfo)`

SetRecipientReceivedAddress sets RecipientReceivedAddress field to given value.

### HasRecipientReceivedAddress

`func (o *RecipientInfo) HasRecipientReceivedAddress() bool`

HasRecipientReceivedAddress returns a boolean if a field has been set.

### GetRecipientSCCPAddress

`func (o *RecipientInfo) GetRecipientSCCPAddress() string`

GetRecipientSCCPAddress returns the RecipientSCCPAddress field if non-nil, zero value otherwise.

### GetRecipientSCCPAddressOk

`func (o *RecipientInfo) GetRecipientSCCPAddressOk() (*string, bool)`

GetRecipientSCCPAddressOk returns a tuple with the RecipientSCCPAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipientSCCPAddress

`func (o *RecipientInfo) SetRecipientSCCPAddress(v string)`

SetRecipientSCCPAddress sets RecipientSCCPAddress field to given value.

### HasRecipientSCCPAddress

`func (o *RecipientInfo) HasRecipientSCCPAddress() bool`

HasRecipientSCCPAddress returns a boolean if a field has been set.

### GetSMDestinationInterface

`func (o *RecipientInfo) GetSMDestinationInterface() SMInterface`

GetSMDestinationInterface returns the SMDestinationInterface field if non-nil, zero value otherwise.

### GetSMDestinationInterfaceOk

`func (o *RecipientInfo) GetSMDestinationInterfaceOk() (*SMInterface, bool)`

GetSMDestinationInterfaceOk returns a tuple with the SMDestinationInterface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSMDestinationInterface

`func (o *RecipientInfo) SetSMDestinationInterface(v SMInterface)`

SetSMDestinationInterface sets SMDestinationInterface field to given value.

### HasSMDestinationInterface

`func (o *RecipientInfo) HasSMDestinationInterface() bool`

HasSMDestinationInterface returns a boolean if a field has been set.

### GetSMrecipientProtocolId

`func (o *RecipientInfo) GetSMrecipientProtocolId() string`

GetSMrecipientProtocolId returns the SMrecipientProtocolId field if non-nil, zero value otherwise.

### GetSMrecipientProtocolIdOk

`func (o *RecipientInfo) GetSMrecipientProtocolIdOk() (*string, bool)`

GetSMrecipientProtocolIdOk returns a tuple with the SMrecipientProtocolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSMrecipientProtocolId

`func (o *RecipientInfo) SetSMrecipientProtocolId(v string)`

SetSMrecipientProtocolId sets SMrecipientProtocolId field to given value.

### HasSMrecipientProtocolId

`func (o *RecipientInfo) HasSMrecipientProtocolId() bool`

HasSMrecipientProtocolId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


