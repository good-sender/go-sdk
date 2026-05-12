# ConsentEmailRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Domain** | **string** | Domain for the email addresses. | 
**RedirectUrl** | Pointer to **string** | URL to which the user will be redirected after providing consent. {email} in the URL will be replaced with the recipient&#39;s email address. | [optional] 
**Emails** | [**[]ConsentEmailEntry**](ConsentEmailEntry.md) | Recipients to request consent from. Each entry may be either a plain email string or a &#x60;{ email, name? }&#x60; object so callers can attach a display name to a specific recipient. Mixing the two forms in one request is allowed.  | 

## Methods

### NewConsentEmailRequest

`func NewConsentEmailRequest(domain string, emails []ConsentEmailEntry, ) *ConsentEmailRequest`

NewConsentEmailRequest instantiates a new ConsentEmailRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConsentEmailRequestWithDefaults

`func NewConsentEmailRequestWithDefaults() *ConsentEmailRequest`

NewConsentEmailRequestWithDefaults instantiates a new ConsentEmailRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomain

`func (o *ConsentEmailRequest) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *ConsentEmailRequest) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *ConsentEmailRequest) SetDomain(v string)`

SetDomain sets Domain field to given value.


### GetRedirectUrl

`func (o *ConsentEmailRequest) GetRedirectUrl() string`

GetRedirectUrl returns the RedirectUrl field if non-nil, zero value otherwise.

### GetRedirectUrlOk

`func (o *ConsentEmailRequest) GetRedirectUrlOk() (*string, bool)`

GetRedirectUrlOk returns a tuple with the RedirectUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUrl

`func (o *ConsentEmailRequest) SetRedirectUrl(v string)`

SetRedirectUrl sets RedirectUrl field to given value.

### HasRedirectUrl

`func (o *ConsentEmailRequest) HasRedirectUrl() bool`

HasRedirectUrl returns a boolean if a field has been set.

### GetEmails

`func (o *ConsentEmailRequest) GetEmails() []ConsentEmailEntry`

GetEmails returns the Emails field if non-nil, zero value otherwise.

### GetEmailsOk

`func (o *ConsentEmailRequest) GetEmailsOk() (*[]ConsentEmailEntry, bool)`

GetEmailsOk returns a tuple with the Emails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmails

`func (o *ConsentEmailRequest) SetEmails(v []ConsentEmailEntry)`

SetEmails sets Emails field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


