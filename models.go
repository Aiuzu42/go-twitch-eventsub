package twitcheventsub

import (
	"encoding/json"
	"time"
)

type Response struct {
	Challenge    string          `json:"challenge"`
	Subscription Subscription    `json:"subscription"`
	Event        json.RawMessage `json:"event"`
	Events       json.RawMessage `json:"events"`
}

type Subscription struct {
	Id        string    `json:"id"`
	Status    string    `json:"status"`
	Type      string    `json:"type"`
	Version   string    `json:"version"`
	Cost      int       `json:"cost"`
	Condition Condition `json:"condition"`
	Transport Transport `json:"transport"`
	CreatedAt time.Time `json:"created_at"`
}

type SubscriptionResponse struct {
	Data         []Subscription `json:"data"`
	Total        int            `json:"total"`
	TotalCost    int            `json:"total_cost"`
	MaxTotalCost int            `json:"max_total_cost"`
	Pagination   Pagination     `json:"pagination"`
}

type SubscriptionRequest struct {
	Type      string    `json:"type"`
	Version   string    `json:"version"`
	Condition Condition `json:"condition"`
	Transport Transport `json:"transport"`
}

type Condition struct {
	BroadcasterUserId     string `json:"broadcaster_user_id"`
	ModeratorUserId       string `json:"moderator_user_id"`
	UserID                string `json:"user_id"`
	FromBroadcasterUserId string `json:"from_broadcaster_user_id"`
	ToBroadcasterUserId   string `json:"to_broadcaster_user_id"`
}

type Transport struct {
	Method   string `json:"method"`
	Callback string `json:"callback"`
	Secret   string `json:"secret"`
}

type Pagination struct {
	Cursor string `json:"cursor"`
}

type ChannelUpdateEvent struct {
	BroadcasterUserID           string   `json:"broadcaster_user_id"`
	BroadcasterUserLogin        string   `json:"broadcaster_user_login"`
	BroadcasterUserName         string   `json:"broadcaster_user_name"`
	Title                       string   `json:"title"`
	Language                    string   `json:"language"`
	CategoryID                  string   `json:"category_id"`
	CategoryName                string   `json:"category_name"`
	ContentClassificationLabels []string `json:"content_classification_labels"`
}

type ChannelFollowEvent struct {
	UserID               string    `json:"user_id"`
	UserLogin            string    `json:"user_login"`
	UserName             string    `json:"user_name"`
	BroadcasterUserID    string    `json:"broadcaster_user_id"`
	BroadcasterUserLogin string    `json:"broadcaster_user_login"`
	BroadcasterUserName  string    `json:"broadcaster_user_name"`
	FollowedAt           time.Time `json:"followed_at"`
}

type ChannelSubscribeEvent struct {
	UserID               string `json:"user_id"`
	UserLogin            string `json:"user_login"`
	UserName             string `json:"user_name"`
	BroadcasterUserID    string `json:"broadcaster_user_id"`
	BroadcasterUserLogin string `json:"broadcaster_user_login"`
	BroadcasterUserName  string `json:"broadcaster_user_name"`
	Tier                 string `json:"tier"`
	IsGift               bool   `json:"is_gift"`
}

type ChannelSubscriptionEndEvent struct {
	UserID               string `json:"user_id"`
	UserLogin            string `json:"user_login"`
	UserName             string `json:"user_name"`
	BroadcasterUserID    string `json:"broadcaster_user_id"`
	BroadcasterUserLogin string `json:"broadcaster_user_login"`
	BroadcasterUserName  string `json:"broadcaster_user_name"`
	Tier                 string `json:"tier"`
	IsGift               bool   `json:"is_gift"`
}

type ChannelSubscriptionGiftEvent struct {
	UserID               *string `json:"user_id"`
	UserLogin            *string `json:"user_login"`
	UserName             *string `json:"user_name"`
	BroadcasterUserID    string  `json:"broadcaster_user_id"`
	BroadcasterUserLogin string  `json:"broadcaster_user_login"`
	BroadcasterUserName  string  `json:"broadcaster_user_name"`
	Total                int     `json:"total"`
	Tier                 string  `json:"tier"`
	CumulativeTotal      *int    `json:"cumulative_total"`
	IsAnonymous          bool    `json:"is_anonymous"`
}

type ChannelSubscriptionMessageEvent struct {
	UserID               string                    `json:"user_id"`
	UserLogin            string                    `json:"user_login"`
	UserName             string                    `json:"user_name"`
	BroadcasterUserID    string                    `json:"broadcaster_user_id"`
	BroadcasterUserLogin string                    `json:"broadcaster_user_login"`
	BroadcasterUserName  string                    `json:"broadcaster_user_name"`
	Tier                 string                    `json:"tier"`
	Message              SubscriptionMessageStruct `json:"message"`
	CumulativeMonths     int                       `json:"cumulative_months"`
	StreakMonths         *int                      `json:"streak_months"`
	DurationMonths       int                       `json:"duration_months"`
}

type SubscriptionMessageStruct struct {
	Text   string   `json:"text"`
	Emotes []Emotes `json:"emotes"`
}

type Emotes struct {
	Begin int    `json:"begin"`
	End   int    `json:"end"`
	ID    string `json:"id"`
}

type ChannelCheerEvent struct {
	IsAnonymous          bool    `json:"is_anonymous"`
	UserID               *string `json:"user_id"`
	UserLogin            *string `json:"user_login"`
	UserName             *string `json:"user_name"`
	BroadcasterUserID    string  `json:"broadcaster_user_id"`
	BroadcasterUserLogin string  `json:"broadcaster_user_login"`
	BroadcasterUserName  string  `json:"broadcaster_user_name"`
	Message              string  `json:"message"`
	Bits                 int     `json:"bits"`
}

type ChannelRaidEvent struct {
	FromBroadcasterUserID    string `json:"from_broadcaster_user_id"`
	FromBroadcasterUserLogin string `json:"from_broadcaster_user_login"`
	FromBroadcasterUserName  string `json:"from_broadcaster_user_name"`
	ToBroadcasterUserID      string `json:"to_broadcaster_user_id"`
	ToBroadcasterUserLogin   string `json:"to_broadcaster_user_login"`
	ToBroadcasterUserName    string `json:"to_broadcaster_user_name"`
	Viewers                  int    `json:"viewers"`
}

type ChannelBanEvent struct {
	UserID               string     `json:"user_id"`
	UserLogin            string     `json:"user_login"`
	UserName             string     `json:"user_name"`
	BroadcasterUserID    string     `json:"broadcaster_user_id"`
	BroadcasterUserLogin string     `json:"broadcaster_user_login"`
	BroadcasterUserName  string     `json:"broadcaster_user_name"`
	ModeratorUserID      string     `json:"moderator_user_id"`
	ModeratorUserLogin   string     `json:"moderator_user_login"`
	ModeratorUserName    string     `json:"moderator_user_name"`
	Reason               string     `json:"reason"`
	BannedAt             time.Time  `json:"banned_at"`
	EndsAt               *time.Time `json:"ends_at"`
	IsPermanent          bool       `json:"is_permanent"`
}

