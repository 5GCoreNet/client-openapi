# PushInfoResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MeKeyMaterial** | **string** | ME Key Material (hex-encoded string) | 
**GbaPushInfo** | **string** | GBA Push Info (hex-encoded string) | 
**UiccKeyMaterial** | Pointer to **string** | UICC key material (hex-encoded string) | [optional] 
**KeyExpiryTime** | Pointer to **time.Time** | string with format &#39;date-time&#39; as defined in OpenAPI. | [optional] 
**BootstrappingInfoCreationTime** | Pointer to **time.Time** | string with format &#39;date-time&#39; as defined in OpenAPI. | [optional] 
**UssList** | Pointer to [**[]UssListItem**](UssListItem.md) |  | [optional] 
**GbaType** | Pointer to [**GbaType**](GbaType.md) |  | [optional] 
**Impi** | Pointer to **string** | IMS Private Identity of the UE | [optional] 
**SecurityFeaturesResponse** | Pointer to [**[]SecFeature**](SecFeature.md) |  | [optional] 

## Methods

### NewPushInfoResponse

`func NewPushInfoResponse(meKeyMaterial string, gbaPushInfo string, ) *PushInfoResponse`

NewPushInfoResponse instantiates a new PushInfoResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPushInfoResponseWithDefaults

`func NewPushInfoResponseWithDefaults() *PushInfoResponse`

NewPushInfoResponseWithDefaults instantiates a new PushInfoResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeKeyMaterial

`func (o *PushInfoResponse) GetMeKeyMaterial() string`

GetMeKeyMaterial returns the MeKeyMaterial field if non-nil, zero value otherwise.

### GetMeKeyMaterialOk

`func (o *PushInfoResponse) GetMeKeyMaterialOk() (*string, bool)`

GetMeKeyMaterialOk returns a tuple with the MeKeyMaterial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeKeyMaterial

`func (o *PushInfoResponse) SetMeKeyMaterial(v string)`

SetMeKeyMaterial sets MeKeyMaterial field to given value.


### GetGbaPushInfo

`func (o *PushInfoResponse) GetGbaPushInfo() string`

GetGbaPushInfo returns the GbaPushInfo field if non-nil, zero value otherwise.

### GetGbaPushInfoOk

`func (o *PushInfoResponse) GetGbaPushInfoOk() (*string, bool)`

GetGbaPushInfoOk returns a tuple with the GbaPushInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGbaPushInfo

`func (o *PushInfoResponse) SetGbaPushInfo(v string)`

SetGbaPushInfo sets GbaPushInfo field to given value.


### GetUiccKeyMaterial

`func (o *PushInfoResponse) GetUiccKeyMaterial() string`

GetUiccKeyMaterial returns the UiccKeyMaterial field if non-nil, zero value otherwise.

### GetUiccKeyMaterialOk

`func (o *PushInfoResponse) GetUiccKeyMaterialOk() (*string, bool)`

GetUiccKeyMaterialOk returns a tuple with the UiccKeyMaterial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUiccKeyMaterial

`func (o *PushInfoResponse) SetUiccKeyMaterial(v string)`

SetUiccKeyMaterial sets UiccKeyMaterial field to given value.

### HasUiccKeyMaterial

`func (o *PushInfoResponse) HasUiccKeyMaterial() bool`

HasUiccKeyMaterial returns a boolean if a field has been set.

### GetKeyExpiryTime

`func (o *PushInfoResponse) GetKeyExpiryTime() time.Time`

GetKeyExpiryTime returns the KeyExpiryTime field if non-nil, zero value otherwise.

### GetKeyExpiryTimeOk

`func (o *PushInfoResponse) GetKeyExpiryTimeOk() (*time.Time, bool)`

GetKeyExpiryTimeOk returns a tuple with the KeyExpiryTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyExpiryTime

`func (o *PushInfoResponse) SetKeyExpiryTime(v time.Time)`

SetKeyExpiryTime sets KeyExpiryTime field to given value.

### HasKeyExpiryTime

`func (o *PushInfoResponse) HasKeyExpiryTime() bool`

HasKeyExpiryTime returns a boolean if a field has been set.

### GetBootstrappingInfoCreationTime

`func (o *PushInfoResponse) GetBootstrappingInfoCreationTime() time.Time`

GetBootstrappingInfoCreationTime returns the BootstrappingInfoCreationTime field if non-nil, zero value otherwise.

### GetBootstrappingInfoCreationTimeOk

`func (o *PushInfoResponse) GetBootstrappingInfoCreationTimeOk() (*time.Time, bool)`

GetBootstrappingInfoCreationTimeOk returns a tuple with the BootstrappingInfoCreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootstrappingInfoCreationTime

`func (o *PushInfoResponse) SetBootstrappingInfoCreationTime(v time.Time)`

SetBootstrappingInfoCreationTime sets BootstrappingInfoCreationTime field to given value.

### HasBootstrappingInfoCreationTime

`func (o *PushInfoResponse) HasBootstrappingInfoCreationTime() bool`

HasBootstrappingInfoCreationTime returns a boolean if a field has been set.

### GetUssList

`func (o *PushInfoResponse) GetUssList() []UssListItem`

GetUssList returns the UssList field if non-nil, zero value otherwise.

### GetUssListOk

`func (o *PushInfoResponse) GetUssListOk() (*[]UssListItem, bool)`

GetUssListOk returns a tuple with the UssList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUssList

`func (o *PushInfoResponse) SetUssList(v []UssListItem)`

SetUssList sets UssList field to given value.

### HasUssList

`func (o *PushInfoResponse) HasUssList() bool`

HasUssList returns a boolean if a field has been set.

### GetGbaType

`func (o *PushInfoResponse) GetGbaType() GbaType`

GetGbaType returns the GbaType field if non-nil, zero value otherwise.

### GetGbaTypeOk

`func (o *PushInfoResponse) GetGbaTypeOk() (*GbaType, bool)`

GetGbaTypeOk returns a tuple with the GbaType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGbaType

`func (o *PushInfoResponse) SetGbaType(v GbaType)`

SetGbaType sets GbaType field to given value.

### HasGbaType

`func (o *PushInfoResponse) HasGbaType() bool`

HasGbaType returns a boolean if a field has been set.

### GetImpi

`func (o *PushInfoResponse) GetImpi() string`

GetImpi returns the Impi field if non-nil, zero value otherwise.

### GetImpiOk

`func (o *PushInfoResponse) GetImpiOk() (*string, bool)`

GetImpiOk returns a tuple with the Impi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImpi

`func (o *PushInfoResponse) SetImpi(v string)`

SetImpi sets Impi field to given value.

### HasImpi

`func (o *PushInfoResponse) HasImpi() bool`

HasImpi returns a boolean if a field has been set.

### GetSecurityFeaturesResponse

`func (o *PushInfoResponse) GetSecurityFeaturesResponse() []SecFeature`

GetSecurityFeaturesResponse returns the SecurityFeaturesResponse field if non-nil, zero value otherwise.

### GetSecurityFeaturesResponseOk

`func (o *PushInfoResponse) GetSecurityFeaturesResponseOk() (*[]SecFeature, bool)`

GetSecurityFeaturesResponseOk returns a tuple with the SecurityFeaturesResponse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityFeaturesResponse

`func (o *PushInfoResponse) SetSecurityFeaturesResponse(v []SecFeature)`

SetSecurityFeaturesResponse sets SecurityFeaturesResponse field to given value.

### HasSecurityFeaturesResponse

`func (o *PushInfoResponse) HasSecurityFeaturesResponse() bool`

HasSecurityFeaturesResponse returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


