# AdvisoryAlmaLinuxUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BsRepoId** | Pointer to [**AdvisoryAlmaObjectID**](AdvisoryAlmaObjectID.md) |  | [optional] 
**Cve** | Pointer to **[]string** |  | [optional] 
**DateAdded** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Fromstr** | Pointer to **string** |  | [optional] 
**Id** | Pointer to [**AdvisoryAlmaObjectID**](AdvisoryAlmaObjectID.md) |  | [optional] 
**IssuedDate** | Pointer to [**AdvisoryAlmaDate**](AdvisoryAlmaDate.md) |  | [optional] 
**Pkglist** | Pointer to [**AdvisoryAlmaPackageList**](AdvisoryAlmaPackageList.md) |  | [optional] 
**Pushcount** | Pointer to **string** |  | [optional] 
**References** | Pointer to [**[]AdvisoryAlmaReference**](AdvisoryAlmaReference.md) |  | [optional] 
**Release** | Pointer to **string** |  | [optional] 
**Rights** | Pointer to **string** |  | [optional] 
**Severity** | Pointer to **string** |  | [optional] 
**Solution** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Summary** | Pointer to **string** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**UpdateUrl** | Pointer to **string** |  | [optional] 
**UpdatedDate** | Pointer to [**AdvisoryAlmaDate**](AdvisoryAlmaDate.md) |  | [optional] 
**UpdateinfoId** | Pointer to **string** |  | [optional] 
**Version** | Pointer to **string** |  | [optional] 

## Methods

### NewAdvisoryAlmaLinuxUpdate

`func NewAdvisoryAlmaLinuxUpdate() *AdvisoryAlmaLinuxUpdate`

NewAdvisoryAlmaLinuxUpdate instantiates a new AdvisoryAlmaLinuxUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdvisoryAlmaLinuxUpdateWithDefaults

`func NewAdvisoryAlmaLinuxUpdateWithDefaults() *AdvisoryAlmaLinuxUpdate`

NewAdvisoryAlmaLinuxUpdateWithDefaults instantiates a new AdvisoryAlmaLinuxUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBsRepoId

`func (o *AdvisoryAlmaLinuxUpdate) GetBsRepoId() AdvisoryAlmaObjectID`

GetBsRepoId returns the BsRepoId field if non-nil, zero value otherwise.

### GetBsRepoIdOk

`func (o *AdvisoryAlmaLinuxUpdate) GetBsRepoIdOk() (*AdvisoryAlmaObjectID, bool)`

GetBsRepoIdOk returns a tuple with the BsRepoId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBsRepoId

`func (o *AdvisoryAlmaLinuxUpdate) SetBsRepoId(v AdvisoryAlmaObjectID)`

SetBsRepoId sets BsRepoId field to given value.

### HasBsRepoId

`func (o *AdvisoryAlmaLinuxUpdate) HasBsRepoId() bool`

HasBsRepoId returns a boolean if a field has been set.

### GetCve

`func (o *AdvisoryAlmaLinuxUpdate) GetCve() []string`

GetCve returns the Cve field if non-nil, zero value otherwise.

### GetCveOk

`func (o *AdvisoryAlmaLinuxUpdate) GetCveOk() (*[]string, bool)`

GetCveOk returns a tuple with the Cve field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCve

`func (o *AdvisoryAlmaLinuxUpdate) SetCve(v []string)`

SetCve sets Cve field to given value.

### HasCve

`func (o *AdvisoryAlmaLinuxUpdate) HasCve() bool`

HasCve returns a boolean if a field has been set.

### GetDateAdded

`func (o *AdvisoryAlmaLinuxUpdate) GetDateAdded() string`

GetDateAdded returns the DateAdded field if non-nil, zero value otherwise.

### GetDateAddedOk

`func (o *AdvisoryAlmaLinuxUpdate) GetDateAddedOk() (*string, bool)`

GetDateAddedOk returns a tuple with the DateAdded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateAdded

`func (o *AdvisoryAlmaLinuxUpdate) SetDateAdded(v string)`

SetDateAdded sets DateAdded field to given value.

### HasDateAdded

`func (o *AdvisoryAlmaLinuxUpdate) HasDateAdded() bool`

HasDateAdded returns a boolean if a field has been set.

### GetDescription