type ChannelUnbanEvent struct {
	UserID               string `json:"user_id"`
	UserLogin            string `json:"user_login"`
	UserName             string `json:"user_name"`
	BroadcasterUserID    string `json:"broadcaster_user_id"`
	BroadcasterUserLogin string `json:"broadcaster_user_login"`
	BroadcasterUserName  string `json:"broadcaster_user_name"`
	ModeratorUserID      string `json:"moderator_user_id"`
	ModeratorUserLogin   string `json:"moderator_user_login"`
	ModeratorUserName    string `json:"moderator_user_name"`
}

type ChannelModeratorAddEvent struct {
	UserID               string `json:"user_id"`
	UserLogin            string `json:"user_login"`
	UserName             string `json:"user_name"`
	BroadcasterUserID    string `json:"broadcaster_user_id"`
	BroadcasterUserLogin string `json:"broadcaster_user_login"`
	BroadcasterUserName  string `json:"broadcaster_user_name"`
}

type ChannelModeratorRemoveEvent struct {
	UserID               string `json:"user_id"`
	UserLogin            string `json:"user_login"`
	UserName             string `json:"user_name"`
	BroadcasterUserID    string `json:"broadcaster_user_id"`
	BroadcasterUserLogin string `json:"broadcaster_user_login"`
	BroadcasterUserName  string `json:"broadcaster_user_name"`
}

type ChannelPointsCustomRewardAddEvent struct {
	ID                                string     `json:"id"`
	BroadcasterUserID                 string     `json:"broadcaster_user_id"`
	BroadcasterUserLogin              string     `json:"broadcaster_user_login"`
	BroadcasterUserName               string     `json:"broadcaster_user_name"`
	IsEnabled                         bool       `json:"is_enabled"`
	IsPaused                          bool       `json:"is_paused"`
	IsInStock                         bool       `json:"is_in_stock"`
	Title                             string     `json:"title"`
	Cost                              int        `json:"cost"`
	Prompt                            string     `json:"prompt"`
	IsUserInputRequired               bool       `json:"is_user_input_required"`
	ShouldRedemptionsSkipRequestQueue bool       `json:"should_redemptions_skip_request_queue"`
	CooldownExpiresAt                 *time.Time `json:"cooldown_expires_at"`
	RedemptionsRedeemedCurrentStream  *int       `json:"redemptions_redeemed_current_stream"`
	MaxPerStream                      MaxPer     `json:"max_per_stream"`
	MaxPerUserPerStream               MaxPer     `json:"max_per_user_per_stream"`
	GlobalCooldown                    Cooldown   `json:"global_cooldown"`
	BackgroundColor                   string     `json:"background_color"`
	Image                             *Image     `json:"image"`
	DefaultImage                      Image      `json:"default_image"`
}

type MaxPer struct {
	IsEnabled bool `json:"is_enabled"`
	Value     int  `json:"value"`
}

type Cooldown struct {
	IsEnabled bool `json:"is_enabled"`
	Seconds   int  `json:"seconds"`
}

type Image struct {
	URL1X string `json:"url_1x"`
	URL2X string `json:"url_2x"`
	URL4X string `json:"url_4x"`
}

type ChannelPointsCustomRewardUpdateEvent struct {
	ID                                string     `json:"id"`
	BroadcasterUserID                 string     `json:"broadcaster_user_id"`
	BroadcasterUserLogin              string     `json:"broadcaster_user_login"`
	BroadcasterUserName               string     `json:"broadcaster_user_name"`
	IsEnabled                         bool       `json:"is_enabled"`
	IsPaused                          bool       `json:"is_paused"`
	IsInStock                         bool       `json:"is_in_stock"`
	Title                             string     `json:"title"`
	Cost                              int        `json:"cost"`
	Prompt                            string     `json:"prompt"`
	IsUserInputRequired               bool       `json:"is_user_input_required"`
	ShouldRedemptionsSkipRequestQueue bool       `json:"should_redemptions_skip_request_queue"`
	CooldownExpiresAt                 *time.Time `json:"cooldown_expires_at"`
	RedemptionsRedeemedCurrentStream  *int       `json:"redemptions_redeemed_current_stream"`
	MaxPerStream                      MaxPer     `json:"max_per_stream"`
	MaxPerUserPerStream               MaxPer     `json:"max_per_user_per_stream"`
	GlobalCooldown                    Cooldown   `json:"global_cooldown"`
	BackgroundColor                   string     `json:"background_color"`
	Image                             *Image     `json:"image"`
	DefaultImage                      Image      `json:"default_image"`
}

type ChannelPointsCustomRewardRemoveEvent struct {
	ID                                string     `json:"id"`
	BroadcasterUserID                 string     `json:"broadcaster_user_id"`
	BroadcasterUserLogin              string     `json:"broadcaster_user_login"`
	BroadcasterUserName               string     `json:"broadcaster_user_name"`
	IsEnabled                         bool       `json:"is_enabled"`
	IsPaused                          bool       `json:"is_paused"`
	IsInStock                         bool       `json:"is_in_stock"`
	Title                             string     `json:"title"`
	Cost                              int        `json:"cost"`
	Prompt                            string     `json:"prompt"`
	IsUserInputRequired               bool       `json:"is_user_input_required"`
	ShouldRedemptionsSkipRequestQueue bool       `json:"should_redemptions_skip_request_queue"`
	CooldownExpiresAt                 *time.Time `json:"cooldown_expires_at"`
	RedemptionsRedeemedCurrentStream  *int       `json:"redemptions_redeemed_current_stream"`
	MaxPerStream                      MaxPer     `json:"max_per_stream"`
	MaxPerUserPerStream               MaxPer     `json:"max_per_user_per_stream"`
	GlobalCooldown                    Cooldown   `json:"global_cooldown"`
	BackgroundColor                   string     `json:"background_color"`
	Image                             *Image     `json:"image"`
	DefaultImage                      Image      `json:"default_image"`
}

type ChannelPointsCustomRewardRedemptionAddEvent struct {
	ID                   string    `json:"id"`
	BroadcasterUserID    string    `json:"broadcaster_user_id"`
	BroadcasterUserLogin string    `json:"broadcaster_user_login"`
	BroadcasterUserName  string    `json:"broadcaster_user_name"`
	UserID               string    `json:"user_id"`
	UserLogin            string    `json:"user_login"`
	UserName             string    `json:"user_name"`
	UserInput            string    `json:"user_input"`
	Status               string    `json:"status"`
	Reward               Reward    `json:"reward"`
	RedeemedAt           time.Time `json:"redeemed_at"`
}

type Reward struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Cost   int    `json:"cost"`
	Prompt string `json:"prompt"`
}

type ChannelPointsCustomRewardRedemptionUpdateEvent struct {
	ID                   string    `json:"id"`
	BroadcasterUserID    string    `json:"broadcaster_user_id"`
	BroadcasterUserLogin string    `json:"broadcaster_user_login"`
	BroadcasterUserName  string    `json:"broadcaster_user_name"`
	UserID               string    `json:"user_id"`
	UserLogin            string    `json:"user_login"`
	UserName             string    `json:"user_name"`
	UserInput            string    `json:"user_input"`
	Status               string    `json:"status"`
	Reward               Reward    `json:"reward"`
	RedeemedAt           time.Time `json:"redeemed_at"`
}

