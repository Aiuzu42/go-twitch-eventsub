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
	debug       bool

	//Notification handling
	onError   func(err error)
	onRevoked func(sub Subscription)
	onDebug   func(msg string)

	//Events
	onAutomodMessageHold                        func(event AutomodMessageHoldEvent)
	onAutomodMessageUpdate                      func(event AutomodMessageUpdateEvent)
	onAutomodSettingsUpdate                     func(event AutomodSettingsUpdateEvent)
	onAutomodTermsUpdate                        func(event AutomodTermsUpdateEvent)
	onChannelUpdate                             func(event ChannelUpdateEvent)
	onChannelFollow                             func(event ChannelFollowEvent)
	onChannelAdBreakBegin                       func(event ChannelAdBreakBeginEvent)
	onChannelChatClear                          func(event ChannelChatClearEvent)
	onChannelChatClearUserMessages              func(event ChannelChatClearUserMessagesEvent)
	onChannelChatMessage                        func(event ChannelChatMessageEvent)
	onChannelChatMessageDelete                  func(event ChannelChatMessageDeleteEvent)
	onChannelChatNotification                   func(event ChannelChatNotificationEvent)
	onChannelChatSettingsUpdate                 func(event ChannelChatSettingsUpdateEvent)
	onChannelChatUserMessageHold                func(event ChannelChatUserMessageHoldEvent)
	onChannelChatUserMessageUpdate              func(event ChannelChatUserMessageUpdateEvent)
	onChannelSubscribe                          func(event ChannelSubscribeEvent)
	onChannelSubscriptionEnd                    func(event ChannelSubscriptionEndEvent)
	onChannelSubscriptionGift                   func(event ChannelSubscriptionGiftEvent)
	onChannelSubscriptionMessage                func(event ChannelSubscriptionMessageEvent)
	onChannelCheer                              func(event ChannelCheerEvent)
	onChannelRaid                               func(event ChannelRaidEvent)
	onChannelBan                                func(event ChannelBanEvent)
	onChannelUnban                              func(event ChannelUnbanEvent)
	onChannelUnbanRequestCreate                 func(event ChannelUnbanRequestCreateEvent)
	onChannelUnbanRequestResolve                func(event ChannelUnbanRequestResolveEvent)
	onChannelModerate                           func(event ChannelModerateEvent)
	onChannelModeratorAdd                       func(event ChannelModeratorAddEvent)
	onChannelModeratorRemove                    func(event ChannelModeratorRemoveEvent)
	onChannelGuestStarSessionBegin              func(event ChannelGuestStarSessionBeginEvent)
	onChannelGuestStarSessionEnd                func(event ChannelGuestStarSessionEndEvent)
	onChannelGuestStarGuestUpdate               func(event ChannelGuestStarGuestUpdateEvent)
	onChannelGuestStarSettingsUpdate            func(event ChannelGuestStarSettingsUpdateEvent)
	onChannelPointsAutomaticRewardRedemptionAdd func(event ChannelPointsAutomaticRewardRedemptionAddEvent)
	onChannelPointsCustomRewardAdd              func(event ChannelPointsCustomRewardAddEvent)
	onChannelPointsCustomRewardUpdate           func(event ChannelPointsCustomRewardUpdateEvent)
	onChannelPointsCustomRewardRemove           func(event ChannelPointsCustomRewardRemoveEvent)
	onChannelPointsCustomRewardRedemptionAdd    func(event ChannelPointsCustomRewardRedemptionAddEvent)
	onChannelPointsCustomRewardRedemptionUpdate func(event ChannelPointsCustomRewardRedemptionUpdateEvent)
	onChannelPollBegin                          func(event ChannelPollBeginEvent)
	onChannelPollProgress                       func(event ChannelPollProgressEvent)
	onChannelPollEnd                            func(event ChannelPollEndEvent)
	onChannelPredictionBegin                    func(event ChannelPredictionBeginEvent)
	onChannelPredictionProgress                 func(event ChannelPredictionProgressEvent)
	onChannelPredictionLock                     func(event ChannelPredictionLockEvent)
	onChannelPredictionEnd                      func(event ChannelPredictionEndEvent)
	onChannelVIPAdd                             func(event ChannelVIPAddEvent)
	onChannelVIPRemove                          func(event ChannelVIPRemoveEvent)
	onCharityCampaignDonate                     func(event CharityCampaignDonateEvent)
	onCharityCampaignStart                      func(event CharityCampaignStartEvent)
	onCharityCampaignProgress                   func(event CharityCampaignProgressEvent)
	onCharityCampaignStop                       func(event CharityCampaignStopEvent)
	onConduitShardDisabled                      func(event ConduitShardDisabledEvent)
	onDropEntitlementGrant                      func(event DropEntitlementGrantEvent)
	onExtensionBitsTransactionCreate            func(event ExtensionBitsTransactionCreateEvent)
	onChannelGoalBegin                          func(event ChannelGoalBeginEvent)
	onChannelGoalProgress                       func(event ChannelGoalProgressEvent)
	onChannelGoalEnd                            func(event ChannelGoalEndEvent)
	onChannelHypeTrainBegin                     func(event ChannelHypeTrainBeginEvent)
	onChannelHypeTrainProgress                  func(event ChannelHypeTrainProgressEvent)
	onChannelHypeTrainEnd                       func(event ChannelHypeTrainEndEvent)
	onChannelShieldModeBegin                    func(event ChannelShieldModeBeginEvent)
	onChannelShieldModeEnd                      func(event ChannelShieldModeEndEvent)
	onChannelShoutOutCreate                     func(event ChannelShoutOutCreateEvent)
	onChannelShoutOutReceived                   func(event ChannelShoutOutReceivedEvent)
	onStreamOnline                              func(event StreamOnlineEvent)
	onStreamOffline                             func(event StreamOfflineEvent)
	onUserAuthorizationGrant                    func(event UserAuthorizationGrantEvent)
	onUserAuthorizationRevoke                   func(event UserAuthorizationRevokeEvent)
	onUserUpdate                                func(event UserUpdateEvent)
	onWhisperReceived                           func(event WhisperReceivedEvent)
	onChannelSuspiciousUserUpdate               func(event ChannelSuspiciousUserUpdateEvent)
}

