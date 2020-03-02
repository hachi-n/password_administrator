package model

import (
	"fmt"
	"log"
)

func init() {
	createServiceTable()
}

const serviceTableName = "services"

type Service struct {
	ID   int
	Name string
}

func FindOrCreateByService(serviceName string) *Service {
	services := SearchServicesByName(serviceName)
	if services != nil {
		return services[0]
	}
	service := CreateServices(serviceName)
	return service
}

func SearchServicesByName(serviceName string) []*Service {
	cmd := fmt.Sprintf(`SELECT * FROM %s WHERE name = ?`, serviceTableName)
	rows, err := DBConnection.Query(cmd, serviceName)
	defer rows.Close()
	if err != nil {
		log.Fatalf("service search query error. %v \n", err)
	}

	var services []*Service
	for rows.Next() {
		service := new(Service)
		rows.Scan(&service.ID, &service.Name)
		services = append(services, service)
	}
	return services
}

func SearchServicesById(id int) *Service {
	cmd := fmt.Sprintf(`SELECT * FROM %s WHERE id = ?`, serviceTableName)
	rows, err := DBConnection.Query(cmd, id)
	defer rows.Close()
	if err != nil {
		log.Fatalf("service search query error. %v \n", err)
	}

	var services []*Service
	for rows.Next() {
		service := new(Service)
		rows.Scan(&service.ID, &service.Name)
		services = append(services, service)
	}
	return services[0]
}

func CreateServices(serviceName string) *Service {
	cmd := fmt.Sprintf(`INSERT INTO %s (name) VALUES ( ? )`, serviceTableName)
	result, err := DBConnection.Exec(cmd, serviceName)
	if err != nil {
		log.Fatalf("service create query error. %v \n", err)
	}

	resultId, _ := result.LastInsertId()

	service := new(Service)
	service.ID = int(resultId)
	service.Name = serviceName

	return service
}
