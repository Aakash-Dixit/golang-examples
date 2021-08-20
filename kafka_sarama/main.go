package main

import (
	"context"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/Shopify/sarama"
)

const (
	broker            = "localhost:9092"
	clientID          = "kafka-app"
	testTopic         = "TestTopic"
	numPartitions     = 1
	replicationFactor = 1
	consumerGroupName = "my-group"
)

var config = sarama.NewConfig()

func main() {
	// setup sarama log to stdout
	//sarama.Logger = log.New(os.Stdout, "", log.Ltime)

	config.Version = sarama.V2_2_0_0

	/* producer config */
	config.Producer.MaxMessageBytes = 100 * 1024 //100kb is max size of each message here
	config.Producer.Retry.Max = 5
	config.Producer.Return.Successes = true
	//batching
	config.Producer.Flush.Messages = 100
	config.Producer.Flush.Bytes = 10240 * 100 //10 kb * 100 messages
	config.Producer.Flush.Frequency = 5 * time.Millisecond

	// using default config for consumer as it seems appropriate

	log.Println("Creating admin client and test topic")
	createAdminClientAndTopic()

	log.Println("Starting sync producer and consumer")
	go produce()
	consume()
}

func createAdminClientAndTopic() {
	// create adminClient
	adminClient, err := sarama.NewClusterAdmin([]string{broker}, config)
	if err != nil {
		log.Fatal("Error while creating cluster admin: ", err.Error())
	}
	defer func() {
		if err = adminClient.Close(); err != nil {
			log.Fatal("error while closing admin client : ", err.Error())
		}
	}()

	// create topic if not exists
	err = adminClient.CreateTopic(testTopic, &sarama.TopicDetail{
		NumPartitions:     numPartitions,
		ReplicationFactor: replicationFactor,
	}, false)

	if err != nil && !strings.Contains(err.Error(), sarama.ErrTopicAlreadyExists.Error()) {
		log.Fatal("Error while creating topic: ", err)
	}
	log.Printf("Topic : %s, created successfully or already exists", testTopic)

	// list topics
	log.Printf("Listing created topics")
	kafkaTopics, err := adminClient.ListTopics()
	for topic, detail := range kafkaTopics {
		log.Println("topic : ", topic, ", topic details : ", detail)
	}
}

func produce() {
	// async producer
	//prd, err := sarama.NewAsyncProducer([]string{broker}, config)

	//counter for creating message
	i := 0

	// sync producer
	producerClient, err := sarama.NewSyncProducer([]string{broker}, config)
	if err != nil {
		log.Fatal("error while creating sync producer")
	}
	defer func() {
		if err = producerClient.Close(); err != nil {
			log.Fatal("error while closing consumer client : ", err.Error())
		}
	}()

	for {
		// Each kafka message has a key and value. The key is used
		// to decide which partition (and consequently, which broker)
		// the message gets published on. The message can also have
		// headers.

		sampleHeaders := []sarama.RecordHeader{
			{
				Key:   []byte("sample-key"),
				Value: []byte("sample-value"),
			},
		}

		sampleMessage := &sarama.ProducerMessage{
			Topic:     testTopic,
			Key:       sarama.StringEncoder(strconv.Itoa(i)),
			Value:     sarama.StringEncoder("this is message" + strconv.Itoa(i)),
			Headers:   sampleHeaders,
			Timestamp: time.Now(),
		}

		partition, offset, err := producerClient.SendMessage(sampleMessage)
		if err != nil {
			log.Println("Error while publish: ", err.Error())
		}

		log.Println("pudlish partition: ", partition)
		log.Println("publish offset: ", offset)

		i++
		// sleep for a millisecond
		time.Sleep(time.Millisecond)
	}
}

func consume() {
	consumerClient, err := sarama.NewConsumerGroup([]string{broker}, consumerGroupName, config)
	if err != nil {
		log.Fatalf("Error creating consumer group client: %v", err)
	}
	defer func() {
		if err = consumerClient.Close(); err != nil {
			log.Fatal("error while closing consumer client : ", err.Error())
		}
	}()

	consumer := &Consumer{}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	for {
		// `Consume` should be called inside an infinite loop, when a
		// server-side rebalance happens, the consumer session will need to be
		// recreated to get the new claims
		if err := consumerClient.Consume(ctx, []string{testTopic}, consumer); err != nil {
			log.Panicf("Error from consumer: %v", err)
		}
		// check if context was cancelled, signaling that the consumer should stop
		if ctx.Err() != nil {
			return
		}
	}
}

// Consumer represents a Sarama consumer group consumer
type Consumer struct {
}

// Setup is run at the beginning of a new session, before ConsumeClaim
func (consumer *Consumer) Setup(sarama.ConsumerGroupSession) error {
	log.Println("Kafka consumer up and ready...")
	return nil
}

// Cleanup is run at the end of a session, once all ConsumeClaim goroutines have exited
func (consumer *Consumer) Cleanup(sarama.ConsumerGroupSession) error {
	log.Println("Perform any required cleanup")
	return nil
}

// ConsumeClaim must start a consumer loop of ConsumerGroupClaim's Messages().
func (consumer *Consumer) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	// NOTE:
	// Do not move the code below to a goroutine.
	// The `ConsumeClaim` itself is called within a goroutine, see:
	// https://github.com/Shopify/sarama/blob/master/consumer_group.go#L27-L29
	for message := range claim.Messages() {
		log.Printf("Message claimed: value = %s, timestamp = %v, topic = %s", string(message.Value), message.Timestamp, message.Topic)
		session.MarkMessage(message, "")
	}
	return nil
}
