// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/team10/app/ent/department"
	"github.com/team10/app/ent/historytaking"
	"github.com/team10/app/ent/nurse"
	"github.com/team10/app/ent/patientrecord"
	"github.com/team10/app/ent/symptomseverity"
)

// HistorytakingCreate is the builder for creating a Historytaking entity.
type HistorytakingCreate struct {
	config
	mutation *HistorytakingMutation
	hooks    []Hook
}

// SetHight sets the hight field.
func (hc *HistorytakingCreate) SetHight(f float32) *HistorytakingCreate {
	hc.mutation.SetHight(f)
	return hc
}

// SetWeight sets the weight field.
func (hc *HistorytakingCreate) SetWeight(f float32) *HistorytakingCreate {
	hc.mutation.SetWeight(f)
	return hc
}

// SetTemp sets the temp field.
func (hc *HistorytakingCreate) SetTemp(f float32) *HistorytakingCreate {
	hc.mutation.SetTemp(f)
	return hc
}

// SetPulse sets the pulse field.
func (hc *HistorytakingCreate) SetPulse(i int) *HistorytakingCreate {
	hc.mutation.SetPulse(i)
	return hc
}

// SetRespiration sets the respiration field.
func (hc *HistorytakingCreate) SetRespiration(i int) *HistorytakingCreate {
	hc.mutation.SetRespiration(i)
	return hc
}

// SetBp sets the bp field.
func (hc *HistorytakingCreate) SetBp(i int) *HistorytakingCreate {
	hc.mutation.SetBp(i)
	return hc
}

// SetOxygen sets the oxygen field.
func (hc *HistorytakingCreate) SetOxygen(s string) *HistorytakingCreate {
	hc.mutation.SetOxygen(s)
	return hc
}

// SetSymptom sets the symptom field.
func (hc *HistorytakingCreate) SetSymptom(s string) *HistorytakingCreate {
	hc.mutation.SetSymptom(s)
	return hc
}

// SetDatetime sets the datetime field.
func (hc *HistorytakingCreate) SetDatetime(t time.Time) *HistorytakingCreate {
	hc.mutation.SetDatetime(t)
	return hc
}

// SetNurseID sets the nurse edge to Nurse by id.
func (hc *HistorytakingCreate) SetNurseID(id int) *HistorytakingCreate {
	hc.mutation.SetNurseID(id)
	return hc
}

// SetNillableNurseID sets the nurse edge to Nurse by id if the given value is not nil.
func (hc *HistorytakingCreate) SetNillableNurseID(id *int) *HistorytakingCreate {
	if id != nil {
		hc = hc.SetNurseID(*id)
	}
	return hc
}

// SetNurse sets the nurse edge to Nurse.
func (hc *HistorytakingCreate) SetNurse(n *Nurse) *HistorytakingCreate {
	return hc.SetNurseID(n.ID)
}

// SetDepartmentID sets the department edge to Department by id.
func (hc *HistorytakingCreate) SetDepartmentID(id int) *HistorytakingCreate {
	hc.mutation.SetDepartmentID(id)
	return hc
}

// SetNillableDepartmentID sets the department edge to Department by id if the given value is not nil.
func (hc *HistorytakingCreate) SetNillableDepartmentID(id *int) *HistorytakingCreate {
	if id != nil {
		hc = hc.SetDepartmentID(*id)
	}
	return hc
}

// SetDepartment sets the department edge to Department.
func (hc *HistorytakingCreate) SetDepartment(d *Department) *HistorytakingCreate {
	return hc.SetDepartmentID(d.ID)
}

// SetSymptomseverityID sets the symptomseverity edge to Symptomseverity by id.
func (hc *HistorytakingCreate) SetSymptomseverityID(id int) *HistorytakingCreate {
	hc.mutation.SetSymptomseverityID(id)
	return hc
}

// SetNillableSymptomseverityID sets the symptomseverity edge to Symptomseverity by id if the given value is not nil.
func (hc *HistorytakingCreate) SetNillableSymptomseverityID(id *int) *HistorytakingCreate {
	if id != nil {
		hc = hc.SetSymptomseverityID(*id)
	}
	return hc
}

// SetSymptomseverity sets the symptomseverity edge to Symptomseverity.
func (hc *HistorytakingCreate) SetSymptomseverity(s *Symptomseverity) *HistorytakingCreate {
	return hc.SetSymptomseverityID(s.ID)
}

// SetPatientrecordID sets the patientrecord edge to Patientrecord by id.
func (hc *HistorytakingCreate) SetPatientrecordID(id int) *HistorytakingCreate {
	hc.mutation.SetPatientrecordID(id)
	return hc
}

// SetNillablePatientrecordID sets the patientrecord edge to Patientrecord by id if the given value is not nil.
func (hc *HistorytakingCreate) SetNillablePatientrecordID(id *int) *HistorytakingCreate {
	if id != nil {
		hc = hc.SetPatientrecordID(*id)
	}
	return hc
}

