# SmContextConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SmalDataRateControl** | Pointer to [**SmallDataRateControl**](SmallDataRateControl.md) |  | [optional] 
**SmallDataRateStatus** | Pointer to [**SmallDataRateStatus**](SmallDataRateStatus.md) |  | [optional] 
**ServPlmnDataRateCtl** | Pointer to **NullableInt32** | When present, this IE shall contain the maximum allowed number of Downlink NAS Data PDUs per deci hour of the serving PLMN, as specified in subclause 5.31.14.2 of 3GPP TS 23.501 [2].   Minimum  10  | [optional] 

## Methods

### NewSmContextConfiguration

`func NewSmContextConfiguration() *SmContextConfiguration`

NewSmContextConfiguration instantiates a new SmContextConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSmContextConfigurationWithDefaults

`func NewSmContextConfigurationWithDefaults() *SmContextConfiguration`

NewSmContextConfigurationWithDefaults instantiates a new SmContextConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSmalDataRateControl

`func (o *SmContextConfiguration) GetSmalDataRateControl() SmallDataRateControl`

GetSmalDataRateControl returns the SmalDataRateControl field if non-nil, zero value otherwise.

### GetSmalDataRateControlOk

`func (o *SmContextConfiguration) GetSmalDataRateControlOk() (*SmallDataRateControl, bool)`

GetSmalDataRateControlOk returns a tuple with the SmalDataRateControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmalDataRateControl

`func (o *SmContextConfiguration) SetSmalDataRateControl(v SmallDataRateControl)`

SetSmalDataRateControl sets SmalDataRateControl field to given value.

### HasSmalDataRateControl

`func (o *SmContextConfiguration) HasSmalDataRateControl() bool`

HasSmalDataRateControl returns a boolean if a field has been set.

### GetSmallDataRateStatus

`func (o *SmContextConfiguration) GetSmallDataRateStatus() SmallDataRateStatus`

GetSmallDataRateStatus returns the SmallDataRateStatus field if non-nil, zero value otherwise.

### GetSmallDataRateStatusOk

`func (o *SmContextConfiguration) GetSmallDataRateStatusOk() (*SmallDataRateStatus, bool)`

GetSmallDataRateStatusOk returns a tuple with the SmallDataRateStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmallDataRateStatus

`func (o *SmContextConfiguration) SetSmallDataRateStatus(v SmallDataRateStatus)`

SetSmallDataRateStatus sets SmallDataRateStatus field to given value.

### HasSmallDataRateStatus

`func (o *SmContextConfiguration) HasSmallDataRateStatus() bool`

HasSmallDataRateStatus returns a boolean if a field has been set.

### GetServPlmnDataRateCtl

`func (o *SmContextConfiguration) GetServPlmnDataRateCtl() int32`

GetServPlmnDataRateCtl returns the ServPlmnDataRateCtl field if non-nil, zero value otherwise.

### GetServPlmnDataRateCtlOk

`func (o *SmContextConfiguration) GetServPlmnDataRateCtlOk() (*int32, bool)`

GetServPlmnDataRateCtlOk returns a tuple with the ServPlmnDataRateCtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServPlmnDataRateCtl

`func (o *SmContextConfiguration) SetServPlmnDataRateCtl(v int32)`

SetServPlmnDataRateCtl sets ServPlmnDataRateCtl field to given value.

### HasServPlmnDataRateCtl

`func (o *SmContextConfiguration) HasServPlmnDataRateCtl() bool`

HasServPlmnDataRateCtl returns a boolean if a field has been set.

### SetServPlmnDataRateCtlNil

`func (o *SmContextConfiguration) SetServPlmnDataRateCtlNil(b bool)`

 SetServPlmnDataRateCtlNil sets the value for ServPlmnDataRateCtl to be an explicit nil

### UnsetServPlmnDataRateCtl
`func (o *SmContextConfiguration) UnsetServPlmnDataRateCtl()`

UnsetServPlmnDataRateCtl ensures that no value is present for ServPlmnDataRateCtl, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


