// Code generated by entc, DO NOT EDIT.

package nurse

import (
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/team10/app/ent/predicate"
)

// ID filters vertices based on their identifier.
func ID(id int) predicate.Nurse {
	return predicate.Nurse(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Nurse {
	return predicate.Nurse(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Nurse {
	return predicate.Nurse(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Nurse {
	return predicate.Nurse(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.Nurse {
	return predicate.Nurse(func(s *sql.Selector) {
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
func IDGT(id int) predicate.Nurse {
	return predicate.Nurse(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Nurse {
	return predicate.Nurse(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Nurse {
	return predicate.Nurse(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Nurse {
	return predicate.Nurse(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Nurse {
	return predicate.Nurse(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// Nursinglicense applies equality check predicate on the "nursinglicense" field. It's identical to NursinglicenseEQ.
func Nursinglicense(v string) predicate.Nurse {
	return predicate.Nurse(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldNursinglicense), v))
	})
}

// Position applies equality check predicate on the "position" field. It's identical to PositionEQ.
func Position(v string) predicate.Nurse {
	return predicate.Nurse(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPosition), v))
	})
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Nurse {
	return predicate.Nurse(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Nurse {
	return predicate.Nurse(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldName), v))
	})
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Nurse {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Nurse(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldName), v...))
	})
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Nurse {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Nurse(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldName), v...))
	})
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Nurse {
	return predicate.Nurse(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldName), v))
	})
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Nurse {
	return predicate.Nurse(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldName), v))
	})
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Nurse {
	return predicate.Nurse(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldName), v))
	})
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Nurse {
	return predicate.Nurse(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldName), v))
	})
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Nurse {
	return predicate.Nurse(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldName), v))
	})
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Nurse {
	return predicate.Nurse(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldName), v))
	})
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Nurse {
	return predicate.Nurse(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldName), v))
	})
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Nurse {
	return predicate.Nurse(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldName), v))
	})
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Nurse {
	return predicate.Nurse(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldName), v))
	})
}

// NursinglicenseEQ applies the EQ predicate on the "nursinglicense" field.
func NursinglicenseEQ(v string) predicate.Nurse {
	return predicate.Nurse(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldNursinglicense), v))
	})
}

// NursinglicenseNEQ applies the NEQ predicate on the "nursinglicense" field.
func NursinglicenseNEQ(v string) predicate.Nurse {
	return predicate.Nurse(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldNursinglicense), v))
	})
}

// NursinglicenseIn applies the In predicate on the "nursinglicense" field.
func NursinglicenseIn(vs ...string) predicate.Nurse {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Nurse(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldNursinglicense), v...))
	})
}

// NursinglicenseNotIn applies the NotIn predicate on the "nursinglicense" field.
func NursinglicenseNotIn(vs ...string) predicate.Nurse {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Nurse(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldNursinglicense), v...))
	})
}

// NursinglicenseGT applies the GT predicate on the "nursinglicense" field.
func NursinglicenseGT(v string) predicate.Nurse {
	return predicate.Nurse(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldNursinglicense), v))
	})
}

// NursinglicenseGTE applies the GTE predicate on the "nursinglicense" field.
func NursinglicenseGTE(v string) predicate.Nurse {
	return predicate.Nurse(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldNursinglicense), v))
	})
}

// NursinglicenseLT applies the LT predicate on the "nursinglicense" field.
func NursinglicenseLT(v string) predicate.Nurse {
	return predicate.Nurse(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldNursinglicense), v))
	})
}

// NursinglicenseLTE applies the LTE predicate on the "nursinglicense" field.
func NursinglicenseLTE(v string) predicate.Nurse {
	return predicate.Nurse(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldNursinglicense), v))
	})
}

// NursinglicenseContains applies the Contains predicate on the "nursinglicense" field.
func NursinglicenseContains(v string) predicate.Nurse {
	return predicate.Nurse(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldNursinglicense), v))
	})
}

// NursinglicenseHasPrefix applies the HasPrefix predicate on the "nursinglicense" field.
func NursinglicenseHasPrefix(v string) predicate.Nurse {
	return predicate.Nurse(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldNursinglicense), v))
	})
}

// NursinglicenseHasSuffix applies the HasSuffix predicate on the "nursinglicense" field.
func NursinglicenseHasSuffix(v string) predicate.Nurse {
	return predicate.Nurse(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldNursinglicense), v))
	})
}

