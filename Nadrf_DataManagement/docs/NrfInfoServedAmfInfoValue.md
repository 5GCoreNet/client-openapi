# NrfInfoServedAmfInfoValue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AmfSetId** | **string** | String identifying the AMF Set ID (10 bits) as specified in clause 2.10.1 of 3GPP TS 23.003.  It is encoded as a string of 3 hexadecimal characters where the first character is limited to  values 0 to 3 (i.e. 10 bits).  | 
**AmfRegionId** | **string** | String identifying the AMF Set ID (10 bits) as specified in clause 2.10.1 of 3GPP TS 23.003.  It is encoded as a string of 3 hexadecimal characters where the first character is limited to  values 0 to 3 (i.e. 10 bits)  | 
**GuamiList** | [**[]Guami**](Guami.md) |  | 
**TaiList** | Pointer to [**[]Tai**](Tai.md) |  | [optional] 
**TaiRangeList** | Pointer to [**[]TaiRange**](TaiRange.md) |  | [optional] 
**BackupInfoAmfFailure** | Pointer to [**[]Guami**](Guami.md) |  | [optional] 
**BackupInfoAmfRemoval** | Pointer to [**[]Guami**](Guami.md) |  | [optional] 
**N2InterfaceAmfInfo** | Pointer to [**N2InterfaceAmfInfo**](N2InterfaceAmfInfo.md) |  | [optional] 
**AmfOnboardingCapability** | Pointer to **bool** |  | [optional] [default to false]

## Methods

### NewNrfInfoServedAmfInfoValue

`func NewNrfInfoServedAmfInfoValue(amfSetId string, amfRegionId string, guamiList []Guami, ) *NrfInfoServedAmfInfoValue`

NewNrfInfoServedAmfInfoValue instantiates a new NrfInfoServedAmfInfoValue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNrfInfoServedAmfInfoValueWithDefaults

`func NewNrfInfoServedAmfInfoValueWithDefaults() *NrfInfoServedAmfInfoValue`

NewNrfInfoServedAmfInfoValueWithDefaults instantiates a new NrfInfoServedAmfInfoValue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmfSetId

`func (o *NrfInfoServedAmfInfoValue) GetAmfSetId() string`

GetAmfSetId returns the AmfSetId field if non-nil, zero value otherwise.

### GetAmfSetIdOk

`func (o *NrfInfoServedAmfInfoValue) GetAmfSetIdOk() (*string, bool)`

GetAmfSetIdOk returns a tuple with the AmfSetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmfSetId

`func (o *NrfInfoServedAmfInfoValue) SetAmfSetId(v string)`

SetAmfSetId sets AmfSetId field to given value.


### GetAmfRegionId

`func (o *NrfInfoServedAmfInfoValue) GetAmfRegionId() string`

GetAmfRegionId returns the AmfRegionId field if non-nil, zero value otherwise.

### GetAmfRegionIdOk

`func (o *NrfInfoServedAmfInfoValue) GetAmfRegionIdOk() (*string, bool)`

GetAmfRegionIdOk returns a tuple with the AmfRegionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmfRegionId

`func (o *NrfInfoServedAmfInfoValue) SetAmfRegionId(v string)`

SetAmfRegionId sets AmfRegionId field to given value.


### GetGuamiList

`func (o *NrfInfoServedAmfInfoValue) GetGuamiList() []Guami`

GetGuamiList returns the GuamiList field if non-nil, zero value otherwise.

### GetGuamiListOk

`func (o *NrfInfoServedAmfInfoValue) GetGuamiListOk() (*[]Guami, bool)`

GetGuamiListOk returns a tuple with the GuamiList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuamiList

`func (o *NrfInfoServedAmfInfoValue) SetGuamiList(v []Guami)`

SetGuamiList sets GuamiList field to given value.


### GetTaiList

`func (o *NrfInfoServedAmfInfoValue) GetTaiList() []Tai`

GetTaiList returns the TaiList field if non-nil, zero value otherwise.

### GetTaiListOk

`func (o *NrfInfoServedAmfInfoValue) GetTaiListOk() (*[]Tai, bool)`

GetTaiListOk returns a tuple with the TaiList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaiList

`func (o *NrfInfoServedAmfInfoValue) SetTaiList(v []Tai)`

SetTaiList sets TaiList field to given value.

### HasTaiList

`func (o *NrfInfoServedAmfInfoValue) HasTaiList() bool`

HasTaiList returns a boolean if a field has been set.

### GetTaiRangeList

`func (o *NrfInfoServedAmfInfoValue) GetTaiRangeList() []TaiRange`

