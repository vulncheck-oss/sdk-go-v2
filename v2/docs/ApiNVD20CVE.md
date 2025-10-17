# ApiNVD20CVE

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CisaActionDue** | Pointer to **string** |  | [optional] 
**CisaExploitAdd** | Pointer to **string** |  | [optional] 
**CisaRequiredAction** | Pointer to **string** |  | [optional] 
**CisaVulnerabilityName** | Pointer to **string** |  | [optional] 
**Configurations** | Pointer to [**[]AdvisoryNVD20Configuration**](AdvisoryNVD20Configuration.md) |  | [optional] 
**Descriptions** | Pointer to [**[]ApiNVD20Description**](ApiNVD20Description.md) |  | [optional] 
**EvaluatorComment** | Pointer to **string** |  | [optional] 
**EvaluatorImpact** | Pointer to **string** |  | [optional] 
**EvaluatorSolution** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**LastModified** | Pointer to **string** |  | [optional] 
**Metrics** | Pointer to [**ApiNVD20Metric**](ApiNVD20Metric.md) |  | [optional] 
**Published** | Pointer to **string** |  | [optional] 
**References** | Pointer to [**[]ApiNVD20Reference**](ApiNVD20Reference.md) |  | [optional] 
**SourceIdentifier** | Pointer to **string** |  | [optional] 
**VcConfigurations** | Pointer to [**[]AdvisoryNVD20Configuration**](AdvisoryNVD20Configuration.md) |  | [optional] 
**VcVulnerableCPEs** | Pointer to **[]string** |  | [optional] 
**VendorComments** | Pointer to [**[]ApiNVD20VendorComment**](ApiNVD20VendorComment.md) |  | [optional] 
**VulnStatus** | Pointer to **string** |  | [optional] 
**Weaknesses** | Pointer to [**[]ApiNVD20Weakness**](ApiNVD20Weakness.md) |  | [optional] 

## Methods

### NewApiNVD20CVE

`func NewApiNVD20CVE() *ApiNVD20CVE`

NewApiNVD20CVE instantiates a new ApiNVD20CVE object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiNVD20CVEWithDefaults

`func NewApiNVD20CVEWithDefaults() *ApiNVD20CVE`

NewApiNVD20CVEWithDefaults instantiates a new ApiNVD20CVE object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCisaActionDue

`func (o *ApiNVD20CVE) GetCisaActionDue() string`

GetCisaActionDue returns the CisaActionDue field if non-nil, zero value otherwise.

### GetCisaActionDueOk

`func (o *ApiNVD20CVE) GetCisaActionDueOk() (*string, bool)`

GetCisaActionDueOk returns a tuple with the CisaActionDue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCisaActionDue

`func (o *ApiNVD20CVE) SetCisaActionDue(v string)`

SetCisaActionDue sets CisaActionDue field to given value.

### HasCisaActionDue

`func (o *ApiNVD20CVE) HasCisaActionDue() bool`

HasCisaActionDue returns a boolean if a field has been set.

### GetCisaExploitAdd

`func (o *ApiNVD20CVE) GetCisaExploitAdd() string`

GetCisaExploitAdd returns the CisaExploitAdd field if non-nil, zero value otherwise.

### GetCisaExploitAddOk

`func (o *ApiNVD20CVE) GetCisaExploitAddOk() (*string, bool)`

GetCisaExploitAddOk returns a tuple with the CisaExploitAdd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCisaExploitAdd

`func (o *ApiNVD20CVE) SetCisaExploitAdd(v string)`

SetCisaExploitAdd sets CisaExploitAdd field to given value.

### HasCisaExploitAdd

`func (o *ApiNVD20CVE) HasCisaExploitAdd() bool`

HasCisaExploitAdd returns a boolean if a field has been set.

### GetCisaRequiredAction

`func (o *ApiNVD20CVE) GetCisaRequiredAction() string`

GetCisaRequiredAction returns the CisaRequiredAction field if non-nil, zero value otherwise.

### GetCisaRequiredActionOk

