package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/Protobuf_GO/src/complexpb"
	"github.com/Protobuf_GO/src/enumpb"
	simplepb "github.com/Protobuf_GO/src/simple"
	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
)

// Play with nested message and lists
func doComplex() {
	sm := complexpb.ComplexMessage{
		Dummy: &complexpb.DummyMessage{
			Id:   12,
			Name: "Hello Dummy",
		},
		List: []*complexpb.DummyMessage{
			&complexpb.DummyMessage{
				Id:   13,
				Name: "Hello Dummy 2",
			},
			&complexpb.DummyMessage{
				Id:   14,
				Name: "Hello Dummy 3",
			},
		},
	}

	fmt.Println(sm)
}

// Play with the enums in proto
func doEnum() {
	em := enumpb.EnumMessage{
		Id:           41,
		DayOfTheWeek: enumpb.DayOfTheWeek_SUNDAY,
	}
	em.DayOfTheWeek = enumpb.DayOfTheWeek_MONDAY
	fmt.Println(em)
}

// Proto file to json  (Generally used for debugging in json)
func toJSON(pb proto.Message) string {
	marshaler := jsonpb.Marshaler{}
	out, err := marshaler.MarshalToString(pb)
	if err != nil {
		log.Fatalln(err)
		return ""
	}

	return out
}

// JSON to Proto
func fromJSON(in string, pb proto.Message) {
	err := jsonpb.UnmarshalString(in, pb)
	if err != nil {
		log.Fatalln(err)
	}
}

// JSON play
func jsonDemo(sm proto.Message) {
	smAsString := toJSON(sm)
	fmt.Println(smAsString)

	sm2 := &simplepb.SimpleMessage{}
	fromJSON(smAsString, sm2)
	fmt.Println("Successfully created proto struct:", sm2)
}

// Play with the proto message objects
func doSimple() *simplepb.SimpleMessage {
	sm := simplepb.SimpleMessage{
		Id:         12345,
		IsSimple:   true,
		Name:       "My Simple Message",
		SampleList: []int32{1, 2, 3, 4},
	}
	fmt.Println(sm)
	sm.Name = "Rename it!"

	fmt.Println(sm)

	fmt.Println("The Id is: ", sm.GetId())
	return &sm
}

func writeToFile(fName string, pb proto.Message) error {
	out, err := proto.Marshal(pb)
	if err != nil {
		log.Fatalln("Can't serialize to bytes", err)
		return err
	}
	if err := ioutil.WriteFile(fName, out, 0644); err != nil {
		log.Fatalln("Can't write to disk", err)
		return err
	}
	fmt.Println("Data has been written")
	return nil
}

func readFromFile(fName string, pb proto.Message) error {
	in, err := ioutil.ReadFile(fName)
	if err != nil {
		log.Fatalln("Unable to read from file or wrong file name", err)
		return err
	}
	if err2 := proto.Unmarshal(in, pb); err2 != nil {
		log.Fatalln("Unable to convert []byte to --> proto.message", err2)
		return err2
	}
	return nil
}

// Play with read and write
func readAndWrite(sm proto.Message) {
	// writeToFile()
	writeToFile("simple.bin", sm)
	//readFromFile)
	sm2 := simplepb.SimpleMessage{} // This is an empty proto message object
	readFromFile("simple.bin", &sm2)
	fmt.Println(sm2)
}

func main() {
	sm := doSimple()

	readAndWrite(sm)
	jsonDemo(sm)
	doEnum()
	doComplex()
}
