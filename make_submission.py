import datetime, os, shutil, subprocess, tarfile, tempfile

now = datetime.datetime.now()
outfilename = "submission_{}.tar.gz".format(now.strftime("%Y_%m_%d_%H%M"))

GOOS = subprocess.check_output(["go", "env", "GOOS"]).decode().strip()
GOARCH = subprocess.check_output(["go", "env", "GOARCH"]).decode().strip()

print("GOOS is {} and GOARCH is {}".format(GOOS, GOARCH))
print("Setting GOOS to linux and GOARCH to amd64")

subprocess.run(["go", "env", "-w", "GOOS=linux"])
subprocess.run(["go", "env", "-w", "GOARCH=amd64"])

print("Building bot")

subprocess.run(["go", "build", "golux2"])

print("Resetting GOOS and GOARCH to their original values")

subprocess.run(["go", "env", "-w", "GOOS={}".format(GOOS)])
subprocess.run(["go", "env", "-w", "GOARCH={}".format(GOARCH)])

print("Building tar.gz file")

def set_permissions(tarinfo):
	tarinfo.mode = 0o777
	return tarinfo

with tempfile.TemporaryDirectory() as tmpdirname:
	shutil.move("golux2", os.path.join(tmpdirname, "golux2"))
	shutil.copy("main.py", os.path.join(tmpdirname, "main.py"))
	with tarfile.open(outfilename, "w:gz") as tar:
		tar.add(tmpdirname, arcname="", filter=set_permissions)

