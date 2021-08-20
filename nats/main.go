package main

import (
	"log"
	"strconv"
	"time"

	"examples/nats/pb"

	"github.com/golang/protobuf/proto"
	"github.com/nats-io/nats-server/v2/server"
	"github.com/nats-io/nats.go"
)

var (
	natsClient *nats.Conn
	subj1      = "TEST1"
	subj2      = "TEST2"
	queueName  = "Test-Queue"
)

// "*" matches any token, at any level of the subject.
// E.g. 'foo.*.bar' will match 'foo.baz.bar'
// ">" matches any length of the tail of a subject, and can only be the last token
// E.g. 'foo.>' will match 'foo.bar', 'foo.bar.baz', 'foo.foo.bar.bax.22'

func publisheTest1() {
	log.Println("Starting publisher on subject : ", subj1)
	var i int32 = 1
	for {
		person := &pb.Person{
			Id:   i,
			Name: "person-" + strconv.Itoa(int(i)),
			Age:  10 + i,
		}
		personBytes, err := proto.Marshal(person)
		if err != nil {
			log.Fatal("Error while marshalling proto message : ", err.Error())
		}

		log.Println("Publishing message on subject : ", subj1)
		err = natsClient.Publish(subj1, personBytes)
		if err != nil {
			log.Fatal("Error while publishing message to nats : ", err.Error())
		}

		i++
		time.Sleep(time.Second)
	}
}

func requesterTest2() {
	log.Println("Starting publishing and requesting on subject : ", subj2)
	var i int32 = 1
	for {
		msg, err := natsClient.Request(subj2, []byte("Request "+strconv.Itoa(int(i))+" on TEST2"), 25*time.Millisecond)
		if err != nil {
			log.Println("Error while publishing request : ", err.Error())
			continue
		}
		log.Println("Received reply on subject : ", msg.Subject, ", Data : ", string(msg.Data))
		i++
		time.Sleep(time.Second)
	}
}

func subscriberTest1() {
	log.Println("Starting subscriber on subject : ", subj1)
	natsClient.Subscribe(subj1, func(msg *nats.Msg) {
		log.Println("Received message on subject : ", subj1)
		person := &pb.Person{}
		err := proto.Unmarshal(msg.Data, person)
		if err != nil {
			log.Fatal("Error while Unmarshalling proto message : ", err.Error())
		}
		log.Println("Id : ", person.Id, " Name : ", person.Name, " Age : ", person.Age)
	})
}

func queueSubscriber1Test1() {
	log.Println("Starting queue subscriber 1 on subject : ", subj1)
	natsClient.QueueSubscribe(subj1, queueName, func(msg *nats.Msg) {
		log.Println("Received message in queue subscriber 1 on subject : ", subj1)
		person := &pb.Person{}
		err := proto.Unmarshal(msg.Data, person)
		if err != nil {
			log.Fatal("Error while Unmarshalling proto message : ", err.Error())
		}
		log.Println("Id : ", person.Id, " Name : ", person.Name, " Age : ", person.Age)
	})
}

func queueSubscriber2Test1() {
	log.Println("Starting  queue subscriber 2 on subject : ", subj1)
	natsClient.QueueSubscribe(subj1, queueName, func(msg *nats.Msg) {
		log.Println("Received message in queue subscriber 2 on subject : ", subj1)
		person := &pb.Person{}
		err := proto.Unmarshal(msg.Data, person)
		if err != nil {
			log.Fatal("Error while Unmarshalling proto message : ", err.Error())
		}
		log.Println("Id : ", person.Id, " Name : ", person.Name, " Age : ", person.Age)
	})
}

func subscriberSyncTest1() {
	log.Println("Starting subscriber on subject : ", subj1)
	sub, err := natsClient.SubscribeSync(subj1)
	if err != nil {
		log.Fatal("Error while sync subscription for sub ", subj1, " : ", err.Error())
	}
	for {
		msg, err := sub.NextMsg(5 * time.Second)
		if err != nil {
			if err == nats.ErrTimeout {
				log.Println("Timeout while waiting for message")
			}
			continue
		}
		log.Println("Received sync message on subject : ", subj1)
		person := &pb.Person{}
		err = proto.Unmarshal(msg.Data, person)
		if err != nil {
			log.Fatal("Error while Unmarshalling proto message : ", err.Error())
		}
		log.Println("Id : ", person.Id, " Name : ", person.Name, " Age : ", person.Age)
	}
}

func replierTest2() {
	var i int32 = 1
	log.Println("Starting subscribing and replying on subject : ", subj2)
	natsClient.Subscribe(subj2, func(msg *nats.Msg) {
		log.Println("Received request : ", string(msg.Data))
		log.Println("Reply subj : ", msg.Reply)
		natsClient.Publish(msg.Reply, []byte("Reply "+strconv.Itoa(int(i))+" on TEST2"))
		i++
	})
}

func main() {
	log.Println("Starting nats server on default port 4222")
	serverOptions := &server.Options{
		Host:           "127.0.0.1",
		Port:           4222,
		MaxPayload:     int32(100 * 1024),       //100 KB
		MaxPending:     int64(10 * 1024 * 1024), //10 MB
		MaxControlLine: int32(4 * 1024),         // total bytes for protocol line including queue group + subject : 4 KB
		NoLog:          false,
		Debug:          true,
		Trace:          true,
		Logtime:        true,
		NoSigs:         true,
	}
	natsServer, err := server.NewServer(serverOptions)
	if err != nil {
		log.Fatal("Error while creating nats server : ", err.Error())
	}
	natsServer.ConfigureLogger()
	go natsServer.Start()
	testNatsConnect := natsServer.ReadyForConnections(20 * time.Second)
	if !testNatsConnect {
		log.Fatal("Unable to start nats server : ", err.Error())
	} else {
		log.Println("Nats server started on localhost:4222")
	}

	log.Println("Creating nats client ...")
	options := []nats.Option{nats.Name("Test nats Client")}
	natsClient, err = nats.Connect(nats.DefaultURL, options...)
	if err != nil {
		log.Fatal("Error while establishing nats client connection : ", err.Error())
	}
	log.Println("Client connection successful")
	defer natsClient.Close()

	//go subscriberTest1()
	//go subscriberSyncTest1()
	//go queueSubscriber1Test1()
	//go queueSubscriber2Test1()

	//publisheTest1()

	go replierTest2()
	time.Sleep(10 * time.Millisecond)
	requesterTest2()
}
