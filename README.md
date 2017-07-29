<p align="center"><img src="http://i.imgur.com/QrhAeBe.png"></p>
<p align="center">FactoCord -- a Factorio to Discord bridge bot for Linux</p>
<p align="center">
<a href="https://goreportcard.com/report/github.com/FactoKit/FactoCord"><img src="https://goreportcard.com/badge/github.com/FactoKit/FactoCord" alt="Go Report Card"></a>
</p>

# Running
*Make sure you have your .env file in the same directory as the executable/binary, you can use .envexample the template*

There are two ways of starting FactoCord

1. Using the start.sh bash script (bash start.sh or ./start.sh) (make sure you chmod +x the script first)
2. Manually running the binary (./FactoCord)

# Installing as a service

To install FactoCord as a service so that it can run on startup, you can use the provided service.sh

*Note you must run service.sh as root/sudo to install it as a service*

Example of running service.sh:
`./service.sh factorio /home/facotrio/factocord/`


# Compiling

`Requires go 1.8 or above`

FactoCord uses the following packages:

- [DiscordGo](https://github.com/bwmarrin/discordgo)
- [godotenv](https://github.com/joho/godotenv/)
- [tails](https://github.com/hpcloud/tail)

To compile just do `go build`


# Error reporting

When FactoCord encounters an error will log to error.log within the same directory as itself.

If you are having an issue make sure to check the error.log to see what the problem is.

If you are unable to solve the issue yourself, please post an issue containing the error.log and I will review and attempt to solve what the problem is.


# Windows Support?

Currently I haven't had any luck getting FactoCord to run correctly on Windows, [see this](https://github.com/FactoKit/FactoCord/issues/3) for information

If a way is found to fix this problem, then Windows support will be added.


# Screenshots

<p><img src="http://i.imgur.com/JsLOVst.png" alt="restart command"></p>
<p><img src="http://i.imgur.com/1cxq54P.png" alt="mod list command"></p>
<p><img src="http://i.imgur.com/qN3NsO6.png" alt="stop command"></p>
<p><img src="http://i.imgur.com/cxjvFG8.png" alt="save command"></p>
<p><img src="http://i.imgur.com/dztOTrk.png" alt="in-game chat being sent to discord, notice how you can mention discord members"></p>
<p><img src="http://i.imgur.com/Npl0vBb.png" alt="discord chat being sent to in-game"></p>


# Special Thanks

  - Brett and MajesticFudgie for making the logo!
  - [UclCommander](https://github.com/UclCommander) for finding me the tails library which made this a lot easier to build.
