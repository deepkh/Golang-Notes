package greetings

import (
    "errors"
    "fmt"
    "math/rand"
    "time"
    "log"
    "runtime"
    _"path"
)

func Trace(format string, a ...interface{}) {
  function, _, line, _ := runtime.Caller(1)
  info := fmt.Sprintf("[%s:%d]",runtime.FuncForPC(function).Name(), line)
  msg := fmt.Sprintf(format, a...)
  log.Println(info, msg)
}

// Hello returns a greeting for the named person.
func Hello(name string) (string, error) {
    // If no name was given, return an error with a message.
    if name == "" {
        return "", errors.New("empty name")
    }

    // If a name was received, return a value that embeds the name 
    // in a greeting message.
    message := fmt.Sprintf("Hi, %v. Welcome!", name)
    return message, nil
}

//Go executes init functions automatically at program startup
// init sets initial values for variables used in the function.
func init() {
    rand.Seed(time.Now().UnixNano())
    Trace("");
}

// randomFormat returns one of a set of greeting messages. The returned
// message is selected at random.
func randomFormat() string {
    // A slice of message formats.
    formats := []string{
        "Hi, %v. Welcome!",
        "Great to see you, %v!",
        "Hail, %v! Well met!",
    }

    // Return a randomly selected message format by specifying
    // a random index for the slice of formats.
    return formats[rand.Intn(len(formats))]
}

// Hellos returns a map that associates each of the named people
// with a greeting message.
func Hellos(names []string) (map[string]string, error) {
    // A map to associate names with messages.
    messages := map[string]string{}
    // Loop through the received slice of names, calling
    // the Hello function to get a message for each name.
    for _, name := range names {
        message, err := Hello(name)
        if err != nil {
            return nil, err
        }
        // In the map, associate the retrieved message with 
        // the name.
        messages[name] = message
    }
    return messages, nil
}

