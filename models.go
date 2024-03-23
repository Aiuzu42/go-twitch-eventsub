package twitcheventsub

import (
	"encoding/json"
	"time"
)

type Response struct {
	Challenge    string          `json:"challenge"`
	Subscription Subscription    `json:"subscription"`
	Event        json.RawMessage `json:"event"`
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
	Subscription Subscription `json:"subscription"`
	Event        struct {
		BroadcasterUserID           string   `json:"broadcaster_user_id"`
		BroadcasterUserLogin        string   `json:"broadcaster_user_login"`
		BroadcasterUserName         string   `json:"broadcaster_user_name"`
		Title                       string   `json:"title"`
		Language                    string   `json:"language"`
		CategoryID                  string   `json:"category_id"`
		CategoryName                string   `json:"category_name"`
		ContentClassificationLabels []string `json:"content_classification_labels"`
	} `json:"event"`
}

type ChannelFollowEvent struct {
	Subscription Subscription `json:"subscription"`
	Event        struct {
		UserID               string    `json:"user_id"`
		UserLogin            string    `json:"user_login"`
		UserName             string    `json:"user_name"`
		BroadcasterUserID    string    `json:"broadcaster_user_id"`
		BroadcasterUserLogin string    `json:"broadcaster_user_login"`
		BroadcasterUserName  string    `json:"broadcaster_user_name"`
		FollowedAt           time.Time `json:"followed_at"`
	} `json:"event"`
}

type ChannelSubscribeEvent struct {
	Subscription Subscription `json:"subscription"`
	Event        struct {
		UserID               string `json:"user_id"`
		UserLogin            string `json:"user_login"`
		UserName             string `json:"user_name"`
		BroadcasterUserID    string `json:"broadcaster_user_id"`
		BroadcasterUserLogin string `json:"broadcaster_user_login"`
		BroadcasterUserName  string `json:"broadcaster_user_name"`
		Tier                 string `json:"tier"`
		IsGift               bool   `json:"is_gift"`
	} `json:"event"`
}

type ChannelSubscriptionEndEvent struct {
	Subscription Subscription `json:"subscription"`
	Event        struct {
		UserID               string `json:"user_id"`
		UserLogin            string `json:"user_login"`
		UserName             string `json:"user_name"`
		BroadcasterUserID    string `json:"broadcaster_user_id"`
		BroadcasterUserLogin string `json:"broadcaster_user_login"`
		BroadcasterUserName  string `json:"broadcaster_user_name"`
		Tier                 string `json:"tier"`
		IsGift               bool   `json:"is_gift"`
	} `json:"event"`
}

type ChannelSubscriptionGiftEvent struct {
	Subscription Subscription `json:"subscription"`
	Event        struct {
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
	} `json:"event"`
}

type ChannelSubscriptionMessageEvent struct {
	Subscription Subscription `json:"subscription"`
	Event        struct {
		UserID               string `json:"user_id"`
		UserLogin            string `json:"user_login"`
		UserName             string `json:"user_name"`
		BroadcasterUserID    string `json:"broadcaster_user_id"`
		BroadcasterUserLogin string `json:"broadcaster_user_login"`
		BroadcasterUserName  string `json:"broadcaster_user_name"`
		Tier                 string `json:"tier"`
		Message              struct {
			Text   string `json:"text"`
			Emotes []struct {
				Begin int    `json:"begin"`
				End   int    `json:"end"`
				ID    string `json:"id"`
			} `json:"emotes"`
		} `json:"message"`
		CumulativeMonths int  `json:"cumulative_months"`
		StreakMonths     *int `json:"streak_months"`
		DurationMonths   int  `json:"duration_months"`
	} `json:"event"`
}

type ChannelCheerEvent struct {
	Subscription Subscription `json:"subscription"`
	Event        struct {
		IsAnonymous          bool    `json:"is_anonymous"`
		UserID               *string `json:"user_id"`
		UserLogin            *string `json:"user_login"`
		UserName             *string `json:"user_name"`
		BroadcasterUserID    string  `json:"broadcaster_user_id"`
		BroadcasterUserLogin string  `json:"broadcaster_user_login"`
		BroadcasterUserName  string  `json:"broadcaster_user_name"`
		Message              string  `json:"message"`
		Bits                 int     `json:"bits"`
	} `json:"event"`
}

type ChannelRaidEvent struct {
	Subscription Subscription `json:"subscription"`
	Event        struct {
		FromBroadcasterUserID    string `json:"from_broadcaster_user_id"`
		FromBroadcasterUserLogin string `json:"from_broadcaster_user_login"`
		FromBroadcasterUserName  string `json:"from_broadcaster_user_name"`
		ToBroadcasterUserID      string `json:"to_broadcaster_user_id"`
		ToBroadcasterUserLogin   string `json:"to_broadcaster_user_login"`
		ToBroadcasterUserName    string `json:"to_broadcaster_user_name"`
		Viewers                  int    `json:"viewers"`
	} `json:"event"`
}

type ChannelBanEvent struct {
	Subscription Subscription `json:"subscription"`
	Event        struct {
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
	} `json:"event"`
}

type ChannelUnbanEvent struct {
	Subscription Subscription `json:"subscription"`
	Event        struct {
		UserID               string `json:"user_id"`
		UserLogin            string `json:"user_login"`
		UserName             string `json:"user_name"`
		BroadcasterUserID    string `json:"broadcaster_user_id"`
		BroadcasterUserLogin string `json:"broadcaster_user_login"`
		BroadcasterUserName  string `json:"broadcaster_user_name"`
		ModeratorUserID      string `json:"moderator_user_id"`
		ModeratorUserLogin   string `json:"moderator_user_login"`
		ModeratorUserName    string `json:"moderator_user_name"`
	} `json:"event"`
}