`func (o *AdvisoryAlmaLinuxUpdate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AdvisoryAlmaLinuxUpdate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AdvisoryAlmaLinuxUpdate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AdvisoryAlmaLinuxUpdate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFromstr

`func (o *AdvisoryAlmaLinuxUpdate) GetFromstr() string`

GetFromstr returns the Fromstr field if non-nil, zero value otherwise.

### GetFromstrOk

`func (o *AdvisoryAlmaLinuxUpdate) GetFromstrOk() (*string, bool)`

GetFromstrOk returns a tuple with the Fromstr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromstr

`func (o *AdvisoryAlmaLinuxUpdate) SetFromstr(v string)`

SetFromstr sets Fromstr field to given value.

### HasFromstr

`func (o *AdvisoryAlmaLinuxUpdate) HasFromstr() bool`

HasFromstr returns a boolean if a field has been set.

### GetId

`func (o *AdvisoryAlmaLinuxUpdate) GetId() AdvisoryAlmaObjectID`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AdvisoryAlmaLinuxUpdate) GetIdOk() (*AdvisoryAlmaObjectID, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AdvisoryAlmaLinuxUpdate) SetId(v AdvisoryAlmaObjectID)`

SetId sets Id field to given value.

### HasId

`func (o *AdvisoryAlmaLinuxUpdate) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIssuedDate

`func (o *AdvisoryAlmaLinuxUpdate) GetIssuedDate() AdvisoryAlmaDate`

GetIssuedDate returns the IssuedDate field if non-nil, zero value otherwise.

### GetIssuedDateOk

`func (o *AdvisoryAlmaLinuxUpdate) GetIssuedDateOk() (*AdvisoryAlmaDate, bool)`

GetIssuedDateOk returns a tuple with the IssuedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuedDate

`func (o *AdvisoryAlmaLinuxUpdate) SetIssuedDate(v AdvisoryAlmaDate)`

SetIssuedDate sets IssuedDate field to given value.

### HasIssuedDate

`func (o *AdvisoryAlmaLinuxUpdate) HasIssuedDate() bool`

HasIssuedDate returns a boolean if a field has been set.

### GetPkglist

`func (o *AdvisoryAlmaLinuxUpdate) GetPkglist() AdvisoryAlmaPackageList`

GetPkglist returns the Pkglist field if non-nil, zero value otherwise.

### GetPkglistOk

`func (o *AdvisoryAlmaLinuxUpdate) GetPkglistOk() (*AdvisoryAlmaPackageList, bool)`

GetPkglistOk returns a tuple with the Pkglist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkglist

`func (o *AdvisoryAlmaLinuxUpdate) SetPkglist(v AdvisoryAlmaPackageList)`

SetPkglist sets Pkglist field to given value.

### HasPkglist

`func (o *AdvisoryAlmaLinuxUpdate) HasPkglist() bool`

HasPkglist returns a boolean if a field has been set.

### GetPushcount

`func (o *AdvisoryAlmaLinuxUpdate) GetPushcount() string`

GetPushcount returns the Pushcount field if non-nil, zero value otherwise.

### GetPushcountOk

`func (o *AdvisoryAlmaLinuxUpdate) GetPushcountOk() (*string, bool)`

GetPushcountOk returns a tuple with the Pushcount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPushcount

`func (o *AdvisoryAlmaLinuxUpdate) SetPushcount(v string)`

SetPushcount sets Pushcount field to given value.

### HasPushcount

`func (o *AdvisoryAlmaLinuxUpdate) HasPushcount() bool`

HasPushcount returns a boolean if a field has been set.

### GetReferences

`func (o *AdvisoryAlmaLinuxUpdate) GetReferences() []AdvisoryAlmaReference`

GetReferences returns the References field if non-nil, zero value otherwise.

### GetReferencesOk

`func (o *AdvisoryAlmaLinuxUpdate) GetReferencesOk() (*[]AdvisoryAlmaReference, bool)`

GetReferencesOk returns a tuple with the References field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferences

`func (o *AdvisoryAlmaLinuxUpdate) SetReferences(v []AdvisoryAlmaReference)`

SetReferences sets References field to given value.

### HasReferences

`func (o *AdvisoryAlmaLinuxUpdate) HasReferences() bool`

HasReferences returns a boolean if a field has been set.

### GetRelease

`func (o *AdvisoryAlmaLinuxUpdate) GetRelease() string`

GetRelease returns the Release field if non-nil, zero value otherwise.

### GetReleaseOk

`func (o *AdvisoryAlmaLinuxUpdate) GetReleaseOk() (*string, bool)`

GetReleaseOk returns a tuple with the Release field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelease

`func (o *AdvisoryAlmaLinuxUpdate) SetRelease(v string)`

SetRelease sets Release field to given value.

### HasRelease

`func (o *AdvisoryAlmaLinuxUpdate) HasRelease() bool`

HasRelease returns a boolean if a field has been set.

### GetRights

`func (o *AdvisoryAlmaLinuxUpdate) GetRights() string`

GetRights returns the Rights field if non-nil, zero value otherwise.

### GetRightsOk

`func (o *AdvisoryAlmaLinuxUpdate) GetRightsOk() (*string, bool)`

GetRightsOk returns a tuple with the Rights field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRights

`func (o *AdvisoryAlmaLinuxUpdate) SetRights(v string)`

SetRights sets Rights field to given value.

### HasRights

`func (o *AdvisoryAlmaLinuxUpdate) HasRights() bool`

HasRights returns a boolean if a field has been set.

### GetSeverity

`func (o *AdvisoryAlmaLinuxUpdate) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *AdvisoryAlmaLinuxUpdate) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *AdvisoryAlmaLinuxUpdate) SetSeverity(v string)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *AdvisoryAlmaLinuxUpdate) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetSolution

