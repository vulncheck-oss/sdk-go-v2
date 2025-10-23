# ApiVulnCheckCanary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Category** | Pointer to **string** |  | [optional] 
**ClientFingerprints** | Pointer to [**ApiClientFingerprints**](ApiClientFingerprints.md) |  | [optional] 
**Cve** | Pointer to **string** |  | [optional] 
**DstCountry** | Pointer to **string** |  | [optional] 
**Http** | Pointer to [**ApiHTTPDetails**](ApiHTTPDetails.md) |  | [optional] 
**Severity** | Pointer to **int32** |  | [optional] 
**Signature** | Pointer to **string** |  | [optional] 
**SignatureId** | Pointer to **int32** |  | [optional] 
**SrcCountry** | Pointer to **string** |  | [optional] 
**SrcIp** | Pointer to **string** |  | [optional] 
**SrcPort** | Pointer to **int32** |  | [optional] 
**Timestamp** | Pointer to **string** |  | [optional] 

## Methods

### NewApiVulnCheckCanary

`func NewApiVulnCheckCanary() *ApiVulnCheckCanary`

NewApiVulnCheckCanary instantiates a new ApiVulnCheckCanary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiVulnCheckCanaryWithDefaults

`func NewApiVulnCheckCanaryWithDefaults() *ApiVulnCheckCanary`

NewApiVulnCheckCanaryWithDefaults instantiates a new ApiVulnCheckCanary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategory

`func (o *ApiVulnCheckCanary) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *ApiVulnCheckCanary) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *ApiVulnCheckCanary) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *ApiVulnCheckCanary) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetClientFingerprints

`func (o *ApiVulnCheckCanary) GetClientFingerprints() ApiClientFingerprints`

GetClientFingerprints returns the ClientFingerprints field if non-nil, zero value otherwise.

### GetClientFingerprintsOk

`func (o *ApiVulnCheckCanary) GetClientFingerprintsOk() (*ApiClientFingerprints, bool)`

GetClientFingerprintsOk returns a tuple with the ClientFingerprints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientFingerprints

`func (o *ApiVulnCheckCanary) SetClientFingerprints(v ApiClientFingerprints)`

SetClientFingerprints sets ClientFingerprints field to given value.

### HasClientFingerprints

`func (o *ApiVulnCheckCanary) HasClientFingerprints() bool`

HasClientFingerprints returns a boolean if a field has been set.

### GetCve

`func (o *ApiVulnCheckCanary) GetCve() string`

GetCve returns the Cve field if non-nil, zero value otherwise.

### GetCveOk

`func (o *ApiVulnCheckCanary) GetCveOk() (*string, bool)`

GetCveOk returns a tuple with the Cve field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCve

`func (o *ApiVulnCheckCanary) SetCve(v string)`

SetCve sets Cve field to given value.

### HasCve

`func (o *ApiVulnCheckCanary) HasCve() bool`

HasCve returns a boolean if a field has been set.

### GetDstCountry

`func (o *ApiVulnCheckCanary) GetDstCountry() string`

GetDstCountry returns the DstCountry field if non-nil, zero value otherwise.

### GetDstCountryOk

`func (o *ApiVulnCheckCanary) GetDstCountryOk() (*string, bool)`

GetDstCountryOk returns a tuple with the DstCountry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDstCountry

`func (o *ApiVulnCheckCanary) SetDstCountry(v string)`

SetDstCountry sets DstCountry field to given value.

### HasDstCountry

`func (o *ApiVulnCheckCanary) HasDstCountry() bool`

HasDstCountry returns a boolean if a field has been set.

### GetHttp

`func (o *ApiVulnCheckCanary) GetHttp() ApiHTTPDetails`

GetHttp returns the Http field if non-nil, zero value otherwise.

### GetHttpOk

`func (o *ApiVulnCheckCanary) GetHttpOk() (*ApiHTTPDetails, bool)`

GetHttpOk returns a tuple with the Http field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttp

`func (o *ApiVulnCheckCanary) SetHttp(v ApiHTTPDetails)`

SetHttp sets Http field to given value.

### HasHttp

`func (o *ApiVulnCheckCanary) HasHttp() bool`

HasHttp returns a boolean if a field has been set.

### GetSeverity

`func (o *ApiVulnCheckCanary) GetSeverity() int32`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *ApiVulnCheckCanary) GetSeverityOk() (*int32, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *ApiVulnCheckCanary) SetSeverity(v int32)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *ApiVulnCheckCanary) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetSignature

`func (o *ApiVulnCheckCanary) GetSignature() string`

GetSignature returns the Signature field if non-nil, zero value otherwise.

### GetSignatureOk

`func (o *ApiVulnCheckCanary) GetSignatureOk() (*string, bool)`

GetSignatureOk returns a tuple with the Signature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignature

`func (o *ApiVulnCheckCanary) SetSignature(v string)`

SetSignature sets Signature field to given value.

### HasSignature

`func (o *ApiVulnCheckCanary) HasSignature() bool`

HasSignature returns a boolean if a field has been set.

### GetSignatureId

`func (o *ApiVulnCheckCanary) GetSignatureId() int32`

GetSignatureId returns the SignatureId field if non-nil, zero value otherwise.

### GetSignatureIdOk

`func (o *ApiVulnCheckCanary) GetSignatureIdOk() (*int32, bool)`

GetSignatureIdOk returns a tuple with the SignatureId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignatureId

`func (o *ApiVulnCheckCanary) SetSignatureId(v int32)`

SetSignatureId sets SignatureId field to given value.

### HasSignatureId

`func (o *ApiVulnCheckCanary) HasSignatureId() bool`

HasSignatureId returns a boolean if a field has been set.

### GetSrcCountry

`func (o *ApiVulnCheckCanary) GetSrcCountry() string`

GetSrcCountry returns the SrcCountry field if non-nil, zero value otherwise.

### GetSrcCountryOk

`func (o *ApiVulnCheckCanary) GetSrcCountryOk() (*string, bool)`

GetSrcCountryOk returns a tuple with the SrcCountry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrcCountry

`func (o *ApiVulnCheckCanary) SetSrcCountry(v string)`

SetSrcCountry sets SrcCountry field to given value.

### HasSrcCountry

`func (o *ApiVulnCheckCanary) HasSrcCountry() bool`

HasSrcCountry returns a boolean if a field has been set.

### GetSrcIp

`func (o *ApiVulnCheckCanary) GetSrcIp() string`

GetSrcIp returns the SrcIp field if non-nil, zero value otherwise.

### GetSrcIpOk

`func (o *ApiVulnCheckCanary) GetSrcIpOk() (*string, bool)`

GetSrcIpOk returns a tuple with the SrcIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrcIp

`func (o *ApiVulnCheckCanary) SetSrcIp(v string)`

SetSrcIp sets SrcIp field to given value.

### HasSrcIp

`func (o *ApiVulnCheckCanary) HasSrcIp() bool`

HasSrcIp returns a boolean if a field has been set.

### GetSrcPort

`func (o *ApiVulnCheckCanary) GetSrcPort() int32`

GetSrcPort returns the SrcPort field if non-nil, zero value otherwise.

### GetSrcPortOk

`func (o *ApiVulnCheckCanary) GetSrcPortOk() (*int32, bool)`

GetSrcPortOk returns a tuple with the SrcPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrcPort

`func (o *ApiVulnCheckCanary) SetSrcPort(v int32)`

SetSrcPort sets SrcPort field to given value.

### HasSrcPort

`func (o *ApiVulnCheckCanary) HasSrcPort() bool`

HasSrcPort returns a boolean if a field has been set.

### GetTimestamp

`func (o *ApiVulnCheckCanary) GetTimestamp() string`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ApiVulnCheckCanary) GetTimestampOk() (*string, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ApiVulnCheckCanary) SetTimestamp(v string)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *ApiVulnCheckCanary) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


