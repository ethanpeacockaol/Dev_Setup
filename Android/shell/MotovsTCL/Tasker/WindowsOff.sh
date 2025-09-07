termux-toast -b red -c blue "Stopping Processes"
sleep 3
pkill -u $(whoami)
