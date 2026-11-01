# NvdcvesNvdCveExtended

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Affected** | Pointer to [**[]ApiNVD20Affected**](ApiNVD20Affected.md) |  | [optional] 
**CisaActionDue** | Pointer to **string** |  | [optional] 
**CisaExploitAdd** | Pointer to **string** |  | [optional] 
**CisaRequiredAction** | Pointer to **string** |  | [optional] 
**CisaVulnerabilityName** | Pointer to **string** |  | [optional] 
**Configurations** | Pointer to [**[]AdvisoryNVD20Configuration**](AdvisoryNVD20Configuration.md) |  | [optional] 
**CveTags** | Pointer to [**[]ApiNVD20CVETag**](ApiNVD20CVETag.md) |  | [optional] 
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
**VcConfigurationsWithPlatforms** | Pointer to [**[]AdvisoryNVD20Configuration**](AdvisoryNVD20Configuration.md) | VCConfigurationsWithPlatforms is VCConfigurations with NVD-style platform (\&quot;running on\&quot;) nodes: a configuration whose product needs a platform is AND{the vulnerable node, a node of vulnerable:false platform CPEs}. Only set when at least one configuration has a platform. | [optional] 
**VcVulnerableCPEs** | Pointer to **[]string** |  | [optional] 
**VendorComments** | Pointer to [**[]ApiNVD20VendorComment**](ApiNVD20VendorComment.md) |  | [optional] 
**VulnStatus** | Pointer to **string** |  | [optional] 
**Weaknesses** | Pointer to [**[]ApiNVD20Weakness**](ApiNVD20Weakness.md) |  | [optional] 

## Methods

### NewNvdcvesNvdCveExtended

`func NewNvdcvesNvdCveExtended() *NvdcvesNvdCveExtended`

NewNvdcvesNvdCveExtended instantiates a new NvdcvesNvdCveExtended object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNvdcvesNvdCveExtendedWithDefaults

`func NewNvdcvesNvdCveExtendedWithDefaults() *NvdcvesNvdCveExtended`

NewNvdcvesNvdCveExtendedWithDefaults instantiates a new NvdcvesNvdCveExtended object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAffected

`func (o *NvdcvesNvdCveExtended) GetAffected() []ApiNVD20Affected`

GetAffected returns the Affected field if non-nil, zero value otherwise.

### GetAffectedOk

`func (o *NvdcvesNvdCveExtended) GetAffectedOk() (*[]ApiNVD20Affected, bool)`

GetAffectedOk returns a tuple with the Affected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffected

`func (o *NvdcvesNvdCveExtended) SetAffected(v []ApiNVD20Affected)`

SetAffected sets Affected field to given value.

### HasAffected

`func (o *NvdcvesNvdCveExtended) HasAffected() bool`

HasAffected returns a boolean if a field has been set.

### GetCisaActionDue

`func (o *NvdcvesNvdCveExtended) GetCisaActionDue() string`

GetCisaActionDue returns the CisaActionDue field if non-nil, zero value otherwise.

### GetCisaActionDueOk

`func (o *NvdcvesNvdCveExtended) GetCisaActionDueOk() (*string, bool)`

GetCisaActionDueOk returns a tuple with the CisaActionDue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCisaActionDue

`func (o *NvdcvesNvdCveExtended) SetCisaActionDue(v string)`

SetCisaActionDue sets CisaActionDue field to given value.

### HasCisaActionDue

`func (o *NvdcvesNvdCveExtended) HasCisaActionDue() bool`

HasCisaActionDue returns a boolean if a field has been set.

### GetCisaExploitAdd

`func (o *NvdcvesNvdCveExtended) GetCisaExploitAdd() string`

GetCisaExploitAdd returns the CisaExploitAdd field if non-nil, zero value otherwise.

### GetCisaExploitAddOk

`func (o *NvdcvesNvdCveExtended) GetCisaExploitAddOk() (*string, bool)`

