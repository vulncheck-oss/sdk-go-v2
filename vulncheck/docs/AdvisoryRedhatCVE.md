# AdvisoryRedhatCVE

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Advisories** | Pointer to **[]string** |  | [optional] 
**AdvisoryCsafVexUrl** | Pointer to **[]string** |  | [optional] 
**AffectedPackages** | Pointer to **[]string** | used for un-marshlling from redhat | [optional] 
**AffectedRelease** | Pointer to [**[]AdvisoryAffectedRel**](AdvisoryAffectedRel.md) |  | [optional] 
**Bugzilla** | Pointer to **string** |  | [optional] 
**BugzillaDescription** | Pointer to **string** |  | [optional] 
**Cve** | Pointer to **[]string** |  | [optional] 
**CveCsafVexUrl** | Pointer to **string** |  | [optional] 
**Cvss3Score** | Pointer to **string** |  | [optional] 
**Cvss3ScoringVector** | Pointer to **string** |  | [optional] 
**CvssScore** | Pointer to **float32** |  | [optional] 
**CvssScoringVector** | Pointer to **string** |  | [optional] 
**Cwe** | Pointer to **string** |  | [optional] 
**PackageState** | Pointer to [**[]AdvisoryPackageStat**](AdvisoryPackageStat.md) |  | [optional] 
**Packages** | Pointer to [**[]AdvisoryVulnCheckPackage**](AdvisoryVulnCheckPackage.md) | used to index into vulncheck OS | [optional] 
**PublicDate** | Pointer to **string** |  | [optional] 
**ResourceUrl** | Pointer to **string** |  | [optional] 
**Severity** | Pointer to **string** |  | [optional] 

## Methods

### NewAdvisoryRedhatCVE

`func NewAdvisoryRedhatCVE() *AdvisoryRedhatCVE`

NewAdvisoryRedhatCVE instantiates a new AdvisoryRedhatCVE object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdvisoryRedhatCVEWithDefaults

`func NewAdvisoryRedhatCVEWithDefaults() *AdvisoryRedhatCVE`

NewAdvisoryRedhatCVEWithDefaults instantiates a new AdvisoryRedhatCVE object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdvisories

`func (o *AdvisoryRedhatCVE) GetAdvisories() []string`

GetAdvisories returns the Advisories field if non-nil, zero value otherwise.

### GetAdvisoriesOk

`func (o *AdvisoryRedhatCVE) GetAdvisoriesOk() (*[]string, bool)`

GetAdvisoriesOk returns a tuple with the Advisories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvisories

`func (o *AdvisoryRedhatCVE) SetAdvisories(v []string)`

SetAdvisories sets Advisories field to given value.

### HasAdvisories

`func (o *AdvisoryRedhatCVE) HasAdvisories() bool`

HasAdvisories returns a boolean if a field has been set.

### GetAdvisoryCsafVexUrl

`func (o *AdvisoryRedhatCVE) GetAdvisoryCsafVexUrl() []string`

GetAdvisoryCsafVexUrl returns the AdvisoryCsafVexUrl field if non-nil, zero value otherwise.

### GetAdvisoryCsafVexUrlOk

`func (o *AdvisoryRedhatCVE) GetAdvisoryCsafVexUrlOk() (*[]string, bool)`

GetAdvisoryCsafVexUrlOk returns a tuple with the AdvisoryCsafVexUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvisoryCsafVexUrl

`func (o *AdvisoryRedhatCVE) SetAdvisoryCsafVexUrl(v []string)`

SetAdvisoryCsafVexUrl sets AdvisoryCsafVexUrl field to given value.

### HasAdvisoryCsafVexUrl

`func (o *AdvisoryRedhatCVE) HasAdvisoryCsafVexUrl() bool`

HasAdvisoryCsafVexUrl returns a boolean if a field has been set.

### GetAffectedPackages

`func (o *AdvisoryRedhatCVE) GetAffectedPackages() []string`

GetAffectedPackages returns the AffectedPackages field if non-nil, zero value otherwise.

### GetAffectedPackagesOk

`func (o *AdvisoryRedhatCVE) GetAffectedPackagesOk() (*[]string, bool)`

GetAffectedPackagesOk returns a tuple with the AffectedPackages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffectedPackages

`func (o *AdvisoryRedhatCVE) SetAffectedPackages(v []string)`

SetAffectedPackages sets AffectedPackages field to given value.

### HasAffectedPackages

`func (o *AdvisoryRedhatCVE) HasAffectedPackages() bool`

HasAffectedPackages returns a boolean if a field has been set.

### GetAffectedRelease

`func (o *AdvisoryRedhatCVE) GetAffectedRelease() []AdvisoryAffectedRel`

GetAffectedRelease returns the AffectedRelease field if non-nil, zero value otherwise.

### GetAffectedReleaseOk

`func (o *AdvisoryRedhatCVE) GetAffectedReleaseOk() (*[]AdvisoryAffectedRel, bool)`

