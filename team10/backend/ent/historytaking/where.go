// Code generated by entc, DO NOT EDIT.

package historytaking

import (
	"time"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/team10/app/ent/predicate"
)

// ID filters vertices based on their identifier.
func ID(id int) predicate.Historytaking {
	return predicate.Historytaking(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Historytaking {
	return predicate.Historytaking(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Historytaking {
	return predicate.Historytaking(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Historytaking {
	return predicate.Historytaking(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.Historytaking {
	return predicate.Historytaking(func(s *sql.Selector) {
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
func IDGT(id int) predicate.Historytaking {
	return predicate.Historytaking(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Historytaking {
	return predicate.Historytaking(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Historytaking {
	return predicate.Historytaking(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Historytaking {
	return predicate.Historytaking(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Hight applies equality check predicate on the "hight" field. It's identical to HightEQ.
func Hight(v float32) predicate.Historytaking {
	return predicate.Historytaking(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldHight), v))
	})
}

// Weight applies equality check predicate on the "weight" field. It's identical to WeightEQ.
func Weight(v float32) predicate.Historytaking {
	return predicate.Historytaking(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldWeight), v))
	})
}

// Temp applies equality check predicate on the "temp" field. It's identical to TempEQ.
func Temp(v float32) predicate.Historytaking {
	return predicate.Historytaking(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTemp), v))
	})
}

// Pulse applies equality check predicate on the "pulse" field. It's identical to PulseEQ.
func Pulse(v int) predicate.Historytaking {
	return predicate.Historytaking(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPulse), v))
	})
}

// Respiration applies equality check predicate on the "respiration" field. It's identical to RespirationEQ.
func Respiration(v int) predicate.Historytaking {
	return predicate.Historytaking(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRespiration), v))
	})
}

// Bp applies equality check predicate on the "bp" field. It's identical to BpEQ.
func Bp(v int) predicate.Historytaking {
	return predicate.Historytaking(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldBp), v))
	})
}

// Oxygen applies equality check predicate on the "oxygen" field. It's identical to OxygenEQ.
func Oxygen(v string) predicate.Historytaking {
	return predicate.Historytaking(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOxygen), v))
	})
}

// Symptom applies equality check predicate on the "symptom" field. It's identical to SymptomEQ.
func Symptom(v string) predicate.Historytaking {
	return predicate.Historytaking(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSymptom), v))
	})
}

// Datetime applies equality check predicate on the "datetime" field. It's identical to DatetimeEQ.
func Datetime(v time.Time) predicate.Historytaking {
	return predicate.Historytaking(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDatetime), v))
	})
}

// HightEQ applies the EQ predicate on the "hight" field.
func HightEQ(v float32) predicate.Historytaking {
	return predicate.Historytaking(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldHight), v))
	})
}

// HightNEQ applies the NEQ predicate on the "hight" field.
func HightNEQ(v float32) predicate.Historytaking {
	return predicate.Historytaking(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldHight), v))
	})
}

// HightIn applies the In predicate on the "hight" field.
func HightIn(vs ...float32) predicate.Historytaking {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Historytaking(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldHight), v...))
	})
}

// HightNotIn applies the NotIn predicate on the "hight" field.
func HightNotIn(vs ...float32) predicate.Historytaking {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Historytaking(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldHight), v...))
	})
}

// HightGT applies the GT predicate on the "hight" field.
func HightGT(v float32) predicate.Historytaking {
	return predicate.Historytaking(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldHight), v))
	})
}

// HightGTE applies the GTE predicate on the "hight" field.
func HightGTE(v float32) predicate.Historytaking {
	return predicate.Historytaking(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldHight), v))
	})
}

// HightLT applies the LT predicate on the "hight" field.
func HightLT(v float32) predicate.Historytaking {
	return predicate.Historytaking(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldHight), v))
	})
}

// HightLTE applies the LTE predicate on the "hight" field.
func HightLTE(v float32) predicate.Historytaking {
	return predicate.Historytaking(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldHight), v))
	})
}

// WeightEQ applies the EQ predicate on the "weight" field.
func WeightEQ(v float32) predicate.Historytaking {
	return predicate.Historytaking(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldWeight), v))
	})
}

