# NFService

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServiceInstanceId** | **string** |  | 
**ServiceName** | [**ServiceName**](ServiceName.md) |  | 
**Versions** | [**[]NFServiceVersion**](NFServiceVersion.md) |  | 
**Scheme** | [**UriScheme**](UriScheme.md) |  | 
**NfServiceStatus** | [**NFServiceStatus**](NFServiceStatus.md) |  | 
**Fqdn** | Pointer to **string** | Fully Qualified Domain Name | [optional] 
**InterPlmnFqdn** | Pointer to **string** | Fully Qualified Domain Name | [optional] 
**IpEndPoints** | Pointer to [**[]IpEndPoint**](IpEndPoint.md) |  | [optional] 
**ApiPrefix** | Pointer to **string** |  | [optional] 
**DefaultNotificationSubscriptions** | Pointer to [**[]DefaultNotificationSubscription**](DefaultNotificationSubscription.md) |  | [optional] 
**AllowedPlmns** | Pointer to [**[]PlmnId**](PlmnId.md) |  | [optional] 
**AllowedSnpns** | Pointer to [**[]PlmnIdNid**](PlmnIdNid.md) |  | [optional] 
**AllowedNfTypes** | Pointer to [**[]NFType**](NFType.md) |  | [optional] 
**AllowedNfDomains** | Pointer to **[]string** |  | [optional] 
**AllowedNssais** | Pointer to [**[]ExtSnssai**](ExtSnssai.md) |  | [optional] 
**AllowedOperationsPerNfType** | Pointer to **map[string][]string** | A map (list of key-value pairs) where NF Type serves as key | [optional] 
**AllowedOperationsPerNfInstance** | Pointer to **map[string][]string** | A map (list of key-value pairs) where NF Instance Id serves as key | [optional] 
**AllowedOperationsPerNfInstanceOverrides** | Pointer to **bool** |  | [optional] [default to false]
**Priority** | Pointer to **int32** |  | [optional] 
**Capacity** | Pointer to **int32** |  | [optional] 
**Load** | Pointer to **int32** |  | [optional] 
**LoadTimeStamp** | Pointer to **time.Time** | string with format &#39;date-time&#39; as defined in OpenAPI. | [optional] 
**RecoveryTime** | Pointer to **time.Time** | string with format &#39;date-time&#39; as defined in OpenAPI. | [optional] 
**SupportedFeatures** | Pointer to **string** | A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \&quot;0\&quot; to \&quot;9\&quot;,  \&quot;a\&quot; to \&quot;f\&quot; or \&quot;A\&quot; to \&quot;F\&quot; and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.  | [optional] 
**NfServiceSetIdList** | Pointer to **[]string** |  | [optional] 
**SNssais** | Pointer to [**[]ExtSnssai**](ExtSnssai.md) |  | [optional] 
**PerPlmnSnssaiList** | Pointer to [**[]PlmnSnssai**](PlmnSnssai.md) |  | [optional] 
**VendorId** | Pointer to **string** | Vendor ID of the NF Service instance (Private Enterprise Number assigned by IANA) | [optional] 
**SupportedVendorSpecificFeatures** | Pointer to [**map[string][]VendorSpecificFeature**](array.md) | A map (list of key-value pairs) where IANA-assigned SMI Network Management Private Enterprise Codes serves as key  | [optional] 
**Oauth2Required** | Pointer to **bool** |  | [optional] 
**PerPlmnOauth2ReqList** | Pointer to [**PlmnOauth2**](PlmnOauth2.md) |  | [optional] 

## Methods

### NewNFService

`func NewNFService(serviceInstanceId string, serviceName ServiceName, versions []NFServiceVersion, scheme UriScheme, nfServiceStatus NFServiceStatus, ) *NFService`

NewNFService instantiates a new NFService object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNFServiceWithDefaults

`func NewNFServiceWithDefaults() *NFService`

NewNFServiceWithDefaults instantiates a new NFService object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServiceInstanceId

`func (o *NFService) GetServiceInstanceId() string`

GetServiceInstanceId returns the ServiceInstanceId field if non-nil, zero value otherwise.

