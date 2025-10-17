# AdvisorySSDAdvisory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Analysis** | Pointer to **string** |  | [optional] 
**Credit** | Pointer to **string** |  | [optional] 
**Cve** | Pointer to **[]string** |  | [optional] 
**DateAdded** | Pointer to **string** |  | [optional] 
**Link** | Pointer to **string** |  | [optional] 
**Poc** | Pointer to **string** | contains actual poc code | [optional] 
**Published** | Pointer to **string** |  | [optional] 
**ResponseRef** | Pointer to **string** |  | [optional] 
**Summary** | Pointer to **string** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 

## Methods

### NewAdvisorySSDAdvisory

`func NewAdvisorySSDAdvisory() *AdvisorySSDAdvisory`

NewAdvisorySSDAdvisory instantiates a new AdvisorySSDAdvisory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdvisorySSDAdvisoryWithDefaults

`func NewAdvisorySSDAdvisoryWithDefaults() *AdvisorySSDAdvisory`

NewAdvisorySSDAdvisoryWithDefaults instantiates a new AdvisorySSDAdvisory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnalysis

`func (o *AdvisorySSDAdvisory) GetAnalysis() string`

GetAnalysis returns the Analysis field if non-nil, zero value otherwise.

### GetAnalysisOk

`func (o *AdvisorySSDAdvisory) GetAnalysisOk() (*string, bool)`

GetAnalysisOk returns a tuple with the Analysis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnalysis

`func (o *AdvisorySSDAdvisory) SetAnalysis(v string)`

SetAnalysis sets Analysis field to given value.

### HasAnalysis

`func (o *AdvisorySSDAdvisory) HasAnalysis() bool`

HasAnalysis returns a boolean if a field has been set.

### GetCredit

`func (o *AdvisorySSDAdvisory) GetCredit() string`

GetCredit returns the Credit field if non-nil, zero value otherwise.

### GetCreditOk

`func (o *AdvisorySSDAdvisory) GetCreditOk() (*string, bool)`

GetCreditOk returns a tuple with the Credit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredit

`func (o *AdvisorySSDAdvisory) SetCredit(v string)`

SetCredit sets Credit field to given value.

### HasCredit

`func (o *AdvisorySSDAdvisory) HasCredit() bool`

HasCredit returns a boolean if a field has been set.

### GetCve

`func (o *AdvisorySSDAdvisory) GetCve() []string`

GetCve returns the Cve field if non-nil, zero value otherwise.

### GetCveOk

`func (o *AdvisorySSDAdvisory) GetCveOk() (*[]string, bool)`

GetCveOk returns a tuple with the Cve field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCve

`func (o *AdvisorySSDAdvisory) SetCve(v []string)`

SetCve sets Cve field to given value.

### HasCve

`func (o *AdvisorySSDAdvisory) HasCve() bool`

HasCve returns a boolean if a field has been set.

### GetDateAdded

`func (o *AdvisorySSDAdvisory) GetDateAdded() string`

GetDateAdded returns the DateAdded field if non-nil, zero value otherwise.

### GetDateAddedOk

`func (o *AdvisorySSDAdvisory) GetDateAddedOk() (*string, bool)`

GetDateAddedOk returns a tuple with the DateAdded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateAdded

`func (o *AdvisorySSDAdvisory) SetDateAdded(v string)`

SetDateAdded sets DateAdded field to given value.

### HasDateAdded

`func (o *AdvisorySSDAdvisory) HasDateAdded() bool`

HasDateAdded returns a boolean if a field has been set.

### GetLink

`func (o *AdvisorySSDAdvisory) GetLink() string`

GetLink returns the Link field if non-nil, zero value otherwise.

### GetLinkOk

`func (o *AdvisorySSDAdvisory) GetLinkOk() (*string, bool)`

GetLinkOk returns a tuple with the Link field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLink

`func (o *AdvisorySSDAdvisory) SetLink(v string)`

SetLink sets Link field to given value.

### HasLink

`func (o *AdvisorySSDAdvisory) HasLink() bool`

HasLink returns a boolean if a field has been set.

### GetPoc

`func (o *AdvisorySSDAdvisory) GetPoc() string`

GetPoc returns the Poc field if non-nil, zero value otherwise.

### GetPocOk

`func (o *AdvisorySSDAdvisory) GetPocOk() (*string, bool)`

GetPocOk returns a tuple with the Poc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoc

`func (o *AdvisorySSDAdvisory) SetPoc(v string)`

SetPoc sets Poc field to given value.

### HasPoc

`func (o *AdvisorySSDAdvisory) HasPoc() bool`

HasPoc returns a boolean if a field has been set.

### GetPublished

`func (o *AdvisorySSDAdvisory) GetPublished() string`

GetPublished returns the Published field if non-nil, zero value otherwise.

### GetPublishedOk

`func (o *AdvisorySSDAdvisory) GetPublishedOk() (*string, bool)`

GetPublishedOk returns a tuple with the Published field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublished

`func (o *AdvisorySSDAdvisory) SetPublished(v string)`

SetPublished sets Published field to given value.

### HasPublished

`func (o *AdvisorySSDAdvisory) HasPublished() bool`

HasPublished returns a boolean if a field has been set.

### GetResponseRef

`func (o *AdvisorySSDAdvisory) GetResponseRef() string`

GetResponseRef returns the ResponseRef field if non-nil, zero value otherwise.

### GetResponseRefOk

`func (o *AdvisorySSDAdvisory) GetResponseRefOk() (*string, bool)`

GetResponseRefOk returns a tuple with the ResponseRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseRef

`func (o *AdvisorySSDAdvisory) SetResponseRef(v string)`

SetResponseRef sets ResponseRef field to given value.

### HasResponseRef

`func (o *AdvisorySSDAdvisory) HasResponseRef() bool`

HasResponseRef returns a boolean if a field has been set.

### GetSummary

`func (o *AdvisorySSDAdvisory) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *AdvisorySSDAdvisory) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *AdvisorySSDAdvisory) SetSummary(v string)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *AdvisorySSDAdvisory) HasSummary() bool`

HasSummary returns a boolean if a field has been set.

### GetTitle

`func (o *AdvisorySSDAdvisory) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *AdvisorySSDAdvisory) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *AdvisorySSDAdvisory) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *AdvisorySSDAdvisory) HasTitle() bool`

HasTitle returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


