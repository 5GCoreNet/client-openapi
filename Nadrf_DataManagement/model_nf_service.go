/*
Nadrf_DataManagement

ADRF Data Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Nadrf_DataManagement

import (
	"encoding/json"
	"time"
)

// NFService Information of a given NF Service Instance; it is part of the NFProfile of an NF Instance 
type NFService struct {
	ServiceInstanceId string `json:"serviceInstanceId"`
	ServiceName ServiceName `json:"serviceName"`
	Versions []NFServiceVersion `json:"versions"`
	Scheme UriScheme `json:"scheme"`
	NfServiceStatus NFServiceStatus `json:"nfServiceStatus"`
	// Fully Qualified Domain Name
	Fqdn *string `json:"fqdn,omitempty"`
	// Fully Qualified Domain Name
	InterPlmnFqdn *string `json:"interPlmnFqdn,omitempty"`
	IpEndPoints []IpEndPoint `json:"ipEndPoints,omitempty"`
	ApiPrefix *string `json:"apiPrefix,omitempty"`
	DefaultNotificationSubscriptions []DefaultNotificationSubscription `json:"defaultNotificationSubscriptions,omitempty"`
	AllowedPlmns []PlmnId `json:"allowedPlmns,omitempty"`
	AllowedSnpns []PlmnIdNid `json:"allowedSnpns,omitempty"`
	AllowedNfTypes []NFType `json:"allowedNfTypes,omitempty"`
	AllowedNfDomains []string `json:"allowedNfDomains,omitempty"`
	AllowedNssais []ExtSnssai `json:"allowedNssais,omitempty"`
	// A map (list of key-value pairs) where NF Type serves as key
	AllowedOperationsPerNfType *map[string][]string `json:"allowedOperationsPerNfType,omitempty"`
	// A map (list of key-value pairs) where NF Instance Id serves as key
	AllowedOperationsPerNfInstance *map[string][]string `json:"allowedOperationsPerNfInstance,omitempty"`
	AllowedOperationsPerNfInstanceOverrides *bool `json:"allowedOperationsPerNfInstanceOverrides,omitempty"`
	Priority *int32 `json:"priority,omitempty"`
	Capacity *int32 `json:"capacity,omitempty"`
	Load *int32 `json:"load,omitempty"`
	// string with format 'date-time' as defined in OpenAPI.
	LoadTimeStamp *time.Time `json:"loadTimeStamp,omitempty"`
	// string with format 'date-time' as defined in OpenAPI.
	RecoveryTime *time.Time `json:"recoveryTime,omitempty"`
	// A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \"0\" to \"9\",  \"a\" to \"f\" or \"A\" to \"F\" and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported. 
	SupportedFeatures *string `json:"supportedFeatures,omitempty"`
	NfServiceSetIdList []string `json:"nfServiceSetIdList,omitempty"`
	SNssais []ExtSnssai `json:"sNssais,omitempty"`
	PerPlmnSnssaiList []PlmnSnssai `json:"perPlmnSnssaiList,omitempty"`
	// Vendor ID of the NF Service instance (Private Enterprise Number assigned by IANA)
	VendorId *string `json:"vendorId,omitempty"`
	// A map (list of key-value pairs) where IANA-assigned SMI Network Management Private Enterprise Codes serves as key 
	SupportedVendorSpecificFeatures *map[string][]VendorSpecificFeature `json:"supportedVendorSpecificFeatures,omitempty"`
	Oauth2Required *bool `json:"oauth2Required,omitempty"`
	PerPlmnOauth2ReqList *PlmnOauth2 `json:"perPlmnOauth2ReqList,omitempty"`
}

// NewNFService instantiates a new NFService object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNFService(serviceInstanceId string, serviceName ServiceName, versions []NFServiceVersion, scheme UriScheme, nfServiceStatus NFServiceStatus) *NFService {
	this := NFService{}
	this.ServiceInstanceId = serviceInstanceId
	this.ServiceName = serviceName
	this.Versions = versions
	this.Scheme = scheme
	this.NfServiceStatus = nfServiceStatus
	var allowedOperationsPerNfInstanceOverrides bool = false
	this.AllowedOperationsPerNfInstanceOverrides = &allowedOperationsPerNfInstanceOverrides
	return &this
}

// NewNFServiceWithDefaults instantiates a new NFService object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNFServiceWithDefaults() *NFService {
	this := NFService{}
	var allowedOperationsPerNfInstanceOverrides bool = false
	this.AllowedOperationsPerNfInstanceOverrides = &allowedOperationsPerNfInstanceOverrides
	return &this
}

// GetServiceInstanceId returns the ServiceInstanceId field value
func (o *NFService) GetServiceInstanceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServiceInstanceId
}

// GetServiceInstanceIdOk returns a tuple with the ServiceInstanceId field value
// and a boolean to check if the value has been set.
func (o *NFService) GetServiceInstanceIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.ServiceInstanceId, true
}

// SetServiceInstanceId sets field value
func (o *NFService) SetServiceInstanceId(v string) {
	o.ServiceInstanceId = v
}

// GetServiceName returns the ServiceName field value
func (o *NFService) GetServiceName() ServiceName {
	if o == nil {
		var ret ServiceName
		return ret
	}

	return o.ServiceName
}

// GetServiceNameOk returns a tuple with the ServiceName field value
// and a boolean to check if the value has been set.
func (o *NFService) GetServiceNameOk() (*ServiceName, bool) {
	if o == nil {
    return nil, false
	}
	return &o.ServiceName, true
}

// SetServiceName sets field value
func (o *NFService) SetServiceName(v ServiceName) {
	o.ServiceName = v
}

// GetVersions returns the Versions field value
func (o *NFService) GetVersions() []NFServiceVersion {
	if o == nil {
		var ret []NFServiceVersion
		return ret
	}

	return o.Versions
}

// GetVersionsOk returns a tuple with the Versions field value
// and a boolean to check if the value has been set.
func (o *NFService) GetVersionsOk() ([]NFServiceVersion, bool) {
	if o == nil {
    return nil, false
	}
	return o.Versions, true
}

// SetVersions sets field value
func (o *NFService) SetVersions(v []NFServiceVersion) {
	o.Versions = v
}

// GetScheme returns the Scheme field value
func (o *NFService) GetScheme() UriScheme {
	if o == nil {
		var ret UriScheme
		return ret
	}

	return o.Scheme
}

// GetSchemeOk returns a tuple with the Scheme field value
// and a boolean to check if the value has been set.
func (o *NFService) GetSchemeOk() (*UriScheme, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Scheme, true
}

// SetScheme sets field value
func (o *NFService) SetScheme(v UriScheme) {
	o.Scheme = v
}

// GetNfServiceStatus returns the NfServiceStatus field value
func (o *NFService) GetNfServiceStatus() NFServiceStatus {
	if o == nil {
		var ret NFServiceStatus
		return ret
	}

	return o.NfServiceStatus
}

// GetNfServiceStatusOk returns a tuple with the NfServiceStatus field value
// and a boolean to check if the value has been set.
func (o *NFService) GetNfServiceStatusOk() (*NFServiceStatus, bool) {
	if o == nil {
    return nil, false
	}
	return &o.NfServiceStatus, true
}

// SetNfServiceStatus sets field value
func (o *NFService) SetNfServiceStatus(v NFServiceStatus) {
	o.NfServiceStatus = v
}

// GetFqdn returns the Fqdn field value if set, zero value otherwise.
func (o *NFService) GetFqdn() string {
	if o == nil || isNil(o.Fqdn) {
		var ret string
		return ret
	}
	return *o.Fqdn
}

// GetFqdnOk returns a tuple with the Fqdn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NFService) GetFqdnOk() (*string, bool) {
	if o == nil || isNil(o.Fqdn) {
    return nil, false
	}
	return o.Fqdn, true
}

// HasFqdn returns a boolean if a field has been set.
func (o *NFService) HasFqdn() bool {
	if o != nil && !isNil(o.Fqdn) {
		return true
	}

	return false
}

// SetFqdn gets a reference to the given string and assigns it to the Fqdn field.
func (o *NFService) SetFqdn(v string) {
	o.Fqdn = &v
}

// GetInterPlmnFqdn returns the InterPlmnFqdn field value if set, zero value otherwise.
func (o *NFService) GetInterPlmnFqdn() string {
	if o == nil || isNil(o.InterPlmnFqdn) {
		var ret string
		return ret
	}
	return *o.InterPlmnFqdn
}

// GetInterPlmnFqdnOk returns a tuple with the InterPlmnFqdn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NFService) GetInterPlmnFqdnOk() (*string, bool) {
	if o == nil || isNil(o.InterPlmnFqdn) {
    return nil, false
	}
	return o.InterPlmnFqdn, true
}

// HasInterPlmnFqdn returns a boolean if a field has been set.
func (o *NFService) HasInterPlmnFqdn() bool {
	if o != nil && !isNil(o.InterPlmnFqdn) {
		return true
	}

	return false
}

// SetInterPlmnFqdn gets a reference to the given string and assigns it to the InterPlmnFqdn field.
func (o *NFService) SetInterPlmnFqdn(v string) {
	o.InterPlmnFqdn = &v
}

// GetIpEndPoints returns the IpEndPoints field value if set, zero value otherwise.
func (o *NFService) GetIpEndPoints() []IpEndPoint {
	if o == nil || isNil(o.IpEndPoints) {
		var ret []IpEndPoint
		return ret
	}
	return o.IpEndPoints
}

// GetIpEndPointsOk returns a tuple with the IpEndPoints field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NFService) GetIpEndPointsOk() ([]IpEndPoint, bool) {
	if o == nil || isNil(o.IpEndPoints) {
    return nil, false
	}
	return o.IpEndPoints, true
}

// HasIpEndPoints returns a boolean if a field has been set.
func (o *NFService) HasIpEndPoints() bool {
	if o != nil && !isNil(o.IpEndPoints) {
		return true
	}

	return false
}

// SetIpEndPoints gets a reference to the given []IpEndPoint and assigns it to the IpEndPoints field.
func (o *NFService) SetIpEndPoints(v []IpEndPoint) {
	o.IpEndPoints = v
}

// GetApiPrefix returns the ApiPrefix field value if set, zero value otherwise.
func (o *NFService) GetApiPrefix() string {
	if o == nil || isNil(o.ApiPrefix) {
		var ret string
		return ret
	}
	return *o.ApiPrefix
}

// GetApiPrefixOk returns a tuple with the ApiPrefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NFService) GetApiPrefixOk() (*string, bool) {
	if o == nil || isNil(o.ApiPrefix) {
    return nil, false
	}
	return o.ApiPrefix, true
}

// HasApiPrefix returns a boolean if a field has been set.
func (o *NFService) HasApiPrefix() bool {
	if o != nil && !isNil(o.ApiPrefix) {
		return true
	}

	return false
}

// SetApiPrefix gets a reference to the given string and assigns it to the ApiPrefix field.
func (o *NFService) SetApiPrefix(v string) {
	o.ApiPrefix = &v
}

// GetDefaultNotificationSubscriptions returns the DefaultNotificationSubscriptions field value if set, zero value otherwise.
func (o *NFService) GetDefaultNotificationSubscriptions() []DefaultNotificationSubscription {
	if o == nil || isNil(o.DefaultNotificationSubscriptions) {
		var ret []DefaultNotificationSubscription
		return ret
	}
	return o.DefaultNotificationSubscriptions
}

// GetDefaultNotificationSubscriptionsOk returns a tuple with the DefaultNotificationSubscriptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NFService) GetDefaultNotificationSubscriptionsOk() ([]DefaultNotificationSubscription, bool) {
	if o == nil || isNil(o.DefaultNotificationSubscriptions) {
    return nil, false
	}
	return o.DefaultNotificationSubscriptions, true
}

// HasDefaultNotificationSubscriptions returns a boolean if a field has been set.
func (o *NFService) HasDefaultNotificationSubscriptions() bool {
	if o != nil && !isNil(o.DefaultNotificationSubscriptions) {
		return true
	}

	return false
}

// SetDefaultNotificationSubscriptions gets a reference to the given []DefaultNotificationSubscription and assigns it to the DefaultNotificationSubscriptions field.
func (o *NFService) SetDefaultNotificationSubscriptions(v []DefaultNotificationSubscription) {
	o.DefaultNotificationSubscriptions = v
}

// GetAllowedPlmns returns the AllowedPlmns field value if set, zero value otherwise.
func (o *NFService) GetAllowedPlmns() []PlmnId {
	if o == nil || isNil(o.AllowedPlmns) {
		var ret []PlmnId
		return ret
	}
	return o.AllowedPlmns
}

// GetAllowedPlmnsOk returns a tuple with the AllowedPlmns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NFService) GetAllowedPlmnsOk() ([]PlmnId, bool) {
	if o == nil || isNil(o.AllowedPlmns) {
    return nil, false
	}
	return o.AllowedPlmns, true
}

// HasAllowedPlmns returns a boolean if a field has been set.
func (o *NFService) HasAllowedPlmns() bool {
	if o != nil && !isNil(o.AllowedPlmns) {
		return true
	}

	return false
}

// SetAllowedPlmns gets a reference to the given []PlmnId and assigns it to the AllowedPlmns field.
func (o *NFService) SetAllowedPlmns(v []PlmnId) {
	o.AllowedPlmns = v
}

// GetAllowedSnpns returns the AllowedSnpns field value if set, zero value otherwise.
func (o *NFService) GetAllowedSnpns() []PlmnIdNid {
	if o == nil || isNil(o.AllowedSnpns) {
		var ret []PlmnIdNid
		return ret
	}
	return o.AllowedSnpns
}

// GetAllowedSnpnsOk returns a tuple with the AllowedSnpns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NFService) GetAllowedSnpnsOk() ([]PlmnIdNid, bool) {
	if o == nil || isNil(o.AllowedSnpns) {
    return nil, false
	}
	return o.AllowedSnpns, true
}

// HasAllowedSnpns returns a boolean if a field has been set.
func (o *NFService) HasAllowedSnpns() bool {
	if o != nil && !isNil(o.AllowedSnpns) {
		return true
	}

	return false
}

// SetAllowedSnpns gets a reference to the given []PlmnIdNid and assigns it to the AllowedSnpns field.
func (o *NFService) SetAllowedSnpns(v []PlmnIdNid) {
	o.AllowedSnpns = v
}

// GetAllowedNfTypes returns the AllowedNfTypes field value if set, zero value otherwise.
func (o *NFService) GetAllowedNfTypes() []NFType {
	if o == nil || isNil(o.AllowedNfTypes) {
		var ret []NFType
		return ret
	}
	return o.AllowedNfTypes
}

// GetAllowedNfTypesOk returns a tuple with the AllowedNfTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NFService) GetAllowedNfTypesOk() ([]NFType, bool) {
	if o == nil || isNil(o.AllowedNfTypes) {
    return nil, false
	}
	return o.AllowedNfTypes, true
}

// HasAllowedNfTypes returns a boolean if a field has been set.
func (o *NFService) HasAllowedNfTypes() bool {
	if o != nil && !isNil(o.AllowedNfTypes) {
		return true
	}

	return false
}

// SetAllowedNfTypes gets a reference to the given []NFType and assigns it to the AllowedNfTypes field.
func (o *NFService) SetAllowedNfTypes(v []NFType) {
	o.AllowedNfTypes = v
}

// GetAllowedNfDomains returns the AllowedNfDomains field value if set, zero value otherwise.
func (o *NFService) GetAllowedNfDomains() []string {
	if o == nil || isNil(o.AllowedNfDomains) {
		var ret []string
		return ret
	}
	return o.AllowedNfDomains
}

// GetAllowedNfDomainsOk returns a tuple with the AllowedNfDomains field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NFService) GetAllowedNfDomainsOk() ([]string, bool) {
	if o == nil || isNil(o.AllowedNfDomains) {
    return nil, false
	}
	return o.AllowedNfDomains, true
}

// HasAllowedNfDomains returns a boolean if a field has been set.
func (o *NFService) HasAllowedNfDomains() bool {
	if o != nil && !isNil(o.AllowedNfDomains) {
		return true
	}

	return false
}

// SetAllowedNfDomains gets a reference to the given []string and assigns it to the AllowedNfDomains field.
func (o *NFService) SetAllowedNfDomains(v []string) {
	o.AllowedNfDomains = v
}

// GetAllowedNssais returns the AllowedNssais field value if set, zero value otherwise.
func (o *NFService) GetAllowedNssais() []ExtSnssai {
	if o == nil || isNil(o.AllowedNssais) {
		var ret []ExtSnssai
		return ret
	}
	return o.AllowedNssais
}

// GetAllowedNssaisOk returns a tuple with the AllowedNssais field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NFService) GetAllowedNssaisOk() ([]ExtSnssai, bool) {
	if o == nil || isNil(o.AllowedNssais) {
    return nil, false
	}
	return o.AllowedNssais, true
}

// HasAllowedNssais returns a boolean if a field has been set.
func (o *NFService) HasAllowedNssais() bool {
	if o != nil && !isNil(o.AllowedNssais) {
		return true
	}

	return false
}

// SetAllowedNssais gets a reference to the given []ExtSnssai and assigns it to the AllowedNssais field.
func (o *NFService) SetAllowedNssais(v []ExtSnssai) {
	o.AllowedNssais = v
}

// GetAllowedOperationsPerNfType returns the AllowedOperationsPerNfType field value if set, zero value otherwise.
func (o *NFService) GetAllowedOperationsPerNfType() map[string][]string {
	if o == nil || isNil(o.AllowedOperationsPerNfType) {
		var ret map[string][]string
		return ret
	}
	return *o.AllowedOperationsPerNfType
}

// GetAllowedOperationsPerNfTypeOk returns a tuple with the AllowedOperationsPerNfType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NFService) GetAllowedOperationsPerNfTypeOk() (*map[string][]string, bool) {
	if o == nil || isNil(o.AllowedOperationsPerNfType) {
    return nil, false
	}
	return o.AllowedOperationsPerNfType, true
}

// HasAllowedOperationsPerNfType returns a boolean if a field has been set.
func (o *NFService) HasAllowedOperationsPerNfType() bool {
	if o != nil && !isNil(o.AllowedOperationsPerNfType) {
		return true
	}

	return false
}

// SetAllowedOperationsPerNfType gets a reference to the given map[string][]string and assigns it to the AllowedOperationsPerNfType field.
func (o *NFService) SetAllowedOperationsPerNfType(v map[string][]string) {
	o.AllowedOperationsPerNfType = &v
}

// GetAllowedOperationsPerNfInstance returns the AllowedOperationsPerNfInstance field value if set, zero value otherwise.
func (o *NFService) GetAllowedOperationsPerNfInstance() map[string][]string {
	if o == nil || isNil(o.AllowedOperationsPerNfInstance) {
		var ret map[string][]string
		return ret
	}
	return *o.AllowedOperationsPerNfInstance
}

// GetAllowedOperationsPerNfInstanceOk returns a tuple with the AllowedOperationsPerNfInstance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NFService) GetAllowedOperationsPerNfInstanceOk() (*map[string][]string, bool) {
	if o == nil || isNil(o.AllowedOperationsPerNfInstance) {
    return nil, false
	}
	return o.AllowedOperationsPerNfInstance, true
}

// HasAllowedOperationsPerNfInstance returns a boolean if a field has been set.
func (o *NFService) HasAllowedOperationsPerNfInstance() bool {
	if o != nil && !isNil(o.AllowedOperationsPerNfInstance) {
		return true
	}

	return false
}

// SetAllowedOperationsPerNfInstance gets a reference to the given map[string][]string and assigns it to the AllowedOperationsPerNfInstance field.
func (o *NFService) SetAllowedOperationsPerNfInstance(v map[string][]string) {
	o.AllowedOperationsPerNfInstance = &v
}

// GetAllowedOperationsPerNfInstanceOverrides returns the AllowedOperationsPerNfInstanceOverrides field value if set, zero value otherwise.
func (o *NFService) GetAllowedOperationsPerNfInstanceOverrides() bool {
	if o == nil || isNil(o.AllowedOperationsPerNfInstanceOverrides) {
		var ret bool
		return ret
	}
	return *o.AllowedOperationsPerNfInstanceOverrides
}

// GetAllowedOperationsPerNfInstanceOverridesOk returns a tuple with the AllowedOperationsPerNfInstanceOverrides field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NFService) GetAllowedOperationsPerNfInstanceOverridesOk() (*bool, bool) {
	if o == nil || isNil(o.AllowedOperationsPerNfInstanceOverrides) {
    return nil, false
	}
	return o.AllowedOperationsPerNfInstanceOverrides, true
}

// HasAllowedOperationsPerNfInstanceOverrides returns a boolean if a field has been set.
func (o *NFService) HasAllowedOperationsPerNfInstanceOverrides() bool {
	if o != nil && !isNil(o.AllowedOperationsPerNfInstanceOverrides) {
		return true
	}

	return false
}

// SetAllowedOperationsPerNfInstanceOverrides gets a reference to the given bool and assigns it to the AllowedOperationsPerNfInstanceOverrides field.
func (o *NFService) SetAllowedOperationsPerNfInstanceOverrides(v bool) {
	o.AllowedOperationsPerNfInstanceOverrides = &v
}

// GetPriority returns the Priority field value if set, zero value otherwise.
func (o *NFService) GetPriority() int32 {
	if o == nil || isNil(o.Priority) {
		var ret int32
		return ret
	}
	return *o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NFService) GetPriorityOk() (*int32, bool) {
	if o == nil || isNil(o.Priority) {
    return nil, false
	}
	return o.Priority, true
}

// HasPriority returns a boolean if a field has been set.
func (o *NFService) HasPriority() bool {
	if o != nil && !isNil(o.Priority) {
		return true
	}

	return false
}

// SetPriority gets a reference to the given int32 and assigns it to the Priority field.
func (o *NFService) SetPriority(v int32) {
	o.Priority = &v
}

// GetCapacity returns the Capacity field value if set, zero value otherwise.
func (o *NFService) GetCapacity() int32 {
	if o == nil || isNil(o.Capacity) {
		var ret int32
		return ret
	}
	return *o.Capacity
}

// GetCapacityOk returns a tuple with the Capacity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NFService) GetCapacityOk() (*int32, bool) {
	if o == nil || isNil(o.Capacity) {
    return nil, false
	}
	return o.Capacity, true
}

// HasCapacity returns a boolean if a field has been set.
func (o *NFService) HasCapacity() bool {
	if o != nil && !isNil(o.Capacity) {
		return true
	}

	return false
}

// SetCapacity gets a reference to the given int32 and assigns it to the Capacity field.
func (o *NFService) SetCapacity(v int32) {
	o.Capacity = &v
}

// GetLoad returns the Load field value if set, zero value otherwise.
func (o *NFService) GetLoad() int32 {
	if o == nil || isNil(o.Load) {
		var ret int32
		return ret
	}
	return *o.Load
}

// GetLoadOk returns a tuple with the Load field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NFService) GetLoadOk() (*int32, bool) {
	if o == nil || isNil(o.Load) {
    return nil, false
	}
	return o.Load, true
}

// HasLoad returns a boolean if a field has been set.
func (o *NFService) HasLoad() bool {
	if o != nil && !isNil(o.Load) {
		return true
	}

	return false
}

// SetLoad gets a reference to the given int32 and assigns it to the Load field.
func (o *NFService) SetLoad(v int32) {
	o.Load = &v
}

// GetLoadTimeStamp returns the LoadTimeStamp field value if set, zero value otherwise.
func (o *NFService) GetLoadTimeStamp() time.Time {
	if o == nil || isNil(o.LoadTimeStamp) {
		var ret time.Time
		return ret
	}
	return *o.LoadTimeStamp
}

// GetLoadTimeStampOk returns a tuple with the LoadTimeStamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NFService) GetLoadTimeStampOk() (*time.Time, bool) {
	if o == nil || isNil(o.LoadTimeStamp) {
    return nil, false
	}
	return o.LoadTimeStamp, true
}

// HasLoadTimeStamp returns a boolean if a field has been set.
func (o *NFService) HasLoadTimeStamp() bool {
	if o != nil && !isNil(o.LoadTimeStamp) {
		return true
	}

	return false
}

// SetLoadTimeStamp gets a reference to the given time.Time and assigns it to the LoadTimeStamp field.
func (o *NFService) SetLoadTimeStamp(v time.Time) {
	o.LoadTimeStamp = &v
}

// GetRecoveryTime returns the RecoveryTime field value if set, zero value otherwise.
func (o *NFService) GetRecoveryTime() time.Time {
	if o == nil || isNil(o.RecoveryTime) {
		var ret time.Time
		return ret
	}
	return *o.RecoveryTime
}

// GetRecoveryTimeOk returns a tuple with the RecoveryTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NFService) GetRecoveryTimeOk() (*time.Time, bool) {
	if o == nil || isNil(o.RecoveryTime) {
    return nil, false
	}
	return o.RecoveryTime, true
}

// HasRecoveryTime returns a boolean if a field has been set.
func (o *NFService) HasRecoveryTime() bool {
	if o != nil && !isNil(o.RecoveryTime) {
		return true
	}

	return false
}

// SetRecoveryTime gets a reference to the given time.Time and assigns it to the RecoveryTime field.
func (o *NFService) SetRecoveryTime(v time.Time) {
	o.RecoveryTime = &v
}

// GetSupportedFeatures returns the SupportedFeatures field value if set, zero value otherwise.
func (o *NFService) GetSupportedFeatures() string {
	if o == nil || isNil(o.SupportedFeatures) {
		var ret string
		return ret
	}
	return *o.SupportedFeatures
}

// GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NFService) GetSupportedFeaturesOk() (*string, bool) {
	if o == nil || isNil(o.SupportedFeatures) {
    return nil, false
	}
	return o.SupportedFeatures, true
}

// HasSupportedFeatures returns a boolean if a field has been set.
func (o *NFService) HasSupportedFeatures() bool {
	if o != nil && !isNil(o.SupportedFeatures) {
		return true
	}

	return false
}

// SetSupportedFeatures gets a reference to the given string and assigns it to the SupportedFeatures field.
func (o *NFService) SetSupportedFeatures(v string) {
	o.SupportedFeatures = &v
}

// GetNfServiceSetIdList returns the NfServiceSetIdList field value if set, zero value otherwise.
func (o *NFService) GetNfServiceSetIdList() []string {
	if o == nil || isNil(o.NfServiceSetIdList) {
		var ret []string
		return ret
	}
	return o.NfServiceSetIdList
}

// GetNfServiceSetIdListOk returns a tuple with the NfServiceSetIdList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NFService) GetNfServiceSetIdListOk() ([]string, bool) {
	if o == nil || isNil(o.NfServiceSetIdList) {
    return nil, false
	}
	return o.NfServiceSetIdList, true
}

// HasNfServiceSetIdList returns a boolean if a field has been set.
func (o *NFService) HasNfServiceSetIdList() bool {
	if o != nil && !isNil(o.NfServiceSetIdList) {
		return true
	}

	return false
}

// SetNfServiceSetIdList gets a reference to the given []string and assigns it to the NfServiceSetIdList field.
func (o *NFService) SetNfServiceSetIdList(v []string) {
	o.NfServiceSetIdList = v
}

// GetSNssais returns the SNssais field value if set, zero value otherwise.
func (o *NFService) GetSNssais() []ExtSnssai {
	if o == nil || isNil(o.SNssais) {
		var ret []ExtSnssai
		return ret
	}
	return o.SNssais
}

// GetSNssaisOk returns a tuple with the SNssais field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NFService) GetSNssaisOk() ([]ExtSnssai, bool) {
	if o == nil || isNil(o.SNssais) {
    return nil, false
	}
	return o.SNssais, true
}

// HasSNssais returns a boolean if a field has been set.
func (o *NFService) HasSNssais() bool {
	if o != nil && !isNil(o.SNssais) {
		return true
	}

	return false
}

// SetSNssais gets a reference to the given []ExtSnssai and assigns it to the SNssais field.
func (o *NFService) SetSNssais(v []ExtSnssai) {
	o.SNssais = v
}

// GetPerPlmnSnssaiList returns the PerPlmnSnssaiList field value if set, zero value otherwise.
func (o *NFService) GetPerPlmnSnssaiList() []PlmnSnssai {
	if o == nil || isNil(o.PerPlmnSnssaiList) {
		var ret []PlmnSnssai
		return ret
	}
	return o.PerPlmnSnssaiList
}

// GetPerPlmnSnssaiListOk returns a tuple with the PerPlmnSnssaiList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NFService) GetPerPlmnSnssaiListOk() ([]PlmnSnssai, bool) {
	if o == nil || isNil(o.PerPlmnSnssaiList) {
    return nil, false
	}
	return o.PerPlmnSnssaiList, true
}

// HasPerPlmnSnssaiList returns a boolean if a field has been set.
func (o *NFService) HasPerPlmnSnssaiList() bool {
	if o != nil && !isNil(o.PerPlmnSnssaiList) {
		return true
	}

	return false
}

// SetPerPlmnSnssaiList gets a reference to the given []PlmnSnssai and assigns it to the PerPlmnSnssaiList field.
func (o *NFService) SetPerPlmnSnssaiList(v []PlmnSnssai) {
	o.PerPlmnSnssaiList = v
}

// GetVendorId returns the VendorId field value if set, zero value otherwise.
func (o *NFService) GetVendorId() string {
	if o == nil || isNil(o.VendorId) {
		var ret string
		return ret
	}
	return *o.VendorId
}

// GetVendorIdOk returns a tuple with the VendorId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NFService) GetVendorIdOk() (*string, bool) {
	if o == nil || isNil(o.VendorId) {
    return nil, false
	}
	return o.VendorId, true
}

// HasVendorId returns a boolean if a field has been set.
func (o *NFService) HasVendorId() bool {
	if o != nil && !isNil(o.VendorId) {
		return true
	}

	return false
}

// SetVendorId gets a reference to the given string and assigns it to the VendorId field.
func (o *NFService) SetVendorId(v string) {
	o.VendorId = &v
}

// GetSupportedVendorSpecificFeatures returns the SupportedVendorSpecificFeatures field value if set, zero value otherwise.
func (o *NFService) GetSupportedVendorSpecificFeatures() map[string][]VendorSpecificFeature {
	if o == nil || isNil(o.SupportedVendorSpecificFeatures) {
		var ret map[string][]VendorSpecificFeature
		return ret
	}
	return *o.SupportedVendorSpecificFeatures
}

// GetSupportedVendorSpecificFeaturesOk returns a tuple with the SupportedVendorSpecificFeatures field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NFService) GetSupportedVendorSpecificFeaturesOk() (*map[string][]VendorSpecificFeature, bool) {
	if o == nil || isNil(o.SupportedVendorSpecificFeatures) {
    return nil, false
	}
	return o.SupportedVendorSpecificFeatures, true
}

// HasSupportedVendorSpecificFeatures returns a boolean if a field has been set.
func (o *NFService) HasSupportedVendorSpecificFeatures() bool {
	if o != nil && !isNil(o.SupportedVendorSpecificFeatures) {
		return true
	}

	return false
}

// SetSupportedVendorSpecificFeatures gets a reference to the given map[string][]VendorSpecificFeature and assigns it to the SupportedVendorSpecificFeatures field.
func (o *NFService) SetSupportedVendorSpecificFeatures(v map[string][]VendorSpecificFeature) {
	o.SupportedVendorSpecificFeatures = &v
}

// GetOauth2Required returns the Oauth2Required field value if set, zero value otherwise.
func (o *NFService) GetOauth2Required() bool {
	if o == nil || isNil(o.Oauth2Required) {
		var ret bool
		return ret
	}
	return *o.Oauth2Required
}

// GetOauth2RequiredOk returns a tuple with the Oauth2Required field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NFService) GetOauth2RequiredOk() (*bool, bool) {
	if o == nil || isNil(o.Oauth2Required) {
    return nil, false
	}
	return o.Oauth2Required, true
}

// HasOauth2Required returns a boolean if a field has been set.
func (o *NFService) HasOauth2Required() bool {
	if o != nil && !isNil(o.Oauth2Required) {
		return true
	}

	return false
}

// SetOauth2Required gets a reference to the given bool and assigns it to the Oauth2Required field.
func (o *NFService) SetOauth2Required(v bool) {
	o.Oauth2Required = &v
}

// GetPerPlmnOauth2ReqList returns the PerPlmnOauth2ReqList field value if set, zero value otherwise.
func (o *NFService) GetPerPlmnOauth2ReqList() PlmnOauth2 {
	if o == nil || isNil(o.PerPlmnOauth2ReqList) {
		var ret PlmnOauth2
		return ret
	}
	return *o.PerPlmnOauth2ReqList
}

// GetPerPlmnOauth2ReqListOk returns a tuple with the PerPlmnOauth2ReqList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NFService) GetPerPlmnOauth2ReqListOk() (*PlmnOauth2, bool) {
	if o == nil || isNil(o.PerPlmnOauth2ReqList) {
    return nil, false
	}
	return o.PerPlmnOauth2ReqList, true
}

// HasPerPlmnOauth2ReqList returns a boolean if a field has been set.
func (o *NFService) HasPerPlmnOauth2ReqList() bool {
	if o != nil && !isNil(o.PerPlmnOauth2ReqList) {
		return true
	}

	return false
}

// SetPerPlmnOauth2ReqList gets a reference to the given PlmnOauth2 and assigns it to the PerPlmnOauth2ReqList field.
func (o *NFService) SetPerPlmnOauth2ReqList(v PlmnOauth2) {
	o.PerPlmnOauth2ReqList = &v
}

func (o NFService) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["serviceInstanceId"] = o.ServiceInstanceId
	}
	if true {
		toSerialize["serviceName"] = o.ServiceName
	}
	if true {
		toSerialize["versions"] = o.Versions
	}
	if true {
		toSerialize["scheme"] = o.Scheme
	}
	if true {
		toSerialize["nfServiceStatus"] = o.NfServiceStatus
	}
	if !isNil(o.Fqdn) {
		toSerialize["fqdn"] = o.Fqdn
	}
	if !isNil(o.InterPlmnFqdn) {
		toSerialize["interPlmnFqdn"] = o.InterPlmnFqdn
	}
	if !isNil(o.IpEndPoints) {
		toSerialize["ipEndPoints"] = o.IpEndPoints
	}
	if !isNil(o.ApiPrefix) {
		toSerialize["apiPrefix"] = o.ApiPrefix
	}
	if !isNil(o.DefaultNotificationSubscriptions) {
		toSerialize["defaultNotificationSubscriptions"] = o.DefaultNotificationSubscriptions
	}
	if !isNil(o.AllowedPlmns) {
		toSerialize["allowedPlmns"] = o.AllowedPlmns
	}
	if !isNil(o.AllowedSnpns) {
		toSerialize["allowedSnpns"] = o.AllowedSnpns
	}
	if !isNil(o.AllowedNfTypes) {
		toSerialize["allowedNfTypes"] = o.AllowedNfTypes
	}
	if !isNil(o.AllowedNfDomains) {
		toSerialize["allowedNfDomains"] = o.AllowedNfDomains
	}
	if !isNil(o.AllowedNssais) {
		toSerialize["allowedNssais"] = o.AllowedNssais
	}
	if !isNil(o.AllowedOperationsPerNfType) {
		toSerialize["allowedOperationsPerNfType"] = o.AllowedOperationsPerNfType
	}
	if !isNil(o.AllowedOperationsPerNfInstance) {
		toSerialize["allowedOperationsPerNfInstance"] = o.AllowedOperationsPerNfInstance
	}
	if !isNil(o.AllowedOperationsPerNfInstanceOverrides) {
		toSerialize["allowedOperationsPerNfInstanceOverrides"] = o.AllowedOperationsPerNfInstanceOverrides
	}
	if !isNil(o.Priority) {
		toSerialize["priority"] = o.Priority
	}
	if !isNil(o.Capacity) {
		toSerialize["capacity"] = o.Capacity
	}
	if !isNil(o.Load) {
		toSerialize["load"] = o.Load
	}
	if !isNil(o.LoadTimeStamp) {
		toSerialize["loadTimeStamp"] = o.LoadTimeStamp
	}
	if !isNil(o.RecoveryTime) {
		toSerialize["recoveryTime"] = o.RecoveryTime
	}
	if !isNil(o.SupportedFeatures) {
		toSerialize["supportedFeatures"] = o.SupportedFeatures
	}
	if !isNil(o.NfServiceSetIdList) {
		toSerialize["nfServiceSetIdList"] = o.NfServiceSetIdList
	}
	if !isNil(o.SNssais) {
		toSerialize["sNssais"] = o.SNssais
	}
	if !isNil(o.PerPlmnSnssaiList) {
		toSerialize["perPlmnSnssaiList"] = o.PerPlmnSnssaiList
	}
	if !isNil(o.VendorId) {
		toSerialize["vendorId"] = o.VendorId
	}
	if !isNil(o.SupportedVendorSpecificFeatures) {
		toSerialize["supportedVendorSpecificFeatures"] = o.SupportedVendorSpecificFeatures
	}
	if !isNil(o.Oauth2Required) {
		toSerialize["oauth2Required"] = o.Oauth2Required
	}
	if !isNil(o.PerPlmnOauth2ReqList) {
		toSerialize["perPlmnOauth2ReqList"] = o.PerPlmnOauth2ReqList
	}
	return json.Marshal(toSerialize)
}

type NullableNFService struct {
	value *NFService
	isSet bool
}

func (v NullableNFService) Get() *NFService {
	return v.value
}

func (v *NullableNFService) Set(val *NFService) {
	v.value = val
	v.isSet = true
}

func (v NullableNFService) IsSet() bool {
	return v.isSet
}

func (v *NullableNFService) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNFService(val *NFService) *NullableNFService {
	return &NullableNFService{value: val, isSet: true}
}

func (v NullableNFService) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNFService) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