type ChannelModeratorAddEvent struct {
	Subscription Subscription `json:"subscription"`
	Event        struct {
		UserID               string `json:"user_id"`
		UserLogin            string `json:"user_login"`
		UserName             string `json:"user_name"`
		BroadcasterUserID    string `json:"broadcaster_user_id"`
		BroadcasterUserLogin string `json:"broadcaster_user_login"`
		BroadcasterUserName  string `json:"broadcaster_user_name"`
	} `json:"event"`
}

type ChannelModeratorRemoveEvent struct {
	Subscription Subscription `json:"subscription"`
	Event        struct {
		UserID               string `json:"user_id"`
		UserLogin            string `json:"user_login"`
		UserName             string `json:"user_name"`
		BroadcasterUserID    string `json:"broadcaster_user_id"`
		BroadcasterUserLogin string `json:"broadcaster_user_login"`
		BroadcasterUserName  string `json:"broadcaster_user_name"`
	} `json:"event"`
}

type ChannelPointsCustomRewardAddEvent struct {
	Subscription Subscription `json:"subscription"`
	Event        struct {
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
		MaxPerStream                      struct {
			IsEnabled bool `json:"is_enabled"`
			Value     int  `json:"value"`
		} `json:"max_per_stream"`
		MaxPerUserPerStream struct {
			IsEnabled bool `json:"is_enabled"`
			Value     int  `json:"value"`
		} `json:"max_per_user_per_stream"`
		GlobalCooldown struct {
			IsEnabled bool `json:"is_enabled"`
			Seconds   int  `json:"seconds"`
		} `json:"global_cooldown"`
		BackgroundColor string `json:"background_color"`
		Image           *Image `json:"image"`
		DefaultImage    Image  `json:"default_image"`
	} `json:"event"`
}

type Image struct {
	URL1X string `json:"url_1x"`
	URL2X string `json:"url_2x"`
	URL4X string `json:"url_4x"`
}

type ChannelPointsCustomRewardUpdateEvent struct {
	Subscription Subscription `json:"subscription"`
	Event        struct {
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
		MaxPerStream                      struct {
			IsEnabled bool `json:"is_enabled"`
			Value     int  `json:"value"`
		} `json:"max_per_stream"`
		MaxPerUserPerStream struct {
			IsEnabled bool `json:"is_enabled"`
			Value     int  `json:"value"`
		} `json:"max_per_user_per_stream"`
		GlobalCooldown struct {
			IsEnabled bool `json:"is_enabled"`
			Seconds   int  `json:"seconds"`
		} `json:"global_cooldown"`
		BackgroundColor string `json:"background_color"`
		Image           *Image `json:"image"`
		DefaultImage    Image  `json:"default_image"`
	} `json:"event"`
}

type ChannelPointsCustomRewardRemoveEvent struct {
	Subscription Subscription `json:"subscription"`
	Event        struct {
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
		MaxPerStream                      struct {
			IsEnabled bool `json:"is_enabled"`
			Value     int  `json:"value"`
		} `json:"max_per_stream"`
		MaxPerUserPerStream struct {
			IsEnabled bool `json:"is_enabled"`
			Value     int  `json:"value"`
		} `json:"max_per_user_per_stream"`
		GlobalCooldown struct {
			IsEnabled bool `json:"is_enabled"`
			Seconds   int  `json:"seconds"`
		} `json:"global_cooldown"`
		BackgroundColor string `json:"background_color"`
		Image           *Image `json:"image"`
		DefaultImage    Image  `json:"default_image"`
	} `json:"event"`
}

type ChannelPointsCustomRewardRedemptionAddEvent struct {
	Subscription Subscription `json:"subscription"`
	Event        struct {
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
	} `json:"event"`
}

type Reward struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Cost   int    `json:"cost"`
	Prompt string `json:"prompt"`
}

type ChannelPointsCustomRewardRedemptionUpdateEvent struct {
	Subscription Subscription `json:"subscription"`
	Event        struct {
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
	} `json:"event"`
}

type ChannelPollBeginEvent struct {
	Subscription Subscription `json:"subscription"`
	Event        struct {
		ID                   string              `json:"id"`
		BroadcasterUserID    string              `json:"broadcaster_user_id"`
		BroadcasterUserLogin string              `json:"broadcaster_user_login"`
		BroadcasterUserName  string              `json:"broadcaster_user_name"`
		Title                string              `json:"title"`
		Choices              []Choices           `json:"choices"`
		ChannelPointsVoting  ChannelPointsVoting `json:"channel_points_voting"`
		StartedAt            time.Time           `json:"started_at"`
		EndsAt               time.Time           `json:"ends_at"`
	} `json:"event"`
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
	Subscription Subscription `json:"subscription"`
	Event        struct {
		ID                   string              `json:"id"`
		BroadcasterUserID    string              `json:"broadcaster_user_id"`
		BroadcasterUserLogin string              `json:"broadcaster_user_login"`
		BroadcasterUserName  string              `json:"broadcaster_user_name"`
		Title                string              `json:"title"`
		Choices              []Choices           `json:"choices"`
		ChannelPointsVoting  ChannelPointsVoting `json:"channel_points_voting"`
		StartedAt            time.Time           `json:"started_at"`
		EndsAt               time.Time           `json:"ends_at"`
	} `json:"event"`
}

type ChannelPollEndEvent struct {
	Subscription Subscription `json:"subscription"`
	Event        struct {
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
	} `json:"event"`
}

