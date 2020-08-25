package config

import (
	"photoalbum/models"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

const (
	DBDriver = "mysql"
	DBUser = "root"
	DBPass = "root"
	//DBPass = ""
	DBName = "albumdb"
)

func GetConnection() (*gorm.DB, error) {
	db, err := gorm.Open(DBDriver, DBUser+":"+DBPass+"@/"+DBName+"?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return db, nil
}

func TestConnection() bool{
	con, err := GetConnection()
	defer con.Close()
	if err != nil {
		log.Fatal("Failed in getting connection!!!")
		return false
	}

	fmt.Println("Connection successfully established")
	return true
}


func InitDB () bool {

	dataSourceName := DBUser + ":" + DBPass+"@tcp(127.0.0.1:3306)/"
	//dataSourceName := DBUser + ":" + DBPass+"@tcp(127.0.0.1:3306)/parseTime=True&loc=Local"
	db, err := gorm.Open(DBDriver, dataSourceName)

	if err != nil {
		log.Println(err)
		log.Fatal("error initalization in database")
		return false
	} else {
		log.Println("database initialized")
	}
	defer db.Close()

	// Create the database. This is a one-time step.
	db.Debug().Exec("DROP DATABASE albumdb")
	db.Debug().Exec("CREATE DATABASE albumdb")
	db.Debug().Exec("USE albumdb")

	//db.SingularTable(true)
	fmt.Println(db.Debug().Table("albums").Value)
	db.Debug().DropTableIfExists(&models.Album{}, &models.Image{}, &models.User{})

	//Create table from defined structrue
	db.Debug().AutoMigrate(&models.User{}, &models.Album{}, &models.Image{})

	//Established relation among tables
	db.Debug().Model(&models.Album{}).AddForeignKey("user_id", "users(user_id)", "CASCADE", "CASCADE")
	db.Debug().Model(&models.Image{}).AddForeignKey("album_id", "albums(album_id)", "CASCADE", "CASCADE")

	db.Debug().Exec("ALTER TABLE images MODIFY imagefile mediumblob not NULL")
	


	return true
}