import os, shutil, subprocess, tarfile

subprocess.run("go env -w GOOS=linux")
subprocess.run("go build bot.go")
subprocess.run("go env -w GOOS=windows")

try:
	os.mkdir("sub")
except:
	pass

shutil.copy("bot", "sub/bot")
shutil.copy("main.py", "sub/main.py")

with tarfile.open("submission.tar.gz", "w:gz") as tar:
	tar.add("sub", arcname="")
