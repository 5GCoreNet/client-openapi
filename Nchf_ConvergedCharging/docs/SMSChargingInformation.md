# SMSChargingInformation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OriginatorInfo** | Pointer to [**OriginatorInfo**](OriginatorInfo.md) |  | [optional] 
**RecipientInfo** | Pointer to [**[]RecipientInfo**](RecipientInfo.md) |  | [optional] 
**UserEquipmentInfo** | Pointer to **string** | String representing a Permanent Equipment Identifier that may contain - an IMEI or IMEISV, as  specified in clause 6.2 of 3GPP TS 23.003; a MAC address for a 5G-RG or FN-RG via  wireline  access, with an indication that this address cannot be trusted for regulatory purpose if this  address cannot be used as an Equipment Identifier of the FN-RG, as specified in clause 4.7.7  of 3GPP TS23.316. Examples are imei-012345678901234 or imeisv-0123456789012345.   | [optional] 
**RoamerInOut** | Pointer to [**RoamerInOut**](RoamerInOut.md) |  | [optional] 
**UserLocationinfo** | Pointer to [**UserLocation**](UserLocation.md) |  | [optional] 
**UetimeZone** | Pointer to **string** | String with format \&quot;time-numoffset\&quot; optionally appended by \&quot;daylightSavingTime\&quot;, where  - \&quot;time-numoffset\&quot; shall represent the time zone adjusted for daylight saving time and be    encoded as time-numoffset as defined in clause 5.6 of IETF RFC 3339;  - \&quot;daylightSavingTime\&quot; shall represent the adjustment that has been made and shall be    encoded as \&quot;+1\&quot; or \&quot;+2\&quot; for a +1 or +2 hours adjustment.   The example is for 8 hours behind UTC, +1 hour adjustment for Daylight Saving Time.  | [optional] 
**RATType** | Pointer to [**RatType**](RatType.md) |  | [optional] 
**SMSCAddress** | Pointer to **string** |  | [optional] 
**SMDataCodingScheme** | Pointer to **int32** |  | [optional] 
**SMMessageType** | Pointer to [**SMMessageType**](SMMessageType.md) |  | [optional] 
**SMReplyPathRequested** | Pointer to [**ReplyPathRequested**](ReplyPathRequested.md) |  | [optional] 
**SMUserDataHeader** | Pointer to **string** |  | [optional] 
**SMStatus** | Pointer to **string** |  | [optional] 
**SMDischargeTime** | Pointer to **time.Time** | string with format &#39;date-time&#39; as defined in OpenAPI. | [optional] 
**NumberofMessagesSent** | Pointer to **int32** | Integer where the allowed values correspond to the value range of an unsigned 32-bit integer.  | [optional] 
**SMServiceType** | Pointer to [**SMServiceType**](SMServiceType.md) |  | [optional] 
**SMSequenceNumber** | Pointer to **int32** | Integer where the allowed values correspond to the value range of an unsigned 32-bit integer.  | [optional] 
**SMSresult** | Pointer to **int32** | Integer where the allowed values correspond to the value range of an unsigned 32-bit integer.  | [optional] 
**SubmissionTime** | Pointer to **time.Time** | string with format &#39;date-time&#39; as defined in OpenAPI. | [optional] 
**SMPriority** | Pointer to [**SMPriority**](SMPriority.md) |  | [optional] 
**MessageReference** | Pointer to **string** |  | [optional] 
**MessageSize** | Pointer to **int32** | Integer where the allowed values correspond to the value range of an unsigned 32-bit integer.  | [optional] 
**MessageClass** | Pointer to [**MessageClass**](MessageClass.md) |  | [optional] 
**DeliveryReportRequested** | Pointer to [**DeliveryReportRequested**](DeliveryReportRequested.md) |  | [optional] 

## Methods

### NewSMSChargingInformation

`func NewSMSChargingInformation() *SMSChargingInformation`

NewSMSChargingInformation instantiates a new SMSChargingInformation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSMSChargingInformationWithDefaults

`func NewSMSChargingInformationWithDefaults() *SMSChargingInformation`

NewSMSChargingInformationWithDefaults instantiates a new SMSChargingInformation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOriginatorInfo

`func (o *SMSChargingInformation) GetOriginatorInfo() OriginatorInfo`

