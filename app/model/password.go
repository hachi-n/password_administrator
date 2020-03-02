package model

import (
	"fmt"
	"github.com/hachi-n/passwd_gen/cli"
	randp "github.com/hachi-n/passwd_gen/lib/random_password"
	"log"
	"strconv"
	"strings"
)

func init() {
	createPasswordTable()
}

const passwordTableName = "passwords"

type Password struct {
	ID         int
	ServiceID  int
	CategoryID int
	Password   string
}

func NewPassword(serviceName, categoryName string) *Password {
	service := FindOrCreateByService(serviceName)
	category := FindOrCreateByCategory(categoryName)

	password := new(Password)
	password.ServiceID = service.ID
	password.CategoryID = category.ID
	password.Password = randp.GeneratePassword(
		cli.Option.Length,
		cli.Option.NumberLength,
		cli.Option.SymbolLength,
	)

	return password
}

func (p *Password) String() string {
	var idStr string
	idStr = strconv.Itoa(p.ID)
	if idStr == "0" {
		idStr = "New Password"
	}

	line := fmt.Sprintf("%s %s %s\n",
		strings.Repeat("*", 25),
		idStr,
		strings.Repeat("*", 25),
	)
	return line + fmt.Sprintf("Service:  %s\n"+"Category: %s \n"+"Password: %s\n",
		SearchServicesById(p.ServiceID).Name,
		SearchCategoriesById(p.CategoryID).Name,
		p.Password)
}

func (p *Password) Save() {
	CreatePassword(p.ServiceID, p.CategoryID, p.Password)
}

func SearchPasswords() []*Password {
	cmd := fmt.Sprintf(`SELECT * FROM %s`, passwordTableName)
	rows, err := DBConnection.Query(cmd)
	defer rows.Close()
	if err != nil {
		log.Fatalf("password search query error %v \n", err)
	}

	var passwords []*Password
	for rows.Next() {
		password := new(Password)
		rows.Scan(&password.ID, &password.ServiceID, &password.CategoryID, &password.Password)
		passwords = append(passwords, password)
	}
	return passwords
}

func CreatePassword(serviceId, categoryId int, password string) *Password {
	cmd := fmt.Sprintf(
		`INSERT INTO %s (service_id, category_id, password) VALUES (?, ?, ?)`,
		passwordTableName,
	)

	result, err := DBConnection.Exec(cmd, serviceId, categoryId, password)
	if err != nil {
		log.Fatalf("password query error. %v \n", err)
	}

	resultId, _ := result.LastInsertId()

	passwd := new(Password)
	passwd.ID = int(resultId)
	passwd.ServiceID = serviceId
	passwd.CategoryID = categoryId
	passwd.Password = password

	return passwd
}
