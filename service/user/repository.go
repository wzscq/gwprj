package user

import (
	"database/sql"
	"github.com/go-sql-driver/mysql"
	"log"
	"time"
)

type UserRole struct {
	ModuleType string
	ResourceType string
	ResourceId string
}

type UserRepository interface {
	GetUserRoles(userID string)([]UserRole)
}

type DefatultUserRepository struct {
	DB *sql.DB
}

func (repo *DefatultUserRepository)GetUserRoles(userID string)([]UserRole){
	var roles []UserRole
	sql:="select MODULETYPE_ as ModuleType,RESOURCETYPE_ as ResourceType,RESOURCEID_ as ResourceId from abi52_eacl_permission where AUTHID_='"+userID+"'"
	log.Println("GetUserRoles sql:"+sql)
	rows, err := repo.DB.Query(sql)
	if err != nil {
		log.Println(err)
		return roles
	}
	
	for rows.Next() {
		var userRole UserRole
		err:= rows.Scan(&userRole.ModuleType,&userRole.ResourceType,&userRole.ResourceId)
		if err != nil {
			log.Println(err)
			return roles
		}
		roles=append(roles,userRole)
	}
	return roles
}

func (repo *DefatultUserRepository)Connect(
	server,user,password,dbName string,
	connMaxLifetime,maxOpenConns,maxIdleConns int){
	// Capture connection properties.
    cfg := mysql.Config{
        User:   user,
        Passwd: password,
        Net:    "tcp",
        Addr:   server,
        DBName: dbName,
		AllowNativePasswords:true,
    }
    // Get a database handle.
    var err error
    repo.DB, err = sql.Open("mysql", cfg.FormatDSN())
    if err != nil {
        log.Fatal(err)
    }

    pingErr := repo.DB.Ping()
    if pingErr != nil {
        log.Fatal(pingErr)
    }

		repo.DB.SetConnMaxLifetime(time.Minute * time.Duration(connMaxLifetime))
		repo.DB.SetMaxOpenConns(maxOpenConns)
		repo.DB.SetMaxIdleConns(maxIdleConns)
    log.Println("connect to mysql server "+server)
}