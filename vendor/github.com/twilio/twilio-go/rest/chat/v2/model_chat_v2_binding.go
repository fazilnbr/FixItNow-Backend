/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Chat
 * This is the public Twilio REST API.
 *
 * NOTE: This class is auto generated by OpenAPI Generator.
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

package openapi

import (
	"time"
)

// ChatV2Binding struct for ChatV2Binding
type ChatV2Binding struct {
	// The unique string that we created to identify the Binding resource.
	Sid *string `json:"sid,omitempty"`
	// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Binding resource.
	AccountSid *string `json:"account_sid,omitempty"`
	// The SID of the [Service](https://www.twilio.com/docs/chat/rest/service-resource) the Binding resource is associated with.
	ServiceSid *string `json:"service_sid,omitempty"`
	// The date and time in GMT when the resource was created specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format.
	DateCreated *time.Time `json:"date_created,omitempty"`
	// The date and time in GMT when the resource was last updated specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format.
	DateUpdated *time.Time `json:"date_updated,omitempty"`
	// The unique endpoint identifier for the Binding. The format of this value depends on the `binding_type`.
	Endpoint *string `json:"endpoint,omitempty"`
	// The application-defined string that uniquely identifies the resource's [User](https://www.twilio.com/docs/chat/rest/user-resource) within the [Service](https://www.twilio.com/docs/chat/rest/service-resource). See [access tokens](https://www.twilio.com/docs/chat/create-tokens) for more info.
	Identity *string `json:"identity,omitempty"`
	// The SID of the [Credential](https://www.twilio.com/docs/chat/rest/credential-resource) for the binding. See [push notification configuration](https://www.twilio.com/docs/chat/push-notification-configuration) for more info.
	CredentialSid *string `json:"credential_sid,omitempty"`
	BindingType   *string `json:"binding_type,omitempty"`
	// The [Programmable Chat message types](https://www.twilio.com/docs/chat/push-notification-configuration#push-types) the binding is subscribed to.
	MessageTypes *[]string `json:"message_types,omitempty"`
	// The absolute URL of the Binding resource.
	Url *string `json:"url,omitempty"`
	// The absolute URLs of the Binding's [User](https://www.twilio.com/docs/chat/rest/user-resource).
	Links *map[string]interface{} `json:"links,omitempty"`
}
