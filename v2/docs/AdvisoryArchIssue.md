# AdvisoryArchIssue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Advisories** | Pointer to **[]string** |  | [optional] 
**Affected** | Pointer to **string** |  | [optional] 
**Cve** | Pointer to **[]string** |  | [optional] 
**DateAdded** | Pointer to **string** |  | [optional] 
**Fixed** | Pointer to **string** |  | [optional] 
**Issues** | Pointer to **[]string** | cves | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Packages** | Pointer to **[]string** |  | [optional] 
**References** | Pointer to **[]string** |  | [optional] 
**Severity** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Ticket** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 

## Methods

### NewAdvisoryArchIssue

`func NewAdvisoryArchIssue() *AdvisoryArchIssue`

NewAdvisoryArchIssue instantiates a new AdvisoryArchIssue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdvisoryArchIssueWithDefaults

`func NewAdvisoryArchIssueWithDefaults() *AdvisoryArchIssue`

NewAdvisoryArchIssueWithDefaults instantiates a new AdvisoryArchIssue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdvisories

`func (o *AdvisoryArchIssue) GetAdvisories() []string`

GetAdvisories returns the Advisories field if non-nil, zero value otherwise.

### GetAdvisoriesOk

`func (o *AdvisoryArchIssue) GetAdvisoriesOk() (*[]string, bool)`

GetAdvisoriesOk returns a tuple with the Advisories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvisories

`func (o *AdvisoryArchIssue) SetAdvisories(v []string)`

SetAdvisories sets Advisories field to given value.

### HasAdvisories

`func (o *AdvisoryArchIssue) HasAdvisories() bool`

HasAdvisories returns a boolean if a field has been set.

### GetAffected

`func (o *AdvisoryArchIssue) GetAffected() string`

GetAffected returns the Affected field if non-nil, zero value otherwise.

### GetAffectedOk

`func (o *AdvisoryArchIssue) GetAffectedOk() (*string, bool)`

GetAffectedOk returns a tuple with the Affected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffected

`func (o *AdvisoryArchIssue) SetAffected(v string)`

SetAffected sets Affected field to given value.

### HasAffected

`func (o *AdvisoryArchIssue) HasAffected() bool`

HasAffected returns a boolean if a field has been set.

### GetCve

`func (o *AdvisoryArchIssue) GetCve() []string`

GetCve returns the Cve field if non-nil, zero value otherwise.

### GetCveOk

`func (o *AdvisoryArchIssue) GetCveOk() (*[]string, bool)`

GetCveOk returns a tuple with the Cve field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCve

`func (o *AdvisoryArchIssue) SetCve(v []string)`

SetCve sets Cve field to given value.

### HasCve

`func (o *AdvisoryArchIssue) HasCve() bool`

HasCve returns a boolean if a field has been set.

### GetDateAdded

`func (o *AdvisoryArchIssue) GetDateAdded() string`

GetDateAdded returns the DateAdded field if non-nil, zero value otherwise.

### GetDateAddedOk

`func (o *AdvisoryArchIssue) GetDateAddedOk() (*string, bool)`

GetDateAddedOk returns a tuple with the DateAdded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateAdded

`func (o *AdvisoryArchIssue) SetDateAdded(v string)`

SetDateAdded sets DateAdded field to given value.

### HasDateAdded

`func (o *AdvisoryArchIssue) HasDateAdded() bool`

HasDateAdded returns a boolean if a field has been set.

### GetFixed

`func (o *AdvisoryArchIssue) GetFixed() string`

GetFixed returns the Fixed field if non-nil, zero value otherwise.

### GetFixedOk

`func (o *AdvisoryArchIssue) GetFixedOk() (*string, bool)`

GetFixedOk returns a tuple with the Fixed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFixed

`func (o *AdvisoryArchIssue) SetFixed(v string)`

SetFixed sets Fixed field to given value.

### HasFixed

`func (o *AdvisoryArchIssue) HasFixed() bool`

HasFixed returns a boolean if a field has been set.

### GetIssues

`func (o *AdvisoryArchIssue) GetIssues() []string`

GetIssues returns the Issues field if non-nil, zero value otherwise.

### GetIssuesOk

`func (o *AdvisoryArchIssue) GetIssuesOk() (*[]string, bool)`

GetIssuesOk returns a tuple with the Issues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssues

`func (o *AdvisoryArchIssue) SetIssues(v []string)`

SetIssues sets Issues field to given value.

### HasIssues

`func (o *AdvisoryArchIssue) HasIssues() bool`

HasIssues returns a boolean if a field has been set.

### GetName

`func (o *AdvisoryArchIssue) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AdvisoryArchIssue) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AdvisoryArchIssue) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AdvisoryArchIssue) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPackages

`func (o *AdvisoryArchIssue) GetPackages() []string`

GetPackages returns the Packages field if non-nil, zero value otherwise.

### GetPackagesOk

`func (o *AdvisoryArchIssue) GetPackagesOk() (*[]string, bool)`

GetPackagesOk returns a tuple with the Packages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackages

`func (o *AdvisoryArchIssue) SetPackages(v []string)`

SetPackages sets Packages field to given value.

### HasPackages

`func (o *AdvisoryArchIssue) HasPackages() bool`

HasPackages returns a boolean if a field has been set.

### GetReferences

`func (o *AdvisoryArchIssue) GetReferences() []string`

GetReferences returns the References field if non-nil, zero value otherwise.

### GetReferencesOk

`func (o *AdvisoryArchIssue) GetReferencesOk() (*[]string, bool)`

GetReferencesOk returns a tuple with the References field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferences

`func (o *AdvisoryArchIssue) SetReferences(v []string)`

SetReferences sets References field to given value.

### HasReferences

`func (o *AdvisoryArchIssue) HasReferences() bool`

HasReferences returns a boolean if a field has been set.

### GetSeverity

`func (o *AdvisoryArchIssue) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *AdvisoryArchIssue) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *AdvisoryArchIssue) SetSeverity(v string)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *AdvisoryArchIssue) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetStatus

`func (o *AdvisoryArchIssue) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AdvisoryArchIssue) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AdvisoryArchIssue) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AdvisoryArchIssue) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTicket

`func (o *AdvisoryArchIssue) GetTicket() string`

GetTicket returns the Ticket field if non-nil, zero value otherwise.

### GetTicketOk

`func (o *AdvisoryArchIssue) GetTicketOk() (*string, bool)`

GetTicketOk returns a tuple with the Ticket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTicket

`func (o *AdvisoryArchIssue) SetTicket(v string)`

SetTicket sets Ticket field to given value.

### HasTicket

`func (o *AdvisoryArchIssue) HasTicket() bool`

HasTicket returns a boolean if a field has been set.

### GetType

`func (o *AdvisoryArchIssue) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AdvisoryArchIssue) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AdvisoryArchIssue) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *AdvisoryArchIssue) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


