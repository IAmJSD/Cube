# Cube. Copyright (C) Jake Gealer <jake@gealer.email> 2017-2018.
# This Source Code Form is subject to the terms of the Mozilla Public
# License, v. 2.0. If a copy of the MPL was not distributed with this
# file, You can obtain one at http://mozilla.org/MPL/2.0/.

import json, os
# Imports go here.

def Plugin(app):

    @app.command("Reloads the bot.", requires_bot_admin=True)
    async def reload(app):
        app.load_plugins()
        app.config = json.load(
            open(os.path.join(app.cube_root, "config.json"), "r")
        )
        await app.say("Reloaded!")
    # Reloads the bot.
