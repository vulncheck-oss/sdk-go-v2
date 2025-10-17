# ApiNVD20CVEExtended

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ALIAS** | Pointer to **string** |  | [optional] 
**STATUS** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **string** | the deep tag instructs deep.Equal to ignore this field (used during OpenSearch loading) | [optional] 
**Categorization** | Pointer to [**ApiCategorizationExtended**](ApiCategorizationExtended.md) |  | [optional] 
**CisaActionDue** | Pointer to **string** |  | [optional] 
**CisaExploitAdd** | Pointer to **string** |  | [optional] 
**CisaRequiredAction** | Pointer to **string** |  | [optional] 
**CisaVulnerabilityName** | Pointer to **string** |  | [optional] 
**Configurations** | Pointer to [**[]AdvisoryNVD20Configuration**](AdvisoryNVD20Configuration.md) |  | [optional] 
**DateAdded** | Pointer to **string** |  | [optional] 
**Descriptions** | Pointer to [**[]ApiNVD20Description**](ApiNVD20Description.md) |  | [optional] 
**DocumentGenerationDate** | Pointer to **string** |  | [optional] 
**EvaluatorComment** | Pointer to **string** |  | [optional] 
**EvaluatorImpact** | Pointer to **string** |  | [optional] 
**EvaluatorSolution** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**LastModified** | Pointer to **string** |  | [optional] 
**Metrics** | Pointer to [**ApiNVD20MetricExtended**](ApiNVD20MetricExtended.md) |  | [optional] 
**MitreAttackTechniques** | Pointer to [**[]ApiMitreAttackTech**](ApiMitreAttackTech.md) |  | [optional] 
**Published** | Pointer to **string** |  | [optional] 
**References** | Pointer to [**[]ApiNVD20ReferenceExtended**](ApiNVD20ReferenceExtended.md) |  | [optional] 
**RelatedAttackPatterns** | Pointer to [**[]ApiRelatedAttackPattern**](ApiRelatedAttackPattern.md) |  | [optional] 
**SourceIdentifier** | Pointer to **string** |  | [optional] 
**VcConfigurations** | Pointer to [**[]AdvisoryNVD20Configuration**](AdvisoryNVD20Configuration.md) |  | [optional] 
**VcVulnerableCPEs** | Pointer to **[]string** |  | [optional] 
**VendorComments** | Pointer to [**[]ApiNVD20VendorComment**](ApiNVD20VendorComment.md) |  | [optional] 
**VulnStatus** | Pointer to **string** |  | [optional] 
**VulncheckKEVExploitAdd** | Pointer to **string** |  | [optional] 
**VulnerableCPEs** | Pointer to **[]string** |  | [optional] 
**Weaknesses** | Pointer to [**[]ApiNVD20WeaknessExtended**](ApiNVD20WeaknessExtended.md) |  | [optional] 

## Methods

### NewApiNVD20CVEExtended

`func NewApiNVD20CVEExtended() *ApiNVD20CVEExtended`

NewApiNVD20CVEExtended instantiates a new ApiNVD20CVEExtended object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiNVD20CVEExtendedWithDefaults

`func NewApiNVD20CVEExtendedWithDefaults() *ApiNVD20CVEExtended`

NewApiNVD20CVEExtendedWithDefaults instantiates a new ApiNVD20CVEExtended object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetALIAS

`func (o *ApiNVD20CVEExtended) GetALIAS() string`

GetALIAS returns the ALIAS field if non-nil, zero value otherwise.

### GetALIASOk

`func (o *ApiNVD20CVEExtended) GetALIASOk() (*string, bool)`

GetALIASOk returns a tuple with the ALIAS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetALIAS

`func (o *ApiNVD20CVEExtended) SetALIAS(v string)`

SetALIAS sets ALIAS field to given value.

### HasALIAS

`func (o *ApiNVD20CVEExtended) HasALIAS() bool`

HasALIAS returns a boolean if a field has been set.

### GetSTATUS

`func (o *ApiNVD20CVEExtended) GetSTATUS() string`

GetSTATUS returns the STATUS field if non-nil, zero value otherwise.

### GetSTATUSOk

`func (o *ApiNVD20CVEExtended) GetSTATUSOk() (*string, bool)`

