// Code generated by entc, DO NOT EDIT.

package wuxing

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/babyname/fate/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id string) predicate.WuXing {
	return predicate.WuXing(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id string) predicate.WuXing {
	return predicate.WuXing(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id string) predicate.WuXing {
	return predicate.WuXing(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...string) predicate.WuXing {
	return predicate.WuXing(func(s *sql.Selector) {
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
func IDNotIn(ids ...string) predicate.WuXing {
	return predicate.WuXing(func(s *sql.Selector) {
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
func IDGT(id string) predicate.WuXing {
	return predicate.WuXing(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id string) predicate.WuXing {
	return predicate.WuXing(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id string) predicate.WuXing {
	return predicate.WuXing(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id string) predicate.WuXing {
	return predicate.WuXing(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Created applies equality check predicate on the "created" field. It's identical to CreatedEQ.
func Created(v time.Time) predicate.WuXing {
	return predicate.WuXing(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreated), v))
	})
}

// Updated applies equality check predicate on the "updated" field. It's identical to UpdatedEQ.
func Updated(v time.Time) predicate.WuXing {
	return predicate.WuXing(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdated), v))
	})
}

// Deleted applies equality check predicate on the "deleted" field. It's identical to DeletedEQ.
func Deleted(v time.Time) predicate.WuXing {
	return predicate.WuXing(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeleted), v))
	})
}

// Version applies equality check predicate on the "version" field. It's identical to VersionEQ.
func Version(v int32) predicate.WuXing {
	return predicate.WuXing(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldVersion), v))
	})
}

// First applies equality check predicate on the "first" field. It's identical to FirstEQ.
func First(v string) predicate.WuXing {
	return predicate.WuXing(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldFirst), v))
	})
}

// Second applies equality check predicate on the "second" field. It's identical to SecondEQ.
func Second(v string) predicate.WuXing {
	return predicate.WuXing(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSecond), v))
	})
}

// Third applies equality check predicate on the "third" field. It's identical to ThirdEQ.
func Third(v string) predicate.WuXing {
	return predicate.WuXing(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldThird), v))
	})
}

// Fortune applies equality check predicate on the "fortune" field. It's identical to FortuneEQ.
func Fortune(v string) predicate.WuXing {
	return predicate.WuXing(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldFortune), v))
	})
}

// CreatedEQ applies the EQ predicate on the "created" field.
func CreatedEQ(v time.Time) predicate.WuXing {
	return predicate.WuXing(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreated), v))
	})
}

// CreatedNEQ applies the NEQ predicate on the "created" field.
func CreatedNEQ(v time.Time) predicate.WuXing {
	return predicate.WuXing(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreated), v))
	})
}

// CreatedIn applies the In predicate on the "created" field.
func CreatedIn(vs ...time.Time) predicate.WuXing {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.WuXing(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCreated), v...))
	})
}

// CreatedNotIn applies the NotIn predicate on the "created" field.
func CreatedNotIn(vs ...time.Time) predicate.WuXing {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.WuXing(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCreated), v...))
	})
}

// CreatedGT applies the GT predicate on the "created" field.
func CreatedGT(v time.Time) predicate.WuXing {
	return predicate.WuXing(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreated), v))
	})
}

// CreatedGTE applies the GTE predicate on the "created" field.
func CreatedGTE(v time.Time) predicate.WuXing {
	return predicate.WuXing(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreated), v))
	})
}

// CreatedLT applies the LT predicate on the "created" field.
func CreatedLT(v time.Time) predicate.WuXing {
	return predicate.WuXing(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreated), v))
	})
}

// CreatedLTE applies the LTE predicate on the "created" field.
func CreatedLTE(v time.Time) predicate.WuXing {
	return predicate.WuXing(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreated), v))
	})
}

// CreatedIsNil applies the IsNil predicate on the "created" field.
func CreatedIsNil() predicate.WuXing {
	return predicate.WuXing(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldCreated)))
	})
}

// CreatedNotNil applies the NotNil predicate on the "created" field.
func CreatedNotNil() predicate.WuXing {
	return predicate.WuXing(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldCreated)))
	})
}

// UpdatedEQ applies the EQ predicate on the "updated" field.
func UpdatedEQ(v time.Time) predicate.WuXing {
	return predicate.WuXing(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdated), v))
	})
}

