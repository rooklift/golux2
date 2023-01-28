package kit

var frame *Frame

func Run(bidder func(*Frame), placer func(*Frame), main_ai func(*Frame)) {
	for {
		frame = make_next_frame()
		if frame.Step == 0 {
			bidder(frame)
			frame.send_bid()
		} else if frame.RealStep() < 0 {
			placer(frame)
			frame.send_placement()
		} else {
			main_ai(frame)
			frame.send_actions()
		}
	}
}
