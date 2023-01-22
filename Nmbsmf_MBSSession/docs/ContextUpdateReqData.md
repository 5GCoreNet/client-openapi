# ContextUpdateReqData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NfcInstanceId** | **string** | String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.   | 
**MbsSessionId** | [**MbsSessionId**](MbsSessionId.md) |  | 
**AreaSessionId** | Pointer to **int32** | Integer where the allowed values correspond to the value range of an unsigned 16-bit integer. | [optional] 
**RequestedAction** | Pointer to [**ContextUpdateAction**](ContextUpdateAction.md) |  | [optional] 
**DlTunnelInfo** | Pointer to **string** | string with format &#39;bytes&#39; as defined in OpenAPI | [optional] 
**N2MbsSmInfo** | Pointer to [**N2MbsSmInfo**](N2MbsSmInfo.md) |  | [optional] 
**RanNodeId** | Pointer to [**GlobalRanNodeId**](GlobalRanNodeId.md) |  | [optional] 
**LeaveInd** | Pointer to **bool** |  | [optional] 

## Methods

### NewContextUpdateReqData

`func NewContextUpdateReqData(nfcInstanceId string, mbsSessionId MbsSessionId, ) *ContextUpdateReqData`

NewContextUpdateReqData instantiates a new ContextUpdateReqData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContextUpdateReqDataWithDefaults

`func NewContextUpdateReqDataWithDefaults() *ContextUpdateReqData`

NewContextUpdateReqDataWithDefaults instantiates a new ContextUpdateReqData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNfcInstanceId

`func (o *ContextUpdateReqData) GetNfcInstanceId() string`

GetNfcInstanceId returns the NfcInstanceId field if non-nil, zero value otherwise.

### GetNfcInstanceIdOk

`func (o *ContextUpdateReqData) GetNfcInstanceIdOk() (*string, bool)`

GetNfcInstanceIdOk returns a tuple with the NfcInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNfcInstanceId

`func (o *ContextUpdateReqData) SetNfcInstanceId(v string)`

SetNfcInstanceId sets NfcInstanceId field to given value.


### GetMbsSessionId

`func (o *ContextUpdateReqData) GetMbsSessionId() MbsSessionId`

GetMbsSessionId returns the MbsSessionId field if non-nil, zero value otherwise.

### GetMbsSessionIdOk

`func (o *ContextUpdateReqData) GetMbsSessionIdOk() (*MbsSessionId, bool)`

GetMbsSessionIdOk returns a tuple with the MbsSessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMbsSessionId

`func (o *ContextUpdateReqData) SetMbsSessionId(v MbsSessionId)`

SetMbsSessionId sets MbsSessionId field to given value.


### GetAreaSessionId

`func (o *ContextUpdateReqData) GetAreaSessionId() int32`

GetAreaSessionId returns the AreaSessionId field if non-nil, zero value otherwise.

### GetAreaSessionIdOk

`func (o *ContextUpdateReqData) GetAreaSessionIdOk() (*int32, bool)`

GetAreaSessionIdOk returns a tuple with the AreaSessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAreaSessionId

`func (o *ContextUpdateReqData) SetAreaSessionId(v int32)`

SetAreaSessionId sets AreaSessionId field to given value.

### HasAreaSessionId

`func (o *ContextUpdateReqData) HasAreaSessionId() bool`

HasAreaSessionId returns a boolean if a field has been set.

### GetRequestedAction

`func (o *ContextUpdateReqData) GetRequestedAction() ContextUpdateAction`

GetRequestedAction returns the RequestedAction field if non-nil, zero value otherwise.

### GetRequestedActionOk

`func (o *ContextUpdateReqData) GetRequestedActionOk() (*ContextUpdateAction, bool)`

GetRequestedActionOk returns a tuple with the RequestedAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedAction

`func (o *ContextUpdateReqData) SetRequestedAction(v ContextUpdateAction)`

SetRequestedAction sets RequestedAction field to given value.

### HasRequestedAction

`func (o *ContextUpdateReqData) HasRequestedAction() bool`

HasRequestedAction returns a boolean if a field has been set.

### GetDlTunnelInfo

`func (o *ContextUpdateReqData) GetDlTunnelInfo() string`

GetDlTunnelInfo returns the DlTunnelInfo field if non-nil, zero value otherwise.

### GetDlTunnelInfoOk

`func (o *ContextUpdateReqData) GetDlTunnelInfoOk() (*string, bool)`

GetDlTunnelInfoOk returns a tuple with the DlTunnelInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDlTunnelInfo

`func (o *ContextUpdateReqData) SetDlTunnelInfo(v string)`

SetDlTunnelInfo sets DlTunnelInfo field to given value.

### HasDlTunnelInfo

`func (o *ContextUpdateReqData) HasDlTunnelInfo() bool`

HasDlTunnelInfo returns a boolean if a field has been set.

### GetN2MbsSmInfo

`func (o *ContextUpdateReqData) GetN2MbsSmInfo() N2MbsSmInfo`

GetN2MbsSmInfo returns the N2MbsSmInfo field if non-nil, zero value otherwise.

### GetN2MbsSmInfoOk

`func (o *ContextUpdateReqData) GetN2MbsSmInfoOk() (*N2MbsSmInfo, bool)`

GetN2MbsSmInfoOk returns a tuple with the N2MbsSmInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetN2MbsSmInfo

`func (o *ContextUpdateReqData) SetN2MbsSmInfo(v N2MbsSmInfo)`

SetN2MbsSmInfo sets N2MbsSmInfo field to given value.

### HasN2MbsSmInfo

`func (o *ContextUpdateReqData) HasN2MbsSmInfo() bool`

HasN2MbsSmInfo returns a boolean if a field has been set.

### GetRanNodeId

`func (o *ContextUpdateReqData) GetRanNodeId() GlobalRanNodeId`

GetRanNodeId returns the RanNodeId field if non-nil, zero value otherwise.

### GetRanNodeIdOk

`func (o *ContextUpdateReqData) GetRanNodeIdOk() (*GlobalRanNodeId, bool)`

GetRanNodeIdOk returns a tuple with the RanNodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRanNodeId

`func (o *ContextUpdateReqData) SetRanNodeId(v GlobalRanNodeId)`

SetRanNodeId sets RanNodeId field to given value.

### HasRanNodeId

`func (o *ContextUpdateReqData) HasRanNodeId() bool`

HasRanNodeId returns a boolean if a field has been set.

### GetLeaveInd

`func (o *ContextUpdateReqData) GetLeaveInd() bool`

GetLeaveInd returns the LeaveInd field if non-nil, zero value otherwise.

### GetLeaveIndOk

`func (o *ContextUpdateReqData) GetLeaveIndOk() (*bool, bool)`

GetLeaveIndOk returns a tuple with the LeaveInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeaveInd

`func (o *ContextUpdateReqData) SetLeaveInd(v bool)`

SetLeaveInd sets LeaveInd field to given value.

### HasLeaveInd

`func (o *ContextUpdateReqData) HasLeaveInd() bool`

HasLeaveInd returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


