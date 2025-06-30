# ApiInitialAccessArtifact

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ArtifactName** | Pointer to **string** | ArtifactName is a title to associate with this artifact. | [optional] 
**ArtifactsURL** | Pointer to **[]string** | ArtifactsURL are URLs to the available artifact. | [optional] 
**BaiduQueries** | Pointer to **[]string** | ... | [optional] 
**BaiduRawQueries** | Pointer to **[]string** | ... | [optional] 
**CensysQueries** | Pointer to **[]string** | CensysQueries are queries for examining potential Internet-exposed devices &amp; applications with Censys in URL form. | [optional] 
**CensysRawQueries** | Pointer to **[]string** | CensysRawQueries are raw queries for examining potential Internet-exposed devices &amp; applications with Censys. | [optional] 
**CloneSSHURL** | Pointer to **string** | CloneSSHURL is the git URL to clone the artifact with. | [optional] 
**DateAdded** | Pointer to **string** | DateAdded is when this artifact entry was first added to the InitialAccess data set. | [optional] 
**Exploit** | Pointer to **bool** | Exploit indicates whether or not an exploit is available in this artifact. | [optional] 
**FofaQueries** | Pointer to **[]string** | FOFAQueries are raw queries for examining potential Internet-exposed devices &amp; applications with FOFA. | [optional] 
**FofaRawQueries** | Pointer to **[]string** |  | [optional] 
**GoogleQueries** | Pointer to **[]string** | google queries | [optional] 
**GoogleRawQueries** | Pointer to **[]string** | raw google queries | [optional] 
**GreynoiseQueries** | Pointer to **[]string** | GreynoiseQueries are queries for finding the vulnerability via honeypot data. | [optional] 
**MitreAttackTechniques** | Pointer to **[]string** | MITRE ATT&amp;CK techniques | [optional] 
**NmapScript** | Pointer to **bool** | NmapScript indicates whether or not an nmap script for scanning environment exists in this artifact. | [optional] 
**Pcap** | Pointer to **bool** | PCAP indicates whether of not a package capture of the exploit PoC exploiting a vulnerable system exists in this artifact. | [optional] 
**Product** | Pointer to **[]string** | Product are the software that has the vulnerability. | [optional] 
**ShodanQueries** | Pointer to **[]string** | ShodanQueries are queries for examining potential Internet-exposed devices &amp; applications with Shodan in URL form. | [optional] 
**ShodanRawQueries** | Pointer to **[]string** | ShodanRawQueries are raw queries for examining potential Internet-exposed devices &amp; applications with Shodan. | [optional] 
**SnortRule** | Pointer to **bool** | SnortRule indicates whether or not a Snort rule designed to detect the exploitation of the vulnerability over the network exists in this artifact. | [optional] 
**SuricataRule** | Pointer to **bool** | SuricataRule indicates whether or not a Suricata rule designed to detect the exploitation of the vulnerability over the network exists in this artifact. | [optional] 
**TargetDocker** | Pointer to **bool** | TargetDocker indicates whether or not there is an available docker image with the vulnerability. | [optional] 
**TargetEncryptedComms** | Pointer to **string** | Encrypted communications? | [optional] 
**TargetService** | Pointer to **string** | TargetService indicates the service (HTTP, FTP, etc) that this exploit targets. | [optional] 
**Vendor** | Pointer to **string** | Vendor of the vulnerable product | [optional] 
**VersionScanner** | Pointer to **bool** | VersionScanner indicates whether or not the exploit PoC can determine if target system is vulnerable without sending exploit payload in this artifact. | [optional] 
**Yara** | Pointer to **bool** | YARA indicates whether or not a YARA rule designed to detect the exploit on an endpoint exists in this artifact. | [optional] 
**Zeroday** | Pointer to **bool** | Zeroday indicates whether or not it is a VulnCheck zeroday. | [optional] 
**ZoomEyeQueries** | Pointer to **[]string** | ZoomEyeQueries are raw queries for examining potential Internet-exposed devices &amp; applications with ZoomEye. | [optional] 
**ZoomEyeRawQueries** | Pointer to **[]string** |  | [optional] 

## Methods

### NewApiInitialAccessArtifact

`func NewApiInitialAccessArtifact() *ApiInitialAccessArtifact`

NewApiInitialAccessArtifact instantiates a new ApiInitialAccessArtifact object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiInitialAccessArtifactWithDefaults

`func NewApiInitialAccessArtifactWithDefaults() *ApiInitialAccessArtifact`

