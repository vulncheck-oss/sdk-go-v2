# AdvisoryEUVD

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Aliases** | Pointer to **[]string** |  | [optional] 
**Assigner** | Pointer to **string** |  | [optional] 
**BaseScore** | Pointer to **float32** |  | [optional] 
**BaseScoreVector** | Pointer to **string** |  | [optional] 
**BaseScoreVersion** | Pointer to **string** |  | [optional] 
**Cve** | Pointer to **[]string** |  | [optional] 
**DateAdded** | Pointer to **string** |  | [optional] 
**DateUpdated** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**EnisaIdProduct** | Pointer to [**[]AdvisoryEnisaIDProduct**](AdvisoryEnisaIDProduct.md) |  | [optional] 
**EnisaIdVendor** | Pointer to [**[]AdvisoryEnisaIDVendor**](AdvisoryEnisaIDVendor.md) |  | [optional] 
**Epss** | Pointer to **float32** |  | [optional] 
**Exploited** | Pointer to **bool** | This field is exploited field from endpoint /api/vulnerabilities. apidocs : https://euvd.enisa.europa.eu/apidoc Note: There are records where exploited_since is populated with a valid date, but it still shows up under non_exploitable data set | [optional] 
**ExploitedSince** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**References** | Pointer to **[]string** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 

## Methods

### NewAdvisoryEUVD

`func NewAdvisoryEUVD() *AdvisoryEUVD`

NewAdvisoryEUVD instantiates a new AdvisoryEUVD object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdvisoryEUVDWithDefaults

`func NewAdvisoryEUVDWithDefaults() *AdvisoryEUVD`

NewAdvisoryEUVDWithDefaults instantiates a new AdvisoryEUVD object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAliases

`func (o *AdvisoryEUVD) GetAliases() []string`

GetAliases returns the Aliases field if non-nil, zero value otherwise.

### GetAliasesOk

`func (o *AdvisoryEUVD) GetAliasesOk() (*[]string, bool)`

GetAliasesOk returns a tuple with the Aliases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAliases

`func (o *AdvisoryEUVD) SetAliases(v []string)`

SetAliases sets Aliases field to given value.

### HasAliases

`func (o *AdvisoryEUVD) HasAliases() bool`

HasAliases returns a boolean if a field has been set.

### GetAssigner

`func (o *AdvisoryEUVD) GetAssigner() string`

GetAssigner returns the Assigner field if non-nil, zero value otherwise.

### GetAssignerOk

`func (o *AdvisoryEUVD) GetAssignerOk() (*string, bool)`

GetAssignerOk returns a tuple with the Assigner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssigner

`func (o *AdvisoryEUVD) SetAssigner(v string)`

SetAssigner sets Assigner field to given value.

### HasAssigner

`func (o *AdvisoryEUVD) HasAssigner() bool`

HasAssigner returns a boolean if a field has been set.

### GetBaseScore

`func (o *AdvisoryEUVD) GetBaseScore() float32`

GetBaseScore returns the BaseScore field if non-nil, zero value otherwise.

### GetBaseScoreOk

`func (o *AdvisoryEUVD) GetBaseScoreOk() (*float32, bool)`

GetBaseScoreOk returns a tuple with the BaseScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseScore

`func (o *AdvisoryEUVD) SetBaseScore(v float32)`

SetBaseScore sets BaseScore field to given value.

### HasBaseScore

`func (o *AdvisoryEUVD) HasBaseScore() bool`

HasBaseScore returns a boolean if a field has been set.

### GetBaseScoreVector

`func (o *AdvisoryEUVD) GetBaseScoreVector() string`

GetBaseScoreVector returns the BaseScoreVector field if non-nil, zero value otherwise.

### GetBaseScoreVectorOk

`func (o *AdvisoryEUVD) GetBaseScoreVectorOk() (*string, bool)`

GetBaseScoreVectorOk returns a tuple with the BaseScoreVector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseScoreVector

`func (o *AdvisoryEUVD) SetBaseScoreVector(v string)`

SetBaseScoreVector sets BaseScoreVector field to given value.

### HasBaseScoreVector

`func (o *AdvisoryEUVD) HasBaseScoreVector() bool`

HasBaseScoreVector returns a boolean if a field has been set.

### GetBaseScoreVersion

`func (o *AdvisoryEUVD) GetBaseScoreVersion() string`

GetBaseScoreVersion returns the BaseScoreVersion field if non-nil, zero value otherwise.

### GetBaseScoreVersionOk

`func (o *AdvisoryEUVD) GetBaseScoreVersionOk() (*string, bool)`

GetBaseScoreVersionOk returns a tuple with the BaseScoreVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseScoreVersion

`func (o *AdvisoryEUVD) SetBaseScoreVersion(v string)`

SetBaseScoreVersion sets BaseScoreVersion field to given value.

### HasBaseScoreVersion

`func (o *AdvisoryEUVD) HasBaseScoreVersion() bool`

HasBaseScoreVersion returns a boolean if a field has been set.

### GetCve

`func (o *AdvisoryEUVD) GetCve() []string`

GetCve returns the Cve field if non-nil, zero value otherwise.

### GetCveOk

`func (o *AdvisoryEUVD) GetCveOk() (*[]string, bool)`

GetCveOk returns a tuple with the Cve field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCve

`func (o *AdvisoryEUVD) SetCve(v []string)`

SetCve sets Cve field to given value.

### HasCve

`func (o *AdvisoryEUVD) HasCve() bool`

HasCve returns a boolean if a field has been set.

### GetDateAdded

`func (o *AdvisoryEUVD) GetDateAdded() string`

GetDateAdded returns the DateAdded field if non-nil, zero value otherwise.

### GetDateAddedOk

