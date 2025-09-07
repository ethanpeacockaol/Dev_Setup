# Variables
dot_termux="$HOME/.termux"
properties="$dot_termux/termux.properties"
permission="\n\nallow-external-apps=true\n"

# Functions
press_enter() {
  read -p 'Press enter to continue...'
}

set_permission() {
  am start \
    -a "android.settings.action.$1" \
    -d package:com.termux
}

allow_overlay() {
  set_permission MANAGE_OVERLAY_PERMISSION
  press_enter
}

allow_battery() {
  set_permission \
    REQUEST_IGNORE_BATTERY_OPTIMIZATIONS

  press_enter
}


# Create Directory ~/.termux
if [ ! -d "$dot_termux" ]; then
  mkdir "$dot_termux"
fi

# Append permission in ~/.termux/termux.properties
if ! grep -qs allow-external-apps "$properties"; then
    echo -e $permission >> "$properties"
fi

# Enable permissions
allow_battery
allow_overlay
