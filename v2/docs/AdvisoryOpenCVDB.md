# AdvisoryOpenCVDB

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AffectedPlatforms** | Pointer to **[]string** |  | [optional] 
**AffectedServices** | Pointer to **[]string** |  | [optional] 
**Cve** | Pointer to **[]string** |  | [optional] 
**DateAdded** | Pointer to **string** |  | [optional] 
**DisclosedAt** | Pointer to **string** |  | [optional] 
**KnownItwExploitation** | Pointer to **bool** |  | [optional] 
**ManualRemediation** | Pointer to **string** |  | [optional] 
**PublishedAt** | Pointer to **string** |  | [optional] 
**References** | Pointer to **[]string** |  | [optional] 
**Summary** | Pointer to **string** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 

## Methods

### NewAdvisoryOpenCVDB

`func NewAdvisoryOpenCVDB() *AdvisoryOpenCVDB`

NewAdvisoryOpenCVDB instantiates a new AdvisoryOpenCVDB object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdvisoryOpenCVDBWithDefaults

`func NewAdvisoryOpenCVDBWithDefaults() *AdvisoryOpenCVDB`

NewAdvisoryOpenCVDBWithDefaults instantiates a new AdvisoryOpenCVDB object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAffectedPlatforms

`func (o *AdvisoryOpenCVDB) GetAffectedPlatforms() []string`

GetAffectedPlatforms returns the AffectedPlatforms field if non-nil, zero value otherwise.

### GetAffectedPlatformsOk

`func (o *AdvisoryOpenCVDB) GetAffectedPlatformsOk() (*[]string, bool)`

GetAffectedPlatformsOk returns a tuple with the AffectedPlatforms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffectedPlatforms

`func (o *AdvisoryOpenCVDB) SetAffectedPlatforms(v []string)`

SetAffectedPlatforms sets AffectedPlatforms field to given value.

### HasAffectedPlatforms

`func (o *AdvisoryOpenCVDB) HasAffectedPlatforms() bool`

HasAffectedPlatforms returns a boolean if a field has been set.

### GetAffectedServices

`func (o *AdvisoryOpenCVDB) GetAffectedServices() []string`

GetAffectedServices returns the AffectedServices field if non-nil, zero value otherwise.

### GetAffectedServicesOk

`func (o *AdvisoryOpenCVDB) GetAffectedServicesOk() (*[]string, bool)`

GetAffectedServicesOk returns a tuple with the AffectedServices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffectedServices

`func (o *AdvisoryOpenCVDB) SetAffectedServices(v []string)`

SetAffectedServices sets AffectedServices field to given value.

### HasAffectedServices

`func (o *AdvisoryOpenCVDB) HasAffectedServices() bool`

HasAffectedServices returns a boolean if a field has been set.

### GetCve

`func (o *AdvisoryOpenCVDB) GetCve() []string`

GetCve returns the Cve field if non-nil, zero value otherwise.

### GetCveOk

`func (o *AdvisoryOpenCVDB) GetCveOk() (*[]string, bool)`

GetCveOk returns a tuple with the Cve field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCve

`func (o *AdvisoryOpenCVDB) SetCve(v []string)`

SetCve sets Cve field to given value.

### HasCve

`func (o *AdvisoryOpenCVDB) HasCve() bool`

HasCve returns a boolean if a field has been set.

### GetDateAdded

`func (o *AdvisoryOpenCVDB) GetDateAdded() string`

GetDateAdded returns the DateAdded field if non-nil, zero value otherwise.

### GetDateAddedOk

`func (o *AdvisoryOpenCVDB) GetDateAddedOk() (*string, bool)`

GetDateAddedOk returns a tuple with the DateAdded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateAdded

`func (o *AdvisoryOpenCVDB) SetDateAdded(v string)`

SetDateAdded sets DateAdded field to given value.

### HasDateAdded

`func (o *AdvisoryOpenCVDB) HasDateAdded() bool`

HasDateAdded returns a boolean if a field has been set.

### GetDisclosedAt

`func (o *AdvisoryOpenCVDB) GetDisclosedAt() string`

