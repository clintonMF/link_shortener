package utils

import (
	"fmt"
	"strings"
)

func CreatedMessage(entity string) string {
	return fmt.Sprintf("%s has been created", entity)
}

func DeletedMessage(entity string, id int) string {
	return fmt.Sprintf("%s with id %d has been deleted", entity, id)
}
func UpdatedMessage(entity string, id int) string {
	return fmt.Sprintf("%s with id %d has been updated", entity, id)
}
func DisenrollMessage(studentName string, courseName string) string {
	return fmt.Sprintf("%s has been disenroled from %s", strings.Title(studentName), courseName)
}
func EnrollMessage(studentName string, courseName string) string {
	return fmt.Sprintf("%s has been enrolled into %s", strings.Title(studentName), courseName)
}
