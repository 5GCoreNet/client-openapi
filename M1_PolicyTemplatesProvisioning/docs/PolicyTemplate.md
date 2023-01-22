# PolicyTemplate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PolicyTemplateId** | **string** | String chosen by the 5GMS AF to serve as an identifier in a resource URI. | 
**State** | [**PolicyTemplateState**](PolicyTemplateState.md) |  | 
**ApiEndPoint** | **string** |  | 
**ApiType** | [**PolicyTemplateApiType**](PolicyTemplateApiType.md) |  | 
**ExternalReference** | **string** |  | 
**QoSSpecification** | Pointer to [**M1QoSSpecification**](M1QoSSpecification.md) |  | [optional] 
**ApplicationSessionContext** | [**PolicyTemplateApplicationSessionContext**](PolicyTemplateApplicationSessionContext.md) |  | 
**ChargingSpecification** | Pointer to [**ChargingSpecification**](ChargingSpecification.md) |  | [optional] 

## Methods

### NewPolicyTemplate

`func NewPolicyTemplate(policyTemplateId string, state PolicyTemplateState, apiEndPoint string, apiType PolicyTemplateApiType, externalReference string, applicationSessionContext PolicyTemplateApplicationSessionContext, ) *PolicyTemplate`

NewPolicyTemplate instantiates a new PolicyTemplate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicyTemplateWithDefaults

`func NewPolicyTemplateWithDefaults() *PolicyTemplate`

NewPolicyTemplateWithDefaults instantiates a new PolicyTemplate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPolicyTemplateId

`func (o *PolicyTemplate) GetPolicyTemplateId() string`

GetPolicyTemplateId returns the PolicyTemplateId field if non-nil, zero value otherwise.

### GetPolicyTemplateIdOk

`func (o *PolicyTemplate) GetPolicyTemplateIdOk() (*string, bool)`

GetPolicyTemplateIdOk returns a tuple with the PolicyTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyTemplateId

`func (o *PolicyTemplate) SetPolicyTemplateId(v string)`

SetPolicyTemplateId sets PolicyTemplateId field to given value.


### GetState

`func (o *PolicyTemplate) GetState() PolicyTemplateState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *PolicyTemplate) GetStateOk() (*PolicyTemplateState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *PolicyTemplate) SetState(v PolicyTemplateState)`

SetState sets State field to given value.


### GetApiEndPoint

`func (o *PolicyTemplate) GetApiEndPoint() string`

GetApiEndPoint returns the ApiEndPoint field if non-nil, zero value otherwise.

### GetApiEndPointOk

`func (o *PolicyTemplate) GetApiEndPointOk() (*string, bool)`

GetApiEndPointOk returns a tuple with the ApiEndPoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiEndPoint

`func (o *PolicyTemplate) SetApiEndPoint(v string)`

SetApiEndPoint sets ApiEndPoint field to given value.


### GetApiType

`func (o *PolicyTemplate) GetApiType() PolicyTemplateApiType`

GetApiType returns the ApiType field if non-nil, zero value otherwise.

### GetApiTypeOk

`func (o *PolicyTemplate) GetApiTypeOk() (*PolicyTemplateApiType, bool)`

GetApiTypeOk returns a tuple with the ApiType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiType

`func (o *PolicyTemplate) SetApiType(v PolicyTemplateApiType)`

SetApiType sets ApiType field to given value.


### GetExternalReference

`func (o *PolicyTemplate) GetExternalReference() string`

GetExternalReference returns the ExternalReference field if non-nil, zero value otherwise.

### GetExternalReferenceOk

`func (o *PolicyTemplate) GetExternalReferenceOk() (*string, bool)`

GetExternalReferenceOk returns a tuple with the ExternalReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalReference

`func (o *PolicyTemplate) SetExternalReference(v string)`

SetExternalReference sets ExternalReference field to given value.


### GetQoSSpecification

`func (o *PolicyTemplate) GetQoSSpecification() M1QoSSpecification`

GetQoSSpecification returns the QoSSpecification field if non-nil, zero value otherwise.

### GetQoSSpecificationOk

`func (o *PolicyTemplate) GetQoSSpecificationOk() (*M1QoSSpecification, bool)`

GetQoSSpecificationOk returns a tuple with the QoSSpecification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQoSSpecification

`func (o *PolicyTemplate) SetQoSSpecification(v M1QoSSpecification)`

SetQoSSpecification sets QoSSpecification field to given value.

### HasQoSSpecification

`func (o *PolicyTemplate) HasQoSSpecification() bool`

HasQoSSpecification returns a boolean if a field has been set.

### GetApplicationSessionContext

`func (o *PolicyTemplate) GetApplicationSessionContext() PolicyTemplateApplicationSessionContext`

GetApplicationSessionContext returns the ApplicationSessionContext field if non-nil, zero value otherwise.

### GetApplicationSessionContextOk

`func (o *PolicyTemplate) GetApplicationSessionContextOk() (*PolicyTemplateApplicationSessionContext, bool)`

GetApplicationSessionContextOk returns a tuple with the ApplicationSessionContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationSessionContext

`func (o *PolicyTemplate) SetApplicationSessionContext(v PolicyTemplateApplicationSessionContext)`

SetApplicationSessionContext sets ApplicationSessionContext field to given value.


### GetChargingSpecification

`func (o *PolicyTemplate) GetChargingSpecification() ChargingSpecification`

GetChargingSpecification returns the ChargingSpecification field if non-nil, zero value otherwise.

### GetChargingSpecificationOk

`func (o *PolicyTemplate) GetChargingSpecificationOk() (*ChargingSpecification, bool)`

GetChargingSpecificationOk returns a tuple with the ChargingSpecification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChargingSpecification

`func (o *PolicyTemplate) SetChargingSpecification(v ChargingSpecification)`

SetChargingSpecification sets ChargingSpecification field to given value.

### HasChargingSpecification

`func (o *PolicyTemplate) HasChargingSpecification() bool`

HasChargingSpecification returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


