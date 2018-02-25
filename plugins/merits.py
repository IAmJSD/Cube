# Cube. Copyright (C) Jake Gealer <jake@gealer.email> 2017-2018.
# This Source Code Form is subject to the terms of the Mozilla Public
# License, v. 2.0. If a copy of the MPL was not distributed with this
# file, You can obtain one at http://mozilla.org/MPL/2.0/.

def Plugin(app):

    @app.command("Allows you to list merits to a user.", usage="[user] (Optional)")
    async def merits(app):
        if len(app.args) != 0 and app.pass_user(app.message.guild, app.args[0]) is None:
            await app.say(embed=app.create_embed("Could not find the user.", "Make sure you tag the user as your first argument or leave it blank.", error=True))
        else:
            if app.args == []:
                user = app.message.author
            else:
                user = app.pass_user(app.message.guild, app.args[0])
            merits = await app.run_mysql("SELECT * FROM merits WHERE user_id = %s AND server_id = %s", (user.id, app.message.guild.id, ), get_many=True)
            embed = app.create_embed(f"Merits for {user.name}:", f"{user.name} has {len(merits)} merit(s) in {app.message.guild.name}.")
            embed.colour = user.colour
            for s in merits:
                if s[4] == "":
                    s[4] = "No merit reason."
                staff = app.pass_user(app.message.guild, s[1])
                if staff is None:
                    staff_formatted = f"{s[1]} [No longer in the server]"
                else:
                    staff_formatted = f"{staff.name} [{staff.id}]"
                embed.add_field(name=f"{s[0]} (Added by {staff_formatted}):", value=s[4], inline=False)
            await app.say(embed=embed)
    # Allows you to list merits to a user.
