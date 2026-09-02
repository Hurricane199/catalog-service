package section

type Monitor struct {
	LogLevel   string `split_words:"true" default:"debug"`
	Enviroment string `default:"development"`
}
