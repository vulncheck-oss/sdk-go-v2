# AdvisoryGitlabAdvisory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AffectedRange** | Pointer to **string** |  | [optional] 
**AffectedVersions** | Pointer to **string** |  | [optional] 
**Cve** | Pointer to **[]string** |  | [optional] 
**CvssV2** | Pointer to **string** |  | [optional] 
**CvssV3** | Pointer to **string** |  | [optional] 
**Cwe** | Pointer to **[]string** |  | [optional] 
**Date** | Pointer to **string** |  | [optional] 
**DateAdded** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Filename** | Pointer to **string** |  | [optional] 
**FixedVersions** | Pointer to **[]string** |  | [optional] 
**Ghsa** | Pointer to **[]string** |  | [optional] 
**GitlabUrl** | Pointer to **string** |  | [optional] 
**Identifier** | Pointer to **string** |  | [optional] 
**Identifiers** | Pointer to **[]string** |  | [optional] 
**NotImpacted** | Pointer to **string** |  | [optional] 
**PackageManager** | Pointer to **string** |  | [optional] 
**PackageName** | Pointer to **string** |  | [optional] 
**PackageSlug** | Pointer to **string** |  | [optional] 
**Pubdate** | Pointer to **string** |  | [optional] 
**Solution** | Pointer to **string** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**Urls** | Pointer to **[]string** |  | [optional] 
**Uuid** | Pointer to **string** |  | [optional] 

## Methods

### NewAdvisoryGitlabAdvisory

`func NewAdvisoryGitlabAdvisory() *AdvisoryGitlabAdvisory`

NewAdvisoryGitlabAdvisory instantiates a new AdvisoryGitlabAdvisory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdvisoryGitlabAdvisoryWithDefaults

`func NewAdvisoryGitlabAdvisoryWithDefaults() *AdvisoryGitlabAdvisory`

NewAdvisoryGitlabAdvisoryWithDefaults instantiates a new AdvisoryGitlabAdvisory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAffectedRange

`func (o *AdvisoryGitlabAdvisory) GetAffectedRange() string`

GetAffectedRange returns the AffectedRange field if non-nil, zero value otherwise.

### GetAffectedRangeOk

`func (o *AdvisoryGitlabAdvisory) GetAffectedRangeOk() (*string, bool)`

GetAffectedRangeOk returns a tuple with the AffectedRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffectedRange

`func (o *AdvisoryGitlabAdvisory) SetAffectedRange(v string)`

SetAffectedRange sets AffectedRange field to given value.

### HasAffectedRange

`func (o *AdvisoryGitlabAdvisory) HasAffectedRange() bool`

HasAffectedRange returns a boolean if a field has been set.

### GetAffectedVersions

`func (o *AdvisoryGitlabAdvisory) GetAffectedVersions() string`

GetAffectedVersions returns the AffectedVersions field if non-nil, zero value otherwise.

### GetAffectedVersionsOk

`func (o *AdvisoryGitlabAdvisory) GetAffectedVersionsOk() (*string, bool)`

GetAffectedVersionsOk returns a tuple with the AffectedVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffectedVersions

`func (o *AdvisoryGitlabAdvisory) SetAffectedVersions(v string)`

SetAffectedVersions sets AffectedVersions field to given value.

### HasAffectedVersions

`func (o *AdvisoryGitlabAdvisory) HasAffectedVersions() bool`

HasAffectedVersions returns a boolean if a field has been set.

### GetCve

`func (o *AdvisoryGitlabAdvisory) GetCve() []string`

GetCve returns the Cve field if non-nil, zero value otherwise.

### GetCveOk

`func (o *AdvisoryGitlabAdvisory) GetCveOk() (*[]string, bool)`

GetCveOk returns a tuple with the Cve field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCve

`func (o *AdvisoryGitlabAdvisory) SetCve(v []string)`

SetCve sets Cve field to given value.

### HasCve

`func (o *AdvisoryGitlabAdvisory) HasCve() bool`

HasCve returns a boolean if a field has been set.