### GetServiceInstanceIdOk

`func (o *NFService) GetServiceInstanceIdOk() (*string, bool)`

GetServiceInstanceIdOk returns a tuple with the ServiceInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceInstanceId

`func (o *NFService) SetServiceInstanceId(v string)`

SetServiceInstanceId sets ServiceInstanceId field to given value.


### GetServiceName

`func (o *NFService) GetServiceName() ServiceName`

GetServiceName returns the ServiceName field if non-nil, zero value otherwise.

### GetServiceNameOk

`func (o *NFService) GetServiceNameOk() (*ServiceName, bool)`

GetServiceNameOk returns a tuple with the ServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceName

`func (o *NFService) SetServiceName(v ServiceName)`

SetServiceName sets ServiceName field to given value.


### GetVersions

`func (o *NFService) GetVersions() []NFServiceVersion`

GetVersions returns the Versions field if non-nil, zero value otherwise.

### GetVersionsOk

`func (o *NFService) GetVersionsOk() (*[]NFServiceVersion, bool)`

GetVersionsOk returns a tuple with the Versions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersions

`func (o *NFService) SetVersions(v []NFServiceVersion)`

SetVersions sets Versions field to given value.


### GetScheme

`func (o *NFService) GetScheme() UriScheme`

GetScheme returns the Scheme field if non-nil, zero value otherwise.

### GetSchemeOk

`func (o *NFService) GetSchemeOk() (*UriScheme, bool)`

GetSchemeOk returns a tuple with the Scheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheme

`func (o *NFService) SetScheme(v UriScheme)`

SetScheme sets Scheme field to given value.


### GetNfServiceStatus

`func (o *NFService) GetNfServiceStatus() NFServiceStatus`

GetNfServiceStatus returns the NfServiceStatus field if non-nil, zero value otherwise.

### GetNfServiceStatusOk

`func (o *NFService) GetNfServiceStatusOk() (*NFServiceStatus, bool)`

GetNfServiceStatusOk returns a tuple with the NfServiceStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNfServiceStatus

`func (o *NFService) SetNfServiceStatus(v NFServiceStatus)`

SetNfServiceStatus sets NfServiceStatus field to given value.


### GetFqdn

`func (o *NFService) GetFqdn() string`

GetFqdn returns the Fqdn field if non-nil, zero value otherwise.

### GetFqdnOk

`func (o *NFService) GetFqdnOk() (*string, bool)`

GetFqdnOk returns a tuple with the Fqdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFqdn

`func (o *NFService) SetFqdn(v string)`

SetFqdn sets Fqdn field to given value.

### HasFqdn

`func (o *NFService) HasFqdn() bool`

HasFqdn returns a boolean if a field has been set.

### GetInterPlmnFqdn

`func (o *NFService) GetInterPlmnFqdn() string`

GetInterPlmnFqdn returns the InterPlmnFqdn field if non-nil, zero value otherwise.

### GetInterPlmnFqdnOk

`func (o *NFService) GetInterPlmnFqdnOk() (*string, bool)`

GetInterPlmnFqdnOk returns a tuple with the InterPlmnFqdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterPlmnFqdn

`func (o *NFService) SetInterPlmnFqdn(v string)`

SetInterPlmnFqdn sets InterPlmnFqdn field to given value.

### HasInterPlmnFqdn

`func (o *NFService) HasInterPlmnFqdn() bool`

HasInterPlmnFqdn returns a boolean if a field has been set.

### GetIpEndPoints

`func (o *NFService) GetIpEndPoints() []IpEndPoint`

GetIpEndPoints returns the IpEndPoints field if non-nil, zero value otherwise.

### GetIpEndPointsOk

`func (o *NFService) GetIpEndPointsOk() (*[]IpEndPoint, bool)`

GetIpEndPointsOk returns a tuple with the IpEndPoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpEndPoints

`func (o *NFService) SetIpEndPoints(v []IpEndPoint)`

SetIpEndPoints sets IpEndPoints field to given value.

### HasIpEndPoints

`func (o *NFService) HasIpEndPoints() bool`

HasIpEndPoints returns a boolean if a field has been set.

