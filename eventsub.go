package twitcheventsub

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"sync"
)

const (
	headerId        = "Twitch-Eventsub-Message-Id"
	headerTimestamp = "Twitch-Eventsub-Message-Timestamp"
	headerSignature = "Twitch-Eventsub-Message-Signature"
	headerType      = "Twitch-Eventsub-Message-Type"
	headerChallenge = "webhook_callback_verification"
	notification    = "notification"
	revocation      = "revocation"
	hmacPrefix      = "sha256="
	parseError      = "parseNotificationError"
)

type Client struct {
	crtPath     string
	keyPath     string
	secret      string
	callback    string
	secretBytes []byte
	stop        *sync.WaitGroup
	srv         *http.Server

	//Notification handling
	onError   func(err error)
	onRevoked func(sub SubscriptionScheme)

	//Events
	onChannelUpdate                             func(event ChannelUpdateEvent)
	onFollow                                    func(event FollowEvent)
	onSubscribe                                 func(event SubscribeEvent)
	onSubscriptionEnd                           func(event SubscriptionEndEvent)
	onSubscriptionGift                          func(event SubscriptionGiftEvent)
	onSubscriptionMessage                       func(event SubscriptionMessageEvent)
	onCheer                                     func(event CheerEvent)
	onRaid                                      func(event RaidEvent)
	onBan                                       func(event BanEvent)
	onUnban                                     func(event UnbanEvent)
	onModeratorAdd                              func(event ModeratorAddEvent)
	onModeratorRemove                           func(event ModeratorRemoveEvent)
	onChannelPointsCustomRewardAdd              func(event ChannelPointsCustomRewardAddEvent)
	onChannelPointsCustomRewardUpdate           func(event ChannelPointsCustomRewardUpdateEvent)
	onChannelPointsCustomRewardRemove           func(event ChannelPointsCustomRewardRemoveEvent)
	onChannelPointsCustomRewardRedemptionAdd    func(event ChannelPointsCustomRewardRedemptionAddEvent)
	onChannelPointsCustomRewardRedemptionUpdate func(event ChannelPointsCustomRewardRedemptionUpdateEvent)
	onPollBegin                                 func(event PollBeginEvent)
	onPollProgress                              func(event PollProgressEvent)
	onPollEnd                                   func(event PollEndEvent)
	onPredictionBegin                           func(event PredictionBeginEvent)
	onPredictionProgress                        func(event PredictionProgressEvent)
	onPredictionLock                            func(event PredictionLockEvent)
	onPredictionEnd                             func(event PredictionEndEvent)
	onCharityDonation                           func(event CharityCampaignDonateEvent)
	onExtensionBitsTransactionCreate            func(event ExtensionBitsTransactionCreateEvent)
	onGoalBegin                                 func(event GoalBeginEvent)
	onGoalProgress                              func(event GoalProgressEvent)
	onGoalEnd                                   func(event GoalEndEvent)
	onHypeTrainBegin                            func(event HypeTrainBeginEvent)
	onHypeTrainProgress                         func(event HypeTrainProgressEvent)
	onHypeTrainEnd                              func(event HypeTrainEndEvent)
	onStreamOnline                              func(event StreamOnlineEvent)
	onStreamOffline                             func(event StreamOfflineEvent)
	onUserAuthorizationGrant                    func(event UserAuthorizationGrantEvent)
	onUserAuthorizationRevoke                   func(event UserAuthorizationRevokeEvent)
	onUserUpdate                                func(event UserUpdateEvent)
}

func NewClient(crtPath, keyPath, secret, callback string) *Client {
	return &Client{crtPath: crtPath, keyPath: keyPath, secret: secret,
		secretBytes: []byte(secret), callback: callback}
}

func (c *Client) StartServer() {
	c.srv = &http.Server{Addr: ":443"}
	c.stop = &sync.WaitGroup{}
	c.stop.Add(1)
	http.HandleFunc("/eventsub", c.handleEvent)
	defer c.stop.Done()
	if err := c.srv.ListenAndServeTLS(c.crtPath, c.keyPath); err != http.ErrServerClosed {
		panic("eventsub start server: " + err.Error())
	}
}

func (c *Client) StopServer() error {
	if err := c.srv.Shutdown(context.Background()); err != nil {
		return err
	}
	c.stop.Wait()
	return nil
}

