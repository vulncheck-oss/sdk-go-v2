# AdvisoryFortinetAdvisory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Acknowledgement** | Pointer to **string** |  | [optional] 
**AffectedProducts** | Pointer to **[]string** |  | [optional] 
**Cve** | Pointer to **[]string** |  | [optional] 
**Cvssv3** | Pointer to **string** |  | [optional] 
**DateAdded** | Pointer to **string** |  | [optional] 
**Irnumber** | Pointer to **string** |  | [optional] 
**Link** | Pointer to **string** |  | [optional] 
**References** | Pointer to **[]string** |  | [optional] 
**Solutions** | Pointer to **[]string** |  | [optional] 
**Summary** | Pointer to **string** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 

## Methods

### NewAdvisoryFortinetAdvisory

`func NewAdvisoryFortinetAdvisory() *AdvisoryFortinetAdvisory`

NewAdvisoryFortinetAdvisory instantiates a new AdvisoryFortinetAdvisory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdvisoryFortinetAdvisoryWithDefaults

`func NewAdvisoryFortinetAdvisoryWithDefaults() *AdvisoryFortinetAdvisory`

NewAdvisoryFortinetAdvisoryWithDefaults instantiates a new AdvisoryFortinetAdvisory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAcknowledgement

`func (o *AdvisoryFortinetAdvisory) GetAcknowledgement() string`

GetAcknowledgement returns the Acknowledgement field if non-nil, zero value otherwise.

### GetAcknowledgementOk

`func (o *AdvisoryFortinetAdvisory) GetAcknowledgementOk() (*string, bool)`

GetAcknowledgementOk returns a tuple with the Acknowledgement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcknowledgement

`func (o *AdvisoryFortinetAdvisory) SetAcknowledgement(v string)`

SetAcknowledgement sets Acknowledgement field to given value.

### HasAcknowledgement

`func (o *AdvisoryFortinetAdvisory) HasAcknowledgement() bool`

HasAcknowledgement returns a boolean if a field has been set.

### GetAffectedProducts

`func (o *AdvisoryFortinetAdvisory) GetAffectedProducts() []string`

GetAffectedProducts returns the AffectedProducts field if non-nil, zero value otherwise.

### GetAffectedProductsOk

`func (o *AdvisoryFortinetAdvisory) GetAffectedProductsOk() (*[]string, bool)`

GetAffectedProductsOk returns a tuple with the AffectedProducts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffectedProducts

`func (o *AdvisoryFortinetAdvisory) SetAffectedProducts(v []string)`

SetAffectedProducts sets AffectedProducts field to given value.

### HasAffectedProducts

`func (o *AdvisoryFortinetAdvisory) HasAffectedProducts() bool`

HasAffectedProducts returns a boolean if a field has been set.

### GetCve

`func (o *AdvisoryFortinetAdvisory) GetCve() []string`

GetCve returns the Cve field if non-nil, zero value otherwise.

### GetCveOk

`func (o *AdvisoryFortinetAdvisory) GetCveOk() (*[]string, bool)`

GetCveOk returns a tuple with the Cve field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCve

`func (o *AdvisoryFortinetAdvisory) SetCve(v []string)`

SetCve sets Cve field to given value.

### HasCve

`func (o *AdvisoryFortinetAdvisory) HasCve() bool`

HasCve returns a boolean if a field has been set.

### GetCvssv3

`func (o *AdvisoryFortinetAdvisory) GetCvssv3() string`

GetCvssv3 returns the Cvssv3 field if non-nil, zero value otherwise.

### GetCvssv3Ok

`func (o *AdvisoryFortinetAdvisory) GetCvssv3Ok() (*string, bool)`

GetCvssv3Ok returns a tuple with the Cvssv3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCvssv3

`func (o *AdvisoryFortinetAdvisory) SetCvssv3(v string)`

SetCvssv3 sets Cvssv3 field to given value.

