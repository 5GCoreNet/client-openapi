# MbsAppSessionCtxt

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MbsSessionId** | [**MbsSessionId**](MbsSessionId.md) |  | 
**MbsServInfo** | Pointer to [**MbsServiceInfo**](MbsServiceInfo.md) |  | [optional] 
**Dnn** | Pointer to **string** | String representing a Data Network as defined in clause 9A of 3GPP TS 23.003;  it shall contain either a DNN Network Identifier, or a full DNN with both the Network  Identifier and Operator Identifier, as specified in 3GPP TS 23.003 clause 9.1.1 and 9.1.2. It shall be coded as string in which the labels are separated by dots  (e.g. \&quot;Label1.Label2.Label3\&quot;).  | [optional] 
**Snssai** | Pointer to [**Snssai**](Snssai.md) |  | [optional] 
**ContactPcfInd** | Pointer to **bool** |  | [optional] [default to false]
**SuppFeat** | Pointer to **string** | A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \&quot;0\&quot; to \&quot;9\&quot;,  \&quot;a\&quot; to \&quot;f\&quot; or \&quot;A\&quot; to \&quot;F\&quot; and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.  | [optional] 

## Methods

### NewMbsAppSessionCtxt

`func NewMbsAppSessionCtxt(mbsSessionId MbsSessionId, ) *MbsAppSessionCtxt`

NewMbsAppSessionCtxt instantiates a new MbsAppSessionCtxt object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMbsAppSessionCtxtWithDefaults

`func NewMbsAppSessionCtxtWithDefaults() *MbsAppSessionCtxt`

NewMbsAppSessionCtxtWithDefaults instantiates a new MbsAppSessionCtxt object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMbsSessionId

`func (o *MbsAppSessionCtxt) GetMbsSessionId() MbsSessionId`

GetMbsSessionId returns the MbsSessionId field if non-nil, zero value otherwise.

### GetMbsSessionIdOk

`func (o *MbsAppSessionCtxt) GetMbsSessionIdOk() (*MbsSessionId, bool)`

GetMbsSessionIdOk returns a tuple with the MbsSessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMbsSessionId

`func (o *MbsAppSessionCtxt) SetMbsSessionId(v MbsSessionId)`

SetMbsSessionId sets MbsSessionId field to given value.


### GetMbsServInfo

`func (o *MbsAppSessionCtxt) GetMbsServInfo() MbsServiceInfo`

GetMbsServInfo returns the MbsServInfo field if non-nil, zero value otherwise.

### GetMbsServInfoOk

`func (o *MbsAppSessionCtxt) GetMbsServInfoOk() (*MbsServiceInfo, bool)`

GetMbsServInfoOk returns a tuple with the MbsServInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMbsServInfo

`func (o *MbsAppSessionCtxt) SetMbsServInfo(v MbsServiceInfo)`

SetMbsServInfo sets MbsServInfo field to given value.

### HasMbsServInfo

`func (o *MbsAppSessionCtxt) HasMbsServInfo() bool`

HasMbsServInfo returns a boolean if a field has been set.

### GetDnn

`func (o *MbsAppSessionCtxt) GetDnn() string`

GetDnn returns the Dnn field if non-nil, zero value otherwise.

### GetDnnOk

`func (o *MbsAppSessionCtxt) GetDnnOk() (*string, bool)`

GetDnnOk returns a tuple with the Dnn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnn

`func (o *MbsAppSessionCtxt) SetDnn(v string)`

SetDnn sets Dnn field to given value.

### HasDnn

`func (o *MbsAppSessionCtxt) HasDnn() bool`

HasDnn returns a boolean if a field has been set.

### GetSnssai

`func (o *MbsAppSessionCtxt) GetSnssai() Snssai`

GetSnssai returns the Snssai field if non-nil, zero value otherwise.

### GetSnssaiOk

`func (o *MbsAppSessionCtxt) GetSnssaiOk() (*Snssai, bool)`

GetSnssaiOk returns a tuple with the Snssai field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnssai

`func (o *MbsAppSessionCtxt) SetSnssai(v Snssai)`

SetSnssai sets Snssai field to given value.

### HasSnssai

`func (o *MbsAppSessionCtxt) HasSnssai() bool`

HasSnssai returns a boolean if a field has been set.

### GetContactPcfInd

`func (o *MbsAppSessionCtxt) GetContactPcfInd() bool`

GetContactPcfInd returns the ContactPcfInd field if non-nil, zero value otherwise.

### GetContactPcfIndOk

`func (o *MbsAppSessionCtxt) GetContactPcfIndOk() (*bool, bool)`

GetContactPcfIndOk returns a tuple with the ContactPcfInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactPcfInd

`func (o *MbsAppSessionCtxt) SetContactPcfInd(v bool)`

SetContactPcfInd sets ContactPcfInd field to given value.

### HasContactPcfInd

`func (o *MbsAppSessionCtxt) HasContactPcfInd() bool`

HasContactPcfInd returns a boolean if a field has been set.

### GetSuppFeat

`func (o *MbsAppSessionCtxt) GetSuppFeat() string`

GetSuppFeat returns the SuppFeat field if non-nil, zero value otherwise.

### GetSuppFeatOk

`func (o *MbsAppSessionCtxt) GetSuppFeatOk() (*string, bool)`

GetSuppFeatOk returns a tuple with the SuppFeat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuppFeat

`func (o *MbsAppSessionCtxt) SetSuppFeat(v string)`

SetSuppFeat sets SuppFeat field to given value.

### HasSuppFeat

`func (o *MbsAppSessionCtxt) HasSuppFeat() bool`

HasSuppFeat returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


