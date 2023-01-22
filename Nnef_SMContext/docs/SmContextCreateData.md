# SmContextCreateData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Supi** | **string** | String identifying a Supi that shall contain either an IMSI, a network specific identifier, a Global Cable Identifier (GCI) or a Global Line Identifier (GLI) as specified in clause  2.2A of 3GPP TS 23.003. It shall be formatted as follows  - for an IMSI \&quot;imsi-&lt;imsi&gt;\&quot;, where &lt;imsi&gt; shall be formatted according to clause 2.2    of 3GPP TS 23.003 that describes an IMSI.  - for a network specific identifier \&quot;nai-&lt;nai&gt;, where &lt;nai&gt; shall be formatted    according to clause 28.7.2 of 3GPP TS 23.003 that describes an NAI.  - for a GCI \&quot;gci-&lt;gci&gt;\&quot;, where &lt;gci&gt; shall be formatted according to clause 28.15.2    of 3GPP TS 23.003.  - for a GLI \&quot;gli-&lt;gli&gt;\&quot;, where &lt;gli&gt; shall be formatted according to clause 28.16.2 of    3GPP TS 23.003.To enable that the value is used as part of an URI, the string shall    only contain characters allowed according to the \&quot;lower-with-hyphen\&quot; naming convention    defined in 3GPP TS 29.501.  | 
**PduSessionId** | **int32** | Unsigned integer identifying a PDU session, within the range 0 to 255, as specified in  clause 11.2.3.1b, bits 1 to 8, of 3GPP TS 24.007. If the PDU Session ID is allocated by the  Core Network for UEs not supporting N1 mode, reserved range 64 to 95 is used. PDU Session ID  within the reserved range is only visible in the Core Network.   | 
**Dnn** | **string** | String representing a Data Network as defined in clause 9A of 3GPP TS 23.003;  it shall contain either a DNN Network Identifier, or a full DNN with both the Network  Identifier and Operator Identifier, as specified in 3GPP TS 23.003 clause 9.1.1 and 9.1.2. It shall be coded as string in which the labels are separated by dots  (e.g. \&quot;Label1.Label2.Label3\&quot;).  | 
**Snssai** | [**Snssai**](Snssai.md) |  | 
**NefId** | **string** | This IE shall contain the NEF ID of the target NEF. | 
**DlNiddEndPoint** | **string** | String providing an URI formatted according to RFC 3986. | 
**NotificationUri** | **string** | String providing an URI formatted according to RFC 3986. | 
**NiddInfo** | Pointer to [**NiddInformation**](NiddInformation.md) |  | [optional] 
**RdsSupport** | Pointer to **bool** | When present, this IE shall indicate the UE capability to support RDS. The value of this IE shall be set as following  - true  UE supports RDS  - false (default)  UE does not support RDS  | [optional] 
**SmContextConfig** | Pointer to [**SmContextConfiguration**](SmContextConfiguration.md) |  | [optional] 
**SupportedFeatures** | Pointer to **string** | A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \&quot;0\&quot; to \&quot;9\&quot;,  \&quot;a\&quot; to \&quot;f\&quot; or \&quot;A\&quot; to \&quot;F\&quot; and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.  | [optional] 

## Methods

### NewSmContextCreateData

`func NewSmContextCreateData(supi string, pduSessionId int32, dnn string, snssai Snssai, nefId string, dlNiddEndPoint string, notificationUri string, ) *SmContextCreateData`

NewSmContextCreateData instantiates a new SmContextCreateData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSmContextCreateDataWithDefaults

`func NewSmContextCreateDataWithDefaults() *SmContextCreateData`

NewSmContextCreateDataWithDefaults instantiates a new SmContextCreateData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSupi

`func (o *SmContextCreateData) GetSupi() string`

GetSupi returns the Supi field if non-nil, zero value otherwise.

### GetSupiOk

`func (o *SmContextCreateData) GetSupiOk() (*string, bool)`

GetSupiOk returns a tuple with the Supi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupi

`func (o *SmContextCreateData) SetSupi(v string)`

SetSupi sets Supi field to given value.


### GetPduSessionId

`func (o *SmContextCreateData) GetPduSessionId() int32`

GetPduSessionId returns the PduSessionId field if non-nil, zero value otherwise.

### GetPduSessionIdOk

`func (o *SmContextCreateData) GetPduSessionIdOk() (*int32, bool)`

GetPduSessionIdOk returns a tuple with the PduSessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPduSessionId

`func (o *SmContextCreateData) SetPduSessionId(v int32)`

SetPduSessionId sets PduSessionId field to given value.


### GetDnn

`func (o *SmContextCreateData) GetDnn() string`

GetDnn returns the Dnn field if non-nil, zero value otherwise.

### GetDnnOk

`func (o *SmContextCreateData) GetDnnOk() (*string, bool)`

