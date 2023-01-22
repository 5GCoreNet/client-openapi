# ManagedNFProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NfInstanceID** | Pointer to **string** |  | [optional] 
**NfType** | Pointer to [**NFType**](NFType.md) |  | [optional] 
**HeartbeatTimer** | Pointer to **int32** |  | [optional] 
**AuthzInfo** | Pointer to **string** |  | [optional] 
**HostAddr** | Pointer to [**HostAddr**](HostAddr.md) |  | [optional] 
**AllowedPLMNs** | Pointer to [**[]PlmnId**](PlmnId.md) |  | [optional] 
**AllowedSNPNs** | Pointer to [**[]SnpnInfo**](SnpnInfo.md) |  | [optional] 
**AllowedNfTypes** | Pointer to [**[]NFType**](NFType.md) |  | [optional] 
**AllowedNfDomains** | Pointer to **[]string** |  | [optional] 
**AllowedNSSAIs** | Pointer to [**[]Snssai**](Snssai.md) |  | [optional] 
**Locality** | Pointer to **string** |  | [optional] 
**NFInfo** | Pointer to [**NFInfo**](NFInfo.md) |  | [optional] 
**Capacity** | Pointer to **int32** |  | [optional] 
**NfSetIdList** | Pointer to **[]string** |  | [optional] 
**ServingScope** | Pointer to **[]string** |  | [optional] 
**LcHSupportInd** | Pointer to **bool** |  | [optional] 
**OlcHSupportInd** | Pointer to **bool** |  | [optional] 
**NfSetRecoveryTimeList** | Pointer to [**[]time.Time**](time.Time.md) |  | [optional] 
**ScpDomains** | Pointer to **[]string** |  | [optional] 
**VendorId** | Pointer to **string** |  | [optional] 

## Methods

### NewManagedNFProfile

`func NewManagedNFProfile() *ManagedNFProfile`

NewManagedNFProfile instantiates a new ManagedNFProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManagedNFProfileWithDefaults

`func NewManagedNFProfileWithDefaults() *ManagedNFProfile`

NewManagedNFProfileWithDefaults instantiates a new ManagedNFProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNfInstanceID

`func (o *ManagedNFProfile) GetNfInstanceID() string`

GetNfInstanceID returns the NfInstanceID field if non-nil, zero value otherwise.

### GetNfInstanceIDOk

`func (o *ManagedNFProfile) GetNfInstanceIDOk() (*string, bool)`

GetNfInstanceIDOk returns a tuple with the NfInstanceID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNfInstanceID

`func (o *ManagedNFProfile) SetNfInstanceID(v string)`

SetNfInstanceID sets NfInstanceID field to given value.

### HasNfInstanceID

`func (o *ManagedNFProfile) HasNfInstanceID() bool`

HasNfInstanceID returns a boolean if a field has been set.

### GetNfType

`func (o *ManagedNFProfile) GetNfType() NFType`

GetNfType returns the NfType field if non-nil, zero value otherwise.

### GetNfTypeOk

`func (o *ManagedNFProfile) GetNfTypeOk() (*NFType, bool)`

GetNfTypeOk returns a tuple with the NfType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNfType

`func (o *ManagedNFProfile) SetNfType(v NFType)`

SetNfType sets NfType field to given value.

### HasNfType

`func (o *ManagedNFProfile) HasNfType() bool`

HasNfType returns a boolean if a field has been set.

### GetHeartbeatTimer

`func (o *ManagedNFProfile) GetHeartbeatTimer() int32`

GetHeartbeatTimer returns the HeartbeatTimer field if non-nil, zero value otherwise.

### GetHeartbeatTimerOk

`func (o *ManagedNFProfile) GetHeartbeatTimerOk() (*int32, bool)`

GetHeartbeatTimerOk returns a tuple with the HeartbeatTimer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeartbeatTimer

`func (o *ManagedNFProfile) SetHeartbeatTimer(v int32)`

SetHeartbeatTimer sets HeartbeatTimer field to given value.

### HasHeartbeatTimer

