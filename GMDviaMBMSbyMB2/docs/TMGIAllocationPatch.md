# TMGIAllocationPatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExternalGroupId** | Pointer to **string** | string containing a local identifier followed by \&quot;@\&quot; and a domain identifier. Both the local identifier and the domain identifier shall be encoded as strings that do not contain any \&quot;@\&quot; characters. See Clauses 4.6.2 and 4.6.3 of 3GPP TS 23.682 for more information. | [optional] 
**MbmsLocArea** | Pointer to [**MbmsLocArea**](MbmsLocArea.md) |  | [optional] 

## Methods

### NewTMGIAllocationPatch

`func NewTMGIAllocationPatch() *TMGIAllocationPatch`

NewTMGIAllocationPatch instantiates a new TMGIAllocationPatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTMGIAllocationPatchWithDefaults

`func NewTMGIAllocationPatchWithDefaults() *TMGIAllocationPatch`

NewTMGIAllocationPatchWithDefaults instantiates a new TMGIAllocationPatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExternalGroupId

`func (o *TMGIAllocationPatch) GetExternalGroupId() string`

GetExternalGroupId returns the ExternalGroupId field if non-nil, zero value otherwise.

### GetExternalGroupIdOk

`func (o *TMGIAllocationPatch) GetExternalGroupIdOk() (*string, bool)`

GetExternalGroupIdOk returns a tuple with the ExternalGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalGroupId

`func (o *TMGIAllocationPatch) SetExternalGroupId(v string)`

SetExternalGroupId sets ExternalGroupId field to given value.

### HasExternalGroupId

`func (o *TMGIAllocationPatch) HasExternalGroupId() bool`

HasExternalGroupId returns a boolean if a field has been set.

### GetMbmsLocArea

`func (o *TMGIAllocationPatch) GetMbmsLocArea() MbmsLocArea`

GetMbmsLocArea returns the MbmsLocArea field if non-nil, zero value otherwise.

### GetMbmsLocAreaOk

`func (o *TMGIAllocationPatch) GetMbmsLocAreaOk() (*MbmsLocArea, bool)`

GetMbmsLocAreaOk returns a tuple with the MbmsLocArea field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMbmsLocArea

`func (o *TMGIAllocationPatch) SetMbmsLocArea(v MbmsLocArea)`

SetMbmsLocArea sets MbmsLocArea field to given value.

### HasMbmsLocArea

`func (o *TMGIAllocationPatch) HasMbmsLocArea() bool`

HasMbmsLocArea returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


