package env

type Env struct {
	EventsDbUri  string `env:"DB_URI" required:"true"`
	EventsDbName string `env:"DB_NAME" required:"true"`
	NatsURL      string `env:"NATS_URL" required:"true"`
}
