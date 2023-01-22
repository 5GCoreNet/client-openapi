# UasId

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GroupId** | Pointer to **string** | string containing a local identifier followed by \&quot;@\&quot; and a domain identifier. Both the local identifier and the domain identifier shall be encoded as strings that do not contain any \&quot;@\&quot; characters. See Clauses 4.6.2 and 4.6.3 of 3GPP TS 23.682 for more information. | [optional] 
**IndividualUasId** | Pointer to [**[]UavId**](UavId.md) |  | [optional] 

## Methods

### NewUasId

`func NewUasId() *UasId`

NewUasId instantiates a new UasId object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUasIdWithDefaults

`func NewUasIdWithDefaults() *UasId`

NewUasIdWithDefaults instantiates a new UasId object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroupId

`func (o *UasId) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *UasId) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *UasId) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *UasId) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetIndividualUasId

`func (o *UasId) GetIndividualUasId() []UavId`

GetIndividualUasId returns the IndividualUasId field if non-nil, zero value otherwise.

### GetIndividualUasIdOk

`func (o *UasId) GetIndividualUasIdOk() (*[]UavId, bool)`

GetIndividualUasIdOk returns a tuple with the IndividualUasId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndividualUasId

`func (o *UasId) SetIndividualUasId(v []UavId)`

SetIndividualUasId sets IndividualUasId field to given value.

### HasIndividualUasId

`func (o *UasId) HasIndividualUasId() bool`

HasIndividualUasId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


