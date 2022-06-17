package service

import (
	"be-test/consumer/model"
	"be-test/consumer/repository"
	"be-test/lib"
	"encoding/json"
	"time"

	"github.com/sirupsen/logrus"
)

type BalanceService interface {
	UpdateBalance(data []byte) error
}

type balanceServiceImpl struct {
	balanceRepository repository.BalanceRepository
}

func NewBalanceService(balanceRepository repository.BalanceRepository) BalanceService {
	return &balanceServiceImpl{
		balanceRepository: balanceRepository,
	}
}

func (s *balanceServiceImpl) UpdateBalance(data []byte) error {
	var api model.BalanceAPI
	err := json.Unmarshal(data, &api)
	if err != nil {
		logrus.Errorf("Error unmarshal service: %v", err)
		return err
	}

	var balance = &model.Balance{}
	err = lib.Merge(api, balance)
	if err != nil {
		logrus.Errorf("Error merger service: %v", err)
		return err
	}

	balance.UpdatedAt = lib.TimePtr(time.Now())
	return s.balanceRepository.Update(balance.WalletID.String(), balance)
}
