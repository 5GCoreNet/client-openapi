# NetworkAssistanceSession

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NaSessionId** | **string** | String chosen by the 5GMS AF to serve as an identifier in a resource URI. | 
**ServiceDataFlowDescription** | Pointer to [**[]ServiceDataFlowDescription**](ServiceDataFlowDescription.md) |  | [optional] 
**PolicyTemplateId** | Pointer to **string** | String chosen by the 5GMS AF to serve as an identifier in a resource URI. | [optional] 
**RequestedQoS** | Pointer to [**M5QoSSpecification**](M5QoSSpecification.md) |  | [optional] 
**RecommendedQoS** | Pointer to [**M5QoSSpecification**](M5QoSSpecification.md) |  | [optional] 
**NotficationURL** | Pointer to **string** | Uniform Resource Locator, comforming with the URI Generic Syntax specified in IETF RFC 3986. | [optional] 

## Methods

### NewNetworkAssistanceSession

`func NewNetworkAssistanceSession(naSessionId string, ) *NetworkAssistanceSession`

NewNetworkAssistanceSession instantiates a new NetworkAssistanceSession object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkAssistanceSessionWithDefaults

`func NewNetworkAssistanceSessionWithDefaults() *NetworkAssistanceSession`

NewNetworkAssistanceSessionWithDefaults instantiates a new NetworkAssistanceSession object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNaSessionId

`func (o *NetworkAssistanceSession) GetNaSessionId() string`

GetNaSessionId returns the NaSessionId field if non-nil, zero value otherwise.

### GetNaSessionIdOk

`func (o *NetworkAssistanceSession) GetNaSessionIdOk() (*string, bool)`

GetNaSessionIdOk returns a tuple with the NaSessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNaSessionId

`func (o *NetworkAssistanceSession) SetNaSessionId(v string)`

SetNaSessionId sets NaSessionId field to given value.


### GetServiceDataFlowDescription

`func (o *NetworkAssistanceSession) GetServiceDataFlowDescription() []ServiceDataFlowDescription`

GetServiceDataFlowDescription returns the ServiceDataFlowDescription field if non-nil, zero value otherwise.

### GetServiceDataFlowDescriptionOk

`func (o *NetworkAssistanceSession) GetServiceDataFlowDescriptionOk() (*[]ServiceDataFlowDescription, bool)`

GetServiceDataFlowDescriptionOk returns a tuple with the ServiceDataFlowDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceDataFlowDescription

`func (o *NetworkAssistanceSession) SetServiceDataFlowDescription(v []ServiceDataFlowDescription)`

SetServiceDataFlowDescription sets ServiceDataFlowDescription field to given value.

### HasServiceDataFlowDescription

`func (o *NetworkAssistanceSession) HasServiceDataFlowDescription() bool`

HasServiceDataFlowDescription returns a boolean if a field has been set.

### GetPolicyTemplateId

`func (o *NetworkAssistanceSession) GetPolicyTemplateId() string`

GetPolicyTemplateId returns the PolicyTemplateId field if non-nil, zero value otherwise.

### GetPolicyTemplateIdOk

`func (o *NetworkAssistanceSession) GetPolicyTemplateIdOk() (*string, bool)`

GetPolicyTemplateIdOk returns a tuple with the PolicyTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyTemplateId

`func (o *NetworkAssistanceSession) SetPolicyTemplateId(v string)`

SetPolicyTemplateId sets PolicyTemplateId field to given value.

### HasPolicyTemplateId

`func (o *NetworkAssistanceSession) HasPolicyTemplateId() bool`

HasPolicyTemplateId returns a boolean if a field has been set.

### GetRequestedQoS

`func (o *NetworkAssistanceSession) GetRequestedQoS() M5QoSSpecification`

GetRequestedQoS returns the RequestedQoS field if non-nil, zero value otherwise.

### GetRequestedQoSOk

`func (o *NetworkAssistanceSession) GetRequestedQoSOk() (*M5QoSSpecification, bool)`

GetRequestedQoSOk returns a tuple with the RequestedQoS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedQoS

`func (o *NetworkAssistanceSession) SetRequestedQoS(v M5QoSSpecification)`

SetRequestedQoS sets RequestedQoS field to given value.

### HasRequestedQoS

`func (o *NetworkAssistanceSession) HasRequestedQoS() bool`

HasRequestedQoS returns a boolean if a field has been set.

### GetRecommendedQoS

`func (o *NetworkAssistanceSession) GetRecommendedQoS() M5QoSSpecification`

GetRecommendedQoS returns the RecommendedQoS field if non-nil, zero value otherwise.

### GetRecommendedQoSOk

`func (o *NetworkAssistanceSession) GetRecommendedQoSOk() (*M5QoSSpecification, bool)`

GetRecommendedQoSOk returns a tuple with the RecommendedQoS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecommendedQoS

`func (o *NetworkAssistanceSession) SetRecommendedQoS(v M5QoSSpecification)`

SetRecommendedQoS sets RecommendedQoS field to given value.

### HasRecommendedQoS

`func (o *NetworkAssistanceSession) HasRecommendedQoS() bool`

HasRecommendedQoS returns a boolean if a field has been set.

### GetNotficationURL

`func (o *NetworkAssistanceSession) GetNotficationURL() string`

GetNotficationURL returns the NotficationURL field if non-nil, zero value otherwise.

### GetNotficationURLOk

`func (o *NetworkAssistanceSession) GetNotficationURLOk() (*string, bool)`

GetNotficationURLOk returns a tuple with the NotficationURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotficationURL

`func (o *NetworkAssistanceSession) SetNotficationURL(v string)`

SetNotficationURL sets NotficationURL field to given value.

### HasNotficationURL

`func (o *NetworkAssistanceSession) HasNotficationURL() bool`

HasNotficationURL returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