// WeightNEQ applies the NEQ predicate on the "weight" field.
func WeightNEQ(v float32) predicate.Historytaking {
	return predicate.Historytaking(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldWeight), v))
	})
}

// WeightIn applies the In predicate on the "weight" field.
func WeightIn(vs ...float32) predicate.Historytaking {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Historytaking(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldWeight), v...))
	})
}

// WeightNotIn applies the NotIn predicate on the "weight" field.
func WeightNotIn(vs ...float32) predicate.Historytaking {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Historytaking(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldWeight), v...))
	})
}

// WeightGT applies the GT predicate on the "weight" field.
func WeightGT(v float32) predicate.Historytaking {
	return predicate.Historytaking(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldWeight), v))
	})
}

// WeightGTE applies the GTE predicate on the "weight" field.
func WeightGTE(v float32) predicate.Historytaking {
	return predicate.Historytaking(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldWeight), v))
	})
}

// WeightLT applies the LT predicate on the "weight" field.
func WeightLT(v float32) predicate.Historytaking {
	return predicate.Historytaking(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldWeight), v))
	})
}

// WeightLTE applies the LTE predicate on the "weight" field.
func WeightLTE(v float32) predicate.Historytaking {
	return predicate.Historytaking(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldWeight), v))
	})
}

// TempEQ applies the EQ predicate on the "temp" field.
func TempEQ(v float32) predicate.Historytaking {
	return predicate.Historytaking(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTemp), v))
	})
}

// TempNEQ applies the NEQ predicate on the "temp" field.
func TempNEQ(v float32) predicate.Historytaking {
	return predicate.Historytaking(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldTemp), v))
	})
}

// TempIn applies the In predicate on the "temp" field.
func TempIn(vs ...float32) predicate.Historytaking {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Historytaking(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldTemp), v...))
	})
}

// TempNotIn applies the NotIn predicate on the "temp" field.
func TempNotIn(vs ...float32) predicate.Historytaking {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Historytaking(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldTemp), v...))
	})
}

// TempGT applies the GT predicate on the "temp" field.
func TempGT(v float32) predicate.Historytaking {
	return predicate.Historytaking(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldTemp), v))
	})
}

// TempGTE applies the GTE predicate on the "temp" field.
func TempGTE(v float32) predicate.Historytaking {
	return predicate.Historytaking(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldTemp), v))
	})
}

// TempLT applies the LT predicate on the "temp" field.
func TempLT(v float32) predicate.Historytaking {
	return predicate.Historytaking(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldTemp), v))
	})
}

// TempLTE applies the LTE predicate on the "temp" field.
func TempLTE(v float32) predicate.Historytaking {
	return predicate.Historytaking(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldTemp), v))
	})
}

// PulseEQ applies the EQ predicate on the "pulse" field.
func PulseEQ(v int) predicate.Historytaking {
	return predicate.Historytaking(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPulse), v))
	})
}

// PulseNEQ applies the NEQ predicate on the "pulse" field.
func PulseNEQ(v int) predicate.Historytaking {
	return predicate.Historytaking(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldPulse), v))
	})
}

// PulseIn applies the In predicate on the "pulse" field.
func PulseIn(vs ...int) predicate.Historytaking {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Historytaking(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldPulse), v...))
	})
}

// PulseNotIn applies the NotIn predicate on the "pulse" field.
func PulseNotIn(vs ...int) predicate.Historytaking {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Historytaking(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldPulse), v...))
	})
}

// PulseGT applies the GT predicate on the "pulse" field.
func PulseGT(v int) predicate.Historytaking {
	return predicate.Historytaking(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldPulse), v))
	})
}

// PulseGTE applies the GTE predicate on the "pulse" field.
func PulseGTE(v int) predicate.Historytaking {
	return predicate.Historytaking(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldPulse), v))
	})
}

// PulseLT applies the LT predicate on the "pulse" field.
func PulseLT(v int) predicate.Historytaking {
	return predicate.Historytaking(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldPulse), v))
	})
}

