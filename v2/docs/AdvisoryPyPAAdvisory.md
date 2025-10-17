# AdvisoryPyPAAdvisory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdvisoryId** | Pointer to **string** | ID is the PYSEC- identifier | [optional] 
**Affected** | Pointer to [**[]AdvisoryPyPAAffected**](AdvisoryPyPAAffected.md) | Affected will list out the vulnerable versions. | [optional] 
**Aliases** | Pointer to **[]string** | Aliases are other identifiers that refer to this, such as a CVE | [optional] 
**Cve** | Pointer to **[]string** |  | [optional] 
**DateAdded** | Pointer to **string** |  | [optional] 
**Details** | Pointer to **string** | Details discuss the vulnerability information | [optional] 
**Modified** | Pointer to **string** | Modified is non-zero Time if entry was updated | [optional] 
**Published** | Pointer to **string** | Published is the datetime when this was released | [optional] 
**References** | Pointer to [**[]AdvisoryPyPAReference**](AdvisoryPyPAReference.md) | References are links to more detailed advisories, fixes, etc. | [optional] 
**WasWithdrawn** | Pointer to **bool** | WasWD indicates if the advisory was withdrawn or not. | [optional] 
**Withdrawn** | Pointer to **string** | Withdrawn is non-zero if this advisory has been withdrawn | [optional] 

## Methods

### NewAdvisoryPyPAAdvisory

`func NewAdvisoryPyPAAdvisory() *AdvisoryPyPAAdvisory`

NewAdvisoryPyPAAdvisory instantiates a new AdvisoryPyPAAdvisory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdvisoryPyPAAdvisoryWithDefaults

`func NewAdvisoryPyPAAdvisoryWithDefaults() *AdvisoryPyPAAdvisory`

NewAdvisoryPyPAAdvisoryWithDefaults instantiates a new AdvisoryPyPAAdvisory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdvisoryId

`func (o *AdvisoryPyPAAdvisory) GetAdvisoryId() string`

GetAdvisoryId returns the AdvisoryId field if non-nil, zero value otherwise.

### GetAdvisoryIdOk

`func (o *AdvisoryPyPAAdvisory) GetAdvisoryIdOk() (*string, bool)`

GetAdvisoryIdOk returns a tuple with the AdvisoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvisoryId

`func (o *AdvisoryPyPAAdvisory) SetAdvisoryId(v string)`

SetAdvisoryId sets AdvisoryId field to given value.

### HasAdvisoryId

`func (o *AdvisoryPyPAAdvisory) HasAdvisoryId() bool`

HasAdvisoryId returns a boolean if a field has been set.

### GetAffected

`func (o *AdvisoryPyPAAdvisory) GetAffected() []AdvisoryPyPAAffected`

GetAffected returns the Affected field if non-nil, zero value otherwise.

### GetAffectedOk

`func (o *AdvisoryPyPAAdvisory) GetAffectedOk() (*[]AdvisoryPyPAAffected, bool)`

GetAffectedOk returns a tuple with the Affected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffected

`func (o *AdvisoryPyPAAdvisory) SetAffected(v []AdvisoryPyPAAffected)`

SetAffected sets Affected field to given value.

### HasAffected

`func (o *AdvisoryPyPAAdvisory) HasAffected() bool`

HasAffected returns a boolean if a field has been set.

### GetAliases

`func (o *AdvisoryPyPAAdvisory) GetAliases() []string`

GetAliases returns the Aliases field if non-nil, zero value otherwise.

### GetAliasesOk

`func (o *AdvisoryPyPAAdvisory) GetAliasesOk() (*[]string, bool)`

GetAliasesOk returns a tuple with the Aliases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAliases

`func (o *AdvisoryPyPAAdvisory) SetAliases(v []string)`

SetAliases sets Aliases field to given value.

### HasAliases

`func (o *AdvisoryPyPAAdvisory) HasAliases() bool`

HasAliases returns a boolean if a field has been set.

### GetCve

`func (o *AdvisoryPyPAAdvisory) GetCve() []string`

GetCve returns the Cve field if non-nil, zero value otherwise.

### GetCveOk

`func (o *AdvisoryPyPAAdvisory) GetCveOk() (*[]string, bool)`

GetCveOk returns a tuple with the Cve field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCve

`func (o *AdvisoryPyPAAdvisory) SetCve(v []string)`

