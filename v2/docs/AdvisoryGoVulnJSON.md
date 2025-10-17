# AdvisoryGoVulnJSON

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdvisoryUrl** | Pointer to **string** |  | [optional] 
**Affected** | Pointer to [**[]AdvisoryGoVulnAffected**](AdvisoryGoVulnAffected.md) |  | [optional] 
**Aliases** | Pointer to **[]string** |  | [optional] 
**Credits** | Pointer to [**[]AdvisoryGoCredits**](AdvisoryGoCredits.md) |  | [optional] 
**Cve** | Pointer to **[]string** |  | [optional] 
**DateAdded** | Pointer to **string** |  | [optional] 
**Details** | Pointer to **string** |  | [optional] 
**Ghsa** | Pointer to **[]string** |  | [optional] 
**GoAdvisoryId** | Pointer to **string** |  | [optional] 
**Modified** | Pointer to **string** |  | [optional] 
**Published** | Pointer to **string** |  | [optional] 
**References** | Pointer to [**[]AdvisoryGoVulnReference**](AdvisoryGoVulnReference.md) |  | [optional] 

## Methods

### NewAdvisoryGoVulnJSON

`func NewAdvisoryGoVulnJSON() *AdvisoryGoVulnJSON`

NewAdvisoryGoVulnJSON instantiates a new AdvisoryGoVulnJSON object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdvisoryGoVulnJSONWithDefaults

`func NewAdvisoryGoVulnJSONWithDefaults() *AdvisoryGoVulnJSON`

NewAdvisoryGoVulnJSONWithDefaults instantiates a new AdvisoryGoVulnJSON object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdvisoryUrl

`func (o *AdvisoryGoVulnJSON) GetAdvisoryUrl() string`

GetAdvisoryUrl returns the AdvisoryUrl field if non-nil, zero value otherwise.

### GetAdvisoryUrlOk

`func (o *AdvisoryGoVulnJSON) GetAdvisoryUrlOk() (*string, bool)`

GetAdvisoryUrlOk returns a tuple with the AdvisoryUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvisoryUrl

`func (o *AdvisoryGoVulnJSON) SetAdvisoryUrl(v string)`

SetAdvisoryUrl sets AdvisoryUrl field to given value.

### HasAdvisoryUrl

`func (o *AdvisoryGoVulnJSON) HasAdvisoryUrl() bool`

HasAdvisoryUrl returns a boolean if a field has been set.

### GetAffected

`func (o *AdvisoryGoVulnJSON) GetAffected() []AdvisoryGoVulnAffected`

GetAffected returns the Affected field if non-nil, zero value otherwise.

### GetAffectedOk

`func (o *AdvisoryGoVulnJSON) GetAffectedOk() (*[]AdvisoryGoVulnAffected, bool)`

GetAffectedOk returns a tuple with the Affected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffected

`func (o *AdvisoryGoVulnJSON) SetAffected(v []AdvisoryGoVulnAffected)`

SetAffected sets Affected field to given value.

### HasAffected

`func (o *AdvisoryGoVulnJSON) HasAffected() bool`

HasAffected returns a boolean if a field has been set.

### GetAliases

`func (o *AdvisoryGoVulnJSON) GetAliases() []string`

GetAliases returns the Aliases field if non-nil, zero value otherwise.

### GetAliasesOk

`func (o *AdvisoryGoVulnJSON) GetAliasesOk() (*[]string, bool)`

GetAliasesOk returns a tuple with the Aliases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAliases

`func (o *AdvisoryGoVulnJSON) SetAliases(v []string)`

SetAliases sets Aliases field to given value.

### HasAliases

`func (o *AdvisoryGoVulnJSON) HasAliases() bool`

HasAliases returns a boolean if a field has been set.

### GetCredits

`func (o *AdvisoryGoVulnJSON) GetCredits() []AdvisoryGoCredits`

GetCredits returns the Credits field if non-nil, zero value otherwise.

### GetCreditsOk

`func (o *AdvisoryGoVulnJSON) GetCreditsOk() (*[]AdvisoryGoCredits, bool)`

GetCreditsOk returns a tuple with the Credits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredits

`func (o *AdvisoryGoVulnJSON) SetCredits(v []AdvisoryGoCredits)`

SetCredits sets Credits field to given value.

### HasCredits

`func (o *AdvisoryGoVulnJSON) HasCredits() bool`

HasCredits returns a boolean if a field has been set.

### GetCve

`func (o *AdvisoryGoVulnJSON) GetCve() []string`

