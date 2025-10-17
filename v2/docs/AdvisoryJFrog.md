# AdvisoryJFrog

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cpes** | Pointer to [**[]AdvisoryNVD20CVECPEMatch**](AdvisoryNVD20CVECPEMatch.md) |  | [optional] 
**Cve** | Pointer to **[]string** |  | [optional] 
**DateAdded** | Pointer to **string** |  | [optional] 
**Product** | Pointer to **string** |  | [optional] 
**Severity** | Pointer to **string** |  | [optional] 
**Summary** | Pointer to **string** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 
**Versions** | Pointer to **[]string** |  | [optional] 

## Methods

### NewAdvisoryJFrog

`func NewAdvisoryJFrog() *AdvisoryJFrog`

NewAdvisoryJFrog instantiates a new AdvisoryJFrog object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdvisoryJFrogWithDefaults

`func NewAdvisoryJFrogWithDefaults() *AdvisoryJFrog`

NewAdvisoryJFrogWithDefaults instantiates a new AdvisoryJFrog object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCpes

`func (o *AdvisoryJFrog) GetCpes() []AdvisoryNVD20CVECPEMatch`

GetCpes returns the Cpes field if non-nil, zero value otherwise.

### GetCpesOk

`func (o *AdvisoryJFrog) GetCpesOk() (*[]AdvisoryNVD20CVECPEMatch, bool)`

GetCpesOk returns a tuple with the Cpes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpes

`func (o *AdvisoryJFrog) SetCpes(v []AdvisoryNVD20CVECPEMatch)`

SetCpes sets Cpes field to given value.

### HasCpes

`func (o *AdvisoryJFrog) HasCpes() bool`

HasCpes returns a boolean if a field has been set.

### GetCve

`func (o *AdvisoryJFrog) GetCve() []string`

GetCve returns the Cve field if non-nil, zero value otherwise.

### GetCveOk

`func (o *AdvisoryJFrog) GetCveOk() (*[]string, bool)`

GetCveOk returns a tuple with the Cve field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCve

`func (o *AdvisoryJFrog) SetCve(v []string)`

SetCve sets Cve field to given value.

### HasCve

`func (o *AdvisoryJFrog) HasCve() bool`

HasCve returns a boolean if a field has been set.

### GetDateAdded

`func (o *AdvisoryJFrog) GetDateAdded() string`

GetDateAdded returns the DateAdded field if non-nil, zero value otherwise.

### GetDateAddedOk

`func (o *AdvisoryJFrog) GetDateAddedOk() (*string, bool)`

GetDateAddedOk returns a tuple with the DateAdded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateAdded

`func (o *AdvisoryJFrog) SetDateAdded(v string)`

SetDateAdded sets DateAdded field to given value.

### HasDateAdded

`func (o *AdvisoryJFrog) HasDateAdded() bool`

HasDateAdded returns a boolean if a field has been set.

### GetProduct

`func (o *AdvisoryJFrog) GetProduct() string`

GetProduct returns the Product field if non-nil, zero value otherwise.

### GetProductOk

`func (o *AdvisoryJFrog) GetProductOk() (*string, bool)`

GetProductOk returns a tuple with the Product field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduct

`func (o *AdvisoryJFrog) SetProduct(v string)`

SetProduct sets Product field to given value.

### HasProduct

`func (o *AdvisoryJFrog) HasProduct() bool`

HasProduct returns a boolean if a field has been set.

### GetSeverity

`func (o *AdvisoryJFrog) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *AdvisoryJFrog) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *AdvisoryJFrog) SetSeverity(v string)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *AdvisoryJFrog) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetSummary

`func (o *AdvisoryJFrog) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *AdvisoryJFrog) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *AdvisoryJFrog) SetSummary(v string)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *AdvisoryJFrog) HasSummary() bool`

HasSummary returns a boolean if a field has been set.

### GetUrl

`func (o *AdvisoryJFrog) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *AdvisoryJFrog) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *AdvisoryJFrog) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *AdvisoryJFrog) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetVersions

`func (o *AdvisoryJFrog) GetVersions() []string`

GetVersions returns the Versions field if non-nil, zero value otherwise.

### GetVersionsOk

`func (o *AdvisoryJFrog) GetVersionsOk() (*[]string, bool)`

GetVersionsOk returns a tuple with the Versions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersions

`func (o *AdvisoryJFrog) SetVersions(v []string)`

SetVersions sets Versions field to given value.

### HasVersions

`func (o *AdvisoryJFrog) HasVersions() bool`

HasVersions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


