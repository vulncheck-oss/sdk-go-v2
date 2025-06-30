# AdvisoryBDUAdvisory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BduId** | Pointer to **string** | BDU:2022-03833 | [optional] 
**Cve** | Pointer to **[]string** | []string{\&quot;CVE-2022-28194\&quot;} | [optional] 
**Cvss** | Pointer to [**AdvisoryBDUCvss**](AdvisoryBDUCvss.md) |  | [optional] 
**Cvss3** | Pointer to [**AdvisoryBDUCvss3**](AdvisoryBDUCvss3.md) |  | [optional] 
**Cwe** | Pointer to **string** | CWE-119 | [optional] 
**DateAdded** | Pointer to **string** |  | [optional] 
**DescriptionRu** | Pointer to **string** | Библиотека libxml2 до версии 2.9.12 не корректно обрабатывает XML-документы, содержащие определенные сущности. В результате могут быть выполнены произвольные команды. | [optional] 
**Environment** | Pointer to [**AdvisoryBDUEnvironment**](AdvisoryBDUEnvironment.md) |  | [optional] 
**ExploitStatusEn** | Pointer to **string** | Exploited | [optional] 
**ExploitStatusRu** | Pointer to **string** | Exploited | [optional] 
**FixStatusEn** | Pointer to **string** | Fixed | [optional] 
**FixStatusRu** | Pointer to **string** | Fixed | [optional] 
**IdentifyDate** | Pointer to **string** | 2022-09-01 | [optional] 
**NameRu** | Pointer to **string** | BDU:2022-03833: Уязвимость модуля Cboot (tegrabl_cbo.c) пакета драйверов микропрограммного обеспечения вычислительных плат NVIDIA Jetson, позволяющая нарушителю выполнить произвольный код или вызвать частичный отказ в обслуживании | [optional] 
**SeverityRu** | Pointer to **string** | High | [optional] 
**SolutionRu** | Pointer to **string** | Обновите драйверы микропрограммного обеспечения вычислительных плат NVIDIA Jetson до версии 32.6.1 или более поздней | [optional] 
**Sources** | Pointer to **[]string** | https://nvd.nist.gov/vuln/detail/CVE-2022-28194 | [optional] 
**TextRu** | Pointer to **string** | Библиотека libxml2 до версии 2.9.12 не корректно обрабатывает XML-документы, содержащие определенные сущности. В результате могут быть выполнены произвольные команды. | [optional] 
**Url** | Pointer to **string** | https://bdu.fstec.ru/vul/2022-03833 | [optional] 
**VulStatusEn** | Pointer to **string** | Exploitable | [optional] 
**VulStatusRu** | Pointer to **string** | Exploitable | [optional] 
**VulnerableSoftware** | Pointer to [**AdvisoryBDUVulnerableSoftware**](AdvisoryBDUVulnerableSoftware.md) |  | [optional] 

## Methods

### NewAdvisoryBDUAdvisory

`func NewAdvisoryBDUAdvisory() *AdvisoryBDUAdvisory`

NewAdvisoryBDUAdvisory instantiates a new AdvisoryBDUAdvisory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdvisoryBDUAdvisoryWithDefaults

`func NewAdvisoryBDUAdvisoryWithDefaults() *AdvisoryBDUAdvisory`

NewAdvisoryBDUAdvisoryWithDefaults instantiates a new AdvisoryBDUAdvisory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBduId

`func (o *AdvisoryBDUAdvisory) GetBduId() string`

GetBduId returns the BduId field if non-nil, zero value otherwise.

### GetBduIdOk

`func (o *AdvisoryBDUAdvisory) GetBduIdOk() (*string, bool)`

GetBduIdOk returns a tuple with the BduId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBduId

`func (o *AdvisoryBDUAdvisory) SetBduId(v string)`

SetBduId sets BduId field to given value.

### HasBduId

`func (o *AdvisoryBDUAdvisory) HasBduId() bool`

HasBduId returns a boolean if a field has been set.

### GetCve

`func (o *AdvisoryBDUAdvisory) GetCve() []string`

GetCve returns the Cve field if non-nil, zero value otherwise.

### GetCveOk

`func (o *AdvisoryBDUAdvisory) GetCveOk() (*[]string, bool)`

GetCveOk returns a tuple with the Cve field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCve

`func (o *AdvisoryBDUAdvisory) SetCve(v []string)`

SetCve sets Cve field to given value.

### HasCve

`func (o *AdvisoryBDUAdvisory) HasCve() bool`

