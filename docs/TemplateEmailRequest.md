# TemplateEmailRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**From** | [**Address**](Address.md) | Sender address (required) | 
**To** | [**Address**](Address.md) | Recipient address (required) | 
**Subject** | **string** | The subject of the email (required) | 
**Template** | [**TemplateEmailRequestTemplate**](TemplateEmailRequestTemplate.md) |  | 

## Methods

### NewTemplateEmailRequest

`func NewTemplateEmailRequest(from Address, to Address, subject string, template TemplateEmailRequestTemplate, ) *TemplateEmailRequest`

NewTemplateEmailRequest instantiates a new TemplateEmailRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTemplateEmailRequestWithDefaults

`func NewTemplateEmailRequestWithDefaults() *TemplateEmailRequest`

NewTemplateEmailRequestWithDefaults instantiates a new TemplateEmailRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFrom

`func (o *TemplateEmailRequest) GetFrom() Address`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *TemplateEmailRequest) GetFromOk() (*Address, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *TemplateEmailRequest) SetFrom(v Address)`

SetFrom sets From field to given value.


### GetTo

`func (o *TemplateEmailRequest) GetTo() Address`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *TemplateEmailRequest) GetToOk() (*Address, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *TemplateEmailRequest) SetTo(v Address)`

SetTo sets To field to given value.


### GetSubject

`func (o *TemplateEmailRequest) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *TemplateEmailRequest) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *TemplateEmailRequest) SetSubject(v string)`

SetSubject sets Subject field to given value.


### GetTemplate

`func (o *TemplateEmailRequest) GetTemplate() TemplateEmailRequestTemplate`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *TemplateEmailRequest) GetTemplateOk() (*TemplateEmailRequestTemplate, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *TemplateEmailRequest) SetTemplate(v TemplateEmailRequestTemplate)`

SetTemplate sets Template field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


