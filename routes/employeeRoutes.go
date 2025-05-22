package routes

import (
	"github.com/Golang-Personal-Projects/GolangTutorial/GoMongoDB/database"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// Employee details information
type Employee struct {
	ID     string  `json:"id,omitempty" bson:"_id,omitempty"`
	Name   string  `json:"name"`
	Salary float64 `json:"salary"`
	Age    float64 `json:"age"`
}

// func createResponse(e Employee) Employee{
// 	return Employee{e.ID, e.Name, e.Salary, e.Age}
// }

func GetEmployees(c *fiber.Ctx) error {
	// initialize the slice
	// var employees []Employee = make([]Employee, 0)

	employees := new([]Employee) // returns a pointer to the slice of Employee struct
	// query
	query := bson.D{{}}

	// run the query against the database collection
	cursor, err := database.Mg.Db.Collection("employees").Find(c.Context(), query)

	if err != nil {
		return c.Status(500).JSON(err.Error())
	}

	err = cursor.All(c.Context(), employees)

	if err != nil {
		return c.Status(500).JSON(err.Error())
	}

	return c.Status(200).JSON(employees)
}

func CreateEmployee(c *fiber.Ctx) error {
	collection := database.Mg.Db.Collection("employees")

	var employee Employee //  employee := new(Employee)
	// employee := new(Employee)

	if err := c.BodyParser(&employee); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	//	employee.ID = ""
	insertion, err := collection.InsertOne(c.Context(), employee)
	if err != nil {
		return c.Status(400).JSON(err.Error())
	}

	filter := bson.D{{Key: "_id", Value: insertion.InsertedID}}

	createdEmployee := &Employee{}
	err = collection.FindOne(c.Context(), filter).Decode(createdEmployee)
	if err != nil {
		c.Status(500).JSON("unable to decode employee")
	}

	return c.Status(200).JSON(createdEmployee)
}

func GetEmployee(c *fiber.Ctx) error {
	idparams := c.Params("id")
	employeeID, err := primitive.ObjectIDFromHex(idparams)
	if err != nil {
		return c.Status(400).JSON("Ensure id is an integer")
	}
	var employee Employee
	if err = c.BodyParser(employee); err != nil {
		c.Status(400).JSON(err.Error())
	}
	query := bson.D{{Key: "_id", Value: employeeID}}
	collection := database.Mg.Db.Collection("employees")
	getRecord := collection.FindOne(c.Context(), query)

	getEmployee := &employee

	getRecord.Decode(getEmployee)
	return c.Status(200).JSON(getEmployee)
}

func UpdateEmployee(c *fiber.Ctx) error {
	params := c.Params("id")

	employeeID, err := primitive.ObjectIDFromHex(params)

	if err != nil {
		return c.Status(400).JSON("Ensure id is an integer")
	}

	var employee Employee // employee := new(Employee)

	if err = c.BodyParser(&employee); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	query := bson.D{{Key: "_id", Value: employeeID}}
	update := bson.D{{
		Key: "$set",
		Value: bson.D{
			{Key: "name", Value: employee.Name},
			{Key: "age", Value: employee.Age},
			{Key: "salary", Value: employee.Salary},
		},
	}}

	err = database.Mg.Db.Collection("employees").FindOneAndUpdate(c.Context(), query, update).Err()
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return c.Status(400).JSON("no documents found")
		}
		return c.Status(500).JSON(err.Error())
	}
	employee.ID = params

	return c.Status(200).JSON(employee)
}

func DeleteEmployee(c *fiber.Ctx) error {
	idparams := c.Params("id")

	employeeID, err := primitive.ObjectIDFromHex(idparams)
	if err != nil {
		return c.Status(400).JSON("make sure id is an integer")
	}

	query := bson.D{{Key: "_id", Value: employeeID}}

	collections := database.Mg.Db.Collection("employees")

	deletedEmployee, err := collections.DeleteOne(c.Context(), query)
	if err != nil {
		return c.Status(400).JSON("employee record could not be deleted")
	}
	noDelRecord := deletedEmployee.DeletedCount
	return c.Status(200).JSON(noDelRecord)

}