// NursinglicenseEqualFold applies the EqualFold predicate on the "nursinglicense" field.
func NursinglicenseEqualFold(v string) predicate.Nurse {
	return predicate.Nurse(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldNursinglicense), v))
	})
}

// NursinglicenseContainsFold applies the ContainsFold predicate on the "nursinglicense" field.
func NursinglicenseContainsFold(v string) predicate.Nurse {
	return predicate.Nurse(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldNursinglicense), v))
	})
}

// PositionEQ applies the EQ predicate on the "position" field.
func PositionEQ(v string) predicate.Nurse {
	return predicate.Nurse(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPosition), v))
	})
}

// PositionNEQ applies the NEQ predicate on the "position" field.
func PositionNEQ(v string) predicate.Nurse {
	return predicate.Nurse(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldPosition), v))
	})
}

// PositionIn applies the In predicate on the "position" field.
func PositionIn(vs ...string) predicate.Nurse {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Nurse(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldPosition), v...))
	})
}

// PositionNotIn applies the NotIn predicate on the "position" field.
func PositionNotIn(vs ...string) predicate.Nurse {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Nurse(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldPosition), v...))
	})
}

// PositionGT applies the GT predicate on the "position" field.
func PositionGT(v string) predicate.Nurse {
	return predicate.Nurse(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldPosition), v))
	})
}

// PositionGTE applies the GTE predicate on the "position" field.
func PositionGTE(v string) predicate.Nurse {
	return predicate.Nurse(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldPosition), v))
	})
}

// PositionLT applies the LT predicate on the "position" field.
func PositionLT(v string) predicate.Nurse {
	return predicate.Nurse(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldPosition), v))
	})
}

// PositionLTE applies the LTE predicate on the "position" field.
func PositionLTE(v string) predicate.Nurse {
	return predicate.Nurse(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldPosition), v))
	})
}

// PositionContains applies the Contains predicate on the "position" field.
func PositionContains(v string) predicate.Nurse {
	return predicate.Nurse(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldPosition), v))
	})
}

// PositionHasPrefix applies the HasPrefix predicate on the "position" field.
func PositionHasPrefix(v string) predicate.Nurse {
	return predicate.Nurse(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldPosition), v))
	})
}

// PositionHasSuffix applies the HasSuffix predicate on the "position" field.
func PositionHasSuffix(v string) predicate.Nurse {
	return predicate.Nurse(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldPosition), v))
	})
}

// PositionEqualFold applies the EqualFold predicate on the "position" field.
func PositionEqualFold(v string) predicate.Nurse {
	return predicate.Nurse(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldPosition), v))
	})
}

// PositionContainsFold applies the ContainsFold predicate on the "position" field.
func PositionContainsFold(v string) predicate.Nurse {
	return predicate.Nurse(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldPosition), v))
	})
}

// HasHistorytaking applies the HasEdge predicate on the "historytaking" edge.
func HasHistorytaking() predicate.Nurse {
	return predicate.Nurse(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(HistorytakingTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, HistorytakingTable, HistorytakingColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasHistorytakingWith applies the HasEdge predicate on the "historytaking" edge with a given conditions (other predicates).
func HasHistorytakingWith(preds ...predicate.Historytaking) predicate.Nurse {
	return predicate.Nurse(func(s *sql.Selector) {
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

// HasUser applies the HasEdge predicate on the "user" edge.
func HasUser() predicate.Nurse {
	return predicate.Nurse(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(UserTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, UserTable, UserColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserWith applies the HasEdge predicate on the "user" edge with a given conditions (other predicates).
func HasUserWith(preds ...predicate.User) predicate.Nurse {
	return predicate.Nurse(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(UserInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, UserTable, UserColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups list of predicates with the AND operator between them.
func And(predicates ...predicate.Nurse) predicate.Nurse {
	return predicate.Nurse(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups list of predicates with the OR operator between them.
func Or(predicates ...predicate.Nurse) predicate.Nurse {
	return predicate.Nurse(func(s *sql.Selector) {
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
func Not(p predicate.Nurse) predicate.Nurse {
	return predicate.Nurse(func(s *sql.Selector) {
		p(s.Not())
	})
}