func NewClient(crtPath, keyPath, secret, callback string) *Client {
	return &Client{crtPath: crtPath, keyPath: keyPath, secret: secret,
		secretBytes: []byte(secret), callback: callback, debug: false}
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
	if c.debug {
		c.onDebug("Received eventsub message")
	}
	body, err := io.ReadAll(req.Body)
	if err != nil {
		c.onError(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	signature := c.getHmac(req.Header.Get(headerId), req.Header.Get(headerTimestamp), body)
	if signature != req.Header.Get(headerSignature) {
		c.onError(errors.New("signatures do not match"))
		w.WriteHeader(http.StatusForbidden)
		return
	}
	var data Response
	if err := json.Unmarshal(body, &data); err != nil {
		c.onError(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	switch req.Header.Get(headerType) {
	case headerChallenge:
		if c.debug {
			c.onDebug("Received challenge message")
		}
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

func (c *Client) parseNotification(data Response) {
	if c.debug {
		c.onDebug(fmt.Sprintf("Received notification of type: %s", data.Subscription.Type))
	}
	switch data.Subscription.Type {
	case "automod.message.hold":
		if c.onAutomodMessageHold == nil {
			break
		}
		var e AutomodMessageHoldEvent
		err := json.Unmarshal(data.Event, &e)
		if err != nil {
			c.onError(fmt.Errorf("%s[automod.message.hold][%s]: %s", parseError, string(data.Event), err.Error()))
			break
		}
		c.onAutomodMessageHold(e)

	case "automod.message.update":
		if c.onAutomodMessageUpdate == nil {
			break
		}
		var e AutomodMessageUpdateEvent
		err := json.Unmarshal(data.Event, &e)
		if err != nil {
			c.onError(fmt.Errorf("%s[automod.message.update][%s]: %s", parseError, string(data.Event), err.Error()))
			break
		}
		c.onAutomodMessageUpdate(e)

	case "automod.settings.update":
		if c.onAutomodSettingsUpdate == nil {
			break
		}
		var e AutomodSettingsUpdateEvent
		err := json.Unmarshal(data.Event, &e)
		if err != nil {
			c.onError(fmt.Errorf("%s[automod.settings.update][%s]: %s", parseError, string(data.Event), err.Error()))
			break
		}
		c.onAutomodSettingsUpdate(e)

	case "automod.terms.update":
		if c.onAutomodTermsUpdate == nil {
			break
		}
		var e AutomodTermsUpdateEvent
		err := json.Unmarshal(data.Event, &e)
		if err != nil {
			c.onError(fmt.Errorf("%s[automod.terms.update][%s]: %s", parseError, string(data.Event), err.Error()))
			break
		}
		c.onAutomodTermsUpdate(e)

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
		if c.onChannelFollow == nil {
			break
		}
		var e ChannelFollowEvent
		err := json.Unmarshal(data.Event, &e)
		if err != nil {
			c.onError(fmt.Errorf("%s[channel.follow][%s]: %s", parseError, string(data.Event), err.Error()))
			break
		}
		c.onChannelFollow(e)

	case "channel.ad_break.begin":
		if c.onChannelAdBreakBegin == nil {
			break
		}
		var e ChannelAdBreakBeginEvent
		err := json.Unmarshal(data.Event, &e)
		if err != nil {
			c.onError(fmt.Errorf("%s[channel.ad_break.begin][%s]: %s", parseError, string(data.Event), err.Error()))
			break
		}
		c.onChannelAdBreakBegin(e)

	case "channel.chat.clear":
		if c.onChannelChatClear == nil {
			break
		}
		var e ChannelChatClearEvent
		err := json.Unmarshal(data.Event, &e)
		if err != nil {
			c.onError(fmt.Errorf("%s[channel.chat.clear][%s]: %s", parseError, string(data.Event), err.Error()))
			break
		}
		c.onChannelChatClear(e)

	case "channel.chat.clear_user_messages":
		if c.onChannelChatClearUserMessages == nil {
			break
		}
		var e ChannelChatClearUserMessagesEvent
		err := json.Unmarshal(data.Event, &e)
		if err != nil {
			c.onError(fmt.Errorf("%s[channel.chat.clear_user_messages][%s]: %s", parseError, string(data.Event), err.Error()))
			break
		}
		c.onChannelChatClearUserMessages(e)

	case "channel.chat.message":
		if c.onChannelChatMessage == nil {
			break
		}
		var e ChannelChatMessageEvent
		err := json.Unmarshal(data.Event, &e)
		if err != nil {
			c.onError(fmt.Errorf("%s[channel.chat.message][%s]: %s", parseError, string(data.Event), err.Error()))
			break
		}
		c.onChannelChatMessage(e)

	case "channel.chat.message_delete":
		if c.onChannelChatMessageDelete == nil {
			break
		}
		var e ChannelChatMessageDeleteEvent
		err := json.Unmarshal(data.Event, &e)
		if err != nil {
			c.onError(fmt.Errorf("%s[channel.chat.message_delete][%s]: %s", parseError, string(data.Event), err.Error()))
			break
		}
		c.onChannelChatMessageDelete(e)

	case "channel.chat.notification":
		if c.onChannelChatNotification == nil {
			break
		}
		var e ChannelChatNotificationEvent
		err := json.Unmarshal(data.Event, &e)
		if err != nil {
			c.onError(fmt.Errorf("%s[channel.chat.notification][%s]: %s", parseError, string(data.Event), err.Error()))
			break
		}
		c.onChannelChatNotification(e)

	case "channel.chat_settings.update":
		if c.onChannelChatSettingsUpdate == nil {
			break
		}
		var e ChannelChatSettingsUpdateEvent
		err := json.Unmarshal(data.Event, &e)
		if err != nil {
			c.onError(fmt.Errorf("%s[channel.chat_settings.update][%s]: %s", parseError, string(data.Event), err.Error()))
			break
		}
		c.onChannelChatSettingsUpdate(e)

	case "channel.chat.user_message_hold":
		if c.onChannelChatUserMessageHold == nil {
			break
		}
		var e ChannelChatUserMessageHoldEvent
		err := json.Unmarshal(data.Event, &e)
		if err != nil {
			c.onError(fmt.Errorf("%s[channel.chat.user_message_hold][%s]: %s", parseError, string(data.Event), err.Error()))
			break
		}
		c.onChannelChatUserMessageHold(e)

	case "channel.chat.user_message_update":
		if c.onChannelChatUserMessageUpdate == nil {
			break
		}
		var e ChannelChatUserMessageUpdateEvent
		err := json.Unmarshal(data.Event, &e)
		if err != nil {
			c.onError(fmt.Errorf("%s[channel.chat.user_message_update][%s]: %s", parseError, string(data.Event), err.Error()))
			break
		}
		c.onChannelChatUserMessageUpdate(e)

	case "channel.subscribe":
		if c.onChannelSubscribe == nil {
			break
		}
		var e ChannelSubscribeEvent
		err := json.Unmarshal(data.Event, &e)
		if err != nil {
			c.onError(fmt.Errorf("%s[channel.subscribe][%s]: %s", parseError, string(data.Event), err.Error()))
			break
		}
		c.onChannelSubscribe(e)

	case "channel.subscription.end":
		if c.onChannelSubscriptionEnd == nil {
			break
		}
		var e ChannelSubscriptionEndEvent
		err := json.Unmarshal(data.Event, &e)
		if err != nil {
			c.onError(fmt.Errorf("%s[channel.subscription.end][%s]: %s", parseError, string(data.Event), err.Error()))
			break
		}
		c.onChannelSubscriptionEnd(e)

	case "channel.subscription.gift":
		if c.onChannelSubscriptionGift == nil {
			break
		}
		var e ChannelSubscriptionGiftEvent
		err := json.Unmarshal(data.Event, &e)
		if err != nil {
			c.onError(fmt.Errorf("%s[channel.subscription.gift][%s]: %s", parseError, string(data.Event), err.Error()))
			break
		}
		c.onChannelSubscriptionGift(e)

	case "channel.subscription.message":
		if c.onChannelSubscriptionMessage == nil {
			break
		}
		var e ChannelSubscriptionMessageEvent
		err := json.Unmarshal(data.Event, &e)
		if err != nil {
			c.onError(fmt.Errorf("%s[channel.subscription.message][%s]: %s", parseError, string(data.Event), err.Error()))
			break
		}
		c.onChannelSubscriptionMessage(e)

	case "channel.cheer":
		if c.onChannelCheer == nil {
			break
		}
		var e ChannelCheerEvent
		err := json.Unmarshal(data.Event, &e)
		if err != nil {
			c.onError(fmt.Errorf("%s[channel.cheer][%s]: %s", parseError, string(data.Event), err.Error()))
			break
		}
		c.onChannelCheer(e)

	case "channel.raid":
		if c.onChannelRaid == nil {
			break
		}
		var e ChannelRaidEvent
		err := json.Unmarshal(data.Event, &e)
		if err != nil {
			c.onError(fmt.Errorf("%s[channel.raid][%s]: %s", parseError, string(data.Event), err.Error()))
			break
		}
		c.onChannelRaid(e)

	case "channel.ban":
		if c.onChannelBan == nil {
			break
		}
		var e ChannelBanEvent
		err := json.Unmarshal(data.Event, &e)
		if err != nil {
			c.onError(fmt.Errorf("%s[channel.ban][%s]: %s", parseError, string(data.Event), err.Error()))
			break
		}
		c.onChannelBan(e)

	case "channel.unban":
		if c.onChannelUnban == nil {
			break
		}
		var e ChannelUnbanEvent
		err := json.Unmarshal(data.Event, &e)
		if err != nil {
			c.onError(fmt.Errorf("%s[channel.unban][%s]: %s", parseError, string(data.Event), err.Error()))
			break
		}
		c.onChannelUnban(e)

	case "channel.unban_request.create":
		if c.onChannelUnbanRequestCreate == nil {
			break
		}
		var e ChannelUnbanRequestCreateEvent
		err := json.Unmarshal(data.Event, &e)
		if err != nil {
			c.onError(fmt.Errorf("%s[channel.unban_request.create][%s]: %s", parseError, string(data.Event), err.Error()))
			break
		}
		c.onChannelUnbanRequestCreate(e)

	case "channel.unban_request.resolve":
		if c.onChannelUnbanRequestResolve == nil {
			break
		}
		var e ChannelUnbanRequestResolveEvent
		err := json.Unmarshal(data.Event, &e)
		if err != nil {
			c.onError(fmt.Errorf("%s[channel.unban_request.resolve][%s]: %s", parseError, string(data.Event), err.Error()))
			break
		}
		c.onChannelUnbanRequestResolve(e)

	case "channel.moderate":
		if c.onChannelModerate == nil {
			break
		}
		var e ChannelModerateEvent
		err := json.Unmarshal(data.Event, &e)
		if err != nil {
			c.onError(fmt.Errorf("%s[channel.moderate][%s]: %s", parseError, string(data.Event), err.Error()))
			break
		}
		c.onChannelModerate(e)

	case "channel.moderator.add":
		if c.onChannelModeratorAdd == nil {
			break
		}
		var e ChannelModeratorAddEvent
		err := json.Unmarshal(data.Event, &e)
		if err != nil {
			c.onError(fmt.Errorf("%s[channel.moderator.add][%s]: %s", parseError, string(data.Event), err.Error()))
			break
		}
		c.onChannelModeratorAdd(e)

	case "channel.moderator.remove":
		if c.onChannelModeratorRemove == nil {
			break
		}
		var e ChannelModeratorRemoveEvent
		err := json.Unmarshal(data.Event, &e)
		if err != nil {
			c.onError(fmt.Errorf("%s[channel.moderator.remove][%s]: %s", parseError, string(data.Event), err.Error()))
			break
		}
		c.onChannelModeratorRemove(e)

	case "channel.guest_star_session.begin":
		if c.onChannelGuestStarSessionBegin == nil {
			break
		}
		var e ChannelGuestStarSessionBeginEvent
		err := json.Unmarshal(data.Event, &e)
		if err != nil {
			c.onError(fmt.Errorf("%s[channel.guest_star_session.begin][%s]: %s", parseError, string(data.Event), err.Error()))
			break
		}
		c.onChannelGuestStarSessionBegin(e)

	case "channel.guest_star_session.end":
		if c.onChannelGuestStarSessionEnd == nil {
			break
		}
		var e ChannelGuestStarSessionEndEvent
		err := json.Unmarshal(data.Event, &e)
		if err != nil {
			c.onError(fmt.Errorf("%s[channel.guest_star_session.end][%s]: %s", parseError, string(data.Event), err.Error()))
			break
		}
		c.onChannelGuestStarSessionEnd(e)

	case "channel.guest_star_guest.update":
		if c.onChannelGuestStarGuestUpdate == nil {
			break
		}
		var e ChannelGuestStarGuestUpdateEvent
		err := json.Unmarshal(data.Event, &e)
		if err != nil {
			c.onError(fmt.Errorf("%s[channel.guest_star_guest.update][%s]: %s", parseError, string(data.Event), err.Error()))
			break
		}
		c.onChannelGuestStarGuestUpdate(e)

	case "channel.guest_star_settings.update":
		if c.onChannelGuestStarSettingsUpdate == nil {
			break
		}
		var e ChannelGuestStarSettingsUpdateEvent
		err := json.Unmarshal(data.Event, &e)
		if err != nil {
			c.onError(fmt.Errorf("%s[channel.guest_star_settings.update][%s]: %s", parseError, string(data.Event), err.Error()))
			break
		}
		c.onChannelGuestStarSettingsUpdate(e)

	case "channel.channel_points_automatic_reward.add":
		if c.onChannelPointsAutomaticRewardRedemptionAdd == nil {
			break
		}
		var e ChannelPointsAutomaticRewardRedemptionAddEvent
		err := json.Unmarshal(data.Event, &e)
		if err != nil {
			c.onError(fmt.Errorf("%s[channel.channel_points_automatic_reward.add][%s]: %s", parseError, string(data.Event), err.Error()))
			break
		}
		c.onChannelPointsAutomaticRewardRedemptionAdd(e)

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
		if c.onChannelPollBegin == nil {
			break
		}
		var e ChannelPollBeginEvent
		err := json.Unmarshal(data.Event, &e)
		if err != nil {
			c.onError(fmt.Errorf("%s[channel.poll.begin][%s]: %s", parseError, string(data.Event), err.Error()))
			break
		}
		c.onChannelPollBegin(e)

	case "channel.poll.progress":
		if c.onChannelPollProgress == nil {
			break
		}
		var e ChannelPollProgressEvent
		err := json.Unmarshal(data.Event, &e)
		if err != nil {
			c.onError(fmt.Errorf("%s[channel.poll.progress][%s]: %s", parseError, string(data.Event), err.Error()))
			break
		}
		c.onChannelPollProgress(e)

	case "channel.poll.end":
		if c.onChannelPollEnd == nil {
			break
		}
		var e ChannelPollEndEvent
		err := json.Unmarshal(data.Event, &e)
		if err != nil {
			c.onError(fmt.Errorf("%s[channel.poll.end][%s]: %s", parseError, string(data.Event), err.Error()))
			break
		}
		c.onChannelPollEnd(e)

	case "channel.prediction.begin":
		if c.onChannelPredictionBegin == nil {
			break
		}
		var e ChannelPredictionBeginEvent
		err := json.Unmarshal(data.Event, &e)
		if err != nil {
			c.onError(fmt.Errorf("%s[channel.prediction.begin][%s]: %s", parseError, string(data.Event), err.Error()))
			break
		}
		c.onChannelPredictionBegin(e)

	case "channel.prediction.progress":
		if c.onChannelPredictionProgress == nil {
			break
		}
		var e ChannelPredictionProgressEvent
		err := json.Unmarshal(data.Event, &e)
		if err != nil {
			c.onError(fmt.Errorf("%s[channel.prediction.progress][%s]: %s", parseError, string(data.Event), err.Error()))
			break
		}
		c.onChannelPredictionProgress(e)

	case "channel.prediction.lock":
		if c.onChannelPredictionLock == nil {
			break
		}
		var e ChannelPredictionLockEvent
		err := json.Unmarshal(data.Event, &e)
		if err != nil {
			c.onError(fmt.Errorf("%s[channel.prediction.lock][%s]: %s", parseError, string(data.Event), err.Error()))
			break
		}
		c.onChannelPredictionLock(e)

	case "channel.prediction.end":
		if c.onChannelPredictionEnd == nil {
			break
		}
		var e ChannelPredictionEndEvent
		err := json.Unmarshal(data.Event, &e)
		if err != nil {
			c.onError(fmt.Errorf("%s[channel.prediction.end][%s]: %s", parseError, string(data.Event), err.Error()))
			break
		}
		c.onChannelPredictionEnd(e)

	case "channel.vip.add":
		if c.onChannelVIPAdd == nil {
			break
		}
		var e ChannelVIPAddEvent
		err := json.Unmarshal(data.Event, &e)
		if err != nil {
			c.onError(fmt.Errorf("%s[channel.vip.add][%s]: %s", parseError, string(data.Event), err.Error()))
			break
		}
		c.onChannelVIPAdd(e)

	case "channel.vip.remove":
		if c.onChannelVIPRemove == nil {
			break
		}
		var e ChannelVIPRemoveEvent
		err := json.Unmarshal(data.Event, &e)
		if err != nil {
			c.onError(fmt.Errorf("%s[channel.vip.remove][%s]: %s", parseError, string(data.Event), err.Error()))
			break
		}
		c.onChannelVIPRemove(e)

	case "channel.charity_campaign.donate":
		if c.onCharityCampaignDonate == nil {
			break
		}
		var e CharityCampaignDonateEvent
		err := json.Unmarshal(data.Event, &e)
		if err != nil {
			c.onError(fmt.Errorf("%s[channel.charity_campaign.donate][%s]: %s", parseError, string(data.Event), err.Error()))
			break
		}
		c.onCharityCampaignDonate(e)

	case "channel.charity_campaign.start":
		if c.onCharityCampaignStart == nil {
			break
		}
		var e CharityCampaignStartEvent
		err := json.Unmarshal(data.Event, &e)
		if err != nil {
			c.onError(fmt.Errorf("%s[channel.charity_campaign.start][%s]: %s", parseError, string(data.Event), err.Error()))
			break
		}
		c.onCharityCampaignStart(e)

	case "channel.charity_campaign.progress":
		if c.onCharityCampaignProgress == nil {
			break
		}
		var e CharityCampaignProgressEvent
		err := json.Unmarshal(data.Event, &e)
		if err != nil {
			c.onError(fmt.Errorf("%s[channel.charity_campaign.progress][%s]: %s", parseError, string(data.Event), err.Error()))
			break
		}
		c.onCharityCampaignProgress(e)

	case "channel.charity_campaign.stop":
		if c.onCharityCampaignStop == nil {
			break
		}
		var e CharityCampaignStopEvent
		err := json.Unmarshal(data.Event, &e)
		if err != nil {
			c.onError(fmt.Errorf("%s[channel.charity_campaign.stop][%s]: %s", parseError, string(data.Event), err.Error()))
			break
		}
		c.onCharityCampaignStop(e)

	case "conduit.shard.disabled":
		if c.onConduitShardDisabled == nil {
			break
		}
		var e ConduitShardDisabledEvent
		err := json.Unmarshal(data.Event, &e)
		if err != nil {
			c.onError(fmt.Errorf("%s[conduit.shard.disabled][%s]: %s", parseError, string(data.Event), err.Error()))
			break
		}
		c.onConduitShardDisabled(e)

	case "drop.entitlement.grant":
		if c.onDropEntitlementGrant == nil {
			break
		}
		var e DropEntitlementGrantEvent
		err := json.Unmarshal(data.Events, &e)
		if err != nil {
			c.onError(fmt.Errorf("%s[drop.entitlement.grant][%s]: %s", parseError, string(data.Events), err.Error()))
			break
		}
		c.onDropEntitlementGrant(e)

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
		if c.onChannelGoalBegin == nil {
			break
		}
		var e ChannelGoalBeginEvent
		err := json.Unmarshal(data.Event, &e)
		if err != nil {
			c.onError(fmt.Errorf("%s[channel.goal.begin][%s]: %s", parseError, string(data.Event), err.Error()))
			break
		}
		c.onChannelGoalBegin(e)

	case "channel.goal.progress":
		if c.onChannelGoalProgress == nil {
			break
		}
		var e ChannelGoalProgressEvent
		err := json.Unmarshal(data.Event, &e)
		if err != nil {
			c.onError(fmt.Errorf("%s[channel.goal.progress][%s]: %s", parseError, string(data.Event), err.Error()))
			break
		}
		c.onChannelGoalProgress(e)

	case "channel.goal.end":
		if c.onChannelGoalEnd == nil {
			break
		}
		var e ChannelGoalEndEvent
		err := json.Unmarshal(data.Event, &e)
		if err != nil {
			c.onError(fmt.Errorf("%s[channel.goal.end][%s]: %s", parseError, string(data.Event), err.Error()))
			break
		}
		c.onChannelGoalEnd(e)

	case "channel.hype_train.begin":
		if c.onChannelHypeTrainBegin == nil {
			break
		}
		var e ChannelHypeTrainBeginEvent
		err := json.Unmarshal(data.Event, &e)
		if err != nil {
			c.onError(fmt.Errorf("%s[channel.hype_train.begin][%s]: %s", parseError, string(data.Event), err.Error()))
			break
		}
		c.onChannelHypeTrainBegin(e)

	case "channel.hype_train.progress":
		if c.onChannelHypeTrainProgress == nil {
			break
		}
		var e ChannelHypeTrainProgressEvent
		err := json.Unmarshal(data.Event, &e)
		if err != nil {
			c.onError(fmt.Errorf("%s[channel.hype_train.progress][%s]: %s", parseError, string(data.Event), err.Error()))
			break
		}
		c.onChannelHypeTrainProgress(e)

	case "channel.hype_train.end":
		if c.onChannelHypeTrainEnd == nil {
			break
		}
		var e ChannelHypeTrainEndEvent
		err := json.Unmarshal(data.Event, &e)
		if err != nil {
			c.onError(fmt.Errorf("%s[channel.hype_train.end][%s]: %s", parseError, string(data.Event), err.Error()))
			break
		}
		c.onChannelHypeTrainEnd(e)

	case "channel.shield_mode.begin":
		if c.onChannelShieldModeBegin == nil {
			break
		}
		var e ChannelShieldModeBeginEvent
		err := json.Unmarshal(data.Event, &e)
		if err != nil {
			c.onError(fmt.Errorf("%s[channel.shield_mode.begin][%s]: %s", parseError, string(data.Event), err.Error()))
			break
		}
		c.onChannelShieldModeBegin(e)

	case "channel.shield_mode.end":
		if c.onChannelShieldModeEnd == nil {
			break
		}
		var e ChannelShieldModeEndEvent
		err := json.Unmarshal(data.Event, &e)
		if err != nil {
			c.onError(fmt.Errorf("%s[channel.shield_mode.end][%s]: %s", parseError, string(data.Event), err.Error()))
			break
		}
		c.onChannelShieldModeEnd(e)

	case "channel.shoutout.create":
		if c.onChannelShoutOutCreate == nil {
			break
		}
		var e ChannelShoutOutCreateEvent
		err := json.Unmarshal(data.Event, &e)
		if err != nil {
			c.onError(fmt.Errorf("%s[channel.shoutout.create][%s]: %s", parseError, string(data.Event), err.Error()))
			break
		}
		c.onChannelShoutOutCreate(e)

	case "channel.shoutout.receive":
		if c.onChannelShoutOutReceived == nil {
			break
		}
		var e ChannelShoutOutReceivedEvent
		err := json.Unmarshal(data.Event, &e)
		if err != nil {
			c.onError(fmt.Errorf("%s[channel.shoutout.receive][%s]: %s", parseError, string(data.Event), err.Error()))
			break
		}
		c.onChannelShoutOutReceived(e)

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

	case "user.whisper.message":
		if c.onWhisperReceived == nil {
			break
		}
		var e WhisperReceivedEvent
		err := json.Unmarshal(data.Event, &e)
		if err != nil {
			c.onError(fmt.Errorf("%s[user.whisper.message][%s]: %s", parseError, string(data.Event), err.Error()))
			break
		}
		c.onWhisperReceived(e)
	case "channel.suspicious_user.update":
		if c.onChannelSuspiciousUserUpdate == nil {
			break
		}
		var e ChannelSuspiciousUserUpdateEvent
		err := json.Unmarshal(data.Event, &e)
		if err != nil {
			c.onError(fmt.Errorf("%s[channel.suspicious_user.update][%s]: %s", parseError, string(data.Event), err.Error()))
			break
		}
		c.onChannelSuspiciousUserUpdate(e)

	default:
		c.onError(fmt.Errorf("%s[default][%s]: Unable to parse event", parseError, string(data.Event)))
	}
}

func (c *Client) OnError(f func(err error)) {
	c.onError = f
}

func (c *Client) OnRevoked(f func(sub Subscription)) {
	c.onRevoked = f
}

func (c *Client) OnDebug(f func(msg string)) {
	c.onDebug = f
}

func (c *Client) OnAutomodMessageHold(f func(event AutomodMessageHoldEvent)) {
	c.onAutomodMessageHold = f
}

func (c *Client) OnAutomodMessageUpdate(f func(event AutomodMessageUpdateEvent)) {
	c.onAutomodMessageUpdate = f
}

func (c *Client) OnAutomodSettingsUpdate(f func(event AutomodSettingsUpdateEvent)) {
	c.onAutomodSettingsUpdate = f
}

func (c *Client) OnAutomodTermsUpdate(f func(event AutomodTermsUpdateEvent)) {
	c.onAutomodTermsUpdate = f
}

func (c *Client) OnChannelUpdate(f func(event ChannelUpdateEvent)) {
	c.onChannelUpdate = f
}

func (c *Client) OnChannelFollow(f func(event ChannelFollowEvent)) {
	c.onChannelFollow = f
}

func (c *Client) OnChannelAdBreakBegin(f func(event ChannelAdBreakBeginEvent)) {
	c.onChannelAdBreakBegin = f
}

func (c *Client) OnChannelChatClear(f func(event ChannelChatClearEvent)) {
	c.onChannelChatClear = f
}

func (c *Client) OnChannelChatClearUserMessages(f func(event ChannelChatClearUserMessagesEvent)) {
	c.onChannelChatClearUserMessages = f
}

func (c *Client) OnChannelChatMessage(f func(event ChannelChatMessageEvent)) {
	c.onChannelChatMessage = f
}

func (c *Client) OnChannelChatMessageDelete(f func(event ChannelChatMessageDeleteEvent)) {
	c.onChannelChatMessageDelete = f
}

func (c *Client) OnChannelChatNotification(f func(event ChannelChatNotificationEvent)) {
	c.onChannelChatNotification = f
}

func (c *Client) OnChannelChatSettingsUpdate(f func(event ChannelChatSettingsUpdateEvent)) {
	c.onChannelChatSettingsUpdate = f
}

func (c *Client) OnChannelChatUserMessageHold(f func(event ChannelChatUserMessageHoldEvent)) {
	c.onChannelChatUserMessageHold = f
}

func (c *Client) OnChannelChatUserMessageUpdate(f func(event ChannelChatUserMessageUpdateEvent)) {
	c.onChannelChatUserMessageUpdate = f
}

func (c *Client) OnChannelSubscribe(f func(event ChannelSubscribeEvent)) {
	c.onChannelSubscribe = f
}

func (c *Client) OnChannelSubscriptionEnd(f func(event ChannelSubscriptionEndEvent)) {
	c.onChannelSubscriptionEnd = f
}

func (c *Client) OnChannelSubscriptionGift(f func(event ChannelSubscriptionGiftEvent)) {
	c.onChannelSubscriptionGift = f
}

func (c *Client) OnChannelSubscriptionMessage(f func(event ChannelSubscriptionMessageEvent)) {
	c.onChannelSubscriptionMessage = f
}

func (c *Client) OnChannelCheer(f func(event ChannelCheerEvent)) {
	c.onChannelCheer = f
}

func (c *Client) OnChannelRaid(f func(event ChannelRaidEvent)) {
	c.onChannelRaid = f
}

func (c *Client) OnChannelBan(f func(event ChannelBanEvent)) {
	c.onChannelBan = f
}

func (c *Client) OnChannelUnban(f func(event ChannelUnbanEvent)) {
	c.onChannelUnban = f
}

func (c *Client) OnChannelUnbanRequestCreate(f func(event ChannelUnbanRequestCreateEvent)) {
	c.onChannelUnbanRequestCreate = f
}

func (c *Client) OnChannelUnbanRequestResolve(f func(event ChannelUnbanRequestResolveEvent)) {
	c.onChannelUnbanRequestResolve = f
}

func (c *Client) OnChannelModerate(f func(event ChannelModerateEvent)) {
	c.onChannelModerate = f
}

func (c *Client) OnChannelModeratorAdd(f func(event ChannelModeratorAddEvent)) {
	c.onChannelModeratorAdd = f
}

func (c *Client) OnChannelModeratorRemove(f func(event ChannelModeratorRemoveEvent)) {
	c.onChannelModeratorRemove = f
}

func (c *Client) OnChannelGuestStarSessionBegin(f func(event ChannelGuestStarSessionBeginEvent)) {
	c.onChannelGuestStarSessionBegin = f
}

func (c *Client) OnChannelGuestStarSessionEnd(f func(event ChannelGuestStarSessionEndEvent)) {
	c.onChannelGuestStarSessionEnd = f
}

func (c *Client) OnChannelGuestStarGuestUpdate(f func(event ChannelGuestStarGuestUpdateEvent)) {
	c.onChannelGuestStarGuestUpdate = f
}

func (c *Client) OnChannelGuestStarSettingsUpdate(f func(event ChannelGuestStarSettingsUpdateEvent)) {
	c.onChannelGuestStarSettingsUpdate = f
}

func (c *Client) OnChannelPointsAutomaticRewardRedemptionAdd(f func(event ChannelPointsAutomaticRewardRedemptionAddEvent)) {
	c.onChannelPointsAutomaticRewardRedemptionAdd = f
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

func (c *Client) OnChannelPollBegin(f func(event ChannelPollBeginEvent)) {
	c.onChannelPollBegin = f
}

func (c *Client) OnChannelPollProgress(f func(event ChannelPollProgressEvent)) {
	c.onChannelPollProgress = f
}

func (c *Client) OnChannelPollEnd(f func(event ChannelPollEndEvent)) {
	c.onChannelPollEnd = f
}

func (c *Client) OnChannelPredictionBegin(f func(event ChannelPredictionBeginEvent)) {
	c.onChannelPredictionBegin = f
}

func (c *Client) OnChannelPredictionProgress(f func(event ChannelPredictionProgressEvent)) {
	c.onChannelPredictionProgress = f
}

func (c *Client) OnChannelPredictionLock(f func(event ChannelPredictionLockEvent)) {
	c.onChannelPredictionLock = f
}

func (c *Client) OnChannelPredictionEnd(f func(event ChannelPredictionEndEvent)) {
	c.onChannelPredictionEnd = f
}

func (c *Client) OnChannelVipAdd(f func(event ChannelVIPAddEvent)) {
	c.onChannelVIPAdd = f
}

func (c *Client) OnChannelVipRemove(f func(event ChannelVIPRemoveEvent)) {
	c.onChannelVIPRemove = f
}

func (c *Client) OnCharityCampaignDonate(f func(event CharityCampaignDonateEvent)) {
	c.onCharityCampaignDonate = f
}

func (c *Client) OnCharityCampaignStart(f func(event CharityCampaignStartEvent)) {
	c.onCharityCampaignStart = f
}

func (c *Client) OnChannelCharityCampaignProgress(f func(event CharityCampaignProgressEvent)) {
	c.onCharityCampaignProgress = f
}

func (c *Client) OnCharityCampaignStop(f func(event CharityCampaignStopEvent)) {
	c.onCharityCampaignStop = f
}

func (c *Client) OnConduitShardDisabled(f func(event ConduitShardDisabledEvent)) {
	c.onConduitShardDisabled = f
}

func (c *Client) OnDropEntitlementGrant(f func(event DropEntitlementGrantEvent)) {
	c.onDropEntitlementGrant = f
}

func (c *Client) OnExtensionBitsTransactionCreate(f func(event ExtensionBitsTransactionCreateEvent)) {
	c.onExtensionBitsTransactionCreate = f
}

func (c *Client) OnChannelGoalBegin(f func(event ChannelGoalBeginEvent)) {
	c.onChannelGoalBegin = f
}

func (c *Client) OnChannelGoalProgress(f func(event ChannelGoalProgressEvent)) {
	c.onChannelGoalProgress = f
}

func (c *Client) OnChannelGoalEnd(f func(event ChannelGoalEndEvent)) {
	c.onChannelGoalEnd = f
}

func (c *Client) OnChannelHypeTrainBegin(f func(event ChannelHypeTrainBeginEvent)) {
	c.onChannelHypeTrainBegin = f
}

func (c *Client) OnChannelHypeTrainProgress(f func(event ChannelHypeTrainProgressEvent)) {
	c.onChannelHypeTrainProgress = f
}

func (c *Client) OnChannelHypeTrainEnd(f func(event ChannelHypeTrainEndEvent)) {
	c.onChannelHypeTrainEnd = f
}

func (c *Client) OnChannelShieldModeBegin(f func(event ChannelShieldModeBeginEvent)) {
	c.onChannelShieldModeBegin = f
}

func (c *Client) OnChannelShieldModeEnd(f func(event ChannelShieldModeEndEvent)) {
	c.onChannelShieldModeEnd = f
}

func (c *Client) OnChannelShoutOutCreate(f func(event ChannelShoutOutCreateEvent)) {
	c.onChannelShoutOutCreate = f
}

func (c *Client) OnChannelShoutOutReceived(f func(event ChannelShoutOutReceivedEvent)) {
	c.onChannelShoutOutReceived = f
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

func (c *Client) OnWhisperReceived(f func(event WhisperReceivedEvent)) {
	c.onWhisperReceived = f
}

func (c *Client) OnChannelSuspiciousUserUpdate(f func(event ChannelSuspiciousUserUpdateEvent)) {
	c.onChannelSuspiciousUserUpdate = f
}

func (c *Client) SetDebug(b bool) {
	c.debug = b
}
