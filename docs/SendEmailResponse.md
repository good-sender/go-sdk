# SendEmailResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sent** | **int32** | Number of emails sent (recipients had consent allowing delivery) | 
**Declined** | **int32** | Number of emails not sent because recipients did not have granted consent to receive emails | 

## Methods

### NewSendEmailResponse

`func NewSendEmailResponse(sent int32, declined int32, ) *SendEmailResponse`

NewSendEmailResponse instantiates a new SendEmailResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSendEmailResponseWithDefaults

`func NewSendEmailResponseWithDefaults() *SendEmailResponse`

NewSendEmailResponseWithDefaults instantiates a new SendEmailResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSent

`func (o *SendEmailResponse) GetSent() int32`

GetSent returns the Sent field if non-nil, zero value otherwise.

### GetSentOk

`func (o *SendEmailResponse) GetSentOk() (*int32, bool)`

GetSentOk returns a tuple with the Sent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSent

`func (o *SendEmailResponse) SetSent(v int32)`

SetSent sets Sent field to given value.


### GetDeclined

`func (o *SendEmailResponse) GetDeclined() int32`

GetDeclined returns the Declined field if non-nil, zero value otherwise.

### GetDeclinedOk

`func (o *SendEmailResponse) GetDeclinedOk() (*int32, bool)`

GetDeclinedOk returns a tuple with the Declined field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeclined

`func (o *SendEmailResponse) SetDeclined(v int32)`

SetDeclined sets Declined field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


