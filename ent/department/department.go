// Code generated by ent, DO NOT EDIT.

package department

const (
	// Label holds the string label denoting the department type in the database.
	Label = "department"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// EdgeEmployees holds the string denoting the employees edge name in mutations.
	EdgeEmployees = "employees"
	// Table holds the table name of the department in the database.
	Table = "departments"
	// EmployeesTable is the table that holds the employees relation/edge.
	EmployeesTable = "employees"
	// EmployeesInverseTable is the table name for the Employee entity.
	// It exists in this package in order to avoid circular dependency with the "employee" package.
	EmployeesInverseTable = "employees"
	// EmployeesColumn is the table column denoting the employees relation/edge.
	EmployeesColumn = "department_employees"
)

// Columns holds all SQL columns for department fields.
var Columns = []string{
	FieldID,
	FieldName,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}
