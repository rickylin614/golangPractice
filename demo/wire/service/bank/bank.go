package bank

type BankService struct{}

func NewBankService() BankService {
	return BankService{}
}

func (b *BankService) GetBank() string {
	return "GetBank"
}
