# AdvisoryVulnCheck

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Affecting** | Pointer to **[]string** |  | [optional] 
**Credit** | Pointer to **[]string** |  | [optional] 
**Cve** | Pointer to **[]string** |  | [optional] 
**Cvss** | Pointer to **string** |  | [optional] 
**CvssV3Vector** | Pointer to **string** |  | [optional] 
**DateAdded** | Pointer to **string** |  | [optional] 
**References** | Pointer to **[]string** |  | [optional] 
**Severity** | Pointer to **string** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 

## Methods

### NewAdvisoryVulnCheck

`func NewAdvisoryVulnCheck() *AdvisoryVulnCheck`

NewAdvisoryVulnCheck instantiates a new AdvisoryVulnCheck object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdvisoryVulnCheckWithDefaults

`func NewAdvisoryVulnCheckWithDefaults() *AdvisoryVulnCheck`

NewAdvisoryVulnCheckWithDefaults instantiates a new AdvisoryVulnCheck object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAffecting

`func (o *AdvisoryVulnCheck) GetAffecting() []string`

GetAffecting returns the Affecting field if non-nil, zero value otherwise.

### GetAffectingOk

`func (o *AdvisoryVulnCheck) GetAffectingOk() (*[]string, bool)`

GetAffectingOk returns a tuple with the Affecting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffecting

`func (o *AdvisoryVulnCheck) SetAffecting(v []string)`

SetAffecting sets Affecting field to given value.

### HasAffecting

`func (o *AdvisoryVulnCheck) HasAffecting() bool`

HasAffecting returns a boolean if a field has been set.

### GetCredit

`func (o *AdvisoryVulnCheck) GetCredit() []string`

GetCredit returns the Credit field if non-nil, zero value otherwise.

### GetCreditOk

`func (o *AdvisoryVulnCheck) GetCreditOk() (*[]string, bool)`

GetCreditOk returns a tuple with the Credit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredit

`func (o *AdvisoryVulnCheck) SetCredit(v []string)`

SetCredit sets Credit field to given value.

### HasCredit

`func (o *AdvisoryVulnCheck) HasCredit() bool`

HasCredit returns a boolean if a field has been set.

### GetCve

`func (o *AdvisoryVulnCheck) GetCve() []string`

GetCve returns the Cve field if non-nil, zero value otherwise.

### GetCveOk

`func (o *AdvisoryVulnCheck) GetCveOk() (*[]string, bool)`

GetCveOk returns a tuple with the Cve field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCve

`func (o *AdvisoryVulnCheck) SetCve(v []string)`

SetCve sets Cve field to given value.

### HasCve

`func (o *AdvisoryVulnCheck) HasCve() bool`

HasCve returns a boolean if a field has been set.

### GetCvss

`func (o *AdvisoryVulnCheck) GetCvss() string`

GetCvss returns the Cvss field if non-nil, zero value otherwise.

### GetCvssOk

`func (o *AdvisoryVulnCheck) GetCvssOk() (*string, bool)`

GetCvssOk returns a tuple with the Cvss field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCvss

`func (o *AdvisoryVulnCheck) SetCvss(v string)`

SetCvss sets Cvss field to given value.

### HasCvss

`func (o *AdvisoryVulnCheck) HasCvss() bool`

HasCvss returns a boolean if a field has been set.

### GetCvssV3Vector

`func (o *AdvisoryVulnCheck) GetCvssV3Vector() string`

GetCvssV3Vector returns the CvssV3Vector field if non-nil, zero value otherwise.

### GetCvssV3VectorOk

`func (o *AdvisoryVulnCheck) GetCvssV3VectorOk() (*string, bool)`

GetCvssV3VectorOk returns a tuple with the CvssV3Vector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCvssV3Vector

`func (o *AdvisoryVulnCheck) SetCvssV3Vector(v string)`

SetCvssV3Vector sets CvssV3Vector field to given value.

### HasCvssV3Vector

`func (o *AdvisoryVulnCheck) HasCvssV3Vector() bool`

HasCvssV3Vector returns a boolean if a field has been set.

### GetDateAdded

`func (o *AdvisoryVulnCheck) GetDateAdded() string`

GetDateAdded returns the DateAdded field if non-nil, zero value otherwise.

### GetDateAddedOk

`func (o *AdvisoryVulnCheck) GetDateAddedOk() (*string, bool)`

GetDateAddedOk returns a tuple with the DateAdded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateAdded

`func (o *AdvisoryVulnCheck) SetDateAdded(v string)`

SetDateAdded sets DateAdded field to given value.

### HasDateAdded

`func (o *AdvisoryVulnCheck) HasDateAdded() bool`

HasDateAdded returns a boolean if a field has been set.

### GetReferences

`func (o *AdvisoryVulnCheck) GetReferences() []string`

GetReferences returns the References field if non-nil, zero value otherwise.

### GetReferencesOk

`func (o *AdvisoryVulnCheck) GetReferencesOk() (*[]string, bool)`

GetReferencesOk returns a tuple with the References field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferences

`func (o *AdvisoryVulnCheck) SetReferences(v []string)`

SetReferences sets References field to given value.

### HasReferences

`func (o *AdvisoryVulnCheck) HasReferences() bool`

HasReferences returns a boolean if a field has been set.

### GetSeverity

`func (o *AdvisoryVulnCheck) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *AdvisoryVulnCheck) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *AdvisoryVulnCheck) SetSeverity(v string)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *AdvisoryVulnCheck) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetTitle

`func (o *AdvisoryVulnCheck) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *AdvisoryVulnCheck) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *AdvisoryVulnCheck) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *AdvisoryVulnCheck) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetType

`func (o *AdvisoryVulnCheck) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AdvisoryVulnCheck) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AdvisoryVulnCheck) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *AdvisoryVulnCheck) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUrl

`func (o *AdvisoryVulnCheck) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *AdvisoryVulnCheck) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *AdvisoryVulnCheck) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *AdvisoryVulnCheck) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


