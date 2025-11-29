package hello

import "fmt"

const (
	English = "English"
	French  = "French"
	Spanish = "Spanish"

	EnglishHelloPrefix = "Hello,"
	FrenchHelloPrefix  = "Bonjour,"
	SpanishHelloPrefix = "Hola,"
)

func Hello(language, name string) string {
	if len(name) == 0 {
		name = defaultName(language)
	}
	return fmt.Sprintf("%s %s!", helloPrefix(language), name)
}

func helloPrefix(language string) string {
	switch language {
	case English:
		return EnglishHelloPrefix
	case French:
		return FrenchHelloPrefix
	case Spanish:
		return SpanishHelloPrefix
	default:
		return EnglishHelloPrefix
	}
}

func defaultName(language string) string {
	switch language {
	case English:
		return "World"
	case French:
		return "Le monde"
	case Spanish:
		return "Mundo"
	default:
		return "World"
	}
}
