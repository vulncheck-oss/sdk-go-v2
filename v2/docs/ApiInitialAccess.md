# ApiInitialAccess

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Artifacts** | Pointer to [**[]ApiInitialAccessArtifact**](ApiInitialAccessArtifact.md) | Artifacts holds the set of available artifacts for this vulnerability, such as exploit, shodan queries, PCAP traces, and others. | [optional] 
**Cve** | Pointer to **string** | CVE identifier for the given initial access record. | [optional] 
**InKEV** | Pointer to **bool** | InKEV is true if this artifact is in CISA&#39;s Known Exploited Vulnerabilities (KEV) data set; otherwise, false. | [optional] 
**InVCKEV** | Pointer to **bool** | InVCKEV is true if this artifact is in VulnCheck&#39;s Known Exploited Vulnerabilities (VCKEV) data set; otherwise, false. | [optional] 
**VulnerableCpes** | Pointer to **[]string** | VulnerableCPEs is the list of vulnerable CPE strings associated with this CVE and artifact(s). | [optional] 

## Methods

### NewApiInitialAccess

`func NewApiInitialAccess() *ApiInitialAccess`

NewApiInitialAccess instantiates a new ApiInitialAccess object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiInitialAccessWithDefaults

`func NewApiInitialAccessWithDefaults() *ApiInitialAccess`

NewApiInitialAccessWithDefaults instantiates a new ApiInitialAccess object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArtifacts

`func (o *ApiInitialAccess) GetArtifacts() []ApiInitialAccessArtifact`

GetArtifacts returns the Artifacts field if non-nil, zero value otherwise.

### GetArtifactsOk

`func (o *ApiInitialAccess) GetArtifactsOk() (*[]ApiInitialAccessArtifact, bool)`

GetArtifactsOk returns a tuple with the Artifacts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifacts

`func (o *ApiInitialAccess) SetArtifacts(v []ApiInitialAccessArtifact)`

SetArtifacts sets Artifacts field to given value.

### HasArtifacts

`func (o *ApiInitialAccess) HasArtifacts() bool`

HasArtifacts returns a boolean if a field has been set.

### GetCve

`func (o *ApiInitialAccess) GetCve() string`

GetCve returns the Cve field if non-nil, zero value otherwise.

### GetCveOk

`func (o *ApiInitialAccess) GetCveOk() (*string, bool)`

GetCveOk returns a tuple with the Cve field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCve

`func (o *ApiInitialAccess) SetCve(v string)`

SetCve sets Cve field to given value.

### HasCve

`func (o *ApiInitialAccess) HasCve() bool`

HasCve returns a boolean if a field has been set.

### GetInKEV

`func (o *ApiInitialAccess) GetInKEV() bool`

GetInKEV returns the InKEV field if non-nil, zero value otherwise.

### GetInKEVOk

`func (o *ApiInitialAccess) GetInKEVOk() (*bool, bool)`

GetInKEVOk returns a tuple with the InKEV field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInKEV

`func (o *ApiInitialAccess) SetInKEV(v bool)`

SetInKEV sets InKEV field to given value.

### HasInKEV

`func (o *ApiInitialAccess) HasInKEV() bool`

HasInKEV returns a boolean if a field has been set.

### GetInVCKEV

`func (o *ApiInitialAccess) GetInVCKEV() bool`

GetInVCKEV returns the InVCKEV field if non-nil, zero value otherwise.

### GetInVCKEVOk

`func (o *ApiInitialAccess) GetInVCKEVOk() (*bool, bool)`

GetInVCKEVOk returns a tuple with the InVCKEV field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInVCKEV

`func (o *ApiInitialAccess) SetInVCKEV(v bool)`

SetInVCKEV sets InVCKEV field to given value.

### HasInVCKEV

`func (o *ApiInitialAccess) HasInVCKEV() bool`

HasInVCKEV returns a boolean if a field has been set.

### GetVulnerableCpes

`func (o *ApiInitialAccess) GetVulnerableCpes() []string`

GetVulnerableCpes returns the VulnerableCpes field if non-nil, zero value otherwise.

### GetVulnerableCpesOk

`func (o *ApiInitialAccess) GetVulnerableCpesOk() (*[]string, bool)`

GetVulnerableCpesOk returns a tuple with the VulnerableCpes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVulnerableCpes

`func (o *ApiInitialAccess) SetVulnerableCpes(v []string)`

SetVulnerableCpes sets VulnerableCpes field to given value.

### HasVulnerableCpes

`func (o *ApiInitialAccess) HasVulnerableCpes() bool`

HasVulnerableCpes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


