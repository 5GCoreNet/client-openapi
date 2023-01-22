# SmContextCreatedData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Supi** | **string** | String identifying a Supi that shall contain either an IMSI, a network specific identifier, a Global Cable Identifier (GCI) or a Global Line Identifier (GLI) as specified in clause  2.2A of 3GPP TS 23.003. It shall be formatted as follows  - for an IMSI \&quot;imsi-&lt;imsi&gt;\&quot;, where &lt;imsi&gt; shall be formatted according to clause 2.2    of 3GPP TS 23.003 that describes an IMSI.  - for a network specific identifier \&quot;nai-&lt;nai&gt;, where &lt;nai&gt; shall be formatted    according to clause 28.7.2 of 3GPP TS 23.003 that describes an NAI.  - for a GCI \&quot;gci-&lt;gci&gt;\&quot;, where &lt;gci&gt; shall be formatted according to clause 28.15.2    of 3GPP TS 23.003.  - for a GLI \&quot;gli-&lt;gli&gt;\&quot;, where &lt;gli&gt; shall be formatted according to clause 28.16.2 of    3GPP TS 23.003.To enable that the value is used as part of an URI, the string shall    only contain characters allowed according to the \&quot;lower-with-hyphen\&quot; naming convention    defined in 3GPP TS 29.501.  | 
**PduSessionId** | **int32** | Unsigned integer identifying a PDU session, within the range 0 to 255, as specified in  clause 11.2.3.1b, bits 1 to 8, of 3GPP TS 24.007. If the PDU Session ID is allocated by the  Core Network for UEs not supporting N1 mode, reserved range 64 to 95 is used. PDU Session ID  within the reserved range is only visible in the Core Network.   | 
**Dnn** | **string** | String representing a Data Network as defined in clause 9A of 3GPP TS 23.003;  it shall contain either a DNN Network Identifier, or a full DNN with both the Network  Identifier and Operator Identifier, as specified in 3GPP TS 23.003 clause 9.1.1 and 9.1.2. It shall be coded as string in which the labels are separated by dots  (e.g. \&quot;Label1.Label2.Label3\&quot;).  | 
**Snssai** | [**Snssai**](Snssai.md) |  | 
**NefId** | **string** | This IE shall contain the NEF ID of the target NEF. | 
**RdsSupport** | Pointer to **bool** | When present, this IE shall indicate the NEF capability to support RDS. The value of this IE shall be set as following   - true  NEF supports RDS   - false (default)  NEF does not support RDS  | [optional] [default to false]
**ExtBufSupport** | Pointer to **bool** | When present, this IE shall indicate whether Extended Buffering applies or not. The value of this IE shall be set as following   - true  Extended Buffering applies   - false (default)  Extended Buffering does not apply  | [optional] [default to false]
**SupportedFeatures** | Pointer to **string** | A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \&quot;0\&quot; to \&quot;9\&quot;,  \&quot;a\&quot; to \&quot;f\&quot; or \&quot;A\&quot; to \&quot;F\&quot; and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.  | [optional] 
**MaxPacketSize** | Pointer to **int32** |  | [optional] 

## Methods

### NewSmContextCreatedData

`func NewSmContextCreatedData(supi string, pduSessionId int32, dnn string, snssai Snssai, nefId string, ) *SmContextCreatedData`

NewSmContextCreatedData instantiates a new SmContextCreatedData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSmContextCreatedDataWithDefaults

`func NewSmContextCreatedDataWithDefaults() *SmContextCreatedData`

NewSmContextCreatedDataWithDefaults instantiates a new SmContextCreatedData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSupi

`func (o *SmContextCreatedData) GetSupi() string`

GetSupi returns the Supi field if non-nil, zero value otherwise.

### GetSupiOk

`func (o *SmContextCreatedData) GetSupiOk() (*string, bool)`

GetSupiOk returns a tuple with the Supi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupi

`func (o *SmContextCreatedData) SetSupi(v string)`

SetSupi sets Supi field to given value.


### GetPduSessionId

`func (o *SmContextCreatedData) GetPduSessionId() int32`

GetPduSessionId returns the PduSessionId field if non-nil, zero value otherwise.

### GetPduSessionIdOk

`func (o *SmContextCreatedData) GetPduSessionIdOk() (*int32, bool)`