type ChannelPredictionBeginEvent struct {
	Subscription Subscription `json:"subscription"`
	Event        struct {
		ID                   string     `json:"id"`
		BroadcasterUserID    string     `json:"broadcaster_user_id"`
		BroadcasterUserLogin string     `json:"broadcaster_user_login"`
		BroadcasterUserName  string     `json:"broadcaster_user_name"`
		Title                string     `json:"title"`
		Outcomes             []Outcomes `json:"outcomes"`
		StartedAt            time.Time  `json:"started_at"`
		LocksAt              time.Time  `json:"locks_at"`
	} `json:"event"`
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
	Subscription Subscription `json:"subscription"`
	Event        struct {
		ID                   string     `json:"id"`
		BroadcasterUserID    string     `json:"broadcaster_user_id"`
		BroadcasterUserLogin string     `json:"broadcaster_user_login"`
		BroadcasterUserName  string     `json:"broadcaster_user_name"`
		Title                string     `json:"title"`
		Outcomes             []Outcomes `json:"outcomes"`
		StartedAt            time.Time  `json:"started_at"`
		LocksAt              time.Time  `json:"locks_at"`
	} `json:"event"`
}

type ChannelPredictionLockEvent struct {
	Subscription Subscription `json:"subscription"`
	Event        struct {
		ID                   string     `json:"id"`
		BroadcasterUserID    string     `json:"broadcaster_user_id"`
		BroadcasterUserLogin string     `json:"broadcaster_user_login"`
		BroadcasterUserName  string     `json:"broadcaster_user_name"`
		Title                string     `json:"title"`
		Outcomes             []Outcomes `json:"outcomes"`
		StartedAt            time.Time  `json:"started_at"`
		LockedAt             time.Time  `json:"locked_at"`
	} `json:"event"`
}

type ChannelPredictionEndEvent struct {
	Subscription Subscription `json:"subscription"`
	Event        struct {
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
	} `json:"event"`
}

type ChannelHypeTrainBeginEvent struct {
	Subscription Subscription `json:"subscription"`
	Event        struct {
		ID                   string         `json:"id"`
		BroadcasterUserID    string         `json:"broadcaster_user_id"`
		BroadcasterUserLogin string         `json:"broadcaster_user_login"`
		BroadcasterUserName  string         `json:"broadcaster_user_name"`
		Total                int            `json:"total"`
		Progress             int            `json:"progress"`
		Goal                 int            `json:"goal"`
		TopContributions     []Contribution `json:"top_contributions"`
		LastContribution     Contribution   `json:"last_contribution"`
		Level                int            `json:"level"`
		StartedAt            time.Time      `json:"started_at"`
		ExpiresAt            time.Time      `json:"expires_at"`
	} `json:"event"`
}

type Contribution struct {
	UserID    string `json:"user_id"`
	UserLogin string `json:"user_login"`
	UserName  string `json:"user_name"`
	Type      string `json:"type"`
	Total     int    `json:"total"`
}

type ChannelHypeTrainProgressEvent struct {
	Subscription Subscription `json:"subscription"`
	Event        struct {
		ID                   string       `json:"id"`
		BroadcasterUserID    string       `json:"broadcaster_user_id"`
		BroadcasterUserLogin string       `json:"broadcaster_user_login"`
		BroadcasterUserName  string       `json:"broadcaster_user_name"`
		Level                int          `json:"level"`
		Total                int          `json:"total"`
		Progress             int          `json:"progress"`
		Goal                 int          `json:"goal"`
		TopContributions     Contribution `json:"top_contributions"`
		LastContribution     Contribution `json:"last_contribution"`
		StartedAt            time.Time    `json:"started_at"`
		ExpiresAt            time.Time    `json:"expires_at"`
	} `json:"event"`
}

type ChannelHypeTrainEndEvent struct {
	Subscription Subscription `json:"subscription"`
	Event        struct {
		ID                   string         `json:"id"`
		BroadcasterUserID    string         `json:"broadcaster_user_id"`
		BroadcasterUserLogin string         `json:"broadcaster_user_login"`
		BroadcasterUserName  string         `json:"broadcaster_user_name"`
		Level                int            `json:"level"`
		Total                int            `json:"total"`
		TopContributions     []Contribution `json:"top_contributions"`
		StartedAt            time.Time      `json:"started_at"`
		EndedAt              time.Time      `json:"ended_at"`
		CooldownEndsAt       time.Time      `json:"cooldown_ends_at"`
	} `json:"event"`
}

type CharityCampaignDonateEvent struct {
	Subscription Subscription `json:"subscription"`
	Event        struct {
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
	} `json:"event"`
}

type Amount struct {
	Value         int    `json:"value"`
	DecimalPlaces int    `json:"decimal_places"`
	Currency      string `json:"currency"`
}

type ExtensionBitsTransactionCreateEvent struct {
	Subscription Subscription `json:"subscription"`
	Event        struct {
		ID                   string  `json:"id"`
		ExtensionClientID    string  `json:"extension_client_id"`
		BroadcasterUserID    string  `json:"broadcaster_user_id"`
		BroadcasterUserLogin string  `json:"broadcaster_user_login"`
		BroadcasterUserName  string  `json:"broadcaster_user_name"`
		UserName             string  `json:"user_name"`
		UserLogin            string  `json:"user_login"`
		UserID               string  `json:"user_id"`
		Product              Product `json:"product"`
	} `json:"event"`
}

type Product struct {
	Name          string `json:"name"`
	Sku           string `json:"sku"`
	Bits          int    `json:"bits"`
	InDevelopment bool   `json:"in_development"`
}

type ChannelGoalBeginEvent struct {
	Subscription Subscription `json:"subscription"`
	Event        struct {
		ID                   string    `json:"id"`
		BroadcasterUserID    string    `json:"broadcaster_user_id"`
		BroadcasterUserName  string    `json:"broadcaster_user_name"`
		BroadcasterUserLogin string    `json:"broadcaster_user_login"`
		Type                 string    `json:"type"`
		Description          string    `json:"description"`
		CurrentAmount        int       `json:"current_amount"`
		TargetAmount         int       `json:"target_amount"`
		StartedAt            time.Time `json:"started_at"`
	} `json:"event"`
}

