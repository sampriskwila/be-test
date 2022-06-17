package test

import (
	"be-test/consumer/model"
	"be-test/lib"
	"encoding/json"
	"testing"
	"time"

	"github.com/sirupsen/logrus"
)

func TestMarshal(t *testing.T) {
	var db = lib.DatabaseConnection()

	var data []model.HistoryDeposit
	data = append(data, model.HistoryDeposit{
		Amount:    lib.Float64Ptr(float64(6000)),
		CreatedAt: lib.TimePtr(time.Now()),
	})

	dataMarshal, err := json.Marshal(data)
	if err != nil {
		logrus.Errorf("Error marshal data: %v", err)
	}
	logrus.Infof("Data marshal: %v", dataMarshal)

	err = db.Put([]byte("history_wallet_id"), dataMarshal, nil)
	if err != nil {
		logrus.Errorf("Error Put: %v", err)
	}

	dataDB, err := db.Get([]byte("history_wallet_id"), nil)
	if err != nil {
		logrus.Errorf("Error Get: %v", err)
	}

	var newData []model.HistoryDeposit
	err = json.Unmarshal(dataDB, &newData)
	if err != nil {
		logrus.Errorf("Error unmarshal data: %v", err)
	}
	logrus.Infof("New Data after unmarshal: %v", newData)

	for _, datum := range newData {
		logrus.Infof("Data for range: %v", datum)
		logrus.Infof("Amount of data range: %v", *datum.Amount)
	}
}
