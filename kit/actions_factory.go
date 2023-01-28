package kit

func (self *Factory) Act(action int) {
	self.Frame.factory_actions[self.UnitId] = action
}

func (self *Factory) Cancel() {
	delete(self.Frame.factory_actions, self.UnitId)
}
