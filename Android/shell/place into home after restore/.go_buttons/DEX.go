import subprocess
import os

session_live = False
result = subprocess.run(["tmux", "has-session", "-t", "x-zoom"], stdout=subprocess.PIPE, stderr=subprocess.PIPE)
if result.returncode == 0:
  session_live = True


if session_live == False:
  os.system("termux-toast -b blue -c white \"Starting Windows (zoom)\"")
  subprocess.run(["tmux", "new", "-d", "-s", "x-zoom"])
  subprocess.run(["tmux", "send-keys", "-t", "x-zoom", "pd sh ubuntu --user xthan --termux-home", "C-m"])
  subprocess.run(["tmux", "send-keys", "-t", "x-zoom", "vncserver -kill", "C-m"])
  subprocess.run(["tmux", "send-keys", "-t", "x-zoom", "rm -rf /tmp/.X*", "C-m"])
  subprocess.run(["tmux", "send-keys", "-t", "x-zoom", "vncserver -xstartup /usr/bin/startlxqt -geometry 720x1600", "C-m"])
  import time
  time.sleep(1)
  os.system("am start --user 0 -n com.realvnc.viewer.android/com.realvnc.viewer.android.app.ConnectionChooserActivity")
  time.sleep(1)
  os.system("tmux attach")
else:
  os.system("termux-toast -b white -c blue \"Opening Desktop zoom\"")
  os.system("am start --user 0 -n com.realvnc.viewer.android/com.realvnc.viewer.android.app.ConnectionChooserActivity")