type ChannelPollBeginEvent struct {
	ID                   string              `json:"id"`
	BroadcasterUserID    string              `json:"broadcaster_user_id"`
	BroadcasterUserLogin string              `json:"broadcaster_user_login"`
	BroadcasterUserName  string              `json:"broadcaster_user_name"`
	Title                string              `json:"title"`
	Choices              []Choices           `json:"choices"`
	ChannelPointsVoting  ChannelPointsVoting `json:"channel_points_voting"`
	StartedAt            time.Time           `json:"started_at"`
	EndsAt               time.Time           `json:"ends_at"`
}

type Choices struct {
	ID                 string `json:"id"`
	Title              string `json:"title"`
	ChannelPointsVotes int    `json:"channel_points_votes"`
	Votes              int    `json:"votes"`
}

type ChannelPointsVoting struct {
	IsEnabled     bool `json:"is_enabled"`
	AmountPerVote int  `json:"amount_per_vote"`
}

type ChannelPollProgressEvent struct {
	ID                   string              `json:"id"`
	BroadcasterUserID    string              `json:"broadcaster_user_id"`
	BroadcasterUserLogin string              `json:"broadcaster_user_login"`
	BroadcasterUserName  string              `json:"broadcaster_user_name"`
	Title                string              `json:"title"`
	Choices              []Choices           `json:"choices"`
	ChannelPointsVoting  ChannelPointsVoting `json:"channel_points_voting"`
	StartedAt            time.Time           `json:"started_at"`
	EndsAt               time.Time           `json:"ends_at"`
}

type ChannelPollEndEvent struct {
	ID                   string              `json:"id"`
	BroadcasterUserID    string              `json:"broadcaster_user_id"`
	BroadcasterUserLogin string              `json:"broadcaster_user_login"`
	BroadcasterUserName  string              `json:"broadcaster_user_name"`
	Title                string              `json:"title"`
	Choices              []Choices           `json:"choices"`
	ChannelPointsVoting  ChannelPointsVoting `json:"channel_points_voting"`
	Status               string              `json:"status"`
	StartedAt            time.Time           `json:"started_at"`
	EndedAt              time.Time           `json:"ended_at"`
}

type ChannelPredictionBeginEvent struct {
	ID                   string     `json:"id"`
	BroadcasterUserID    string     `json:"broadcaster_user_id"`
	BroadcasterUserLogin string     `json:"broadcaster_user_login"`
	BroadcasterUserName  string     `json:"broadcaster_user_name"`
	Title                string     `json:"title"`
	Outcomes             []Outcomes `json:"outcomes"`
	StartedAt            time.Time  `json:"started_at"`
	LocksAt              time.Time  `json:"locks_at"`
}

type Outcomes struct {
	ID            string        `json:"id"`
	Title         string        `json:"title"`
	Color         string        `json:"color"`
	Users         int           `json:"users"`
	ChannelPoints int           `json:"channel_points"`
	TopPredictors TopPredictors `json:"top_predictors"`
}

type TopPredictors struct {
	UserName          string `json:"user_name"`
	UserLogin         string `json:"user_login"`
	UserID            string `json:"user_id"`
	ChannelPointsWon  *int   `json:"channel_points_won"`
	ChannelPointsUsed int    `json:"channel_points_used"`
}

type ChannelPredictionProgressEvent struct {
	ID                   string     `json:"id"`
	BroadcasterUserID    string     `json:"broadcaster_user_id"`
	BroadcasterUserLogin string     `json:"broadcaster_user_login"`
	BroadcasterUserName  string     `json:"broadcaster_user_name"`
	Title                string     `json:"title"`
	Outcomes             []Outcomes `json:"outcomes"`
	StartedAt            time.Time  `json:"started_at"`
	LocksAt              time.Time  `json:"locks_at"`
}

type ChannelPredictionLockEvent struct {
	ID                   string     `json:"id"`
	BroadcasterUserID    string     `json:"broadcaster_user_id"`
	BroadcasterUserLogin string     `json:"broadcaster_user_login"`
	BroadcasterUserName  string     `json:"broadcaster_user_name"`
	Title                string     `json:"title"`
	Outcomes             []Outcomes `json:"outcomes"`
	StartedAt            time.Time  `json:"started_at"`
	LockedAt             time.Time  `json:"locked_at"`
}

type ChannelPredictionEndEvent struct {
	ID                   string     `json:"id"`
	BroadcasterUserID    string     `json:"broadcaster_user_id"`
	BroadcasterUserLogin string     `json:"broadcaster_user_login"`
	BroadcasterUserName  string     `json:"broadcaster_user_name"`
	Title                string     `json:"title"`
	WinningOutcomeID     string     `json:"winning_outcome_id"`
	Outcomes             []Outcomes `json:"outcomes"`
	Status               string     `json:"status"`
	StartedAt            time.Time  `json:"started_at"`
	EndedAt              time.Time  `json:"ended_at"`
}

type ChannelHypeTrainBeginEventV2 struct {
	ID                      string                   `json:"id"`
	BroadcasterUserID       string                   `json:"broadcaster_user_id"`
	BroadcasterUserLogin    string                   `json:"broadcaster_user_login"`
	BroadcasterUserName     string                   `json:"broadcaster_user_name"`
	Total                   int                      `json:"total"`
	Progress                int                      `json:"progress"`
	Goal                    int                      `json:"goal"`
	TopContributions        []Contribution           `json:"top_contributions"`
	Level                   int                      `json:"level"`
	AllTimeHighLevel        int                      `json:"all_time_high_level"`
	AllTimeHighTotal        int                      `json:"all_time_high_total"`
	SharedTrainParticipants []SharedTrainParticipant `json:"shared_train_participants"`
	StartedAt               time.Time                `json:"started_at"`
	ExpiresAt               time.Time                `json:"expires_at"`
	Type                    string                   `json:"type"`
	IsSharedTrain           bool                     `json:"is_shared_train"`
}

type Contribution struct {
	UserID    string `json:"user_id"`
	UserLogin string `json:"user_login"`
	UserName  string `json:"user_name"`
	Type      string `json:"type"`
	Total     int    `json:"total"`
}

type SharedTrainParticipant struct {
	BroadcasterUserID    string `json:"broadcaster_user_id"`
	BroadcasterUserLogin string `json:"broadcaster_user_login"`
	BroadcasterUserName  string `json:"broadcaster_user_name"`
}

type ChannelHypeTrainProgressEventV2 struct {
	ID                      string                   `json:"id"`
	BroadcasterUserID       string                   `json:"broadcaster_user_id"`
	BroadcasterUserLogin    string                   `json:"broadcaster_user_login"`
	BroadcasterUserName     string                   `json:"broadcaster_user_name"`
	Total                   int                      `json:"total"`
	Progress                int                      `json:"progress"`
	Goal                    int                      `json:"goal"`
	TopContributions        []Contribution           `json:"top_contributions"`
	Level                   int                      `json:"level"`
	SharedTrainParticipants []SharedTrainParticipant `json:"shared_train_participants"`
	StartedAt               time.Time                `json:"started_at"`
	ExpiresAt               time.Time                `json:"expires_at"`
	Type                    string                   `json:"type"`
	IsSharedTrain           bool                     `json:"is_shared_train"`
}

