import datetime, os, shutil, subprocess, tarfile, tempfile

now = datetime.datetime.now()
outfilename = "submission_{}.tar.gz".format(now.strftime("%Y_%m_%d_%H%M"))

subprocess.run("go env -w GOOS=linux")
subprocess.run("go build golux2")
subprocess.run("go env -w GOOS=windows")

def set_permissions(tarinfo):
	tarinfo.mode = 0o777
	return tarinfo

with tempfile.TemporaryDirectory() as tmpdirname:

	shutil.move("golux2", os.path.join(tmpdirname, "golux2"))
	shutil.copy("main.py", os.path.join(tmpdirname, "main.py"))

	with tarfile.open(outfilename, "w:gz") as tar:
		tar.add(tmpdirname, arcname="", filter=set_permissions)
