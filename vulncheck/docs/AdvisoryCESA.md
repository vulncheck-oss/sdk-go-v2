# AdvisoryCESA

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Arch** | Pointer to **[]string** |  | [optional] 
**Cve** | Pointer to **[]string** |  | [optional] 
**DateAdded** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**IssueDate** | Pointer to **string** |  | [optional] 
**OsRelease** | Pointer to **string** |  | [optional] 
**Packages** | Pointer to [**[]AdvisoryCentosPackage**](AdvisoryCentosPackage.md) |  | [optional] 
**References** | Pointer to **[]string** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 

## Methods

### NewAdvisoryCESA

`func NewAdvisoryCESA() *AdvisoryCESA`

NewAdvisoryCESA instantiates a new AdvisoryCESA object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdvisoryCESAWithDefaults

`func NewAdvisoryCESAWithDefaults() *AdvisoryCESA`

NewAdvisoryCESAWithDefaults instantiates a new AdvisoryCESA object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArch

`func (o *AdvisoryCESA) GetArch() []string`

GetArch returns the Arch field if non-nil, zero value otherwise.

### GetArchOk

`func (o *AdvisoryCESA) GetArchOk() (*[]string, bool)`

GetArchOk returns a tuple with the Arch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArch

`func (o *AdvisoryCESA) SetArch(v []string)`

SetArch sets Arch field to given value.

### HasArch

`func (o *AdvisoryCESA) HasArch() bool`

HasArch returns a boolean if a field has been set.

### GetCve

`func (o *AdvisoryCESA) GetCve() []string`

GetCve returns the Cve field if non-nil, zero value otherwise.

### GetCveOk

`func (o *AdvisoryCESA) GetCveOk() (*[]string, bool)`

GetCveOk returns a tuple with the Cve field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCve

`func (o *AdvisoryCESA) SetCve(v []string)`

SetCve sets Cve field to given value.

### HasCve

`func (o *AdvisoryCESA) HasCve() bool`

HasCve returns a boolean if a field has been set.

### GetDateAdded

`func (o *AdvisoryCESA) GetDateAdded() string`

GetDateAdded returns the DateAdded field if non-nil, zero value otherwise.

### GetDateAddedOk

`func (o *AdvisoryCESA) GetDateAddedOk() (*string, bool)`

GetDateAddedOk returns a tuple with the DateAdded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateAdded

`func (o *AdvisoryCESA) SetDateAdded(v string)`

SetDateAdded sets DateAdded field to given value.

### HasDateAdded

`func (o *AdvisoryCESA) HasDateAdded() bool`

HasDateAdded returns a boolean if a field has been set.

### GetId

`func (o *AdvisoryCESA) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AdvisoryCESA) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AdvisoryCESA) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AdvisoryCESA) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIssueDate

`func (o *AdvisoryCESA) GetIssueDate() string`

GetIssueDate returns the IssueDate field if non-nil, zero value otherwise.

### GetIssueDateOk

`func (o *AdvisoryCESA) GetIssueDateOk() (*string, bool)`

GetIssueDateOk returns a tuple with the IssueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueDate

`func (o *AdvisoryCESA) SetIssueDate(v string)`

SetIssueDate sets IssueDate field to given value.

### HasIssueDate

`func (o *AdvisoryCESA) HasIssueDate() bool`

HasIssueDate returns a boolean if a field has been set.

### GetOsRelease

`func (o *AdvisoryCESA) GetOsRelease() string`

GetOsRelease returns the OsRelease field if non-nil, zero value otherwise.

### GetOsReleaseOk

`func (o *AdvisoryCESA) GetOsReleaseOk() (*string, bool)`

GetOsReleaseOk returns a tuple with the OsRelease field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsRelease

`func (o *AdvisoryCESA) SetOsRelease(v string)`

SetOsRelease sets OsRelease field to given value.

### HasOsRelease

`func (o *AdvisoryCESA) HasOsRelease() bool`

HasOsRelease returns a boolean if a field has been set.

### GetPackages

`func (o *AdvisoryCESA) GetPackages() []AdvisoryCentosPackage`

GetPackages returns the Packages field if non-nil, zero value otherwise.

### GetPackagesOk

`func (o *AdvisoryCESA) GetPackagesOk() (*[]AdvisoryCentosPackage, bool)`

GetPackagesOk returns a tuple with the Packages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackages

`func (o *AdvisoryCESA) SetPackages(v []AdvisoryCentosPackage)`

SetPackages sets Packages field to given value.

### HasPackages

`func (o *AdvisoryCESA) HasPackages() bool`

HasPackages returns a boolean if a field has been set.

### GetReferences

`func (o *AdvisoryCESA) GetReferences() []string`

GetReferences returns the References field if non-nil, zero value otherwise.

### GetReferencesOk

`func (o *AdvisoryCESA) GetReferencesOk() (*[]string, bool)`

GetReferencesOk returns a tuple with the References field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferences

`func (o *AdvisoryCESA) SetReferences(v []string)`

SetReferences sets References field to given value.

### HasReferences

`func (o *AdvisoryCESA) HasReferences() bool`

HasReferences returns a boolean if a field has been set.

### GetTitle

`func (o *AdvisoryCESA) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *AdvisoryCESA) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *AdvisoryCESA) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *AdvisoryCESA) HasTitle() bool`

HasTitle returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