### GetCvssV2

`func (o *AdvisoryGitlabAdvisory) GetCvssV2() string`

GetCvssV2 returns the CvssV2 field if non-nil, zero value otherwise.

### GetCvssV2Ok

`func (o *AdvisoryGitlabAdvisory) GetCvssV2Ok() (*string, bool)`

GetCvssV2Ok returns a tuple with the CvssV2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCvssV2

`func (o *AdvisoryGitlabAdvisory) SetCvssV2(v string)`

SetCvssV2 sets CvssV2 field to given value.

### HasCvssV2

`func (o *AdvisoryGitlabAdvisory) HasCvssV2() bool`

HasCvssV2 returns a boolean if a field has been set.

### GetCvssV3

`func (o *AdvisoryGitlabAdvisory) GetCvssV3() string`

GetCvssV3 returns the CvssV3 field if non-nil, zero value otherwise.

### GetCvssV3Ok

`func (o *AdvisoryGitlabAdvisory) GetCvssV3Ok() (*string, bool)`

GetCvssV3Ok returns a tuple with the CvssV3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCvssV3

`func (o *AdvisoryGitlabAdvisory) SetCvssV3(v string)`

SetCvssV3 sets CvssV3 field to given value.

### HasCvssV3

`func (o *AdvisoryGitlabAdvisory) HasCvssV3() bool`

HasCvssV3 returns a boolean if a field has been set.

### GetCwe

`func (o *AdvisoryGitlabAdvisory) GetCwe() []string`

GetCwe returns the Cwe field if non-nil, zero value otherwise.

### GetCweOk

`func (o *AdvisoryGitlabAdvisory) GetCweOk() (*[]string, bool)`

GetCweOk returns a tuple with the Cwe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCwe

`func (o *AdvisoryGitlabAdvisory) SetCwe(v []string)`

SetCwe sets Cwe field to given value.

### HasCwe

`func (o *AdvisoryGitlabAdvisory) HasCwe() bool`

HasCwe returns a boolean if a field has been set.

### GetDate

`func (o *AdvisoryGitlabAdvisory) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *AdvisoryGitlabAdvisory) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *AdvisoryGitlabAdvisory) SetDate(v string)`

SetDate sets Date field to given value.

### HasDate

`func (o *AdvisoryGitlabAdvisory) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetDateAdded

`func (o *AdvisoryGitlabAdvisory) GetDateAdded() string`

GetDateAdded returns the DateAdded field if non-nil, zero value otherwise.

### GetDateAddedOk

`func (o *AdvisoryGitlabAdvisory) GetDateAddedOk() (*string, bool)`

GetDateAddedOk returns a tuple with the DateAdded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateAdded

`func (o *AdvisoryGitlabAdvisory) SetDateAdded(v string)`

SetDateAdded sets DateAdded field to given value.

### HasDateAdded

`func (o *AdvisoryGitlabAdvisory) HasDateAdded() bool`

HasDateAdded returns a boolean if a field has been set.

### GetDescription

`func (o *AdvisoryGitlabAdvisory) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AdvisoryGitlabAdvisory) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AdvisoryGitlabAdvisory) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AdvisoryGitlabAdvisory) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFilename

`func (o *AdvisoryGitlabAdvisory) GetFilename() string`

GetFilename returns the Filename field if non-nil, zero value otherwise.

### GetFilenameOk

`func (o *AdvisoryGitlabAdvisory) GetFilenameOk() (*string, bool)`

GetFilenameOk returns a tuple with the Filename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilename

`func (o *AdvisoryGitlabAdvisory) SetFilename(v string)`

SetFilename sets Filename field to given value.

### HasFilename

`func (o *AdvisoryGitlabAdvisory) HasFilename() bool`

HasFilename returns a boolean if a field has been set.

### GetFixedVersions

`func (o *AdvisoryGitlabAdvisory) GetFixedVersions() []string`

GetFixedVersions returns the FixedVersions field if non-nil, zero value otherwise.

### GetFixedVersionsOk

