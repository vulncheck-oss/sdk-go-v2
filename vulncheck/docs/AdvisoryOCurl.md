# AdvisoryOCurl

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Affected** | Pointer to [**[]AdvisoryCurlAffected**](AdvisoryCurlAffected.md) |  | [optional] 
**Aliases** | Pointer to **[]string** |  | [optional] 
**Credits** | Pointer to [**[]AdvisoryCurlCredit**](AdvisoryCurlCredit.md) |  | [optional] 
**DatabaseSpecific** | Pointer to [**AdvisoryDBSpecific**](AdvisoryDBSpecific.md) |  | [optional] 
**Details** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Modified** | Pointer to **string** |  | [optional] 
**Published** | Pointer to **string** |  | [optional] 
**SchemaVersion** | Pointer to **string** |  | [optional] 
**Summary** | Pointer to **string** |  | [optional] 

## Methods

### NewAdvisoryOCurl

`func NewAdvisoryOCurl() *AdvisoryOCurl`

NewAdvisoryOCurl instantiates a new AdvisoryOCurl object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdvisoryOCurlWithDefaults

`func NewAdvisoryOCurlWithDefaults() *AdvisoryOCurl`

NewAdvisoryOCurlWithDefaults instantiates a new AdvisoryOCurl object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAffected

`func (o *AdvisoryOCurl) GetAffected() []AdvisoryCurlAffected`

GetAffected returns the Affected field if non-nil, zero value otherwise.

### GetAffectedOk

`func (o *AdvisoryOCurl) GetAffectedOk() (*[]AdvisoryCurlAffected, bool)`

GetAffectedOk returns a tuple with the Affected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffected

`func (o *AdvisoryOCurl) SetAffected(v []AdvisoryCurlAffected)`

SetAffected sets Affected field to given value.

### HasAffected

`func (o *AdvisoryOCurl) HasAffected() bool`

HasAffected returns a boolean if a field has been set.

### GetAliases

`func (o *AdvisoryOCurl) GetAliases() []string`

GetAliases returns the Aliases field if non-nil, zero value otherwise.

### GetAliasesOk

`func (o *AdvisoryOCurl) GetAliasesOk() (*[]string, bool)`

GetAliasesOk returns a tuple with the Aliases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAliases

`func (o *AdvisoryOCurl) SetAliases(v []string)`

SetAliases sets Aliases field to given value.

### HasAliases

`func (o *AdvisoryOCurl) HasAliases() bool`

HasAliases returns a boolean if a field has been set.

### GetCredits

`func (o *AdvisoryOCurl) GetCredits() []AdvisoryCurlCredit`

GetCredits returns the Credits field if non-nil, zero value otherwise.

### GetCreditsOk

`func (o *AdvisoryOCurl) GetCreditsOk() (*[]AdvisoryCurlCredit, bool)`

GetCreditsOk returns a tuple with the Credits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredits

`func (o *AdvisoryOCurl) SetCredits(v []AdvisoryCurlCredit)`

SetCredits sets Credits field to given value.

### HasCredits

`func (o *AdvisoryOCurl) HasCredits() bool`

HasCredits returns a boolean if a field has been set.

### GetDatabaseSpecific

`func (o *AdvisoryOCurl) GetDatabaseSpecific() AdvisoryDBSpecific`

GetDatabaseSpecific returns the DatabaseSpecific field if non-nil, zero value otherwise.

### GetDatabaseSpecificOk

`func (o *AdvisoryOCurl) GetDatabaseSpecificOk() (*AdvisoryDBSpecific, bool)`

GetDatabaseSpecificOk returns a tuple with the DatabaseSpecific field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseSpecific

`func (o *AdvisoryOCurl) SetDatabaseSpecific(v AdvisoryDBSpecific)`

SetDatabaseSpecific sets DatabaseSpecific field to given value.

### HasDatabaseSpecific

`func (o *AdvisoryOCurl) HasDatabaseSpecific() bool`

HasDatabaseSpecific returns a boolean if a field has been set.

### GetDetails

`func (o *AdvisoryOCurl) GetDetails() string`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *AdvisoryOCurl) GetDetailsOk() (*string, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *AdvisoryOCurl) SetDetails(v string)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *AdvisoryOCurl) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### GetId

`func (o *AdvisoryOCurl) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AdvisoryOCurl) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AdvisoryOCurl) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AdvisoryOCurl) HasId() bool`

HasId returns a boolean if a field has been set.

### GetModified

`func (o *AdvisoryOCurl) GetModified() string`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *AdvisoryOCurl) GetModifiedOk() (*string, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *AdvisoryOCurl) SetModified(v string)`

SetModified sets Modified field to given value.

### HasModified

`func (o *AdvisoryOCurl) HasModified() bool`

HasModified returns a boolean if a field has been set.

### GetPublished

`func (o *AdvisoryOCurl) GetPublished() string`

GetPublished returns the Published field if non-nil, zero value otherwise.

### GetPublishedOk

`func (o *AdvisoryOCurl) GetPublishedOk() (*string, bool)`

GetPublishedOk returns a tuple with the Published field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublished

`func (o *AdvisoryOCurl) SetPublished(v string)`

SetPublished sets Published field to given value.

### HasPublished

`func (o *AdvisoryOCurl) HasPublished() bool`

HasPublished returns a boolean if a field has been set.

### GetSchemaVersion

`func (o *AdvisoryOCurl) GetSchemaVersion() string`

GetSchemaVersion returns the SchemaVersion field if non-nil, zero value otherwise.

### GetSchemaVersionOk

`func (o *AdvisoryOCurl) GetSchemaVersionOk() (*string, bool)`

GetSchemaVersionOk returns a tuple with the SchemaVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaVersion

`func (o *AdvisoryOCurl) SetSchemaVersion(v string)`

SetSchemaVersion sets SchemaVersion field to given value.

### HasSchemaVersion

`func (o *AdvisoryOCurl) HasSchemaVersion() bool`

HasSchemaVersion returns a boolean if a field has been set.

### GetSummary

`func (o *AdvisoryOCurl) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *AdvisoryOCurl) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *AdvisoryOCurl) SetSummary(v string)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *AdvisoryOCurl) HasSummary() bool`

HasSummary returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


