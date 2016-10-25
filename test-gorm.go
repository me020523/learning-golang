package main

import (
	"fmt"
	"log"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

type GenderType string

const (
	Male   GenderType = "M"
	Female GenderType = "F"
)

type Employee struct {
	EmpNo     int        `gorm:"primary_key"`
	BirthDate time.Time  `gorm:"not null"`
	FirstName string     `gorm:"type:varchar(14);not null"`
	LastName  string     `gorm:"type:varchar(16);not null"`
	Gender    GenderType `gorm:"type:ENUM('M','F');not null"`
	HireDate  time.Time  `gorm:"not null"`
}

func initConnection(driver, source string) error {
	var err error = nil
	db, err = gorm.Open(driver, source)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	err := initConnection("mysql", "root:sbc@241553@tcp(bmx.it-cloud.cn:3306)/employees?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Fatal(err.Error())
	}
	emp := Employee{}
	db.Where("emp_no = ?", 141239).Find(&emp)
	fmt.Print(emp)
}
