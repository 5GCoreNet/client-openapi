# NotificationItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventType** | [**EventType**](EventType.md) |  | 
**UeIpv4Addr** | Pointer to **string** | String identifying a IPv4 address formatted in the &#39;dotted decimal&#39; notation as defined in RFC 1166.  | [optional] 
**UeIpv6Prefix** | Pointer to [**Ipv6Prefix**](Ipv6Prefix.md) |  | [optional] 
**UeMacAddr** | Pointer to **string** | String identifying a MAC address formatted in the hexadecimal notation according to clause 1.1 and clause 2.1 of RFC 7042.  | [optional] 
**Dnn** | Pointer to **string** | String representing a Data Network as defined in clause 9A of 3GPP TS 23.003;  it shall contain either a DNN Network Identifier, or a full DNN with both the Network  Identifier and Operator Identifier, as specified in 3GPP TS 23.003 clause 9.1.1 and 9.1.2. It shall be coded as string in which the labels are separated by dots  (e.g. \&quot;Label1.Label2.Label3\&quot;).  | [optional] 
**Snssai** | Pointer to [**Snssai**](Snssai.md) |  | [optional] 
**Gpsi** | Pointer to **string** | String identifying a Gpsi shall contain either an External Id or an MSISDN.  It shall be formatted as follows -External Identifier&#x3D; \&quot;extid-&#39;extid&#39;, where &#39;extid&#39;  shall be formatted according to clause 19.7.2 of 3GPP TS 23.003 that describes an  External Identifier.   | [optional] 
**TimeStamp** | **time.Time** | string with format &#39;date-time&#39; as defined in OpenAPI. | 
**StartTime** | Pointer to **time.Time** | string with format &#39;date-time&#39; as defined in OpenAPI. | [optional] 
**QosMonitoringMeasurement** | Pointer to [**QosMonitoringMeasurement**](QosMonitoringMeasurement.md) |  | [optional] 

## Methods

### NewNotificationItem

`func NewNotificationItem(eventType EventType, timeStamp time.Time, ) *NotificationItem`

NewNotificationItem instantiates a new NotificationItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationItemWithDefaults

`func NewNotificationItemWithDefaults() *NotificationItem`

NewNotificationItemWithDefaults instantiates a new NotificationItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventType

