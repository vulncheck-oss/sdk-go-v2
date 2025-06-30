# AdvisoryVulnCheckKEV

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Timestamp** | Pointer to **string** |  | [optional] 
**CisaDateAdded** | Pointer to **string** |  | [optional] 
**Cve** | Pointer to **[]string** |  | [optional] 
**Cwes** | Pointer to **[]string** |  | [optional] 
**DateAdded** | Pointer to **string** |  | [optional] 
**DueDate** | Pointer to **string** |  | [optional] 
**KnownRansomwareCampaignUse** | Pointer to **string** |  | [optional] 
**Product** | Pointer to **string** |  | [optional] 
**RequiredAction** | Pointer to **string** |  | [optional] 
**ShortDescription** | Pointer to **string** |  | [optional] 
**VendorProject** | Pointer to **string** |  | [optional] 
**VulncheckReportedExploitation** | Pointer to [**[]AdvisoryReportedExploit**](AdvisoryReportedExploit.md) |  | [optional] 
**VulncheckXdb** | Pointer to [**[]AdvisoryXDB**](AdvisoryXDB.md) |  | [optional] 
**VulnerabilityName** | Pointer to **string** |  | [optional] 

## Methods

### NewAdvisoryVulnCheckKEV

`func NewAdvisoryVulnCheckKEV() *AdvisoryVulnCheckKEV`

NewAdvisoryVulnCheckKEV instantiates a new AdvisoryVulnCheckKEV object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdvisoryVulnCheckKEVWithDefaults

`func NewAdvisoryVulnCheckKEVWithDefaults() *AdvisoryVulnCheckKEV`

NewAdvisoryVulnCheckKEVWithDefaults instantiates a new AdvisoryVulnCheckKEV object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimestamp

`func (o *AdvisoryVulnCheckKEV) GetTimestamp() string`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *AdvisoryVulnCheckKEV) GetTimestampOk() (*string, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *AdvisoryVulnCheckKEV) SetTimestamp(v string)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *AdvisoryVulnCheckKEV) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetCisaDateAdded

`func (o *AdvisoryVulnCheckKEV) GetCisaDateAdded() string`

GetCisaDateAdded returns the CisaDateAdded field if non-nil, zero value otherwise.

### GetCisaDateAddedOk

`func (o *AdvisoryVulnCheckKEV) GetCisaDateAddedOk() (*string, bool)`

GetCisaDateAddedOk returns a tuple with the CisaDateAdded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCisaDateAdded

`func (o *AdvisoryVulnCheckKEV) SetCisaDateAdded(v string)`

SetCisaDateAdded sets CisaDateAdded field to given value.

### HasCisaDateAdded

`func (o *AdvisoryVulnCheckKEV) HasCisaDateAdded() bool`

HasCisaDateAdded returns a boolean if a field has been set.

### GetCve

`func (o *AdvisoryVulnCheckKEV) GetCve() []string`

GetCve returns the Cve field if non-nil, zero value otherwise.

### GetCveOk

`func (o *AdvisoryVulnCheckKEV) GetCveOk() (*[]string, bool)`

GetCveOk returns a tuple with the Cve field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCve

`func (o *AdvisoryVulnCheckKEV) SetCve(v []string)`

SetCve sets Cve field to given value.

### HasCve

`func (o *AdvisoryVulnCheckKEV) HasCve() bool`

HasCve returns a boolean if a field has been set.

### GetCwes

`func (o *AdvisoryVulnCheckKEV) GetCwes() []string`

GetCwes returns the Cwes field if non-nil, zero value otherwise.

### GetCwesOk

`func (o *AdvisoryVulnCheckKEV) GetCwesOk() (*[]string, bool)`

GetCwesOk returns a tuple with the Cwes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCwes

`func (o *AdvisoryVulnCheckKEV) SetCwes(v []string)`

SetCwes sets Cwes field to given value.

### HasCwes

