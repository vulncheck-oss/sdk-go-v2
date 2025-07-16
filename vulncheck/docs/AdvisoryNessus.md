# AdvisoryNessus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cpe** | Pointer to **[]string** |  | [optional] 
**Cve** | Pointer to **[]string** |  | [optional] 
**DateAdded** | Pointer to **string** |  | [optional] 
**ExploitabilityEase** | Pointer to **string** | seems like only 3 vals for this | [optional] 
**Filename** | Pointer to **string** |  | [optional] 
**Iava** | Pointer to **[]string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**ScriptId** | Pointer to **int32** |  | [optional] 
**Updated** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 

## Methods

### NewAdvisoryNessus

`func NewAdvisoryNessus() *AdvisoryNessus`

NewAdvisoryNessus instantiates a new AdvisoryNessus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdvisoryNessusWithDefaults

`func NewAdvisoryNessusWithDefaults() *AdvisoryNessus`

NewAdvisoryNessusWithDefaults instantiates a new AdvisoryNessus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCpe

`func (o *AdvisoryNessus) GetCpe() []string`

GetCpe returns the Cpe field if non-nil, zero value otherwise.

### GetCpeOk

`func (o *AdvisoryNessus) GetCpeOk() (*[]string, bool)`

GetCpeOk returns a tuple with the Cpe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpe

`func (o *AdvisoryNessus) SetCpe(v []string)`

SetCpe sets Cpe field to given value.

### HasCpe

`func (o *AdvisoryNessus) HasCpe() bool`

HasCpe returns a boolean if a field has been set.

### GetCve

`func (o *AdvisoryNessus) GetCve() []string`

GetCve returns the Cve field if non-nil, zero value otherwise.

### GetCveOk

`func (o *AdvisoryNessus) GetCveOk() (*[]string, bool)`

GetCveOk returns a tuple with the Cve field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCve

`func (o *AdvisoryNessus) SetCve(v []string)`

SetCve sets Cve field to given value.

### HasCve

`func (o *AdvisoryNessus) HasCve() bool`

HasCve returns a boolean if a field has been set.

### GetDateAdded

`func (o *AdvisoryNessus) GetDateAdded() string`

GetDateAdded returns the DateAdded field if non-nil, zero value otherwise.

### GetDateAddedOk

`func (o *AdvisoryNessus) GetDateAddedOk() (*string, bool)`

GetDateAddedOk returns a tuple with the DateAdded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateAdded

`func (o *AdvisoryNessus) SetDateAdded(v string)`

SetDateAdded sets DateAdded field to given value.

### HasDateAdded

`func (o *AdvisoryNessus) HasDateAdded() bool`

HasDateAdded returns a boolean if a field has been set.

### GetExploitabilityEase

`func (o *AdvisoryNessus) GetExploitabilityEase() string`

GetExploitabilityEase returns the ExploitabilityEase field if non-nil, zero value otherwise.

### GetExploitabilityEaseOk

`func (o *AdvisoryNessus) GetExploitabilityEaseOk() (*string, bool)`

GetExploitabilityEaseOk returns a tuple with the ExploitabilityEase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExploitabilityEase

`func (o *AdvisoryNessus) SetExploitabilityEase(v string)`

SetExploitabilityEase sets ExploitabilityEase field to given value.

### HasExploitabilityEase

`func (o *AdvisoryNessus) HasExploitabilityEase() bool`

HasExploitabilityEase returns a boolean if a field has been set.

### GetFilename

`func (o *AdvisoryNessus) GetFilename() string`

GetFilename returns the Filename field if non-nil, zero value otherwise.

### GetFilenameOk

`func (o *AdvisoryNessus) GetFilenameOk() (*string, bool)`

GetFilenameOk returns a tuple with the Filename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilename

`func (o *AdvisoryNessus) SetFilename(v string)`

SetFilename sets Filename field to given value.

### HasFilename

`func (o *AdvisoryNessus) HasFilename() bool`

HasFilename returns a boolean if a field has been set.

### GetIava

`func (o *AdvisoryNessus) GetIava() []string`

GetIava returns the Iava field if non-nil, zero value otherwise.

### GetIavaOk

`func (o *AdvisoryNessus) GetIavaOk() (*[]string, bool)`

GetIavaOk returns a tuple with the Iava field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIava

`func (o *AdvisoryNessus) SetIava(v []string)`

SetIava sets Iava field to given value.

### HasIava

`func (o *AdvisoryNessus) HasIava() bool`

HasIava returns a boolean if a field has been set.

### GetName

`func (o *AdvisoryNessus) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AdvisoryNessus) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AdvisoryNessus) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AdvisoryNessus) HasName() bool`

HasName returns a boolean if a field has been set.

### GetScriptId

`func (o *AdvisoryNessus) GetScriptId() int32`

GetScriptId returns the ScriptId field if non-nil, zero value otherwise.

### GetScriptIdOk

`func (o *AdvisoryNessus) GetScriptIdOk() (*int32, bool)`

GetScriptIdOk returns a tuple with the ScriptId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScriptId

`func (o *AdvisoryNessus) SetScriptId(v int32)`

SetScriptId sets ScriptId field to given value.

### HasScriptId

`func (o *AdvisoryNessus) HasScriptId() bool`

HasScriptId returns a boolean if a field has been set.

### GetUpdated

`func (o *AdvisoryNessus) GetUpdated() string`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *AdvisoryNessus) GetUpdatedOk() (*string, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *AdvisoryNessus) SetUpdated(v string)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *AdvisoryNessus) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *AdvisoryNessus) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *AdvisoryNessus) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *AdvisoryNessus) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *AdvisoryNessus) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUrl

`func (o *AdvisoryNessus) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *AdvisoryNessus) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *AdvisoryNessus) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *AdvisoryNessus) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