`func (o *NotificationItem) GetEventType() EventType`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *NotificationItem) GetEventTypeOk() (*EventType, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *NotificationItem) SetEventType(v EventType)`

SetEventType sets EventType field to given value.


### GetUeIpv4Addr

`func (o *NotificationItem) GetUeIpv4Addr() string`

GetUeIpv4Addr returns the UeIpv4Addr field if non-nil, zero value otherwise.

### GetUeIpv4AddrOk

`func (o *NotificationItem) GetUeIpv4AddrOk() (*string, bool)`

GetUeIpv4AddrOk returns a tuple with the UeIpv4Addr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUeIpv4Addr

`func (o *NotificationItem) SetUeIpv4Addr(v string)`

SetUeIpv4Addr sets UeIpv4Addr field to given value.

### HasUeIpv4Addr

`func (o *NotificationItem) HasUeIpv4Addr() bool`

HasUeIpv4Addr returns a boolean if a field has been set.

### GetUeIpv6Prefix

`func (o *NotificationItem) GetUeIpv6Prefix() Ipv6Prefix`

GetUeIpv6Prefix returns the UeIpv6Prefix field if non-nil, zero value otherwise.

### GetUeIpv6PrefixOk

`func (o *NotificationItem) GetUeIpv6PrefixOk() (*Ipv6Prefix, bool)`

GetUeIpv6PrefixOk returns a tuple with the UeIpv6Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUeIpv6Prefix

`func (o *NotificationItem) SetUeIpv6Prefix(v Ipv6Prefix)`

SetUeIpv6Prefix sets UeIpv6Prefix field to given value.

### HasUeIpv6Prefix

`func (o *NotificationItem) HasUeIpv6Prefix() bool`

HasUeIpv6Prefix returns a boolean if a field has been set.

### GetUeMacAddr

`func (o *NotificationItem) GetUeMacAddr() string`

GetUeMacAddr returns the UeMacAddr field if non-nil, zero value otherwise.

### GetUeMacAddrOk

`func (o *NotificationItem) GetUeMacAddrOk() (*string, bool)`

GetUeMacAddrOk returns a tuple with the UeMacAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUeMacAddr

`func (o *NotificationItem) SetUeMacAddr(v string)`

SetUeMacAddr sets UeMacAddr field to given value.

### HasUeMacAddr

`func (o *NotificationItem) HasUeMacAddr() bool`

HasUeMacAddr returns a boolean if a field has been set.

### GetDnn

`func (o *NotificationItem) GetDnn() string`

GetDnn returns the Dnn field if non-nil, zero value otherwise.

### GetDnnOk

`func (o *NotificationItem) GetDnnOk() (*string, bool)`

GetDnnOk returns a tuple with the Dnn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnn

`func (o *NotificationItem) SetDnn(v string)`

SetDnn sets Dnn field to given value.

### HasDnn

`func (o *NotificationItem) HasDnn() bool`

HasDnn returns a boolean if a field has been set.

### GetSnssai

`func (o *NotificationItem) GetSnssai() Snssai`

GetSnssai returns the Snssai field if non-nil, zero value otherwise.

### GetSnssaiOk

`func (o *NotificationItem) GetSnssaiOk() (*Snssai, bool)`

GetSnssaiOk returns a tuple with the Snssai field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnssai

`func (o *NotificationItem) SetSnssai(v Snssai)`

SetSnssai sets Snssai field to given value.

### HasSnssai

`func (o *NotificationItem) HasSnssai() bool`

HasSnssai returns a boolean if a field has been set.

### GetGpsi

`func (o *NotificationItem) GetGpsi() string`

GetGpsi returns the Gpsi field if non-nil, zero value otherwise.

### GetGpsiOk

`func (o *NotificationItem) GetGpsiOk() (*string, bool)`

GetGpsiOk returns a tuple with the Gpsi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpsi

`func (o *NotificationItem) SetGpsi(v string)`

SetGpsi sets Gpsi field to given value.

### HasGpsi

`func (o *NotificationItem) HasGpsi() bool`

HasGpsi returns a boolean if a field has been set.

### GetTimeStamp

`func (o *NotificationItem) GetTimeStamp() time.Time`

GetTimeStamp returns the TimeStamp field if non-nil, zero value otherwise.

### GetTimeStampOk

`func (o *NotificationItem) GetTimeStampOk() (*time.Time, bool)`

GetTimeStampOk returns a tuple with the TimeStamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeStamp

`func (o *NotificationItem) SetTimeStamp(v time.Time)`

SetTimeStamp sets TimeStamp field to given value.


### GetStartTime

`func (o *NotificationItem) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *NotificationItem) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *NotificationItem) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *NotificationItem) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetQosMonitoringMeasurement

`func (o *NotificationItem) GetQosMonitoringMeasurement() QosMonitoringMeasurement`

GetQosMonitoringMeasurement returns the QosMonitoringMeasurement field if non-nil, zero value otherwise.

### GetQosMonitoringMeasurementOk

`func (o *NotificationItem) GetQosMonitoringMeasurementOk() (*QosMonitoringMeasurement, bool)`

GetQosMonitoringMeasurementOk returns a tuple with the QosMonitoringMeasurement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQosMonitoringMeasurement

`func (o *NotificationItem) SetQosMonitoringMeasurement(v QosMonitoringMeasurement)`

SetQosMonitoringMeasurement sets QosMonitoringMeasurement field to given value.

### HasQosMonitoringMeasurement

`func (o *NotificationItem) HasQosMonitoringMeasurement() bool`

HasQosMonitoringMeasurement returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


