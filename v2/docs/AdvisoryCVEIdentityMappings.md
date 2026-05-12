# AdvisoryCVEIdentityMappings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cve** | Pointer to **[]string** |  | [optional] 
**Mappings** | Pointer to [**[]AdvisoryCVEMapping**](AdvisoryCVEMapping.md) |  | [optional] 

## Methods

### NewAdvisoryCVEIdentityMappings

`func NewAdvisoryCVEIdentityMappings() *AdvisoryCVEIdentityMappings`

NewAdvisoryCVEIdentityMappings instantiates a new AdvisoryCVEIdentityMappings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdvisoryCVEIdentityMappingsWithDefaults

`func NewAdvisoryCVEIdentityMappingsWithDefaults() *AdvisoryCVEIdentityMappings`

NewAdvisoryCVEIdentityMappingsWithDefaults instantiates a new AdvisoryCVEIdentityMappings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCve

`func (o *AdvisoryCVEIdentityMappings) GetCve() []string`

GetCve returns the Cve field if non-nil, zero value otherwise.

### GetCveOk

`func (o *AdvisoryCVEIdentityMappings) GetCveOk() (*[]string, bool)`

GetCveOk returns a tuple with the Cve field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCve

`func (o *AdvisoryCVEIdentityMappings) SetCve(v []string)`

SetCve sets Cve field to given value.

### HasCve

`func (o *AdvisoryCVEIdentityMappings) HasCve() bool`

HasCve returns a boolean if a field has been set.

### GetMappings

`func (o *AdvisoryCVEIdentityMappings) GetMappings() []AdvisoryCVEMapping`

GetMappings returns the Mappings field if non-nil, zero value otherwise.

### GetMappingsOk

`func (o *AdvisoryCVEIdentityMappings) GetMappingsOk() (*[]AdvisoryCVEMapping, bool)`

GetMappingsOk returns a tuple with the Mappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMappings

`func (o *AdvisoryCVEIdentityMappings) SetMappings(v []AdvisoryCVEMapping)`

SetMappings sets Mappings field to given value.

### HasMappings

`func (o *AdvisoryCVEIdentityMappings) HasMappings() bool`

HasMappings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


