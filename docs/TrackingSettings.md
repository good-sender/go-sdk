# TrackingSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Opens** | Pointer to **bool** | Whether to track email opens | [optional] 
**Clicks** | Pointer to **bool** | Whether to track link clicks | [optional] 
**Unsubscribes** | Pointer to **bool** | Whether to track unsubscribes | [optional] 
**UnsubscribeGroupId** | Pointer to **NullableInt64** | Optional unsubscribe group ID. If not specified, uses global unsubscribe list. This setting is ignored if unsubscribes is false. | [optional] 

## Methods

### NewTrackingSettings

`func NewTrackingSettings() *TrackingSettings`

NewTrackingSettings instantiates a new TrackingSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrackingSettingsWithDefaults

`func NewTrackingSettingsWithDefaults() *TrackingSettings`

NewTrackingSettingsWithDefaults instantiates a new TrackingSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOpens

`func (o *TrackingSettings) GetOpens() bool`

GetOpens returns the Opens field if non-nil, zero value otherwise.

### GetOpensOk

`func (o *TrackingSettings) GetOpensOk() (*bool, bool)`

GetOpensOk returns a tuple with the Opens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpens

`func (o *TrackingSettings) SetOpens(v bool)`

SetOpens sets Opens field to given value.

### HasOpens

`func (o *TrackingSettings) HasOpens() bool`

HasOpens returns a boolean if a field has been set.

### GetClicks

`func (o *TrackingSettings) GetClicks() bool`

GetClicks returns the Clicks field if non-nil, zero value otherwise.

### GetClicksOk

`func (o *TrackingSettings) GetClicksOk() (*bool, bool)`

GetClicksOk returns a tuple with the Clicks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClicks

`func (o *TrackingSettings) SetClicks(v bool)`

SetClicks sets Clicks field to given value.

### HasClicks

`func (o *TrackingSettings) HasClicks() bool`

HasClicks returns a boolean if a field has been set.

### GetUnsubscribes

`func (o *TrackingSettings) GetUnsubscribes() bool`

GetUnsubscribes returns the Unsubscribes field if non-nil, zero value otherwise.

### GetUnsubscribesOk

`func (o *TrackingSettings) GetUnsubscribesOk() (*bool, bool)`

GetUnsubscribesOk returns a tuple with the Unsubscribes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnsubscribes

`func (o *TrackingSettings) SetUnsubscribes(v bool)`

SetUnsubscribes sets Unsubscribes field to given value.

### HasUnsubscribes

`func (o *TrackingSettings) HasUnsubscribes() bool`

HasUnsubscribes returns a boolean if a field has been set.

### GetUnsubscribeGroupId

`func (o *TrackingSettings) GetUnsubscribeGroupId() int64`

GetUnsubscribeGroupId returns the UnsubscribeGroupId field if non-nil, zero value otherwise.

### GetUnsubscribeGroupIdOk

`func (o *TrackingSettings) GetUnsubscribeGroupIdOk() (*int64, bool)`

GetUnsubscribeGroupIdOk returns a tuple with the UnsubscribeGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnsubscribeGroupId

`func (o *TrackingSettings) SetUnsubscribeGroupId(v int64)`

SetUnsubscribeGroupId sets UnsubscribeGroupId field to given value.

### HasUnsubscribeGroupId

`func (o *TrackingSettings) HasUnsubscribeGroupId() bool`

HasUnsubscribeGroupId returns a boolean if a field has been set.

### SetUnsubscribeGroupIdNil

`func (o *TrackingSettings) SetUnsubscribeGroupIdNil(b bool)`

 SetUnsubscribeGroupIdNil sets the value for UnsubscribeGroupId to be an explicit nil

### UnsetUnsubscribeGroupId
`func (o *TrackingSettings) UnsetUnsubscribeGroupId()`

UnsetUnsubscribeGroupId ensures that no value is present for UnsubscribeGroupId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


