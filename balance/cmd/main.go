package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	_ "github.com/go-sql-driver/mysql"
	"github.com/marcohnp/fullcycle_eda/internal/database"
	"github.com/marcohnp/fullcycle_eda/internal/usecase/get_balance"
	"github.com/marcohnp/fullcycle_eda/internal/usecase/save_balance"
	"github.com/marcohnp/fullcycle_eda/internal/web"
	"github.com/marcohnp/fullcycle_eda/internal/web/webserver"
	"github.com/marcohnp/fullcycle_eda/pkg/kafka"
)

func main() {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", "root", "root", "mysql", "3306", "wallet"))

	if err != nil {
		panic(err)
	}
	defer db.Close()

	getBalanceUsecase := get_balance.NewGetBalanceUsecase(database.NewBalanceDB(db))
	saveBalanceUsecase := save_balance.NewSaveUseCase(database.NewBalanceDB(db))

	configMap := ckafka.ConfigMap{
		"bootstrap.servers": "kafka:29092",
		"group.id":          "wallet",
		"auto.offset.reset": "earliest",
	}

	topics := []string{"balances"}

	consumer := kafka.NewConsumer(&configMap, topics)

	consumerMsgChan := make(chan *ckafka.Message)

	go consumer.Consume(consumerMsgChan)

	consumerFunc := func(msgChan chan *ckafka.Message, saveBalanceUseCase *save_balance.SaveBalanceUsecase) {
		for {
			fmt.Println("Waiting for messages from kafka...")
			msg := <-msgChan
			kafkaMsg := KafkaMsgDto{}
			err := json.Unmarshal(msg.Value, &kafkaMsg)
			kafkaPayload := kafkaMsg.Payload
			if err != nil {
				fmt.Println(err.Error())
			}
			input := save_balance.SaveBalanceInputDto{
				AccountId: kafkaPayload.AccountIDFrom,
				Balance:   float64(kafkaPayload.BalanceAccountIDFrom),
			}
			_, err = saveBalanceUseCase.Execute(input)
			if err != nil {
				fmt.Println(err.Error())
			}
			input = save_balance.SaveBalanceInputDto{
				AccountId: kafkaPayload.AccountIDTo,
				Balance:   float64(kafkaPayload.BalanceAccountIDTo),
			}
			_, err = saveBalanceUseCase.Execute(input)
			if err != nil {
				fmt.Println(err.Error())
			}
		}
	}

	go consumerFunc(consumerMsgChan, saveBalanceUsecase)

	getBalanceHandler := web.NewWebBalanceHandler(*getBalanceUsecase)

	webserver := webserver.NewWebServer(":3003")
	webserver.AddHandler("GET", "/balances/{id}", getBalanceHandler.GetBalance)
	fmt.Println("Server is balance running...")
	webserver.Start()
}

type PayloadKafkaDto struct {
	AccountIDFrom        string  `json:"account_id_from"`
	BalanceAccountIDFrom float64 `json:"balance_account_id_from"`
	AccountIDTo          string  `json:"account_id_to"`
	BalanceAccountIDTo   float64 `json:"balance_account_id_to"`
}

type KafkaMsgDto struct {
	Name    string
	Payload PayloadKafkaDto
}

func (kafkaPayload *PayloadKafkaDto) String() string {
	return fmt.Sprintf("AccountIDFrom: %s, BalanceAccountIDFrom: %f, AccountIDTo: %s, BalanceAccountIDTo: %f", kafkaPayload.AccountIDFrom, kafkaPayload.BalanceAccountIDFrom, kafkaPayload.AccountIDTo, kafkaPayload.BalanceAccountIDTo)
}
