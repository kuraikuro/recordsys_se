// Code generated by entc, DO NOT EDIT.

package hook

import (
	"context"
	"fmt"

	"github.com/team10/app/ent"
)

// The AbilitypatientrightsFunc type is an adapter to allow the use of ordinary
// function as Abilitypatientrights mutator.
type AbilitypatientrightsFunc func(context.Context, *ent.AbilitypatientrightsMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f AbilitypatientrightsFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.AbilitypatientrightsMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.AbilitypatientrightsMutation", m)
	}
	return f(ctx, mv)
}

// The BillFunc type is an adapter to allow the use of ordinary
// function as Bill mutator.
type BillFunc func(context.Context, *ent.BillMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f BillFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.BillMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.BillMutation", m)
	}
	return f(ctx, mv)
}

// The DepartmentFunc type is an adapter to allow the use of ordinary
// function as Department mutator.
type DepartmentFunc func(context.Context, *ent.DepartmentMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f DepartmentFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.DepartmentMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.DepartmentMutation", m)
	}
	return f(ctx, mv)
}

// The DoctorinfoFunc type is an adapter to allow the use of ordinary
// function as Doctorinfo mutator.
type DoctorinfoFunc func(context.Context, *ent.DoctorinfoMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f DoctorinfoFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.DoctorinfoMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.DoctorinfoMutation", m)
	}
	return f(ctx, mv)
}

// The EducationlevelFunc type is an adapter to allow the use of ordinary
// function as Educationlevel mutator.
type EducationlevelFunc func(context.Context, *ent.EducationlevelMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f EducationlevelFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.EducationlevelMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.EducationlevelMutation", m)
	}
	return f(ctx, mv)
}

// The FinancierFunc type is an adapter to allow the use of ordinary
// function as Financier mutator.
type FinancierFunc func(context.Context, *ent.FinancierMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f FinancierFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.FinancierMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.FinancierMutation", m)
	}
	return f(ctx, mv)
}

// The GenderFunc type is an adapter to allow the use of ordinary
// function as Gender mutator.
type GenderFunc func(context.Context, *ent.GenderMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f GenderFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.GenderMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.GenderMutation", m)
	}
	return f(ctx, mv)
}

// The HistorytakingFunc type is an adapter to allow the use of ordinary
// function as Historytaking mutator.
type HistorytakingFunc func(context.Context, *ent.HistorytakingMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f HistorytakingFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.HistorytakingMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.HistorytakingMutation", m)
	}
	return f(ctx, mv)
}

// The InsuranceFunc type is an adapter to allow the use of ordinary
// function as Insurance mutator.
type InsuranceFunc func(context.Context, *ent.InsuranceMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f InsuranceFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.InsuranceMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.InsuranceMutation", m)
	}
	return f(ctx, mv)
}

// The MedicalrecordstaffFunc type is an adapter to allow the use of ordinary
// function as Medicalrecordstaff mutator.
type MedicalrecordstaffFunc func(context.Context, *ent.MedicalrecordstaffMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f MedicalrecordstaffFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.MedicalrecordstaffMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.MedicalrecordstaffMutation", m)
	}
	return f(ctx, mv)
}

// The NurseFunc type is an adapter to allow the use of ordinary
// function as Nurse mutator.
type NurseFunc func(context.Context, *ent.NurseMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f NurseFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.NurseMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.NurseMutation", m)
	}
	return f(ctx, mv)
}

// The OfficeroomFunc type is an adapter to allow the use of ordinary
// function as Officeroom mutator.
type OfficeroomFunc func(context.Context, *ent.OfficeroomMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f OfficeroomFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.OfficeroomMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.OfficeroomMutation", m)
	}
	return f(ctx, mv)
}

// The PatientrecordFunc type is an adapter to allow the use of ordinary
// function as Patientrecord mutator.
type PatientrecordFunc func(context.Context, *ent.PatientrecordMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f PatientrecordFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.PatientrecordMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.PatientrecordMutation", m)
	}
	return f(ctx, mv)
}

