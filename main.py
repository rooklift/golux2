import atexit, json, os, subprocess, sys, threading

# Spawn...

cwd = os.path.dirname(__file__)
if cwd == "":
	cwd ="./"

proc = subprocess.Popen(["bot.exe"], stdin=subprocess.PIPE, stdout=subprocess.PIPE, stderr=subprocess.PIPE, cwd=cwd)
atexit.register(lambda: proc.kill())

# Bot output will be sent to the referee...

def bot_output_handler(in_stream, out_stream):
	while True:
		s = in_stream.readline().decode("utf8")
		out_stream.write(s)
		out_stream.flush()

threading.Thread(target = bot_output_handler, daemon = True, args = [proc.stdout, sys.stdout]).start()
threading.Thread(target = bot_output_handler, daemon = True, args = [proc.stderr, sys.stderr]).start()

# Referee output will be sent to the bot, modified a bit.

msg_old = None
msg = None
cfg = None

i = -1

while True:

	i += 1
	s = sys.stdin.readline()

	msg_old = msg
	msg = json.loads(s)

	if "info" in msg and "env_cfg" in msg["info"] and len(msg["info"]["env_cfg"]) > 0:
		cfg = msg["info"]["env_cfg"]
	else:
		if cfg:
			msg["info"]["env_cfg"] = cfg

	if i > 0:
		fixup_board = msg_old["obs"]["board"]
		for key1 in ["rubble", "lichen", "lichen_strains"]:
			for key2, value in msg["obs"]["board"][key1].items():
				x, y = [int(z) for z in key2.split(",")]
				fixup_board[key1][x][y] = value
		msg["obs"]["board"] = fixup_board

	out_s = json.dumps(msg) + "\n"

	proc.stdin.write(out_s.encode("utf8"))
	proc.stdin.flush()

