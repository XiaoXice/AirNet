package database

import (
	"database/sql"
	"fmt"
	"github.com/XiaoXice/AirNet/common/logger"
	_ "github.com/mattn/go-sqlite3"
)

type DataBase struct {
	Type   string
	Path   string
	Logger logger.Logger
	db     *sql.DB
}

func (d *DataBase) Connect() bool {
	db, err := sql.Open(d.Type, d.Path)
	if err != nil {
		d.Logger.Handler(fmt.Sprintf("Can't connect %s on %s", d.Type, d.Path))
		return false
	}
	sqlTable := `
	CREATE TABLE IF NOT EXISTS Config(
	  HashInfo BLOB NOT NULL PRIMARY KEY ,
	  PublicKey BLOB NOT NULL ,
	  PrivateKey BLOB NOT NULL ,
	  Description TEXT NULL ,
	  Type INTEGER NOT NULL 
    );
	CREATE TABLE IF NOT EXISTS Port(
	  LocalPort INTEGER NOT NULL PRIMARY KEY ,
	  RemotePort INTEGER NOT NULL ,
	  RemoteAddress TEXT NOT NULL 
	);
	CREATE TABLE IF NOT EXISTS NodeList(
	  HashInfo BLOB NOT NULL PRIMARY KEY ,
	  Address BLOB NOT NULL ,
	  Port INTEGER NOT NULL ,
	  NeedHole INTEGER NOT NULL ,
	  CanConnection INTEGER NOT NULL ,
	  Description TEXT NULL ,
	  Type INTEGER NOT NULL ,
	  PublicKey BLOB NOT NULL ,
	  LastCheckTime TIMESTAMP NOT NULL
	);
	CREATE TABLE IF NOT EXISTS Router(
	  Remote BLOB NOT NULL ,
	  Next BLOB NOT NULL ,
	  Weight FLOAT NOT NULL ,
	  PRIMARY KEY (Remote,Next)
	);
`
	_, err = db.Exec(sqlTable)
	if err != nil {
		d.Logger.Handler("Init database error: " + err.Error())
		return false
	}
	d.db = db
	return true
}

func (d *DataBase) CanUse() bool {
	if d.db == nil {
		return false
	} else if d.db.Ping() != nil {
		return false
	} else {
		return true
	}
}
