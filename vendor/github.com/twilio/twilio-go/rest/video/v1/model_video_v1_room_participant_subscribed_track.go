/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Video
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

// VideoV1RoomParticipantSubscribedTrack struct for VideoV1RoomParticipantSubscribedTrack
type VideoV1RoomParticipantSubscribedTrack struct {
	// The unique string that we created to identify the RoomParticipantSubscribedTrack resource.
	Sid *string `json:"sid,omitempty"`
	// The SID of the participant that subscribes to the track.
	ParticipantSid *string `json:"participant_sid,omitempty"`
	// The SID of the participant that publishes the track.
	PublisherSid *string `json:"publisher_sid,omitempty"`
	// The SID of the room where the track is published.
	RoomSid *string `json:"room_sid,omitempty"`
	// The track name. Must have no more than 128 characters and be unique among the participant's published tracks.
	Name *string `json:"name,omitempty"`
	// The date and time in GMT when the resource was created specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format.
	DateCreated *time.Time `json:"date_created,omitempty"`
	// The date and time in GMT when the resource was last updated specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format.
	DateUpdated *time.Time `json:"date_updated,omitempty"`
	// Whether the track is enabled.
	Enabled *bool   `json:"enabled,omitempty"`
	Kind    *string `json:"kind,omitempty"`
	// The absolute URL of the resource.
	Url *string `json:"url,omitempty"`
}
