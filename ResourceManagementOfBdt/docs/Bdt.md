# Bdt

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Self** | Pointer to **string** | string formatted according to IETF RFC 3986 identifying a referenced resource. | [optional] 
**SupportedFeatures** | Pointer to **string** | A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \&quot;0\&quot; to \&quot;9\&quot;,  \&quot;a\&quot; to \&quot;f\&quot; or \&quot;A\&quot; to \&quot;F\&quot; and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.  | [optional] 
**VolumePerUE** | [**UsageThreshold**](UsageThreshold.md) |  | 
**NumberOfUEs** | **int32** | Identifies the number of UEs. | 
**DesiredTimeWindow** | [**TimeWindow**](TimeWindow.md) |  | 
**LocationArea** | Pointer to [**LocationArea**](LocationArea.md) |  | [optional] 
**LocationArea5G** | Pointer to [**LocationArea5G**](LocationArea5G.md) |  | [optional] 
**ReferenceId** | Pointer to **string** | string identifying a BDT Reference ID as defined in clause 5.3.3 of 3GPP TS 29.154. | [optional] 
**TransferPolicies** | Pointer to [**[]TransferPolicy**](TransferPolicy.md) | Identifies an offered transfer policy. | [optional] [readonly] 
**SelectedPolicy** | Pointer to **int32** | Identity of the selected background data transfer policy. Shall not be present in initial message exchange, can be provided by NF service consumer in a subsequent message exchange. | [optional] 
**ExternalGroupId** | Pointer to **string** | string containing a local identifier followed by \&quot;@\&quot; and a domain identifier. Both the local identifier and the domain identifier shall be encoded as strings that do not contain any \&quot;@\&quot; characters. See Clauses 4.6.2 and 4.6.3 of 3GPP TS 23.682 for more information. | [optional] 
**NotificationDestination** | Pointer to **string** | string formatted according to IETF RFC 3986 identifying a referenced resource. | [optional] 
**WarnNotifEnabled** | Pointer to **bool** | Indicates whether the BDT warning notification is enabled (true) or not (false). Default value is false.  | [optional] 
**TrafficDes** | Pointer to **string** | Identify a traffic descriptor as defined in Figure 5.2.2 of 3GPP TS 24.526, octets v+5 to w. | [optional] 

## Methods

### NewBdt

`func NewBdt(volumePerUE UsageThreshold, numberOfUEs int32, desiredTimeWindow TimeWindow, ) *Bdt`

NewBdt instantiates a new Bdt object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBdtWithDefaults

`func NewBdtWithDefaults() *Bdt`

NewBdtWithDefaults instantiates a new Bdt object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSelf

`func (o *Bdt) GetSelf() string`

GetSelf returns the Self field if non-nil, zero value otherwise.

### GetSelfOk

`func (o *Bdt) GetSelfOk() (*string, bool)`

GetSelfOk returns a tuple with the Self field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelf

`func (o *Bdt) SetSelf(v string)`

SetSelf sets Self field to given value.

### HasSelf

`func (o *Bdt) HasSelf() bool`

HasSelf returns a boolean if a field has been set.

### GetSupportedFeatures

`func (o *Bdt) GetSupportedFeatures() string`

GetSupportedFeatures returns the SupportedFeatures field if non-nil, zero value otherwise.

### GetSupportedFeaturesOk

`func (o *Bdt) GetSupportedFeaturesOk() (*string, bool)`

GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedFeatures

`func (o *Bdt) SetSupportedFeatures(v string)`

SetSupportedFeatures sets SupportedFeatures field to given value.

### HasSupportedFeatures

`func (o *Bdt) HasSupportedFeatures() bool`

HasSupportedFeatures returns a boolean if a field has been set.

### GetVolumePerUE

`func (o *Bdt) GetVolumePerUE() UsageThreshold`

GetVolumePerUE returns the VolumePerUE field if non-nil, zero value otherwise.

### GetVolumePerUEOk

`func (o *Bdt) GetVolumePerUEOk() (*UsageThreshold, bool)`

GetVolumePerUEOk returns a tuple with the VolumePerUE field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumePerUE

`func (o *Bdt) SetVolumePerUE(v UsageThreshold)`

SetVolumePerUE sets VolumePerUE field to given value.


### GetNumberOfUEs

`func (o *Bdt) GetNumberOfUEs() int32`

GetNumberOfUEs returns the NumberOfUEs field if non-nil, zero value otherwise.

### GetNumberOfUEsOk

`func (o *Bdt) GetNumberOfUEsOk() (*int32, bool)`

GetNumberOfUEsOk returns a tuple with the NumberOfUEs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfUEs

`func (o *Bdt) SetNumberOfUEs(v int32)`

SetNumberOfUEs sets NumberOfUEs field to given value.


### GetDesiredTimeWindow

`func (o *Bdt) GetDesiredTimeWindow() TimeWindow`

GetDesiredTimeWindow returns the DesiredTimeWindow field if non-nil, zero value otherwise.

### GetDesiredTimeWindowOk

`func (o *Bdt) GetDesiredTimeWindowOk() (*TimeWindow, bool)`

GetDesiredTimeWindowOk returns a tuple with the DesiredTimeWindow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesiredTimeWindow

`func (o *Bdt) SetDesiredTimeWindow(v TimeWindow)`

SetDesiredTimeWindow sets DesiredTimeWindow field to given value.


### GetLocationArea

`func (o *Bdt) GetLocationArea() LocationArea`

GetLocationArea returns the LocationArea field if non-nil, zero value otherwise.

### GetLocationAreaOk

