# AccessTransferInformation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessTransferType** | Pointer to [**AccessTransferType**](AccessTransferType.md) |  | [optional] 
**AccessNetworkInformation** | Pointer to **[]string** |  | [optional] 
**CellularNetworkInformation** | Pointer to **string** |  | [optional] 
**InterUETransfer** | Pointer to [**UETransferType**](UETransferType.md) |  | [optional] 
**UserEquipmentInfo** | Pointer to **string** | String representing a Permanent Equipment Identifier that may contain - an IMEI or IMEISV, as  specified in clause 6.2 of 3GPP TS 23.003; a MAC address for a 5G-RG or FN-RG via  wireline  access, with an indication that this address cannot be trusted for regulatory purpose if this  address cannot be used as an Equipment Identifier of the FN-RG, as specified in clause 4.7.7  of 3GPP TS23.316. Examples are imei-012345678901234 or imeisv-0123456789012345.   | [optional] 
**InstanceId** | Pointer to **string** |  | [optional] 
**RelatedIMSChargingIdentifier** | Pointer to **string** |  | [optional] 
**RelatedIMSChargingIdentifierNode** | Pointer to [**IMSAddress**](IMSAddress.md) |  | [optional] 
**ChangeTime** | Pointer to **time.Time** | string with format &#39;date-time&#39; as defined in OpenAPI. | [optional] 

## Methods

### NewAccessTransferInformation

`func NewAccessTransferInformation() *AccessTransferInformation`

NewAccessTransferInformation instantiates a new AccessTransferInformation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccessTransferInformationWithDefaults

`func NewAccessTransferInformationWithDefaults() *AccessTransferInformation`

NewAccessTransferInformationWithDefaults instantiates a new AccessTransferInformation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessTransferType

`func (o *AccessTransferInformation) GetAccessTransferType() AccessTransferType`

GetAccessTransferType returns the AccessTransferType field if non-nil, zero value otherwise.

### GetAccessTransferTypeOk

`func (o *AccessTransferInformation) GetAccessTransferTypeOk() (*AccessTransferType, bool)`

GetAccessTransferTypeOk returns a tuple with the AccessTransferType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessTransferType

`func (o *AccessTransferInformation) SetAccessTransferType(v AccessTransferType)`

SetAccessTransferType sets AccessTransferType field to given value.

### HasAccessTransferType

`func (o *AccessTransferInformation) HasAccessTransferType() bool`

HasAccessTransferType returns a boolean if a field has been set.

### GetAccessNetworkInformation

`func (o *AccessTransferInformation) GetAccessNetworkInformation() []string`

GetAccessNetworkInformation returns the AccessNetworkInformation field if non-nil, zero value otherwise.

### GetAccessNetworkInformationOk

`func (o *AccessTransferInformation) GetAccessNetworkInformationOk() (*[]string, bool)`

GetAccessNetworkInformationOk returns a tuple with the AccessNetworkInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessNetworkInformation

`func (o *AccessTransferInformation) SetAccessNetworkInformation(v []string)`

SetAccessNetworkInformation sets AccessNetworkInformation field to given value.

### HasAccessNetworkInformation

`func (o *AccessTransferInformation) HasAccessNetworkInformation() bool`

HasAccessNetworkInformation returns a boolean if a field has been set.

### GetCellularNetworkInformation

`func (o *AccessTransferInformation) GetCellularNetworkInformation() string`

GetCellularNetworkInformation returns the CellularNetworkInformation field if non-nil, zero value otherwise.

### GetCellularNetworkInformationOk

`func (o *AccessTransferInformation) GetCellularNetworkInformationOk() (*string, bool)`

GetCellularNetworkInformationOk returns a tuple with the CellularNetworkInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCellularNetworkInformation

`func (o *AccessTransferInformation) SetCellularNetworkInformation(v string)`

SetCellularNetworkInformation sets CellularNetworkInformation field to given value.

### HasCellularNetworkInformation

`func (o *AccessTransferInformation) HasCellularNetworkInformation() bool`

HasCellularNetworkInformation returns a boolean if a field has been set.

### GetInterUETransfer

`func (o *AccessTransferInformation) GetInterUETransfer() UETransferType`

GetInterUETransfer returns the InterUETransfer field if non-nil, zero value otherwise.

### GetInterUETransferOk

`func (o *AccessTransferInformation) GetInterUETransferOk() (*UETransferType, bool)`

GetInterUETransferOk returns a tuple with the InterUETransfer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterUETransfer

