#include <iostream>
#include <fstream>
#include <string>
#include <addressbook/addressbook.pb.h>
using namespace std;

static void CreateAddressBook(addressbook::AddressBook *address_book) {
  for (int i=0; i<3; i++) {
    addressbook::Person *person = address_book->add_person();
    person->set_id(i);
    *person->mutable_name() = string("John-")+ std::to_string(i);
    person->set_email(string("JohnMail-")+ std::to_string(i));
    for (int j=0; j<3; j++) {
      addressbook::Person::PhoneNumber* phone_number = person->add_phones();
      phone_number->set_number(string("JohnNumber-")+ std::to_string(i) + std::to_string(j));
      phone_number->set_type(static_cast<addressbook::Person_PhoneType>(j%3));
    }
  }
}

// Main function:  Reads the entire address book from a file,
//   adds one person based on user input, then writes it back out to the same
//   file.
int main(int argc, char* argv[]) {
  // Verify that the version of the library that we linked against is
  // compatible with the version of the headers we compiled against.
  GOOGLE_PROTOBUF_VERIFY_VERSION;

  if (argc != 2) {
    cerr << "Usage:  " << argv[0] << " ADDRESS_BOOK_FILE" << endl;
    return -1;
  }

  //read & unmarshal
  addressbook::AddressBook address_book;
  {
    // Read the existing address book.
    fstream input(argv[1], ios::in | ios::binary);
    if (!input) {
      cout << argv[1] << ": File not found.  Creating a new file." << endl;
      CreateAddressBook(&address_book);
    } else if (!address_book.ParseFromIstream(&input)) {
      cerr << "Failed to parse address book." << endl;
      return -1;
    }
  }
  
  //dump
  for (int i=0; i< address_book.person_size(); i++) {
    const addressbook::Person &person = address_book.person(i);
    printf("\n%d\n", i);
    printf("\tid:%d\n", person.id());
    printf("\tname:%s\n", person.name().c_str());
    printf("\temail:%s\n", person.email().c_str());
    for (int j=0; j<person.phones_size(); j++) {
      const addressbook::Person::PhoneNumber& phone_number = person.phones(j);
      printf("\t\tnumber:%s\n", phone_number.number().c_str());
      printf("\t\ttype:%d\n", phone_number.type());
    }
  }

  //marshal & write 
  {
    // Write the new address book back to disk.
    fstream output(argv[1], ios::out | ios::trunc | ios::binary);
    if (!address_book.SerializeToOstream(&output)) {
      cerr << "Failed to write address book." << endl;
      return -1;
    }
  }

  // Optional:  Delete all global objects allocated by libprotobuf.
  google::protobuf::ShutdownProtobufLibrary();
  return 0;
}

