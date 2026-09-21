# ApiC2Community

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AsName** | Pointer to **string** |  | [optional] 
**Asn** | Pointer to **string** |  | [optional] 
**Classifications** | Pointer to **[]string** |  | [optional] 
**Country** | Pointer to **string** |  | [optional] 
**CountryCode** | Pointer to **string** |  | [optional] 
**FirstSeen** | Pointer to **string** |  | [optional] 
**Hostname** | Pointer to **string** |  | [optional] 
**Ip** | Pointer to **string** |  | [optional] 
**LastSeen** | Pointer to **string** |  | [optional] 
**Port** | Pointer to **int32** |  | [optional] 
**Source** | Pointer to **[]string** |  | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 

## Methods

### NewApiC2Community

`func NewApiC2Community() *ApiC2Community`

NewApiC2Community instantiates a new ApiC2Community object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiC2CommunityWithDefaults

`func NewApiC2CommunityWithDefaults() *ApiC2Community`

NewApiC2CommunityWithDefaults instantiates a new ApiC2Community object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAsName

`func (o *ApiC2Community) GetAsName() string`

GetAsName returns the AsName field if non-nil, zero value otherwise.

### GetAsNameOk

`func (o *ApiC2Community) GetAsNameOk() (*string, bool)`

GetAsNameOk returns a tuple with the AsName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsName

`func (o *ApiC2Community) SetAsName(v string)`

SetAsName sets AsName field to given value.

### HasAsName

`func (o *ApiC2Community) HasAsName() bool`

HasAsName returns a boolean if a field has been set.

### GetAsn

`func (o *ApiC2Community) GetAsn() string`

GetAsn returns the Asn field if non-nil, zero value otherwise.

### GetAsnOk

`func (o *ApiC2Community) GetAsnOk() (*string, bool)`

GetAsnOk returns a tuple with the Asn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsn

`func (o *ApiC2Community) SetAsn(v string)`

SetAsn sets Asn field to given value.

### HasAsn

`func (o *ApiC2Community) HasAsn() bool`

HasAsn returns a boolean if a field has been set.

### GetClassifications

`func (o *ApiC2Community) GetClassifications() []string`

GetClassifications returns the Classifications field if non-nil, zero value otherwise.

### GetClassificationsOk

`func (o *ApiC2Community) GetClassificationsOk() (*[]string, bool)`

GetClassificationsOk returns a tuple with the Classifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassifications

`func (o *ApiC2Community) SetClassifications(v []string)`

SetClassifications sets Classifications field to given value.

### HasClassifications

`func (o *ApiC2Community) HasClassifications() bool`

HasClassifications returns a boolean if a field has been set.

### GetCountry

`func (o *ApiC2Community) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *ApiC2Community) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *ApiC2Community) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *ApiC2Community) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetCountryCode

`func (o *ApiC2Community) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *ApiC2Community) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *ApiC2Community) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.

### HasCountryCode

`func (o *ApiC2Community) HasCountryCode() bool`

HasCountryCode returns a boolean if a field has been set.

### GetFirstSeen

`func (o *ApiC2Community) GetFirstSeen() string`

GetFirstSeen returns the FirstSeen field if non-nil, zero value otherwise.

### GetFirstSeenOk

`func (o *ApiC2Community) GetFirstSeenOk() (*string, bool)`

GetFirstSeenOk returns a tuple with the FirstSeen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstSeen

`func (o *ApiC2Community) SetFirstSeen(v string)`

SetFirstSeen sets FirstSeen field to given value.

### HasFirstSeen

`func (o *ApiC2Community) HasFirstSeen() bool`

HasFirstSeen returns a boolean if a field has been set.

### GetHostname

`func (o *ApiC2Community) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *ApiC2Community) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *ApiC2Community) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *ApiC2Community) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetIp

`func (o *ApiC2Community) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *ApiC2Community) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *ApiC2Community) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *ApiC2Community) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetLastSeen

`func (o *ApiC2Community) GetLastSeen() string`

GetLastSeen returns the LastSeen field if non-nil, zero value otherwise.

### GetLastSeenOk

`func (o *ApiC2Community) GetLastSeenOk() (*string, bool)`

GetLastSeenOk returns a tuple with the LastSeen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSeen

`func (o *ApiC2Community) SetLastSeen(v string)`

SetLastSeen sets LastSeen field to given value.

### HasLastSeen

`func (o *ApiC2Community) HasLastSeen() bool`

HasLastSeen returns a boolean if a field has been set.

### GetPort

`func (o *ApiC2Community) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *ApiC2Community) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *ApiC2Community) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *ApiC2Community) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetSource

`func (o *ApiC2Community) GetSource() []string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *ApiC2Community) GetSourceOk() (*[]string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *ApiC2Community) SetSource(v []string)`

SetSource sets Source field to given value.

### HasSource

`func (o *ApiC2Community) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *ApiC2Community) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ApiC2Community) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ApiC2Community) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ApiC2Community) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


