# AdvisoryGHAdvisoryJSONLean

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Classification** | Pointer to **string** |  | [optional] 
**Cve** | Pointer to **[]string** |  | [optional] 
**Cvss** | Pointer to [**AdvisoryGHCvss**](AdvisoryGHCvss.md) |  | [optional] 
**Cwes** | Pointer to [**AdvisoryCwes**](AdvisoryCwes.md) |  | [optional] 
**DatabaseId** | Pointer to **int32** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**GhsaId** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Identifiers** | Pointer to [**[]AdvisoryGHIdentifier**](AdvisoryGHIdentifier.md) |  | [optional] 
**NotificationsPermalink** | Pointer to **string** |  | [optional] 
**Origin** | Pointer to **string** |  | [optional] 
**Permalink** | Pointer to **string** |  | [optional] 
**PublishedAt** | Pointer to **string** |  | [optional] 
**References** | Pointer to [**[]AdvisoryGHReference**](AdvisoryGHReference.md) |  | [optional] 
**Severity** | Pointer to **string** |  | [optional] 
**Summary** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 
**Vulnerabilities** | Pointer to [**AdvisoryGHVulnerabilities**](AdvisoryGHVulnerabilities.md) |  | [optional] 
**WithdrawnAt** | Pointer to **string** |  | [optional] 

## Methods

### NewAdvisoryGHAdvisoryJSONLean

`func NewAdvisoryGHAdvisoryJSONLean() *AdvisoryGHAdvisoryJSONLean`

NewAdvisoryGHAdvisoryJSONLean instantiates a new AdvisoryGHAdvisoryJSONLean object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdvisoryGHAdvisoryJSONLeanWithDefaults

`func NewAdvisoryGHAdvisoryJSONLeanWithDefaults() *AdvisoryGHAdvisoryJSONLean`

NewAdvisoryGHAdvisoryJSONLeanWithDefaults instantiates a new AdvisoryGHAdvisoryJSONLean object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassification

`func (o *AdvisoryGHAdvisoryJSONLean) GetClassification() string`

GetClassification returns the Classification field if non-nil, zero value otherwise.

### GetClassificationOk

`func (o *AdvisoryGHAdvisoryJSONLean) GetClassificationOk() (*string, bool)`

GetClassificationOk returns a tuple with the Classification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassification

`func (o *AdvisoryGHAdvisoryJSONLean) SetClassification(v string)`

SetClassification sets Classification field to given value.

### HasClassification

`func (o *AdvisoryGHAdvisoryJSONLean) HasClassification() bool`

HasClassification returns a boolean if a field has been set.

### GetCve

`func (o *AdvisoryGHAdvisoryJSONLean) GetCve() []string`

GetCve returns the Cve field if non-nil, zero value otherwise.

### GetCveOk

`func (o *AdvisoryGHAdvisoryJSONLean) GetCveOk() (*[]string, bool)`

GetCveOk returns a tuple with the Cve field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCve

`func (o *AdvisoryGHAdvisoryJSONLean) SetCve(v []string)`

SetCve sets Cve field to given value.

### HasCve

`func (o *AdvisoryGHAdvisoryJSONLean) HasCve() bool`

HasCve returns a boolean if a field has been set.

### GetCvss

`func (o *AdvisoryGHAdvisoryJSONLean) GetCvss() AdvisoryGHCvss`

GetCvss returns the Cvss field if non-nil, zero value otherwise.

### GetCvssOk

`func (o *AdvisoryGHAdvisoryJSONLean) GetCvssOk() (*AdvisoryGHCvss, bool)`

GetCvssOk returns a tuple with the Cvss field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCvss

`func (o *AdvisoryGHAdvisoryJSONLean) SetCvss(v AdvisoryGHCvss)`

SetCvss sets Cvss field to given value.

### HasCvss

`func (o *AdvisoryGHAdvisoryJSONLean) HasCvss() bool`

HasCvss returns a boolean if a field has been set.

### GetCwes

`func (o *AdvisoryGHAdvisoryJSONLean) GetCwes() AdvisoryCwes`

GetCwes returns the Cwes field if non-nil, zero value otherwise.

### GetCwesOk

`func (o *AdvisoryGHAdvisoryJSONLean) GetCwesOk() (*AdvisoryCwes, bool)`

GetCwesOk returns a tuple with the Cwes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCwes

`func (o *AdvisoryGHAdvisoryJSONLean) SetCwes(v AdvisoryCwes)`

SetCwes sets Cwes field to given value.

### HasCwes

