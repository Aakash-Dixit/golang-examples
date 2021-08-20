package main

import (
	"examples/protobuf/pb"
	"log"

	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
)

func main() {
	p1 := &pb.Person{
		Id:    1234,
		Name:  "John",
		Email: "john@example.com",
		Phones: []*pb.Person_PhoneNumber{
			{
				Number: "555-4321",
				Type:   pb.Person_HOME,
			},
		},
		LastUpdated: ptypes.TimestampNow(),
	}
	p2 := &pb.Person{
		Id:    5678,
		Name:  "Jack",
		Email: "jack@example.com",
		Phones: []*pb.Person_PhoneNumber{
			{
				Number: "666-4321",
				Type:   pb.Person_WORK,
			},
		},
		LastUpdated: ptypes.TimestampNow(),
	}
	people := []*pb.Person{p1, p2}
	addressBook := &pb.AddressBook{People: people}
	addressBookBytes, err := proto.Marshal(addressBook)
	if err != nil {
		log.Fatal("Error while marshalling proto obj : ", err.Error())
		return
	}

	newAddressBook := &pb.AddressBook{}
	err = proto.Unmarshal(addressBookBytes, newAddressBook)
	newPeople := newAddressBook.GetPeople()
	for _, newPerson := range newPeople {
		log.Println("Name : ", newPerson.GetName())
		log.Println("ID : ", newPerson.GetId())
		log.Println("Email : ", newPerson.GetEmail())
		log.Println("LastUpdated : ", newPerson.GetLastUpdated().GetSeconds())
		phoneNumbers := newPerson.GetPhones()
		for _, phNo := range phoneNumbers {
			log.Println("Number : ", phNo.GetNumber())
			log.Println("Type : ", phNo.GetType())
		}
	}
}