`func (o *AdvisoryGitlabAdvisory) GetFixedVersionsOk() (*[]string, bool)`

GetFixedVersionsOk returns a tuple with the FixedVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFixedVersions

`func (o *AdvisoryGitlabAdvisory) SetFixedVersions(v []string)`

SetFixedVersions sets FixedVersions field to given value.

### HasFixedVersions

`func (o *AdvisoryGitlabAdvisory) HasFixedVersions() bool`

HasFixedVersions returns a boolean if a field has been set.

### GetGhsa

`func (o *AdvisoryGitlabAdvisory) GetGhsa() []string`

GetGhsa returns the Ghsa field if non-nil, zero value otherwise.

### GetGhsaOk

`func (o *AdvisoryGitlabAdvisory) GetGhsaOk() (*[]string, bool)`

GetGhsaOk returns a tuple with the Ghsa field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGhsa

`func (o *AdvisoryGitlabAdvisory) SetGhsa(v []string)`

SetGhsa sets Ghsa field to given value.

### HasGhsa

`func (o *AdvisoryGitlabAdvisory) HasGhsa() bool`

HasGhsa returns a boolean if a field has been set.

### GetGitlabUrl

`func (o *AdvisoryGitlabAdvisory) GetGitlabUrl() string`

GetGitlabUrl returns the GitlabUrl field if non-nil, zero value otherwise.

### GetGitlabUrlOk

`func (o *AdvisoryGitlabAdvisory) GetGitlabUrlOk() (*string, bool)`

GetGitlabUrlOk returns a tuple with the GitlabUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitlabUrl

`func (o *AdvisoryGitlabAdvisory) SetGitlabUrl(v string)`

SetGitlabUrl sets GitlabUrl field to given value.

### HasGitlabUrl

`func (o *AdvisoryGitlabAdvisory) HasGitlabUrl() bool`

HasGitlabUrl returns a boolean if a field has been set.

### GetIdentifier