`func (o *Bdt) GetLocationAreaOk() (*LocationArea, bool)`

GetLocationAreaOk returns a tuple with the LocationArea field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationArea

`func (o *Bdt) SetLocationArea(v LocationArea)`

SetLocationArea sets LocationArea field to given value.

### HasLocationArea

`func (o *Bdt) HasLocationArea() bool`

HasLocationArea returns a boolean if a field has been set.

### GetLocationArea5G

`func (o *Bdt) GetLocationArea5G() LocationArea5G`

GetLocationArea5G returns the LocationArea5G field if non-nil, zero value otherwise.

### GetLocationArea5GOk

`func (o *Bdt) GetLocationArea5GOk() (*LocationArea5G, bool)`

GetLocationArea5GOk returns a tuple with the LocationArea5G field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationArea5G

`func (o *Bdt) SetLocationArea5G(v LocationArea5G)`

SetLocationArea5G sets LocationArea5G field to given value.

### HasLocationArea5G

`func (o *Bdt) HasLocationArea5G() bool`

HasLocationArea5G returns a boolean if a field has been set.

### GetReferenceId

`func (o *Bdt) GetReferenceId() string`

GetReferenceId returns the ReferenceId field if non-nil, zero value otherwise.

### GetReferenceIdOk

`func (o *Bdt) GetReferenceIdOk() (*string, bool)`

GetReferenceIdOk returns a tuple with the ReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceId

`func (o *Bdt) SetReferenceId(v string)`

SetReferenceId sets ReferenceId field to given value.

### HasReferenceId

`func (o *Bdt) HasReferenceId() bool`

HasReferenceId returns a boolean if a field has been set.

### GetTransferPolicies

`func (o *Bdt) GetTransferPolicies() []TransferPolicy`

GetTransferPolicies returns the TransferPolicies field if non-nil, zero value otherwise.

### GetTransferPoliciesOk

`func (o *Bdt) GetTransferPoliciesOk() (*[]TransferPolicy, bool)`

GetTransferPoliciesOk returns a tuple with the TransferPolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransferPolicies

`func (o *Bdt) SetTransferPolicies(v []TransferPolicy)`

SetTransferPolicies sets TransferPolicies field to given value.

### HasTransferPolicies

`func (o *Bdt) HasTransferPolicies() bool`

HasTransferPolicies returns a boolean if a field has been set.

### GetSelectedPolicy

`func (o *Bdt) GetSelectedPolicy() int32`

GetSelectedPolicy returns the SelectedPolicy field if non-nil, zero value otherwise.

### GetSelectedPolicyOk

`func (o *Bdt) GetSelectedPolicyOk() (*int32, bool)`

GetSelectedPolicyOk returns a tuple with the SelectedPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectedPolicy

`func (o *Bdt) SetSelectedPolicy(v int32)`

SetSelectedPolicy sets SelectedPolicy field to given value.

### HasSelectedPolicy

`func (o *Bdt) HasSelectedPolicy() bool`

HasSelectedPolicy returns a boolean if a field has been set.

### GetExternalGroupId

`func (o *Bdt) GetExternalGroupId() string`

GetExternalGroupId returns the ExternalGroupId field if non-nil, zero value otherwise.

### GetExternalGroupIdOk

`func (o *Bdt) GetExternalGroupIdOk() (*string, bool)`

GetExternalGroupIdOk returns a tuple with the ExternalGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalGroupId

`func (o *Bdt) SetExternalGroupId(v string)`

SetExternalGroupId sets ExternalGroupId field to given value.

### HasExternalGroupId

`func (o *Bdt) HasExternalGroupId() bool`

HasExternalGroupId returns a boolean if a field has been set.

### GetNotificationDestination

`func (o *Bdt) GetNotificationDestination() string`

GetNotificationDestination returns the NotificationDestination field if non-nil, zero value otherwise.

### GetNotificationDestinationOk

`func (o *Bdt) GetNotificationDestinationOk() (*string, bool)`

GetNotificationDestinationOk returns a tuple with the NotificationDestination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationDestination

`func (o *Bdt) SetNotificationDestination(v string)`

SetNotificationDestination sets NotificationDestination field to given value.

### HasNotificationDestination

`func (o *Bdt) HasNotificationDestination() bool`

HasNotificationDestination returns a boolean if a field has been set.

### GetWarnNotifEnabled

`func (o *Bdt) GetWarnNotifEnabled() bool`

GetWarnNotifEnabled returns the WarnNotifEnabled field if non-nil, zero value otherwise.

### GetWarnNotifEnabledOk

`func (o *Bdt) GetWarnNotifEnabledOk() (*bool, bool)`

GetWarnNotifEnabledOk returns a tuple with the WarnNotifEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnNotifEnabled

`func (o *Bdt) SetWarnNotifEnabled(v bool)`

SetWarnNotifEnabled sets WarnNotifEnabled field to given value.

### HasWarnNotifEnabled

`func (o *Bdt) HasWarnNotifEnabled() bool`

HasWarnNotifEnabled returns a boolean if a field has been set.

### GetTrafficDes

`func (o *Bdt) GetTrafficDes() string`

GetTrafficDes returns the TrafficDes field if non-nil, zero value otherwise.

### GetTrafficDesOk

`func (o *Bdt) GetTrafficDesOk() (*string, bool)`

GetTrafficDesOk returns a tuple with the TrafficDes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrafficDes

`func (o *Bdt) SetTrafficDes(v string)`

SetTrafficDes sets TrafficDes field to given value.

### HasTrafficDes

`func (o *Bdt) HasTrafficDes() bool`

HasTrafficDes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


