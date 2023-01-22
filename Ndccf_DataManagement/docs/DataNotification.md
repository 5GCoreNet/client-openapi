# DataNotification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AmfEventNotifs** | Pointer to [**[]AmfEventNotification**](AmfEventNotification.md) | List of notifications of AMF events. | [optional] 
**SmfEventNotifs** | Pointer to [**[]NsmfEventExposureNotification**](NsmfEventExposureNotification.md) | List of notifications of SMF events. | [optional] 
**UdmEventNotifs** | Pointer to [**[]MonitoringReport**](MonitoringReport.md) | List of notifications of UDM events. | [optional] 
**NefEventNotifs** | Pointer to [**[]NefEventExposureNotif**](NefEventExposureNotif.md) | List of notifications of NEF events. | [optional] 
**AfEventNotifs** | Pointer to [**[]AfEventExposureNotif**](AfEventExposureNotif.md) | List of notifications of AF events. | [optional] 
**NrfEventNotifs** | Pointer to [**[]NotificationData**](NotificationData.md) | List of notifications of NRF events. | [optional] 
**NsacfEventNotifs** | Pointer to [**[]SACEventReport**](SACEventReport.md) | List of notifications of NSACF events. | [optional] 

## Methods

### NewDataNotification

`func NewDataNotification() *DataNotification`

NewDataNotification instantiates a new DataNotification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataNotificationWithDefaults

`func NewDataNotificationWithDefaults() *DataNotification`

NewDataNotificationWithDefaults instantiates a new DataNotification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmfEventNotifs

`func (o *DataNotification) GetAmfEventNotifs() []AmfEventNotification`

GetAmfEventNotifs returns the AmfEventNotifs field if non-nil, zero value otherwise.

### GetAmfEventNotifsOk

`func (o *DataNotification) GetAmfEventNotifsOk() (*[]AmfEventNotification, bool)`

GetAmfEventNotifsOk returns a tuple with the AmfEventNotifs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmfEventNotifs

`func (o *DataNotification) SetAmfEventNotifs(v []AmfEventNotification)`

SetAmfEventNotifs sets AmfEventNotifs field to given value.

### HasAmfEventNotifs

`func (o *DataNotification) HasAmfEventNotifs() bool`

HasAmfEventNotifs returns a boolean if a field has been set.

### GetSmfEventNotifs

`func (o *DataNotification) GetSmfEventNotifs() []NsmfEventExposureNotification`

GetSmfEventNotifs returns the SmfEventNotifs field if non-nil, zero value otherwise.

### GetSmfEventNotifsOk

`func (o *DataNotification) GetSmfEventNotifsOk() (*[]NsmfEventExposureNotification, bool)`

GetSmfEventNotifsOk returns a tuple with the SmfEventNotifs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmfEventNotifs

`func (o *DataNotification) SetSmfEventNotifs(v []NsmfEventExposureNotification)`

SetSmfEventNotifs sets SmfEventNotifs field to given value.

### HasSmfEventNotifs

`func (o *DataNotification) HasSmfEventNotifs() bool`

HasSmfEventNotifs returns a boolean if a field has been set.

### GetUdmEventNotifs

`func (o *DataNotification) GetUdmEventNotifs() []MonitoringReport`

GetUdmEventNotifs returns the UdmEventNotifs field if non-nil, zero value otherwise.

### GetUdmEventNotifsOk

`func (o *DataNotification) GetUdmEventNotifsOk() (*[]MonitoringReport, bool)`

GetUdmEventNotifsOk returns a tuple with the UdmEventNotifs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUdmEventNotifs

`func (o *DataNotification) SetUdmEventNotifs(v []MonitoringReport)`

SetUdmEventNotifs sets UdmEventNotifs field to given value.

### HasUdmEventNotifs

`func (o *DataNotification) HasUdmEventNotifs() bool`

HasUdmEventNotifs returns a boolean if a field has been set.

### GetNefEventNotifs

`func (o *DataNotification) GetNefEventNotifs() []NefEventExposureNotif`

GetNefEventNotifs returns the NefEventNotifs field if non-nil, zero value otherwise.

### GetNefEventNotifsOk

`func (o *DataNotification) GetNefEventNotifsOk() (*[]NefEventExposureNotif, bool)`

GetNefEventNotifsOk returns a tuple with the NefEventNotifs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNefEventNotifs

`func (o *DataNotification) SetNefEventNotifs(v []NefEventExposureNotif)`

SetNefEventNotifs sets NefEventNotifs field to given value.

### HasNefEventNotifs

`func (o *DataNotification) HasNefEventNotifs() bool`

HasNefEventNotifs returns a boolean if a field has been set.

### GetAfEventNotifs

`func (o *DataNotification) GetAfEventNotifs() []AfEventExposureNotif`

GetAfEventNotifs returns the AfEventNotifs field if non-nil, zero value otherwise.

### GetAfEventNotifsOk

`func (o *DataNotification) GetAfEventNotifsOk() (*[]AfEventExposureNotif, bool)`

GetAfEventNotifsOk returns a tuple with the AfEventNotifs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAfEventNotifs

`func (o *DataNotification) SetAfEventNotifs(v []AfEventExposureNotif)`

SetAfEventNotifs sets AfEventNotifs field to given value.

### HasAfEventNotifs

`func (o *DataNotification) HasAfEventNotifs() bool`

HasAfEventNotifs returns a boolean if a field has been set.

### GetNrfEventNotifs

`func (o *DataNotification) GetNrfEventNotifs() []NotificationData`

GetNrfEventNotifs returns the NrfEventNotifs field if non-nil, zero value otherwise.

### GetNrfEventNotifsOk

`func (o *DataNotification) GetNrfEventNotifsOk() (*[]NotificationData, bool)`

GetNrfEventNotifsOk returns a tuple with the NrfEventNotifs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNrfEventNotifs

`func (o *DataNotification) SetNrfEventNotifs(v []NotificationData)`

SetNrfEventNotifs sets NrfEventNotifs field to given value.

### HasNrfEventNotifs

`func (o *DataNotification) HasNrfEventNotifs() bool`

HasNrfEventNotifs returns a boolean if a field has been set.

### GetNsacfEventNotifs

`func (o *DataNotification) GetNsacfEventNotifs() []SACEventReport`

GetNsacfEventNotifs returns the NsacfEventNotifs field if non-nil, zero value otherwise.

### GetNsacfEventNotifsOk

`func (o *DataNotification) GetNsacfEventNotifsOk() (*[]SACEventReport, bool)`

GetNsacfEventNotifsOk returns a tuple with the NsacfEventNotifs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNsacfEventNotifs

`func (o *DataNotification) SetNsacfEventNotifs(v []SACEventReport)`

SetNsacfEventNotifs sets NsacfEventNotifs field to given value.

### HasNsacfEventNotifs

`func (o *DataNotification) HasNsacfEventNotifs() bool`

HasNsacfEventNotifs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


