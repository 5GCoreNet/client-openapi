# SubscriptionData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NfStatusNotificationUri** | **string** |  | 
**ReqNfInstanceId** | Pointer to **string** | String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.   | [optional] 
**SubscrCond** | Pointer to [**SubscrCond**](SubscrCond.md) |  | [optional] 
**SubscriptionId** | **string** |  | [readonly] 
**ValidityTime** | Pointer to **time.Time** | string with format &#39;date-time&#39; as defined in OpenAPI. | [optional] 
**ReqNotifEvents** | Pointer to [**[]NotificationEventType**](NotificationEventType.md) |  | [optional] 
**PlmnId** | Pointer to [**PlmnId**](PlmnId.md) |  | [optional] 
**Nid** | Pointer to **string** | This represents the Network Identifier, which together with a PLMN ID is used to identify an SNPN (see 3GPP TS 23.003 and 3GPP TS 23.501 clause 5.30.2.1).   | [optional] 
**NotifCondition** | Pointer to [**NotifCondition**](NotifCondition.md) |  | [optional] 
**ReqNfType** | Pointer to [**NFType**](NFType.md) |  | [optional] 
**ReqNfFqdn** | Pointer to **string** | Fully Qualified Domain Name | [optional] 
**ReqSnssais** | Pointer to [**[]ExtSnssai**](ExtSnssai.md) |  | [optional] 
**ReqPerPlmnSnssais** | Pointer to [**[]PlmnSnssai**](PlmnSnssai.md) |  | [optional] 
**ReqPlmnList** | Pointer to [**[]PlmnId**](PlmnId.md) |  | [optional] 
**ReqSnpnList** | Pointer to [**[]PlmnIdNid**](PlmnIdNid.md) |  | [optional] 
**ServingScope** | Pointer to **[]string** |  | [optional] 
**RequesterFeatures** | Pointer to **string** |  | [optional] 
**NrfSupportedFeatures** | Pointer to **string** |  | [optional] [readonly] 
**HnrfUri** | Pointer to **string** | String providing an URI formatted according to RFC 3986. | [optional] 
**OnboardingCapability** | Pointer to **bool** |  | [optional] [default to false]
**TargetHni** | Pointer to **string** | Fully Qualified Domain Name | [optional] 
**PreferredLocality** | Pointer to **string** |  | [optional] 
**ExtPreferredLocality** | Pointer to [**map[string][]LocalityDescription**](array.md) | A map (list of key-value pairs) where the key of the map represents the relative priority, for the requester, of each locality description among the list of locality descriptions in this query parameter, encoded as \&quot;1\&quot; (highest priority\&quot;), \&quot;2\&quot;, \&quot;3\&quot;, …,  \&quot;n\&quot; (lowest priority)  | [optional] 

## Methods

### NewSubscriptionData

`func NewSubscriptionData(nfStatusNotificationUri string, subscriptionId string, ) *SubscriptionData`

NewSubscriptionData instantiates a new SubscriptionData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriptionDataWithDefaults

`func NewSubscriptionDataWithDefaults() *SubscriptionData`

NewSubscriptionDataWithDefaults instantiates a new SubscriptionData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNfStatusNotificationUri

`func (o *SubscriptionData) GetNfStatusNotificationUri() string`

GetNfStatusNotificationUri returns the NfStatusNotificationUri field if non-nil, zero value otherwise.

### GetNfStatusNotificationUriOk

`func (o *SubscriptionData) GetNfStatusNotificationUriOk() (*string, bool)`

GetNfStatusNotificationUriOk returns a tuple with the NfStatusNotificationUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNfStatusNotificationUri

`func (o *SubscriptionData) SetNfStatusNotificationUri(v string)`

SetNfStatusNotificationUri sets NfStatusNotificationUri field to given value.


### GetReqNfInstanceId

`func (o *SubscriptionData) GetReqNfInstanceId() string`

GetReqNfInstanceId returns the ReqNfInstanceId field if non-nil, zero value otherwise.

### GetReqNfInstanceIdOk

`func (o *SubscriptionData) GetReqNfInstanceIdOk() (*string, bool)`

GetReqNfInstanceIdOk returns a tuple with the ReqNfInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReqNfInstanceId

`func (o *SubscriptionData) SetReqNfInstanceId(v string)`

SetReqNfInstanceId sets ReqNfInstanceId field to given value.

### HasReqNfInstanceId

`func (o *SubscriptionData) HasReqNfInstanceId() bool`

HasReqNfInstanceId returns a boolean if a field has been set.

### GetSubscrCond

`func (o *SubscriptionData) GetSubscrCond() SubscrCond`