// UpdatedNEQ applies the NEQ predicate on the "updated" field.
func UpdatedNEQ(v time.Time) predicate.WuXing {
	return predicate.WuXing(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdated), v))
	})
}

// UpdatedIn applies the In predicate on the "updated" field.
func UpdatedIn(vs ...time.Time) predicate.WuXing {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.WuXing(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldUpdated), v...))
	})
}

// UpdatedNotIn applies the NotIn predicate on the "updated" field.
func UpdatedNotIn(vs ...time.Time) predicate.WuXing {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.WuXing(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldUpdated), v...))
	})
}

// UpdatedGT applies the GT predicate on the "updated" field.
func UpdatedGT(v time.Time) predicate.WuXing {
	return predicate.WuXing(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdated), v))
	})
}

// UpdatedGTE applies the GTE predicate on the "updated" field.
func UpdatedGTE(v time.Time) predicate.WuXing {
	return predicate.WuXing(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdated), v))
	})
}

// UpdatedLT applies the LT predicate on the "updated" field.
func UpdatedLT(v time.Time) predicate.WuXing {
	return predicate.WuXing(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdated), v))
	})
}

// UpdatedLTE applies the LTE predicate on the "updated" field.
func UpdatedLTE(v time.Time) predicate.WuXing {
	return predicate.WuXing(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdated), v))
	})
}

// UpdatedIsNil applies the IsNil predicate on the "updated" field.
func UpdatedIsNil() predicate.WuXing {
	return predicate.WuXing(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldUpdated)))
	})
}

// UpdatedNotNil applies the NotNil predicate on the "updated" field.
func UpdatedNotNil() predicate.WuXing {
	return predicate.WuXing(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldUpdated)))
	})
}

// DeletedEQ applies the EQ predicate on the "deleted" field.
func DeletedEQ(v time.Time) predicate.WuXing {
	return predicate.WuXing(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeleted), v))
	})
}

// DeletedNEQ applies the NEQ predicate on the "deleted" field.
func DeletedNEQ(v time.Time) predicate.WuXing {
	return predicate.WuXing(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDeleted), v))
	})
}

// DeletedIn applies the In predicate on the "deleted" field.
func DeletedIn(vs ...time.Time) predicate.WuXing {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.WuXing(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldDeleted), v...))
	})
}

// DeletedNotIn applies the NotIn predicate on the "deleted" field.
func DeletedNotIn(vs ...time.Time) predicate.WuXing {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.WuXing(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldDeleted), v...))
	})
}

// DeletedGT applies the GT predicate on the "deleted" field.
func DeletedGT(v time.Time) predicate.WuXing {
	return predicate.WuXing(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDeleted), v))
	})
}

// DeletedGTE applies the GTE predicate on the "deleted" field.
func DeletedGTE(v time.Time) predicate.WuXing {
	return predicate.WuXing(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDeleted), v))
	})
}

// DeletedLT applies the LT predicate on the "deleted" field.
func DeletedLT(v time.Time) predicate.WuXing {
	return predicate.WuXing(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDeleted), v))
	})
}

// DeletedLTE applies the LTE predicate on the "deleted" field.
func DeletedLTE(v time.Time) predicate.WuXing {
	return predicate.WuXing(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDeleted), v))
	})
}

// DeletedIsNil applies the IsNil predicate on the "deleted" field.
func DeletedIsNil() predicate.WuXing {
	return predicate.WuXing(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldDeleted)))
	})
}

// DeletedNotNil applies the NotNil predicate on the "deleted" field.
func DeletedNotNil() predicate.WuXing {
	return predicate.WuXing(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldDeleted)))
	})
}

// VersionEQ applies the EQ predicate on the "version" field.
func VersionEQ(v int32) predicate.WuXing {
	return predicate.WuXing(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldVersion), v))
	})
}

// VersionNEQ applies the NEQ predicate on the "version" field.
func VersionNEQ(v int32) predicate.WuXing {
	return predicate.WuXing(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldVersion), v))
	})
}

// VersionIn applies the In predicate on the "version" field.
func VersionIn(vs ...int32) predicate.WuXing {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.WuXing(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldVersion), v...))
	})
}

