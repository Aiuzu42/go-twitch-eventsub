# Go Twitch EventSub

A Go library for handling Twitch EventSub webhook notifications.

## Installation

```bash
go get github.com/Aiuzu42/go-twitch-eventsub/v2
```

## Usage

```go
package main

import (
    "log"
    "net/http"
    twitcheventsub "github.com/Aiuzu42/go-twitch-eventsub/v2"
)

func main() {
    client := twitcheventsub.NewClient("your-secret", "https://your-callback-url.com")
    
    // Set up error handler
    client.OnError(func(err error) {
        log.Println("Error:", err)
    })
    
    // Set up event handlers
    client.OnChannelFollow(func(event twitcheventsub.ChannelFollowEvent) {
        log.Printf("%s followed %s\n", event.UserName, event.BroadcasterUserName)
    })
    
    // Handle webhooks
    http.HandleFunc("/webhook", client.HandleEvent)
    http.ListenAndServe(":8080", nil)
}
```

## Supported Events

Below is a complete list of Twitch EventSub subscription types. Events marked with ✅ are implemented in this library.

### Automod Events

- ✅ `automod.message.hold` (v1)
- ⬜ `automod.message.hold` (v2)
- ✅ `automod.message.update` (v1)
- ⬜ `automod.message.update` (v2)
- ✅ `automod.settings.update` (v1)
- ✅ `automod.terms.update` (v1)

### Channel Events

- ✅ `channel.update` (v2)
- ✅ `channel.follow` (v2)
- ✅ `channel.ad_break.begin` (v1)
- ✅ `channel.bits.use` (v1)

### Chat Events

- ✅ `channel.chat.clear` (v1)
- ✅ `channel.chat.clear_user_messages` (v1)
- ✅ `channel.chat.message` (v1)
- ✅ `channel.chat.message_delete` (v1)
- ✅ `channel.chat.notification` (v1)
- ✅ `channel.chat_settings.update` (v1)
- ✅ `channel.chat.user_message_hold` (v1)
- ✅ `channel.chat.user_message_update` (v1)

### Shared Chat Events

- ⬜ `channel.shared_chat.begin` (v1)
- ⬜ `channel.shared_chat.update` (v1)
- ⬜ `channel.shared_chat.end` (v1)

### Subscription Events

- ✅ `channel.subscribe` (v1)
- ✅ `channel.subscription.end` (v1)
- ✅ `channel.subscription.gift` (v1)
- ✅ `channel.subscription.message` (v1)

### Bits & Cheers

- ✅ `channel.cheer` (v1)

### Raids

- ✅ `channel.raid` (v1)

### Moderation Events

- ✅ `channel.ban` (v1)
- ✅ `channel.unban` (v1)
- ✅ `channel.unban_request.create` (v1)
- ✅ `channel.unban_request.resolve` (v1)
- ✅ `channel.moderate` (v1)
- ✅ `channel.moderate` (v2)
- ✅ `channel.moderator.add` (v1)
- ✅ `channel.moderator.remove` (v1)
- ✅ `channel.suspicious_user.message` (v1)
- ✅ `channel.suspicious_user.update` (v1)
- ✅ `channel.warning.acknowledge` (v1)
- ✅ `channel.warning.send` (v1)

### Guest Star Events

- ✅ `channel.guest_star_session.begin` (beta)
- ✅ `channel.guest_star_session.end` (beta)
- ✅ `channel.guest_star_guest.update` (beta)
- ✅ `channel.guest_star_settings.update` (beta)

### Channel Points Events

- ✅ `channel.channel_points_automatic_reward_redemption.add` (v1)
- ⬜ `channel.channel_points_automatic_reward_redemption.add` (v2)
- ✅ `channel.channel_points_custom_reward.add` (v1)
- ✅ `channel.channel_points_custom_reward.update` (v1)
- ✅ `channel.channel_points_custom_reward.remove` (v1)
- ✅ `channel.channel_points_custom_reward_redemption.add` (v1)
- ✅ `channel.channel_points_custom_reward_redemption.update` (v1)

### Power-ups Events

- ✅ `channel.custom_power_up_redemption.add` (beta)

### Poll Events

- ✅ `channel.poll.begin` (v1)
- ✅ `channel.poll.progress` (v1)
- ✅ `channel.poll.end` (v1)

### Prediction Events

- ✅ `channel.prediction.begin` (v1)
- ✅ `channel.prediction.progress` (v1)
- ✅ `channel.prediction.lock` (v1)
- ✅ `channel.prediction.end` (v1)

### VIP Events

- ✅ `channel.vip.add` (v1)
- ✅ `channel.vip.remove` (v1)

### Charity Events

- ✅ `channel.charity_campaign.donate` (v1)
- ✅ `channel.charity_campaign.start` (v1)
- ✅ `channel.charity_campaign.progress` (v1)
- ✅ `channel.charity_campaign.stop` (v1)

### Conduit Events

- ✅ `conduit.shard.disabled` (v1)

### Drop Events

- ✅ `drop.entitlement.grant` (v1)

### Extension Events

- ✅ `extension.bits_transaction.create` (v1)

### Goal Events

- ✅ `channel.goal.begin` (v1)
- ✅ `channel.goal.progress` (v1)
- ✅ `channel.goal.end` (v1)

### Hype Train Events

- ✅ `channel.hype_train.begin` (v2)
- ✅ `channel.hype_train.progress` (v2)
- ✅ `channel.hype_train.end` (v2)

### Shield Mode Events

- ✅ `channel.shield_mode.begin` (v1)
- ✅ `channel.shield_mode.end` (v1)

### Shoutout Events

- ✅ `channel.shoutout.create` (v1)
- ✅ `channel.shoutout.receive` (v1)

### Stream Events

- ✅ `stream.online` (v1)
- ✅ `stream.offline` (v1)

### User Events

- ✅ `user.authorization.grant` (v1)
- ✅ `user.authorization.revoke` (v1)
- ✅ `user.update` (v1)
- ✅ `user.whisper.message` (v1)

## Event Implementation Status

**Total Events**: 85+  
**Implemented**: 79  
**Not Implemented**: 6

### Not Yet Implemented

- `automod.message.hold` (v2)
- `automod.message.update` (v2)
- `channel.channel_points_automatic_reward_redemption.add` (v2)
- `channel.shared_chat.begin` (v1)
- `channel.shared_chat.update` (v1)
- `channel.shared_chat.end` (v1)

## Contributing

Contributions are welcome! If you'd like to add support for missing events, please submit a pull request.

## License

See [LICENSE](LICENSE) for details.

## Resources

- [Twitch EventSub Documentation](https://dev.twitch.tv/docs/eventsub/)
- [EventSub Subscription Types](https://dev.twitch.tv/docs/eventsub/eventsub-subscription-types/)
