# EnableGroupReachabilityReqData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UeInfoList** | [**[]UeInfo**](UeInfo.md) |  | 
**Tmgi** | [**Tmgi**](Tmgi.md) |  | 
**ReachabilityNotifyUri** | Pointer to **string** | String providing an URI formatted according to RFC 3986. | [optional] 
**MbsServiceAreaInfoList** | Pointer to [**[]MbsServiceAreaInfo**](MbsServiceAreaInfo.md) |  | [optional] 
**Arp** | Pointer to [**Arp**](Arp.md) |  | [optional] 
**Var5qi** | Pointer to **int32** | Unsigned integer representing a 5G QoS Identifier (see clause 5.7.2.1 of 3GPP TS 23.501, within the range 0 to 255.  | [optional] 
**SupportedFeatures** | Pointer to **string** | A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \&quot;0\&quot; to \&quot;9\&quot;,  \&quot;a\&quot; to \&quot;f\&quot; or \&quot;A\&quot; to \&quot;F\&quot; and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.  | [optional] 

## Methods

### NewEnableGroupReachabilityReqData

`func NewEnableGroupReachabilityReqData(ueInfoList []UeInfo, tmgi Tmgi, ) *EnableGroupReachabilityReqData`

NewEnableGroupReachabilityReqData instantiates a new EnableGroupReachabilityReqData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnableGroupReachabilityReqDataWithDefaults

`func NewEnableGroupReachabilityReqDataWithDefaults() *EnableGroupReachabilityReqData`

NewEnableGroupReachabilityReqDataWithDefaults instantiates a new EnableGroupReachabilityReqData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUeInfoList

`func (o *EnableGroupReachabilityReqData) GetUeInfoList() []UeInfo`

GetUeInfoList returns the UeInfoList field if non-nil, zero value otherwise.

### GetUeInfoListOk

`func (o *EnableGroupReachabilityReqData) GetUeInfoListOk() (*[]UeInfo, bool)`

GetUeInfoListOk returns a tuple with the UeInfoList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUeInfoList

`func (o *EnableGroupReachabilityReqData) SetUeInfoList(v []UeInfo)`

SetUeInfoList sets UeInfoList field to given value.


### GetTmgi

`func (o *EnableGroupReachabilityReqData) GetTmgi() Tmgi`

GetTmgi returns the Tmgi field if non-nil, zero value otherwise.

### GetTmgiOk

`func (o *EnableGroupReachabilityReqData) GetTmgiOk() (*Tmgi, bool)`

GetTmgiOk returns a tuple with the Tmgi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTmgi

`func (o *EnableGroupReachabilityReqData) SetTmgi(v Tmgi)`

SetTmgi sets Tmgi field to given value.


### GetReachabilityNotifyUri

`func (o *EnableGroupReachabilityReqData) GetReachabilityNotifyUri() string`

GetReachabilityNotifyUri returns the ReachabilityNotifyUri field if non-nil, zero value otherwise.

### GetReachabilityNotifyUriOk

`func (o *EnableGroupReachabilityReqData) GetReachabilityNotifyUriOk() (*string, bool)`

GetReachabilityNotifyUriOk returns a tuple with the ReachabilityNotifyUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReachabilityNotifyUri

`func (o *EnableGroupReachabilityReqData) SetReachabilityNotifyUri(v string)`

SetReachabilityNotifyUri sets ReachabilityNotifyUri field to given value.

### HasReachabilityNotifyUri

`func (o *EnableGroupReachabilityReqData) HasReachabilityNotifyUri() bool`

HasReachabilityNotifyUri returns a boolean if a field has been set.

### GetMbsServiceAreaInfoList

`func (o *EnableGroupReachabilityReqData) GetMbsServiceAreaInfoList() []MbsServiceAreaInfo`

GetMbsServiceAreaInfoList returns the MbsServiceAreaInfoList field if non-nil, zero value otherwise.

### GetMbsServiceAreaInfoListOk

`func (o *EnableGroupReachabilityReqData) GetMbsServiceAreaInfoListOk() (*[]MbsServiceAreaInfo, bool)`

GetMbsServiceAreaInfoListOk returns a tuple with the MbsServiceAreaInfoList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMbsServiceAreaInfoList

`func (o *EnableGroupReachabilityReqData) SetMbsServiceAreaInfoList(v []MbsServiceAreaInfo)`

SetMbsServiceAreaInfoList sets MbsServiceAreaInfoList field to given value.

### HasMbsServiceAreaInfoList

`func (o *EnableGroupReachabilityReqData) HasMbsServiceAreaInfoList() bool`

HasMbsServiceAreaInfoList returns a boolean if a field has been set.

### GetArp

`func (o *EnableGroupReachabilityReqData) GetArp() Arp`

GetArp returns the Arp field if non-nil, zero value otherwise.

### GetArpOk

`func (o *EnableGroupReachabilityReqData) GetArpOk() (*Arp, bool)`

GetArpOk returns a tuple with the Arp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArp

`func (o *EnableGroupReachabilityReqData) SetArp(v Arp)`

SetArp sets Arp field to given value.

### HasArp

`func (o *EnableGroupReachabilityReqData) HasArp() bool`

HasArp returns a boolean if a field has been set.

### GetVar5qi

`func (o *EnableGroupReachabilityReqData) GetVar5qi() int32`

GetVar5qi returns the Var5qi field if non-nil, zero value otherwise.

### GetVar5qiOk

`func (o *EnableGroupReachabilityReqData) GetVar5qiOk() (*int32, bool)`

GetVar5qiOk returns a tuple with the Var5qi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar5qi

`func (o *EnableGroupReachabilityReqData) SetVar5qi(v int32)`

SetVar5qi sets Var5qi field to given value.

### HasVar5qi

`func (o *EnableGroupReachabilityReqData) HasVar5qi() bool`

HasVar5qi returns a boolean if a field has been set.

### GetSupportedFeatures

`func (o *EnableGroupReachabilityReqData) GetSupportedFeatures() string`

GetSupportedFeatures returns the SupportedFeatures field if non-nil, zero value otherwise.

### GetSupportedFeaturesOk

`func (o *EnableGroupReachabilityReqData) GetSupportedFeaturesOk() (*string, bool)`

GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedFeatures

`func (o *EnableGroupReachabilityReqData) SetSupportedFeatures(v string)`

SetSupportedFeatures sets SupportedFeatures field to given value.

### HasSupportedFeatures

`func (o *EnableGroupReachabilityReqData) HasSupportedFeatures() bool`

HasSupportedFeatures returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