GetSTATUSOk returns a tuple with the STATUS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSTATUS

`func (o *ApiNVD20CVEExtended) SetSTATUS(v string)`

SetSTATUS sets STATUS field to given value.

### HasSTATUS

`func (o *ApiNVD20CVEExtended) HasSTATUS() bool`

HasSTATUS returns a boolean if a field has been set.

### GetTimestamp

`func (o *ApiNVD20CVEExtended) GetTimestamp() string`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ApiNVD20CVEExtended) GetTimestampOk() (*string, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ApiNVD20CVEExtended) SetTimestamp(v string)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *ApiNVD20CVEExtended) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetCategorization

`func (o *ApiNVD20CVEExtended) GetCategorization() ApiCategorizationExtended`

GetCategorization returns the Categorization field if non-nil, zero value otherwise.

### GetCategorizationOk

`func (o *ApiNVD20CVEExtended) GetCategorizationOk() (*ApiCategorizationExtended, bool)`

GetCategorizationOk returns a tuple with the Categorization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategorization

`func (o *ApiNVD20CVEExtended) SetCategorization(v ApiCategorizationExtended)`

SetCategorization sets Categorization field to given value.

### HasCategorization

`func (o *ApiNVD20CVEExtended) HasCategorization() bool`

HasCategorization returns a boolean if a field has been set.

### GetCisaActionDue

`func (o *ApiNVD20CVEExtended) GetCisaActionDue() string`

GetCisaActionDue returns the CisaActionDue field if non-nil, zero value otherwise.

### GetCisaActionDueOk

`func (o *ApiNVD20CVEExtended) GetCisaActionDueOk() (*string, bool)`

GetCisaActionDueOk returns a tuple with the CisaActionDue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCisaActionDue

`func (o *ApiNVD20CVEExtended) SetCisaActionDue(v string)`

SetCisaActionDue sets CisaActionDue field to given value.

### HasCisaActionDue

`func (o *ApiNVD20CVEExtended) HasCisaActionDue() bool`

HasCisaActionDue returns a boolean if a field has been set.

### GetCisaExploitAdd

`func (o *ApiNVD20CVEExtended) GetCisaExploitAdd() string`

GetCisaExploitAdd returns the CisaExploitAdd field if non-nil, zero value otherwise.

### GetCisaExploitAddOk

`func (o *ApiNVD20CVEExtended) GetCisaExploitAddOk() (*string, bool)`

GetCisaExploitAddOk returns a tuple with the CisaExploitAdd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCisaExploitAdd

`func (o *ApiNVD20CVEExtended) SetCisaExploitAdd(v string)`

SetCisaExploitAdd sets CisaExploitAdd field to given value.

### HasCisaExploitAdd

`func (o *ApiNVD20CVEExtended) HasCisaExploitAdd() bool`

HasCisaExploitAdd returns a boolean if a field has been set.

### GetCisaRequiredAction

`func (o *ApiNVD20CVEExtended) GetCisaRequiredAction() string`

GetCisaRequiredAction returns the CisaRequiredAction field if non-nil, zero value otherwise.

### GetCisaRequiredActionOk

`func (o *ApiNVD20CVEExtended) GetCisaRequiredActionOk() (*string, bool)`

GetCisaRequiredActionOk returns a tuple with the CisaRequiredAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCisaRequiredAction

`func (o *ApiNVD20CVEExtended) SetCisaRequiredAction(v string)`

SetCisaRequiredAction sets CisaRequiredAction field to given value.

### HasCisaRequiredAction

`func (o *ApiNVD20CVEExtended) HasCisaRequiredAction() bool`

HasCisaRequiredAction returns a boolean if a field has been set.

### GetCisaVulnerabilityName

`func (o *ApiNVD20CVEExtended) GetCisaVulnerabilityName() string`

GetCisaVulnerabilityName returns the CisaVulnerabilityName field if non-nil, zero value otherwise.

### GetCisaVulnerabilityNameOk

`func (o *ApiNVD20CVEExtended) GetCisaVulnerabilityNameOk() (*string, bool)`

GetCisaVulnerabilityNameOk returns a tuple with the CisaVulnerabilityName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCisaVulnerabilityName

`func (o *ApiNVD20CVEExtended) SetCisaVulnerabilityName(v string)`

