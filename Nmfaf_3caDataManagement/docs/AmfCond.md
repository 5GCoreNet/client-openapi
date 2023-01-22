# AmfCond

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AmfSetId** | Pointer to **string** | String identifying the AMF Set ID (10 bits) as specified in clause 2.10.1 of 3GPP TS 23.003.  It is encoded as a string of 3 hexadecimal characters where the first character is limited to  values 0 to 3 (i.e. 10 bits).  | [optional] 
**AmfRegionId** | Pointer to **string** | String identifying the AMF Set ID (10 bits) as specified in clause 2.10.1 of 3GPP TS 23.003.  It is encoded as a string of 3 hexadecimal characters where the first character is limited to  values 0 to 3 (i.e. 10 bits)  | [optional] 

## Methods

### NewAmfCond

`func NewAmfCond() *AmfCond`

NewAmfCond instantiates a new AmfCond object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAmfCondWithDefaults

`func NewAmfCondWithDefaults() *AmfCond`

NewAmfCondWithDefaults instantiates a new AmfCond object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmfSetId

`func (o *AmfCond) GetAmfSetId() string`

GetAmfSetId returns the AmfSetId field if non-nil, zero value otherwise.

### GetAmfSetIdOk

`func (o *AmfCond) GetAmfSetIdOk() (*string, bool)`

GetAmfSetIdOk returns a tuple with the AmfSetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmfSetId

`func (o *AmfCond) SetAmfSetId(v string)`

SetAmfSetId sets AmfSetId field to given value.

### HasAmfSetId

`func (o *AmfCond) HasAmfSetId() bool`

HasAmfSetId returns a boolean if a field has been set.

### GetAmfRegionId

`func (o *AmfCond) GetAmfRegionId() string`

GetAmfRegionId returns the AmfRegionId field if non-nil, zero value otherwise.

### GetAmfRegionIdOk

`func (o *AmfCond) GetAmfRegionIdOk() (*string, bool)`

GetAmfRegionIdOk returns a tuple with the AmfRegionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmfRegionId

`func (o *AmfCond) SetAmfRegionId(v string)`

SetAmfRegionId sets AmfRegionId field to given value.

### HasAmfRegionId

`func (o *AmfCond) HasAmfRegionId() bool`

HasAmfRegionId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


