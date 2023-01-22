# LocationInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AgeOfLocationInfo** | Pointer to **int32** | Unsigned integer identifying a period of time in units of minutes. | [optional] 
**CellId** | Pointer to **string** | Indicates the Cell Global Identification of the user which identifies the cell the UE is registered. | [optional] 
**EnodeBId** | Pointer to **string** | Indicates the eNodeB in which the UE is currently located. | [optional] 
**RoutingAreaId** | Pointer to **string** | Identifies the Routing Area Identity of the user where the UE is located. | [optional] 
**TrackingAreaId** | Pointer to **string** | Identifies the Tracking Area Identity of the user where the UE is located. | [optional] 
**PlmnId** | Pointer to **string** | Identifies the PLMN Identity of the user where the UE is located. | [optional] 
**TwanId** | Pointer to **string** | Identifies the TWAN Identity of the user where the UE is located. | [optional] 
**GeographicArea** | Pointer to [**GeographicArea**](GeographicArea.md) |  | [optional] 
**CivicAddress** | Pointer to [**CivicAddress**](CivicAddress.md) |  | [optional] 
**PositionMethod** | Pointer to [**PositioningMethod**](PositioningMethod.md) |  | [optional] 
**QosFulfilInd** | Pointer to [**AccuracyFulfilmentIndicator**](AccuracyFulfilmentIndicator.md) |  | [optional] 
**UeVelocity** | Pointer to [**VelocityEstimate**](VelocityEstimate.md) |  | [optional] 
**LdrType** | Pointer to [**LdrType**](LdrType.md) |  | [optional] 
**AchievedQos** | Pointer to [**MinorLocationQoS**](MinorLocationQoS.md) |  | [optional] 

## Methods

### NewLocationInfo

`func NewLocationInfo() *LocationInfo`

NewLocationInfo instantiates a new LocationInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLocationInfoWithDefaults

`func NewLocationInfoWithDefaults() *LocationInfo`

NewLocationInfoWithDefaults instantiates a new LocationInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAgeOfLocationInfo

`func (o *LocationInfo) GetAgeOfLocationInfo() int32`

GetAgeOfLocationInfo returns the AgeOfLocationInfo field if non-nil, zero value otherwise.

### GetAgeOfLocationInfoOk

`func (o *LocationInfo) GetAgeOfLocationInfoOk() (*int32, bool)`

GetAgeOfLocationInfoOk returns a tuple with the AgeOfLocationInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgeOfLocationInfo

`func (o *LocationInfo) SetAgeOfLocationInfo(v int32)`

SetAgeOfLocationInfo sets AgeOfLocationInfo field to given value.

### HasAgeOfLocationInfo

`func (o *LocationInfo) HasAgeOfLocationInfo() bool`

HasAgeOfLocationInfo returns a boolean if a field has been set.

### GetCellId

`func (o *LocationInfo) GetCellId() string`

GetCellId returns the CellId field if non-nil, zero value otherwise.

### GetCellIdOk

`func (o *LocationInfo) GetCellIdOk() (*string, bool)`

GetCellIdOk returns a tuple with the CellId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCellId

`func (o *LocationInfo) SetCellId(v string)`

SetCellId sets CellId field to given value.

### HasCellId

`func (o *LocationInfo) HasCellId() bool`

HasCellId returns a boolean if a field has been set.

### GetEnodeBId

`func (o *LocationInfo) GetEnodeBId() string`

GetEnodeBId returns the EnodeBId field if non-nil, zero value otherwise.

### GetEnodeBIdOk

`func (o *LocationInfo) GetEnodeBIdOk() (*string, bool)`

GetEnodeBIdOk returns a tuple with the EnodeBId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnodeBId

`func (o *LocationInfo) SetEnodeBId(v string)`

SetEnodeBId sets EnodeBId field to given value.

### HasEnodeBId

`func (o *LocationInfo) HasEnodeBId() bool`

HasEnodeBId returns a boolean if a field has been set.

### GetRoutingAreaId

`func (o *LocationInfo) GetRoutingAreaId() string`

GetRoutingAreaId returns the RoutingAreaId field if non-nil, zero value otherwise.

### GetRoutingAreaIdOk

`func (o *LocationInfo) GetRoutingAreaIdOk() (*string, bool)`

GetRoutingAreaIdOk returns a tuple with the RoutingAreaId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutingAreaId

`func (o *LocationInfo) SetRoutingAreaId(v string)`

SetRoutingAreaId sets RoutingAreaId field to given value.

### HasRoutingAreaId

`func (o *LocationInfo) HasRoutingAreaId() bool`

HasRoutingAreaId returns a boolean if a field has been set.

### GetTrackingAreaId

`func (o *LocationInfo) GetTrackingAreaId() string`

GetTrackingAreaId returns the TrackingAreaId field if non-nil, zero value otherwise.

### GetTrackingAreaIdOk

`func (o *LocationInfo) GetTrackingAreaIdOk() (*string, bool)`

GetTrackingAreaIdOk returns a tuple with the TrackingAreaId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackingAreaId

`func (o *LocationInfo) SetTrackingAreaId(v string)`

SetTrackingAreaId sets TrackingAreaId field to given value.

### HasTrackingAreaId

`func (o *LocationInfo) HasTrackingAreaId() bool`

HasTrackingAreaId returns a boolean if a field has been set.

### GetPlmnId

`func (o *LocationInfo) GetPlmnId() string`

GetPlmnId returns the PlmnId field if non-nil, zero value otherwise.

### GetPlmnIdOk

`func (o *LocationInfo) GetPlmnIdOk() (*string, bool)`

