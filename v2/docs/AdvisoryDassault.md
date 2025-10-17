# AdvisoryDassault

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AffectedProducts** | Pointer to **string** |  | [optional] 
**AffectedVersions** | Pointer to **string** |  | [optional] 
**Cve** | Pointer to **[]string** |  | [optional] 
**DateAdded** | Pointer to **string** |  | [optional] 
**Details** | Pointer to **string** |  | [optional] 
**Severity** | Pointer to **string** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 

## Methods

### NewAdvisoryDassault

`func NewAdvisoryDassault() *AdvisoryDassault`

NewAdvisoryDassault instantiates a new AdvisoryDassault object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdvisoryDassaultWithDefaults

`func NewAdvisoryDassaultWithDefaults() *AdvisoryDassault`

NewAdvisoryDassaultWithDefaults instantiates a new AdvisoryDassault object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAffectedProducts

`func (o *AdvisoryDassault) GetAffectedProducts() string`

GetAffectedProducts returns the AffectedProducts field if non-nil, zero value otherwise.

### GetAffectedProductsOk

`func (o *AdvisoryDassault) GetAffectedProductsOk() (*string, bool)`

GetAffectedProductsOk returns a tuple with the AffectedProducts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffectedProducts

`func (o *AdvisoryDassault) SetAffectedProducts(v string)`

SetAffectedProducts sets AffectedProducts field to given value.

### HasAffectedProducts

`func (o *AdvisoryDassault) HasAffectedProducts() bool`

HasAffectedProducts returns a boolean if a field has been set.

### GetAffectedVersions

`func (o *AdvisoryDassault) GetAffectedVersions() string`

GetAffectedVersions returns the AffectedVersions field if non-nil, zero value otherwise.

### GetAffectedVersionsOk

`func (o *AdvisoryDassault) GetAffectedVersionsOk() (*string, bool)`

GetAffectedVersionsOk returns a tuple with the AffectedVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffectedVersions

`func (o *AdvisoryDassault) SetAffectedVersions(v string)`

SetAffectedVersions sets AffectedVersions field to given value.

### HasAffectedVersions

`func (o *AdvisoryDassault) HasAffectedVersions() bool`

HasAffectedVersions returns a boolean if a field has been set.

### GetCve

`func (o *AdvisoryDassault) GetCve() []string`

GetCve returns the Cve field if non-nil, zero value otherwise.

### GetCveOk

`func (o *AdvisoryDassault) GetCveOk() (*[]string, bool)`

GetCveOk returns a tuple with the Cve field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCve

`func (o *AdvisoryDassault) SetCve(v []string)`

SetCve sets Cve field to given value.

### HasCve

`func (o *AdvisoryDassault) HasCve() bool`

HasCve returns a boolean if a field has been set.

### GetDateAdded

`func (o *AdvisoryDassault) GetDateAdded() string`

GetDateAdded returns the DateAdded field if non-nil, zero value otherwise.

### GetDateAddedOk

`func (o *AdvisoryDassault) GetDateAddedOk() (*string, bool)`

GetDateAddedOk returns a tuple with the DateAdded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateAdded

`func (o *AdvisoryDassault) SetDateAdded(v string)`

SetDateAdded sets DateAdded field to given value.

### HasDateAdded

`func (o *AdvisoryDassault) HasDateAdded() bool`

HasDateAdded returns a boolean if a field has been set.

### GetDetails

`func (o *AdvisoryDassault) GetDetails() string`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *AdvisoryDassault) GetDetailsOk() (*string, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *AdvisoryDassault) SetDetails(v string)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *AdvisoryDassault) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### GetSeverity

`func (o *AdvisoryDassault) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *AdvisoryDassault) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *AdvisoryDassault) SetSeverity(v string)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *AdvisoryDassault) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetTitle

`func (o *AdvisoryDassault) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *AdvisoryDassault) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *AdvisoryDassault) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *AdvisoryDassault) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *AdvisoryDassault) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *AdvisoryDassault) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *AdvisoryDassault) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *AdvisoryDassault) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUrl

`func (o *AdvisoryDassault) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *AdvisoryDassault) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *AdvisoryDassault) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *AdvisoryDassault) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


