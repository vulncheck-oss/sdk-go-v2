# PurlBatchVulnFinding

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cves** | Pointer to **[]string** | list of associated CVE &#39;s | [optional] 
**Purl** | Pointer to **string** | the purl, ex. hex/coherence@0.1.2 | [optional] 
**PurlStruct** | Pointer to [**PurlPackageURLJSON**](PurlPackageURLJSON.md) |  | [optional] 
**ResearchAttributes** | Pointer to [**ApiOSSPackageResearchAttributes**](ApiOSSPackageResearchAttributes.md) |  | [optional] 
**Vulnerabilities** | Pointer to [**[]ApiOSSPackageVulnerability**](ApiOSSPackageVulnerability.md) | list of associated vulnerabilities | [optional] 

## Methods

### NewPurlBatchVulnFinding

`func NewPurlBatchVulnFinding() *PurlBatchVulnFinding`

NewPurlBatchVulnFinding instantiates a new PurlBatchVulnFinding object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPurlBatchVulnFindingWithDefaults

`func NewPurlBatchVulnFindingWithDefaults() *PurlBatchVulnFinding`

NewPurlBatchVulnFindingWithDefaults instantiates a new PurlBatchVulnFinding object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCves

`func (o *PurlBatchVulnFinding) GetCves() []string`

GetCves returns the Cves field if non-nil, zero value otherwise.

### GetCvesOk

`func (o *PurlBatchVulnFinding) GetCvesOk() (*[]string, bool)`

GetCvesOk returns a tuple with the Cves field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCves

`func (o *PurlBatchVulnFinding) SetCves(v []string)`

SetCves sets Cves field to given value.

### HasCves

`func (o *PurlBatchVulnFinding) HasCves() bool`

HasCves returns a boolean if a field has been set.

### GetPurl

`func (o *PurlBatchVulnFinding) GetPurl() string`

GetPurl returns the Purl field if non-nil, zero value otherwise.

### GetPurlOk

`func (o *PurlBatchVulnFinding) GetPurlOk() (*string, bool)`

GetPurlOk returns a tuple with the Purl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurl

`func (o *PurlBatchVulnFinding) SetPurl(v string)`

SetPurl sets Purl field to given value.

### HasPurl

`func (o *PurlBatchVulnFinding) HasPurl() bool`

HasPurl returns a boolean if a field has been set.

### GetPurlStruct

`func (o *PurlBatchVulnFinding) GetPurlStruct() PurlPackageURLJSON`

GetPurlStruct returns the PurlStruct field if non-nil, zero value otherwise.

### GetPurlStructOk

`func (o *PurlBatchVulnFinding) GetPurlStructOk() (*PurlPackageURLJSON, bool)`

GetPurlStructOk returns a tuple with the PurlStruct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurlStruct

`func (o *PurlBatchVulnFinding) SetPurlStruct(v PurlPackageURLJSON)`

SetPurlStruct sets PurlStruct field to given value.

### HasPurlStruct

`func (o *PurlBatchVulnFinding) HasPurlStruct() bool`

HasPurlStruct returns a boolean if a field has been set.

### GetResearchAttributes

`func (o *PurlBatchVulnFinding) GetResearchAttributes() ApiOSSPackageResearchAttributes`

GetResearchAttributes returns the ResearchAttributes field if non-nil, zero value otherwise.

### GetResearchAttributesOk

`func (o *PurlBatchVulnFinding) GetResearchAttributesOk() (*ApiOSSPackageResearchAttributes, bool)`

GetResearchAttributesOk returns a tuple with the ResearchAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResearchAttributes

`func (o *PurlBatchVulnFinding) SetResearchAttributes(v ApiOSSPackageResearchAttributes)`

SetResearchAttributes sets ResearchAttributes field to given value.

### HasResearchAttributes

`func (o *PurlBatchVulnFinding) HasResearchAttributes() bool`

HasResearchAttributes returns a boolean if a field has been set.

### GetVulnerabilities

`func (o *PurlBatchVulnFinding) GetVulnerabilities() []ApiOSSPackageVulnerability`

GetVulnerabilities returns the Vulnerabilities field if non-nil, zero value otherwise.

### GetVulnerabilitiesOk

`func (o *PurlBatchVulnFinding) GetVulnerabilitiesOk() (*[]ApiOSSPackageVulnerability, bool)`

GetVulnerabilitiesOk returns a tuple with the Vulnerabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVulnerabilities

`func (o *PurlBatchVulnFinding) SetVulnerabilities(v []ApiOSSPackageVulnerability)`

SetVulnerabilities sets Vulnerabilities field to given value.

### HasVulnerabilities

`func (o *PurlBatchVulnFinding) HasVulnerabilities() bool`

HasVulnerabilities returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


