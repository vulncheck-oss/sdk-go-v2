# AdvisoryLexmarkAdvisory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AffectedProducts** | Pointer to [**[]AdvisoryAffectedProduct**](AdvisoryAffectedProduct.md) |  | [optional] 
**Cve** | Pointer to **[]string** |  | [optional] 
**DateAdded** | Pointer to **string** |  | [optional] 
**Details** | Pointer to **string** |  | [optional] 
**Impact** | Pointer to **string** |  | [optional] 
**LastUpdate** | Pointer to **string** |  | [optional] 
**Link** | Pointer to **string** |  | [optional] 
**PublicReleaseDate** | Pointer to **string** |  | [optional] 
**References** | Pointer to **[]string** |  | [optional] 
**Revision** | Pointer to **string** |  | [optional] 
**Summary** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 
**Workarounds** | Pointer to **string** |  | [optional] 

## Methods

### NewAdvisoryLexmarkAdvisory

`func NewAdvisoryLexmarkAdvisory() *AdvisoryLexmarkAdvisory`

NewAdvisoryLexmarkAdvisory instantiates a new AdvisoryLexmarkAdvisory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdvisoryLexmarkAdvisoryWithDefaults

`func NewAdvisoryLexmarkAdvisoryWithDefaults() *AdvisoryLexmarkAdvisory`

NewAdvisoryLexmarkAdvisoryWithDefaults instantiates a new AdvisoryLexmarkAdvisory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAffectedProducts

`func (o *AdvisoryLexmarkAdvisory) GetAffectedProducts() []AdvisoryAffectedProduct`

GetAffectedProducts returns the AffectedProducts field if non-nil, zero value otherwise.

### GetAffectedProductsOk

`func (o *AdvisoryLexmarkAdvisory) GetAffectedProductsOk() (*[]AdvisoryAffectedProduct, bool)`

GetAffectedProductsOk returns a tuple with the AffectedProducts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffectedProducts

`func (o *AdvisoryLexmarkAdvisory) SetAffectedProducts(v []AdvisoryAffectedProduct)`

SetAffectedProducts sets AffectedProducts field to given value.

### HasAffectedProducts

`func (o *AdvisoryLexmarkAdvisory) HasAffectedProducts() bool`

HasAffectedProducts returns a boolean if a field has been set.

### GetCve

`func (o *AdvisoryLexmarkAdvisory) GetCve() []string`

GetCve returns the Cve field if non-nil, zero value otherwise.

### GetCveOk

`func (o *AdvisoryLexmarkAdvisory) GetCveOk() (*[]string, bool)`

GetCveOk returns a tuple with the Cve field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCve

`func (o *AdvisoryLexmarkAdvisory) SetCve(v []string)`

SetCve sets Cve field to given value.

### HasCve

`func (o *AdvisoryLexmarkAdvisory) HasCve() bool`

HasCve returns a boolean if a field has been set.

### GetDateAdded

`func (o *AdvisoryLexmarkAdvisory) GetDateAdded() string`

GetDateAdded returns the DateAdded field if non-nil, zero value otherwise.

### GetDateAddedOk

`func (o *AdvisoryLexmarkAdvisory) GetDateAddedOk() (*string, bool)`

GetDateAddedOk returns a tuple with the DateAdded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateAdded

`func (o *AdvisoryLexmarkAdvisory) SetDateAdded(v string)`

SetDateAdded sets DateAdded field to given value.

### HasDateAdded

`func (o *AdvisoryLexmarkAdvisory) HasDateAdded() bool`

HasDateAdded returns a boolean if a field has been set.

### GetDetails

`func (o *AdvisoryLexmarkAdvisory) GetDetails() string`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *AdvisoryLexmarkAdvisory) GetDetailsOk() (*string, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *AdvisoryLexmarkAdvisory) SetDetails(v string)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *AdvisoryLexmarkAdvisory) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### GetImpact

`func (o *AdvisoryLexmarkAdvisory) GetImpact() string`

GetImpact returns the Impact field if non-nil, zero value otherwise.

### GetImpactOk

`func (o *AdvisoryLexmarkAdvisory) GetImpactOk() (*string, bool)`

GetImpactOk returns a tuple with the Impact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImpact

`func (o *AdvisoryLexmarkAdvisory) SetImpact(v string)`

SetImpact sets Impact field to given value.

### HasImpact

`func (o *AdvisoryLexmarkAdvisory) HasImpact() bool`