GetOriginatorInfo returns the OriginatorInfo field if non-nil, zero value otherwise.

### GetOriginatorInfoOk

`func (o *SMSChargingInformation) GetOriginatorInfoOk() (*OriginatorInfo, bool)`

GetOriginatorInfoOk returns a tuple with the OriginatorInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginatorInfo

`func (o *SMSChargingInformation) SetOriginatorInfo(v OriginatorInfo)`

SetOriginatorInfo sets OriginatorInfo field to given value.

### HasOriginatorInfo

`func (o *SMSChargingInformation) HasOriginatorInfo() bool`

HasOriginatorInfo returns a boolean if a field has been set.

### GetRecipientInfo

`func (o *SMSChargingInformation) GetRecipientInfo() []RecipientInfo`

GetRecipientInfo returns the RecipientInfo field if non-nil, zero value otherwise.

### GetRecipientInfoOk

`func (o *SMSChargingInformation) GetRecipientInfoOk() (*[]RecipientInfo, bool)`

GetRecipientInfoOk returns a tuple with the RecipientInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipientInfo

`func (o *SMSChargingInformation) SetRecipientInfo(v []RecipientInfo)`

SetRecipientInfo sets RecipientInfo field to given value.

### HasRecipientInfo

`func (o *SMSChargingInformation) HasRecipientInfo() bool`

HasRecipientInfo returns a boolean if a field has been set.

### GetUserEquipmentInfo

`func (o *SMSChargingInformation) GetUserEquipmentInfo() string`

GetUserEquipmentInfo returns the UserEquipmentInfo field if non-nil, zero value otherwise.

### GetUserEquipmentInfoOk

`func (o *SMSChargingInformation) GetUserEquipmentInfoOk() (*string, bool)`

GetUserEquipmentInfoOk returns a tuple with the UserEquipmentInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserEquipmentInfo

`func (o *SMSChargingInformation) SetUserEquipmentInfo(v string)`

SetUserEquipmentInfo sets UserEquipmentInfo field to given value.

### HasUserEquipmentInfo

`func (o *SMSChargingInformation) HasUserEquipmentInfo() bool`

HasUserEquipmentInfo returns a boolean if a field has been set.

### GetRoamerInOut

`func (o *SMSChargingInformation) GetRoamerInOut() RoamerInOut`

GetRoamerInOut returns the RoamerInOut field if non-nil, zero value otherwise.

### GetRoamerInOutOk

`func (o *SMSChargingInformation) GetRoamerInOutOk() (*RoamerInOut, bool)`

GetRoamerInOutOk returns a tuple with the RoamerInOut field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoamerInOut

`func (o *SMSChargingInformation) SetRoamerInOut(v RoamerInOut)`

SetRoamerInOut sets RoamerInOut field to given value.

### HasRoamerInOut

`func (o *SMSChargingInformation) HasRoamerInOut() bool`

HasRoamerInOut returns a boolean if a field has been set.

### GetUserLocationinfo

`func (o *SMSChargingInformation) GetUserLocationinfo() UserLocation`

GetUserLocationinfo returns the UserLocationinfo field if non-nil, zero value otherwise.

### GetUserLocationinfoOk

`func (o *SMSChargingInformation) GetUserLocationinfoOk() (*UserLocation, bool)`

GetUserLocationinfoOk returns a tuple with the UserLocationinfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserLocationinfo

`func (o *SMSChargingInformation) SetUserLocationinfo(v UserLocation)`

SetUserLocationinfo sets UserLocationinfo field to given value.

### HasUserLocationinfo

`func (o *SMSChargingInformation) HasUserLocationinfo() bool`

HasUserLocationinfo returns a boolean if a field has been set.

### GetUetimeZone

`func (o *SMSChargingInformation) GetUetimeZone() string`

GetUetimeZone returns the UetimeZone field if non-nil, zero value otherwise.

### GetUetimeZoneOk

`func (o *SMSChargingInformation) GetUetimeZoneOk() (*string, bool)`

GetUetimeZoneOk returns a tuple with the UetimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUetimeZone

`func (o *SMSChargingInformation) SetUetimeZone(v string)`

SetUetimeZone sets UetimeZone field to given value.

### HasUetimeZone

`func (o *SMSChargingInformation) HasUetimeZone() bool`