NewApiInitialAccessArtifactWithDefaults instantiates a new ApiInitialAccessArtifact object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArtifactName

`func (o *ApiInitialAccessArtifact) GetArtifactName() string`

GetArtifactName returns the ArtifactName field if non-nil, zero value otherwise.

### GetArtifactNameOk

`func (o *ApiInitialAccessArtifact) GetArtifactNameOk() (*string, bool)`

GetArtifactNameOk returns a tuple with the ArtifactName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifactName

`func (o *ApiInitialAccessArtifact) SetArtifactName(v string)`

SetArtifactName sets ArtifactName field to given value.

### HasArtifactName

`func (o *ApiInitialAccessArtifact) HasArtifactName() bool`

HasArtifactName returns a boolean if a field has been set.

### GetArtifactsURL

`func (o *ApiInitialAccessArtifact) GetArtifactsURL() []string`

GetArtifactsURL returns the ArtifactsURL field if non-nil, zero value otherwise.

### GetArtifactsURLOk

`func (o *ApiInitialAccessArtifact) GetArtifactsURLOk() (*[]string, bool)`

GetArtifactsURLOk returns a tuple with the ArtifactsURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifactsURL

`func (o *ApiInitialAccessArtifact) SetArtifactsURL(v []string)`

SetArtifactsURL sets ArtifactsURL field to given value.

### HasArtifactsURL

`func (o *ApiInitialAccessArtifact) HasArtifactsURL() bool`

HasArtifactsURL returns a boolean if a field has been set.

### GetBaiduQueries

`func (o *ApiInitialAccessArtifact) GetBaiduQueries() []string`

GetBaiduQueries returns the BaiduQueries field if non-nil, zero value otherwise.

### GetBaiduQueriesOk

`func (o *ApiInitialAccessArtifact) GetBaiduQueriesOk() (*[]string, bool)`

GetBaiduQueriesOk returns a tuple with the BaiduQueries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaiduQueries

`func (o *ApiInitialAccessArtifact) SetBaiduQueries(v []string)`

SetBaiduQueries sets BaiduQueries field to given value.

### HasBaiduQueries

`func (o *ApiInitialAccessArtifact) HasBaiduQueries() bool`

HasBaiduQueries returns a boolean if a field has been set.

### GetBaiduRawQueries

`func (o *ApiInitialAccessArtifact) GetBaiduRawQueries() []string`

GetBaiduRawQueries returns the BaiduRawQueries field if non-nil, zero value otherwise.

### GetBaiduRawQueriesOk

`func (o *ApiInitialAccessArtifact) GetBaiduRawQueriesOk() (*[]string, bool)`

GetBaiduRawQueriesOk returns a tuple with the BaiduRawQueries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaiduRawQueries

`func (o *ApiInitialAccessArtifact) SetBaiduRawQueries(v []string)`

SetBaiduRawQueries sets BaiduRawQueries field to given value.

### HasBaiduRawQueries

`func (o *ApiInitialAccessArtifact) HasBaiduRawQueries() bool`

HasBaiduRawQueries returns a boolean if a field has been set.

### GetCensysQueries

`func (o *ApiInitialAccessArtifact) GetCensysQueries() []string`

GetCensysQueries returns the CensysQueries field if non-nil, zero value otherwise.

### GetCensysQueriesOk

`func (o *ApiInitialAccessArtifact) GetCensysQueriesOk() (*[]string, bool)`

GetCensysQueriesOk returns a tuple with the CensysQueries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCensysQueries

`func (o *ApiInitialAccessArtifact) SetCensysQueries(v []string)`

SetCensysQueries sets CensysQueries field to given value.

### HasCensysQueries

`func (o *ApiInitialAccessArtifact) HasCensysQueries() bool`

HasCensysQueries returns a boolean if a field has been set.

### GetCensysRawQueries

`func (o *ApiInitialAccessArtifact) GetCensysRawQueries() []string`

GetCensysRawQueries returns the CensysRawQueries field if non-nil, zero value otherwise.

### GetCensysRawQueriesOk

`func (o *ApiInitialAccessArtifact) GetCensysRawQueriesOk() (*[]string, bool)`

GetCensysRawQueriesOk returns a tuple with the CensysRawQueries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCensysRawQueries

`func (o *ApiInitialAccessArtifact) SetCensysRawQueries(v []string)`

SetCensysRawQueries sets CensysRawQueries field to given value.

### HasCensysRawQueries

