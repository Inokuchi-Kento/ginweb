// Code generated by ent, DO NOT EDIT.

package employee

import (
	"ginweb/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Employee {
	return predicate.Employee(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Employee {
	return predicate.Employee(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Employee {
	return predicate.Employee(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Employee {
	return predicate.Employee(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Employee {
	return predicate.Employee(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Employee {
	return predicate.Employee(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Employee {
	return predicate.Employee(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Employee {
	return predicate.Employee(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Employee {
	return predicate.Employee(sql.FieldLTE(FieldID, id))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Employee {
	return predicate.Employee(sql.FieldEQ(FieldName, v))
}

// Age applies equality check predicate on the "age" field. It's identical to AgeEQ.
func Age(v int) predicate.Employee {
	return predicate.Employee(sql.FieldEQ(FieldAge, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Employee {
	return predicate.Employee(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Employee {
	return predicate.Employee(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Employee {
	return predicate.Employee(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Employee {
	return predicate.Employee(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Employee {
	return predicate.Employee(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Employee {
	return predicate.Employee(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Employee {
	return predicate.Employee(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Employee {
	return predicate.Employee(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Employee {
	return predicate.Employee(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Employee {
	return predicate.Employee(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Employee {
	return predicate.Employee(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Employee {
	return predicate.Employee(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Employee {
	return predicate.Employee(sql.FieldContainsFold(FieldName, v))
}

// AgeEQ applies the EQ predicate on the "age" field.
func AgeEQ(v int) predicate.Employee {
	return predicate.Employee(sql.FieldEQ(FieldAge, v))
}

// AgeNEQ applies the NEQ predicate on the "age" field.
func AgeNEQ(v int) predicate.Employee {
	return predicate.Employee(sql.FieldNEQ(FieldAge, v))
}

// AgeIn applies the In predicate on the "age" field.
func AgeIn(vs ...int) predicate.Employee {
	return predicate.Employee(sql.FieldIn(FieldAge, vs...))
}

// AgeNotIn applies the NotIn predicate on the "age" field.
func AgeNotIn(vs ...int) predicate.Employee {
	return predicate.Employee(sql.FieldNotIn(FieldAge, vs...))
}

// AgeGT applies the GT predicate on the "age" field.
func AgeGT(v int) predicate.Employee {
	return predicate.Employee(sql.FieldGT(FieldAge, v))
}

// AgeGTE applies the GTE predicate on the "age" field.
func AgeGTE(v int) predicate.Employee {
	return predicate.Employee(sql.FieldGTE(FieldAge, v))
}

// AgeLT applies the LT predicate on the "age" field.
func AgeLT(v int) predicate.Employee {
	return predicate.Employee(sql.FieldLT(FieldAge, v))
}

// AgeLTE applies the LTE predicate on the "age" field.
func AgeLTE(v int) predicate.Employee {
	return predicate.Employee(sql.FieldLTE(FieldAge, v))
}

// HasDepartment applies the HasEdge predicate on the "department" edge.
func HasDepartment() predicate.Employee {
	return predicate.Employee(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, DepartmentTable, DepartmentColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasDepartmentWith applies the HasEdge predicate on the "department" edge with a given conditions (other predicates).
func HasDepartmentWith(preds ...predicate.Department) predicate.Employee {
	return predicate.Employee(func(s *sql.Selector) {
		step := newDepartmentStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Employee) predicate.Employee {
	return predicate.Employee(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Employee) predicate.Employee {
	return predicate.Employee(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Employee) predicate.Employee {
	return predicate.Employee(func(s *sql.Selector) {
		p(s.Not())
	})
}
