package twitcheventsub

import (
	"encoding/json"
	"time"
)

type ResponseScheme struct {
	Challenge    string             `json:"challenge"`
	Subscription SubscriptionScheme `json:"subscription"`
	Event        json.RawMessage    `json:"event"`
}

type SubscriptionScheme struct {
	Id        string          `json:"id"`
	Status    string          `json:"status"`
	Type      string          `json:"type"`
	Version   string          `json:"version"`
	Cost      int             `json:"cost"`
	Condition ConditionSchema `json:"condition"`
	Transport TransportSchema `json:"transport"`
	CreatedAt time.Time       `json:"created_at"`
}

type SubscriptionResponse struct {
	Data         []SubscriptionScheme `json:"data"`
	Total        int                  `json:"total"`
	TotalCost    int                  `json:"total_cost"`
	MaxTotalCost int                  `json:"max_total_cost"`
	Pagination   PaginationSchema     `json:"pagination"`
}

type SubscriptionRequest struct {
	Type      string          `json:"type"`
	Version   string          `json:"version"`
	Condition ConditionSchema `json:"condition"`
	Transport TransportSchema `json:"transport"`
}

type ConditionSchema struct {
	BroadcasterUserId string `json:"broadcaster_user_id"`
}

type TransportSchema struct {
	Method   string `json:"method"`
	Callback string `json:"callback"`
	Secret   string `json:"secret"`
}

type PaginationSchema struct {
	Cursor string `json:"cursor"`
}

type ChannelUpdateEvent struct {
	BroadcasterUserID    string `json:"broadcaster_user_id"`
	BroadcasterUserLogin string `json:"broadcaster_user_login"`
	BroadcasterUserName  string `json:"broadcaster_user_name"`
	Title                string `json:"title"`
	Language             string `json:"language"`
	CategoryID           string `json:"category_id"`
	CategoryName         string `json:"category_name"`
	IsMature             bool   `json:"is_mature"`
}

type FollowEvent struct {
	UserID               string    `json:"user_id"`
	UserLogin            string    `json:"user_login"`
	UserName             string    `json:"user_name"`
	BroadcasterUserID    string    `json:"broadcaster_user_id"`
	BroadcasterUserLogin string    `json:"broadcaster_user_login"`
	BroadcasterUserName  string    `json:"broadcaster_user_name"`
	FollowedAt           time.Time `json:"followed_at"`
}

type SubscribeEvent struct {
	UserID               string `json:"user_id"`
	UserLogin            string `json:"user_login"`
	UserName             string `json:"user_name"`
	BroadcasterUserID    string `json:"broadcaster_user_id"`
	BroadcasterUserLogin string `json:"broadcaster_user_login"`
	BroadcasterUserName  string `json:"broadcaster_user_name"`
	Tier                 string `json:"tier"`
	IsGift               bool   `json:"is_gift"`
}

type SubscriptionEndEvent struct {
	UserID               string `json:"user_id"`
	UserLogin            string `json:"user_login"`
	UserName             string `json:"user_name"`
	BroadcasterUserID    string `json:"broadcaster_user_id"`
	BroadcasterUserLogin string `json:"broadcaster_user_login"`
	BroadcasterUserName  string `json:"broadcaster_user_name"`
	Tier                 string `json:"tier"`
	IsGift               bool   `json:"is_gift"`
}

type SubscriptionGiftEvent struct {
	UserID               string `json:"user_id"`
	UserLogin            string `json:"user_login"`
	UserName             string `json:"user_name"`
	BroadcasterUserID    string `json:"broadcaster_user_id"`
	BroadcasterUserLogin string `json:"broadcaster_user_login"`
	BroadcasterUserName  string `json:"broadcaster_user_name"`
	Total                int    `json:"total"`
	Tier                 string `json:"tier"`
	CumulativeTotal      *int   `json:"cumulative_total"`
	IsAnonymous          bool   `json:"is_anonymous"`
}

type SubscriptionMessageEvent struct {
	UserID               string        `json:"user_id"`
	UserLogin            string        `json:"user_login"`
	UserName             string        `json:"user_name"`
	BroadcasterUserID    string        `json:"broadcaster_user_id"`
	BroadcasterUserLogin string        `json:"broadcaster_user_login"`
	BroadcasterUserName  string        `json:"broadcaster_user_name"`
	Tier                 string        `json:"tier"`
	Message              MessageSchema `json:"message"`
	CumulativeMonths     int           `json:"cumulative_months"`
	StreakMonths         *int          `json:"streak_months"`
	DurationMonths       int           `json:"duration_months"`
}

