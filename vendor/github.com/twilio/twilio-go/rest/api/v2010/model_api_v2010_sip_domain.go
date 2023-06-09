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

// ApiV2010SipDomain struct for ApiV2010SipDomain
type ApiV2010SipDomain struct {
	// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the SipDomain resource.
	AccountSid *string `json:"account_sid,omitempty"`
	// The API version used to process the call.
	ApiVersion *string `json:"api_version,omitempty"`
	// The types of authentication you have mapped to your domain. Can be: `IP_ACL` and `CREDENTIAL_LIST`. If you have both defined for your domain, both will be returned in a comma delimited string. If `auth_type` is not defined, the domain will not be able to receive any traffic.
	AuthType *string `json:"auth_type,omitempty"`
	// The date and time in GMT that the resource was created specified in [RFC 2822](https://www.ietf.org/rfc/rfc2822.txt) format.
	DateCreated *string `json:"date_created,omitempty"`
	// The date and time in GMT that the resource was last updated specified in [RFC 2822](https://www.ietf.org/rfc/rfc2822.txt) format.
	DateUpdated *string `json:"date_updated,omitempty"`
	// The unique address you reserve on Twilio to which you route your SIP traffic. Domain names can contain letters, digits, and \"-\" and must end with `sip.twilio.com`.
	DomainName *string `json:"domain_name,omitempty"`
	// The string that you assigned to describe the resource.
	FriendlyName *string `json:"friendly_name,omitempty"`
	// The unique string that that we created to identify the SipDomain resource.
	Sid *string `json:"sid,omitempty"`
	// The URI of the resource, relative to `https://api.twilio.com`.
	Uri *string `json:"uri,omitempty"`
	// The HTTP method we use to call `voice_fallback_url`. Can be: `GET` or `POST`.
	VoiceFallbackMethod *string `json:"voice_fallback_method,omitempty"`
	// The URL that we call when an error occurs while retrieving or executing the TwiML requested from `voice_url`.
	VoiceFallbackUrl *string `json:"voice_fallback_url,omitempty"`
	// The HTTP method we use to call `voice_url`. Can be: `GET` or `POST`.
	VoiceMethod *string `json:"voice_method,omitempty"`
	// The HTTP method we use to call `voice_status_callback_url`. Either `GET` or `POST`.
	VoiceStatusCallbackMethod *string `json:"voice_status_callback_method,omitempty"`
	// The URL that we call to pass status parameters (such as call ended) to your application.
	VoiceStatusCallbackUrl *string `json:"voice_status_callback_url,omitempty"`
	// The URL we call using the `voice_method` when the domain receives a call.
	VoiceUrl *string `json:"voice_url,omitempty"`
	// A list of mapping resources associated with the SIP Domain resource identified by their relative URIs.
	SubresourceUris *map[string]interface{} `json:"subresource_uris,omitempty"`
	// Whether to allow SIP Endpoints to register with the domain to receive calls.
	SipRegistration *bool `json:"sip_registration,omitempty"`
	// Whether emergency calling is enabled for the domain. If enabled, allows emergency calls on the domain from phone numbers with validated addresses.
	EmergencyCallingEnabled *bool `json:"emergency_calling_enabled,omitempty"`
	// Whether secure SIP is enabled for the domain. If enabled, TLS will be enforced and SRTP will be negotiated on all incoming calls to this sip domain.
	Secure *bool `json:"secure,omitempty"`
	// The SID of the BYOC Trunk(Bring Your Own Carrier) resource that the Sip Domain will be associated with.
	ByocTrunkSid *string `json:"byoc_trunk_sid,omitempty"`
	// Whether an emergency caller sid is configured for the domain. If present, this phone number will be used as the callback for the emergency call.
	EmergencyCallerSid *string `json:"emergency_caller_sid,omitempty"`
}