type ChannelGoalProgressEvent struct {
	Subscription Subscription `json:"subscription"`
	Event        struct {
		ID                   string    `json:"id"`
		BroadcasterUserID    string    `json:"broadcaster_user_id"`
		BroadcasterUserName  string    `json:"broadcaster_user_name"`
		BroadcasterUserLogin string    `json:"broadcaster_user_login"`
		Type                 string    `json:"type"`
		Description          string    `json:"description"`
		CurrentAmount        int       `json:"current_amount"`
		TargetAmount         int       `json:"target_amount"`
		StartedAt            time.Time `json:"started_at"`
	} `json:"event"`
}

type ChannelGoalEndEvent struct {
	Subscription Subscription `json:"subscription"`
	Event        struct {
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
	} `json:"event"`
}

type StreamOnlineEvent struct {
	Subscription Subscription `json:"subscription"`
	Event        struct {
		ID                   string    `json:"id"`
		BroadcasterUserID    string    `json:"broadcaster_user_id"`
		BroadcasterUserLogin string    `json:"broadcaster_user_login"`
		BroadcasterUserName  string    `json:"broadcaster_user_name"`
		Type                 string    `json:"type"`
		StartedAt            time.Time `json:"started_at"`
	} `json:"event"`
}

type StreamOfflineEvent struct {
	Subscription Subscription `json:"subscription"`
	Event        struct {
		BroadcasterUserID    string `json:"broadcaster_user_id"`
		BroadcasterUserLogin string `json:"broadcaster_user_login"`
		BroadcasterUserName  string `json:"broadcaster_user_name"`
	} `json:"event"`
}

type UserAuthorizationGrantEvent struct {
	Subscription Subscription `json:"subscription"`
	Event        struct {
		ClientID  string `json:"client_id"`
		UserID    string `json:"user_id"`
		UserLogin string `json:"user_login"`
		UserName  string `json:"user_name"`
	} `json:"event"`
}

type UserAuthorizationRevokeEvent struct {
	Subscription Subscription `json:"subscription"`
	Event        struct {
		ClientID  string  `json:"client_id"`
		UserID    string  `json:"user_id"`
		UserLogin *string `json:"user_login"`
		UserName  *string `json:"user_name"`
	} `json:"event"`
}

type UserUpdateEvent struct {
	Subscription Subscription `json:"subscription"`
	Event        struct {
		UserID        string `json:"user_id"`
		UserLogin     string `json:"user_login"`
		UserName      string `json:"user_name"`
		Email         string `json:"email"`
		EmailVerified bool   `json:"email_verified"`
		Description   string `json:"description"`
	} `json:"event"`
}

type AutomodMessageHoldEvent struct {
	Subscription Subscription `json:"subscription"`
	Event        struct {
		BroadcasterUserID    string    `json:"broadcaster_user_id"`
		BroadcasterUserName  string    `json:"broadcaster_user_name"`
		BroadcasterUserLogin string    `json:"broadcaster_user_login"`
		UserID               string    `json:"user_id"`
		UserName             string    `json:"user_name"`
		UserLogin            string    `json:"user_login"`
		MessageID            string    `json:"message_id"`
		Message              string    `json:"message"`
		Level                int       `json:"level"`
		Category             string    `json:"category"`
		HeldAt               time.Time `json:"held_at"`
		Fragments            struct {
			Emotes     []Emote     `json:"emotes"`
			Cheermotes []Cheermote `json:"cheermotes"`
		} `json:"fragments"`
	} `json:"event"`
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
	Subscription Subscription `json:"subscription"`
	Event        struct {
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
		Fragments            struct {
			Emotes     []Emote     `json:"emotes"`
			Cheermotes []Cheermote `json:"cheermotes"`
		} `json:"fragments"`
	} `json:"event"`
}

type AutomodSettingsUpdateEvent struct {
	Subscription Subscription `json:"subscription"`
	Event        struct {
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
	} `json:"event"`
}

type AutomodTermsUpdateEvent struct {
	Subscription Subscription `json:"subscription"`
	Event        struct {
		BroadcasterUserID    string   `json:"broadcaster_user_id"`
		BroadcasterUserName  string   `json:"broadcaster_user_name"`
		BroadcasterUserLogin string   `json:"broadcaster_user_login"`
		ModeratorUserID      string   `json:"moderator_user_id"`
		ModeratorUserLogin   string   `json:"moderator_user_login"`
		ModeratorUserName    string   `json:"moderator_user_name"`
		Action               string   `json:"action"`
		FromAutomod          bool     `json:"from_automod"`
		Terms                []string `json:"terms"`
	} `json:"event"`
}

type ChannelAdBreakBeginEvent struct {
	Subscription Subscription `json:"subscription"`
	Event        struct {
		DurationSeconds      string    `json:"duration_seconds"`
		StartedAt            time.Time `json:"started_at"`
		IsAutomatic          string    `json:"is_automatic"`
		BroadcasterUserID    string    `json:"broadcaster_user_id"`
		BroadcasterUserLogin string    `json:"broadcaster_user_login"`
		BroadcasterUserName  string    `json:"broadcaster_user_name"`
		RequesterUserID      string    `json:"requester_user_id"`
		RequesterUserLogin   string    `json:"requester_user_login"`
		RequesterUserName    string    `json:"requester_user_name"`
	} `json:"event"`
}

