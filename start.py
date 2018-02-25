# Cube. Copyright (C) Jake Gealer <jake@gealer.email> 2017-2018.
# This Source Code Form is subject to the terms of the Mozilla Public
# License, v. 2.0. If a copy of the MPL was not distributed with this
# file, You can obtain one at http://mozilla.org/MPL/2.0/.

from internal.dclient import load_dclient
from pluginbase import PluginBase
import discord, logging, inspect, json, os, aiomysql, re
# Imports go here.

print(r"""
   ______      __
  / ____/_  __/ /_  ___
 / /   / / / / __ \/ _ \
/ /___/ /_/ / /_/ /  __/
\____/\__,_/_.___/\___/

This bot is based on the Cube core created by JakeMakesStuff on GitHub and licensed under the MPL-2.0 license.
""")
# ASCII!!!!!!11111111111111!

try:
    import asyncio, uvloop, concurrent.futures
    asyncio.set_event_loop_policy(uvloop.EventLoopPolicy())
    loop = asyncio.get_event_loop()
    pool = concurrent.futures.ThreadPoolExecutor()
    loop.set_default_executor(pool)
    using_uvloop = True
except:
    using_uvloop = False
# Uses uvloop if it is installed.

class Application:

    def __init__(self, using_uvloop): self.using_uvloop = using_uvloop
    # Sets whether the bot is using uvloop.

    def pass_user(self, server, user_string):
        try:
            if isinstance(user_string, str):
                uid = ''.join(re.findall("[0-9]", user_string))
            else:
                uid = user_string
            return server.get_member(int(uid))
        except:
            return None
    # Attempts to pass a user.

    def pass_channel(self, channel_string, server = None):
        try:
            if isinstance(channel_string, str):
                cid = int(''.join(re.findall("[0-9]", channel_string)))
            else:
                cid = channel_string
            if server is None:
                return self.dclient.get_channel(cid)
            else:
                return server.get_channel(cid)
        except:
            return None
    # Attempts to pass a channel.

    message = None
    # Sets the message to a Nonetype.

    cube_root = os.path.dirname(os.path.abspath(__file__))
    # The exact path of Cube's root folder.

    config = json.load(
        open(os.path.join(cube_root, "config.json"), "r")
    )
    # Loads the config.

    def create_embed(self, title, description = "", success = False, error = False):
        if success:
            c = 0x00ff00
        elif error:
            c = 0xff0000
        else:
            c = 0x000000
        e = discord.Embed(title=title, description=description, color=c)
        e.set_footer(text=f'{self.config["bot_name"]} v{self.config["version"]}')
        return e
    # Creates a branded embed.

    logger = logging.getLogger(config["bot_name"].lower())
    # Sets the logger.

    async def load_mysql(self):
        self.mysql_pool = await aiomysql.create_pool(host=self.config["mysql_hostname"],
                                               user=self.config["mysql_username"],
                                               password=self.config["mysql_password"],
                                               db=self.config["mysql_db"],
                                               charset='utf8mb4')
        self.logger.info("Initialised MySQL. Ready to accept connections!")
    # Loads MySQL on boot.

    async def run_mysql(self, *args, **kwargs):
        async with self.mysql_pool.acquire() as conn:
            async with conn.cursor() as c:
                await c.execute(*args)
                r = None
                if kwargs.get('get_one', False):
                    r = await c.fetchone()
                elif kwargs.get('get_many', False):
                    r = await c.fetchall()
                if kwargs.get('commit', False):
                    await conn.commit()
        return r
    # Runs any MySQL scripts.

    async def attempt_log(self, server_id, *args, **kwargs):
        logging_channel = await self.run_mysql("SELECT * FROM `logging_channels` WHERE `server_id` = %s",
        (server_id,), get_one=True)
        if not logging_channel is None:
            c = self.dclient.get_channel(logging_channel[1])
            try:
                await c.send(*args, **kwargs)
            except:
                pass
    # Attempts to send a message to the logging channel.

    async def get_prefix(self, server_id):
        prefix_sql = await self.run_mysql("SELECT * FROM `custom_prefixes` WHERE `server_id` = %s", (server_id,), get_one=True)
        if prefix_sql is None:
            return self.config["default_prefix"]
        else:
            return prefix_sql[1]
    # Gets the prefix.

    discord_callbacks = dict()
    discord_commands = dict()
    plugins = dict()
    # Defines the dictionarys.

    dclient = discord.AutoShardedClient()
    # Sets the Discord client.

    def command(self,
    description = "No description given.",
    requires_staff = False,
    requires_management = False,
    requires_bot_admin = False,
    name = None,
    usage = None):
        def deco(func):
            cmd_attr = {"description" : description,
                        "requires_staff" : requires_staff,
                        "requires_management" : requires_management,
                        "requires_bot_admin" : requires_bot_admin,
                        "usage" : usage}
            func.cmd_attr = cmd_attr
            if name is None:
                func_name = func.__name__
            else:
                func_name = name
            if not inspect.iscoroutinefunction(func):
                raise Exception("Command not async.")
            self.discord_commands[func_name] = func
        return deco
    # The command decorator.

    async def say(self, *args, **kwargs):
        if not self.message == None:
            msg = await self.message.channel.send(*args, **kwargs)
            return msg
        else:
            raise Exception("Message not found. Is this a message based command?")
    # The say function for message based commands.

    async def whisper(self, *args, **kwargs):
        if not self.message == None:
            msg = await self.message.author.send(*args, **kwargs)
            return msg
        else:
            raise Exception("Message not found. Is this a message based command?")
    # The whisper function for message based commands.

    def event(self, func):
        if func.__name__ not in self.discord_callbacks:
            self.discord_callbacks[func.__name__] = []
        if not inspect.iscoroutinefunction(func):
            raise Exception("Event not async.")
        self.discord_callbacks[func.__name__].append(func)
    # The event decorator.

    def load_plugins(self):
        self.discord_callbacks = dict()
        self.discord_commands = dict()
        self.plugins = dict()
        self.plugin_base = PluginBase(package='__main__.plugins')
        self.plugin_source = self.plugin_base.make_plugin_source(searchpath=[os.path.join(self.cube_root, "./plugins")])
        for plugin in self.plugin_source.list_plugins():
            self.plugins[plugin] = self.plugin_source.load_plugin(plugin)
            self.plugins[plugin].Plugin(self)
            self.logger.info(f'Loaded "{plugin}" into the plugin cache.')
    # Loads the plugins.

    def run(self):
        if isinstance(self.config["owner_server_id"], str):
            self.config["owner_server_id"] = int(self.config["owner_server_id"])
        if isinstance(self.config["owner_user_id"], str):
            self.config["owner_user_id"] = int(self.config["owner_user_id"])
        logging.basicConfig(level=logging.INFO)
        self.load_plugins()
        self.dclient.loop.create_task(self.load_mysql())
        self.dclient.run(self.config["token"])
    # Function to start the bot.

app = Application(using_uvloop)
# Defines the application.

load_dclient(app)
# Loads the events into the Discord client.

app.run()
# Starts the bot.
