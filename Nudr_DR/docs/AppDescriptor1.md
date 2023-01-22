# AppDescriptor1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OsId** | **string** | Represents the Operating System of the served UE. | 
**AppIds** | **map[string]string** | Identifies applications that are running on the UE&#39;s operating system. Any string value can be used as a key of the map.  | 

## Methods

### NewAppDescriptor1

`func NewAppDescriptor1(osId string, appIds map[string]string, ) *AppDescriptor1`

NewAppDescriptor1 instantiates a new AppDescriptor1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppDescriptor1WithDefaults

`func NewAppDescriptor1WithDefaults() *AppDescriptor1`

NewAppDescriptor1WithDefaults instantiates a new AppDescriptor1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOsId

`func (o *AppDescriptor1) GetOsId() string`

GetOsId returns the OsId field if non-nil, zero value otherwise.

### GetOsIdOk

`func (o *AppDescriptor1) GetOsIdOk() (*string, bool)`

GetOsIdOk returns a tuple with the OsId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsId

`func (o *AppDescriptor1) SetOsId(v string)`

SetOsId sets OsId field to given value.


### GetAppIds

`func (o *AppDescriptor1) GetAppIds() map[string]string`

GetAppIds returns the AppIds field if non-nil, zero value otherwise.

### GetAppIdsOk

`func (o *AppDescriptor1) GetAppIdsOk() (*map[string]string, bool)`

GetAppIdsOk returns a tuple with the AppIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppIds

`func (o *AppDescriptor1) SetAppIds(v map[string]string)`

SetAppIds sets AppIds field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


