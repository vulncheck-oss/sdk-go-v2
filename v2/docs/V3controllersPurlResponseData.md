# V3controllersPurlResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cves** | Pointer to **[]string** | list of associated CVE &#39;s | [optional] 
**Vulnerabilities** | Pointer to [**[]ApiOSSPackageVulnerability**](ApiOSSPackageVulnerability.md) | list of associated vulnerabilities | [optional] 

## Methods

### NewV3controllersPurlResponseData

`func NewV3controllersPurlResponseData() *V3controllersPurlResponseData`

NewV3controllersPurlResponseData instantiates a new V3controllersPurlResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV3controllersPurlResponseDataWithDefaults

`func NewV3controllersPurlResponseDataWithDefaults() *V3controllersPurlResponseData`

NewV3controllersPurlResponseDataWithDefaults instantiates a new V3controllersPurlResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCves

`func (o *V3controllersPurlResponseData) GetCves() []string`

GetCves returns the Cves field if non-nil, zero value otherwise.

### GetCvesOk

`func (o *V3controllersPurlResponseData) GetCvesOk() (*[]string, bool)`

GetCvesOk returns a tuple with the Cves field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCves

`func (o *V3controllersPurlResponseData) SetCves(v []string)`

SetCves sets Cves field to given value.

### HasCves

`func (o *V3controllersPurlResponseData) HasCves() bool`

HasCves returns a boolean if a field has been set.

### GetVulnerabilities

`func (o *V3controllersPurlResponseData) GetVulnerabilities() []ApiOSSPackageVulnerability`

GetVulnerabilities returns the Vulnerabilities field if non-nil, zero value otherwise.

### GetVulnerabilitiesOk

`func (o *V3controllersPurlResponseData) GetVulnerabilitiesOk() (*[]ApiOSSPackageVulnerability, bool)`

GetVulnerabilitiesOk returns a tuple with the Vulnerabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVulnerabilities

`func (o *V3controllersPurlResponseData) SetVulnerabilities(v []ApiOSSPackageVulnerability)`

SetVulnerabilities sets Vulnerabilities field to given value.

### HasVulnerabilities

`func (o *V3controllersPurlResponseData) HasVulnerabilities() bool`

HasVulnerabilities returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