HasUetimeZone returns a boolean if a field has been set.

### GetRATType

`func (o *SMSChargingInformation) GetRATType() RatType`

GetRATType returns the RATType field if non-nil, zero value otherwise.

### GetRATTypeOk

`func (o *SMSChargingInformation) GetRATTypeOk() (*RatType, bool)`

GetRATTypeOk returns a tuple with the RATType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRATType

`func (o *SMSChargingInformation) SetRATType(v RatType)`

SetRATType sets RATType field to given value.

### HasRATType

`func (o *SMSChargingInformation) HasRATType() bool`

HasRATType returns a boolean if a field has been set.

### GetSMSCAddress

`func (o *SMSChargingInformation) GetSMSCAddress() string`

GetSMSCAddress returns the SMSCAddress field if non-nil, zero value otherwise.

### GetSMSCAddressOk

`func (o *SMSChargingInformation) GetSMSCAddressOk() (*string, bool)`

GetSMSCAddressOk returns a tuple with the SMSCAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSMSCAddress

`func (o *SMSChargingInformation) SetSMSCAddress(v string)`

SetSMSCAddress sets SMSCAddress field to given value.

### HasSMSCAddress

`func (o *SMSChargingInformation) HasSMSCAddress() bool`

HasSMSCAddress returns a boolean if a field has been set.

### GetSMDataCodingScheme

`func (o *SMSChargingInformation) GetSMDataCodingScheme() int32`

GetSMDataCodingScheme returns the SMDataCodingScheme field if non-nil, zero value otherwise.

### GetSMDataCodingSchemeOk

`func (o *SMSChargingInformation) GetSMDataCodingSchemeOk() (*int32, bool)`

GetSMDataCodingSchemeOk returns a tuple with the SMDataCodingScheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSMDataCodingScheme

`func (o *SMSChargingInformation) SetSMDataCodingScheme(v int32)`

SetSMDataCodingScheme sets SMDataCodingScheme field to given value.

### HasSMDataCodingScheme

`func (o *SMSChargingInformation) HasSMDataCodingScheme() bool`

HasSMDataCodingScheme returns a boolean if a field has been set.

### GetSMMessageType

`func (o *SMSChargingInformation) GetSMMessageType() SMMessageType`

GetSMMessageType returns the SMMessageType field if non-nil, zero value otherwise.

### GetSMMessageTypeOk

`func (o *SMSChargingInformation) GetSMMessageTypeOk() (*SMMessageType, bool)`

GetSMMessageTypeOk returns a tuple with the SMMessageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSMMessageType

`func (o *SMSChargingInformation) SetSMMessageType(v SMMessageType)`

SetSMMessageType sets SMMessageType field to given value.

### HasSMMessageType

`func (o *SMSChargingInformation) HasSMMessageType() bool`

HasSMMessageType returns a boolean if a field has been set.

### GetSMReplyPathRequested

`func (o *SMSChargingInformation) GetSMReplyPathRequested() ReplyPathRequested`

GetSMReplyPathRequested returns the SMReplyPathRequested field if non-nil, zero value otherwise.

### GetSMReplyPathRequestedOk

`func (o *SMSChargingInformation) GetSMReplyPathRequestedOk() (*ReplyPathRequested, bool)`

GetSMReplyPathRequestedOk returns a tuple with the SMReplyPathRequested field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSMReplyPathRequested

`func (o *SMSChargingInformation) SetSMReplyPathRequested(v ReplyPathRequested)`

SetSMReplyPathRequested sets SMReplyPathRequested field to given value.

### HasSMReplyPathRequested

`func (o *SMSChargingInformation) HasSMReplyPathRequested() bool`

HasSMReplyPathRequested returns a boolean if a field has been set.

### GetSMUserDataHeader

`func (o *SMSChargingInformation) GetSMUserDataHeader() string`

GetSMUserDataHeader returns the SMUserDataHeader field if non-nil, zero value otherwise.

### GetSMUserDataHeaderOk

`func (o *SMSChargingInformation) GetSMUserDataHeaderOk() (*string, bool)`

GetSMUserDataHeaderOk returns a tuple with the SMUserDataHeader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSMUserDataHeader

`func (o *SMSChargingInformation) SetSMUserDataHeader(v string)`

SetSMUserDataHeader sets SMUserDataHeader field to given value.

