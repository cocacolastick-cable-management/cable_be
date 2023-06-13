package env

type Env struct {
	DbDsn string `env:"DB_DSN"`

	JwtSecret string `env:"JWT_SECRET"`

	SmtpEmail    string `env:"SMTP_EMAIL"`
	SmtpHost     string `env:"SMTP_HOST"`
	SmtpPort     string `env:"SMTP_PORT"`
	SmtpPassword string `env:"SMTP_PASSWORD"`
}