type MessageSchema struct {
	Text   string         `json:"text"`
	Emotes []EmotesSchema `json:"emotes"`
}

type EmotesSchema struct {
	Begin int    `json:"begin"`
	End   int    `json:"end"`
	Id    string `json:"id"`
}

type CheerEvent struct {
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

type RaidEvent struct {
	FromBroadcasterUserID    string `json:"from_broadcaster_user_id"`
	FromBroadcasterUserLogin string `json:"from_broadcaster_user_login"`
	FromBroadcasterUserName  string `json:"from_broadcaster_user_name"`
	ToBroadcasterUserID      string `json:"to_broadcaster_user_id"`
	ToBroadcasterUserLogin   string `json:"to_broadcaster_user_login"`
	ToBroadcasterUserName    string `json:"to_broadcaster_user_name"`
	Viewers                  int    `json:"viewers"`
}

type BanEvent struct {
	UserID               string    `json:"user_id"`
	UserLogin            string    `json:"user_login"`
	UserName             string    `json:"user_name"`
	BroadcasterUserID    string    `json:"broadcaster_user_id"`
	BroadcasterUserLogin string    `json:"broadcaster_user_login"`
	BroadcasterUserName  string    `json:"broadcaster_user_name"`
	ModeratorUserID      string    `json:"moderator_user_id"`
	ModeratorUserLogin   string    `json:"moderator_user_login"`
	ModeratorUserName    string    `json:"moderator_user_name"`
	Reason               string    `json:"reason"`
	BannedAt             time.Time `json:"banned_at"`
	EndsAt               time.Time `json:"ends_at"`
	IsPermanent          bool      `json:"is_permanent"`
}

type UnbanEvent struct {
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

type ModeratorAddEvent struct {
	UserID               string `json:"user_id"`
	UserLogin            string `json:"user_login"`
	UserName             string `json:"user_name"`
	BroadcasterUserID    string `json:"broadcaster_user_id"`
	BroadcasterUserLogin string `json:"broadcaster_user_login"`
	BroadcasterUserName  string `json:"broadcaster_user_name"`
}

type ModeratorRemoveEvent struct {
	UserID               string `json:"user_id"`
	UserLogin            string `json:"user_login"`
	UserName             string `json:"user_name"`
	BroadcasterUserID    string `json:"broadcaster_user_id"`
	BroadcasterUserLogin string `json:"broadcaster_user_login"`
	BroadcasterUserName  string `json:"broadcaster_user_name"`
}

type ChannelPointsCustomRewardAddEvent struct {
	ID                                string               `json:"id"`
	BroadcasterUserID                 string               `json:"broadcaster_user_id"`
	BroadcasterUserLogin              string               `json:"broadcaster_user_login"`
	BroadcasterUserName               string               `json:"broadcaster_user_name"`
	IsEnabled                         bool                 `json:"is_enabled"`
	IsPaused                          bool                 `json:"is_paused"`
	IsInStock                         bool                 `json:"is_in_stock"`
	Title                             string               `json:"title"`
	Cost                              int                  `json:"cost"`
	Prompt                            string               `json:"prompt"`
	IsUserInputRequired               bool                 `json:"is_user_input_required"`
	ShouldRedemptionsSkipRequestQueue bool                 `json:"should_redemptions_skip_request_queue"`
	CooldownExpiresAt                 *string              `json:"cooldown_expires_at"`
	RedemptionsRedeemedCurrentStream  *int                 `json:"redemptions_redeemed_current_stream"`
	MaxPerStream                      MaxPerStreamSchema   `json:"max_per_stream"`
	MaxPerUserPerStream               MaxPerStreamSchema   `json:"max_per_user_per_stream"`
	GlobalCooldown                    GlobalCooldownSchema `json:"global_cooldown"`
	BackgroundColor                   string               `json:"background_color"`
	Image                             ImageSchema          `json:"image"`
	DefaultImage                      ImageSchema          `json:"default_image"`
}

type MaxPerStreamSchema struct {
	IsEnabled bool `json:"is_enabled"`
	Value     int  `json:"value"`
}

type GlobalCooldownSchema struct {
	IsEnabled bool `json:"is_enabled"`
	Seconds   int  `json:"seconds"`
}

type ImageSchema struct {
	URL1X string `json:"url_1x"`
	URL2X string `json:"url_2x"`
	URL4X string `json:"url_4x"`
}

type ChannelPointsCustomRewardUpdateEvent struct {
	ID                                string               `json:"id"`
	BroadcasterUserID                 string               `json:"broadcaster_user_id"`
	BroadcasterUserLogin              string               `json:"broadcaster_user_login"`
	BroadcasterUserName               string               `json:"broadcaster_user_name"`
	IsEnabled                         bool                 `json:"is_enabled"`
	IsPaused                          bool                 `json:"is_paused"`
	IsInStock                         bool                 `json:"is_in_stock"`
	Title                             string               `json:"title"`
	Cost                              int                  `json:"cost"`
	Prompt                            string               `json:"prompt"`
	IsUserInputRequired               bool                 `json:"is_user_input_required"`
	ShouldRedemptionsSkipRequestQueue bool                 `json:"should_redemptions_skip_request_queue"`
	CooldownExpiresAt                 *string              `json:"cooldown_expires_at"`
	RedemptionsRedeemedCurrentStream  *int                 `json:"redemptions_redeemed_current_stream"`
	MaxPerStream                      MaxPerStreamSchema   `json:"max_per_stream"`
	MaxPerUserPerStream               MaxPerStreamSchema   `json:"max_per_user_per_stream"`
	GlobalCooldown                    GlobalCooldownSchema `json:"global_cooldown"`
	BackgroundColor                   string               `json:"background_color"`
	Image                             ImageSchema          `json:"image"`
	DefaultImage                      ImageSchema          `json:"default_image"`
}

type ChannelPointsCustomRewardRemoveEvent struct {
	ID                                string               `json:"id"`
	BroadcasterUserID                 string               `json:"broadcaster_user_id"`
	BroadcasterUserLogin              string               `json:"broadcaster_user_login"`
	BroadcasterUserName               string               `json:"broadcaster_user_name"`
	IsEnabled                         bool                 `json:"is_enabled"`
	IsPaused                          bool                 `json:"is_paused"`
	IsInStock                         bool                 `json:"is_in_stock"`
	Title                             string               `json:"title"`
	Cost                              int                  `json:"cost"`
	Prompt                            string               `json:"prompt"`
	IsUserInputRequired               bool                 `json:"is_user_input_required"`
	ShouldRedemptionsSkipRequestQueue bool                 `json:"should_redemptions_skip_request_queue"`
	CooldownExpiresAt                 *string              `json:"cooldown_expires_at"`
	RedemptionsRedeemedCurrentStream  *int                 `json:"redemptions_redeemed_current_stream"`
	MaxPerStream                      MaxPerStreamSchema   `json:"max_per_stream"`
	MaxPerUserPerStream               MaxPerStreamSchema   `json:"max_per_user_per_stream"`
	GlobalCooldown                    GlobalCooldownSchema `json:"global_cooldown"`
	BackgroundColor                   string               `json:"background_color"`
	Image                             ImageSchema          `json:"image"`
	DefaultImage                      ImageSchema          `json:"default_image"`
}

type ChannelPointsCustomRewardRedemptionAddEvent struct {
	ID                   string       `json:"id"`
	BroadcasterUserID    string       `json:"broadcaster_user_id"`
	BroadcasterUserLogin string       `json:"broadcaster_user_login"`
	BroadcasterUserName  string       `json:"broadcaster_user_name"`
	UserID               string       `json:"user_id"`
	UserLogin            string       `json:"user_login"`
	UserName             string       `json:"user_name"`
	UserInput            string       `json:"user_input"`
	Status               string       `json:"status"`
	Reward               RewardSchema `json:"reward"`
	RedeemedAt           time.Time    `json:"redeemed_at"`
}

type RewardSchema struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Cost   int    `json:"cost"`
	Prompt string `json:"prompt"`
}

type ChannelPointsCustomRewardRedemptionUpdateEvent struct {
	ID                   string       `json:"id"`
	BroadcasterUserID    string       `json:"broadcaster_user_id"`
	BroadcasterUserLogin string       `json:"broadcaster_user_login"`
	BroadcasterUserName  string       `json:"broadcaster_user_name"`
	UserID               string       `json:"user_id"`
	UserLogin            string       `json:"user_login"`
	UserName             string       `json:"user_name"`
	UserInput            string       `json:"user_input"`
	Status               string       `json:"status"` // Either fulfilled or cancelled
	Reward               RewardSchema `json:"reward"`
	RedeemedAt           time.Time    `json:"redeemed_at"`
}

type PollBeginEvent struct {
	ID                   string          `json:"id"`
	BroadcasterUserID    string          `json:"broadcaster_user_id"`
	BroadcasterUserLogin string          `json:"broadcaster_user_login"`
	BroadcasterUserName  string          `json:"broadcaster_user_name"`
	Title                string          `json:"title"`
	Choices              []ChoicesSchema `json:"choices"`
	BitsVoting           VotingSchema    `json:"bits_voting"`
	ChannelPointsVoting  VotingSchema    `json:"channel_points_voting"`
	StartedAt            time.Time       `json:"started_at"`
	EndsAt               time.Time       `json:"ends_at"`
}

type ChoicesSchema struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}

