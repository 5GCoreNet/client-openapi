# AmfInfo

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

### NewAmfInfo

`func NewAmfInfo(amfSetId string, amfRegionId string, guamiList []Guami, ) *AmfInfo`

NewAmfInfo instantiates a new AmfInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAmfInfoWithDefaults

`func NewAmfInfoWithDefaults() *AmfInfo`

NewAmfInfoWithDefaults instantiates a new AmfInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmfSetId

`func (o *AmfInfo) GetAmfSetId() string`

GetAmfSetId returns the AmfSetId field if non-nil, zero value otherwise.

### GetAmfSetIdOk

`func (o *AmfInfo) GetAmfSetIdOk() (*string, bool)`

GetAmfSetIdOk returns a tuple with the AmfSetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmfSetId

`func (o *AmfInfo) SetAmfSetId(v string)`

SetAmfSetId sets AmfSetId field to given value.


### GetAmfRegionId

`func (o *AmfInfo) GetAmfRegionId() string`

GetAmfRegionId returns the AmfRegionId field if non-nil, zero value otherwise.

### GetAmfRegionIdOk

`func (o *AmfInfo) GetAmfRegionIdOk() (*string, bool)`

GetAmfRegionIdOk returns a tuple with the AmfRegionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmfRegionId

`func (o *AmfInfo) SetAmfRegionId(v string)`

SetAmfRegionId sets AmfRegionId field to given value.


### GetGuamiList

`func (o *AmfInfo) GetGuamiList() []Guami`

GetGuamiList returns the GuamiList field if non-nil, zero value otherwise.

### GetGuamiListOk

`func (o *AmfInfo) GetGuamiListOk() (*[]Guami, bool)`

GetGuamiListOk returns a tuple with the GuamiList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuamiList

`func (o *AmfInfo) SetGuamiList(v []Guami)`

SetGuamiList sets GuamiList field to given value.


### GetTaiList

`func (o *AmfInfo) GetTaiList() []Tai`

GetTaiList returns the TaiList field if non-nil, zero value otherwise.

### GetTaiListOk

`func (o *AmfInfo) GetTaiListOk() (*[]Tai, bool)`

GetTaiListOk returns a tuple with the TaiList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaiList

`func (o *AmfInfo) SetTaiList(v []Tai)`

SetTaiList sets TaiList field to given value.

### HasTaiList

`func (o *AmfInfo) HasTaiList() bool`

HasTaiList returns a boolean if a field has been set.

### GetTaiRangeList

`func (o *AmfInfo) GetTaiRangeList() []TaiRange`

GetTaiRangeList returns the TaiRangeList field if non-nil, zero value otherwise.

### GetTaiRangeListOk

`func (o *AmfInfo) GetTaiRangeListOk() (*[]TaiRange, bool)`

GetTaiRangeListOk returns a tuple with the TaiRangeList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaiRangeList

`func (o *AmfInfo) SetTaiRangeList(v []TaiRange)`

SetTaiRangeList sets TaiRangeList field to given value.

### HasTaiRangeList

`func (o *AmfInfo) HasTaiRangeList() bool`

HasTaiRangeList returns a boolean if a field has been set.

### GetBackupInfoAmfFailure

`func (o *AmfInfo) GetBackupInfoAmfFailure() []Guami`

GetBackupInfoAmfFailure returns the BackupInfoAmfFailure field if non-nil, zero value otherwise.

### GetBackupInfoAmfFailureOk

`func (o *AmfInfo) GetBackupInfoAmfFailureOk() (*[]Guami, bool)`

GetBackupInfoAmfFailureOk returns a tuple with the BackupInfoAmfFailure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupInfoAmfFailure

`func (o *AmfInfo) SetBackupInfoAmfFailure(v []Guami)`

SetBackupInfoAmfFailure sets BackupInfoAmfFailure field to given value.

### HasBackupInfoAmfFailure

`func (o *AmfInfo) HasBackupInfoAmfFailure() bool`

HasBackupInfoAmfFailure returns a boolean if a field has been set.

### GetBackupInfoAmfRemoval

`func (o *AmfInfo) GetBackupInfoAmfRemoval() []Guami`

GetBackupInfoAmfRemoval returns the BackupInfoAmfRemoval field if non-nil, zero value otherwise.

### GetBackupInfoAmfRemovalOk

`func (o *AmfInfo) GetBackupInfoAmfRemovalOk() (*[]Guami, bool)`

GetBackupInfoAmfRemovalOk returns a tuple with the BackupInfoAmfRemoval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupInfoAmfRemoval

`func (o *AmfInfo) SetBackupInfoAmfRemoval(v []Guami)`

SetBackupInfoAmfRemoval sets BackupInfoAmfRemoval field to given value.

### HasBackupInfoAmfRemoval

`func (o *AmfInfo) HasBackupInfoAmfRemoval() bool`

HasBackupInfoAmfRemoval returns a boolean if a field has been set.

### GetN2InterfaceAmfInfo

`func (o *AmfInfo) GetN2InterfaceAmfInfo() N2InterfaceAmfInfo`

GetN2InterfaceAmfInfo returns the N2InterfaceAmfInfo field if non-nil, zero value otherwise.

### GetN2InterfaceAmfInfoOk

`func (o *AmfInfo) GetN2InterfaceAmfInfoOk() (*N2InterfaceAmfInfo, bool)`

GetN2InterfaceAmfInfoOk returns a tuple with the N2InterfaceAmfInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetN2InterfaceAmfInfo

`func (o *AmfInfo) SetN2InterfaceAmfInfo(v N2InterfaceAmfInfo)`

SetN2InterfaceAmfInfo sets N2InterfaceAmfInfo field to given value.

### HasN2InterfaceAmfInfo

`func (o *AmfInfo) HasN2InterfaceAmfInfo() bool`

HasN2InterfaceAmfInfo returns a boolean if a field has been set.

### GetAmfOnboardingCapability

`func (o *AmfInfo) GetAmfOnboardingCapability() bool`

GetAmfOnboardingCapability returns the AmfOnboardingCapability field if non-nil, zero value otherwise.

### GetAmfOnboardingCapabilityOk

`func (o *AmfInfo) GetAmfOnboardingCapabilityOk() (*bool, bool)`

GetAmfOnboardingCapabilityOk returns a tuple with the AmfOnboardingCapability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmfOnboardingCapability

`func (o *AmfInfo) SetAmfOnboardingCapability(v bool)`

SetAmfOnboardingCapability sets AmfOnboardingCapability field to given value.

### HasAmfOnboardingCapability

`func (o *AmfInfo) HasAmfOnboardingCapability() bool`

HasAmfOnboardingCapability returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