type ChannelChatClearEvent struct {
	Subscription Subscription `json:"subscription"`
	Event        struct {
		BroadcasterUserID    string `json:"broadcaster_user_id"`
		BroadcasterUserName  string `json:"broadcaster_user_name"`
		BroadcasterUserLogin string `json:"broadcaster_user_login"`
	} `json:"event"`
}

type ChannelChatClearUserMessagesEvent struct {
	Subscription Subscription `json:"subscription"`
	Event        struct {
		BroadcasterUserID    string `json:"broadcaster_user_id"`
		BroadcasterUserName  string `json:"broadcaster_user_name"`
		BroadcasterUserLogin string `json:"broadcaster_user_login"`
		TargetUserID         string `json:"target_user_id"`
		TargetUserName       string `json:"target_user_name"`
		TargetUserLogin      string `json:"target_user_login"`
	} `json:"event"`
}

type ChannelChatMessageEvent struct {
	Subscription Subscription `json:"subscription"`
	Event        struct {
		BroadcasterUserID    string `json:"broadcaster_user_id"`
		BroadcasterUserLogin string `json:"broadcaster_user_login"`
		BroadcasterUserName  string `json:"broadcaster_user_name"`
		ChatterUserID        string `json:"chatter_user_id"`
		ChatterUserLogin     string `json:"chatter_user_login"`
		ChatterUserName      string `json:"chatter_user_name"`
		MessageID            string `json:"message_id"`
		Message              struct {
			Text      string `json:"text"`
			Fragments []struct {
				Type      string     `json:"type"`
				Text      string     `json:"text"`
				Cheermote *Cheermote `json:"cheermote"`
				Emote     *struct {
					ID         string `json:"id"`
					EmoteSetID string `json:"emote_set_id"`
					Format     string `json:"format"`
				} `json:"emote"`
				Mention *struct {
					UserID    string `json:"user_id"`
					UserName  string `json:"user_name"`
					UserLogin string `json:"user_login"`
				} `json:"mention"`
			} `json:"fragments"`
		} `json:"message"`
		Color  string `json:"color"`
		Badges []struct {
			SetID string `json:"set_id"`
			ID    string `json:"id"`
			Info  string `json:"info"`
		} `json:"badges"`
		MessageType string `json:"message_type"`
		Cheer       *struct {
			Bits int `json:"bits"`
		} `json:"cheer"`
		Reply *struct {
			ParentMessageID   string `json:"parent_message_id"`
			ParentMessageBody string `json:"parent_message_body"`
			ParentUserID      string `json:"parent_user_id"`
			ParentUserName    string `json:"parent_user_name"`
			ParentUserLogin   string `json:"parent_user_login"`
			ThreadMessageID   string `json:"thread_message_id"`
			ThreadUserID      string `json:"thread_user_id"`
			ThreadUserName    string `json:"thread_user_name"`
			ThreadUserLogin   string `json:"thread_user_login"`
		} `json:"reply"`
		ChannelPointsCustomRewardID *string `json:"channel_points_custom_reward_id"`
	} `json:"event"`
}

type ChannelChatMessageDeleteEvent struct {
	Subscription Subscription `json:"subscription"`
	Event        struct {
		BroadcasterUserID    string `json:"broadcaster_user_id"`
		BroadcasterUserName  string `json:"broadcaster_user_name"`
		BroadcasterUserLogin string `json:"broadcaster_user_login"`
		TargetUserID         string `json:"target_user_id"`
		TargetUserName       string `json:"target_user_name"`
		TargetUserLogin      string `json:"target_user_login"`
		MessageID            string `json:"message_id"`
	} `json:"event"`
}

