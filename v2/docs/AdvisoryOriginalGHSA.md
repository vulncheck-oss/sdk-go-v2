# AdvisoryOriginalGHSA

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Affected** | Pointer to [**[]AdvisoryGHSAAffected**](AdvisoryGHSAAffected.md) |  | [optional] 
**Aliases** | Pointer to **[]string** |  | [optional] 
**DatabaseSpecific** | Pointer to [**AdvisoryGHSADatabaseSpecific**](AdvisoryGHSADatabaseSpecific.md) |  | [optional] 
**Details** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Modified** | Pointer to **string** |  | [optional] 
**Published** | Pointer to **string** |  | [optional] 
**References** | Pointer to [**[]AdvisoryGHSAReference**](AdvisoryGHSAReference.md) |  | [optional] 
**SchemaVersion** | Pointer to **string** |  | [optional] 
**Severity** | Pointer to [**[]AdvisoryGHSASeverity**](AdvisoryGHSASeverity.md) |  | [optional] 
**Summary** | Pointer to **string** |  | [optional] 

## Methods

### NewAdvisoryOriginalGHSA

`func NewAdvisoryOriginalGHSA() *AdvisoryOriginalGHSA`

NewAdvisoryOriginalGHSA instantiates a new AdvisoryOriginalGHSA object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdvisoryOriginalGHSAWithDefaults

`func NewAdvisoryOriginalGHSAWithDefaults() *AdvisoryOriginalGHSA`

NewAdvisoryOriginalGHSAWithDefaults instantiates a new AdvisoryOriginalGHSA object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAffected

`func (o *AdvisoryOriginalGHSA) GetAffected() []AdvisoryGHSAAffected`

GetAffected returns the Affected field if non-nil, zero value otherwise.

### GetAffectedOk

`func (o *AdvisoryOriginalGHSA) GetAffectedOk() (*[]AdvisoryGHSAAffected, bool)`

GetAffectedOk returns a tuple with the Affected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffected

`func (o *AdvisoryOriginalGHSA) SetAffected(v []AdvisoryGHSAAffected)`

SetAffected sets Affected field to given value.

### HasAffected

`func (o *AdvisoryOriginalGHSA) HasAffected() bool`

HasAffected returns a boolean if a field has been set.

### GetAliases

`func (o *AdvisoryOriginalGHSA) GetAliases() []string`

GetAliases returns the Aliases field if non-nil, zero value otherwise.

### GetAliasesOk

`func (o *AdvisoryOriginalGHSA) GetAliasesOk() (*[]string, bool)`

GetAliasesOk returns a tuple with the Aliases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAliases

`func (o *AdvisoryOriginalGHSA) SetAliases(v []string)`

SetAliases sets Aliases field to given value.

### HasAliases

`func (o *AdvisoryOriginalGHSA) HasAliases() bool`

HasAliases returns a boolean if a field has been set.

### GetDatabaseSpecific

`func (o *AdvisoryOriginalGHSA) GetDatabaseSpecific() AdvisoryGHSADatabaseSpecific`

GetDatabaseSpecific returns the DatabaseSpecific field if non-nil, zero value otherwise.

### GetDatabaseSpecificOk

`func (o *AdvisoryOriginalGHSA) GetDatabaseSpecificOk() (*AdvisoryGHSADatabaseSpecific, bool)`

GetDatabaseSpecificOk returns a tuple with the DatabaseSpecific field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseSpecific

`func (o *AdvisoryOriginalGHSA) SetDatabaseSpecific(v AdvisoryGHSADatabaseSpecific)`

SetDatabaseSpecific sets DatabaseSpecific field to given value.

### HasDatabaseSpecific

`func (o *AdvisoryOriginalGHSA) HasDatabaseSpecific() bool`

HasDatabaseSpecific returns a boolean if a field has been set.

### GetDetails

`func (o *AdvisoryOriginalGHSA) GetDetails() string`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *AdvisoryOriginalGHSA) GetDetailsOk() (*string, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *AdvisoryOriginalGHSA) SetDetails(v string)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *AdvisoryOriginalGHSA) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### GetId

`func (o *AdvisoryOriginalGHSA) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AdvisoryOriginalGHSA) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AdvisoryOriginalGHSA) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AdvisoryOriginalGHSA) HasId() bool`

HasId returns a boolean if a field has been set.

### GetModified

`func (o *AdvisoryOriginalGHSA) GetModified() string`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *AdvisoryOriginalGHSA) GetModifiedOk() (*string, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *AdvisoryOriginalGHSA) SetModified(v string)`

SetModified sets Modified field to given value.

### HasModified

`func (o *AdvisoryOriginalGHSA) HasModified() bool`

HasModified returns a boolean if a field has been set.

### GetPublished

`func (o *AdvisoryOriginalGHSA) GetPublished() string`

GetPublished returns the Published field if non-nil, zero value otherwise.

### GetPublishedOk

`func (o *AdvisoryOriginalGHSA) GetPublishedOk() (*string, bool)`

GetPublishedOk returns a tuple with the Published field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublished

`func (o *AdvisoryOriginalGHSA) SetPublished(v string)`

SetPublished sets Published field to given value.

### HasPublished

`func (o *AdvisoryOriginalGHSA) HasPublished() bool`

HasPublished returns a boolean if a field has been set.

### GetReferences

`func (o *AdvisoryOriginalGHSA) GetReferences() []AdvisoryGHSAReference`

GetReferences returns the References field if non-nil, zero value otherwise.

### GetReferencesOk

`func (o *AdvisoryOriginalGHSA) GetReferencesOk() (*[]AdvisoryGHSAReference, bool)`

GetReferencesOk returns a tuple with the References field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferences

`func (o *AdvisoryOriginalGHSA) SetReferences(v []AdvisoryGHSAReference)`

SetReferences sets References field to given value.

### HasReferences

`func (o *AdvisoryOriginalGHSA) HasReferences() bool`

HasReferences returns a boolean if a field has been set.

### GetSchemaVersion

`func (o *AdvisoryOriginalGHSA) GetSchemaVersion() string`

GetSchemaVersion returns the SchemaVersion field if non-nil, zero value otherwise.

### GetSchemaVersionOk

`func (o *AdvisoryOriginalGHSA) GetSchemaVersionOk() (*string, bool)`

GetSchemaVersionOk returns a tuple with the SchemaVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaVersion

`func (o *AdvisoryOriginalGHSA) SetSchemaVersion(v string)`

SetSchemaVersion sets SchemaVersion field to given value.

### HasSchemaVersion

`func (o *AdvisoryOriginalGHSA) HasSchemaVersion() bool`

HasSchemaVersion returns a boolean if a field has been set.

### GetSeverity

`func (o *AdvisoryOriginalGHSA) GetSeverity() []AdvisoryGHSASeverity`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *AdvisoryOriginalGHSA) GetSeverityOk() (*[]AdvisoryGHSASeverity, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *AdvisoryOriginalGHSA) SetSeverity(v []AdvisoryGHSASeverity)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *AdvisoryOriginalGHSA) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetSummary

`func (o *AdvisoryOriginalGHSA) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *AdvisoryOriginalGHSA) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *AdvisoryOriginalGHSA) SetSummary(v string)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *AdvisoryOriginalGHSA) HasSummary() bool`

HasSummary returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