GetDisclosedAt returns the DisclosedAt field if non-nil, zero value otherwise.

### GetDisclosedAtOk

`func (o *AdvisoryOpenCVDB) GetDisclosedAtOk() (*string, bool)`

GetDisclosedAtOk returns a tuple with the DisclosedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisclosedAt

`func (o *AdvisoryOpenCVDB) SetDisclosedAt(v string)`

SetDisclosedAt sets DisclosedAt field to given value.

### HasDisclosedAt

`func (o *AdvisoryOpenCVDB) HasDisclosedAt() bool`

HasDisclosedAt returns a boolean if a field has been set.

### GetKnownItwExploitation

`func (o *AdvisoryOpenCVDB) GetKnownItwExploitation() bool`

GetKnownItwExploitation returns the KnownItwExploitation field if non-nil, zero value otherwise.

### GetKnownItwExploitationOk

`func (o *AdvisoryOpenCVDB) GetKnownItwExploitationOk() (*bool, bool)`

GetKnownItwExploitationOk returns a tuple with the KnownItwExploitation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKnownItwExploitation

`func (o *AdvisoryOpenCVDB) SetKnownItwExploitation(v bool)`

SetKnownItwExploitation sets KnownItwExploitation field to given value.

### HasKnownItwExploitation

`func (o *AdvisoryOpenCVDB) HasKnownItwExploitation() bool`

HasKnownItwExploitation returns a boolean if a field has been set.

### GetManualRemediation

`func (o *AdvisoryOpenCVDB) GetManualRemediation() string`

GetManualRemediation returns the ManualRemediation field if non-nil, zero value otherwise.

### GetManualRemediationOk

`func (o *AdvisoryOpenCVDB) GetManualRemediationOk() (*string, bool)`

GetManualRemediationOk returns a tuple with the ManualRemediation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManualRemediation

`func (o *AdvisoryOpenCVDB) SetManualRemediation(v string)`

SetManualRemediation sets ManualRemediation field to given value.

### HasManualRemediation

`func (o *AdvisoryOpenCVDB) HasManualRemediation() bool`

HasManualRemediation returns a boolean if a field has been set.

### GetPublishedAt

`func (o *AdvisoryOpenCVDB) GetPublishedAt() string`

GetPublishedAt returns the PublishedAt field if non-nil, zero value otherwise.

### GetPublishedAtOk

`func (o *AdvisoryOpenCVDB) GetPublishedAtOk() (*string, bool)`

GetPublishedAtOk returns a tuple with the PublishedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishedAt

`func (o *AdvisoryOpenCVDB) SetPublishedAt(v string)`

SetPublishedAt sets PublishedAt field to given value.

### HasPublishedAt

`func (o *AdvisoryOpenCVDB) HasPublishedAt() bool`

HasPublishedAt returns a boolean if a field has been set.

### GetReferences

`func (o *AdvisoryOpenCVDB) GetReferences() []string`

GetReferences returns the References field if non-nil, zero value otherwise.

### GetReferencesOk

`func (o *AdvisoryOpenCVDB) GetReferencesOk() (*[]string, bool)`

GetReferencesOk returns a tuple with the References field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferences

`func (o *AdvisoryOpenCVDB) SetReferences(v []string)`

SetReferences sets References field to given value.

### HasReferences

`func (o *AdvisoryOpenCVDB) HasReferences() bool`

HasReferences returns a boolean if a field has been set.

### GetSummary

`func (o *AdvisoryOpenCVDB) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *AdvisoryOpenCVDB) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *AdvisoryOpenCVDB) SetSummary(v string)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *AdvisoryOpenCVDB) HasSummary() bool`

HasSummary returns a boolean if a field has been set.

### GetTitle

`func (o *AdvisoryOpenCVDB) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *AdvisoryOpenCVDB) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *AdvisoryOpenCVDB) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *AdvisoryOpenCVDB) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetUrl

`func (o *AdvisoryOpenCVDB) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *AdvisoryOpenCVDB) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *AdvisoryOpenCVDB) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *AdvisoryOpenCVDB) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