`func (o *ManagedNFProfile) HasHeartbeatTimer() bool`

HasHeartbeatTimer returns a boolean if a field has been set.

### GetAuthzInfo

`func (o *ManagedNFProfile) GetAuthzInfo() string`

GetAuthzInfo returns the AuthzInfo field if non-nil, zero value otherwise.

### GetAuthzInfoOk

`func (o *ManagedNFProfile) GetAuthzInfoOk() (*string, bool)`

GetAuthzInfoOk returns a tuple with the AuthzInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthzInfo

`func (o *ManagedNFProfile) SetAuthzInfo(v string)`

SetAuthzInfo sets AuthzInfo field to given value.

### HasAuthzInfo

`func (o *ManagedNFProfile) HasAuthzInfo() bool`

HasAuthzInfo returns a boolean if a field has been set.

### GetHostAddr

`func (o *ManagedNFProfile) GetHostAddr() HostAddr`

GetHostAddr returns the HostAddr field if non-nil, zero value otherwise.

### GetHostAddrOk

`func (o *ManagedNFProfile) GetHostAddrOk() (*HostAddr, bool)`

GetHostAddrOk returns a tuple with the HostAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostAddr

`func (o *ManagedNFProfile) SetHostAddr(v HostAddr)`

SetHostAddr sets HostAddr field to given value.

### HasHostAddr

`func (o *ManagedNFProfile) HasHostAddr() bool`

HasHostAddr returns a boolean if a field has been set.

### GetAllowedPLMNs

`func (o *ManagedNFProfile) GetAllowedPLMNs() []PlmnId`

GetAllowedPLMNs returns the AllowedPLMNs field if non-nil, zero value otherwise.

### GetAllowedPLMNsOk

`func (o *ManagedNFProfile) GetAllowedPLMNsOk() (*[]PlmnId, bool)`

GetAllowedPLMNsOk returns a tuple with the AllowedPLMNs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedPLMNs

`func (o *ManagedNFProfile) SetAllowedPLMNs(v []PlmnId)`

SetAllowedPLMNs sets AllowedPLMNs field to given value.

### HasAllowedPLMNs

`func (o *ManagedNFProfile) HasAllowedPLMNs() bool`

HasAllowedPLMNs returns a boolean if a field has been set.

### GetAllowedSNPNs

`func (o *ManagedNFProfile) GetAllowedSNPNs() []SnpnInfo`

GetAllowedSNPNs returns the AllowedSNPNs field if non-nil, zero value otherwise.

### GetAllowedSNPNsOk

`func (o *ManagedNFProfile) GetAllowedSNPNsOk() (*[]SnpnInfo, bool)`

GetAllowedSNPNsOk returns a tuple with the AllowedSNPNs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedSNPNs

`func (o *ManagedNFProfile) SetAllowedSNPNs(v []SnpnInfo)`

SetAllowedSNPNs sets AllowedSNPNs field to given value.

### HasAllowedSNPNs

`func (o *ManagedNFProfile) HasAllowedSNPNs() bool`

HasAllowedSNPNs returns a boolean if a field has been set.

### GetAllowedNfTypes

`func (o *ManagedNFProfile) GetAllowedNfTypes() []NFType`

GetAllowedNfTypes returns the AllowedNfTypes field if non-nil, zero value otherwise.

### GetAllowedNfTypesOk

`func (o *ManagedNFProfile) GetAllowedNfTypesOk() (*[]NFType, bool)`

GetAllowedNfTypesOk returns a tuple with the AllowedNfTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedNfTypes

`func (o *ManagedNFProfile) SetAllowedNfTypes(v []NFType)`

SetAllowedNfTypes sets AllowedNfTypes field to given value.

### HasAllowedNfTypes

`func (o *ManagedNFProfile) HasAllowedNfTypes() bool`

HasAllowedNfTypes returns a boolean if a field has been set.

### GetAllowedNfDomains

`func (o *ManagedNFProfile) GetAllowedNfDomains() []string`

GetAllowedNfDomains returns the AllowedNfDomains field if non-nil, zero value otherwise.

