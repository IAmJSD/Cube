import discord, pymysql, logging, inspect, json, time, threading
from copy import copy
from pluginbase import PluginBase
# Imports go here.

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

class Application():

    def pass_channel(self, string : str):
        return self.dclient.get_channel(string.lstrip('<#').rstrip('>'))
    # Converts str to discord.channel.

    def pass_user(self, string : str, server):
        return server.get_member(string.lstrip('<@').lstrip('!').rstrip('>'))
    # Converts str to discord.user.

    message = None
    # Defines the message.

    args = None
    # Defines the arguments.

    config = json.load(open("config.json", "r"))
    # Defines the config.

    premade_ver = config["bot_name"] + " v" + config["version"]
    # Premakes a string with the name and version.

    mysql_connection = pymysql.connect(host=config["mysql_hostname"],
                                       user=config["mysql_username"],
                                       password=config["mysql_password"],
                                       db=config["mysql_db"],
                                       charset='utf8mb4',
                                       cursorclass=pymysql.cursors.DictCursor)
    # Defines the MySQL connection.

    def get_prefix(self, server_id):
        with self.mysql_connection.cursor() as cursor:
            sql = "SELECT * FROM `custom_prefixes` WHERE `server_id` = %s"
            cursor.execute(sql, (server_id,))
            result = cursor.fetchone()
            cursor.close()
        if result == None:
            return self.config["default_prefix"]
        else:
            return result["prefix"]
    # Fetches the bot prefix.

    def __init__(self, dclient, logger, using_uvloop):
        self.dclient = dclient
        self.logger = logger
        self.using_uvloop = using_uvloop
    # Initialises the script.

    def load(self):

        self.discord_commands = dict()
        self.discord_callbacks = dict()
        self.plugins = dict()
        # Resets the dicts.

        self.plugin_base = PluginBase(package='__main__.plugins')
        # Defines the plugin base for Discord functions.

        self.plugin_source = self.plugin_base.make_plugin_source(
            searchpath=['./plugins'])
        # Defines the plugin source for Discord functions.

        listening_events = ["on_message", "on_ready", "on_message_delete", "on_reaction_add", "on_reaction_remove", "on_channel_delete", "on_channel_create",
        "on_channel_update", "on_member_join", "on_member_remove", "on_member_update", "on_server_join", "on_server_remove", "on_server_update"]
        # The list of events to look for and seperate.

        for plugin in self.plugin_source.list_plugins():
            self.plugins[plugin] = self.plugin_source.load_plugin(plugin)
            for i in self.plugins[plugin].__dict__.keys():
                if callable(self.plugins[plugin].__dict__[i]):
                    if i == "execute_on_init":
                        self.plugins[plugin].__dict__[i](self)
                    elif not inspect.iscoroutinefunction(self.plugins[plugin].__dict__[i]):
                        raise Exception("Definition is not async.")
                    elif i in listening_events:
                        if not i in self.discord_callbacks:
                            self.discord_callbacks[i] = [self.plugins[plugin].__dict__[i]]
                        else:
                            self.discord_callbacks[i].append(self.plugins[plugin].__dict__[i])
                    else:
                        if i in self.discord_commands:
                            raise Exception("A command with the same name already exists.")
                        else:
                            self.discord_commands[i] = self.plugins[plugin].__dict__[i]
            logger.info("Loaded '" + str(plugin) + "' into the plugin cache.")
        # Loads each Discord related plugin into the correct dictionary.

    # Loads the plugins.

    async def say(self, *args, **kwargs):
        if not self.message == None:
            msg = await self.dclient.send_message(self.message.channel, *args, **kwargs)
            return msg
        else:
            raise Exception("Message not found. Is this a message based command?")
    # The say function for message based commands.

    async def whisper(self, *args, **kwargs):
        if not self.message == None:
            msg = await self.dclient.send_message(self.message.author, *args, **kwargs)
            return msg
        else:
            raise Exception("Message not found. Is this a message based command?")
    # The whisper function for message based commands.

    async def attempt_log(self, server_id, embed):
        logging_count_sql = "SELECT COUNT(*) AS COUNT FROM logging_channels WHERE server_id = %s"
        logging_get_sql = "SELECT * FROM logging_channels WHERE server_id = %s"
        with self.mysql_connection.cursor() as cursor:
            cursor.execute(logging_count_sql, (server_id, ))
            if cursor.fetchone()["COUNT"] != 0:
                cursor.execute(logging_get_sql, (server_id, ))
                logging_channel = self.pass_channel(cursor.fetchone()["channel_id"])
                cursor.close()
                if logging_channel != None:
                    embed.set_footer(text=self.premade_ver)
                    try:
                        await self.dclient.send_message(logging_channel, embed=embed)
                    except:
                        pass
    # Attempts to send a embed to the logging channel.

