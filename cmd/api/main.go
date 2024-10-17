package main

import (
	"context"
	"database/sql"
	"flag"
	"log/slog"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/joho/godotenv"
	"juanmagc99.comic-chest/internal/data"
	"juanmagc99.comic-chest/internal/mailer"
)

const version = "1.0.0"

type config struct {
	port   int
	domain string
	env    string
	db     struct {
		dsn string
	}
	smtp struct {
		host     string
		port     int
		username string
		password string
		sender   string
	}
	cors struct {
		trustedOrigins []string
	}
}

type application struct {
	config config
	logger *slog.Logger
	models data.Models
	mailer mailer.Mailer
}

func main() {

	var cfg config

	godotenv.Load()

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	smtpport, err := strconv.Atoi(os.Getenv("SMTP_PORT"))
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	flag.IntVar(&cfg.port, "port", 4000, "API server port")
	flag.StringVar(&cfg.domain, "domain", "localhost", "API server port")
	flag.StringVar(&cfg.env, "env", "development", "Environment (development|staging|production)")
	flag.StringVar(&cfg.db.dsn, "db-dsn", os.Getenv("DB_DSN"), "PostgreSQL DSN")

	flag.StringVar(&cfg.smtp.host, "smtp-host", os.Getenv("SMTP_HOST"), "SMTP host")
	flag.IntVar(&cfg.smtp.port, "smtp-port", smtpport, "SMTP port")
	flag.StringVar(&cfg.smtp.username, "smtp-username", os.Getenv("SMTP_USERNAME"), "SMTP username")
	flag.StringVar(&cfg.smtp.password, "smtp-password", os.Getenv("SMTP_PASSWORD"), "SMTP password")
	flag.StringVar(&cfg.smtp.sender, "smtp-sender", os.Getenv("SMTP_SENDER"), "SMTP sender")

	flag.Func("cors-trusted-origins", "Trusted CORS origins (space separated)", func(val string) error {
		cfg.cors.trustedOrigins = strings.Fields(val)
		return nil
	})

	flag.Parse()

	db, err := openDB(cfg)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	app := &application{
		config: cfg,
		logger: logger,
		models: data.NewModels(db),
		mailer: mailer.New(cfg.smtp.host, cfg.smtp.port, cfg.smtp.username, cfg.smtp.password, cfg.smtp.sender),
	}

	err = app.server()
	logger.Error(err.Error())
	os.Exit(1)
}

func openDB(cfg config) (*sql.DB, error) {
	db, err := sql.Open("postgres", cfg.db.dsn)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}