### GetAllowedNfDomainsOk

`func (o *ManagedNFProfile) GetAllowedNfDomainsOk() (*[]string, bool)`

GetAllowedNfDomainsOk returns a tuple with the AllowedNfDomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedNfDomains

`func (o *ManagedNFProfile) SetAllowedNfDomains(v []string)`

SetAllowedNfDomains sets AllowedNfDomains field to given value.

### HasAllowedNfDomains

`func (o *ManagedNFProfile) HasAllowedNfDomains() bool`

HasAllowedNfDomains returns a boolean if a field has been set.

### GetAllowedNSSAIs

`func (o *ManagedNFProfile) GetAllowedNSSAIs() []Snssai`

GetAllowedNSSAIs returns the AllowedNSSAIs field if non-nil, zero value otherwise.

### GetAllowedNSSAIsOk

`func (o *ManagedNFProfile) GetAllowedNSSAIsOk() (*[]Snssai, bool)`

GetAllowedNSSAIsOk returns a tuple with the AllowedNSSAIs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedNSSAIs

`func (o *ManagedNFProfile) SetAllowedNSSAIs(v []Snssai)`

SetAllowedNSSAIs sets AllowedNSSAIs field to given value.

### HasAllowedNSSAIs

`func (o *ManagedNFProfile) HasAllowedNSSAIs() bool`

HasAllowedNSSAIs returns a boolean if a field has been set.

### GetLocality

`func (o *ManagedNFProfile) GetLocality() string`

GetLocality returns the Locality field if non-nil, zero value otherwise.

### GetLocalityOk

`func (o *ManagedNFProfile) GetLocalityOk() (*string, bool)`

GetLocalityOk returns a tuple with the Locality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocality

`func (o *ManagedNFProfile) SetLocality(v string)`

SetLocality sets Locality field to given value.

### HasLocality

`func (o *ManagedNFProfile) HasLocality() bool`

HasLocality returns a boolean if a field has been set.

### GetNFInfo

`func (o *ManagedNFProfile) GetNFInfo() NFInfo`

GetNFInfo returns the NFInfo field if non-nil, zero value otherwise.

### GetNFInfoOk

`func (o *ManagedNFProfile) GetNFInfoOk() (*NFInfo, bool)`

GetNFInfoOk returns a tuple with the NFInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNFInfo

`func (o *ManagedNFProfile) SetNFInfo(v NFInfo)`

SetNFInfo sets NFInfo field to given value.

### HasNFInfo

`func (o *ManagedNFProfile) HasNFInfo() bool`

HasNFInfo returns a boolean if a field has been set.

### GetCapacity

`func (o *ManagedNFProfile) GetCapacity() int32`

GetCapacity returns the Capacity field if non-nil, zero value otherwise.

### GetCapacityOk

`func (o *ManagedNFProfile) GetCapacityOk() (*int32, bool)`

GetCapacityOk returns a tuple with the Capacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacity

`func (o *ManagedNFProfile) SetCapacity(v int32)`

SetCapacity sets Capacity field to given value.

### HasCapacity

`func (o *ManagedNFProfile) HasCapacity() bool`

HasCapacity returns a boolean if a field has been set.

### GetNfSetIdList

`func (o *ManagedNFProfile) GetNfSetIdList() []string`

GetNfSetIdList returns the NfSetIdList field if non-nil, zero value otherwise.

### GetNfSetIdListOk

`func (o *ManagedNFProfile) GetNfSetIdListOk() (*[]string, bool)`

GetNfSetIdListOk returns a tuple with the NfSetIdList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNfSetIdList

`func (o *ManagedNFProfile) SetNfSetIdList(v []string)`

SetNfSetIdList sets NfSetIdList field to given value.

### HasNfSetIdList

`func (o *ManagedNFProfile) HasNfSetIdList() bool`

HasNfSetIdList returns a boolean if a field has been set.

### GetServingScope

`func (o *ManagedNFProfile) GetServingScope() []string`

GetServingScope returns the ServingScope field if non-nil, zero value otherwise.

### GetServingScopeOk

