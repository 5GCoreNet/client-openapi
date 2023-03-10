# BdtData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AspId** | **string** |  | 
**TransPolicy** | [**TransferPolicy**](TransferPolicy.md) |  | 
**BdtRefId** | Pointer to **string** | string identifying a BDT Reference ID as defined in clause 5.3.3 of 3GPP TS 29.154. | [optional] 
**NwAreaInfo** | Pointer to [**NetworkAreaInfo**](NetworkAreaInfo.md) |  | [optional] 
**NumOfUes** | Pointer to **int32** | Unsigned Integer, i.e. only value 0 and integers above 0 are permissible. | [optional] 
**VolPerUe** | Pointer to [**UsageThreshold**](UsageThreshold.md) |  | [optional] 
**Dnn** | Pointer to **string** | String representing a Data Network as defined in clause 9A of 3GPP TS 23.003;  it shall contain either a DNN Network Identifier, or a full DNN with both the Network  Identifier and Operator Identifier, as specified in 3GPP TS 23.003 clause 9.1.1 and 9.1.2. It shall be coded as string in which the labels are separated by dots  (e.g. \&quot;Label1.Label2.Label3\&quot;).  | [optional] 
**Snssai** | Pointer to [**Snssai**](Snssai.md) |  | [optional] 
**TrafficDes** | Pointer to **string** | Identify a traffic descriptor as defined in Figure 5.2.2 of 3GPP TS 24.526, octets v+5 to w. | [optional] 
**BdtpStatus** | Pointer to [**BdtPolicyStatus**](BdtPolicyStatus.md) |  | [optional] 
**SuppFeat** | Pointer to **string** | A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \&quot;0\&quot; to \&quot;9\&quot;,  \&quot;a\&quot; to \&quot;f\&quot; or \&quot;A\&quot; to \&quot;F\&quot; and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.  | [optional] 
**ResetIds** | Pointer to **[]string** |  | [optional] 

## Methods

### NewBdtData

`func NewBdtData(aspId string, transPolicy TransferPolicy, ) *BdtData`

NewBdtData instantiates a new BdtData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBdtDataWithDefaults

`func NewBdtDataWithDefaults() *BdtData`

NewBdtDataWithDefaults instantiates a new BdtData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAspId

`func (o *BdtData) GetAspId() string`

GetAspId returns the AspId field if non-nil, zero value otherwise.

### GetAspIdOk

`func (o *BdtData) GetAspIdOk() (*string, bool)`

GetAspIdOk returns a tuple with the AspId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAspId

`func (o *BdtData) SetAspId(v string)`

SetAspId sets AspId field to given value.


### GetTransPolicy

`func (o *BdtData) GetTransPolicy() TransferPolicy`

GetTransPolicy returns the TransPolicy field if non-nil, zero value otherwise.

### GetTransPolicyOk

`func (o *BdtData) GetTransPolicyOk() (*TransferPolicy, bool)`

GetTransPolicyOk returns a tuple with the TransPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransPolicy

`func (o *BdtData) SetTransPolicy(v TransferPolicy)`

SetTransPolicy sets TransPolicy field to given value.


### GetBdtRefId

`func (o *BdtData) GetBdtRefId() string`

GetBdtRefId returns the BdtRefId field if non-nil, zero value otherwise.

### GetBdtRefIdOk

`func (o *BdtData) GetBdtRefIdOk() (*string, bool)`

GetBdtRefIdOk returns a tuple with the BdtRefId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBdtRefId

`func (o *BdtData) SetBdtRefId(v string)`

SetBdtRefId sets BdtRefId field to given value.

### HasBdtRefId

`func (o *BdtData) HasBdtRefId() bool`

HasBdtRefId returns a boolean if a field has been set.

### GetNwAreaInfo

`func (o *BdtData) GetNwAreaInfo() NetworkAreaInfo`

GetNwAreaInfo returns the NwAreaInfo field if non-nil, zero value otherwise.

### GetNwAreaInfoOk

`func (o *BdtData) GetNwAreaInfoOk() (*NetworkAreaInfo, bool)`

GetNwAreaInfoOk returns a tuple with the NwAreaInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNwAreaInfo

`func (o *BdtData) SetNwAreaInfo(v NetworkAreaInfo)`

SetNwAreaInfo sets NwAreaInfo field to given value.

### HasNwAreaInfo

`func (o *BdtData) HasNwAreaInfo() bool`

HasNwAreaInfo returns a boolean if a field has been set.

### GetNumOfUes

`func (o *BdtData) GetNumOfUes() int32`

GetNumOfUes returns the NumOfUes field if non-nil, zero value otherwise.

