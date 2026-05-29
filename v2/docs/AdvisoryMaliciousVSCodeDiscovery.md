# AdvisoryMaliciousVSCodeDiscovery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Date** | Pointer to **string** | Date the timestamp when the discovery was published (via the \&quot;reference\&quot;) | [optional] 
**Marketplace** | Pointer to **[]string** | Marketplace list of names of the marketplaces where the extension versions are available | [optional] 
**Reference** | Pointer to **string** | Reference location where the initial public release of the vulnerability was pulled | [optional] 
**Type** | Pointer to **string** | Type VulnCheck vulnerability classification | [optional] 
**Versions** | Pointer to **[]string** | Versions list of vulnerable versions | [optional] 

## Methods

### NewAdvisoryMaliciousVSCodeDiscovery

`func NewAdvisoryMaliciousVSCodeDiscovery() *AdvisoryMaliciousVSCodeDiscovery`

NewAdvisoryMaliciousVSCodeDiscovery instantiates a new AdvisoryMaliciousVSCodeDiscovery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdvisoryMaliciousVSCodeDiscoveryWithDefaults

`func NewAdvisoryMaliciousVSCodeDiscoveryWithDefaults() *AdvisoryMaliciousVSCodeDiscovery`

NewAdvisoryMaliciousVSCodeDiscoveryWithDefaults instantiates a new AdvisoryMaliciousVSCodeDiscovery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDate

`func (o *AdvisoryMaliciousVSCodeDiscovery) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *AdvisoryMaliciousVSCodeDiscovery) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *AdvisoryMaliciousVSCodeDiscovery) SetDate(v string)`

SetDate sets Date field to given value.

### HasDate

`func (o *AdvisoryMaliciousVSCodeDiscovery) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetMarketplace

`func (o *AdvisoryMaliciousVSCodeDiscovery) GetMarketplace() []string`

GetMarketplace returns the Marketplace field if non-nil, zero value otherwise.

### GetMarketplaceOk

`func (o *AdvisoryMaliciousVSCodeDiscovery) GetMarketplaceOk() (*[]string, bool)`

GetMarketplaceOk returns a tuple with the Marketplace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarketplace

`func (o *AdvisoryMaliciousVSCodeDiscovery) SetMarketplace(v []string)`

SetMarketplace sets Marketplace field to given value.

### HasMarketplace

`func (o *AdvisoryMaliciousVSCodeDiscovery) HasMarketplace() bool`

HasMarketplace returns a boolean if a field has been set.

### GetReference

`func (o *AdvisoryMaliciousVSCodeDiscovery) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *AdvisoryMaliciousVSCodeDiscovery) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *AdvisoryMaliciousVSCodeDiscovery) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *AdvisoryMaliciousVSCodeDiscovery) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetType

`func (o *AdvisoryMaliciousVSCodeDiscovery) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AdvisoryMaliciousVSCodeDiscovery) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AdvisoryMaliciousVSCodeDiscovery) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *AdvisoryMaliciousVSCodeDiscovery) HasType() bool`

HasType returns a boolean if a field has been set.

### GetVersions

`func (o *AdvisoryMaliciousVSCodeDiscovery) GetVersions() []string`

GetVersions returns the Versions field if non-nil, zero value otherwise.

### GetVersionsOk

`func (o *AdvisoryMaliciousVSCodeDiscovery) GetVersionsOk() (*[]string, bool)`

GetVersionsOk returns a tuple with the Versions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersions

`func (o *AdvisoryMaliciousVSCodeDiscovery) SetVersions(v []string)`

SetVersions sets Versions field to given value.

### HasVersions

`func (o *AdvisoryMaliciousVSCodeDiscovery) HasVersions() bool`

HasVersions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


