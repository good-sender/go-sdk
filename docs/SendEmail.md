# SendEmail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**From** | [**Address**](Address.md) | Sender address (required) | 
**To** | [**[]Address**](Address.md) | To recipients. At least one recipient (to, cc, or bcc) is required. Maximum 1000 recipients per email. | 
**Subject** | **string** | The subject of the email (required) | [default to ""]
**TextContent** | Pointer to **string** | Plain text content | [optional] 
**HtmlContent** | Pointer to **string** | HTML content | [optional] 
**MarkdownContent** | Pointer to **string** | Markdown content. When provided, text_content and html_content are ignored. The raw markdown is used as text_content and rendered to HTML for html_content.  | [optional] 
**TemplateId** | Pointer to **string** | Template ID for templated emails | [optional] 
**TemplateData** | Pointer to **map[string]interface{}** | Data to populate template variables | [optional] 
**Attachments** | Pointer to [**[]Attachment**](Attachment.md) | Email attachments | [optional] 
**Headers** | Pointer to **map[string]string** | Custom email headers | [optional] 
**ReplyTo** | Pointer to [**Address**](Address.md) | Reply-to address | [optional] 
**SendTime** | Pointer to **int64** | Unix timestamp for when to send the email. Must not be more than 72 hours in the future. If 0, sends immediately. | [optional] 
**WebhookData** | Pointer to **map[string]string** | Custom data to include in webhook events. Maximum 10 keys, key length 50 chars, value length 100 chars. | [optional] 
**Tag** | Pointer to **string** | Custom tag for tracking. Maximum 100 characters. | [optional] 
**Tracking** | Pointer to [**TrackingSettings**](TrackingSettings.md) | Email tracking settings | [optional] 

## Methods

### NewSendEmail

`func NewSendEmail(from Address, to []Address, subject string, ) *SendEmail`

NewSendEmail instantiates a new SendEmail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSendEmailWithDefaults

`func NewSendEmailWithDefaults() *SendEmail`

NewSendEmailWithDefaults instantiates a new SendEmail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFrom

`func (o *SendEmail) GetFrom() Address`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *SendEmail) GetFromOk() (*Address, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *SendEmail) SetFrom(v Address)`

SetFrom sets From field to given value.


### GetTo

`func (o *SendEmail) GetTo() []Address`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *SendEmail) GetToOk() (*[]Address, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *SendEmail) SetTo(v []Address)`

SetTo sets To field to given value.


### GetSubject

`func (o *SendEmail) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *SendEmail) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *SendEmail) SetSubject(v string)`

SetSubject sets Subject field to given value.


### GetTextContent

`func (o *SendEmail) GetTextContent() string`

GetTextContent returns the TextContent field if non-nil, zero value otherwise.

### GetTextContentOk

`func (o *SendEmail) GetTextContentOk() (*string, bool)`

GetTextContentOk returns a tuple with the TextContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTextContent

`func (o *SendEmail) SetTextContent(v string)`

SetTextContent sets TextContent field to given value.

### HasTextContent

`func (o *SendEmail) HasTextContent() bool`

HasTextContent returns a boolean if a field has been set.

### GetHtmlContent

`func (o *SendEmail) GetHtmlContent() string`

GetHtmlContent returns the HtmlContent field if non-nil, zero value otherwise.

### GetHtmlContentOk

`func (o *SendEmail) GetHtmlContentOk() (*string, bool)`

GetHtmlContentOk returns a tuple with the HtmlContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtmlContent

`func (o *SendEmail) SetHtmlContent(v string)`

SetHtmlContent sets HtmlContent field to given value.

### HasHtmlContent

`func (o *SendEmail) HasHtmlContent() bool`

HasHtmlContent returns a boolean if a field has been set.

### GetMarkdownContent

`func (o *SendEmail) GetMarkdownContent() string`

GetMarkdownContent returns the MarkdownContent field if non-nil, zero value otherwise.

### GetMarkdownContentOk

`func (o *SendEmail) GetMarkdownContentOk() (*string, bool)`

GetMarkdownContentOk returns a tuple with the MarkdownContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarkdownContent

`func (o *SendEmail) SetMarkdownContent(v string)`

SetMarkdownContent sets MarkdownContent field to given value.

### HasMarkdownContent

