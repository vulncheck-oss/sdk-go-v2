# AdvisoryIpIntelRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Asn** | Pointer to **string** |  | [optional] 
**City** | Pointer to **string** |  | [optional] 
**Country** | Pointer to **string** |  | [optional] 
**CountryCode** | Pointer to **string** |  | [optional] 
**Cve** | Pointer to **[]string** |  | [optional] 
**FeedIds** | Pointer to **[]string** |  | [optional] 
**Hostnames** | Pointer to **[]string** |  | [optional] 
**Ip** | Pointer to **string** |  | [optional] 
**LastSeen** | Pointer to **string** |  | [optional] 
**Matches** | Pointer to **[]string** |  | [optional] 
**Port** | Pointer to **int32** |  | [optional] 
**Ssl** | Pointer to **bool** |  | [optional] 
**Type** | Pointer to [**AdvisoryRecordType**](AdvisoryRecordType.md) |  | [optional] 

## Methods

### NewAdvisoryIpIntelRecord

`func NewAdvisoryIpIntelRecord() *AdvisoryIpIntelRecord`

NewAdvisoryIpIntelRecord instantiates a new AdvisoryIpIntelRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdvisoryIpIntelRecordWithDefaults

`func NewAdvisoryIpIntelRecordWithDefaults() *AdvisoryIpIntelRecord`

NewAdvisoryIpIntelRecordWithDefaults instantiates a new AdvisoryIpIntelRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAsn

`func (o *AdvisoryIpIntelRecord) GetAsn() string`

GetAsn returns the Asn field if non-nil, zero value otherwise.

### GetAsnOk

`func (o *AdvisoryIpIntelRecord) GetAsnOk() (*string, bool)`

GetAsnOk returns a tuple with the Asn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsn

`func (o *AdvisoryIpIntelRecord) SetAsn(v string)`

SetAsn sets Asn field to given value.

### HasAsn

`func (o *AdvisoryIpIntelRecord) HasAsn() bool`

HasAsn returns a boolean if a field has been set.

### GetCity

`func (o *AdvisoryIpIntelRecord) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *AdvisoryIpIntelRecord) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *AdvisoryIpIntelRecord) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *AdvisoryIpIntelRecord) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetCountry

`func (o *AdvisoryIpIntelRecord) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *AdvisoryIpIntelRecord) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *AdvisoryIpIntelRecord) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *AdvisoryIpIntelRecord) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetCountryCode

`func (o *AdvisoryIpIntelRecord) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *AdvisoryIpIntelRecord) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *AdvisoryIpIntelRecord) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.

### HasCountryCode

`func (o *AdvisoryIpIntelRecord) HasCountryCode() bool`

HasCountryCode returns a boolean if a field has been set.

### GetCve

`func (o *AdvisoryIpIntelRecord) GetCve() []string`

GetCve returns the Cve field if non-nil, zero value otherwise.

### GetCveOk

`func (o *AdvisoryIpIntelRecord) GetCveOk() (*[]string, bool)`

GetCveOk returns a tuple with the Cve field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCve

`func (o *AdvisoryIpIntelRecord) SetCve(v []string)`

SetCve sets Cve field to given value.

### HasCve

`func (o *AdvisoryIpIntelRecord) HasCve() bool`

HasCve returns a boolean if a field has been set.

### GetFeedIds

`func (o *AdvisoryIpIntelRecord) GetFeedIds() []string`

GetFeedIds returns the FeedIds field if non-nil, zero value otherwise.

### GetFeedIdsOk

`func (o *AdvisoryIpIntelRecord) GetFeedIdsOk() (*[]string, bool)`

GetFeedIdsOk returns a tuple with the FeedIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeedIds

`func (o *AdvisoryIpIntelRecord) SetFeedIds(v []string)`

SetFeedIds sets FeedIds field to given value.

### HasFeedIds

`func (o *AdvisoryIpIntelRecord) HasFeedIds() bool`

HasFeedIds returns a boolean if a field has been set.

### GetHostnames

`func (o *AdvisoryIpIntelRecord) GetHostnames() []string`

GetHostnames returns the Hostnames field if non-nil, zero value otherwise.

### GetHostnamesOk

`func (o *AdvisoryIpIntelRecord) GetHostnamesOk() (*[]string, bool)`

GetHostnamesOk returns a tuple with the Hostnames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostnames

`func (o *AdvisoryIpIntelRecord) SetHostnames(v []string)`

SetHostnames sets Hostnames field to given value.

### HasHostnames

`func (o *AdvisoryIpIntelRecord) HasHostnames() bool`

HasHostnames returns a boolean if a field has been set.

### GetIp

`func (o *AdvisoryIpIntelRecord) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *AdvisoryIpIntelRecord) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *AdvisoryIpIntelRecord) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *AdvisoryIpIntelRecord) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetLastSeen

`func (o *AdvisoryIpIntelRecord) GetLastSeen() string`

GetLastSeen returns the LastSeen field if non-nil, zero value otherwise.

### GetLastSeenOk

`func (o *AdvisoryIpIntelRecord) GetLastSeenOk() (*string, bool)`

GetLastSeenOk returns a tuple with the LastSeen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSeen

`func (o *AdvisoryIpIntelRecord) SetLastSeen(v string)`

SetLastSeen sets LastSeen field to given value.

### HasLastSeen

`func (o *AdvisoryIpIntelRecord) HasLastSeen() bool`

HasLastSeen returns a boolean if a field has been set.

### GetMatches

`func (o *AdvisoryIpIntelRecord) GetMatches() []string`

GetMatches returns the Matches field if non-nil, zero value otherwise.

### GetMatchesOk

`func (o *AdvisoryIpIntelRecord) GetMatchesOk() (*[]string, bool)`

GetMatchesOk returns a tuple with the Matches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatches

`func (o *AdvisoryIpIntelRecord) SetMatches(v []string)`

SetMatches sets Matches field to given value.

### HasMatches

`func (o *AdvisoryIpIntelRecord) HasMatches() bool`

HasMatches returns a boolean if a field has been set.

### GetPort

`func (o *AdvisoryIpIntelRecord) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *AdvisoryIpIntelRecord) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *AdvisoryIpIntelRecord) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *AdvisoryIpIntelRecord) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetSsl

`func (o *AdvisoryIpIntelRecord) GetSsl() bool`

GetSsl returns the Ssl field if non-nil, zero value otherwise.

### GetSslOk

`func (o *AdvisoryIpIntelRecord) GetSslOk() (*bool, bool)`

GetSslOk returns a tuple with the Ssl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsl

`func (o *AdvisoryIpIntelRecord) SetSsl(v bool)`

SetSsl sets Ssl field to given value.

### HasSsl

`func (o *AdvisoryIpIntelRecord) HasSsl() bool`

HasSsl returns a boolean if a field has been set.

### GetType

`func (o *AdvisoryIpIntelRecord) GetType() AdvisoryRecordType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AdvisoryIpIntelRecord) GetTypeOk() (*AdvisoryRecordType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AdvisoryIpIntelRecord) SetType(v AdvisoryRecordType)`

SetType sets Type field to given value.

### HasType

`func (o *AdvisoryIpIntelRecord) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


