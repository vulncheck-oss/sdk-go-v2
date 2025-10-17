# AdvisoryRustsecFrontMatterAdvisory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Aliases** | Pointer to **[]string** | Vulnerability aliases, e.g. CVE IDs (optional but recommended) Request a CVE for your RustSec vulns: https://iwantacve.org/ | [optional] 
**Categories** | Pointer to **[]string** | Optional: Categories this advisory falls under. Valid categories are: \&quot;code-execution\&quot;, \&quot;crypto-failure\&quot;, \&quot;denial-of-service\&quot;, \&quot;file-disclosure\&quot; \&quot;format-injection\&quot;, \&quot;memory-corruption\&quot;, \&quot;memory-exposure\&quot;, \&quot;privilege-escalation\&quot; | [optional] 
**Cvss** | Pointer to **string** | Optional: a Common Vulnerability Scoring System score. More information can be found on the CVSS website, https://www.first.org/cvss/. | [optional] 
**Date** | Pointer to **string** | Disclosure date of the advisory as an RFC 3339 date (mandatory) | [optional] 
**Informational** | Pointer to **string** | Optional: Indicates the type of informational security  advisory  - \&quot;unsound\&quot; for soundness issues  - \&quot;unmaintained\&quot; for crates that are no longer maintained  - \&quot;notice\&quot; for other informational notices | [optional] 
**Keywords** | Pointer to **[]string** | Freeform keywords which describe this vulnerability, similar to Cargo (optional) | [optional] 
**Package** | Pointer to **string** | Name of the affected crate (mandatory) | [optional] 
**References** | Pointer to **[]string** | URL to additional helpful references regarding the advisory (optional) | [optional] 
**Related** | Pointer to **[]string** | Related vulnerabilities (optional) e.g. CVE for a C library wrapped by a -sys crate) | [optional] 
**RustsecId** | Pointer to **string** | Identifier for the advisory (mandatory). Will be assigned a \&quot;RUSTSEC-YYYY-NNNN\&quot; identifier e.g. RUSTSEC-2018-0001. Please use \&quot;RUSTSEC-0000-0000\&quot; in PRs. | [optional] 
**Url** | Pointer to **string** | URL to a long-form description of this issue, e.g. a GitHub issue/PR, a change log entry, or a blogpost announcing the release (optional) | [optional] 
**Withdrawn** | Pointer to **string** | Whether the advisory is withdrawn (optional) | [optional] 

## Methods

### NewAdvisoryRustsecFrontMatterAdvisory

`func NewAdvisoryRustsecFrontMatterAdvisory() *AdvisoryRustsecFrontMatterAdvisory`

NewAdvisoryRustsecFrontMatterAdvisory instantiates a new AdvisoryRustsecFrontMatterAdvisory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdvisoryRustsecFrontMatterAdvisoryWithDefaults

`func NewAdvisoryRustsecFrontMatterAdvisoryWithDefaults() *AdvisoryRustsecFrontMatterAdvisory`

NewAdvisoryRustsecFrontMatterAdvisoryWithDefaults instantiates a new AdvisoryRustsecFrontMatterAdvisory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAliases

`func (o *AdvisoryRustsecFrontMatterAdvisory) GetAliases() []string`

GetAliases returns the Aliases field if non-nil, zero value otherwise.

### GetAliasesOk

`func (o *AdvisoryRustsecFrontMatterAdvisory) GetAliasesOk() (*[]string, bool)`

GetAliasesOk returns a tuple with the Aliases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAliases

`func (o *AdvisoryRustsecFrontMatterAdvisory) SetAliases(v []string)`

SetAliases sets Aliases field to given value.

### HasAliases

`func (o *AdvisoryRustsecFrontMatterAdvisory) HasAliases() bool`

HasAliases returns a boolean if a field has been set.

### GetCategories

