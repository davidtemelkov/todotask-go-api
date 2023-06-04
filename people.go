package main

import (
	"net/http"

	"errors"

	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Person struct {
	ID   uint
	NAME string
	AGE  uint
}

func getPeople(c *gin.Context) {
	id := c.Query("id")

	if id == "" {
		var people []Person
		result := db.Find(&people)

		if result.Error != nil {
			panic(result.Error)
		}

		c.IndentedJSON(http.StatusOK, people)

		return
	}

	uintNum, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	uintId := uint(uintNum)

	person, err := getPersonByID(uintId)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, person)
}

func createPeople(c *gin.Context) {
	var newPerson Person

	if err := c.BindJSON(&newPerson); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.Create(&newPerson).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusCreated, newPerson)
}

func editPeople(c *gin.Context) {
	var newPerson Person

	if err := c.BindJSON(&newPerson); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var person Person
	person, err := getPersonByID(newPerson.ID)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	person.NAME = newPerson.NAME
	person.AGE = newPerson.AGE

	if err := db.Save(&person).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, person)
}

func getPersonByID(id uint) (Person, error) {
	var person Person
	result := db.First(&person, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return person, errors.New("Person not found")
		}
		return person, result.Error
	}
	return person, nil
}
