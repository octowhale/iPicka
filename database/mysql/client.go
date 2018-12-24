package mysql

import (
	"fmt"

	"github.com/sirupsen/logrus"

	"github.com/go-sql-driver/mysql"
	// _ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type MysqlAgent struct {
	User     string
	Password string
	Host     string
	Port     string
	DBName   string
}

func NewMysqlAgent(user, password, host, port, dbname string) (*MysqlAgent, error) {

	agent := &MysqlAgent{
		Host:     host,
		Port:     port,
		DBName:   dbname,
		User:     user,
		Password: password,
	}
	return agent, nil
}

func (m *MysqlAgent) InitClient() (*sqlx.DB, error) {
	// user:password@tcp(127.0.0.1:3306)/test"

	// mydsn := "root:SMdemT2Pm@tcp(172.18.8.88:60333)/demo"

	// c := &mysql.Config{User: m.User}
	c := mysql.Config{
		User:                 m.User,
		Passwd:               m.Password,
		Addr:                 fmt.Sprintf("%s:%s", m.Host, m.Port),
		DBName:               m.DBName,
		Net:                  "tcp",
		AllowNativePasswords: true,
	}

	mydsn := c.FormatDSN()
	logrus.Debugln(mydsn)

	// mydsn := fmt.Sprintf("%s:%s@tcp(%s:%v)/%s", m.User, m.Password, m.Host, m.Port, m.DBName)

	db, err := sqlx.Connect("mysql", c.FormatDSN())
	if err != nil {
		return nil, err
	}

	err = m.InitDB(db)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func (m *MysqlAgent) InitDB(db *sqlx.DB) error {
	// db, _ := m.InitClient()
	// defer db.Close()
	stmt := `CREATE TABLE IF NOT EXISTS url (
		md5 varchar(32) NOT NULL COMMENT 'file md5sum',
		url varchar(255) NOT NULL COMMENT 'url value',
		PRIMARY KEY (md5),
		UNIQUE KEY md5 (md5) USING HASH COMMENT 'md5 index'
	  ) ENGINE=InnoDB DEFAULT CHARSET=utf8;`

	_, err := db.Exec(stmt)

	if err != nil {
		logrus.Errorln(err)
		return err
	}

	return nil
}

type FileURL struct {
	Md5 string `db:"md5"`
	Url string `db:"url"`
}

func (m *MysqlAgent) Get(k string) (s string, err error) {

	stmt := fmt.Sprintf("SELECT url FROM url WHERE md5='%s';", k)
	// logrus.Debugln("stmt = ", stmt, k)

	db, _ := m.InitClient()
	// defer db.Close()

	fileurl := FileURL{}

	rows := db.QueryRowx(stmt)
	err = rows.StructScan(&fileurl)
	if err != nil {
		// logrus.Errorln("Rows.StructScan Error:", err)
		return "", err
	}

	// logrus.Infoln(fileurl.Url)
	return fileurl.Url, nil
}

func (m *MysqlAgent) Set(k, v string) (err error) {

	stmt := fmt.Sprintf("INSERT INTO url(md5, url) VALUES ('%s', '%s');", k, v)

	db, _ := m.InitClient()

	rows, err := db.Exec(stmt)
	if err != nil {
		// logrus.Errorln("db.Exec Error:", err)
		return err
	}

	_, err = rows.RowsAffected()
	if err != nil {
		// logrus.Errorln("rows.RowsAffected Error:", err)
		return err
	}

	return nil
}