`func (o *ManagedNFProfile) GetServingScopeOk() (*[]string, bool)`

GetServingScopeOk returns a tuple with the ServingScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServingScope

`func (o *ManagedNFProfile) SetServingScope(v []string)`

SetServingScope sets ServingScope field to given value.

### HasServingScope

`func (o *ManagedNFProfile) HasServingScope() bool`

HasServingScope returns a boolean if a field has been set.

### GetLcHSupportInd

`func (o *ManagedNFProfile) GetLcHSupportInd() bool`

GetLcHSupportInd returns the LcHSupportInd field if non-nil, zero value otherwise.

### GetLcHSupportIndOk

`func (o *ManagedNFProfile) GetLcHSupportIndOk() (*bool, bool)`

GetLcHSupportIndOk returns a tuple with the LcHSupportInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLcHSupportInd

`func (o *ManagedNFProfile) SetLcHSupportInd(v bool)`

SetLcHSupportInd sets LcHSupportInd field to given value.

### HasLcHSupportInd

`func (o *ManagedNFProfile) HasLcHSupportInd() bool`

HasLcHSupportInd returns a boolean if a field has been set.

### GetOlcHSupportInd

`func (o *ManagedNFProfile) GetOlcHSupportInd() bool`

GetOlcHSupportInd returns the OlcHSupportInd field if non-nil, zero value otherwise.

### GetOlcHSupportIndOk

`func (o *ManagedNFProfile) GetOlcHSupportIndOk() (*bool, bool)`

GetOlcHSupportIndOk returns a tuple with the OlcHSupportInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOlcHSupportInd

`func (o *ManagedNFProfile) SetOlcHSupportInd(v bool)`

SetOlcHSupportInd sets OlcHSupportInd field to given value.

### HasOlcHSupportInd

`func (o *ManagedNFProfile) HasOlcHSupportInd() bool`

HasOlcHSupportInd returns a boolean if a field has been set.

### GetNfSetRecoveryTimeList

`func (o *ManagedNFProfile) GetNfSetRecoveryTimeList() []time.Time`

GetNfSetRecoveryTimeList returns the NfSetRecoveryTimeList field if non-nil, zero value otherwise.

### GetNfSetRecoveryTimeListOk

`func (o *ManagedNFProfile) GetNfSetRecoveryTimeListOk() (*[]time.Time, bool)`

GetNfSetRecoveryTimeListOk returns a tuple with the NfSetRecoveryTimeList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNfSetRecoveryTimeList

`func (o *ManagedNFProfile) SetNfSetRecoveryTimeList(v []time.Time)`

SetNfSetRecoveryTimeList sets NfSetRecoveryTimeList field to given value.

### HasNfSetRecoveryTimeList

`func (o *ManagedNFProfile) HasNfSetRecoveryTimeList() bool`

HasNfSetRecoveryTimeList returns a boolean if a field has been set.

### GetScpDomains

`func (o *ManagedNFProfile) GetScpDomains() []string`

GetScpDomains returns the ScpDomains field if non-nil, zero value otherwise.

### GetScpDomainsOk

`func (o *ManagedNFProfile) GetScpDomainsOk() (*[]string, bool)`

GetScpDomainsOk returns a tuple with the ScpDomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScpDomains

`func (o *ManagedNFProfile) SetScpDomains(v []string)`

SetScpDomains sets ScpDomains field to given value.

### HasScpDomains

`func (o *ManagedNFProfile) HasScpDomains() bool`

HasScpDomains returns a boolean if a field has been set.

### GetVendorId

`func (o *ManagedNFProfile) GetVendorId() string`

GetVendorId returns the VendorId field if non-nil, zero value otherwise.

### GetVendorIdOk

`func (o *ManagedNFProfile) GetVendorIdOk() (*string, bool)`

GetVendorIdOk returns a tuple with the VendorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorId

`func (o *ManagedNFProfile) SetVendorId(v string)`

SetVendorId sets VendorId field to given value.

### HasVendorId

`func (o *ManagedNFProfile) HasVendorId() bool`

HasVendorId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


