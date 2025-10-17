# AdvisoryWatchGuard

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdvisoryId** | Pointer to **string** |  | [optional] 
**Affected** | Pointer to **string** |  | [optional] 
**Cve** | Pointer to **[]string** |  | [optional] 
**DateAdded** | Pointer to **string** |  | [optional] 
**Resolution** | Pointer to **string** |  | [optional] 
**Score** | Pointer to **string** |  | [optional] 
**Summary** | Pointer to **string** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 

## Methods

### NewAdvisoryWatchGuard

`func NewAdvisoryWatchGuard() *AdvisoryWatchGuard`

NewAdvisoryWatchGuard instantiates a new AdvisoryWatchGuard object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdvisoryWatchGuardWithDefaults

`func NewAdvisoryWatchGuardWithDefaults() *AdvisoryWatchGuard`

NewAdvisoryWatchGuardWithDefaults instantiates a new AdvisoryWatchGuard object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdvisoryId

`func (o *AdvisoryWatchGuard) GetAdvisoryId() string`

GetAdvisoryId returns the AdvisoryId field if non-nil, zero value otherwise.

### GetAdvisoryIdOk

`func (o *AdvisoryWatchGuard) GetAdvisoryIdOk() (*string, bool)`

GetAdvisoryIdOk returns a tuple with the AdvisoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvisoryId

`func (o *AdvisoryWatchGuard) SetAdvisoryId(v string)`

SetAdvisoryId sets AdvisoryId field to given value.

### HasAdvisoryId

`func (o *AdvisoryWatchGuard) HasAdvisoryId() bool`

HasAdvisoryId returns a boolean if a field has been set.

### GetAffected

`func (o *AdvisoryWatchGuard) GetAffected() string`

GetAffected returns the Affected field if non-nil, zero value otherwise.

### GetAffectedOk

`func (o *AdvisoryWatchGuard) GetAffectedOk() (*string, bool)`

GetAffectedOk returns a tuple with the Affected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffected

`func (o *AdvisoryWatchGuard) SetAffected(v string)`

SetAffected sets Affected field to given value.

### HasAffected

`func (o *AdvisoryWatchGuard) HasAffected() bool`

HasAffected returns a boolean if a field has been set.

### GetCve

`func (o *AdvisoryWatchGuard) GetCve() []string`

GetCve returns the Cve field if non-nil, zero value otherwise.

### GetCveOk

`func (o *AdvisoryWatchGuard) GetCveOk() (*[]string, bool)`

GetCveOk returns a tuple with the Cve field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCve

`func (o *AdvisoryWatchGuard) SetCve(v []string)`

SetCve sets Cve field to given value.

### HasCve

`func (o *AdvisoryWatchGuard) HasCve() bool`

HasCve returns a boolean if a field has been set.

### GetDateAdded

`func (o *AdvisoryWatchGuard) GetDateAdded() string`

GetDateAdded returns the DateAdded field if non-nil, zero value otherwise.

### GetDateAddedOk

`func (o *AdvisoryWatchGuard) GetDateAddedOk() (*string, bool)`

GetDateAddedOk returns a tuple with the DateAdded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateAdded

`func (o *AdvisoryWatchGuard) SetDateAdded(v string)`

SetDateAdded sets DateAdded field to given value.

### HasDateAdded

`func (o *AdvisoryWatchGuard) HasDateAdded() bool`

HasDateAdded returns a boolean if a field has been set.

### GetResolution

`func (o *AdvisoryWatchGuard) GetResolution() string`

GetResolution returns the Resolution field if non-nil, zero value otherwise.

### GetResolutionOk

`func (o *AdvisoryWatchGuard) GetResolutionOk() (*string, bool)`

GetResolutionOk returns a tuple with the Resolution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolution

`func (o *AdvisoryWatchGuard) SetResolution(v string)`

SetResolution sets Resolution field to given value.

### HasResolution

`func (o *AdvisoryWatchGuard) HasResolution() bool`

HasResolution returns a boolean if a field has been set.

### GetScore

`func (o *AdvisoryWatchGuard) GetScore() string`

GetScore returns the Score field if non-nil, zero value otherwise.

### GetScoreOk

`func (o *AdvisoryWatchGuard) GetScoreOk() (*string, bool)`

GetScoreOk returns a tuple with the Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScore

`func (o *AdvisoryWatchGuard) SetScore(v string)`

SetScore sets Score field to given value.

### HasScore

`func (o *AdvisoryWatchGuard) HasScore() bool`

HasScore returns a boolean if a field has been set.

### GetSummary

`func (o *AdvisoryWatchGuard) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *AdvisoryWatchGuard) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *AdvisoryWatchGuard) SetSummary(v string)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *AdvisoryWatchGuard) HasSummary() bool`

HasSummary returns a boolean if a field has been set.

### GetTitle

`func (o *AdvisoryWatchGuard) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *AdvisoryWatchGuard) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *AdvisoryWatchGuard) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *AdvisoryWatchGuard) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetUrl

`func (o *AdvisoryWatchGuard) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *AdvisoryWatchGuard) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *AdvisoryWatchGuard) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *AdvisoryWatchGuard) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