// PulseLTE applies the LTE predicate on the "pulse" field.
func PulseLTE(v int) predicate.Historytaking {
	return predicate.Historytaking(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldPulse), v))
	})
}

// RespirationEQ applies the EQ predicate on the "respiration" field.
func RespirationEQ(v int) predicate.Historytaking {
	return predicate.Historytaking(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRespiration), v))
	})
}

// RespirationNEQ applies the NEQ predicate on the "respiration" field.
func RespirationNEQ(v int) predicate.Historytaking {
	return predicate.Historytaking(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldRespiration), v))
	})
}

// RespirationIn applies the In predicate on the "respiration" field.
func RespirationIn(vs ...int) predicate.Historytaking {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Historytaking(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldRespiration), v...))
	})
}

// RespirationNotIn applies the NotIn predicate on the "respiration" field.
func RespirationNotIn(vs ...int) predicate.Historytaking {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Historytaking(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldRespiration), v...))
	})
}

// RespirationGT applies the GT predicate on the "respiration" field.
func RespirationGT(v int) predicate.Historytaking {
	return predicate.Historytaking(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldRespiration), v))
	})
}

// RespirationGTE applies the GTE predicate on the "respiration" field.
func RespirationGTE(v int) predicate.Historytaking {
	return predicate.Historytaking(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldRespiration), v))
	})
}

// RespirationLT applies the LT predicate on the "respiration" field.
func RespirationLT(v int) predicate.Historytaking {
	return predicate.Historytaking(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldRespiration), v))
	})
}

// RespirationLTE applies the LTE predicate on the "respiration" field.
func RespirationLTE(v int) predicate.Historytaking {
	return predicate.Historytaking(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldRespiration), v))
	})
}

// BpEQ applies the EQ predicate on the "bp" field.
func BpEQ(v int) predicate.Historytaking {
	return predicate.Historytaking(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldBp), v))
	})
}

// BpNEQ applies the NEQ predicate on the "bp" field.
func BpNEQ(v int) predicate.Historytaking {
	return predicate.Historytaking(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldBp), v))
	})
}

// BpIn applies the In predicate on the "bp" field.
func BpIn(vs ...int) predicate.Historytaking {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Historytaking(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldBp), v...))
	})
}

// BpNotIn applies the NotIn predicate on the "bp" field.
func BpNotIn(vs ...int) predicate.Historytaking {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Historytaking(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldBp), v...))
	})
}

// BpGT applies the GT predicate on the "bp" field.
func BpGT(v int) predicate.Historytaking {
	return predicate.Historytaking(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldBp), v))
	})
}

// BpGTE applies the GTE predicate on the "bp" field.
func BpGTE(v int) predicate.Historytaking {
	return predicate.Historytaking(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldBp), v))
	})
}

// BpLT applies the LT predicate on the "bp" field.
func BpLT(v int) predicate.Historytaking {
	return predicate.Historytaking(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldBp), v))
	})
}

// BpLTE applies the LTE predicate on the "bp" field.
func BpLTE(v int) predicate.Historytaking {
	return predicate.Historytaking(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldBp), v))
	})
}

// OxygenEQ applies the EQ predicate on the "oxygen" field.
func OxygenEQ(v string) predicate.Historytaking {
	return predicate.Historytaking(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOxygen), v))
	})
}

// OxygenNEQ applies the NEQ predicate on the "oxygen" field.
func OxygenNEQ(v string) predicate.Historytaking {
	return predicate.Historytaking(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldOxygen), v))
	})
}

// OxygenIn applies the In predicate on the "oxygen" field.
func OxygenIn(vs ...string) predicate.Historytaking {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Historytaking(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldOxygen), v...))
	})
}

// OxygenNotIn applies the NotIn predicate on the "oxygen" field.
func OxygenNotIn(vs ...string) predicate.Historytaking {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Historytaking(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldOxygen), v...))
	})
}

// OxygenGT applies the GT predicate on the "oxygen" field.
func OxygenGT(v string) predicate.Historytaking {
	return predicate.Historytaking(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldOxygen), v))
	})
}

// OxygenGTE applies the GTE predicate on the "oxygen" field.
func OxygenGTE(v string) predicate.Historytaking {
	return predicate.Historytaking(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldOxygen), v))
	})
}

