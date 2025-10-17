# AdvisoryCycle

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Codename** | Pointer to **string** |  | [optional] 
**Cycle** | Pointer to **string** |  | [optional] 
**Discontinued** | Pointer to **map[string]interface{}** |  | [optional] 
**Eol** | Pointer to **map[string]interface{}** |  | [optional] 
**ExtendedSupport** | Pointer to **map[string]interface{}** |  | [optional] 
**Latest** | Pointer to **string** |  | [optional] 
**LatestReleaseDate** | Pointer to **string** |  | [optional] 
**Link** | Pointer to **string** |  | [optional] 
**Lts** | Pointer to **map[string]interface{}** |  | [optional] 
**ReleaseDate** | Pointer to **string** |  | [optional] 
**ReleaseLabel** | Pointer to **string** |  | [optional] 
**Support** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewAdvisoryCycle

`func NewAdvisoryCycle() *AdvisoryCycle`

NewAdvisoryCycle instantiates a new AdvisoryCycle object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdvisoryCycleWithDefaults

`func NewAdvisoryCycleWithDefaults() *AdvisoryCycle`

NewAdvisoryCycleWithDefaults instantiates a new AdvisoryCycle object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCodename

`func (o *AdvisoryCycle) GetCodename() string`

GetCodename returns the Codename field if non-nil, zero value otherwise.

### GetCodenameOk

`func (o *AdvisoryCycle) GetCodenameOk() (*string, bool)`

GetCodenameOk returns a tuple with the Codename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodename

`func (o *AdvisoryCycle) SetCodename(v string)`

SetCodename sets Codename field to given value.

### HasCodename

`func (o *AdvisoryCycle) HasCodename() bool`

HasCodename returns a boolean if a field has been set.

### GetCycle

`func (o *AdvisoryCycle) GetCycle() string`

GetCycle returns the Cycle field if non-nil, zero value otherwise.

### GetCycleOk

`func (o *AdvisoryCycle) GetCycleOk() (*string, bool)`

GetCycleOk returns a tuple with the Cycle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCycle

`func (o *AdvisoryCycle) SetCycle(v string)`

SetCycle sets Cycle field to given value.

### HasCycle

`func (o *AdvisoryCycle) HasCycle() bool`

HasCycle returns a boolean if a field has been set.

### GetDiscontinued

`func (o *AdvisoryCycle) GetDiscontinued() map[string]interface{}`

GetDiscontinued returns the Discontinued field if non-nil, zero value otherwise.

### GetDiscontinuedOk

`func (o *AdvisoryCycle) GetDiscontinuedOk() (*map[string]interface{}, bool)`

GetDiscontinuedOk returns a tuple with the Discontinued field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscontinued

`func (o *AdvisoryCycle) SetDiscontinued(v map[string]interface{})`

SetDiscontinued sets Discontinued field to given value.

### HasDiscontinued

`func (o *AdvisoryCycle) HasDiscontinued() bool`

HasDiscontinued returns a boolean if a field has been set.

### GetEol

`func (o *AdvisoryCycle) GetEol() map[string]interface{}`

GetEol returns the Eol field if non-nil, zero value otherwise.

### GetEolOk

`func (o *AdvisoryCycle) GetEolOk() (*map[string]interface{}, bool)`

GetEolOk returns a tuple with the Eol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEol

`func (o *AdvisoryCycle) SetEol(v map[string]interface{})`

SetEol sets Eol field to given value.

### HasEol

`func (o *AdvisoryCycle) HasEol() bool`

HasEol returns a boolean if a field has been set.

### GetExtendedSupport

`func (o *AdvisoryCycle) GetExtendedSupport() map[string]interface{}`

GetExtendedSupport returns the ExtendedSupport field if non-nil, zero value otherwise.

### GetExtendedSupportOk

`func (o *AdvisoryCycle) GetExtendedSupportOk() (*map[string]interface{}, bool)`

GetExtendedSupportOk returns a tuple with the ExtendedSupport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtendedSupport

`func (o *AdvisoryCycle) SetExtendedSupport(v map[string]interface{})`

SetExtendedSupport sets ExtendedSupport field to given value.

### HasExtendedSupport

`func (o *AdvisoryCycle) HasExtendedSupport() bool`

HasExtendedSupport returns a boolean if a field has been set.

