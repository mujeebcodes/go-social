package main

import (
	"github.com/mujeebcodes/go-social/internal/auth"
	"github.com/mujeebcodes/go-social/internal/db"
	"github.com/mujeebcodes/go-social/internal/env"
	"github.com/mujeebcodes/go-social/internal/mailer"
	"github.com/mujeebcodes/go-social/internal/store"
	"go.uber.org/zap"
	"time"
)

const version = "0.0.1"

//	@title			Go-Social API
//	@description	API for Go-Social, a social network
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@BasePath	/v1
//
//@securityDefinitions.apikey ApiKeyAuth
//@in 	header
//@name Authorization
//@description

func main() {

	cfg := config{
		addr:        env.GetString("PORT", ":8080"),
		apiURL:      env.GetString("EXTERNAL_URL", "localhost:4000"),
		frontendURL: env.GetString("FRONTEND_URL", "localhost:8000"),

		db: dbConfig{
			addr:         env.GetString("DB_ADDR", ""),
			maxOpenConns: env.GetInt("MAX_OPEN_CONNS", 30),
			maxIdleConns: env.GetInt("MAX_IDLE_CONNS", 30),
			maxIdleTime:  env.GetString("MAX_IDLE_TIME", "15m"),
		},
		env: env.GetString("ENV", "development"),
		mail: mailConfig{
			exp:       time.Hour * 24 * 3,
			fromEmail: env.GetString("FROM_EMAIL", ""),
			mailTrap: mailTrapConfig{
				apiKey:          env.GetString("MAIL_TRAP_API_KEY", ""),
				sandboxUsername: env.GetString("MAIL_TRAP_SANDBOX_USERNAME", ""),
				sandboxPassword: env.GetString("MAIL_TRAP_SANDBOX_PASSWORD", ""),
			},
		},
		auth: authConfig{
			basic: basicConfig{
				user: env.GetString("AUTH_BASIC_USER", ""),
				pass: env.GetString("AUTH_BASIC_PASS", ""),
			},
			token: tokenConfig{
				secret: env.GetString("AUTH_TOKEN_SECRET", "example"),
				exp:    time.Hour * 24 * 3,
				iss:    "go-social",
			},
		},
	}

	//Logger
	logger := zap.Must(zap.NewProduction()).Sugar()
	defer logger.Sync()

	//database
	db, err := db.New(cfg.db.addr, cfg.db.maxOpenConns, cfg.db.maxIdleConns, cfg.db.maxIdleTime)
	if err != nil {
		logger.Fatal(err)
	}

	defer db.Close()
	logger.Info("database connection pool established")

	store := store.NewStorage(db)
	mailtrap, err := mailer.NewMailTrapClient(cfg.mail.mailTrap.apiKey, cfg.mail.fromEmail, cfg.mail.mailTrap.sandboxUsername, cfg.mail.mailTrap.sandboxPassword)
	if err != nil {
		logger.Fatal(err)
	}
	
	jwtAuthenticator := auth.NewJWTAuthenticator(cfg.auth.token.secret, cfg.auth.token.iss, cfg.auth.token.iss)

	app := &application{
		config:        cfg,
		store:         store,
		logger:        logger,
		mailer:        mailtrap,
		authenticator: jwtAuthenticator,
	}

	mux := app.mount()
	logger.Fatal(app.run(mux))
}
