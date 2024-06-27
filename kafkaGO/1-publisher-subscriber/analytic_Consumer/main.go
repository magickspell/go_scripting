package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/IBM/sarama"
	"log"
	"time"
)

type ProductInfo struct {
	SKU   int64
	Price float64
	Cnt   int64
}

type OrderInfo struct {
	UserID    int64
	CreatedAt time.Time
	Products  []ProductInfo
}

type Consumer struct {
	ID int
}

func (consumer *Consumer) Setup(sarama.ConsumerGroupSession) error {
	return nil
}

func (consumer *Consumer) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

func (consumer *Consumer) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for message := range claim.Messages() {
		var msg OrderInfo
		err := json.Unmarshal(message.Value, &msg)
		if err != nil {
			fmt.Printf("Error unmarshaling message: %s\n", err)
		}
		fmt.Printf("Msg: %v, Consumer: %d\n", msg, consumer.ID)

		session.MarkMessage(message, "done")
	}
	return nil
}

func subscribe(ctx context.Context, topic string, consumerGroup sarama.ConsumerGroup) error {
	consumer := Consumer{}

	go func() {
		for {
			if err := consumerGroup.Consume(ctx, []string{topic}, &consumer); err != nil {
				fmt.Printf("Error from consumerGroup consume: %s\n", err)
			}
			if ctx.Err() != nil {
				return
			}
		}
	}()
	return nil
}

var brokers = []string{"127.0.0.1:9095", "127.0.0.1:9096", "127.0.0.1:9097"}

func StartConsuming(ctx context.Context) error {
	config := sarama.NewConfig()

	config.Consumer.Group.Rebalance.Strategy = sarama.BalanceStrategyRoundRobin
	config.Consumer.Offsets.Initial = sarama.OffsetOldest

	consumerGroup, err := sarama.NewConsumerGroup(brokers, "analytic_consumer", config)
	if err != nil {
		fmt.Printf("consumerGroup error: %s\n", err)
		return err
	}
	//_ = subscribe(ctx, "orders", consumerGroup)
	//time.Sleep(time.Second)

	return subscribe(ctx, "orders", consumerGroup)
}

func main() {
	ctx := context.Background()

	err := StartConsuming(ctx)
	if err != nil {
		log.Fatal(err)
		return
	}

	for {
	}
}
