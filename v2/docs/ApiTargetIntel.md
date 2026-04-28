# ApiTargetIntel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Timestamp** | Pointer to **string** |  | [optional] 
**AsDomain** | Pointer to **string** |  | [optional] 
**AsName** | Pointer to **string** |  | [optional] 
**Asn** | Pointer to **string** |  | [optional] 
**Confidence** | Pointer to **float32** |  | [optional] 
**Country** | Pointer to **string** |  | [optional] 
**CountryCode** | Pointer to **string** |  | [optional] 
**Cpe** | Pointer to **string** |  | [optional] 
**Cves** | Pointer to **[]string** |  | [optional] 
**DateAdded** | Pointer to **string** |  | [optional] 
**FingerprintMetadata** | Pointer to **[]int32** |  | [optional] 
**Ip** | Pointer to **string** |  | [optional] 
**MatchTarget** | Pointer to **string** |  | [optional] 
**Port** | Pointer to **int32** |  | [optional] 
**Product** | Pointer to **string** |  | [optional] 
**Protocol** | Pointer to **string** |  | [optional] 
**RuleId** | Pointer to **string** |  | [optional] 
**RuleSource** | Pointer to **string** |  | [optional] 
**ServiceMetadata** | Pointer to **[]int32** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Vendor** | Pointer to **string** |  | [optional] 
**Version** | Pointer to **string** |  | [optional] 

## Methods

### NewApiTargetIntel

`func NewApiTargetIntel() *ApiTargetIntel`

NewApiTargetIntel instantiates a new ApiTargetIntel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiTargetIntelWithDefaults

`func NewApiTargetIntelWithDefaults() *ApiTargetIntel`

NewApiTargetIntelWithDefaults instantiates a new ApiTargetIntel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimestamp

`func (o *ApiTargetIntel) GetTimestamp() string`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ApiTargetIntel) GetTimestampOk() (*string, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ApiTargetIntel) SetTimestamp(v string)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *ApiTargetIntel) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetAsDomain

`func (o *ApiTargetIntel) GetAsDomain() string`

GetAsDomain returns the AsDomain field if non-nil, zero value otherwise.

### GetAsDomainOk

`func (o *ApiTargetIntel) GetAsDomainOk() (*string, bool)`

GetAsDomainOk returns a tuple with the AsDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsDomain

`func (o *ApiTargetIntel) SetAsDomain(v string)`

SetAsDomain sets AsDomain field to given value.

### HasAsDomain

`func (o *ApiTargetIntel) HasAsDomain() bool`

HasAsDomain returns a boolean if a field has been set.

### GetAsName

`func (o *ApiTargetIntel) GetAsName() string`

GetAsName returns the AsName field if non-nil, zero value otherwise.

### GetAsNameOk

`func (o *ApiTargetIntel) GetAsNameOk() (*string, bool)`

GetAsNameOk returns a tuple with the AsName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsName

`func (o *ApiTargetIntel) SetAsName(v string)`

SetAsName sets AsName field to given value.

### HasAsName

`func (o *ApiTargetIntel) HasAsName() bool`

HasAsName returns a boolean if a field has been set.

### GetAsn

`func (o *ApiTargetIntel) GetAsn() string`

GetAsn returns the Asn field if non-nil, zero value otherwise.

### GetAsnOk

`func (o *ApiTargetIntel) GetAsnOk() (*string, bool)`

GetAsnOk returns a tuple with the Asn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsn

`func (o *ApiTargetIntel) SetAsn(v string)`

SetAsn sets Asn field to given value.

### HasAsn

`func (o *ApiTargetIntel) HasAsn() bool`

HasAsn returns a boolean if a field has been set.

### GetConfidence

`func (o *ApiTargetIntel) GetConfidence() float32`

GetConfidence returns the Confidence field if non-nil, zero value otherwise.

### GetConfidenceOk

`func (o *ApiTargetIntel) GetConfidenceOk() (*float32, bool)`

GetConfidenceOk returns a tuple with the Confidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfidence

`func (o *ApiTargetIntel) SetConfidence(v float32)`

SetConfidence sets Confidence field to given value.

### HasConfidence

`func (o *ApiTargetIntel) HasConfidence() bool`

HasConfidence returns a boolean if a field has been set.

### GetCountry

`func (o *ApiTargetIntel) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *ApiTargetIntel) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *ApiTargetIntel) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *ApiTargetIntel) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetCountryCode

`func (o *ApiTargetIntel) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *ApiTargetIntel) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *ApiTargetIntel) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.

### HasCountryCode

`func (o *ApiTargetIntel) HasCountryCode() bool`

HasCountryCode returns a boolean if a field has been set.

### GetCpe

`func (o *ApiTargetIntel) GetCpe() string`

GetCpe returns the Cpe field if non-nil, zero value otherwise.

### GetCpeOk

`func (o *ApiTargetIntel) GetCpeOk() (*string, bool)`

GetCpeOk returns a tuple with the Cpe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpe

`func (o *ApiTargetIntel) SetCpe(v string)`

SetCpe sets Cpe field to given value.

### HasCpe

`func (o *ApiTargetIntel) HasCpe() bool`

HasCpe returns a boolean if a field has been set.

### GetCves

`func (o *ApiTargetIntel) GetCves() []string`

GetCves returns the Cves field if non-nil, zero value otherwise.

### GetCvesOk

`func (o *ApiTargetIntel) GetCvesOk() (*[]string, bool)`

GetCvesOk returns a tuple with the Cves field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCves

`func (o *ApiTargetIntel) SetCves(v []string)`

SetCves sets Cves field to given value.

### HasCves

`func (o *ApiTargetIntel) HasCves() bool`

