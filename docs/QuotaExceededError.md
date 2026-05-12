# QuotaExceededError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | **string** | Machine-readable error code. | 
**Kind** | **string** | Whether the daily or monthly quota was exhausted. | 
**Message** | **string** | Human-readable error message. | 
**Limit** | **int32** | Quota limit that was reached. | 
**Used** | **int32** | Number of emails already used against the quota. | 
**ResetAt** | **time.Time** | Timestamp at which the quota window resets. | 

## Methods

### NewQuotaExceededError

`func NewQuotaExceededError(code string, kind string, message string, limit int32, used int32, resetAt time.Time, ) *QuotaExceededError`

NewQuotaExceededError instantiates a new QuotaExceededError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuotaExceededErrorWithDefaults

`func NewQuotaExceededErrorWithDefaults() *QuotaExceededError`

NewQuotaExceededErrorWithDefaults instantiates a new QuotaExceededError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *QuotaExceededError) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *QuotaExceededError) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *QuotaExceededError) SetCode(v string)`

SetCode sets Code field to given value.


### GetKind

`func (o *QuotaExceededError) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *QuotaExceededError) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *QuotaExceededError) SetKind(v string)`

SetKind sets Kind field to given value.


### GetMessage

`func (o *QuotaExceededError) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *QuotaExceededError) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *QuotaExceededError) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetLimit

`func (o *QuotaExceededError) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *QuotaExceededError) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *QuotaExceededError) SetLimit(v int32)`

SetLimit sets Limit field to given value.


### GetUsed

`func (o *QuotaExceededError) GetUsed() int32`

GetUsed returns the Used field if non-nil, zero value otherwise.

### GetUsedOk

`func (o *QuotaExceededError) GetUsedOk() (*int32, bool)`

GetUsedOk returns a tuple with the Used field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsed

`func (o *QuotaExceededError) SetUsed(v int32)`

SetUsed sets Used field to given value.


### GetResetAt

`func (o *QuotaExceededError) GetResetAt() time.Time`

GetResetAt returns the ResetAt field if non-nil, zero value otherwise.

### GetResetAtOk

`func (o *QuotaExceededError) GetResetAtOk() (*time.Time, bool)`

GetResetAtOk returns a tuple with the ResetAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResetAt

`func (o *QuotaExceededError) SetResetAt(v time.Time)`

SetResetAt sets ResetAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