### GetLatest

`func (o *AdvisoryCycle) GetLatest() string`

GetLatest returns the Latest field if non-nil, zero value otherwise.

### GetLatestOk

`func (o *AdvisoryCycle) GetLatestOk() (*string, bool)`

GetLatestOk returns a tuple with the Latest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatest

`func (o *AdvisoryCycle) SetLatest(v string)`

SetLatest sets Latest field to given value.

### HasLatest

`func (o *AdvisoryCycle) HasLatest() bool`

HasLatest returns a boolean if a field has been set.

### GetLatestReleaseDate

`func (o *AdvisoryCycle) GetLatestReleaseDate() string`

GetLatestReleaseDate returns the LatestReleaseDate field if non-nil, zero value otherwise.

### GetLatestReleaseDateOk

`func (o *AdvisoryCycle) GetLatestReleaseDateOk() (*string, bool)`

GetLatestReleaseDateOk returns a tuple with the LatestReleaseDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestReleaseDate

`func (o *AdvisoryCycle) SetLatestReleaseDate(v string)`

SetLatestReleaseDate sets LatestReleaseDate field to given value.

### HasLatestReleaseDate

`func (o *AdvisoryCycle) HasLatestReleaseDate() bool`

HasLatestReleaseDate returns a boolean if a field has been set.

### GetLink

`func (o *AdvisoryCycle) GetLink() string`

GetLink returns the Link field if non-nil, zero value otherwise.

### GetLinkOk

`func (o *AdvisoryCycle) GetLinkOk() (*string, bool)`

GetLinkOk returns a tuple with the Link field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLink

`func (o *AdvisoryCycle) SetLink(v string)`

SetLink sets Link field to given value.

### HasLink

`func (o *AdvisoryCycle) HasLink() bool`

HasLink returns a boolean if a field has been set.

### GetLts

`func (o *AdvisoryCycle) GetLts() map[string]interface{}`

GetLts returns the Lts field if non-nil, zero value otherwise.

### GetLtsOk

`func (o *AdvisoryCycle) GetLtsOk() (*map[string]interface{}, bool)`

GetLtsOk returns a tuple with the Lts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLts

`func (o *AdvisoryCycle) SetLts(v map[string]interface{})`

SetLts sets Lts field to given value.

### HasLts

`func (o *AdvisoryCycle) HasLts() bool`

HasLts returns a boolean if a field has been set.

### GetReleaseDate

`func (o *AdvisoryCycle) GetReleaseDate() string`

GetReleaseDate returns the ReleaseDate field if non-nil, zero value otherwise.

### GetReleaseDateOk

`func (o *AdvisoryCycle) GetReleaseDateOk() (*string, bool)`

GetReleaseDateOk returns a tuple with the ReleaseDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseDate

`func (o *AdvisoryCycle) SetReleaseDate(v string)`

SetReleaseDate sets ReleaseDate field to given value.

### HasReleaseDate

`func (o *AdvisoryCycle) HasReleaseDate() bool`

HasReleaseDate returns a boolean if a field has been set.

### GetReleaseLabel

`func (o *AdvisoryCycle) GetReleaseLabel() string`

GetReleaseLabel returns the ReleaseLabel field if non-nil, zero value otherwise.

### GetReleaseLabelOk

`func (o *AdvisoryCycle) GetReleaseLabelOk() (*string, bool)`

GetReleaseLabelOk returns a tuple with the ReleaseLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseLabel

`func (o *AdvisoryCycle) SetReleaseLabel(v string)`

SetReleaseLabel sets ReleaseLabel field to given value.

### HasReleaseLabel

`func (o *AdvisoryCycle) HasReleaseLabel() bool`

HasReleaseLabel returns a boolean if a field has been set.

### GetSupport

`func (o *AdvisoryCycle) GetSupport() map[string]interface{}`

GetSupport returns the Support field if non-nil, zero value otherwise.

### GetSupportOk

`func (o *AdvisoryCycle) GetSupportOk() (*map[string]interface{}, bool)`

GetSupportOk returns a tuple with the Support field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupport

`func (o *AdvisoryCycle) SetSupport(v map[string]interface{})`

SetSupport sets Support field to given value.

### HasSupport

`func (o *AdvisoryCycle) HasSupport() bool`

HasSupport returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


