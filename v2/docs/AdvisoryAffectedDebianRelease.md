# AdvisoryAffectedDebianRelease

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FixedVersion** | Pointer to **string** |  | [optional] 
**Nodsa** | Pointer to **string** |  | [optional] 
**NodsaReason** | Pointer to **string** |  | [optional] 
**ReleaseName** | Pointer to **string** |  | [optional] 
**Repositories** | Pointer to [**[]AdvisoryAffectedDebianRepository**](AdvisoryAffectedDebianRepository.md) |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Urgency** | Pointer to **string** |  | [optional] 

## Methods

### NewAdvisoryAffectedDebianRelease

`func NewAdvisoryAffectedDebianRelease() *AdvisoryAffectedDebianRelease`

NewAdvisoryAffectedDebianRelease instantiates a new AdvisoryAffectedDebianRelease object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdvisoryAffectedDebianReleaseWithDefaults

`func NewAdvisoryAffectedDebianReleaseWithDefaults() *AdvisoryAffectedDebianRelease`

NewAdvisoryAffectedDebianReleaseWithDefaults instantiates a new AdvisoryAffectedDebianRelease object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFixedVersion

`func (o *AdvisoryAffectedDebianRelease) GetFixedVersion() string`

GetFixedVersion returns the FixedVersion field if non-nil, zero value otherwise.

### GetFixedVersionOk

`func (o *AdvisoryAffectedDebianRelease) GetFixedVersionOk() (*string, bool)`

GetFixedVersionOk returns a tuple with the FixedVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFixedVersion

`func (o *AdvisoryAffectedDebianRelease) SetFixedVersion(v string)`

SetFixedVersion sets FixedVersion field to given value.

### HasFixedVersion

`func (o *AdvisoryAffectedDebianRelease) HasFixedVersion() bool`

HasFixedVersion returns a boolean if a field has been set.

### GetNodsa

`func (o *AdvisoryAffectedDebianRelease) GetNodsa() string`

GetNodsa returns the Nodsa field if non-nil, zero value otherwise.

### GetNodsaOk

`func (o *AdvisoryAffectedDebianRelease) GetNodsaOk() (*string, bool)`

GetNodsaOk returns a tuple with the Nodsa field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodsa

`func (o *AdvisoryAffectedDebianRelease) SetNodsa(v string)`

SetNodsa sets Nodsa field to given value.

### HasNodsa

`func (o *AdvisoryAffectedDebianRelease) HasNodsa() bool`

HasNodsa returns a boolean if a field has been set.

### GetNodsaReason

`func (o *AdvisoryAffectedDebianRelease) GetNodsaReason() string`

GetNodsaReason returns the NodsaReason field if non-nil, zero value otherwise.

### GetNodsaReasonOk

`func (o *AdvisoryAffectedDebianRelease) GetNodsaReasonOk() (*string, bool)`

GetNodsaReasonOk returns a tuple with the NodsaReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodsaReason

`func (o *AdvisoryAffectedDebianRelease) SetNodsaReason(v string)`

SetNodsaReason sets NodsaReason field to given value.

### HasNodsaReason

`func (o *AdvisoryAffectedDebianRelease) HasNodsaReason() bool`

HasNodsaReason returns a boolean if a field has been set.

### GetReleaseName

`func (o *AdvisoryAffectedDebianRelease) GetReleaseName() string`

GetReleaseName returns the ReleaseName field if non-nil, zero value otherwise.

### GetReleaseNameOk

`func (o *AdvisoryAffectedDebianRelease) GetReleaseNameOk() (*string, bool)`

GetReleaseNameOk returns a tuple with the ReleaseName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseName

`func (o *AdvisoryAffectedDebianRelease) SetReleaseName(v string)`

SetReleaseName sets ReleaseName field to given value.

### HasReleaseName

`func (o *AdvisoryAffectedDebianRelease) HasReleaseName() bool`

HasReleaseName returns a boolean if a field has been set.

### GetRepositories

`func (o *AdvisoryAffectedDebianRelease) GetRepositories() []AdvisoryAffectedDebianRepository`

GetRepositories returns the Repositories field if non-nil, zero value otherwise.

### GetRepositoriesOk

`func (o *AdvisoryAffectedDebianRelease) GetRepositoriesOk() (*[]AdvisoryAffectedDebianRepository, bool)`

GetRepositoriesOk returns a tuple with the Repositories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepositories

`func (o *AdvisoryAffectedDebianRelease) SetRepositories(v []AdvisoryAffectedDebianRepository)`

SetRepositories sets Repositories field to given value.

### HasRepositories

`func (o *AdvisoryAffectedDebianRelease) HasRepositories() bool`

HasRepositories returns a boolean if a field has been set.

### GetStatus

`func (o *AdvisoryAffectedDebianRelease) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AdvisoryAffectedDebianRelease) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AdvisoryAffectedDebianRelease) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AdvisoryAffectedDebianRelease) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUrgency

`func (o *AdvisoryAffectedDebianRelease) GetUrgency() string`

GetUrgency returns the Urgency field if non-nil, zero value otherwise.

### GetUrgencyOk

`func (o *AdvisoryAffectedDebianRelease) GetUrgencyOk() (*string, bool)`

GetUrgencyOk returns a tuple with the Urgency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrgency

`func (o *AdvisoryAffectedDebianRelease) SetUrgency(v string)`

SetUrgency sets Urgency field to given value.

### HasUrgency

`func (o *AdvisoryAffectedDebianRelease) HasUrgency() bool`

HasUrgency returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