type ChannelHypeTrainEndEventV2 struct {
	ID                      string                   `json:"id"`
	BroadcasterUserID       string                   `json:"broadcaster_user_id"`
	BroadcasterUserLogin    string                   `json:"broadcaster_user_login"`
	BroadcasterUserName     string                   `json:"broadcaster_user_name"`
	Total                   int                      `json:"total"`
	TopContributions        []Contribution           `json:"top_contributions"`
	Level                   int                      `json:"level"`
	SharedTrainParticipants []SharedTrainParticipant `json:"shared_train_participants"`
	StartedAt               time.Time                `json:"started_at"`
	CooldownEndsAt          time.Time                `json:"cooldown_ends_at"`
	EndedAt                 time.Time                `json:"ended_at"`
	Type                    string                   `json:"type"`
	IsSharedTrain           bool                     `json:"is_shared_train"`
}

type CharityCampaignDonateEvent struct {
	ID                   string `json:"id"`
	CampaignID           string `json:"campaign_id"`
	BroadcasterUserID    string `json:"broadcaster_user_id"`
	BroadcasterUserName  string `json:"broadcaster_user_name"`
	BroadcasterUserLogin string `json:"broadcaster_user_login"`
	UserID               string `json:"user_id"`
	UserLogin            string `json:"user_login"`
	UserName             string `json:"user_name"`
	CharityName          string `json:"charity_name"`
	CharityDescription   string `json:"charity_description"`
	CharityLogo          string `json:"charity_logo"`
	CharityWebsite       string `json:"charity_website"`
	Amount               Amount `json:"amount"`
}

type Amount struct {
	Value         int    `json:"value"`
	DecimalPlaces int    `json:"decimal_places"`
	Currency      string `json:"currency"`
}

type ExtensionBitsTransactionCreateEvent struct {
	ID                   string  `json:"id"`
	ExtensionClientID    string  `json:"extension_client_id"`
	BroadcasterUserID    string  `json:"broadcaster_user_id"`
	BroadcasterUserLogin string  `json:"broadcaster_user_login"`
	BroadcasterUserName  string  `json:"broadcaster_user_name"`
	UserName             string  `json:"user_name"`
	UserLogin            string  `json:"user_login"`
	UserID               string  `json:"user_id"`
	Product              Product `json:"product"`
}

type Product struct {
	Name          string `json:"name"`
	Sku           string `json:"sku"`
	Bits          int    `json:"bits"`
	InDevelopment bool   `json:"in_development"`
}

type ChannelGoalBeginEvent struct {
	ID                   string    `json:"id"`
	BroadcasterUserID    string    `json:"broadcaster_user_id"`
	BroadcasterUserName  string    `json:"broadcaster_user_name"`
	BroadcasterUserLogin string    `json:"broadcaster_user_login"`
	Type                 string    `json:"type"`
	Description          string    `json:"description"`
	CurrentAmount        int       `json:"current_amount"`
	TargetAmount         int       `json:"target_amount"`
	StartedAt            time.Time `json:"started_at"`
}

type ChannelGoalProgressEvent struct {
	ID                   string    `json:"id"`
	BroadcasterUserID    string    `json:"broadcaster_user_id"`
	BroadcasterUserName  string    `json:"broadcaster_user_name"`
	BroadcasterUserLogin string    `json:"broadcaster_user_login"`
	Type                 string    `json:"type"`
	Description          string    `json:"description"`
	CurrentAmount        int       `json:"current_amount"`
	TargetAmount         int       `json:"target_amount"`
	StartedAt            time.Time `json:"started_at"`
}

type ChannelGoalEndEvent struct {
	ID                   string    `json:"id"`
	BroadcasterUserID    string    `json:"broadcaster_user_id"`
	BroadcasterUserName  string    `json:"broadcaster_user_name"`
	BroadcasterUserLogin string    `json:"broadcaster_user_login"`
	Type                 string    `json:"type"`
	Description          string    `json:"description"`
	IsAchieved           bool      `json:"is_achieved"`
	CurrentAmount        int       `json:"current_amount"`
	TargetAmount         int       `json:"target_amount"`
	StartedAt            time.Time `json:"started_at"`
	EndedAt              time.Time `json:"ended_at"`
}

type StreamOnlineEvent struct {
	ID                   string    `json:"id"`
	BroadcasterUserID    string    `json:"broadcaster_user_id"`
	BroadcasterUserLogin string    `json:"broadcaster_user_login"`
	BroadcasterUserName  string    `json:"broadcaster_user_name"`
	Type                 string    `json:"type"`
	StartedAt            time.Time `json:"started_at"`
}

type StreamOfflineEvent struct {
	BroadcasterUserID    string `json:"broadcaster_user_id"`
	BroadcasterUserLogin string `json:"broadcaster_user_login"`
	BroadcasterUserName  string `json:"broadcaster_user_name"`
}

type UserAuthorizationGrantEvent struct {
	ClientID  string `json:"client_id"`
	UserID    string `json:"user_id"`
	UserLogin string `json:"user_login"`
	UserName  string `json:"user_name"`
}

type UserAuthorizationRevokeEvent struct {
	ClientID  string  `json:"client_id"`
	UserID    string  `json:"user_id"`
	UserLogin *string `json:"user_login"`
	UserName  *string `json:"user_name"`
}

type UserUpdateEvent struct {
	UserID        string `json:"user_id"`
	UserLogin     string `json:"user_login"`
	UserName      string `json:"user_name"`
	Email         string `json:"email"`
	EmailVerified bool   `json:"email_verified"`
	Description   string `json:"description"`
}

type AutomodMessageHoldEvent struct {
	BroadcasterUserID    string                    `json:"broadcaster_user_id"`
	BroadcasterUserName  string                    `json:"broadcaster_user_name"`
	BroadcasterUserLogin string                    `json:"broadcaster_user_login"`
	UserID               string                    `json:"user_id"`
	UserName             string                    `json:"user_name"`
	UserLogin            string                    `json:"user_login"`
	MessageID            string                    `json:"message_id"`
	Message              AutomodMessageHoldMessage `json:"message"`
	Level                int                       `json:"level"`
	Category             string                    `json:"category"`
	HeldAt               time.Time                 `json:"held_at"`
}

type AutomodMessageHoldMessage struct {
	Text      string     `json:"text"`
	Fragments []Fragment `json:"fragments"`
}

type Fragment struct {
	Text       string      `json:"text"`
	TextType   string      `json:"type"`
	Emotes     []Emote     `json:"emotes"`
	Cheermotes []Cheermote `json:"cheermotes"`
}

type Cheermote struct {
	Text   string `json:"text"`
	Amount int    `json:"amount"`
	Prefix string `json:"prefix"`
	Tier   int    `json:"tier"`
}

type Emote struct {
	Text  string `json:"text"`
	ID    string `json:"id"`
	SetID string `json:"set-id"`
}

type AutomodMessageUpdateEvent struct {
	BroadcasterUserID    string    `json:"broadcaster_user_id"`
	BroadcasterUserName  string    `json:"broadcaster_user_name"`
	BroadcasterUserLogin string    `json:"broadcaster_user_login"`
	UserID               string    `json:"user_id"`
	UserName             string    `json:"user_name"`
	UserLogin            string    `json:"user_login"`
	ModeratorUserID      string    `json:"moderator_user_id"`
	ModeratorUserLogin   string    `json:"moderator_user_login"`
	ModeratorUserName    string    `json:"moderator_user_name"`
	MessageID            string    `json:"message_id"`
	Message              string    `json:"message"`
	Level                int       `json:"level"`
	Category             string    `json:"category"`
	Status               string    `json:"status"`
	HeldAt               time.Time `json:"held_at"`
	Fragments            Fragment  `json:"fragments"`
}

