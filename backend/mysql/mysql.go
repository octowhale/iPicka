package mysql

import (
	"fmt"

	log "github.com/sirupsen/logrus"

	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type MysqlAgent struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	pool     *sqlx.DB
}

func NewMysqlAgent(host, port, user, password, dbname string) (*MysqlAgent, error) {

	return &MysqlAgent{
		Host:     host,
		Port:     port,
		User:     user,
		Password: password,
		DBName:   dbname,
	}, nil
}

func (my *MysqlAgent) Conn() *sqlx.DB {

	if my.pool != nil {
		mycfg := &mysql.Config{
			User:   my.User,
			Passwd: my.Password,
			Addr:   fmt.Sprintf("%s:%s", my.Host, my.Port),
			Net:    "tcp",
			DBName: my.DBName,
		}
		// mydsn := mycfg.FormatDSN()

		// mydsn := (&mysql.Config{}).FormatDSN()
		// myCnf.FormatDSN()
		var err error
		my.pool, err = sqlx.Connect("mysql", mycfg.FormatDSN())
		if err != nil {
			log.Errorln(err)
			panic(err)
		}
	}
	return my.pool
}

type TableURL struct {
	Md5 string `db:md5`
	Url string `db:"url"`
}

func (my *MysqlAgent) Get(k string) (s string, err error) {

	return
}

func (my *MysqlAgent) Set(k, v string) (ok bool, err error) {

	return
}
