# ConsentEmailRecipient

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | **string** | Recipient email address. Leading and trailing whitespace are stripped server-side. | 
**Name** | Pointer to **NullableString** | Optional display name. Pass a non-empty string to set or replace the stored name. Passing &#x60;null&#x60;, an empty string, or omitting the field on a re-submitted recipient leaves any previously-stored name unchanged — clearing must be done via the dashboard / authenticated edit flow.  | [optional] 

## Methods

### NewConsentEmailRecipient

`func NewConsentEmailRecipient(email string, ) *ConsentEmailRecipient`

NewConsentEmailRecipient instantiates a new ConsentEmailRecipient object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConsentEmailRecipientWithDefaults

`func NewConsentEmailRecipientWithDefaults() *ConsentEmailRecipient`

NewConsentEmailRecipientWithDefaults instantiates a new ConsentEmailRecipient object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *ConsentEmailRecipient) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *ConsentEmailRecipient) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *ConsentEmailRecipient) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetName

`func (o *ConsentEmailRecipient) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ConsentEmailRecipient) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ConsentEmailRecipient) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ConsentEmailRecipient) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *ConsentEmailRecipient) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ConsentEmailRecipient) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