`func (o *AdvisoryEUVD) GetDateAddedOk() (*string, bool)`

GetDateAddedOk returns a tuple with the DateAdded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateAdded

`func (o *AdvisoryEUVD) SetDateAdded(v string)`

SetDateAdded sets DateAdded field to given value.

### HasDateAdded

`func (o *AdvisoryEUVD) HasDateAdded() bool`

HasDateAdded returns a boolean if a field has been set.

### GetDateUpdated

`func (o *AdvisoryEUVD) GetDateUpdated() string`

GetDateUpdated returns the DateUpdated field if non-nil, zero value otherwise.

### GetDateUpdatedOk

`func (o *AdvisoryEUVD) GetDateUpdatedOk() (*string, bool)`

GetDateUpdatedOk returns a tuple with the DateUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateUpdated

`func (o *AdvisoryEUVD) SetDateUpdated(v string)`

SetDateUpdated sets DateUpdated field to given value.

### HasDateUpdated

`func (o *AdvisoryEUVD) HasDateUpdated() bool`

HasDateUpdated returns a boolean if a field has been set.

### GetDescription

`func (o *AdvisoryEUVD) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AdvisoryEUVD) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AdvisoryEUVD) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AdvisoryEUVD) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnisaIdProduct

`func (o *AdvisoryEUVD) GetEnisaIdProduct() []AdvisoryEnisaIDProduct`

GetEnisaIdProduct returns the EnisaIdProduct field if non-nil, zero value otherwise.

### GetEnisaIdProductOk

`func (o *AdvisoryEUVD) GetEnisaIdProductOk() (*[]AdvisoryEnisaIDProduct, bool)`

GetEnisaIdProductOk returns a tuple with the EnisaIdProduct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnisaIdProduct

`func (o *AdvisoryEUVD) SetEnisaIdProduct(v []AdvisoryEnisaIDProduct)`

SetEnisaIdProduct sets EnisaIdProduct field to given value.

### HasEnisaIdProduct

`func (o *AdvisoryEUVD) HasEnisaIdProduct() bool`

HasEnisaIdProduct returns a boolean if a field has been set.

### GetEnisaIdVendor

`func (o *AdvisoryEUVD) GetEnisaIdVendor() []AdvisoryEnisaIDVendor`

GetEnisaIdVendor returns the EnisaIdVendor field if non-nil, zero value otherwise.

### GetEnisaIdVendorOk

`func (o *AdvisoryEUVD) GetEnisaIdVendorOk() (*[]AdvisoryEnisaIDVendor, bool)`

GetEnisaIdVendorOk returns a tuple with the EnisaIdVendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnisaIdVendor

`func (o *AdvisoryEUVD) SetEnisaIdVendor(v []AdvisoryEnisaIDVendor)`

SetEnisaIdVendor sets EnisaIdVendor field to given value.

### HasEnisaIdVendor

`func (o *AdvisoryEUVD) HasEnisaIdVendor() bool`

HasEnisaIdVendor returns a boolean if a field has been set.

### GetEpss

`func (o *AdvisoryEUVD) GetEpss() float32`

GetEpss returns the Epss field if non-nil, zero value otherwise.

### GetEpssOk

`func (o *AdvisoryEUVD) GetEpssOk() (*float32, bool)`

GetEpssOk returns a tuple with the Epss field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEpss

`func (o *AdvisoryEUVD) SetEpss(v float32)`

SetEpss sets Epss field to given value.

### HasEpss

`func (o *AdvisoryEUVD) HasEpss() bool`

HasEpss returns a boolean if a field has been set.

### GetExploited

`func (o *AdvisoryEUVD) GetExploited() bool`

GetExploited returns the Exploited field if non-nil, zero value otherwise.

### GetExploitedOk

`func (o *AdvisoryEUVD) GetExploitedOk() (*bool, bool)`

GetExploitedOk returns a tuple with the Exploited field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExploited

`func (o *AdvisoryEUVD) SetExploited(v bool)`

SetExploited sets Exploited field to given value.

### HasExploited

`func (o *AdvisoryEUVD) HasExploited() bool`

HasExploited returns a boolean if a field has been set.

### GetExploitedSince

`func (o *AdvisoryEUVD) GetExploitedSince() string`

GetExploitedSince returns the ExploitedSince field if non-nil, zero value otherwise.

### GetExploitedSinceOk

`func (o *AdvisoryEUVD) GetExploitedSinceOk() (*string, bool)`

GetExploitedSinceOk returns a tuple with the ExploitedSince field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExploitedSince

`func (o *AdvisoryEUVD) SetExploitedSince(v string)`

SetExploitedSince sets ExploitedSince field to given value.

### HasExploitedSince

`func (o *AdvisoryEUVD) HasExploitedSince() bool`

HasExploitedSince returns a boolean if a field has been set.

### GetId

`func (o *AdvisoryEUVD) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AdvisoryEUVD) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AdvisoryEUVD) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AdvisoryEUVD) HasId() bool`

HasId returns a boolean if a field has been set.

### GetReferences

`func (o *AdvisoryEUVD) GetReferences() []string`

GetReferences returns the References field if non-nil, zero value otherwise.

### GetReferencesOk

`func (o *AdvisoryEUVD) GetReferencesOk() (*[]string, bool)`

GetReferencesOk returns a tuple with the References field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferences

`func (o *AdvisoryEUVD) SetReferences(v []string)`

SetReferences sets References field to given value.

### HasReferences

`func (o *AdvisoryEUVD) HasReferences() bool`

HasReferences returns a boolean if a field has been set.

### GetUrl

`func (o *AdvisoryEUVD) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *AdvisoryEUVD) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *AdvisoryEUVD) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *AdvisoryEUVD) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