SetCisaVulnerabilityName sets CisaVulnerabilityName field to given value.

### HasCisaVulnerabilityName

`func (o *ApiNVD20CVEExtended) HasCisaVulnerabilityName() bool`

HasCisaVulnerabilityName returns a boolean if a field has been set.

### GetConfigurations

`func (o *ApiNVD20CVEExtended) GetConfigurations() []AdvisoryNVD20Configuration`

GetConfigurations returns the Configurations field if non-nil, zero value otherwise.

### GetConfigurationsOk

`func (o *ApiNVD20CVEExtended) GetConfigurationsOk() (*[]AdvisoryNVD20Configuration, bool)`

GetConfigurationsOk returns a tuple with the Configurations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurations

`func (o *ApiNVD20CVEExtended) SetConfigurations(v []AdvisoryNVD20Configuration)`

SetConfigurations sets Configurations field to given value.

### HasConfigurations

`func (o *ApiNVD20CVEExtended) HasConfigurations() bool`

HasConfigurations returns a boolean if a field has been set.

### GetDateAdded

`func (o *ApiNVD20CVEExtended) GetDateAdded() string`

GetDateAdded returns the DateAdded field if non-nil, zero value otherwise.

### GetDateAddedOk

`func (o *ApiNVD20CVEExtended) GetDateAddedOk() (*string, bool)`

GetDateAddedOk returns a tuple with the DateAdded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateAdded

`func (o *ApiNVD20CVEExtended) SetDateAdded(v string)`

SetDateAdded sets DateAdded field to given value.

### HasDateAdded

`func (o *ApiNVD20CVEExtended) HasDateAdded() bool`

HasDateAdded returns a boolean if a field has been set.

### GetDescriptions

`func (o *ApiNVD20CVEExtended) GetDescriptions() []ApiNVD20Description`

GetDescriptions returns the Descriptions field if non-nil, zero value otherwise.

### GetDescriptionsOk

`func (o *ApiNVD20CVEExtended) GetDescriptionsOk() (*[]ApiNVD20Description, bool)`

GetDescriptionsOk returns a tuple with the Descriptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescriptions

`func (o *ApiNVD20CVEExtended) SetDescriptions(v []ApiNVD20Description)`

SetDescriptions sets Descriptions field to given value.

### HasDescriptions

`func (o *ApiNVD20CVEExtended) HasDescriptions() bool`

HasDescriptions returns a boolean if a field has been set.

### GetDocumentGenerationDate

`func (o *ApiNVD20CVEExtended) GetDocumentGenerationDate() string`

GetDocumentGenerationDate returns the DocumentGenerationDate field if non-nil, zero value otherwise.

### GetDocumentGenerationDateOk

`func (o *ApiNVD20CVEExtended) GetDocumentGenerationDateOk() (*string, bool)`

GetDocumentGenerationDateOk returns a tuple with the DocumentGenerationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentGenerationDate

`func (o *ApiNVD20CVEExtended) SetDocumentGenerationDate(v string)`

SetDocumentGenerationDate sets DocumentGenerationDate field to given value.

### HasDocumentGenerationDate

`func (o *ApiNVD20CVEExtended) HasDocumentGenerationDate() bool`

HasDocumentGenerationDate returns a boolean if a field has been set.

### GetEvaluatorComment

`func (o *ApiNVD20CVEExtended) GetEvaluatorComment() string`

GetEvaluatorComment returns the EvaluatorComment field if non-nil, zero value otherwise.

### GetEvaluatorCommentOk

`func (o *ApiNVD20CVEExtended) GetEvaluatorCommentOk() (*string, bool)`

GetEvaluatorCommentOk returns a tuple with the EvaluatorComment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvaluatorComment

`func (o *ApiNVD20CVEExtended) SetEvaluatorComment(v string)`

SetEvaluatorComment sets EvaluatorComment field to given value.

### HasEvaluatorComment

`func (o *ApiNVD20CVEExtended) HasEvaluatorComment() bool`

HasEvaluatorComment returns a boolean if a field has been set.

### GetEvaluatorImpact

`func (o *ApiNVD20CVEExtended) GetEvaluatorImpact() string`

