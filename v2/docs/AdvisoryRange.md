# AdvisoryRange

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Events** | Pointer to [**[]AdvisoryEvent**](AdvisoryEvent.md) |  | [optional] 
**Repo** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 

## Methods

### NewAdvisoryRange

`func NewAdvisoryRange() *AdvisoryRange`

NewAdvisoryRange instantiates a new AdvisoryRange object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdvisoryRangeWithDefaults

`func NewAdvisoryRangeWithDefaults() *AdvisoryRange`

NewAdvisoryRangeWithDefaults instantiates a new AdvisoryRange object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvents

`func (o *AdvisoryRange) GetEvents() []AdvisoryEvent`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *AdvisoryRange) GetEventsOk() (*[]AdvisoryEvent, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *AdvisoryRange) SetEvents(v []AdvisoryEvent)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *AdvisoryRange) HasEvents() bool`

HasEvents returns a boolean if a field has been set.

### GetRepo

`func (o *AdvisoryRange) GetRepo() string`

GetRepo returns the Repo field if non-nil, zero value otherwise.

### GetRepoOk

`func (o *AdvisoryRange) GetRepoOk() (*string, bool)`

GetRepoOk returns a tuple with the Repo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepo

`func (o *AdvisoryRange) SetRepo(v string)`

SetRepo sets Repo field to given value.

### HasRepo

`func (o *AdvisoryRange) HasRepo() bool`

HasRepo returns a boolean if a field has been set.

### GetType

`func (o *AdvisoryRange) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AdvisoryRange) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AdvisoryRange) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *AdvisoryRange) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