GetSubscrCond returns the SubscrCond field if non-nil, zero value otherwise.

### GetSubscrCondOk

`func (o *SubscriptionData) GetSubscrCondOk() (*SubscrCond, bool)`

GetSubscrCondOk returns a tuple with the SubscrCond field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscrCond

`func (o *SubscriptionData) SetSubscrCond(v SubscrCond)`

SetSubscrCond sets SubscrCond field to given value.

### HasSubscrCond

`func (o *SubscriptionData) HasSubscrCond() bool`

HasSubscrCond returns a boolean if a field has been set.

### GetSubscriptionId

`func (o *SubscriptionData) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *SubscriptionData) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *SubscriptionData) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.


### GetValidityTime

`func (o *SubscriptionData) GetValidityTime() time.Time`

GetValidityTime returns the ValidityTime field if non-nil, zero value otherwise.

### GetValidityTimeOk

`func (o *SubscriptionData) GetValidityTimeOk() (*time.Time, bool)`

GetValidityTimeOk returns a tuple with the ValidityTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidityTime

`func (o *SubscriptionData) SetValidityTime(v time.Time)`

SetValidityTime sets ValidityTime field to given value.

### HasValidityTime

`func (o *SubscriptionData) HasValidityTime() bool`

HasValidityTime returns a boolean if a field has been set.

### GetReqNotifEvents

`func (o *SubscriptionData) GetReqNotifEvents() []NotificationEventType`

GetReqNotifEvents returns the ReqNotifEvents field if non-nil, zero value otherwise.

### GetReqNotifEventsOk

`func (o *SubscriptionData) GetReqNotifEventsOk() (*[]NotificationEventType, bool)`

GetReqNotifEventsOk returns a tuple with the ReqNotifEvents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReqNotifEvents

`func (o *SubscriptionData) SetReqNotifEvents(v []NotificationEventType)`

SetReqNotifEvents sets ReqNotifEvents field to given value.

### HasReqNotifEvents

`func (o *SubscriptionData) HasReqNotifEvents() bool`

HasReqNotifEvents returns a boolean if a field has been set.

### GetPlmnId

`func (o *SubscriptionData) GetPlmnId() PlmnId`

GetPlmnId returns the PlmnId field if non-nil, zero value otherwise.

### GetPlmnIdOk

`func (o *SubscriptionData) GetPlmnIdOk() (*PlmnId, bool)`

GetPlmnIdOk returns a tuple with the PlmnId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlmnId

`func (o *SubscriptionData) SetPlmnId(v PlmnId)`

SetPlmnId sets PlmnId field to given value.

### HasPlmnId

`func (o *SubscriptionData) HasPlmnId() bool`

HasPlmnId returns a boolean if a field has been set.

### GetNid

`func (o *SubscriptionData) GetNid() string`

GetNid returns the Nid field if non-nil, zero value otherwise.

### GetNidOk

`func (o *SubscriptionData) GetNidOk() (*string, bool)`

GetNidOk returns a tuple with the Nid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNid

`func (o *SubscriptionData) SetNid(v string)`

SetNid sets Nid field to given value.

### HasNid

`func (o *SubscriptionData) HasNid() bool`

HasNid returns a boolean if a field has been set.

### GetNotifCondition

`func (o *SubscriptionData) GetNotifCondition() NotifCondition`

GetNotifCondition returns the NotifCondition field if non-nil, zero value otherwise.

### GetNotifConditionOk

`func (o *SubscriptionData) GetNotifConditionOk() (*NotifCondition, bool)`

GetNotifConditionOk returns a tuple with the NotifCondition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifCondition

`func (o *SubscriptionData) SetNotifCondition(v NotifCondition)`

SetNotifCondition sets NotifCondition field to given value.

### HasNotifCondition

`func (o *SubscriptionData) HasNotifCondition() bool`

HasNotifCondition returns a boolean if a field has been set.

### GetReqNfType

`func (o *SubscriptionData) GetReqNfType() NFType`

GetReqNfType returns the ReqNfType field if non-nil, zero value otherwise.

### GetReqNfTypeOk

`func (o *SubscriptionData) GetReqNfTypeOk() (*NFType, bool)`

GetReqNfTypeOk returns a tuple with the ReqNfType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReqNfType

`func (o *SubscriptionData) SetReqNfType(v NFType)`

SetReqNfType sets ReqNfType field to given value.

### HasReqNfType

`func (o *SubscriptionData) HasReqNfType() bool`

HasReqNfType returns a boolean if a field has been set.

### GetReqNfFqdn

`func (o *SubscriptionData) GetReqNfFqdn() string`

GetReqNfFqdn returns the ReqNfFqdn field if non-nil, zero value otherwise.

### GetReqNfFqdnOk

`func (o *SubscriptionData) GetReqNfFqdnOk() (*string, bool)`

GetReqNfFqdnOk returns a tuple with the ReqNfFqdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReqNfFqdn

`func (o *SubscriptionData) SetReqNfFqdn(v string)`

SetReqNfFqdn sets ReqNfFqdn field to given value.

### HasReqNfFqdn

`func (o *SubscriptionData) HasReqNfFqdn() bool`

HasReqNfFqdn returns a boolean if a field has been set.

### GetReqSnssais

`func (o *SubscriptionData) GetReqSnssais() []ExtSnssai`

GetReqSnssais returns the ReqSnssais field if non-nil, zero value otherwise.

### GetReqSnssaisOk

`func (o *SubscriptionData) GetReqSnssaisOk() (*[]ExtSnssai, bool)`

GetReqSnssaisOk returns a tuple with the ReqSnssais field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReqSnssais

`func (o *SubscriptionData) SetReqSnssais(v []ExtSnssai)`

SetReqSnssais sets ReqSnssais field to given value.

### HasReqSnssais

`func (o *SubscriptionData) HasReqSnssais() bool`

HasReqSnssais returns a boolean if a field has been set.

### GetReqPerPlmnSnssais

`func (o *SubscriptionData) GetReqPerPlmnSnssais() []PlmnSnssai`

GetReqPerPlmnSnssais returns the ReqPerPlmnSnssais field if non-nil, zero value otherwise.

### GetReqPerPlmnSnssaisOk

`func (o *SubscriptionData) GetReqPerPlmnSnssaisOk() (*[]PlmnSnssai, bool)`

GetReqPerPlmnSnssaisOk returns a tuple with the ReqPerPlmnSnssais field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReqPerPlmnSnssais

`func (o *SubscriptionData) SetReqPerPlmnSnssais(v []PlmnSnssai)`

SetReqPerPlmnSnssais sets ReqPerPlmnSnssais field to given value.

### HasReqPerPlmnSnssais

`func (o *SubscriptionData) HasReqPerPlmnSnssais() bool`

HasReqPerPlmnSnssais returns a boolean if a field has been set.

### GetReqPlmnList

`func (o *SubscriptionData) GetReqPlmnList() []PlmnId`

GetReqPlmnList returns the ReqPlmnList field if non-nil, zero value otherwise.

### GetReqPlmnListOk

`func (o *SubscriptionData) GetReqPlmnListOk() (*[]PlmnId, bool)`

GetReqPlmnListOk returns a tuple with the ReqPlmnList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReqPlmnList

`func (o *SubscriptionData) SetReqPlmnList(v []PlmnId)`

SetReqPlmnList sets ReqPlmnList field to given value.

### HasReqPlmnList

`func (o *SubscriptionData) HasReqPlmnList() bool`

HasReqPlmnList returns a boolean if a field has been set.

### GetReqSnpnList

`func (o *SubscriptionData) GetReqSnpnList() []PlmnIdNid`

GetReqSnpnList returns the ReqSnpnList field if non-nil, zero value otherwise.

### GetReqSnpnListOk

`func (o *SubscriptionData) GetReqSnpnListOk() (*[]PlmnIdNid, bool)`

GetReqSnpnListOk returns a tuple with the ReqSnpnList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReqSnpnList

`func (o *SubscriptionData) SetReqSnpnList(v []PlmnIdNid)`

SetReqSnpnList sets ReqSnpnList field to given value.

### HasReqSnpnList

`func (o *SubscriptionData) HasReqSnpnList() bool`

HasReqSnpnList returns a boolean if a field has been set.

### GetServingScope

`func (o *SubscriptionData) GetServingScope() []string`

GetServingScope returns the ServingScope field if non-nil, zero value otherwise.

### GetServingScopeOk

`func (o *SubscriptionData) GetServingScopeOk() (*[]string, bool)`

GetServingScopeOk returns a tuple with the ServingScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServingScope

`func (o *SubscriptionData) SetServingScope(v []string)`

SetServingScope sets ServingScope field to given value.

### HasServingScope

`func (o *SubscriptionData) HasServingScope() bool`

HasServingScope returns a boolean if a field has been set.

### GetRequesterFeatures

`func (o *SubscriptionData) GetRequesterFeatures() string`

GetRequesterFeatures returns the RequesterFeatures field if non-nil, zero value otherwise.

### GetRequesterFeaturesOk

`func (o *SubscriptionData) GetRequesterFeaturesOk() (*string, bool)`

GetRequesterFeaturesOk returns a tuple with the RequesterFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequesterFeatures

`func (o *SubscriptionData) SetRequesterFeatures(v string)`

SetRequesterFeatures sets RequesterFeatures field to given value.

### HasRequesterFeatures

`func (o *SubscriptionData) HasRequesterFeatures() bool`

HasRequesterFeatures returns a boolean if a field has been set.

### GetNrfSupportedFeatures

`func (o *SubscriptionData) GetNrfSupportedFeatures() string`

GetNrfSupportedFeatures returns the NrfSupportedFeatures field if non-nil, zero value otherwise.

### GetNrfSupportedFeaturesOk

`func (o *SubscriptionData) GetNrfSupportedFeaturesOk() (*string, bool)`

GetNrfSupportedFeaturesOk returns a tuple with the NrfSupportedFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNrfSupportedFeatures

`func (o *SubscriptionData) SetNrfSupportedFeatures(v string)`

SetNrfSupportedFeatures sets NrfSupportedFeatures field to given value.

### HasNrfSupportedFeatures

`func (o *SubscriptionData) HasNrfSupportedFeatures() bool`

HasNrfSupportedFeatures returns a boolean if a field has been set.

### GetHnrfUri

`func (o *SubscriptionData) GetHnrfUri() string`

GetHnrfUri returns the HnrfUri field if non-nil, zero value otherwise.

### GetHnrfUriOk

`func (o *SubscriptionData) GetHnrfUriOk() (*string, bool)`

GetHnrfUriOk returns a tuple with the HnrfUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHnrfUri

`func (o *SubscriptionData) SetHnrfUri(v string)`

SetHnrfUri sets HnrfUri field to given value.

### HasHnrfUri

`func (o *SubscriptionData) HasHnrfUri() bool`

HasHnrfUri returns a boolean if a field has been set.

### GetOnboardingCapability

`func (o *SubscriptionData) GetOnboardingCapability() bool`

GetOnboardingCapability returns the OnboardingCapability field if non-nil, zero value otherwise.

### GetOnboardingCapabilityOk

`func (o *SubscriptionData) GetOnboardingCapabilityOk() (*bool, bool)`

GetOnboardingCapabilityOk returns a tuple with the OnboardingCapability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnboardingCapability

`func (o *SubscriptionData) SetOnboardingCapability(v bool)`

SetOnboardingCapability sets OnboardingCapability field to given value.

### HasOnboardingCapability

`func (o *SubscriptionData) HasOnboardingCapability() bool`

HasOnboardingCapability returns a boolean if a field has been set.

### GetTargetHni

`func (o *SubscriptionData) GetTargetHni() string`

GetTargetHni returns the TargetHni field if non-nil, zero value otherwise.

### GetTargetHniOk

`func (o *SubscriptionData) GetTargetHniOk() (*string, bool)`

GetTargetHniOk returns a tuple with the TargetHni field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetHni

`func (o *SubscriptionData) SetTargetHni(v string)`

SetTargetHni sets TargetHni field to given value.

### HasTargetHni

`func (o *SubscriptionData) HasTargetHni() bool`

HasTargetHni returns a boolean if a field has been set.

### GetPreferredLocality

`func (o *SubscriptionData) GetPreferredLocality() string`

GetPreferredLocality returns the PreferredLocality field if non-nil, zero value otherwise.

### GetPreferredLocalityOk

`func (o *SubscriptionData) GetPreferredLocalityOk() (*string, bool)`

GetPreferredLocalityOk returns a tuple with the PreferredLocality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferredLocality

`func (o *SubscriptionData) SetPreferredLocality(v string)`

SetPreferredLocality sets PreferredLocality field to given value.

### HasPreferredLocality

`func (o *SubscriptionData) HasPreferredLocality() bool`

HasPreferredLocality returns a boolean if a field has been set.

### GetExtPreferredLocality

`func (o *SubscriptionData) GetExtPreferredLocality() map[string][]LocalityDescription`

GetExtPreferredLocality returns the ExtPreferredLocality field if non-nil, zero value otherwise.

### GetExtPreferredLocalityOk

`func (o *SubscriptionData) GetExtPreferredLocalityOk() (*map[string][]LocalityDescription, bool)`

GetExtPreferredLocalityOk returns a tuple with the ExtPreferredLocality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtPreferredLocality

`func (o *SubscriptionData) SetExtPreferredLocality(v map[string][]LocalityDescription)`

SetExtPreferredLocality sets ExtPreferredLocality field to given value.

### HasExtPreferredLocality

`func (o *SubscriptionData) HasExtPreferredLocality() bool`

HasExtPreferredLocality returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