`func (o *ApiNVD20CVE) GetCisaRequiredActionOk() (*string, bool)`

GetCisaRequiredActionOk returns a tuple with the CisaRequiredAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCisaRequiredAction

`func (o *ApiNVD20CVE) SetCisaRequiredAction(v string)`

SetCisaRequiredAction sets CisaRequiredAction field to given value.

### HasCisaRequiredAction

`func (o *ApiNVD20CVE) HasCisaRequiredAction() bool`

HasCisaRequiredAction returns a boolean if a field has been set.

### GetCisaVulnerabilityName

`func (o *ApiNVD20CVE) GetCisaVulnerabilityName() string`

GetCisaVulnerabilityName returns the CisaVulnerabilityName field if non-nil, zero value otherwise.

### GetCisaVulnerabilityNameOk

`func (o *ApiNVD20CVE) GetCisaVulnerabilityNameOk() (*string, bool)`

GetCisaVulnerabilityNameOk returns a tuple with the CisaVulnerabilityName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCisaVulnerabilityName

`func (o *ApiNVD20CVE) SetCisaVulnerabilityName(v string)`

SetCisaVulnerabilityName sets CisaVulnerabilityName field to given value.

### HasCisaVulnerabilityName

`func (o *ApiNVD20CVE) HasCisaVulnerabilityName() bool`

HasCisaVulnerabilityName returns a boolean if a field has been set.

### GetConfigurations

`func (o *ApiNVD20CVE) GetConfigurations() []AdvisoryNVD20Configuration`

GetConfigurations returns the Configurations field if non-nil, zero value otherwise.

### GetConfigurationsOk

`func (o *ApiNVD20CVE) GetConfigurationsOk() (*[]AdvisoryNVD20Configuration, bool)`

GetConfigurationsOk returns a tuple with the Configurations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurations

`func (o *ApiNVD20CVE) SetConfigurations(v []AdvisoryNVD20Configuration)`

SetConfigurations sets Configurations field to given value.

### HasConfigurations

`func (o *ApiNVD20CVE) HasConfigurations() bool`

HasConfigurations returns a boolean if a field has been set.

### GetDescriptions

`func (o *ApiNVD20CVE) GetDescriptions() []ApiNVD20Description`

GetDescriptions returns the Descriptions field if non-nil, zero value otherwise.

### GetDescriptionsOk

`func (o *ApiNVD20CVE) GetDescriptionsOk() (*[]ApiNVD20Description, bool)`

GetDescriptionsOk returns a tuple with the Descriptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescriptions

`func (o *ApiNVD20CVE) SetDescriptions(v []ApiNVD20Description)`

SetDescriptions sets Descriptions field to given value.

### HasDescriptions

`func (o *ApiNVD20CVE) HasDescriptions() bool`

HasDescriptions returns a boolean if a field has been set.

### GetEvaluatorComment

`func (o *ApiNVD20CVE) GetEvaluatorComment() string`

GetEvaluatorComment returns the EvaluatorComment field if non-nil, zero value otherwise.

### GetEvaluatorCommentOk

`func (o *ApiNVD20CVE) GetEvaluatorCommentOk() (*string, bool)`

GetEvaluatorCommentOk returns a tuple with the EvaluatorComment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvaluatorComment

`func (o *ApiNVD20CVE) SetEvaluatorComment(v string)`

SetEvaluatorComment sets EvaluatorComment field to given value.

### HasEvaluatorComment

`func (o *ApiNVD20CVE) HasEvaluatorComment() bool`

HasEvaluatorComment returns a boolean if a field has been set.

### GetEvaluatorImpact

`func (o *ApiNVD20CVE) GetEvaluatorImpact() string`

GetEvaluatorImpact returns the EvaluatorImpact field if non-nil, zero value otherwise.

### GetEvaluatorImpactOk

`func (o *ApiNVD20CVE) GetEvaluatorImpactOk() (*string, bool)`

GetEvaluatorImpactOk returns a tuple with the EvaluatorImpact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvaluatorImpact

`func (o *ApiNVD20CVE) SetEvaluatorImpact(v string)`