`func (o *SendEmail) HasMarkdownContent() bool`

HasMarkdownContent returns a boolean if a field has been set.

### GetTemplateId

`func (o *SendEmail) GetTemplateId() string`

GetTemplateId returns the TemplateId field if non-nil, zero value otherwise.

### GetTemplateIdOk

`func (o *SendEmail) GetTemplateIdOk() (*string, bool)`

GetTemplateIdOk returns a tuple with the TemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateId

`func (o *SendEmail) SetTemplateId(v string)`

SetTemplateId sets TemplateId field to given value.

### HasTemplateId

`func (o *SendEmail) HasTemplateId() bool`

HasTemplateId returns a boolean if a field has been set.

### GetTemplateData

`func (o *SendEmail) GetTemplateData() map[string]interface{}`

GetTemplateData returns the TemplateData field if non-nil, zero value otherwise.

### GetTemplateDataOk

`func (o *SendEmail) GetTemplateDataOk() (*map[string]interface{}, bool)`

GetTemplateDataOk returns a tuple with the TemplateData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateData

`func (o *SendEmail) SetTemplateData(v map[string]interface{})`

SetTemplateData sets TemplateData field to given value.

### HasTemplateData

`func (o *SendEmail) HasTemplateData() bool`

HasTemplateData returns a boolean if a field has been set.

### GetAttachments

`func (o *SendEmail) GetAttachments() []Attachment`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *SendEmail) GetAttachmentsOk() (*[]Attachment, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *SendEmail) SetAttachments(v []Attachment)`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *SendEmail) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.

### GetHeaders

`func (o *SendEmail) GetHeaders() map[string]string`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *SendEmail) GetHeadersOk() (*map[string]string, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *SendEmail) SetHeaders(v map[string]string)`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *SendEmail) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.

### GetReplyTo

`func (o *SendEmail) GetReplyTo() Address`

GetReplyTo returns the ReplyTo field if non-nil, zero value otherwise.

### GetReplyToOk

`func (o *SendEmail) GetReplyToOk() (*Address, bool)`

GetReplyToOk returns a tuple with the ReplyTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplyTo

`func (o *SendEmail) SetReplyTo(v Address)`

SetReplyTo sets ReplyTo field to given value.

### HasReplyTo

`func (o *SendEmail) HasReplyTo() bool`

HasReplyTo returns a boolean if a field has been set.

### GetSendTime

`func (o *SendEmail) GetSendTime() int64`

GetSendTime returns the SendTime field if non-nil, zero value otherwise.

### GetSendTimeOk

`func (o *SendEmail) GetSendTimeOk() (*int64, bool)`

GetSendTimeOk returns a tuple with the SendTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendTime

`func (o *SendEmail) SetSendTime(v int64)`

SetSendTime sets SendTime field to given value.

### HasSendTime

`func (o *SendEmail) HasSendTime() bool`

HasSendTime returns a boolean if a field has been set.

### GetWebhookData

`func (o *SendEmail) GetWebhookData() map[string]string`

GetWebhookData returns the WebhookData field if non-nil, zero value otherwise.

### GetWebhookDataOk

`func (o *SendEmail) GetWebhookDataOk() (*map[string]string, bool)`

GetWebhookDataOk returns a tuple with the WebhookData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookData

`func (o *SendEmail) SetWebhookData(v map[string]string)`

SetWebhookData sets WebhookData field to given value.

### HasWebhookData

`func (o *SendEmail) HasWebhookData() bool`

HasWebhookData returns a boolean if a field has been set.

### GetTag

`func (o *SendEmail) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *SendEmail) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *SendEmail) SetTag(v string)`

SetTag sets Tag field to given value.

### HasTag

`func (o *SendEmail) HasTag() bool`

HasTag returns a boolean if a field has been set.

### GetTracking

`func (o *SendEmail) GetTracking() TrackingSettings`

GetTracking returns the Tracking field if non-nil, zero value otherwise.

### GetTrackingOk

`func (o *SendEmail) GetTrackingOk() (*TrackingSettings, bool)`

GetTrackingOk returns a tuple with the Tracking field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTracking

`func (o *SendEmail) SetTracking(v TrackingSettings)`

SetTracking sets Tracking field to given value.

### HasTracking

`func (o *SendEmail) HasTracking() bool`

HasTracking returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


