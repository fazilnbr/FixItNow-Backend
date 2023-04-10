/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Conversations
 * This is the public Twilio REST API.
 *
 * NOTE: This class is auto generated by OpenAPI Generator.
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

package openapi

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"

	"github.com/twilio/twilio-go/client"
)

// Optional parameters for the method 'CreateConversationScopedWebhook'
type CreateConversationScopedWebhookParams struct {
	//
	Target *string `json:"Target,omitempty"`
	// The absolute url the webhook request should be sent to.
	ConfigurationUrl *string `json:"Configuration.Url,omitempty"`
	//
	ConfigurationMethod *string `json:"Configuration.Method,omitempty"`
	// The list of events, firing webhook event for this Conversation.
	ConfigurationFilters *[]string `json:"Configuration.Filters,omitempty"`
	// The list of keywords, firing webhook event for this Conversation.
	ConfigurationTriggers *[]string `json:"Configuration.Triggers,omitempty"`
	// The studio flow SID, where the webhook should be sent to.
	ConfigurationFlowSid *string `json:"Configuration.FlowSid,omitempty"`
	// The message index for which and it's successors the webhook will be replayed. Not set by default
	ConfigurationReplayAfter *int `json:"Configuration.ReplayAfter,omitempty"`
}

func (params *CreateConversationScopedWebhookParams) SetTarget(Target string) *CreateConversationScopedWebhookParams {
	params.Target = &Target
	return params
}
func (params *CreateConversationScopedWebhookParams) SetConfigurationUrl(ConfigurationUrl string) *CreateConversationScopedWebhookParams {
	params.ConfigurationUrl = &ConfigurationUrl
	return params
}
func (params *CreateConversationScopedWebhookParams) SetConfigurationMethod(ConfigurationMethod string) *CreateConversationScopedWebhookParams {
	params.ConfigurationMethod = &ConfigurationMethod
	return params
}
func (params *CreateConversationScopedWebhookParams) SetConfigurationFilters(ConfigurationFilters []string) *CreateConversationScopedWebhookParams {
	params.ConfigurationFilters = &ConfigurationFilters
	return params
}
func (params *CreateConversationScopedWebhookParams) SetConfigurationTriggers(ConfigurationTriggers []string) *CreateConversationScopedWebhookParams {
	params.ConfigurationTriggers = &ConfigurationTriggers
	return params
}
func (params *CreateConversationScopedWebhookParams) SetConfigurationFlowSid(ConfigurationFlowSid string) *CreateConversationScopedWebhookParams {
	params.ConfigurationFlowSid = &ConfigurationFlowSid
	return params
}
func (params *CreateConversationScopedWebhookParams) SetConfigurationReplayAfter(ConfigurationReplayAfter int) *CreateConversationScopedWebhookParams {
	params.ConfigurationReplayAfter = &ConfigurationReplayAfter
	return params
}