### GetApiPrefix

`func (o *NFService) GetApiPrefix() string`

GetApiPrefix returns the ApiPrefix field if non-nil, zero value otherwise.

### GetApiPrefixOk

`func (o *NFService) GetApiPrefixOk() (*string, bool)`

GetApiPrefixOk returns a tuple with the ApiPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiPrefix

`func (o *NFService) SetApiPrefix(v string)`

SetApiPrefix sets ApiPrefix field to given value.

### HasApiPrefix

`func (o *NFService) HasApiPrefix() bool`

HasApiPrefix returns a boolean if a field has been set.

### GetDefaultNotificationSubscriptions

`func (o *NFService) GetDefaultNotificationSubscriptions() []DefaultNotificationSubscription`

GetDefaultNotificationSubscriptions returns the DefaultNotificationSubscriptions field if non-nil, zero value otherwise.

### GetDefaultNotificationSubscriptionsOk

`func (o *NFService) GetDefaultNotificationSubscriptionsOk() (*[]DefaultNotificationSubscription, bool)`

GetDefaultNotificationSubscriptionsOk returns a tuple with the DefaultNotificationSubscriptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultNotificationSubscriptions

`func (o *NFService) SetDefaultNotificationSubscriptions(v []DefaultNotificationSubscription)`

SetDefaultNotificationSubscriptions sets DefaultNotificationSubscriptions field to given value.

### HasDefaultNotificationSubscriptions

`func (o *NFService) HasDefaultNotificationSubscriptions() bool`

HasDefaultNotificationSubscriptions returns a boolean if a field has been set.

### GetAllowedPlmns

`func (o *NFService) GetAllowedPlmns() []PlmnId`

GetAllowedPlmns returns the AllowedPlmns field if non-nil, zero value otherwise.

### GetAllowedPlmnsOk

`func (o *NFService) GetAllowedPlmnsOk() (*[]PlmnId, bool)`

GetAllowedPlmnsOk returns a tuple with the AllowedPlmns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedPlmns

`func (o *NFService) SetAllowedPlmns(v []PlmnId)`

SetAllowedPlmns sets AllowedPlmns field to given value.

### HasAllowedPlmns

`func (o *NFService) HasAllowedPlmns() bool`

HasAllowedPlmns returns a boolean if a field has been set.

### GetAllowedSnpns

`func (o *NFService) GetAllowedSnpns() []PlmnIdNid`

GetAllowedSnpns returns the AllowedSnpns field if non-nil, zero value otherwise.

### GetAllowedSnpnsOk

`func (o *NFService) GetAllowedSnpnsOk() (*[]PlmnIdNid, bool)`

GetAllowedSnpnsOk returns a tuple with the AllowedSnpns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedSnpns

`func (o *NFService) SetAllowedSnpns(v []PlmnIdNid)`

SetAllowedSnpns sets AllowedSnpns field to given value.

### HasAllowedSnpns

`func (o *NFService) HasAllowedSnpns() bool`

HasAllowedSnpns returns a boolean if a field has been set.

### GetAllowedNfTypes

`func (o *NFService) GetAllowedNfTypes() []NFType`

GetAllowedNfTypes returns the AllowedNfTypes field if non-nil, zero value otherwise.

### GetAllowedNfTypesOk

`func (o *NFService) GetAllowedNfTypesOk() (*[]NFType, bool)`

GetAllowedNfTypesOk returns a tuple with the AllowedNfTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedNfTypes

`func (o *NFService) SetAllowedNfTypes(v []NFType)`

SetAllowedNfTypes sets AllowedNfTypes field to given value.

### HasAllowedNfTypes

`func (o *NFService) HasAllowedNfTypes() bool`

HasAllowedNfTypes returns a boolean if a field has been set.

### GetAllowedNfDomains

`func (o *NFService) GetAllowedNfDomains() []string`

GetAllowedNfDomains returns the AllowedNfDomains field if non-nil, zero value otherwise.

### GetAllowedNfDomainsOk

`func (o *NFService) GetAllowedNfDomainsOk() (*[]string, bool)`

GetAllowedNfDomainsOk returns a tuple with the AllowedNfDomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedNfDomains

