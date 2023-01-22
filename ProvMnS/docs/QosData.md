# QosData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**QosId** | Pointer to **string** |  | [optional] 
**FiveQIValue** | Pointer to **int32** |  | [optional] 
**MaxbrUl** | Pointer to **NullableString** | This data type is defined in the same way as the &#39;BitRate&#39; data type, but with the OpenAPI &#39;nullable: true&#39; property.  | [optional] 
**MaxbrDl** | Pointer to **NullableString** | This data type is defined in the same way as the &#39;BitRate&#39; data type, but with the OpenAPI &#39;nullable: true&#39; property.  | [optional] 
**GbrUl** | Pointer to **NullableString** | This data type is defined in the same way as the &#39;BitRate&#39; data type, but with the OpenAPI &#39;nullable: true&#39; property.  | [optional] 
**GbrDl** | Pointer to **NullableString** | This data type is defined in the same way as the &#39;BitRate&#39; data type, but with the OpenAPI &#39;nullable: true&#39; property.  | [optional] 
**Arp** | Pointer to [**Arp**](Arp.md) |  | [optional] 
**QosNotificationControl** | Pointer to **bool** |  | [optional] 
**ReflectiveQos** | Pointer to **bool** |  | [optional] 
**SharingKeyDl** | Pointer to **string** |  | [optional] 
**SharingKeyUl** | Pointer to **string** |  | [optional] 
**MaxPacketLossRateDl** | Pointer to **NullableInt32** | This data type is defined in the same way as the &#39;PacketLossRate&#39; data type, but with the OpenAPI &#39;nullable: true&#39; property.  | [optional] 
**MaxPacketLossRateUl** | Pointer to **NullableInt32** | This data type is defined in the same way as the &#39;PacketLossRate&#39; data type, but with the OpenAPI &#39;nullable: true&#39; property.  | [optional] 
**ExtMaxDataBurstVol** | Pointer to **NullableInt32** | This data type is defined in the same way as the &#39;ExtMaxDataBurstVol&#39; data type, but with the OpenAPI &#39;nullable: true&#39; property.  | [optional] 

## Methods

### NewQosData

`func NewQosData() *QosData`

NewQosData instantiates a new QosData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQosDataWithDefaults

`func NewQosDataWithDefaults() *QosData`

NewQosDataWithDefaults instantiates a new QosData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQosId

`func (o *QosData) GetQosId() string`

GetQosId returns the QosId field if non-nil, zero value otherwise.

### GetQosIdOk

`func (o *QosData) GetQosIdOk() (*string, bool)`

GetQosIdOk returns a tuple with the QosId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQosId

`func (o *QosData) SetQosId(v string)`

SetQosId sets QosId field to given value.

### HasQosId

`func (o *QosData) HasQosId() bool`

HasQosId returns a boolean if a field has been set.

### GetFiveQIValue

`func (o *QosData) GetFiveQIValue() int32`

GetFiveQIValue returns the FiveQIValue field if non-nil, zero value otherwise.

### GetFiveQIValueOk

`func (o *QosData) GetFiveQIValueOk() (*int32, bool)`

GetFiveQIValueOk returns a tuple with the FiveQIValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiveQIValue

`func (o *QosData) SetFiveQIValue(v int32)`

SetFiveQIValue sets FiveQIValue field to given value.

### HasFiveQIValue

`func (o *QosData) HasFiveQIValue() bool`

HasFiveQIValue returns a boolean if a field has been set.

### GetMaxbrUl

`func (o *QosData) GetMaxbrUl() string`

GetMaxbrUl returns the MaxbrUl field if non-nil, zero value otherwise.

### GetMaxbrUlOk

`func (o *QosData) GetMaxbrUlOk() (*string, bool)`

GetMaxbrUlOk returns a tuple with the MaxbrUl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxbrUl

`func (o *QosData) SetMaxbrUl(v string)`

SetMaxbrUl sets MaxbrUl field to given value.

### HasMaxbrUl

`func (o *QosData) HasMaxbrUl() bool`

HasMaxbrUl returns a boolean if a field has been set.

### SetMaxbrUlNil

`func (o *QosData) SetMaxbrUlNil(b bool)`

 SetMaxbrUlNil sets the value for MaxbrUl to be an explicit nil

### UnsetMaxbrUl
`func (o *QosData) UnsetMaxbrUl()`

UnsetMaxbrUl ensures that no value is present for MaxbrUl, not even an explicit nil
### GetMaxbrDl

`func (o *QosData) GetMaxbrDl() string`

GetMaxbrDl returns the MaxbrDl field if non-nil, zero value otherwise.

### GetMaxbrDlOk

`func (o *QosData) GetMaxbrDlOk() (*string, bool)`

