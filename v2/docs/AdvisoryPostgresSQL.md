# AdvisoryPostgresSQL

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cve** | Pointer to **[]string** |  | [optional] 
**DateAdded** | Pointer to **string** |  | [optional] 
**PgFix** | Pointer to [**[]AdvisoryPGFix**](AdvisoryPGFix.md) |  | [optional] 
**Summary** | Pointer to **string** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 

## Methods

### NewAdvisoryPostgresSQL

`func NewAdvisoryPostgresSQL() *AdvisoryPostgresSQL`

NewAdvisoryPostgresSQL instantiates a new AdvisoryPostgresSQL object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdvisoryPostgresSQLWithDefaults

`func NewAdvisoryPostgresSQLWithDefaults() *AdvisoryPostgresSQL`

NewAdvisoryPostgresSQLWithDefaults instantiates a new AdvisoryPostgresSQL object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCve

`func (o *AdvisoryPostgresSQL) GetCve() []string`

GetCve returns the Cve field if non-nil, zero value otherwise.

### GetCveOk

`func (o *AdvisoryPostgresSQL) GetCveOk() (*[]string, bool)`

GetCveOk returns a tuple with the Cve field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCve

`func (o *AdvisoryPostgresSQL) SetCve(v []string)`

SetCve sets Cve field to given value.

### HasCve

`func (o *AdvisoryPostgresSQL) HasCve() bool`

HasCve returns a boolean if a field has been set.

### GetDateAdded

`func (o *AdvisoryPostgresSQL) GetDateAdded() string`

GetDateAdded returns the DateAdded field if non-nil, zero value otherwise.

### GetDateAddedOk

`func (o *AdvisoryPostgresSQL) GetDateAddedOk() (*string, bool)`

GetDateAddedOk returns a tuple with the DateAdded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateAdded

`func (o *AdvisoryPostgresSQL) SetDateAdded(v string)`

SetDateAdded sets DateAdded field to given value.

### HasDateAdded

`func (o *AdvisoryPostgresSQL) HasDateAdded() bool`

HasDateAdded returns a boolean if a field has been set.

### GetPgFix

`func (o *AdvisoryPostgresSQL) GetPgFix() []AdvisoryPGFix`

GetPgFix returns the PgFix field if non-nil, zero value otherwise.

### GetPgFixOk

`func (o *AdvisoryPostgresSQL) GetPgFixOk() (*[]AdvisoryPGFix, bool)`

GetPgFixOk returns a tuple with the PgFix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPgFix

`func (o *AdvisoryPostgresSQL) SetPgFix(v []AdvisoryPGFix)`

SetPgFix sets PgFix field to given value.

### HasPgFix

`func (o *AdvisoryPostgresSQL) HasPgFix() bool`

HasPgFix returns a boolean if a field has been set.

### GetSummary

`func (o *AdvisoryPostgresSQL) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *AdvisoryPostgresSQL) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *AdvisoryPostgresSQL) SetSummary(v string)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *AdvisoryPostgresSQL) HasSummary() bool`

HasSummary returns a boolean if a field has been set.

### GetTitle

`func (o *AdvisoryPostgresSQL) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *AdvisoryPostgresSQL) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *AdvisoryPostgresSQL) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *AdvisoryPostgresSQL) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetUrl

`func (o *AdvisoryPostgresSQL) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *AdvisoryPostgresSQL) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *AdvisoryPostgresSQL) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *AdvisoryPostgresSQL) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


