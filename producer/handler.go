package producer

import (
	"be-test/consumer/model"
	"be-test/lib"
	"be-test/producer/service"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type BalanceHandler struct {
	kafka   KafkaProducer
	service service.BalanceService
}

func NewBalanceHandler(kafka KafkaProducer, service service.BalanceService) *BalanceHandler {
	return &BalanceHandler{kafka: kafka, service: service}
}

func (h *BalanceHandler) Get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	walletID := []byte(vars["id"])

	balance, err := h.service.GetBalance(string(walletID))
	if err != nil {
		logrus.Errorf("Error handler get balance: %v", err)

		w.WriteHeader(http.StatusNotFound)
		lib.ResponseBuilder(w, "Data not found")
		return
	}

	if balance == nil {
		logrus.Errorf("Error handler balance null")

		w.WriteHeader(http.StatusNotFound)
		lib.ResponseBuilder(w, "Data not found")
		return
	}

	w.WriteHeader(http.StatusOK)
	lib.ResponseBuilder(w, balance)
}

func (h *BalanceHandler) Post(w http.ResponseWriter, r *http.Request) {
	api := &model.BalanceAPI{}
	balance := model.Balance{}

	err := json.NewDecoder(r.Body).Decode(api)
	if err != nil {
		errMsg := fmt.Sprintf("Request decoder error: %v", err)

		w.WriteHeader(http.StatusBadRequest)
		lib.ResponseBuilder(w, errMsg)
		return
	}

	err = lib.Merge(api, &balance)
	if err != nil {
		errMsg := fmt.Sprintf("Request merger error: %v", err)

		w.WriteHeader(http.StatusBadRequest)
		lib.ResponseBuilder(w, errMsg)
		return
	}

	balanceByte, err := json.Marshal(balance)
	if err != nil {
		errMsg := fmt.Sprintf("Request decoder error: %v", err)

		w.WriteHeader(http.StatusBadRequest)
		lib.ResponseBuilder(w, errMsg)
		return
	}

	err = h.kafka.SendMessage("deposits", lib.BalancePost.ToKey(), balanceByte)
	if err != nil {
		errMsg := fmt.Sprintf("Create customer error : %v", err)

		w.WriteHeader(http.StatusInternalServerError)
		lib.ResponseBuilder(w, errMsg)
		return
	}

	w.WriteHeader(http.StatusCreated)
	lib.ResponseBuilder(w, "customer created successfully")
}