GetCisaExploitAddOk returns a tuple with the CisaExploitAdd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCisaExploitAdd

`func (o *NvdcvesNvdCveExtended) SetCisaExploitAdd(v string)`

SetCisaExploitAdd sets CisaExploitAdd field to given value.

### HasCisaExploitAdd

`func (o *NvdcvesNvdCveExtended) HasCisaExploitAdd() bool`

HasCisaExploitAdd returns a boolean if a field has been set.

### GetCisaRequiredAction

`func (o *NvdcvesNvdCveExtended) GetCisaRequiredAction() string`

GetCisaRequiredAction returns the CisaRequiredAction field if non-nil, zero value otherwise.

### GetCisaRequiredActionOk

`func (o *NvdcvesNvdCveExtended) GetCisaRequiredActionOk() (*string, bool)`

GetCisaRequiredActionOk returns a tuple with the CisaRequiredAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCisaRequiredAction

`func (o *NvdcvesNvdCveExtended) SetCisaRequiredAction(v string)`

SetCisaRequiredAction sets CisaRequiredAction field to given value.

### HasCisaRequiredAction

`func (o *NvdcvesNvdCveExtended) HasCisaRequiredAction() bool`

HasCisaRequiredAction returns a boolean if a field has been set.

### GetCisaVulnerabilityName

`func (o *NvdcvesNvdCveExtended) GetCisaVulnerabilityName() string`

GetCisaVulnerabilityName returns the CisaVulnerabilityName field if non-nil, zero value otherwise.

### GetCisaVulnerabilityNameOk

`func (o *NvdcvesNvdCveExtended) GetCisaVulnerabilityNameOk() (*string, bool)`

GetCisaVulnerabilityNameOk returns a tuple with the CisaVulnerabilityName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCisaVulnerabilityName

`func (o *NvdcvesNvdCveExtended) SetCisaVulnerabilityName(v string)`

SetCisaVulnerabilityName sets CisaVulnerabilityName field to given value.

### HasCisaVulnerabilityName

`func (o *NvdcvesNvdCveExtended) HasCisaVulnerabilityName() bool`

HasCisaVulnerabilityName returns a boolean if a field has been set.

### GetConfigurations

`func (o *NvdcvesNvdCveExtended) GetConfigurations() []AdvisoryNVD20Configuration`

GetConfigurations returns the Configurations field if non-nil, zero value otherwise.

### GetConfigurationsOk

`func (o *NvdcvesNvdCveExtended) GetConfigurationsOk() (*[]AdvisoryNVD20Configuration, bool)`

GetConfigurationsOk returns a tuple with the Configurations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurations

`func (o *NvdcvesNvdCveExtended) SetConfigurations(v []AdvisoryNVD20Configuration)`

SetConfigurations sets Configurations field to given value.

### HasConfigurations

`func (o *NvdcvesNvdCveExtended) HasConfigurations() bool`

HasConfigurations returns a boolean if a field has been set.

### GetCveTags

`func (o *NvdcvesNvdCveExtended) GetCveTags() []ApiNVD20CVETag`

GetCveTags returns the CveTags field if non-nil, zero value otherwise.

### GetCveTagsOk

`func (o *NvdcvesNvdCveExtended) GetCveTagsOk() (*[]ApiNVD20CVETag, bool)`

GetCveTagsOk returns a tuple with the CveTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCveTags

`func (o *NvdcvesNvdCveExtended) SetCveTags(v []ApiNVD20CVETag)`

SetCveTags sets CveTags field to given value.

### HasCveTags

`func (o *NvdcvesNvdCveExtended) HasCveTags() bool`

HasCveTags returns a boolean if a field has been set.

### GetDescriptions

`func (o *NvdcvesNvdCveExtended) GetDescriptions() []ApiNVD20Description`

GetDescriptions returns the Descriptions field if non-nil, zero value otherwise.

