# NSMChargingInformation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ManagementOperation** | [**ManagementOperation**](ManagementOperation.md) |  | 
**IdNetworkSliceInstance** | Pointer to **string** |  | [optional] 
**ListOfserviceProfileChargingInformation** | Pointer to [**[]ServiceProfileChargingInformation**](ServiceProfileChargingInformation.md) |  | [optional] 
**ManagementOperationStatus** | Pointer to [**ManagementOperationStatus**](ManagementOperationStatus.md) |  | [optional] 

## Methods

### NewNSMChargingInformation

`func NewNSMChargingInformation(managementOperation ManagementOperation, ) *NSMChargingInformation`

NewNSMChargingInformation instantiates a new NSMChargingInformation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNSMChargingInformationWithDefaults

`func NewNSMChargingInformationWithDefaults() *NSMChargingInformation`

NewNSMChargingInformationWithDefaults instantiates a new NSMChargingInformation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetManagementOperation

`func (o *NSMChargingInformation) GetManagementOperation() ManagementOperation`

GetManagementOperation returns the ManagementOperation field if non-nil, zero value otherwise.

### GetManagementOperationOk

`func (o *NSMChargingInformation) GetManagementOperationOk() (*ManagementOperation, bool)`

GetManagementOperationOk returns a tuple with the ManagementOperation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementOperation

`func (o *NSMChargingInformation) SetManagementOperation(v ManagementOperation)`

SetManagementOperation sets ManagementOperation field to given value.


### GetIdNetworkSliceInstance

`func (o *NSMChargingInformation) GetIdNetworkSliceInstance() string`

GetIdNetworkSliceInstance returns the IdNetworkSliceInstance field if non-nil, zero value otherwise.

### GetIdNetworkSliceInstanceOk

`func (o *NSMChargingInformation) GetIdNetworkSliceInstanceOk() (*string, bool)`

GetIdNetworkSliceInstanceOk returns a tuple with the IdNetworkSliceInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdNetworkSliceInstance

`func (o *NSMChargingInformation) SetIdNetworkSliceInstance(v string)`

SetIdNetworkSliceInstance sets IdNetworkSliceInstance field to given value.

### HasIdNetworkSliceInstance

`func (o *NSMChargingInformation) HasIdNetworkSliceInstance() bool`

HasIdNetworkSliceInstance returns a boolean if a field has been set.

### GetListOfserviceProfileChargingInformation

`func (o *NSMChargingInformation) GetListOfserviceProfileChargingInformation() []ServiceProfileChargingInformation`

GetListOfserviceProfileChargingInformation returns the ListOfserviceProfileChargingInformation field if non-nil, zero value otherwise.

### GetListOfserviceProfileChargingInformationOk

`func (o *NSMChargingInformation) GetListOfserviceProfileChargingInformationOk() (*[]ServiceProfileChargingInformation, bool)`

GetListOfserviceProfileChargingInformationOk returns a tuple with the ListOfserviceProfileChargingInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListOfserviceProfileChargingInformation

`func (o *NSMChargingInformation) SetListOfserviceProfileChargingInformation(v []ServiceProfileChargingInformation)`

SetListOfserviceProfileChargingInformation sets ListOfserviceProfileChargingInformation field to given value.

### HasListOfserviceProfileChargingInformation

`func (o *NSMChargingInformation) HasListOfserviceProfileChargingInformation() bool`

HasListOfserviceProfileChargingInformation returns a boolean if a field has been set.

### GetManagementOperationStatus

`func (o *NSMChargingInformation) GetManagementOperationStatus() ManagementOperationStatus`

GetManagementOperationStatus returns the ManagementOperationStatus field if non-nil, zero value otherwise.

### GetManagementOperationStatusOk

`func (o *NSMChargingInformation) GetManagementOperationStatusOk() (*ManagementOperationStatus, bool)`

GetManagementOperationStatusOk returns a tuple with the ManagementOperationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementOperationStatus

`func (o *NSMChargingInformation) SetManagementOperationStatus(v ManagementOperationStatus)`

SetManagementOperationStatus sets ManagementOperationStatus field to given value.

### HasManagementOperationStatus

`func (o *NSMChargingInformation) HasManagementOperationStatus() bool`

HasManagementOperationStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


