import sys
import os
import subprocess

if len(sys.argv) == 1:
	# set default port number
	port = 8888
else:
	port = sys.argv[1]

url = f"http://localhost:{port}"
cmd1 = "cd ~/Desktop"
cmd2 = f"python -m notebook --port={port}"


#os.system(f"{cmd1} && {cmd2}")
#os.system(cmd3)



session_name = f"juno-{port}]"
session_live = False
result = subprocess.run(["tmux", "has-session", "-t", f"{session_name}"], stdout=subprocess.PIPE, stderr=subprocess.PIPE)
if result.returncode == 0:
  session_live = True

if session_live:
  os.system(f"termux-toast -b black -c red \"Notebook ({port}) alread running\"")
  os.system("am start --user 0 -n com.teslacoilsw.launcher/com.teslacoilsw.launcher.NovaLauncher")
else:
  os.system(f"termux-notification --title \"Jupyter Notebook started on port {port} \" --content \"python notebook running\"")
  os.system(f"termux-toast -b black -c purple \"Jupyter {port}\"")
  subprocess.run(["tmux", "new", "-d", "-s", f"{session_name}"])
  subprocess.run(["tmux", "send-keys", "-t", f"{session_name}", f"{cmd1} && {cmd2}" , "C-m"])
  #os.system("am start --user 0 -n com.teslacoilsw.launcher/com.teslacoilsw.launcher.NovaLauncher")

os.system(f"tmux attach -t {session_name}")