type AutomodSettingsUpdateEvent struct {
	Data []struct {
		BroadcasterID           string `json:"broadcaster_id"`
		ModeratorID             string `json:"moderator_id"`
		OverallLevel            *int   `json:"overall_level"`
		Disability              int    `json:"disability"`
		Aggression              int    `json:"aggression"`
		SexualitySexOrGender    int    `json:"sexuality_sex_or_gender"`
		Misogyny                int    `json:"misogyny"`
		Bullying                int    `json:"bullying"`
		Swearing                int    `json:"swearing"`
		RaceEthnicityOrReligion int    `json:"race_ethnicity_or_religion"`
		SexBasedTerms           int    `json:"sex_based_terms"`
	} `json:"data"`
}

type AutomodTermsUpdateEvent struct {
	BroadcasterUserID    string   `json:"broadcaster_user_id"`
	BroadcasterUserName  string   `json:"broadcaster_user_name"`
	BroadcasterUserLogin string   `json:"broadcaster_user_login"`
	ModeratorUserID      string   `json:"moderator_user_id"`
	ModeratorUserLogin   string   `json:"moderator_user_login"`
	ModeratorUserName    string   `json:"moderator_user_name"`
	Action               string   `json:"action"`
	FromAutomod          bool     `json:"from_automod"`
	Terms                []string `json:"terms"`
}

type ChannelAdBreakBeginEvent struct {
	DurationSeconds      string    `json:"duration_seconds"`
	StartedAt            time.Time `json:"started_at"`
	IsAutomatic          string    `json:"is_automatic"`
	BroadcasterUserID    string    `json:"broadcaster_user_id"`
	BroadcasterUserLogin string    `json:"broadcaster_user_login"`
	BroadcasterUserName  string    `json:"broadcaster_user_name"`
	RequesterUserID      string    `json:"requester_user_id"`
	RequesterUserLogin   string    `json:"requester_user_login"`
	RequesterUserName    string    `json:"requester_user_name"`
}

type ChannelChatClearEvent struct {
	BroadcasterUserID    string `json:"broadcaster_user_id"`
	BroadcasterUserName  string `json:"broadcaster_user_name"`
	BroadcasterUserLogin string `json:"broadcaster_user_login"`
}

type ChannelChatClearUserMessagesEvent struct {
	BroadcasterUserID    string `json:"broadcaster_user_id"`
	BroadcasterUserName  string `json:"broadcaster_user_name"`
	BroadcasterUserLogin string `json:"broadcaster_user_login"`
	TargetUserID         string `json:"target_user_id"`
	TargetUserName       string `json:"target_user_name"`
	TargetUserLogin      string `json:"target_user_login"`
}

type ChannelChatMessageEvent struct {
	BroadcasterUserID           string        `json:"broadcaster_user_id"`
	BroadcasterUserLogin        string        `json:"broadcaster_user_login"`
	BroadcasterUserName         string        `json:"broadcaster_user_name"`
	ChatterUserID               string        `json:"chatter_user_id"`
	ChatterUserLogin            string        `json:"chatter_user_login"`
	ChatterUserName             string        `json:"chatter_user_name"`
	MessageID                   string        `json:"message_id"`
	Message                     Message       `json:"message"`
	Color                       string        `json:"color"`
	Badges                      []Badges      `json:"badges"`
	MessageType                 string        `json:"message_type"`
	Cheer                       *MessageCheer `json:"cheer"`
	Reply                       *Reply        `json:"reply"`
	ChannelPointsCustomRewardID *string       `json:"channel_points_custom_reward_id"`
}

type Message struct {
	Text      string            `json:"text"`
	Fragments []MessageFragment `json:"fragments"`
}

type MessageFragment struct {
	Type      string          `json:"type"`
	Text      string          `json:"text"`
	Cheermote *Cheermote      `json:"cheermote"`
	Emote     *MessageEmote   `json:"emote"`
	Mention   *MessageMention `json:"mention"`
}

type MessageEmote struct {
	ID         string   `json:"id"`
	EmoteSetID string   `json:"emote_set_id"`
	Format     []string `json:"format"`
}

type MessageMention struct {
	UserID    string `json:"user_id"`
	UserName  string `json:"user_name"`
	UserLogin string `json:"user_login"`
}

type Badges struct {
	SetID string `json:"set_id"`
	ID    string `json:"id"`
	Info  string `json:"info"`
}

type MessageCheer struct {
	Bits int `json:"bits"`
}

type Reply struct {
	ParentMessageID   string `json:"parent_message_id"`
	ParentMessageBody string `json:"parent_message_body"`
	ParentUserID      string `json:"parent_user_id"`
	ParentUserName    string `json:"parent_user_name"`
	ParentUserLogin   string `json:"parent_user_login"`
	ThreadMessageID   string `json:"thread_message_id"`
	ThreadUserID      string `json:"thread_user_id"`
	ThreadUserName    string `json:"thread_user_name"`
	ThreadUserLogin   string `json:"thread_user_login"`
}

type ChannelChatMessageDeleteEvent struct {
	BroadcasterUserID    string `json:"broadcaster_user_id"`
	BroadcasterUserName  string `json:"broadcaster_user_name"`
	BroadcasterUserLogin string `json:"broadcaster_user_login"`
	TargetUserID         string `json:"target_user_id"`
	TargetUserName       string `json:"target_user_name"`
	TargetUserLogin      string `json:"target_user_login"`
	MessageID            string `json:"message_id"`
}

type ChannelChatNotificationEvent struct {
	BroadcasterUserID    string                        `json:"broadcaster_user_id"`
	BroadcasterUserName  string                        `json:"broadcaster_user_name"`
	BroadcasterUserLogin string                        `json:"broadcaster_user_login"`
	ChatterUserID        string                        `json:"chatter_user_id"`
	ChatterUserLogin     string                        `json:"chatter_user_login"`
	ChatterUserName      string                        `json:"chatter_user_name"`
	ChatterIsAnonymous   bool                          `json:"chatter_is_anonymous"`
	Color                string                        `json:"color"`
	Badges               []Badges                      `json:"badges"`
	SystemMessage        string                        `json:"system_message"`
	MessageID            string                        `json:"message_id"`
	Message              NotificationMessage           `json:"message"`
	NoticeType           string                        `json:"notice_type"`
	Sub                  *NotificationSub              `json:"sub"`
	Resub                *NotificationResub            `json:"resub"`
	SubGift              *NotificationSubGift          `json:"sub_gift"`
	CommunitySubGift     *NotificationCommunitySubGift `json:"community_sub_gift"`
	GiftPaidUpgrade      *NotificationGiftPaidUpgrade  `json:"gift_paid_upgrade"`
	PrimePaidUpgrade     *NotificationPrimePaidUpgrade `json:"prime_paid_upgrade"`
	PayItForward         *NotificationPayItForward     `json:"pay_it_forward"`
	Raid                 *NotificationRaid             `json:"raid"`
	Unraid               any                           `json:"unraid"`
	Announcement         *NotificationAnnouncement     `json:"announcement"`
	BitsBadgeTier        *NotificationBitsBadgeTier    `json:"bits_badge_tier"`
	CharityDonation      *NotificationCharityDonation  `json:"charity_donation"`
}