type VotingSchema struct {
	IsEnabled     bool `json:"is_enabled"`
	AmountPerVote int  `json:"amount_per_vote"`
}

type PollProgressEvent struct {
	ID                   string                  `json:"id"`
	BroadcasterUserID    string                  `json:"broadcaster_user_id"`
	BroadcasterUserLogin string                  `json:"broadcaster_user_login"`
	BroadcasterUserName  string                  `json:"broadcaster_user_name"`
	Title                string                  `json:"title"`
	Choices              []ChoicesProgressSchema `json:"choices"`
	BitsVoting           VotingSchema            `json:"bits_voting"`
	ChannelPointsVoting  VotingSchema            `json:"channel_points_voting"`
	StartedAt            time.Time               `json:"started_at"`
	EndsAt               time.Time               `json:"ends_at"`
}

type ChoicesProgressSchema struct {
	ID                 string `json:"id"`
	Title              string `json:"title"`
	BitsVotes          int    `json:"bits_votes"`
	ChannelPointsVotes int    `json:"channel_points_votes"`
	Votes              int    `json:"votes"`
}

type PollEndEvent struct {
	ID                   string                  `json:"id"`
	BroadcasterUserID    string                  `json:"broadcaster_user_id"`
	BroadcasterUserLogin string                  `json:"broadcaster_user_login"`
	BroadcasterUserName  string                  `json:"broadcaster_user_name"`
	Title                string                  `json:"title"`
	Choices              []ChoicesProgressSchema `json:"choices"`
	BitsVoting           VotingSchema            `json:"bits_voting"`
	ChannelPointsVoting  VotingSchema            `json:"channel_points_voting"`
	Status               string                  `json:"status"`
	StartedAt            time.Time               `json:"started_at"`
	EndedAt              time.Time               `json:"ended_at"`
}