### GetNumOfUesOk

`func (o *BdtData) GetNumOfUesOk() (*int32, bool)`

GetNumOfUesOk returns a tuple with the NumOfUes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumOfUes

`func (o *BdtData) SetNumOfUes(v int32)`

SetNumOfUes sets NumOfUes field to given value.

### HasNumOfUes

`func (o *BdtData) HasNumOfUes() bool`

HasNumOfUes returns a boolean if a field has been set.

### GetVolPerUe

`func (o *BdtData) GetVolPerUe() UsageThreshold`

GetVolPerUe returns the VolPerUe field if non-nil, zero value otherwise.

### GetVolPerUeOk

`func (o *BdtData) GetVolPerUeOk() (*UsageThreshold, bool)`

GetVolPerUeOk returns a tuple with the VolPerUe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolPerUe

`func (o *BdtData) SetVolPerUe(v UsageThreshold)`

SetVolPerUe sets VolPerUe field to given value.

### HasVolPerUe

`func (o *BdtData) HasVolPerUe() bool`

HasVolPerUe returns a boolean if a field has been set.

### GetDnn

`func (o *BdtData) GetDnn() string`

GetDnn returns the Dnn field if non-nil, zero value otherwise.

### GetDnnOk

`func (o *BdtData) GetDnnOk() (*string, bool)`

GetDnnOk returns a tuple with the Dnn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnn

`func (o *BdtData) SetDnn(v string)`

SetDnn sets Dnn field to given value.

### HasDnn

`func (o *BdtData) HasDnn() bool`

HasDnn returns a boolean if a field has been set.

### GetSnssai

`func (o *BdtData) GetSnssai() Snssai`

GetSnssai returns the Snssai field if non-nil, zero value otherwise.

### GetSnssaiOk

`func (o *BdtData) GetSnssaiOk() (*Snssai, bool)`

GetSnssaiOk returns a tuple with the Snssai field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnssai

`func (o *BdtData) SetSnssai(v Snssai)`

SetSnssai sets Snssai field to given value.

### HasSnssai

`func (o *BdtData) HasSnssai() bool`

HasSnssai returns a boolean if a field has been set.

### GetTrafficDes

`func (o *BdtData) GetTrafficDes() string`

GetTrafficDes returns the TrafficDes field if non-nil, zero value otherwise.

### GetTrafficDesOk

`func (o *BdtData) GetTrafficDesOk() (*string, bool)`

GetTrafficDesOk returns a tuple with the TrafficDes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrafficDes

`func (o *BdtData) SetTrafficDes(v string)`

SetTrafficDes sets TrafficDes field to given value.

### HasTrafficDes

`func (o *BdtData) HasTrafficDes() bool`

HasTrafficDes returns a boolean if a field has been set.

### GetBdtpStatus

`func (o *BdtData) GetBdtpStatus() BdtPolicyStatus`

GetBdtpStatus returns the BdtpStatus field if non-nil, zero value otherwise.

### GetBdtpStatusOk

`func (o *BdtData) GetBdtpStatusOk() (*BdtPolicyStatus, bool)`

GetBdtpStatusOk returns a tuple with the BdtpStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBdtpStatus

`func (o *BdtData) SetBdtpStatus(v BdtPolicyStatus)`

SetBdtpStatus sets BdtpStatus field to given value.

### HasBdtpStatus

`func (o *BdtData) HasBdtpStatus() bool`

HasBdtpStatus returns a boolean if a field has been set.

### GetSuppFeat

`func (o *BdtData) GetSuppFeat() string`

GetSuppFeat returns the SuppFeat field if non-nil, zero value otherwise.

### GetSuppFeatOk

`func (o *BdtData) GetSuppFeatOk() (*string, bool)`

GetSuppFeatOk returns a tuple with the SuppFeat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuppFeat

`func (o *BdtData) SetSuppFeat(v string)`

SetSuppFeat sets SuppFeat field to given value.

### HasSuppFeat

`func (o *BdtData) HasSuppFeat() bool`

HasSuppFeat returns a boolean if a field has been set.

### GetResetIds

`func (o *BdtData) GetResetIds() []string`

GetResetIds returns the ResetIds field if non-nil, zero value otherwise.

### GetResetIdsOk

`func (o *BdtData) GetResetIdsOk() (*[]string, bool)`

GetResetIdsOk returns a tuple with the ResetIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResetIds

`func (o *BdtData) SetResetIds(v []string)`

SetResetIds sets ResetIds field to given value.

### HasResetIds

`func (o *BdtData) HasResetIds() bool`

HasResetIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