`func (o *NFService) SetAllowedNfDomains(v []string)`

SetAllowedNfDomains sets AllowedNfDomains field to given value.

### HasAllowedNfDomains

`func (o *NFService) HasAllowedNfDomains() bool`

HasAllowedNfDomains returns a boolean if a field has been set.

### GetAllowedNssais

`func (o *NFService) GetAllowedNssais() []ExtSnssai`

GetAllowedNssais returns the AllowedNssais field if non-nil, zero value otherwise.

### GetAllowedNssaisOk

`func (o *NFService) GetAllowedNssaisOk() (*[]ExtSnssai, bool)`

GetAllowedNssaisOk returns a tuple with the AllowedNssais field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedNssais

`func (o *NFService) SetAllowedNssais(v []ExtSnssai)`

SetAllowedNssais sets AllowedNssais field to given value.

### HasAllowedNssais

`func (o *NFService) HasAllowedNssais() bool`

HasAllowedNssais returns a boolean if a field has been set.

### GetAllowedOperationsPerNfType

`func (o *NFService) GetAllowedOperationsPerNfType() map[string][]string`

GetAllowedOperationsPerNfType returns the AllowedOperationsPerNfType field if non-nil, zero value otherwise.

### GetAllowedOperationsPerNfTypeOk

`func (o *NFService) GetAllowedOperationsPerNfTypeOk() (*map[string][]string, bool)`

GetAllowedOperationsPerNfTypeOk returns a tuple with the AllowedOperationsPerNfType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedOperationsPerNfType

`func (o *NFService) SetAllowedOperationsPerNfType(v map[string][]string)`

SetAllowedOperationsPerNfType sets AllowedOperationsPerNfType field to given value.

### HasAllowedOperationsPerNfType

`func (o *NFService) HasAllowedOperationsPerNfType() bool`

HasAllowedOperationsPerNfType returns a boolean if a field has been set.

### GetAllowedOperationsPerNfInstance

`func (o *NFService) GetAllowedOperationsPerNfInstance() map[string][]string`

GetAllowedOperationsPerNfInstance returns the AllowedOperationsPerNfInstance field if non-nil, zero value otherwise.

### GetAllowedOperationsPerNfInstanceOk

`func (o *NFService) GetAllowedOperationsPerNfInstanceOk() (*map[string][]string, bool)`

GetAllowedOperationsPerNfInstanceOk returns a tuple with the AllowedOperationsPerNfInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedOperationsPerNfInstance

`func (o *NFService) SetAllowedOperationsPerNfInstance(v map[string][]string)`

SetAllowedOperationsPerNfInstance sets AllowedOperationsPerNfInstance field to given value.

### HasAllowedOperationsPerNfInstance

`func (o *NFService) HasAllowedOperationsPerNfInstance() bool`

HasAllowedOperationsPerNfInstance returns a boolean if a field has been set.

### GetAllowedOperationsPerNfInstanceOverrides

`func (o *NFService) GetAllowedOperationsPerNfInstanceOverrides() bool`

GetAllowedOperationsPerNfInstanceOverrides returns the AllowedOperationsPerNfInstanceOverrides field if non-nil, zero value otherwise.

### GetAllowedOperationsPerNfInstanceOverridesOk

`func (o *NFService) GetAllowedOperationsPerNfInstanceOverridesOk() (*bool, bool)`

GetAllowedOperationsPerNfInstanceOverridesOk returns a tuple with the AllowedOperationsPerNfInstanceOverrides field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedOperationsPerNfInstanceOverrides

`func (o *NFService) SetAllowedOperationsPerNfInstanceOverrides(v bool)`

SetAllowedOperationsPerNfInstanceOverrides sets AllowedOperationsPerNfInstanceOverrides field to given value.

### HasAllowedOperationsPerNfInstanceOverrides

`func (o *NFService) HasAllowedOperationsPerNfInstanceOverrides() bool`

HasAllowedOperationsPerNfInstanceOverrides returns a boolean if a field has been set.

### GetPriority

