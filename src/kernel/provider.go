package kernel

func RegisterProvider(app ApplicationInterface) {
	container := app.GetContainer()

	// first merge configs
	container.GetConfig()

	// register global http client

	// register log service

}
