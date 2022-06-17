package producer

import (
	"be-test/lib"
	"be-test/producer/repository"
	"be-test/producer/service"
	"net/http"

	"github.com/Shopify/sarama"
	"github.com/gorilla/mux"
)

func RestRouter(producers sarama.SyncProducer) *mux.Router {
	r := mux.NewRouter()
	api := r.PathPrefix("/api/v1").Subrouter()

	customerRouter(api, producers)
	return r
}

func customerRouter(r *mux.Router, producers sarama.SyncProducer) {
	kafka := KafkaProducer{
		Producer: producers,
	}

	var db = lib.DatabaseConnection()
	var balanceRepository = repository.NewBalanceRepository(db)
	var balanceService = service.NewBalanceService(balanceRepository)
	var balanceHandler = NewBalanceHandler(kafka, balanceService)

	r.HandleFunc("/balance/{id}", balanceHandler.Get).Methods(http.MethodGet)
	r.HandleFunc("/balance", balanceHandler.Post).Methods(http.MethodPost)
}