### HasSMUserDataHeader

`func (o *SMSChargingInformation) HasSMUserDataHeader() bool`

HasSMUserDataHeader returns a boolean if a field has been set.

### GetSMStatus

`func (o *SMSChargingInformation) GetSMStatus() string`

GetSMStatus returns the SMStatus field if non-nil, zero value otherwise.

### GetSMStatusOk

`func (o *SMSChargingInformation) GetSMStatusOk() (*string, bool)`

GetSMStatusOk returns a tuple with the SMStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSMStatus

`func (o *SMSChargingInformation) SetSMStatus(v string)`

SetSMStatus sets SMStatus field to given value.

### HasSMStatus

`func (o *SMSChargingInformation) HasSMStatus() bool`

HasSMStatus returns a boolean if a field has been set.

### GetSMDischargeTime

`func (o *SMSChargingInformation) GetSMDischargeTime() time.Time`

GetSMDischargeTime returns the SMDischargeTime field if non-nil, zero value otherwise.

### GetSMDischargeTimeOk

`func (o *SMSChargingInformation) GetSMDischargeTimeOk() (*time.Time, bool)`

GetSMDischargeTimeOk returns a tuple with the SMDischargeTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSMDischargeTime

`func (o *SMSChargingInformation) SetSMDischargeTime(v time.Time)`

SetSMDischargeTime sets SMDischargeTime field to given value.

### HasSMDischargeTime

`func (o *SMSChargingInformation) HasSMDischargeTime() bool`

HasSMDischargeTime returns a boolean if a field has been set.

### GetNumberofMessagesSent

`func (o *SMSChargingInformation) GetNumberofMessagesSent() int32`

GetNumberofMessagesSent returns the NumberofMessagesSent field if non-nil, zero value otherwise.

### GetNumberofMessagesSentOk

`func (o *SMSChargingInformation) GetNumberofMessagesSentOk() (*int32, bool)`

GetNumberofMessagesSentOk returns a tuple with the NumberofMessagesSent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberofMessagesSent

`func (o *SMSChargingInformation) SetNumberofMessagesSent(v int32)`

SetNumberofMessagesSent sets NumberofMessagesSent field to given value.

### HasNumberofMessagesSent

`func (o *SMSChargingInformation) HasNumberofMessagesSent() bool`

HasNumberofMessagesSent returns a boolean if a field has been set.

### GetSMServiceType

`func (o *SMSChargingInformation) GetSMServiceType() SMServiceType`

GetSMServiceType returns the SMServiceType field if non-nil, zero value otherwise.

### GetSMServiceTypeOk

`func (o *SMSChargingInformation) GetSMServiceTypeOk() (*SMServiceType, bool)`

GetSMServiceTypeOk returns a tuple with the SMServiceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSMServiceType

`func (o *SMSChargingInformation) SetSMServiceType(v SMServiceType)`

SetSMServiceType sets SMServiceType field to given value.

### HasSMServiceType

`func (o *SMSChargingInformation) HasSMServiceType() bool`

HasSMServiceType returns a boolean if a field has been set.

### GetSMSequenceNumber

`func (o *SMSChargingInformation) GetSMSequenceNumber() int32`

GetSMSequenceNumber returns the SMSequenceNumber field if non-nil, zero value otherwise.

### GetSMSequenceNumberOk

`func (o *SMSChargingInformation) GetSMSequenceNumberOk() (*int32, bool)`

GetSMSequenceNumberOk returns a tuple with the SMSequenceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSMSequenceNumber

`func (o *SMSChargingInformation) SetSMSequenceNumber(v int32)`

SetSMSequenceNumber sets SMSequenceNumber field to given value.

### HasSMSequenceNumber

`func (o *SMSChargingInformation) HasSMSequenceNumber() bool`

HasSMSequenceNumber returns a boolean if a field has been set.

### GetSMSresult

`func (o *SMSChargingInformation) GetSMSresult() int32`

GetSMSresult returns the SMSresult field if non-nil, zero value otherwise.

### GetSMSresultOk

`func (o *SMSChargingInformation) GetSMSresultOk() (*int32, bool)`

GetSMSresultOk returns a tuple with the SMSresult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSMSresult

`func (o *SMSChargingInformation) SetSMSresult(v int32)`

