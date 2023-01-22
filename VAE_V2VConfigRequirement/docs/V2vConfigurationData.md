# V2vConfigurationData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GroupId** | Pointer to **string** | Represents the group ID for which a V2X message is addressed. | [optional] 
**ServiceId** | Pointer to **string** | Represents the V2X service ID to which a V2X message belongs. | [optional] 
**CanUeIds** | Pointer to **[]string** |  | [optional] 
**AppQosReq** | Pointer to [**AppplicationQosRequirement**](AppplicationQosRequirement.md) |  | [optional] 
**SuppFeat** | Pointer to **string** | A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \&quot;0\&quot; to \&quot;9\&quot;,  \&quot;a\&quot; to \&quot;f\&quot; or \&quot;A\&quot; to \&quot;F\&quot; and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.  | [optional] 

## Methods

### NewV2vConfigurationData

`func NewV2vConfigurationData() *V2vConfigurationData`

NewV2vConfigurationData instantiates a new V2vConfigurationData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2vConfigurationDataWithDefaults

`func NewV2vConfigurationDataWithDefaults() *V2vConfigurationData`

NewV2vConfigurationDataWithDefaults instantiates a new V2vConfigurationData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroupId

`func (o *V2vConfigurationData) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *V2vConfigurationData) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *V2vConfigurationData) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *V2vConfigurationData) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetServiceId

`func (o *V2vConfigurationData) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *V2vConfigurationData) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *V2vConfigurationData) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.

### HasServiceId

`func (o *V2vConfigurationData) HasServiceId() bool`

HasServiceId returns a boolean if a field has been set.

### GetCanUeIds

`func (o *V2vConfigurationData) GetCanUeIds() []string`

GetCanUeIds returns the CanUeIds field if non-nil, zero value otherwise.

### GetCanUeIdsOk

`func (o *V2vConfigurationData) GetCanUeIdsOk() (*[]string, bool)`

GetCanUeIdsOk returns a tuple with the CanUeIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanUeIds

`func (o *V2vConfigurationData) SetCanUeIds(v []string)`

SetCanUeIds sets CanUeIds field to given value.

### HasCanUeIds

`func (o *V2vConfigurationData) HasCanUeIds() bool`

HasCanUeIds returns a boolean if a field has been set.

### GetAppQosReq

`func (o *V2vConfigurationData) GetAppQosReq() AppplicationQosRequirement`

GetAppQosReq returns the AppQosReq field if non-nil, zero value otherwise.

### GetAppQosReqOk

`func (o *V2vConfigurationData) GetAppQosReqOk() (*AppplicationQosRequirement, bool)`

GetAppQosReqOk returns a tuple with the AppQosReq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppQosReq

`func (o *V2vConfigurationData) SetAppQosReq(v AppplicationQosRequirement)`

SetAppQosReq sets AppQosReq field to given value.

### HasAppQosReq

`func (o *V2vConfigurationData) HasAppQosReq() bool`

HasAppQosReq returns a boolean if a field has been set.

### GetSuppFeat

`func (o *V2vConfigurationData) GetSuppFeat() string`

GetSuppFeat returns the SuppFeat field if non-nil, zero value otherwise.

### GetSuppFeatOk

`func (o *V2vConfigurationData) GetSuppFeatOk() (*string, bool)`

GetSuppFeatOk returns a tuple with the SuppFeat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuppFeat

`func (o *V2vConfigurationData) SetSuppFeat(v string)`

SetSuppFeat sets SuppFeat field to given value.

### HasSuppFeat

`func (o *V2vConfigurationData) HasSuppFeat() bool`

HasSuppFeat returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