HasCve returns a boolean if a field has been set.

### GetCvss

`func (o *AdvisoryBDUAdvisory) GetCvss() AdvisoryBDUCvss`

GetCvss returns the Cvss field if non-nil, zero value otherwise.

### GetCvssOk

`func (o *AdvisoryBDUAdvisory) GetCvssOk() (*AdvisoryBDUCvss, bool)`

GetCvssOk returns a tuple with the Cvss field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCvss

`func (o *AdvisoryBDUAdvisory) SetCvss(v AdvisoryBDUCvss)`

SetCvss sets Cvss field to given value.

### HasCvss

`func (o *AdvisoryBDUAdvisory) HasCvss() bool`

HasCvss returns a boolean if a field has been set.

### GetCvss3

`func (o *AdvisoryBDUAdvisory) GetCvss3() AdvisoryBDUCvss3`

GetCvss3 returns the Cvss3 field if non-nil, zero value otherwise.

### GetCvss3Ok

`func (o *AdvisoryBDUAdvisory) GetCvss3Ok() (*AdvisoryBDUCvss3, bool)`

GetCvss3Ok returns a tuple with the Cvss3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCvss3

`func (o *AdvisoryBDUAdvisory) SetCvss3(v AdvisoryBDUCvss3)`

SetCvss3 sets Cvss3 field to given value.

### HasCvss3

`func (o *AdvisoryBDUAdvisory) HasCvss3() bool`

HasCvss3 returns a boolean if a field has been set.

### GetCwe

`func (o *AdvisoryBDUAdvisory) GetCwe() string`

GetCwe returns the Cwe field if non-nil, zero value otherwise.

### GetCweOk

`func (o *AdvisoryBDUAdvisory) GetCweOk() (*string, bool)`

GetCweOk returns a tuple with the Cwe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCwe

`func (o *AdvisoryBDUAdvisory) SetCwe(v string)`

SetCwe sets Cwe field to given value.

### HasCwe

`func (o *AdvisoryBDUAdvisory) HasCwe() bool`

HasCwe returns a boolean if a field has been set.

### GetDateAdded

`func (o *AdvisoryBDUAdvisory) GetDateAdded() string`

GetDateAdded returns the DateAdded field if non-nil, zero value otherwise.

### GetDateAddedOk

`func (o *AdvisoryBDUAdvisory) GetDateAddedOk() (*string, bool)`

GetDateAddedOk returns a tuple with the DateAdded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateAdded

`func (o *AdvisoryBDUAdvisory) SetDateAdded(v string)`

SetDateAdded sets DateAdded field to given value.

### HasDateAdded

`func (o *AdvisoryBDUAdvisory) HasDateAdded() bool`

HasDateAdded returns a boolean if a field has been set.

### GetDescriptionRu

`func (o *AdvisoryBDUAdvisory) GetDescriptionRu() string`

GetDescriptionRu returns the DescriptionRu field if non-nil, zero value otherwise.

### GetDescriptionRuOk

`func (o *AdvisoryBDUAdvisory) GetDescriptionRuOk() (*string, bool)`

GetDescriptionRuOk returns a tuple with the DescriptionRu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescriptionRu

`func (o *AdvisoryBDUAdvisory) SetDescriptionRu(v string)`

SetDescriptionRu sets DescriptionRu field to given value.

### HasDescriptionRu

`func (o *AdvisoryBDUAdvisory) HasDescriptionRu() bool`

HasDescriptionRu returns a boolean if a field has been set.

### GetEnvironment

`func (o *AdvisoryBDUAdvisory) GetEnvironment() AdvisoryBDUEnvironment`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *AdvisoryBDUAdvisory) GetEnvironmentOk() (*AdvisoryBDUEnvironment, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *AdvisoryBDUAdvisory) SetEnvironment(v AdvisoryBDUEnvironment)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *AdvisoryBDUAdvisory) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetExploitStatusEn

`func (o *AdvisoryBDUAdvisory) GetExploitStatusEn() string`

GetExploitStatusEn returns the ExploitStatusEn field if non-nil, zero value otherwise.

### GetExploitStatusEnOk

`func (o *AdvisoryBDUAdvisory) GetExploitStatusEnOk() (*string, bool)`

GetExploitStatusEnOk returns a tuple with the ExploitStatusEn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExploitStatusEn

`func (o *AdvisoryBDUAdvisory) SetExploitStatusEn(v string)`

SetExploitStatusEn sets ExploitStatusEn field to given value.

### HasExploitStatusEn

