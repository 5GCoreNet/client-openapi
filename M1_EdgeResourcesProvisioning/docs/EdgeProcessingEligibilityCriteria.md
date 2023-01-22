# EdgeProcessingEligibilityCriteria

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServiceDataFlowDescriptions** | [**[]ServiceDataFlowDescription**](ServiceDataFlowDescription.md) |  | 
**UeLocations** | [**[]LocationArea5G**](LocationArea5G.md) |  | 
**TimeWindows** | [**[]TimeWindow**](TimeWindow.md) |  | 
**AppRequest** | **bool** |  | 

## Methods

### NewEdgeProcessingEligibilityCriteria

`func NewEdgeProcessingEligibilityCriteria(serviceDataFlowDescriptions []ServiceDataFlowDescription, ueLocations []LocationArea5G, timeWindows []TimeWindow, appRequest bool, ) *EdgeProcessingEligibilityCriteria`

NewEdgeProcessingEligibilityCriteria instantiates a new EdgeProcessingEligibilityCriteria object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEdgeProcessingEligibilityCriteriaWithDefaults

`func NewEdgeProcessingEligibilityCriteriaWithDefaults() *EdgeProcessingEligibilityCriteria`

NewEdgeProcessingEligibilityCriteriaWithDefaults instantiates a new EdgeProcessingEligibilityCriteria object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServiceDataFlowDescriptions

`func (o *EdgeProcessingEligibilityCriteria) GetServiceDataFlowDescriptions() []ServiceDataFlowDescription`

GetServiceDataFlowDescriptions returns the ServiceDataFlowDescriptions field if non-nil, zero value otherwise.

### GetServiceDataFlowDescriptionsOk

`func (o *EdgeProcessingEligibilityCriteria) GetServiceDataFlowDescriptionsOk() (*[]ServiceDataFlowDescription, bool)`

GetServiceDataFlowDescriptionsOk returns a tuple with the ServiceDataFlowDescriptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceDataFlowDescriptions

`func (o *EdgeProcessingEligibilityCriteria) SetServiceDataFlowDescriptions(v []ServiceDataFlowDescription)`

SetServiceDataFlowDescriptions sets ServiceDataFlowDescriptions field to given value.


### GetUeLocations

`func (o *EdgeProcessingEligibilityCriteria) GetUeLocations() []LocationArea5G`

GetUeLocations returns the UeLocations field if non-nil, zero value otherwise.

### GetUeLocationsOk

`func (o *EdgeProcessingEligibilityCriteria) GetUeLocationsOk() (*[]LocationArea5G, bool)`

GetUeLocationsOk returns a tuple with the UeLocations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUeLocations

`func (o *EdgeProcessingEligibilityCriteria) SetUeLocations(v []LocationArea5G)`

SetUeLocations sets UeLocations field to given value.


### GetTimeWindows

`func (o *EdgeProcessingEligibilityCriteria) GetTimeWindows() []TimeWindow`

GetTimeWindows returns the TimeWindows field if non-nil, zero value otherwise.

### GetTimeWindowsOk

`func (o *EdgeProcessingEligibilityCriteria) GetTimeWindowsOk() (*[]TimeWindow, bool)`

GetTimeWindowsOk returns a tuple with the TimeWindows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeWindows

`func (o *EdgeProcessingEligibilityCriteria) SetTimeWindows(v []TimeWindow)`

SetTimeWindows sets TimeWindows field to given value.


### GetAppRequest

`func (o *EdgeProcessingEligibilityCriteria) GetAppRequest() bool`

GetAppRequest returns the AppRequest field if non-nil, zero value otherwise.

### GetAppRequestOk

`func (o *EdgeProcessingEligibilityCriteria) GetAppRequestOk() (*bool, bool)`

GetAppRequestOk returns a tuple with the AppRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppRequest

`func (o *EdgeProcessingEligibilityCriteria) SetAppRequest(v bool)`

SetAppRequest sets AppRequest field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


