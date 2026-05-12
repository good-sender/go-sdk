# Attachment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FileName** | Pointer to **string** | File name for the attachment | [optional] [default to ""]
**Content** | Pointer to **string** | Base64 encoded content | [optional] 
**ContentType** | **string** | MIME content type (required) | [default to ""]
**InlineId** | Pointer to **string** | Inline attachment ID | [optional] [default to ""]

## Methods

### NewAttachment

`func NewAttachment(contentType string, ) *Attachment`

NewAttachment instantiates a new Attachment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAttachmentWithDefaults

`func NewAttachmentWithDefaults() *Attachment`

NewAttachmentWithDefaults instantiates a new Attachment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFileName

`func (o *Attachment) GetFileName() string`

GetFileName returns the FileName field if non-nil, zero value otherwise.

### GetFileNameOk

`func (o *Attachment) GetFileNameOk() (*string, bool)`

GetFileNameOk returns a tuple with the FileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileName

`func (o *Attachment) SetFileName(v string)`

SetFileName sets FileName field to given value.

### HasFileName

`func (o *Attachment) HasFileName() bool`

HasFileName returns a boolean if a field has been set.

### GetContent

`func (o *Attachment) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *Attachment) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *Attachment) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *Attachment) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetContentType

`func (o *Attachment) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *Attachment) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *Attachment) SetContentType(v string)`

SetContentType sets ContentType field to given value.


### GetInlineId

`func (o *Attachment) GetInlineId() string`

GetInlineId returns the InlineId field if non-nil, zero value otherwise.

### GetInlineIdOk

`func (o *Attachment) GetInlineIdOk() (*string, bool)`

GetInlineIdOk returns a tuple with the InlineId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInlineId

`func (o *Attachment) SetInlineId(v string)`

SetInlineId sets InlineId field to given value.

### HasInlineId

`func (o *Attachment) HasInlineId() bool`

HasInlineId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