// The PatientrightsFunc type is an adapter to allow the use of ordinary
// function as Patientrights mutator.
type PatientrightsFunc func(context.Context, *ent.PatientrightsMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f PatientrightsFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.PatientrightsMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.PatientrightsMutation", m)
	}
	return f(ctx, mv)
}

// The PatientrightstypeFunc type is an adapter to allow the use of ordinary
// function as Patientrightstype mutator.
type PatientrightstypeFunc func(context.Context, *ent.PatientrightstypeMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f PatientrightstypeFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.PatientrightstypeMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.PatientrightstypeMutation", m)
	}
	return f(ctx, mv)
}

// The PaytypeFunc type is an adapter to allow the use of ordinary
// function as Paytype mutator.
type PaytypeFunc func(context.Context, *ent.PaytypeMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f PaytypeFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.PaytypeMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.PaytypeMutation", m)
	}
	return f(ctx, mv)
}

// The PrenameFunc type is an adapter to allow the use of ordinary
// function as Prename mutator.
type PrenameFunc func(context.Context, *ent.PrenameMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f PrenameFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.PrenameMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.PrenameMutation", m)
	}
	return f(ctx, mv)
}

// The RegistrarFunc type is an adapter to allow the use of ordinary
// function as Registrar mutator.
type RegistrarFunc func(context.Context, *ent.RegistrarMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f RegistrarFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.RegistrarMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.RegistrarMutation", m)
	}
	return f(ctx, mv)
}

// The SymptomseverityFunc type is an adapter to allow the use of ordinary
// function as Symptomseverity mutator.
type SymptomseverityFunc func(context.Context, *ent.SymptomseverityMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f SymptomseverityFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.SymptomseverityMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.SymptomseverityMutation", m)
	}
	return f(ctx, mv)
}

// The TreatmentFunc type is an adapter to allow the use of ordinary
// function as Treatment mutator.
type TreatmentFunc func(context.Context, *ent.TreatmentMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f TreatmentFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.TreatmentMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.TreatmentMutation", m)
	}
	return f(ctx, mv)
}

// The TypetreatmentFunc type is an adapter to allow the use of ordinary
// function as Typetreatment mutator.
type TypetreatmentFunc func(context.Context, *ent.TypetreatmentMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f TypetreatmentFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.TypetreatmentMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.TypetreatmentMutation", m)
	}
	return f(ctx, mv)
}

// The UnpaybillFunc type is an adapter to allow the use of ordinary
// function as Unpaybill mutator.
type UnpaybillFunc func(context.Context, *ent.UnpaybillMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f UnpaybillFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.UnpaybillMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.UnpaybillMutation", m)
	}
	return f(ctx, mv)
}

// The UserFunc type is an adapter to allow the use of ordinary
// function as User mutator.
type UserFunc func(context.Context, *ent.UserMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f UserFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.UserMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.UserMutation", m)
	}
	return f(ctx, mv)
}

// The UserstatusFunc type is an adapter to allow the use of ordinary
// function as Userstatus mutator.
type UserstatusFunc func(context.Context, *ent.UserstatusMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f UserstatusFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.UserstatusMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.UserstatusMutation", m)
	}
	return f(ctx, mv)
}

// Condition is a hook condition function.
type Condition func(context.Context, ent.Mutation) bool

// And groups conditions with the AND operator.
func And(first, second Condition, rest ...Condition) Condition {
	return func(ctx context.Context, m ent.Mutation) bool {
		if !first(ctx, m) || !second(ctx, m) {
			return false
		}
		for _, cond := range rest {
			if !cond(ctx, m) {
				return false
			}
		}
		return true
	}
}

// Or groups conditions with the OR operator.
func Or(first, second Condition, rest ...Condition) Condition {
	return func(ctx context.Context, m ent.Mutation) bool {
		if first(ctx, m) || second(ctx, m) {
			return true
		}
		for _, cond := range rest {
			if cond(ctx, m) {
				return true
			}
		}
		return false
	}
}

