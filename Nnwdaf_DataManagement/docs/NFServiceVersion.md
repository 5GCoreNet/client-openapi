# NFServiceVersion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersionInUri** | **string** |  | 
**ApiFullVersion** | **string** |  | 
**Expiry** | Pointer to **time.Time** | string with format &#39;date-time&#39; as defined in OpenAPI. | [optional] 

## Methods

### NewNFServiceVersion

`func NewNFServiceVersion(apiVersionInUri string, apiFullVersion string, ) *NFServiceVersion`

NewNFServiceVersion instantiates a new NFServiceVersion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNFServiceVersionWithDefaults

`func NewNFServiceVersionWithDefaults() *NFServiceVersion`

NewNFServiceVersionWithDefaults instantiates a new NFServiceVersion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersionInUri

`func (o *NFServiceVersion) GetApiVersionInUri() string`

GetApiVersionInUri returns the ApiVersionInUri field if non-nil, zero value otherwise.

### GetApiVersionInUriOk

`func (o *NFServiceVersion) GetApiVersionInUriOk() (*string, bool)`

GetApiVersionInUriOk returns a tuple with the ApiVersionInUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersionInUri

`func (o *NFServiceVersion) SetApiVersionInUri(v string)`

SetApiVersionInUri sets ApiVersionInUri field to given value.


### GetApiFullVersion

`func (o *NFServiceVersion) GetApiFullVersion() string`

GetApiFullVersion returns the ApiFullVersion field if non-nil, zero value otherwise.

### GetApiFullVersionOk

`func (o *NFServiceVersion) GetApiFullVersionOk() (*string, bool)`

GetApiFullVersionOk returns a tuple with the ApiFullVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiFullVersion

`func (o *NFServiceVersion) SetApiFullVersion(v string)`

SetApiFullVersion sets ApiFullVersion field to given value.


### GetExpiry

`func (o *NFServiceVersion) GetExpiry() time.Time`

GetExpiry returns the Expiry field if non-nil, zero value otherwise.

### GetExpiryOk

`func (o *NFServiceVersion) GetExpiryOk() (*time.Time, bool)`

GetExpiryOk returns a tuple with the Expiry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiry

`func (o *NFServiceVersion) SetExpiry(v time.Time)`

SetExpiry sets Expiry field to given value.

### HasExpiry

`func (o *NFServiceVersion) HasExpiry() bool`

HasExpiry returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