SetSMSresult sets SMSresult field to given value.

### HasSMSresult

`func (o *SMSChargingInformation) HasSMSresult() bool`

HasSMSresult returns a boolean if a field has been set.

### GetSubmissionTime

`func (o *SMSChargingInformation) GetSubmissionTime() time.Time`

GetSubmissionTime returns the SubmissionTime field if non-nil, zero value otherwise.

### GetSubmissionTimeOk

`func (o *SMSChargingInformation) GetSubmissionTimeOk() (*time.Time, bool)`

GetSubmissionTimeOk returns a tuple with the SubmissionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmissionTime

`func (o *SMSChargingInformation) SetSubmissionTime(v time.Time)`

SetSubmissionTime sets SubmissionTime field to given value.

### HasSubmissionTime

`func (o *SMSChargingInformation) HasSubmissionTime() bool`

HasSubmissionTime returns a boolean if a field has been set.

### GetSMPriority

`func (o *SMSChargingInformation) GetSMPriority() SMPriority`

GetSMPriority returns the SMPriority field if non-nil, zero value otherwise.

### GetSMPriorityOk

`func (o *SMSChargingInformation) GetSMPriorityOk() (*SMPriority, bool)`

GetSMPriorityOk returns a tuple with the SMPriority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSMPriority

`func (o *SMSChargingInformation) SetSMPriority(v SMPriority)`

SetSMPriority sets SMPriority field to given value.

### HasSMPriority

`func (o *SMSChargingInformation) HasSMPriority() bool`

HasSMPriority returns a boolean if a field has been set.

### GetMessageReference

`func (o *SMSChargingInformation) GetMessageReference() string`

GetMessageReference returns the MessageReference field if non-nil, zero value otherwise.

### GetMessageReferenceOk

`func (o *SMSChargingInformation) GetMessageReferenceOk() (*string, bool)`

GetMessageReferenceOk returns a tuple with the MessageReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageReference

`func (o *SMSChargingInformation) SetMessageReference(v string)`

SetMessageReference sets MessageReference field to given value.

### HasMessageReference

`func (o *SMSChargingInformation) HasMessageReference() bool`

HasMessageReference returns a boolean if a field has been set.

### GetMessageSize

`func (o *SMSChargingInformation) GetMessageSize() int32`

GetMessageSize returns the MessageSize field if non-nil, zero value otherwise.

### GetMessageSizeOk

`func (o *SMSChargingInformation) GetMessageSizeOk() (*int32, bool)`

GetMessageSizeOk returns a tuple with the MessageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageSize

`func (o *SMSChargingInformation) SetMessageSize(v int32)`

SetMessageSize sets MessageSize field to given value.

### HasMessageSize

`func (o *SMSChargingInformation) HasMessageSize() bool`

HasMessageSize returns a boolean if a field has been set.

### GetMessageClass

`func (o *SMSChargingInformation) GetMessageClass() MessageClass`

GetMessageClass returns the MessageClass field if non-nil, zero value otherwise.

### GetMessageClassOk

`func (o *SMSChargingInformation) GetMessageClassOk() (*MessageClass, bool)`

GetMessageClassOk returns a tuple with the MessageClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageClass

`func (o *SMSChargingInformation) SetMessageClass(v MessageClass)`

SetMessageClass sets MessageClass field to given value.

### HasMessageClass

`func (o *SMSChargingInformation) HasMessageClass() bool`

HasMessageClass returns a boolean if a field has been set.

### GetDeliveryReportRequested

`func (o *SMSChargingInformation) GetDeliveryReportRequested() DeliveryReportRequested`

GetDeliveryReportRequested returns the DeliveryReportRequested field if non-nil, zero value otherwise.

### GetDeliveryReportRequestedOk

`func (o *SMSChargingInformation) GetDeliveryReportRequestedOk() (*DeliveryReportRequested, bool)`

GetDeliveryReportRequestedOk returns a tuple with the DeliveryReportRequested field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryReportRequested

`func (o *SMSChargingInformation) SetDeliveryReportRequested(v DeliveryReportRequested)`

SetDeliveryReportRequested sets DeliveryReportRequested field to given value.

### HasDeliveryReportRequested

`func (o *SMSChargingInformation) HasDeliveryReportRequested() bool`

HasDeliveryReportRequested returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