type NotificationCheermote struct {
	Prefix string `json:"prefix"`
	Bits   int    `json:"bits"`
	Tier   int    `json:"tier"`
}

type NotificationEmote struct {
	ID         string   `json:"id"`
	EmoteSetID string   `json:"emote_set_id"`
	OwnerID    string   `json:"owner_id"`
	Format     []string `json:"format"`
}

type NotificationMention struct {
	UserID    string `json:"user_id"`
	UserName  string `json:"user_name"`
	UserLogin string `json:"user_login"`
}

type NotificationFragment struct {
	Type      string                 `json:"type"`
	Text      string                 `json:"text"`
	Cheermote *NotificationCheermote `json:"cheermote"`
	Emote     *NotificationEmote     `json:"emote"`
	Mention   *NotificationMention   `json:"mention"`
}

type NotificationMessage struct {
	Text      string                 `json:"text"`
	Fragments []NotificationFragment `json:"fragments"`
}

type NotificationSub struct {
	SubTier        string `json:"sub_tier"`
	IsPrime        bool   `json:"is_prime"`
	DurationMonths int    `json:"duration_months"`
}

type NotificationResub struct {
	CumulativeMonths  int     `json:"cumulative_months"`
	DurationMonths    int     `json:"duration_months"`
	StreakMonths      int     `json:"streak_months"`
	SubTier           string  `json:"sub_tier"`
	IsPrime           *bool   `json:"is_prime"`
	IsGift            bool    `json:"is_gift"`
	GifterIsAnonymous *bool   `json:"gifter_is_anonymous"`
	GifterUserID      *string `json:"gifter_user_id"`
	GifterUserName    *string `json:"gifter_user_name"`
	GifterUserLogin   *string `json:"gifter_user_login"`
}

type NotificationSubGift struct {
	DurationMonths     int     `json:"duration_months"`
	CumulativeTotal    *int    `json:"cumulative_months"`
	RecipientUserID    string  `json:"recipient_user_id"`
	RecipientUserName  string  `json:"recipient_user_name"`
	RecipientUserLogin string  `json:"recipient_user_login"`
	SubTier            string  `json:"sub_tier"`
	CommunityGiftID    *string `json:"community_gift_id"`
}

type NotificationCommunitySubGift struct {
	ID              string `json:"id"`
	Total           int    `json:"total"`
	SubTier         string `json:"sub_tier"`
	CumulativeTotal *int   `json:"cumulative_total"`
}

type NotificationGiftPaidUpgrade struct {
	GifterIsAnonymous bool    `json:"gifter_is_anonymous"`
	GifterUserID      *string `json:"gifter_user_id"`
	GifterUserName    *string `json:"gifter_user_name"`
	GifterUserLogin   *string `json:"gifter_user_login"`
}

type NotificationPrimePaidUpgrade struct {
	SubTier string `json:"sub_tier"`
}

type NotificationPayItForward struct {
	GifterIsAnonymous bool    `json:"gifter_is_anonymous"`
	GifterUserID      *string `json:"gifter_user_id"`
	GifterUserName    *string `json:"gifter_user_name"`
	GifterUserLogin   *string `json:"gifter_user_login"`
}

type NotificationRaid struct {
	UserID          string `json:"user_id"`
	UserName        string `json:"user_name"`
	UserLogin       string `json:"user_login"`
	ViewerCount     int    `json:"viewer_count"`
	ProfileImageURL string `json:"profile_image_url"`
}

type NotificationAnnouncement struct {
	Color string `json:"color"`
}

type NotificationBitsBadgeTier struct {
	Tier int `json:"tier"`
}

type NotificationCharityDonation struct {
	CharityName string `json:"charity_name"`
	Amount      Amount `json:"amount"`
}

type ChannelChatSettingsUpdateEvent struct {
	BroadcasterUserID           string `json:"broadcaster_user_id"`
	BroadcasterUserLogin        string `json:"broadcaster_user_login"`
	BroadcasterUserName         string `json:"broadcaster_user_name"`
	EmoteMode                   bool   `json:"emote_mode"`
	FollowerMode                bool   `json:"follower_mode"`
	FollowerModeDurationMinutes *int   `json:"follower_mode_duration_minutes"`
	SlowMode                    bool   `json:"slow_mode"`
	SlowModeWaitTimeSeconds     int    `json:"slow_mode_wait_time_seconds"`
	SubscriberMode              bool   `json:"subscriber_mode"`
	UniqueChatMode              bool   `json:"unique_chat_mode"`
}

type ChannelChatUserMessageHoldEvent struct {
	BroadcasterUserID    string  `json:"broadcaster_user_id"`
	BroadcasterUserLogin string  `json:"broadcaster_user_login"`
	BroadcasterUserName  string  `json:"broadcaster_user_name"`
	UserID               string  `json:"user_id"`
	UserLogin            string  `json:"user_login"`
	UserName             string  `json:"user_name"`
	MessageID            string  `json:"message_id"`
	Message              Message `json:"message"`
}

type ChannelChatUserMessageUpdateEvent struct {
	BroadcasterUserID    string  `json:"broadcaster_user_id"`
	BroadcasterUserLogin string  `json:"broadcaster_user_login"`
	BroadcasterUserName  string  `json:"broadcaster_user_name"`
	UserID               string  `json:"user_id"`
	UserLogin            string  `json:"user_login"`
	UserName             string  `json:"user_name"`
	Status               string  `json:"status"`
	MessageID            string  `json:"message_id"`
	Message              Message `json:"message"`
}

type ChannelUnbanRequestCreateEvent struct {
	ID                   string    `json:"id"`
	BroadcasterUserID    string    `json:"broadcaster_user_id"`
	BroadcasterUserLogin string    `json:"broadcaster_user_login"`
	BroadcasterUserName  string    `json:"broadcaster_user_name"`
	UserID               string    `json:"user_id"`
	UserLogin            string    `json:"user_login"`
	UserName             string    `json:"user_name"`
	Text                 string    `json:"text"`
	CreatedAt            time.Time `json:"created_at"`
}

type ChannelUnbanRequestResolveEvent struct {
	ID                   string `json:"id"`
	BroadcasterUserID    string `json:"broadcaster_user_id"`
	BroadcasterUserLogin string `json:"broadcaster_user_login"`
	BroadcasterUserName  string `json:"broadcaster_user_name"`
	ModeratorUserID      string `json:"moderator_user_id"`
	ModeratorUserLogin   string `json:"moderator_user_login"`
	ModeratorUserName    string `json:"moderator_user_name"`
	UserID               string `json:"user_id"`
	UserLogin            string `json:"user_login"`
	UserName             string `json:"user_name"`
	ResolutionText       string `json:"resolution_text"`
	Status               string `json:"status"`
}

