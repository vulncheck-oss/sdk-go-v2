# AdvisoryWRT

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Advisory** | Pointer to **string** |  | [optional] 
**AffectedVersions** | Pointer to **string** |  | [optional] 
**Credits** | Pointer to **string** |  | [optional] 
**Cve** | Pointer to **[]string** |  | [optional] 
**DateAdded** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Mitigations** | Pointer to **string** |  | [optional] 
**References** | Pointer to **[]string** |  | [optional] 
**Requirements** | Pointer to **string** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 

## Methods

### NewAdvisoryWRT

`func NewAdvisoryWRT() *AdvisoryWRT`

NewAdvisoryWRT instantiates a new AdvisoryWRT object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdvisoryWRTWithDefaults

`func NewAdvisoryWRTWithDefaults() *AdvisoryWRT`

NewAdvisoryWRTWithDefaults instantiates a new AdvisoryWRT object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdvisory

`func (o *AdvisoryWRT) GetAdvisory() string`

GetAdvisory returns the Advisory field if non-nil, zero value otherwise.

### GetAdvisoryOk

`func (o *AdvisoryWRT) GetAdvisoryOk() (*string, bool)`

GetAdvisoryOk returns a tuple with the Advisory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvisory

`func (o *AdvisoryWRT) SetAdvisory(v string)`

SetAdvisory sets Advisory field to given value.

### HasAdvisory

`func (o *AdvisoryWRT) HasAdvisory() bool`

HasAdvisory returns a boolean if a field has been set.

### GetAffectedVersions

`func (o *AdvisoryWRT) GetAffectedVersions() string`

GetAffectedVersions returns the AffectedVersions field if non-nil, zero value otherwise.

### GetAffectedVersionsOk

`func (o *AdvisoryWRT) GetAffectedVersionsOk() (*string, bool)`

GetAffectedVersionsOk returns a tuple with the AffectedVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffectedVersions

`func (o *AdvisoryWRT) SetAffectedVersions(v string)`

SetAffectedVersions sets AffectedVersions field to given value.

### HasAffectedVersions

`func (o *AdvisoryWRT) HasAffectedVersions() bool`

HasAffectedVersions returns a boolean if a field has been set.

### GetCredits

`func (o *AdvisoryWRT) GetCredits() string`

GetCredits returns the Credits field if non-nil, zero value otherwise.

### GetCreditsOk

`func (o *AdvisoryWRT) GetCreditsOk() (*string, bool)`

GetCreditsOk returns a tuple with the Credits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredits

`func (o *AdvisoryWRT) SetCredits(v string)`

SetCredits sets Credits field to given value.

### HasCredits

`func (o *AdvisoryWRT) HasCredits() bool`

HasCredits returns a boolean if a field has been set.

### GetCve

`func (o *AdvisoryWRT) GetCve() []string`

GetCve returns the Cve field if non-nil, zero value otherwise.

### GetCveOk

`func (o *AdvisoryWRT) GetCveOk() (*[]string, bool)`

GetCveOk returns a tuple with the Cve field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCve

`func (o *AdvisoryWRT) SetCve(v []string)`

SetCve sets Cve field to given value.

### HasCve

`func (o *AdvisoryWRT) HasCve() bool`

HasCve returns a boolean if a field has been set.

### GetDateAdded

`func (o *AdvisoryWRT) GetDateAdded() string`

GetDateAdded returns the DateAdded field if non-nil, zero value otherwise.

### GetDateAddedOk

`func (o *AdvisoryWRT) GetDateAddedOk() (*string, bool)`

GetDateAddedOk returns a tuple with the DateAdded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateAdded

`func (o *AdvisoryWRT) SetDateAdded(v string)`

SetDateAdded sets DateAdded field to given value.

### HasDateAdded

`func (o *AdvisoryWRT) HasDateAdded() bool`

HasDateAdded returns a boolean if a field has been set.

### GetDescription

`func (o *AdvisoryWRT) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AdvisoryWRT) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AdvisoryWRT) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AdvisoryWRT) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMitigations

`func (o *AdvisoryWRT) GetMitigations() string`

GetMitigations returns the Mitigations field if non-nil, zero value otherwise.

### GetMitigationsOk

`func (o *AdvisoryWRT) GetMitigationsOk() (*string, bool)`

GetMitigationsOk returns a tuple with the Mitigations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMitigations

`func (o *AdvisoryWRT) SetMitigations(v string)`

SetMitigations sets Mitigations field to given value.

### HasMitigations

`func (o *AdvisoryWRT) HasMitigations() bool`

HasMitigations returns a boolean if a field has been set.

### GetReferences

`func (o *AdvisoryWRT) GetReferences() []string`

GetReferences returns the References field if non-nil, zero value otherwise.

### GetReferencesOk

`func (o *AdvisoryWRT) GetReferencesOk() (*[]string, bool)`

GetReferencesOk returns a tuple with the References field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferences

`func (o *AdvisoryWRT) SetReferences(v []string)`

SetReferences sets References field to given value.

### HasReferences

`func (o *AdvisoryWRT) HasReferences() bool`

HasReferences returns a boolean if a field has been set.

### GetRequirements

`func (o *AdvisoryWRT) GetRequirements() string`

GetRequirements returns the Requirements field if non-nil, zero value otherwise.

### GetRequirementsOk

`func (o *AdvisoryWRT) GetRequirementsOk() (*string, bool)`

GetRequirementsOk returns a tuple with the Requirements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequirements

`func (o *AdvisoryWRT) SetRequirements(v string)`

SetRequirements sets Requirements field to given value.

### HasRequirements

`func (o *AdvisoryWRT) HasRequirements() bool`

HasRequirements returns a boolean if a field has been set.

### GetTitle

`func (o *AdvisoryWRT) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *AdvisoryWRT) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *AdvisoryWRT) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *AdvisoryWRT) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetUrl

`func (o *AdvisoryWRT) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *AdvisoryWRT) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *AdvisoryWRT) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *AdvisoryWRT) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