`func (o *ApiInitialAccessArtifact) HasCensysRawQueries() bool`

HasCensysRawQueries returns a boolean if a field has been set.

### GetCloneSSHURL

`func (o *ApiInitialAccessArtifact) GetCloneSSHURL() string`

GetCloneSSHURL returns the CloneSSHURL field if non-nil, zero value otherwise.

### GetCloneSSHURLOk

`func (o *ApiInitialAccessArtifact) GetCloneSSHURLOk() (*string, bool)`

GetCloneSSHURLOk returns a tuple with the CloneSSHURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloneSSHURL

`func (o *ApiInitialAccessArtifact) SetCloneSSHURL(v string)`

SetCloneSSHURL sets CloneSSHURL field to given value.

### HasCloneSSHURL

`func (o *ApiInitialAccessArtifact) HasCloneSSHURL() bool`

HasCloneSSHURL returns a boolean if a field has been set.

### GetDateAdded

`func (o *ApiInitialAccessArtifact) GetDateAdded() string`

GetDateAdded returns the DateAdded field if non-nil, zero value otherwise.

### GetDateAddedOk

`func (o *ApiInitialAccessArtifact) GetDateAddedOk() (*string, bool)`

GetDateAddedOk returns a tuple with the DateAdded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateAdded

`func (o *ApiInitialAccessArtifact) SetDateAdded(v string)`

SetDateAdded sets DateAdded field to given value.

### HasDateAdded

`func (o *ApiInitialAccessArtifact) HasDateAdded() bool`

HasDateAdded returns a boolean if a field has been set.

### GetExploit

`func (o *ApiInitialAccessArtifact) GetExploit() bool`

GetExploit returns the Exploit field if non-nil, zero value otherwise.

### GetExploitOk

`func (o *ApiInitialAccessArtifact) GetExploitOk() (*bool, bool)`

GetExploitOk returns a tuple with the Exploit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExploit

`func (o *ApiInitialAccessArtifact) SetExploit(v bool)`

SetExploit sets Exploit field to given value.

### HasExploit

`func (o *ApiInitialAccessArtifact) HasExploit() bool`

HasExploit returns a boolean if a field has been set.

### GetFofaQueries

`func (o *ApiInitialAccessArtifact) GetFofaQueries() []string`

GetFofaQueries returns the FofaQueries field if non-nil, zero value otherwise.

### GetFofaQueriesOk

`func (o *ApiInitialAccessArtifact) GetFofaQueriesOk() (*[]string, bool)`

GetFofaQueriesOk returns a tuple with the FofaQueries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFofaQueries

`func (o *ApiInitialAccessArtifact) SetFofaQueries(v []string)`

SetFofaQueries sets FofaQueries field to given value.

### HasFofaQueries

`func (o *ApiInitialAccessArtifact) HasFofaQueries() bool`

HasFofaQueries returns a boolean if a field has been set.

### GetFofaRawQueries

`func (o *ApiInitialAccessArtifact) GetFofaRawQueries() []string`

GetFofaRawQueries returns the FofaRawQueries field if non-nil, zero value otherwise.

### GetFofaRawQueriesOk

`func (o *ApiInitialAccessArtifact) GetFofaRawQueriesOk() (*[]string, bool)`

GetFofaRawQueriesOk returns a tuple with the FofaRawQueries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFofaRawQueries

`func (o *ApiInitialAccessArtifact) SetFofaRawQueries(v []string)`

SetFofaRawQueries sets FofaRawQueries field to given value.

### HasFofaRawQueries

`func (o *ApiInitialAccessArtifact) HasFofaRawQueries() bool`

HasFofaRawQueries returns a boolean if a field has been set.

### GetGoogleQueries

`func (o *ApiInitialAccessArtifact) GetGoogleQueries() []string`

GetGoogleQueries returns the GoogleQueries field if non-nil, zero value otherwise.

### GetGoogleQueriesOk

`func (o *ApiInitialAccessArtifact) GetGoogleQueriesOk() (*[]string, bool)`

GetGoogleQueriesOk returns a tuple with the GoogleQueries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoogleQueries

`func (o *ApiInitialAccessArtifact) SetGoogleQueries(v []string)`

SetGoogleQueries sets GoogleQueries field to given value.

### HasGoogleQueries

`func (o *ApiInitialAccessArtifact) HasGoogleQueries() bool`

HasGoogleQueries returns a boolean if a field has been set.

### GetGoogleRawQueries

