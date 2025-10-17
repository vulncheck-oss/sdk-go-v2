# AdvisoryEOLReleaseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AlreadyEol** | Pointer to **bool** |  | [optional] 
**Branch** | Pointer to **string** | Alpine Linux | [optional] 
**BranchUrl** | Pointer to **string** | Alpine Linux | [optional] 
**Codename** | Pointer to **string** |  | [optional] 
**Cpe** | Pointer to **string** |  | [optional] 
**EolDate** | Pointer to **string** |  | [optional] 
**EolDateExtendedSupport** | Pointer to **string** | Oracle Linux, Solaris | [optional] 
**EolDatePremierSupport** | Pointer to **string** | Oracle Linux, Solaris | [optional] 
**EolEltsDate** | Pointer to **string** |  | [optional] 
**EolLtsDate** | Pointer to **string** |  | [optional] 
**GitBranch** | Pointer to **string** | Alpine Linux | [optional] 
**GitBranchUrl** | Pointer to **string** | Alpine Linux | [optional] 
**Lts** | Pointer to **bool** | Ubuntu | [optional] 
**MinorReleases** | Pointer to **[]string** | Alpine Linux | [optional] 
**Product** | Pointer to **string** |  | [optional] 
**ReleaseDate** | Pointer to **string** |  | [optional] 
**ReleaseName** | Pointer to **string** |  | [optional] 
**SourceUrl** | Pointer to **string** |  | [optional] 
**TechnologyLevel** | Pointer to **string** | AIX | [optional] 
**Vendor** | Pointer to **string** |  | [optional] 
**Version** | Pointer to **string** |  | [optional] 
**VersionApi** | Pointer to **string** | Android | [optional] 
**VersionDarwin** | Pointer to **string** | macOS | [optional] 
**VersionSunos** | Pointer to **string** | Solaris | [optional] 
**WindowsCurrentBuild** | Pointer to **string** | Microsoft Windows | [optional] 
**WindowsDisplayVersion** | Pointer to **string** | Microsoft Windows | [optional] 
**WindowsEditionId** | Pointer to **string** | Microsoft Windows | [optional] 
**WindowsInsiderPreview** | Pointer to **bool** | Microsoft Windows | [optional] 

## Methods

### NewAdvisoryEOLReleaseData

`func NewAdvisoryEOLReleaseData() *AdvisoryEOLReleaseData`

NewAdvisoryEOLReleaseData instantiates a new AdvisoryEOLReleaseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdvisoryEOLReleaseDataWithDefaults

`func NewAdvisoryEOLReleaseDataWithDefaults() *AdvisoryEOLReleaseData`

NewAdvisoryEOLReleaseDataWithDefaults instantiates a new AdvisoryEOLReleaseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlreadyEol

`func (o *AdvisoryEOLReleaseData) GetAlreadyEol() bool`

GetAlreadyEol returns the AlreadyEol field if non-nil, zero value otherwise.

### GetAlreadyEolOk

`func (o *AdvisoryEOLReleaseData) GetAlreadyEolOk() (*bool, bool)`

GetAlreadyEolOk returns a tuple with the AlreadyEol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlreadyEol

`func (o *AdvisoryEOLReleaseData) SetAlreadyEol(v bool)`

SetAlreadyEol sets AlreadyEol field to given value.

### HasAlreadyEol

`func (o *AdvisoryEOLReleaseData) HasAlreadyEol() bool`

HasAlreadyEol returns a boolean if a field has been set.

### GetBranch

`func (o *AdvisoryEOLReleaseData) GetBranch() string`

GetBranch returns the Branch field if non-nil, zero value otherwise.

### GetBranchOk

`func (o *AdvisoryEOLReleaseData) GetBranchOk() (*string, bool)`

GetBranchOk returns a tuple with the Branch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranch

`func (o *AdvisoryEOLReleaseData) SetBranch(v string)`

SetBranch sets Branch field to given value.

### HasBranch

`func (o *AdvisoryEOLReleaseData) HasBranch() bool`

HasBranch returns a boolean if a field has been set.

### GetBranchUrl

`func (o *AdvisoryEOLReleaseData) GetBranchUrl() string`

GetBranchUrl returns the BranchUrl field if non-nil, zero value otherwise.

### GetBranchUrlOk

`func (o *AdvisoryEOLReleaseData) GetBranchUrlOk() (*string, bool)`

GetBranchUrlOk returns a tuple with the BranchUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranchUrl

