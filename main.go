package main
import(
	"strconv"
	"net/http"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type User struct{
	Name string `json:"name"`
	Email string `json:"email"`
	ClassCollection []Class `json:"class"`
}

type UserCollection struct{
	Users []User `json:"data"`
}

type Class struct{
	Id int `json:"id"`
	Class string `json:"class"`
}


func main(){
	e := echo.New()
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	e.GET("/", func(c echo.Context) error{
		userCollect := UserCollection{}
		
		var dataClass []Class
		for a := 0; a < 3; a++{
			class := Class{Id: a, Class: "Class " + strconv.Itoa(a)}
			dataClass = append(dataClass, class)
		}

		for i := 0; i < 10; i++{

			user := User{Name: "A " + strconv.Itoa(i), Email: "damasukmakd@gmail.com", ClassCollection: dataClass}
			userCollect.Users = append(userCollect.Users, user) 
		}
		
		if err := c.Bind(&userCollect); err != nil{
			return err
		}
		
		data := userCollect.Users
		return c.JSON(http.StatusCreated, data)
	})	
	e.Logger.Fatal(e.Start(":1234"))
}