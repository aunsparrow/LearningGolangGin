package Services

import (
	Model "api/Migration"
	Repo "api/Repositories/TestRepository"
	Repo2 "api/Repositories/TestRepository2"
	"fmt"
)

func TestService() ([]Model.User, error) {
	testData, err := Repo.GetAll()
	_ = testData
	if err != nil {
		fmt.Println(err.Error())
	}
	testData2, err := Repo2.GetAll()
	fmt.Println(testData2)
	return testData, err
}
