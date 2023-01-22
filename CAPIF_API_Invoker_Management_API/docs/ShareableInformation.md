# ShareableInformation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsShareable** | **bool** | Set to \&quot;true\&quot; indicates that the service API and/or the service API category can be shared to the list of CAPIF provider domain information. Otherwise set to \&quot;false\&quot;.  | 
**CapifProvDoms** | Pointer to **[]string** | List of CAPIF provider domains to which the service API information to be shared.  | [optional] 

## Methods

### NewShareableInformation

`func NewShareableInformation(isShareable bool, ) *ShareableInformation`

NewShareableInformation instantiates a new ShareableInformation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShareableInformationWithDefaults

`func NewShareableInformationWithDefaults() *ShareableInformation`

NewShareableInformationWithDefaults instantiates a new ShareableInformation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsShareable

`func (o *ShareableInformation) GetIsShareable() bool`

GetIsShareable returns the IsShareable field if non-nil, zero value otherwise.

### GetIsShareableOk

`func (o *ShareableInformation) GetIsShareableOk() (*bool, bool)`

GetIsShareableOk returns a tuple with the IsShareable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsShareable

`func (o *ShareableInformation) SetIsShareable(v bool)`

SetIsShareable sets IsShareable field to given value.


### GetCapifProvDoms

`func (o *ShareableInformation) GetCapifProvDoms() []string`

GetCapifProvDoms returns the CapifProvDoms field if non-nil, zero value otherwise.

### GetCapifProvDomsOk

`func (o *ShareableInformation) GetCapifProvDomsOk() (*[]string, bool)`

GetCapifProvDomsOk returns a tuple with the CapifProvDoms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapifProvDoms

`func (o *ShareableInformation) SetCapifProvDoms(v []string)`

SetCapifProvDoms sets CapifProvDoms field to given value.

### HasCapifProvDoms

`func (o *ShareableInformation) HasCapifProvDoms() bool`

HasCapifProvDoms returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