GetCve returns the Cve field if non-nil, zero value otherwise.

### GetCveOk

`func (o *AdvisoryGoVulnJSON) GetCveOk() (*[]string, bool)`

GetCveOk returns a tuple with the Cve field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCve

`func (o *AdvisoryGoVulnJSON) SetCve(v []string)`

SetCve sets Cve field to given value.

### HasCve

`func (o *AdvisoryGoVulnJSON) HasCve() bool`

HasCve returns a boolean if a field has been set.

### GetDateAdded

`func (o *AdvisoryGoVulnJSON) GetDateAdded() string`

GetDateAdded returns the DateAdded field if non-nil, zero value otherwise.

### GetDateAddedOk

`func (o *AdvisoryGoVulnJSON) GetDateAddedOk() (*string, bool)`

GetDateAddedOk returns a tuple with the DateAdded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateAdded

`func (o *AdvisoryGoVulnJSON) SetDateAdded(v string)`

SetDateAdded sets DateAdded field to given value.

### HasDateAdded

`func (o *AdvisoryGoVulnJSON) HasDateAdded() bool`

HasDateAdded returns a boolean if a field has been set.

### GetDetails

`func (o *AdvisoryGoVulnJSON) GetDetails() string`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *AdvisoryGoVulnJSON) GetDetailsOk() (*string, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *AdvisoryGoVulnJSON) SetDetails(v string)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *AdvisoryGoVulnJSON) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### GetGhsa

`func (o *AdvisoryGoVulnJSON) GetGhsa() []string`

GetGhsa returns the Ghsa field if non-nil, zero value otherwise.

### GetGhsaOk

`func (o *AdvisoryGoVulnJSON) GetGhsaOk() (*[]string, bool)`

GetGhsaOk returns a tuple with the Ghsa field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGhsa

`func (o *AdvisoryGoVulnJSON) SetGhsa(v []string)`

SetGhsa sets Ghsa field to given value.

### HasGhsa

`func (o *AdvisoryGoVulnJSON) HasGhsa() bool`

HasGhsa returns a boolean if a field has been set.

### GetGoAdvisoryId

`func (o *AdvisoryGoVulnJSON) GetGoAdvisoryId() string`

GetGoAdvisoryId returns the GoAdvisoryId field if non-nil, zero value otherwise.

### GetGoAdvisoryIdOk

`func (o *AdvisoryGoVulnJSON) GetGoAdvisoryIdOk() (*string, bool)`

GetGoAdvisoryIdOk returns a tuple with the GoAdvisoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoAdvisoryId

`func (o *AdvisoryGoVulnJSON) SetGoAdvisoryId(v string)`

SetGoAdvisoryId sets GoAdvisoryId field to given value.

### HasGoAdvisoryId

`func (o *AdvisoryGoVulnJSON) HasGoAdvisoryId() bool`

HasGoAdvisoryId returns a boolean if a field has been set.

### GetModified

`func (o *AdvisoryGoVulnJSON) GetModified() string`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *AdvisoryGoVulnJSON) GetModifiedOk() (*string, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *AdvisoryGoVulnJSON) SetModified(v string)`

SetModified sets Modified field to given value.

### HasModified

`func (o *AdvisoryGoVulnJSON) HasModified() bool`

HasModified returns a boolean if a field has been set.

### GetPublished

`func (o *AdvisoryGoVulnJSON) GetPublished() string`

GetPublished returns the Published field if non-nil, zero value otherwise.

### GetPublishedOk

`func (o *AdvisoryGoVulnJSON) GetPublishedOk() (*string, bool)`

GetPublishedOk returns a tuple with the Published field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublished

`func (o *AdvisoryGoVulnJSON) SetPublished(v string)`

SetPublished sets Published field to given value.

### HasPublished

`func (o *AdvisoryGoVulnJSON) HasPublished() bool`

HasPublished returns a boolean if a field has been set.

### GetReferences

`func (o *AdvisoryGoVulnJSON) GetReferences() []AdvisoryGoVulnReference`

GetReferences returns the References field if non-nil, zero value otherwise.

### GetReferencesOk

`func (o *AdvisoryGoVulnJSON) GetReferencesOk() (*[]AdvisoryGoVulnReference, bool)`

GetReferencesOk returns a tuple with the References field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferences

`func (o *AdvisoryGoVulnJSON) SetReferences(v []AdvisoryGoVulnReference)`

SetReferences sets References field to given value.

### HasReferences

`func (o *AdvisoryGoVulnJSON) HasReferences() bool`

HasReferences returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