`func (o *AdvisoryGHAdvisoryJSONLean) HasCwes() bool`

HasCwes returns a boolean if a field has been set.

### GetDatabaseId

`func (o *AdvisoryGHAdvisoryJSONLean) GetDatabaseId() int32`

GetDatabaseId returns the DatabaseId field if non-nil, zero value otherwise.

### GetDatabaseIdOk

`func (o *AdvisoryGHAdvisoryJSONLean) GetDatabaseIdOk() (*int32, bool)`

GetDatabaseIdOk returns a tuple with the DatabaseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseId

`func (o *AdvisoryGHAdvisoryJSONLean) SetDatabaseId(v int32)`

SetDatabaseId sets DatabaseId field to given value.

### HasDatabaseId

`func (o *AdvisoryGHAdvisoryJSONLean) HasDatabaseId() bool`

HasDatabaseId returns a boolean if a field has been set.

### GetDescription

`func (o *AdvisoryGHAdvisoryJSONLean) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AdvisoryGHAdvisoryJSONLean) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AdvisoryGHAdvisoryJSONLean) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AdvisoryGHAdvisoryJSONLean) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetGhsaId

`func (o *AdvisoryGHAdvisoryJSONLean) GetGhsaId() string`

GetGhsaId returns the GhsaId field if non-nil, zero value otherwise.

### GetGhsaIdOk

`func (o *AdvisoryGHAdvisoryJSONLean) GetGhsaIdOk() (*string, bool)`

GetGhsaIdOk returns a tuple with the GhsaId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGhsaId

`func (o *AdvisoryGHAdvisoryJSONLean) SetGhsaId(v string)`

SetGhsaId sets GhsaId field to given value.

### HasGhsaId

`func (o *AdvisoryGHAdvisoryJSONLean) HasGhsaId() bool`

HasGhsaId returns a boolean if a field has been set.

### GetId

`func (o *AdvisoryGHAdvisoryJSONLean) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AdvisoryGHAdvisoryJSONLean) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AdvisoryGHAdvisoryJSONLean) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AdvisoryGHAdvisoryJSONLean) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIdentifiers

`func (o *AdvisoryGHAdvisoryJSONLean) GetIdentifiers() []AdvisoryGHIdentifier`

GetIdentifiers returns the Identifiers field if non-nil, zero value otherwise.

### GetIdentifiersOk

`func (o *AdvisoryGHAdvisoryJSONLean) GetIdentifiersOk() (*[]AdvisoryGHIdentifier, bool)`

GetIdentifiersOk returns a tuple with the Identifiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifiers

`func (o *AdvisoryGHAdvisoryJSONLean) SetIdentifiers(v []AdvisoryGHIdentifier)`

SetIdentifiers sets Identifiers field to given value.

### HasIdentifiers

`func (o *AdvisoryGHAdvisoryJSONLean) HasIdentifiers() bool`

HasIdentifiers returns a boolean if a field has been set.

### GetNotificationsPermalink

`func (o *AdvisoryGHAdvisoryJSONLean) GetNotificationsPermalink() string`

GetNotificationsPermalink returns the NotificationsPermalink field if non-nil, zero value otherwise.

### GetNotificationsPermalinkOk

`func (o *AdvisoryGHAdvisoryJSONLean) GetNotificationsPermalinkOk() (*string, bool)`

GetNotificationsPermalinkOk returns a tuple with the NotificationsPermalink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationsPermalink

`func (o *AdvisoryGHAdvisoryJSONLean) SetNotificationsPermalink(v string)`

SetNotificationsPermalink sets NotificationsPermalink field to given value.

### HasNotificationsPermalink

`func (o *AdvisoryGHAdvisoryJSONLean) HasNotificationsPermalink() bool`

HasNotificationsPermalink returns a boolean if a field has been set.

### GetOrigin

`func (o *AdvisoryGHAdvisoryJSONLean) GetOrigin() string`

GetOrigin returns the Origin field if non-nil, zero value otherwise.

### GetOriginOk

`func (o *AdvisoryGHAdvisoryJSONLean) GetOriginOk() (*string, bool)`

GetOriginOk returns a tuple with the Origin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigin

`func (o *AdvisoryGHAdvisoryJSONLean) SetOrigin(v string)`

SetOrigin sets Origin field to given value.

### HasOrigin

`func (o *AdvisoryGHAdvisoryJSONLean) HasOrigin() bool`

HasOrigin returns a boolean if a field has been set.

### GetPermalink

`func (o *AdvisoryGHAdvisoryJSONLean) GetPermalink() string`

