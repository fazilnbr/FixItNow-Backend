/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Api
 * This is the public Twilio REST API.
 *
 * NOTE: This class is auto generated by OpenAPI Generator.
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

package openapi

// ApiV2010OutgoingCallerId struct for ApiV2010OutgoingCallerId
type ApiV2010OutgoingCallerId struct {
	// The unique string that that we created to identify the OutgoingCallerId resource.
	Sid *string `json:"sid,omitempty"`
	// The date and time in GMT that the resource was created specified in [RFC 2822](https://www.ietf.org/rfc/rfc2822.txt) format.
	DateCreated *string `json:"date_created,omitempty"`
	// The date and time in GMT that the resource was last updated specified in [RFC 2822](https://www.ietf.org/rfc/rfc2822.txt) format.
	DateUpdated *string `json:"date_updated,omitempty"`
	// The string that you assigned to describe the resource.
	FriendlyName *string `json:"friendly_name,omitempty"`
	// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the OutgoingCallerId resource.
	AccountSid *string `json:"account_sid,omitempty"`
	// The phone number in [E.164](https://www.twilio.com/docs/glossary/what-e164) format, which consists of a + followed by the country code and subscriber number.
	PhoneNumber *string `json:"phone_number,omitempty"`
	// The URI of the resource, relative to `https://api.twilio.com`.
	Uri *string `json:"uri,omitempty"`
}
