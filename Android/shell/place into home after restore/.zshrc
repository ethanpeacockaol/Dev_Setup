pulseaudio --start \
    --load="module-native-protocol-tcp auth-ip-acl=127.0.0.1 auth-anonymous=1" \
    --exit-idle-time=-1


alias heystop="pkill -u $(whoami)"
#alias juno="cd ~/Desktop && python -m notebook"
alias juno="python ~/.go_buttons/start_juno.go"


clear && echo Yellow-TCL5g
alias root="pd sh ubuntu --user xthan"

