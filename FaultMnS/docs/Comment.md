# Comment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CommentTime** | Pointer to **time.Time** |  | [optional] 
**CommentUserId** | Pointer to **string** |  | [optional] 
**CommentSystemId** | Pointer to **string** |  | [optional] 
**CommentText** | Pointer to **string** |  | [optional] 

## Methods

### NewComment

`func NewComment() *Comment`

NewComment instantiates a new Comment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommentWithDefaults

`func NewCommentWithDefaults() *Comment`

NewCommentWithDefaults instantiates a new Comment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommentTime

`func (o *Comment) GetCommentTime() time.Time`

GetCommentTime returns the CommentTime field if non-nil, zero value otherwise.

### GetCommentTimeOk

`func (o *Comment) GetCommentTimeOk() (*time.Time, bool)`

GetCommentTimeOk returns a tuple with the CommentTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommentTime

`func (o *Comment) SetCommentTime(v time.Time)`

SetCommentTime sets CommentTime field to given value.

### HasCommentTime

`func (o *Comment) HasCommentTime() bool`

HasCommentTime returns a boolean if a field has been set.

### GetCommentUserId

`func (o *Comment) GetCommentUserId() string`

GetCommentUserId returns the CommentUserId field if non-nil, zero value otherwise.

### GetCommentUserIdOk

`func (o *Comment) GetCommentUserIdOk() (*string, bool)`

GetCommentUserIdOk returns a tuple with the CommentUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommentUserId

`func (o *Comment) SetCommentUserId(v string)`

SetCommentUserId sets CommentUserId field to given value.

### HasCommentUserId

`func (o *Comment) HasCommentUserId() bool`

HasCommentUserId returns a boolean if a field has been set.

### GetCommentSystemId

`func (o *Comment) GetCommentSystemId() string`

GetCommentSystemId returns the CommentSystemId field if non-nil, zero value otherwise.

### GetCommentSystemIdOk

`func (o *Comment) GetCommentSystemIdOk() (*string, bool)`

GetCommentSystemIdOk returns a tuple with the CommentSystemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommentSystemId

`func (o *Comment) SetCommentSystemId(v string)`

SetCommentSystemId sets CommentSystemId field to given value.

### HasCommentSystemId

`func (o *Comment) HasCommentSystemId() bool`

HasCommentSystemId returns a boolean if a field has been set.

### GetCommentText

`func (o *Comment) GetCommentText() string`

GetCommentText returns the CommentText field if non-nil, zero value otherwise.

### GetCommentTextOk

`func (o *Comment) GetCommentTextOk() (*string, bool)`

GetCommentTextOk returns a tuple with the CommentText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommentText

`func (o *Comment) SetCommentText(v string)`

SetCommentText sets CommentText field to given value.

### HasCommentText

`func (o *Comment) HasCommentText() bool`

HasCommentText returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


