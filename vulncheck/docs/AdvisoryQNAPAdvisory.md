# AdvisoryQNAPAdvisory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Affected** | Pointer to **string** |  | [optional] 
**BulletinId** | Pointer to **string** |  | [optional] 
**Cve** | Pointer to **[]string** |  | [optional] 
**DateAdded** | Pointer to **string** |  | [optional] 
**LastUpdateDate** | Pointer to **string** |  | [optional] 
**Link** | Pointer to **string** |  | [optional] 
**Severity** | Pointer to **string** |  | [optional] 
**SeverityNumber** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Summary** | Pointer to **string** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 

## Methods

### NewAdvisoryQNAPAdvisory

`func NewAdvisoryQNAPAdvisory() *AdvisoryQNAPAdvisory`

NewAdvisoryQNAPAdvisory instantiates a new AdvisoryQNAPAdvisory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdvisoryQNAPAdvisoryWithDefaults

`func NewAdvisoryQNAPAdvisoryWithDefaults() *AdvisoryQNAPAdvisory`

NewAdvisoryQNAPAdvisoryWithDefaults instantiates a new AdvisoryQNAPAdvisory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAffected

`func (o *AdvisoryQNAPAdvisory) GetAffected() string`

GetAffected returns the Affected field if non-nil, zero value otherwise.

### GetAffectedOk

`func (o *AdvisoryQNAPAdvisory) GetAffectedOk() (*string, bool)`

GetAffectedOk returns a tuple with the Affected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffected

`func (o *AdvisoryQNAPAdvisory) SetAffected(v string)`

SetAffected sets Affected field to given value.

### HasAffected

`func (o *AdvisoryQNAPAdvisory) HasAffected() bool`

HasAffected returns a boolean if a field has been set.

### GetBulletinId

`func (o *AdvisoryQNAPAdvisory) GetBulletinId() string`

GetBulletinId returns the BulletinId field if non-nil, zero value otherwise.

### GetBulletinIdOk

`func (o *AdvisoryQNAPAdvisory) GetBulletinIdOk() (*string, bool)`

GetBulletinIdOk returns a tuple with the BulletinId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBulletinId

`func (o *AdvisoryQNAPAdvisory) SetBulletinId(v string)`

SetBulletinId sets BulletinId field to given value.

### HasBulletinId

`func (o *AdvisoryQNAPAdvisory) HasBulletinId() bool`

HasBulletinId returns a boolean if a field has been set.

### GetCve

`func (o *AdvisoryQNAPAdvisory) GetCve() []string`

GetCve returns the Cve field if non-nil, zero value otherwise.

### GetCveOk

`func (o *AdvisoryQNAPAdvisory) GetCveOk() (*[]string, bool)`

GetCveOk returns a tuple with the Cve field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCve

`func (o *AdvisoryQNAPAdvisory) SetCve(v []string)`

SetCve sets Cve field to given value.

### HasCve

`func (o *AdvisoryQNAPAdvisory) HasCve() bool`

HasCve returns a boolean if a field has been set.

### GetDateAdded

`func (o *AdvisoryQNAPAdvisory) GetDateAdded() string`

GetDateAdded returns the DateAdded field if non-nil, zero value otherwise.

### GetDateAddedOk

`func (o *AdvisoryQNAPAdvisory) GetDateAddedOk() (*string, bool)`

GetDateAddedOk returns a tuple with the DateAdded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateAdded

`func (o *AdvisoryQNAPAdvisory) SetDateAdded(v string)`

SetDateAdded sets DateAdded field to given value.

### HasDateAdded

`func (o *AdvisoryQNAPAdvisory) HasDateAdded() bool`

HasDateAdded returns a boolean if a field has been set.

### GetLastUpdateDate

`func (o *AdvisoryQNAPAdvisory) GetLastUpdateDate() string`

GetLastUpdateDate returns the LastUpdateDate field if non-nil, zero value otherwise.

### GetLastUpdateDateOk

`func (o *AdvisoryQNAPAdvisory) GetLastUpdateDateOk() (*string, bool)`

GetLastUpdateDateOk returns a tuple with the LastUpdateDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdateDate

`func (o *AdvisoryQNAPAdvisory) SetLastUpdateDate(v string)`

SetLastUpdateDate sets LastUpdateDate field to given value.

### HasLastUpdateDate

`func (o *AdvisoryQNAPAdvisory) HasLastUpdateDate() bool`

HasLastUpdateDate returns a boolean if a field has been set.

### GetLink

`func (o *AdvisoryQNAPAdvisory) GetLink() string`

GetLink returns the Link field if non-nil, zero value otherwise.

### GetLinkOk

`func (o *AdvisoryQNAPAdvisory) GetLinkOk() (*string, bool)`

GetLinkOk returns a tuple with the Link field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLink

`func (o *AdvisoryQNAPAdvisory) SetLink(v string)`

SetLink sets Link field to given value.

### HasLink

`func (o *AdvisoryQNAPAdvisory) HasLink() bool`

HasLink returns a boolean if a field has been set.

### GetSeverity

`func (o *AdvisoryQNAPAdvisory) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *AdvisoryQNAPAdvisory) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *AdvisoryQNAPAdvisory) SetSeverity(v string)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *AdvisoryQNAPAdvisory) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetSeverityNumber

`func (o *AdvisoryQNAPAdvisory) GetSeverityNumber() string`

GetSeverityNumber returns the SeverityNumber field if non-nil, zero value otherwise.

### GetSeverityNumberOk

`func (o *AdvisoryQNAPAdvisory) GetSeverityNumberOk() (*string, bool)`

GetSeverityNumberOk returns a tuple with the SeverityNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverityNumber

`func (o *AdvisoryQNAPAdvisory) SetSeverityNumber(v string)`

SetSeverityNumber sets SeverityNumber field to given value.

### HasSeverityNumber

`func (o *AdvisoryQNAPAdvisory) HasSeverityNumber() bool`

HasSeverityNumber returns a boolean if a field has been set.

### GetStatus

`func (o *AdvisoryQNAPAdvisory) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AdvisoryQNAPAdvisory) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AdvisoryQNAPAdvisory) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AdvisoryQNAPAdvisory) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSummary

`func (o *AdvisoryQNAPAdvisory) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *AdvisoryQNAPAdvisory) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *AdvisoryQNAPAdvisory) SetSummary(v string)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *AdvisoryQNAPAdvisory) HasSummary() bool`

HasSummary returns a boolean if a field has been set.

### GetTitle

`func (o *AdvisoryQNAPAdvisory) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *AdvisoryQNAPAdvisory) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *AdvisoryQNAPAdvisory) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *AdvisoryQNAPAdvisory) HasTitle() bool`

HasTitle returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