dclient = discord.Client()
# Defines the Discord Client.

logger = logging.getLogger("cube")
# Defines the logger.

app = Application(dclient, logger, using_uvloop)
# Defines the "app".

@dclient.event
async def on_ready():
    if "on_ready" in app.discord_callbacks:
        for i in app.discord_callbacks["on_ready"]:
            try:
                await i(app)
            except Exception as e:
                logger.error(e)
# Defines on_ready.

async def cmd_handler(msg):
    if not msg.channel.is_private:
        prefix = app.get_prefix(msg.server.id)
        if msg.content.startswith(prefix):
            cmd = msg.content.lstrip(prefix).split(' ')[0].lower()
            if cmd in app.discord_commands:
                allowed_to_run = True
                if hasattr(app.discord_commands[cmd], "requires_staff"):
                    y = app.discord_commands[cmd].requires_staff
                    if y:
                        x = False
                        for role in msg.author.roles:
                            if "staff" in role.name.lower():
                                x = True
                        allowed_to_run = x
                elif hasattr(app.discord_commands[cmd], "requires_management"):
                    y = app.discord_commands[cmd].requires_management
                    if y:
                        x = False
                        for role in msg.author.roles:
                            if "managers" in role.name.lower() or "management" in role.name.lower():
                                x = True
                        allowed_to_run = x
                elif hasattr(app.discord_commands[cmd], "requires_bot_admin"):
                    y = app.discord_commands[cmd].requires_bot_admin
                    if y:
                        sql = "SELECT * FROM `bot_admins` WHERE `user_id` = %s"
                        with app.mysql_connection.cursor() as cursor:
                            cursor.execute(sql, (msg.author.id, ))
                            allowed_to_run = not cursor.fetchone() is None
                            cursor.close()
                if allowed_to_run:
                    app_ctx = copy(app)
                    app_ctx.message = msg
                    args = [x for x in msg.content.split(' ') if x != ""]
                    del args[0]
                    app_ctx.args = args
                    try:
                        await app.discord_commands[cmd](app_ctx)
                    except Exception as e:
                        try:
                            embed=discord.Embed(title="Bug alert!",
                                            description="I had a problem executing that command.```{}```Simply reply `yes` in the next 30 seconds for this to be reported or anything else to ignore this.".format(e),
                                            color=0xff0000)
                            embed.set_footer(text=app.premade_ver)
                            await dclient.send_message(msg.channel, embed=embed)
                            msg2 = await dclient.wait_for_message(author=msg.author, timeout=30)
                            if msg2.content.lower() == "yes":
                                owner = dclient.get_server(app.config["owner_server_id"]).get_member(app.config["owner_user_id"])
                                formatted_msg = "New command bug report!\n\nMessage that triggered: `{}`\n\nError: `{}`".format(msg.content, e)
                                await dclient.send_message(owner, formatted_msg)
                                await dclient.send_message(msg.channel, "Bug report submitted! Thanks for making Cube better.")
                        except:
                            pass
                else:
                    app_ctx = copy(app)
                    app_ctx.message = msg
                    embed=discord.Embed(title="Uh oh!",
                                        description="You do not have permission to run this command.",
                                        color=0xff0000)
                    embed.set_footer(text=app.premade_ver)
                    try:
                        await app_ctx.say(embed=embed)
                    except:
                        pass
            else:
                sql = "SELECT * FROM `custom_commands` WHERE `server_id` = %s AND `command` = %s"
                with app.mysql_connection.cursor() as cursor:
                    cursor.execute(sql, (msg.server.id, cmd, ))
                    result = cursor.fetchone()
                    cursor.close()
                if not result is None:
                    try:
                        args = msg.content.split(' ')
                        del args[0]
                        await dclient.send_message(msg.channel, result["response"].replace("$args$", ' '.join(args).replace("@everyone", "\@everyone").replace("@here", "\@here")))
                    except:
                        pass
# Defines the command handler.

