<img src="https://i.imgur.com/fG94ZeC.png" width="300" height="100">

---

**A bot designed to run the economy of your Discord server with very little configuration required. You can invite the bot [here](https://discordapp.com/oauth2/authorize?client_id=660941659613298689&scope=bot&permissions=1546775744).**

Cube is a bot which is designed to be able to be lightweight, stable and give the user a bot which can run the economy. As of right now, the bot is in beta stages meaning it is constantly being developed. The bot uses Redis for its database, allowing for fast access to information and, therefore, lightning fast response times (usually well under 1ms). We don't want users waiting after all, this bot is meant to be fun!

We accept pull requests for Cube as long as they meet the following criteria:
- **Are fun features:** We are **NOT** a moderation bot, features should be fun for the end user! If you want a moderation bot, we suggest [Auttaja](https://auttaja.io/).
- **Are written well and meet usual Go coding protocol:** We don't want code which will cause instability or performance hits for the bot.
- **Puts anything that takes a lot of time into threads:** Go is a multi-threaded language, we do not want to hang one thread for a users request.
- **Is economy/XP based if it's a feature:** We are an economy bot!

Note that Cube is automatically deployed from the master branch. Therefore, any PR's there will be much more stringent.

## Setup
After building Cube, please note that credits.json needs to be in the CWD which you are running Cube from. If you are using Docker, this is sorted for you!

Additionally, the following environment variables need to be set:
- `TOKEN` - The Discord bot token which is in use.
- `DEFAULT_PREFIX` - The default prefix for the bot to use.
- `REDIS_ADDR` - The address for the Redis host.
- `SENTRY_DSN` - The Sentry DSN which will be used.

The following environment variables are optional:
- `REDIS_PASSWORD` - The password for the Redis instance. This is obviously required if your Redis instance is password protected.
- `SHARD_COUNT` - The amount of shards.
- `SHARD_ID` - The ID of the current shard.
- `POD_NAME` - This overrides `SHARD_ID` and is designed for use with [Marver](https://github.com/Auttaja-OpenSource/Marver) in a Kubernetes cluster.
