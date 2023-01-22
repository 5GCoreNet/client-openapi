# GuamiRm

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PlmnId** | [**PlmnIdNid**](PlmnIdNid.md) |  | 
**AmfId** | **string** | String identifying the AMF ID composed of AMF Region ID (8 bits), AMF Set ID (10 bits) and AMF  Pointer (6 bits) as specified in clause 2.10.1 of 3GPP TS 23.003. It is encoded as a string of  6 hexadecimal characters (i.e., 24 bits).   | 

## Methods

### NewGuamiRm

`func NewGuamiRm(plmnId PlmnIdNid, amfId string, ) *GuamiRm`

NewGuamiRm instantiates a new GuamiRm object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGuamiRmWithDefaults

`func NewGuamiRmWithDefaults() *GuamiRm`

NewGuamiRmWithDefaults instantiates a new GuamiRm object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPlmnId

`func (o *GuamiRm) GetPlmnId() PlmnIdNid`

GetPlmnId returns the PlmnId field if non-nil, zero value otherwise.

### GetPlmnIdOk

`func (o *GuamiRm) GetPlmnIdOk() (*PlmnIdNid, bool)`

GetPlmnIdOk returns a tuple with the PlmnId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlmnId

`func (o *GuamiRm) SetPlmnId(v PlmnIdNid)`

SetPlmnId sets PlmnId field to given value.


### GetAmfId

`func (o *GuamiRm) GetAmfId() string`

GetAmfId returns the AmfId field if non-nil, zero value otherwise.

### GetAmfIdOk

`func (o *GuamiRm) GetAmfIdOk() (*string, bool)`

GetAmfIdOk returns a tuple with the AmfId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmfId

`func (o *GuamiRm) SetAmfId(v string)`

SetAmfId sets AmfId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