### GetDescriptionsOk

`func (o *NvdcvesNvdCveExtended) GetDescriptionsOk() (*[]ApiNVD20Description, bool)`

GetDescriptionsOk returns a tuple with the Descriptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescriptions

`func (o *NvdcvesNvdCveExtended) SetDescriptions(v []ApiNVD20Description)`

SetDescriptions sets Descriptions field to given value.

### HasDescriptions

`func (o *NvdcvesNvdCveExtended) HasDescriptions() bool`

HasDescriptions returns a boolean if a field has been set.

### GetEvaluatorComment

`func (o *NvdcvesNvdCveExtended) GetEvaluatorComment() string`

GetEvaluatorComment returns the EvaluatorComment field if non-nil, zero value otherwise.

### GetEvaluatorCommentOk

`func (o *NvdcvesNvdCveExtended) GetEvaluatorCommentOk() (*string, bool)`

GetEvaluatorCommentOk returns a tuple with the EvaluatorComment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvaluatorComment

`func (o *NvdcvesNvdCveExtended) SetEvaluatorComment(v string)`

SetEvaluatorComment sets EvaluatorComment field to given value.

### HasEvaluatorComment

`func (o *NvdcvesNvdCveExtended) HasEvaluatorComment() bool`

HasEvaluatorComment returns a boolean if a field has been set.

### GetEvaluatorImpact

`func (o *NvdcvesNvdCveExtended) GetEvaluatorImpact() string`

GetEvaluatorImpact returns the EvaluatorImpact field if non-nil, zero value otherwise.

### GetEvaluatorImpactOk

`func (o *NvdcvesNvdCveExtended) GetEvaluatorImpactOk() (*string, bool)`

GetEvaluatorImpactOk returns a tuple with the EvaluatorImpact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvaluatorImpact

`func (o *NvdcvesNvdCveExtended) SetEvaluatorImpact(v string)`

SetEvaluatorImpact sets EvaluatorImpact field to given value.

### HasEvaluatorImpact

`func (o *NvdcvesNvdCveExtended) HasEvaluatorImpact() bool`

HasEvaluatorImpact returns a boolean if a field has been set.

### GetEvaluatorSolution

`func (o *NvdcvesNvdCveExtended) GetEvaluatorSolution() string`

GetEvaluatorSolution returns the EvaluatorSolution field if non-nil, zero value otherwise.

### GetEvaluatorSolutionOk

`func (o *NvdcvesNvdCveExtended) GetEvaluatorSolutionOk() (*string, bool)`

GetEvaluatorSolutionOk returns a tuple with the EvaluatorSolution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvaluatorSolution

`func (o *NvdcvesNvdCveExtended) SetEvaluatorSolution(v string)`

SetEvaluatorSolution sets EvaluatorSolution field to given value.

### HasEvaluatorSolution

`func (o *NvdcvesNvdCveExtended) HasEvaluatorSolution() bool`

HasEvaluatorSolution returns a boolean if a field has been set.

### GetId