@dclient.event
async def on_message(msg):
    if not msg.author == dclient.user:
        if "on_message" in app.discord_callbacks:
            for i in app.discord_callbacks["on_message"]:
                try:
                    app_ctx = copy(app)
                    app_ctx.message = msg
                    await i(app_ctx)
                except Exception as e:
                    logger.error(e)
        await cmd_handler(msg)
# Defines on_message.

@dclient.event
async def on_message_delete(msg):
    if not msg.author == dclient.user:
        if "on_message_delete" in app.discord_callbacks:
            for i in app.discord_callbacks["on_message_delete"]:
                try:
                    app_ctx = copy(app)
                    app_ctx.message = msg
                    await i(app_ctx)
                except Exception as e:
                    logger.error(e)
# Defines on_message_delete.

@dclient.event
async def on_reaction_add(reaction, user):
    if "on_reaction_add" in app.discord_callbacks:
        for i in app.discord_callbacks["on_reaction_add"]:
            try:
                app_ctx = copy(app)
                app_ctx.message = reaction.message
                await i(app_ctx, reaction, user)
            except Exception as e:
                logger.error(e)
# Defines on_reaction_add.

@dclient.event
async def on_reaction_remove(reaction, user):
    if "on_reaction_remove" in app.discord_callbacks:
        for i in app.discord_callbacks["on_reaction_remove"]:
            try:
                app_ctx = copy(app)
                app_ctx.message = reaction.message
                await i(app_ctx, reaction, user)
            except Exception as e:
                logger.error(e)
# Defines on_reaction_remove.

@dclient.event
async def on_channel_delete(channel):
    if "on_channel_delete" in app.discord_callbacks:
        for i in app.discord_callbacks["on_channel_delete"]:
            try:
                await i(app, channel)
            except Exception as e:
                logger.error(e)
# Defines on_channel_delete.

@dclient.event
async def on_channel_create(channel):
    if "on_channel_create" in app.discord_callbacks:
        for i in app.discord_callbacks["on_channel_create"]:
            try:
                await i(app, channel)
            except Exception as e:
                logger.error(e)
# Defines on_channel_create.

@dclient.event
async def on_channel_update(before, after):
    if "on_channel_update" in app.discord_callbacks:
        for i in app.discord_callbacks["on_channel_update"]:
            try:
                await i(app, before, after)
            except Exception as e:
                logger.error(e)
# Defines on_channel_update.

@dclient.event
async def on_member_join(member):
    if "on_member_join" in app.discord_callbacks:
        for i in app.discord_callbacks["on_member_join"]:
            try:
                await i(app, member)
            except Exception as e:
                logger.error(e)
# Defines on_member_join.

@dclient.event
async def on_member_remove(member):
    if "on_member_remove" in app.discord_callbacks:
        for i in app.discord_callbacks["on_member_remove"]:
            try:
                await i(app, member)
            except Exception as e:
                logger.error(e)
# Defines on_member_remove.

@dclient.event
async def on_member_update(before, after):
    if "on_member_update" in app.discord_callbacks:
        for i in app.discord_callbacks["on_member_update"]:
            try:
                await i(app, before, after)
            except Exception as e:
                logger.error(e)
# Defines on_member_update.

@dclient.event
async def on_server_join(server):
    if "on_server_join" in app.discord_callbacks:
        for i in app.discord_callbacks["on_server_join"]:
            try:
                await i(app, server)
            except Exception as e:
                logger.error(e)
# Defines on_server_join.

@dclient.event
async def on_server_remove(server):
    if "on_server_remove" in app.discord_callbacks:
        for i in app.discord_callbacks["on_server_remove"]:
            try:
                await i(app, server)
            except Exception as e:
                logger.error(e)
# Defines on_server_remove.

@dclient.event
async def on_server_update(before, after):
    if "on_server_update" in app.discord_callbacks:
        for i in app.discord_callbacks["on_server_update"]:
            try:
                await i(app, before, after)
            except Exception as e:
                logger.error(e)
# Defines on_server_update.

def mysql_keep_alive(app):
    while True:
        with app.mysql_connection.cursor() as cursor:
            cursor.execute("SELECT COUNT(*)")
            cursor.close()
        time.sleep(30)
# Every 30 seconds, run a small SQL command to stop pymysql from dying (weird bug I was personally facing).

def main():
    threading.Thread(target=mysql_keep_alive, args=(app,)).start()
    logging.basicConfig(level=logging.INFO)
    app.load()
    app.dclient.run(app.config["token"])
# Defines the main definition.

if __name__ == '__main__':
    main()
# Starts the script.
