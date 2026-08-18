# ApiTargetIntel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AsDomain** | Pointer to **string** |  | [optional] 
**AsName** | Pointer to **string** |  | [optional] 
**Asn** | Pointer to **string** |  | [optional] 
**Classifications** | Pointer to **[]string** |  | [optional] 
**ContainsCve** | Pointer to **bool** | Deprecated: use Summary.ContainsCVE instead. | [optional] 
**Country** | Pointer to **string** |  | [optional] 
**CountryCode** | Pointer to **string** |  | [optional] 
**Cpe** | Pointer to **[]string** |  | [optional] 
**Cve** | Pointer to **[]string** | Deprecated: use Fingerprints[].CVEs for per-fingerprint attribution, or Summary for the aggregate count. | [optional] 
**CveConfirmed** | Pointer to [**[]ApiCVEConfirmed**](ApiCVEConfirmed.md) | Deprecated: use Fingerprints[].CVEs for per-fingerprint attribution, or Summary for the aggregate confirmed count. | [optional] 
**DateAdded** | Pointer to **string** |  | [optional] 
**Fingerprints** | Pointer to [**[]ApiFingerprint**](ApiFingerprint.md) |  | [optional] 
**Hostname** | Pointer to **string** |  | [optional] 
**Ip** | Pointer to **string** |  | [optional] 
**Metadata** | Pointer to **[]int32** |  | [optional] 
**Port** | Pointer to **int32** |  | [optional] 
**Product** | Pointer to **[]string** |  | [optional] 
**Protocol** | Pointer to **string** |  | [optional] 
**Summary** | Pointer to [**ApiTargetIntelSummary**](ApiTargetIntelSummary.md) |  | [optional] 
**Timestamp** | Pointer to **string** |  | [optional] 
**Transport** | Pointer to **string** |  | [optional] 
**Vendor** | Pointer to **[]string** |  | [optional] 
**Version** | Pointer to **[]string** |  | [optional] 

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

### GetClassifications

`func (o *ApiTargetIntel) GetClassifications() []string`

GetClassifications returns the Classifications field if non-nil, zero value otherwise.

### GetClassificationsOk

`func (o *ApiTargetIntel) GetClassificationsOk() (*[]string, bool)`

GetClassificationsOk returns a tuple with the Classifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassifications

`func (o *ApiTargetIntel) SetClassifications(v []string)`

SetClassifications sets Classifications field to given value.

### HasClassifications

`func (o *ApiTargetIntel) HasClassifications() bool`

HasClassifications returns a boolean if a field has been set.

### GetContainsCve

`func (o *ApiTargetIntel) GetContainsCve() bool`

GetContainsCve returns the ContainsCve field if non-nil, zero value otherwise.

### GetContainsCveOk

`func (o *ApiTargetIntel) GetContainsCveOk() (*bool, bool)`

GetContainsCveOk returns a tuple with the ContainsCve field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainsCve

`func (o *ApiTargetIntel) SetContainsCve(v bool)`

SetContainsCve sets ContainsCve field to given value.

### HasContainsCve

`func (o *ApiTargetIntel) HasContainsCve() bool`

HasContainsCve returns a boolean if a field has been set.

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

`func (o *ApiTargetIntel) GetCpe() []string`

GetCpe returns the Cpe field if non-nil, zero value otherwise.

### GetCpeOk

`func (o *ApiTargetIntel) GetCpeOk() (*[]string, bool)`

GetCpeOk returns a tuple with the Cpe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpe

`func (o *ApiTargetIntel) SetCpe(v []string)`

SetCpe sets Cpe field to given value.

### HasCpe

`func (o *ApiTargetIntel) HasCpe() bool`

HasCpe returns a boolean if a field has been set.

### GetCve

`func (o *ApiTargetIntel) GetCve() []string`

GetCve returns the Cve field if non-nil, zero value otherwise.

### GetCveOk

`func (o *ApiTargetIntel) GetCveOk() (*[]string, bool)`

GetCveOk returns a tuple with the Cve field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCve

