# NFProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NfInstanceId** | **string** | String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.   | 
**NfInstanceName** | Pointer to **string** |  | [optional] 
**NfType** | [**NFType**](NFType.md) |  | 
**NfStatus** | [**NfStatus**](NfStatus.md) |  | 
**CollocatedNfInstances** | Pointer to [**[]CollocatedNfInstance**](CollocatedNfInstance.md) |  | [optional] 
**HeartBeatTimer** | Pointer to **int32** |  | [optional] 
**PlmnList** | Pointer to [**[]PlmnId**](PlmnId.md) |  | [optional] 
**SnpnList** | Pointer to [**[]PlmnIdNid**](PlmnIdNid.md) |  | [optional] 
**SNssais** | Pointer to [**[]ExtSnssai**](ExtSnssai.md) |  | [optional] 
**PerPlmnSnssaiList** | Pointer to [**[]PlmnSnssai**](PlmnSnssai.md) |  | [optional] 
**NsiList** | Pointer to **[]string** |  | [optional] 
**Fqdn** | Pointer to **string** | Fully Qualified Domain Name | [optional] 
**InterPlmnFqdn** | Pointer to **string** | Fully Qualified Domain Name | [optional] 
**Ipv4Addresses** | Pointer to **[]string** |  | [optional] 
**Ipv6Addresses** | Pointer to [**[]Ipv6Addr**](Ipv6Addr.md) |  | [optional] 
**AllowedPlmns** | Pointer to [**[]PlmnId**](PlmnId.md) |  | [optional] 
**AllowedSnpns** | Pointer to [**[]PlmnIdNid**](PlmnIdNid.md) |  | [optional] 
**AllowedNfTypes** | Pointer to [**[]NFType**](NFType.md) |  | [optional] 
**AllowedNfDomains** | Pointer to **[]string** |  | [optional] 
**AllowedNssais** | Pointer to [**[]ExtSnssai**](ExtSnssai.md) |  | [optional] 
**Priority** | Pointer to **int32** |  | [optional] 
**Capacity** | Pointer to **int32** |  | [optional] 
**Load** | Pointer to **int32** |  | [optional] 
**LoadTimeStamp** | Pointer to **time.Time** | string with format &#39;date-time&#39; as defined in OpenAPI. | [optional] 
**Locality** | Pointer to **string** |  | [optional] 
**ExtLocality** | Pointer to **map[string]string** | A map (list of key-value pairs) where a (unique) valid JSON string serves as key representing a type of locality  | [optional] 
**UdrInfo** | Pointer to [**UdrInfo**](UdrInfo.md) |  | [optional] 
**UdrInfoList** | Pointer to [**map[string]UdrInfo**](UdrInfo.md) | A map (list of key-value pairs) where a (unique) valid JSON string serves as key of UdrInfo  | [optional] 
**UdmInfo** | Pointer to [**UdmInfo**](UdmInfo.md) |  | [optional] 
**UdmInfoList** | Pointer to [**map[string]UdmInfo**](UdmInfo.md) | A map (list of key-value pairs) where a (unique) valid JSON string serves as key of UdmInfo  | [optional] 
**AusfInfo** | Pointer to [**AusfInfo**](AusfInfo.md) |  | [optional] 
**AusfInfoList** | Pointer to [**map[string]AusfInfo**](AusfInfo.md) | A map (list of key-value pairs) where a (unique) valid JSON string serves as key of AusfInfo  | [optional] 
**AmfInfo** | Pointer to [**AmfInfo**](AmfInfo.md) |  | [optional] 
**AmfInfoList** | Pointer to [**map[string]AmfInfo**](AmfInfo.md) | A map (list of key-value pairs) where a (unique) valid JSON string serves as key of AmfInfo  | [optional] 
**SmfInfo** | Pointer to [**SmfInfo**](SmfInfo.md) |  | [optional] 
**SmfInfoList** | Pointer to [**map[string]SmfInfo**](SmfInfo.md) | A map (list of key-value pairs) where a (unique) valid JSON string serves as key of SmfInfo  | [optional] 
**UpfInfo** | Pointer to [**UpfInfo**](UpfInfo.md) |  | [optional] 
**UpfInfoList** | Pointer to [**map[string]UpfInfo**](UpfInfo.md) | A map (list of key-value pairs) where a (unique) valid JSON string serves as key of UpfInfo  | [optional] 
**PcfInfo** | Pointer to [**PcfInfo**](PcfInfo.md) |  | [optional] 
**PcfInfoList** | Pointer to [**map[string]PcfInfo**](PcfInfo.md) | A map (list of key-value pairs) where a (unique) valid JSON string serves as key of PcfInfo  | [optional] 
**BsfInfo** | Pointer to [**BsfInfo**](BsfInfo.md) |  | [optional] 
**BsfInfoList** | Pointer to [**map[string]BsfInfo**](BsfInfo.md) | A map (list of key-value pairs) where a (unique) valid JSON string serves as key of BsfInfo  | [optional] 
**ChfInfo** | Pointer to [**ChfInfo**](ChfInfo.md) |  | [optional] 
**ChfInfoList** | Pointer to [**map[string]ChfInfo**](ChfInfo.md) | A map (list of key-value pairs) where a (unique) valid JSON string serves as key of ChfInfo  | [optional] 
**NefInfo** | Pointer to [**NefInfo**](NefInfo.md) |  | [optional] 
**NrfInfo** | Pointer to [**NrfInfo**](NrfInfo.md) |  | [optional] 
**UdsfInfo** | Pointer to [**UdsfInfo**](UdsfInfo.md) |  | [optional] 
**UdsfInfoList** | Pointer to [**map[string]UdsfInfo**](UdsfInfo.md) | A map (list of key-value pairs) where a (unique) valid JSON string serves as key of UdsfInfo  | [optional] 
**NwdafInfo** | Pointer to [**NwdafInfo**](NwdafInfo.md) |  | [optional] 
**NwdafInfoList** | Pointer to [**map[string]NwdafInfo**](NwdafInfo.md) | A map (list of key-value pairs) where a (unique) valid JSON string serves as key of NwdafInfo  | [optional] 
**PcscfInfoList** | Pointer to [**map[string]PcscfInfo**](PcscfInfo.md) | A map (list of key-value pairs) where a (unique) valid JSON string serves as key of PcscfInfo  | [optional] 
**HssInfoList** | Pointer to [**map[string]HssInfo**](HssInfo.md) | A map (list of key-value pairs) where a (unique) valid JSON string serves as key of HssInfo  | [optional] 
**CustomInfo** | Pointer to **map[string]interface{}** |  | [optional] 
**RecoveryTime** | Pointer to **time.Time** | string with format &#39;date-time&#39; as defined in OpenAPI. | [optional] 
**NfServicePersistence** | Pointer to **bool** |  | [optional] [default to false]
**NfServices** | Pointer to [**[]NFService**](NFService.md) |  | [optional] 
**NfServiceList** | Pointer to [**map[string]NFService**](NFService.md) | A map (list of key-value pairs) where serviceInstanceId serves as key of NFService  | [optional] 
**NfProfileChangesSupportInd** | Pointer to **bool** |  | [optional] [default to false]
**NfProfileChangesInd** | Pointer to **bool** |  | [optional] [readonly] [default to false]
**DefaultNotificationSubscriptions** | Pointer to [**[]DefaultNotificationSubscription**](DefaultNotificationSubscription.md) |  | [optional] 
**LmfInfo** | Pointer to [**LmfInfo**](LmfInfo.md) |  | [optional] 
**GmlcInfo** | Pointer to [**GmlcInfo**](GmlcInfo.md) |  | [optional] 
**NfSetIdList** | Pointer to **[]string** |  | [optional] 
**ServingScope** | Pointer to **[]string** |  | [optional] 
**LcHSupportInd** | Pointer to **bool** |  | [optional] [default to false]
**OlcHSupportInd** | Pointer to **bool** |  | [optional] [default to false]
**NfSetRecoveryTimeList** | Pointer to [**map[string]time.Time**](time.Time.md) | A map (list of key-value pairs) where NfSetId serves as key of DateTime | [optional] 
**ServiceSetRecoveryTimeList** | Pointer to [**map[string]time.Time**](time.Time.md) | A map (list of key-value pairs) where NfServiceSetId serves as key of DateTime  | [optional] 
**ScpDomains** | Pointer to **[]string** |  | [optional] 
**ScpInfo** | Pointer to [**ScpInfo**](ScpInfo.md) |  | [optional] 
**SeppInfo** | Pointer to [**SeppInfo**](SeppInfo.md) |  | [optional] 
**VendorId** | Pointer to **string** | Vendor ID of the NF Service instance (Private Enterprise Number assigned by IANA) | [optional] 
**SupportedVendorSpecificFeatures** | Pointer to [**map[string][]VendorSpecificFeature**](array.md) | The key of the map is the IANA-assigned SMI Network Management Private Enterprise Codes  | [optional] 
**AanfInfoList** | Pointer to [**map[string]AanfInfo**](AanfInfo.md) | A map (list of key-value pairs) where a (unique) valid JSON string serves as key of AanfInfo  | [optional] 
**Var5gDdnmfInfo** | Pointer to [**Model5GDdnmfInfo**](Model5GDdnmfInfo.md) |  | [optional] 
**MfafInfo** | Pointer to [**MfafInfo**](MfafInfo.md) |  | [optional] 
**EasdfInfoList** | Pointer to [**map[string]EasdfInfo**](EasdfInfo.md) | A map (list of key-value pairs) where a (unique) valid JSON string serves as key of EasdfInfo  | [optional] 
**DccfInfo** | Pointer to [**DccfInfo**](DccfInfo.md) |  | [optional] 
**NsacfInfoList** | Pointer to [**map[string]NsacfInfo**](NsacfInfo.md) | A map (list of key-value pairs) where a (unique) valid JSON string serves as key of NsacfInfo  | [optional] 
**MbSmfInfoList** | Pointer to [**map[string]MbSmfInfo**](MbSmfInfo.md) | A map (list of key-value pairs) where a (unique) valid JSON string serves as key of MbSmfInfo  | [optional] 
**TsctsfInfoList** | Pointer to [**map[string]TsctsfInfo**](TsctsfInfo.md) | A map (list of key-value pairs) where a (unique) valid JSON string serves as key of TsctsfInfo  | [optional] 
**MbUpfInfoList** | Pointer to [**map[string]MbUpfInfo**](MbUpfInfo.md) | A map (list of key-value pairs) where a (unique) valid JSON string serves as key of MbUpfInfo  | [optional] 
**TrustAfInfo** | Pointer to [**TrustAfInfo**](TrustAfInfo.md) |  | [optional] 
**NssaafInfo** | Pointer to [**NssaafInfo**](NssaafInfo.md) |  | [optional] 
**HniList** | Pointer to **[]string** |  | [optional] 
**IwmscInfo** | Pointer to [**IwmscInfo**](IwmscInfo.md) |  | [optional] 
**MnpfInfo** | Pointer to [**MnpfInfo**](MnpfInfo.md) |  | [optional] 