`func (o *AdvisoryAlmaLinuxUpdate) GetSolution() string`

GetSolution returns the Solution field if non-nil, zero value otherwise.

### GetSolutionOk

`func (o *AdvisoryAlmaLinuxUpdate) GetSolutionOk() (*string, bool)`

GetSolutionOk returns a tuple with the Solution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSolution

`func (o *AdvisoryAlmaLinuxUpdate) SetSolution(v string)`

SetSolution sets Solution field to given value.

### HasSolution

`func (o *AdvisoryAlmaLinuxUpdate) HasSolution() bool`

HasSolution returns a boolean if a field has been set.

### GetStatus

`func (o *AdvisoryAlmaLinuxUpdate) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AdvisoryAlmaLinuxUpdate) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AdvisoryAlmaLinuxUpdate) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AdvisoryAlmaLinuxUpdate) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSummary

`func (o *AdvisoryAlmaLinuxUpdate) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *AdvisoryAlmaLinuxUpdate) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *AdvisoryAlmaLinuxUpdate) SetSummary(v string)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *AdvisoryAlmaLinuxUpdate) HasSummary() bool`

HasSummary returns a boolean if a field has been set.

### GetTitle

`func (o *AdvisoryAlmaLinuxUpdate) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *AdvisoryAlmaLinuxUpdate) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *AdvisoryAlmaLinuxUpdate) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *AdvisoryAlmaLinuxUpdate) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetType

`func (o *AdvisoryAlmaLinuxUpdate) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AdvisoryAlmaLinuxUpdate) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AdvisoryAlmaLinuxUpdate) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *AdvisoryAlmaLinuxUpdate) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUpdateUrl

`func (o *AdvisoryAlmaLinuxUpdate) GetUpdateUrl() string`

GetUpdateUrl returns the UpdateUrl field if non-nil, zero value otherwise.

### GetUpdateUrlOk

`func (o *AdvisoryAlmaLinuxUpdate) GetUpdateUrlOk() (*string, bool)`

GetUpdateUrlOk returns a tuple with the UpdateUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateUrl

`func (o *AdvisoryAlmaLinuxUpdate) SetUpdateUrl(v string)`

SetUpdateUrl sets UpdateUrl field to given value.

### HasUpdateUrl

`func (o *AdvisoryAlmaLinuxUpdate) HasUpdateUrl() bool`

HasUpdateUrl returns a boolean if a field has been set.

### GetUpdatedDate

`func (o *AdvisoryAlmaLinuxUpdate) GetUpdatedDate() AdvisoryAlmaDate`

GetUpdatedDate returns the UpdatedDate field if non-nil, zero value otherwise.

### GetUpdatedDateOk

`func (o *AdvisoryAlmaLinuxUpdate) GetUpdatedDateOk() (*AdvisoryAlmaDate, bool)`

GetUpdatedDateOk returns a tuple with the UpdatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedDate

`func (o *AdvisoryAlmaLinuxUpdate) SetUpdatedDate(v AdvisoryAlmaDate)`

SetUpdatedDate sets UpdatedDate field to given value.

### HasUpdatedDate

`func (o *AdvisoryAlmaLinuxUpdate) HasUpdatedDate() bool`

HasUpdatedDate returns a boolean if a field has been set.

### GetUpdateinfoId

`func (o *AdvisoryAlmaLinuxUpdate) GetUpdateinfoId() string`

GetUpdateinfoId returns the UpdateinfoId field if non-nil, zero value otherwise.

### GetUpdateinfoIdOk

`func (o *AdvisoryAlmaLinuxUpdate) GetUpdateinfoIdOk() (*string, bool)`

GetUpdateinfoIdOk returns a tuple with the UpdateinfoId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateinfoId

`func (o *AdvisoryAlmaLinuxUpdate) SetUpdateinfoId(v string)`

SetUpdateinfoId sets UpdateinfoId field to given value.

### HasUpdateinfoId

`func (o *AdvisoryAlmaLinuxUpdate) HasUpdateinfoId() bool`

HasUpdateinfoId returns a boolean if a field has been set.

### GetVersion

`func (o *AdvisoryAlmaLinuxUpdate) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *AdvisoryAlmaLinuxUpdate) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *AdvisoryAlmaLinuxUpdate) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *AdvisoryAlmaLinuxUpdate) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