// VersionNotIn applies the NotIn predicate on the "version" field.
func VersionNotIn(vs ...int32) predicate.WuXing {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.WuXing(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldVersion), v...))
	})
}

// VersionGT applies the GT predicate on the "version" field.
func VersionGT(v int32) predicate.WuXing {
	return predicate.WuXing(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldVersion), v))
	})
}

// VersionGTE applies the GTE predicate on the "version" field.
func VersionGTE(v int32) predicate.WuXing {
	return predicate.WuXing(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldVersion), v))
	})
}

// VersionLT applies the LT predicate on the "version" field.
func VersionLT(v int32) predicate.WuXing {
	return predicate.WuXing(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldVersion), v))
	})
}

// VersionLTE applies the LTE predicate on the "version" field.
func VersionLTE(v int32) predicate.WuXing {
	return predicate.WuXing(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldVersion), v))
	})
}

// VersionIsNil applies the IsNil predicate on the "version" field.
func VersionIsNil() predicate.WuXing {
	return predicate.WuXing(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldVersion)))
	})
}

// VersionNotNil applies the NotNil predicate on the "version" field.
func VersionNotNil() predicate.WuXing {
	return predicate.WuXing(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldVersion)))
	})
}

// FirstEQ applies the EQ predicate on the "first" field.
func FirstEQ(v string) predicate.WuXing {
	return predicate.WuXing(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldFirst), v))
	})
}

// FirstNEQ applies the NEQ predicate on the "first" field.
func FirstNEQ(v string) predicate.WuXing {
	return predicate.WuXing(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldFirst), v))
	})
}

// FirstIn applies the In predicate on the "first" field.
func FirstIn(vs ...string) predicate.WuXing {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.WuXing(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldFirst), v...))
	})
}

// FirstNotIn applies the NotIn predicate on the "first" field.
func FirstNotIn(vs ...string) predicate.WuXing {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.WuXing(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldFirst), v...))
	})
}

// FirstGT applies the GT predicate on the "first" field.
func FirstGT(v string) predicate.WuXing {
	return predicate.WuXing(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldFirst), v))
	})
}

// FirstGTE applies the GTE predicate on the "first" field.
func FirstGTE(v string) predicate.WuXing {
	return predicate.WuXing(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldFirst), v))
	})
}

// FirstLT applies the LT predicate on the "first" field.
func FirstLT(v string) predicate.WuXing {
	return predicate.WuXing(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldFirst), v))
	})
}

// FirstLTE applies the LTE predicate on the "first" field.
func FirstLTE(v string) predicate.WuXing {
	return predicate.WuXing(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldFirst), v))
	})
}

// FirstContains applies the Contains predicate on the "first" field.
func FirstContains(v string) predicate.WuXing {
	return predicate.WuXing(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldFirst), v))
	})
}

// FirstHasPrefix applies the HasPrefix predicate on the "first" field.
func FirstHasPrefix(v string) predicate.WuXing {
	return predicate.WuXing(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldFirst), v))
	})
}

// FirstHasSuffix applies the HasSuffix predicate on the "first" field.
func FirstHasSuffix(v string) predicate.WuXing {
	return predicate.WuXing(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldFirst), v))
	})
}

// FirstIsNil applies the IsNil predicate on the "first" field.
func FirstIsNil() predicate.WuXing {
	return predicate.WuXing(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldFirst)))
	})
}

// FirstNotNil applies the NotNil predicate on the "first" field.
func FirstNotNil() predicate.WuXing {
	return predicate.WuXing(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldFirst)))
	})
}

// FirstEqualFold applies the EqualFold predicate on the "first" field.
func FirstEqualFold(v string) predicate.WuXing {
	return predicate.WuXing(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldFirst), v))
	})
}

// FirstContainsFold applies the ContainsFold predicate on the "first" field.
func FirstContainsFold(v string) predicate.WuXing {
	return predicate.WuXing(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldFirst), v))
	})
}

// SecondEQ applies the EQ predicate on the "second" field.
func SecondEQ(v string) predicate.WuXing {
	return predicate.WuXing(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSecond), v))
	})
}

