# Setting Up Cube

To initially start the install, Cube requires the following:
- A really stable internet connection, ideally with 1-2ms ping and decent speeds.
- A MySQL/MariaDB database (settting up of the database will be described later).
- Python 3.5 or above (Pip installation/usage will be described later).

## Database Setup
In order for Cube to be setup, you need to setup the database for it. Firstly, install MySQL or MariaDB (I personally use MariaDB). After you have done this, make a account with a **strong password** which can be accessed from outside the LAN for easy management. Enable utf8mb4 as your database character set, create/select a database with the name you want and run the code in `prepare_db.sql` as a SQL query. Google should help you out with this if you get stuck since this is just the basics of setting up a MySQL database. You can then set the MySQL information in the `config.json`.

## Pip Setup
Pip is a very essential package manager for anything running Python 3.5 or above. To verify you have it installed, run `py -m pip` (or sometimes `python3.5` if you have Linux). If you get a error about pip not being found, run https://bootstrap.pypa.io/get-pip.py as administrator and try again. From there in the root directory run `py -m pip install -r requirements.txt` (or sometimes `python3.5` if you have Linux). If you are on Linux you can also install uvloop from Pip for added performance.

## Bot Account Setup
Go to Discord Developers and then create a bot account, you can then copy and paste the ID into https://discordapi.com/permissions.html to generate the invite and then put the token into `config.json`.

## Setting the Owner User/Server ID.
To do this, simply put in `owner_server_id` the ID of a server you will be in and in with the bot and in `owner_user_id` the ID of your user.

## Starting
After finishing filling out the remaining items in `config.json` you can start the bot by starting `cube.py` with your Python 3.5+ interpreter. If you are using Linux, you may want to use a `screen` so that you can leave the SSH session without killing the bot.
