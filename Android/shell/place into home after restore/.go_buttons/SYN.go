import subprocess
import os

session_live = False
result = subprocess.run(["tmux", "has-session", "-t", "x"], stdout=subprocess.PIPE, stderr=subprocess.PIPE)
if result.returncode == 0:
  session_live = True
  
if session_live:
  os.system("termux-toast -b yellow -c blue \"Opening Desktop\"")
  os.system("am start --user 0 -n com.termux.x11/com.termux.x11.MainActivity")
else:
  os.system(f"termux-notification --title \"Windows running\" --content \":)\"")
  os.system("termux-toast -b white -c blue \"Starting Windows\"")
  subprocess.run(["tmux", "new", "-d", "-s", "x"])
  subprocess.run(["tmux", "send-keys", "-t", "x", "termux-x11 :1 &" , "C-m"])
  subprocess.run(["tmux", "send-keys", "-t", "x", "pd sh ubuntu --user xthan --shared-tmp --termux-home", "C-m"])
  subprocess.run(["tmux", "send-keys", "-t", "x", "export DISPLAY=:1", "C-m"])
  subprocess.run(["tmux", "send-keys", "-t", "x", "startlxqt", "C-m"])
  import time
  time.sleep(1)
  os.system("am start --user 0 -n com.termux.x11/com.termux.x11.MainActivity")
  os.system("tmux attach -t x")
