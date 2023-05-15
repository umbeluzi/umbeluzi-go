package umbeluzi

type Client struct {
	Transfers        *TransfersService
	Providers        *ProvidersService
	ReceivingMethods *ReceivingMethodsService
}
