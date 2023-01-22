# AppDescriptorRm

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppIds** | Pointer to **map[string]string** | Identifies applications that are running on the UE&#39;s operating system. Any string value can be used as a key of the map.  | [optional] 

## Methods

### NewAppDescriptorRm

`func NewAppDescriptorRm() *AppDescriptorRm`

NewAppDescriptorRm instantiates a new AppDescriptorRm object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppDescriptorRmWithDefaults

`func NewAppDescriptorRmWithDefaults() *AppDescriptorRm`

NewAppDescriptorRmWithDefaults instantiates a new AppDescriptorRm object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppIds

`func (o *AppDescriptorRm) GetAppIds() map[string]string`

GetAppIds returns the AppIds field if non-nil, zero value otherwise.

### GetAppIdsOk

`func (o *AppDescriptorRm) GetAppIdsOk() (*map[string]string, bool)`

GetAppIdsOk returns a tuple with the AppIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppIds

`func (o *AppDescriptorRm) SetAppIds(v map[string]string)`

SetAppIds sets AppIds field to given value.

### HasAppIds

`func (o *AppDescriptorRm) HasAppIds() bool`

HasAppIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