HasCves returns a boolean if a field has been set.

### GetDateAdded

`func (o *ApiTargetIntel) GetDateAdded() string`

GetDateAdded returns the DateAdded field if non-nil, zero value otherwise.

### GetDateAddedOk

`func (o *ApiTargetIntel) GetDateAddedOk() (*string, bool)`

GetDateAddedOk returns a tuple with the DateAdded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateAdded

`func (o *ApiTargetIntel) SetDateAdded(v string)`

SetDateAdded sets DateAdded field to given value.

### HasDateAdded

`func (o *ApiTargetIntel) HasDateAdded() bool`

HasDateAdded returns a boolean if a field has been set.

### GetFingerprintMetadata

`func (o *ApiTargetIntel) GetFingerprintMetadata() []int32`

GetFingerprintMetadata returns the FingerprintMetadata field if non-nil, zero value otherwise.

### GetFingerprintMetadataOk

`func (o *ApiTargetIntel) GetFingerprintMetadataOk() (*[]int32, bool)`

GetFingerprintMetadataOk returns a tuple with the FingerprintMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFingerprintMetadata

`func (o *ApiTargetIntel) SetFingerprintMetadata(v []int32)`

SetFingerprintMetadata sets FingerprintMetadata field to given value.

### HasFingerprintMetadata

`func (o *ApiTargetIntel) HasFingerprintMetadata() bool`

HasFingerprintMetadata returns a boolean if a field has been set.

### GetIp

`func (o *ApiTargetIntel) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *ApiTargetIntel) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *ApiTargetIntel) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *ApiTargetIntel) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetMatchTarget

`func (o *ApiTargetIntel) GetMatchTarget() string`

GetMatchTarget returns the MatchTarget field if non-nil, zero value otherwise.

### GetMatchTargetOk

`func (o *ApiTargetIntel) GetMatchTargetOk() (*string, bool)`

GetMatchTargetOk returns a tuple with the MatchTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchTarget

`func (o *ApiTargetIntel) SetMatchTarget(v string)`

SetMatchTarget sets MatchTarget field to given value.

### HasMatchTarget

`func (o *ApiTargetIntel) HasMatchTarget() bool`

HasMatchTarget returns a boolean if a field has been set.

### GetPort

`func (o *ApiTargetIntel) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *ApiTargetIntel) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *ApiTargetIntel) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *ApiTargetIntel) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetProduct

`func (o *ApiTargetIntel) GetProduct() string`

GetProduct returns the Product field if non-nil, zero value otherwise.

### GetProductOk

`func (o *ApiTargetIntel) GetProductOk() (*string, bool)`

GetProductOk returns a tuple with the Product field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduct

`func (o *ApiTargetIntel) SetProduct(v string)`

SetProduct sets Product field to given value.

### HasProduct

`func (o *ApiTargetIntel) HasProduct() bool`

HasProduct returns a boolean if a field has been set.

### GetProtocol

`func (o *ApiTargetIntel) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *ApiTargetIntel) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *ApiTargetIntel) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *ApiTargetIntel) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetRuleId

`func (o *ApiTargetIntel) GetRuleId() string`

GetRuleId returns the RuleId field if non-nil, zero value otherwise.

### GetRuleIdOk

`func (o *ApiTargetIntel) GetRuleIdOk() (*string, bool)`

GetRuleIdOk returns a tuple with the RuleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleId

`func (o *ApiTargetIntel) SetRuleId(v string)`

SetRuleId sets RuleId field to given value.

### HasRuleId

`func (o *ApiTargetIntel) HasRuleId() bool`

HasRuleId returns a boolean if a field has been set.

### GetRuleSource

`func (o *ApiTargetIntel) GetRuleSource() string`

GetRuleSource returns the RuleSource field if non-nil, zero value otherwise.

### GetRuleSourceOk

`func (o *ApiTargetIntel) GetRuleSourceOk() (*string, bool)`

GetRuleSourceOk returns a tuple with the RuleSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleSource

`func (o *ApiTargetIntel) SetRuleSource(v string)`

SetRuleSource sets RuleSource field to given value.

### HasRuleSource

`func (o *ApiTargetIntel) HasRuleSource() bool`

HasRuleSource returns a boolean if a field has been set.

### GetServiceMetadata

`func (o *ApiTargetIntel) GetServiceMetadata() []int32`

GetServiceMetadata returns the ServiceMetadata field if non-nil, zero value otherwise.

### GetServiceMetadataOk

`func (o *ApiTargetIntel) GetServiceMetadataOk() (*[]int32, bool)`

GetServiceMetadataOk returns a tuple with the ServiceMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceMetadata

`func (o *ApiTargetIntel) SetServiceMetadata(v []int32)`

SetServiceMetadata sets ServiceMetadata field to given value.

### HasServiceMetadata

`func (o *ApiTargetIntel) HasServiceMetadata() bool`

HasServiceMetadata returns a boolean if a field has been set.

### GetType

`func (o *ApiTargetIntel) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ApiTargetIntel) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ApiTargetIntel) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ApiTargetIntel) HasType() bool`

HasType returns a boolean if a field has been set.

### GetVendor

`func (o *ApiTargetIntel) GetVendor() string`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *ApiTargetIntel) GetVendorOk() (*string, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *ApiTargetIntel) SetVendor(v string)`

SetVendor sets Vendor field to given value.

### HasVendor

`func (o *ApiTargetIntel) HasVendor() bool`

HasVendor returns a boolean if a field has been set.

### GetVersion

`func (o *ApiTargetIntel) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ApiTargetIntel) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ApiTargetIntel) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ApiTargetIntel) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