## Methods

### NewNFProfile

`func NewNFProfile(nfInstanceId string, nfType NFType, nfStatus NfStatus, ) *NFProfile`

NewNFProfile instantiates a new NFProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNFProfileWithDefaults

`func NewNFProfileWithDefaults() *NFProfile`

NewNFProfileWithDefaults instantiates a new NFProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNfInstanceId

`func (o *NFProfile) GetNfInstanceId() string`

GetNfInstanceId returns the NfInstanceId field if non-nil, zero value otherwise.

### GetNfInstanceIdOk

`func (o *NFProfile) GetNfInstanceIdOk() (*string, bool)`

GetNfInstanceIdOk returns a tuple with the NfInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNfInstanceId

`func (o *NFProfile) SetNfInstanceId(v string)`

SetNfInstanceId sets NfInstanceId field to given value.


### GetNfInstanceName

`func (o *NFProfile) GetNfInstanceName() string`

GetNfInstanceName returns the NfInstanceName field if non-nil, zero value otherwise.

### GetNfInstanceNameOk

`func (o *NFProfile) GetNfInstanceNameOk() (*string, bool)`

GetNfInstanceNameOk returns a tuple with the NfInstanceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNfInstanceName

`func (o *NFProfile) SetNfInstanceName(v string)`

SetNfInstanceName sets NfInstanceName field to given value.

### HasNfInstanceName

`func (o *NFProfile) HasNfInstanceName() bool`

HasNfInstanceName returns a boolean if a field has been set.

### GetNfType

`func (o *NFProfile) GetNfType() NFType`

GetNfType returns the NfType field if non-nil, zero value otherwise.

### GetNfTypeOk

`func (o *NFProfile) GetNfTypeOk() (*NFType, bool)`

GetNfTypeOk returns a tuple with the NfType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNfType

`func (o *NFProfile) SetNfType(v NFType)`

SetNfType sets NfType field to given value.


### GetNfStatus

`func (o *NFProfile) GetNfStatus() NfStatus`

GetNfStatus returns the NfStatus field if non-nil, zero value otherwise.

### GetNfStatusOk

`func (o *NFProfile) GetNfStatusOk() (*NfStatus, bool)`

GetNfStatusOk returns a tuple with the NfStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNfStatus

`func (o *NFProfile) SetNfStatus(v NfStatus)`

SetNfStatus sets NfStatus field to given value.


### GetCollocatedNfInstances

`func (o *NFProfile) GetCollocatedNfInstances() []CollocatedNfInstance`

GetCollocatedNfInstances returns the CollocatedNfInstances field if non-nil, zero value otherwise.

### GetCollocatedNfInstancesOk

`func (o *NFProfile) GetCollocatedNfInstancesOk() (*[]CollocatedNfInstance, bool)`

GetCollocatedNfInstancesOk returns a tuple with the CollocatedNfInstances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollocatedNfInstances

`func (o *NFProfile) SetCollocatedNfInstances(v []CollocatedNfInstance)`

SetCollocatedNfInstances sets CollocatedNfInstances field to given value.

### HasCollocatedNfInstances

`func (o *NFProfile) HasCollocatedNfInstances() bool`

HasCollocatedNfInstances returns a boolean if a field has been set.

### GetHeartBeatTimer

`func (o *NFProfile) GetHeartBeatTimer() int32`

GetHeartBeatTimer returns the HeartBeatTimer field if non-nil, zero value otherwise.

### GetHeartBeatTimerOk

`func (o *NFProfile) GetHeartBeatTimerOk() (*int32, bool)`

GetHeartBeatTimerOk returns a tuple with the HeartBeatTimer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeartBeatTimer

`func (o *NFProfile) SetHeartBeatTimer(v int32)`

SetHeartBeatTimer sets HeartBeatTimer field to given value.

### HasHeartBeatTimer

`func (o *NFProfile) HasHeartBeatTimer() bool`

HasHeartBeatTimer returns a boolean if a field has been set.

### GetPlmnList

`func (o *NFProfile) GetPlmnList() []PlmnId`

GetPlmnList returns the PlmnList field if non-nil, zero value otherwise.

### GetPlmnListOk

`func (o *NFProfile) GetPlmnListOk() (*[]PlmnId, bool)`

GetPlmnListOk returns a tuple with the PlmnList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlmnList

`func (o *NFProfile) SetPlmnList(v []PlmnId)`

SetPlmnList sets PlmnList field to given value.

### HasPlmnList

`func (o *NFProfile) HasPlmnList() bool`

HasPlmnList returns a boolean if a field has been set.

### GetSnpnList

`func (o *NFProfile) GetSnpnList() []PlmnIdNid`

GetSnpnList returns the SnpnList field if non-nil, zero value otherwise.

### GetSnpnListOk

`func (o *NFProfile) GetSnpnListOk() (*[]PlmnIdNid, bool)`

GetSnpnListOk returns a tuple with the SnpnList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnpnList

`func (o *NFProfile) SetSnpnList(v []PlmnIdNid)`

SetSnpnList sets SnpnList field to given value.

### HasSnpnList

`func (o *NFProfile) HasSnpnList() bool`

HasSnpnList returns a boolean if a field has been set.

### GetSNssais

`func (o *NFProfile) GetSNssais() []ExtSnssai`

GetSNssais returns the SNssais field if non-nil, zero value otherwise.

### GetSNssaisOk

`func (o *NFProfile) GetSNssaisOk() (*[]ExtSnssai, bool)`

GetSNssaisOk returns a tuple with the SNssais field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSNssais

`func (o *NFProfile) SetSNssais(v []ExtSnssai)`

SetSNssais sets SNssais field to given value.

### HasSNssais

`func (o *NFProfile) HasSNssais() bool`

HasSNssais returns a boolean if a field has been set.

### GetPerPlmnSnssaiList

`func (o *NFProfile) GetPerPlmnSnssaiList() []PlmnSnssai`

GetPerPlmnSnssaiList returns the PerPlmnSnssaiList field if non-nil, zero value otherwise.

### GetPerPlmnSnssaiListOk

`func (o *NFProfile) GetPerPlmnSnssaiListOk() (*[]PlmnSnssai, bool)`

GetPerPlmnSnssaiListOk returns a tuple with the PerPlmnSnssaiList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerPlmnSnssaiList

`func (o *NFProfile) SetPerPlmnSnssaiList(v []PlmnSnssai)`

SetPerPlmnSnssaiList sets PerPlmnSnssaiList field to given value.

### HasPerPlmnSnssaiList