func (c *Client) handleEvent(w http.ResponseWriter, req *http.Request) {
	body, err := io.ReadAll(req.Body)
	if err != nil {
		c.onError(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	signature := c.getHmac(req.Header.Get(headerId), req.Header.Get(headerTimestamp), body)
	if signature != req.Header.Get(headerSignature) {
		c.onError(errors.New("Signatures do not match"))
		w.WriteHeader(http.StatusForbidden)
		return
	}
	var data ResponseScheme
	if err := json.Unmarshal(body, &data); err != nil {
		c.onError(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	switch req.Header.Get(headerType) {
	case headerChallenge:
		w.Write([]byte(data.Challenge))
	case notification:
		go c.parseNotification(data)
	case revocation:
		go c.onRevoked(data.Subscription)
	}
}

func (c *Client) getHmac(id, timestamp string, body []byte) string {
	message := []byte(id + timestamp)
	message = append(message, body...)
	hash := hmac.New(sha256.New, c.secretBytes)
	hash.Write(message)
	return hmacPrefix + hex.EncodeToString(hash.Sum(nil))
}

func (c *Client) parseNotification(data ResponseScheme) {
	switch data.Subscription.Type {
	case "channel.update":
		if c.onChannelUpdate == nil {
			break
		}
		var e ChannelUpdateEvent
		err := json.Unmarshal(data.Event, &e)
		if err != nil {
			c.onError(fmt.Errorf("%s[channel.update][%s]: %s", parseError, string(data.Event), err.Error()))
			break
		}
		c.onChannelUpdate(e)
	case "channel.follow":
		if c.onFollow == nil {
			break
		}
		var e FollowEvent
		err := json.Unmarshal(data.Event, &e)
		if err != nil {
			c.onError(fmt.Errorf("%s[channel.follow][%s]: %s", parseError, string(data.Event), err.Error()))
			break
		}
		c.onFollow(e)
	case "channel.subscribe":
		if c.onSubscribe == nil {
			break
		}
		var e SubscribeEvent
		err := json.Unmarshal(data.Event, &e)
		if err != nil {
			c.onError(fmt.Errorf("%s[channel.subscribe][%s]: %s", parseError, string(data.Event), err.Error()))
			break
		}
		c.onSubscribe(e)
	case "channel.subscription.end":
		if c.onSubscriptionEnd == nil {
			break
		}
		var e SubscriptionEndEvent
		err := json.Unmarshal(data.Event, &e)
		if err != nil {
			c.onError(fmt.Errorf("%s[channel.subscription.end][%s]: %s", parseError, string(data.Event), err.Error()))
			break
		}
		c.onSubscriptionEnd(e)
	case "channel.subscription.gift":
		if c.onSubscriptionGift == nil {
			break
		}
		var e SubscriptionGiftEvent
		err := json.Unmarshal(data.Event, &e)
		if err != nil {
			c.onError(fmt.Errorf("%s[channel.subscription.gift][%s]: %s", parseError, string(data.Event), err.Error()))
			break
		}
		c.onSubscriptionGift(e)
	case "channel.subscription.message":
		if c.onSubscriptionMessage == nil {
			break
		}
		var e SubscriptionMessageEvent
		err := json.Unmarshal(data.Event, &e)
		if err != nil {
			c.onError(fmt.Errorf("%s[channel.subscription.message][%s]: %s", parseError, string(data.Event), err.Error()))
			break
		}
		c.onSubscriptionMessage(e)
	case "channel.cheer":
		if c.onCheer == nil {
			break
		}
		var e CheerEvent
		err := json.Unmarshal(data.Event, &e)
		if err != nil {
			c.onError(fmt.Errorf("%s[channel.cheer][%s]: %s", parseError, string(data.Event), err.Error()))
			break
		}
		c.onCheer(e)
	case "channel.raid":
		if c.onRaid == nil {
			break
		}
		var e RaidEvent
		err := json.Unmarshal(data.Event, &e)
		if err != nil {
			c.onError(fmt.Errorf("%s[channel.raid][%s]: %s", parseError, string(data.Event), err.Error()))
			break
		}
		c.onRaid(e)
	case "channel.ban":
		if c.onBan == nil {
			break
		}
		var e BanEvent
		err := json.Unmarshal(data.Event, &e)
		if err != nil {
			c.onError(fmt.Errorf("%s[channel.ban][%s]: %s", parseError, string(data.Event), err.Error()))
			break
		}
		c.onBan(e)
	case "channel.unban":
		if c.onUnban == nil {
			break
		}
		var e UnbanEvent
		err := json.Unmarshal(data.Event, &e)
		if err != nil {
			c.onError(fmt.Errorf("%s[channel.unban][%s]: %s", parseError, string(data.Event), err.Error()))
			break
		}
		c.onUnban(e)
	case "channel.moderator.add":
		if c.onModeratorAdd == nil {
			break
		}
		var e ModeratorAddEvent
		err := json.Unmarshal(data.Event, &e)
		if err != nil {
			c.onError(fmt.Errorf("%s[channel.moderator.add][%s]: %s", parseError, string(data.Event), err.Error()))
			break
		}
		c.onModeratorAdd(e)
	case "channel.moderator.remove":
		if c.onModeratorRemove == nil {
			break
		}
		var e ModeratorRemoveEvent
		err := json.Unmarshal(data.Event, &e)
		if err != nil {
			c.onError(fmt.Errorf("%s[channel.moderator.remove][%s]: %s", parseError, string(data.Event), err.Error()))
			break
		}
		c.onModeratorRemove(e)
	case "channel.channel_points_custom_reward.add":
		if c.onChannelPointsCustomRewardAdd == nil {
			break
		}
		var e ChannelPointsCustomRewardAddEvent
		err := json.Unmarshal(data.Event, &e)
		if err != nil {
			c.onError(fmt.Errorf("%s[channel.channel_points_custom_reward.add][%s]: %s", parseError, string(data.Event), err.Error()))
			break
		}
		c.onChannelPointsCustomRewardAdd(e)
	case "channel.channel_points_custom_reward.update":
		if c.onChannelPointsCustomRewardUpdate == nil {
			break
		}
		var e ChannelPointsCustomRewardUpdateEvent
		err := json.Unmarshal(data.Event, &e)
		if err != nil {
			c.onError(fmt.Errorf("%s[channel.channel_points_custom_reward.update][%s]: %s", parseError, string(data.Event), err.Error()))
			break
		}
		c.onChannelPointsCustomRewardUpdate(e)
	case "channel.channel_points_custom_reward.remove":
		if c.onChannelPointsCustomRewardRemove == nil {
			break
		}
		var e ChannelPointsCustomRewardRemoveEvent
		err := json.Unmarshal(data.Event, &e)
		if err != nil {
			c.onError(fmt.Errorf("%s[channel.channel_points_custom_reward.remove][%s]: %s", parseError, string(data.Event), err.Error()))
			break
		}
		c.onChannelPointsCustomRewardRemove(e)
	case "channel.channel_points_custom_reward_redemption.add":
		if c.onChannelPointsCustomRewardRedemptionAdd == nil {
			break
		}
		var e ChannelPointsCustomRewardRedemptionAddEvent
		err := json.Unmarshal(data.Event, &e)
		if err != nil {
			c.onError(fmt.Errorf("%s[channel.channel_points_custom_reward_redemption.add][%s]: %s", parseError, string(data.Event), err.Error()))
			break
		}
		c.onChannelPointsCustomRewardRedemptionAdd(e)
	case "channel.channel_points_custom_reward_redemption.update":
		if c.onChannelPointsCustomRewardRedemptionUpdate == nil {
			break
		}
		var e ChannelPointsCustomRewardRedemptionUpdateEvent
		err := json.Unmarshal(data.Event, &e)
		if err != nil {
			c.onError(fmt.Errorf("%s[channel.channel_points_custom_reward_redemption.update][%s]: %s", parseError, string(data.Event), err.Error()))
			break
		}
		c.onChannelPointsCustomRewardRedemptionUpdate(e)
	case "channel.poll.begin":
		if c.onPollBegin == nil {
			break
		}
		var e PollBeginEvent
		err := json.Unmarshal(data.Event, &e)
		if err != nil {
			c.onError(fmt.Errorf("%s[channel.poll.begin][%s]: %s", parseError, string(data.Event), err.Error()))
			break
		}
		c.onPollBegin(e)
	case "channel.poll.progress":
		if c.onPollProgress == nil {
			break
		}
		var e PollProgressEvent
		err := json.Unmarshal(data.Event, &e)
		if err != nil {
			c.onError(fmt.Errorf("%s[channel.poll.progress][%s]: %s", parseError, string(data.Event), err.Error()))
			break
		}
		c.onPollProgress(e)
	case "channel.poll.end":
		if c.onPollEnd == nil {
			break
		}
		var e PollEndEvent
		err := json.Unmarshal(data.Event, &e)
		if err != nil {
			c.onError(fmt.Errorf("%s[channel.poll.end][%s]: %s", parseError, string(data.Event), err.Error()))
			break
		}
		c.onPollEnd(e)
	case "channel.prediction.begin":
		if c.onPredictionBegin == nil {
			break
		}
		var e PredictionBeginEvent
		err := json.Unmarshal(data.Event, &e)
		if err != nil {
			c.onError(fmt.Errorf("%s[channel.prediction.begin][%s]: %s", parseError, string(data.Event), err.Error()))
			break
		}
		c.onPredictionBegin(e)
	case "channel.prediction.progress":
		if c.onPredictionProgress == nil {
			break
		}
		var e PredictionProgressEvent
		err := json.Unmarshal(data.Event, &e)
		if err != nil {
			c.onError(fmt.Errorf("%s[channel.prediction.progress][%s]: %s", parseError, string(data.Event), err.Error()))
			break
		}
		c.onPredictionProgress(e)
	case "channel.prediction.lock":
		if c.onPredictionLock == nil {
			break
		}
		var e PredictionLockEvent
		err := json.Unmarshal(data.Event, &e)
		if err != nil {
			c.onError(fmt.Errorf("%s[channel.prediction.lock][%s]: %s", parseError, string(data.Event), err.Error()))
			break
		}
		c.onPredictionLock(e)
	case "channel.prediction.end":
		if c.onPredictionEnd == nil {
			break
		}
		var e PredictionEndEvent
		err := json.Unmarshal(data.Event, &e)
		if err != nil {
			c.onError(fmt.Errorf("%s[channel.prediction.end][%s]: %s", parseError, string(data.Event), err.Error()))
			break
		}
		c.onPredictionEnd(e)
	case "channel.charity_campaign.donate":
		if c.onCharityDonation == nil {
			break
		}
		var e CharityCampaignDonateEvent
		err := json.Unmarshal(data.Event, &e)
		if err != nil {
			c.onError(fmt.Errorf("%s[channel.charity_campaign.donate][%s]: %s", parseError, string(data.Event), err.Error()))
			break
		}
		c.onCharityDonation(e)
	case "extension.bits_transaction.create":
		if c.onExtensionBitsTransactionCreate == nil {
			break
		}
		var e ExtensionBitsTransactionCreateEvent
		err := json.Unmarshal(data.Event, &e)
		if err != nil {
			c.onError(fmt.Errorf("%s[extension.bits_transaction.create][%s]: %s", parseError, string(data.Event), err.Error()))
			break
		}
		c.onExtensionBitsTransactionCreate(e)
	case "channel.goal.begin":
		if c.onGoalBegin == nil {
			break
		}
		var e GoalBeginEvent
		err := json.Unmarshal(data.Event, &e)
		if err != nil {
			c.onError(fmt.Errorf("%s[channel.goal.begin][%s]: %s", parseError, string(data.Event), err.Error()))
			break
		}
		c.onGoalBegin(e)
	case "channel.goal.progress":
		if c.onGoalProgress == nil {
			break
		}
		var e GoalProgressEvent
		err := json.Unmarshal(data.Event, &e)
		if err != nil {
			c.onError(fmt.Errorf("%s[channel.goal.progress][%s]: %s", parseError, string(data.Event), err.Error()))
			break
		}
		c.onGoalProgress(e)
	case "channel.goal.end":
		if c.onGoalEnd == nil {
			break
		}
		var e GoalEndEvent
		err := json.Unmarshal(data.Event, &e)
		if err != nil {
			c.onError(fmt.Errorf("%s[channel.goal.end][%s]: %s", parseError, string(data.Event), err.Error()))
			break
		}
		c.onGoalEnd(e)
	case "channel.hype_train.begin":
		if c.onHypeTrainBegin == nil {
			break
		}
		var e HypeTrainBeginEvent
		err := json.Unmarshal(data.Event, &e)
		if err != nil {
			c.onError(fmt.Errorf("%s[channel.hype_train.begin][%s]: %s", parseError, string(data.Event), err.Error()))
			break
		}
		c.onHypeTrainBegin(e)
	case "channel.hype_train.progress":
		if c.onHypeTrainProgress == nil {
			break
		}
		var e HypeTrainProgressEvent
		err := json.Unmarshal(data.Event, &e)
		if err != nil {
			c.onError(fmt.Errorf("%s[channel.hype_train.progress][%s]: %s", parseError, string(data.Event), err.Error()))
			break
		}
		c.onHypeTrainProgress(e)
	case "channel.hype_train.end":
		if c.onHypeTrainEnd == nil {
			break
		}
		var e HypeTrainEndEvent
		err := json.Unmarshal(data.Event, &e)
		if err != nil {
			c.onError(fmt.Errorf("%s[channel.hype_train.end][%s]: %s", parseError, string(data.Event), err.Error()))
			break
		}
		c.onHypeTrainEnd(e)
	case "stream.online":
		if c.onStreamOnline == nil {
			break
		}
		var e StreamOnlineEvent
		err := json.Unmarshal(data.Event, &e)
		if err != nil {
			c.onError(fmt.Errorf("%s[stream.online][%s]: %s", parseError, string(data.Event), err.Error()))
			break
		}
		c.onStreamOnline(e)
	case "stream.offline":
		if c.onStreamOffline == nil {
			break
		}
		var e StreamOfflineEvent
		err := json.Unmarshal(data.Event, &e)
		if err != nil {
			c.onError(fmt.Errorf("%s[stream.offline][%s]: %s", parseError, string(data.Event), err.Error()))
			break
		}
		c.onStreamOffline(e)
	case "user.authorization.grant":
		if c.onUserAuthorizationGrant == nil {
			break
		}
		var e UserAuthorizationGrantEvent
		err := json.Unmarshal(data.Event, &e)
		if err != nil {
			c.onError(fmt.Errorf("%s[user.authorization.grant][%s]: %s", parseError, string(data.Event), err.Error()))
			break
		}
		c.onUserAuthorizationGrant(e)
	case "user.authorization.revoke":
		if c.onUserAuthorizationRevoke == nil {
			break
		}
		var e UserAuthorizationRevokeEvent
		err := json.Unmarshal(data.Event, &e)
		if err != nil {
			c.onError(fmt.Errorf("%s[user.authorization.revoke][%s]: %s", parseError, string(data.Event), err.Error()))
			break
		}
		c.onUserAuthorizationRevoke(e)
	case "user.update":
		if c.onUserUpdate == nil {
			break
		}
		var e UserUpdateEvent
		err := json.Unmarshal(data.Event, &e)
		if err != nil {
			c.onError(fmt.Errorf("%s[user.update][%s]: %s", parseError, string(data.Event), err.Error()))
			break
		}
		c.onUserUpdate(e)
	default:
		c.onError(fmt.Errorf("%s[default][%s]: Unable to parse event", parseError, string(data.Event)))
	}
}

func (c *Client) OnError(f func(err error)) {
	c.onError = f
}

func (c *Client) OnRevoked(f func(sub SubscriptionScheme)) {
	c.onRevoked = f
}

func (c *Client) OnChannelUpdate(f func(event ChannelUpdateEvent)) {
	c.onChannelUpdate = f
}

func (c *Client) OnFollow(f func(event FollowEvent)) {
	c.onFollow = f
}

func (c *Client) OnSubscribe(f func(event SubscribeEvent)) {
	c.onSubscribe = f
}

func (c *Client) OnSubscriptionEnd(f func(event SubscriptionEndEvent)) {
	c.onSubscriptionEnd = f
}

func (c *Client) OnSubscriptionGift(f func(event SubscriptionGiftEvent)) {
	c.onSubscriptionGift = f
}

func (c *Client) OnSubscriptionMessage(f func(event SubscriptionMessageEvent)) {
	c.onSubscriptionMessage = f
}

func (c *Client) OnCheer(f func(event CheerEvent)) {
	c.onCheer = f
}

func (c *Client) OnRaid(f func(event RaidEvent)) {
	c.onRaid = f
}

func (c *Client) OnBan(f func(event BanEvent)) {
	c.onBan = f
}

func (c *Client) OnUnban(f func(event UnbanEvent)) {
	c.onUnban = f
}

func (c *Client) OnModeratorAdd(f func(event ModeratorAddEvent)) {
	c.onModeratorAdd = f
}

func (c *Client) OnModeratorRemove(f func(event ModeratorRemoveEvent)) {
	c.onModeratorRemove = f
}

func (c *Client) OnChannelPointsCustomRewardAdd(f func(event ChannelPointsCustomRewardAddEvent)) {
	c.onChannelPointsCustomRewardAdd = f
}

func (c *Client) OnChannelPointsCustomRewardUpdate(f func(event ChannelPointsCustomRewardUpdateEvent)) {
	c.onChannelPointsCustomRewardUpdate = f
}

func (c *Client) OnChannelPointsCustomRewardRemove(f func(event ChannelPointsCustomRewardRemoveEvent)) {
	c.onChannelPointsCustomRewardRemove = f
}

func (c *Client) OnChannelPointsCustomRewardRedemptionAdd(f func(event ChannelPointsCustomRewardRedemptionAddEvent)) {
	c.onChannelPointsCustomRewardRedemptionAdd = f
}

func (c *Client) OnChannelPointsCustomRewardRedemptionUpdate(f func(event ChannelPointsCustomRewardRedemptionUpdateEvent)) {
	c.onChannelPointsCustomRewardRedemptionUpdate = f
}

func (c *Client) OnPollBegin(f func(event PollBeginEvent)) {
	c.onPollBegin = f
}

func (c *Client) OnPollProgress(f func(event PollProgressEvent)) {
	c.onPollProgress = f
}

func (c *Client) OnPollEnd(f func(event PollEndEvent)) {
	c.onPollEnd = f
}

func (c *Client) OnPredictionBegin(f func(event PredictionBeginEvent)) {
	c.onPredictionBegin = f
}

func (c *Client) OnPredictionProgress(f func(event PredictionProgressEvent)) {
	c.onPredictionProgress = f
}

func (c *Client) OnPredictionLock(f func(event PredictionLockEvent)) {
	c.onPredictionLock = f
}

func (c *Client) OnPredictionEnd(f func(event PredictionEndEvent)) {
	c.onPredictionEnd = f
}

func (c *Client) OnCharityDonation(f func(event CharityCampaignDonateEvent)) {
	c.onCharityDonation = f
}

func (c *Client) OnExtensionBitsTransactionCreate(f func(event ExtensionBitsTransactionCreateEvent)) {
	c.onExtensionBitsTransactionCreate = f
}

func (c *Client) OnGoalBegin(f func(event GoalBeginEvent)) {
	c.onGoalBegin = f
}

func (c *Client) OnGoalProgress(f func(event GoalProgressEvent)) {
	c.onGoalProgress = f
}

func (c *Client) OnGoalEnd(f func(event GoalEndEvent)) {
	c.onGoalEnd = f
}

func (c *Client) OnHypeTrainBegin(f func(event HypeTrainBeginEvent)) {
	c.onHypeTrainBegin = f
}

func (c *Client) OnHypeTrainProgress(f func(event HypeTrainProgressEvent)) {
	c.onHypeTrainProgress = f
}

func (c *Client) OnHypeTrainEnd(f func(event HypeTrainEndEvent)) {
	c.onHypeTrainEnd = f
}

func (c *Client) OnStreamOnline(f func(event StreamOnlineEvent)) {
	c.onStreamOnline = f
}

func (c *Client) OnStreamOffline(f func(event StreamOfflineEvent)) {
	c.onStreamOffline = f
}

func (c *Client) OnUserAuthorizationGrant(f func(event UserAuthorizationGrantEvent)) {
	c.onUserAuthorizationGrant = f
}

func (c *Client) OnUserAuthorizationRevoke(f func(event UserAuthorizationRevokeEvent)) {
	c.onUserAuthorizationRevoke = f
}

func (c *Client) OnUserUpdate(f func(event UserUpdateEvent)) {
	c.onUserUpdate = f
}
