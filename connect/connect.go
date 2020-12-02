package connect

import (
	// "gorm.io/gorm"
	// "gorm.io/driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"go_api_rest/structures"
)

var connection *gorm.DB

const engine_sql string = "mysql"
const username 	 string = "root"
const password 	 string = "123456"
const database 	 string = "taller"

func InitializeDatabase(){
	connection = ConnectORM( CreateString())
	log.Println("Open connection")
}

func CloseConnection()  {
	connection.Close()
	log.Println("Closed connection")
}

func ConnectORM(stringConnection string) *gorm.DB {
	connection, err := gorm.Open( engine_sql, stringConnection )
	// gorm.Open(mysql.Open(CreateString()), &gorm.Config{}) 
	if err != nil {
		log.Fatal(err)
		return nil
	}
	return connection	
}

func CreateUser(user structures.User) structures.User {
	connection.Create(&user)
	return user
}

func UpdateUser(id string , user structures.User) structures.User {
	current_user := structures.User{}
	connection.Where("id = ?",id).First(&current_user)
	current_user.Username 	= user.Username
	//current_user.First_Name = user.First_Name
	//current_user.Last_Name 	= user.Last_Name
	connection.Save(&current_user)
	return current_user
}

func DeleteUser(id string) {
	current_user := structures.User{}
	connection.Where("id = ?",id).First(&current_user)
	connection.Delete(&current_user)
}

func CreateString() string {
	var te = username + ":" + password + "@tcp(go_mysql)/" + database
	log.Println(te)
	return te
}

func GetUser(id string) structures.User {
	user := structures.User{}
	connection.Where("id = ?", id).First(&user)
	return user
}

func GetUsers() []structures.User {
	users := []	structures.User{}
	connection.Find(&users)
	return users
}