type ChannelChatNotificationEvent struct {
	Subscription Subscription `json:"subscription"`
	Event        struct {
		BroadcasterUserID    string `json:"broadcaster_user_id"`
		BroadcasterUserName  string `json:"broadcaster_user_name"`
		BroadcasterUserLogin string `json:"broadcaster_user_login"`
		ChatterUserID        string `json:"chatter_user_id"`
		ChatterUserLogin     string `json:"chatter_user_login"`
		ChatterUserName      string `json:"chatter_user_name"`
		ChatterIsAnonymous   bool   `json:"chatter_is_anonymous"`
		Color                string `json:"color"`
		Badges               []struct {
			SetID string `json:"set_id"`
			ID    string `json:"id"`
			Info  string `json:"info"`
		} `json:"badges"`
		SystemMessage string `json:"system_message"`
		MessageID     string `json:"message_id"`
		Message       struct {
			Text      string `json:"text"`
			Fragments []struct {
				Type      string `json:"type"`
				Text      string `json:"text"`
				Cheermote *struct {
					Prefix string `json:"prefix"`
					Bits   int    `json:"bits"`
					Tier   int    `json:"tier"`
				} `json:"cheermote"`
				Emote *struct {
					ID         string   `json:"id"`
					EmoteSetID string   `json:"emote_set_id"`
					OwnerID    string   `json:"owner_id"`
					Format     []string `json:"format"`
				} `json:"emote"`
				Mention *struct {
					UserID    string `json:"user_id"`
					UserName  string `json:"user_name"`
					UserLogin string `json:"user_login"`
				} `json:"mention"`
			} `json:"fragments"`
		} `json:"message"`
		NoticeType string `json:"notice_type"`
		Sub        *struct {
			SubTier        string `json:"sub_tier"`
			IsPrime        bool   `json:"is_prime"`
			DurationMonths int    `json:"duration_months"`
		} `json:"sub"`
		Resub *struct {
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
		} `json:"resub"`
		SubGift *struct {
			DurationMonths     int     `json:"duration_months"`
			CumulativeTotal    *int    `json:"cumulative_months"`
			RecipientUserID    string  `json:"recipient_user_id"`
			RecipientUserName  string  `json:"recipient_user_name"`
			RecipientUserLogin string  `json:"recipient_user_login"`
			SubTier            string  `json:"sub_tier"`
			CommunityGiftID    *string `json:"community_gift_id"`
		} `json:"sub_gift"`
		CommunitySubGift *struct {
			ID              string `json:"id"`
			Total           int    `json:"total"`
			SubTier         string `json:"sub_tier"`
			CumulativeTotal *int   `json:"cumulative_total"`
		} `json:"community_sub_gift"`
		GiftPaidUpgrade *struct {
			GifterIsAnonymous bool    `json:"gifter_is_anonymous"`
			GifterUserID      *string `json:"gifter_user_id"`
			GifterUserName    *string `json:"gifter_user_name"`
			GifterUserLogin   *string `json:"gifter_user_login"`
		} `json:"gift_paid_upgrade"`
		PrimePaidUpgrade *struct {
			SubTier string `json:"sub_tier"`
		} `json:"prime_paid_upgrade"`
		PayItForward *struct {
			GifterIsAnonymous bool    `json:"gifter_is_anonymous"`
			GifterUserID      *string `json:"gifter_user_id"`
			GifterUserName    *string `json:"gifter_user_name"`
			GifterUserLogin   *string `json:"gifter_user_login"`
		} `json:"pay_it_forward"`
		Raid *struct {
			UserID          string `json:"user_id"`
			UserName        string `json:"user_name"`
			UserLogin       string `json:"user_login"`
			ViewerCount     int    `json:"viewer_count"`
			ProfileImageURL string `json:"profile_image_url"`
		} `json:"raid"`
		Unraid       any `json:"unraid"`
		Announcement struct {
			Color string `json:"color"`
		} `json:"announcement"`
		BitsBadgeTier *struct {
			Tier int `json:"tier"`
		} `json:"bits_badge_tier"`
		CharityDonation *struct {
			CharityName string `json:"charity_name"`
			Amount      struct {
				Value        int    `json:"value"`
				DecimalPlace int    `json:"decimal_place"`
				Currency     string `json:"currency"`
			} `json:"amount"`
		} `json:"charity_donation"`
	} `json:"event"`
}

type ChannelChatSettingsUpdateEvent struct {
	Subscription Subscription `json:"subscription"`
	Event        struct {
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
	} `json:"event"`
}

type ChannelChatUserMessageHoldEvent struct {
	Subscription Subscription `json:"subscription"`
	Event        struct {
		BroadcasterUserID    string `json:"broadcaster_user_id"`
		BroadcasterUserLogin string `json:"broadcaster_user_login"`
		BroadcasterUserName  string `json:"broadcaster_user_name"`
		UserID               string `json:"user_id"`
		UserLogin            string `json:"user_login"`
		UserName             string `json:"user_name"`
		MessageID            string `json:"message_id"`
		Message              struct {
			Text      string `json:"text"`
			Fragments []struct {
				Type      string `json:"type"`
				Text      string `json:"text"`
				Cheermote *struct {
					Prefix string `json:"prefix"`
					Bits   int    `json:"bits"`
					Tier   int    `json:"tier"`
				} `json:"cheermote"`
				Emote *struct {
					ID         string `json:"id"`
					EmoteSetID string `json:"emote_set_id"`
				} `json:"emote"`
			} `json:"fragments"`
		} `json:"message"`
	} `json:"event"`
}

type ChannelChatUserMessageUpdateEvent struct {
	Subscription Subscription `json:"subscription"`
	Event        struct {
		BroadcasterUserID    string `json:"broadcaster_user_id"`
		BroadcasterUserLogin string `json:"broadcaster_user_login"`
		BroadcasterUserName  string `json:"broadcaster_user_name"`
		UserID               string `json:"user_id"`
		UserLogin            string `json:"user_login"`
		UserName             string `json:"user_name"`
		Status               string `json:"status"`
		MessageID            string `json:"message_id"`
		Message              struct {
			Text      string `json:"text"`
			Fragments []struct {
				Type      string `json:"type"`
				Text      string `json:"text"`
				Cheermote *struct {
					Prefix string `json:"prefix"`
					Bits   int    `json:"bits"`
					Tier   int    `json:"tier"`
				} `json:"cheermote"`
				Emote *struct {
					ID         string `json:"id"`
					EmoteSetID string `json:"emote_set_id"`
				} `json:"emote"`
			} `json:"fragments"`
		} `json:"message"`
	} `json:"event"`
}

type ChannelUnbanRequestCreateEvent struct {
	Subscription Subscription `json:"subscription"`
	Event        struct {
		ID                   string    `json:"id"`
		BroadcasterUserID    string    `json:"broadcaster_user_id"`
		BroadcasterUserLogin string    `json:"broadcaster_user_login"`
		BroadcasterUserName  string    `json:"broadcaster_user_name"`
		UserID               string    `json:"user_id"`
		UserLogin            string    `json:"user_login"`
		UserName             string    `json:"user_name"`
		Text                 string    `json:"text"`
		CreatedAt            time.Time `json:"created_at"`
	} `json:"event"`
}

type ChannelUnbanRequestResolveEvent struct {
	Subscription Subscription `json:"subscription"`
	Event        struct {
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
	} `json:"event"`
}