type ChannelModerateEvent struct {
	BroadcasterUserID    string                 `json:"broadcaster_user_id"`
	BroadcasterUserLogin string                 `json:"broadcaster_user_login"`
	BroadcasterUserName  string                 `json:"broadcaster_user_name"`
	ModeratorUserID      string                 `json:"moderator_user_id"`
	ModeratorUserLogin   string                 `json:"moderator_user_login"`
	ModeratorUserName    string                 `json:"moderator_user_name"`
	Action               string                 `json:"action"`
	Followers            *FollowDurationMinutes `json:"followers"`
	Slow                 *Slow                  `json:"slow"`
	Vip                  *UserData              `json:"vip"`
	Unvip                *UserData              `json:"unvip"`
	Mod                  *UserData              `json:"mod"`
	Unmod                *UserData              `json:"unmod"`
	Ban                  *EventBan              `json:"ban"`
	Unban                *UserData              `json:"unban"`
	Timeout              *Timeout               `json:"timeout"`
	Untimeout            *UserData              `json:"untimeout"`
	Raid                 *EventRaid             `json:"raid"`
	Unraid               *UserData              `json:"unraid"`
	Delete               *Delete                `json:"delete"`
	AutomodTerms         *AutoModTerms          `json:"automod_terms"`
	UnbanRequest         *UnBanRequest          `json:"unban_request"`
}

type FollowDurationMinutes struct {
	FollowDurationMinutes int `json:"follow_duration_minutes"`
}

type Slow struct {
	WaitTimeSeconds int `json:"wait_time_seconds"`
}

type UserData struct {
	UserID    string `json:"user_id"`
	UserLogin string `json:"user_login"`
	UserName  string `json:"user_name"`
}

type EventBan struct {
	UserID    string  `json:"user_id"`
	UserLogin string  `json:"user_login"`
	UserName  string  `json:"user_name"`
	Reason    *string `json:"reason"`
}

type Timeout struct {
	UserID    string    `json:"user_id"`
	UserLogin string    `json:"user_login"`
	UserName  string    `json:"user_name"`
	Reason    *string   `json:"reason"`
	ExpiresAt time.Time `json:"expires_at"`
}

type EventRaid struct {
	UserID      string `json:"user_id"`
	UserLogin   string `json:"user_login"`
	UserName    string `json:"user_name"`
	ViewerCount int    `json:"viewer_count"`
}

type Delete struct {
	UserID      string `json:"user_id"`
	UserLogin   string `json:"user_login"`
	UserName    string `json:"user_name"`
	MessageID   string `json:"message_id"`
	MessageBody string `json:"message_body"`
}

type AutoModTerms struct {
	Action      string   `json:"action"`
	List        string   `json:"list"`
	Terms       []string `json:"terms"`
	FromAutomod bool     `json:"from_automod"`
}

type UnBanRequest struct {
	IsApproved       bool   `json:"is_approved"`
	UserID           string `json:"user_id"`
	UserLogin        string `json:"user_login"`
	UserName         string `json:"user_name"`
	ModeratorMessage string `json:"moderator_message"`
}

type ChannelGuestStarSessionBeginEvent struct {
	BroadcasterUserID    string    `json:"broadcaster_user_id"`
	BroadcasterUserName  string    `json:"broadcaster_user_name"`
	BroadcasterUserLogin string    `json:"broadcaster_user_login"`
	ModeratorUserID      string    `json:"moderator_user_id"`
	ModeratorUserName    string    `json:"moderator_user_name"`
	ModeratorUserLogin   string    `json:"moderator_user_login"`
	SessionID            string    `json:"session_id"`
	StartedAt            time.Time `json:"started_at"`
}

type ChannelGuestStarSessionEndEvent struct {
	BroadcasterUserID    string    `json:"broadcaster_user_id"`
	BroadcasterUserName  string    `json:"broadcaster_user_name"`
	BroadcasterUserLogin string    `json:"broadcaster_user_login"`
	ModeratorUserID      string    `json:"moderator_user_id"`
	ModeratorUserName    string    `json:"moderator_user_name"`
	ModeratorUserLogin   string    `json:"moderator_user_login"`
	SessionID            string    `json:"session_id"`
	StartedAt            time.Time `json:"started_at"`
	EndedAt              time.Time `json:"ended_at"`
}

type ChannelGuestStarGuestUpdateEvent struct {
	BroadcasterUserID    string  `json:"broadcaster_user_id"`
	BroadcasterUserName  string  `json:"broadcaster_user_name"`
	BroadcasterUserLogin string  `json:"broadcaster_user_login"`
	SessionID            string  `json:"session_id"`
	ModeratorUserID      *string `json:"moderator_user_id"`
	ModeratorUserName    *string `json:"moderator_user_name"`
	ModeratorUserLogin   *string `json:"moderator_user_login"`
	GuestUserID          *string `json:"guest_user_id"`
	GuestUserName        *string `json:"guest_user_name"`
	GuestUserLogin       *string `json:"guest_user_login"`
	SlotID               *string `json:"slot_id"`
	State                *string `json:"state"`
	HostVideoEnabled     *bool   `json:"host_video_enabled"`
	HostAudioEnabled     *bool   `json:"host_audio_enabled"`
	HostVolume           *int    `json:"host_volume"`
	HostUserID           string  `json:"host_user_id"`
	HostUserName         string  `json:"host_user_name"`
	HostUserLogin        string  `json:"host_user_login"`
}

type ChannelGuestStarSettingsUpdateEvent struct {
	BroadcasterUserID           string `json:"broadcaster_user_id"`
	BroadcasterUserName         string `json:"broadcaster_user_name"`
	BroadcasterUserLogin        string `json:"broadcaster_user_login"`
	IsModeratorSendLiveEnabled  bool   `json:"is_moderator_send_live_enabled"`
	SlotCount                   int    `json:"slot_count"`
	IsBrowserSourceAudioEnabled bool   `json:"is_browser_source_audio_enabled"`
	GroupLayout                 string `json:"group_layout"`
}

type ChannelPointsAutomaticRewardRedemptionAddEvent struct {
	BroadcasterUserID    string                     `json:"broadcaster_user_id"`
	BroadcasterUserName  string                     `json:"broadcaster_user_name"`
	BroadcasterUserLogin string                     `json:"broadcaster_user_login"`
	UserID               string                     `json:"user_id"`
	UserName             string                     `json:"user_name"`
	UserLogin            string                     `json:"user_login"`
	ID                   string                     `json:"id"`
	Reward               ChannelPointsReward        `json:"reward"`
	Message              ChannelPointsRewardMessage `json:"message"`
	UserInput            string                     `json:"user_input"`
	RedeemedAt           time.Time                  `json:"redeemed_at"`
}

type ChannelPointsRewardMessage struct {
	Text   string                     `json:"text"`
	Emotes []ChannelPointsRewardEmote `json:"emotes"`
}

type ChannelPointsRewardEmote struct {
	ID    string `json:"id"`
	Begin int    `json:"begin"`
	End   int    `json:"end"`
}

type ChannelPointsReward struct {
	Type          string               `json:"type"`
	Cost          int                  `json:"cost"`
	UnlockedEmote *RewardUnlockedEmote `json:"unlocked_emote"`
}

type RewardUnlockedEmote struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type ChannelVIPAddEvent struct {
	UserID               string `json:"user_id"`
	UserLogin            string `json:"user_login"`
	UserName             string `json:"user_name"`
	BroadcasterUserID    string `json:"broadcaster_user_id"`
	BroadcasterUserLogin string `json:"broadcaster_user_login"`
	BroadcasterUserName  string `json:"broadcaster_user_name"`
}