SetEvaluatorImpact sets EvaluatorImpact field to given value.

### HasEvaluatorImpact

`func (o *ApiNVD20CVE) HasEvaluatorImpact() bool`

HasEvaluatorImpact returns a boolean if a field has been set.

### GetEvaluatorSolution

`func (o *ApiNVD20CVE) GetEvaluatorSolution() string`

GetEvaluatorSolution returns the EvaluatorSolution field if non-nil, zero value otherwise.

### GetEvaluatorSolutionOk

`func (o *ApiNVD20CVE) GetEvaluatorSolutionOk() (*string, bool)`

GetEvaluatorSolutionOk returns a tuple with the EvaluatorSolution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvaluatorSolution

`func (o *ApiNVD20CVE) SetEvaluatorSolution(v string)`

SetEvaluatorSolution sets EvaluatorSolution field to given value.

### HasEvaluatorSolution

`func (o *ApiNVD20CVE) HasEvaluatorSolution() bool`

HasEvaluatorSolution returns a boolean if a field has been set.

### GetId

`func (o *ApiNVD20CVE) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApiNVD20CVE) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApiNVD20CVE) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ApiNVD20CVE) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastModified

`func (o *ApiNVD20CVE) GetLastModified() string`

GetLastModified returns the LastModified field if non-nil, zero value otherwise.

### GetLastModifiedOk

`func (o *ApiNVD20CVE) GetLastModifiedOk() (*string, bool)`

GetLastModifiedOk returns a tuple with the LastModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModified

`func (o *ApiNVD20CVE) SetLastModified(v string)`

SetLastModified sets LastModified field to given value.

### HasLastModified

`func (o *ApiNVD20CVE) HasLastModified() bool`

HasLastModified returns a boolean if a field has been set.

### GetMetrics

`func (o *ApiNVD20CVE) GetMetrics() ApiNVD20Metric`

GetMetrics returns the Metrics field if non-nil, zero value otherwise.

### GetMetricsOk

`func (o *ApiNVD20CVE) GetMetricsOk() (*ApiNVD20Metric, bool)`

GetMetricsOk returns a tuple with the Metrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetrics

`func (o *ApiNVD20CVE) SetMetrics(v ApiNVD20Metric)`

SetMetrics sets Metrics field to given value.

### HasMetrics

`func (o *ApiNVD20CVE) HasMetrics() bool`

HasMetrics returns a boolean if a field has been set.

### GetPublished

`func (o *ApiNVD20CVE) GetPublished() string`

GetPublished returns the Published field if non-nil, zero value otherwise.

### GetPublishedOk

`func (o *ApiNVD20CVE) GetPublishedOk() (*string, bool)`

GetPublishedOk returns a tuple with the Published field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublished

`func (o *ApiNVD20CVE) SetPublished(v string)`

SetPublished sets Published field to given value.

### HasPublished

`func (o *ApiNVD20CVE) HasPublished() bool`

HasPublished returns a boolean if a field has been set.

### GetReferences

`func (o *ApiNVD20CVE) GetReferences() []ApiNVD20Reference`

GetReferences returns the References field if non-nil, zero value otherwise.

### GetReferencesOk

`func (o *ApiNVD20CVE) GetReferencesOk() (*[]ApiNVD20Reference, bool)`

GetReferencesOk returns a tuple with the References field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferences

`func (o *ApiNVD20CVE) SetReferences(v []ApiNVD20Reference)`

SetReferences sets References field to given value.

### HasReferences

`func (o *ApiNVD20CVE) HasReferences() bool`

HasReferences returns a boolean if a field has been set.

### GetSourceIdentifier

`func (o *ApiNVD20CVE) GetSourceIdentifier() string`

GetSourceIdentifier returns the SourceIdentifier field if non-nil, zero value otherwise.

### GetSourceIdentifierOk

`func (o *ApiNVD20CVE) GetSourceIdentifierOk() (*string, bool)`

GetSourceIdentifierOk returns a tuple with the SourceIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceIdentifier

`func (o *ApiNVD20CVE) SetSourceIdentifier(v string)`

SetSourceIdentifier sets SourceIdentifier field to given value.

### HasSourceIdentifier

`func (o *ApiNVD20CVE) HasSourceIdentifier() bool`

HasSourceIdentifier returns a boolean if a field has been set.

### GetVcConfigurations

`func (o *ApiNVD20CVE) GetVcConfigurations() []AdvisoryNVD20Configuration`

GetVcConfigurations returns the VcConfigurations field if non-nil, zero value otherwise.

### GetVcConfigurationsOk

`func (o *ApiNVD20CVE) GetVcConfigurationsOk() (*[]AdvisoryNVD20Configuration, bool)`

GetVcConfigurationsOk returns a tuple with the VcConfigurations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcConfigurations

`func (o *ApiNVD20CVE) SetVcConfigurations(v []AdvisoryNVD20Configuration)`

SetVcConfigurations sets VcConfigurations field to given value.

### HasVcConfigurations

`func (o *ApiNVD20CVE) HasVcConfigurations() bool`

HasVcConfigurations returns a boolean if a field has been set.

### GetVcVulnerableCPEs

`func (o *ApiNVD20CVE) GetVcVulnerableCPEs() []string`

GetVcVulnerableCPEs returns the VcVulnerableCPEs field if non-nil, zero value otherwise.

### GetVcVulnerableCPEsOk

`func (o *ApiNVD20CVE) GetVcVulnerableCPEsOk() (*[]string, bool)`

GetVcVulnerableCPEsOk returns a tuple with the VcVulnerableCPEs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcVulnerableCPEs

`func (o *ApiNVD20CVE) SetVcVulnerableCPEs(v []string)`

SetVcVulnerableCPEs sets VcVulnerableCPEs field to given value.

### HasVcVulnerableCPEs

`func (o *ApiNVD20CVE) HasVcVulnerableCPEs() bool`

HasVcVulnerableCPEs returns a boolean if a field has been set.

### GetVendorComments

`func (o *ApiNVD20CVE) GetVendorComments() []ApiNVD20VendorComment`

GetVendorComments returns the VendorComments field if non-nil, zero value otherwise.

### GetVendorCommentsOk

`func (o *ApiNVD20CVE) GetVendorCommentsOk() (*[]ApiNVD20VendorComment, bool)`

GetVendorCommentsOk returns a tuple with the VendorComments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorComments

`func (o *ApiNVD20CVE) SetVendorComments(v []ApiNVD20VendorComment)`

SetVendorComments sets VendorComments field to given value.

### HasVendorComments

`func (o *ApiNVD20CVE) HasVendorComments() bool`

HasVendorComments returns a boolean if a field has been set.

### GetVulnStatus

`func (o *ApiNVD20CVE) GetVulnStatus() string`

GetVulnStatus returns the VulnStatus field if non-nil, zero value otherwise.

### GetVulnStatusOk

`func (o *ApiNVD20CVE) GetVulnStatusOk() (*string, bool)`

GetVulnStatusOk returns a tuple with the VulnStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVulnStatus

`func (o *ApiNVD20CVE) SetVulnStatus(v string)`

SetVulnStatus sets VulnStatus field to given value.

### HasVulnStatus

`func (o *ApiNVD20CVE) HasVulnStatus() bool`

HasVulnStatus returns a boolean if a field has been set.

### GetWeaknesses

`func (o *ApiNVD20CVE) GetWeaknesses() []ApiNVD20Weakness`

GetWeaknesses returns the Weaknesses field if non-nil, zero value otherwise.

### GetWeaknessesOk

`func (o *ApiNVD20CVE) GetWeaknessesOk() (*[]ApiNVD20Weakness, bool)`

GetWeaknessesOk returns a tuple with the Weaknesses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeaknesses

`func (o *ApiNVD20CVE) SetWeaknesses(v []ApiNVD20Weakness)`

SetWeaknesses sets Weaknesses field to given value.

### HasWeaknesses

`func (o *ApiNVD20CVE) HasWeaknesses() bool`

HasWeaknesses returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