type ChannelModerateEvent struct {
	Subscription Subscription `json:"subscription"`
	Event        struct {
		BroadcasterUserID    string `json:"broadcaster_user_id"`
		BroadcasterUserLogin string `json:"broadcaster_user_login"`
		BroadcasterUserName  string `json:"broadcaster_user_name"`
		ModeratorUserID      string `json:"moderator_user_id"`
		ModeratorUserLogin   string `json:"moderator_user_login"`
		ModeratorUserName    string `json:"moderator_user_name"`
		Action               string `json:"action"`
		Followers            *struct {
			FollowDurationMinutes int `json:"follow_duration_minutes"`
		} `json:"followers"`
		Slow *struct {
			WaitTimeSeconds int `json:"wait_time_seconds"`
		} `json:"slow"`
		Vip *struct {
			UserID    string `json:"user_id"`
			UserLogin string `json:"user_login"`
			UserName  string `json:"user_name"`
		} `json:"vip"`
		Unvip *struct {
			UserID    string `json:"user_id"`
			UserLogin string `json:"user_login"`
			UserName  string `json:"user_name"`
		} `json:"unvip"`
		Mod *struct {
			UserID    string `json:"user_id"`
			UserLogin string `json:"user_login"`
			UserName  string `json:"user_name"`
		} `json:"mod"`
		Unmod *struct {
			UserID    string `json:"user_id"`
			UserLogin string `json:"user_login"`
			UserName  string `json:"user_name"`
		} `json:"unmod"`
		Ban *struct {
			UserID    string  `json:"user_id"`
			UserLogin string  `json:"user_login"`
			UserName  string  `json:"user_name"`
			Reason    *string `json:"reason"`
		} `json:"ban"`
		Unban *struct {
			UserID    string `json:"user_id"`
			UserLogin string `json:"user_login"`
			UserName  string `json:"user_name"`
		} `json:"unban"`
		Timeout *struct {
			UserID    string    `json:"user_id"`
			UserLogin string    `json:"user_login"`
			UserName  string    `json:"user_name"`
			Reason    *string   `json:"reason"`
			ExpiresAt time.Time `json:"expires_at"`
		} `json:"timeout"`
		Untimeout *struct {
			UserID    string `json:"user_id"`
			UserLogin string `json:"user_login"`
			UserName  string `json:"user_name"`
		} `json:"untimeout"`
		Raid *struct {
			UserID      string `json:"user_id"`
			UserLogin   string `json:"user_login"`
			UserName    string `json:"user_name"`
			ViewerCount int    `json:"viewer_count"`
		} `json:"raid"`
		Unraid *struct {
			UserID    string `json:"user_id"`
			UserLogin string `json:"user_login"`
			UserName  string `json:"user_name"`
		} `json:"unraid"`
		Delete *struct {
			UserID      string `json:"user_id"`
			UserLogin   string `json:"user_login"`
			UserName    string `json:"user_name"`
			MessageID   string `json:"message_id"`
			MessageBody string `json:"message_body"`
		} `json:"delete"`
		AutomodTerms *struct {
			Action      string   `json:"action"`
			List        string   `json:"list"`
			Terms       []string `json:"terms"`
			FromAutomod bool     `json:"from_automod"`
		} `json:"automod_terms"`
		UnbanRequest *struct {
			IsApproved       bool   `json:"is_approved"`
			UserID           string `json:"user_id"`
			UserLogin        string `json:"user_login"`
			UserName         string `json:"user_name"`
			ModeratorMessage string `json:"moderator_message"`
		} `json:"unban_request"`
	} `json:"event"`
}

type ChannelGuestStarSessionBeginEvent struct {
	Subscription Subscription `json:"subscription"`
	Event        struct {
		BroadcasterUserID    string    `json:"broadcaster_user_id"`
		BroadcasterUserName  string    `json:"broadcaster_user_name"`
		BroadcasterUserLogin string    `json:"broadcaster_user_login"`
		ModeratorUserID      string    `json:"moderator_user_id"`
		ModeratorUserName    string    `json:"moderator_user_name"`
		ModeratorUserLogin   string    `json:"moderator_user_login"`
		SessionID            string    `json:"session_id"`
		StartedAt            time.Time `json:"started_at"`
	} `json:"event"`
}

type ChannelGuestStarSessionEndEvent struct {
	Subscription Subscription `json:"subscription"`
	Event        struct {
		BroadcasterUserID    string    `json:"broadcaster_user_id"`
		BroadcasterUserName  string    `json:"broadcaster_user_name"`
		BroadcasterUserLogin string    `json:"broadcaster_user_login"`
		ModeratorUserID      string    `json:"moderator_user_id"`
		ModeratorUserName    string    `json:"moderator_user_name"`
		ModeratorUserLogin   string    `json:"moderator_user_login"`
		SessionID            string    `json:"session_id"`
		StartedAt            time.Time `json:"started_at"`
		EndedAt              time.Time `json:"ended_at"`
	} `json:"event"`
}

type ChannelGuestStarGuestUpdateEvent struct {
	Subscription Subscription `json:"subscription"`
	Event        struct {
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
	} `json:"event"`
}

type ChannelGuestStarSettingsUpdateEvent struct {
	Subscription Subscription `json:"subscription"`
	Event        struct {
		BroadcasterUserID           string `json:"broadcaster_user_id"`
		BroadcasterUserName         string `json:"broadcaster_user_name"`
		BroadcasterUserLogin        string `json:"broadcaster_user_login"`
		IsModeratorSendLiveEnabled  bool   `json:"is_moderator_send_live_enabled"`
		SlotCount                   int    `json:"slot_count"`
		IsBrowserSourceAudioEnabled bool   `json:"is_browser_source_audio_enabled"`
		GroupLayout                 string `json:"group_layout"`
	} `json:"event"`
}

