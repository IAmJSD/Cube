# Cube. Copyright (C) Jake Gealer <jake@gealer.email> 2017-2018.
# This Source Code Form is subject to the terms of the Mozilla Public
# License, v. 2.0. If a copy of the MPL was not distributed with this
# file, You can obtain one at http://mozilla.org/MPL/2.0/.

def Plugin(app):

    @app.command("Sends a setup guide for the bot.")
    async def setup(app):
        await app.say("https://gist.github.com/JakeMakesStuff/7bef3a9f3f6cc0f1f8dd96efac427f1d")
    # Sends a setup guide for the bot.
