# RestNvdChange

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | Pointer to **string** |  | [optional] 
**CveChangeId** | Pointer to **string** |  | [optional] 
**CveId** | Pointer to **string** |  | [optional] 
**Details** | Pointer to [**[]RestNvdChangeDetail**](RestNvdChangeDetail.md) |  | [optional] 
**EventName** | Pointer to **string** |  | [optional] 
**SourceIdentifier** | Pointer to **string** |  | [optional] 

## Methods

### NewRestNvdChange

`func NewRestNvdChange() *RestNvdChange`

NewRestNvdChange instantiates a new RestNvdChange object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRestNvdChangeWithDefaults

`func NewRestNvdChangeWithDefaults() *RestNvdChange`

NewRestNvdChangeWithDefaults instantiates a new RestNvdChange object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *RestNvdChange) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *RestNvdChange) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *RestNvdChange) SetCreated(v string)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *RestNvdChange) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetCveChangeId

`func (o *RestNvdChange) GetCveChangeId() string`

GetCveChangeId returns the CveChangeId field if non-nil, zero value otherwise.

### GetCveChangeIdOk

`func (o *RestNvdChange) GetCveChangeIdOk() (*string, bool)`

GetCveChangeIdOk returns a tuple with the CveChangeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCveChangeId

`func (o *RestNvdChange) SetCveChangeId(v string)`

SetCveChangeId sets CveChangeId field to given value.

### HasCveChangeId

`func (o *RestNvdChange) HasCveChangeId() bool`

HasCveChangeId returns a boolean if a field has been set.

### GetCveId

`func (o *RestNvdChange) GetCveId() string`

GetCveId returns the CveId field if non-nil, zero value otherwise.

### GetCveIdOk

`func (o *RestNvdChange) GetCveIdOk() (*string, bool)`

GetCveIdOk returns a tuple with the CveId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCveId

`func (o *RestNvdChange) SetCveId(v string)`

SetCveId sets CveId field to given value.

### HasCveId

`func (o *RestNvdChange) HasCveId() bool`

HasCveId returns a boolean if a field has been set.

### GetDetails

`func (o *RestNvdChange) GetDetails() []RestNvdChangeDetail`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *RestNvdChange) GetDetailsOk() (*[]RestNvdChangeDetail, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *RestNvdChange) SetDetails(v []RestNvdChangeDetail)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *RestNvdChange) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### GetEventName

`func (o *RestNvdChange) GetEventName() string`

GetEventName returns the EventName field if non-nil, zero value otherwise.

### GetEventNameOk

`func (o *RestNvdChange) GetEventNameOk() (*string, bool)`

GetEventNameOk returns a tuple with the EventName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventName

`func (o *RestNvdChange) SetEventName(v string)`

SetEventName sets EventName field to given value.

### HasEventName

`func (o *RestNvdChange) HasEventName() bool`

HasEventName returns a boolean if a field has been set.

### GetSourceIdentifier

`func (o *RestNvdChange) GetSourceIdentifier() string`

GetSourceIdentifier returns the SourceIdentifier field if non-nil, zero value otherwise.

### GetSourceIdentifierOk

`func (o *RestNvdChange) GetSourceIdentifierOk() (*string, bool)`

GetSourceIdentifierOk returns a tuple with the SourceIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceIdentifier

`func (o *RestNvdChange) SetSourceIdentifier(v string)`

SetSourceIdentifier sets SourceIdentifier field to given value.

### HasSourceIdentifier

`func (o *RestNvdChange) HasSourceIdentifier() bool`

HasSourceIdentifier returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