// Not negates a given condition.
func Not(cond Condition) Condition {
	return func(ctx context.Context, m ent.Mutation) bool {
		return !cond(ctx, m)
	}
}

// HasOp is a condition testing mutation operation.
func HasOp(op ent.Op) Condition {
	return func(_ context.Context, m ent.Mutation) bool {
		return m.Op().Is(op)
	}
}

// HasAddedFields is a condition validating `.AddedField` on fields.
func HasAddedFields(field string, fields ...string) Condition {
	return func(_ context.Context, m ent.Mutation) bool {
		if _, exists := m.AddedField(field); !exists {
			return false
		}
		for _, field := range fields {
			if _, exists := m.AddedField(field); !exists {
				return false
			}
		}
		return true
	}
}

// HasClearedFields is a condition validating `.FieldCleared` on fields.
func HasClearedFields(field string, fields ...string) Condition {
	return func(_ context.Context, m ent.Mutation) bool {
		if exists := m.FieldCleared(field); !exists {
			return false
		}
		for _, field := range fields {
			if exists := m.FieldCleared(field); !exists {
				return false
			}
		}
		return true
	}
}

// HasFields is a condition validating `.Field` on fields.
func HasFields(field string, fields ...string) Condition {
	return func(_ context.Context, m ent.Mutation) bool {
		if _, exists := m.Field(field); !exists {
			return false
		}
		for _, field := range fields {
			if _, exists := m.Field(field); !exists {
				return false
			}
		}
		return true
	}
}

// If executes the given hook under condition.
//
//	Hook.If(ComputeAverage, And(HasFields(...), HasAddedFields(...)))
//
func If(hk ent.Hook, cond Condition) ent.Hook {
	return func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			if cond(ctx, m) {
				return hk(next).Mutate(ctx, m)
			}
			return next.Mutate(ctx, m)
		})
	}
}

// On executes the given hook only for the given operation.
//
//	hook.On(Log, ent.Delete|ent.Create)
//
func On(hk ent.Hook, op ent.Op) ent.Hook {
	return If(hk, HasOp(op))
}

// Unless skips the given hook only for the given operation.
//
//	hook.Unless(Log, ent.Update|ent.UpdateOne)
//
func Unless(hk ent.Hook, op ent.Op) ent.Hook {
	return If(hk, Not(HasOp(op)))
}

// Reject returns a hook that rejects all operations that match op.
//
//	func (T) Hooks() []ent.Hook {
//		return []ent.Hook{
//			Reject(ent.Delete|ent.Update),
//		}
//	}
//
func Reject(op ent.Op) ent.Hook {
	hk := func(ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(_ context.Context, m ent.Mutation) (ent.Value, error) {
			return nil, fmt.Errorf("%s operation is not allowed", m.Op())
		})
	}
	return On(hk, op)
}

// Chain acts as a list of hooks and is effectively immutable.
// Once created, it will always hold the same set of hooks in the same order.
type Chain struct {
	hooks []ent.Hook
}

// NewChain creates a new chain of hooks.
func NewChain(hooks ...ent.Hook) Chain {
	return Chain{append([]ent.Hook(nil), hooks...)}
}

// Hook chains the list of hooks and returns the final hook.
func (c Chain) Hook() ent.Hook {
	return func(mutator ent.Mutator) ent.Mutator {
		for i := len(c.hooks) - 1; i >= 0; i-- {
			mutator = c.hooks[i](mutator)
		}
		return mutator
	}
}

// Append extends a chain, adding the specified hook
// as the last ones in the mutation flow.
func (c Chain) Append(hooks ...ent.Hook) Chain {
	newHooks := make([]ent.Hook, 0, len(c.hooks)+len(hooks))
	newHooks = append(newHooks, c.hooks...)
	newHooks = append(newHooks, hooks...)
	return Chain{newHooks}
}

// Extend extends a chain, adding the specified chain
// as the last ones in the mutation flow.
func (c Chain) Extend(chain Chain) Chain {
	return c.Append(chain.hooks...)
}
