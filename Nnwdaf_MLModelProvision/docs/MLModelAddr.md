# MLModelAddr

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MLModelUrl** | Pointer to **string** | String providing an URI formatted according to RFC 3986. | [optional] 
**MlFileFqdn** | Pointer to **string** | The FQDN of the ML Model file. | [optional] 

## Methods

### NewMLModelAddr

`func NewMLModelAddr() *MLModelAddr`

NewMLModelAddr instantiates a new MLModelAddr object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMLModelAddrWithDefaults

`func NewMLModelAddrWithDefaults() *MLModelAddr`

NewMLModelAddrWithDefaults instantiates a new MLModelAddr object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMLModelUrl

`func (o *MLModelAddr) GetMLModelUrl() string`

GetMLModelUrl returns the MLModelUrl field if non-nil, zero value otherwise.

### GetMLModelUrlOk

`func (o *MLModelAddr) GetMLModelUrlOk() (*string, bool)`

GetMLModelUrlOk returns a tuple with the MLModelUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMLModelUrl

`func (o *MLModelAddr) SetMLModelUrl(v string)`

SetMLModelUrl sets MLModelUrl field to given value.

### HasMLModelUrl

`func (o *MLModelAddr) HasMLModelUrl() bool`

HasMLModelUrl returns a boolean if a field has been set.

### GetMlFileFqdn

`func (o *MLModelAddr) GetMlFileFqdn() string`

GetMlFileFqdn returns the MlFileFqdn field if non-nil, zero value otherwise.

### GetMlFileFqdnOk

`func (o *MLModelAddr) GetMlFileFqdnOk() (*string, bool)`

GetMlFileFqdnOk returns a tuple with the MlFileFqdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMlFileFqdn

`func (o *MLModelAddr) SetMlFileFqdn(v string)`

SetMlFileFqdn sets MlFileFqdn field to given value.

### HasMlFileFqdn

`func (o *MLModelAddr) HasMlFileFqdn() bool`

HasMlFileFqdn returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


