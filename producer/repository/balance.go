package repository

import (
	"be-test/consumer/model"
	"be-test/lib"
	"encoding/json"
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/syndtr/goleveldb/leveldb"
)

type BalanceRepository interface {
	GetBalance(walletID string) (*model.Balance, error)
}

type balanceRepositoryImpl struct {
	db *leveldb.DB
}

func NewBalanceRepository(db *leveldb.DB) BalanceRepository {
	return &balanceRepositoryImpl{db: db}
}

func (r *balanceRepositoryImpl) GetBalance(walletID string) (balance *model.Balance, err error) {
	data, err := r.db.Get([]byte(walletID), nil)
	if err != nil {
		logrus.Errorf("Error get balance: %v", err)
		return nil, err
	}

	balance = &model.Balance{}
	err = json.Unmarshal(data, balance)
	if err != nil {
		logrus.Errorf("Error unmarshal balance repository: %v", err)
		return nil, err
	}

	thresholdKey := fmt.Sprintf("%s%s", lib.BalanceThreshold, walletID)
	dataThreshold, err := r.db.Get([]byte(thresholdKey), nil)
	if err != nil {
		logrus.Errorf("Error get threshold: %v", err)
		return nil, err
	}

	threshold := &model.Threshold{}
	err = json.Unmarshal(dataThreshold, threshold)
	if err != nil {
		logrus.Errorf("Error unmarshal threshold repository: %v", err)
		return nil, err
	}

	balance.IsThreshold = threshold.IsThreshold
	logrus.Infof("Balance: %v", balance)
	return
}
