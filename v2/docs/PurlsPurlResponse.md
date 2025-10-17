# PurlsPurlResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Artifacts** | Pointer to **map[string]interface{}** |  | [optional] 
**Cves** | Pointer to **[]string** |  | [optional] 
**Licenses** | Pointer to **[]string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**PublishedDate** | Pointer to **string** |  | [optional] 
**Purl** | Pointer to **[]string** |  | [optional] 
**Version** | Pointer to **string** |  | [optional] 
**Vulnerabilities** | Pointer to [**[]PurlsVulnerability**](PurlsVulnerability.md) |  | [optional] 

## Methods

### NewPurlsPurlResponse

`func NewPurlsPurlResponse() *PurlsPurlResponse`

NewPurlsPurlResponse instantiates a new PurlsPurlResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPurlsPurlResponseWithDefaults

`func NewPurlsPurlResponseWithDefaults() *PurlsPurlResponse`

NewPurlsPurlResponseWithDefaults instantiates a new PurlsPurlResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArtifacts

`func (o *PurlsPurlResponse) GetArtifacts() map[string]interface{}`

GetArtifacts returns the Artifacts field if non-nil, zero value otherwise.

### GetArtifactsOk

`func (o *PurlsPurlResponse) GetArtifactsOk() (*map[string]interface{}, bool)`

GetArtifactsOk returns a tuple with the Artifacts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifacts

`func (o *PurlsPurlResponse) SetArtifacts(v map[string]interface{})`

SetArtifacts sets Artifacts field to given value.

### HasArtifacts

`func (o *PurlsPurlResponse) HasArtifacts() bool`

HasArtifacts returns a boolean if a field has been set.

### GetCves

`func (o *PurlsPurlResponse) GetCves() []string`

GetCves returns the Cves field if non-nil, zero value otherwise.

### GetCvesOk

`func (o *PurlsPurlResponse) GetCvesOk() (*[]string, bool)`

GetCvesOk returns a tuple with the Cves field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCves

`func (o *PurlsPurlResponse) SetCves(v []string)`

SetCves sets Cves field to given value.

### HasCves

`func (o *PurlsPurlResponse) HasCves() bool`

HasCves returns a boolean if a field has been set.

### GetLicenses

`func (o *PurlsPurlResponse) GetLicenses() []string`

GetLicenses returns the Licenses field if non-nil, zero value otherwise.

### GetLicensesOk

`func (o *PurlsPurlResponse) GetLicensesOk() (*[]string, bool)`

GetLicensesOk returns a tuple with the Licenses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenses

`func (o *PurlsPurlResponse) SetLicenses(v []string)`

SetLicenses sets Licenses field to given value.

### HasLicenses

`func (o *PurlsPurlResponse) HasLicenses() bool`

HasLicenses returns a boolean if a field has been set.

### GetName

`func (o *PurlsPurlResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PurlsPurlResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PurlsPurlResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PurlsPurlResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPublishedDate

`func (o *PurlsPurlResponse) GetPublishedDate() string`

GetPublishedDate returns the PublishedDate field if non-nil, zero value otherwise.

### GetPublishedDateOk

`func (o *PurlsPurlResponse) GetPublishedDateOk() (*string, bool)`

GetPublishedDateOk returns a tuple with the PublishedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishedDate

`func (o *PurlsPurlResponse) SetPublishedDate(v string)`

SetPublishedDate sets PublishedDate field to given value.

### HasPublishedDate

`func (o *PurlsPurlResponse) HasPublishedDate() bool`

HasPublishedDate returns a boolean if a field has been set.

### GetPurl

`func (o *PurlsPurlResponse) GetPurl() []string`

GetPurl returns the Purl field if non-nil, zero value otherwise.

### GetPurlOk

`func (o *PurlsPurlResponse) GetPurlOk() (*[]string, bool)`

GetPurlOk returns a tuple with the Purl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurl

`func (o *PurlsPurlResponse) SetPurl(v []string)`

SetPurl sets Purl field to given value.

### HasPurl

`func (o *PurlsPurlResponse) HasPurl() bool`

HasPurl returns a boolean if a field has been set.

### GetVersion

`func (o *PurlsPurlResponse) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *PurlsPurlResponse) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *PurlsPurlResponse) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *PurlsPurlResponse) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetVulnerabilities

`func (o *PurlsPurlResponse) GetVulnerabilities() []PurlsVulnerability`

GetVulnerabilities returns the Vulnerabilities field if non-nil, zero value otherwise.

### GetVulnerabilitiesOk

`func (o *PurlsPurlResponse) GetVulnerabilitiesOk() (*[]PurlsVulnerability, bool)`

GetVulnerabilitiesOk returns a tuple with the Vulnerabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVulnerabilities

`func (o *PurlsPurlResponse) SetVulnerabilities(v []PurlsVulnerability)`

SetVulnerabilities sets Vulnerabilities field to given value.

### HasVulnerabilities

`func (o *PurlsPurlResponse) HasVulnerabilities() bool`

HasVulnerabilities returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


