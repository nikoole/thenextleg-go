package thenextleg

import "time"

type ImagineRequest struct {
	Msg             string `json:"msg"`
	Ref             string `json:"ref"`
	WebhookOverride string `json:"webhookOverride"`
	IgnorePrefilter bool   `json:"ignorePrefilter"`
}

type ImagineResponse struct {
	Success   bool      `json:"success"`
	MessageId string    `json:"messageId"`
	Msg       string    `json:"msg"`
	CreatedAt time.Time `json:"createdAt"`
}

type ImagineResult struct {
	AccountId            string    `json:"accountId"`
	CreatedAt            time.Time `json:"createdAt"`
	Buttons              []string  `json:"buttons"`
	Type                 string    `json:"type"`
	ImageUrl             string    `json:"imageUrl"`
	ImageUrls            []string  `json:"imageUrls"`
	ButtonMessageId      string    `json:"buttonMessageId"`
	OriginatingMessageId string    `json:"originatingMessageId"`
	Content              string    `json:"content"`
	Ref                  string    `json:"ref"`
	ResponseAt           time.Time `json:"responseAt"`
}

type ImagineProgress struct {
	Progress int           `json:"progress"`
	Response ImagineResult `json:"response"`
}

func (self *ImagineResult) IsStatusResponse() bool {
	if _, ok := StatusMap[Status(self.Content)]; ok {
		return true
	}
	return false
}

type DescribeRequest struct {
	Url             string `json:"url"`
	Ref             string `json:"ref"`
	WebhookOverride string `json:"webhookOverride"`
	IgnorePrefilter bool   `json:"ignorePrefilter"`
}

type DescribeResponse struct {
	Success   bool      `json:"success"`
	MessageId string    `json:"messageId"`
	CreatedAt time.Time `json:"createdAt"`
}

type DescribeResult struct {
	AccountId            string    `json:"accountId"`
	CreatedAt            time.Time `json:"createdAt"`
	OriginatingMessageId string    `json:"originatingMessageId"`
	Content              []string  `json:"content"`
	Ref                  string    `json:"ref"`
	Type                 string    `json:"type"`
	ResponseAt           time.Time `json:"responseAt"`
}

type DescribeProgress struct {
	Progress int            `json:"progress"`
	Response DescribeResult `json:"response"`
}

type BlendRequest struct {
	Urls            []string `json:"urls"`
	Ref             string   `json:"ref"`
	WebhookOverride string   `json:"webhookOverride"`
}

type BlendResponse struct {
	Success   bool      `json:"success"`
	MessageId string    `json:"messageId"`
	Msg       string    `json:"msg"`
	CreatedAt time.Time `json:"createdAt"`
}

type BlendResult struct {
	CreatedAt            time.Time `json:"createdAt"`
	Buttons              []string  `json:"buttons"`
	Type                 string    `json:"type"`
	ImageUrl             string    `json:"imageUrl"`
	ImageUrls            []string  `json:"imageUrls"`
	ButtonMessageId      string    `json:"buttonMessageId"`
	OriginatingMessageId string    `json:"originatingMessageId"`
	Content              string    `json:"content"`
	Ref                  string    `json:"ref"`
	ResponseAt           time.Time `json:"responseAt"`
}

type BlendProgress struct {
	Progress int         `json:"progress"`
	Response BlendResult `json:"response"`
}

type FaceSwapRequest struct {
	SourceImg string `json:"sourceImg"`
	TargetImg string `json:"targetImg"`
}

type GetImageRequest struct {
	ImgUrl string `json:"imgUrl"`
}

type InfoRequest struct {
	Ref             string `json:"ref"`
	WebhookOverride string `json:"webhookOverride"`
}

type InfoResponse struct {
	Success   bool      `json:"success"`
	MessageId string    `json:"messageId"`
	CreatedAt time.Time `json:"createdAt"`
}

type InfoResult struct {
	CreatedAt            time.Time `json:"createdAt"`
	OriginatingMessageId string    `json:"originatingMessageId"`
	Content              struct {
		FastTimeRemaining string `json:"Fast Time Remaining"`
		JobMode           string `json:"Job Mode"`
	} `json:"content"`
	Type       string    `json:"type"`
	Ref        string    `json:"ref"`
	ResponseAt time.Time `json:"responseAt"`
}

type InfoProgress struct {
	Progress int        `json:"progress"`
	Response InfoResult `json:"response"`
}

type IsThisNaughtyRequest struct {
	Msg string `json:"msg"`
}

type IsThisNaughtyResponse struct {
	IsNaughty bool   `json:"isNaughty"`
	Phrase    string `json:"phrase"`
}

type Status string

const (
	StatusAlreadyRequestedUpscale            Status = "ALREADY_REQUESTED_UPSCALE"
	StatusBotTookTooLongToProcessYourCommand Status = "BOT_TOOK_TOO_LONG_TO_PROCESS_YOUR_COMMAND"
	StatusAppealAccepted                     Status = "APPEAL_ACCEPTED"
	StatusAppealRejected                     Status = "APPEAL_REJECTED"
	StatusBannedPrompt                       Status = "BANNED_PROMPT"
	StatusAppealBlocked                      Status = "BLOCKED"
	StatusButtonNotFound                     Status = "BUTTON_NOT_FOUND"
	StatusFailedToProcessYourCommand         Status = "FAILED_TO_PROCESS_YOUR_COMMAND"
	StatusFailedToRequest                    Status = "FAILED_TO_REQUEST"
	StatusImageBlocked                       Status = "IMAGE_BLOCKED"
	StatusInternalError                      Status = "INTERNAL_ERROR"
	StatusInvalidLink                        Status = "INVALID_LINK"
	StatusInvalidParameter                   Status = "INVALID_PARAMETER"
	StatusJobActionRestricted                Status = "JOB_ACTION_RESTRICTED"
	StatusJobQueued                          Status = "JOB_QUEUED"
	StatusModerationOutage                   Status = "MODERATION_OUTAGE"
	StatusNoFastHours                        Status = "NO_FAST_HOURS"
	StatusPleaseSubscribeToMJInYourDashboard Status = "PLEASE_SUBSCRIBE_TO_MJ_IN_YOUR_DASHBOARD"
	StatusQueueFull                          Status = "QUEUE_FULL"
)

var StatusMap = map[Status]bool{
	StatusAlreadyRequestedUpscale:            true,
	StatusBotTookTooLongToProcessYourCommand: true,
	StatusAppealAccepted:                     true,
	StatusAppealRejected:                     true,
	StatusBannedPrompt:                       true,
	StatusAppealBlocked:                      true,
	StatusButtonNotFound:                     true,
	StatusFailedToProcessYourCommand:         true,
	StatusFailedToRequest:                    true,
	StatusImageBlocked:                       true,
	StatusInternalError:                      true,
	StatusInvalidLink:                        true,
	StatusInvalidParameter:                   true,
	StatusJobActionRestricted:                true,
	StatusJobQueued:                          true,
	StatusModerationOutage:                   true,
	StatusNoFastHours:                        true,
	StatusPleaseSubscribeToMJInYourDashboard: true,
	StatusQueueFull:                          true,
}
