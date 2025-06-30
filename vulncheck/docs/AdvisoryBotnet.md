# AdvisoryBotnet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BotnetName** | Pointer to **string** |  | [optional] 
**Cve** | Pointer to **[]string** |  | [optional] 
**CveReferences** | Pointer to [**[]AdvisoryCVEReference**](AdvisoryCVEReference.md) |  | [optional] 
**DateAdded** | Pointer to **string** |  | [optional] 
**MalpediaUrl** | Pointer to **string** |  | [optional] 
**Tools** | Pointer to [**[]AdvisoryTool**](AdvisoryTool.md) |  | [optional] 

## Methods

### NewAdvisoryBotnet

`func NewAdvisoryBotnet() *AdvisoryBotnet`

NewAdvisoryBotnet instantiates a new AdvisoryBotnet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdvisoryBotnetWithDefaults

`func NewAdvisoryBotnetWithDefaults() *AdvisoryBotnet`

NewAdvisoryBotnetWithDefaults instantiates a new AdvisoryBotnet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBotnetName

`func (o *AdvisoryBotnet) GetBotnetName() string`

GetBotnetName returns the BotnetName field if non-nil, zero value otherwise.

### GetBotnetNameOk

`func (o *AdvisoryBotnet) GetBotnetNameOk() (*string, bool)`

GetBotnetNameOk returns a tuple with the BotnetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBotnetName

`func (o *AdvisoryBotnet) SetBotnetName(v string)`

SetBotnetName sets BotnetName field to given value.

### HasBotnetName

`func (o *AdvisoryBotnet) HasBotnetName() bool`

HasBotnetName returns a boolean if a field has been set.

### GetCve

`func (o *AdvisoryBotnet) GetCve() []string`

GetCve returns the Cve field if non-nil, zero value otherwise.

### GetCveOk

`func (o *AdvisoryBotnet) GetCveOk() (*[]string, bool)`

GetCveOk returns a tuple with the Cve field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCve

`func (o *AdvisoryBotnet) SetCve(v []string)`

SetCve sets Cve field to given value.

### HasCve

`func (o *AdvisoryBotnet) HasCve() bool`

HasCve returns a boolean if a field has been set.

### GetCveReferences

`func (o *AdvisoryBotnet) GetCveReferences() []AdvisoryCVEReference`

GetCveReferences returns the CveReferences field if non-nil, zero value otherwise.

### GetCveReferencesOk

`func (o *AdvisoryBotnet) GetCveReferencesOk() (*[]AdvisoryCVEReference, bool)`

GetCveReferencesOk returns a tuple with the CveReferences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCveReferences

`func (o *AdvisoryBotnet) SetCveReferences(v []AdvisoryCVEReference)`

SetCveReferences sets CveReferences field to given value.

### HasCveReferences

`func (o *AdvisoryBotnet) HasCveReferences() bool`

HasCveReferences returns a boolean if a field has been set.

### GetDateAdded

`func (o *AdvisoryBotnet) GetDateAdded() string`

GetDateAdded returns the DateAdded field if non-nil, zero value otherwise.

### GetDateAddedOk

`func (o *AdvisoryBotnet) GetDateAddedOk() (*string, bool)`

GetDateAddedOk returns a tuple with the DateAdded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateAdded

`func (o *AdvisoryBotnet) SetDateAdded(v string)`

SetDateAdded sets DateAdded field to given value.

### HasDateAdded

`func (o *AdvisoryBotnet) HasDateAdded() bool`

HasDateAdded returns a boolean if a field has been set.

### GetMalpediaUrl

`func (o *AdvisoryBotnet) GetMalpediaUrl() string`

GetMalpediaUrl returns the MalpediaUrl field if non-nil, zero value otherwise.

### GetMalpediaUrlOk

`func (o *AdvisoryBotnet) GetMalpediaUrlOk() (*string, bool)`

GetMalpediaUrlOk returns a tuple with the MalpediaUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMalpediaUrl

`func (o *AdvisoryBotnet) SetMalpediaUrl(v string)`

SetMalpediaUrl sets MalpediaUrl field to given value.

### HasMalpediaUrl

`func (o *AdvisoryBotnet) HasMalpediaUrl() bool`

HasMalpediaUrl returns a boolean if a field has been set.

### GetTools

`func (o *AdvisoryBotnet) GetTools() []AdvisoryTool`

GetTools returns the Tools field if non-nil, zero value otherwise.

### GetToolsOk

`func (o *AdvisoryBotnet) GetToolsOk() (*[]AdvisoryTool, bool)`

GetToolsOk returns a tuple with the Tools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTools

`func (o *AdvisoryBotnet) SetTools(v []AdvisoryTool)`

SetTools sets Tools field to given value.

### HasTools

`func (o *AdvisoryBotnet) HasTools() bool`

HasTools returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


