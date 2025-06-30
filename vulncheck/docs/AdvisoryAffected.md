# AdvisoryAffected

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DatabaseSpecific** | Pointer to **map[string]interface{}** | The meaning of the values within the object is entirely defined by the database | [optional] 
**EcosystemSpecific** | Pointer to **map[string]interface{}** | The meaning of the values within the object is entirely defined by the ecosystem | [optional] 
**Package** | Pointer to [**AdvisoryOSVPackage**](AdvisoryOSVPackage.md) |  | [optional] 
**Ranges** | Pointer to [**[]AdvisoryRange**](AdvisoryRange.md) |  | [optional] 
**Severity** | Pointer to [**[]AdvisorySeverity**](AdvisorySeverity.md) |  | [optional] 
**Versions** | Pointer to **[]string** |  | [optional] 

## Methods

### NewAdvisoryAffected

`func NewAdvisoryAffected() *AdvisoryAffected`

NewAdvisoryAffected instantiates a new AdvisoryAffected object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdvisoryAffectedWithDefaults

`func NewAdvisoryAffectedWithDefaults() *AdvisoryAffected`

NewAdvisoryAffectedWithDefaults instantiates a new AdvisoryAffected object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatabaseSpecific

`func (o *AdvisoryAffected) GetDatabaseSpecific() map[string]interface{}`

GetDatabaseSpecific returns the DatabaseSpecific field if non-nil, zero value otherwise.

### GetDatabaseSpecificOk

`func (o *AdvisoryAffected) GetDatabaseSpecificOk() (*map[string]interface{}, bool)`

GetDatabaseSpecificOk returns a tuple with the DatabaseSpecific field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseSpecific

`func (o *AdvisoryAffected) SetDatabaseSpecific(v map[string]interface{})`

SetDatabaseSpecific sets DatabaseSpecific field to given value.

### HasDatabaseSpecific

`func (o *AdvisoryAffected) HasDatabaseSpecific() bool`

HasDatabaseSpecific returns a boolean if a field has been set.

### GetEcosystemSpecific

`func (o *AdvisoryAffected) GetEcosystemSpecific() map[string]interface{}`

GetEcosystemSpecific returns the EcosystemSpecific field if non-nil, zero value otherwise.

### GetEcosystemSpecificOk

`func (o *AdvisoryAffected) GetEcosystemSpecificOk() (*map[string]interface{}, bool)`

GetEcosystemSpecificOk returns a tuple with the EcosystemSpecific field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEcosystemSpecific

`func (o *AdvisoryAffected) SetEcosystemSpecific(v map[string]interface{})`

SetEcosystemSpecific sets EcosystemSpecific field to given value.

### HasEcosystemSpecific

`func (o *AdvisoryAffected) HasEcosystemSpecific() bool`

HasEcosystemSpecific returns a boolean if a field has been set.

### GetPackage

`func (o *AdvisoryAffected) GetPackage() AdvisoryOSVPackage`

GetPackage returns the Package field if non-nil, zero value otherwise.

### GetPackageOk

`func (o *AdvisoryAffected) GetPackageOk() (*AdvisoryOSVPackage, bool)`

GetPackageOk returns a tuple with the Package field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackage

`func (o *AdvisoryAffected) SetPackage(v AdvisoryOSVPackage)`

SetPackage sets Package field to given value.

### HasPackage

`func (o *AdvisoryAffected) HasPackage() bool`

HasPackage returns a boolean if a field has been set.

### GetRanges

`func (o *AdvisoryAffected) GetRanges() []AdvisoryRange`

GetRanges returns the Ranges field if non-nil, zero value otherwise.

### GetRangesOk

`func (o *AdvisoryAffected) GetRangesOk() (*[]AdvisoryRange, bool)`

GetRangesOk returns a tuple with the Ranges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRanges

`func (o *AdvisoryAffected) SetRanges(v []AdvisoryRange)`

SetRanges sets Ranges field to given value.

### HasRanges

`func (o *AdvisoryAffected) HasRanges() bool`

HasRanges returns a boolean if a field has been set.

### GetSeverity

`func (o *AdvisoryAffected) GetSeverity() []AdvisorySeverity`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *AdvisoryAffected) GetSeverityOk() (*[]AdvisorySeverity, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *AdvisoryAffected) SetSeverity(v []AdvisorySeverity)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *AdvisoryAffected) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetVersions

`func (o *AdvisoryAffected) GetVersions() []string`

GetVersions returns the Versions field if non-nil, zero value otherwise.

### GetVersionsOk

`func (o *AdvisoryAffected) GetVersionsOk() (*[]string, bool)`

GetVersionsOk returns a tuple with the Versions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersions

`func (o *AdvisoryAffected) SetVersions(v []string)`

SetVersions sets Versions field to given value.

### HasVersions

`func (o *AdvisoryAffected) HasVersions() bool`

HasVersions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


