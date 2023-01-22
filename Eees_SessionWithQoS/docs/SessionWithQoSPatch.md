# SessionWithQoSPatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IpFlows** | Pointer to **[]string** | Contains the flow description for the Uplink and/or Downlink IP flows. | [optional] 
**QosReference** | Pointer to **string** | Identifies a pre-defined QoS information. | [optional] 
**AltQosReference** | Pointer to **[]string** | Identifies an ordered list of pre-defined QoS information. The lower the index of the array for a given entry, the higher the priority.  | [optional] 
**Events** | Pointer to [**[]UserPlaneEvent**](UserPlaneEvent.md) | Indicates the events subscribed by the EAS. | [optional] 
**SponsorInformation** | Pointer to [**SponsorInformation**](SponsorInformation.md) |  | [optional] 
**QosMonInfo** | Pointer to [**QosMonitoringInformationRm**](QosMonitoringInformationRm.md) |  | [optional] 
**NotificationDestination** | Pointer to **string** | string providing an URI formatted according to IETF RFC 3986. | [optional] 
**MaxbrUl** | Pointer to **NullableString** | This data type is defined in the same way as the &#39;BitRate&#39; data type, but with the OpenAPI &#39;nullable: true&#39; property.  | [optional] 
**MaxbrDl** | Pointer to **NullableString** | This data type is defined in the same way as the &#39;BitRate&#39; data type, but with the OpenAPI &#39;nullable: true&#39; property.  | [optional] 
**DisUeNotif** | Pointer to **bool** |  | [optional] 

## Methods

### NewSessionWithQoSPatch

`func NewSessionWithQoSPatch() *SessionWithQoSPatch`

NewSessionWithQoSPatch instantiates a new SessionWithQoSPatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSessionWithQoSPatchWithDefaults

`func NewSessionWithQoSPatchWithDefaults() *SessionWithQoSPatch`

NewSessionWithQoSPatchWithDefaults instantiates a new SessionWithQoSPatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIpFlows

`func (o *SessionWithQoSPatch) GetIpFlows() []string`

GetIpFlows returns the IpFlows field if non-nil, zero value otherwise.

### GetIpFlowsOk

`func (o *SessionWithQoSPatch) GetIpFlowsOk() (*[]string, bool)`

GetIpFlowsOk returns a tuple with the IpFlows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpFlows

`func (o *SessionWithQoSPatch) SetIpFlows(v []string)`

SetIpFlows sets IpFlows field to given value.

### HasIpFlows

`func (o *SessionWithQoSPatch) HasIpFlows() bool`

HasIpFlows returns a boolean if a field has been set.

### GetQosReference

`func (o *SessionWithQoSPatch) GetQosReference() string`

GetQosReference returns the QosReference field if non-nil, zero value otherwise.

### GetQosReferenceOk

`func (o *SessionWithQoSPatch) GetQosReferenceOk() (*string, bool)`

GetQosReferenceOk returns a tuple with the QosReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQosReference

`func (o *SessionWithQoSPatch) SetQosReference(v string)`

SetQosReference sets QosReference field to given value.

### HasQosReference

`func (o *SessionWithQoSPatch) HasQosReference() bool`

HasQosReference returns a boolean if a field has been set.

### GetAltQosReference

`func (o *SessionWithQoSPatch) GetAltQosReference() []string`

GetAltQosReference returns the AltQosReference field if non-nil, zero value otherwise.

### GetAltQosReferenceOk

`func (o *SessionWithQoSPatch) GetAltQosReferenceOk() (*[]string, bool)`

GetAltQosReferenceOk returns a tuple with the AltQosReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAltQosReference

`func (o *SessionWithQoSPatch) SetAltQosReference(v []string)`

SetAltQosReference sets AltQosReference field to given value.

### HasAltQosReference

`func (o *SessionWithQoSPatch) HasAltQosReference() bool`

HasAltQosReference returns a boolean if a field has been set.

### GetEvents

