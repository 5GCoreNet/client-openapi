# MbsSecurityContext

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**KeyList** | [**map[string]MbsKeyInfo**](MbsKeyInfo.md) | A map (list of key-value pairs) where a (unique) valid JSON string serves as key of MbsSecurityContext | 

## Methods

### NewMbsSecurityContext

`func NewMbsSecurityContext(keyList map[string]MbsKeyInfo, ) *MbsSecurityContext`

NewMbsSecurityContext instantiates a new MbsSecurityContext object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMbsSecurityContextWithDefaults

`func NewMbsSecurityContextWithDefaults() *MbsSecurityContext`

NewMbsSecurityContextWithDefaults instantiates a new MbsSecurityContext object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKeyList

`func (o *MbsSecurityContext) GetKeyList() map[string]MbsKeyInfo`

GetKeyList returns the KeyList field if non-nil, zero value otherwise.

### GetKeyListOk

`func (o *MbsSecurityContext) GetKeyListOk() (*map[string]MbsKeyInfo, bool)`

GetKeyListOk returns a tuple with the KeyList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyList

`func (o *MbsSecurityContext) SetKeyList(v map[string]MbsKeyInfo)`

SetKeyList sets KeyList field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