`func (o *AdvisoryBDUAdvisory) HasExploitStatusEn() bool`

HasExploitStatusEn returns a boolean if a field has been set.

### GetExploitStatusRu

`func (o *AdvisoryBDUAdvisory) GetExploitStatusRu() string`

GetExploitStatusRu returns the ExploitStatusRu field if non-nil, zero value otherwise.

### GetExploitStatusRuOk

`func (o *AdvisoryBDUAdvisory) GetExploitStatusRuOk() (*string, bool)`

GetExploitStatusRuOk returns a tuple with the ExploitStatusRu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExploitStatusRu

`func (o *AdvisoryBDUAdvisory) SetExploitStatusRu(v string)`

SetExploitStatusRu sets ExploitStatusRu field to given value.

### HasExploitStatusRu

`func (o *AdvisoryBDUAdvisory) HasExploitStatusRu() bool`

HasExploitStatusRu returns a boolean if a field has been set.

### GetFixStatusEn

`func (o *AdvisoryBDUAdvisory) GetFixStatusEn() string`

GetFixStatusEn returns the FixStatusEn field if non-nil, zero value otherwise.

### GetFixStatusEnOk

`func (o *AdvisoryBDUAdvisory) GetFixStatusEnOk() (*string, bool)`

GetFixStatusEnOk returns a tuple with the FixStatusEn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFixStatusEn

`func (o *AdvisoryBDUAdvisory) SetFixStatusEn(v string)`

SetFixStatusEn sets FixStatusEn field to given value.

### HasFixStatusEn

`func (o *AdvisoryBDUAdvisory) HasFixStatusEn() bool`

HasFixStatusEn returns a boolean if a field has been set.

### GetFixStatusRu

`func (o *AdvisoryBDUAdvisory) GetFixStatusRu() string`

GetFixStatusRu returns the FixStatusRu field if non-nil, zero value otherwise.

### GetFixStatusRuOk

`func (o *AdvisoryBDUAdvisory) GetFixStatusRuOk() (*string, bool)`

GetFixStatusRuOk returns a tuple with the FixStatusRu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFixStatusRu

`func (o *AdvisoryBDUAdvisory) SetFixStatusRu(v string)`

SetFixStatusRu sets FixStatusRu field to given value.

### HasFixStatusRu

`func (o *AdvisoryBDUAdvisory) HasFixStatusRu() bool`

HasFixStatusRu returns a boolean if a field has been set.

### GetIdentifyDate

`func (o *AdvisoryBDUAdvisory) GetIdentifyDate() string`

GetIdentifyDate returns the IdentifyDate field if non-nil, zero value otherwise.

### GetIdentifyDateOk

`func (o *AdvisoryBDUAdvisory) GetIdentifyDateOk() (*string, bool)`

GetIdentifyDateOk returns a tuple with the IdentifyDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifyDate

`func (o *AdvisoryBDUAdvisory) SetIdentifyDate(v string)`

SetIdentifyDate sets IdentifyDate field to given value.

### HasIdentifyDate

`func (o *AdvisoryBDUAdvisory) HasIdentifyDate() bool`

HasIdentifyDate returns a boolean if a field has been set.

### GetNameRu

`func (o *AdvisoryBDUAdvisory) GetNameRu() string`

GetNameRu returns the NameRu field if non-nil, zero value otherwise.

### GetNameRuOk

`func (o *AdvisoryBDUAdvisory) GetNameRuOk() (*string, bool)`

GetNameRuOk returns a tuple with the NameRu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameRu

`func (o *AdvisoryBDUAdvisory) SetNameRu(v string)`

SetNameRu sets NameRu field to given value.

### HasNameRu

`func (o *AdvisoryBDUAdvisory) HasNameRu() bool`

HasNameRu returns a boolean if a field has been set.

### GetSeverityRu

`func (o *AdvisoryBDUAdvisory) GetSeverityRu() string`

GetSeverityRu returns the SeverityRu field if non-nil, zero value otherwise.

### GetSeverityRuOk

`func (o *AdvisoryBDUAdvisory) GetSeverityRuOk() (*string, bool)`

GetSeverityRuOk returns a tuple with the SeverityRu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverityRu

`func (o *AdvisoryBDUAdvisory) SetSeverityRu(v string)`

SetSeverityRu sets SeverityRu field to given value.

### HasSeverityRu

`func (o *AdvisoryBDUAdvisory) HasSeverityRu() bool`

HasSeverityRu returns a boolean if a field has been set.

### GetSolutionRu

