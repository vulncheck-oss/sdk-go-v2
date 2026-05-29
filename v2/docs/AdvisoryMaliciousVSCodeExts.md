# AdvisoryMaliciousVSCodeExts

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Discoveries** | Pointer to [**[]AdvisoryMaliciousVSCodeDiscovery**](AdvisoryMaliciousVSCodeDiscovery.md) | Discoveries list of individual vulnerability reports, sorted chronologically by discovery date | [optional] 
**FirstSeen** | Pointer to **string** | FirstSeen the earliest date the vulnerability was observed or created in our system | [optional] 
**LastUpdated** | Pointer to **string** | LastUpdated the most recent date when any discovery was added for this extension | [optional] 
**Marketplace** | Pointer to **[]string** | Marketplace list of names of the marketplaces where the extension versions are available | [optional] 
**Name** | Pointer to **string** | Name is the name of the extension | [optional] 
**Publisher** | Pointer to **string** | Publisher name of the publisher of the extension | [optional] 
**References** | Pointer to **[]string** | Reference list of locations where the public releases about the vulnerabilities were pulled | [optional] 
**Types** | Pointer to **[]string** | Types VulnCheck vulnerability classifications found for this extension | [optional] 
**Versions** | Pointer to **[]string** | Versions list of vulnerable versions | [optional] 

## Methods

### NewAdvisoryMaliciousVSCodeExts

`func NewAdvisoryMaliciousVSCodeExts() *AdvisoryMaliciousVSCodeExts`

NewAdvisoryMaliciousVSCodeExts instantiates a new AdvisoryMaliciousVSCodeExts object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdvisoryMaliciousVSCodeExtsWithDefaults

`func NewAdvisoryMaliciousVSCodeExtsWithDefaults() *AdvisoryMaliciousVSCodeExts`

NewAdvisoryMaliciousVSCodeExtsWithDefaults instantiates a new AdvisoryMaliciousVSCodeExts object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDiscoveries

`func (o *AdvisoryMaliciousVSCodeExts) GetDiscoveries() []AdvisoryMaliciousVSCodeDiscovery`

GetDiscoveries returns the Discoveries field if non-nil, zero value otherwise.

### GetDiscoveriesOk

`func (o *AdvisoryMaliciousVSCodeExts) GetDiscoveriesOk() (*[]AdvisoryMaliciousVSCodeDiscovery, bool)`

GetDiscoveriesOk returns a tuple with the Discoveries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoveries

`func (o *AdvisoryMaliciousVSCodeExts) SetDiscoveries(v []AdvisoryMaliciousVSCodeDiscovery)`

SetDiscoveries sets Discoveries field to given value.

### HasDiscoveries

`func (o *AdvisoryMaliciousVSCodeExts) HasDiscoveries() bool`

HasDiscoveries returns a boolean if a field has been set.

### GetFirstSeen

`func (o *AdvisoryMaliciousVSCodeExts) GetFirstSeen() string`

GetFirstSeen returns the FirstSeen field if non-nil, zero value otherwise.

### GetFirstSeenOk

`func (o *AdvisoryMaliciousVSCodeExts) GetFirstSeenOk() (*string, bool)`

GetFirstSeenOk returns a tuple with the FirstSeen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstSeen

`func (o *AdvisoryMaliciousVSCodeExts) SetFirstSeen(v string)`

SetFirstSeen sets FirstSeen field to given value.

### HasFirstSeen

`func (o *AdvisoryMaliciousVSCodeExts) HasFirstSeen() bool`

HasFirstSeen returns a boolean if a field has been set.

### GetLastUpdated

`func (o *AdvisoryMaliciousVSCodeExts) GetLastUpdated() string`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *AdvisoryMaliciousVSCodeExts) GetLastUpdatedOk() (*string, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *AdvisoryMaliciousVSCodeExts) SetLastUpdated(v string)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *AdvisoryMaliciousVSCodeExts) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetMarketplace

`func (o *AdvisoryMaliciousVSCodeExts) GetMarketplace() []string`

GetMarketplace returns the Marketplace field if non-nil, zero value otherwise.

### GetMarketplaceOk

`func (o *AdvisoryMaliciousVSCodeExts) GetMarketplaceOk() (*[]string, bool)`

GetMarketplaceOk returns a tuple with the Marketplace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarketplace

`func (o *AdvisoryMaliciousVSCodeExts) SetMarketplace(v []string)`

SetMarketplace sets Marketplace field to given value.

### HasMarketplace

`func (o *AdvisoryMaliciousVSCodeExts) HasMarketplace() bool`

HasMarketplace returns a boolean if a field has been set.

### GetName

`func (o *AdvisoryMaliciousVSCodeExts) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AdvisoryMaliciousVSCodeExts) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AdvisoryMaliciousVSCodeExts) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AdvisoryMaliciousVSCodeExts) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPublisher

`func (o *AdvisoryMaliciousVSCodeExts) GetPublisher() string`

GetPublisher returns the Publisher field if non-nil, zero value otherwise.

### GetPublisherOk

`func (o *AdvisoryMaliciousVSCodeExts) GetPublisherOk() (*string, bool)`

GetPublisherOk returns a tuple with the Publisher field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublisher

`func (o *AdvisoryMaliciousVSCodeExts) SetPublisher(v string)`

SetPublisher sets Publisher field to given value.

### HasPublisher

`func (o *AdvisoryMaliciousVSCodeExts) HasPublisher() bool`

HasPublisher returns a boolean if a field has been set.

### GetReferences

`func (o *AdvisoryMaliciousVSCodeExts) GetReferences() []string`

GetReferences returns the References field if non-nil, zero value otherwise.

### GetReferencesOk

`func (o *AdvisoryMaliciousVSCodeExts) GetReferencesOk() (*[]string, bool)`

GetReferencesOk returns a tuple with the References field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferences

`func (o *AdvisoryMaliciousVSCodeExts) SetReferences(v []string)`

SetReferences sets References field to given value.

### HasReferences

`func (o *AdvisoryMaliciousVSCodeExts) HasReferences() bool`

HasReferences returns a boolean if a field has been set.

### GetTypes

`func (o *AdvisoryMaliciousVSCodeExts) GetTypes() []string`

GetTypes returns the Types field if non-nil, zero value otherwise.

### GetTypesOk

`func (o *AdvisoryMaliciousVSCodeExts) GetTypesOk() (*[]string, bool)`

GetTypesOk returns a tuple with the Types field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypes

`func (o *AdvisoryMaliciousVSCodeExts) SetTypes(v []string)`

SetTypes sets Types field to given value.

### HasTypes

`func (o *AdvisoryMaliciousVSCodeExts) HasTypes() bool`

HasTypes returns a boolean if a field has been set.

### GetVersions

`func (o *AdvisoryMaliciousVSCodeExts) GetVersions() []string`

GetVersions returns the Versions field if non-nil, zero value otherwise.

### GetVersionsOk

`func (o *AdvisoryMaliciousVSCodeExts) GetVersionsOk() (*[]string, bool)`

GetVersionsOk returns a tuple with the Versions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersions

`func (o *AdvisoryMaliciousVSCodeExts) SetVersions(v []string)`

SetVersions sets Versions field to given value.

### HasVersions

`func (o *AdvisoryMaliciousVSCodeExts) HasVersions() bool`

HasVersions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


