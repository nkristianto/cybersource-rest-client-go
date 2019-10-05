// Code generated by go-swagger; DO NOT EDIT.

package tokenization

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"
)

// TokenizeReader is a Reader for the Tokenize structure.
type TokenizeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TokenizeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewTokenizeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewTokenizeDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewTokenizeOK creates a TokenizeOK with default headers values
func NewTokenizeOK() *TokenizeOK {
	return &TokenizeOK{}
}

/*TokenizeOK handles this case with default header values.

Created payment token.
*/
type TokenizeOK struct {
	Payload *TokenizeOKBody
}

func (o *TokenizeOK) Error() string {
	return fmt.Sprintf("[POST /flex/v1/tokens][%d] tokenizeOK  %+v", 200, o.Payload)
}

func (o *TokenizeOK) GetPayload() *TokenizeOKBody {
	return o.Payload
}

func (o *TokenizeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(TokenizeOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTokenizeDefault creates a TokenizeDefault with default headers values
func NewTokenizeDefault(code int) *TokenizeDefault {
	return &TokenizeDefault{
		_statusCode: code,
	}
}

/*TokenizeDefault handles this case with default header values.

Error creating token.
*/
type TokenizeDefault struct {
	_statusCode int

	Payload *TokenizeDefaultBody
}

// Code gets the status code for the tokenize default response
func (o *TokenizeDefault) Code() int {
	return o._statusCode
}

func (o *TokenizeDefault) Error() string {
	return fmt.Sprintf("[POST /flex/v1/tokens][%d] tokenize default  %+v", o._statusCode, o.Payload)
}

func (o *TokenizeDefault) GetPayload() *TokenizeDefaultBody {
	return o.Payload
}

func (o *TokenizeDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(TokenizeDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*TokenizeBody tokenize body
swagger:model TokenizeBody
*/
type TokenizeBody struct {

	// card info
	CardInfo *TokenizeParamsBodyCardInfo `json:"cardInfo,omitempty"`

	// Unique identifier for the generated token. This is obtained from the Generate Key request. See the [Java Script and Java examples](http://apps.cybersource.com/library/documentation/dev_guides/Secure_Acceptance_Flex/Key/html) on how to import the key and encrypt using the imported key.
	// Required: true
	KeyID *string `json:"keyId"`
}

// Validate validates this tokenize body
func (o *TokenizeBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateCardInfo(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateKeyID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *TokenizeBody) validateCardInfo(formats strfmt.Registry) error {

	if swag.IsZero(o.CardInfo) { // not required
		return nil
	}

	if o.CardInfo != nil {
		if err := o.CardInfo.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("tokenizeRequest" + "." + "cardInfo")
			}
			return err
		}
	}

	return nil
}

func (o *TokenizeBody) validateKeyID(formats strfmt.Registry) error {

	if err := validate.Required("tokenizeRequest"+"."+"keyId", "body", o.KeyID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *TokenizeBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *TokenizeBody) UnmarshalBinary(b []byte) error {
	var res TokenizeBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*TokenizeDefaultBody tokenize default body
swagger:model TokenizeDefaultBody
*/
type TokenizeDefaultBody struct {

	// links
	Links *TokenizeDefaultBodyLinks `json:"_links,omitempty"`

	// response status
	ResponseStatus *TokenizeDefaultBodyResponseStatus `json:"responseStatus,omitempty"`
}

// Validate validates this tokenize default body
func (o *TokenizeDefaultBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateResponseStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *TokenizeDefaultBody) validateLinks(formats strfmt.Registry) error {

	if swag.IsZero(o.Links) { // not required
		return nil
	}

	if o.Links != nil {
		if err := o.Links.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("tokenize default" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

func (o *TokenizeDefaultBody) validateResponseStatus(formats strfmt.Registry) error {

	if swag.IsZero(o.ResponseStatus) { // not required
		return nil
	}

	if o.ResponseStatus != nil {
		if err := o.ResponseStatus.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("tokenize default" + "." + "responseStatus")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *TokenizeDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *TokenizeDefaultBody) UnmarshalBinary(b []byte) error {
	var res TokenizeDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*TokenizeDefaultBodyLinks tokenize default body links
swagger:model TokenizeDefaultBodyLinks
*/
type TokenizeDefaultBodyLinks struct {

	// documentation
	Documentation []*TokenizeDefaultBodyLinksDocumentationItems0 `json:"documentation"`

	// next
	Next []*TokenizeDefaultBodyLinksNextItems0 `json:"next"`

	// self
	Self *TokenizeDefaultBodyLinksSelf `json:"self,omitempty"`
}

// Validate validates this tokenize default body links
func (o *TokenizeDefaultBodyLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDocumentation(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateNext(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *TokenizeDefaultBodyLinks) validateDocumentation(formats strfmt.Registry) error {

	if swag.IsZero(o.Documentation) { // not required
		return nil
	}

	for i := 0; i < len(o.Documentation); i++ {
		if swag.IsZero(o.Documentation[i]) { // not required
			continue
		}

		if o.Documentation[i] != nil {
			if err := o.Documentation[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("tokenize default" + "." + "_links" + "." + "documentation" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *TokenizeDefaultBodyLinks) validateNext(formats strfmt.Registry) error {

	if swag.IsZero(o.Next) { // not required
		return nil
	}

	for i := 0; i < len(o.Next); i++ {
		if swag.IsZero(o.Next[i]) { // not required
			continue
		}

		if o.Next[i] != nil {
			if err := o.Next[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("tokenize default" + "." + "_links" + "." + "next" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *TokenizeDefaultBodyLinks) validateSelf(formats strfmt.Registry) error {

	if swag.IsZero(o.Self) { // not required
		return nil
	}

	if o.Self != nil {
		if err := o.Self.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("tokenize default" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *TokenizeDefaultBodyLinks) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *TokenizeDefaultBodyLinks) UnmarshalBinary(b []byte) error {
	var res TokenizeDefaultBodyLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*TokenizeDefaultBodyLinksDocumentationItems0 tokenize default body links documentation items0
swagger:model TokenizeDefaultBodyLinksDocumentationItems0
*/
type TokenizeDefaultBodyLinksDocumentationItems0 struct {

	// URI of the linked resource.
	Href string `json:"href,omitempty"`

	// HTTP method of the linked resource.
	Method string `json:"method,omitempty"`

	// Label of the linked resource.
	Title string `json:"title,omitempty"`
}

// Validate validates this tokenize default body links documentation items0
func (o *TokenizeDefaultBodyLinksDocumentationItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *TokenizeDefaultBodyLinksDocumentationItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *TokenizeDefaultBodyLinksDocumentationItems0) UnmarshalBinary(b []byte) error {
	var res TokenizeDefaultBodyLinksDocumentationItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*TokenizeDefaultBodyLinksNextItems0 tokenize default body links next items0
swagger:model TokenizeDefaultBodyLinksNextItems0
*/
type TokenizeDefaultBodyLinksNextItems0 struct {

	// URI of the linked resource.
	Href string `json:"href,omitempty"`

	// HTTP method of the linked resource.
	Method string `json:"method,omitempty"`

	// Label of the linked resource.
	Title string `json:"title,omitempty"`
}

// Validate validates this tokenize default body links next items0
func (o *TokenizeDefaultBodyLinksNextItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *TokenizeDefaultBodyLinksNextItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *TokenizeDefaultBodyLinksNextItems0) UnmarshalBinary(b []byte) error {
	var res TokenizeDefaultBodyLinksNextItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*TokenizeDefaultBodyLinksSelf tokenize default body links self
swagger:model TokenizeDefaultBodyLinksSelf
*/
type TokenizeDefaultBodyLinksSelf struct {

	// URI of the linked resource.
	Href string `json:"href,omitempty"`

	// HTTP method of the linked resource.
	Method string `json:"method,omitempty"`

	// Label of the linked resource.
	Title string `json:"title,omitempty"`
}

// Validate validates this tokenize default body links self
func (o *TokenizeDefaultBodyLinksSelf) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *TokenizeDefaultBodyLinksSelf) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *TokenizeDefaultBodyLinksSelf) UnmarshalBinary(b []byte) error {
	var res TokenizeDefaultBodyLinksSelf
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*TokenizeDefaultBodyResponseStatus tokenize default body response status
swagger:model TokenizeDefaultBodyResponseStatus
*/
type TokenizeDefaultBodyResponseStatus struct {

	// API correlation ID.
	CorrelationID string `json:"correlationId,omitempty"`

	// details
	Details []*TokenizeDefaultBodyResponseStatusDetailsItems0 `json:"details"`

	// Error Message.
	Message string `json:"message,omitempty"`

	// Error Reason Code.
	Reason string `json:"reason,omitempty"`

	// HTTP Status code.
	Status float64 `json:"status,omitempty"`
}

// Validate validates this tokenize default body response status
func (o *TokenizeDefaultBodyResponseStatus) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *TokenizeDefaultBodyResponseStatus) validateDetails(formats strfmt.Registry) error {

	if swag.IsZero(o.Details) { // not required
		return nil
	}

	for i := 0; i < len(o.Details); i++ {
		if swag.IsZero(o.Details[i]) { // not required
			continue
		}

		if o.Details[i] != nil {
			if err := o.Details[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("tokenize default" + "." + "responseStatus" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *TokenizeDefaultBodyResponseStatus) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *TokenizeDefaultBodyResponseStatus) UnmarshalBinary(b []byte) error {
	var res TokenizeDefaultBodyResponseStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*TokenizeDefaultBodyResponseStatusDetailsItems0 tokenize default body response status details items0
swagger:model TokenizeDefaultBodyResponseStatusDetailsItems0
*/
type TokenizeDefaultBodyResponseStatusDetailsItems0 struct {

	// Field name referred to for validation issues.
	Location string `json:"location,omitempty"`

	// Description or code of any error response.
	Message string `json:"message,omitempty"`
}

// Validate validates this tokenize default body response status details items0
func (o *TokenizeDefaultBodyResponseStatusDetailsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *TokenizeDefaultBodyResponseStatusDetailsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *TokenizeDefaultBodyResponseStatusDetailsItems0) UnmarshalBinary(b []byte) error {
	var res TokenizeDefaultBodyResponseStatusDetailsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*TokenizeOKBody flexV1TokensPost200Response
swagger:model TokenizeOKBody
*/
type TokenizeOKBody struct {

	// The card type.
	CardType string `json:"cardType,omitempty"`

	// discoverable services
	DiscoverableServices map[string]interface{} `json:"discoverableServices,omitempty"`

	// The Key ID.
	KeyID string `json:"keyId,omitempty"`

	// The masked card number displaying the first 6 digits and the last 4 digits.
	MaskedPan string `json:"maskedPan,omitempty"`

	// Flex-generated digital signature. To ensure the values have not been tampered with while passing through the client, verify this server-side using the public key generated from the /keys resource.
	Signature string `json:"signature,omitempty"`

	// Indicates which fields from the response make up the data that is used when verifying the response signature. See the [sample code] (https://github.com/CyberSource/cybersource-flex-samples/blob/master/java/spring-boot/src/main/java/com/cybersource/flex/application/CheckoutController.java) on how to verify the signature.
	SignedFields string `json:"signedFields,omitempty"`

	// The UTC date and time in milliseconds at which the signature was generated.
	Timestamp int64 `json:"timestamp,omitempty"`

	// The generated token. The token replaces card data and is used as the Subscription ID in the CyberSource Simple Order API or SCMP API.
	Token string `json:"token,omitempty"`
}

// Validate validates this tokenize o k body
func (o *TokenizeOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *TokenizeOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *TokenizeOKBody) UnmarshalBinary(b []byte) error {
	var res TokenizeOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*TokenizeParamsBodyCardInfo tokenize params body card info
swagger:model TokenizeParamsBodyCardInfo
*/
type TokenizeParamsBodyCardInfo struct {

	// Two digit expiration month
	CardExpirationMonth string `json:"cardExpirationMonth,omitempty"`

	// Four digit expiration year
	CardExpirationYear string `json:"cardExpirationYear,omitempty"`

	// Encrypted or plain text card number. If the encryption type of “None” was used in the Generate Key request, this value can be set to the plaintext card number/Personal Account Number (PAN). If the encryption type of RsaOaep256 was used in the Generate Key request, this value needs to be the RSA OAEP 256 encrypted card number. The card number should be encrypted on the cardholders’ device. The [WebCrypto API] (https://github.com/CyberSource/cybersource-flex-samples/blob/master/java/spring-boot/src/main/resources/public/flex.js) can be used with the JWK obtained in the Generate Key request.
	// Required: true
	CardNumber *string `json:"cardNumber"`

	// Card Type. This field is required. Refer to the CyberSource Credit Card Services documentation for supported card types.
	// Required: true
	CardType *string `json:"cardType"`
}

// Validate validates this tokenize params body card info
func (o *TokenizeParamsBodyCardInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateCardNumber(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateCardType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *TokenizeParamsBodyCardInfo) validateCardNumber(formats strfmt.Registry) error {

	if err := validate.Required("tokenizeRequest"+"."+"cardInfo"+"."+"cardNumber", "body", o.CardNumber); err != nil {
		return err
	}

	return nil
}

func (o *TokenizeParamsBodyCardInfo) validateCardType(formats strfmt.Registry) error {

	if err := validate.Required("tokenizeRequest"+"."+"cardInfo"+"."+"cardType", "body", o.CardType); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *TokenizeParamsBodyCardInfo) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *TokenizeParamsBodyCardInfo) UnmarshalBinary(b []byte) error {
	var res TokenizeParamsBodyCardInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
