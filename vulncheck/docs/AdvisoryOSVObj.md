# AdvisoryOSVObj

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Affected** | Pointer to [**[]AdvisoryAffected**](AdvisoryAffected.md) | collection based on https://ossf.github.io/osv-schema/ | [optional] 
**Aliases** | Pointer to **[]string** |  | [optional] 
**Details** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Modified** | Pointer to **string** |  | [optional] 
**Published** | Pointer to **string** |  | [optional] 
**References** | Pointer to [**[]AdvisoryOSVReference**](AdvisoryOSVReference.md) |  | [optional] 
**Related** | Pointer to **[]string** |  | [optional] 
**Summary** | Pointer to **string** |  | [optional] 
**Withdrawn** | Pointer to **string** |  | [optional] 

## Methods

### NewAdvisoryOSVObj

`func NewAdvisoryOSVObj() *AdvisoryOSVObj`

NewAdvisoryOSVObj instantiates a new AdvisoryOSVObj object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdvisoryOSVObjWithDefaults

`func NewAdvisoryOSVObjWithDefaults() *AdvisoryOSVObj`

NewAdvisoryOSVObjWithDefaults instantiates a new AdvisoryOSVObj object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAffected

`func (o *AdvisoryOSVObj) GetAffected() []AdvisoryAffected`

GetAffected returns the Affected field if non-nil, zero value otherwise.

### GetAffectedOk

`func (o *AdvisoryOSVObj) GetAffectedOk() (*[]AdvisoryAffected, bool)`

GetAffectedOk returns a tuple with the Affected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffected

`func (o *AdvisoryOSVObj) SetAffected(v []AdvisoryAffected)`

SetAffected sets Affected field to given value.

### HasAffected

`func (o *AdvisoryOSVObj) HasAffected() bool`

HasAffected returns a boolean if a field has been set.

### GetAliases

`func (o *AdvisoryOSVObj) GetAliases() []string`

GetAliases returns the Aliases field if non-nil, zero value otherwise.

### GetAliasesOk

`func (o *AdvisoryOSVObj) GetAliasesOk() (*[]string, bool)`

GetAliasesOk returns a tuple with the Aliases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAliases

`func (o *AdvisoryOSVObj) SetAliases(v []string)`

SetAliases sets Aliases field to given value.

### HasAliases

`func (o *AdvisoryOSVObj) HasAliases() bool`

HasAliases returns a boolean if a field has been set.

### GetDetails

`func (o *AdvisoryOSVObj) GetDetails() string`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *AdvisoryOSVObj) GetDetailsOk() (*string, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *AdvisoryOSVObj) SetDetails(v string)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *AdvisoryOSVObj) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### GetId

`func (o *AdvisoryOSVObj) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AdvisoryOSVObj) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AdvisoryOSVObj) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AdvisoryOSVObj) HasId() bool`

HasId returns a boolean if a field has been set.

### GetModified

`func (o *AdvisoryOSVObj) GetModified() string`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *AdvisoryOSVObj) GetModifiedOk() (*string, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *AdvisoryOSVObj) SetModified(v string)`

SetModified sets Modified field to given value.

### HasModified

`func (o *AdvisoryOSVObj) HasModified() bool`

HasModified returns a boolean if a field has been set.

### GetPublished

`func (o *AdvisoryOSVObj) GetPublished() string`

GetPublished returns the Published field if non-nil, zero value otherwise.

### GetPublishedOk

`func (o *AdvisoryOSVObj) GetPublishedOk() (*string, bool)`

GetPublishedOk returns a tuple with the Published field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublished

`func (o *AdvisoryOSVObj) SetPublished(v string)`

SetPublished sets Published field to given value.

### HasPublished

`func (o *AdvisoryOSVObj) HasPublished() bool`

HasPublished returns a boolean if a field has been set.

### GetReferences

`func (o *AdvisoryOSVObj) GetReferences() []AdvisoryOSVReference`

GetReferences returns the References field if non-nil, zero value otherwise.

### GetReferencesOk

`func (o *AdvisoryOSVObj) GetReferencesOk() (*[]AdvisoryOSVReference, bool)`

GetReferencesOk returns a tuple with the References field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferences

`func (o *AdvisoryOSVObj) SetReferences(v []AdvisoryOSVReference)`

SetReferences sets References field to given value.

### HasReferences

`func (o *AdvisoryOSVObj) HasReferences() bool`

HasReferences returns a boolean if a field has been set.

### GetRelated

`func (o *AdvisoryOSVObj) GetRelated() []string`

GetRelated returns the Related field if non-nil, zero value otherwise.

### GetRelatedOk

`func (o *AdvisoryOSVObj) GetRelatedOk() (*[]string, bool)`

GetRelatedOk returns a tuple with the Related field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelated

`func (o *AdvisoryOSVObj) SetRelated(v []string)`

SetRelated sets Related field to given value.

### HasRelated

`func (o *AdvisoryOSVObj) HasRelated() bool`

HasRelated returns a boolean if a field has been set.

### GetSummary

`func (o *AdvisoryOSVObj) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *AdvisoryOSVObj) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *AdvisoryOSVObj) SetSummary(v string)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *AdvisoryOSVObj) HasSummary() bool`

HasSummary returns a boolean if a field has been set.

### GetWithdrawn

`func (o *AdvisoryOSVObj) GetWithdrawn() string`

GetWithdrawn returns the Withdrawn field if non-nil, zero value otherwise.

### GetWithdrawnOk

`func (o *AdvisoryOSVObj) GetWithdrawnOk() (*string, bool)`

GetWithdrawnOk returns a tuple with the Withdrawn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithdrawn

`func (o *AdvisoryOSVObj) SetWithdrawn(v string)`

SetWithdrawn sets Withdrawn field to given value.

### HasWithdrawn

`func (o *AdvisoryOSVObj) HasWithdrawn() bool`

HasWithdrawn returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


