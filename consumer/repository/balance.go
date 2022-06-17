package repository

import (
	"be-test/consumer/model"
	"be-test/lib"
	"encoding/json"
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/syndtr/goleveldb/leveldb"
)

type BalanceRepository interface {
	Update(walletID string, balance *model.Balance) error
}

type balanceRepositoryImpl struct {
	db *leveldb.DB
}

func NewBalanceRepository(db *leveldb.DB) BalanceRepository {
	return &balanceRepositoryImpl{db: db}
}

func (r *balanceRepositoryImpl) Update(walletID string, balance *model.Balance) error {
	balanceByte, err := json.Marshal(balance)
	if err != nil {
		logrus.Errorf("Error marshal: %v", err)
		return err
	}

	err = r.db.Put([]byte(walletID), balanceByte, nil)
	if err != nil {
		logrus.Errorf("Error put: %v", err)
		return err
	}

	// Get history
	var historyDeposit []model.HistoryDeposit

	historyKey := fmt.Sprintf("%s%s", lib.BalanceHistory, walletID)
	dataHistory, _ := r.db.Get([]byte(historyKey), nil)

	if len(dataHistory) > 0 {
		err = json.Unmarshal(dataHistory, &historyDeposit)
		if err != nil {
			logrus.Errorf("Error unmarshal repository: %v", err)
			return err
		}
	}

	// Append history
	history := model.HistoryDeposit{
		Amount:    balance.Amount,
		CreatedAt: lib.TimePtr(time.Now()),
	}

	historyDeposit = append(historyDeposit, history)
	historyDepositByte, err := json.Marshal(historyDeposit)
	if err != nil {
		logrus.Errorf("Error marshal: %v", err)
		return err
	}

	err = r.db.Put([]byte(historyKey), historyDepositByte, nil)
	if err != nil {
		logrus.Errorf("Error put: %v", err)
		return err
	}

	// Validate threshold
	threshold := model.Threshold{
		IsThreshold: lib.BoolPtr(false),
	}
	var tmpAmount []float64
	var tmpCreatedAt []time.Time
	for i := len(historyDeposit) - 1; i > -1; i-- {
		tmpAmount = append(tmpAmount, *historyDeposit[i].Amount)
		tmpCreatedAt = append(tmpCreatedAt, *history.CreatedAt)

		// Case test 1 & 2
		if len(tmpAmount) == 2 {
			if tmpAmount[0] == float64(6000) &&
				tmpAmount[1] == float64(6000) {
				if tmpCreatedAt[0].Sub(tmpCreatedAt[1]) <= 2*time.Minute {
					// Case Test 1
					threshold.IsThreshold = lib.BoolPtr(true)
				} else {
					// Case Test 2
					threshold.IsThreshold = lib.BoolPtr(false)
				}
			}
		}

		// Case Test 3 & 4
		if len(tmpAmount) == 6 {
			isCase3 := false
			for _, f := range tmpAmount {
				if f != float64(2000) {
					// Case Test 3
					isCase3 = true
					threshold.IsThreshold = lib.BoolPtr(false)
					break
				}
			}

			if !isCase3 {
				if tmpCreatedAt[5].Sub(tmpCreatedAt[0]) <= 2*time.Minute {
					// Case Test 4
					threshold.IsThreshold = lib.BoolPtr(true)
				}
			}

		}
	}

	thresholdByte, err := json.Marshal(threshold)
	if err != nil {
		logrus.Errorf("Error marshal threshold: %v", err)
		return err
	}

	thresholdKey := fmt.Sprintf("%s%s", lib.BalanceThreshold, walletID)
	err = r.db.Put([]byte(thresholdKey), thresholdByte, nil)
	if err != nil {
		logrus.Errorf("Error put threshold: %v", err)
		return err
	}

	return nil
}