`func (o *NFProfile) HasPerPlmnSnssaiList() bool`

HasPerPlmnSnssaiList returns a boolean if a field has been set.

### GetNsiList

`func (o *NFProfile) GetNsiList() []string`

GetNsiList returns the NsiList field if non-nil, zero value otherwise.

### GetNsiListOk

`func (o *NFProfile) GetNsiListOk() (*[]string, bool)`

GetNsiListOk returns a tuple with the NsiList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNsiList

`func (o *NFProfile) SetNsiList(v []string)`

SetNsiList sets NsiList field to given value.

### HasNsiList

`func (o *NFProfile) HasNsiList() bool`

HasNsiList returns a boolean if a field has been set.

### GetFqdn

`func (o *NFProfile) GetFqdn() string`

GetFqdn returns the Fqdn field if non-nil, zero value otherwise.

### GetFqdnOk

`func (o *NFProfile) GetFqdnOk() (*string, bool)`

GetFqdnOk returns a tuple with the Fqdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFqdn

`func (o *NFProfile) SetFqdn(v string)`

SetFqdn sets Fqdn field to given value.

### HasFqdn

`func (o *NFProfile) HasFqdn() bool`

HasFqdn returns a boolean if a field has been set.

### GetInterPlmnFqdn

`func (o *NFProfile) GetInterPlmnFqdn() string`

GetInterPlmnFqdn returns the InterPlmnFqdn field if non-nil, zero value otherwise.

### GetInterPlmnFqdnOk

`func (o *NFProfile) GetInterPlmnFqdnOk() (*string, bool)`

GetInterPlmnFqdnOk returns a tuple with the InterPlmnFqdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterPlmnFqdn

`func (o *NFProfile) SetInterPlmnFqdn(v string)`

SetInterPlmnFqdn sets InterPlmnFqdn field to given value.

### HasInterPlmnFqdn

`func (o *NFProfile) HasInterPlmnFqdn() bool`

HasInterPlmnFqdn returns a boolean if a field has been set.

### GetIpv4Addresses

`func (o *NFProfile) GetIpv4Addresses() []string`

GetIpv4Addresses returns the Ipv4Addresses field if non-nil, zero value otherwise.

### GetIpv4AddressesOk

`func (o *NFProfile) GetIpv4AddressesOk() (*[]string, bool)`

GetIpv4AddressesOk returns a tuple with the Ipv4Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4Addresses

`func (o *NFProfile) SetIpv4Addresses(v []string)`

SetIpv4Addresses sets Ipv4Addresses field to given value.

### HasIpv4Addresses

`func (o *NFProfile) HasIpv4Addresses() bool`

HasIpv4Addresses returns a boolean if a field has been set.

### GetIpv6Addresses

`func (o *NFProfile) GetIpv6Addresses() []Ipv6Addr`

GetIpv6Addresses returns the Ipv6Addresses field if non-nil, zero value otherwise.

### GetIpv6AddressesOk

`func (o *NFProfile) GetIpv6AddressesOk() (*[]Ipv6Addr, bool)`

GetIpv6AddressesOk returns a tuple with the Ipv6Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6Addresses

`func (o *NFProfile) SetIpv6Addresses(v []Ipv6Addr)`

SetIpv6Addresses sets Ipv6Addresses field to given value.

### HasIpv6Addresses

`func (o *NFProfile) HasIpv6Addresses() bool`

HasIpv6Addresses returns a boolean if a field has been set.

### GetAllowedPlmns

`func (o *NFProfile) GetAllowedPlmns() []PlmnId`

GetAllowedPlmns returns the AllowedPlmns field if non-nil, zero value otherwise.

### GetAllowedPlmnsOk

`func (o *NFProfile) GetAllowedPlmnsOk() (*[]PlmnId, bool)`

GetAllowedPlmnsOk returns a tuple with the AllowedPlmns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedPlmns

`func (o *NFProfile) SetAllowedPlmns(v []PlmnId)`

SetAllowedPlmns sets AllowedPlmns field to given value.

### HasAllowedPlmns

`func (o *NFProfile) HasAllowedPlmns() bool`

HasAllowedPlmns returns a boolean if a field has been set.

### GetAllowedSnpns

`func (o *NFProfile) GetAllowedSnpns() []PlmnIdNid`

GetAllowedSnpns returns the AllowedSnpns field if non-nil, zero value otherwise.

### GetAllowedSnpnsOk

`func (o *NFProfile) GetAllowedSnpnsOk() (*[]PlmnIdNid, bool)`

GetAllowedSnpnsOk returns a tuple with the AllowedSnpns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedSnpns

`func (o *NFProfile) SetAllowedSnpns(v []PlmnIdNid)`

SetAllowedSnpns sets AllowedSnpns field to given value.

### HasAllowedSnpns

`func (o *NFProfile) HasAllowedSnpns() bool`

HasAllowedSnpns returns a boolean if a field has been set.

### GetAllowedNfTypes

`func (o *NFProfile) GetAllowedNfTypes() []NFType`

GetAllowedNfTypes returns the AllowedNfTypes field if non-nil, zero value otherwise.

### GetAllowedNfTypesOk

`func (o *NFProfile) GetAllowedNfTypesOk() (*[]NFType, bool)`

GetAllowedNfTypesOk returns a tuple with the AllowedNfTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedNfTypes

`func (o *NFProfile) SetAllowedNfTypes(v []NFType)`

SetAllowedNfTypes sets AllowedNfTypes field to given value.

### HasAllowedNfTypes

`func (o *NFProfile) HasAllowedNfTypes() bool`

HasAllowedNfTypes returns a boolean if a field has been set.

### GetAllowedNfDomains

`func (o *NFProfile) GetAllowedNfDomains() []string`

GetAllowedNfDomains returns the AllowedNfDomains field if non-nil, zero value otherwise.

### GetAllowedNfDomainsOk

`func (o *NFProfile) GetAllowedNfDomainsOk() (*[]string, bool)`

GetAllowedNfDomainsOk returns a tuple with the AllowedNfDomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedNfDomains

`func (o *NFProfile) SetAllowedNfDomains(v []string)`

SetAllowedNfDomains sets AllowedNfDomains field to given value.

### HasAllowedNfDomains

`func (o *NFProfile) HasAllowedNfDomains() bool`

HasAllowedNfDomains returns a boolean if a field has been set.

### GetAllowedNssais

`func (o *NFProfile) GetAllowedNssais() []ExtSnssai`

GetAllowedNssais returns the AllowedNssais field if non-nil, zero value otherwise.

### GetAllowedNssaisOk

`func (o *NFProfile) GetAllowedNssaisOk() (*[]ExtSnssai, bool)`

GetAllowedNssaisOk returns a tuple with the AllowedNssais field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedNssais

`func (o *NFProfile) SetAllowedNssais(v []ExtSnssai)`

SetAllowedNssais sets AllowedNssais field to given value.

### HasAllowedNssais

`func (o *NFProfile) HasAllowedNssais() bool`

HasAllowedNssais returns a boolean if a field has been set.

### GetPriority

