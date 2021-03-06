// Code generated by entc, DO NOT EDIT.

package ent

import (
	"github.com/team10/app/ent/bill"
	"github.com/team10/app/ent/department"
	"github.com/team10/app/ent/doctorinfo"
	"github.com/team10/app/ent/educationlevel"
	"github.com/team10/app/ent/financier"
	"github.com/team10/app/ent/officeroom"
	"github.com/team10/app/ent/paytype"
	"github.com/team10/app/ent/prename"
	"github.com/team10/app/ent/schema"
	"github.com/team10/app/ent/treatment"
	"github.com/team10/app/ent/typetreatment"
	"github.com/team10/app/ent/unpaybill"
	"github.com/team10/app/ent/user"
)

// The init function reads all schema descriptors with runtime
// code (default values, validators or hooks) and stitches it
// to their package variables.
func init() {
	billFields := schema.Bill{}.Fields()
	_ = billFields
	// billDescAmount is the schema descriptor for Amount field.
	billDescAmount := billFields[0].Descriptor()
	// bill.AmountValidator is a validator for the "Amount" field. It is called by the builders before save.
	bill.AmountValidator = billDescAmount.Validators[0].(func(string) error)
	departmentFields := schema.Department{}.Fields()
	_ = departmentFields
	// departmentDescDepartment is the schema descriptor for department field.
	departmentDescDepartment := departmentFields[0].Descriptor()
	// department.DepartmentValidator is a validator for the "department" field. It is called by the builders before save.
	department.DepartmentValidator = departmentDescDepartment.Validators[0].(func(string) error)
	doctorinfoFields := schema.Doctorinfo{}.Fields()
	_ = doctorinfoFields
	// doctorinfoDescDoctorname is the schema descriptor for doctorname field.
	doctorinfoDescDoctorname := doctorinfoFields[0].Descriptor()
	// doctorinfo.DoctornameValidator is a validator for the "doctorname" field. It is called by the builders before save.
	doctorinfo.DoctornameValidator = doctorinfoDescDoctorname.Validators[0].(func(string) error)
	// doctorinfoDescDoctorsurname is the schema descriptor for doctorsurname field.
	doctorinfoDescDoctorsurname := doctorinfoFields[1].Descriptor()
	// doctorinfo.DoctorsurnameValidator is a validator for the "doctorsurname" field. It is called by the builders before save.
	doctorinfo.DoctorsurnameValidator = doctorinfoDescDoctorsurname.Validators[0].(func(string) error)
	// doctorinfoDescTelephonenumber is the schema descriptor for telephonenumber field.
	doctorinfoDescTelephonenumber := doctorinfoFields[2].Descriptor()
	// doctorinfo.TelephonenumberValidator is a validator for the "telephonenumber" field. It is called by the builders before save.
	doctorinfo.TelephonenumberValidator = doctorinfoDescTelephonenumber.Validators[0].(func(string) error)
	// doctorinfoDescLicensenumber is the schema descriptor for licensenumber field.
	doctorinfoDescLicensenumber := doctorinfoFields[3].Descriptor()
	// doctorinfo.LicensenumberValidator is a validator for the "licensenumber" field. It is called by the builders before save.
	doctorinfo.LicensenumberValidator = doctorinfoDescLicensenumber.Validators[0].(func(string) error)
	educationlevelFields := schema.Educationlevel{}.Fields()
	_ = educationlevelFields
	// educationlevelDescLevel is the schema descriptor for level field.
	educationlevelDescLevel := educationlevelFields[0].Descriptor()
	// educationlevel.LevelValidator is a validator for the "level" field. It is called by the builders before save.
	educationlevel.LevelValidator = educationlevelDescLevel.Validators[0].(func(string) error)
	financierFields := schema.Financier{}.Fields()
	_ = financierFields
	// financierDescName is the schema descriptor for name field.
	financierDescName := financierFields[0].Descriptor()
	// financier.NameValidator is a validator for the "name" field. It is called by the builders before save.
	financier.NameValidator = financierDescName.Validators[0].(func(string) error)
	officeroomFields := schema.Officeroom{}.Fields()
	_ = officeroomFields
	// officeroomDescRoomnumber is the schema descriptor for roomnumber field.
	officeroomDescRoomnumber := officeroomFields[0].Descriptor()
	// officeroom.RoomnumberValidator is a validator for the "roomnumber" field. It is called by the builders before save.
	officeroom.RoomnumberValidator = officeroomDescRoomnumber.Validators[0].(func(string) error)
	paytypeFields := schema.Paytype{}.Fields()
	_ = paytypeFields
	// paytypeDescPaytype is the schema descriptor for paytype field.
	paytypeDescPaytype := paytypeFields[0].Descriptor()
	// paytype.PaytypeValidator is a validator for the "paytype" field. It is called by the builders before save.
	paytype.PaytypeValidator = paytypeDescPaytype.Validators[0].(func(string) error)
	prenameFields := schema.Prename{}.Fields()
	_ = prenameFields
	// prenameDescPrefix is the schema descriptor for prefix field.
	prenameDescPrefix := prenameFields[0].Descriptor()
	// prename.PrefixValidator is a validator for the "prefix" field. It is called by the builders before save.
	prename.PrefixValidator = prenameDescPrefix.Validators[0].(func(string) error)
	treatmentFields := schema.Treatment{}.Fields()
	_ = treatmentFields
	// treatmentDescTreatment is the schema descriptor for Treatment field.
	treatmentDescTreatment := treatmentFields[0].Descriptor()
	// treatment.TreatmentValidator is a validator for the "Treatment" field. It is called by the builders before save.
	treatment.TreatmentValidator = treatmentDescTreatment.Validators[0].(func(string) error)
	typetreatmentFields := schema.Typetreatment{}.Fields()
	_ = typetreatmentFields
	// typetreatmentDescTypetreatment is the schema descriptor for Typetreatment field.
	typetreatmentDescTypetreatment := typetreatmentFields[0].Descriptor()
	// typetreatment.TypetreatmentValidator is a validator for the "Typetreatment" field. It is called by the builders before save.
	typetreatment.TypetreatmentValidator = typetreatmentDescTypetreatment.Validators[0].(func(string) error)
	unpaybillFields := schema.Unpaybill{}.Fields()
	_ = unpaybillFields
	// unpaybillDescStatus is the schema descriptor for Status field.
	unpaybillDescStatus := unpaybillFields[0].Descriptor()
	// unpaybill.StatusValidator is a validator for the "Status" field. It is called by the builders before save.
	unpaybill.StatusValidator = unpaybillDescStatus.Validators[0].(func(string) error)
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescEmail is the schema descriptor for email field.
	userDescEmail := userFields[0].Descriptor()
	// user.EmailValidator is a validator for the "email" field. It is called by the builders before save.
	user.EmailValidator = userDescEmail.Validators[0].(func(string) error)
	// userDescPassword is the schema descriptor for password field.
	userDescPassword := userFields[1].Descriptor()
	// user.PasswordValidator is a validator for the "password" field. It is called by the builders before save.
	user.PasswordValidator = userDescPassword.Validators[0].(func(string) error)
}