// SetPatientrecord sets the patientrecord edge to Patientrecord.
func (hc *HistorytakingCreate) SetPatientrecord(p *Patientrecord) *HistorytakingCreate {
	return hc.SetPatientrecordID(p.ID)
}

// Mutation returns the HistorytakingMutation object of the builder.
func (hc *HistorytakingCreate) Mutation() *HistorytakingMutation {
	return hc.mutation
}

// Save creates the Historytaking in the database.
func (hc *HistorytakingCreate) Save(ctx context.Context) (*Historytaking, error) {
	if _, ok := hc.mutation.Hight(); !ok {
		return nil, &ValidationError{Name: "hight", err: errors.New("ent: missing required field \"hight\"")}
	}
	if _, ok := hc.mutation.Weight(); !ok {
		return nil, &ValidationError{Name: "weight", err: errors.New("ent: missing required field \"weight\"")}
	}
	if _, ok := hc.mutation.Temp(); !ok {
		return nil, &ValidationError{Name: "temp", err: errors.New("ent: missing required field \"temp\"")}
	}
	if _, ok := hc.mutation.Pulse(); !ok {
		return nil, &ValidationError{Name: "pulse", err: errors.New("ent: missing required field \"pulse\"")}
	}
	if _, ok := hc.mutation.Respiration(); !ok {
		return nil, &ValidationError{Name: "respiration", err: errors.New("ent: missing required field \"respiration\"")}
	}
	if _, ok := hc.mutation.Bp(); !ok {
		return nil, &ValidationError{Name: "bp", err: errors.New("ent: missing required field \"bp\"")}
	}
	if _, ok := hc.mutation.Oxygen(); !ok {
		return nil, &ValidationError{Name: "oxygen", err: errors.New("ent: missing required field \"oxygen\"")}
	}
	if _, ok := hc.mutation.Symptom(); !ok {
		return nil, &ValidationError{Name: "symptom", err: errors.New("ent: missing required field \"symptom\"")}
	}
	if _, ok := hc.mutation.Datetime(); !ok {
		return nil, &ValidationError{Name: "datetime", err: errors.New("ent: missing required field \"datetime\"")}
	}
	var (
		err  error
		node *Historytaking
	)
	if len(hc.hooks) == 0 {
		node, err = hc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*HistorytakingMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			hc.mutation = mutation
			node, err = hc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(hc.hooks) - 1; i >= 0; i-- {
			mut = hc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, hc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (hc *HistorytakingCreate) SaveX(ctx context.Context) *Historytaking {
	v, err := hc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (hc *HistorytakingCreate) sqlSave(ctx context.Context) (*Historytaking, error) {
	h, _spec := hc.createSpec()
	if err := sqlgraph.CreateNode(ctx, hc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	h.ID = int(id)
	return h, nil
}

func (hc *HistorytakingCreate) createSpec() (*Historytaking, *sqlgraph.CreateSpec) {
	var (
		h     = &Historytaking{config: hc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: historytaking.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: historytaking.FieldID,
			},
		}
	)
	if value, ok := hc.mutation.Hight(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat32,
			Value:  value,
			Column: historytaking.FieldHight,
		})
		h.Hight = value
	}
	if value, ok := hc.mutation.Weight(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat32,
			Value:  value,
			Column: historytaking.FieldWeight,
		})
		h.Weight = value
	}
	if value, ok := hc.mutation.Temp(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat32,
			Value:  value,
			Column: historytaking.FieldTemp,
		})
		h.Temp = value
	}
	if value, ok := hc.mutation.Pulse(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: historytaking.FieldPulse,
		})
		h.Pulse = value
	}
	if value, ok := hc.mutation.Respiration(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: historytaking.FieldRespiration,
		})
		h.Respiration = value
	}
	if value, ok := hc.mutation.Bp(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: historytaking.FieldBp,
		})
		h.Bp = value
	}
	if value, ok := hc.mutation.Oxygen(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: historytaking.FieldOxygen,
		})
		h.Oxygen = value
	}
	if value, ok := hc.mutation.Symptom(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: historytaking.FieldSymptom,
		})
		h.Symptom = value
	}
	if value, ok := hc.mutation.Datetime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: historytaking.FieldDatetime,
		})
		h.Datetime = value
	}
	if nodes := hc.mutation.NurseIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   historytaking.NurseTable,
			Columns: []string{historytaking.NurseColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: nurse.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := hc.mutation.DepartmentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   historytaking.DepartmentTable,
			Columns: []string{historytaking.DepartmentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: department.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := hc.mutation.SymptomseverityIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   historytaking.SymptomseverityTable,
			Columns: []string{historytaking.SymptomseverityColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: symptomseverity.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := hc.mutation.PatientrecordIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   historytaking.PatientrecordTable,
			Columns: []string{historytaking.PatientrecordColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: patientrecord.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return h, _spec
}