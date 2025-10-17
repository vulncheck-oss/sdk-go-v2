# AdvisoryRockyAdvisory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AffectedProducts** | Pointer to **[]string** |  | [optional] 
**BuildReferences** | Pointer to **[]string** |  | [optional] 
**Cves** | Pointer to [**[]AdvisoryRockyCve**](AdvisoryRockyCve.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Fixes** | Pointer to [**[]AdvisoryRockyFix**](AdvisoryRockyFix.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**PublishedAt** | Pointer to **string** |  | [optional] 
**RebootSuggested** | Pointer to **bool** |  | [optional] 
**References** | Pointer to **[]string** |  | [optional] 
**Rpms** | Pointer to [**map[string]AdvisoryRockyVersion**](AdvisoryRockyVersion.md) |  | [optional] 
**Severity** | Pointer to **string** |  | [optional] 
**ShortCode** | Pointer to **string** |  | [optional] 
**Solution** | Pointer to **string** |  | [optional] 
**Synopsis** | Pointer to **string** |  | [optional] 
**Topic** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 

## Methods

### NewAdvisoryRockyAdvisory

`func NewAdvisoryRockyAdvisory() *AdvisoryRockyAdvisory`

NewAdvisoryRockyAdvisory instantiates a new AdvisoryRockyAdvisory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdvisoryRockyAdvisoryWithDefaults

`func NewAdvisoryRockyAdvisoryWithDefaults() *AdvisoryRockyAdvisory`

NewAdvisoryRockyAdvisoryWithDefaults instantiates a new AdvisoryRockyAdvisory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAffectedProducts

`func (o *AdvisoryRockyAdvisory) GetAffectedProducts() []string`

GetAffectedProducts returns the AffectedProducts field if non-nil, zero value otherwise.

### GetAffectedProductsOk

`func (o *AdvisoryRockyAdvisory) GetAffectedProductsOk() (*[]string, bool)`

GetAffectedProductsOk returns a tuple with the AffectedProducts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffectedProducts

`func (o *AdvisoryRockyAdvisory) SetAffectedProducts(v []string)`

SetAffectedProducts sets AffectedProducts field to given value.

### HasAffectedProducts

`func (o *AdvisoryRockyAdvisory) HasAffectedProducts() bool`

HasAffectedProducts returns a boolean if a field has been set.

### GetBuildReferences

`func (o *AdvisoryRockyAdvisory) GetBuildReferences() []string`

GetBuildReferences returns the BuildReferences field if non-nil, zero value otherwise.

### GetBuildReferencesOk

`func (o *AdvisoryRockyAdvisory) GetBuildReferencesOk() (*[]string, bool)`

GetBuildReferencesOk returns a tuple with the BuildReferences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildReferences

`func (o *AdvisoryRockyAdvisory) SetBuildReferences(v []string)`

SetBuildReferences sets BuildReferences field to given value.

### HasBuildReferences

`func (o *AdvisoryRockyAdvisory) HasBuildReferences() bool`

HasBuildReferences returns a boolean if a field has been set.

### GetCves

`func (o *AdvisoryRockyAdvisory) GetCves() []AdvisoryRockyCve`

GetCves returns the Cves field if non-nil, zero value otherwise.

### GetCvesOk

`func (o *AdvisoryRockyAdvisory) GetCvesOk() (*[]AdvisoryRockyCve, bool)`

GetCvesOk returns a tuple with the Cves field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCves

`func (o *AdvisoryRockyAdvisory) SetCves(v []AdvisoryRockyCve)`

SetCves sets Cves field to given value.

### HasCves

`func (o *AdvisoryRockyAdvisory) HasCves() bool`

HasCves returns a boolean if a field has been set.

### GetDescription

`func (o *AdvisoryRockyAdvisory) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AdvisoryRockyAdvisory) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AdvisoryRockyAdvisory) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AdvisoryRockyAdvisory) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFixes

`func (o *AdvisoryRockyAdvisory) GetFixes() []AdvisoryRockyFix`

GetFixes returns the Fixes field if non-nil, zero value otherwise.

### GetFixesOk

`func (o *AdvisoryRockyAdvisory) GetFixesOk() (*[]AdvisoryRockyFix, bool)`

GetFixesOk returns a tuple with the Fixes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFixes

`func (o *AdvisoryRockyAdvisory) SetFixes(v []AdvisoryRockyFix)`

SetFixes sets Fixes field to given value.

### HasFixes

`func (o *AdvisoryRockyAdvisory) HasFixes() bool`

HasFixes returns a boolean if a field has been set.

### GetName

`func (o *AdvisoryRockyAdvisory) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AdvisoryRockyAdvisory) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AdvisoryRockyAdvisory) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AdvisoryRockyAdvisory) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPublishedAt

`func (o *AdvisoryRockyAdvisory) GetPublishedAt() string`

GetPublishedAt returns the PublishedAt field if non-nil, zero value otherwise.

### GetPublishedAtOk

`func (o *AdvisoryRockyAdvisory) GetPublishedAtOk() (*string, bool)`

GetPublishedAtOk returns a tuple with the PublishedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishedAt

`func (o *AdvisoryRockyAdvisory) SetPublishedAt(v string)`

