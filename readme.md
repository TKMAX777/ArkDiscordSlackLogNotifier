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

```
wget https://github.com/TKMAX777/ArkDiscordSlackLogNotifier/releases/latest/download/ArkDiscordSlackLogNotifier_Linux_x86_64.tar.gz
tar -xzvf ArkDiscordSlackLogNotifier_Linux_x86_64.tar.gz
```

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
SLACK_SEND_KILLED=yes/no
SLACK_SEND_OTHER=yes/no
SLACK_SEND_ALL=yes/no

SLACK_JOIN_AND_LEFT_STATE_EMOJI=:tools: 
SLACK_JOIN_EMOJI=:revolving_heart:
SLACK_KILLED_EMOJI=
SLACK_LEFT_EMOJI=:wave: 
SLACK_OTHER_EMOJI=

DISCORD_HOOK_URI=
DISCORD_AVATER_URI=
DISCORD_USERNAME=

DISCORD_SEND_JOIN_AND_LEFT_STATE=yes/no
DISCORD_SEND_JOIN_AND_LEFT=yes/no
DISCORD_SEND_KILLED=yes/no
DISCORD_SEND_OTHER=yes/no
DISCORD_SEND_ALL=yes/no

DISCORD_JOIN_AND_LEFT_STATE_EMOJI=:tools: 
DISCORD_JOIN_EMOJI=:revolving_heart:
DISCORD_KILLED_EMOJI=
DISCORD_LEFT_EMOJI=:wave: 
DISCORD_OTHER_EMOJI=
```
