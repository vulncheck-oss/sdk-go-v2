# AdvisoryContainerOS

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cve** | Pointer to **[]string** |  | [optional] 
**DateAdded** | Pointer to **string** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**Updates** | Pointer to [**[]AdvisoryCOSUpdate**](AdvisoryCOSUpdate.md) |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 

## Methods

### NewAdvisoryContainerOS

`func NewAdvisoryContainerOS() *AdvisoryContainerOS`

NewAdvisoryContainerOS instantiates a new AdvisoryContainerOS object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdvisoryContainerOSWithDefaults

`func NewAdvisoryContainerOSWithDefaults() *AdvisoryContainerOS`

NewAdvisoryContainerOSWithDefaults instantiates a new AdvisoryContainerOS object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCve

`func (o *AdvisoryContainerOS) GetCve() []string`

GetCve returns the Cve field if non-nil, zero value otherwise.

### GetCveOk

`func (o *AdvisoryContainerOS) GetCveOk() (*[]string, bool)`

GetCveOk returns a tuple with the Cve field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCve

`func (o *AdvisoryContainerOS) SetCve(v []string)`

SetCve sets Cve field to given value.

### HasCve

`func (o *AdvisoryContainerOS) HasCve() bool`

HasCve returns a boolean if a field has been set.

### GetDateAdded

`func (o *AdvisoryContainerOS) GetDateAdded() string`

GetDateAdded returns the DateAdded field if non-nil, zero value otherwise.

### GetDateAddedOk

`func (o *AdvisoryContainerOS) GetDateAddedOk() (*string, bool)`

GetDateAddedOk returns a tuple with the DateAdded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateAdded

`func (o *AdvisoryContainerOS) SetDateAdded(v string)`

SetDateAdded sets DateAdded field to given value.

### HasDateAdded

`func (o *AdvisoryContainerOS) HasDateAdded() bool`

HasDateAdded returns a boolean if a field has been set.

### GetTitle

`func (o *AdvisoryContainerOS) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *AdvisoryContainerOS) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *AdvisoryContainerOS) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *AdvisoryContainerOS) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetUpdates

`func (o *AdvisoryContainerOS) GetUpdates() []AdvisoryCOSUpdate`

GetUpdates returns the Updates field if non-nil, zero value otherwise.

### GetUpdatesOk

`func (o *AdvisoryContainerOS) GetUpdatesOk() (*[]AdvisoryCOSUpdate, bool)`

GetUpdatesOk returns a tuple with the Updates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdates

`func (o *AdvisoryContainerOS) SetUpdates(v []AdvisoryCOSUpdate)`

SetUpdates sets Updates field to given value.

### HasUpdates

`func (o *AdvisoryContainerOS) HasUpdates() bool`

HasUpdates returns a boolean if a field has been set.

### GetUrl

`func (o *AdvisoryContainerOS) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *AdvisoryContainerOS) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *AdvisoryContainerOS) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *AdvisoryContainerOS) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