GetAffectedReleaseOk returns a tuple with the AffectedRelease field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffectedRelease

`func (o *AdvisoryRedhatCVE) SetAffectedRelease(v []AdvisoryAffectedRel)`

SetAffectedRelease sets AffectedRelease field to given value.

### HasAffectedRelease

`func (o *AdvisoryRedhatCVE) HasAffectedRelease() bool`

HasAffectedRelease returns a boolean if a field has been set.

### GetBugzilla

`func (o *AdvisoryRedhatCVE) GetBugzilla() string`

GetBugzilla returns the Bugzilla field if non-nil, zero value otherwise.

### GetBugzillaOk

`func (o *AdvisoryRedhatCVE) GetBugzillaOk() (*string, bool)`

GetBugzillaOk returns a tuple with the Bugzilla field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBugzilla

`func (o *AdvisoryRedhatCVE) SetBugzilla(v string)`

SetBugzilla sets Bugzilla field to given value.

### HasBugzilla

`func (o *AdvisoryRedhatCVE) HasBugzilla() bool`

HasBugzilla returns a boolean if a field has been set.

### GetBugzillaDescription

`func (o *AdvisoryRedhatCVE) GetBugzillaDescription() string`

GetBugzillaDescription returns the BugzillaDescription field if non-nil, zero value otherwise.

### GetBugzillaDescriptionOk

`func (o *AdvisoryRedhatCVE) GetBugzillaDescriptionOk() (*string, bool)`

GetBugzillaDescriptionOk returns a tuple with the BugzillaDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBugzillaDescription

`func (o *AdvisoryRedhatCVE) SetBugzillaDescription(v string)`

SetBugzillaDescription sets BugzillaDescription field to given value.

### HasBugzillaDescription

`func (o *AdvisoryRedhatCVE) HasBugzillaDescription() bool`

HasBugzillaDescription returns a boolean if a field has been set.

### GetCve

`func (o *AdvisoryRedhatCVE) GetCve() []string`

GetCve returns the Cve field if non-nil, zero value otherwise.

### GetCveOk

`func (o *AdvisoryRedhatCVE) GetCveOk() (*[]string, bool)`

GetCveOk returns a tuple with the Cve field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCve

`func (o *AdvisoryRedhatCVE) SetCve(v []string)`

SetCve sets Cve field to given value.

### HasCve

`func (o *AdvisoryRedhatCVE) HasCve() bool`

HasCve returns a boolean if a field has been set.

### GetCveCsafVexUrl

`func (o *AdvisoryRedhatCVE) GetCveCsafVexUrl() string`

GetCveCsafVexUrl returns the CveCsafVexUrl field if non-nil, zero value otherwise.

### GetCveCsafVexUrlOk

`func (o *AdvisoryRedhatCVE) GetCveCsafVexUrlOk() (*string, bool)`

GetCveCsafVexUrlOk returns a tuple with the CveCsafVexUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCveCsafVexUrl

`func (o *AdvisoryRedhatCVE) SetCveCsafVexUrl(v string)`

SetCveCsafVexUrl sets CveCsafVexUrl field to given value.

### HasCveCsafVexUrl

`func (o *AdvisoryRedhatCVE) HasCveCsafVexUrl() bool`

HasCveCsafVexUrl returns a boolean if a field has been set.

### GetCvss3Score

`func (o *AdvisoryRedhatCVE) GetCvss3Score() string`

GetCvss3Score returns the Cvss3Score field if non-nil, zero value otherwise.

### GetCvss3ScoreOk

`func (o *AdvisoryRedhatCVE) GetCvss3ScoreOk() (*string, bool)`

GetCvss3ScoreOk returns a tuple with the Cvss3Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCvss3Score

`func (o *AdvisoryRedhatCVE) SetCvss3Score(v string)`

SetCvss3Score sets Cvss3Score field to given value.

### HasCvss3Score

`func (o *AdvisoryRedhatCVE) HasCvss3Score() bool`

HasCvss3Score returns a boolean if a field has been set.

### GetCvss3ScoringVector

`func (o *AdvisoryRedhatCVE) GetCvss3ScoringVector() string`

GetCvss3ScoringVector returns the Cvss3ScoringVector field if non-nil, zero value otherwise.

### GetCvss3ScoringVectorOk

`func (o *AdvisoryRedhatCVE) GetCvss3ScoringVectorOk() (*string, bool)`

GetCvss3ScoringVectorOk returns a tuple with the Cvss3ScoringVector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCvss3ScoringVector

`func (o *AdvisoryRedhatCVE) SetCvss3ScoringVector(v string)`

SetCvss3ScoringVector sets Cvss3ScoringVector field to given value.

### HasCvss3ScoringVector

`func (o *AdvisoryRedhatCVE) HasCvss3ScoringVector() bool`

HasCvss3ScoringVector returns a boolean if a field has been set.

### GetCvssScore

`func (o *AdvisoryRedhatCVE) GetCvssScore() float32`

