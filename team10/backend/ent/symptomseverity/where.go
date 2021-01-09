// Code generated by entc, DO NOT EDIT.

package symptomseverity

import (
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/team10/app/ent/predicate"
)

// ID filters vertices based on their identifier.
func ID(id int) predicate.Symptomseverity {
	return predicate.Symptomseverity(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Symptomseverity {
	return predicate.Symptomseverity(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Symptomseverity {
	return predicate.Symptomseverity(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Symptomseverity {
	return predicate.Symptomseverity(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Symptomseverity {
	return predicate.Symptomseverity(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Symptomseverity {
	return predicate.Symptomseverity(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Symptomseverity {
	return predicate.Symptomseverity(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Symptomseverity {
	return predicate.Symptomseverity(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Symptomseverity {
	return predicate.Symptomseverity(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Symptomseverity applies equality check predicate on the "symptomseverity" field. It's identical to SymptomseverityEQ.
func Symptomseverity(v string) predicate.Symptomseverity {
	return predicate.Symptomseverity(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSymptomseverity), v))
	})
}

// SymptomseverityEQ applies the EQ predicate on the "symptomseverity" field.
func SymptomseverityEQ(v string) predicate.Symptomseverity {
	return predicate.Symptomseverity(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSymptomseverity), v))
	})
}

// SymptomseverityNEQ applies the NEQ predicate on the "symptomseverity" field.
func SymptomseverityNEQ(v string) predicate.Symptomseverity {
	return predicate.Symptomseverity(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldSymptomseverity), v))
	})
}

// SymptomseverityIn applies the In predicate on the "symptomseverity" field.
func SymptomseverityIn(vs ...string) predicate.Symptomseverity {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Symptomseverity(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldSymptomseverity), v...))
	})
}

// SymptomseverityNotIn applies the NotIn predicate on the "symptomseverity" field.
func SymptomseverityNotIn(vs ...string) predicate.Symptomseverity {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Symptomseverity(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldSymptomseverity), v...))
	})
}

// SymptomseverityGT applies the GT predicate on the "symptomseverity" field.
func SymptomseverityGT(v string) predicate.Symptomseverity {
	return predicate.Symptomseverity(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldSymptomseverity), v))
	})
}

// SymptomseverityGTE applies the GTE predicate on the "symptomseverity" field.
func SymptomseverityGTE(v string) predicate.Symptomseverity {
	return predicate.Symptomseverity(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldSymptomseverity), v))
	})
}

// SymptomseverityLT applies the LT predicate on the "symptomseverity" field.
func SymptomseverityLT(v string) predicate.Symptomseverity {
	return predicate.Symptomseverity(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldSymptomseverity), v))
	})
}

// SymptomseverityLTE applies the LTE predicate on the "symptomseverity" field.
func SymptomseverityLTE(v string) predicate.Symptomseverity {
	return predicate.Symptomseverity(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldSymptomseverity), v))
	})
}

// SymptomseverityContains applies the Contains predicate on the "symptomseverity" field.
func SymptomseverityContains(v string) predicate.Symptomseverity {
	return predicate.Symptomseverity(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldSymptomseverity), v))
	})
}

// SymptomseverityHasPrefix applies the HasPrefix predicate on the "symptomseverity" field.
func SymptomseverityHasPrefix(v string) predicate.Symptomseverity {
	return predicate.Symptomseverity(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldSymptomseverity), v))
	})
}

// SymptomseverityHasSuffix applies the HasSuffix predicate on the "symptomseverity" field.
func SymptomseverityHasSuffix(v string) predicate.Symptomseverity {
	return predicate.Symptomseverity(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldSymptomseverity), v))
	})
}

// SymptomseverityEqualFold applies the EqualFold predicate on the "symptomseverity" field.
func SymptomseverityEqualFold(v string) predicate.Symptomseverity {
	return predicate.Symptomseverity(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldSymptomseverity), v))
	})
}

// SymptomseverityContainsFold applies the ContainsFold predicate on the "symptomseverity" field.
func SymptomseverityContainsFold(v string) predicate.Symptomseverity {
	return predicate.Symptomseverity(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldSymptomseverity), v))
	})
}

// HasHistorytaking applies the HasEdge predicate on the "historytaking" edge.
func HasHistorytaking() predicate.Symptomseverity {
	return predicate.Symptomseverity(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(HistorytakingTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, HistorytakingTable, HistorytakingColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasHistorytakingWith applies the HasEdge predicate on the "historytaking" edge with a given conditions (other predicates).
func HasHistorytakingWith(preds ...predicate.Historytaking) predicate.Symptomseverity {
	return predicate.Symptomseverity(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(HistorytakingInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, HistorytakingTable, HistorytakingColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups list of predicates with the AND operator between them.
func And(predicates ...predicate.Symptomseverity) predicate.Symptomseverity {
	return predicate.Symptomseverity(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups list of predicates with the OR operator between them.
func Or(predicates ...predicate.Symptomseverity) predicate.Symptomseverity {
	return predicate.Symptomseverity(func(s *sql.Selector) {
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
func Not(p predicate.Symptomseverity) predicate.Symptomseverity {
	return predicate.Symptomseverity(func(s *sql.Selector) {
		p(s.Not())
	})
}