`func (o *ApiInitialAccessArtifact) GetGoogleRawQueries() []string`

GetGoogleRawQueries returns the GoogleRawQueries field if non-nil, zero value otherwise.

### GetGoogleRawQueriesOk

`func (o *ApiInitialAccessArtifact) GetGoogleRawQueriesOk() (*[]string, bool)`

GetGoogleRawQueriesOk returns a tuple with the GoogleRawQueries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoogleRawQueries

`func (o *ApiInitialAccessArtifact) SetGoogleRawQueries(v []string)`

SetGoogleRawQueries sets GoogleRawQueries field to given value.

### HasGoogleRawQueries

`func (o *ApiInitialAccessArtifact) HasGoogleRawQueries() bool`

HasGoogleRawQueries returns a boolean if a field has been set.

### GetGreynoiseQueries

`func (o *ApiInitialAccessArtifact) GetGreynoiseQueries() []string`

GetGreynoiseQueries returns the GreynoiseQueries field if non-nil, zero value otherwise.

### GetGreynoiseQueriesOk

`func (o *ApiInitialAccessArtifact) GetGreynoiseQueriesOk() (*[]string, bool)`

GetGreynoiseQueriesOk returns a tuple with the GreynoiseQueries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGreynoiseQueries

`func (o *ApiInitialAccessArtifact) SetGreynoiseQueries(v []string)`

SetGreynoiseQueries sets GreynoiseQueries field to given value.

### HasGreynoiseQueries

`func (o *ApiInitialAccessArtifact) HasGreynoiseQueries() bool`

HasGreynoiseQueries returns a boolean if a field has been set.

### GetMitreAttackTechniques

`func (o *ApiInitialAccessArtifact) GetMitreAttackTechniques() []string`

GetMitreAttackTechniques returns the MitreAttackTechniques field if non-nil, zero value otherwise.

### GetMitreAttackTechniquesOk

`func (o *ApiInitialAccessArtifact) GetMitreAttackTechniquesOk() (*[]string, bool)`

GetMitreAttackTechniquesOk returns a tuple with the MitreAttackTechniques field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMitreAttackTechniques

`func (o *ApiInitialAccessArtifact) SetMitreAttackTechniques(v []string)`

SetMitreAttackTechniques sets MitreAttackTechniques field to given value.

### HasMitreAttackTechniques

`func (o *ApiInitialAccessArtifact) HasMitreAttackTechniques() bool`

HasMitreAttackTechniques returns a boolean if a field has been set.

### GetNmapScript

`func (o *ApiInitialAccessArtifact) GetNmapScript() bool`

GetNmapScript returns the NmapScript field if non-nil, zero value otherwise.

### GetNmapScriptOk

`func (o *ApiInitialAccessArtifact) GetNmapScriptOk() (*bool, bool)`

GetNmapScriptOk returns a tuple with the NmapScript field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNmapScript

`func (o *ApiInitialAccessArtifact) SetNmapScript(v bool)`

SetNmapScript sets NmapScript field to given value.

### HasNmapScript

`func (o *ApiInitialAccessArtifact) HasNmapScript() bool`

HasNmapScript returns a boolean if a field has been set.

### GetPcap

`func (o *ApiInitialAccessArtifact) GetPcap() bool`

GetPcap returns the Pcap field if non-nil, zero value otherwise.

### GetPcapOk

`func (o *ApiInitialAccessArtifact) GetPcapOk() (*bool, bool)`

GetPcapOk returns a tuple with the Pcap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPcap

`func (o *ApiInitialAccessArtifact) SetPcap(v bool)`

SetPcap sets Pcap field to given value.

### HasPcap

`func (o *ApiInitialAccessArtifact) HasPcap() bool`

HasPcap returns a boolean if a field has been set.

### GetProduct

`func (o *ApiInitialAccessArtifact) GetProduct() []string`

GetProduct returns the Product field if non-nil, zero value otherwise.

### GetProductOk

`func (o *ApiInitialAccessArtifact) GetProductOk() (*[]string, bool)`

GetProductOk returns a tuple with the Product field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduct

`func (o *ApiInitialAccessArtifact) SetProduct(v []string)`

SetProduct sets Product field to given value.

### HasProduct

`func (o *ApiInitialAccessArtifact) HasProduct() bool`

HasProduct returns a boolean if a field has been set.

### GetShodanQueries

`func (o *ApiInitialAccessArtifact) GetShodanQueries() []string`

GetShodanQueries returns the ShodanQueries field if non-nil, zero value otherwise.