GetMaxbrDlOk returns a tuple with the MaxbrDl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxbrDl

`func (o *QosData) SetMaxbrDl(v string)`

SetMaxbrDl sets MaxbrDl field to given value.

### HasMaxbrDl

`func (o *QosData) HasMaxbrDl() bool`

HasMaxbrDl returns a boolean if a field has been set.

### SetMaxbrDlNil

`func (o *QosData) SetMaxbrDlNil(b bool)`

 SetMaxbrDlNil sets the value for MaxbrDl to be an explicit nil

### UnsetMaxbrDl
`func (o *QosData) UnsetMaxbrDl()`

UnsetMaxbrDl ensures that no value is present for MaxbrDl, not even an explicit nil
### GetGbrUl

`func (o *QosData) GetGbrUl() string`

GetGbrUl returns the GbrUl field if non-nil, zero value otherwise.

### GetGbrUlOk

`func (o *QosData) GetGbrUlOk() (*string, bool)`

GetGbrUlOk returns a tuple with the GbrUl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGbrUl

`func (o *QosData) SetGbrUl(v string)`

SetGbrUl sets GbrUl field to given value.

### HasGbrUl

`func (o *QosData) HasGbrUl() bool`

HasGbrUl returns a boolean if a field has been set.

### SetGbrUlNil

`func (o *QosData) SetGbrUlNil(b bool)`

 SetGbrUlNil sets the value for GbrUl to be an explicit nil

### UnsetGbrUl
`func (o *QosData) UnsetGbrUl()`

UnsetGbrUl ensures that no value is present for GbrUl, not even an explicit nil
### GetGbrDl

`func (o *QosData) GetGbrDl() string`

GetGbrDl returns the GbrDl field if non-nil, zero value otherwise.

### GetGbrDlOk

`func (o *QosData) GetGbrDlOk() (*string, bool)`

GetGbrDlOk returns a tuple with the GbrDl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGbrDl

`func (o *QosData) SetGbrDl(v string)`

SetGbrDl sets GbrDl field to given value.

### HasGbrDl

`func (o *QosData) HasGbrDl() bool`

HasGbrDl returns a boolean if a field has been set.

### SetGbrDlNil

`func (o *QosData) SetGbrDlNil(b bool)`

 SetGbrDlNil sets the value for GbrDl to be an explicit nil

### UnsetGbrDl
`func (o *QosData) UnsetGbrDl()`

UnsetGbrDl ensures that no value is present for GbrDl, not even an explicit nil
### GetArp

`func (o *QosData) GetArp() Arp`

GetArp returns the Arp field if non-nil, zero value otherwise.

### GetArpOk

`func (o *QosData) GetArpOk() (*Arp, bool)`

GetArpOk returns a tuple with the Arp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArp

`func (o *QosData) SetArp(v Arp)`

SetArp sets Arp field to given value.

### HasArp

`func (o *QosData) HasArp() bool`

HasArp returns a boolean if a field has been set.

### GetQosNotificationControl

`func (o *QosData) GetQosNotificationControl() bool`

GetQosNotificationControl returns the QosNotificationControl field if non-nil, zero value otherwise.

### GetQosNotificationControlOk

`func (o *QosData) GetQosNotificationControlOk() (*bool, bool)`

GetQosNotificationControlOk returns a tuple with the QosNotificationControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQosNotificationControl

`func (o *QosData) SetQosNotificationControl(v bool)`

SetQosNotificationControl sets QosNotificationControl field to given value.

### HasQosNotificationControl

`func (o *QosData) HasQosNotificationControl() bool`

HasQosNotificationControl returns a boolean if a field has been set.

### GetReflectiveQos

`func (o *QosData) GetReflectiveQos() bool`

GetReflectiveQos returns the ReflectiveQos field if non-nil, zero value otherwise.

### GetReflectiveQosOk

`func (o *QosData) GetReflectiveQosOk() (*bool, bool)`

GetReflectiveQosOk returns a tuple with the ReflectiveQos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReflectiveQos

`func (o *QosData) SetReflectiveQos(v bool)`

SetReflectiveQos sets ReflectiveQos field to given value.

### HasReflectiveQos

`func (o *QosData) HasReflectiveQos() bool`

HasReflectiveQos returns a boolean if a field has been set.

### GetSharingKeyDl

`func (o *QosData) GetSharingKeyDl() string`

GetSharingKeyDl returns the SharingKeyDl field if non-nil, zero value otherwise.

### GetSharingKeyDlOk

`func (o *QosData) GetSharingKeyDlOk() (*string, bool)`

GetSharingKeyDlOk returns a tuple with the SharingKeyDl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharingKeyDl

`func (o *QosData) SetSharingKeyDl(v string)`