SetCve sets Cve field to given value.

### HasCve

`func (o *AdvisoryPyPAAdvisory) HasCve() bool`

HasCve returns a boolean if a field has been set.

### GetDateAdded

`func (o *AdvisoryPyPAAdvisory) GetDateAdded() string`

GetDateAdded returns the DateAdded field if non-nil, zero value otherwise.

### GetDateAddedOk

`func (o *AdvisoryPyPAAdvisory) GetDateAddedOk() (*string, bool)`

GetDateAddedOk returns a tuple with the DateAdded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateAdded

`func (o *AdvisoryPyPAAdvisory) SetDateAdded(v string)`

SetDateAdded sets DateAdded field to given value.

### HasDateAdded

`func (o *AdvisoryPyPAAdvisory) HasDateAdded() bool`

HasDateAdded returns a boolean if a field has been set.

### GetDetails

`func (o *AdvisoryPyPAAdvisory) GetDetails() string`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *AdvisoryPyPAAdvisory) GetDetailsOk() (*string, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *AdvisoryPyPAAdvisory) SetDetails(v string)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *AdvisoryPyPAAdvisory) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### GetModified

`func (o *AdvisoryPyPAAdvisory) GetModified() string`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *AdvisoryPyPAAdvisory) GetModifiedOk() (*string, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *AdvisoryPyPAAdvisory) SetModified(v string)`

SetModified sets Modified field to given value.

### HasModified

`func (o *AdvisoryPyPAAdvisory) HasModified() bool`

HasModified returns a boolean if a field has been set.

### GetPublished

`func (o *AdvisoryPyPAAdvisory) GetPublished() string`

GetPublished returns the Published field if non-nil, zero value otherwise.

### GetPublishedOk

`func (o *AdvisoryPyPAAdvisory) GetPublishedOk() (*string, bool)`

GetPublishedOk returns a tuple with the Published field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublished

`func (o *AdvisoryPyPAAdvisory) SetPublished(v string)`

SetPublished sets Published field to given value.

### HasPublished

`func (o *AdvisoryPyPAAdvisory) HasPublished() bool`

HasPublished returns a boolean if a field has been set.

### GetReferences

`func (o *AdvisoryPyPAAdvisory) GetReferences() []AdvisoryPyPAReference`

GetReferences returns the References field if non-nil, zero value otherwise.

### GetReferencesOk

`func (o *AdvisoryPyPAAdvisory) GetReferencesOk() (*[]AdvisoryPyPAReference, bool)`

GetReferencesOk returns a tuple with the References field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferences

`func (o *AdvisoryPyPAAdvisory) SetReferences(v []AdvisoryPyPAReference)`

SetReferences sets References field to given value.

### HasReferences

`func (o *AdvisoryPyPAAdvisory) HasReferences() bool`

HasReferences returns a boolean if a field has been set.

### GetWasWithdrawn

`func (o *AdvisoryPyPAAdvisory) GetWasWithdrawn() bool`

GetWasWithdrawn returns the WasWithdrawn field if non-nil, zero value otherwise.

### GetWasWithdrawnOk

`func (o *AdvisoryPyPAAdvisory) GetWasWithdrawnOk() (*bool, bool)`

GetWasWithdrawnOk returns a tuple with the WasWithdrawn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWasWithdrawn

`func (o *AdvisoryPyPAAdvisory) SetWasWithdrawn(v bool)`

SetWasWithdrawn sets WasWithdrawn field to given value.

### HasWasWithdrawn

`func (o *AdvisoryPyPAAdvisory) HasWasWithdrawn() bool`

HasWasWithdrawn returns a boolean if a field has been set.

### GetWithdrawn

`func (o *AdvisoryPyPAAdvisory) GetWithdrawn() string`

GetWithdrawn returns the Withdrawn field if non-nil, zero value otherwise.

### GetWithdrawnOk

`func (o *AdvisoryPyPAAdvisory) GetWithdrawnOk() (*string, bool)`

GetWithdrawnOk returns a tuple with the Withdrawn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithdrawn

`func (o *AdvisoryPyPAAdvisory) SetWithdrawn(v string)`

SetWithdrawn sets Withdrawn field to given value.

### HasWithdrawn

`func (o *AdvisoryPyPAAdvisory) HasWithdrawn() bool`

HasWithdrawn returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