HasImpact returns a boolean if a field has been set.

### GetLastUpdate

`func (o *AdvisoryLexmarkAdvisory) GetLastUpdate() string`

GetLastUpdate returns the LastUpdate field if non-nil, zero value otherwise.

### GetLastUpdateOk

`func (o *AdvisoryLexmarkAdvisory) GetLastUpdateOk() (*string, bool)`

GetLastUpdateOk returns a tuple with the LastUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdate

`func (o *AdvisoryLexmarkAdvisory) SetLastUpdate(v string)`

SetLastUpdate sets LastUpdate field to given value.

### HasLastUpdate

`func (o *AdvisoryLexmarkAdvisory) HasLastUpdate() bool`

HasLastUpdate returns a boolean if a field has been set.

### GetLink

`func (o *AdvisoryLexmarkAdvisory) GetLink() string`

GetLink returns the Link field if non-nil, zero value otherwise.

### GetLinkOk

`func (o *AdvisoryLexmarkAdvisory) GetLinkOk() (*string, bool)`

GetLinkOk returns a tuple with the Link field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLink

`func (o *AdvisoryLexmarkAdvisory) SetLink(v string)`

SetLink sets Link field to given value.

### HasLink

`func (o *AdvisoryLexmarkAdvisory) HasLink() bool`

HasLink returns a boolean if a field has been set.

### GetPublicReleaseDate

`func (o *AdvisoryLexmarkAdvisory) GetPublicReleaseDate() string`

GetPublicReleaseDate returns the PublicReleaseDate field if non-nil, zero value otherwise.

### GetPublicReleaseDateOk

`func (o *AdvisoryLexmarkAdvisory) GetPublicReleaseDateOk() (*string, bool)`

GetPublicReleaseDateOk returns a tuple with the PublicReleaseDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicReleaseDate

`func (o *AdvisoryLexmarkAdvisory) SetPublicReleaseDate(v string)`

SetPublicReleaseDate sets PublicReleaseDate field to given value.

### HasPublicReleaseDate

`func (o *AdvisoryLexmarkAdvisory) HasPublicReleaseDate() bool`

HasPublicReleaseDate returns a boolean if a field has been set.

### GetReferences

`func (o *AdvisoryLexmarkAdvisory) GetReferences() []string`

GetReferences returns the References field if non-nil, zero value otherwise.

### GetReferencesOk

`func (o *AdvisoryLexmarkAdvisory) GetReferencesOk() (*[]string, bool)`

GetReferencesOk returns a tuple with the References field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferences

`func (o *AdvisoryLexmarkAdvisory) SetReferences(v []string)`

SetReferences sets References field to given value.

### HasReferences

`func (o *AdvisoryLexmarkAdvisory) HasReferences() bool`

HasReferences returns a boolean if a field has been set.

### GetRevision

`func (o *AdvisoryLexmarkAdvisory) GetRevision() string`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *AdvisoryLexmarkAdvisory) GetRevisionOk() (*string, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *AdvisoryLexmarkAdvisory) SetRevision(v string)`

SetRevision sets Revision field to given value.

### HasRevision

`func (o *AdvisoryLexmarkAdvisory) HasRevision() bool`

HasRevision returns a boolean if a field has been set.

### GetSummary

`func (o *AdvisoryLexmarkAdvisory) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *AdvisoryLexmarkAdvisory) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *AdvisoryLexmarkAdvisory) SetSummary(v string)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *AdvisoryLexmarkAdvisory) HasSummary() bool`

HasSummary returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *AdvisoryLexmarkAdvisory) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *AdvisoryLexmarkAdvisory) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *AdvisoryLexmarkAdvisory) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *AdvisoryLexmarkAdvisory) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetWorkarounds

`func (o *AdvisoryLexmarkAdvisory) GetWorkarounds() string`

GetWorkarounds returns the Workarounds field if non-nil, zero value otherwise.

### GetWorkaroundsOk

`func (o *AdvisoryLexmarkAdvisory) GetWorkaroundsOk() (*string, bool)`

GetWorkaroundsOk returns a tuple with the Workarounds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkarounds

`func (o *AdvisoryLexmarkAdvisory) SetWorkarounds(v string)`

SetWorkarounds sets Workarounds field to given value.

### HasWorkarounds

`func (o *AdvisoryLexmarkAdvisory) HasWorkarounds() bool`

HasWorkarounds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