`func (o *AdvisoryGitlabAdvisory) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *AdvisoryGitlabAdvisory) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *AdvisoryGitlabAdvisory) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.

### HasIdentifier

`func (o *AdvisoryGitlabAdvisory) HasIdentifier() bool`

HasIdentifier returns a boolean if a field has been set.

### GetIdentifiers

`func (o *AdvisoryGitlabAdvisory) GetIdentifiers() []string`

GetIdentifiers returns the Identifiers field if non-nil, zero value otherwise.

### GetIdentifiersOk

`func (o *AdvisoryGitlabAdvisory) GetIdentifiersOk() (*[]string, bool)`

GetIdentifiersOk returns a tuple with the Identifiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifiers

`func (o *AdvisoryGitlabAdvisory) SetIdentifiers(v []string)`

SetIdentifiers sets Identifiers field to given value.

### HasIdentifiers

`func (o *AdvisoryGitlabAdvisory) HasIdentifiers() bool`

HasIdentifiers returns a boolean if a field has been set.

### GetNotImpacted

`func (o *AdvisoryGitlabAdvisory) GetNotImpacted() string`

GetNotImpacted returns the NotImpacted field if non-nil, zero value otherwise.

### GetNotImpactedOk

`func (o *AdvisoryGitlabAdvisory) GetNotImpactedOk() (*string, bool)`

GetNotImpactedOk returns a tuple with the NotImpacted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotImpacted

`func (o *AdvisoryGitlabAdvisory) SetNotImpacted(v string)`

SetNotImpacted sets NotImpacted field to given value.

### HasNotImpacted

`func (o *AdvisoryGitlabAdvisory) HasNotImpacted() bool`

HasNotImpacted returns a boolean if a field has been set.

### GetPackageManager

`func (o *AdvisoryGitlabAdvisory) GetPackageManager() string`

GetPackageManager returns the PackageManager field if non-nil, zero value otherwise.

### GetPackageManagerOk

`func (o *AdvisoryGitlabAdvisory) GetPackageManagerOk() (*string, bool)`

GetPackageManagerOk returns a tuple with the PackageManager field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageManager

`func (o *AdvisoryGitlabAdvisory) SetPackageManager(v string)`

SetPackageManager sets PackageManager field to given value.

### HasPackageManager

`func (o *AdvisoryGitlabAdvisory) HasPackageManager() bool`

HasPackageManager returns a boolean if a field has been set.

### GetPackageName

`func (o *AdvisoryGitlabAdvisory) GetPackageName() string`

GetPackageName returns the PackageName field if non-nil, zero value otherwise.

### GetPackageNameOk

`func (o *AdvisoryGitlabAdvisory) GetPackageNameOk() (*string, bool)`

GetPackageNameOk returns a tuple with the PackageName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageName

`func (o *AdvisoryGitlabAdvisory) SetPackageName(v string)`

SetPackageName sets PackageName field to given value.

### HasPackageName

`func (o *AdvisoryGitlabAdvisory) HasPackageName() bool`

HasPackageName returns a boolean if a field has been set.

### GetPackageSlug

`func (o *AdvisoryGitlabAdvisory) GetPackageSlug() string`

GetPackageSlug returns the PackageSlug field if non-nil, zero value otherwise.

### GetPackageSlugOk

`func (o *AdvisoryGitlabAdvisory) GetPackageSlugOk() (*string, bool)`

GetPackageSlugOk returns a tuple with the PackageSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageSlug

`func (o *AdvisoryGitlabAdvisory) SetPackageSlug(v string)`

SetPackageSlug sets PackageSlug field to given value.

### HasPackageSlug

`func (o *AdvisoryGitlabAdvisory) HasPackageSlug() bool`

HasPackageSlug returns a boolean if a field has been set.

### GetPubdate

`func (o *AdvisoryGitlabAdvisory) GetPubdate() string`

GetPubdate returns the Pubdate field if non-nil, zero value otherwise.

### GetPubdateOk

`func (o *AdvisoryGitlabAdvisory) GetPubdateOk() (*string, bool)`

GetPubdateOk returns a tuple with the Pubdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPubdate

`func (o *AdvisoryGitlabAdvisory) SetPubdate(v string)`

SetPubdate sets Pubdate field to given value.

### HasPubdate

`func (o *AdvisoryGitlabAdvisory) HasPubdate() bool`

HasPubdate returns a boolean if a field has been set.

### GetSolution

`func (o *AdvisoryGitlabAdvisory) GetSolution() string`

GetSolution returns the Solution field if non-nil, zero value otherwise.

### GetSolutionOk

`func (o *AdvisoryGitlabAdvisory) GetSolutionOk() (*string, bool)`

GetSolutionOk returns a tuple with the Solution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSolution

`func (o *AdvisoryGitlabAdvisory) SetSolution(v string)`

SetSolution sets Solution field to given value.

### HasSolution

`func (o *AdvisoryGitlabAdvisory) HasSolution() bool`

HasSolution returns a boolean if a field has been set.

### GetTitle

`func (o *AdvisoryGitlabAdvisory) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *AdvisoryGitlabAdvisory) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *AdvisoryGitlabAdvisory) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *AdvisoryGitlabAdvisory) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetUrls

`func (o *AdvisoryGitlabAdvisory) GetUrls() []string`

GetUrls returns the Urls field if non-nil, zero value otherwise.

### GetUrlsOk

`func (o *AdvisoryGitlabAdvisory) GetUrlsOk() (*[]string, bool)`

GetUrlsOk returns a tuple with the Urls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrls

`func (o *AdvisoryGitlabAdvisory) SetUrls(v []string)`

SetUrls sets Urls field to given value.

### HasUrls

`func (o *AdvisoryGitlabAdvisory) HasUrls() bool`

HasUrls returns a boolean if a field has been set.

### GetUuid

`func (o *AdvisoryGitlabAdvisory) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *AdvisoryGitlabAdvisory) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *AdvisoryGitlabAdvisory) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *AdvisoryGitlabAdvisory) HasUuid() bool`

HasUuid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


