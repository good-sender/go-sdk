# DomainVerification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Verified** | **bool** | Overall verification status. True only when every required DNS record is in place. | 
**TrackingVerified** | **bool** | Whether the tracking subdomain CNAME is in place. | 
**ReturnPathVerified** | **bool** | Whether the return-path subdomain CNAME is in place. | 
**Dkim1Verified** | **bool** | Whether the first DKIM record is in place. | 
**Dkim2Verified** | **bool** | Whether the second DKIM record is in place. | 

## Methods

### NewDomainVerification

`func NewDomainVerification(verified bool, trackingVerified bool, returnPathVerified bool, dkim1Verified bool, dkim2Verified bool, ) *DomainVerification`

NewDomainVerification instantiates a new DomainVerification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDomainVerificationWithDefaults

`func NewDomainVerificationWithDefaults() *DomainVerification`

NewDomainVerificationWithDefaults instantiates a new DomainVerification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVerified

`func (o *DomainVerification) GetVerified() bool`

GetVerified returns the Verified field if non-nil, zero value otherwise.

### GetVerifiedOk

`func (o *DomainVerification) GetVerifiedOk() (*bool, bool)`

GetVerifiedOk returns a tuple with the Verified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerified

`func (o *DomainVerification) SetVerified(v bool)`

SetVerified sets Verified field to given value.


### GetTrackingVerified

`func (o *DomainVerification) GetTrackingVerified() bool`

GetTrackingVerified returns the TrackingVerified field if non-nil, zero value otherwise.

### GetTrackingVerifiedOk

`func (o *DomainVerification) GetTrackingVerifiedOk() (*bool, bool)`

GetTrackingVerifiedOk returns a tuple with the TrackingVerified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackingVerified

`func (o *DomainVerification) SetTrackingVerified(v bool)`

SetTrackingVerified sets TrackingVerified field to given value.


### GetReturnPathVerified

`func (o *DomainVerification) GetReturnPathVerified() bool`

GetReturnPathVerified returns the ReturnPathVerified field if non-nil, zero value otherwise.

### GetReturnPathVerifiedOk

`func (o *DomainVerification) GetReturnPathVerifiedOk() (*bool, bool)`

GetReturnPathVerifiedOk returns a tuple with the ReturnPathVerified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnPathVerified

`func (o *DomainVerification) SetReturnPathVerified(v bool)`

SetReturnPathVerified sets ReturnPathVerified field to given value.


### GetDkim1Verified

`func (o *DomainVerification) GetDkim1Verified() bool`

GetDkim1Verified returns the Dkim1Verified field if non-nil, zero value otherwise.

### GetDkim1VerifiedOk

`func (o *DomainVerification) GetDkim1VerifiedOk() (*bool, bool)`

GetDkim1VerifiedOk returns a tuple with the Dkim1Verified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDkim1Verified

`func (o *DomainVerification) SetDkim1Verified(v bool)`

SetDkim1Verified sets Dkim1Verified field to given value.


### GetDkim2Verified

`func (o *DomainVerification) GetDkim2Verified() bool`

GetDkim2Verified returns the Dkim2Verified field if non-nil, zero value otherwise.

### GetDkim2VerifiedOk

`func (o *DomainVerification) GetDkim2VerifiedOk() (*bool, bool)`

GetDkim2VerifiedOk returns a tuple with the Dkim2Verified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDkim2Verified

`func (o *DomainVerification) SetDkim2Verified(v bool)`

SetDkim2Verified sets Dkim2Verified field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