// Create a new webhook scoped to the conversation
func (c *ApiService) CreateConversationScopedWebhook(ConversationSid string, params *CreateConversationScopedWebhookParams) (*ConversationsV1ConversationScopedWebhook, error) {
	path := "/v1/Conversations/{ConversationSid}/Webhooks"
	path = strings.Replace(path, "{"+"ConversationSid"+"}", ConversationSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.Target != nil {
		data.Set("Target", *params.Target)
	}
	if params != nil && params.ConfigurationUrl != nil {
		data.Set("Configuration.Url", *params.ConfigurationUrl)
	}
	if params != nil && params.ConfigurationMethod != nil {
		data.Set("Configuration.Method", *params.ConfigurationMethod)
	}
	if params != nil && params.ConfigurationFilters != nil {
		for _, item := range *params.ConfigurationFilters {
			data.Add("Configuration.Filters", item)
		}
	}
	if params != nil && params.ConfigurationTriggers != nil {
		for _, item := range *params.ConfigurationTriggers {
			data.Add("Configuration.Triggers", item)
		}
	}
	if params != nil && params.ConfigurationFlowSid != nil {
		data.Set("Configuration.FlowSid", *params.ConfigurationFlowSid)
	}
	if params != nil && params.ConfigurationReplayAfter != nil {
		data.Set("Configuration.ReplayAfter", fmt.Sprint(*params.ConfigurationReplayAfter))
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ConversationsV1ConversationScopedWebhook{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Remove an existing webhook scoped to the conversation
func (c *ApiService) DeleteConversationScopedWebhook(ConversationSid string, Sid string) error {
	path := "/v1/Conversations/{ConversationSid}/Webhooks/{Sid}"
	path = strings.Replace(path, "{"+"ConversationSid"+"}", ConversationSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Delete(c.baseURL+path, data, headers)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return nil
}

// Fetch the configuration of a conversation-scoped webhook
func (c *ApiService) FetchConversationScopedWebhook(ConversationSid string, Sid string) (*ConversationsV1ConversationScopedWebhook, error) {
	path := "/v1/Conversations/{ConversationSid}/Webhooks/{Sid}"
	path = strings.Replace(path, "{"+"ConversationSid"+"}", ConversationSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ConversationsV1ConversationScopedWebhook{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListConversationScopedWebhook'
type ListConversationScopedWebhookParams struct {
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListConversationScopedWebhookParams) SetPageSize(PageSize int) *ListConversationScopedWebhookParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListConversationScopedWebhookParams) SetLimit(Limit int) *ListConversationScopedWebhookParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of ConversationScopedWebhook records from the API. Request is executed immediately.
func (c *ApiService) PageConversationScopedWebhook(ConversationSid string, params *ListConversationScopedWebhookParams, pageToken, pageNumber string) (*ListConversationScopedWebhookResponse, error) {
	path := "/v1/Conversations/{ConversationSid}/Webhooks"

	path = strings.Replace(path, "{"+"ConversationSid"+"}", ConversationSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.PageSize != nil {
		data.Set("PageSize", fmt.Sprint(*params.PageSize))
	}

	if pageToken != "" {
		data.Set("PageToken", pageToken)
	}
	if pageNumber != "" {
		data.Set("Page", pageNumber)
	}

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListConversationScopedWebhookResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists ConversationScopedWebhook records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListConversationScopedWebhook(ConversationSid string, params *ListConversationScopedWebhookParams) ([]ConversationsV1ConversationScopedWebhook, error) {
	response, errors := c.StreamConversationScopedWebhook(ConversationSid, params)

	records := make([]ConversationsV1ConversationScopedWebhook, 0)
	for record := range response {
		records = append(records, record)
	}

	if err := <-errors; err != nil {
		return nil, err
	}

	return records, nil
}

// Streams ConversationScopedWebhook records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamConversationScopedWebhook(ConversationSid string, params *ListConversationScopedWebhookParams) (chan ConversationsV1ConversationScopedWebhook, chan error) {
	if params == nil {
		params = &ListConversationScopedWebhookParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	recordChannel := make(chan ConversationsV1ConversationScopedWebhook, 1)
	errorChannel := make(chan error, 1)

	response, err := c.PageConversationScopedWebhook(ConversationSid, params, "", "")
	if err != nil {
		errorChannel <- err
		close(recordChannel)
		close(errorChannel)
	} else {
		go c.streamConversationScopedWebhook(response, params, recordChannel, errorChannel)
	}

	return recordChannel, errorChannel
}

func (c *ApiService) streamConversationScopedWebhook(response *ListConversationScopedWebhookResponse, params *ListConversationScopedWebhookParams, recordChannel chan ConversationsV1ConversationScopedWebhook, errorChannel chan error) {
	curRecord := 1

	for response != nil {
		responseRecords := response.Webhooks
		for item := range responseRecords {
			recordChannel <- responseRecords[item]
			curRecord += 1
			if params.Limit != nil && *params.Limit < curRecord {
				close(recordChannel)
				close(errorChannel)
				return
			}
		}

		record, err := client.GetNext(c.baseURL, response, c.getNextListConversationScopedWebhookResponse)
		if err != nil {
			errorChannel <- err
			break
		} else if record == nil {
			break
		}

		response = record.(*ListConversationScopedWebhookResponse)
	}

	close(recordChannel)
	close(errorChannel)
}

func (c *ApiService) getNextListConversationScopedWebhookResponse(nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(nextPageUrl, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListConversationScopedWebhookResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}

// Optional parameters for the method 'UpdateConversationScopedWebhook'
type UpdateConversationScopedWebhookParams struct {
	// The absolute url the webhook request should be sent to.
	ConfigurationUrl *string `json:"Configuration.Url,omitempty"`
	//
	ConfigurationMethod *string `json:"Configuration.Method,omitempty"`
	// The list of events, firing webhook event for this Conversation.
	ConfigurationFilters *[]string `json:"Configuration.Filters,omitempty"`
	// The list of keywords, firing webhook event for this Conversation.
	ConfigurationTriggers *[]string `json:"Configuration.Triggers,omitempty"`
	// The studio flow SID, where the webhook should be sent to.
	ConfigurationFlowSid *string `json:"Configuration.FlowSid,omitempty"`
}

func (params *UpdateConversationScopedWebhookParams) SetConfigurationUrl(ConfigurationUrl string) *UpdateConversationScopedWebhookParams {
	params.ConfigurationUrl = &ConfigurationUrl
	return params
}
func (params *UpdateConversationScopedWebhookParams) SetConfigurationMethod(ConfigurationMethod string) *UpdateConversationScopedWebhookParams {
	params.ConfigurationMethod = &ConfigurationMethod
	return params
}
func (params *UpdateConversationScopedWebhookParams) SetConfigurationFilters(ConfigurationFilters []string) *UpdateConversationScopedWebhookParams {
	params.ConfigurationFilters = &ConfigurationFilters
	return params
}
func (params *UpdateConversationScopedWebhookParams) SetConfigurationTriggers(ConfigurationTriggers []string) *UpdateConversationScopedWebhookParams {
	params.ConfigurationTriggers = &ConfigurationTriggers
	return params
}
func (params *UpdateConversationScopedWebhookParams) SetConfigurationFlowSid(ConfigurationFlowSid string) *UpdateConversationScopedWebhookParams {
	params.ConfigurationFlowSid = &ConfigurationFlowSid
	return params
}

// Update an existing conversation-scoped webhook
func (c *ApiService) UpdateConversationScopedWebhook(ConversationSid string, Sid string, params *UpdateConversationScopedWebhookParams) (*ConversationsV1ConversationScopedWebhook, error) {
	path := "/v1/Conversations/{ConversationSid}/Webhooks/{Sid}"
	path = strings.Replace(path, "{"+"ConversationSid"+"}", ConversationSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.ConfigurationUrl != nil {
		data.Set("Configuration.Url", *params.ConfigurationUrl)
	}
	if params != nil && params.ConfigurationMethod != nil {
		data.Set("Configuration.Method", *params.ConfigurationMethod)
	}
	if params != nil && params.ConfigurationFilters != nil {
		for _, item := range *params.ConfigurationFilters {
			data.Add("Configuration.Filters", item)
		}
	}
	if params != nil && params.ConfigurationTriggers != nil {
		for _, item := range *params.ConfigurationTriggers {
			data.Add("Configuration.Triggers", item)
		}
	}
	if params != nil && params.ConfigurationFlowSid != nil {
		data.Set("Configuration.FlowSid", *params.ConfigurationFlowSid)
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ConversationsV1ConversationScopedWebhook{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