// SecondNEQ applies the NEQ predicate on the "second" field.
func SecondNEQ(v string) predicate.WuXing {
	return predicate.WuXing(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldSecond), v))
	})
}

// SecondIn applies the In predicate on the "second" field.
func SecondIn(vs ...string) predicate.WuXing {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.WuXing(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldSecond), v...))
	})
}

// SecondNotIn applies the NotIn predicate on the "second" field.
func SecondNotIn(vs ...string) predicate.WuXing {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.WuXing(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldSecond), v...))
	})
}

// SecondGT applies the GT predicate on the "second" field.
func SecondGT(v string) predicate.WuXing {
	return predicate.WuXing(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldSecond), v))
	})
}

// SecondGTE applies the GTE predicate on the "second" field.
func SecondGTE(v string) predicate.WuXing {
	return predicate.WuXing(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldSecond), v))
	})
}

// SecondLT applies the LT predicate on the "second" field.
func SecondLT(v string) predicate.WuXing {
	return predicate.WuXing(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldSecond), v))
	})
}

// SecondLTE applies the LTE predicate on the "second" field.
func SecondLTE(v string) predicate.WuXing {
	return predicate.WuXing(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldSecond), v))
	})
}

// SecondContains applies the Contains predicate on the "second" field.
func SecondContains(v string) predicate.WuXing {
	return predicate.WuXing(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldSecond), v))
	})
}

// SecondHasPrefix applies the HasPrefix predicate on the "second" field.
func SecondHasPrefix(v string) predicate.WuXing {
	return predicate.WuXing(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldSecond), v))
	})
}

// SecondHasSuffix applies the HasSuffix predicate on the "second" field.
func SecondHasSuffix(v string) predicate.WuXing {
	return predicate.WuXing(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldSecond), v))
	})
}

// SecondIsNil applies the IsNil predicate on the "second" field.
func SecondIsNil() predicate.WuXing {
	return predicate.WuXing(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldSecond)))
	})
}

// SecondNotNil applies the NotNil predicate on the "second" field.
func SecondNotNil() predicate.WuXing {
	return predicate.WuXing(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldSecond)))
	})
}

// SecondEqualFold applies the EqualFold predicate on the "second" field.
func SecondEqualFold(v string) predicate.WuXing {
	return predicate.WuXing(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldSecond), v))
	})
}

// SecondContainsFold applies the ContainsFold predicate on the "second" field.
func SecondContainsFold(v string) predicate.WuXing {
	return predicate.WuXing(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldSecond), v))
	})
}

// ThirdEQ applies the EQ predicate on the "third" field.
func ThirdEQ(v string) predicate.WuXing {
	return predicate.WuXing(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldThird), v))
	})
}

// ThirdNEQ applies the NEQ predicate on the "third" field.
func ThirdNEQ(v string) predicate.WuXing {
	return predicate.WuXing(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldThird), v))
	})
}

// ThirdIn applies the In predicate on the "third" field.
func ThirdIn(vs ...string) predicate.WuXing {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.WuXing(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldThird), v...))
	})
}

// ThirdNotIn applies the NotIn predicate on the "third" field.
func ThirdNotIn(vs ...string) predicate.WuXing {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.WuXing(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldThird), v...))
	})
}

// ThirdGT applies the GT predicate on the "third" field.
func ThirdGT(v string) predicate.WuXing {
	return predicate.WuXing(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldThird), v))
	})
}

// ThirdGTE applies the GTE predicate on the "third" field.
func ThirdGTE(v string) predicate.WuXing {
	return predicate.WuXing(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldThird), v))
	})
}

// ThirdLT applies the LT predicate on the "third" field.
func ThirdLT(v string) predicate.WuXing {
	return predicate.WuXing(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldThird), v))
	})
}

// ThirdLTE applies the LTE predicate on the "third" field.
func ThirdLTE(v string) predicate.WuXing {
	return predicate.WuXing(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldThird), v))
	})
}

// ThirdContains applies the Contains predicate on the "third" field.
func ThirdContains(v string) predicate.WuXing {
	return predicate.WuXing(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldThird), v))
	})
}

// ThirdHasPrefix applies the HasPrefix predicate on the "third" field.
func ThirdHasPrefix(v string) predicate.WuXing {
	return predicate.WuXing(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldThird), v))
	})
}

