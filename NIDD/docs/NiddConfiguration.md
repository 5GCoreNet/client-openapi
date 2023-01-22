# NiddConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Self** | Pointer to **string** | string formatted according to IETF RFC 3986 identifying a referenced resource. | [optional] 
**SupportedFeatures** | Pointer to **string** | A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \&quot;0\&quot; to \&quot;9\&quot;,  \&quot;a\&quot; to \&quot;f\&quot; or \&quot;A\&quot; to \&quot;F\&quot; and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.  | [optional] 
**MtcProviderId** | Pointer to **string** | Identifies the MTC Service Provider and/or MTC Application. | [optional] 
**ExternalId** | Pointer to **string** | string containing a local identifier followed by \&quot;@\&quot; and a domain identifier. Both the local identifier and the domain identifier shall be encoded as strings that do not contain any \&quot;@\&quot; characters. See Clause 4.6.2 of 3GPP TS 23.682 for more information. | [optional] 
**Msisdn** | Pointer to **string** | string formatted according to clause 3.3 of 3GPP TS 23.003 that describes an MSISDN. | [optional] 
**ExternalGroupId** | Pointer to **string** | string containing a local identifier followed by \&quot;@\&quot; and a domain identifier. Both the local identifier and the domain identifier shall be encoded as strings that do not contain any \&quot;@\&quot; characters. See Clauses 4.6.2 and 4.6.3 of 3GPP TS 23.682 for more information. | [optional] 
**Duration** | Pointer to **time.Time** | string with format \&quot;date-time\&quot; as defined in OpenAPI. | [optional] 
**ReliableDataService** | Pointer to **bool** | Indicates whether the reliable data service (as defined in clause 4.5.14.3 of 3GPP TS  23.682) acknowledgement is requested (true) or not (false). Default value is false.  | [optional] 
**RdsPorts** | Pointer to [**[]RdsPort**](RdsPort.md) | Indicates the static port configuration that is used for reliable data transfer between specific applications using RDS (as defined in clause 5.2.4 and 5.2.5 of 3GPP TS 24.250). | [optional] 
**PdnEstablishmentOption** | Pointer to [**PdnEstablishmentOptions**](PdnEstablishmentOptions.md) |  | [optional] 
**NotificationDestination** | **string** | string formatted according to IETF RFC 3986 identifying a referenced resource. | 
**RequestTestNotification** | Pointer to **bool** | Set to true by the SCS/AS to request the SCEF to send a test notification as defined in clause 5.2.5.3. Set to false or omitted otherwise. | [optional] 
**WebsockNotifConfig** | Pointer to [**WebsockNotifConfig**](WebsockNotifConfig.md) |  | [optional] 
**MaximumPacketSize** | Pointer to **int32** | The Maximum Packet Size is the maximum NIDD packet size that was transferred to the UE by the SCEF in the PCO, see clause 4.5.14.1 of 3GPP TS 23.682. If no maximum packet size was provided to the UE by the SCEF, the SCEF sends a default configured max packet size to SCS/AS. Unit  bit. | [optional] [readonly] 
**NiddDownlinkDataTransfers** | Pointer to [**[]NiddDownlinkDataTransfer**](NiddDownlinkDataTransfer.md) | The downlink data deliveries that needed to be executed by the SCEF. The cardinality of the property shall be 0..1 in the request and 0..N in the response (i.e. response may contain multiple buffered MT NIDD). | [optional] 
**Status** | Pointer to [**NiddStatus**](NiddStatus.md) |  | [optional] 

## Methods

### NewNiddConfiguration

`func NewNiddConfiguration(notificationDestination string, ) *NiddConfiguration`

NewNiddConfiguration instantiates a new NiddConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNiddConfigurationWithDefaults

`func NewNiddConfigurationWithDefaults() *NiddConfiguration`

NewNiddConfigurationWithDefaults instantiates a new NiddConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSelf

`func (o *NiddConfiguration) GetSelf() string`

GetSelf returns the Self field if non-nil, zero value otherwise.

### GetSelfOk

`func (o *NiddConfiguration) GetSelfOk() (*string, bool)`

GetSelfOk returns a tuple with the Self field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelf

`func (o *NiddConfiguration) SetSelf(v string)`

SetSelf sets Self field to given value.

### HasSelf

`func (o *NiddConfiguration) HasSelf() bool`

HasSelf returns a boolean if a field has been set.

### GetSupportedFeatures

`func (o *NiddConfiguration) GetSupportedFeatures() string`

GetSupportedFeatures returns the SupportedFeatures field if non-nil, zero value otherwise.

### GetSupportedFeaturesOk

