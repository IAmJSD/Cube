# Cube. Copyright (C) Jake Gealer <jake@gealer.email> 2017-2018.
# This Source Code Form is subject to the terms of the Mozilla Public
# License, v. 2.0. If a copy of the MPL was not distributed with this
# file, You can obtain one at http://mozilla.org/MPL/2.0/.

def Plugin(app):

    @app.command("Allows you to nuke messages.", requires_staff=True, usage="[amount] [user] (User not required if you want to nuke from all users)")
    async def nuke(app):
        if app.args == []:
            await app.say(embed=app.create_embed("Could not find arguments.",
                                "Please provide arguments for this command.",
                                                                error=True))
        else:
            try:
                amount = int(app.args[0])
            except:
                await app.say(embed=app.create_embed("Could not convert your first argument to a number.", "Please make sure the number is the first argument.", error=True))
                return
            user = None
            if len(app.args) >= 2:
                user = app.pass_user(app.message.guild, app.args[1])
                if user is None:
                    await app.say("Could not pass the second argument as a user! Do you still want to run on messages by any user? Simply say yes if so.")
                    try:
                        def check(m):
                            return m.author == app.message.author and m.channel == app.message.channel
                        msg2 = await app.dclient.wait_for('message', check=check, timeout=30)
                        if not msg2.content.strip(' ').lower() == "yes":
                            return
                    except:
                        return
            def check(message):
                if user is None or user == message.author:
                    return True
                else:
                    return False
            if amount > 100:
                amount = 100
            nuked_msgs = await app.message.channel.purge(limit=amount, check=check)
            if user is None:
                user_bit = " "
            else:
                user_bit = f" by `{user.name} ({user.id})` "
            embed = app.create_embed("Messages nuked:", f"{app.message.author.name} nuked {amount} message(s){user_bit}in {app.message.channel.mention}.", success=True)
            await app.say(embed=embed)
            await app.attempt_log(app.message.guild.id, embed=embed)
    # Allows you to nuke messages.