`func (o *AdvisoryEOLReleaseData) SetBranchUrl(v string)`

SetBranchUrl sets BranchUrl field to given value.

### HasBranchUrl

`func (o *AdvisoryEOLReleaseData) HasBranchUrl() bool`

HasBranchUrl returns a boolean if a field has been set.

### GetCodename

`func (o *AdvisoryEOLReleaseData) GetCodename() string`

GetCodename returns the Codename field if non-nil, zero value otherwise.

### GetCodenameOk

`func (o *AdvisoryEOLReleaseData) GetCodenameOk() (*string, bool)`

GetCodenameOk returns a tuple with the Codename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodename

`func (o *AdvisoryEOLReleaseData) SetCodename(v string)`

SetCodename sets Codename field to given value.

### HasCodename

`func (o *AdvisoryEOLReleaseData) HasCodename() bool`

HasCodename returns a boolean if a field has been set.

### GetCpe

`func (o *AdvisoryEOLReleaseData) GetCpe() string`

GetCpe returns the Cpe field if non-nil, zero value otherwise.

### GetCpeOk

`func (o *AdvisoryEOLReleaseData) GetCpeOk() (*string, bool)`

GetCpeOk returns a tuple with the Cpe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpe

`func (o *AdvisoryEOLReleaseData) SetCpe(v string)`

SetCpe sets Cpe field to given value.

### HasCpe

`func (o *AdvisoryEOLReleaseData) HasCpe() bool`

HasCpe returns a boolean if a field has been set.

### GetEolDate

`func (o *AdvisoryEOLReleaseData) GetEolDate() string`

GetEolDate returns the EolDate field if non-nil, zero value otherwise.

### GetEolDateOk

`func (o *AdvisoryEOLReleaseData) GetEolDateOk() (*string, bool)`

GetEolDateOk returns a tuple with the EolDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEolDate

`func (o *AdvisoryEOLReleaseData) SetEolDate(v string)`

SetEolDate sets EolDate field to given value.

### HasEolDate

`func (o *AdvisoryEOLReleaseData) HasEolDate() bool`

HasEolDate returns a boolean if a field has been set.

### GetEolDateExtendedSupport

`func (o *AdvisoryEOLReleaseData) GetEolDateExtendedSupport() string`

GetEolDateExtendedSupport returns the EolDateExtendedSupport field if non-nil, zero value otherwise.

### GetEolDateExtendedSupportOk

`func (o *AdvisoryEOLReleaseData) GetEolDateExtendedSupportOk() (*string, bool)`

GetEolDateExtendedSupportOk returns a tuple with the EolDateExtendedSupport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEolDateExtendedSupport

`func (o *AdvisoryEOLReleaseData) SetEolDateExtendedSupport(v string)`

SetEolDateExtendedSupport sets EolDateExtendedSupport field to given value.

### HasEolDateExtendedSupport

`func (o *AdvisoryEOLReleaseData) HasEolDateExtendedSupport() bool`

HasEolDateExtendedSupport returns a boolean if a field has been set.

### GetEolDatePremierSupport

`func (o *AdvisoryEOLReleaseData) GetEolDatePremierSupport() string`

GetEolDatePremierSupport returns the EolDatePremierSupport field if non-nil, zero value otherwise.

### GetEolDatePremierSupportOk

`func (o *AdvisoryEOLReleaseData) GetEolDatePremierSupportOk() (*string, bool)`

GetEolDatePremierSupportOk returns a tuple with the EolDatePremierSupport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEolDatePremierSupport

`func (o *AdvisoryEOLReleaseData) SetEolDatePremierSupport(v string)`

SetEolDatePremierSupport sets EolDatePremierSupport field to given value.

### HasEolDatePremierSupport

`func (o *AdvisoryEOLReleaseData) HasEolDatePremierSupport() bool`

HasEolDatePremierSupport returns a boolean if a field has been set.

### GetEolEltsDate

`func (o *AdvisoryEOLReleaseData) GetEolEltsDate() string`

GetEolEltsDate returns the EolEltsDate field if non-nil, zero value otherwise.

### GetEolEltsDateOk

`func (o *AdvisoryEOLReleaseData) GetEolEltsDateOk() (*string, bool)`

GetEolEltsDateOk returns a tuple with the EolEltsDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEolEltsDate

`func (o *AdvisoryEOLReleaseData) SetEolEltsDate(v string)`