### HasCvssv3

`func (o *AdvisoryFortinetAdvisory) HasCvssv3() bool`

HasCvssv3 returns a boolean if a field has been set.

### GetDateAdded

`func (o *AdvisoryFortinetAdvisory) GetDateAdded() string`

GetDateAdded returns the DateAdded field if non-nil, zero value otherwise.

### GetDateAddedOk

`func (o *AdvisoryFortinetAdvisory) GetDateAddedOk() (*string, bool)`

GetDateAddedOk returns a tuple with the DateAdded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateAdded

`func (o *AdvisoryFortinetAdvisory) SetDateAdded(v string)`

SetDateAdded sets DateAdded field to given value.

### HasDateAdded

`func (o *AdvisoryFortinetAdvisory) HasDateAdded() bool`

HasDateAdded returns a boolean if a field has been set.

### GetIrnumber

`func (o *AdvisoryFortinetAdvisory) GetIrnumber() string`

GetIrnumber returns the Irnumber field if non-nil, zero value otherwise.

### GetIrnumberOk

`func (o *AdvisoryFortinetAdvisory) GetIrnumberOk() (*string, bool)`

GetIrnumberOk returns a tuple with the Irnumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIrnumber

`func (o *AdvisoryFortinetAdvisory) SetIrnumber(v string)`

SetIrnumber sets Irnumber field to given value.

### HasIrnumber

`func (o *AdvisoryFortinetAdvisory) HasIrnumber() bool`

HasIrnumber returns a boolean if a field has been set.

### GetLink

`func (o *AdvisoryFortinetAdvisory) GetLink() string`

GetLink returns the Link field if non-nil, zero value otherwise.

### GetLinkOk

`func (o *AdvisoryFortinetAdvisory) GetLinkOk() (*string, bool)`

GetLinkOk returns a tuple with the Link field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLink

`func (o *AdvisoryFortinetAdvisory) SetLink(v string)`

SetLink sets Link field to given value.

### HasLink

`func (o *AdvisoryFortinetAdvisory) HasLink() bool`

HasLink returns a boolean if a field has been set.

### GetReferences

`func (o *AdvisoryFortinetAdvisory) GetReferences() []string`

GetReferences returns the References field if non-nil, zero value otherwise.

### GetReferencesOk

`func (o *AdvisoryFortinetAdvisory) GetReferencesOk() (*[]string, bool)`

GetReferencesOk returns a tuple with the References field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferences

`func (o *AdvisoryFortinetAdvisory) SetReferences(v []string)`

SetReferences sets References field to given value.

### HasReferences

`func (o *AdvisoryFortinetAdvisory) HasReferences() bool`

HasReferences returns a boolean if a field has been set.

### GetSolutions

`func (o *AdvisoryFortinetAdvisory) GetSolutions() []string`

GetSolutions returns the Solutions field if non-nil, zero value otherwise.

### GetSolutionsOk

`func (o *AdvisoryFortinetAdvisory) GetSolutionsOk() (*[]string, bool)`

GetSolutionsOk returns a tuple with the Solutions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSolutions

`func (o *AdvisoryFortinetAdvisory) SetSolutions(v []string)`

SetSolutions sets Solutions field to given value.

### HasSolutions

`func (o *AdvisoryFortinetAdvisory) HasSolutions() bool`

HasSolutions returns a boolean if a field has been set.

### GetSummary

`func (o *AdvisoryFortinetAdvisory) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *AdvisoryFortinetAdvisory) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *AdvisoryFortinetAdvisory) SetSummary(v string)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *AdvisoryFortinetAdvisory) HasSummary() bool`

HasSummary returns a boolean if a field has been set.

### GetTitle

`func (o *AdvisoryFortinetAdvisory) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *AdvisoryFortinetAdvisory) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *AdvisoryFortinetAdvisory) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *AdvisoryFortinetAdvisory) HasTitle() bool`

HasTitle returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


