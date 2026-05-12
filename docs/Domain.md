# Domain

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Domain** | **string** | The domain name. | 
**Tracking** | **string** | Subdomain used for click tracking. | 
**ReturnPath** | **string** | Subdomain used for the return path. | 
**RequireTls** | **bool** | Whether outbound mail from this domain must be sent over TLS. | 
**Verification** | [**DomainVerification**](DomainVerification.md) |  | 

## Methods

### NewDomain

`func NewDomain(domain string, tracking string, returnPath string, requireTls bool, verification DomainVerification, ) *Domain`

NewDomain instantiates a new Domain object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDomainWithDefaults

`func NewDomainWithDefaults() *Domain`

NewDomainWithDefaults instantiates a new Domain object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomain

`func (o *Domain) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *Domain) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *Domain) SetDomain(v string)`

SetDomain sets Domain field to given value.


### GetTracking

`func (o *Domain) GetTracking() string`

GetTracking returns the Tracking field if non-nil, zero value otherwise.

### GetTrackingOk

`func (o *Domain) GetTrackingOk() (*string, bool)`

GetTrackingOk returns a tuple with the Tracking field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTracking

`func (o *Domain) SetTracking(v string)`

SetTracking sets Tracking field to given value.


### GetReturnPath

`func (o *Domain) GetReturnPath() string`

GetReturnPath returns the ReturnPath field if non-nil, zero value otherwise.

### GetReturnPathOk

`func (o *Domain) GetReturnPathOk() (*string, bool)`

GetReturnPathOk returns a tuple with the ReturnPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnPath

`func (o *Domain) SetReturnPath(v string)`

SetReturnPath sets ReturnPath field to given value.


### GetRequireTls

`func (o *Domain) GetRequireTls() bool`

GetRequireTls returns the RequireTls field if non-nil, zero value otherwise.

### GetRequireTlsOk

`func (o *Domain) GetRequireTlsOk() (*bool, bool)`

GetRequireTlsOk returns a tuple with the RequireTls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireTls

`func (o *Domain) SetRequireTls(v bool)`

SetRequireTls sets RequireTls field to given value.


### GetVerification

`func (o *Domain) GetVerification() DomainVerification`

GetVerification returns the Verification field if non-nil, zero value otherwise.

### GetVerificationOk

`func (o *Domain) GetVerificationOk() (*DomainVerification, bool)`

GetVerificationOk returns a tuple with the Verification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerification

`func (o *Domain) SetVerification(v DomainVerification)`

SetVerification sets Verification field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