SetEolEltsDate sets EolEltsDate field to given value.

### HasEolEltsDate

`func (o *AdvisoryEOLReleaseData) HasEolEltsDate() bool`

HasEolEltsDate returns a boolean if a field has been set.

### GetEolLtsDate

`func (o *AdvisoryEOLReleaseData) GetEolLtsDate() string`

GetEolLtsDate returns the EolLtsDate field if non-nil, zero value otherwise.

### GetEolLtsDateOk

`func (o *AdvisoryEOLReleaseData) GetEolLtsDateOk() (*string, bool)`

GetEolLtsDateOk returns a tuple with the EolLtsDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEolLtsDate

`func (o *AdvisoryEOLReleaseData) SetEolLtsDate(v string)`

SetEolLtsDate sets EolLtsDate field to given value.

### HasEolLtsDate

`func (o *AdvisoryEOLReleaseData) HasEolLtsDate() bool`

HasEolLtsDate returns a boolean if a field has been set.

### GetGitBranch

`func (o *AdvisoryEOLReleaseData) GetGitBranch() string`

GetGitBranch returns the GitBranch field if non-nil, zero value otherwise.

### GetGitBranchOk

`func (o *AdvisoryEOLReleaseData) GetGitBranchOk() (*string, bool)`

GetGitBranchOk returns a tuple with the GitBranch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitBranch

`func (o *AdvisoryEOLReleaseData) SetGitBranch(v string)`

SetGitBranch sets GitBranch field to given value.

### HasGitBranch

`func (o *AdvisoryEOLReleaseData) HasGitBranch() bool`

HasGitBranch returns a boolean if a field has been set.

### GetGitBranchUrl

`func (o *AdvisoryEOLReleaseData) GetGitBranchUrl() string`

GetGitBranchUrl returns the GitBranchUrl field if non-nil, zero value otherwise.

### GetGitBranchUrlOk

`func (o *AdvisoryEOLReleaseData) GetGitBranchUrlOk() (*string, bool)`

GetGitBranchUrlOk returns a tuple with the GitBranchUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitBranchUrl

`func (o *AdvisoryEOLReleaseData) SetGitBranchUrl(v string)`

SetGitBranchUrl sets GitBranchUrl field to given value.

### HasGitBranchUrl

`func (o *AdvisoryEOLReleaseData) HasGitBranchUrl() bool`

HasGitBranchUrl returns a boolean if a field has been set.

### GetLts

`func (o *AdvisoryEOLReleaseData) GetLts() bool`

GetLts returns the Lts field if non-nil, zero value otherwise.

### GetLtsOk

`func (o *AdvisoryEOLReleaseData) GetLtsOk() (*bool, bool)`

GetLtsOk returns a tuple with the Lts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLts

`func (o *AdvisoryEOLReleaseData) SetLts(v bool)`

SetLts sets Lts field to given value.

### HasLts

`func (o *AdvisoryEOLReleaseData) HasLts() bool`

HasLts returns a boolean if a field has been set.

### GetMinorReleases

`func (o *AdvisoryEOLReleaseData) GetMinorReleases() []string`

GetMinorReleases returns the MinorReleases field if non-nil, zero value otherwise.

### GetMinorReleasesOk

`func (o *AdvisoryEOLReleaseData) GetMinorReleasesOk() (*[]string, bool)`

GetMinorReleasesOk returns a tuple with the MinorReleases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinorReleases

`func (o *AdvisoryEOLReleaseData) SetMinorReleases(v []string)`

SetMinorReleases sets MinorReleases field to given value.

### HasMinorReleases

`func (o *AdvisoryEOLReleaseData) HasMinorReleases() bool`

HasMinorReleases returns a boolean if a field has been set.

### GetProduct

`func (o *AdvisoryEOLReleaseData) GetProduct() string`

GetProduct returns the Product field if non-nil, zero value otherwise.

### GetProductOk

`func (o *AdvisoryEOLReleaseData) GetProductOk() (*string, bool)`

GetProductOk returns a tuple with the Product field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduct

`func (o *AdvisoryEOLReleaseData) SetProduct(v string)`

SetProduct sets Product field to given value.

### HasProduct

`func (o *AdvisoryEOLReleaseData) HasProduct() bool`

HasProduct returns a boolean if a field has been set.

### GetReleaseDate

`func (o *AdvisoryEOLReleaseData) GetReleaseDate() string`

GetReleaseDate returns the ReleaseDate field if non-nil, zero value otherwise.

