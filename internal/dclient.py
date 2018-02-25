# Cube. Copyright (C) Jake Gealer <jake@gealer.email> 2017-2018.
# This Source Code Form is subject to the terms of the Mozilla Public
# License, v. 2.0. If a copy of the MPL was not distributed with this
# file, You can obtain one at http://mozilla.org/MPL/2.0/.

from copy import copy
from internal.handlers import cmd_handler, server_restriction_handler
import discord
# Imports go here.

def load_dclient(app):

    @app.dclient.event
    async def on_ready():
        if "on_ready" in app.discord_callbacks:
            for i in app.discord_callbacks["on_ready"]:
                try:
                    await i(app)
                except Exception as e:
                    app.logger.error(e)
    # Defines on_ready.

    @app.dclient.event
    async def on_message(msg):
        if not msg.author.bot and not isinstance(msg.channel, discord.abc.PrivateChannel):
            if "on_message" in app.discord_callbacks:
                for i in app.discord_callbacks["on_message"]:
                    try:
                        app_ctx = copy(app)
                        app_ctx.message = msg
                        await i(app_ctx)
                    except Exception as e:
                        app.logger.error(e)
            await cmd_handler(app, msg)
    # Defines on_message.

    @app.dclient.event
    async def on_message_delete(msg):
        if not msg.author.bot:
            if "on_message_delete" in app.discord_callbacks:
                for i in app.discord_callbacks["on_message_delete"]:
                    try:
                        app_ctx = copy(app)
                        app_ctx.message = msg
                        await i(app_ctx)
                    except Exception as e:
                        app.logger.error(e)
    # Defines on_message_delete.

    @app.dclient.event
    async def on_reaction_add(reaction, user):
        if "on_reaction_add" in app.discord_callbacks:
            for i in app.discord_callbacks["on_reaction_add"]:
                try:
                    app_ctx = copy(app)
                    app_ctx.message = reaction.message
                    await i(app_ctx, reaction, user)
                except Exception as e:
                    app.logger.error(e)
    # Defines on_reaction_add.

    @app.dclient.event
    async def on_reaction_remove(reaction, user):
        if "on_reaction_remove" in app.discord_callbacks:
            for i in app.discord_callbacks["on_reaction_remove"]:
                try:
                    app_ctx = copy(app)
                    app_ctx.message = reaction.message
                    await i(app_ctx, reaction, user)
                except Exception as e:
                    app.logger.error(e)
    # Defines on_reaction_remove.

    @app.dclient.event
    async def on_channel_delete(channel):
        if "on_channel_delete" in app.discord_callbacks:
            for i in app.discord_callbacks["on_channel_delete"]:
                try:
                    await i(app, channel)
                except Exception as e:
                    app.logger.error(e)
    # Defines on_channel_delete.

    @app.dclient.event
    async def on_channel_create(channel):
        if "on_channel_create" in app.discord_callbacks:
            for i in app.discord_callbacks["on_channel_create"]:
                try:
                    await i(app, channel)
                except Exception as e:
                    app.logger.error(e)
    # Defines on_channel_create.

    @app.dclient.event
    async def on_channel_update(before, after):
        if "on_channel_update" in app.discord_callbacks:
            for i in app.discord_callbacks["on_channel_update"]:
                try:
                    await i(app, before, after)
                except Exception as e:
                    app.logger.error(e)
    # Defines on_channel_update.

    @app.dclient.event
    async def on_member_join(member):
        if "on_member_join" in app.discord_callbacks:
            for i in app.discord_callbacks["on_member_join"]:
                try:
                    await i(app, member)
                except Exception as e:
                    app.logger.error(e)
    # Defines on_member_join.

    @app.dclient.event
    async def on_member_remove(member):
        if "on_member_remove" in app.discord_callbacks:
            for i in app.discord_callbacks["on_member_remove"]:
                try:
                    await i(app, member)
                except Exception as e:
                    app.logger.error(e)
    # Defines on_member_remove.

    @app.dclient.event
    async def on_member_update(before, after):
        if "on_member_update" in app.discord_callbacks:
            for i in app.discord_callbacks["on_member_update"]:
                try:
                    await i(app, before, after)
                except Exception as e:
                    app.logger.error(e)
    # Defines on_member_update.

    @app.dclient.event
    async def on_guild_join(server):
        if await server_restriction_handler(app, server):
            return
        if "on_guild_join" in app.discord_callbacks:
            for i in app.discord_callbacks["on_guild_join"]:
                try:
                    await i(app, server)
                except Exception as e:
                    app.logger.error(e)
    # Defines on_server_join.

    @app.dclient.event
    async def on_guild_remove(server):
        if "on_guild_remove" in app.discord_callbacks:
            for i in app.discord_callbacks["on_guild_remove"]:
                try:
                    await i(app, server)
                except Exception as e:
                    app.logger.error(e)
    # Defines on_server_remove.

    @app.dclient.event
    async def on_guild_update(before, after):
        if "on_guild_update" in app.discord_callbacks:
            for i in app.discord_callbacks["on_guild_update"]:
                try:
                    await i(app, before, after)
                except Exception as e:
                    app.logger.error(e)
    # Defines on_server_update.
