package controller

func (eh *EndpointHandler) AddHandlers() {
	eh.Dg.AddHandler(eh.MessageCreate)
}