`func (o *AdvisoryVulnCheckKEV) HasCwes() bool`

HasCwes returns a boolean if a field has been set.

### GetDateAdded

`func (o *AdvisoryVulnCheckKEV) GetDateAdded() string`

GetDateAdded returns the DateAdded field if non-nil, zero value otherwise.

### GetDateAddedOk

`func (o *AdvisoryVulnCheckKEV) GetDateAddedOk() (*string, bool)`

GetDateAddedOk returns a tuple with the DateAdded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateAdded

`func (o *AdvisoryVulnCheckKEV) SetDateAdded(v string)`

SetDateAdded sets DateAdded field to given value.

### HasDateAdded

`func (o *AdvisoryVulnCheckKEV) HasDateAdded() bool`

HasDateAdded returns a boolean if a field has been set.

### GetDueDate

`func (o *AdvisoryVulnCheckKEV) GetDueDate() string`

GetDueDate returns the DueDate field if non-nil, zero value otherwise.

### GetDueDateOk

`func (o *AdvisoryVulnCheckKEV) GetDueDateOk() (*string, bool)`

GetDueDateOk returns a tuple with the DueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueDate

`func (o *AdvisoryVulnCheckKEV) SetDueDate(v string)`

SetDueDate sets DueDate field to given value.

### HasDueDate

`func (o *AdvisoryVulnCheckKEV) HasDueDate() bool`

HasDueDate returns a boolean if a field has been set.

### GetKnownRansomwareCampaignUse

`func (o *AdvisoryVulnCheckKEV) GetKnownRansomwareCampaignUse() string`

GetKnownRansomwareCampaignUse returns the KnownRansomwareCampaignUse field if non-nil, zero value otherwise.

### GetKnownRansomwareCampaignUseOk

`func (o *AdvisoryVulnCheckKEV) GetKnownRansomwareCampaignUseOk() (*string, bool)`

GetKnownRansomwareCampaignUseOk returns a tuple with the KnownRansomwareCampaignUse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKnownRansomwareCampaignUse

`func (o *AdvisoryVulnCheckKEV) SetKnownRansomwareCampaignUse(v string)`

SetKnownRansomwareCampaignUse sets KnownRansomwareCampaignUse field to given value.

### HasKnownRansomwareCampaignUse

`func (o *AdvisoryVulnCheckKEV) HasKnownRansomwareCampaignUse() bool`

HasKnownRansomwareCampaignUse returns a boolean if a field has been set.

### GetProduct

`func (o *AdvisoryVulnCheckKEV) GetProduct() string`

GetProduct returns the Product field if non-nil, zero value otherwise.

### GetProductOk

`func (o *AdvisoryVulnCheckKEV) GetProductOk() (*string, bool)`

GetProductOk returns a tuple with the Product field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduct

`func (o *AdvisoryVulnCheckKEV) SetProduct(v string)`

SetProduct sets Product field to given value.

### HasProduct

`func (o *AdvisoryVulnCheckKEV) HasProduct() bool`

HasProduct returns a boolean if a field has been set.

### GetRequiredAction

`func (o *AdvisoryVulnCheckKEV) GetRequiredAction() string`

GetRequiredAction returns the RequiredAction field if non-nil, zero value otherwise.

### GetRequiredActionOk

`func (o *AdvisoryVulnCheckKEV) GetRequiredActionOk() (*string, bool)`

GetRequiredActionOk returns a tuple with the RequiredAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredAction

`func (o *AdvisoryVulnCheckKEV) SetRequiredAction(v string)`

SetRequiredAction sets RequiredAction field to given value.

### HasRequiredAction

`func (o *AdvisoryVulnCheckKEV) HasRequiredAction() bool`

HasRequiredAction returns a boolean if a field has been set.

### GetShortDescription

`func (o *AdvisoryVulnCheckKEV) GetShortDescription() string`

GetShortDescription returns the ShortDescription field if non-nil, zero value otherwise.

