# AdvisoryPaloAltoAdvisory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Affected** | Pointer to **string** |  | [optional] 
**ApplicableVersions** | Pointer to **string** |  | [optional] 
**AttackComplexity** | Pointer to **string** |  | [optional] 
**AttackRequirements** | Pointer to **string** |  | [optional] 
**AttackVector** | Pointer to **string** |  | [optional] 
**AvailabilityImpact** | Pointer to **string** |  | [optional] 
**ConfidentialityImpact** | Pointer to **string** |  | [optional] 
**Cve** | Pointer to **[]string** |  | [optional] 
**CvssbaseScore** | Pointer to **string** |  | [optional] 
**DatePublished** | Pointer to **string** |  | [optional] 
**DateUpdated** | Pointer to **string** |  | [optional] 
**DateAdded** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**IntegrityImpact** | Pointer to **string** |  | [optional] 
**PrivilegesRequired** | Pointer to **string** |  | [optional] 
**Problem** | Pointer to **string** |  | [optional] 
**Product** | Pointer to **string** |  | [optional] 
**Scope** | Pointer to **string** |  | [optional] 
**Severity** | Pointer to **string** |  | [optional] 
**Solution** | Pointer to **string** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**Unaffected** | Pointer to **string** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 
**UserInteraction** | Pointer to **string** |  | [optional] 
**Workaround** | Pointer to **string** |  | [optional] 

## Methods

### NewAdvisoryPaloAltoAdvisory

`func NewAdvisoryPaloAltoAdvisory() *AdvisoryPaloAltoAdvisory`

NewAdvisoryPaloAltoAdvisory instantiates a new AdvisoryPaloAltoAdvisory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdvisoryPaloAltoAdvisoryWithDefaults

`func NewAdvisoryPaloAltoAdvisoryWithDefaults() *AdvisoryPaloAltoAdvisory`

NewAdvisoryPaloAltoAdvisoryWithDefaults instantiates a new AdvisoryPaloAltoAdvisory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAffected

`func (o *AdvisoryPaloAltoAdvisory) GetAffected() string`

GetAffected returns the Affected field if non-nil, zero value otherwise.

### GetAffectedOk

`func (o *AdvisoryPaloAltoAdvisory) GetAffectedOk() (*string, bool)`

GetAffectedOk returns a tuple with the Affected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffected

`func (o *AdvisoryPaloAltoAdvisory) SetAffected(v string)`

SetAffected sets Affected field to given value.

### HasAffected

`func (o *AdvisoryPaloAltoAdvisory) HasAffected() bool`

HasAffected returns a boolean if a field has been set.

### GetApplicableVersions

`func (o *AdvisoryPaloAltoAdvisory) GetApplicableVersions() string`

GetApplicableVersions returns the ApplicableVersions field if non-nil, zero value otherwise.

### GetApplicableVersionsOk

`func (o *AdvisoryPaloAltoAdvisory) GetApplicableVersionsOk() (*string, bool)`

GetApplicableVersionsOk returns a tuple with the ApplicableVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicableVersions

`func (o *AdvisoryPaloAltoAdvisory) SetApplicableVersions(v string)`

SetApplicableVersions sets ApplicableVersions field to given value.

### HasApplicableVersions

`func (o *AdvisoryPaloAltoAdvisory) HasApplicableVersions() bool`

HasApplicableVersions returns a boolean if a field has been set.

### GetAttackComplexity

`func (o *AdvisoryPaloAltoAdvisory) GetAttackComplexity() string`

GetAttackComplexity returns the AttackComplexity field if non-nil, zero value otherwise.

### GetAttackComplexityOk

`func (o *AdvisoryPaloAltoAdvisory) GetAttackComplexityOk() (*string, bool)`

GetAttackComplexityOk returns a tuple with the AttackComplexity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttackComplexity

`func (o *AdvisoryPaloAltoAdvisory) SetAttackComplexity(v string)`

SetAttackComplexity sets AttackComplexity field to given value.

### HasAttackComplexity

`func (o *AdvisoryPaloAltoAdvisory) HasAttackComplexity() bool`

HasAttackComplexity returns a boolean if a field has been set.

### GetAttackRequirements

`func (o *AdvisoryPaloAltoAdvisory) GetAttackRequirements() string`

GetAttackRequirements returns the AttackRequirements field if non-nil, zero value otherwise.

### GetAttackRequirementsOk

