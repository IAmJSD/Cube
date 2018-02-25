# Cube. Copyright (C) Jake Gealer <jake@gealer.email> 2017-2018.
# This Source Code Form is subject to the terms of the Mozilla Public
# License, v. 2.0. If a copy of the MPL was not distributed with this
# file, You can obtain one at http://mozilla.org/MPL/2.0/.

import discord
# Imports go here.

def Plugin(app):

    @app.command("Allows bot admins to change the game.", requires_bot_admin=True, usage="[game]")
    async def game(app):
        await app.dclient.change_presence(game=discord.Game(name=' '.join(app.args)))
        await app.say("Game set to {}".format(' '.join(app.args)))
        app.logger.info("Game set to {}".format(' '.join(app.args)))
    # Allows bot admins to change the game.

    @app.event
    async def on_ready(app):
        await app.dclient.change_presence(game=discord.Game(name=app.config["game"]))
        app.logger.info('Game set to "{}".'.format(app.config["game"]))
    # Sets the game on bot boot.