type ChannelPointsAutomaticRewardRedemptionAddEvent struct {
	Subscription Subscription `json:"subscription"`
	Event        struct {
		BroadcasterUserID    string `json:"broadcaster_user_id"`
		BroadcasterUserName  string `json:"broadcaster_user_name"`
		BroadcasterUserLogin string `json:"broadcaster_user_login"`
		UserID               string `json:"user_id"`
		UserName             string `json:"user_name"`
		UserLogin            string `json:"user_login"`
		ID                   string `json:"id"`
		Reward               struct {
			Type          string `json:"type"`
			Cost          int    `json:"cost"`
			UnlockedEmote *struct {
				ID   string `json:"id"`
				Name string `json:"name"`
			} `json:"unlocked_emote"`
		} `json:"reward"`
		Message struct {
			Text   string `json:"text"`
			Emotes []struct {
				ID    string `json:"id"`
				Begin int    `json:"begin"`
				End   int    `json:"end"`
			} `json:"emotes"`
		} `json:"message"`
		UserInput  string    `json:"user_input"`
		RedeemedAt time.Time `json:"redeemed_at"`
	} `json:"event"`
}

type ChannelVIPAddEvent struct {
	Subscription Subscription `json:"subscription"`
	Event        struct {
		UserID               string `json:"user_id"`
		UserLogin            string `json:"user_login"`
		UserName             string `json:"user_name"`
		BroadcasterUserID    string `json:"broadcaster_user_id"`
		BroadcasterUserLogin string `json:"broadcaster_user_login"`
		BroadcasterUserName  string `json:"broadcaster_user_name"`
	} `json:"event"`
}

type ChannelVIPRemoveEvent struct {
	Subscription Subscription `json:"subscription"`
	Event        struct {
		UserID               string `json:"user_id"`
		UserLogin            string `json:"user_login"`
		UserName             string `json:"user_name"`
		BroadcasterUserID    string `json:"broadcaster_user_id"`
		BroadcasterUserLogin string `json:"broadcaster_user_login"`
		BroadcasterUserName  string `json:"broadcaster_user_name"`
	} `json:"event"`
}

type CharityCampaignStartEvent struct {
	Subscription Subscription `json:"subscription"`
	Event        struct {
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
	} `json:"event"`
}

type CharityCampaignProgressEvent struct {
	Subscription Subscription `json:"subscription"`
	Event        struct {
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
	} `json:"event"`
}

type CharityCampaignStopEvent struct {
	Subscription Subscription `json:"subscription"`
	Event        struct {
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
	} `json:"event"`
}

type ConduitShardDisabledEvent struct {
	Subscription Subscription `json:"subscription"`
	Event        struct {
		ConduitID string `json:"conduit_id"`
		ShardID   string `json:"shard_id"`
		Status    string `json:"status"`
		Transport struct {
			Method         string     `json:"method"`
			SessionID      *string    `json:"session_id"`
			Callback       *string    `json:"callback"`
			ConnectedAt    *time.Time `json:"connected_at"`
			DisconnectedAt *time.Time `json:"disconnected_at"`
		} `json:"transport"`
	} `json:"event"`
}

type DropEntitlementGrantEvent struct {
	Subscription Subscription `json:"subscription"`
	Events       []struct {
		ID   string `json:"id"`
		Data struct {
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
		} `json:"data"`
	} `json:"events"`
}

type ChannelShieldModeBeginEvent struct {
	Subscription Subscription `json:"subscription"`
	Event        struct {
		BroadcasterUserID    string    `json:"broadcaster_user_id"`
		BroadcasterUserName  string    `json:"broadcaster_user_name"`
		BroadcasterUserLogin string    `json:"broadcaster_user_login"`
		ModeratorUserID      string    `json:"moderator_user_id"`
		ModeratorUserName    string    `json:"moderator_user_name"`
		ModeratorUserLogin   string    `json:"moderator_user_login"`
		StartedAt            time.Time `json:"started_at"`
	} `json:"event"`
}

type ChannelShieldModeEndEvent struct {
	Subscription Subscription `json:"subscription"`
	Event        struct {
		BroadcasterUserID    string    `json:"broadcaster_user_id"`
		BroadcasterUserName  string    `json:"broadcaster_user_name"`
		BroadcasterUserLogin string    `json:"broadcaster_user_login"`
		ModeratorUserID      string    `json:"moderator_user_id"`
		ModeratorUserName    string    `json:"moderator_user_name"`
		ModeratorUserLogin   string    `json:"moderator_user_login"`
		EndedAt              time.Time `json:"ended_at"`
	} `json:"event"`
}

type ChannelShoutOutCreateEvent struct {
	Subscription Subscription `json:"subscription"`
	Event        struct {
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
	} `json:"event"`
}

type ChannelShoutOutReceivedEvent struct {
	Subscription Subscription `json:"subscription"`
	Event        struct {
		BroadcasterUserID        string    `json:"broadcaster_user_id"`
		BroadcasterUserName      string    `json:"broadcaster_user_name"`
		BroadcasterUserLogin     string    `json:"broadcaster_user_login"`
		FromBroadcasterUserID    string    `json:"from_broadcaster_user_id"`
		FromBroadcasterUserName  string    `json:"from_broadcaster_user_name"`
		FromBroadcasterUserLogin string    `json:"from_broadcaster_user_login"`
		ViewerCount              int       `json:"viewer_count"`
		StartedAt                time.Time `json:"started_at"`
	} `json:"event"`
}

type WhisperReceivedEvent struct {
	Subscription Subscription `json:"subscription"`
	Event        struct {
		FromUserID    string `json:"from_user_id"`
		FromUserLogin string `json:"from_user_login"`
		FromUserName  string `json:"from_user_name"`
		ToUserID      string `json:"to_user_id"`
		ToUserLogin   string `json:"to_user_login"`
		ToUserName    string `json:"to_user_name"`
		WhisperID     string `json:"whisper_id"`
		Whisper       struct {
			Text string `json:"text"`
		} `json:"whisper"`
	} `json:"event"`
}
