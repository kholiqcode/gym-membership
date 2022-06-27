package main

import (
	"gym-membership-api/controllers"
	"gym-membership-api/entity"
	"gym-membership-api/service"
	"log"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:@tcp(127.0.0.1:3306)/gym-membership-api?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Database Connection Error!")
	}

	db.AutoMigrate(&entity.User{})
	// gyms, err := gymRepository.FindAll()
	// if err != nil{
	// 	fmt.Println("**************************")
	// 	fmt.Println("* Error Finding Book Record! *")
	// 	fmt.Println("**************************")
	// }

	// for _, models := range gyms{
	// 	fmt.Println("Tittle ", models.Tittle)
	// }

	// gym, err := gymRepository.FindByID(2)
	// if err != nil{
	// 	fmt.Println("**************************")
	// 	fmt.Println("* Error Finding Book Record! *")
	// 	fmt.Println("**************************")
	// }
	// 	fmt.Println("Tittle ", gym.Tittle)

	// gymRequest := models.Users{
	// 	Tittle:      "Maylida Dwi",
	// 	Description: "I love May",
	// 	Price:       "90000",
	// 	Rating:      4,
	// 	Discount:    0,
	// }

	// gymRepository.Create(gym)
	// gymService.Create(gymRequest)
	
	// //CRUD
	// CREATE DATA
	// gym := models.Gym{}
	// gym.Tittle = "Maylida"
	// gym.Price = 90000
	// gym.Discount = 10
	// gym.Rating = 5
	// gym.Description = "Buku diary Maylida"

	// err = db.Create(&gym).Error
	// if err != nil {
	// 	fmt.Println("**************************")
	// 	fmt.Println("* Error Creating Book Record! *")
	// 	fmt.Println("**************************")
	// }

	// READ DATA
	// var gym models.Gym
	// err = db.Debug().First(&gym, 2).Error
	// if err != nil {
	// 	fmt.Println("**************************")
	// 	fmt.Println("* Error Finding Book Record! *")
	// 	fmt.Println("**************************")
	// }

	// READ MANY DATA
	// var gyms []models.Gym
	// err = db.Debug().Find(&gyms).Error
	// if err != nil {
	// 	fmt.Println("**************************")
	// 	fmt.Println("* Error Finding Book Record! *")
	// 	fmt.Println("**************************")
	// }

	// for _, g := range gyms{
	// fmt.Println("Tittle :", g.Tittle)
	// fmt.Printf("Gym Object %v", g)
	// }

	// READ DATA BY CONDITION
	// var gyms []models.Gym
	// err = db.Debug().Where("tittle = ?", "Maylida").Find(&gyms).Error
	// if err != nil {
	// 	fmt.Println("**************************")
	// 	fmt.Println("* Error Finding Book Record! *")
	// 	fmt.Println("**************************")
	// }

	// for _, g := range gyms{
	// fmt.Println("Tittle :", g.Tittle)
	// fmt.Printf("Gym Object %v", g)
	// }

	// UPDATE DATA
	// var gym models.Gym
	// err = db.Debug().Where("id = ?", 1).First(&gym).Error
	// if err != nil {
	// 	fmt.Println("**************************")
	// 	fmt.Println("* Error Finding Book Record! *")
	// 	fmt.Println("**************************")
	// }

	// gym.Tittle = "Al Tsaqif"
	// gym.Description = "Buku diary Al Tsaqif"
	// err = db.Save(&gym).Error
	// if err != nil {
	// 	fmt.Println("**************************")
	// 	fmt.Println("* Error Updating Book Record! *")
	// 	fmt.Println("**************************")
	// }

	// DELETE DATA
	// var gym models.Gym
	// err = db.Debug().Where("id = ?", 1).First(&gym).Error
	// if err != nil {
	// 	fmt.Println("**************************")
	// 	fmt.Println("* Error Finding Book Record! *")
	// 	fmt.Println("**************************")
	// }

	// err = db.Delete(&gym).Error
	// if err != nil {
	// 	fmt.Println("**************************")
	// 	fmt.Println("* Error Deleting Book Record! *")
	// 	fmt.Println("**************************")
	// }
	/****************************************/
	// e := echo.New()

	gymRepository := service.NewRepository(db)
	// gymFileRepository := models.NewFileRepository()
	GymService := service.NewService(gymRepository)
	GymControllers := controllers.NewGymControllers(GymService)
	e := gin.Default()
	v1 := e.Group("/v1")

	// Method GET
	// Table Users
	v1.GET("/", GymControllers.GetRootControllers)
	v1.GET("/users/list", GymControllers.GetAllUser)
	v1.GET("/users/:id", GymControllers.GetUserById)

	// Method POST
	// Table Users
	v1.POST("/users", GymControllers.CreateUser_POST)

	// Method PUT
	// Table Users
	v1.PUT("/users/:id", GymControllers.UpdateUserByID_PUT)
	v1.PATCH("/users/:id", GymControllers.UpdateUserByID_PATCH)

	// Method DELETE
	// Table Users
	v1.DELETE("/users/:id", GymControllers.DeleteUserByID)

	e.Run(":8080")
}

//main
//controllers
//service -> bisnis logic
//repository -> database
//db
//mysql


