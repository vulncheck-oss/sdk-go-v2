# AdvisoryThreatActorWithExternalObjects

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Country** | Pointer to **string** |  | [optional] 
**CveReferences** | Pointer to [**[]AdvisoryCVEReference**](AdvisoryCVEReference.md) |  | [optional] 
**DateAdded** | Pointer to **string** |  | [optional] 
**MalpediaUrl** | Pointer to **string** |  | [optional] 
**MispId** | Pointer to **string** |  | [optional] 
**MispThreatActor** | Pointer to [**AdvisoryMISPValueNoID**](AdvisoryMISPValueNoID.md) |  | [optional] 
**MitreAttackGroup** | Pointer to [**AdvisoryMITREAttackGroupNoID**](AdvisoryMITREAttackGroupNoID.md) |  | [optional] 
**MitreGroupCti** | Pointer to [**AdvisoryMitreGroupCTI**](AdvisoryMitreGroupCTI.md) |  | [optional] 
**MitreId** | Pointer to **string** |  | [optional] 
**ThreatActorName** | Pointer to **string** |  | [optional] 
**Tools** | Pointer to [**[]AdvisoryTool**](AdvisoryTool.md) |  | [optional] 
**VendorNamesForThreatActors** | Pointer to [**[]AdvisoryVendorNameForThreatActor**](AdvisoryVendorNameForThreatActor.md) |  | [optional] 
**VendorsAndProductsTargeted** | Pointer to [**[]AdvisoryVendorProduct**](AdvisoryVendorProduct.md) |  | [optional] 

## Methods

### NewAdvisoryThreatActorWithExternalObjects

`func NewAdvisoryThreatActorWithExternalObjects() *AdvisoryThreatActorWithExternalObjects`

NewAdvisoryThreatActorWithExternalObjects instantiates a new AdvisoryThreatActorWithExternalObjects object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdvisoryThreatActorWithExternalObjectsWithDefaults

`func NewAdvisoryThreatActorWithExternalObjectsWithDefaults() *AdvisoryThreatActorWithExternalObjects`

NewAdvisoryThreatActorWithExternalObjectsWithDefaults instantiates a new AdvisoryThreatActorWithExternalObjects object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCountry

`func (o *AdvisoryThreatActorWithExternalObjects) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *AdvisoryThreatActorWithExternalObjects) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *AdvisoryThreatActorWithExternalObjects) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *AdvisoryThreatActorWithExternalObjects) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetCveReferences

`func (o *AdvisoryThreatActorWithExternalObjects) GetCveReferences() []AdvisoryCVEReference`

GetCveReferences returns the CveReferences field if non-nil, zero value otherwise.

### GetCveReferencesOk

`func (o *AdvisoryThreatActorWithExternalObjects) GetCveReferencesOk() (*[]AdvisoryCVEReference, bool)`

GetCveReferencesOk returns a tuple with the CveReferences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCveReferences

`func (o *AdvisoryThreatActorWithExternalObjects) SetCveReferences(v []AdvisoryCVEReference)`

SetCveReferences sets CveReferences field to given value.

### HasCveReferences

`func (o *AdvisoryThreatActorWithExternalObjects) HasCveReferences() bool`

HasCveReferences returns a boolean if a field has been set.

### GetDateAdded

`func (o *AdvisoryThreatActorWithExternalObjects) GetDateAdded() string`

GetDateAdded returns the DateAdded field if non-nil, zero value otherwise.

### GetDateAddedOk

`func (o *AdvisoryThreatActorWithExternalObjects) GetDateAddedOk() (*string, bool)`

GetDateAddedOk returns a tuple with the DateAdded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateAdded

`func (o *AdvisoryThreatActorWithExternalObjects) SetDateAdded(v string)`

SetDateAdded sets DateAdded field to given value.

### HasDateAdded

`func (o *AdvisoryThreatActorWithExternalObjects) HasDateAdded() bool`

HasDateAdded returns a boolean if a field has been set.

### GetMalpediaUrl

`func (o *AdvisoryThreatActorWithExternalObjects) GetMalpediaUrl() string`

GetMalpediaUrl returns the MalpediaUrl field if non-nil, zero value otherwise.

### GetMalpediaUrlOk

`func (o *AdvisoryThreatActorWithExternalObjects) GetMalpediaUrlOk() (*string, bool)`

GetMalpediaUrlOk returns a tuple with the MalpediaUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMalpediaUrl

`func (o *AdvisoryThreatActorWithExternalObjects) SetMalpediaUrl(v string)`

SetMalpediaUrl sets MalpediaUrl field to given value.

### HasMalpediaUrl

`func (o *AdvisoryThreatActorWithExternalObjects) HasMalpediaUrl() bool`

HasMalpediaUrl returns a boolean if a field has been set.

### GetMispId

`func (o *AdvisoryThreatActorWithExternalObjects) GetMispId() string`

GetMispId returns the MispId field if non-nil, zero value otherwise.

### GetMispIdOk

`func (o *AdvisoryThreatActorWithExternalObjects) GetMispIdOk() (*string, bool)`

GetMispIdOk returns a tuple with the MispId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMispId

`func (o *AdvisoryThreatActorWithExternalObjects) SetMispId(v string)`

SetMispId sets MispId field to given value.

### HasMispId

`func (o *AdvisoryThreatActorWithExternalObjects) HasMispId() bool`

HasMispId returns a boolean if a field has been set.

### GetMispThreatActor

`func (o *AdvisoryThreatActorWithExternalObjects) GetMispThreatActor() AdvisoryMISPValueNoID`

GetMispThreatActor returns the MispThreatActor field if non-nil, zero value otherwise.

### GetMispThreatActorOk

`func (o *AdvisoryThreatActorWithExternalObjects) GetMispThreatActorOk() (*AdvisoryMISPValueNoID, bool)`

GetMispThreatActorOk returns a tuple with the MispThreatActor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMispThreatActor

`func (o *AdvisoryThreatActorWithExternalObjects) SetMispThreatActor(v AdvisoryMISPValueNoID)`

SetMispThreatActor sets MispThreatActor field to given value.

### HasMispThreatActor

`func (o *AdvisoryThreatActorWithExternalObjects) HasMispThreatActor() bool`

HasMispThreatActor returns a boolean if a field has been set.

### GetMitreAttackGroup

`func (o *AdvisoryThreatActorWithExternalObjects) GetMitreAttackGroup() AdvisoryMITREAttackGroupNoID`

GetMitreAttackGroup returns the MitreAttackGroup field if non-nil, zero value otherwise.

### GetMitreAttackGroupOk

`func (o *AdvisoryThreatActorWithExternalObjects) GetMitreAttackGroupOk() (*AdvisoryMITREAttackGroupNoID, bool)`

GetMitreAttackGroupOk returns a tuple with the MitreAttackGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMitreAttackGroup

`func (o *AdvisoryThreatActorWithExternalObjects) SetMitreAttackGroup(v AdvisoryMITREAttackGroupNoID)`

SetMitreAttackGroup sets MitreAttackGroup field to given value.

### HasMitreAttackGroup

`func (o *AdvisoryThreatActorWithExternalObjects) HasMitreAttackGroup() bool`

HasMitreAttackGroup returns a boolean if a field has been set.

### GetMitreGroupCti

`func (o *AdvisoryThreatActorWithExternalObjects) GetMitreGroupCti() AdvisoryMitreGroupCTI`

GetMitreGroupCti returns the MitreGroupCti field if non-nil, zero value otherwise.

### GetMitreGroupCtiOk

`func (o *AdvisoryThreatActorWithExternalObjects) GetMitreGroupCtiOk() (*AdvisoryMitreGroupCTI, bool)`

GetMitreGroupCtiOk returns a tuple with the MitreGroupCti field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMitreGroupCti

`func (o *AdvisoryThreatActorWithExternalObjects) SetMitreGroupCti(v AdvisoryMitreGroupCTI)`

SetMitreGroupCti sets MitreGroupCti field to given value.

### HasMitreGroupCti

`func (o *AdvisoryThreatActorWithExternalObjects) HasMitreGroupCti() bool`

HasMitreGroupCti returns a boolean if a field has been set.

### GetMitreId

`func (o *AdvisoryThreatActorWithExternalObjects) GetMitreId() string`

GetMitreId returns the MitreId field if non-nil, zero value otherwise.

### GetMitreIdOk

`func (o *AdvisoryThreatActorWithExternalObjects) GetMitreIdOk() (*string, bool)`

GetMitreIdOk returns a tuple with the MitreId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMitreId

`func (o *AdvisoryThreatActorWithExternalObjects) SetMitreId(v string)`

SetMitreId sets MitreId field to given value.

### HasMitreId

`func (o *AdvisoryThreatActorWithExternalObjects) HasMitreId() bool`

HasMitreId returns a boolean if a field has been set.

### GetThreatActorName

`func (o *AdvisoryThreatActorWithExternalObjects) GetThreatActorName() string`

GetThreatActorName returns the ThreatActorName field if non-nil, zero value otherwise.

### GetThreatActorNameOk

`func (o *AdvisoryThreatActorWithExternalObjects) GetThreatActorNameOk() (*string, bool)`

GetThreatActorNameOk returns a tuple with the ThreatActorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreatActorName

`func (o *AdvisoryThreatActorWithExternalObjects) SetThreatActorName(v string)`

SetThreatActorName sets ThreatActorName field to given value.

### HasThreatActorName

`func (o *AdvisoryThreatActorWithExternalObjects) HasThreatActorName() bool`

HasThreatActorName returns a boolean if a field has been set.

### GetTools

`func (o *AdvisoryThreatActorWithExternalObjects) GetTools() []AdvisoryTool`

GetTools returns the Tools field if non-nil, zero value otherwise.

### GetToolsOk

`func (o *AdvisoryThreatActorWithExternalObjects) GetToolsOk() (*[]AdvisoryTool, bool)`

GetToolsOk returns a tuple with the Tools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTools

`func (o *AdvisoryThreatActorWithExternalObjects) SetTools(v []AdvisoryTool)`

SetTools sets Tools field to given value.

### HasTools

`func (o *AdvisoryThreatActorWithExternalObjects) HasTools() bool`

HasTools returns a boolean if a field has been set.

### GetVendorNamesForThreatActors

`func (o *AdvisoryThreatActorWithExternalObjects) GetVendorNamesForThreatActors() []AdvisoryVendorNameForThreatActor`

GetVendorNamesForThreatActors returns the VendorNamesForThreatActors field if non-nil, zero value otherwise.

### GetVendorNamesForThreatActorsOk

`func (o *AdvisoryThreatActorWithExternalObjects) GetVendorNamesForThreatActorsOk() (*[]AdvisoryVendorNameForThreatActor, bool)`

GetVendorNamesForThreatActorsOk returns a tuple with the VendorNamesForThreatActors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorNamesForThreatActors

`func (o *AdvisoryThreatActorWithExternalObjects) SetVendorNamesForThreatActors(v []AdvisoryVendorNameForThreatActor)`

SetVendorNamesForThreatActors sets VendorNamesForThreatActors field to given value.

### HasVendorNamesForThreatActors

`func (o *AdvisoryThreatActorWithExternalObjects) HasVendorNamesForThreatActors() bool`

HasVendorNamesForThreatActors returns a boolean if a field has been set.

### GetVendorsAndProductsTargeted

`func (o *AdvisoryThreatActorWithExternalObjects) GetVendorsAndProductsTargeted() []AdvisoryVendorProduct`

GetVendorsAndProductsTargeted returns the VendorsAndProductsTargeted field if non-nil, zero value otherwise.

### GetVendorsAndProductsTargetedOk

`func (o *AdvisoryThreatActorWithExternalObjects) GetVendorsAndProductsTargetedOk() (*[]AdvisoryVendorProduct, bool)`

GetVendorsAndProductsTargetedOk returns a tuple with the VendorsAndProductsTargeted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorsAndProductsTargeted

`func (o *AdvisoryThreatActorWithExternalObjects) SetVendorsAndProductsTargeted(v []AdvisoryVendorProduct)`

SetVendorsAndProductsTargeted sets VendorsAndProductsTargeted field to given value.

### HasVendorsAndProductsTargeted

`func (o *AdvisoryThreatActorWithExternalObjects) HasVendorsAndProductsTargeted() bool`

HasVendorsAndProductsTargeted returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


