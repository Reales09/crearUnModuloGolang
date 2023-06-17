package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

// Hello devuelve un saludo para la persona especificada

func Hello(name string) (string, error) {

	if name == "" {
		return name, errors.New("Nombre vacio")
	}

	message := fmt.Sprintf(randomFormat(), name)

	return message, nil
}

//Hellos returns a map that associate  each of the named people
//with a greeting message

func Hellos(names []string) (map[string]string, error) {
	messages := make(map[string]string)

	for _, name := range names {
		message, err := Hello(name)

		if err != nil {
			return nil, err
		}
		messages[name] = message
	}

	return messages, nil

}

func randomFormat() string {
	formats := []string{
		"¡Hola, %v! ¡Bienvenido!",
		"Que bueno verte, %v!",
		"¡Saludo, %v! ¡Encantado de conocerte!",
	}

	return formats[rand.Intn(len(formats))]
}