type ChannelVIPRemoveEvent struct {
	UserID               string `json:"user_id"`
	UserLogin            string `json:"user_login"`
	UserName             string `json:"user_name"`
	BroadcasterUserID    string `json:"broadcaster_user_id"`
	BroadcasterUserLogin string `json:"broadcaster_user_login"`
	BroadcasterUserName  string `json:"broadcaster_user_name"`
}

type CharityCampaignStartEvent struct {
	ID                 string    `json:"id"`
	BroadcasterID      string    `json:"broadcaster_id"`
	BroadcasterName    string    `json:"broadcaster_name"`
	BroadcasterLogin   string    `json:"broadcaster_login"`
	CharityName        string    `json:"charity_name"`
	CharityDescription string    `json:"charity_description"`
	CharityLogo        string    `json:"charity_logo"`
	CharityWebsite     string    `json:"charity_website"`
	CurrentAmount      Amount    `json:"current_amount"`
	TargetAmount       Amount    `json:"target_amount"`
	StartedAt          time.Time `json:"started_at"`
}

type CharityCampaignProgressEvent struct {
	ID                 string `json:"id"`
	BroadcasterID      string `json:"broadcaster_id"`
	BroadcasterName    string `json:"broadcaster_name"`
	BroadcasterLogin   string `json:"broadcaster_login"`
	CharityName        string `json:"charity_name"`
	CharityDescription string `json:"charity_description"`
	CharityLogo        string `json:"charity_logo"`
	CharityWebsite     string `json:"charity_website"`
	CurrentAmount      Amount `json:"current_amount"`
	TargetAmount       Amount `json:"target_amount"`
}

type CharityCampaignStopEvent struct {
	ID                 string    `json:"id"`
	BroadcasterID      string    `json:"broadcaster_id"`
	BroadcasterName    string    `json:"broadcaster_name"`
	BroadcasterLogin   string    `json:"broadcaster_login"`
	CharityName        string    `json:"charity_name"`
	CharityDescription string    `json:"charity_description"`
	CharityLogo        string    `json:"charity_logo"`
	CharityWebsite     string    `json:"charity_website"`
	CurrentAmount      Amount    `json:"current_amount"`
	TargetAmount       Amount    `json:"target_amount"`
	StoppedAt          time.Time `json:"stopped_at"`
}

type ConduitShardDisabledEvent struct {
	ConduitID string           `json:"conduit_id"`
	ShardID   string           `json:"shard_id"`
	Status    string           `json:"status"`
	Transport ConduitTransport `json:"transport"`
}

type ConduitTransport struct {
	Method         string     `json:"method"`
	SessionID      *string    `json:"session_id"`
	Callback       *string    `json:"callback"`
	ConnectedAt    *time.Time `json:"connected_at"`
	DisconnectedAt *time.Time `json:"disconnected_at"`
}

type DropEntitlementGrantEvent struct {
	Events []DropEvent `json:"events"`
}

type DropEvent struct {
	ID   string   `json:"id"`
	Data DropData `json:"data"`
}

type DropData struct {
	OrganizationID string    `json:"organization_id"`
	CategoryID     string    `json:"category_id"`
	CategoryName   string    `json:"category_name"`
	CampaignID     string    `json:"campaign_id"`
	UserID         string    `json:"user_id"`
	UserName       string    `json:"user_name"`
	UserLogin      string    `json:"user_login"`
	EntitlementID  string    `json:"entitlement_id"`
	BenefitID      string    `json:"benefit_id"`
	CreatedAt      time.Time `json:"created_at"`
}

type ChannelShieldModeBeginEvent struct {
	BroadcasterUserID    string    `json:"broadcaster_user_id"`
	BroadcasterUserName  string    `json:"broadcaster_user_name"`
	BroadcasterUserLogin string    `json:"broadcaster_user_login"`
	ModeratorUserID      string    `json:"moderator_user_id"`
	ModeratorUserName    string    `json:"moderator_user_name"`
	ModeratorUserLogin   string    `json:"moderator_user_login"`
	StartedAt            time.Time `json:"started_at"`
}

type ChannelShieldModeEndEvent struct {
	BroadcasterUserID    string    `json:"broadcaster_user_id"`
	BroadcasterUserName  string    `json:"broadcaster_user_name"`
	BroadcasterUserLogin string    `json:"broadcaster_user_login"`
	ModeratorUserID      string    `json:"moderator_user_id"`
	ModeratorUserName    string    `json:"moderator_user_name"`
	ModeratorUserLogin   string    `json:"moderator_user_login"`
	EndedAt              time.Time `json:"ended_at"`
}

type ChannelShoutOutCreateEvent struct {
	BroadcasterUserID      string    `json:"broadcaster_user_id"`
	BroadcasterUserName    string    `json:"broadcaster_user_name"`
	BroadcasterUserLogin   string    `json:"broadcaster_user_login"`
	ModeratorUserID        string    `json:"moderator_user_id"`
	ModeratorUserName      string    `json:"moderator_user_name"`
	ModeratorUserLogin     string    `json:"moderator_user_login"`
	ToBroadcasterUserID    string    `json:"to_broadcaster_user_id"`
	ToBroadcasterUserName  string    `json:"to_broadcaster_user_name"`
	ToBroadcasterUserLogin string    `json:"to_broadcaster_user_login"`
	StartedAt              time.Time `json:"started_at"`
	ViewerCount            int       `json:"viewer_count"`
	CooldownEndsAt         time.Time `json:"cooldown_ends_at"`
	TargetCooldownEndsAt   time.Time `json:"target_cooldown_ends_at"`
}

type ChannelShoutOutReceivedEvent struct {
	BroadcasterUserID        string    `json:"broadcaster_user_id"`
	BroadcasterUserName      string    `json:"broadcaster_user_name"`
	BroadcasterUserLogin     string    `json:"broadcaster_user_login"`
	FromBroadcasterUserID    string    `json:"from_broadcaster_user_id"`
	FromBroadcasterUserName  string    `json:"from_broadcaster_user_name"`
	FromBroadcasterUserLogin string    `json:"from_broadcaster_user_login"`
	ViewerCount              int       `json:"viewer_count"`
	StartedAt                time.Time `json:"started_at"`
}

type WhisperReceivedEvent struct {
	FromUserID    string  `json:"from_user_id"`
	FromUserLogin string  `json:"from_user_login"`
	FromUserName  string  `json:"from_user_name"`
	ToUserID      string  `json:"to_user_id"`
	ToUserLogin   string  `json:"to_user_login"`
	ToUserName    string  `json:"to_user_name"`
	WhisperID     string  `json:"whisper_id"`
	Whisper       Whisper `json:"whisper"`
}

type Whisper struct {
	Text string `json:"text"`
}

type ChannelSuspiciousUserUpdateEvent struct {
	BroadcasterUserID    string `json:"broadcaster_user_id"`
	BroadcasterUserName  string `json:"broadcaster_user_name"`
	BroadcasterUserLogin string `json:"broadcaster_user_login"`
	ModeratorUserID      string `json:"moderator_user_id"`
	ModeratorUserName    string `json:"moderator_user_name"`
	ModeratorUserLogin   string `json:"moderator_user_login"`
	UserID               string `json:"user_id"`
	UserName             string `json:"user_name"`
	UserLogin            string `json:"user_login"`
	LowTrustStatus       string `json:"low_trust_status"`
}
