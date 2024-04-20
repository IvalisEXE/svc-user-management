package postgresql

import (
	"log"
	"time"

	"github.com/jmoiron/sqlx"

	_ "github.com/lib/pq"
)

// Driver list
const (
	DriverPostgres = "postgres"
)

// Db object
var (
	Master *DB
)

var (
	DefaultRetryInterval int = 5
	DefaultMaxIddleConn  int = 10
	DefaultMaxConn       int = 10
)

type (
	//DSNConfig for database source name
	DSNConfig struct {
		DSN string
	}

	//DBConfig for databases configuration
	DBConfig struct {
		MasterDSN     string `json:"master_dsn" mapstructure:"master_dsn"`
		RetryInterval int    `json:"retry_interval" mapstructure:"retry_interval"`
		MaxIdleConn   int    `json:"max_idle" mapstructure:"max_idle"`
		MaxConn       int    `json:"max_con" mapstructure:"max_con"`
	}

	//DB configuration
	DB struct {
		DBConnection  *sqlx.DB
		DBString      string
		RetryInterval int
		MaxIdleConn   int
		MaxConn       int
		doneChannel   chan bool
	}

	Architecture struct {
		Master *sqlx.DB
	}
)

func (s *Architecture) GetMaster() *sqlx.DB {
	return s.Master
}

func New(cfg DBConfig, dbDriver string) *Architecture {
	masterDSN := cfg.MasterDSN

	retryInterval := cfg.RetryInterval
	maxIddleConn := cfg.MaxIdleConn
	maxConn := cfg.MaxConn

	if retryInterval == 0 {
		retryInterval = DefaultRetryInterval
	}
	if maxIddleConn == 0 {
		maxIddleConn = DefaultMaxIddleConn
	}
	if maxConn == 0 {
		maxConn = DefaultMaxConn
	}

	Master = &DB{
		DBString:      masterDSN,
		RetryInterval: retryInterval,
		MaxIdleConn:   maxIddleConn,
		MaxConn:       maxConn,
		doneChannel:   make(chan bool),
	}

	err := Master.ConnectAndMonitor(dbDriver)
	if err != nil {
		log.Fatal("Could not initiate Master DB connection: " + err.Error())
		return &Architecture{}
	}

	log.Println("Connected to db postgresql successfully")
	return &Architecture{Master: Master.DBConnection}
}

// Connect to database
func (d *DB) Connect(driver string) error {
	var db *sqlx.DB
	var err error

	db, err = sqlx.Open(driver, d.DBString)

	if err != nil {
		log.Println("[Error]: DB open connection error", err.Error())
	} else {
		d.DBConnection = db
		err = db.Ping()
		if err != nil {
			log.Println("[Error]: DB connection error", err.Error())
		}
		return err
	}

	db.SetMaxOpenConns(d.MaxConn)
	db.SetMaxIdleConns(d.MaxIdleConn)

	return err
}

// ConnectAndMonitor to database
func (d *DB) ConnectAndMonitor(driver string) error {
	err := d.Connect(driver)

	if err != nil {
		log.Printf("Not connected to database %s, trying", d.DBString)
		return err
	}

	ticker := time.NewTicker(time.Duration(d.RetryInterval) * time.Second)
	go func() error {
		for {
			select {
			case <-ticker.C:
				if d.DBConnection == nil {
					d.Connect(driver)
				} else {
					err := d.DBConnection.Ping()
					if err != nil {
						log.Println("[Error]: DB reconnect error", err.Error())
						return err
					}
				}
			case <-d.doneChannel:
				return nil
			}
		}
	}()
	return nil
}
