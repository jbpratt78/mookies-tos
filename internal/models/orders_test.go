// Code generated by SQLBoiler 4.4.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"bytes"
	"context"
	"reflect"
	"testing"

	"github.com/volatiletech/randomize"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/strmangle"
)

var (
	// Relationships sometimes use the reflection helper queries.Equal/queries.Assign
	// so force a package dependency in case they don't.
	_ = queries.Equal
)

func testOrders(t *testing.T) {
	t.Parallel()

	query := Orders()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testOrdersDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Order{}
	if err = randomize.Struct(seed, o, orderDBTypes, true, orderColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Order struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := o.Delete(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Orders().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testOrdersQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Order{}
	if err = randomize.Struct(seed, o, orderDBTypes, true, orderColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Order struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := Orders().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Orders().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testOrdersSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Order{}
	if err = randomize.Struct(seed, o, orderDBTypes, true, orderColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Order struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := OrderSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Orders().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testOrdersExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Order{}
	if err = randomize.Struct(seed, o, orderDBTypes, true, orderColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Order struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := OrderExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if Order exists: %s", err)
	}
	if !e {
		t.Errorf("Expected OrderExists to return true, but got false.")
	}
}

func testOrdersFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Order{}
	if err = randomize.Struct(seed, o, orderDBTypes, true, orderColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Order struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	orderFound, err := FindOrder(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if orderFound == nil {
		t.Error("want a record, got nil")
	}
}

func testOrdersBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Order{}
	if err = randomize.Struct(seed, o, orderDBTypes, true, orderColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Order struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = Orders().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testOrdersOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Order{}
	if err = randomize.Struct(seed, o, orderDBTypes, true, orderColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Order struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := Orders().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testOrdersAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	orderOne := &Order{}
	orderTwo := &Order{}
	if err = randomize.Struct(seed, orderOne, orderDBTypes, false, orderColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Order struct: %s", err)
	}
	if err = randomize.Struct(seed, orderTwo, orderDBTypes, false, orderColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Order struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = orderOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = orderTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Orders().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testOrdersCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	orderOne := &Order{}
	orderTwo := &Order{}
	if err = randomize.Struct(seed, orderOne, orderDBTypes, false, orderColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Order struct: %s", err)
	}
	if err = randomize.Struct(seed, orderTwo, orderDBTypes, false, orderColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Order struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = orderOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = orderTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Orders().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func testOrdersInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Order{}
	if err = randomize.Struct(seed, o, orderDBTypes, true, orderColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Order struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Orders().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testOrdersInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Order{}
	if err = randomize.Struct(seed, o, orderDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Order struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(orderColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := Orders().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testOrderToManyOrderItemOptions(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Order
	var b, c OrderItemOption

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, orderDBTypes, true, orderColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Order struct: %s", err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, orderItemOptionDBTypes, false, orderItemOptionColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, orderItemOptionDBTypes, false, orderItemOptionColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	queries.Assign(&b.OrderID, a.ID)
	queries.Assign(&c.OrderID, a.ID)
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := a.OrderItemOptions().All(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range check {
		if queries.Equal(v.OrderID, b.OrderID) {
			bFound = true
		}
		if queries.Equal(v.OrderID, c.OrderID) {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := OrderSlice{&a}
	if err = a.L.LoadOrderItemOptions(ctx, tx, false, (*[]*Order)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.OrderItemOptions); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.OrderItemOptions = nil
	if err = a.L.LoadOrderItemOptions(ctx, tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.OrderItemOptions); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", check)
	}
}

func testOrderToManyOrderItems(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Order
	var b, c OrderItem

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, orderDBTypes, true, orderColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Order struct: %s", err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, orderItemDBTypes, false, orderItemColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, orderItemDBTypes, false, orderItemColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	queries.Assign(&b.OrderID, a.ID)
	queries.Assign(&c.OrderID, a.ID)
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := a.OrderItems().All(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range check {
		if queries.Equal(v.OrderID, b.OrderID) {
			bFound = true
		}
		if queries.Equal(v.OrderID, c.OrderID) {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := OrderSlice{&a}
	if err = a.L.LoadOrderItems(ctx, tx, false, (*[]*Order)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.OrderItems); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.OrderItems = nil
	if err = a.L.LoadOrderItems(ctx, tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.OrderItems); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", check)
	}
}

func testOrderToManyAddOpOrderItemOptions(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Order
	var b, c, d, e OrderItemOption

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, orderDBTypes, false, strmangle.SetComplement(orderPrimaryKeyColumns, orderColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*OrderItemOption{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, orderItemOptionDBTypes, false, strmangle.SetComplement(orderItemOptionPrimaryKeyColumns, orderItemOptionColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	foreignersSplitByInsertion := [][]*OrderItemOption{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddOrderItemOptions(ctx, tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if !queries.Equal(a.ID, first.OrderID) {
			t.Error("foreign key was wrong value", a.ID, first.OrderID)
		}
		if !queries.Equal(a.ID, second.OrderID) {
			t.Error("foreign key was wrong value", a.ID, second.OrderID)
		}

		if first.R.Order != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.Order != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.OrderItemOptions[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.OrderItemOptions[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.OrderItemOptions().Count(ctx, tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}
func testOrderToManyAddOpOrderItems(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Order
	var b, c, d, e OrderItem

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, orderDBTypes, false, strmangle.SetComplement(orderPrimaryKeyColumns, orderColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*OrderItem{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, orderItemDBTypes, false, strmangle.SetComplement(orderItemPrimaryKeyColumns, orderItemColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	foreignersSplitByInsertion := [][]*OrderItem{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddOrderItems(ctx, tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if !queries.Equal(a.ID, first.OrderID) {
			t.Error("foreign key was wrong value", a.ID, first.OrderID)
		}
		if !queries.Equal(a.ID, second.OrderID) {
			t.Error("foreign key was wrong value", a.ID, second.OrderID)
		}

		if first.R.Order != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.Order != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.OrderItems[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.OrderItems[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.OrderItems().Count(ctx, tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}

func testOrdersReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Order{}
	if err = randomize.Struct(seed, o, orderDBTypes, true, orderColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Order struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = o.Reload(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testOrdersReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Order{}
	if err = randomize.Struct(seed, o, orderDBTypes, true, orderColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Order struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := OrderSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testOrdersSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Order{}
	if err = randomize.Struct(seed, o, orderDBTypes, true, orderColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Order struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Orders().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	orderDBTypes = map[string]string{`ID`: `INTEGER`, `SubmittedAt`: `INTEGER`, `CompletedAt`: `INTEGER`}
	_            = bytes.MinRead
)

func testOrdersUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(orderPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(orderAllColumns) == len(orderPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Order{}
	if err = randomize.Struct(seed, o, orderDBTypes, true, orderColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Order struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Orders().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, orderDBTypes, true, orderPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Order struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testOrdersSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(orderAllColumns) == len(orderPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Order{}
	if err = randomize.Struct(seed, o, orderDBTypes, true, orderColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Order struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Orders().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, orderDBTypes, true, orderPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Order struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(orderAllColumns, orderPrimaryKeyColumns) {
		fields = orderAllColumns
	} else {
		fields = strmangle.SetComplement(
			orderAllColumns,
			orderPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	typ := reflect.TypeOf(o).Elem()
	n := typ.NumField()

	updateMap := M{}
	for _, col := range fields {
		for i := 0; i < n; i++ {
			f := typ.Field(i)
			if f.Tag.Get("boil") == col {
				updateMap[col] = value.Field(i).Interface()
			}
		}
	}

	slice := OrderSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}