// ThirdHasSuffix applies the HasSuffix predicate on the "third" field.
func ThirdHasSuffix(v string) predicate.WuXing {
	return predicate.WuXing(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldThird), v))
	})
}

// ThirdIsNil applies the IsNil predicate on the "third" field.
func ThirdIsNil() predicate.WuXing {
	return predicate.WuXing(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldThird)))
	})
}

// ThirdNotNil applies the NotNil predicate on the "third" field.
func ThirdNotNil() predicate.WuXing {
	return predicate.WuXing(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldThird)))
	})
}

// ThirdEqualFold applies the EqualFold predicate on the "third" field.
func ThirdEqualFold(v string) predicate.WuXing {
	return predicate.WuXing(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldThird), v))
	})
}

// ThirdContainsFold applies the ContainsFold predicate on the "third" field.
func ThirdContainsFold(v string) predicate.WuXing {
	return predicate.WuXing(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldThird), v))
	})
}

// FortuneEQ applies the EQ predicate on the "fortune" field.
func FortuneEQ(v string) predicate.WuXing {
	return predicate.WuXing(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldFortune), v))
	})
}

// FortuneNEQ applies the NEQ predicate on the "fortune" field.
func FortuneNEQ(v string) predicate.WuXing {
	return predicate.WuXing(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldFortune), v))
	})
}

// FortuneIn applies the In predicate on the "fortune" field.
func FortuneIn(vs ...string) predicate.WuXing {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.WuXing(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldFortune), v...))
	})
}

// FortuneNotIn applies the NotIn predicate on the "fortune" field.
func FortuneNotIn(vs ...string) predicate.WuXing {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.WuXing(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldFortune), v...))
	})
}

// FortuneGT applies the GT predicate on the "fortune" field.
func FortuneGT(v string) predicate.WuXing {
	return predicate.WuXing(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldFortune), v))
	})
}

// FortuneGTE applies the GTE predicate on the "fortune" field.
func FortuneGTE(v string) predicate.WuXing {
	return predicate.WuXing(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldFortune), v))
	})
}

// FortuneLT applies the LT predicate on the "fortune" field.
func FortuneLT(v string) predicate.WuXing {
	return predicate.WuXing(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldFortune), v))
	})
}

// FortuneLTE applies the LTE predicate on the "fortune" field.
func FortuneLTE(v string) predicate.WuXing {
	return predicate.WuXing(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldFortune), v))
	})
}

// FortuneContains applies the Contains predicate on the "fortune" field.
func FortuneContains(v string) predicate.WuXing {
	return predicate.WuXing(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldFortune), v))
	})
}

// FortuneHasPrefix applies the HasPrefix predicate on the "fortune" field.
func FortuneHasPrefix(v string) predicate.WuXing {
	return predicate.WuXing(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldFortune), v))
	})
}

// FortuneHasSuffix applies the HasSuffix predicate on the "fortune" field.
func FortuneHasSuffix(v string) predicate.WuXing {
	return predicate.WuXing(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldFortune), v))
	})
}

// FortuneIsNil applies the IsNil predicate on the "fortune" field.
func FortuneIsNil() predicate.WuXing {
	return predicate.WuXing(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldFortune)))
	})
}

// FortuneNotNil applies the NotNil predicate on the "fortune" field.
func FortuneNotNil() predicate.WuXing {
	return predicate.WuXing(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldFortune)))
	})
}

// FortuneEqualFold applies the EqualFold predicate on the "fortune" field.
func FortuneEqualFold(v string) predicate.WuXing {
	return predicate.WuXing(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldFortune), v))
	})
}

// FortuneContainsFold applies the ContainsFold predicate on the "fortune" field.
func FortuneContainsFold(v string) predicate.WuXing {
	return predicate.WuXing(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldFortune), v))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.WuXing) predicate.WuXing {
	return predicate.WuXing(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.WuXing) predicate.WuXing {
	return predicate.WuXing(func(s *sql.Selector) {
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
func Not(p predicate.WuXing) predicate.WuXing {
	return predicate.WuXing(func(s *sql.Selector) {
		p(s.Not())
	})
}
