package main

import (
	"encoding/json"
	"log"
)

type Message struct {
	Name string
	Body string
	Time int64
}
type SubMessage struct {
	SName string `json,yaml:"name"`
	SBody string `json:"body"`
	STime int    `json:"time"`
}
type Animal interface {
	Say()
}

type HumanAnimal struct {
	Name string
	Age  int
}

func (human HumanAnimal) Say() {
	log.Println("I am a human")
}

type DogAnimal struct {
	Name string
}

func (dog DogAnimal) Say() {
	log.Println("I am a dog")
}

func UnmarshalToInterface() {
	jsonString := `
        [
            {
                "Name": "Ted",
                "Age": 12
            }
,
            {
                "Name": "Marry",
                "Age": 11
            }
        ]
    `
	var animals []*HumanAnimal
	err := json.Unmarshal([]byte(jsonString), &animals)
	if err != nil {
		log.Fatal(err)
	}

	for _, animal := range animals {
		animal.Say()
	}
}

func Marshal() {
	m := Message{"Alice", "Hello", 1294706395881547000}
	b, err := json.Marshal(m)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(string(b))
}
func Unmarshal() {
	jsonString := []byte(`{"name":"Bob", "body":"Hello", "time":1294706395881547000, "other":"world"}`)
	var m Message
	err := json.Unmarshal(jsonString, &m)
	if err != nil {
		log.Fatal(err)
	}
	b, err := json.Marshal(m)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(string(b))
}
func UnmarshalMap() {
	jsonString := []byte(`{"name":"Bob", "body":"Hello", "time":1294706395881547000, "other":"world"}`)
	var m map[string]interface{}
	err := json.Unmarshal(jsonString, &m)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%v\n", m)
}
func SubUnmarshal() {
	jsonString := []byte(`{"name":"Bob", "body":"Hello", "time":1294706395881547000, "other":"world"}`)
	var m SubMessage
	err := json.Unmarshal(jsonString, &m)
	if err != nil {
		log.Fatal(err)
	}
	b, err := json.Marshal(m)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(string(b))
}
func main() {
	UnmarshalMap()
}
