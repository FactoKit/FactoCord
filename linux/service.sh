# !/bin/bash
# Installs FactoCord as a systemd service
# Script by Allen Lydiard (FM1337)


# First let's check if the user is running this script as root.
if [ "$EUID" -ne 0 ]; then
	echo "Sorry to install this service, you need to run this script as root"
	echo "To run it as root do sudo bash service.sh (or sudo ./service.sh)"
	exit 1
fi


# We then check for the amount of arguments.
if [ $# -eq 0 ]; then
	echo "You didn't supply any arguments!"
	echo "You must supply the following arguments"
	echo "1. The user to run the service as: e.g factorio"
	echo "2. The working directory: e.g /home/factorio/factocord"
	exit 1
fi


# Followed by checking if the user to run the service as exists.
if ! id "$1" >/dev/null 2>&1; then
	echo "Sorry the user you want to run the service as does not exist!"
	exit 1
fi

# We also need to check if the working directory exists.
if [ ! -d "$2" ]; then
	echo "Sorry the working directory you supplied does not exist."
	echo "Example of working directory: /home/factorio/factocord/"
	exit 1
fi

# Moving on we need to check if the start.sh script exists in the working directory.
if [ ! -f "$2/start.sh" ]; then
	echo "Sorry the start.sh script isn't in the working directory."
	echo "You can fetch a copy of it from here https://github.com/FactoKit/FactoCord/tree/master/linux"
	exit 1
fi

# After we verified start.sh exist, we need to check if the FactoCord binary exists.
if [ ! -f "$2/FactoCord" ]; then
	echo "Sorry the FactoCord binary is missing from your working directory"
	echo "You can get a copy of it from here https://github.com/FactoKit/FactoCord/releases"
	exit 1
fi

# Last but not least, we check if the factocord service is already installed and ask if it should be
# reinstalled.
if [ -f "/etc/systemd/system/factocord.service" ]; then
	while true; do
		echo "It looks like Factocord is already installed a service"
		read -p "Would you like to reinstall it? " yn
		case $yn in
			[Yy]* ) systemctl disable factocord.service; break;;
			[Nn]* ) exit 0;;
			* ) echo "Please enter yes or no.";;
		esac
	done
fi


ServiceVar="[Unit]
Description=FactoCord Bot
Wants=network-online.target
After=network.target network-online.target

[Service]
User=$1
WorkingDirectory=$2
PIDFile=$2/factocord.pid
ExecStart=/bin/bash start.sh
Restart=on-failure
StartLimitInterval=600

[Install]
WantedBy=multi-user.target"

echo "$ServiceVar" > /etc/systemd/system/factocord.service

systemctl daemon-reload
systemctl enable factocord.service
systemctl start factocord.service
