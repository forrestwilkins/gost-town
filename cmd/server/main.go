// Used for reference: https://gist.github.com/mashingan/4212d447f857cfdfbbba4f5436b779ac

package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/labstack/echo"
)

type User struct {
	gorm.Model `json:"model"`
	Name       string `json:"name"`
	Email      string `json:"email"`
}

type dbops struct {
	db *gorm.DB
}

func (d dbops) findAll(users *[]User) error {
	return d.db.Find(users).Error
}

func (d dbops) create(user *User) error {
	return d.db.Create(user).Error
}

func (d dbops) findByPage(users *[]User, page, view int) error {
	return d.db.Limit(view).Offset(view * (page - 1)).Find(&users).Error

}

func (d dbops) updateByName(name, email string) error {
	var user User
	d.db.Where("name=?", name).Find(&user)
	user.Email = email
	return d.db.Save(&user).Error
}

func (d dbops) deleteByName(name string) error {
	var user User
	d.db.Where("name=?", name).Find(&user)
	return d.db.Delete(&user).Error
}

func HandlerFunc(msg string) func(echo.Context) error {
	return func(c echo.Context) error {
		return c.String(http.StatusOK, msg)
	}
}

func allUsers(dbobj dbops) func(echo.Context) error {
	return func(c echo.Context) error {
		var users []User
		dbobj.findAll(&users)
		fmt.Println("{}", users)

		return c.JSON(http.StatusOK, users)
	}
}

func newUser(dbobj dbops) func(echo.Context) error {
	return func(c echo.Context) error {
		name := c.Param("name")
		email := c.Param("email")
		dbobj.create(&User{Name: name, Email: email})
		return c.String(http.StatusOK, name+" user successfully created")
	}
}

func deleteUser(dbobj dbops) func(echo.Context) error {
	return func(c echo.Context) error {
		name := c.Param("name")

		dbobj.deleteByName(name)

		return c.String(http.StatusOK, name+" user successfully deleted")
	}
}

func updateUser(dbobj dbops) func(echo.Context) error {
	return func(c echo.Context) error {
		name := c.Param("name")
		email := c.Param("email")
		dbobj.updateByName(name, email)
		return c.String(http.StatusOK, name+" user successfully updated")
	}
}

func usersByPage(dbobj dbops) func(echo.Context) error {
	return func(c echo.Context) error {
		limit, _ := strconv.Atoi(c.QueryParam("limit"))
		page, _ := strconv.Atoi(c.QueryParam("page"))
		var result []User
		dbobj.findByPage(&result, page, limit)
		return c.JSON(http.StatusOK, result)
	}
}

func handleRequest(dbgorm *gorm.DB) {
	e := echo.New()
	db := dbops{dbgorm}

	e.GET("/users", allUsers(db))
	e.GET("/user", usersByPage(db))
	e.POST("/user/:name/:email", newUser(db))
	e.DELETE("/user/:name", deleteUser(db))
	e.PUT("/user/:name/:email", updateUser(db))

	e.Logger.Fatal(e.Start(":3000"))
}

func initialMigration(db *gorm.DB) {

	db.AutoMigrate(&User{})
}

func main() {
	fmt.Println("Go ORM tutorial")
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		fmt.Println(err.Error())
		panic("failed to connect database")
	}
	defer db.Close()
	initialMigration(db)
	handleRequest(db)
}
