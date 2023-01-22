# OriginatorInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OriginatorSUPI** | Pointer to **string** | String identifying a Supi that shall contain either an IMSI, a network specific identifier, a Global Cable Identifier (GCI) or a Global Line Identifier (GLI) as specified in clause  2.2A of 3GPP TS 23.003. It shall be formatted as follows  - for an IMSI \&quot;imsi-&lt;imsi&gt;\&quot;, where &lt;imsi&gt; shall be formatted according to clause 2.2    of 3GPP TS 23.003 that describes an IMSI.  - for a network specific identifier \&quot;nai-&lt;nai&gt;, where &lt;nai&gt; shall be formatted    according to clause 28.7.2 of 3GPP TS 23.003 that describes an NAI.  - for a GCI \&quot;gci-&lt;gci&gt;\&quot;, where &lt;gci&gt; shall be formatted according to clause 28.15.2    of 3GPP TS 23.003.  - for a GLI \&quot;gli-&lt;gli&gt;\&quot;, where &lt;gli&gt; shall be formatted according to clause 28.16.2 of    3GPP TS 23.003.To enable that the value is used as part of an URI, the string shall    only contain characters allowed according to the \&quot;lower-with-hyphen\&quot; naming convention    defined in 3GPP TS 29.501.  | [optional] 
**OriginatorGPSI** | Pointer to **string** | String identifying a Gpsi shall contain either an External Id or an MSISDN.  It shall be formatted as follows -External Identifier&#x3D; \&quot;extid-&#39;extid&#39;, where &#39;extid&#39;  shall be formatted according to clause 19.7.2 of 3GPP TS 23.003 that describes an  External Identifier.   | [optional] 
**OriginatorOtherAddress** | Pointer to [**SMAddressInfo**](SMAddressInfo.md) |  | [optional] 
**OriginatorReceivedAddress** | Pointer to [**SMAddressInfo**](SMAddressInfo.md) |  | [optional] 
**OriginatorSCCPAddress** | Pointer to **string** |  | [optional] 
**SMOriginatorInterface** | Pointer to [**SMInterface**](SMInterface.md) |  | [optional] 
**SMOriginatorProtocolId** | Pointer to **string** |  | [optional] 

## Methods

### NewOriginatorInfo

`func NewOriginatorInfo() *OriginatorInfo`

NewOriginatorInfo instantiates a new OriginatorInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOriginatorInfoWithDefaults

`func NewOriginatorInfoWithDefaults() *OriginatorInfo`

NewOriginatorInfoWithDefaults instantiates a new OriginatorInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOriginatorSUPI

`func (o *OriginatorInfo) GetOriginatorSUPI() string`

GetOriginatorSUPI returns the OriginatorSUPI field if non-nil, zero value otherwise.

### GetOriginatorSUPIOk

`func (o *OriginatorInfo) GetOriginatorSUPIOk() (*string, bool)`

GetOriginatorSUPIOk returns a tuple with the OriginatorSUPI field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginatorSUPI

`func (o *OriginatorInfo) SetOriginatorSUPI(v string)`

SetOriginatorSUPI sets OriginatorSUPI field to given value.

### HasOriginatorSUPI

`func (o *OriginatorInfo) HasOriginatorSUPI() bool`

HasOriginatorSUPI returns a boolean if a field has been set.

### GetOriginatorGPSI

`func (o *OriginatorInfo) GetOriginatorGPSI() string`

GetOriginatorGPSI returns the OriginatorGPSI field if non-nil, zero value otherwise.

### GetOriginatorGPSIOk

`func (o *OriginatorInfo) GetOriginatorGPSIOk() (*string, bool)`

GetOriginatorGPSIOk returns a tuple with the OriginatorGPSI field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginatorGPSI

`func (o *OriginatorInfo) SetOriginatorGPSI(v string)`

SetOriginatorGPSI sets OriginatorGPSI field to given value.

### HasOriginatorGPSI

`func (o *OriginatorInfo) HasOriginatorGPSI() bool`

HasOriginatorGPSI returns a boolean if a field has been set.

### GetOriginatorOtherAddress

`func (o *OriginatorInfo) GetOriginatorOtherAddress() SMAddressInfo`