### GetShortDescriptionOk

`func (o *AdvisoryVulnCheckKEV) GetShortDescriptionOk() (*string, bool)`

GetShortDescriptionOk returns a tuple with the ShortDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortDescription

`func (o *AdvisoryVulnCheckKEV) SetShortDescription(v string)`

SetShortDescription sets ShortDescription field to given value.

### HasShortDescription

`func (o *AdvisoryVulnCheckKEV) HasShortDescription() bool`

HasShortDescription returns a boolean if a field has been set.

### GetVendorProject

`func (o *AdvisoryVulnCheckKEV) GetVendorProject() string`

GetVendorProject returns the VendorProject field if non-nil, zero value otherwise.

### GetVendorProjectOk

`func (o *AdvisoryVulnCheckKEV) GetVendorProjectOk() (*string, bool)`

GetVendorProjectOk returns a tuple with the VendorProject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorProject

`func (o *AdvisoryVulnCheckKEV) SetVendorProject(v string)`

SetVendorProject sets VendorProject field to given value.

### HasVendorProject

`func (o *AdvisoryVulnCheckKEV) HasVendorProject() bool`

HasVendorProject returns a boolean if a field has been set.

### GetVulncheckReportedExploitation

`func (o *AdvisoryVulnCheckKEV) GetVulncheckReportedExploitation() []AdvisoryReportedExploit`

GetVulncheckReportedExploitation returns the VulncheckReportedExploitation field if non-nil, zero value otherwise.

### GetVulncheckReportedExploitationOk

`func (o *AdvisoryVulnCheckKEV) GetVulncheckReportedExploitationOk() (*[]AdvisoryReportedExploit, bool)`

GetVulncheckReportedExploitationOk returns a tuple with the VulncheckReportedExploitation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVulncheckReportedExploitation

`func (o *AdvisoryVulnCheckKEV) SetVulncheckReportedExploitation(v []AdvisoryReportedExploit)`

SetVulncheckReportedExploitation sets VulncheckReportedExploitation field to given value.

### HasVulncheckReportedExploitation

`func (o *AdvisoryVulnCheckKEV) HasVulncheckReportedExploitation() bool`

HasVulncheckReportedExploitation returns a boolean if a field has been set.

### GetVulncheckXdb

`func (o *AdvisoryVulnCheckKEV) GetVulncheckXdb() []AdvisoryXDB`

GetVulncheckXdb returns the VulncheckXdb field if non-nil, zero value otherwise.

### GetVulncheckXdbOk

`func (o *AdvisoryVulnCheckKEV) GetVulncheckXdbOk() (*[]AdvisoryXDB, bool)`

GetVulncheckXdbOk returns a tuple with the VulncheckXdb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVulncheckXdb

`func (o *AdvisoryVulnCheckKEV) SetVulncheckXdb(v []AdvisoryXDB)`

SetVulncheckXdb sets VulncheckXdb field to given value.

### HasVulncheckXdb

`func (o *AdvisoryVulnCheckKEV) HasVulncheckXdb() bool`

HasVulncheckXdb returns a boolean if a field has been set.

### GetVulnerabilityName

`func (o *AdvisoryVulnCheckKEV) GetVulnerabilityName() string`

GetVulnerabilityName returns the VulnerabilityName field if non-nil, zero value otherwise.

### GetVulnerabilityNameOk

`func (o *AdvisoryVulnCheckKEV) GetVulnerabilityNameOk() (*string, bool)`

GetVulnerabilityNameOk returns a tuple with the VulnerabilityName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVulnerabilityName

`func (o *AdvisoryVulnCheckKEV) SetVulnerabilityName(v string)`

SetVulnerabilityName sets VulnerabilityName field to given value.

### HasVulnerabilityName

`func (o *AdvisoryVulnCheckKEV) HasVulnerabilityName() bool`

HasVulnerabilityName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


