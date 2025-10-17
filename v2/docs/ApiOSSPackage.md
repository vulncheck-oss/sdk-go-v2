# ApiOSSPackage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Artifacts** | Pointer to [**ApiOSSPackageArtifacts**](ApiOSSPackageArtifacts.md) |  | [optional] 
**Cves** | Pointer to **[]string** |  | [optional] 
**Licenses** | Pointer to **[]string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**PublishedDate** | Pointer to **string** |  | [optional] 
**Purl** | Pointer to **[]string** |  | [optional] 
**ResearchAttributes** | Pointer to [**ApiOSSPackageResearchAttributes**](ApiOSSPackageResearchAttributes.md) |  | [optional] 
**Version** | Pointer to **string** |  | [optional] 
**Vulnerabilities** | Pointer to [**[]ApiOSSPackageVulnerability**](ApiOSSPackageVulnerability.md) |  | [optional] 

## Methods

### NewApiOSSPackage

`func NewApiOSSPackage() *ApiOSSPackage`

NewApiOSSPackage instantiates a new ApiOSSPackage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiOSSPackageWithDefaults

`func NewApiOSSPackageWithDefaults() *ApiOSSPackage`

NewApiOSSPackageWithDefaults instantiates a new ApiOSSPackage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArtifacts

`func (o *ApiOSSPackage) GetArtifacts() ApiOSSPackageArtifacts`

GetArtifacts returns the Artifacts field if non-nil, zero value otherwise.

### GetArtifactsOk

`func (o *ApiOSSPackage) GetArtifactsOk() (*ApiOSSPackageArtifacts, bool)`

GetArtifactsOk returns a tuple with the Artifacts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifacts

`func (o *ApiOSSPackage) SetArtifacts(v ApiOSSPackageArtifacts)`

SetArtifacts sets Artifacts field to given value.

### HasArtifacts

`func (o *ApiOSSPackage) HasArtifacts() bool`

HasArtifacts returns a boolean if a field has been set.

### GetCves

`func (o *ApiOSSPackage) GetCves() []string`

GetCves returns the Cves field if non-nil, zero value otherwise.

### GetCvesOk

`func (o *ApiOSSPackage) GetCvesOk() (*[]string, bool)`

GetCvesOk returns a tuple with the Cves field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCves

`func (o *ApiOSSPackage) SetCves(v []string)`

SetCves sets Cves field to given value.

### HasCves

`func (o *ApiOSSPackage) HasCves() bool`

HasCves returns a boolean if a field has been set.

### GetLicenses

`func (o *ApiOSSPackage) GetLicenses() []string`

GetLicenses returns the Licenses field if non-nil, zero value otherwise.

### GetLicensesOk

`func (o *ApiOSSPackage) GetLicensesOk() (*[]string, bool)`

GetLicensesOk returns a tuple with the Licenses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenses

`func (o *ApiOSSPackage) SetLicenses(v []string)`

SetLicenses sets Licenses field to given value.

### HasLicenses

`func (o *ApiOSSPackage) HasLicenses() bool`

HasLicenses returns a boolean if a field has been set.

### GetName

`func (o *ApiOSSPackage) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApiOSSPackage) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApiOSSPackage) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ApiOSSPackage) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPublishedDate

`func (o *ApiOSSPackage) GetPublishedDate() string`

GetPublishedDate returns the PublishedDate field if non-nil, zero value otherwise.

### GetPublishedDateOk

`func (o *ApiOSSPackage) GetPublishedDateOk() (*string, bool)`

GetPublishedDateOk returns a tuple with the PublishedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishedDate

`func (o *ApiOSSPackage) SetPublishedDate(v string)`

SetPublishedDate sets PublishedDate field to given value.

### HasPublishedDate

`func (o *ApiOSSPackage) HasPublishedDate() bool`

HasPublishedDate returns a boolean if a field has been set.

### GetPurl

`func (o *ApiOSSPackage) GetPurl() []string`

GetPurl returns the Purl field if non-nil, zero value otherwise.

### GetPurlOk

`func (o *ApiOSSPackage) GetPurlOk() (*[]string, bool)`

GetPurlOk returns a tuple with the Purl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurl

`func (o *ApiOSSPackage) SetPurl(v []string)`

SetPurl sets Purl field to given value.

### HasPurl

`func (o *ApiOSSPackage) HasPurl() bool`

HasPurl returns a boolean if a field has been set.

### GetResearchAttributes

`func (o *ApiOSSPackage) GetResearchAttributes() ApiOSSPackageResearchAttributes`

GetResearchAttributes returns the ResearchAttributes field if non-nil, zero value otherwise.

### GetResearchAttributesOk

`func (o *ApiOSSPackage) GetResearchAttributesOk() (*ApiOSSPackageResearchAttributes, bool)`

GetResearchAttributesOk returns a tuple with the ResearchAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResearchAttributes

`func (o *ApiOSSPackage) SetResearchAttributes(v ApiOSSPackageResearchAttributes)`

SetResearchAttributes sets ResearchAttributes field to given value.

### HasResearchAttributes

`func (o *ApiOSSPackage) HasResearchAttributes() bool`

HasResearchAttributes returns a boolean if a field has been set.

### GetVersion

`func (o *ApiOSSPackage) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ApiOSSPackage) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ApiOSSPackage) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ApiOSSPackage) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetVulnerabilities

`func (o *ApiOSSPackage) GetVulnerabilities() []ApiOSSPackageVulnerability`

GetVulnerabilities returns the Vulnerabilities field if non-nil, zero value otherwise.

### GetVulnerabilitiesOk

`func (o *ApiOSSPackage) GetVulnerabilitiesOk() (*[]ApiOSSPackageVulnerability, bool)`

GetVulnerabilitiesOk returns a tuple with the Vulnerabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVulnerabilities

`func (o *ApiOSSPackage) SetVulnerabilities(v []ApiOSSPackageVulnerability)`

SetVulnerabilities sets Vulnerabilities field to given value.

### HasVulnerabilities

`func (o *ApiOSSPackage) HasVulnerabilities() bool`

HasVulnerabilities returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