GetOriginatorOtherAddress returns the OriginatorOtherAddress field if non-nil, zero value otherwise.

### GetOriginatorOtherAddressOk

`func (o *OriginatorInfo) GetOriginatorOtherAddressOk() (*SMAddressInfo, bool)`

GetOriginatorOtherAddressOk returns a tuple with the OriginatorOtherAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginatorOtherAddress

`func (o *OriginatorInfo) SetOriginatorOtherAddress(v SMAddressInfo)`

SetOriginatorOtherAddress sets OriginatorOtherAddress field to given value.

### HasOriginatorOtherAddress

`func (o *OriginatorInfo) HasOriginatorOtherAddress() bool`

HasOriginatorOtherAddress returns a boolean if a field has been set.

### GetOriginatorReceivedAddress

`func (o *OriginatorInfo) GetOriginatorReceivedAddress() SMAddressInfo`

GetOriginatorReceivedAddress returns the OriginatorReceivedAddress field if non-nil, zero value otherwise.

### GetOriginatorReceivedAddressOk

`func (o *OriginatorInfo) GetOriginatorReceivedAddressOk() (*SMAddressInfo, bool)`

GetOriginatorReceivedAddressOk returns a tuple with the OriginatorReceivedAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginatorReceivedAddress

`func (o *OriginatorInfo) SetOriginatorReceivedAddress(v SMAddressInfo)`

SetOriginatorReceivedAddress sets OriginatorReceivedAddress field to given value.

### HasOriginatorReceivedAddress

`func (o *OriginatorInfo) HasOriginatorReceivedAddress() bool`

HasOriginatorReceivedAddress returns a boolean if a field has been set.

### GetOriginatorSCCPAddress

`func (o *OriginatorInfo) GetOriginatorSCCPAddress() string`

GetOriginatorSCCPAddress returns the OriginatorSCCPAddress field if non-nil, zero value otherwise.

### GetOriginatorSCCPAddressOk

`func (o *OriginatorInfo) GetOriginatorSCCPAddressOk() (*string, bool)`

GetOriginatorSCCPAddressOk returns a tuple with the OriginatorSCCPAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginatorSCCPAddress

`func (o *OriginatorInfo) SetOriginatorSCCPAddress(v string)`

SetOriginatorSCCPAddress sets OriginatorSCCPAddress field to given value.

### HasOriginatorSCCPAddress

`func (o *OriginatorInfo) HasOriginatorSCCPAddress() bool`

HasOriginatorSCCPAddress returns a boolean if a field has been set.

### GetSMOriginatorInterface

`func (o *OriginatorInfo) GetSMOriginatorInterface() SMInterface`

GetSMOriginatorInterface returns the SMOriginatorInterface field if non-nil, zero value otherwise.

### GetSMOriginatorInterfaceOk

`func (o *OriginatorInfo) GetSMOriginatorInterfaceOk() (*SMInterface, bool)`

GetSMOriginatorInterfaceOk returns a tuple with the SMOriginatorInterface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSMOriginatorInterface

`func (o *OriginatorInfo) SetSMOriginatorInterface(v SMInterface)`

SetSMOriginatorInterface sets SMOriginatorInterface field to given value.

### HasSMOriginatorInterface

`func (o *OriginatorInfo) HasSMOriginatorInterface() bool`

HasSMOriginatorInterface returns a boolean if a field has been set.

### GetSMOriginatorProtocolId

`func (o *OriginatorInfo) GetSMOriginatorProtocolId() string`

GetSMOriginatorProtocolId returns the SMOriginatorProtocolId field if non-nil, zero value otherwise.

### GetSMOriginatorProtocolIdOk

`func (o *OriginatorInfo) GetSMOriginatorProtocolIdOk() (*string, bool)`

GetSMOriginatorProtocolIdOk returns a tuple with the SMOriginatorProtocolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSMOriginatorProtocolId

`func (o *OriginatorInfo) SetSMOriginatorProtocolId(v string)`

SetSMOriginatorProtocolId sets SMOriginatorProtocolId field to given value.

### HasSMOriginatorProtocolId

`func (o *OriginatorInfo) HasSMOriginatorProtocolId() bool`

HasSMOriginatorProtocolId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


