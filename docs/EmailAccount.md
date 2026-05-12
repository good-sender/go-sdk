# EmailAccount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | **string** | Recipient email address. | 
**Name** | Pointer to **NullableString** | Optional display name for the recipient. Used in the To header on the consent email when present, and surfaced through the Dashboard. May be null when no name has been provided. | [optional] 
**Domain** | **string** | Domain part of the email address. | 
**ConsentStatus** | **string** | Status of the recipient&#39;s consent for receiving emails. &#39;pending&#39; &#x3D; awaiting consent email send, &#39;requested&#39; &#x3D; consent email dispatched, &#39;failed&#39; &#x3D; consent email delivery failed, &#39;granted&#39; &#x3D; recipient consented to receive emails, &#39;denied&#39; &#x3D; recipient declined to receive emails. | 
**EngagementStatus** | Pointer to **string** | Status of the recipient&#39;s engagement with the emails. | [optional] 

## Methods

### NewEmailAccount

`func NewEmailAccount(email string, domain string, consentStatus string, ) *EmailAccount`

NewEmailAccount instantiates a new EmailAccount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmailAccountWithDefaults

`func NewEmailAccountWithDefaults() *EmailAccount`

NewEmailAccountWithDefaults instantiates a new EmailAccount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *EmailAccount) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *EmailAccount) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *EmailAccount) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetName

`func (o *EmailAccount) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EmailAccount) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EmailAccount) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *EmailAccount) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *EmailAccount) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *EmailAccount) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDomain

`func (o *EmailAccount) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *EmailAccount) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *EmailAccount) SetDomain(v string)`

SetDomain sets Domain field to given value.


### GetConsentStatus

`func (o *EmailAccount) GetConsentStatus() string`

GetConsentStatus returns the ConsentStatus field if non-nil, zero value otherwise.

### GetConsentStatusOk

`func (o *EmailAccount) GetConsentStatusOk() (*string, bool)`

GetConsentStatusOk returns a tuple with the ConsentStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsentStatus

`func (o *EmailAccount) SetConsentStatus(v string)`

SetConsentStatus sets ConsentStatus field to given value.


### GetEngagementStatus

`func (o *EmailAccount) GetEngagementStatus() string`

GetEngagementStatus returns the EngagementStatus field if non-nil, zero value otherwise.

### GetEngagementStatusOk

`func (o *EmailAccount) GetEngagementStatusOk() (*string, bool)`

GetEngagementStatusOk returns a tuple with the EngagementStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngagementStatus

`func (o *EmailAccount) SetEngagementStatus(v string)`

SetEngagementStatus sets EngagementStatus field to given value.

### HasEngagementStatus

`func (o *EmailAccount) HasEngagementStatus() bool`

HasEngagementStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