GetPduSessionIdOk returns a tuple with the PduSessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPduSessionId

`func (o *SmContextCreatedData) SetPduSessionId(v int32)`

SetPduSessionId sets PduSessionId field to given value.


### GetDnn

`func (o *SmContextCreatedData) GetDnn() string`

GetDnn returns the Dnn field if non-nil, zero value otherwise.

### GetDnnOk

`func (o *SmContextCreatedData) GetDnnOk() (*string, bool)`

GetDnnOk returns a tuple with the Dnn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnn

`func (o *SmContextCreatedData) SetDnn(v string)`

SetDnn sets Dnn field to given value.


### GetSnssai

`func (o *SmContextCreatedData) GetSnssai() Snssai`

GetSnssai returns the Snssai field if non-nil, zero value otherwise.

### GetSnssaiOk

`func (o *SmContextCreatedData) GetSnssaiOk() (*Snssai, bool)`

GetSnssaiOk returns a tuple with the Snssai field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnssai

`func (o *SmContextCreatedData) SetSnssai(v Snssai)`

SetSnssai sets Snssai field to given value.


### GetNefId

`func (o *SmContextCreatedData) GetNefId() string`

GetNefId returns the NefId field if non-nil, zero value otherwise.

### GetNefIdOk

`func (o *SmContextCreatedData) GetNefIdOk() (*string, bool)`

GetNefIdOk returns a tuple with the NefId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNefId

`func (o *SmContextCreatedData) SetNefId(v string)`

SetNefId sets NefId field to given value.


### GetRdsSupport

`func (o *SmContextCreatedData) GetRdsSupport() bool`

GetRdsSupport returns the RdsSupport field if non-nil, zero value otherwise.

### GetRdsSupportOk

`func (o *SmContextCreatedData) GetRdsSupportOk() (*bool, bool)`

GetRdsSupportOk returns a tuple with the RdsSupport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRdsSupport

`func (o *SmContextCreatedData) SetRdsSupport(v bool)`

SetRdsSupport sets RdsSupport field to given value.

### HasRdsSupport

`func (o *SmContextCreatedData) HasRdsSupport() bool`

HasRdsSupport returns a boolean if a field has been set.

### GetExtBufSupport

`func (o *SmContextCreatedData) GetExtBufSupport() bool`

GetExtBufSupport returns the ExtBufSupport field if non-nil, zero value otherwise.

### GetExtBufSupportOk

`func (o *SmContextCreatedData) GetExtBufSupportOk() (*bool, bool)`

GetExtBufSupportOk returns a tuple with the ExtBufSupport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtBufSupport

`func (o *SmContextCreatedData) SetExtBufSupport(v bool)`

SetExtBufSupport sets ExtBufSupport field to given value.

### HasExtBufSupport

`func (o *SmContextCreatedData) HasExtBufSupport() bool`

HasExtBufSupport returns a boolean if a field has been set.

### GetSupportedFeatures

`func (o *SmContextCreatedData) GetSupportedFeatures() string`

GetSupportedFeatures returns the SupportedFeatures field if non-nil, zero value otherwise.

### GetSupportedFeaturesOk

`func (o *SmContextCreatedData) GetSupportedFeaturesOk() (*string, bool)`

GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedFeatures

`func (o *SmContextCreatedData) SetSupportedFeatures(v string)`

SetSupportedFeatures sets SupportedFeatures field to given value.

### HasSupportedFeatures

`func (o *SmContextCreatedData) HasSupportedFeatures() bool`

HasSupportedFeatures returns a boolean if a field has been set.

### GetMaxPacketSize

`func (o *SmContextCreatedData) GetMaxPacketSize() int32`

GetMaxPacketSize returns the MaxPacketSize field if non-nil, zero value otherwise.

### GetMaxPacketSizeOk

`func (o *SmContextCreatedData) GetMaxPacketSizeOk() (*int32, bool)`

GetMaxPacketSizeOk returns a tuple with the MaxPacketSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPacketSize

`func (o *SmContextCreatedData) SetMaxPacketSize(v int32)`

SetMaxPacketSize sets MaxPacketSize field to given value.

### HasMaxPacketSize

`func (o *SmContextCreatedData) HasMaxPacketSize() bool`

HasMaxPacketSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