`func (o *AdvisoryRustsecFrontMatterAdvisory) GetCategories() []string`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *AdvisoryRustsecFrontMatterAdvisory) GetCategoriesOk() (*[]string, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategories

`func (o *AdvisoryRustsecFrontMatterAdvisory) SetCategories(v []string)`

SetCategories sets Categories field to given value.

### HasCategories

`func (o *AdvisoryRustsecFrontMatterAdvisory) HasCategories() bool`

HasCategories returns a boolean if a field has been set.

### GetCvss

`func (o *AdvisoryRustsecFrontMatterAdvisory) GetCvss() string`

GetCvss returns the Cvss field if non-nil, zero value otherwise.

### GetCvssOk

`func (o *AdvisoryRustsecFrontMatterAdvisory) GetCvssOk() (*string, bool)`

GetCvssOk returns a tuple with the Cvss field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCvss

`func (o *AdvisoryRustsecFrontMatterAdvisory) SetCvss(v string)`

SetCvss sets Cvss field to given value.

### HasCvss

`func (o *AdvisoryRustsecFrontMatterAdvisory) HasCvss() bool`

HasCvss returns a boolean if a field has been set.

### GetDate

`func (o *AdvisoryRustsecFrontMatterAdvisory) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *AdvisoryRustsecFrontMatterAdvisory) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *AdvisoryRustsecFrontMatterAdvisory) SetDate(v string)`

SetDate sets Date field to given value.

### HasDate

`func (o *AdvisoryRustsecFrontMatterAdvisory) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetInformational

`func (o *AdvisoryRustsecFrontMatterAdvisory) GetInformational() string`

GetInformational returns the Informational field if non-nil, zero value otherwise.

### GetInformationalOk

`func (o *AdvisoryRustsecFrontMatterAdvisory) GetInformationalOk() (*string, bool)`

GetInformationalOk returns a tuple with the Informational field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInformational

`func (o *AdvisoryRustsecFrontMatterAdvisory) SetInformational(v string)`

SetInformational sets Informational field to given value.

### HasInformational

`func (o *AdvisoryRustsecFrontMatterAdvisory) HasInformational() bool`

HasInformational returns a boolean if a field has been set.

### GetKeywords

`func (o *AdvisoryRustsecFrontMatterAdvisory) GetKeywords() []string`

GetKeywords returns the Keywords field if non-nil, zero value otherwise.

### GetKeywordsOk

`func (o *AdvisoryRustsecFrontMatterAdvisory) GetKeywordsOk() (*[]string, bool)`

GetKeywordsOk returns a tuple with the Keywords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeywords

`func (o *AdvisoryRustsecFrontMatterAdvisory) SetKeywords(v []string)`

SetKeywords sets Keywords field to given value.

### HasKeywords

`func (o *AdvisoryRustsecFrontMatterAdvisory) HasKeywords() bool`

HasKeywords returns a boolean if a field has been set.

### GetPackage

`func (o *AdvisoryRustsecFrontMatterAdvisory) GetPackage() string`

GetPackage returns the Package field if non-nil, zero value otherwise.

### GetPackageOk

`func (o *AdvisoryRustsecFrontMatterAdvisory) GetPackageOk() (*string, bool)`

GetPackageOk returns a tuple with the Package field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackage

`func (o *AdvisoryRustsecFrontMatterAdvisory) SetPackage(v string)`

SetPackage sets Package field to given value.

### HasPackage

`func (o *AdvisoryRustsecFrontMatterAdvisory) HasPackage() bool`

HasPackage returns a boolean if a field has been set.

### GetReferences

`func (o *AdvisoryRustsecFrontMatterAdvisory) GetReferences() []string`

GetReferences returns the References field if non-nil, zero value otherwise.

### GetReferencesOk

`func (o *AdvisoryRustsecFrontMatterAdvisory) GetReferencesOk() (*[]string, bool)`

GetReferencesOk returns a tuple with the References field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferences

`func (o *AdvisoryRustsecFrontMatterAdvisory) SetReferences(v []string)`

SetReferences sets References field to given value.

### HasReferences

`func (o *AdvisoryRustsecFrontMatterAdvisory) HasReferences() bool`

HasReferences returns a boolean if a field has been set.

### GetRelated

`func (o *AdvisoryRustsecFrontMatterAdvisory) GetRelated() []string`

GetRelated returns the Related field if non-nil, zero value otherwise.

### GetRelatedOk

`func (o *AdvisoryRustsecFrontMatterAdvisory) GetRelatedOk() (*[]string, bool)`

GetRelatedOk returns a tuple with the Related field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelated

`func (o *AdvisoryRustsecFrontMatterAdvisory) SetRelated(v []string)`

SetRelated sets Related field to given value.

### HasRelated

`func (o *AdvisoryRustsecFrontMatterAdvisory) HasRelated() bool`

HasRelated returns a boolean if a field has been set.

### GetRustsecId

`func (o *AdvisoryRustsecFrontMatterAdvisory) GetRustsecId() string`

GetRustsecId returns the RustsecId field if non-nil, zero value otherwise.

### GetRustsecIdOk

`func (o *AdvisoryRustsecFrontMatterAdvisory) GetRustsecIdOk() (*string, bool)`

GetRustsecIdOk returns a tuple with the RustsecId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRustsecId

`func (o *AdvisoryRustsecFrontMatterAdvisory) SetRustsecId(v string)`

SetRustsecId sets RustsecId field to given value.

### HasRustsecId

`func (o *AdvisoryRustsecFrontMatterAdvisory) HasRustsecId() bool`

HasRustsecId returns a boolean if a field has been set.

### GetUrl

`func (o *AdvisoryRustsecFrontMatterAdvisory) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *AdvisoryRustsecFrontMatterAdvisory) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *AdvisoryRustsecFrontMatterAdvisory) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *AdvisoryRustsecFrontMatterAdvisory) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetWithdrawn

`func (o *AdvisoryRustsecFrontMatterAdvisory) GetWithdrawn() string`

GetWithdrawn returns the Withdrawn field if non-nil, zero value otherwise.

### GetWithdrawnOk

`func (o *AdvisoryRustsecFrontMatterAdvisory) GetWithdrawnOk() (*string, bool)`

GetWithdrawnOk returns a tuple with the Withdrawn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithdrawn

`func (o *AdvisoryRustsecFrontMatterAdvisory) SetWithdrawn(v string)`

SetWithdrawn sets Withdrawn field to given value.

### HasWithdrawn

`func (o *AdvisoryRustsecFrontMatterAdvisory) HasWithdrawn() bool`

HasWithdrawn returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


