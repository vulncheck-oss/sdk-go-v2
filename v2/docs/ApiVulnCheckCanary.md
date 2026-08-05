# ApiVulnCheckCanary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**C2Frequency3d** | Pointer to [**[]ApiC2Frequency**](ApiC2Frequency.md) |  | [optional] 
**C2Location** | Pointer to **[]string** |  | [optional] 
**Category** | Pointer to **string** |  | [optional] 
**ClientFingerprints** | Pointer to [**ApiClientFingerprints**](ApiClientFingerprints.md) |  | [optional] 
**Cve** | Pointer to **string** |  | [optional] 
**DstCountry** | Pointer to **string** |  | [optional] 
**Http** | Pointer to [**ApiHTTPDetails**](ApiHTTPDetails.md) |  | [optional] 
**Payload** | Pointer to **string** |  | [optional] 
**PayloadTlsh** | Pointer to **string** |  | [optional] 
**PayloadTooling** | Pointer to **[]string** |  | [optional] 
**Severity** | Pointer to **int32** |  | [optional] 
**Signature** | Pointer to **string** |  | [optional] 
**SignatureId** | Pointer to **int32** |  | [optional] 
**SrcAsDomain** | Pointer to **string** |  | [optional] 
**SrcAsName** | Pointer to **string** |  | [optional] 
**SrcAsn** | Pointer to **string** |  | [optional] 
**SrcCountry** | Pointer to **string** |  | [optional] 
**SrcIp** | Pointer to **string** |  | [optional] 
**SrcIpFreq3d** | Pointer to **int32** |  | [optional] 
**SrcIpFreq3dCanary** | Pointer to **int32** |  | [optional] 
**SrcIpTypeFindings** | Pointer to **[]string** |  | [optional] 
**SrcPort** | Pointer to **int32** |  | [optional] 
**TechVertical** | Pointer to **[]string** |  | [optional] 
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

### GetC2Frequency3d

`func (o *ApiVulnCheckCanary) GetC2Frequency3d() []ApiC2Frequency`

GetC2Frequency3d returns the C2Frequency3d field if non-nil, zero value otherwise.

### GetC2Frequency3dOk

`func (o *ApiVulnCheckCanary) GetC2Frequency3dOk() (*[]ApiC2Frequency, bool)`

GetC2Frequency3dOk returns a tuple with the C2Frequency3d field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetC2Frequency3d

`func (o *ApiVulnCheckCanary) SetC2Frequency3d(v []ApiC2Frequency)`

SetC2Frequency3d sets C2Frequency3d field to given value.

### HasC2Frequency3d

`func (o *ApiVulnCheckCanary) HasC2Frequency3d() bool`

HasC2Frequency3d returns a boolean if a field has been set.

### GetC2Location

`func (o *ApiVulnCheckCanary) GetC2Location() []string`

GetC2Location returns the C2Location field if non-nil, zero value otherwise.

### GetC2LocationOk

`func (o *ApiVulnCheckCanary) GetC2LocationOk() (*[]string, bool)`

GetC2LocationOk returns a tuple with the C2Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetC2Location

`func (o *ApiVulnCheckCanary) SetC2Location(v []string)`

SetC2Location sets C2Location field to given value.

### HasC2Location

`func (o *ApiVulnCheckCanary) HasC2Location() bool`

HasC2Location returns a boolean if a field has been set.

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

### GetPayload