`func (o *AdvisoryBDUAdvisory) GetSolutionRu() string`

GetSolutionRu returns the SolutionRu field if non-nil, zero value otherwise.

### GetSolutionRuOk

`func (o *AdvisoryBDUAdvisory) GetSolutionRuOk() (*string, bool)`

GetSolutionRuOk returns a tuple with the SolutionRu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSolutionRu

`func (o *AdvisoryBDUAdvisory) SetSolutionRu(v string)`

SetSolutionRu sets SolutionRu field to given value.

### HasSolutionRu

`func (o *AdvisoryBDUAdvisory) HasSolutionRu() bool`

HasSolutionRu returns a boolean if a field has been set.

### GetSources

`func (o *AdvisoryBDUAdvisory) GetSources() []string`

GetSources returns the Sources field if non-nil, zero value otherwise.

### GetSourcesOk

`func (o *AdvisoryBDUAdvisory) GetSourcesOk() (*[]string, bool)`

GetSourcesOk returns a tuple with the Sources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSources

`func (o *AdvisoryBDUAdvisory) SetSources(v []string)`

SetSources sets Sources field to given value.

### HasSources

`func (o *AdvisoryBDUAdvisory) HasSources() bool`

HasSources returns a boolean if a field has been set.

### GetTextRu

`func (o *AdvisoryBDUAdvisory) GetTextRu() string`

GetTextRu returns the TextRu field if non-nil, zero value otherwise.

### GetTextRuOk

`func (o *AdvisoryBDUAdvisory) GetTextRuOk() (*string, bool)`

GetTextRuOk returns a tuple with the TextRu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTextRu

`func (o *AdvisoryBDUAdvisory) SetTextRu(v string)`

SetTextRu sets TextRu field to given value.

### HasTextRu

`func (o *AdvisoryBDUAdvisory) HasTextRu() bool`

HasTextRu returns a boolean if a field has been set.

### GetUrl

`func (o *AdvisoryBDUAdvisory) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *AdvisoryBDUAdvisory) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *AdvisoryBDUAdvisory) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *AdvisoryBDUAdvisory) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetVulStatusEn

`func (o *AdvisoryBDUAdvisory) GetVulStatusEn() string`

GetVulStatusEn returns the VulStatusEn field if non-nil, zero value otherwise.

### GetVulStatusEnOk

`func (o *AdvisoryBDUAdvisory) GetVulStatusEnOk() (*string, bool)`

GetVulStatusEnOk returns a tuple with the VulStatusEn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVulStatusEn

`func (o *AdvisoryBDUAdvisory) SetVulStatusEn(v string)`

SetVulStatusEn sets VulStatusEn field to given value.

### HasVulStatusEn

`func (o *AdvisoryBDUAdvisory) HasVulStatusEn() bool`

HasVulStatusEn returns a boolean if a field has been set.

### GetVulStatusRu

`func (o *AdvisoryBDUAdvisory) GetVulStatusRu() string`

GetVulStatusRu returns the VulStatusRu field if non-nil, zero value otherwise.

### GetVulStatusRuOk

`func (o *AdvisoryBDUAdvisory) GetVulStatusRuOk() (*string, bool)`

GetVulStatusRuOk returns a tuple with the VulStatusRu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVulStatusRu

`func (o *AdvisoryBDUAdvisory) SetVulStatusRu(v string)`

SetVulStatusRu sets VulStatusRu field to given value.

### HasVulStatusRu

`func (o *AdvisoryBDUAdvisory) HasVulStatusRu() bool`

HasVulStatusRu returns a boolean if a field has been set.

### GetVulnerableSoftware

`func (o *AdvisoryBDUAdvisory) GetVulnerableSoftware() AdvisoryBDUVulnerableSoftware`

GetVulnerableSoftware returns the VulnerableSoftware field if non-nil, zero value otherwise.

### GetVulnerableSoftwareOk

`func (o *AdvisoryBDUAdvisory) GetVulnerableSoftwareOk() (*AdvisoryBDUVulnerableSoftware, bool)`

GetVulnerableSoftwareOk returns a tuple with the VulnerableSoftware field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVulnerableSoftware

`func (o *AdvisoryBDUAdvisory) SetVulnerableSoftware(v AdvisoryBDUVulnerableSoftware)`

SetVulnerableSoftware sets VulnerableSoftware field to given value.

### HasVulnerableSoftware

`func (o *AdvisoryBDUAdvisory) HasVulnerableSoftware() bool`

HasVulnerableSoftware returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