GetDnnOk returns a tuple with the Dnn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnn

`func (o *SmContextCreateData) SetDnn(v string)`

SetDnn sets Dnn field to given value.


### GetSnssai

`func (o *SmContextCreateData) GetSnssai() Snssai`

GetSnssai returns the Snssai field if non-nil, zero value otherwise.

### GetSnssaiOk

`func (o *SmContextCreateData) GetSnssaiOk() (*Snssai, bool)`

GetSnssaiOk returns a tuple with the Snssai field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnssai

`func (o *SmContextCreateData) SetSnssai(v Snssai)`

SetSnssai sets Snssai field to given value.


### GetNefId

`func (o *SmContextCreateData) GetNefId() string`

GetNefId returns the NefId field if non-nil, zero value otherwise.

### GetNefIdOk

`func (o *SmContextCreateData) GetNefIdOk() (*string, bool)`

GetNefIdOk returns a tuple with the NefId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNefId

`func (o *SmContextCreateData) SetNefId(v string)`

SetNefId sets NefId field to given value.


### GetDlNiddEndPoint

`func (o *SmContextCreateData) GetDlNiddEndPoint() string`

GetDlNiddEndPoint returns the DlNiddEndPoint field if non-nil, zero value otherwise.

### GetDlNiddEndPointOk

`func (o *SmContextCreateData) GetDlNiddEndPointOk() (*string, bool)`

GetDlNiddEndPointOk returns a tuple with the DlNiddEndPoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDlNiddEndPoint

`func (o *SmContextCreateData) SetDlNiddEndPoint(v string)`

SetDlNiddEndPoint sets DlNiddEndPoint field to given value.


### GetNotificationUri

`func (o *SmContextCreateData) GetNotificationUri() string`

GetNotificationUri returns the NotificationUri field if non-nil, zero value otherwise.

### GetNotificationUriOk

`func (o *SmContextCreateData) GetNotificationUriOk() (*string, bool)`

GetNotificationUriOk returns a tuple with the NotificationUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationUri

`func (o *SmContextCreateData) SetNotificationUri(v string)`

SetNotificationUri sets NotificationUri field to given value.


### GetNiddInfo

`func (o *SmContextCreateData) GetNiddInfo() NiddInformation`

GetNiddInfo returns the NiddInfo field if non-nil, zero value otherwise.

### GetNiddInfoOk

`func (o *SmContextCreateData) GetNiddInfoOk() (*NiddInformation, bool)`

GetNiddInfoOk returns a tuple with the NiddInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNiddInfo

`func (o *SmContextCreateData) SetNiddInfo(v NiddInformation)`

SetNiddInfo sets NiddInfo field to given value.

### HasNiddInfo

`func (o *SmContextCreateData) HasNiddInfo() bool`

HasNiddInfo returns a boolean if a field has been set.

### GetRdsSupport

`func (o *SmContextCreateData) GetRdsSupport() bool`

GetRdsSupport returns the RdsSupport field if non-nil, zero value otherwise.

### GetRdsSupportOk

`func (o *SmContextCreateData) GetRdsSupportOk() (*bool, bool)`

GetRdsSupportOk returns a tuple with the RdsSupport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRdsSupport

`func (o *SmContextCreateData) SetRdsSupport(v bool)`

SetRdsSupport sets RdsSupport field to given value.

### HasRdsSupport

`func (o *SmContextCreateData) HasRdsSupport() bool`

HasRdsSupport returns a boolean if a field has been set.

### GetSmContextConfig

`func (o *SmContextCreateData) GetSmContextConfig() SmContextConfiguration`

GetSmContextConfig returns the SmContextConfig field if non-nil, zero value otherwise.

### GetSmContextConfigOk

`func (o *SmContextCreateData) GetSmContextConfigOk() (*SmContextConfiguration, bool)`

GetSmContextConfigOk returns a tuple with the SmContextConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmContextConfig

`func (o *SmContextCreateData) SetSmContextConfig(v SmContextConfiguration)`

SetSmContextConfig sets SmContextConfig field to given value.

### HasSmContextConfig

`func (o *SmContextCreateData) HasSmContextConfig() bool`

HasSmContextConfig returns a boolean if a field has been set.

### GetSupportedFeatures

`func (o *SmContextCreateData) GetSupportedFeatures() string`

GetSupportedFeatures returns the SupportedFeatures field if non-nil, zero value otherwise.

### GetSupportedFeaturesOk

`func (o *SmContextCreateData) GetSupportedFeaturesOk() (*string, bool)`

GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedFeatures

`func (o *SmContextCreateData) SetSupportedFeatures(v string)`

SetSupportedFeatures sets SupportedFeatures field to given value.

### HasSupportedFeatures

`func (o *SmContextCreateData) HasSupportedFeatures() bool`

HasSupportedFeatures returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