GetEvaluatorImpact returns the EvaluatorImpact field if non-nil, zero value otherwise.

### GetEvaluatorImpactOk

`func (o *ApiNVD20CVEExtended) GetEvaluatorImpactOk() (*string, bool)`

GetEvaluatorImpactOk returns a tuple with the EvaluatorImpact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvaluatorImpact

`func (o *ApiNVD20CVEExtended) SetEvaluatorImpact(v string)`

SetEvaluatorImpact sets EvaluatorImpact field to given value.

### HasEvaluatorImpact

`func (o *ApiNVD20CVEExtended) HasEvaluatorImpact() bool`

HasEvaluatorImpact returns a boolean if a field has been set.

### GetEvaluatorSolution

`func (o *ApiNVD20CVEExtended) GetEvaluatorSolution() string`

GetEvaluatorSolution returns the EvaluatorSolution field if non-nil, zero value otherwise.

### GetEvaluatorSolutionOk

`func (o *ApiNVD20CVEExtended) GetEvaluatorSolutionOk() (*string, bool)`

GetEvaluatorSolutionOk returns a tuple with the EvaluatorSolution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvaluatorSolution

`func (o *ApiNVD20CVEExtended) SetEvaluatorSolution(v string)`

SetEvaluatorSolution sets EvaluatorSolution field to given value.

### HasEvaluatorSolution

`func (o *ApiNVD20CVEExtended) HasEvaluatorSolution() bool`

HasEvaluatorSolution returns a boolean if a field has been set.

### GetId

