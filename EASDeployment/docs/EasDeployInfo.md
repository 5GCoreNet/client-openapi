# EasDeployInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Self** | Pointer to **string** | string formatted according to IETF RFC 3986 identifying a referenced resource. | [optional] 
**AfServiceId** | Pointer to **string** |  | [optional] 
**FqdnPatternList** | [**[]FqdnPatternMatchingRule**](FqdnPatternMatchingRule.md) |  | 
**AppId** | Pointer to **string** |  | [optional] 
**Dnn** | Pointer to **string** | String representing a Data Network as defined in clause 9A of 3GPP TS 23.003;  it shall contain either a DNN Network Identifier, or a full DNN with both the Network  Identifier and Operator Identifier, as specified in 3GPP TS 23.003 clause 9.1.1 and 9.1.2. It shall be coded as string in which the labels are separated by dots  (e.g. \&quot;Label1.Label2.Label3\&quot;).  | [optional] 
**Snssai** | Pointer to [**Snssai**](Snssai.md) |  | [optional] 
**ExternalGroupId** | Pointer to **string** | string containing a local identifier followed by \&quot;@\&quot; and a domain identifier. Both the local identifier and the domain identifier shall be encoded as strings that do not contain any \&quot;@\&quot; characters. See Clauses 4.6.2 and 4.6.3 of 3GPP TS 23.682 for more information. | [optional] 
**DnaiInfos** | Pointer to [**map[string]DnaiInformation**](DnaiInformation.md) | list of DNS server identifier (consisting of IP address and port) and/or IP address(s) of the EAS in the local DN for each DNAI. The key of map is the DNAI.  | [optional] 

## Methods

### NewEasDeployInfo

`func NewEasDeployInfo(fqdnPatternList []FqdnPatternMatchingRule, ) *EasDeployInfo`

NewEasDeployInfo instantiates a new EasDeployInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEasDeployInfoWithDefaults

`func NewEasDeployInfoWithDefaults() *EasDeployInfo`

NewEasDeployInfoWithDefaults instantiates a new EasDeployInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSelf

`func (o *EasDeployInfo) GetSelf() string`

GetSelf returns the Self field if non-nil, zero value otherwise.

### GetSelfOk

`func (o *EasDeployInfo) GetSelfOk() (*string, bool)`

GetSelfOk returns a tuple with the Self field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelf

`func (o *EasDeployInfo) SetSelf(v string)`

SetSelf sets Self field to given value.

### HasSelf

`func (o *EasDeployInfo) HasSelf() bool`

HasSelf returns a boolean if a field has been set.

### GetAfServiceId

`func (o *EasDeployInfo) GetAfServiceId() string`

GetAfServiceId returns the AfServiceId field if non-nil, zero value otherwise.

### GetAfServiceIdOk

`func (o *EasDeployInfo) GetAfServiceIdOk() (*string, bool)`

GetAfServiceIdOk returns a tuple with the AfServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAfServiceId

`func (o *EasDeployInfo) SetAfServiceId(v string)`

SetAfServiceId sets AfServiceId field to given value.

### HasAfServiceId

`func (o *EasDeployInfo) HasAfServiceId() bool`

HasAfServiceId returns a boolean if a field has been set.

### GetFqdnPatternList

`func (o *EasDeployInfo) GetFqdnPatternList() []FqdnPatternMatchingRule`

GetFqdnPatternList returns the FqdnPatternList field if non-nil, zero value otherwise.

### GetFqdnPatternListOk

`func (o *EasDeployInfo) GetFqdnPatternListOk() (*[]FqdnPatternMatchingRule, bool)`

GetFqdnPatternListOk returns a tuple with the FqdnPatternList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFqdnPatternList

`func (o *EasDeployInfo) SetFqdnPatternList(v []FqdnPatternMatchingRule)`

SetFqdnPatternList sets FqdnPatternList field to given value.


### GetAppId

`func (o *EasDeployInfo) GetAppId() string`

GetAppId returns the AppId field if non-nil, zero value otherwise.

### GetAppIdOk

`func (o *EasDeployInfo) GetAppIdOk() (*string, bool)`

GetAppIdOk returns a tuple with the AppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppId

`func (o *EasDeployInfo) SetAppId(v string)`

SetAppId sets AppId field to given value.

### HasAppId

`func (o *EasDeployInfo) HasAppId() bool`

HasAppId returns a boolean if a field has been set.

### GetDnn

`func (o *EasDeployInfo) GetDnn() string`

GetDnn returns the Dnn field if non-nil, zero value otherwise.

### GetDnnOk

`func (o *EasDeployInfo) GetDnnOk() (*string, bool)`

GetDnnOk returns a tuple with the Dnn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnn

`func (o *EasDeployInfo) SetDnn(v string)`

SetDnn sets Dnn field to given value.

### HasDnn

`func (o *EasDeployInfo) HasDnn() bool`

HasDnn returns a boolean if a field has been set.

### GetSnssai

`func (o *EasDeployInfo) GetSnssai() Snssai`

GetSnssai returns the Snssai field if non-nil, zero value otherwise.

### GetSnssaiOk

`func (o *EasDeployInfo) GetSnssaiOk() (*Snssai, bool)`

GetSnssaiOk returns a tuple with the Snssai field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnssai

`func (o *EasDeployInfo) SetSnssai(v Snssai)`

SetSnssai sets Snssai field to given value.

### HasSnssai

`func (o *EasDeployInfo) HasSnssai() bool`

HasSnssai returns a boolean if a field has been set.

### GetExternalGroupId

`func (o *EasDeployInfo) GetExternalGroupId() string`

GetExternalGroupId returns the ExternalGroupId field if non-nil, zero value otherwise.

### GetExternalGroupIdOk

`func (o *EasDeployInfo) GetExternalGroupIdOk() (*string, bool)`

GetExternalGroupIdOk returns a tuple with the ExternalGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalGroupId

`func (o *EasDeployInfo) SetExternalGroupId(v string)`

SetExternalGroupId sets ExternalGroupId field to given value.

### HasExternalGroupId

`func (o *EasDeployInfo) HasExternalGroupId() bool`

HasExternalGroupId returns a boolean if a field has been set.

### GetDnaiInfos

`func (o *EasDeployInfo) GetDnaiInfos() map[string]DnaiInformation`

GetDnaiInfos returns the DnaiInfos field if non-nil, zero value otherwise.

### GetDnaiInfosOk

`func (o *EasDeployInfo) GetDnaiInfosOk() (*map[string]DnaiInformation, bool)`

GetDnaiInfosOk returns a tuple with the DnaiInfos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnaiInfos

`func (o *EasDeployInfo) SetDnaiInfos(v map[string]DnaiInformation)`

SetDnaiInfos sets DnaiInfos field to given value.

### HasDnaiInfos

`func (o *EasDeployInfo) HasDnaiInfos() bool`

HasDnaiInfos returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


