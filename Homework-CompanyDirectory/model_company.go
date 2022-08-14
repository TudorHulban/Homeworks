package company

import (
	"errors"
	"strconv"
	"strings"
)

type Employee struct {
	FullName string
	Manager  *Employee
	ID       int
}

type Company struct {
	team []*Employee
}

var (
	errNoEmployee = errors.New("no employee(s)")
	errNoManager  = errors.New("no manager")
)

const newLine = " \n"

func NewCompany() *Company {
	return &Company{
		team: []*Employee{},
	}
}

// Add Method should add CEO first.
func (c *Company) Add(employee *Employee) error {
	if employee.Manager == nil && len(c.team) != 0 {
		return errNoManager
	}

	c.team = append(c.team, employee)
	employee.ID = len(c.team)

	return nil
}

// Hierarchy Method to return lineage up to root employee.
func (c *Company) Hierarchy(empl *Employee) []*Employee {
	var res []*Employee
	currentEmpl := empl

	for currentEmpl.Manager != nil {
		res = append(res, currentEmpl)

		currentEmpl = currentEmpl.Manager
	}

	return res
}

// FirstLineage Method to return first common manager of the two passed employees.
func (c *Company) FirstManager(e1, e2 *Employee) (*Employee, error) {
	if e1 == nil || e2 == nil {
		return nil, errNoEmployee
	}

	if len(c.team) == 0 {
		return nil, errors.New("build team first")
	}

	if e1 == e2 {
		return e1.Manager, nil
	}

	h1 := c.Hierarchy(e1)
	h2 := c.Hierarchy(e2)

	for i := len(h1) - 1; i >= 0; i-- {
		for j := len(h2) - 1; j >= 0; j-- {
			if h2[j] == h1[i] {
				return h1[i], nil
			}
		}
	}

	return c.team[0], errors.New("common manager is CEO")
}

// List Helper to inspect data.
func List(empl ...*Employee) ([]string, error) {
	if len(empl) == 0 {
		return nil, errors.New("no employees passed")
	}

	list := []string{strings.Join([]string{"ID", "Full Name", "Manager", newLine}, "|")}

	for _, emp := range empl {
		var manager string

		if emp.Manager != nil {
			manager = emp.Manager.FullName
		}

		line := strings.Join([]string{strconv.Itoa(emp.ID), emp.FullName, manager, newLine}, "|")
		list = append(list, line)
	}

	return list, nil
}
