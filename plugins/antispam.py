# Cube. Copyright (C) Jake Gealer <jake@gealer.email> 2017-2018.
# This Source Code Form is subject to the terms of the Mozilla Public
# License, v. 2.0. If a copy of the MPL was not distributed with this
# file, You can obtain one at http://mozilla.org/MPL/2.0/.

import time
# Imports go here.

def Plugin(app):

    app.time_since_antispam_refresh = time.perf_counter()
    app.msg_authors_since_antispam_refresh = {}
    # Some predefined stuff.

    @app.event
    async def on_message(app):
        if app.time_since_antispam_refresh - time.perf_counter() >= 2:
            app.time_since_antispam_refresh = time.perf_counter()
            app.msg_authors_since_antispam_refresh = {}
        mid = app.message.author.id
        if mid in app.msg_authors_since_antispam_refresh:
            app.msg_authors_since_antispam_refresh[mid] = app.msg_authors_since_antispam_refresh[mid] + 1
        else:
            app.msg_authors_since_antispam_refresh[mid] = 1
        if app.msg_authors_since_antispam_refresh[mid] >= 20:
            m = None
            for r in app.message.guild.roles:
                if m is None and "muted" in r.name.lower():
                    m = r
            if not m is None:
                try:
                    app.msg_authors_since_antispam_refresh[mid] = 0
                    roles_without_everyone = [r for r in app.message.author.roles if r.name != "@everyone"]
                    await app.message.author.remove_roles(*roles_without_everyone, reason="AUTOMUTE: Spam")
                    await app.message.author.add_roles(m, reason="AUTOMUTE: Spam")
                    embed = app.create_embed("Antispam engaged:", f"`{app.message.author.name} ({app.message.author.id})` was caught spamming. I ran my antispam on them.", error=True)
                    try:
                        await app.message.author.send(f"You have been muted for spamming in `{app.message.guild.name}`.")
                    except:
                        pass
                    await app.attempt_log(app.message.guild.id, embed=embed)
                except:
                    pass
    # The antispam stuff.