`func (o *ApiTargetIntel) SetCve(v []string)`

SetCve sets Cve field to given value.

### HasCve

`func (o *ApiTargetIntel) HasCve() bool`

HasCve returns a boolean if a field has been set.

### GetCveConfirmed

`func (o *ApiTargetIntel) GetCveConfirmed() []ApiCVEConfirmed`

GetCveConfirmed returns the CveConfirmed field if non-nil, zero value otherwise.

### GetCveConfirmedOk

`func (o *ApiTargetIntel) GetCveConfirmedOk() (*[]ApiCVEConfirmed, bool)`

GetCveConfirmedOk returns a tuple with the CveConfirmed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCveConfirmed

`func (o *ApiTargetIntel) SetCveConfirmed(v []ApiCVEConfirmed)`

SetCveConfirmed sets CveConfirmed field to given value.

### HasCveConfirmed

`func (o *ApiTargetIntel) HasCveConfirmed() bool`

HasCveConfirmed returns a boolean if a field has been set.

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

### GetFingerprints

`func (o *ApiTargetIntel) GetFingerprints() []ApiFingerprint`

GetFingerprints returns the Fingerprints field if non-nil, zero value otherwise.

### GetFingerprintsOk

`func (o *ApiTargetIntel) GetFingerprintsOk() (*[]ApiFingerprint, bool)`

GetFingerprintsOk returns a tuple with the Fingerprints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFingerprints

`func (o *ApiTargetIntel) SetFingerprints(v []ApiFingerprint)`

SetFingerprints sets Fingerprints field to given value.

### HasFingerprints

`func (o *ApiTargetIntel) HasFingerprints() bool`

HasFingerprints returns a boolean if a field has been set.

### GetHostname

`func (o *ApiTargetIntel) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *ApiTargetIntel) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *ApiTargetIntel) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *ApiTargetIntel) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

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

### GetMetadata

`func (o *ApiTargetIntel) GetMetadata() []int32`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ApiTargetIntel) GetMetadataOk() (*[]int32, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ApiTargetIntel) SetMetadata(v []int32)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ApiTargetIntel) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

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

`func (o *ApiTargetIntel) GetProduct() []string`

GetProduct returns the Product field if non-nil, zero value otherwise.

### GetProductOk

`func (o *ApiTargetIntel) GetProductOk() (*[]string, bool)`

GetProductOk returns a tuple with the Product field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduct

`func (o *ApiTargetIntel) SetProduct(v []string)`

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

### GetSummary

`func (o *ApiTargetIntel) GetSummary() ApiTargetIntelSummary`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *ApiTargetIntel) GetSummaryOk() (*ApiTargetIntelSummary, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *ApiTargetIntel) SetSummary(v ApiTargetIntelSummary)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *ApiTargetIntel) HasSummary() bool`

HasSummary returns a boolean if a field has been set.

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

### GetTransport

`func (o *ApiTargetIntel) GetTransport() string`

GetTransport returns the Transport field if non-nil, zero value otherwise.

### GetTransportOk

`func (o *ApiTargetIntel) GetTransportOk() (*string, bool)`

GetTransportOk returns a tuple with the Transport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransport

`func (o *ApiTargetIntel) SetTransport(v string)`

SetTransport sets Transport field to given value.

### HasTransport

`func (o *ApiTargetIntel) HasTransport() bool`

HasTransport returns a boolean if a field has been set.

### GetVendor

`func (o *ApiTargetIntel) GetVendor() []string`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *ApiTargetIntel) GetVendorOk() (*[]string, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *ApiTargetIntel) SetVendor(v []string)`

SetVendor sets Vendor field to given value.

### HasVendor

`func (o *ApiTargetIntel) HasVendor() bool`

HasVendor returns a boolean if a field has been set.

### GetVersion

`func (o *ApiTargetIntel) GetVersion() []string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ApiTargetIntel) GetVersionOk() (*[]string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ApiTargetIntel) SetVersion(v []string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ApiTargetIntel) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


