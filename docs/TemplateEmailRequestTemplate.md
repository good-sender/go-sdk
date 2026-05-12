# TemplateEmailRequestTemplate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TemplateId** | **string** | The ID of the template to use | 
**Variables** | Pointer to **map[string]string** | Key-value pairs to populate template variables | [optional] 

## Methods

### NewTemplateEmailRequestTemplate

`func NewTemplateEmailRequestTemplate(templateId string, ) *TemplateEmailRequestTemplate`

NewTemplateEmailRequestTemplate instantiates a new TemplateEmailRequestTemplate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTemplateEmailRequestTemplateWithDefaults

`func NewTemplateEmailRequestTemplateWithDefaults() *TemplateEmailRequestTemplate`

NewTemplateEmailRequestTemplateWithDefaults instantiates a new TemplateEmailRequestTemplate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTemplateId

`func (o *TemplateEmailRequestTemplate) GetTemplateId() string`

GetTemplateId returns the TemplateId field if non-nil, zero value otherwise.

### GetTemplateIdOk

`func (o *TemplateEmailRequestTemplate) GetTemplateIdOk() (*string, bool)`

GetTemplateIdOk returns a tuple with the TemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateId

`func (o *TemplateEmailRequestTemplate) SetTemplateId(v string)`

SetTemplateId sets TemplateId field to given value.


### GetVariables

`func (o *TemplateEmailRequestTemplate) GetVariables() map[string]string`

GetVariables returns the Variables field if non-nil, zero value otherwise.

### GetVariablesOk

`func (o *TemplateEmailRequestTemplate) GetVariablesOk() (*map[string]string, bool)`

GetVariablesOk returns a tuple with the Variables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariables

`func (o *TemplateEmailRequestTemplate) SetVariables(v map[string]string)`

SetVariables sets Variables field to given value.

### HasVariables

`func (o *TemplateEmailRequestTemplate) HasVariables() bool`

HasVariables returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


