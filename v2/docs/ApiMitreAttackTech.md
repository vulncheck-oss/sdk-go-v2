# ApiMitreAttackTech

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**D3fendmapping** | Pointer to [**[]ApiMitreMitigation2D3fendMapping**](ApiMitreMitigation2D3fendMapping.md) |  | [optional] 
**Detections** | Pointer to [**[]ApiMitreDetectionTech**](ApiMitreDetectionTech.md) |  | [optional] 
**Domain** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Mitigations** | Pointer to [**[]ApiMitreMitigationTech**](ApiMitreMitigationTech.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Subtechnique** | Pointer to **bool** |  | [optional] 
**Tactics** | Pointer to **[]string** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 

## Methods

### NewApiMitreAttackTech

`func NewApiMitreAttackTech() *ApiMitreAttackTech`

NewApiMitreAttackTech instantiates a new ApiMitreAttackTech object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiMitreAttackTechWithDefaults

`func NewApiMitreAttackTechWithDefaults() *ApiMitreAttackTech`

NewApiMitreAttackTechWithDefaults instantiates a new ApiMitreAttackTech object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetD3fendmapping

`func (o *ApiMitreAttackTech) GetD3fendmapping() []ApiMitreMitigation2D3fendMapping`

GetD3fendmapping returns the D3fendmapping field if non-nil, zero value otherwise.

### GetD3fendmappingOk

`func (o *ApiMitreAttackTech) GetD3fendmappingOk() (*[]ApiMitreMitigation2D3fendMapping, bool)`

GetD3fendmappingOk returns a tuple with the D3fendmapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetD3fendmapping

`func (o *ApiMitreAttackTech) SetD3fendmapping(v []ApiMitreMitigation2D3fendMapping)`

SetD3fendmapping sets D3fendmapping field to given value.

### HasD3fendmapping

`func (o *ApiMitreAttackTech) HasD3fendmapping() bool`

HasD3fendmapping returns a boolean if a field has been set.

### GetDetections

`func (o *ApiMitreAttackTech) GetDetections() []ApiMitreDetectionTech`

GetDetections returns the Detections field if non-nil, zero value otherwise.

### GetDetectionsOk

`func (o *ApiMitreAttackTech) GetDetectionsOk() (*[]ApiMitreDetectionTech, bool)`

GetDetectionsOk returns a tuple with the Detections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetections

`func (o *ApiMitreAttackTech) SetDetections(v []ApiMitreDetectionTech)`

SetDetections sets Detections field to given value.

### HasDetections

`func (o *ApiMitreAttackTech) HasDetections() bool`

HasDetections returns a boolean if a field has been set.

### GetDomain

`func (o *ApiMitreAttackTech) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *ApiMitreAttackTech) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *ApiMitreAttackTech) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *ApiMitreAttackTech) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetId

`func (o *ApiMitreAttackTech) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApiMitreAttackTech) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApiMitreAttackTech) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ApiMitreAttackTech) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMitigations

`func (o *ApiMitreAttackTech) GetMitigations() []ApiMitreMitigationTech`

GetMitigations returns the Mitigations field if non-nil, zero value otherwise.

### GetMitigationsOk

`func (o *ApiMitreAttackTech) GetMitigationsOk() (*[]ApiMitreMitigationTech, bool)`

GetMitigationsOk returns a tuple with the Mitigations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMitigations

`func (o *ApiMitreAttackTech) SetMitigations(v []ApiMitreMitigationTech)`

SetMitigations sets Mitigations field to given value.

### HasMitigations

`func (o *ApiMitreAttackTech) HasMitigations() bool`

HasMitigations returns a boolean if a field has been set.

### GetName

`func (o *ApiMitreAttackTech) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApiMitreAttackTech) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApiMitreAttackTech) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ApiMitreAttackTech) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSubtechnique

`func (o *ApiMitreAttackTech) GetSubtechnique() bool`

GetSubtechnique returns the Subtechnique field if non-nil, zero value otherwise.

### GetSubtechniqueOk

`func (o *ApiMitreAttackTech) GetSubtechniqueOk() (*bool, bool)`

GetSubtechniqueOk returns a tuple with the Subtechnique field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtechnique

`func (o *ApiMitreAttackTech) SetSubtechnique(v bool)`

SetSubtechnique sets Subtechnique field to given value.

### HasSubtechnique

`func (o *ApiMitreAttackTech) HasSubtechnique() bool`

HasSubtechnique returns a boolean if a field has been set.

### GetTactics

`func (o *ApiMitreAttackTech) GetTactics() []string`

GetTactics returns the Tactics field if non-nil, zero value otherwise.

### GetTacticsOk

`func (o *ApiMitreAttackTech) GetTacticsOk() (*[]string, bool)`

GetTacticsOk returns a tuple with the Tactics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTactics

`func (o *ApiMitreAttackTech) SetTactics(v []string)`

SetTactics sets Tactics field to given value.

### HasTactics

`func (o *ApiMitreAttackTech) HasTactics() bool`

HasTactics returns a boolean if a field has been set.

### GetUrl

`func (o *ApiMitreAttackTech) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ApiMitreAttackTech) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ApiMitreAttackTech) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *ApiMitreAttackTech) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


