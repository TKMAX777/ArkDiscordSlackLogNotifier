# ArkDiscordSlackLogNotifier

## About

This program monitors the Ark server logs and notifies Discord or Slack when players enter or leave the room, etc.

## Content

<!-- TOC -->

- [ArkDiscordSlackLogNotifier](#arkdiscordslacklognotifier)
    - [About](#about)
    - [Content](#content)
    - [Install](#install)
    - [Setup](#setup)
        - [Environment Variables](#environment-variables)

<!-- /TOC -->

## Install


## Setup

### Environment Variables

```
ARK_LOG_DIRECTORY=/path/to/ShooterGame/Saved/Logs

SLACK_TOKEN=xoxb-********

SLACK_CHANNEL_ID=C******
SLACK_AVATER_URI=http://********.png
SLACK_USERNAME=ARK SERVER

SLACK_SEND_JOIN_AND_LEFT_STATE=yes/no
SLACK_SEND_JOIN_AND_LEFT=yes/no
SLACK_SEND_OTHER=yes/no
SLACK_SEND_ALL=yes/no

SLACK_JOIN_AND_LEFT_STATE_EMOJI=
SLACK_JOIN_EMOJI=
SLACK_LEFT_EMOJI=
SLACK_OTHER_EMOJI=

DISCORD_HOOK_URI=
DISCORD_AVATER_URI=
DISCORD_USERNAME=

DISCORD_SEND_JOIN_AND_LEFT_STATE=yes/no
DISCORD_SEND_JOIN_AND_LEFT=yes/no
DISCORD_SEND_OTHER=yes/no
DISCORD_SEND_ALL=yes/no

DISCORD_JOIN_AND_LEFT_STATE_EMOJI=
DISCORD_JOIN_EMOJI=
DISCORD_LEFT_EMOJI=
DISCORD_OTHER_EMOJI=
```