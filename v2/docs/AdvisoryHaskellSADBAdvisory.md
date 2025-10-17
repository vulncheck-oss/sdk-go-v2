# AdvisoryHaskellSADBAdvisory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdvisoryId** | Pointer to **string** |  | [optional] 
**AffectedPackages** | Pointer to [**[]AdvisoryHaskellAffected**](AdvisoryHaskellAffected.md) |  | [optional] 
**Aliases** | Pointer to **[]string** |  | [optional] 
**Cve** | Pointer to **[]string** |  | [optional] 
**Cwes** | Pointer to **[]int32** |  | [optional] 
**DateAdded** | Pointer to **string** |  | [optional] 
**Keywords** | Pointer to **[]string** |  | [optional] 
**References** | Pointer to **map[string][]string** |  | [optional] 
**RelatedVulns** | Pointer to **[]string** |  | [optional] 

## Methods

### NewAdvisoryHaskellSADBAdvisory

`func NewAdvisoryHaskellSADBAdvisory() *AdvisoryHaskellSADBAdvisory`

NewAdvisoryHaskellSADBAdvisory instantiates a new AdvisoryHaskellSADBAdvisory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdvisoryHaskellSADBAdvisoryWithDefaults

`func NewAdvisoryHaskellSADBAdvisoryWithDefaults() *AdvisoryHaskellSADBAdvisory`

NewAdvisoryHaskellSADBAdvisoryWithDefaults instantiates a new AdvisoryHaskellSADBAdvisory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdvisoryId

`func (o *AdvisoryHaskellSADBAdvisory) GetAdvisoryId() string`

GetAdvisoryId returns the AdvisoryId field if non-nil, zero value otherwise.

### GetAdvisoryIdOk

`func (o *AdvisoryHaskellSADBAdvisory) GetAdvisoryIdOk() (*string, bool)`

GetAdvisoryIdOk returns a tuple with the AdvisoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvisoryId

`func (o *AdvisoryHaskellSADBAdvisory) SetAdvisoryId(v string)`

SetAdvisoryId sets AdvisoryId field to given value.

### HasAdvisoryId

`func (o *AdvisoryHaskellSADBAdvisory) HasAdvisoryId() bool`

HasAdvisoryId returns a boolean if a field has been set.

### GetAffectedPackages

`func (o *AdvisoryHaskellSADBAdvisory) GetAffectedPackages() []AdvisoryHaskellAffected`

GetAffectedPackages returns the AffectedPackages field if non-nil, zero value otherwise.

### GetAffectedPackagesOk

`func (o *AdvisoryHaskellSADBAdvisory) GetAffectedPackagesOk() (*[]AdvisoryHaskellAffected, bool)`

GetAffectedPackagesOk returns a tuple with the AffectedPackages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffectedPackages

`func (o *AdvisoryHaskellSADBAdvisory) SetAffectedPackages(v []AdvisoryHaskellAffected)`

SetAffectedPackages sets AffectedPackages field to given value.

### HasAffectedPackages

`func (o *AdvisoryHaskellSADBAdvisory) HasAffectedPackages() bool`

HasAffectedPackages returns a boolean if a field has been set.

### GetAliases

`func (o *AdvisoryHaskellSADBAdvisory) GetAliases() []string`

GetAliases returns the Aliases field if non-nil, zero value otherwise.

### GetAliasesOk

`func (o *AdvisoryHaskellSADBAdvisory) GetAliasesOk() (*[]string, bool)`

GetAliasesOk returns a tuple with the Aliases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAliases

`func (o *AdvisoryHaskellSADBAdvisory) SetAliases(v []string)`

SetAliases sets Aliases field to given value.

### HasAliases

`func (o *AdvisoryHaskellSADBAdvisory) HasAliases() bool`

HasAliases returns a boolean if a field has been set.

### GetCve

`func (o *AdvisoryHaskellSADBAdvisory) GetCve() []string`

GetCve returns the Cve field if non-nil, zero value otherwise.

