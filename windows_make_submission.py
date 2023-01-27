import os, shutil, subprocess, tarfile, tempfile

subprocess.run("go env -w GOOS=linux")
subprocess.run("go build bot.go")
subprocess.run("go env -w GOOS=windows")

def set_permissions(tarinfo):
	tarinfo.mode = 0o777
	return tarinfo

with tempfile.TemporaryDirectory() as tmpdirname:

	shutil.move("bot", os.path.join(tmpdirname, "bot"))
	shutil.copy("main.py", os.path.join(tmpdirname, "main.py"))

	with tarfile.open("submission.tar.gz", "w:gz") as tar:
		tar.add(tmpdirname, arcname="", filter=set_permissions)