GetPlmnIdOk returns a tuple with the PlmnId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlmnId

`func (o *LocationInfo) SetPlmnId(v string)`

SetPlmnId sets PlmnId field to given value.

### HasPlmnId

`func (o *LocationInfo) HasPlmnId() bool`

HasPlmnId returns a boolean if a field has been set.

### GetTwanId

`func (o *LocationInfo) GetTwanId() string`

GetTwanId returns the TwanId field if non-nil, zero value otherwise.

### GetTwanIdOk

`func (o *LocationInfo) GetTwanIdOk() (*string, bool)`

GetTwanIdOk returns a tuple with the TwanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwanId

`func (o *LocationInfo) SetTwanId(v string)`

SetTwanId sets TwanId field to given value.

### HasTwanId

`func (o *LocationInfo) HasTwanId() bool`

HasTwanId returns a boolean if a field has been set.

### GetGeographicArea

`func (o *LocationInfo) GetGeographicArea() GeographicArea`

GetGeographicArea returns the GeographicArea field if non-nil, zero value otherwise.

### GetGeographicAreaOk

`func (o *LocationInfo) GetGeographicAreaOk() (*GeographicArea, bool)`

GetGeographicAreaOk returns a tuple with the GeographicArea field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeographicArea

`func (o *LocationInfo) SetGeographicArea(v GeographicArea)`

SetGeographicArea sets GeographicArea field to given value.

### HasGeographicArea

`func (o *LocationInfo) HasGeographicArea() bool`

HasGeographicArea returns a boolean if a field has been set.

### GetCivicAddress

`func (o *LocationInfo) GetCivicAddress() CivicAddress`

GetCivicAddress returns the CivicAddress field if non-nil, zero value otherwise.

### GetCivicAddressOk

`func (o *LocationInfo) GetCivicAddressOk() (*CivicAddress, bool)`

GetCivicAddressOk returns a tuple with the CivicAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCivicAddress

`func (o *LocationInfo) SetCivicAddress(v CivicAddress)`

SetCivicAddress sets CivicAddress field to given value.

### HasCivicAddress

`func (o *LocationInfo) HasCivicAddress() bool`

HasCivicAddress returns a boolean if a field has been set.

### GetPositionMethod

`func (o *LocationInfo) GetPositionMethod() PositioningMethod`

GetPositionMethod returns the PositionMethod field if non-nil, zero value otherwise.

### GetPositionMethodOk

`func (o *LocationInfo) GetPositionMethodOk() (*PositioningMethod, bool)`

GetPositionMethodOk returns a tuple with the PositionMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPositionMethod

`func (o *LocationInfo) SetPositionMethod(v PositioningMethod)`

SetPositionMethod sets PositionMethod field to given value.

### HasPositionMethod

`func (o *LocationInfo) HasPositionMethod() bool`

HasPositionMethod returns a boolean if a field has been set.

### GetQosFulfilInd

`func (o *LocationInfo) GetQosFulfilInd() AccuracyFulfilmentIndicator`

GetQosFulfilInd returns the QosFulfilInd field if non-nil, zero value otherwise.

### GetQosFulfilIndOk

`func (o *LocationInfo) GetQosFulfilIndOk() (*AccuracyFulfilmentIndicator, bool)`

GetQosFulfilIndOk returns a tuple with the QosFulfilInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQosFulfilInd

`func (o *LocationInfo) SetQosFulfilInd(v AccuracyFulfilmentIndicator)`

SetQosFulfilInd sets QosFulfilInd field to given value.

### HasQosFulfilInd

`func (o *LocationInfo) HasQosFulfilInd() bool`

HasQosFulfilInd returns a boolean if a field has been set.

### GetUeVelocity

`func (o *LocationInfo) GetUeVelocity() VelocityEstimate`

GetUeVelocity returns the UeVelocity field if non-nil, zero value otherwise.

### GetUeVelocityOk

`func (o *LocationInfo) GetUeVelocityOk() (*VelocityEstimate, bool)`

GetUeVelocityOk returns a tuple with the UeVelocity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUeVelocity

`func (o *LocationInfo) SetUeVelocity(v VelocityEstimate)`

SetUeVelocity sets UeVelocity field to given value.

### HasUeVelocity

`func (o *LocationInfo) HasUeVelocity() bool`

HasUeVelocity returns a boolean if a field has been set.

### GetLdrType

`func (o *LocationInfo) GetLdrType() LdrType`

GetLdrType returns the LdrType field if non-nil, zero value otherwise.

### GetLdrTypeOk

`func (o *LocationInfo) GetLdrTypeOk() (*LdrType, bool)`

GetLdrTypeOk returns a tuple with the LdrType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdrType

`func (o *LocationInfo) SetLdrType(v LdrType)`

SetLdrType sets LdrType field to given value.

### HasLdrType

`func (o *LocationInfo) HasLdrType() bool`

HasLdrType returns a boolean if a field has been set.

### GetAchievedQos

`func (o *LocationInfo) GetAchievedQos() MinorLocationQoS`

GetAchievedQos returns the AchievedQos field if non-nil, zero value otherwise.

### GetAchievedQosOk

`func (o *LocationInfo) GetAchievedQosOk() (*MinorLocationQoS, bool)`

GetAchievedQosOk returns a tuple with the AchievedQos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAchievedQos

`func (o *LocationInfo) SetAchievedQos(v MinorLocationQoS)`

SetAchievedQos sets AchievedQos field to given value.

### HasAchievedQos

`func (o *LocationInfo) HasAchievedQos() bool`

HasAchievedQos returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