// OxygenLT applies the LT predicate on the "oxygen" field.
func OxygenLT(v string) predicate.Historytaking {
	return predicate.Historytaking(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldOxygen), v))
	})
}

// OxygenLTE applies the LTE predicate on the "oxygen" field.
func OxygenLTE(v string) predicate.Historytaking {
	return predicate.Historytaking(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldOxygen), v))
	})
}

// OxygenContains applies the Contains predicate on the "oxygen" field.
func OxygenContains(v string) predicate.Historytaking {
	return predicate.Historytaking(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldOxygen), v))
	})
}

// OxygenHasPrefix applies the HasPrefix predicate on the "oxygen" field.
func OxygenHasPrefix(v string) predicate.Historytaking {
	return predicate.Historytaking(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldOxygen), v))
	})
}

// OxygenHasSuffix applies the HasSuffix predicate on the "oxygen" field.
func OxygenHasSuffix(v string) predicate.Historytaking {
	return predicate.Historytaking(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldOxygen), v))
	})
}

// OxygenEqualFold applies the EqualFold predicate on the "oxygen" field.
func OxygenEqualFold(v string) predicate.Historytaking {
	return predicate.Historytaking(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldOxygen), v))
	})
}

// OxygenContainsFold applies the ContainsFold predicate on the "oxygen" field.
func OxygenContainsFold(v string) predicate.Historytaking {
	return predicate.Historytaking(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldOxygen), v))
	})
}

// SymptomEQ applies the EQ predicate on the "symptom" field.
func SymptomEQ(v string) predicate.Historytaking {
	return predicate.Historytaking(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSymptom), v))
	})
}

// SymptomNEQ applies the NEQ predicate on the "symptom" field.
func SymptomNEQ(v string) predicate.Historytaking {
	return predicate.Historytaking(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldSymptom), v))
	})
}

// SymptomIn applies the In predicate on the "symptom" field.
func SymptomIn(vs ...string) predicate.Historytaking {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Historytaking(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldSymptom), v...))
	})
}

// SymptomNotIn applies the NotIn predicate on the "symptom" field.
func SymptomNotIn(vs ...string) predicate.Historytaking {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Historytaking(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldSymptom), v...))
	})
}

// SymptomGT applies the GT predicate on the "symptom" field.
func SymptomGT(v string) predicate.Historytaking {
	return predicate.Historytaking(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldSymptom), v))
	})
}

// SymptomGTE applies the GTE predicate on the "symptom" field.
func SymptomGTE(v string) predicate.Historytaking {
	return predicate.Historytaking(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldSymptom), v))
	})
}

// SymptomLT applies the LT predicate on the "symptom" field.
func SymptomLT(v string) predicate.Historytaking {
	return predicate.Historytaking(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldSymptom), v))
	})
}

// SymptomLTE applies the LTE predicate on the "symptom" field.
func SymptomLTE(v string) predicate.Historytaking {
	return predicate.Historytaking(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldSymptom), v))
	})
}

// SymptomContains applies the Contains predicate on the "symptom" field.
func SymptomContains(v string) predicate.Historytaking {
	return predicate.Historytaking(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldSymptom), v))
	})
}

// SymptomHasPrefix applies the HasPrefix predicate on the "symptom" field.
func SymptomHasPrefix(v string) predicate.Historytaking {
	return predicate.Historytaking(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldSymptom), v))
	})
}

// SymptomHasSuffix applies the HasSuffix predicate on the "symptom" field.
func SymptomHasSuffix(v string) predicate.Historytaking {
	return predicate.Historytaking(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldSymptom), v))
	})
}

// SymptomEqualFold applies the EqualFold predicate on the "symptom" field.
func SymptomEqualFold(v string) predicate.Historytaking {
	return predicate.Historytaking(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldSymptom), v))
	})
}

// SymptomContainsFold applies the ContainsFold predicate on the "symptom" field.
func SymptomContainsFold(v string) predicate.Historytaking {
	return predicate.Historytaking(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldSymptom), v))
	})
}