### GetShodanQueriesOk

`func (o *ApiInitialAccessArtifact) GetShodanQueriesOk() (*[]string, bool)`

GetShodanQueriesOk returns a tuple with the ShodanQueries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShodanQueries

`func (o *ApiInitialAccessArtifact) SetShodanQueries(v []string)`

SetShodanQueries sets ShodanQueries field to given value.

### HasShodanQueries

`func (o *ApiInitialAccessArtifact) HasShodanQueries() bool`

HasShodanQueries returns a boolean if a field has been set.

### GetShodanRawQueries

`func (o *ApiInitialAccessArtifact) GetShodanRawQueries() []string`

GetShodanRawQueries returns the ShodanRawQueries field if non-nil, zero value otherwise.

### GetShodanRawQueriesOk

`func (o *ApiInitialAccessArtifact) GetShodanRawQueriesOk() (*[]string, bool)`

GetShodanRawQueriesOk returns a tuple with the ShodanRawQueries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShodanRawQueries

`func (o *ApiInitialAccessArtifact) SetShodanRawQueries(v []string)`

SetShodanRawQueries sets ShodanRawQueries field to given value.

### HasShodanRawQueries

`func (o *ApiInitialAccessArtifact) HasShodanRawQueries() bool`

HasShodanRawQueries returns a boolean if a field has been set.

### GetSnortRule

`func (o *ApiInitialAccessArtifact) GetSnortRule() bool`

GetSnortRule returns the SnortRule field if non-nil, zero value otherwise.

### GetSnortRuleOk

`func (o *ApiInitialAccessArtifact) GetSnortRuleOk() (*bool, bool)`

GetSnortRuleOk returns a tuple with the SnortRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnortRule

`func (o *ApiInitialAccessArtifact) SetSnortRule(v bool)`

SetSnortRule sets SnortRule field to given value.

### HasSnortRule

`func (o *ApiInitialAccessArtifact) HasSnortRule() bool`

HasSnortRule returns a boolean if a field has been set.

### GetSuricataRule

`func (o *ApiInitialAccessArtifact) GetSuricataRule() bool`

GetSuricataRule returns the SuricataRule field if non-nil, zero value otherwise.

### GetSuricataRuleOk

`func (o *ApiInitialAccessArtifact) GetSuricataRuleOk() (*bool, bool)`

GetSuricataRuleOk returns a tuple with the SuricataRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuricataRule

`func (o *ApiInitialAccessArtifact) SetSuricataRule(v bool)`

SetSuricataRule sets SuricataRule field to given value.

### HasSuricataRule

`func (o *ApiInitialAccessArtifact) HasSuricataRule() bool`

HasSuricataRule returns a boolean if a field has been set.

### GetTargetDocker

`func (o *ApiInitialAccessArtifact) GetTargetDocker() bool`

GetTargetDocker returns the TargetDocker field if non-nil, zero value otherwise.

### GetTargetDockerOk

`func (o *ApiInitialAccessArtifact) GetTargetDockerOk() (*bool, bool)`

GetTargetDockerOk returns a tuple with the TargetDocker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetDocker

`func (o *ApiInitialAccessArtifact) SetTargetDocker(v bool)`

SetTargetDocker sets TargetDocker field to given value.

### HasTargetDocker

`func (o *ApiInitialAccessArtifact) HasTargetDocker() bool`

HasTargetDocker returns a boolean if a field has been set.

### GetTargetEncryptedComms

`func (o *ApiInitialAccessArtifact) GetTargetEncryptedComms() string`

GetTargetEncryptedComms returns the TargetEncryptedComms field if non-nil, zero value otherwise.

### GetTargetEncryptedCommsOk

`func (o *ApiInitialAccessArtifact) GetTargetEncryptedCommsOk() (*string, bool)`

GetTargetEncryptedCommsOk returns a tuple with the TargetEncryptedComms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetEncryptedComms

`func (o *ApiInitialAccessArtifact) SetTargetEncryptedComms(v string)`

SetTargetEncryptedComms sets TargetEncryptedComms field to given value.

### HasTargetEncryptedComms

`func (o *ApiInitialAccessArtifact) HasTargetEncryptedComms() bool`

HasTargetEncryptedComms returns a boolean if a field has been set.

### GetTargetService

`func (o *ApiInitialAccessArtifact) GetTargetService() string`

GetTargetService returns the TargetService field if non-nil, zero value otherwise.

### GetTargetServiceOk

