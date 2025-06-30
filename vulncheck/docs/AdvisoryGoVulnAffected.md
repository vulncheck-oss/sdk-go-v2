# AdvisoryGoVulnAffected

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DatabaseSpecific** | Pointer to [**AdvisoryGoVulnDatabaseSpecific**](AdvisoryGoVulnDatabaseSpecific.md) |  | [optional] 
**EcosystemSpecific** | Pointer to [**AdvisoryGoVulnEcosystemSpecific**](AdvisoryGoVulnEcosystemSpecific.md) |  | [optional] 
**Package** | Pointer to [**AdvisoryGoVulnPackage**](AdvisoryGoVulnPackage.md) |  | [optional] 
**Ranges** | Pointer to [**[]AdvisoryGoVulnRanges**](AdvisoryGoVulnRanges.md) |  | [optional] 

## Methods

### NewAdvisoryGoVulnAffected

`func NewAdvisoryGoVulnAffected() *AdvisoryGoVulnAffected`

NewAdvisoryGoVulnAffected instantiates a new AdvisoryGoVulnAffected object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdvisoryGoVulnAffectedWithDefaults

`func NewAdvisoryGoVulnAffectedWithDefaults() *AdvisoryGoVulnAffected`

NewAdvisoryGoVulnAffectedWithDefaults instantiates a new AdvisoryGoVulnAffected object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatabaseSpecific

`func (o *AdvisoryGoVulnAffected) GetDatabaseSpecific() AdvisoryGoVulnDatabaseSpecific`

GetDatabaseSpecific returns the DatabaseSpecific field if non-nil, zero value otherwise.

### GetDatabaseSpecificOk

`func (o *AdvisoryGoVulnAffected) GetDatabaseSpecificOk() (*AdvisoryGoVulnDatabaseSpecific, bool)`

GetDatabaseSpecificOk returns a tuple with the DatabaseSpecific field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseSpecific

`func (o *AdvisoryGoVulnAffected) SetDatabaseSpecific(v AdvisoryGoVulnDatabaseSpecific)`

SetDatabaseSpecific sets DatabaseSpecific field to given value.

### HasDatabaseSpecific

`func (o *AdvisoryGoVulnAffected) HasDatabaseSpecific() bool`

HasDatabaseSpecific returns a boolean if a field has been set.

### GetEcosystemSpecific

`func (o *AdvisoryGoVulnAffected) GetEcosystemSpecific() AdvisoryGoVulnEcosystemSpecific`

GetEcosystemSpecific returns the EcosystemSpecific field if non-nil, zero value otherwise.

### GetEcosystemSpecificOk

`func (o *AdvisoryGoVulnAffected) GetEcosystemSpecificOk() (*AdvisoryGoVulnEcosystemSpecific, bool)`

GetEcosystemSpecificOk returns a tuple with the EcosystemSpecific field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEcosystemSpecific

`func (o *AdvisoryGoVulnAffected) SetEcosystemSpecific(v AdvisoryGoVulnEcosystemSpecific)`

SetEcosystemSpecific sets EcosystemSpecific field to given value.

### HasEcosystemSpecific

`func (o *AdvisoryGoVulnAffected) HasEcosystemSpecific() bool`

HasEcosystemSpecific returns a boolean if a field has been set.

### GetPackage

`func (o *AdvisoryGoVulnAffected) GetPackage() AdvisoryGoVulnPackage`

GetPackage returns the Package field if non-nil, zero value otherwise.

### GetPackageOk

`func (o *AdvisoryGoVulnAffected) GetPackageOk() (*AdvisoryGoVulnPackage, bool)`

GetPackageOk returns a tuple with the Package field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackage

`func (o *AdvisoryGoVulnAffected) SetPackage(v AdvisoryGoVulnPackage)`

SetPackage sets Package field to given value.

### HasPackage

`func (o *AdvisoryGoVulnAffected) HasPackage() bool`

HasPackage returns a boolean if a field has been set.

### GetRanges

`func (o *AdvisoryGoVulnAffected) GetRanges() []AdvisoryGoVulnRanges`

GetRanges returns the Ranges field if non-nil, zero value otherwise.

### GetRangesOk

`func (o *AdvisoryGoVulnAffected) GetRangesOk() (*[]AdvisoryGoVulnRanges, bool)`

GetRangesOk returns a tuple with the Ranges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRanges

`func (o *AdvisoryGoVulnAffected) SetRanges(v []AdvisoryGoVulnRanges)`

SetRanges sets Ranges field to given value.

### HasRanges

`func (o *AdvisoryGoVulnAffected) HasRanges() bool`

HasRanges returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