GetTaiRangeList returns the TaiRangeList field if non-nil, zero value otherwise.

### GetTaiRangeListOk

`func (o *NrfInfoServedAmfInfoValue) GetTaiRangeListOk() (*[]TaiRange, bool)`

GetTaiRangeListOk returns a tuple with the TaiRangeList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaiRangeList

`func (o *NrfInfoServedAmfInfoValue) SetTaiRangeList(v []TaiRange)`

SetTaiRangeList sets TaiRangeList field to given value.

### HasTaiRangeList

`func (o *NrfInfoServedAmfInfoValue) HasTaiRangeList() bool`

HasTaiRangeList returns a boolean if a field has been set.

### GetBackupInfoAmfFailure

`func (o *NrfInfoServedAmfInfoValue) GetBackupInfoAmfFailure() []Guami`

GetBackupInfoAmfFailure returns the BackupInfoAmfFailure field if non-nil, zero value otherwise.

### GetBackupInfoAmfFailureOk

`func (o *NrfInfoServedAmfInfoValue) GetBackupInfoAmfFailureOk() (*[]Guami, bool)`

GetBackupInfoAmfFailureOk returns a tuple with the BackupInfoAmfFailure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupInfoAmfFailure

`func (o *NrfInfoServedAmfInfoValue) SetBackupInfoAmfFailure(v []Guami)`

SetBackupInfoAmfFailure sets BackupInfoAmfFailure field to given value.

### HasBackupInfoAmfFailure

`func (o *NrfInfoServedAmfInfoValue) HasBackupInfoAmfFailure() bool`

HasBackupInfoAmfFailure returns a boolean if a field has been set.

### GetBackupInfoAmfRemoval

`func (o *NrfInfoServedAmfInfoValue) GetBackupInfoAmfRemoval() []Guami`

GetBackupInfoAmfRemoval returns the BackupInfoAmfRemoval field if non-nil, zero value otherwise.

### GetBackupInfoAmfRemovalOk

`func (o *NrfInfoServedAmfInfoValue) GetBackupInfoAmfRemovalOk() (*[]Guami, bool)`

GetBackupInfoAmfRemovalOk returns a tuple with the BackupInfoAmfRemoval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupInfoAmfRemoval

`func (o *NrfInfoServedAmfInfoValue) SetBackupInfoAmfRemoval(v []Guami)`

SetBackupInfoAmfRemoval sets BackupInfoAmfRemoval field to given value.

### HasBackupInfoAmfRemoval

`func (o *NrfInfoServedAmfInfoValue) HasBackupInfoAmfRemoval() bool`

HasBackupInfoAmfRemoval returns a boolean if a field has been set.

### GetN2InterfaceAmfInfo

`func (o *NrfInfoServedAmfInfoValue) GetN2InterfaceAmfInfo() N2InterfaceAmfInfo`

GetN2InterfaceAmfInfo returns the N2InterfaceAmfInfo field if non-nil, zero value otherwise.

### GetN2InterfaceAmfInfoOk

`func (o *NrfInfoServedAmfInfoValue) GetN2InterfaceAmfInfoOk() (*N2InterfaceAmfInfo, bool)`

GetN2InterfaceAmfInfoOk returns a tuple with the N2InterfaceAmfInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetN2InterfaceAmfInfo

`func (o *NrfInfoServedAmfInfoValue) SetN2InterfaceAmfInfo(v N2InterfaceAmfInfo)`

SetN2InterfaceAmfInfo sets N2InterfaceAmfInfo field to given value.

### HasN2InterfaceAmfInfo

`func (o *NrfInfoServedAmfInfoValue) HasN2InterfaceAmfInfo() bool`

HasN2InterfaceAmfInfo returns a boolean if a field has been set.

### GetAmfOnboardingCapability

`func (o *NrfInfoServedAmfInfoValue) GetAmfOnboardingCapability() bool`

GetAmfOnboardingCapability returns the AmfOnboardingCapability field if non-nil, zero value otherwise.

### GetAmfOnboardingCapabilityOk

`func (o *NrfInfoServedAmfInfoValue) GetAmfOnboardingCapabilityOk() (*bool, bool)`

GetAmfOnboardingCapabilityOk returns a tuple with the AmfOnboardingCapability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmfOnboardingCapability

`func (o *NrfInfoServedAmfInfoValue) SetAmfOnboardingCapability(v bool)`

SetAmfOnboardingCapability sets AmfOnboardingCapability field to given value.

### HasAmfOnboardingCapability

`func (o *NrfInfoServedAmfInfoValue) HasAmfOnboardingCapability() bool`

HasAmfOnboardingCapability returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