`func (o *AdvisoryPaloAltoAdvisory) GetAttackRequirementsOk() (*string, bool)`

GetAttackRequirementsOk returns a tuple with the AttackRequirements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttackRequirements

`func (o *AdvisoryPaloAltoAdvisory) SetAttackRequirements(v string)`

SetAttackRequirements sets AttackRequirements field to given value.

### HasAttackRequirements

`func (o *AdvisoryPaloAltoAdvisory) HasAttackRequirements() bool`

HasAttackRequirements returns a boolean if a field has been set.

### GetAttackVector

`func (o *AdvisoryPaloAltoAdvisory) GetAttackVector() string`

GetAttackVector returns the AttackVector field if non-nil, zero value otherwise.

### GetAttackVectorOk

`func (o *AdvisoryPaloAltoAdvisory) GetAttackVectorOk() (*string, bool)`

GetAttackVectorOk returns a tuple with the AttackVector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttackVector

`func (o *AdvisoryPaloAltoAdvisory) SetAttackVector(v string)`

SetAttackVector sets AttackVector field to given value.

### HasAttackVector

`func (o *AdvisoryPaloAltoAdvisory) HasAttackVector() bool`

HasAttackVector returns a boolean if a field has been set.

### GetAvailabilityImpact

`func (o *AdvisoryPaloAltoAdvisory) GetAvailabilityImpact() string`

GetAvailabilityImpact returns the AvailabilityImpact field if non-nil, zero value otherwise.

### GetAvailabilityImpactOk

`func (o *AdvisoryPaloAltoAdvisory) GetAvailabilityImpactOk() (*string, bool)`

GetAvailabilityImpactOk returns a tuple with the AvailabilityImpact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityImpact

`func (o *AdvisoryPaloAltoAdvisory) SetAvailabilityImpact(v string)`

SetAvailabilityImpact sets AvailabilityImpact field to given value.

### HasAvailabilityImpact

`func (o *AdvisoryPaloAltoAdvisory) HasAvailabilityImpact() bool`

HasAvailabilityImpact returns a boolean if a field has been set.

### GetConfidentialityImpact

`func (o *AdvisoryPaloAltoAdvisory) GetConfidentialityImpact() string`

GetConfidentialityImpact returns the ConfidentialityImpact field if non-nil, zero value otherwise.

### GetConfidentialityImpactOk

`func (o *AdvisoryPaloAltoAdvisory) GetConfidentialityImpactOk() (*string, bool)`

GetConfidentialityImpactOk returns a tuple with the ConfidentialityImpact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfidentialityImpact

`func (o *AdvisoryPaloAltoAdvisory) SetConfidentialityImpact(v string)`

SetConfidentialityImpact sets ConfidentialityImpact field to given value.

### HasConfidentialityImpact

`func (o *AdvisoryPaloAltoAdvisory) HasConfidentialityImpact() bool`

HasConfidentialityImpact returns a boolean if a field has been set.

### GetCve

`func (o *AdvisoryPaloAltoAdvisory) GetCve() []string`

GetCve returns the Cve field if non-nil, zero value otherwise.

### GetCveOk

`func (o *AdvisoryPaloAltoAdvisory) GetCveOk() (*[]string, bool)`

GetCveOk returns a tuple with the Cve field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCve

`func (o *AdvisoryPaloAltoAdvisory) SetCve(v []string)`

SetCve sets Cve field to given value.

### HasCve

`func (o *AdvisoryPaloAltoAdvisory) HasCve() bool`

HasCve returns a boolean if a field has been set.

### GetCvssbaseScore

`func (o *AdvisoryPaloAltoAdvisory) GetCvssbaseScore() string`

GetCvssbaseScore returns the CvssbaseScore field if non-nil, zero value otherwise.

### GetCvssbaseScoreOk

`func (o *AdvisoryPaloAltoAdvisory) GetCvssbaseScoreOk() (*string, bool)`

GetCvssbaseScoreOk returns a tuple with the CvssbaseScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCvssbaseScore

`func (o *AdvisoryPaloAltoAdvisory) SetCvssbaseScore(v string)`

SetCvssbaseScore sets CvssbaseScore field to given value.

### HasCvssbaseScore

`func (o *AdvisoryPaloAltoAdvisory) HasCvssbaseScore() bool`

HasCvssbaseScore returns a boolean if a field has been set.

### GetDatePublished

