# PduACRequestInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Supi** | **string** | String identifying a Supi that shall contain either an IMSI, a network specific identifier, a Global Cable Identifier (GCI) or a Global Line Identifier (GLI) as specified in clause  2.2A of 3GPP TS 23.003. It shall be formatted as follows  - for an IMSI \&quot;imsi-&lt;imsi&gt;\&quot;, where &lt;imsi&gt; shall be formatted according to clause 2.2    of 3GPP TS 23.003 that describes an IMSI.  - for a network specific identifier \&quot;nai-&lt;nai&gt;, where &lt;nai&gt; shall be formatted    according to clause 28.7.2 of 3GPP TS 23.003 that describes an NAI.  - for a GCI \&quot;gci-&lt;gci&gt;\&quot;, where &lt;gci&gt; shall be formatted according to clause 28.15.2    of 3GPP TS 23.003.  - for a GLI \&quot;gli-&lt;gli&gt;\&quot;, where &lt;gli&gt; shall be formatted according to clause 28.16.2 of    3GPP TS 23.003.To enable that the value is used as part of an URI, the string shall    only contain characters allowed according to the \&quot;lower-with-hyphen\&quot; naming convention    defined in 3GPP TS 29.501.  | 
**AnType** | [**AccessType**](AccessType.md) |  | 
**PduSessionId** | **int32** | Unsigned integer identifying a PDU session, within the range 0 to 255, as specified in  clause 11.2.3.1b, bits 1 to 8, of 3GPP TS 24.007. If the PDU Session ID is allocated by the  Core Network for UEs not supporting N1 mode, reserved range 64 to 95 is used. PDU Session ID  within the reserved range is only visible in the Core Network.   | 
**AcuOperationList** | [**[]AcuOperationItem**](AcuOperationItem.md) |  | 
**AdditionalAnType** | Pointer to [**AccessType**](AccessType.md) |  | [optional] 

## Methods

### NewPduACRequestInfo

`func NewPduACRequestInfo(supi string, anType AccessType, pduSessionId int32, acuOperationList []AcuOperationItem, ) *PduACRequestInfo`

NewPduACRequestInfo instantiates a new PduACRequestInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPduACRequestInfoWithDefaults

`func NewPduACRequestInfoWithDefaults() *PduACRequestInfo`

NewPduACRequestInfoWithDefaults instantiates a new PduACRequestInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSupi

`func (o *PduACRequestInfo) GetSupi() string`

GetSupi returns the Supi field if non-nil, zero value otherwise.

### GetSupiOk

`func (o *PduACRequestInfo) GetSupiOk() (*string, bool)`

GetSupiOk returns a tuple with the Supi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupi

`func (o *PduACRequestInfo) SetSupi(v string)`

SetSupi sets Supi field to given value.


### GetAnType

`func (o *PduACRequestInfo) GetAnType() AccessType`

GetAnType returns the AnType field if non-nil, zero value otherwise.

### GetAnTypeOk

`func (o *PduACRequestInfo) GetAnTypeOk() (*AccessType, bool)`

GetAnTypeOk returns a tuple with the AnType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnType

`func (o *PduACRequestInfo) SetAnType(v AccessType)`

SetAnType sets AnType field to given value.


### GetPduSessionId

`func (o *PduACRequestInfo) GetPduSessionId() int32`

GetPduSessionId returns the PduSessionId field if non-nil, zero value otherwise.

### GetPduSessionIdOk

`func (o *PduACRequestInfo) GetPduSessionIdOk() (*int32, bool)`

GetPduSessionIdOk returns a tuple with the PduSessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPduSessionId

`func (o *PduACRequestInfo) SetPduSessionId(v int32)`

SetPduSessionId sets PduSessionId field to given value.


### GetAcuOperationList

`func (o *PduACRequestInfo) GetAcuOperationList() []AcuOperationItem`

GetAcuOperationList returns the AcuOperationList field if non-nil, zero value otherwise.

### GetAcuOperationListOk

`func (o *PduACRequestInfo) GetAcuOperationListOk() (*[]AcuOperationItem, bool)`

GetAcuOperationListOk returns a tuple with the AcuOperationList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcuOperationList

`func (o *PduACRequestInfo) SetAcuOperationList(v []AcuOperationItem)`

SetAcuOperationList sets AcuOperationList field to given value.


### GetAdditionalAnType

`func (o *PduACRequestInfo) GetAdditionalAnType() AccessType`

GetAdditionalAnType returns the AdditionalAnType field if non-nil, zero value otherwise.

### GetAdditionalAnTypeOk

`func (o *PduACRequestInfo) GetAdditionalAnTypeOk() (*AccessType, bool)`

GetAdditionalAnTypeOk returns a tuple with the AdditionalAnType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalAnType

`func (o *PduACRequestInfo) SetAdditionalAnType(v AccessType)`

SetAdditionalAnType sets AdditionalAnType field to given value.

### HasAdditionalAnType

`func (o *PduACRequestInfo) HasAdditionalAnType() bool`

HasAdditionalAnType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


