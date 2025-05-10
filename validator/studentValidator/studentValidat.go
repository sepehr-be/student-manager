package studentvalidator

import (
	"fmt"
	"reflect"
	"student-app/dto"
)

func CourseValidator(c string) error {
	if reflect.TypeOf(c).Kind() == reflect.String {
		return nil
	}
	return fmt.Errorf("invalid inpute")
}

func StudentValidator(req dto.NewStudentRequest) error {
	if req.Name == "" {
		return fmt.Errorf("name cannot be empty")
	}

	if req.Age < 5 || req.Age > 100 {
		return fmt.Errorf("age must be between 5 and 100")
	}

	if req.Grade < 0 || req.Grade > 100 {
		return fmt.Errorf("grade must be between 0 and 100")
	}

	if len(req.Course) == 0 {
		return fmt.Errorf("course list cannot be empty")
	}

	return nil

}