`func (o *AccessTransferInformation) SetInterUETransfer(v UETransferType)`

SetInterUETransfer sets InterUETransfer field to given value.

### HasInterUETransfer

`func (o *AccessTransferInformation) HasInterUETransfer() bool`

HasInterUETransfer returns a boolean if a field has been set.

### GetUserEquipmentInfo

`func (o *AccessTransferInformation) GetUserEquipmentInfo() string`

GetUserEquipmentInfo returns the UserEquipmentInfo field if non-nil, zero value otherwise.

### GetUserEquipmentInfoOk

`func (o *AccessTransferInformation) GetUserEquipmentInfoOk() (*string, bool)`

GetUserEquipmentInfoOk returns a tuple with the UserEquipmentInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserEquipmentInfo

`func (o *AccessTransferInformation) SetUserEquipmentInfo(v string)`

SetUserEquipmentInfo sets UserEquipmentInfo field to given value.

### HasUserEquipmentInfo

`func (o *AccessTransferInformation) HasUserEquipmentInfo() bool`

HasUserEquipmentInfo returns a boolean if a field has been set.

### GetInstanceId

`func (o *AccessTransferInformation) GetInstanceId() string`

GetInstanceId returns the InstanceId field if non-nil, zero value otherwise.

### GetInstanceIdOk

`func (o *AccessTransferInformation) GetInstanceIdOk() (*string, bool)`

GetInstanceIdOk returns a tuple with the InstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceId

`func (o *AccessTransferInformation) SetInstanceId(v string)`

SetInstanceId sets InstanceId field to given value.

### HasInstanceId

`func (o *AccessTransferInformation) HasInstanceId() bool`

HasInstanceId returns a boolean if a field has been set.

### GetRelatedIMSChargingIdentifier

`func (o *AccessTransferInformation) GetRelatedIMSChargingIdentifier() string`

GetRelatedIMSChargingIdentifier returns the RelatedIMSChargingIdentifier field if non-nil, zero value otherwise.

### GetRelatedIMSChargingIdentifierOk

`func (o *AccessTransferInformation) GetRelatedIMSChargingIdentifierOk() (*string, bool)`

GetRelatedIMSChargingIdentifierOk returns a tuple with the RelatedIMSChargingIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelatedIMSChargingIdentifier

`func (o *AccessTransferInformation) SetRelatedIMSChargingIdentifier(v string)`

SetRelatedIMSChargingIdentifier sets RelatedIMSChargingIdentifier field to given value.

### HasRelatedIMSChargingIdentifier

`func (o *AccessTransferInformation) HasRelatedIMSChargingIdentifier() bool`

HasRelatedIMSChargingIdentifier returns a boolean if a field has been set.

### GetRelatedIMSChargingIdentifierNode

`func (o *AccessTransferInformation) GetRelatedIMSChargingIdentifierNode() IMSAddress`

GetRelatedIMSChargingIdentifierNode returns the RelatedIMSChargingIdentifierNode field if non-nil, zero value otherwise.

### GetRelatedIMSChargingIdentifierNodeOk

`func (o *AccessTransferInformation) GetRelatedIMSChargingIdentifierNodeOk() (*IMSAddress, bool)`

GetRelatedIMSChargingIdentifierNodeOk returns a tuple with the RelatedIMSChargingIdentifierNode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelatedIMSChargingIdentifierNode

`func (o *AccessTransferInformation) SetRelatedIMSChargingIdentifierNode(v IMSAddress)`

SetRelatedIMSChargingIdentifierNode sets RelatedIMSChargingIdentifierNode field to given value.

### HasRelatedIMSChargingIdentifierNode

`func (o *AccessTransferInformation) HasRelatedIMSChargingIdentifierNode() bool`

HasRelatedIMSChargingIdentifierNode returns a boolean if a field has been set.

### GetChangeTime

`func (o *AccessTransferInformation) GetChangeTime() time.Time`

GetChangeTime returns the ChangeTime field if non-nil, zero value otherwise.

### GetChangeTimeOk

`func (o *AccessTransferInformation) GetChangeTimeOk() (*time.Time, bool)`

GetChangeTimeOk returns a tuple with the ChangeTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeTime

`func (o *AccessTransferInformation) SetChangeTime(v time.Time)`

SetChangeTime sets ChangeTime field to given value.

### HasChangeTime

`func (o *AccessTransferInformation) HasChangeTime() bool`

HasChangeTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