`func (o *ApiNVD20CVEExtended) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApiNVD20CVEExtended) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApiNVD20CVEExtended) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ApiNVD20CVEExtended) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastModified

`func (o *ApiNVD20CVEExtended) GetLastModified() string`

GetLastModified returns the LastModified field if non-nil, zero value otherwise.

### GetLastModifiedOk

`func (o *ApiNVD20CVEExtended) GetLastModifiedOk() (*string, bool)`

GetLastModifiedOk returns a tuple with the LastModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModified

`func (o *ApiNVD20CVEExtended) SetLastModified(v string)`

SetLastModified sets LastModified field to given value.

### HasLastModified

`func (o *ApiNVD20CVEExtended) HasLastModified() bool`

HasLastModified returns a boolean if a field has been set.

### GetMetrics

`func (o *ApiNVD20CVEExtended) GetMetrics() ApiNVD20MetricExtended`

GetMetrics returns the Metrics field if non-nil, zero value otherwise.

### GetMetricsOk

`func (o *ApiNVD20CVEExtended) GetMetricsOk() (*ApiNVD20MetricExtended, bool)`

GetMetricsOk returns a tuple with the Metrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetrics

`func (o *ApiNVD20CVEExtended) SetMetrics(v ApiNVD20MetricExtended)`

SetMetrics sets Metrics field to given value.

### HasMetrics

`func (o *ApiNVD20CVEExtended) HasMetrics() bool`

HasMetrics returns a boolean if a field has been set.

### GetMitreAttackTechniques

`func (o *ApiNVD20CVEExtended) GetMitreAttackTechniques() []ApiMitreAttackTech`

GetMitreAttackTechniques returns the MitreAttackTechniques field if non-nil, zero value otherwise.

### GetMitreAttackTechniquesOk

`func (o *ApiNVD20CVEExtended) GetMitreAttackTechniquesOk() (*[]ApiMitreAttackTech, bool)`

GetMitreAttackTechniquesOk returns a tuple with the MitreAttackTechniques field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMitreAttackTechniques

`func (o *ApiNVD20CVEExtended) SetMitreAttackTechniques(v []ApiMitreAttackTech)`

SetMitreAttackTechniques sets MitreAttackTechniques field to given value.

### HasMitreAttackTechniques

`func (o *ApiNVD20CVEExtended) HasMitreAttackTechniques() bool`

HasMitreAttackTechniques returns a boolean if a field has been set.

### GetPublished

`func (o *ApiNVD20CVEExtended) GetPublished() string`

GetPublished returns the Published field if non-nil, zero value otherwise.

### GetPublishedOk

`func (o *ApiNVD20CVEExtended) GetPublishedOk() (*string, bool)`

GetPublishedOk returns a tuple with the Published field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublished

`func (o *ApiNVD20CVEExtended) SetPublished(v string)`

SetPublished sets Published field to given value.

### HasPublished

`func (o *ApiNVD20CVEExtended) HasPublished() bool`

HasPublished returns a boolean if a field has been set.

### GetReferences

`func (o *ApiNVD20CVEExtended) GetReferences() []ApiNVD20ReferenceExtended`

GetReferences returns the References field if non-nil, zero value otherwise.

### GetReferencesOk

`func (o *ApiNVD20CVEExtended) GetReferencesOk() (*[]ApiNVD20ReferenceExtended, bool)`

GetReferencesOk returns a tuple with the References field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferences

`func (o *ApiNVD20CVEExtended) SetReferences(v []ApiNVD20ReferenceExtended)`

SetReferences sets References field to given value.

### HasReferences

`func (o *ApiNVD20CVEExtended) HasReferences() bool`

HasReferences returns a boolean if a field has been set.

### GetRelatedAttackPatterns

`func (o *ApiNVD20CVEExtended) GetRelatedAttackPatterns() []ApiRelatedAttackPattern`

GetRelatedAttackPatterns returns the RelatedAttackPatterns field if non-nil, zero value otherwise.

### GetRelatedAttackPatternsOk

`func (o *ApiNVD20CVEExtended) GetRelatedAttackPatternsOk() (*[]ApiRelatedAttackPattern, bool)`

GetRelatedAttackPatternsOk returns a tuple with the RelatedAttackPatterns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelatedAttackPatterns

`func (o *ApiNVD20CVEExtended) SetRelatedAttackPatterns(v []ApiRelatedAttackPattern)`

SetRelatedAttackPatterns sets RelatedAttackPatterns field to given value.

### HasRelatedAttackPatterns

`func (o *ApiNVD20CVEExtended) HasRelatedAttackPatterns() bool`

HasRelatedAttackPatterns returns a boolean if a field has been set.

### GetSourceIdentifier

`func (o *ApiNVD20CVEExtended) GetSourceIdentifier() string`

GetSourceIdentifier returns the SourceIdentifier field if non-nil, zero value otherwise.

### GetSourceIdentifierOk

`func (o *ApiNVD20CVEExtended) GetSourceIdentifierOk() (*string, bool)`

GetSourceIdentifierOk returns a tuple with the SourceIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceIdentifier

`func (o *ApiNVD20CVEExtended) SetSourceIdentifier(v string)`

SetSourceIdentifier sets SourceIdentifier field to given value.

### HasSourceIdentifier

`func (o *ApiNVD20CVEExtended) HasSourceIdentifier() bool`

HasSourceIdentifier returns a boolean if a field has been set.

### GetVcConfigurations

`func (o *ApiNVD20CVEExtended) GetVcConfigurations() []AdvisoryNVD20Configuration`

GetVcConfigurations returns the VcConfigurations field if non-nil, zero value otherwise.

### GetVcConfigurationsOk

`func (o *ApiNVD20CVEExtended) GetVcConfigurationsOk() (*[]AdvisoryNVD20Configuration, bool)`

GetVcConfigurationsOk returns a tuple with the VcConfigurations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcConfigurations

`func (o *ApiNVD20CVEExtended) SetVcConfigurations(v []AdvisoryNVD20Configuration)`

SetVcConfigurations sets VcConfigurations field to given value.

### HasVcConfigurations

`func (o *ApiNVD20CVEExtended) HasVcConfigurations() bool`

HasVcConfigurations returns a boolean if a field has been set.

### GetVcVulnerableCPEs

`func (o *ApiNVD20CVEExtended) GetVcVulnerableCPEs() []string`

GetVcVulnerableCPEs returns the VcVulnerableCPEs field if non-nil, zero value otherwise.

### GetVcVulnerableCPEsOk

`func (o *ApiNVD20CVEExtended) GetVcVulnerableCPEsOk() (*[]string, bool)`

GetVcVulnerableCPEsOk returns a tuple with the VcVulnerableCPEs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcVulnerableCPEs

`func (o *ApiNVD20CVEExtended) SetVcVulnerableCPEs(v []string)`

SetVcVulnerableCPEs sets VcVulnerableCPEs field to given value.

### HasVcVulnerableCPEs

`func (o *ApiNVD20CVEExtended) HasVcVulnerableCPEs() bool`

HasVcVulnerableCPEs returns a boolean if a field has been set.

### GetVendorComments

`func (o *ApiNVD20CVEExtended) GetVendorComments() []ApiNVD20VendorComment`

GetVendorComments returns the VendorComments field if non-nil, zero value otherwise.

### GetVendorCommentsOk

`func (o *ApiNVD20CVEExtended) GetVendorCommentsOk() (*[]ApiNVD20VendorComment, bool)`

GetVendorCommentsOk returns a tuple with the VendorComments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorComments

`func (o *ApiNVD20CVEExtended) SetVendorComments(v []ApiNVD20VendorComment)`

SetVendorComments sets VendorComments field to given value.

### HasVendorComments

`func (o *ApiNVD20CVEExtended) HasVendorComments() bool`

HasVendorComments returns a boolean if a field has been set.

### GetVulnStatus

`func (o *ApiNVD20CVEExtended) GetVulnStatus() string`

GetVulnStatus returns the VulnStatus field if non-nil, zero value otherwise.

### GetVulnStatusOk

`func (o *ApiNVD20CVEExtended) GetVulnStatusOk() (*string, bool)`

GetVulnStatusOk returns a tuple with the VulnStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVulnStatus

`func (o *ApiNVD20CVEExtended) SetVulnStatus(v string)`

SetVulnStatus sets VulnStatus field to given value.

### HasVulnStatus

`func (o *ApiNVD20CVEExtended) HasVulnStatus() bool`

HasVulnStatus returns a boolean if a field has been set.

### GetVulncheckKEVExploitAdd

`func (o *ApiNVD20CVEExtended) GetVulncheckKEVExploitAdd() string`

GetVulncheckKEVExploitAdd returns the VulncheckKEVExploitAdd field if non-nil, zero value otherwise.

### GetVulncheckKEVExploitAddOk

`func (o *ApiNVD20CVEExtended) GetVulncheckKEVExploitAddOk() (*string, bool)`

GetVulncheckKEVExploitAddOk returns a tuple with the VulncheckKEVExploitAdd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVulncheckKEVExploitAdd

`func (o *ApiNVD20CVEExtended) SetVulncheckKEVExploitAdd(v string)`

SetVulncheckKEVExploitAdd sets VulncheckKEVExploitAdd field to given value.

### HasVulncheckKEVExploitAdd

`func (o *ApiNVD20CVEExtended) HasVulncheckKEVExploitAdd() bool`

HasVulncheckKEVExploitAdd returns a boolean if a field has been set.

### GetVulnerableCPEs

`func (o *ApiNVD20CVEExtended) GetVulnerableCPEs() []string`

GetVulnerableCPEs returns the VulnerableCPEs field if non-nil, zero value otherwise.

### GetVulnerableCPEsOk

`func (o *ApiNVD20CVEExtended) GetVulnerableCPEsOk() (*[]string, bool)`

GetVulnerableCPEsOk returns a tuple with the VulnerableCPEs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVulnerableCPEs

`func (o *ApiNVD20CVEExtended) SetVulnerableCPEs(v []string)`

SetVulnerableCPEs sets VulnerableCPEs field to given value.

### HasVulnerableCPEs

`func (o *ApiNVD20CVEExtended) HasVulnerableCPEs() bool`

HasVulnerableCPEs returns a boolean if a field has been set.

### GetWeaknesses

`func (o *ApiNVD20CVEExtended) GetWeaknesses() []ApiNVD20WeaknessExtended`

GetWeaknesses returns the Weaknesses field if non-nil, zero value otherwise.

### GetWeaknessesOk

`func (o *ApiNVD20CVEExtended) GetWeaknessesOk() (*[]ApiNVD20WeaknessExtended, bool)`

GetWeaknessesOk returns a tuple with the Weaknesses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeaknesses

`func (o *ApiNVD20CVEExtended) SetWeaknesses(v []ApiNVD20WeaknessExtended)`

SetWeaknesses sets Weaknesses field to given value.

### HasWeaknesses

`func (o *ApiNVD20CVEExtended) HasWeaknesses() bool`

HasWeaknesses returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