type PredictionBeginEvent struct {
	ID                   string           `json:"id"`
	BroadcasterUserID    string           `json:"broadcaster_user_id"`
	BroadcasterUserLogin string           `json:"broadcaster_user_login"`
	BroadcasterUserName  string           `json:"broadcaster_user_name"`
	Title                string           `json:"title"`
	Outcomes             []OutcomesSchema `json:"outcomes"`
	StartedAt            time.Time        `json:"started_at"`
	LocksAt              time.Time        `json:"locks_at"`
}

type OutcomesSchema struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	Color string `json:"color"`
}

type PredictionProgressEvent struct {
	ID                   string                   `json:"id"`
	BroadcasterUserID    string                   `json:"broadcaster_user_id"`
	BroadcasterUserLogin string                   `json:"broadcaster_user_login"`
	BroadcasterUserName  string                   `json:"broadcaster_user_name"`
	Title                string                   `json:"title"`
	Outcomes             []OutcomesProgressSchema `json:"outcomes"`
	StartedAt            time.Time                `json:"started_at"`
	LocksAt              time.Time                `json:"locks_at"`
}

type OutcomesProgressSchema struct {
	ID            string            `json:"id"`
	Title         string            `json:"title"`
	Color         string            `json:"color"` // can be blue or pink
	Users         int               `json:"users,omitempty"`
	ChannelPoints int               `json:"channel_points,omitempty"`
	TopPredictors []PredictorSchema `json:"top_predictors"` // contains up to 10 users
}

type PredictorSchema struct {
	UserName          string `json:"user_name"`
	UserLogin         string `json:"user_login"`
	UserID            string `json:"user_id"`
	ChannelPointsWon  *int   `json:"channel_points_won"` // null if result is refund or loss
	ChannelPointsUsed int    `json:"channel_points_used"`
}

type PredictionLockEvent struct {
	ID                   string                   `json:"id"`
	BroadcasterUserID    string                   `json:"broadcaster_user_id"`
	BroadcasterUserLogin string                   `json:"broadcaster_user_login"`
	BroadcasterUserName  string                   `json:"broadcaster_user_name"`
	Title                string                   `json:"title"`
	Outcomes             []OutcomesProgressSchema `json:"outcomes"`
	StartedAt            time.Time                `json:"started_at"`
	LockedAt             time.Time                `json:"locked_at"`
}

