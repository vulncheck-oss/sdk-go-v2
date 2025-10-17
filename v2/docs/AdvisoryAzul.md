# AdvisoryAzul

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BaseScore** | Pointer to **string** |  | [optional] 
**Cve** | Pointer to **[]string** |  | [optional] 
**DateAdded** | Pointer to **string** |  | [optional] 
**PrimeVersion** | Pointer to [**[]AdvisoryPrimeVersion**](AdvisoryPrimeVersion.md) |  | [optional] 
**Release** | Pointer to **string** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 
**ZuluVersion** | Pointer to [**[]AdvisoryZuluVersion**](AdvisoryZuluVersion.md) |  | [optional] 

## Methods

### NewAdvisoryAzul

`func NewAdvisoryAzul() *AdvisoryAzul`

NewAdvisoryAzul instantiates a new AdvisoryAzul object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdvisoryAzulWithDefaults

`func NewAdvisoryAzulWithDefaults() *AdvisoryAzul`

NewAdvisoryAzulWithDefaults instantiates a new AdvisoryAzul object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBaseScore

`func (o *AdvisoryAzul) GetBaseScore() string`

GetBaseScore returns the BaseScore field if non-nil, zero value otherwise.

### GetBaseScoreOk

`func (o *AdvisoryAzul) GetBaseScoreOk() (*string, bool)`

GetBaseScoreOk returns a tuple with the BaseScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseScore

`func (o *AdvisoryAzul) SetBaseScore(v string)`

SetBaseScore sets BaseScore field to given value.

### HasBaseScore

`func (o *AdvisoryAzul) HasBaseScore() bool`

HasBaseScore returns a boolean if a field has been set.

### GetCve

`func (o *AdvisoryAzul) GetCve() []string`

GetCve returns the Cve field if non-nil, zero value otherwise.

### GetCveOk

`func (o *AdvisoryAzul) GetCveOk() (*[]string, bool)`

GetCveOk returns a tuple with the Cve field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCve

`func (o *AdvisoryAzul) SetCve(v []string)`

SetCve sets Cve field to given value.

### HasCve

`func (o *AdvisoryAzul) HasCve() bool`

HasCve returns a boolean if a field has been set.

### GetDateAdded

`func (o *AdvisoryAzul) GetDateAdded() string`

GetDateAdded returns the DateAdded field if non-nil, zero value otherwise.

### GetDateAddedOk

`func (o *AdvisoryAzul) GetDateAddedOk() (*string, bool)`

GetDateAddedOk returns a tuple with the DateAdded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateAdded

`func (o *AdvisoryAzul) SetDateAdded(v string)`

SetDateAdded sets DateAdded field to given value.

### HasDateAdded

`func (o *AdvisoryAzul) HasDateAdded() bool`

HasDateAdded returns a boolean if a field has been set.

### GetPrimeVersion

`func (o *AdvisoryAzul) GetPrimeVersion() []AdvisoryPrimeVersion`

GetPrimeVersion returns the PrimeVersion field if non-nil, zero value otherwise.

### GetPrimeVersionOk

`func (o *AdvisoryAzul) GetPrimeVersionOk() (*[]AdvisoryPrimeVersion, bool)`

GetPrimeVersionOk returns a tuple with the PrimeVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimeVersion

`func (o *AdvisoryAzul) SetPrimeVersion(v []AdvisoryPrimeVersion)`

SetPrimeVersion sets PrimeVersion field to given value.

### HasPrimeVersion

`func (o *AdvisoryAzul) HasPrimeVersion() bool`

HasPrimeVersion returns a boolean if a field has been set.

### GetRelease

`func (o *AdvisoryAzul) GetRelease() string`

GetRelease returns the Release field if non-nil, zero value otherwise.

### GetReleaseOk

`func (o *AdvisoryAzul) GetReleaseOk() (*string, bool)`

GetReleaseOk returns a tuple with the Release field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelease

`func (o *AdvisoryAzul) SetRelease(v string)`

SetRelease sets Release field to given value.

### HasRelease

`func (o *AdvisoryAzul) HasRelease() bool`

HasRelease returns a boolean if a field has been set.

### GetUrl

`func (o *AdvisoryAzul) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *AdvisoryAzul) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *AdvisoryAzul) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *AdvisoryAzul) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetZuluVersion

`func (o *AdvisoryAzul) GetZuluVersion() []AdvisoryZuluVersion`

GetZuluVersion returns the ZuluVersion field if non-nil, zero value otherwise.

### GetZuluVersionOk

`func (o *AdvisoryAzul) GetZuluVersionOk() (*[]AdvisoryZuluVersion, bool)`

GetZuluVersionOk returns a tuple with the ZuluVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZuluVersion

`func (o *AdvisoryAzul) SetZuluVersion(v []AdvisoryZuluVersion)`

SetZuluVersion sets ZuluVersion field to given value.

### HasZuluVersion

`func (o *AdvisoryAzul) HasZuluVersion() bool`

HasZuluVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