`func (o *SessionWithQoSPatch) GetEvents() []UserPlaneEvent`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *SessionWithQoSPatch) GetEventsOk() (*[]UserPlaneEvent, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *SessionWithQoSPatch) SetEvents(v []UserPlaneEvent)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *SessionWithQoSPatch) HasEvents() bool`

HasEvents returns a boolean if a field has been set.

### GetSponsorInformation

`func (o *SessionWithQoSPatch) GetSponsorInformation() SponsorInformation`

GetSponsorInformation returns the SponsorInformation field if non-nil, zero value otherwise.

### GetSponsorInformationOk

`func (o *SessionWithQoSPatch) GetSponsorInformationOk() (*SponsorInformation, bool)`

GetSponsorInformationOk returns a tuple with the SponsorInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSponsorInformation

`func (o *SessionWithQoSPatch) SetSponsorInformation(v SponsorInformation)`

SetSponsorInformation sets SponsorInformation field to given value.

### HasSponsorInformation

`func (o *SessionWithQoSPatch) HasSponsorInformation() bool`

HasSponsorInformation returns a boolean if a field has been set.

### GetQosMonInfo

`func (o *SessionWithQoSPatch) GetQosMonInfo() QosMonitoringInformationRm`

GetQosMonInfo returns the QosMonInfo field if non-nil, zero value otherwise.

### GetQosMonInfoOk

`func (o *SessionWithQoSPatch) GetQosMonInfoOk() (*QosMonitoringInformationRm, bool)`

GetQosMonInfoOk returns a tuple with the QosMonInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQosMonInfo

`func (o *SessionWithQoSPatch) SetQosMonInfo(v QosMonitoringInformationRm)`

SetQosMonInfo sets QosMonInfo field to given value.

### HasQosMonInfo

`func (o *SessionWithQoSPatch) HasQosMonInfo() bool`

HasQosMonInfo returns a boolean if a field has been set.

### GetNotificationDestination

`func (o *SessionWithQoSPatch) GetNotificationDestination() string`

GetNotificationDestination returns the NotificationDestination field if non-nil, zero value otherwise.

### GetNotificationDestinationOk

`func (o *SessionWithQoSPatch) GetNotificationDestinationOk() (*string, bool)`

GetNotificationDestinationOk returns a tuple with the NotificationDestination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationDestination

`func (o *SessionWithQoSPatch) SetNotificationDestination(v string)`

SetNotificationDestination sets NotificationDestination field to given value.

### HasNotificationDestination

`func (o *SessionWithQoSPatch) HasNotificationDestination() bool`

HasNotificationDestination returns a boolean if a field has been set.

### GetMaxbrUl

`func (o *SessionWithQoSPatch) GetMaxbrUl() string`

GetMaxbrUl returns the MaxbrUl field if non-nil, zero value otherwise.

### GetMaxbrUlOk

`func (o *SessionWithQoSPatch) GetMaxbrUlOk() (*string, bool)`

GetMaxbrUlOk returns a tuple with the MaxbrUl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxbrUl

`func (o *SessionWithQoSPatch) SetMaxbrUl(v string)`

SetMaxbrUl sets MaxbrUl field to given value.

### HasMaxbrUl

`func (o *SessionWithQoSPatch) HasMaxbrUl() bool`

HasMaxbrUl returns a boolean if a field has been set.

### SetMaxbrUlNil

`func (o *SessionWithQoSPatch) SetMaxbrUlNil(b bool)`

 SetMaxbrUlNil sets the value for MaxbrUl to be an explicit nil

### UnsetMaxbrUl
`func (o *SessionWithQoSPatch) UnsetMaxbrUl()`

UnsetMaxbrUl ensures that no value is present for MaxbrUl, not even an explicit nil
### GetMaxbrDl

`func (o *SessionWithQoSPatch) GetMaxbrDl() string`

GetMaxbrDl returns the MaxbrDl field if non-nil, zero value otherwise.

### GetMaxbrDlOk

`func (o *SessionWithQoSPatch) GetMaxbrDlOk() (*string, bool)`

GetMaxbrDlOk returns a tuple with the MaxbrDl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxbrDl

`func (o *SessionWithQoSPatch) SetMaxbrDl(v string)`

SetMaxbrDl sets MaxbrDl field to given value.

### HasMaxbrDl

`func (o *SessionWithQoSPatch) HasMaxbrDl() bool`

HasMaxbrDl returns a boolean if a field has been set.

### SetMaxbrDlNil

`func (o *SessionWithQoSPatch) SetMaxbrDlNil(b bool)`

 SetMaxbrDlNil sets the value for MaxbrDl to be an explicit nil

### UnsetMaxbrDl
`func (o *SessionWithQoSPatch) UnsetMaxbrDl()`

UnsetMaxbrDl ensures that no value is present for MaxbrDl, not even an explicit nil
### GetDisUeNotif

`func (o *SessionWithQoSPatch) GetDisUeNotif() bool`

GetDisUeNotif returns the DisUeNotif field if non-nil, zero value otherwise.

### GetDisUeNotifOk

`func (o *SessionWithQoSPatch) GetDisUeNotifOk() (*bool, bool)`

GetDisUeNotifOk returns a tuple with the DisUeNotif field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisUeNotif

`func (o *SessionWithQoSPatch) SetDisUeNotif(v bool)`

SetDisUeNotif sets DisUeNotif field to given value.

### HasDisUeNotif

`func (o *SessionWithQoSPatch) HasDisUeNotif() bool`

HasDisUeNotif returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


