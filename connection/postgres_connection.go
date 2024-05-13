package connection

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/joho/godotenv"
)

func config() *pgxpool.Config {
	err := godotenv.Load("../.env.dev")
	if err != nil {
		log.Fatal(err.Error())
	}

	db_driver := os.Getenv("DB_DRIVER")
	db_host := os.Getenv("DB_HOST")
	db_port := os.Getenv("DB_PORT")
	db_user := os.Getenv("DB_USER")
	db_password := os.Getenv("DB_PASSWORD")
	db_name := os.Getenv("DB_NAME")

	const defaultMaxConns = int32(50)
	const defaultMinConns = int32(10)
	const defaultMaxConnLifetime = time.Hour
	const defaultMaxConnIdleTime = time.Minute * 30
	const defaultHealthCheckPeriod = time.Minute
	const defaultConnectTimeout = time.Second * 5

	dsn := fmt.Sprintf("%s://%s:%s@%s:%s/%s", db_driver, db_user, db_password, db_host, db_port, db_name)

	config, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		log.Fatal(err.Error())
	}

	config.MaxConns = defaultMaxConns
	config.MinConns = defaultMinConns
	config.MaxConnLifetime = defaultMaxConnLifetime
	config.MaxConnIdleTime = defaultMaxConnIdleTime
	config.HealthCheckPeriod = defaultHealthCheckPeriod
	config.ConnConfig.ConnectTimeout = defaultConnectTimeout

	return config
}

func GetConnection() (*pgxpool.Pool, error) {
	ctx := context.Background()
	config := config()

	conn, err := pgxpool.Connect(ctx, config.ConnString())
	if err != nil {
		log.Fatal(err.Error())
		return conn, err
	}

	return conn, nil
}
