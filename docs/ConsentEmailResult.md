# ConsentEmailResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Emails** | [**[]EmailAccount**](EmailAccount.md) |  | 

## Methods

### NewConsentEmailResult

`func NewConsentEmailResult(emails []EmailAccount, ) *ConsentEmailResult`

NewConsentEmailResult instantiates a new ConsentEmailResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConsentEmailResultWithDefaults

`func NewConsentEmailResultWithDefaults() *ConsentEmailResult`

NewConsentEmailResultWithDefaults instantiates a new ConsentEmailResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmails

`func (o *ConsentEmailResult) GetEmails() []EmailAccount`

GetEmails returns the Emails field if non-nil, zero value otherwise.

### GetEmailsOk

`func (o *ConsentEmailResult) GetEmailsOk() (*[]EmailAccount, bool)`

GetEmailsOk returns a tuple with the Emails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmails

`func (o *ConsentEmailResult) SetEmails(v []EmailAccount)`

SetEmails sets Emails field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


