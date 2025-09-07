import subprocess
import os
import socket
s = socket.socket(socket.AF_INET, socket.SOCK_DGRAM)
s.connect(("8.8.8.8", 80)) # Google's public DNS server
local_ip = s.getsockname()[0]

session_live = False
result = subprocess.run(["tmux", "has-session", "-t", "serve"], stdout=subprocess.PIPE, stderr=subprocess.PIPE)
if result.returncode == 0:
  session_live = True
  
if session_live:
  os.system("termux-toast -b white -c green \"SSH Server Already Running :8022\"")
else:
  os.system(f"termux-notification --title \"SSH Server 8022 {os.getlogin()}\" --content \"{local_ip}\"")
  os.system(f"termux-toast -b yellow -c red \"Starting SSH Server {os.getlogin()} port 8022\"")
  subprocess.run(["tmux", "new", "-d", "-s", "serve"])
  subprocess.run(["tmux", "send-keys", "-t", "serve", "sshd" , "C-m"])
  subprocess.run(["tmux", "send-keys", "-t", "serve", "neofetch", "C-m"])
  
  import time
  time.sleep(1)
  os.system("am start --user 0 -n com.teslacoilsw.launcher/com.teslacoilsw.launcher.NovaLauncher")
  os.system("tmux attach -t serve")
  
