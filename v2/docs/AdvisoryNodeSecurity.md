# AdvisoryNodeSecurity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AffectedEnvironments** | Pointer to **[]string** |  | [optional] 
**Author** | Pointer to [**AdvisoryNodeAuthor**](AdvisoryNodeAuthor.md) |  | [optional] 
**CoordinatingVendor** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**Cve** | Pointer to **[]string** |  | [optional] 
**CvssScore** | Pointer to **float32** |  | [optional] 
**CvssVector** | Pointer to **string** |  | [optional] 
**DateAdded** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **int32** |  | [optional] 
**ModuleName** | Pointer to **string** |  | [optional] 
**Overview** | Pointer to **string** |  | [optional] 
**PatchedVersions** | Pointer to **string** |  | [optional] 
**PublishDate** | Pointer to **string** |  | [optional] 
**Recommendation** | Pointer to **string** |  | [optional] 
**References** | Pointer to **[]string** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 
**VulnerableVersions** | Pointer to **string** |  | [optional] 

## Methods

### NewAdvisoryNodeSecurity

`func NewAdvisoryNodeSecurity() *AdvisoryNodeSecurity`

NewAdvisoryNodeSecurity instantiates a new AdvisoryNodeSecurity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdvisoryNodeSecurityWithDefaults

`func NewAdvisoryNodeSecurityWithDefaults() *AdvisoryNodeSecurity`

NewAdvisoryNodeSecurityWithDefaults instantiates a new AdvisoryNodeSecurity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAffectedEnvironments

`func (o *AdvisoryNodeSecurity) GetAffectedEnvironments() []string`

GetAffectedEnvironments returns the AffectedEnvironments field if non-nil, zero value otherwise.

### GetAffectedEnvironmentsOk

`func (o *AdvisoryNodeSecurity) GetAffectedEnvironmentsOk() (*[]string, bool)`

GetAffectedEnvironmentsOk returns a tuple with the AffectedEnvironments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffectedEnvironments

`func (o *AdvisoryNodeSecurity) SetAffectedEnvironments(v []string)`

SetAffectedEnvironments sets AffectedEnvironments field to given value.

### HasAffectedEnvironments

`func (o *AdvisoryNodeSecurity) HasAffectedEnvironments() bool`

HasAffectedEnvironments returns a boolean if a field has been set.

### GetAuthor

`func (o *AdvisoryNodeSecurity) GetAuthor() AdvisoryNodeAuthor`

GetAuthor returns the Author field if non-nil, zero value otherwise.

### GetAuthorOk

`func (o *AdvisoryNodeSecurity) GetAuthorOk() (*AdvisoryNodeAuthor, bool)`

GetAuthorOk returns a tuple with the Author field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthor

`func (o *AdvisoryNodeSecurity) SetAuthor(v AdvisoryNodeAuthor)`

SetAuthor sets Author field to given value.

### HasAuthor

`func (o *AdvisoryNodeSecurity) HasAuthor() bool`

HasAuthor returns a boolean if a field has been set.

### GetCoordinatingVendor

`func (o *AdvisoryNodeSecurity) GetCoordinatingVendor() string`

GetCoordinatingVendor returns the CoordinatingVendor field if non-nil, zero value otherwise.

### GetCoordinatingVendorOk

`func (o *AdvisoryNodeSecurity) GetCoordinatingVendorOk() (*string, bool)`

GetCoordinatingVendorOk returns a tuple with the CoordinatingVendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoordinatingVendor

`func (o *AdvisoryNodeSecurity) SetCoordinatingVendor(v string)`

SetCoordinatingVendor sets CoordinatingVendor field to given value.

### HasCoordinatingVendor

`func (o *AdvisoryNodeSecurity) HasCoordinatingVendor() bool`

HasCoordinatingVendor returns a boolean if a field has been set.

### GetCreatedAt

