// Code generated by ent, DO NOT EDIT.

package vendor

import (
	"spoolman-go/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Vendor {
	return predicate.Vendor(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Vendor {
	return predicate.Vendor(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Vendor {
	return predicate.Vendor(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Vendor {
	return predicate.Vendor(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Vendor {
	return predicate.Vendor(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Vendor {
	return predicate.Vendor(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Vendor {
	return predicate.Vendor(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Vendor {
	return predicate.Vendor(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Vendor {
	return predicate.Vendor(sql.FieldLTE(FieldID, id))
}

// Registered applies equality check predicate on the "registered" field. It's identical to RegisteredEQ.
func Registered(v time.Time) predicate.Vendor {
	return predicate.Vendor(sql.FieldEQ(FieldRegistered, v))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Vendor {
	return predicate.Vendor(sql.FieldEQ(FieldName, v))
}

// EmptySpoolWeight applies equality check predicate on the "empty_spool_weight" field. It's identical to EmptySpoolWeightEQ.
func EmptySpoolWeight(v float64) predicate.Vendor {
	return predicate.Vendor(sql.FieldEQ(FieldEmptySpoolWeight, v))
}

// Comment applies equality check predicate on the "comment" field. It's identical to CommentEQ.
func Comment(v string) predicate.Vendor {
	return predicate.Vendor(sql.FieldEQ(FieldComment, v))
}

// ExternalID applies equality check predicate on the "external_id" field. It's identical to ExternalIDEQ.
func ExternalID(v string) predicate.Vendor {
	return predicate.Vendor(sql.FieldEQ(FieldExternalID, v))
}

// RegisteredEQ applies the EQ predicate on the "registered" field.
func RegisteredEQ(v time.Time) predicate.Vendor {
	return predicate.Vendor(sql.FieldEQ(FieldRegistered, v))
}

// RegisteredNEQ applies the NEQ predicate on the "registered" field.
func RegisteredNEQ(v time.Time) predicate.Vendor {
	return predicate.Vendor(sql.FieldNEQ(FieldRegistered, v))
}

// RegisteredIn applies the In predicate on the "registered" field.
func RegisteredIn(vs ...time.Time) predicate.Vendor {
	return predicate.Vendor(sql.FieldIn(FieldRegistered, vs...))
}

// RegisteredNotIn applies the NotIn predicate on the "registered" field.
func RegisteredNotIn(vs ...time.Time) predicate.Vendor {
	return predicate.Vendor(sql.FieldNotIn(FieldRegistered, vs...))
}

// RegisteredGT applies the GT predicate on the "registered" field.
func RegisteredGT(v time.Time) predicate.Vendor {
	return predicate.Vendor(sql.FieldGT(FieldRegistered, v))
}

// RegisteredGTE applies the GTE predicate on the "registered" field.
func RegisteredGTE(v time.Time) predicate.Vendor {
	return predicate.Vendor(sql.FieldGTE(FieldRegistered, v))
}

// RegisteredLT applies the LT predicate on the "registered" field.
func RegisteredLT(v time.Time) predicate.Vendor {
	return predicate.Vendor(sql.FieldLT(FieldRegistered, v))
}

// RegisteredLTE applies the LTE predicate on the "registered" field.
func RegisteredLTE(v time.Time) predicate.Vendor {
	return predicate.Vendor(sql.FieldLTE(FieldRegistered, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Vendor {
	return predicate.Vendor(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Vendor {
	return predicate.Vendor(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Vendor {
	return predicate.Vendor(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Vendor {
	return predicate.Vendor(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Vendor {
	return predicate.Vendor(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Vendor {
	return predicate.Vendor(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Vendor {
	return predicate.Vendor(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Vendor {
	return predicate.Vendor(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Vendor {
	return predicate.Vendor(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Vendor {
	return predicate.Vendor(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Vendor {
	return predicate.Vendor(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Vendor {
	return predicate.Vendor(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Vendor {
	return predicate.Vendor(sql.FieldContainsFold(FieldName, v))
}

// EmptySpoolWeightEQ applies the EQ predicate on the "empty_spool_weight" field.
func EmptySpoolWeightEQ(v float64) predicate.Vendor {
	return predicate.Vendor(sql.FieldEQ(FieldEmptySpoolWeight, v))
}

// EmptySpoolWeightNEQ applies the NEQ predicate on the "empty_spool_weight" field.
func EmptySpoolWeightNEQ(v float64) predicate.Vendor {
	return predicate.Vendor(sql.FieldNEQ(FieldEmptySpoolWeight, v))
}

// EmptySpoolWeightIn applies the In predicate on the "empty_spool_weight" field.
func EmptySpoolWeightIn(vs ...float64) predicate.Vendor {
	return predicate.Vendor(sql.FieldIn(FieldEmptySpoolWeight, vs...))
}

// EmptySpoolWeightNotIn applies the NotIn predicate on the "empty_spool_weight" field.
func EmptySpoolWeightNotIn(vs ...float64) predicate.Vendor {
	return predicate.Vendor(sql.FieldNotIn(FieldEmptySpoolWeight, vs...))
}

// EmptySpoolWeightGT applies the GT predicate on the "empty_spool_weight" field.
func EmptySpoolWeightGT(v float64) predicate.Vendor {
	return predicate.Vendor(sql.FieldGT(FieldEmptySpoolWeight, v))
}

// EmptySpoolWeightGTE applies the GTE predicate on the "empty_spool_weight" field.
func EmptySpoolWeightGTE(v float64) predicate.Vendor {
	return predicate.Vendor(sql.FieldGTE(FieldEmptySpoolWeight, v))
}

// EmptySpoolWeightLT applies the LT predicate on the "empty_spool_weight" field.
func EmptySpoolWeightLT(v float64) predicate.Vendor {
	return predicate.Vendor(sql.FieldLT(FieldEmptySpoolWeight, v))
}

// EmptySpoolWeightLTE applies the LTE predicate on the "empty_spool_weight" field.
func EmptySpoolWeightLTE(v float64) predicate.Vendor {
	return predicate.Vendor(sql.FieldLTE(FieldEmptySpoolWeight, v))
}

// EmptySpoolWeightIsNil applies the IsNil predicate on the "empty_spool_weight" field.
func EmptySpoolWeightIsNil() predicate.Vendor {
	return predicate.Vendor(sql.FieldIsNull(FieldEmptySpoolWeight))
}

// EmptySpoolWeightNotNil applies the NotNil predicate on the "empty_spool_weight" field.
func EmptySpoolWeightNotNil() predicate.Vendor {
	return predicate.Vendor(sql.FieldNotNull(FieldEmptySpoolWeight))
}

// CommentEQ applies the EQ predicate on the "comment" field.
func CommentEQ(v string) predicate.Vendor {
	return predicate.Vendor(sql.FieldEQ(FieldComment, v))
}

// CommentNEQ applies the NEQ predicate on the "comment" field.
func CommentNEQ(v string) predicate.Vendor {
	return predicate.Vendor(sql.FieldNEQ(FieldComment, v))
}

// CommentIn applies the In predicate on the "comment" field.
func CommentIn(vs ...string) predicate.Vendor {
	return predicate.Vendor(sql.FieldIn(FieldComment, vs...))
}

// CommentNotIn applies the NotIn predicate on the "comment" field.
func CommentNotIn(vs ...string) predicate.Vendor {
	return predicate.Vendor(sql.FieldNotIn(FieldComment, vs...))
}

// CommentGT applies the GT predicate on the "comment" field.
func CommentGT(v string) predicate.Vendor {
	return predicate.Vendor(sql.FieldGT(FieldComment, v))
}

// CommentGTE applies the GTE predicate on the "comment" field.
func CommentGTE(v string) predicate.Vendor {
	return predicate.Vendor(sql.FieldGTE(FieldComment, v))
}

// CommentLT applies the LT predicate on the "comment" field.
func CommentLT(v string) predicate.Vendor {
	return predicate.Vendor(sql.FieldLT(FieldComment, v))
}

// CommentLTE applies the LTE predicate on the "comment" field.
func CommentLTE(v string) predicate.Vendor {
	return predicate.Vendor(sql.FieldLTE(FieldComment, v))
}

// CommentContains applies the Contains predicate on the "comment" field.
func CommentContains(v string) predicate.Vendor {
	return predicate.Vendor(sql.FieldContains(FieldComment, v))
}

// CommentHasPrefix applies the HasPrefix predicate on the "comment" field.
func CommentHasPrefix(v string) predicate.Vendor {
	return predicate.Vendor(sql.FieldHasPrefix(FieldComment, v))
}

// CommentHasSuffix applies the HasSuffix predicate on the "comment" field.
func CommentHasSuffix(v string) predicate.Vendor {
	return predicate.Vendor(sql.FieldHasSuffix(FieldComment, v))
}

// CommentIsNil applies the IsNil predicate on the "comment" field.
func CommentIsNil() predicate.Vendor {
	return predicate.Vendor(sql.FieldIsNull(FieldComment))
}

// CommentNotNil applies the NotNil predicate on the "comment" field.
func CommentNotNil() predicate.Vendor {
	return predicate.Vendor(sql.FieldNotNull(FieldComment))
}

// CommentEqualFold applies the EqualFold predicate on the "comment" field.
func CommentEqualFold(v string) predicate.Vendor {
	return predicate.Vendor(sql.FieldEqualFold(FieldComment, v))
}

// CommentContainsFold applies the ContainsFold predicate on the "comment" field.
func CommentContainsFold(v string) predicate.Vendor {
	return predicate.Vendor(sql.FieldContainsFold(FieldComment, v))
}

// ExternalIDEQ applies the EQ predicate on the "external_id" field.
func ExternalIDEQ(v string) predicate.Vendor {
	return predicate.Vendor(sql.FieldEQ(FieldExternalID, v))
}

// ExternalIDNEQ applies the NEQ predicate on the "external_id" field.
func ExternalIDNEQ(v string) predicate.Vendor {
	return predicate.Vendor(sql.FieldNEQ(FieldExternalID, v))
}

// ExternalIDIn applies the In predicate on the "external_id" field.
func ExternalIDIn(vs ...string) predicate.Vendor {
	return predicate.Vendor(sql.FieldIn(FieldExternalID, vs...))
}

// ExternalIDNotIn applies the NotIn predicate on the "external_id" field.
func ExternalIDNotIn(vs ...string) predicate.Vendor {
	return predicate.Vendor(sql.FieldNotIn(FieldExternalID, vs...))
}

// ExternalIDGT applies the GT predicate on the "external_id" field.
func ExternalIDGT(v string) predicate.Vendor {
	return predicate.Vendor(sql.FieldGT(FieldExternalID, v))
}

// ExternalIDGTE applies the GTE predicate on the "external_id" field.
func ExternalIDGTE(v string) predicate.Vendor {
	return predicate.Vendor(sql.FieldGTE(FieldExternalID, v))
}

// ExternalIDLT applies the LT predicate on the "external_id" field.
func ExternalIDLT(v string) predicate.Vendor {
	return predicate.Vendor(sql.FieldLT(FieldExternalID, v))
}

// ExternalIDLTE applies the LTE predicate on the "external_id" field.
func ExternalIDLTE(v string) predicate.Vendor {
	return predicate.Vendor(sql.FieldLTE(FieldExternalID, v))
}

// ExternalIDContains applies the Contains predicate on the "external_id" field.
func ExternalIDContains(v string) predicate.Vendor {
	return predicate.Vendor(sql.FieldContains(FieldExternalID, v))
}

// ExternalIDHasPrefix applies the HasPrefix predicate on the "external_id" field.
func ExternalIDHasPrefix(v string) predicate.Vendor {
	return predicate.Vendor(sql.FieldHasPrefix(FieldExternalID, v))
}

// ExternalIDHasSuffix applies the HasSuffix predicate on the "external_id" field.
func ExternalIDHasSuffix(v string) predicate.Vendor {
	return predicate.Vendor(sql.FieldHasSuffix(FieldExternalID, v))
}

// ExternalIDIsNil applies the IsNil predicate on the "external_id" field.
func ExternalIDIsNil() predicate.Vendor {
	return predicate.Vendor(sql.FieldIsNull(FieldExternalID))
}

// ExternalIDNotNil applies the NotNil predicate on the "external_id" field.
func ExternalIDNotNil() predicate.Vendor {
	return predicate.Vendor(sql.FieldNotNull(FieldExternalID))
}

// ExternalIDEqualFold applies the EqualFold predicate on the "external_id" field.
func ExternalIDEqualFold(v string) predicate.Vendor {
	return predicate.Vendor(sql.FieldEqualFold(FieldExternalID, v))
}

// ExternalIDContainsFold applies the ContainsFold predicate on the "external_id" field.
func ExternalIDContainsFold(v string) predicate.Vendor {
	return predicate.Vendor(sql.FieldContainsFold(FieldExternalID, v))
}

// HasFilaments applies the HasEdge predicate on the "filaments" edge.
func HasFilaments() predicate.Vendor {
	return predicate.Vendor(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, FilamentsTable, FilamentsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasFilamentsWith applies the HasEdge predicate on the "filaments" edge with a given conditions (other predicates).
func HasFilamentsWith(preds ...predicate.Filament) predicate.Vendor {
	return predicate.Vendor(func(s *sql.Selector) {
		step := newFilamentsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasExtra applies the HasEdge predicate on the "extra" edge.
func HasExtra() predicate.Vendor {
	return predicate.Vendor(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ExtraTable, ExtraColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasExtraWith applies the HasEdge predicate on the "extra" edge with a given conditions (other predicates).
func HasExtraWith(preds ...predicate.VendorField) predicate.Vendor {
	return predicate.Vendor(func(s *sql.Selector) {
		step := newExtraStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Vendor) predicate.Vendor {
	return predicate.Vendor(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Vendor) predicate.Vendor {
	return predicate.Vendor(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Vendor) predicate.Vendor {
	return predicate.Vendor(sql.NotPredicates(p))
}
