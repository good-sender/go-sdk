# EmailListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Emails** | [**[]EmailAccount**](EmailAccount.md) |  | 
**NextCursor** | Pointer to **string** | Cursor to retrieve the next page of results. Omitted if there are no more results. | [optional] 

## Methods

### NewEmailListResponse

`func NewEmailListResponse(emails []EmailAccount, ) *EmailListResponse`

NewEmailListResponse instantiates a new EmailListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmailListResponseWithDefaults

`func NewEmailListResponseWithDefaults() *EmailListResponse`

NewEmailListResponseWithDefaults instantiates a new EmailListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmails

`func (o *EmailListResponse) GetEmails() []EmailAccount`

GetEmails returns the Emails field if non-nil, zero value otherwise.

### GetEmailsOk

`func (o *EmailListResponse) GetEmailsOk() (*[]EmailAccount, bool)`

GetEmailsOk returns a tuple with the Emails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmails

`func (o *EmailListResponse) SetEmails(v []EmailAccount)`

SetEmails sets Emails field to given value.


### GetNextCursor

`func (o *EmailListResponse) GetNextCursor() string`

GetNextCursor returns the NextCursor field if non-nil, zero value otherwise.

### GetNextCursorOk

`func (o *EmailListResponse) GetNextCursorOk() (*string, bool)`

GetNextCursorOk returns a tuple with the NextCursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextCursor

`func (o *EmailListResponse) SetNextCursor(v string)`

SetNextCursor sets NextCursor field to given value.

### HasNextCursor

`func (o *EmailListResponse) HasNextCursor() bool`

HasNextCursor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