`func (o *AdvisoryNodeSecurity) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AdvisoryNodeSecurity) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AdvisoryNodeSecurity) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *AdvisoryNodeSecurity) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCve

`func (o *AdvisoryNodeSecurity) GetCve() []string`

GetCve returns the Cve field if non-nil, zero value otherwise.

### GetCveOk

`func (o *AdvisoryNodeSecurity) GetCveOk() (*[]string, bool)`

GetCveOk returns a tuple with the Cve field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCve

`func (o *AdvisoryNodeSecurity) SetCve(v []string)`

SetCve sets Cve field to given value.

### HasCve

`func (o *AdvisoryNodeSecurity) HasCve() bool`

HasCve returns a boolean if a field has been set.

### GetCvssScore

`func (o *AdvisoryNodeSecurity) GetCvssScore() float32`

GetCvssScore returns the CvssScore field if non-nil, zero value otherwise.

### GetCvssScoreOk

`func (o *AdvisoryNodeSecurity) GetCvssScoreOk() (*float32, bool)`

GetCvssScoreOk returns a tuple with the CvssScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCvssScore

`func (o *AdvisoryNodeSecurity) SetCvssScore(v float32)`

SetCvssScore sets CvssScore field to given value.

### HasCvssScore

`func (o *AdvisoryNodeSecurity) HasCvssScore() bool`

HasCvssScore returns a boolean if a field has been set.

### GetCvssVector

`func (o *AdvisoryNodeSecurity) GetCvssVector() string`

GetCvssVector returns the CvssVector field if non-nil, zero value otherwise.

### GetCvssVectorOk

`func (o *AdvisoryNodeSecurity) GetCvssVectorOk() (*string, bool)`

GetCvssVectorOk returns a tuple with the CvssVector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCvssVector

`func (o *AdvisoryNodeSecurity) SetCvssVector(v string)`

SetCvssVector sets CvssVector field to given value.

### HasCvssVector

`func (o *AdvisoryNodeSecurity) HasCvssVector() bool`

HasCvssVector returns a boolean if a field has been set.

### GetDateAdded

`func (o *AdvisoryNodeSecurity) GetDateAdded() string`

GetDateAdded returns the DateAdded field if non-nil, zero value otherwise.

### GetDateAddedOk

`func (o *AdvisoryNodeSecurity) GetDateAddedOk() (*string, bool)`

GetDateAddedOk returns a tuple with the DateAdded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateAdded

`func (o *AdvisoryNodeSecurity) SetDateAdded(v string)`

SetDateAdded sets DateAdded field to given value.

### HasDateAdded

`func (o *AdvisoryNodeSecurity) HasDateAdded() bool`

HasDateAdded returns a boolean if a field has been set.

### GetId

`func (o *AdvisoryNodeSecurity) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AdvisoryNodeSecurity) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AdvisoryNodeSecurity) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *AdvisoryNodeSecurity) HasId() bool`

HasId returns a boolean if a field has been set.

### GetModuleName

`func (o *AdvisoryNodeSecurity) GetModuleName() string`

GetModuleName returns the ModuleName field if non-nil, zero value otherwise.

### GetModuleNameOk

`func (o *AdvisoryNodeSecurity) GetModuleNameOk() (*string, bool)`

GetModuleNameOk returns a tuple with the ModuleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModuleName

`func (o *AdvisoryNodeSecurity) SetModuleName(v string)`

SetModuleName sets ModuleName field to given value.

### HasModuleName

`func (o *AdvisoryNodeSecurity) HasModuleName() bool`

HasModuleName returns a boolean if a field has been set.

### GetOverview

`func (o *AdvisoryNodeSecurity) GetOverview() string`

GetOverview returns the Overview field if non-nil, zero value otherwise.

### GetOverviewOk

`func (o *AdvisoryNodeSecurity) GetOverviewOk() (*string, bool)`

GetOverviewOk returns a tuple with the Overview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverview

`func (o *AdvisoryNodeSecurity) SetOverview(v string)`

SetOverview sets Overview field to given value.

### HasOverview

`func (o *AdvisoryNodeSecurity) HasOverview() bool`

HasOverview returns a boolean if a field has been set.

### GetPatchedVersions

`func (o *AdvisoryNodeSecurity) GetPatchedVersions() string`

GetPatchedVersions returns the PatchedVersions field if non-nil, zero value otherwise.

### GetPatchedVersionsOk

`func (o *AdvisoryNodeSecurity) GetPatchedVersionsOk() (*string, bool)`

GetPatchedVersionsOk returns a tuple with the PatchedVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatchedVersions

`func (o *AdvisoryNodeSecurity) SetPatchedVersions(v string)`

SetPatchedVersions sets PatchedVersions field to given value.

### HasPatchedVersions

`func (o *AdvisoryNodeSecurity) HasPatchedVersions() bool`

HasPatchedVersions returns a boolean if a field has been set.

### GetPublishDate

`func (o *AdvisoryNodeSecurity) GetPublishDate() string`

GetPublishDate returns the PublishDate field if non-nil, zero value otherwise.

### GetPublishDateOk

`func (o *AdvisoryNodeSecurity) GetPublishDateOk() (*string, bool)`

GetPublishDateOk returns a tuple with the PublishDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishDate

`func (o *AdvisoryNodeSecurity) SetPublishDate(v string)`

SetPublishDate sets PublishDate field to given value.

### HasPublishDate

`func (o *AdvisoryNodeSecurity) HasPublishDate() bool`

HasPublishDate returns a boolean if a field has been set.

### GetRecommendation

`func (o *AdvisoryNodeSecurity) GetRecommendation() string`

GetRecommendation returns the Recommendation field if non-nil, zero value otherwise.

### GetRecommendationOk

`func (o *AdvisoryNodeSecurity) GetRecommendationOk() (*string, bool)`

GetRecommendationOk returns a tuple with the Recommendation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecommendation

`func (o *AdvisoryNodeSecurity) SetRecommendation(v string)`

SetRecommendation sets Recommendation field to given value.

### HasRecommendation

`func (o *AdvisoryNodeSecurity) HasRecommendation() bool`

HasRecommendation returns a boolean if a field has been set.

### GetReferences

`func (o *AdvisoryNodeSecurity) GetReferences() []string`

GetReferences returns the References field if non-nil, zero value otherwise.

### GetReferencesOk

`func (o *AdvisoryNodeSecurity) GetReferencesOk() (*[]string, bool)`

GetReferencesOk returns a tuple with the References field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferences

`func (o *AdvisoryNodeSecurity) SetReferences(v []string)`

SetReferences sets References field to given value.

### HasReferences

`func (o *AdvisoryNodeSecurity) HasReferences() bool`

HasReferences returns a boolean if a field has been set.

### GetTitle

`func (o *AdvisoryNodeSecurity) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *AdvisoryNodeSecurity) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *AdvisoryNodeSecurity) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *AdvisoryNodeSecurity) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *AdvisoryNodeSecurity) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *AdvisoryNodeSecurity) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *AdvisoryNodeSecurity) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *AdvisoryNodeSecurity) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUrl

`func (o *AdvisoryNodeSecurity) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *AdvisoryNodeSecurity) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *AdvisoryNodeSecurity) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *AdvisoryNodeSecurity) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetVulnerableVersions

`func (o *AdvisoryNodeSecurity) GetVulnerableVersions() string`

GetVulnerableVersions returns the VulnerableVersions field if non-nil, zero value otherwise.

### GetVulnerableVersionsOk

`func (o *AdvisoryNodeSecurity) GetVulnerableVersionsOk() (*string, bool)`

GetVulnerableVersionsOk returns a tuple with the VulnerableVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVulnerableVersions

`func (o *AdvisoryNodeSecurity) SetVulnerableVersions(v string)`

SetVulnerableVersions sets VulnerableVersions field to given value.

### HasVulnerableVersions

`func (o *AdvisoryNodeSecurity) HasVulnerableVersions() bool`

HasVulnerableVersions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


