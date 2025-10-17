# ApiReferencesExtended

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReferenceData** | Pointer to [**[]ApiReferenceDataExtended**](ApiReferenceDataExtended.md) | ExploitData     []NormalizedExploit       &#x60;json:\&quot;exploit_data,omitempty\&quot;&#x60;   ThreatActorData []ThreatActorExtended     &#x60;json:\&quot;threat_actor_data,omitempty\&quot;&#x60;   RansomwareData  []RansomwareReferenceData &#x60;json:\&quot;ransomware_data,omitempty\&quot;&#x60;   AdvisoryData    []AdvisoryExtended        &#x60;json:\&quot;advisory_data,omitempty\&quot;&#x60;   IdentifierData  []IdentifierExtended      &#x60;json:\&quot;identifier_data,omitempty\&quot;&#x60; | [optional] 

## Methods

### NewApiReferencesExtended

`func NewApiReferencesExtended() *ApiReferencesExtended`

NewApiReferencesExtended instantiates a new ApiReferencesExtended object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiReferencesExtendedWithDefaults

`func NewApiReferencesExtendedWithDefaults() *ApiReferencesExtended`

NewApiReferencesExtendedWithDefaults instantiates a new ApiReferencesExtended object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReferenceData

`func (o *ApiReferencesExtended) GetReferenceData() []ApiReferenceDataExtended`

GetReferenceData returns the ReferenceData field if non-nil, zero value otherwise.

### GetReferenceDataOk

`func (o *ApiReferencesExtended) GetReferenceDataOk() (*[]ApiReferenceDataExtended, bool)`

GetReferenceDataOk returns a tuple with the ReferenceData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceData

`func (o *ApiReferencesExtended) SetReferenceData(v []ApiReferenceDataExtended)`

SetReferenceData sets ReferenceData field to given value.

### HasReferenceData

`func (o *ApiReferencesExtended) HasReferenceData() bool`

HasReferenceData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