SetSharingKeyDl sets SharingKeyDl field to given value.

### HasSharingKeyDl

`func (o *QosData) HasSharingKeyDl() bool`

HasSharingKeyDl returns a boolean if a field has been set.

### GetSharingKeyUl

`func (o *QosData) GetSharingKeyUl() string`

GetSharingKeyUl returns the SharingKeyUl field if non-nil, zero value otherwise.

### GetSharingKeyUlOk

`func (o *QosData) GetSharingKeyUlOk() (*string, bool)`

GetSharingKeyUlOk returns a tuple with the SharingKeyUl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharingKeyUl

`func (o *QosData) SetSharingKeyUl(v string)`

SetSharingKeyUl sets SharingKeyUl field to given value.

### HasSharingKeyUl

`func (o *QosData) HasSharingKeyUl() bool`

HasSharingKeyUl returns a boolean if a field has been set.

### GetMaxPacketLossRateDl

`func (o *QosData) GetMaxPacketLossRateDl() int32`

GetMaxPacketLossRateDl returns the MaxPacketLossRateDl field if non-nil, zero value otherwise.

### GetMaxPacketLossRateDlOk

`func (o *QosData) GetMaxPacketLossRateDlOk() (*int32, bool)`

GetMaxPacketLossRateDlOk returns a tuple with the MaxPacketLossRateDl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPacketLossRateDl

`func (o *QosData) SetMaxPacketLossRateDl(v int32)`

SetMaxPacketLossRateDl sets MaxPacketLossRateDl field to given value.

### HasMaxPacketLossRateDl

`func (o *QosData) HasMaxPacketLossRateDl() bool`

HasMaxPacketLossRateDl returns a boolean if a field has been set.

### SetMaxPacketLossRateDlNil

`func (o *QosData) SetMaxPacketLossRateDlNil(b bool)`

 SetMaxPacketLossRateDlNil sets the value for MaxPacketLossRateDl to be an explicit nil

### UnsetMaxPacketLossRateDl
`func (o *QosData) UnsetMaxPacketLossRateDl()`

UnsetMaxPacketLossRateDl ensures that no value is present for MaxPacketLossRateDl, not even an explicit nil
### GetMaxPacketLossRateUl

`func (o *QosData) GetMaxPacketLossRateUl() int32`

GetMaxPacketLossRateUl returns the MaxPacketLossRateUl field if non-nil, zero value otherwise.

### GetMaxPacketLossRateUlOk

`func (o *QosData) GetMaxPacketLossRateUlOk() (*int32, bool)`

GetMaxPacketLossRateUlOk returns a tuple with the MaxPacketLossRateUl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPacketLossRateUl

`func (o *QosData) SetMaxPacketLossRateUl(v int32)`

SetMaxPacketLossRateUl sets MaxPacketLossRateUl field to given value.

### HasMaxPacketLossRateUl

`func (o *QosData) HasMaxPacketLossRateUl() bool`

HasMaxPacketLossRateUl returns a boolean if a field has been set.

### SetMaxPacketLossRateUlNil

`func (o *QosData) SetMaxPacketLossRateUlNil(b bool)`

 SetMaxPacketLossRateUlNil sets the value for MaxPacketLossRateUl to be an explicit nil

### UnsetMaxPacketLossRateUl
`func (o *QosData) UnsetMaxPacketLossRateUl()`

UnsetMaxPacketLossRateUl ensures that no value is present for MaxPacketLossRateUl, not even an explicit nil
### GetExtMaxDataBurstVol

`func (o *QosData) GetExtMaxDataBurstVol() int32`

GetExtMaxDataBurstVol returns the ExtMaxDataBurstVol field if non-nil, zero value otherwise.

### GetExtMaxDataBurstVolOk

`func (o *QosData) GetExtMaxDataBurstVolOk() (*int32, bool)`

GetExtMaxDataBurstVolOk returns a tuple with the ExtMaxDataBurstVol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtMaxDataBurstVol

`func (o *QosData) SetExtMaxDataBurstVol(v int32)`

SetExtMaxDataBurstVol sets ExtMaxDataBurstVol field to given value.

### HasExtMaxDataBurstVol

`func (o *QosData) HasExtMaxDataBurstVol() bool`

HasExtMaxDataBurstVol returns a boolean if a field has been set.

### SetExtMaxDataBurstVolNil

`func (o *QosData) SetExtMaxDataBurstVolNil(b bool)`

 SetExtMaxDataBurstVolNil sets the value for ExtMaxDataBurstVol to be an explicit nil

### UnsetExtMaxDataBurstVol
`func (o *QosData) UnsetExtMaxDataBurstVol()`

UnsetExtMaxDataBurstVol ensures that no value is present for ExtMaxDataBurstVol, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