GetPermalink returns the Permalink field if non-nil, zero value otherwise.

### GetPermalinkOk

`func (o *AdvisoryGHAdvisoryJSONLean) GetPermalinkOk() (*string, bool)`

GetPermalinkOk returns a tuple with the Permalink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermalink

`func (o *AdvisoryGHAdvisoryJSONLean) SetPermalink(v string)`

SetPermalink sets Permalink field to given value.

### HasPermalink

`func (o *AdvisoryGHAdvisoryJSONLean) HasPermalink() bool`

HasPermalink returns a boolean if a field has been set.

### GetPublishedAt

`func (o *AdvisoryGHAdvisoryJSONLean) GetPublishedAt() string`

GetPublishedAt returns the PublishedAt field if non-nil, zero value otherwise.

### GetPublishedAtOk

`func (o *AdvisoryGHAdvisoryJSONLean) GetPublishedAtOk() (*string, bool)`

GetPublishedAtOk returns a tuple with the PublishedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishedAt

`func (o *AdvisoryGHAdvisoryJSONLean) SetPublishedAt(v string)`

SetPublishedAt sets PublishedAt field to given value.

### HasPublishedAt

`func (o *AdvisoryGHAdvisoryJSONLean) HasPublishedAt() bool`

HasPublishedAt returns a boolean if a field has been set.

### GetReferences

`func (o *AdvisoryGHAdvisoryJSONLean) GetReferences() []AdvisoryGHReference`

GetReferences returns the References field if non-nil, zero value otherwise.

### GetReferencesOk

`func (o *AdvisoryGHAdvisoryJSONLean) GetReferencesOk() (*[]AdvisoryGHReference, bool)`

GetReferencesOk returns a tuple with the References field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferences

`func (o *AdvisoryGHAdvisoryJSONLean) SetReferences(v []AdvisoryGHReference)`

SetReferences sets References field to given value.

### HasReferences

`func (o *AdvisoryGHAdvisoryJSONLean) HasReferences() bool`

HasReferences returns a boolean if a field has been set.

### GetSeverity

`func (o *AdvisoryGHAdvisoryJSONLean) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *AdvisoryGHAdvisoryJSONLean) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *AdvisoryGHAdvisoryJSONLean) SetSeverity(v string)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *AdvisoryGHAdvisoryJSONLean) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetSummary

`func (o *AdvisoryGHAdvisoryJSONLean) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *AdvisoryGHAdvisoryJSONLean) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *AdvisoryGHAdvisoryJSONLean) SetSummary(v string)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *AdvisoryGHAdvisoryJSONLean) HasSummary() bool`

HasSummary returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *AdvisoryGHAdvisoryJSONLean) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *AdvisoryGHAdvisoryJSONLean) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *AdvisoryGHAdvisoryJSONLean) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *AdvisoryGHAdvisoryJSONLean) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetVulnerabilities

`func (o *AdvisoryGHAdvisoryJSONLean) GetVulnerabilities() AdvisoryGHVulnerabilities`

GetVulnerabilities returns the Vulnerabilities field if non-nil, zero value otherwise.

### GetVulnerabilitiesOk

`func (o *AdvisoryGHAdvisoryJSONLean) GetVulnerabilitiesOk() (*AdvisoryGHVulnerabilities, bool)`

GetVulnerabilitiesOk returns a tuple with the Vulnerabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVulnerabilities

`func (o *AdvisoryGHAdvisoryJSONLean) SetVulnerabilities(v AdvisoryGHVulnerabilities)`

SetVulnerabilities sets Vulnerabilities field to given value.

### HasVulnerabilities

`func (o *AdvisoryGHAdvisoryJSONLean) HasVulnerabilities() bool`

HasVulnerabilities returns a boolean if a field has been set.

### GetWithdrawnAt

`func (o *AdvisoryGHAdvisoryJSONLean) GetWithdrawnAt() string`

GetWithdrawnAt returns the WithdrawnAt field if non-nil, zero value otherwise.

### GetWithdrawnAtOk

`func (o *AdvisoryGHAdvisoryJSONLean) GetWithdrawnAtOk() (*string, bool)`

GetWithdrawnAtOk returns a tuple with the WithdrawnAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithdrawnAt

`func (o *AdvisoryGHAdvisoryJSONLean) SetWithdrawnAt(v string)`

SetWithdrawnAt sets WithdrawnAt field to given value.

### HasWithdrawnAt

`func (o *AdvisoryGHAdvisoryJSONLean) HasWithdrawnAt() bool`

HasWithdrawnAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


