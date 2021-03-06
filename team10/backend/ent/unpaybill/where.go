// Code generated by entc, DO NOT EDIT.

package unpaybill

import (
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/team10/app/ent/predicate"
)

// ID filters vertices based on their identifier.
func ID(id int) predicate.Unpaybill {
	return predicate.Unpaybill(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Unpaybill {
	return predicate.Unpaybill(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Unpaybill {
	return predicate.Unpaybill(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Unpaybill {
	return predicate.Unpaybill(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.Unpaybill {
	return predicate.Unpaybill(func(s *sql.Selector) {
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
func IDGT(id int) predicate.Unpaybill {
	return predicate.Unpaybill(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Unpaybill {
	return predicate.Unpaybill(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Unpaybill {
	return predicate.Unpaybill(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Unpaybill {
	return predicate.Unpaybill(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Status applies equality check predicate on the "Status" field. It's identical to StatusEQ.
func Status(v string) predicate.Unpaybill {
	return predicate.Unpaybill(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStatus), v))
	})
}

// StatusEQ applies the EQ predicate on the "Status" field.
func StatusEQ(v string) predicate.Unpaybill {
	return predicate.Unpaybill(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStatus), v))
	})
}

// StatusNEQ applies the NEQ predicate on the "Status" field.
func StatusNEQ(v string) predicate.Unpaybill {
	return predicate.Unpaybill(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldStatus), v))
	})
}

// StatusIn applies the In predicate on the "Status" field.
func StatusIn(vs ...string) predicate.Unpaybill {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Unpaybill(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldStatus), v...))
	})
}

// StatusNotIn applies the NotIn predicate on the "Status" field.
func StatusNotIn(vs ...string) predicate.Unpaybill {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Unpaybill(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldStatus), v...))
	})
}

// StatusGT applies the GT predicate on the "Status" field.
func StatusGT(v string) predicate.Unpaybill {
	return predicate.Unpaybill(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldStatus), v))
	})
}

// StatusGTE applies the GTE predicate on the "Status" field.
func StatusGTE(v string) predicate.Unpaybill {
	return predicate.Unpaybill(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldStatus), v))
	})
}

// StatusLT applies the LT predicate on the "Status" field.
func StatusLT(v string) predicate.Unpaybill {
	return predicate.Unpaybill(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldStatus), v))
	})
}

// StatusLTE applies the LTE predicate on the "Status" field.
func StatusLTE(v string) predicate.Unpaybill {
	return predicate.Unpaybill(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldStatus), v))
	})
}

// StatusContains applies the Contains predicate on the "Status" field.
func StatusContains(v string) predicate.Unpaybill {
	return predicate.Unpaybill(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldStatus), v))
	})
}

// StatusHasPrefix applies the HasPrefix predicate on the "Status" field.
func StatusHasPrefix(v string) predicate.Unpaybill {
	return predicate.Unpaybill(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldStatus), v))
	})
}

// StatusHasSuffix applies the HasSuffix predicate on the "Status" field.
func StatusHasSuffix(v string) predicate.Unpaybill {
	return predicate.Unpaybill(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldStatus), v))
	})
}

// StatusEqualFold applies the EqualFold predicate on the "Status" field.
func StatusEqualFold(v string) predicate.Unpaybill {
	return predicate.Unpaybill(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldStatus), v))
	})
}

// StatusContainsFold applies the ContainsFold predicate on the "Status" field.
func StatusContainsFold(v string) predicate.Unpaybill {
	return predicate.Unpaybill(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldStatus), v))
	})
}

// HasTreatment applies the HasEdge predicate on the "treatment" edge.
func HasTreatment() predicate.Unpaybill {
	return predicate.Unpaybill(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(TreatmentTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, TreatmentTable, TreatmentColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasTreatmentWith applies the HasEdge predicate on the "treatment" edge with a given conditions (other predicates).
func HasTreatmentWith(preds ...predicate.Treatment) predicate.Unpaybill {
	return predicate.Unpaybill(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(TreatmentInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, TreatmentTable, TreatmentColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasBills applies the HasEdge predicate on the "bills" edge.
func HasBills() predicate.Unpaybill {
	return predicate.Unpaybill(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(BillsTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, BillsTable, BillsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasBillsWith applies the HasEdge predicate on the "bills" edge with a given conditions (other predicates).
func HasBillsWith(preds ...predicate.Bill) predicate.Unpaybill {
	return predicate.Unpaybill(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(BillsInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, BillsTable, BillsColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups list of predicates with the AND operator between them.
func And(predicates ...predicate.Unpaybill) predicate.Unpaybill {
	return predicate.Unpaybill(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups list of predicates with the OR operator between them.
func Or(predicates ...predicate.Unpaybill) predicate.Unpaybill {
	return predicate.Unpaybill(func(s *sql.Selector) {
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
func Not(p predicate.Unpaybill) predicate.Unpaybill {
	return predicate.Unpaybill(func(s *sql.Selector) {
		p(s.Not())
	})
}
