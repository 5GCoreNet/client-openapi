# AcsConfigurationDataPatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AcsInfo** | Pointer to [**AcsInfo**](AcsInfo.md) |  | [optional] 
**MtcProviderId** | Pointer to **string** | String uniquely identifying MTC provider information. | [optional] 

## Methods

### NewAcsConfigurationDataPatch

`func NewAcsConfigurationDataPatch() *AcsConfigurationDataPatch`

NewAcsConfigurationDataPatch instantiates a new AcsConfigurationDataPatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAcsConfigurationDataPatchWithDefaults

`func NewAcsConfigurationDataPatchWithDefaults() *AcsConfigurationDataPatch`

NewAcsConfigurationDataPatchWithDefaults instantiates a new AcsConfigurationDataPatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAcsInfo

`func (o *AcsConfigurationDataPatch) GetAcsInfo() AcsInfo`

GetAcsInfo returns the AcsInfo field if non-nil, zero value otherwise.

### GetAcsInfoOk

`func (o *AcsConfigurationDataPatch) GetAcsInfoOk() (*AcsInfo, bool)`

GetAcsInfoOk returns a tuple with the AcsInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcsInfo

`func (o *AcsConfigurationDataPatch) SetAcsInfo(v AcsInfo)`

SetAcsInfo sets AcsInfo field to given value.

### HasAcsInfo

`func (o *AcsConfigurationDataPatch) HasAcsInfo() bool`

HasAcsInfo returns a boolean if a field has been set.

### GetMtcProviderId

`func (o *AcsConfigurationDataPatch) GetMtcProviderId() string`

GetMtcProviderId returns the MtcProviderId field if non-nil, zero value otherwise.

### GetMtcProviderIdOk

`func (o *AcsConfigurationDataPatch) GetMtcProviderIdOk() (*string, bool)`

GetMtcProviderIdOk returns a tuple with the MtcProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtcProviderId

`func (o *AcsConfigurationDataPatch) SetMtcProviderId(v string)`

SetMtcProviderId sets MtcProviderId field to given value.

### HasMtcProviderId

`func (o *AcsConfigurationDataPatch) HasMtcProviderId() bool`

HasMtcProviderId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


