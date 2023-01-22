# EasDeployInfoData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppId** | Pointer to **string** |  | [optional] 
**DnaiInfos** | Pointer to [**map[string]DnaiInformation**](DnaiInformation.md) | list of DNS server identifier (consisting of IP address and port) and/or IP address(s)  of the EAS in the local DN for each DNAI. The key of map is the DNAI.  | [optional] 
**Dnn** | Pointer to **string** | String representing a Data Network as defined in clause 9A of 3GPP TS 23.003;  it shall contain either a DNN Network Identifier, or a full DNN with both the Network  Identifier and Operator Identifier, as specified in 3GPP TS 23.003 clause 9.1.1 and 9.1.2. It shall be coded as string in which the labels are separated by dots  (e.g. \&quot;Label1.Label2.Label3\&quot;).  | [optional] 
**FqdnPatternList** | [**[]FqdnPatternMatchingRule**](FqdnPatternMatchingRule.md) |  | 
**InternalGroupId** | Pointer to **string** | String identifying a group of devices network internal globally unique ID which identifies a set of IMSIs, as specified in clause 19.9 of 3GPP TS 23.003.   | [optional] 
**Snssai** | Pointer to [**Snssai**](Snssai.md) |  | [optional] 

## Methods

### NewEasDeployInfoData

`func NewEasDeployInfoData(fqdnPatternList []FqdnPatternMatchingRule, ) *EasDeployInfoData`

NewEasDeployInfoData instantiates a new EasDeployInfoData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEasDeployInfoDataWithDefaults

`func NewEasDeployInfoDataWithDefaults() *EasDeployInfoData`

NewEasDeployInfoDataWithDefaults instantiates a new EasDeployInfoData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppId

`func (o *EasDeployInfoData) GetAppId() string`

GetAppId returns the AppId field if non-nil, zero value otherwise.

### GetAppIdOk

`func (o *EasDeployInfoData) GetAppIdOk() (*string, bool)`

GetAppIdOk returns a tuple with the AppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppId

`func (o *EasDeployInfoData) SetAppId(v string)`

SetAppId sets AppId field to given value.

### HasAppId

`func (o *EasDeployInfoData) HasAppId() bool`

HasAppId returns a boolean if a field has been set.

### GetDnaiInfos

`func (o *EasDeployInfoData) GetDnaiInfos() map[string]DnaiInformation`

GetDnaiInfos returns the DnaiInfos field if non-nil, zero value otherwise.

### GetDnaiInfosOk

`func (o *EasDeployInfoData) GetDnaiInfosOk() (*map[string]DnaiInformation, bool)`

GetDnaiInfosOk returns a tuple with the DnaiInfos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnaiInfos

`func (o *EasDeployInfoData) SetDnaiInfos(v map[string]DnaiInformation)`

SetDnaiInfos sets DnaiInfos field to given value.

### HasDnaiInfos

`func (o *EasDeployInfoData) HasDnaiInfos() bool`

HasDnaiInfos returns a boolean if a field has been set.

### GetDnn

`func (o *EasDeployInfoData) GetDnn() string`

GetDnn returns the Dnn field if non-nil, zero value otherwise.

### GetDnnOk

`func (o *EasDeployInfoData) GetDnnOk() (*string, bool)`

GetDnnOk returns a tuple with the Dnn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnn

`func (o *EasDeployInfoData) SetDnn(v string)`

SetDnn sets Dnn field to given value.

### HasDnn

`func (o *EasDeployInfoData) HasDnn() bool`

HasDnn returns a boolean if a field has been set.

### GetFqdnPatternList

`func (o *EasDeployInfoData) GetFqdnPatternList() []FqdnPatternMatchingRule`

GetFqdnPatternList returns the FqdnPatternList field if non-nil, zero value otherwise.

### GetFqdnPatternListOk

`func (o *EasDeployInfoData) GetFqdnPatternListOk() (*[]FqdnPatternMatchingRule, bool)`

GetFqdnPatternListOk returns a tuple with the FqdnPatternList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFqdnPatternList

`func (o *EasDeployInfoData) SetFqdnPatternList(v []FqdnPatternMatchingRule)`

SetFqdnPatternList sets FqdnPatternList field to given value.


### GetInternalGroupId

`func (o *EasDeployInfoData) GetInternalGroupId() string`

GetInternalGroupId returns the InternalGroupId field if non-nil, zero value otherwise.

### GetInternalGroupIdOk

`func (o *EasDeployInfoData) GetInternalGroupIdOk() (*string, bool)`

GetInternalGroupIdOk returns a tuple with the InternalGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalGroupId

`func (o *EasDeployInfoData) SetInternalGroupId(v string)`

SetInternalGroupId sets InternalGroupId field to given value.

### HasInternalGroupId

`func (o *EasDeployInfoData) HasInternalGroupId() bool`

HasInternalGroupId returns a boolean if a field has been set.

### GetSnssai

`func (o *EasDeployInfoData) GetSnssai() Snssai`

GetSnssai returns the Snssai field if non-nil, zero value otherwise.

### GetSnssaiOk

`func (o *EasDeployInfoData) GetSnssaiOk() (*Snssai, bool)`

GetSnssaiOk returns a tuple with the Snssai field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnssai

`func (o *EasDeployInfoData) SetSnssai(v Snssai)`

SetSnssai sets Snssai field to given value.

### HasSnssai

`func (o *EasDeployInfoData) HasSnssai() bool`

HasSnssai returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