`func (o *ApiVulnCheckCanary) GetPayload() string`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *ApiVulnCheckCanary) GetPayloadOk() (*string, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *ApiVulnCheckCanary) SetPayload(v string)`

SetPayload sets Payload field to given value.

### HasPayload

`func (o *ApiVulnCheckCanary) HasPayload() bool`

HasPayload returns a boolean if a field has been set.

### GetPayloadTlsh

`func (o *ApiVulnCheckCanary) GetPayloadTlsh() string`

GetPayloadTlsh returns the PayloadTlsh field if non-nil, zero value otherwise.

### GetPayloadTlshOk

`func (o *ApiVulnCheckCanary) GetPayloadTlshOk() (*string, bool)`

GetPayloadTlshOk returns a tuple with the PayloadTlsh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayloadTlsh

`func (o *ApiVulnCheckCanary) SetPayloadTlsh(v string)`

SetPayloadTlsh sets PayloadTlsh field to given value.

### HasPayloadTlsh

`func (o *ApiVulnCheckCanary) HasPayloadTlsh() bool`

HasPayloadTlsh returns a boolean if a field has been set.

### GetPayloadTooling

`func (o *ApiVulnCheckCanary) GetPayloadTooling() []string`

GetPayloadTooling returns the PayloadTooling field if non-nil, zero value otherwise.

### GetPayloadToolingOk

`func (o *ApiVulnCheckCanary) GetPayloadToolingOk() (*[]string, bool)`

GetPayloadToolingOk returns a tuple with the PayloadTooling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayloadTooling

`func (o *ApiVulnCheckCanary) SetPayloadTooling(v []string)`

SetPayloadTooling sets PayloadTooling field to given value.

### HasPayloadTooling

`func (o *ApiVulnCheckCanary) HasPayloadTooling() bool`

HasPayloadTooling returns a boolean if a field has been set.

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

### GetSrcAsDomain

`func (o *ApiVulnCheckCanary) GetSrcAsDomain() string`

GetSrcAsDomain returns the SrcAsDomain field if non-nil, zero value otherwise.

### GetSrcAsDomainOk

`func (o *ApiVulnCheckCanary) GetSrcAsDomainOk() (*string, bool)`

GetSrcAsDomainOk returns a tuple with the SrcAsDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrcAsDomain

`func (o *ApiVulnCheckCanary) SetSrcAsDomain(v string)`

SetSrcAsDomain sets SrcAsDomain field to given value.

### HasSrcAsDomain

`func (o *ApiVulnCheckCanary) HasSrcAsDomain() bool`

HasSrcAsDomain returns a boolean if a field has been set.

### GetSrcAsName

`func (o *ApiVulnCheckCanary) GetSrcAsName() string`

GetSrcAsName returns the SrcAsName field if non-nil, zero value otherwise.

### GetSrcAsNameOk

`func (o *ApiVulnCheckCanary) GetSrcAsNameOk() (*string, bool)`

GetSrcAsNameOk returns a tuple with the SrcAsName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrcAsName

`func (o *ApiVulnCheckCanary) SetSrcAsName(v string)`

SetSrcAsName sets SrcAsName field to given value.

### HasSrcAsName

`func (o *ApiVulnCheckCanary) HasSrcAsName() bool`

HasSrcAsName returns a boolean if a field has been set.

### GetSrcAsn

`func (o *ApiVulnCheckCanary) GetSrcAsn() string`

GetSrcAsn returns the SrcAsn field if non-nil, zero value otherwise.

### GetSrcAsnOk

`func (o *ApiVulnCheckCanary) GetSrcAsnOk() (*string, bool)`

GetSrcAsnOk returns a tuple with the SrcAsn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrcAsn

`func (o *ApiVulnCheckCanary) SetSrcAsn(v string)`

SetSrcAsn sets SrcAsn field to given value.

### HasSrcAsn

`func (o *ApiVulnCheckCanary) HasSrcAsn() bool`

HasSrcAsn returns a boolean if a field has been set.

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

### GetSrcIpFreq3d

`func (o *ApiVulnCheckCanary) GetSrcIpFreq3d() int32`

GetSrcIpFreq3d returns the SrcIpFreq3d field if non-nil, zero value otherwise.

### GetSrcIpFreq3dOk

`func (o *ApiVulnCheckCanary) GetSrcIpFreq3dOk() (*int32, bool)`

GetSrcIpFreq3dOk returns a tuple with the SrcIpFreq3d field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrcIpFreq3d

`func (o *ApiVulnCheckCanary) SetSrcIpFreq3d(v int32)`

SetSrcIpFreq3d sets SrcIpFreq3d field to given value.

### HasSrcIpFreq3d

`func (o *ApiVulnCheckCanary) HasSrcIpFreq3d() bool`

HasSrcIpFreq3d returns a boolean if a field has been set.

### GetSrcIpFreq3dCanary

`func (o *ApiVulnCheckCanary) GetSrcIpFreq3dCanary() int32`

GetSrcIpFreq3dCanary returns the SrcIpFreq3dCanary field if non-nil, zero value otherwise.

### GetSrcIpFreq3dCanaryOk

`func (o *ApiVulnCheckCanary) GetSrcIpFreq3dCanaryOk() (*int32, bool)`

GetSrcIpFreq3dCanaryOk returns a tuple with the SrcIpFreq3dCanary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrcIpFreq3dCanary

`func (o *ApiVulnCheckCanary) SetSrcIpFreq3dCanary(v int32)`

SetSrcIpFreq3dCanary sets SrcIpFreq3dCanary field to given value.

### HasSrcIpFreq3dCanary

`func (o *ApiVulnCheckCanary) HasSrcIpFreq3dCanary() bool`

HasSrcIpFreq3dCanary returns a boolean if a field has been set.

### GetSrcIpTypeFindings

`func (o *ApiVulnCheckCanary) GetSrcIpTypeFindings() []string`

GetSrcIpTypeFindings returns the SrcIpTypeFindings field if non-nil, zero value otherwise.

### GetSrcIpTypeFindingsOk

`func (o *ApiVulnCheckCanary) GetSrcIpTypeFindingsOk() (*[]string, bool)`

GetSrcIpTypeFindingsOk returns a tuple with the SrcIpTypeFindings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrcIpTypeFindings

`func (o *ApiVulnCheckCanary) SetSrcIpTypeFindings(v []string)`

SetSrcIpTypeFindings sets SrcIpTypeFindings field to given value.

### HasSrcIpTypeFindings

`func (o *ApiVulnCheckCanary) HasSrcIpTypeFindings() bool`

HasSrcIpTypeFindings returns a boolean if a field has been set.

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

### GetTechVertical

`func (o *ApiVulnCheckCanary) GetTechVertical() []string`

GetTechVertical returns the TechVertical field if non-nil, zero value otherwise.

### GetTechVerticalOk

`func (o *ApiVulnCheckCanary) GetTechVerticalOk() (*[]string, bool)`

GetTechVerticalOk returns a tuple with the TechVertical field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTechVertical

`func (o *ApiVulnCheckCanary) SetTechVertical(v []string)`

SetTechVertical sets TechVertical field to given value.

### HasTechVertical

`func (o *ApiVulnCheckCanary) HasTechVertical() bool`

HasTechVertical returns a boolean if a field has been set.

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