`func (o *NiddConfiguration) GetSupportedFeaturesOk() (*string, bool)`

GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedFeatures

`func (o *NiddConfiguration) SetSupportedFeatures(v string)`

SetSupportedFeatures sets SupportedFeatures field to given value.

### HasSupportedFeatures

`func (o *NiddConfiguration) HasSupportedFeatures() bool`

HasSupportedFeatures returns a boolean if a field has been set.

### GetMtcProviderId

`func (o *NiddConfiguration) GetMtcProviderId() string`

GetMtcProviderId returns the MtcProviderId field if non-nil, zero value otherwise.

### GetMtcProviderIdOk

`func (o *NiddConfiguration) GetMtcProviderIdOk() (*string, bool)`

GetMtcProviderIdOk returns a tuple with the MtcProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtcProviderId

`func (o *NiddConfiguration) SetMtcProviderId(v string)`

SetMtcProviderId sets MtcProviderId field to given value.

### HasMtcProviderId

`func (o *NiddConfiguration) HasMtcProviderId() bool`

HasMtcProviderId returns a boolean if a field has been set.

### GetExternalId

`func (o *NiddConfiguration) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *NiddConfiguration) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *NiddConfiguration) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *NiddConfiguration) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetMsisdn

`func (o *NiddConfiguration) GetMsisdn() string`

GetMsisdn returns the Msisdn field if non-nil, zero value otherwise.

### GetMsisdnOk

`func (o *NiddConfiguration) GetMsisdnOk() (*string, bool)`

GetMsisdnOk returns a tuple with the Msisdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsisdn

`func (o *NiddConfiguration) SetMsisdn(v string)`

SetMsisdn sets Msisdn field to given value.

### HasMsisdn

`func (o *NiddConfiguration) HasMsisdn() bool`

HasMsisdn returns a boolean if a field has been set.

### GetExternalGroupId

`func (o *NiddConfiguration) GetExternalGroupId() string`

GetExternalGroupId returns the ExternalGroupId field if non-nil, zero value otherwise.

### GetExternalGroupIdOk

`func (o *NiddConfiguration) GetExternalGroupIdOk() (*string, bool)`

GetExternalGroupIdOk returns a tuple with the ExternalGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalGroupId

`func (o *NiddConfiguration) SetExternalGroupId(v string)`

SetExternalGroupId sets ExternalGroupId field to given value.

### HasExternalGroupId

`func (o *NiddConfiguration) HasExternalGroupId() bool`

HasExternalGroupId returns a boolean if a field has been set.

### GetDuration

`func (o *NiddConfiguration) GetDuration() time.Time`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *NiddConfiguration) GetDurationOk() (*time.Time, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *NiddConfiguration) SetDuration(v time.Time)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *NiddConfiguration) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetReliableDataService

`func (o *NiddConfiguration) GetReliableDataService() bool`

GetReliableDataService returns the ReliableDataService field if non-nil, zero value otherwise.

### GetReliableDataServiceOk

`func (o *NiddConfiguration) GetReliableDataServiceOk() (*bool, bool)`

GetReliableDataServiceOk returns a tuple with the ReliableDataService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReliableDataService

`func (o *NiddConfiguration) SetReliableDataService(v bool)`

SetReliableDataService sets ReliableDataService field to given value.

### HasReliableDataService

`func (o *NiddConfiguration) HasReliableDataService() bool`

HasReliableDataService returns a boolean if a field has been set.

### GetRdsPorts

`func (o *NiddConfiguration) GetRdsPorts() []RdsPort`

GetRdsPorts returns the RdsPorts field if non-nil, zero value otherwise.

### GetRdsPortsOk

`func (o *NiddConfiguration) GetRdsPortsOk() (*[]RdsPort, bool)`

GetRdsPortsOk returns a tuple with the RdsPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRdsPorts

`func (o *NiddConfiguration) SetRdsPorts(v []RdsPort)`

SetRdsPorts sets RdsPorts field to given value.

### HasRdsPorts

`func (o *NiddConfiguration) HasRdsPorts() bool`

HasRdsPorts returns a boolean if a field has been set.

### GetPdnEstablishmentOption

`func (o *NiddConfiguration) GetPdnEstablishmentOption() PdnEstablishmentOptions`

GetPdnEstablishmentOption returns the PdnEstablishmentOption field if non-nil, zero value otherwise.

### GetPdnEstablishmentOptionOk

`func (o *NiddConfiguration) GetPdnEstablishmentOptionOk() (*PdnEstablishmentOptions, bool)`

GetPdnEstablishmentOptionOk returns a tuple with the PdnEstablishmentOption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPdnEstablishmentOption

`func (o *NiddConfiguration) SetPdnEstablishmentOption(v PdnEstablishmentOptions)`

SetPdnEstablishmentOption sets PdnEstablishmentOption field to given value.

### HasPdnEstablishmentOption

`func (o *NiddConfiguration) HasPdnEstablishmentOption() bool`

HasPdnEstablishmentOption returns a boolean if a field has been set.

### GetNotificationDestination

`func (o *NiddConfiguration) GetNotificationDestination() string`

GetNotificationDestination returns the NotificationDestination field if non-nil, zero value otherwise.

### GetNotificationDestinationOk

`func (o *NiddConfiguration) GetNotificationDestinationOk() (*string, bool)`

GetNotificationDestinationOk returns a tuple with the NotificationDestination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationDestination

`func (o *NiddConfiguration) SetNotificationDestination(v string)`

SetNotificationDestination sets NotificationDestination field to given value.


### GetRequestTestNotification

`func (o *NiddConfiguration) GetRequestTestNotification() bool`

GetRequestTestNotification returns the RequestTestNotification field if non-nil, zero value otherwise.

### GetRequestTestNotificationOk

`func (o *NiddConfiguration) GetRequestTestNotificationOk() (*bool, bool)`

GetRequestTestNotificationOk returns a tuple with the RequestTestNotification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestTestNotification

`func (o *NiddConfiguration) SetRequestTestNotification(v bool)`

SetRequestTestNotification sets RequestTestNotification field to given value.

### HasRequestTestNotification

`func (o *NiddConfiguration) HasRequestTestNotification() bool`

HasRequestTestNotification returns a boolean if a field has been set.

### GetWebsockNotifConfig

`func (o *NiddConfiguration) GetWebsockNotifConfig() WebsockNotifConfig`

GetWebsockNotifConfig returns the WebsockNotifConfig field if non-nil, zero value otherwise.

### GetWebsockNotifConfigOk

`func (o *NiddConfiguration) GetWebsockNotifConfigOk() (*WebsockNotifConfig, bool)`

GetWebsockNotifConfigOk returns a tuple with the WebsockNotifConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsockNotifConfig

`func (o *NiddConfiguration) SetWebsockNotifConfig(v WebsockNotifConfig)`

SetWebsockNotifConfig sets WebsockNotifConfig field to given value.

### HasWebsockNotifConfig

`func (o *NiddConfiguration) HasWebsockNotifConfig() bool`

HasWebsockNotifConfig returns a boolean if a field has been set.

### GetMaximumPacketSize

`func (o *NiddConfiguration) GetMaximumPacketSize() int32`

GetMaximumPacketSize returns the MaximumPacketSize field if non-nil, zero value otherwise.

### GetMaximumPacketSizeOk

`func (o *NiddConfiguration) GetMaximumPacketSizeOk() (*int32, bool)`

GetMaximumPacketSizeOk returns a tuple with the MaximumPacketSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumPacketSize

`func (o *NiddConfiguration) SetMaximumPacketSize(v int32)`

SetMaximumPacketSize sets MaximumPacketSize field to given value.

### HasMaximumPacketSize

`func (o *NiddConfiguration) HasMaximumPacketSize() bool`

HasMaximumPacketSize returns a boolean if a field has been set.

### GetNiddDownlinkDataTransfers

`func (o *NiddConfiguration) GetNiddDownlinkDataTransfers() []NiddDownlinkDataTransfer`

GetNiddDownlinkDataTransfers returns the NiddDownlinkDataTransfers field if non-nil, zero value otherwise.

### GetNiddDownlinkDataTransfersOk

`func (o *NiddConfiguration) GetNiddDownlinkDataTransfersOk() (*[]NiddDownlinkDataTransfer, bool)`

GetNiddDownlinkDataTransfersOk returns a tuple with the NiddDownlinkDataTransfers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNiddDownlinkDataTransfers

`func (o *NiddConfiguration) SetNiddDownlinkDataTransfers(v []NiddDownlinkDataTransfer)`

SetNiddDownlinkDataTransfers sets NiddDownlinkDataTransfers field to given value.

### HasNiddDownlinkDataTransfers

`func (o *NiddConfiguration) HasNiddDownlinkDataTransfers() bool`

HasNiddDownlinkDataTransfers returns a boolean if a field has been set.

### GetStatus

`func (o *NiddConfiguration) GetStatus() NiddStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *NiddConfiguration) GetStatusOk() (*NiddStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *NiddConfiguration) SetStatus(v NiddStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *NiddConfiguration) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