`func (o *NvdcvesNvdCveExtended) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NvdcvesNvdCveExtended) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NvdcvesNvdCveExtended) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *NvdcvesNvdCveExtended) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastModified

`func (o *NvdcvesNvdCveExtended) GetLastModified() string`

GetLastModified returns the LastModified field if non-nil, zero value otherwise.

### GetLastModifiedOk

`func (o *NvdcvesNvdCveExtended) GetLastModifiedOk() (*string, bool)`

GetLastModifiedOk returns a tuple with the LastModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModified

`func (o *NvdcvesNvdCveExtended) SetLastModified(v string)`

SetLastModified sets LastModified field to given value.

### HasLastModified

`func (o *NvdcvesNvdCveExtended) HasLastModified() bool`

HasLastModified returns a boolean if a field has been set.

### GetMetrics

`func (o *NvdcvesNvdCveExtended) GetMetrics() ApiNVD20Metric`

GetMetrics returns the Metrics field if non-nil, zero value otherwise.

### GetMetricsOk

`func (o *NvdcvesNvdCveExtended) GetMetricsOk() (*ApiNVD20Metric, bool)`

GetMetricsOk returns a tuple with the Metrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetrics

`func (o *NvdcvesNvdCveExtended) SetMetrics(v ApiNVD20Metric)`

SetMetrics sets Metrics field to given value.

### HasMetrics

`func (o *NvdcvesNvdCveExtended) HasMetrics() bool`

HasMetrics returns a boolean if a field has been set.

### GetPublished

`func (o *NvdcvesNvdCveExtended) GetPublished() string`

GetPublished returns the Published field if non-nil, zero value otherwise.

### GetPublishedOk

`func (o *NvdcvesNvdCveExtended) GetPublishedOk() (*string, bool)`

GetPublishedOk returns a tuple with the Published field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublished

`func (o *NvdcvesNvdCveExtended) SetPublished(v string)`

SetPublished sets Published field to given value.

### HasPublished

`func (o *NvdcvesNvdCveExtended) HasPublished() bool`

HasPublished returns a boolean if a field has been set.

### GetReferences

`func (o *NvdcvesNvdCveExtended) GetReferences() []ApiNVD20Reference`

GetReferences returns the References field if non-nil, zero value otherwise.

### GetReferencesOk

`func (o *NvdcvesNvdCveExtended) GetReferencesOk() (*[]ApiNVD20Reference, bool)`

GetReferencesOk returns a tuple with the References field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferences

`func (o *NvdcvesNvdCveExtended) SetReferences(v []ApiNVD20Reference)`

SetReferences sets References field to given value.

### HasReferences

`func (o *NvdcvesNvdCveExtended) HasReferences() bool`

HasReferences returns a boolean if a field has been set.

### GetSourceIdentifier

`func (o *NvdcvesNvdCveExtended) GetSourceIdentifier() string`

GetSourceIdentifier returns the SourceIdentifier field if non-nil, zero value otherwise.

### GetSourceIdentifierOk

`func (o *NvdcvesNvdCveExtended) GetSourceIdentifierOk() (*string, bool)`

GetSourceIdentifierOk returns a tuple with the SourceIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceIdentifier

`func (o *NvdcvesNvdCveExtended) SetSourceIdentifier(v string)`

SetSourceIdentifier sets SourceIdentifier field to given value.

### HasSourceIdentifier

`func (o *NvdcvesNvdCveExtended) HasSourceIdentifier() bool`

HasSourceIdentifier returns a boolean if a field has been set.

### GetVcConfigurations

`func (o *NvdcvesNvdCveExtended) GetVcConfigurations() []AdvisoryNVD20Configuration`

GetVcConfigurations returns the VcConfigurations field if non-nil, zero value otherwise.

### GetVcConfigurationsOk

`func (o *NvdcvesNvdCveExtended) GetVcConfigurationsOk() (*[]AdvisoryNVD20Configuration, bool)`

GetVcConfigurationsOk returns a tuple with the VcConfigurations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcConfigurations

`func (o *NvdcvesNvdCveExtended) SetVcConfigurations(v []AdvisoryNVD20Configuration)`

SetVcConfigurations sets VcConfigurations field to given value.

### HasVcConfigurations

`func (o *NvdcvesNvdCveExtended) HasVcConfigurations() bool`

HasVcConfigurations returns a boolean if a field has been set.

### GetVcConfigurationsWithPlatforms

`func (o *NvdcvesNvdCveExtended) GetVcConfigurationsWithPlatforms() []AdvisoryNVD20Configuration`

GetVcConfigurationsWithPlatforms returns the VcConfigurationsWithPlatforms field if non-nil, zero value otherwise.

### GetVcConfigurationsWithPlatformsOk

`func (o *NvdcvesNvdCveExtended) GetVcConfigurationsWithPlatformsOk() (*[]AdvisoryNVD20Configuration, bool)`

GetVcConfigurationsWithPlatformsOk returns a tuple with the VcConfigurationsWithPlatforms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcConfigurationsWithPlatforms

`func (o *NvdcvesNvdCveExtended) SetVcConfigurationsWithPlatforms(v []AdvisoryNVD20Configuration)`

SetVcConfigurationsWithPlatforms sets VcConfigurationsWithPlatforms field to given value.

### HasVcConfigurationsWithPlatforms

`func (o *NvdcvesNvdCveExtended) HasVcConfigurationsWithPlatforms() bool`

HasVcConfigurationsWithPlatforms returns a boolean if a field has been set.

### GetVcVulnerableCPEs

`func (o *NvdcvesNvdCveExtended) GetVcVulnerableCPEs() []string`

GetVcVulnerableCPEs returns the VcVulnerableCPEs field if non-nil, zero value otherwise.

### GetVcVulnerableCPEsOk

`func (o *NvdcvesNvdCveExtended) GetVcVulnerableCPEsOk() (*[]string, bool)`

GetVcVulnerableCPEsOk returns a tuple with the VcVulnerableCPEs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcVulnerableCPEs

`func (o *NvdcvesNvdCveExtended) SetVcVulnerableCPEs(v []string)`

SetVcVulnerableCPEs sets VcVulnerableCPEs field to given value.

### HasVcVulnerableCPEs

`func (o *NvdcvesNvdCveExtended) HasVcVulnerableCPEs() bool`

HasVcVulnerableCPEs returns a boolean if a field has been set.

### GetVendorComments

`func (o *NvdcvesNvdCveExtended) GetVendorComments() []ApiNVD20VendorComment`

GetVendorComments returns the VendorComments field if non-nil, zero value otherwise.

### GetVendorCommentsOk

`func (o *NvdcvesNvdCveExtended) GetVendorCommentsOk() (*[]ApiNVD20VendorComment, bool)`

GetVendorCommentsOk returns a tuple with the VendorComments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorComments

`func (o *NvdcvesNvdCveExtended) SetVendorComments(v []ApiNVD20VendorComment)`

SetVendorComments sets VendorComments field to given value.

### HasVendorComments

`func (o *NvdcvesNvdCveExtended) HasVendorComments() bool`

HasVendorComments returns a boolean if a field has been set.

### GetVulnStatus

`func (o *NvdcvesNvdCveExtended) GetVulnStatus() string`

GetVulnStatus returns the VulnStatus field if non-nil, zero value otherwise.

### GetVulnStatusOk

`func (o *NvdcvesNvdCveExtended) GetVulnStatusOk() (*string, bool)`

GetVulnStatusOk returns a tuple with the VulnStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVulnStatus

`func (o *NvdcvesNvdCveExtended) SetVulnStatus(v string)`

SetVulnStatus sets VulnStatus field to given value.

### HasVulnStatus

`func (o *NvdcvesNvdCveExtended) HasVulnStatus() bool`

HasVulnStatus returns a boolean if a field has been set.

### GetWeaknesses

`func (o *NvdcvesNvdCveExtended) GetWeaknesses() []ApiNVD20Weakness`

GetWeaknesses returns the Weaknesses field if non-nil, zero value otherwise.

### GetWeaknessesOk

`func (o *NvdcvesNvdCveExtended) GetWeaknessesOk() (*[]ApiNVD20Weakness, bool)`

GetWeaknessesOk returns a tuple with the Weaknesses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeaknesses

`func (o *NvdcvesNvdCveExtended) SetWeaknesses(v []ApiNVD20Weakness)`

SetWeaknesses sets Weaknesses field to given value.

### HasWeaknesses

`func (o *NvdcvesNvdCveExtended) HasWeaknesses() bool`

HasWeaknesses returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


