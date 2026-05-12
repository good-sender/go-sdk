# SendEmailRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Emails** | [**[]SendEmail**](SendEmail.md) | List of emails to send. Cannot be empty. | 

## Methods

### NewSendEmailRequest

`func NewSendEmailRequest(emails []SendEmail, ) *SendEmailRequest`

NewSendEmailRequest instantiates a new SendEmailRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSendEmailRequestWithDefaults

`func NewSendEmailRequestWithDefaults() *SendEmailRequest`

NewSendEmailRequestWithDefaults instantiates a new SendEmailRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmails

`func (o *SendEmailRequest) GetEmails() []SendEmail`

GetEmails returns the Emails field if non-nil, zero value otherwise.

### GetEmailsOk

`func (o *SendEmailRequest) GetEmailsOk() (*[]SendEmail, bool)`

GetEmailsOk returns a tuple with the Emails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmails

`func (o *SendEmailRequest) SetEmails(v []SendEmail)`

SetEmails sets Emails field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


