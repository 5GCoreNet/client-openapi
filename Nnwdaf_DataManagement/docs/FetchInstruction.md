# FetchInstruction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FetchUri** | **string** | String providing an URI formatted according to RFC 3986. | 
**FetchCorrIds** | **[]string** | The fetch correlation identifier(s) of the MFAF Data or Analytics. | 
**Expiry** | Pointer to **time.Time** | string with format &#39;date-time&#39; as defined in OpenAPI. | [optional] 

## Methods

### NewFetchInstruction

`func NewFetchInstruction(fetchUri string, fetchCorrIds []string, ) *FetchInstruction`

NewFetchInstruction instantiates a new FetchInstruction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFetchInstructionWithDefaults

`func NewFetchInstructionWithDefaults() *FetchInstruction`

NewFetchInstructionWithDefaults instantiates a new FetchInstruction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFetchUri

`func (o *FetchInstruction) GetFetchUri() string`

GetFetchUri returns the FetchUri field if non-nil, zero value otherwise.

### GetFetchUriOk

`func (o *FetchInstruction) GetFetchUriOk() (*string, bool)`

GetFetchUriOk returns a tuple with the FetchUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFetchUri

`func (o *FetchInstruction) SetFetchUri(v string)`

SetFetchUri sets FetchUri field to given value.


### GetFetchCorrIds

`func (o *FetchInstruction) GetFetchCorrIds() []string`

GetFetchCorrIds returns the FetchCorrIds field if non-nil, zero value otherwise.

### GetFetchCorrIdsOk

`func (o *FetchInstruction) GetFetchCorrIdsOk() (*[]string, bool)`

GetFetchCorrIdsOk returns a tuple with the FetchCorrIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFetchCorrIds

`func (o *FetchInstruction) SetFetchCorrIds(v []string)`

SetFetchCorrIds sets FetchCorrIds field to given value.


### GetExpiry

`func (o *FetchInstruction) GetExpiry() time.Time`

GetExpiry returns the Expiry field if non-nil, zero value otherwise.

### GetExpiryOk

`func (o *FetchInstruction) GetExpiryOk() (*time.Time, bool)`

GetExpiryOk returns a tuple with the Expiry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiry

`func (o *FetchInstruction) SetExpiry(v time.Time)`

SetExpiry sets Expiry field to given value.

### HasExpiry

`func (o *FetchInstruction) HasExpiry() bool`

HasExpiry returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


