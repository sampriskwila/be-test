package service

import (
	"be-test/consumer/model"
	"be-test/producer/repository"
)

type BalanceService interface {
	GetBalance(walletID string) (*model.Balance, error)
}

type balanceServiceImpl struct {
	balanceRepository repository.BalanceRepository
}

func NewBalanceService(balanceRepository repository.BalanceRepository) BalanceService {
	return &balanceServiceImpl{
		balanceRepository: balanceRepository,
	}
}

func (s *balanceServiceImpl) GetBalance(walletID string) (*model.Balance, error) {
	return s.balanceRepository.GetBalance(walletID)
}
