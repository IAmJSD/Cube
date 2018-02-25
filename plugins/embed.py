# Cube. Copyright (C) Jake Gealer <jake@gealer.email> 2017-2018.
# This Source Code Form is subject to the terms of the Mozilla Public
# License, v. 2.0. If a copy of the MPL was not distributed with this
# file, You can obtain one at http://mozilla.org/MPL/2.0/.

import asyncio, discord
# Imports go here.

def Plugin(app):

    @app.command("Enters a mode where you can create embeds.", requires_staff=True)
    async def embed(app):
        try:
            await app.say('''```Entering Cube's embed creation mode. Type "help" for help or "exit" to exit this mode.```''')
            e = discord.Embed()
            x = True
            while x:
                def check(m):
                    return m.author == app.message.author and m.channel == app.message.channel
                msg2 = await app.dclient.wait_for('message', check=check, timeout=300)
                if msg2 is None:
                    x = False
                else:
                    cmd = msg2.content.split(' ')[0].lower()
                    args = [x for x in msg2.content.split(' ')[1:] if x != ""]
                    if cmd == "preview":
                        await app.say(embed=e)
                    elif cmd == "exit":
                        await app.say("```Exiting embed creation mode.```")
                        x = False
                    elif cmd == "title":
                        if args == []:
                            await app.say("```No argument found with the title.```")
                        else:
                            args = ' '.join(args)
                            e.title = args
                            await app.say("```Set title.```")
                    elif cmd == "description":
                        if args == []:
                            await app.say("```No argument found with the description.```")
                        else:
                            args = ' '.join(args)
                            e.description = args
                            await app.say("```Set description.```")
                    elif cmd == "colour" or cmd == "color":
                        if args == []:
                            clr = ""
                        else:
                            clr = args[0].lower()
                        if clr == "red":
                            e.colour = 0xff0000
                            await app.say("```Set colour.```")
                        elif clr == "green":
                            e.colour = 0x00ff00
                            await app.say("```Set colour.```")
                        elif clr == "blue":
                            e.colour = 0x0080ff
                            await app.say("```Set colour.```")
                        elif clr == "orange":
                            e.colour = 0xff8000
                            await app.say("```Set colour.```")
                        elif clr == "pink":
                            e.colour = 0xff80ff
                            await app.say("```Set colour.```")
                        else:
                            await app.say("```Colour not found.```")
                    elif cmd == "field":
                        if args == []:
                            await app.say("```No arguments found.```")
                        else:
                            join_args = ' '.join(args)
                            pipe_split_args = [x.strip(' ') for x in join_args.split("|")]
                            if len(pipe_split_args) == 1:
                                await app.say("```2 arguments were not found. Make sure you seperated them with a pipe (|) character.```")
                            else:
                                e.add_field(name=pipe_split_args[0], value=pipe_split_args[1], inline=True)
                                await app.say("```Set field.```")
                    elif cmd == "footer":
                        if args == []:
                            await app.say("```No argument found with the footer.```")
                        else:
                            args = ' '.join(args)
                            e.set_footer(text=args)
                            await app.say("```Set footer.```")
                    elif cmd == "thumbnail_url":
                        if args == []:
                            await app.say("```No argument found with the footer.```")
                        else:
                            join_args = ' '.join(args)
                            e.set_thumbnail(url=join_args)
                            await app.say("```Set the thumbnail URL.```")
                    elif cmd == "send":
                        if args == []:
                            await app.say("```No argument found to send the embed to.```")
                        else:
                            c = app.message.guild.get_channel(int(args[0].lstrip('<#').rstrip('>')))
                            if c is None:
                                await app.say("```Channel not found.```")
                            else:
                                try:
                                    await c.send(embed=e)
                                    await app.say("```Embed sent and embed mode closed.```")
                                    x = False
                                except:
                                    await app.say("```Could not send embed.```")
                    elif cmd == "help":
                        await app.say("```preview - Allows you to preview the embed.\nexit - Closes this mode.\ntitle [title] - Sets the title.\ndescription [description] - Sets the description.\ncolour [colour] - Allows you to set a colour. Choose between red, green, blue, orange and pink.\nfield [name] | [description] - Allows you to set a field.\nfooter [footer] - Allows you to set a footer.\nthumbnail_url [url] - Allows you to set a URL.\nsend [channel] - Allows you to send it to a channel and then closes the embed creation mode.```")
                    else:
                        await app.say('```Command not found. Run "help" for a command list.```')
        except asyncio.TimeoutError:
            pass
    # Enters a mode where you can create embeds.