// DatetimeEQ applies the EQ predicate on the "datetime" field.
func DatetimeEQ(v time.Time) predicate.Historytaking {
	return predicate.Historytaking(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDatetime), v))
	})
}

// DatetimeNEQ applies the NEQ predicate on the "datetime" field.
func DatetimeNEQ(v time.Time) predicate.Historytaking {
	return predicate.Historytaking(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDatetime), v))
	})
}

// DatetimeIn applies the In predicate on the "datetime" field.
func DatetimeIn(vs ...time.Time) predicate.Historytaking {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Historytaking(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldDatetime), v...))
	})
}

// DatetimeNotIn applies the NotIn predicate on the "datetime" field.
func DatetimeNotIn(vs ...time.Time) predicate.Historytaking {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Historytaking(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldDatetime), v...))
	})
}

// DatetimeGT applies the GT predicate on the "datetime" field.
func DatetimeGT(v time.Time) predicate.Historytaking {
	return predicate.Historytaking(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDatetime), v))
	})
}

// DatetimeGTE applies the GTE predicate on the "datetime" field.
func DatetimeGTE(v time.Time) predicate.Historytaking {
	return predicate.Historytaking(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDatetime), v))
	})
}

// DatetimeLT applies the LT predicate on the "datetime" field.
func DatetimeLT(v time.Time) predicate.Historytaking {
	return predicate.Historytaking(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDatetime), v))
	})
}

// DatetimeLTE applies the LTE predicate on the "datetime" field.
func DatetimeLTE(v time.Time) predicate.Historytaking {
	return predicate.Historytaking(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDatetime), v))
	})
}

// HasNurse applies the HasEdge predicate on the "nurse" edge.
func HasNurse() predicate.Historytaking {
	return predicate.Historytaking(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(NurseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, NurseTable, NurseColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasNurseWith applies the HasEdge predicate on the "nurse" edge with a given conditions (other predicates).
func HasNurseWith(preds ...predicate.Nurse) predicate.Historytaking {
	return predicate.Historytaking(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(NurseInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, NurseTable, NurseColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasDepartment applies the HasEdge predicate on the "department" edge.
func HasDepartment() predicate.Historytaking {
	return predicate.Historytaking(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(DepartmentTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, DepartmentTable, DepartmentColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasDepartmentWith applies the HasEdge predicate on the "department" edge with a given conditions (other predicates).
func HasDepartmentWith(preds ...predicate.Department) predicate.Historytaking {
	return predicate.Historytaking(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(DepartmentInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, DepartmentTable, DepartmentColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasSymptomseverity applies the HasEdge predicate on the "symptomseverity" edge.
func HasSymptomseverity() predicate.Historytaking {
	return predicate.Historytaking(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(SymptomseverityTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, SymptomseverityTable, SymptomseverityColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasSymptomseverityWith applies the HasEdge predicate on the "symptomseverity" edge with a given conditions (other predicates).
func HasSymptomseverityWith(preds ...predicate.Symptomseverity) predicate.Historytaking {
	return predicate.Historytaking(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(SymptomseverityInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, SymptomseverityTable, SymptomseverityColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasPatientrecord applies the HasEdge predicate on the "patientrecord" edge.
func HasPatientrecord() predicate.Historytaking {
	return predicate.Historytaking(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(PatientrecordTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, PatientrecordTable, PatientrecordColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPatientrecordWith applies the HasEdge predicate on the "patientrecord" edge with a given conditions (other predicates).
func HasPatientrecordWith(preds ...predicate.Patientrecord) predicate.Historytaking {
	return predicate.Historytaking(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(PatientrecordInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, PatientrecordTable, PatientrecordColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups list of predicates with the AND operator between them.
func And(predicates ...predicate.Historytaking) predicate.Historytaking {
	return predicate.Historytaking(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups list of predicates with the OR operator between them.
func Or(predicates ...predicate.Historytaking) predicate.Historytaking {
	return predicate.Historytaking(func(s *sql.Selector) {
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
func Not(p predicate.Historytaking) predicate.Historytaking {
	return predicate.Historytaking(func(s *sql.Selector) {
		p(s.Not())
	})
}
