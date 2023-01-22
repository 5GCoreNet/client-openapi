# PduidInformation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Expiry** | **time.Time** | string with format &#39;date-time&#39; as defined in OpenAPI. | 
**Pduid** | **string** | Contains the PDUID. | 

## Methods

### NewPduidInformation

`func NewPduidInformation(expiry time.Time, pduid string, ) *PduidInformation`

NewPduidInformation instantiates a new PduidInformation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPduidInformationWithDefaults

`func NewPduidInformationWithDefaults() *PduidInformation`

NewPduidInformationWithDefaults instantiates a new PduidInformation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpiry

`func (o *PduidInformation) GetExpiry() time.Time`

GetExpiry returns the Expiry field if non-nil, zero value otherwise.

### GetExpiryOk

`func (o *PduidInformation) GetExpiryOk() (*time.Time, bool)`

GetExpiryOk returns a tuple with the Expiry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiry

`func (o *PduidInformation) SetExpiry(v time.Time)`

SetExpiry sets Expiry field to given value.


### GetPduid

`func (o *PduidInformation) GetPduid() string`

GetPduid returns the Pduid field if non-nil, zero value otherwise.

### GetPduidOk

`func (o *PduidInformation) GetPduidOk() (*string, bool)`

GetPduidOk returns a tuple with the Pduid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPduid

`func (o *PduidInformation) SetPduid(v string)`

SetPduid sets Pduid field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


