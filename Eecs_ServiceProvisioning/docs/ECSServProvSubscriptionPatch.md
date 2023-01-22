# ECSServProvSubscriptionPatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AcProfs** | Pointer to [**[]ACProfile**](ACProfile.md) | Information about services the EEC wants to connect to. | [optional] 
**ExpTime** | Pointer to **time.Time** | string with format \&quot;date-time\&quot; as defined in OpenAPI. | [optional] 
**EecSvcContSupp** | Pointer to [**[]ACRScenario**](ACRScenario.md) | Indicates if the EEC supports service continuity or not, also indicates which ACR scenarios are supported by the EEC.  | [optional] 
**ConnInfo** | Pointer to [**[]ConnectivityInfo**](ConnectivityInfo.md) | List of connectivity information for the UE. | [optional] 

## Methods

### NewECSServProvSubscriptionPatch

`func NewECSServProvSubscriptionPatch() *ECSServProvSubscriptionPatch`

NewECSServProvSubscriptionPatch instantiates a new ECSServProvSubscriptionPatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewECSServProvSubscriptionPatchWithDefaults

`func NewECSServProvSubscriptionPatchWithDefaults() *ECSServProvSubscriptionPatch`

NewECSServProvSubscriptionPatchWithDefaults instantiates a new ECSServProvSubscriptionPatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAcProfs

`func (o *ECSServProvSubscriptionPatch) GetAcProfs() []ACProfile`

GetAcProfs returns the AcProfs field if non-nil, zero value otherwise.

### GetAcProfsOk

`func (o *ECSServProvSubscriptionPatch) GetAcProfsOk() (*[]ACProfile, bool)`

GetAcProfsOk returns a tuple with the AcProfs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcProfs

`func (o *ECSServProvSubscriptionPatch) SetAcProfs(v []ACProfile)`

SetAcProfs sets AcProfs field to given value.

### HasAcProfs

`func (o *ECSServProvSubscriptionPatch) HasAcProfs() bool`

HasAcProfs returns a boolean if a field has been set.

### GetExpTime

`func (o *ECSServProvSubscriptionPatch) GetExpTime() time.Time`

GetExpTime returns the ExpTime field if non-nil, zero value otherwise.

### GetExpTimeOk

`func (o *ECSServProvSubscriptionPatch) GetExpTimeOk() (*time.Time, bool)`

GetExpTimeOk returns a tuple with the ExpTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpTime

`func (o *ECSServProvSubscriptionPatch) SetExpTime(v time.Time)`

SetExpTime sets ExpTime field to given value.

### HasExpTime

`func (o *ECSServProvSubscriptionPatch) HasExpTime() bool`

HasExpTime returns a boolean if a field has been set.

### GetEecSvcContSupp

`func (o *ECSServProvSubscriptionPatch) GetEecSvcContSupp() []ACRScenario`

GetEecSvcContSupp returns the EecSvcContSupp field if non-nil, zero value otherwise.

### GetEecSvcContSuppOk

`func (o *ECSServProvSubscriptionPatch) GetEecSvcContSuppOk() (*[]ACRScenario, bool)`

GetEecSvcContSuppOk returns a tuple with the EecSvcContSupp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEecSvcContSupp

`func (o *ECSServProvSubscriptionPatch) SetEecSvcContSupp(v []ACRScenario)`

SetEecSvcContSupp sets EecSvcContSupp field to given value.

### HasEecSvcContSupp

`func (o *ECSServProvSubscriptionPatch) HasEecSvcContSupp() bool`

HasEecSvcContSupp returns a boolean if a field has been set.

### GetConnInfo

`func (o *ECSServProvSubscriptionPatch) GetConnInfo() []ConnectivityInfo`

GetConnInfo returns the ConnInfo field if non-nil, zero value otherwise.

### GetConnInfoOk

`func (o *ECSServProvSubscriptionPatch) GetConnInfoOk() (*[]ConnectivityInfo, bool)`

GetConnInfoOk returns a tuple with the ConnInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnInfo

`func (o *ECSServProvSubscriptionPatch) SetConnInfo(v []ConnectivityInfo)`

SetConnInfo sets ConnInfo field to given value.

### HasConnInfo

`func (o *ECSServProvSubscriptionPatch) HasConnInfo() bool`

HasConnInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


