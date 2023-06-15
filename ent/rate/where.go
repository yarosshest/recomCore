// Code generated by ent, DO NOT EDIT.

package rate

import (
	"recomCore/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Rate {
	return predicate.Rate(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Rate {
	return predicate.Rate(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Rate {
	return predicate.Rate(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Rate {
	return predicate.Rate(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Rate {
	return predicate.Rate(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Rate {
	return predicate.Rate(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Rate {
	return predicate.Rate(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Rate {
	return predicate.Rate(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Rate {
	return predicate.Rate(sql.FieldLTE(FieldID, id))
}

// Rate applies equality check predicate on the "rate" field. It's identical to RateEQ.
func Rate(v bool) predicate.Rate {
	return predicate.Rate(sql.FieldEQ(FieldRate, v))
}

// RateEQ applies the EQ predicate on the "rate" field.
func RateEQ(v bool) predicate.Rate {
	return predicate.Rate(sql.FieldEQ(FieldRate, v))
}

// RateNEQ applies the NEQ predicate on the "rate" field.
func RateNEQ(v bool) predicate.Rate {
	return predicate.Rate(sql.FieldNEQ(FieldRate, v))
}

// HasOwner applies the HasEdge predicate on the "owner" edge.
func HasOwner() predicate.Rate {
	return predicate.Rate(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, OwnerTable, OwnerColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasOwnerWith applies the HasEdge predicate on the "owner" edge with a given conditions (other predicates).
func HasOwnerWith(preds ...predicate.User) predicate.Rate {
	return predicate.Rate(func(s *sql.Selector) {
		step := newOwnerStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasSubject applies the HasEdge predicate on the "subject" edge.
func HasSubject() predicate.Rate {
	return predicate.Rate(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, SubjectTable, SubjectColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasSubjectWith applies the HasEdge predicate on the "subject" edge with a given conditions (other predicates).
func HasSubjectWith(preds ...predicate.Product) predicate.Rate {
	return predicate.Rate(func(s *sql.Selector) {
		step := newSubjectStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Rate) predicate.Rate {
	return predicate.Rate(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Rate) predicate.Rate {
	return predicate.Rate(func(s *sql.Selector) {
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
func Not(p predicate.Rate) predicate.Rate {
	return predicate.Rate(func(s *sql.Selector) {
		p(s.Not())
	})
}