SetPublishedAt sets PublishedAt field to given value.

### HasPublishedAt

`func (o *AdvisoryRockyAdvisory) HasPublishedAt() bool`

HasPublishedAt returns a boolean if a field has been set.

### GetRebootSuggested

`func (o *AdvisoryRockyAdvisory) GetRebootSuggested() bool`

GetRebootSuggested returns the RebootSuggested field if non-nil, zero value otherwise.

### GetRebootSuggestedOk

`func (o *AdvisoryRockyAdvisory) GetRebootSuggestedOk() (*bool, bool)`

GetRebootSuggestedOk returns a tuple with the RebootSuggested field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRebootSuggested

`func (o *AdvisoryRockyAdvisory) SetRebootSuggested(v bool)`

SetRebootSuggested sets RebootSuggested field to given value.

### HasRebootSuggested

`func (o *AdvisoryRockyAdvisory) HasRebootSuggested() bool`

HasRebootSuggested returns a boolean if a field has been set.

### GetReferences

`func (o *AdvisoryRockyAdvisory) GetReferences() []string`

GetReferences returns the References field if non-nil, zero value otherwise.

### GetReferencesOk

`func (o *AdvisoryRockyAdvisory) GetReferencesOk() (*[]string, bool)`

GetReferencesOk returns a tuple with the References field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferences

`func (o *AdvisoryRockyAdvisory) SetReferences(v []string)`

SetReferences sets References field to given value.

### HasReferences

`func (o *AdvisoryRockyAdvisory) HasReferences() bool`

HasReferences returns a boolean if a field has been set.

### GetRpms

`func (o *AdvisoryRockyAdvisory) GetRpms() map[string]AdvisoryRockyVersion`

GetRpms returns the Rpms field if non-nil, zero value otherwise.

### GetRpmsOk

`func (o *AdvisoryRockyAdvisory) GetRpmsOk() (*map[string]AdvisoryRockyVersion, bool)`

GetRpmsOk returns a tuple with the Rpms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRpms

`func (o *AdvisoryRockyAdvisory) SetRpms(v map[string]AdvisoryRockyVersion)`

SetRpms sets Rpms field to given value.

### HasRpms

`func (o *AdvisoryRockyAdvisory) HasRpms() bool`

HasRpms returns a boolean if a field has been set.

### GetSeverity

`func (o *AdvisoryRockyAdvisory) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *AdvisoryRockyAdvisory) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *AdvisoryRockyAdvisory) SetSeverity(v string)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *AdvisoryRockyAdvisory) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetShortCode

`func (o *AdvisoryRockyAdvisory) GetShortCode() string`

GetShortCode returns the ShortCode field if non-nil, zero value otherwise.

### GetShortCodeOk

`func (o *AdvisoryRockyAdvisory) GetShortCodeOk() (*string, bool)`

GetShortCodeOk returns a tuple with the ShortCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortCode

`func (o *AdvisoryRockyAdvisory) SetShortCode(v string)`

SetShortCode sets ShortCode field to given value.

### HasShortCode

`func (o *AdvisoryRockyAdvisory) HasShortCode() bool`

HasShortCode returns a boolean if a field has been set.

### GetSolution

`func (o *AdvisoryRockyAdvisory) GetSolution() string`

GetSolution returns the Solution field if non-nil, zero value otherwise.

### GetSolutionOk

`func (o *AdvisoryRockyAdvisory) GetSolutionOk() (*string, bool)`

GetSolutionOk returns a tuple with the Solution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSolution

`func (o *AdvisoryRockyAdvisory) SetSolution(v string)`

SetSolution sets Solution field to given value.

### HasSolution

`func (o *AdvisoryRockyAdvisory) HasSolution() bool`

HasSolution returns a boolean if a field has been set.

### GetSynopsis

`func (o *AdvisoryRockyAdvisory) GetSynopsis() string`

GetSynopsis returns the Synopsis field if non-nil, zero value otherwise.

### GetSynopsisOk

`func (o *AdvisoryRockyAdvisory) GetSynopsisOk() (*string, bool)`

GetSynopsisOk returns a tuple with the Synopsis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSynopsis

`func (o *AdvisoryRockyAdvisory) SetSynopsis(v string)`

SetSynopsis sets Synopsis field to given value.

### HasSynopsis

`func (o *AdvisoryRockyAdvisory) HasSynopsis() bool`

HasSynopsis returns a boolean if a field has been set.

### GetTopic

`func (o *AdvisoryRockyAdvisory) GetTopic() string`

GetTopic returns the Topic field if non-nil, zero value otherwise.

### GetTopicOk

`func (o *AdvisoryRockyAdvisory) GetTopicOk() (*string, bool)`

GetTopicOk returns a tuple with the Topic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopic

`func (o *AdvisoryRockyAdvisory) SetTopic(v string)`

SetTopic sets Topic field to given value.

### HasTopic

`func (o *AdvisoryRockyAdvisory) HasTopic() bool`

HasTopic returns a boolean if a field has been set.

### GetType

`func (o *AdvisoryRockyAdvisory) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AdvisoryRockyAdvisory) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AdvisoryRockyAdvisory) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *AdvisoryRockyAdvisory) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


