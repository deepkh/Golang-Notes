package main

import(
  _"fmt"
  pb "addressbook"
  "strconv"
  "io/ioutil"
  "log"
  "os"
  "github.com/golang/protobuf/proto"
)


func CreateAddressBook(ab *pb.AddressBook, n int)  {
  var i int32;
  for i=0; i<3; i++ {
    person := &pb.Person{
        Id:    i,
        Name:  "John-"+strconv.Itoa(int(i)),
        Email: "jdoe@example.com",
    };

    for j:=0; j<3; j++ {
      person.Phones = append(person.Phones,
        &pb.Person_PhoneNumber{Number: "12345678#"+strconv.Itoa(int(j)), Type: pb.Person_PhoneType(j%3)},
      );
    }

    ab.Person = append(ab.Person, person);
  }
  log.Printf("%v\n", len(ab.Person));
}

type AddressBookLoader struct {
  
};

func NewAddressBookLoader() *AddressBookLoader {
  return &AddressBookLoader {};
}

func (p *AddressBookLoader) Unmarshal(in []byte, ab *pb.AddressBook) error {
  var err error;
  if err = proto.Unmarshal(in, ab); err != nil {
    log.Printf("failed to unmarshal address book: %v\n", err);
  }
  return err;
}

func (p *AddressBookLoader) ReadFromFile(filename string, ab *pb.AddressBook) error {
  var in []byte;
  var err error;

  // Read the existing address book.
  in, err = ioutil.ReadFile(filename);
  if err != nil {
    if os.IsNotExist(err) {
      log.Printf("failed to read file (FileNotFound) %v \n", filename);
    } else {
      log.Printf("failed to read file: %v", err);
    }
    return err;
  }

  return p.Unmarshal(in, ab);
}

func (p *AddressBookLoader) Marshal(ab *pb.AddressBook) ([]byte, error) {
  return proto.Marshal(ab);
}

func (p *AddressBookLoader) WriteToFile(filename string, ab *pb.AddressBook) error {
  var buf []byte;
  var err error = nil;
  
  buf, err = p.Marshal(ab);
  if err != nil {
    return err;
  }

  if err = ioutil.WriteFile(filename, buf, 0644); err != nil {
    log.Fatalln("Failed to write address book:", err);
  }
  return err;
}

func main() {
  ab := &pb.AddressBook{};
  loader := NewAddressBookLoader();
  var filename string = "";

  if len(os.Args) != 2 {
    log.Printf("usage: %v filename\n", os.Args[0]);
    return;
  }

  filename = os.Args[1];

  //read
  if err := loader.ReadFromFile(filename, ab); err != nil {
    log.Printf("create new one\n");
    CreateAddressBook(ab, 10);
  } else {
    log.Printf("file read %v ok\n", filename);
  }

  //dump
  for i:=0; i<len(ab.Person); i++ {
    log.Printf("\n\n%d\n", i);
    log.Printf("\t%v %v\n", ab.Person[i].GetId(), ab.GetPerson()[i]);
  }

  //write
  if err := loader.WriteToFile(filename, ab); err != nil {
    log.Printf("failed to loader.WriteToFile %v\n", err);
  }
}