type PredictionEndEvent struct {
	ID                   string                   `json:"id"`
	BroadcasterUserID    string                   `json:"broadcaster_user_id"`
	BroadcasterUserLogin string                   `json:"broadcaster_user_login"`
	BroadcasterUserName  string                   `json:"broadcaster_user_name"`
	Title                string                   `json:"title"`
	WinningOutcomeID     string                   `json:"winning_outcome_id"`
	Outcomes             []OutcomesProgressSchema `json:"outcomes"`
	Status               string                   `json:"status"` // valid values: resolved, canceled
	StartedAt            time.Time                `json:"started_at"`
	EndedAt              time.Time                `json:"ended_at"`
}

type HypeTrainBeginEvent struct {
	ID                   string                       `json:"id"`
	BroadcasterUserID    string                       `json:"broadcaster_user_id"`
	BroadcasterUserLogin string                       `json:"broadcaster_user_login"`
	BroadcasterUserName  string                       `json:"broadcaster_user_name"`
	Total                int                          `json:"total"`
	Progress             int                          `json:"progress"`
	Goal                 int                          `json:"goal"`
	TopContributions     []HypeTrainContributorSchema `json:"top_contributions"`
	LastContribution     HypeTrainContributorSchema   `json:"last_contribution"`
	Level                int                          `json:"level"`
	StartedAt            time.Time                    `json:"started_at"`
	ExpiresAt            time.Time                    `json:"expires_at"`
}

type HypeTrainContributorSchema struct {
	UserID    string `json:"user_id"`
	UserLogin string `json:"user_login"`
	UserName  string `json:"user_name"`
	Type      string `json:"type"`
	Total     int    `json:"total"`
}

type HypeTrainProgressEvent struct {
	ID                   string                       `json:"id"`
	BroadcasterUserID    string                       `json:"broadcaster_user_id"`
	BroadcasterUserLogin string                       `json:"broadcaster_user_login"`
	BroadcasterUserName  string                       `json:"broadcaster_user_name"`
	Level                int                          `json:"level"`
	Total                int                          `json:"total"`
	Progress             int                          `json:"progress"`
	Goal                 int                          `json:"goal"`
	TopContributions     []HypeTrainContributorSchema `json:"top_contributions"`
	LastContribution     HypeTrainContributorSchema   `json:"last_contribution"`
	StartedAt            time.Time                    `json:"started_at"`
	ExpiresAt            time.Time                    `json:"expires_at"`
}

type HypeTrainEndEvent struct {
	ID                   string                       `json:"id"`
	BroadcasterUserID    string                       `json:"broadcaster_user_id"`
	BroadcasterUserLogin string                       `json:"broadcaster_user_login"`
	BroadcasterUserName  string                       `json:"broadcaster_user_name"`
	Level                int                          `json:"level"`
	Total                int                          `json:"total"`
	TopContributions     []HypeTrainContributorSchema `json:"top_contributions"`
	StartedAt            time.Time                    `json:"started_at"`
	EndedAt              time.Time                    `json:"ended_at"`
	CooldownEndsAt       time.Time                    `json:"cooldown_ends_at"`
}

type CharityCampaignDonateEvent struct {
	CampaignID       string       `json:"campaign_id"`
	BroadcasterID    string       `json:"broadcaster_id"`
	BroadcasterName  string       `json:"broadcaster_name"`
	BroadcasterLogin string       `json:"broadcaster_login"`
	UserID           string       `json:"user_id"`
	UserName         string       `json:"user_name"`
	UserLogin        string       `json:"user_login"`
	Amount           AmountSchema `json:"amount"`
}

type AmountSchema struct {
	Value         int    `json:"value"`
	DecimalPlaces int    `json:"decimal_places"`
	Currency      string `json:"currency"`
}

type ExtensionBitsTransactionCreateEvent struct {
	ID                   string        `json:"id"`
	ExtensionClientID    string        `json:"extension_client_id"`
	BroadcasterUserID    string        `json:"broadcaster_user_id"`
	BroadcasterUserLogin string        `json:"broadcaster_user_login"`
	BroadcasterUserName  string        `json:"broadcaster_user_name"`
	UserName             string        `json:"user_name"`
	UserLogin            string        `json:"user_login"`
	UserID               string        `json:"user_id"`
	Product              ProductSchema `json:"product"`
}

type ProductSchema struct {
	Name          string `json:"name"`
	Sku           string `json:"sku"`
	Bits          int    `json:"bits"`
	InDevelopment bool   `json:"in_development"`
}

type GoalBeginEvent struct {
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

type GoalProgressEvent struct {
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

type GoalEndEvent struct {
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
	EmailVerified bool   `json:"email_verified"` // Requires user:read:email scope
	Description   string `json:"description"`
}