`func (o *NFProfile) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *NFProfile) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *NFProfile) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *NFProfile) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetCapacity

`func (o *NFProfile) GetCapacity() int32`

GetCapacity returns the Capacity field if non-nil, zero value otherwise.

### GetCapacityOk

`func (o *NFProfile) GetCapacityOk() (*int32, bool)`

GetCapacityOk returns a tuple with the Capacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacity

`func (o *NFProfile) SetCapacity(v int32)`

SetCapacity sets Capacity field to given value.

### HasCapacity

`func (o *NFProfile) HasCapacity() bool`

HasCapacity returns a boolean if a field has been set.

### GetLoad

`func (o *NFProfile) GetLoad() int32`

GetLoad returns the Load field if non-nil, zero value otherwise.

### GetLoadOk

`func (o *NFProfile) GetLoadOk() (*int32, bool)`

GetLoadOk returns a tuple with the Load field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoad

`func (o *NFProfile) SetLoad(v int32)`

SetLoad sets Load field to given value.

### HasLoad

`func (o *NFProfile) HasLoad() bool`

HasLoad returns a boolean if a field has been set.

### GetLoadTimeStamp

`func (o *NFProfile) GetLoadTimeStamp() time.Time`

GetLoadTimeStamp returns the LoadTimeStamp field if non-nil, zero value otherwise.

### GetLoadTimeStampOk

`func (o *NFProfile) GetLoadTimeStampOk() (*time.Time, bool)`

GetLoadTimeStampOk returns a tuple with the LoadTimeStamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadTimeStamp

`func (o *NFProfile) SetLoadTimeStamp(v time.Time)`

SetLoadTimeStamp sets LoadTimeStamp field to given value.

### HasLoadTimeStamp

`func (o *NFProfile) HasLoadTimeStamp() bool`

HasLoadTimeStamp returns a boolean if a field has been set.

### GetLocality

`func (o *NFProfile) GetLocality() string`

GetLocality returns the Locality field if non-nil, zero value otherwise.

### GetLocalityOk

`func (o *NFProfile) GetLocalityOk() (*string, bool)`

GetLocalityOk returns a tuple with the Locality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocality

`func (o *NFProfile) SetLocality(v string)`

SetLocality sets Locality field to given value.

### HasLocality

`func (o *NFProfile) HasLocality() bool`

HasLocality returns a boolean if a field has been set.

### GetExtLocality

`func (o *NFProfile) GetExtLocality() map[string]string`

GetExtLocality returns the ExtLocality field if non-nil, zero value otherwise.

### GetExtLocalityOk

`func (o *NFProfile) GetExtLocalityOk() (*map[string]string, bool)`

GetExtLocalityOk returns a tuple with the ExtLocality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtLocality

`func (o *NFProfile) SetExtLocality(v map[string]string)`

SetExtLocality sets ExtLocality field to given value.

### HasExtLocality

`func (o *NFProfile) HasExtLocality() bool`

HasExtLocality returns a boolean if a field has been set.

### GetUdrInfo

`func (o *NFProfile) GetUdrInfo() UdrInfo`

GetUdrInfo returns the UdrInfo field if non-nil, zero value otherwise.

### GetUdrInfoOk

`func (o *NFProfile) GetUdrInfoOk() (*UdrInfo, bool)`

GetUdrInfoOk returns a tuple with the UdrInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUdrInfo

`func (o *NFProfile) SetUdrInfo(v UdrInfo)`

SetUdrInfo sets UdrInfo field to given value.

### HasUdrInfo

`func (o *NFProfile) HasUdrInfo() bool`

HasUdrInfo returns a boolean if a field has been set.

### GetUdrInfoList

`func (o *NFProfile) GetUdrInfoList() map[string]UdrInfo`

GetUdrInfoList returns the UdrInfoList field if non-nil, zero value otherwise.

### GetUdrInfoListOk

`func (o *NFProfile) GetUdrInfoListOk() (*map[string]UdrInfo, bool)`

GetUdrInfoListOk returns a tuple with the UdrInfoList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUdrInfoList

`func (o *NFProfile) SetUdrInfoList(v map[string]UdrInfo)`

SetUdrInfoList sets UdrInfoList field to given value.

### HasUdrInfoList

`func (o *NFProfile) HasUdrInfoList() bool`

HasUdrInfoList returns a boolean if a field has been set.

### GetUdmInfo

`func (o *NFProfile) GetUdmInfo() UdmInfo`

GetUdmInfo returns the UdmInfo field if non-nil, zero value otherwise.

### GetUdmInfoOk

`func (o *NFProfile) GetUdmInfoOk() (*UdmInfo, bool)`

GetUdmInfoOk returns a tuple with the UdmInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUdmInfo

`func (o *NFProfile) SetUdmInfo(v UdmInfo)`

SetUdmInfo sets UdmInfo field to given value.

### HasUdmInfo

`func (o *NFProfile) HasUdmInfo() bool`

HasUdmInfo returns a boolean if a field has been set.

### GetUdmInfoList

`func (o *NFProfile) GetUdmInfoList() map[string]UdmInfo`

GetUdmInfoList returns the UdmInfoList field if non-nil, zero value otherwise.

### GetUdmInfoListOk

`func (o *NFProfile) GetUdmInfoListOk() (*map[string]UdmInfo, bool)`

GetUdmInfoListOk returns a tuple with the UdmInfoList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUdmInfoList

`func (o *NFProfile) SetUdmInfoList(v map[string]UdmInfo)`

SetUdmInfoList sets UdmInfoList field to given value.

### HasUdmInfoList

`func (o *NFProfile) HasUdmInfoList() bool`

HasUdmInfoList returns a boolean if a field has been set.

### GetAusfInfo

`func (o *NFProfile) GetAusfInfo() AusfInfo`

GetAusfInfo returns the AusfInfo field if non-nil, zero value otherwise.

### GetAusfInfoOk

`func (o *NFProfile) GetAusfInfoOk() (*AusfInfo, bool)`

GetAusfInfoOk returns a tuple with the AusfInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAusfInfo

`func (o *NFProfile) SetAusfInfo(v AusfInfo)`

SetAusfInfo sets AusfInfo field to given value.

### HasAusfInfo

`func (o *NFProfile) HasAusfInfo() bool`

HasAusfInfo returns a boolean if a field has been set.

### GetAusfInfoList

`func (o *NFProfile) GetAusfInfoList() map[string]AusfInfo`

GetAusfInfoList returns the AusfInfoList field if non-nil, zero value otherwise.

### GetAusfInfoListOk

`func (o *NFProfile) GetAusfInfoListOk() (*map[string]AusfInfo, bool)`

GetAusfInfoListOk returns a tuple with the AusfInfoList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAusfInfoList

`func (o *NFProfile) SetAusfInfoList(v map[string]AusfInfo)`

SetAusfInfoList sets AusfInfoList field to given value.

### HasAusfInfoList

`func (o *NFProfile) HasAusfInfoList() bool`

HasAusfInfoList returns a boolean if a field has been set.

### GetAmfInfo

`func (o *NFProfile) GetAmfInfo() AmfInfo`

GetAmfInfo returns the AmfInfo field if non-nil, zero value otherwise.

### GetAmfInfoOk

`func (o *NFProfile) GetAmfInfoOk() (*AmfInfo, bool)`

GetAmfInfoOk returns a tuple with the AmfInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmfInfo

`func (o *NFProfile) SetAmfInfo(v AmfInfo)`

SetAmfInfo sets AmfInfo field to given value.

### HasAmfInfo

`func (o *NFProfile) HasAmfInfo() bool`

HasAmfInfo returns a boolean if a field has been set.

### GetAmfInfoList

`func (o *NFProfile) GetAmfInfoList() map[string]AmfInfo`

GetAmfInfoList returns the AmfInfoList field if non-nil, zero value otherwise.

### GetAmfInfoListOk

`func (o *NFProfile) GetAmfInfoListOk() (*map[string]AmfInfo, bool)`

GetAmfInfoListOk returns a tuple with the AmfInfoList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmfInfoList

`func (o *NFProfile) SetAmfInfoList(v map[string]AmfInfo)`

SetAmfInfoList sets AmfInfoList field to given value.

### HasAmfInfoList

`func (o *NFProfile) HasAmfInfoList() bool`

HasAmfInfoList returns a boolean if a field has been set.

### GetSmfInfo

`func (o *NFProfile) GetSmfInfo() SmfInfo`

GetSmfInfo returns the SmfInfo field if non-nil, zero value otherwise.

### GetSmfInfoOk

`func (o *NFProfile) GetSmfInfoOk() (*SmfInfo, bool)`

GetSmfInfoOk returns a tuple with the SmfInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmfInfo

`func (o *NFProfile) SetSmfInfo(v SmfInfo)`

SetSmfInfo sets SmfInfo field to given value.

### HasSmfInfo

`func (o *NFProfile) HasSmfInfo() bool`

HasSmfInfo returns a boolean if a field has been set.

### GetSmfInfoList

`func (o *NFProfile) GetSmfInfoList() map[string]SmfInfo`

GetSmfInfoList returns the SmfInfoList field if non-nil, zero value otherwise.

### GetSmfInfoListOk

`func (o *NFProfile) GetSmfInfoListOk() (*map[string]SmfInfo, bool)`

GetSmfInfoListOk returns a tuple with the SmfInfoList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmfInfoList

`func (o *NFProfile) SetSmfInfoList(v map[string]SmfInfo)`

SetSmfInfoList sets SmfInfoList field to given value.

### HasSmfInfoList

`func (o *NFProfile) HasSmfInfoList() bool`

HasSmfInfoList returns a boolean if a field has been set.

### GetUpfInfo

`func (o *NFProfile) GetUpfInfo() UpfInfo`

GetUpfInfo returns the UpfInfo field if non-nil, zero value otherwise.

### GetUpfInfoOk

`func (o *NFProfile) GetUpfInfoOk() (*UpfInfo, bool)`

GetUpfInfoOk returns a tuple with the UpfInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpfInfo

`func (o *NFProfile) SetUpfInfo(v UpfInfo)`

SetUpfInfo sets UpfInfo field to given value.

### HasUpfInfo

`func (o *NFProfile) HasUpfInfo() bool`

HasUpfInfo returns a boolean if a field has been set.

### GetUpfInfoList

`func (o *NFProfile) GetUpfInfoList() map[string]UpfInfo`

GetUpfInfoList returns the UpfInfoList field if non-nil, zero value otherwise.

### GetUpfInfoListOk

`func (o *NFProfile) GetUpfInfoListOk() (*map[string]UpfInfo, bool)`

GetUpfInfoListOk returns a tuple with the UpfInfoList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpfInfoList

`func (o *NFProfile) SetUpfInfoList(v map[string]UpfInfo)`

SetUpfInfoList sets UpfInfoList field to given value.

### HasUpfInfoList

`func (o *NFProfile) HasUpfInfoList() bool`

HasUpfInfoList returns a boolean if a field has been set.

### GetPcfInfo

`func (o *NFProfile) GetPcfInfo() PcfInfo`

GetPcfInfo returns the PcfInfo field if non-nil, zero value otherwise.

### GetPcfInfoOk

`func (o *NFProfile) GetPcfInfoOk() (*PcfInfo, bool)`

GetPcfInfoOk returns a tuple with the PcfInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPcfInfo

`func (o *NFProfile) SetPcfInfo(v PcfInfo)`

SetPcfInfo sets PcfInfo field to given value.

### HasPcfInfo

`func (o *NFProfile) HasPcfInfo() bool`

HasPcfInfo returns a boolean if a field has been set.

### GetPcfInfoList

`func (o *NFProfile) GetPcfInfoList() map[string]PcfInfo`

GetPcfInfoList returns the PcfInfoList field if non-nil, zero value otherwise.

### GetPcfInfoListOk

`func (o *NFProfile) GetPcfInfoListOk() (*map[string]PcfInfo, bool)`

GetPcfInfoListOk returns a tuple with the PcfInfoList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPcfInfoList

`func (o *NFProfile) SetPcfInfoList(v map[string]PcfInfo)`

SetPcfInfoList sets PcfInfoList field to given value.

### HasPcfInfoList

`func (o *NFProfile) HasPcfInfoList() bool`

HasPcfInfoList returns a boolean if a field has been set.

### GetBsfInfo

`func (o *NFProfile) GetBsfInfo() BsfInfo`

GetBsfInfo returns the BsfInfo field if non-nil, zero value otherwise.

### GetBsfInfoOk

`func (o *NFProfile) GetBsfInfoOk() (*BsfInfo, bool)`

GetBsfInfoOk returns a tuple with the BsfInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBsfInfo

`func (o *NFProfile) SetBsfInfo(v BsfInfo)`

SetBsfInfo sets BsfInfo field to given value.

### HasBsfInfo

`func (o *NFProfile) HasBsfInfo() bool`

HasBsfInfo returns a boolean if a field has been set.

### GetBsfInfoList

`func (o *NFProfile) GetBsfInfoList() map[string]BsfInfo`

GetBsfInfoList returns the BsfInfoList field if non-nil, zero value otherwise.

### GetBsfInfoListOk

`func (o *NFProfile) GetBsfInfoListOk() (*map[string]BsfInfo, bool)`

GetBsfInfoListOk returns a tuple with the BsfInfoList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBsfInfoList

`func (o *NFProfile) SetBsfInfoList(v map[string]BsfInfo)`

SetBsfInfoList sets BsfInfoList field to given value.

### HasBsfInfoList

`func (o *NFProfile) HasBsfInfoList() bool`

HasBsfInfoList returns a boolean if a field has been set.

### GetChfInfo

`func (o *NFProfile) GetChfInfo() ChfInfo`

GetChfInfo returns the ChfInfo field if non-nil, zero value otherwise.

### GetChfInfoOk

`func (o *NFProfile) GetChfInfoOk() (*ChfInfo, bool)`

GetChfInfoOk returns a tuple with the ChfInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChfInfo

`func (o *NFProfile) SetChfInfo(v ChfInfo)`

SetChfInfo sets ChfInfo field to given value.

### HasChfInfo

`func (o *NFProfile) HasChfInfo() bool`

HasChfInfo returns a boolean if a field has been set.

### GetChfInfoList

`func (o *NFProfile) GetChfInfoList() map[string]ChfInfo`

GetChfInfoList returns the ChfInfoList field if non-nil, zero value otherwise.

### GetChfInfoListOk

`func (o *NFProfile) GetChfInfoListOk() (*map[string]ChfInfo, bool)`

GetChfInfoListOk returns a tuple with the ChfInfoList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChfInfoList

`func (o *NFProfile) SetChfInfoList(v map[string]ChfInfo)`

SetChfInfoList sets ChfInfoList field to given value.

### HasChfInfoList

`func (o *NFProfile) HasChfInfoList() bool`

HasChfInfoList returns a boolean if a field has been set.

### GetNefInfo

`func (o *NFProfile) GetNefInfo() NefInfo`

GetNefInfo returns the NefInfo field if non-nil, zero value otherwise.

### GetNefInfoOk

`func (o *NFProfile) GetNefInfoOk() (*NefInfo, bool)`

GetNefInfoOk returns a tuple with the NefInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNefInfo

`func (o *NFProfile) SetNefInfo(v NefInfo)`

SetNefInfo sets NefInfo field to given value.

### HasNefInfo

`func (o *NFProfile) HasNefInfo() bool`

HasNefInfo returns a boolean if a field has been set.

### GetNrfInfo

`func (o *NFProfile) GetNrfInfo() NrfInfo`

GetNrfInfo returns the NrfInfo field if non-nil, zero value otherwise.

### GetNrfInfoOk

`func (o *NFProfile) GetNrfInfoOk() (*NrfInfo, bool)`

GetNrfInfoOk returns a tuple with the NrfInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNrfInfo

`func (o *NFProfile) SetNrfInfo(v NrfInfo)`

SetNrfInfo sets NrfInfo field to given value.

### HasNrfInfo

`func (o *NFProfile) HasNrfInfo() bool`

HasNrfInfo returns a boolean if a field has been set.

### GetUdsfInfo

`func (o *NFProfile) GetUdsfInfo() UdsfInfo`

GetUdsfInfo returns the UdsfInfo field if non-nil, zero value otherwise.

### GetUdsfInfoOk

`func (o *NFProfile) GetUdsfInfoOk() (*UdsfInfo, bool)`

GetUdsfInfoOk returns a tuple with the UdsfInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUdsfInfo

`func (o *NFProfile) SetUdsfInfo(v UdsfInfo)`

SetUdsfInfo sets UdsfInfo field to given value.

### HasUdsfInfo

`func (o *NFProfile) HasUdsfInfo() bool`

HasUdsfInfo returns a boolean if a field has been set.

### GetUdsfInfoList

`func (o *NFProfile) GetUdsfInfoList() map[string]UdsfInfo`

GetUdsfInfoList returns the UdsfInfoList field if non-nil, zero value otherwise.

### GetUdsfInfoListOk

`func (o *NFProfile) GetUdsfInfoListOk() (*map[string]UdsfInfo, bool)`

GetUdsfInfoListOk returns a tuple with the UdsfInfoList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUdsfInfoList

`func (o *NFProfile) SetUdsfInfoList(v map[string]UdsfInfo)`

SetUdsfInfoList sets UdsfInfoList field to given value.

### HasUdsfInfoList

`func (o *NFProfile) HasUdsfInfoList() bool`

HasUdsfInfoList returns a boolean if a field has been set.

### GetNwdafInfo

`func (o *NFProfile) GetNwdafInfo() NwdafInfo`

GetNwdafInfo returns the NwdafInfo field if non-nil, zero value otherwise.

### GetNwdafInfoOk

`func (o *NFProfile) GetNwdafInfoOk() (*NwdafInfo, bool)`

GetNwdafInfoOk returns a tuple with the NwdafInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNwdafInfo

`func (o *NFProfile) SetNwdafInfo(v NwdafInfo)`

SetNwdafInfo sets NwdafInfo field to given value.

### HasNwdafInfo

`func (o *NFProfile) HasNwdafInfo() bool`

HasNwdafInfo returns a boolean if a field has been set.

### GetNwdafInfoList

`func (o *NFProfile) GetNwdafInfoList() map[string]NwdafInfo`

GetNwdafInfoList returns the NwdafInfoList field if non-nil, zero value otherwise.

### GetNwdafInfoListOk

`func (o *NFProfile) GetNwdafInfoListOk() (*map[string]NwdafInfo, bool)`

GetNwdafInfoListOk returns a tuple with the NwdafInfoList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNwdafInfoList

`func (o *NFProfile) SetNwdafInfoList(v map[string]NwdafInfo)`

SetNwdafInfoList sets NwdafInfoList field to given value.

### HasNwdafInfoList

`func (o *NFProfile) HasNwdafInfoList() bool`

HasNwdafInfoList returns a boolean if a field has been set.

### GetPcscfInfoList

`func (o *NFProfile) GetPcscfInfoList() map[string]PcscfInfo`

GetPcscfInfoList returns the PcscfInfoList field if non-nil, zero value otherwise.

### GetPcscfInfoListOk

`func (o *NFProfile) GetPcscfInfoListOk() (*map[string]PcscfInfo, bool)`

GetPcscfInfoListOk returns a tuple with the PcscfInfoList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPcscfInfoList

`func (o *NFProfile) SetPcscfInfoList(v map[string]PcscfInfo)`

SetPcscfInfoList sets PcscfInfoList field to given value.

### HasPcscfInfoList

`func (o *NFProfile) HasPcscfInfoList() bool`

HasPcscfInfoList returns a boolean if a field has been set.

### GetHssInfoList

`func (o *NFProfile) GetHssInfoList() map[string]HssInfo`

GetHssInfoList returns the HssInfoList field if non-nil, zero value otherwise.

### GetHssInfoListOk

`func (o *NFProfile) GetHssInfoListOk() (*map[string]HssInfo, bool)`

GetHssInfoListOk returns a tuple with the HssInfoList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHssInfoList

`func (o *NFProfile) SetHssInfoList(v map[string]HssInfo)`

SetHssInfoList sets HssInfoList field to given value.

### HasHssInfoList

`func (o *NFProfile) HasHssInfoList() bool`

HasHssInfoList returns a boolean if a field has been set.

### GetCustomInfo

`func (o *NFProfile) GetCustomInfo() map[string]interface{}`

GetCustomInfo returns the CustomInfo field if non-nil, zero value otherwise.

### GetCustomInfoOk

`func (o *NFProfile) GetCustomInfoOk() (*map[string]interface{}, bool)`

GetCustomInfoOk returns a tuple with the CustomInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomInfo

`func (o *NFProfile) SetCustomInfo(v map[string]interface{})`

SetCustomInfo sets CustomInfo field to given value.

### HasCustomInfo

`func (o *NFProfile) HasCustomInfo() bool`

HasCustomInfo returns a boolean if a field has been set.

### GetRecoveryTime

`func (o *NFProfile) GetRecoveryTime() time.Time`

GetRecoveryTime returns the RecoveryTime field if non-nil, zero value otherwise.

### GetRecoveryTimeOk

`func (o *NFProfile) GetRecoveryTimeOk() (*time.Time, bool)`

GetRecoveryTimeOk returns a tuple with the RecoveryTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoveryTime

`func (o *NFProfile) SetRecoveryTime(v time.Time)`

SetRecoveryTime sets RecoveryTime field to given value.

### HasRecoveryTime

`func (o *NFProfile) HasRecoveryTime() bool`

HasRecoveryTime returns a boolean if a field has been set.

### GetNfServicePersistence

`func (o *NFProfile) GetNfServicePersistence() bool`

GetNfServicePersistence returns the NfServicePersistence field if non-nil, zero value otherwise.

### GetNfServicePersistenceOk

`func (o *NFProfile) GetNfServicePersistenceOk() (*bool, bool)`

GetNfServicePersistenceOk returns a tuple with the NfServicePersistence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNfServicePersistence

`func (o *NFProfile) SetNfServicePersistence(v bool)`

SetNfServicePersistence sets NfServicePersistence field to given value.

### HasNfServicePersistence

`func (o *NFProfile) HasNfServicePersistence() bool`

HasNfServicePersistence returns a boolean if a field has been set.

### GetNfServices

`func (o *NFProfile) GetNfServices() []NFService`

GetNfServices returns the NfServices field if non-nil, zero value otherwise.

### GetNfServicesOk

`func (o *NFProfile) GetNfServicesOk() (*[]NFService, bool)`

GetNfServicesOk returns a tuple with the NfServices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNfServices

`func (o *NFProfile) SetNfServices(v []NFService)`

SetNfServices sets NfServices field to given value.

### HasNfServices

`func (o *NFProfile) HasNfServices() bool`

HasNfServices returns a boolean if a field has been set.

### GetNfServiceList

`func (o *NFProfile) GetNfServiceList() map[string]NFService`

GetNfServiceList returns the NfServiceList field if non-nil, zero value otherwise.

### GetNfServiceListOk

`func (o *NFProfile) GetNfServiceListOk() (*map[string]NFService, bool)`

GetNfServiceListOk returns a tuple with the NfServiceList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNfServiceList

`func (o *NFProfile) SetNfServiceList(v map[string]NFService)`

SetNfServiceList sets NfServiceList field to given value.

### HasNfServiceList

`func (o *NFProfile) HasNfServiceList() bool`

HasNfServiceList returns a boolean if a field has been set.

### GetNfProfileChangesSupportInd

`func (o *NFProfile) GetNfProfileChangesSupportInd() bool`

GetNfProfileChangesSupportInd returns the NfProfileChangesSupportInd field if non-nil, zero value otherwise.

### GetNfProfileChangesSupportIndOk

`func (o *NFProfile) GetNfProfileChangesSupportIndOk() (*bool, bool)`

GetNfProfileChangesSupportIndOk returns a tuple with the NfProfileChangesSupportInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNfProfileChangesSupportInd

`func (o *NFProfile) SetNfProfileChangesSupportInd(v bool)`

SetNfProfileChangesSupportInd sets NfProfileChangesSupportInd field to given value.

### HasNfProfileChangesSupportInd

`func (o *NFProfile) HasNfProfileChangesSupportInd() bool`

HasNfProfileChangesSupportInd returns a boolean if a field has been set.

### GetNfProfileChangesInd

`func (o *NFProfile) GetNfProfileChangesInd() bool`

GetNfProfileChangesInd returns the NfProfileChangesInd field if non-nil, zero value otherwise.

### GetNfProfileChangesIndOk

`func (o *NFProfile) GetNfProfileChangesIndOk() (*bool, bool)`

GetNfProfileChangesIndOk returns a tuple with the NfProfileChangesInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNfProfileChangesInd

`func (o *NFProfile) SetNfProfileChangesInd(v bool)`

SetNfProfileChangesInd sets NfProfileChangesInd field to given value.

### HasNfProfileChangesInd

`func (o *NFProfile) HasNfProfileChangesInd() bool`

HasNfProfileChangesInd returns a boolean if a field has been set.

### GetDefaultNotificationSubscriptions

`func (o *NFProfile) GetDefaultNotificationSubscriptions() []DefaultNotificationSubscription`

GetDefaultNotificationSubscriptions returns the DefaultNotificationSubscriptions field if non-nil, zero value otherwise.

### GetDefaultNotificationSubscriptionsOk

`func (o *NFProfile) GetDefaultNotificationSubscriptionsOk() (*[]DefaultNotificationSubscription, bool)`

GetDefaultNotificationSubscriptionsOk returns a tuple with the DefaultNotificationSubscriptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultNotificationSubscriptions

`func (o *NFProfile) SetDefaultNotificationSubscriptions(v []DefaultNotificationSubscription)`

SetDefaultNotificationSubscriptions sets DefaultNotificationSubscriptions field to given value.

### HasDefaultNotificationSubscriptions

`func (o *NFProfile) HasDefaultNotificationSubscriptions() bool`

HasDefaultNotificationSubscriptions returns a boolean if a field has been set.

### GetLmfInfo

`func (o *NFProfile) GetLmfInfo() LmfInfo`

GetLmfInfo returns the LmfInfo field if non-nil, zero value otherwise.

### GetLmfInfoOk

`func (o *NFProfile) GetLmfInfoOk() (*LmfInfo, bool)`

GetLmfInfoOk returns a tuple with the LmfInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLmfInfo

`func (o *NFProfile) SetLmfInfo(v LmfInfo)`

SetLmfInfo sets LmfInfo field to given value.

### HasLmfInfo

`func (o *NFProfile) HasLmfInfo() bool`

HasLmfInfo returns a boolean if a field has been set.

### GetGmlcInfo

`func (o *NFProfile) GetGmlcInfo() GmlcInfo`

GetGmlcInfo returns the GmlcInfo field if non-nil, zero value otherwise.

### GetGmlcInfoOk

`func (o *NFProfile) GetGmlcInfoOk() (*GmlcInfo, bool)`

GetGmlcInfoOk returns a tuple with the GmlcInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGmlcInfo

`func (o *NFProfile) SetGmlcInfo(v GmlcInfo)`

SetGmlcInfo sets GmlcInfo field to given value.

### HasGmlcInfo

`func (o *NFProfile) HasGmlcInfo() bool`

HasGmlcInfo returns a boolean if a field has been set.

### GetNfSetIdList

`func (o *NFProfile) GetNfSetIdList() []string`

GetNfSetIdList returns the NfSetIdList field if non-nil, zero value otherwise.

### GetNfSetIdListOk

`func (o *NFProfile) GetNfSetIdListOk() (*[]string, bool)`

GetNfSetIdListOk returns a tuple with the NfSetIdList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNfSetIdList

`func (o *NFProfile) SetNfSetIdList(v []string)`

SetNfSetIdList sets NfSetIdList field to given value.

### HasNfSetIdList

`func (o *NFProfile) HasNfSetIdList() bool`

HasNfSetIdList returns a boolean if a field has been set.

### GetServingScope

`func (o *NFProfile) GetServingScope() []string`

GetServingScope returns the ServingScope field if non-nil, zero value otherwise.

### GetServingScopeOk

`func (o *NFProfile) GetServingScopeOk() (*[]string, bool)`

GetServingScopeOk returns a tuple with the ServingScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServingScope

`func (o *NFProfile) SetServingScope(v []string)`

SetServingScope sets ServingScope field to given value.

### HasServingScope

`func (o *NFProfile) HasServingScope() bool`

HasServingScope returns a boolean if a field has been set.

### GetLcHSupportInd

`func (o *NFProfile) GetLcHSupportInd() bool`

GetLcHSupportInd returns the LcHSupportInd field if non-nil, zero value otherwise.

### GetLcHSupportIndOk

`func (o *NFProfile) GetLcHSupportIndOk() (*bool, bool)`

GetLcHSupportIndOk returns a tuple with the LcHSupportInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLcHSupportInd

`func (o *NFProfile) SetLcHSupportInd(v bool)`

SetLcHSupportInd sets LcHSupportInd field to given value.

### HasLcHSupportInd

`func (o *NFProfile) HasLcHSupportInd() bool`

HasLcHSupportInd returns a boolean if a field has been set.

### GetOlcHSupportInd

`func (o *NFProfile) GetOlcHSupportInd() bool`

GetOlcHSupportInd returns the OlcHSupportInd field if non-nil, zero value otherwise.

### GetOlcHSupportIndOk

`func (o *NFProfile) GetOlcHSupportIndOk() (*bool, bool)`

GetOlcHSupportIndOk returns a tuple with the OlcHSupportInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOlcHSupportInd

`func (o *NFProfile) SetOlcHSupportInd(v bool)`

SetOlcHSupportInd sets OlcHSupportInd field to given value.

### HasOlcHSupportInd

`func (o *NFProfile) HasOlcHSupportInd() bool`

HasOlcHSupportInd returns a boolean if a field has been set.

### GetNfSetRecoveryTimeList

`func (o *NFProfile) GetNfSetRecoveryTimeList() map[string]time.Time`

GetNfSetRecoveryTimeList returns the NfSetRecoveryTimeList field if non-nil, zero value otherwise.

### GetNfSetRecoveryTimeListOk

`func (o *NFProfile) GetNfSetRecoveryTimeListOk() (*map[string]time.Time, bool)`

GetNfSetRecoveryTimeListOk returns a tuple with the NfSetRecoveryTimeList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNfSetRecoveryTimeList

`func (o *NFProfile) SetNfSetRecoveryTimeList(v map[string]time.Time)`

SetNfSetRecoveryTimeList sets NfSetRecoveryTimeList field to given value.

### HasNfSetRecoveryTimeList

`func (o *NFProfile) HasNfSetRecoveryTimeList() bool`

HasNfSetRecoveryTimeList returns a boolean if a field has been set.

### GetServiceSetRecoveryTimeList

`func (o *NFProfile) GetServiceSetRecoveryTimeList() map[string]time.Time`

GetServiceSetRecoveryTimeList returns the ServiceSetRecoveryTimeList field if non-nil, zero value otherwise.

### GetServiceSetRecoveryTimeListOk

`func (o *NFProfile) GetServiceSetRecoveryTimeListOk() (*map[string]time.Time, bool)`

GetServiceSetRecoveryTimeListOk returns a tuple with the ServiceSetRecoveryTimeList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceSetRecoveryTimeList

`func (o *NFProfile) SetServiceSetRecoveryTimeList(v map[string]time.Time)`

SetServiceSetRecoveryTimeList sets ServiceSetRecoveryTimeList field to given value.

### HasServiceSetRecoveryTimeList

`func (o *NFProfile) HasServiceSetRecoveryTimeList() bool`

HasServiceSetRecoveryTimeList returns a boolean if a field has been set.

### GetScpDomains

`func (o *NFProfile) GetScpDomains() []string`

GetScpDomains returns the ScpDomains field if non-nil, zero value otherwise.

### GetScpDomainsOk

`func (o *NFProfile) GetScpDomainsOk() (*[]string, bool)`

GetScpDomainsOk returns a tuple with the ScpDomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScpDomains

`func (o *NFProfile) SetScpDomains(v []string)`

SetScpDomains sets ScpDomains field to given value.

### HasScpDomains

`func (o *NFProfile) HasScpDomains() bool`

HasScpDomains returns a boolean if a field has been set.

### GetScpInfo

`func (o *NFProfile) GetScpInfo() ScpInfo`

GetScpInfo returns the ScpInfo field if non-nil, zero value otherwise.

### GetScpInfoOk

`func (o *NFProfile) GetScpInfoOk() (*ScpInfo, bool)`

GetScpInfoOk returns a tuple with the ScpInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScpInfo

`func (o *NFProfile) SetScpInfo(v ScpInfo)`

SetScpInfo sets ScpInfo field to given value.

### HasScpInfo

`func (o *NFProfile) HasScpInfo() bool`

HasScpInfo returns a boolean if a field has been set.

### GetSeppInfo

`func (o *NFProfile) GetSeppInfo() SeppInfo`

GetSeppInfo returns the SeppInfo field if non-nil, zero value otherwise.

### GetSeppInfoOk

`func (o *NFProfile) GetSeppInfoOk() (*SeppInfo, bool)`

GetSeppInfoOk returns a tuple with the SeppInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeppInfo

`func (o *NFProfile) SetSeppInfo(v SeppInfo)`

SetSeppInfo sets SeppInfo field to given value.

### HasSeppInfo

`func (o *NFProfile) HasSeppInfo() bool`

HasSeppInfo returns a boolean if a field has been set.

### GetVendorId

`func (o *NFProfile) GetVendorId() string`

GetVendorId returns the VendorId field if non-nil, zero value otherwise.

### GetVendorIdOk

`func (o *NFProfile) GetVendorIdOk() (*string, bool)`

GetVendorIdOk returns a tuple with the VendorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorId

`func (o *NFProfile) SetVendorId(v string)`

SetVendorId sets VendorId field to given value.

### HasVendorId

`func (o *NFProfile) HasVendorId() bool`

HasVendorId returns a boolean if a field has been set.

### GetSupportedVendorSpecificFeatures

`func (o *NFProfile) GetSupportedVendorSpecificFeatures() map[string][]VendorSpecificFeature`

GetSupportedVendorSpecificFeatures returns the SupportedVendorSpecificFeatures field if non-nil, zero value otherwise.

### GetSupportedVendorSpecificFeaturesOk

`func (o *NFProfile) GetSupportedVendorSpecificFeaturesOk() (*map[string][]VendorSpecificFeature, bool)`

GetSupportedVendorSpecificFeaturesOk returns a tuple with the SupportedVendorSpecificFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedVendorSpecificFeatures

`func (o *NFProfile) SetSupportedVendorSpecificFeatures(v map[string][]VendorSpecificFeature)`

SetSupportedVendorSpecificFeatures sets SupportedVendorSpecificFeatures field to given value.

### HasSupportedVendorSpecificFeatures

`func (o *NFProfile) HasSupportedVendorSpecificFeatures() bool`

HasSupportedVendorSpecificFeatures returns a boolean if a field has been set.

### GetAanfInfoList

`func (o *NFProfile) GetAanfInfoList() map[string]AanfInfo`

GetAanfInfoList returns the AanfInfoList field if non-nil, zero value otherwise.

### GetAanfInfoListOk

`func (o *NFProfile) GetAanfInfoListOk() (*map[string]AanfInfo, bool)`

GetAanfInfoListOk returns a tuple with the AanfInfoList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAanfInfoList

`func (o *NFProfile) SetAanfInfoList(v map[string]AanfInfo)`

SetAanfInfoList sets AanfInfoList field to given value.

### HasAanfInfoList

`func (o *NFProfile) HasAanfInfoList() bool`

HasAanfInfoList returns a boolean if a field has been set.

### GetVar5gDdnmfInfo

`func (o *NFProfile) GetVar5gDdnmfInfo() Model5GDdnmfInfo`

GetVar5gDdnmfInfo returns the Var5gDdnmfInfo field if non-nil, zero value otherwise.

### GetVar5gDdnmfInfoOk

`func (o *NFProfile) GetVar5gDdnmfInfoOk() (*Model5GDdnmfInfo, bool)`

GetVar5gDdnmfInfoOk returns a tuple with the Var5gDdnmfInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar5gDdnmfInfo

`func (o *NFProfile) SetVar5gDdnmfInfo(v Model5GDdnmfInfo)`

SetVar5gDdnmfInfo sets Var5gDdnmfInfo field to given value.

### HasVar5gDdnmfInfo

`func (o *NFProfile) HasVar5gDdnmfInfo() bool`

HasVar5gDdnmfInfo returns a boolean if a field has been set.

### GetMfafInfo

`func (o *NFProfile) GetMfafInfo() MfafInfo`

GetMfafInfo returns the MfafInfo field if non-nil, zero value otherwise.

### GetMfafInfoOk

`func (o *NFProfile) GetMfafInfoOk() (*MfafInfo, bool)`

GetMfafInfoOk returns a tuple with the MfafInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMfafInfo

`func (o *NFProfile) SetMfafInfo(v MfafInfo)`

SetMfafInfo sets MfafInfo field to given value.

### HasMfafInfo

`func (o *NFProfile) HasMfafInfo() bool`

HasMfafInfo returns a boolean if a field has been set.

### GetEasdfInfoList

`func (o *NFProfile) GetEasdfInfoList() map[string]EasdfInfo`

GetEasdfInfoList returns the EasdfInfoList field if non-nil, zero value otherwise.

### GetEasdfInfoListOk

`func (o *NFProfile) GetEasdfInfoListOk() (*map[string]EasdfInfo, bool)`

GetEasdfInfoListOk returns a tuple with the EasdfInfoList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEasdfInfoList

`func (o *NFProfile) SetEasdfInfoList(v map[string]EasdfInfo)`

SetEasdfInfoList sets EasdfInfoList field to given value.

### HasEasdfInfoList

`func (o *NFProfile) HasEasdfInfoList() bool`

HasEasdfInfoList returns a boolean if a field has been set.

### GetDccfInfo

`func (o *NFProfile) GetDccfInfo() DccfInfo`

GetDccfInfo returns the DccfInfo field if non-nil, zero value otherwise.

### GetDccfInfoOk

`func (o *NFProfile) GetDccfInfoOk() (*DccfInfo, bool)`

GetDccfInfoOk returns a tuple with the DccfInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDccfInfo

`func (o *NFProfile) SetDccfInfo(v DccfInfo)`

SetDccfInfo sets DccfInfo field to given value.

### HasDccfInfo

`func (o *NFProfile) HasDccfInfo() bool`

HasDccfInfo returns a boolean if a field has been set.

### GetNsacfInfoList

`func (o *NFProfile) GetNsacfInfoList() map[string]NsacfInfo`

GetNsacfInfoList returns the NsacfInfoList field if non-nil, zero value otherwise.

### GetNsacfInfoListOk

`func (o *NFProfile) GetNsacfInfoListOk() (*map[string]NsacfInfo, bool)`

GetNsacfInfoListOk returns a tuple with the NsacfInfoList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNsacfInfoList

`func (o *NFProfile) SetNsacfInfoList(v map[string]NsacfInfo)`

SetNsacfInfoList sets NsacfInfoList field to given value.

### HasNsacfInfoList

`func (o *NFProfile) HasNsacfInfoList() bool`

HasNsacfInfoList returns a boolean if a field has been set.

### GetMbSmfInfoList

`func (o *NFProfile) GetMbSmfInfoList() map[string]MbSmfInfo`

GetMbSmfInfoList returns the MbSmfInfoList field if non-nil, zero value otherwise.

### GetMbSmfInfoListOk

`func (o *NFProfile) GetMbSmfInfoListOk() (*map[string]MbSmfInfo, bool)`

GetMbSmfInfoListOk returns a tuple with the MbSmfInfoList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMbSmfInfoList

`func (o *NFProfile) SetMbSmfInfoList(v map[string]MbSmfInfo)`

SetMbSmfInfoList sets MbSmfInfoList field to given value.

### HasMbSmfInfoList

`func (o *NFProfile) HasMbSmfInfoList() bool`

HasMbSmfInfoList returns a boolean if a field has been set.

### GetTsctsfInfoList

`func (o *NFProfile) GetTsctsfInfoList() map[string]TsctsfInfo`

GetTsctsfInfoList returns the TsctsfInfoList field if non-nil, zero value otherwise.

### GetTsctsfInfoListOk

`func (o *NFProfile) GetTsctsfInfoListOk() (*map[string]TsctsfInfo, bool)`

GetTsctsfInfoListOk returns a tuple with the TsctsfInfoList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTsctsfInfoList

`func (o *NFProfile) SetTsctsfInfoList(v map[string]TsctsfInfo)`

SetTsctsfInfoList sets TsctsfInfoList field to given value.

### HasTsctsfInfoList

`func (o *NFProfile) HasTsctsfInfoList() bool`

HasTsctsfInfoList returns a boolean if a field has been set.

### GetMbUpfInfoList

`func (o *NFProfile) GetMbUpfInfoList() map[string]MbUpfInfo`

GetMbUpfInfoList returns the MbUpfInfoList field if non-nil, zero value otherwise.

### GetMbUpfInfoListOk

`func (o *NFProfile) GetMbUpfInfoListOk() (*map[string]MbUpfInfo, bool)`

GetMbUpfInfoListOk returns a tuple with the MbUpfInfoList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMbUpfInfoList

`func (o *NFProfile) SetMbUpfInfoList(v map[string]MbUpfInfo)`

SetMbUpfInfoList sets MbUpfInfoList field to given value.

### HasMbUpfInfoList

`func (o *NFProfile) HasMbUpfInfoList() bool`

HasMbUpfInfoList returns a boolean if a field has been set.

### GetTrustAfInfo

`func (o *NFProfile) GetTrustAfInfo() TrustAfInfo`

GetTrustAfInfo returns the TrustAfInfo field if non-nil, zero value otherwise.

### GetTrustAfInfoOk

`func (o *NFProfile) GetTrustAfInfoOk() (*TrustAfInfo, bool)`

GetTrustAfInfoOk returns a tuple with the TrustAfInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustAfInfo

`func (o *NFProfile) SetTrustAfInfo(v TrustAfInfo)`

SetTrustAfInfo sets TrustAfInfo field to given value.

### HasTrustAfInfo

`func (o *NFProfile) HasTrustAfInfo() bool`

HasTrustAfInfo returns a boolean if a field has been set.

### GetNssaafInfo

`func (o *NFProfile) GetNssaafInfo() NssaafInfo`

GetNssaafInfo returns the NssaafInfo field if non-nil, zero value otherwise.

### GetNssaafInfoOk

`func (o *NFProfile) GetNssaafInfoOk() (*NssaafInfo, bool)`

GetNssaafInfoOk returns a tuple with the NssaafInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNssaafInfo

`func (o *NFProfile) SetNssaafInfo(v NssaafInfo)`

SetNssaafInfo sets NssaafInfo field to given value.

### HasNssaafInfo

`func (o *NFProfile) HasNssaafInfo() bool`

HasNssaafInfo returns a boolean if a field has been set.

### GetHniList

`func (o *NFProfile) GetHniList() []string`

GetHniList returns the HniList field if non-nil, zero value otherwise.

### GetHniListOk

`func (o *NFProfile) GetHniListOk() (*[]string, bool)`

GetHniListOk returns a tuple with the HniList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHniList

`func (o *NFProfile) SetHniList(v []string)`

SetHniList sets HniList field to given value.

### HasHniList

`func (o *NFProfile) HasHniList() bool`

HasHniList returns a boolean if a field has been set.

### GetIwmscInfo

`func (o *NFProfile) GetIwmscInfo() IwmscInfo`

GetIwmscInfo returns the IwmscInfo field if non-nil, zero value otherwise.

### GetIwmscInfoOk

`func (o *NFProfile) GetIwmscInfoOk() (*IwmscInfo, bool)`

GetIwmscInfoOk returns a tuple with the IwmscInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIwmscInfo

`func (o *NFProfile) SetIwmscInfo(v IwmscInfo)`

SetIwmscInfo sets IwmscInfo field to given value.

### HasIwmscInfo

`func (o *NFProfile) HasIwmscInfo() bool`

HasIwmscInfo returns a boolean if a field has been set.

### GetMnpfInfo

`func (o *NFProfile) GetMnpfInfo() MnpfInfo`

GetMnpfInfo returns the MnpfInfo field if non-nil, zero value otherwise.

### GetMnpfInfoOk

`func (o *NFProfile) GetMnpfInfoOk() (*MnpfInfo, bool)`

GetMnpfInfoOk returns a tuple with the MnpfInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMnpfInfo

`func (o *NFProfile) SetMnpfInfo(v MnpfInfo)`

SetMnpfInfo sets MnpfInfo field to given value.

### HasMnpfInfo

`func (o *NFProfile) HasMnpfInfo() bool`

HasMnpfInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