`func (o *ApiInitialAccessArtifact) GetTargetServiceOk() (*string, bool)`

GetTargetServiceOk returns a tuple with the TargetService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetService

`func (o *ApiInitialAccessArtifact) SetTargetService(v string)`

SetTargetService sets TargetService field to given value.

### HasTargetService

`func (o *ApiInitialAccessArtifact) HasTargetService() bool`

HasTargetService returns a boolean if a field has been set.

### GetVendor

`func (o *ApiInitialAccessArtifact) GetVendor() string`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *ApiInitialAccessArtifact) GetVendorOk() (*string, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *ApiInitialAccessArtifact) SetVendor(v string)`

SetVendor sets Vendor field to given value.

### HasVendor

`func (o *ApiInitialAccessArtifact) HasVendor() bool`

HasVendor returns a boolean if a field has been set.

### GetVersionScanner

`func (o *ApiInitialAccessArtifact) GetVersionScanner() bool`

GetVersionScanner returns the VersionScanner field if non-nil, zero value otherwise.

### GetVersionScannerOk

`func (o *ApiInitialAccessArtifact) GetVersionScannerOk() (*bool, bool)`

GetVersionScannerOk returns a tuple with the VersionScanner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionScanner

`func (o *ApiInitialAccessArtifact) SetVersionScanner(v bool)`

SetVersionScanner sets VersionScanner field to given value.

### HasVersionScanner

`func (o *ApiInitialAccessArtifact) HasVersionScanner() bool`

HasVersionScanner returns a boolean if a field has been set.

### GetYara

`func (o *ApiInitialAccessArtifact) GetYara() bool`

GetYara returns the Yara field if non-nil, zero value otherwise.

### GetYaraOk

`func (o *ApiInitialAccessArtifact) GetYaraOk() (*bool, bool)`

GetYaraOk returns a tuple with the Yara field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYara

`func (o *ApiInitialAccessArtifact) SetYara(v bool)`

SetYara sets Yara field to given value.

### HasYara

`func (o *ApiInitialAccessArtifact) HasYara() bool`

HasYara returns a boolean if a field has been set.

### GetZeroday

`func (o *ApiInitialAccessArtifact) GetZeroday() bool`

GetZeroday returns the Zeroday field if non-nil, zero value otherwise.

### GetZerodayOk

`func (o *ApiInitialAccessArtifact) GetZerodayOk() (*bool, bool)`

GetZerodayOk returns a tuple with the Zeroday field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZeroday

`func (o *ApiInitialAccessArtifact) SetZeroday(v bool)`

SetZeroday sets Zeroday field to given value.

### HasZeroday

`func (o *ApiInitialAccessArtifact) HasZeroday() bool`

HasZeroday returns a boolean if a field has been set.

### GetZoomEyeQueries

`func (o *ApiInitialAccessArtifact) GetZoomEyeQueries() []string`

GetZoomEyeQueries returns the ZoomEyeQueries field if non-nil, zero value otherwise.

### GetZoomEyeQueriesOk

`func (o *ApiInitialAccessArtifact) GetZoomEyeQueriesOk() (*[]string, bool)`

GetZoomEyeQueriesOk returns a tuple with the ZoomEyeQueries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoomEyeQueries

`func (o *ApiInitialAccessArtifact) SetZoomEyeQueries(v []string)`

SetZoomEyeQueries sets ZoomEyeQueries field to given value.

### HasZoomEyeQueries

`func (o *ApiInitialAccessArtifact) HasZoomEyeQueries() bool`

HasZoomEyeQueries returns a boolean if a field has been set.

### GetZoomEyeRawQueries

`func (o *ApiInitialAccessArtifact) GetZoomEyeRawQueries() []string`

GetZoomEyeRawQueries returns the ZoomEyeRawQueries field if non-nil, zero value otherwise.

### GetZoomEyeRawQueriesOk

`func (o *ApiInitialAccessArtifact) GetZoomEyeRawQueriesOk() (*[]string, bool)`

GetZoomEyeRawQueriesOk returns a tuple with the ZoomEyeRawQueries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoomEyeRawQueries

`func (o *ApiInitialAccessArtifact) SetZoomEyeRawQueries(v []string)`

SetZoomEyeRawQueries sets ZoomEyeRawQueries field to given value.

### HasZoomEyeRawQueries

`func (o *ApiInitialAccessArtifact) HasZoomEyeRawQueries() bool`

HasZoomEyeRawQueries returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