`func (o *AdvisoryPaloAltoAdvisory) GetDatePublished() string`

GetDatePublished returns the DatePublished field if non-nil, zero value otherwise.

### GetDatePublishedOk

`func (o *AdvisoryPaloAltoAdvisory) GetDatePublishedOk() (*string, bool)`

GetDatePublishedOk returns a tuple with the DatePublished field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatePublished

`func (o *AdvisoryPaloAltoAdvisory) SetDatePublished(v string)`

SetDatePublished sets DatePublished field to given value.

### HasDatePublished

`func (o *AdvisoryPaloAltoAdvisory) HasDatePublished() bool`

HasDatePublished returns a boolean if a field has been set.

### GetDateUpdated

`func (o *AdvisoryPaloAltoAdvisory) GetDateUpdated() string`

GetDateUpdated returns the DateUpdated field if non-nil, zero value otherwise.

### GetDateUpdatedOk

`func (o *AdvisoryPaloAltoAdvisory) GetDateUpdatedOk() (*string, bool)`

GetDateUpdatedOk returns a tuple with the DateUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateUpdated

`func (o *AdvisoryPaloAltoAdvisory) SetDateUpdated(v string)`

SetDateUpdated sets DateUpdated field to given value.

### HasDateUpdated

`func (o *AdvisoryPaloAltoAdvisory) HasDateUpdated() bool`

HasDateUpdated returns a boolean if a field has been set.

### GetDateAdded

`func (o *AdvisoryPaloAltoAdvisory) GetDateAdded() string`

GetDateAdded returns the DateAdded field if non-nil, zero value otherwise.

### GetDateAddedOk

`func (o *AdvisoryPaloAltoAdvisory) GetDateAddedOk() (*string, bool)`

GetDateAddedOk returns a tuple with the DateAdded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateAdded

`func (o *AdvisoryPaloAltoAdvisory) SetDateAdded(v string)`

SetDateAdded sets DateAdded field to given value.

### HasDateAdded

`func (o *AdvisoryPaloAltoAdvisory) HasDateAdded() bool`

HasDateAdded returns a boolean if a field has been set.

### GetId

`func (o *AdvisoryPaloAltoAdvisory) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AdvisoryPaloAltoAdvisory) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AdvisoryPaloAltoAdvisory) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AdvisoryPaloAltoAdvisory) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIntegrityImpact

`func (o *AdvisoryPaloAltoAdvisory) GetIntegrityImpact() string`

GetIntegrityImpact returns the IntegrityImpact field if non-nil, zero value otherwise.

### GetIntegrityImpactOk

`func (o *AdvisoryPaloAltoAdvisory) GetIntegrityImpactOk() (*string, bool)`

GetIntegrityImpactOk returns a tuple with the IntegrityImpact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrityImpact

`func (o *AdvisoryPaloAltoAdvisory) SetIntegrityImpact(v string)`

SetIntegrityImpact sets IntegrityImpact field to given value.

### HasIntegrityImpact

`func (o *AdvisoryPaloAltoAdvisory) HasIntegrityImpact() bool`

HasIntegrityImpact returns a boolean if a field has been set.

### GetPrivilegesRequired

`func (o *AdvisoryPaloAltoAdvisory) GetPrivilegesRequired() string`

GetPrivilegesRequired returns the PrivilegesRequired field if non-nil, zero value otherwise.

### GetPrivilegesRequiredOk

`func (o *AdvisoryPaloAltoAdvisory) GetPrivilegesRequiredOk() (*string, bool)`

GetPrivilegesRequiredOk returns a tuple with the PrivilegesRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivilegesRequired

`func (o *AdvisoryPaloAltoAdvisory) SetPrivilegesRequired(v string)`

SetPrivilegesRequired sets PrivilegesRequired field to given value.

### HasPrivilegesRequired

`func (o *AdvisoryPaloAltoAdvisory) HasPrivilegesRequired() bool`

HasPrivilegesRequired returns a boolean if a field has been set.

### GetProblem

`func (o *AdvisoryPaloAltoAdvisory) GetProblem() string`

GetProblem returns the Problem field if non-nil, zero value otherwise.

### GetProblemOk

`func (o *AdvisoryPaloAltoAdvisory) GetProblemOk() (*string, bool)`

GetProblemOk returns a tuple with the Problem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProblem

`func (o *AdvisoryPaloAltoAdvisory) SetProblem(v string)`

SetProblem sets Problem field to given value.

### HasProblem

`func (o *AdvisoryPaloAltoAdvisory) HasProblem() bool`

HasProblem returns a boolean if a field has been set.

### GetProduct

`func (o *AdvisoryPaloAltoAdvisory) GetProduct() string`

GetProduct returns the Product field if non-nil, zero value otherwise.

### GetProductOk

`func (o *AdvisoryPaloAltoAdvisory) GetProductOk() (*string, bool)`

GetProductOk returns a tuple with the Product field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduct

`func (o *AdvisoryPaloAltoAdvisory) SetProduct(v string)`

SetProduct sets Product field to given value.

### HasProduct

`func (o *AdvisoryPaloAltoAdvisory) HasProduct() bool`

HasProduct returns a boolean if a field has been set.

### GetScope

`func (o *AdvisoryPaloAltoAdvisory) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *AdvisoryPaloAltoAdvisory) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *AdvisoryPaloAltoAdvisory) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *AdvisoryPaloAltoAdvisory) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetSeverity

`func (o *AdvisoryPaloAltoAdvisory) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *AdvisoryPaloAltoAdvisory) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *AdvisoryPaloAltoAdvisory) SetSeverity(v string)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *AdvisoryPaloAltoAdvisory) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetSolution

`func (o *AdvisoryPaloAltoAdvisory) GetSolution() string`

GetSolution returns the Solution field if non-nil, zero value otherwise.

### GetSolutionOk

`func (o *AdvisoryPaloAltoAdvisory) GetSolutionOk() (*string, bool)`

GetSolutionOk returns a tuple with the Solution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSolution

`func (o *AdvisoryPaloAltoAdvisory) SetSolution(v string)`

SetSolution sets Solution field to given value.

### HasSolution

`func (o *AdvisoryPaloAltoAdvisory) HasSolution() bool`

HasSolution returns a boolean if a field has been set.

### GetTitle

`func (o *AdvisoryPaloAltoAdvisory) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *AdvisoryPaloAltoAdvisory) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *AdvisoryPaloAltoAdvisory) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *AdvisoryPaloAltoAdvisory) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetUnaffected

`func (o *AdvisoryPaloAltoAdvisory) GetUnaffected() string`

GetUnaffected returns the Unaffected field if non-nil, zero value otherwise.

### GetUnaffectedOk

`func (o *AdvisoryPaloAltoAdvisory) GetUnaffectedOk() (*string, bool)`

GetUnaffectedOk returns a tuple with the Unaffected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnaffected

`func (o *AdvisoryPaloAltoAdvisory) SetUnaffected(v string)`

SetUnaffected sets Unaffected field to given value.

### HasUnaffected

`func (o *AdvisoryPaloAltoAdvisory) HasUnaffected() bool`

HasUnaffected returns a boolean if a field has been set.

### GetUrl

`func (o *AdvisoryPaloAltoAdvisory) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *AdvisoryPaloAltoAdvisory) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *AdvisoryPaloAltoAdvisory) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *AdvisoryPaloAltoAdvisory) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetUserInteraction

`func (o *AdvisoryPaloAltoAdvisory) GetUserInteraction() string`

GetUserInteraction returns the UserInteraction field if non-nil, zero value otherwise.

### GetUserInteractionOk

`func (o *AdvisoryPaloAltoAdvisory) GetUserInteractionOk() (*string, bool)`

GetUserInteractionOk returns a tuple with the UserInteraction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserInteraction

`func (o *AdvisoryPaloAltoAdvisory) SetUserInteraction(v string)`

SetUserInteraction sets UserInteraction field to given value.

### HasUserInteraction

`func (o *AdvisoryPaloAltoAdvisory) HasUserInteraction() bool`

HasUserInteraction returns a boolean if a field has been set.

### GetWorkaround

`func (o *AdvisoryPaloAltoAdvisory) GetWorkaround() string`

GetWorkaround returns the Workaround field if non-nil, zero value otherwise.

### GetWorkaroundOk

`func (o *AdvisoryPaloAltoAdvisory) GetWorkaroundOk() (*string, bool)`

GetWorkaroundOk returns a tuple with the Workaround field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkaround

`func (o *AdvisoryPaloAltoAdvisory) SetWorkaround(v string)`

SetWorkaround sets Workaround field to given value.

### HasWorkaround

`func (o *AdvisoryPaloAltoAdvisory) HasWorkaround() bool`

HasWorkaround returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


