# AdvisoryMispMeta

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AttributionConfidence** | Pointer to **string** |  | [optional] 
**CfrSuspectedStateSponsor** | Pointer to **string** |  | [optional] 
**CfrSuspectedVictims** | Pointer to **[]string** |  | [optional] 
**CfrTargetCategory** | Pointer to **[]string** |  | [optional] 
**CfrTypeOfIncident** | Pointer to **[]string** |  | [optional] 
**Country** | Pointer to **string** |  | [optional] 
**Refs** | Pointer to **[]string** |  | [optional] 
**Synonyms** | Pointer to **[]string** |  | [optional] 

## Methods

### NewAdvisoryMispMeta

`func NewAdvisoryMispMeta() *AdvisoryMispMeta`

NewAdvisoryMispMeta instantiates a new AdvisoryMispMeta object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdvisoryMispMetaWithDefaults

`func NewAdvisoryMispMetaWithDefaults() *AdvisoryMispMeta`

NewAdvisoryMispMetaWithDefaults instantiates a new AdvisoryMispMeta object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttributionConfidence

`func (o *AdvisoryMispMeta) GetAttributionConfidence() string`

GetAttributionConfidence returns the AttributionConfidence field if non-nil, zero value otherwise.

### GetAttributionConfidenceOk

`func (o *AdvisoryMispMeta) GetAttributionConfidenceOk() (*string, bool)`

GetAttributionConfidenceOk returns a tuple with the AttributionConfidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributionConfidence

`func (o *AdvisoryMispMeta) SetAttributionConfidence(v string)`

SetAttributionConfidence sets AttributionConfidence field to given value.

### HasAttributionConfidence

`func (o *AdvisoryMispMeta) HasAttributionConfidence() bool`

HasAttributionConfidence returns a boolean if a field has been set.

### GetCfrSuspectedStateSponsor

`func (o *AdvisoryMispMeta) GetCfrSuspectedStateSponsor() string`

GetCfrSuspectedStateSponsor returns the CfrSuspectedStateSponsor field if non-nil, zero value otherwise.

### GetCfrSuspectedStateSponsorOk

`func (o *AdvisoryMispMeta) GetCfrSuspectedStateSponsorOk() (*string, bool)`

GetCfrSuspectedStateSponsorOk returns a tuple with the CfrSuspectedStateSponsor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCfrSuspectedStateSponsor

`func (o *AdvisoryMispMeta) SetCfrSuspectedStateSponsor(v string)`

SetCfrSuspectedStateSponsor sets CfrSuspectedStateSponsor field to given value.

### HasCfrSuspectedStateSponsor

`func (o *AdvisoryMispMeta) HasCfrSuspectedStateSponsor() bool`

HasCfrSuspectedStateSponsor returns a boolean if a field has been set.

### GetCfrSuspectedVictims

`func (o *AdvisoryMispMeta) GetCfrSuspectedVictims() []string`

GetCfrSuspectedVictims returns the CfrSuspectedVictims field if non-nil, zero value otherwise.

### GetCfrSuspectedVictimsOk

`func (o *AdvisoryMispMeta) GetCfrSuspectedVictimsOk() (*[]string, bool)`

GetCfrSuspectedVictimsOk returns a tuple with the CfrSuspectedVictims field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCfrSuspectedVictims

`func (o *AdvisoryMispMeta) SetCfrSuspectedVictims(v []string)`

SetCfrSuspectedVictims sets CfrSuspectedVictims field to given value.

### HasCfrSuspectedVictims

`func (o *AdvisoryMispMeta) HasCfrSuspectedVictims() bool`

HasCfrSuspectedVictims returns a boolean if a field has been set.

### GetCfrTargetCategory

`func (o *AdvisoryMispMeta) GetCfrTargetCategory() []string`

GetCfrTargetCategory returns the CfrTargetCategory field if non-nil, zero value otherwise.

### GetCfrTargetCategoryOk

`func (o *AdvisoryMispMeta) GetCfrTargetCategoryOk() (*[]string, bool)`

GetCfrTargetCategoryOk returns a tuple with the CfrTargetCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCfrTargetCategory

`func (o *AdvisoryMispMeta) SetCfrTargetCategory(v []string)`

SetCfrTargetCategory sets CfrTargetCategory field to given value.

### HasCfrTargetCategory

`func (o *AdvisoryMispMeta) HasCfrTargetCategory() bool`

HasCfrTargetCategory returns a boolean if a field has been set.

### GetCfrTypeOfIncident

`func (o *AdvisoryMispMeta) GetCfrTypeOfIncident() []string`

GetCfrTypeOfIncident returns the CfrTypeOfIncident field if non-nil, zero value otherwise.

### GetCfrTypeOfIncidentOk

`func (o *AdvisoryMispMeta) GetCfrTypeOfIncidentOk() (*[]string, bool)`

GetCfrTypeOfIncidentOk returns a tuple with the CfrTypeOfIncident field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCfrTypeOfIncident

`func (o *AdvisoryMispMeta) SetCfrTypeOfIncident(v []string)`

SetCfrTypeOfIncident sets CfrTypeOfIncident field to given value.

### HasCfrTypeOfIncident

`func (o *AdvisoryMispMeta) HasCfrTypeOfIncident() bool`

HasCfrTypeOfIncident returns a boolean if a field has been set.

### GetCountry

`func (o *AdvisoryMispMeta) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *AdvisoryMispMeta) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *AdvisoryMispMeta) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *AdvisoryMispMeta) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetRefs

`func (o *AdvisoryMispMeta) GetRefs() []string`

GetRefs returns the Refs field if non-nil, zero value otherwise.

### GetRefsOk

`func (o *AdvisoryMispMeta) GetRefsOk() (*[]string, bool)`

GetRefsOk returns a tuple with the Refs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefs

`func (o *AdvisoryMispMeta) SetRefs(v []string)`

SetRefs sets Refs field to given value.

### HasRefs

`func (o *AdvisoryMispMeta) HasRefs() bool`

HasRefs returns a boolean if a field has been set.

### GetSynonyms

`func (o *AdvisoryMispMeta) GetSynonyms() []string`

GetSynonyms returns the Synonyms field if non-nil, zero value otherwise.

### GetSynonymsOk

`func (o *AdvisoryMispMeta) GetSynonymsOk() (*[]string, bool)`

GetSynonymsOk returns a tuple with the Synonyms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSynonyms

`func (o *AdvisoryMispMeta) SetSynonyms(v []string)`

SetSynonyms sets Synonyms field to given value.

### HasSynonyms

`func (o *AdvisoryMispMeta) HasSynonyms() bool`

HasSynonyms returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


