package twitcheventsub

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
)

type EventType string

const (
	baseUrl                                             = "https://api.twitch.tv/helix/eventsub/subscriptions"
	Update                                    EventType = "channel.update"
	Follow                                              = "channel.follow"
	Subscribe                                           = "channel.subscribe"
	SubscriptionEnd                                     = "channel.subscription.end"
	SubscriptionGift                                    = "channel.subscription.gift"
	SubscriptionMessage                                 = "channel.subscription.message"
	Cheer                                               = "channel.cheer"
	Raid                                                = "channel.raid"
	Ban                                                 = "channel.ban"
	Unban                                               = "channel.unban"
	ModeratorAdd                                        = "channel.moderator.add"
	ModeratorRemove                                     = "channel.moderator.remove"
	ChannelPointsCustomRewardAdd                        = "channel.channel_points_custom_reward.add"
	ChannelPointsCustomRewardUpdate                     = "channel.channel_points_custom_reward.update"
	ChannelPointsCustomRewardRemove                     = "channel.channel_points_custom_reward.remove"
	ChannelPointsCustomRewardRedemptionAdd              = "channel.channel_points_custom_reward_redemption.add"
	ChannelPointsCustomRewardRedemptionUpdate           = "channel.channel_points_custom_reward_redemption.update"
	PollBegin                                           = "channel.poll.begin"
	PollProgress                                        = "channel.poll.progress"
	PollEnd                                             = "channel.poll.end"
	PredictionBegin                                     = "channel.prediction.begin"
	PredictionProgress                                  = "channel.prediction.progress"
	PredictionLock                                      = "channel.prediction.lock"
	PredictionEnd                                       = "channel.prediction.end"
	CharityCampaignDonate                               = "channel.charity_campaign.donate"
	ExtensionBitsTransactionCreate                      = "extension.bits_transaction.create"
	GoalBegin                                           = "channel.goal.begin"
	GoalProgress                                        = "channel.goal.progress"
	GoalEnd                                             = "channel.goal.end"
	HypeTrainBegin                                      = "channel.hype_train.begin"
	HypeTrainProgress                                   = "channel.hype_train.progress"
	HypeTrainEnd                                        = "channel.hype_train.end"
	StreamOnline                                        = "stream.online"
	StreamOffline                                       = "stream.offline"
	UserAuthorizationGrant                              = "user.authorization.grant"
	UserAuthorizationRevoke                             = "user.authorization.revoke"
	UserUpdate                                          = "user.update"
	ChannelSuspiciousUserUpdate                         = "channel.suspicious_user.update"
	ChannelBitsUse                                      = "channel.bits.use"
	ChannelSuspiciousUserMessage                        = "channel.suspicious_user.message"
	ChannelWarningAcknowledge                           = "channel.warning.acknowledge"
	ChannelWarningSend                                  = "channel.warning.send"
	UserWhisperMessage                                  = "user.whisper.message"
)

func (c *Client) SubscribeToEvent(event EventType, broadcasterId, token, clientId string) (SubscriptionResponse, error) {
	subReq := SubscriptionRequest{Type: string(event), Version: "1",
		Condition: Condition{BroadcasterUserId: broadcasterId},
		Transport: Transport{Method: "webhook", Callback: c.callback, Secret: c.secret}}
	payload, err := json.Marshal(subReq)
	if err != nil {
		return SubscriptionResponse{}, errors.New("error encoding: " + err.Error())
	}
	client := http.Client{}
	req, _ := http.NewRequest(http.MethodPost, baseUrl, bytes.NewBuffer(payload))
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Client-Id", clientId)
	req.Header.Set("Content-Type", "application/json")
	res, err := client.Do(req)
	if err != nil {
		return SubscriptionResponse{}, errors.New("error sending request: " + err.Error())
	}
	defer res.Body.Close()
	if res.StatusCode == http.StatusAccepted {
		var response SubscriptionResponse
		err := json.NewDecoder(res.Body).Decode(&response)
		if err != nil {
			return SubscriptionResponse{}, errors.New("error decoding: " + err.Error())
		}
		return response, nil
	}
	return SubscriptionResponse{}, errors.New(res.Status)
}

// DeleteSubscription Requires an application OAuth access token.
func (c *Client) DeleteSubscription(id, token, clientId string) error {
	client := http.Client{}
	url := fmt.Sprintf("%s?id=%s", baseUrl, url.QueryEscape(id))
	req, _ := http.NewRequest(http.MethodDelete, url, nil)
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Client-Id", clientId)
	res, err := client.Do(req)
	if err != nil {
		return errors.New("error sending request: " + err.Error())
	}
	defer res.Body.Close()
	if res.StatusCode == http.StatusNoContent {
		return nil
	}
	return errors.New(res.Status)
}

// GetSubscriptions Requires an application OAuth access token.
func (c *Client) GetSubscriptions(token, clientId, subType, userId, after string, status []string) (SubscriptionResponse, error) {
	v := url.Values{}
	for _, s := range status {
		v.Add("status", s)
	}
	if subType != "" {
		v.Set("type", subType)
	}
	if userId != "" {
		v.Set("user_id", userId)
	}
	if after != "" {
		v.Set("after", after)
	}
	client := http.Client{}
	url := fmt.Sprintf("%s?%s", baseUrl, v.Encode())
	req, _ := http.NewRequest(http.MethodGet, url, nil)
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Client-Id", clientId)
	req.Header.Set("Content-Type", "application/json")
	res, err := client.Do(req)
	if err != nil {
		return SubscriptionResponse{}, errors.New("error sending request: " + err.Error())
	}
	defer res.Body.Close()
	if res.StatusCode == http.StatusOK {
		var response SubscriptionResponse
		err := json.NewDecoder(res.Body).Decode(&response)
		if err != nil {
			return SubscriptionResponse{}, errors.New("error decoding: " + err.Error())
		}
		return response, nil
	}
	return SubscriptionResponse{}, errors.New(res.Status)
}
