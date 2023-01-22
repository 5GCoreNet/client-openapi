# BootstrappingInfoResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MeKeyMaterial** | **string** | ME Key Material (hex-encoded string) | 
**UiccKeyMaterial** | Pointer to **string** | UICC key material (hex-encoded string) | [optional] 
**KeyExpiryTime** | Pointer to **time.Time** | string with format &#39;date-time&#39; as defined in OpenAPI. | [optional] 
**BootstrappingInfoCreationTime** | Pointer to **time.Time** | string with format &#39;date-time&#39; as defined in OpenAPI. | [optional] 
**UssList** | Pointer to [**[]UssListItem**](UssListItem.md) |  | [optional] 
**GbaType** | Pointer to [**GbaType**](GbaType.md) |  | [optional] 
**Impi** | Pointer to **string** | IMS Private Identity of the UE | [optional] 

## Methods

### NewBootstrappingInfoResponse

`func NewBootstrappingInfoResponse(meKeyMaterial string, ) *BootstrappingInfoResponse`

NewBootstrappingInfoResponse instantiates a new BootstrappingInfoResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBootstrappingInfoResponseWithDefaults

`func NewBootstrappingInfoResponseWithDefaults() *BootstrappingInfoResponse`

NewBootstrappingInfoResponseWithDefaults instantiates a new BootstrappingInfoResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeKeyMaterial

`func (o *BootstrappingInfoResponse) GetMeKeyMaterial() string`

GetMeKeyMaterial returns the MeKeyMaterial field if non-nil, zero value otherwise.

### GetMeKeyMaterialOk

`func (o *BootstrappingInfoResponse) GetMeKeyMaterialOk() (*string, bool)`

GetMeKeyMaterialOk returns a tuple with the MeKeyMaterial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeKeyMaterial

`func (o *BootstrappingInfoResponse) SetMeKeyMaterial(v string)`

SetMeKeyMaterial sets MeKeyMaterial field to given value.


### GetUiccKeyMaterial

`func (o *BootstrappingInfoResponse) GetUiccKeyMaterial() string`

GetUiccKeyMaterial returns the UiccKeyMaterial field if non-nil, zero value otherwise.

### GetUiccKeyMaterialOk

`func (o *BootstrappingInfoResponse) GetUiccKeyMaterialOk() (*string, bool)`

GetUiccKeyMaterialOk returns a tuple with the UiccKeyMaterial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUiccKeyMaterial

`func (o *BootstrappingInfoResponse) SetUiccKeyMaterial(v string)`

SetUiccKeyMaterial sets UiccKeyMaterial field to given value.

### HasUiccKeyMaterial

`func (o *BootstrappingInfoResponse) HasUiccKeyMaterial() bool`

HasUiccKeyMaterial returns a boolean if a field has been set.

### GetKeyExpiryTime

`func (o *BootstrappingInfoResponse) GetKeyExpiryTime() time.Time`

GetKeyExpiryTime returns the KeyExpiryTime field if non-nil, zero value otherwise.

### GetKeyExpiryTimeOk

`func (o *BootstrappingInfoResponse) GetKeyExpiryTimeOk() (*time.Time, bool)`

GetKeyExpiryTimeOk returns a tuple with the KeyExpiryTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyExpiryTime

`func (o *BootstrappingInfoResponse) SetKeyExpiryTime(v time.Time)`

SetKeyExpiryTime sets KeyExpiryTime field to given value.

### HasKeyExpiryTime

`func (o *BootstrappingInfoResponse) HasKeyExpiryTime() bool`

HasKeyExpiryTime returns a boolean if a field has been set.

### GetBootstrappingInfoCreationTime

`func (o *BootstrappingInfoResponse) GetBootstrappingInfoCreationTime() time.Time`

GetBootstrappingInfoCreationTime returns the BootstrappingInfoCreationTime field if non-nil, zero value otherwise.

### GetBootstrappingInfoCreationTimeOk

`func (o *BootstrappingInfoResponse) GetBootstrappingInfoCreationTimeOk() (*time.Time, bool)`

GetBootstrappingInfoCreationTimeOk returns a tuple with the BootstrappingInfoCreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootstrappingInfoCreationTime

`func (o *BootstrappingInfoResponse) SetBootstrappingInfoCreationTime(v time.Time)`

SetBootstrappingInfoCreationTime sets BootstrappingInfoCreationTime field to given value.

### HasBootstrappingInfoCreationTime

`func (o *BootstrappingInfoResponse) HasBootstrappingInfoCreationTime() bool`

HasBootstrappingInfoCreationTime returns a boolean if a field has been set.

### GetUssList

`func (o *BootstrappingInfoResponse) GetUssList() []UssListItem`

GetUssList returns the UssList field if non-nil, zero value otherwise.

### GetUssListOk

`func (o *BootstrappingInfoResponse) GetUssListOk() (*[]UssListItem, bool)`

GetUssListOk returns a tuple with the UssList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUssList

`func (o *BootstrappingInfoResponse) SetUssList(v []UssListItem)`

SetUssList sets UssList field to given value.

### HasUssList

`func (o *BootstrappingInfoResponse) HasUssList() bool`

HasUssList returns a boolean if a field has been set.

### GetGbaType

`func (o *BootstrappingInfoResponse) GetGbaType() GbaType`

GetGbaType returns the GbaType field if non-nil, zero value otherwise.

### GetGbaTypeOk

`func (o *BootstrappingInfoResponse) GetGbaTypeOk() (*GbaType, bool)`

GetGbaTypeOk returns a tuple with the GbaType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGbaType

`func (o *BootstrappingInfoResponse) SetGbaType(v GbaType)`

SetGbaType sets GbaType field to given value.

### HasGbaType

`func (o *BootstrappingInfoResponse) HasGbaType() bool`

HasGbaType returns a boolean if a field has been set.

### GetImpi

`func (o *BootstrappingInfoResponse) GetImpi() string`

GetImpi returns the Impi field if non-nil, zero value otherwise.

### GetImpiOk

`func (o *BootstrappingInfoResponse) GetImpiOk() (*string, bool)`

GetImpiOk returns a tuple with the Impi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImpi

`func (o *BootstrappingInfoResponse) SetImpi(v string)`

SetImpi sets Impi field to given value.

### HasImpi

`func (o *BootstrappingInfoResponse) HasImpi() bool`

HasImpi returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


