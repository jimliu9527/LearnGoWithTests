package hello

const (
	englishHelloPrefix = "Hello "
	spanishHelloPrefix = "Hola "
	frenchHelloPrefix  = "Bonjour "
	spanish            = "Spanish"
	french             = "French"
	defaultName        = "World"
)

func Hello(name string, language string) string {
	if name == "" {
		name = defaultName
	}
	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}
