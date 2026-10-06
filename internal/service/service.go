package service

// Service create.
type Service struct {
	UserRepository          UserRepository
	PurchasesRepository     PurchasesRepository
	CoinTransfersRepository CoinTransfersRepository
	MerchRepository         MerchRepository
	Transactor              Transactor
	Produce                 Produce

	jwtSecret string
}

// New constructor for service create.
func New(
	userRepository UserRepository,
	coinTransfersRepository CoinTransfersRepository,
	purchasesRepository PurchasesRepository,
	merchRepository MerchRepository,
	transactor Transactor,
	produce Produce,

	jwtSecret string,
) *Service {
	return &Service{
		UserRepository:          userRepository,
		PurchasesRepository:     purchasesRepository,
		CoinTransfersRepository: coinTransfersRepository,
		MerchRepository:         merchRepository,
		Transactor:              transactor,
		Produce:                 produce,

		jwtSecret: jwtSecret,
	}
}