GetCvssScore returns the CvssScore field if non-nil, zero value otherwise.

### GetCvssScoreOk

`func (o *AdvisoryRedhatCVE) GetCvssScoreOk() (*float32, bool)`

GetCvssScoreOk returns a tuple with the CvssScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCvssScore

`func (o *AdvisoryRedhatCVE) SetCvssScore(v float32)`

SetCvssScore sets CvssScore field to given value.

### HasCvssScore

`func (o *AdvisoryRedhatCVE) HasCvssScore() bool`

HasCvssScore returns a boolean if a field has been set.

### GetCvssScoringVector

`func (o *AdvisoryRedhatCVE) GetCvssScoringVector() string`

GetCvssScoringVector returns the CvssScoringVector field if non-nil, zero value otherwise.

### GetCvssScoringVectorOk

`func (o *AdvisoryRedhatCVE) GetCvssScoringVectorOk() (*string, bool)`

GetCvssScoringVectorOk returns a tuple with the CvssScoringVector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCvssScoringVector

`func (o *AdvisoryRedhatCVE) SetCvssScoringVector(v string)`

SetCvssScoringVector sets CvssScoringVector field to given value.

### HasCvssScoringVector

`func (o *AdvisoryRedhatCVE) HasCvssScoringVector() bool`

HasCvssScoringVector returns a boolean if a field has been set.

### GetCwe

`func (o *AdvisoryRedhatCVE) GetCwe() string`

GetCwe returns the Cwe field if non-nil, zero value otherwise.

### GetCweOk

`func (o *AdvisoryRedhatCVE) GetCweOk() (*string, bool)`

GetCweOk returns a tuple with the Cwe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCwe

`func (o *AdvisoryRedhatCVE) SetCwe(v string)`

SetCwe sets Cwe field to given value.

### HasCwe

`func (o *AdvisoryRedhatCVE) HasCwe() bool`

HasCwe returns a boolean if a field has been set.

### GetPackageState

`func (o *AdvisoryRedhatCVE) GetPackageState() []AdvisoryPackageStat`

GetPackageState returns the PackageState field if non-nil, zero value otherwise.

### GetPackageStateOk

`func (o *AdvisoryRedhatCVE) GetPackageStateOk() (*[]AdvisoryPackageStat, bool)`

GetPackageStateOk returns a tuple with the PackageState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageState

`func (o *AdvisoryRedhatCVE) SetPackageState(v []AdvisoryPackageStat)`

SetPackageState sets PackageState field to given value.

### HasPackageState

`func (o *AdvisoryRedhatCVE) HasPackageState() bool`

HasPackageState returns a boolean if a field has been set.

### GetPackages

`func (o *AdvisoryRedhatCVE) GetPackages() []AdvisoryVulnCheckPackage`

GetPackages returns the Packages field if non-nil, zero value otherwise.

### GetPackagesOk

`func (o *AdvisoryRedhatCVE) GetPackagesOk() (*[]AdvisoryVulnCheckPackage, bool)`

GetPackagesOk returns a tuple with the Packages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackages

`func (o *AdvisoryRedhatCVE) SetPackages(v []AdvisoryVulnCheckPackage)`

SetPackages sets Packages field to given value.

### HasPackages

`func (o *AdvisoryRedhatCVE) HasPackages() bool`

HasPackages returns a boolean if a field has been set.

### GetPublicDate

`func (o *AdvisoryRedhatCVE) GetPublicDate() string`

GetPublicDate returns the PublicDate field if non-nil, zero value otherwise.

### GetPublicDateOk

`func (o *AdvisoryRedhatCVE) GetPublicDateOk() (*string, bool)`

GetPublicDateOk returns a tuple with the PublicDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicDate

`func (o *AdvisoryRedhatCVE) SetPublicDate(v string)`

SetPublicDate sets PublicDate field to given value.

### HasPublicDate

`func (o *AdvisoryRedhatCVE) HasPublicDate() bool`

HasPublicDate returns a boolean if a field has been set.

### GetResourceUrl

`func (o *AdvisoryRedhatCVE) GetResourceUrl() string`

GetResourceUrl returns the ResourceUrl field if non-nil, zero value otherwise.

### GetResourceUrlOk

`func (o *AdvisoryRedhatCVE) GetResourceUrlOk() (*string, bool)`

GetResourceUrlOk returns a tuple with the ResourceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceUrl

`func (o *AdvisoryRedhatCVE) SetResourceUrl(v string)`

SetResourceUrl sets ResourceUrl field to given value.

### HasResourceUrl

`func (o *AdvisoryRedhatCVE) HasResourceUrl() bool`

HasResourceUrl returns a boolean if a field has been set.

### GetSeverity

`func (o *AdvisoryRedhatCVE) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *AdvisoryRedhatCVE) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *AdvisoryRedhatCVE) SetSeverity(v string)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *AdvisoryRedhatCVE) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