`func (o *NFService) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *NFService) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *NFService) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *NFService) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetCapacity

`func (o *NFService) GetCapacity() int32`

GetCapacity returns the Capacity field if non-nil, zero value otherwise.

### GetCapacityOk

`func (o *NFService) GetCapacityOk() (*int32, bool)`

GetCapacityOk returns a tuple with the Capacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacity

`func (o *NFService) SetCapacity(v int32)`

SetCapacity sets Capacity field to given value.

### HasCapacity

`func (o *NFService) HasCapacity() bool`

HasCapacity returns a boolean if a field has been set.

### GetLoad

`func (o *NFService) GetLoad() int32`

GetLoad returns the Load field if non-nil, zero value otherwise.

### GetLoadOk

`func (o *NFService) GetLoadOk() (*int32, bool)`

GetLoadOk returns a tuple with the Load field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoad

`func (o *NFService) SetLoad(v int32)`

SetLoad sets Load field to given value.

### HasLoad

`func (o *NFService) HasLoad() bool`

HasLoad returns a boolean if a field has been set.

### GetLoadTimeStamp

`func (o *NFService) GetLoadTimeStamp() time.Time`

GetLoadTimeStamp returns the LoadTimeStamp field if non-nil, zero value otherwise.

### GetLoadTimeStampOk

`func (o *NFService) GetLoadTimeStampOk() (*time.Time, bool)`

GetLoadTimeStampOk returns a tuple with the LoadTimeStamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadTimeStamp

`func (o *NFService) SetLoadTimeStamp(v time.Time)`

SetLoadTimeStamp sets LoadTimeStamp field to given value.

### HasLoadTimeStamp

`func (o *NFService) HasLoadTimeStamp() bool`

HasLoadTimeStamp returns a boolean if a field has been set.

### GetRecoveryTime

`func (o *NFService) GetRecoveryTime() time.Time`

GetRecoveryTime returns the RecoveryTime field if non-nil, zero value otherwise.

### GetRecoveryTimeOk

`func (o *NFService) GetRecoveryTimeOk() (*time.Time, bool)`

GetRecoveryTimeOk returns a tuple with the RecoveryTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoveryTime

`func (o *NFService) SetRecoveryTime(v time.Time)`

SetRecoveryTime sets RecoveryTime field to given value.

### HasRecoveryTime

`func (o *NFService) HasRecoveryTime() bool`

HasRecoveryTime returns a boolean if a field has been set.

### GetSupportedFeatures

`func (o *NFService) GetSupportedFeatures() string`

GetSupportedFeatures returns the SupportedFeatures field if non-nil, zero value otherwise.

### GetSupportedFeaturesOk

`func (o *NFService) GetSupportedFeaturesOk() (*string, bool)`

GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedFeatures

`func (o *NFService) SetSupportedFeatures(v string)`

SetSupportedFeatures sets SupportedFeatures field to given value.

### HasSupportedFeatures

`func (o *NFService) HasSupportedFeatures() bool`

HasSupportedFeatures returns a boolean if a field has been set.

### GetNfServiceSetIdList

`func (o *NFService) GetNfServiceSetIdList() []string`

GetNfServiceSetIdList returns the NfServiceSetIdList field if non-nil, zero value otherwise.

### GetNfServiceSetIdListOk

`func (o *NFService) GetNfServiceSetIdListOk() (*[]string, bool)`

GetNfServiceSetIdListOk returns a tuple with the NfServiceSetIdList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNfServiceSetIdList

`func (o *NFService) SetNfServiceSetIdList(v []string)`

SetNfServiceSetIdList sets NfServiceSetIdList field to given value.

### HasNfServiceSetIdList

`func (o *NFService) HasNfServiceSetIdList() bool`

HasNfServiceSetIdList returns a boolean if a field has been set.

### GetSNssais

`func (o *NFService) GetSNssais() []ExtSnssai`

GetSNssais returns the SNssais field if non-nil, zero value otherwise.

### GetSNssaisOk

`func (o *NFService) GetSNssaisOk() (*[]ExtSnssai, bool)`

GetSNssaisOk returns a tuple with the SNssais field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSNssais

`func (o *NFService) SetSNssais(v []ExtSnssai)`

SetSNssais sets SNssais field to given value.

### HasSNssais

`func (o *NFService) HasSNssais() bool`

HasSNssais returns a boolean if a field has been set.

### GetPerPlmnSnssaiList

`func (o *NFService) GetPerPlmnSnssaiList() []PlmnSnssai`

GetPerPlmnSnssaiList returns the PerPlmnSnssaiList field if non-nil, zero value otherwise.

### GetPerPlmnSnssaiListOk

`func (o *NFService) GetPerPlmnSnssaiListOk() (*[]PlmnSnssai, bool)`

GetPerPlmnSnssaiListOk returns a tuple with the PerPlmnSnssaiList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerPlmnSnssaiList

`func (o *NFService) SetPerPlmnSnssaiList(v []PlmnSnssai)`

SetPerPlmnSnssaiList sets PerPlmnSnssaiList field to given value.

### HasPerPlmnSnssaiList

`func (o *NFService) HasPerPlmnSnssaiList() bool`

HasPerPlmnSnssaiList returns a boolean if a field has been set.

### GetVendorId

`func (o *NFService) GetVendorId() string`

GetVendorId returns the VendorId field if non-nil, zero value otherwise.

### GetVendorIdOk

`func (o *NFService) GetVendorIdOk() (*string, bool)`

GetVendorIdOk returns a tuple with the VendorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorId

`func (o *NFService) SetVendorId(v string)`

SetVendorId sets VendorId field to given value.

### HasVendorId

`func (o *NFService) HasVendorId() bool`

HasVendorId returns a boolean if a field has been set.

### GetSupportedVendorSpecificFeatures

`func (o *NFService) GetSupportedVendorSpecificFeatures() map[string][]VendorSpecificFeature`

GetSupportedVendorSpecificFeatures returns the SupportedVendorSpecificFeatures field if non-nil, zero value otherwise.

### GetSupportedVendorSpecificFeaturesOk

`func (o *NFService) GetSupportedVendorSpecificFeaturesOk() (*map[string][]VendorSpecificFeature, bool)`

GetSupportedVendorSpecificFeaturesOk returns a tuple with the SupportedVendorSpecificFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedVendorSpecificFeatures

`func (o *NFService) SetSupportedVendorSpecificFeatures(v map[string][]VendorSpecificFeature)`

SetSupportedVendorSpecificFeatures sets SupportedVendorSpecificFeatures field to given value.

### HasSupportedVendorSpecificFeatures

`func (o *NFService) HasSupportedVendorSpecificFeatures() bool`

HasSupportedVendorSpecificFeatures returns a boolean if a field has been set.

### GetOauth2Required

`func (o *NFService) GetOauth2Required() bool`

GetOauth2Required returns the Oauth2Required field if non-nil, zero value otherwise.

### GetOauth2RequiredOk

`func (o *NFService) GetOauth2RequiredOk() (*bool, bool)`

GetOauth2RequiredOk returns a tuple with the Oauth2Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauth2Required

`func (o *NFService) SetOauth2Required(v bool)`

SetOauth2Required sets Oauth2Required field to given value.

### HasOauth2Required

`func (o *NFService) HasOauth2Required() bool`

HasOauth2Required returns a boolean if a field has been set.

### GetPerPlmnOauth2ReqList

`func (o *NFService) GetPerPlmnOauth2ReqList() PlmnOauth2`

GetPerPlmnOauth2ReqList returns the PerPlmnOauth2ReqList field if non-nil, zero value otherwise.

### GetPerPlmnOauth2ReqListOk

`func (o *NFService) GetPerPlmnOauth2ReqListOk() (*PlmnOauth2, bool)`

GetPerPlmnOauth2ReqListOk returns a tuple with the PerPlmnOauth2ReqList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerPlmnOauth2ReqList

`func (o *NFService) SetPerPlmnOauth2ReqList(v PlmnOauth2)`

SetPerPlmnOauth2ReqList sets PerPlmnOauth2ReqList field to given value.

### HasPerPlmnOauth2ReqList

`func (o *NFService) HasPerPlmnOauth2ReqList() bool`

HasPerPlmnOauth2ReqList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