### GetCveOk

`func (o *AdvisoryHaskellSADBAdvisory) GetCveOk() (*[]string, bool)`

GetCveOk returns a tuple with the Cve field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCve

`func (o *AdvisoryHaskellSADBAdvisory) SetCve(v []string)`

SetCve sets Cve field to given value.

### HasCve

`func (o *AdvisoryHaskellSADBAdvisory) HasCve() bool`

HasCve returns a boolean if a field has been set.

### GetCwes

`func (o *AdvisoryHaskellSADBAdvisory) GetCwes() []int32`

GetCwes returns the Cwes field if non-nil, zero value otherwise.

### GetCwesOk

`func (o *AdvisoryHaskellSADBAdvisory) GetCwesOk() (*[]int32, bool)`

GetCwesOk returns a tuple with the Cwes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCwes

`func (o *AdvisoryHaskellSADBAdvisory) SetCwes(v []int32)`

SetCwes sets Cwes field to given value.

### HasCwes

`func (o *AdvisoryHaskellSADBAdvisory) HasCwes() bool`

HasCwes returns a boolean if a field has been set.

### GetDateAdded

`func (o *AdvisoryHaskellSADBAdvisory) GetDateAdded() string`

GetDateAdded returns the DateAdded field if non-nil, zero value otherwise.

### GetDateAddedOk

`func (o *AdvisoryHaskellSADBAdvisory) GetDateAddedOk() (*string, bool)`

GetDateAddedOk returns a tuple with the DateAdded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateAdded

`func (o *AdvisoryHaskellSADBAdvisory) SetDateAdded(v string)`

SetDateAdded sets DateAdded field to given value.

### HasDateAdded

`func (o *AdvisoryHaskellSADBAdvisory) HasDateAdded() bool`

HasDateAdded returns a boolean if a field has been set.

### GetKeywords

`func (o *AdvisoryHaskellSADBAdvisory) GetKeywords() []string`

GetKeywords returns the Keywords field if non-nil, zero value otherwise.

### GetKeywordsOk

`func (o *AdvisoryHaskellSADBAdvisory) GetKeywordsOk() (*[]string, bool)`

GetKeywordsOk returns a tuple with the Keywords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeywords

`func (o *AdvisoryHaskellSADBAdvisory) SetKeywords(v []string)`

SetKeywords sets Keywords field to given value.

### HasKeywords

`func (o *AdvisoryHaskellSADBAdvisory) HasKeywords() bool`

HasKeywords returns a boolean if a field has been set.

### GetReferences

`func (o *AdvisoryHaskellSADBAdvisory) GetReferences() map[string][]string`

GetReferences returns the References field if non-nil, zero value otherwise.

### GetReferencesOk

`func (o *AdvisoryHaskellSADBAdvisory) GetReferencesOk() (*map[string][]string, bool)`

GetReferencesOk returns a tuple with the References field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferences

`func (o *AdvisoryHaskellSADBAdvisory) SetReferences(v map[string][]string)`

SetReferences sets References field to given value.

### HasReferences

`func (o *AdvisoryHaskellSADBAdvisory) HasReferences() bool`

HasReferences returns a boolean if a field has been set.

### GetRelatedVulns

`func (o *AdvisoryHaskellSADBAdvisory) GetRelatedVulns() []string`

GetRelatedVulns returns the RelatedVulns field if non-nil, zero value otherwise.

### GetRelatedVulnsOk

`func (o *AdvisoryHaskellSADBAdvisory) GetRelatedVulnsOk() (*[]string, bool)`

GetRelatedVulnsOk returns a tuple with the RelatedVulns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelatedVulns

`func (o *AdvisoryHaskellSADBAdvisory) SetRelatedVulns(v []string)`

SetRelatedVulns sets RelatedVulns field to given value.

### HasRelatedVulns

`func (o *AdvisoryHaskellSADBAdvisory) HasRelatedVulns() bool`

HasRelatedVulns returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