### GetReleaseDateOk

`func (o *AdvisoryEOLReleaseData) GetReleaseDateOk() (*string, bool)`

GetReleaseDateOk returns a tuple with the ReleaseDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseDate

`func (o *AdvisoryEOLReleaseData) SetReleaseDate(v string)`

SetReleaseDate sets ReleaseDate field to given value.

### HasReleaseDate

`func (o *AdvisoryEOLReleaseData) HasReleaseDate() bool`

HasReleaseDate returns a boolean if a field has been set.

### GetReleaseName

`func (o *AdvisoryEOLReleaseData) GetReleaseName() string`

GetReleaseName returns the ReleaseName field if non-nil, zero value otherwise.

### GetReleaseNameOk

`func (o *AdvisoryEOLReleaseData) GetReleaseNameOk() (*string, bool)`

GetReleaseNameOk returns a tuple with the ReleaseName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseName

`func (o *AdvisoryEOLReleaseData) SetReleaseName(v string)`

SetReleaseName sets ReleaseName field to given value.

### HasReleaseName

`func (o *AdvisoryEOLReleaseData) HasReleaseName() bool`

HasReleaseName returns a boolean if a field has been set.

### GetSourceUrl

`func (o *AdvisoryEOLReleaseData) GetSourceUrl() string`

GetSourceUrl returns the SourceUrl field if non-nil, zero value otherwise.

### GetSourceUrlOk

`func (o *AdvisoryEOLReleaseData) GetSourceUrlOk() (*string, bool)`

GetSourceUrlOk returns a tuple with the SourceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceUrl

`func (o *AdvisoryEOLReleaseData) SetSourceUrl(v string)`

SetSourceUrl sets SourceUrl field to given value.

### HasSourceUrl

`func (o *AdvisoryEOLReleaseData) HasSourceUrl() bool`

HasSourceUrl returns a boolean if a field has been set.

### GetTechnologyLevel

`func (o *AdvisoryEOLReleaseData) GetTechnologyLevel() string`

GetTechnologyLevel returns the TechnologyLevel field if non-nil, zero value otherwise.

### GetTechnologyLevelOk

`func (o *AdvisoryEOLReleaseData) GetTechnologyLevelOk() (*string, bool)`

GetTechnologyLevelOk returns a tuple with the TechnologyLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTechnologyLevel

`func (o *AdvisoryEOLReleaseData) SetTechnologyLevel(v string)`

SetTechnologyLevel sets TechnologyLevel field to given value.

### HasTechnologyLevel

`func (o *AdvisoryEOLReleaseData) HasTechnologyLevel() bool`

HasTechnologyLevel returns a boolean if a field has been set.

### GetVendor

`func (o *AdvisoryEOLReleaseData) GetVendor() string`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *AdvisoryEOLReleaseData) GetVendorOk() (*string, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *AdvisoryEOLReleaseData) SetVendor(v string)`

SetVendor sets Vendor field to given value.

### HasVendor

`func (o *AdvisoryEOLReleaseData) HasVendor() bool`

HasVendor returns a boolean if a field has been set.

### GetVersion

`func (o *AdvisoryEOLReleaseData) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *AdvisoryEOLReleaseData) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *AdvisoryEOLReleaseData) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *AdvisoryEOLReleaseData) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetVersionApi

`func (o *AdvisoryEOLReleaseData) GetVersionApi() string`

GetVersionApi returns the VersionApi field if non-nil, zero value otherwise.

### GetVersionApiOk

`func (o *AdvisoryEOLReleaseData) GetVersionApiOk() (*string, bool)`

GetVersionApiOk returns a tuple with the VersionApi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionApi

`func (o *AdvisoryEOLReleaseData) SetVersionApi(v string)`

SetVersionApi sets VersionApi field to given value.

### HasVersionApi

`func (o *AdvisoryEOLReleaseData) HasVersionApi() bool`

HasVersionApi returns a boolean if a field has been set.

### GetVersionDarwin

`func (o *AdvisoryEOLReleaseData) GetVersionDarwin() string`

GetVersionDarwin returns the VersionDarwin field if non-nil, zero value otherwise.

### GetVersionDarwinOk

`func (o *AdvisoryEOLReleaseData) GetVersionDarwinOk() (*string, bool)`

GetVersionDarwinOk returns a tuple with the VersionDarwin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionDarwin

`func (o *AdvisoryEOLReleaseData) SetVersionDarwin(v string)`

SetVersionDarwin sets VersionDarwin field to given value.

### HasVersionDarwin

`func (o *AdvisoryEOLReleaseData) HasVersionDarwin() bool`

HasVersionDarwin returns a boolean if a field has been set.

### GetVersionSunos

`func (o *AdvisoryEOLReleaseData) GetVersionSunos() string`

GetVersionSunos returns the VersionSunos field if non-nil, zero value otherwise.

### GetVersionSunosOk

`func (o *AdvisoryEOLReleaseData) GetVersionSunosOk() (*string, bool)`

GetVersionSunosOk returns a tuple with the VersionSunos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionSunos

`func (o *AdvisoryEOLReleaseData) SetVersionSunos(v string)`

SetVersionSunos sets VersionSunos field to given value.

### HasVersionSunos

`func (o *AdvisoryEOLReleaseData) HasVersionSunos() bool`

HasVersionSunos returns a boolean if a field has been set.

### GetWindowsCurrentBuild

`func (o *AdvisoryEOLReleaseData) GetWindowsCurrentBuild() string`

GetWindowsCurrentBuild returns the WindowsCurrentBuild field if non-nil, zero value otherwise.

### GetWindowsCurrentBuildOk

`func (o *AdvisoryEOLReleaseData) GetWindowsCurrentBuildOk() (*string, bool)`

GetWindowsCurrentBuildOk returns a tuple with the WindowsCurrentBuild field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindowsCurrentBuild

`func (o *AdvisoryEOLReleaseData) SetWindowsCurrentBuild(v string)`

SetWindowsCurrentBuild sets WindowsCurrentBuild field to given value.

### HasWindowsCurrentBuild

`func (o *AdvisoryEOLReleaseData) HasWindowsCurrentBuild() bool`

HasWindowsCurrentBuild returns a boolean if a field has been set.

### GetWindowsDisplayVersion

`func (o *AdvisoryEOLReleaseData) GetWindowsDisplayVersion() string`

GetWindowsDisplayVersion returns the WindowsDisplayVersion field if non-nil, zero value otherwise.

### GetWindowsDisplayVersionOk

`func (o *AdvisoryEOLReleaseData) GetWindowsDisplayVersionOk() (*string, bool)`

GetWindowsDisplayVersionOk returns a tuple with the WindowsDisplayVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindowsDisplayVersion

`func (o *AdvisoryEOLReleaseData) SetWindowsDisplayVersion(v string)`

SetWindowsDisplayVersion sets WindowsDisplayVersion field to given value.

### HasWindowsDisplayVersion

`func (o *AdvisoryEOLReleaseData) HasWindowsDisplayVersion() bool`

HasWindowsDisplayVersion returns a boolean if a field has been set.

### GetWindowsEditionId

`func (o *AdvisoryEOLReleaseData) GetWindowsEditionId() string`

GetWindowsEditionId returns the WindowsEditionId field if non-nil, zero value otherwise.

### GetWindowsEditionIdOk

`func (o *AdvisoryEOLReleaseData) GetWindowsEditionIdOk() (*string, bool)`

GetWindowsEditionIdOk returns a tuple with the WindowsEditionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindowsEditionId

`func (o *AdvisoryEOLReleaseData) SetWindowsEditionId(v string)`

SetWindowsEditionId sets WindowsEditionId field to given value.

### HasWindowsEditionId

`func (o *AdvisoryEOLReleaseData) HasWindowsEditionId() bool`

HasWindowsEditionId returns a boolean if a field has been set.

### GetWindowsInsiderPreview

`func (o *AdvisoryEOLReleaseData) GetWindowsInsiderPreview() bool`

GetWindowsInsiderPreview returns the WindowsInsiderPreview field if non-nil, zero value otherwise.

### GetWindowsInsiderPreviewOk

`func (o *AdvisoryEOLReleaseData) GetWindowsInsiderPreviewOk() (*bool, bool)`

GetWindowsInsiderPreviewOk returns a tuple with the WindowsInsiderPreview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindowsInsiderPreview

`func (o *AdvisoryEOLReleaseData) SetWindowsInsiderPreview(v bool)`

SetWindowsInsiderPreview sets WindowsInsiderPreview field to given value.

### HasWindowsInsiderPreview

`func (o *AdvisoryEOLReleaseData) HasWindowsInsiderPreview() bool`

HasWindowsInsiderPreview returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


