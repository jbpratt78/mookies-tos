// Code generated by SQLBoiler 4.4.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strings"
	"sync"
	"time"

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// Order is an object representing the database table.
type Order struct {
	ID          null.Int64 `boil:"id" json:"id,omitempty" toml:"id" yaml:"id,omitempty"`
	CreatedAt   int64      `boil:"created_at" json:"createdAt" toml:"createdAt" yaml:"createdAt"`
	CompletedAt null.Int64 `boil:"completed_at" json:"completedAt,omitempty" toml:"completedAt" yaml:"completedAt,omitempty"`

	R *orderR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L orderL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var OrderColumns = struct {
	ID          string
	CreatedAt   string
	CompletedAt string
}{
	ID:          "id",
	CreatedAt:   "created_at",
	CompletedAt: "completed_at",
}

// Generated where

var OrderWhere = struct {
	ID          whereHelpernull_Int64
	CreatedAt   whereHelperint64
	CompletedAt whereHelpernull_Int64
}{
	ID:          whereHelpernull_Int64{field: "\"orders\".\"id\""},
	CreatedAt:   whereHelperint64{field: "\"orders\".\"created_at\""},
	CompletedAt: whereHelpernull_Int64{field: "\"orders\".\"completed_at\""},
}

// OrderRels is where relationship names are stored.
var OrderRels = struct {
	OrderItemOptions string
	OrderItems       string
}{
	OrderItemOptions: "OrderItemOptions",
	OrderItems:       "OrderItems",
}

// orderR is where relationships are stored.
type orderR struct {
	OrderItemOptions OrderItemOptionSlice `boil:"OrderItemOptions" json:"OrderItemOptions" toml:"OrderItemOptions" yaml:"OrderItemOptions"`
	OrderItems       OrderItemSlice       `boil:"OrderItems" json:"OrderItems" toml:"OrderItems" yaml:"OrderItems"`
}

// NewStruct creates a new relationship struct
func (*orderR) NewStruct() *orderR {
	return &orderR{}
}

// orderL is where Load methods for each relationship are stored.
type orderL struct{}

var (
	orderAllColumns            = []string{"id", "created_at", "completed_at"}
	orderColumnsWithoutDefault = []string{}
	orderColumnsWithDefault    = []string{"id", "created_at", "completed_at"}
	orderPrimaryKeyColumns     = []string{"id"}
)

type (
	// OrderSlice is an alias for a slice of pointers to Order.
	// This should generally be used opposed to []Order.
	OrderSlice []*Order

	orderQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	orderType                 = reflect.TypeOf(&Order{})
	orderMapping              = queries.MakeStructMapping(orderType)
	orderPrimaryKeyMapping, _ = queries.BindMapping(orderType, orderMapping, orderPrimaryKeyColumns)
	orderInsertCacheMut       sync.RWMutex
	orderInsertCache          = make(map[string]insertCache)
	orderUpdateCacheMut       sync.RWMutex
	orderUpdateCache          = make(map[string]updateCache)
	orderUpsertCacheMut       sync.RWMutex
	orderUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

// OneG returns a single order record from the query using the global executor.
func (q orderQuery) OneG(ctx context.Context) (*Order, error) {
	return q.One(ctx, boil.GetContextDB())
}

// One returns a single order record from the query.
func (q orderQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Order, error) {
	o := &Order{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for orders")
	}

	return o, nil
}

// AllG returns all Order records from the query using the global executor.
func (q orderQuery) AllG(ctx context.Context) (OrderSlice, error) {
	return q.All(ctx, boil.GetContextDB())
}

// All returns all Order records from the query.
func (q orderQuery) All(ctx context.Context, exec boil.ContextExecutor) (OrderSlice, error) {
	var o []*Order

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Order slice")
	}

	return o, nil
}

// CountG returns the count of all Order records in the query, and panics on error.
func (q orderQuery) CountG(ctx context.Context) (int64, error) {
	return q.Count(ctx, boil.GetContextDB())
}

// Count returns the count of all Order records in the query.
func (q orderQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count orders rows")
	}

	return count, nil
}

// ExistsG checks if the row exists in the table, and panics on error.
func (q orderQuery) ExistsG(ctx context.Context) (bool, error) {
	return q.Exists(ctx, boil.GetContextDB())
}

// Exists checks if the row exists in the table.
func (q orderQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if orders exists")
	}

	return count > 0, nil
}

// OrderItemOptions retrieves all the order_item_option's OrderItemOptions with an executor.
func (o *Order) OrderItemOptions(mods ...qm.QueryMod) orderItemOptionQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"order_item_options\".\"order_id\"=?", o.ID),
	)

	query := OrderItemOptions(queryMods...)
	queries.SetFrom(query.Query, "\"order_item_options\"")

	if len(queries.GetSelect(query.Query)) == 0 {
		queries.SetSelect(query.Query, []string{"\"order_item_options\".*"})
	}

	return query
}

// OrderItems retrieves all the order_item's OrderItems with an executor.
func (o *Order) OrderItems(mods ...qm.QueryMod) orderItemQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"order_items\".\"order_id\"=?", o.ID),
	)

	query := OrderItems(queryMods...)
	queries.SetFrom(query.Query, "\"order_items\"")

	if len(queries.GetSelect(query.Query)) == 0 {
		queries.SetSelect(query.Query, []string{"\"order_items\".*"})
	}

	return query
}

// LoadOrderItemOptions allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (orderL) LoadOrderItemOptions(ctx context.Context, e boil.ContextExecutor, singular bool, maybeOrder interface{}, mods queries.Applicator) error {
	var slice []*Order
	var object *Order

	if singular {
		object = maybeOrder.(*Order)
	} else {
		slice = *maybeOrder.(*[]*Order)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &orderR{}
		}
		args = append(args, object.ID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &orderR{}
			}

			for _, a := range args {
				if queries.Equal(a, obj.ID) {
					continue Outer
				}
			}

			args = append(args, obj.ID)
		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`order_item_options`),
		qm.WhereIn(`order_item_options.order_id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load order_item_options")
	}

	var resultSlice []*OrderItemOption
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice order_item_options")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on order_item_options")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for order_item_options")
	}

	if singular {
		object.R.OrderItemOptions = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &orderItemOptionR{}
			}
			foreign.R.Order = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if queries.Equal(local.ID, foreign.OrderID) {
				local.R.OrderItemOptions = append(local.R.OrderItemOptions, foreign)
				if foreign.R == nil {
					foreign.R = &orderItemOptionR{}
				}
				foreign.R.Order = local
				break
			}
		}
	}

	return nil
}

// LoadOrderItems allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (orderL) LoadOrderItems(ctx context.Context, e boil.ContextExecutor, singular bool, maybeOrder interface{}, mods queries.Applicator) error {
	var slice []*Order
	var object *Order

	if singular {
		object = maybeOrder.(*Order)
	} else {
		slice = *maybeOrder.(*[]*Order)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &orderR{}
		}
		args = append(args, object.ID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &orderR{}
			}

			for _, a := range args {
				if queries.Equal(a, obj.ID) {
					continue Outer
				}
			}

			args = append(args, obj.ID)
		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`order_items`),
		qm.WhereIn(`order_items.order_id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load order_items")
	}

	var resultSlice []*OrderItem
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice order_items")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on order_items")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for order_items")
	}

	if singular {
		object.R.OrderItems = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &orderItemR{}
			}
			foreign.R.Order = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if queries.Equal(local.ID, foreign.OrderID) {
				local.R.OrderItems = append(local.R.OrderItems, foreign)
				if foreign.R == nil {
					foreign.R = &orderItemR{}
				}
				foreign.R.Order = local
				break
			}
		}
	}

	return nil
}

// AddOrderItemOptionsG adds the given related objects to the existing relationships
// of the order, optionally inserting them as new records.
// Appends related to o.R.OrderItemOptions.
// Sets related.R.Order appropriately.
// Uses the global database handle.
func (o *Order) AddOrderItemOptionsG(ctx context.Context, insert bool, related ...*OrderItemOption) error {
	return o.AddOrderItemOptions(ctx, boil.GetContextDB(), insert, related...)
}

// AddOrderItemOptions adds the given related objects to the existing relationships
// of the order, optionally inserting them as new records.
// Appends related to o.R.OrderItemOptions.
// Sets related.R.Order appropriately.
func (o *Order) AddOrderItemOptions(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*OrderItemOption) error {
	var err error
	for _, rel := range related {
		if insert {
			queries.Assign(&rel.OrderID, o.ID)
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE \"order_item_options\" SET %s WHERE %s",
				strmangle.SetParamNames("\"", "\"", 0, []string{"order_id"}),
				strmangle.WhereClause("\"", "\"", 0, orderItemOptionPrimaryKeyColumns),
			)
			values := []interface{}{o.ID, rel.ID}

			if boil.IsDebug(ctx) {
				writer := boil.DebugWriterFrom(ctx)
				fmt.Fprintln(writer, updateQuery)
				fmt.Fprintln(writer, values)
			}
			if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}

			queries.Assign(&rel.OrderID, o.ID)
		}
	}

	if o.R == nil {
		o.R = &orderR{
			OrderItemOptions: related,
		}
	} else {
		o.R.OrderItemOptions = append(o.R.OrderItemOptions, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &orderItemOptionR{
				Order: o,
			}
		} else {
			rel.R.Order = o
		}
	}
	return nil
}

// AddOrderItemsG adds the given related objects to the existing relationships
// of the order, optionally inserting them as new records.
// Appends related to o.R.OrderItems.
// Sets related.R.Order appropriately.
// Uses the global database handle.
func (o *Order) AddOrderItemsG(ctx context.Context, insert bool, related ...*OrderItem) error {
	return o.AddOrderItems(ctx, boil.GetContextDB(), insert, related...)
}

// AddOrderItems adds the given related objects to the existing relationships
// of the order, optionally inserting them as new records.
// Appends related to o.R.OrderItems.
// Sets related.R.Order appropriately.
func (o *Order) AddOrderItems(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*OrderItem) error {
	var err error
	for _, rel := range related {
		if insert {
			queries.Assign(&rel.OrderID, o.ID)
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE \"order_items\" SET %s WHERE %s",
				strmangle.SetParamNames("\"", "\"", 0, []string{"order_id"}),
				strmangle.WhereClause("\"", "\"", 0, orderItemPrimaryKeyColumns),
			)
			values := []interface{}{o.ID, rel.ID}

			if boil.IsDebug(ctx) {
				writer := boil.DebugWriterFrom(ctx)
				fmt.Fprintln(writer, updateQuery)
				fmt.Fprintln(writer, values)
			}
			if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}

			queries.Assign(&rel.OrderID, o.ID)
		}
	}

	if o.R == nil {
		o.R = &orderR{
			OrderItems: related,
		}
	} else {
		o.R.OrderItems = append(o.R.OrderItems, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &orderItemR{
				Order: o,
			}
		} else {
			rel.R.Order = o
		}
	}
	return nil
}

// Orders retrieves all the records using an executor.
func Orders(mods ...qm.QueryMod) orderQuery {
	mods = append(mods, qm.From("\"orders\""))
	return orderQuery{NewQuery(mods...)}
}

// FindOrderG retrieves a single record by ID.
func FindOrderG(ctx context.Context, iD null.Int64, selectCols ...string) (*Order, error) {
	return FindOrder(ctx, boil.GetContextDB(), iD, selectCols...)
}

// FindOrder retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindOrder(ctx context.Context, exec boil.ContextExecutor, iD null.Int64, selectCols ...string) (*Order, error) {
	orderObj := &Order{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"orders\" where \"id\"=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, orderObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from orders")
	}

	return orderObj, nil
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *Order) InsertG(ctx context.Context, columns boil.Columns) error {
	return o.Insert(ctx, boil.GetContextDB(), columns)
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Order) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no orders provided for insertion")
	}

	var err error
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if queries.MustTime(o.CreatedAt).IsZero() {
			queries.SetScanner(&o.CreatedAt, currTime)
		}
	}

	nzDefaults := queries.NonZeroDefaultSet(orderColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	orderInsertCacheMut.RLock()
	cache, cached := orderInsertCache[key]
	orderInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			orderAllColumns,
			orderColumnsWithDefault,
			orderColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(orderType, orderMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(orderType, orderMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"orders\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"orders\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT \"%s\" FROM \"orders\" WHERE %s", strings.Join(returnColumns, "\",\""), strmangle.WhereClause("\"", "\"", 0, orderPrimaryKeyColumns))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}
	_, err = exec.ExecContext(ctx, cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into orders")
	}

	var identifierCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	identifierCols = []interface{}{
		o.ID,
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, identifierCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, identifierCols...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for orders")
	}

CacheNoHooks:
	if !cached {
		orderInsertCacheMut.Lock()
		orderInsertCache[key] = cache
		orderInsertCacheMut.Unlock()
	}

	return nil
}

// UpdateG a single Order record using the global executor.
// See Update for more documentation.
func (o *Order) UpdateG(ctx context.Context, columns boil.Columns) (int64, error) {
	return o.Update(ctx, boil.GetContextDB(), columns)
}

// Update uses an executor to update the Order.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Order) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	key := makeCacheKey(columns, nil)
	orderUpdateCacheMut.RLock()
	cache, cached := orderUpdateCache[key]
	orderUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			orderAllColumns,
			orderPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update orders, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"orders\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 0, wl),
			strmangle.WhereClause("\"", "\"", 0, orderPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(orderType, orderMapping, append(wl, orderPrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, values)
	}
	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update orders row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for orders")
	}

	if !cached {
		orderUpdateCacheMut.Lock()
		orderUpdateCache[key] = cache
		orderUpdateCacheMut.Unlock()
	}

	return rowsAff, nil
}

// UpdateAllG updates all rows with the specified column values.
func (q orderQuery) UpdateAllG(ctx context.Context, cols M) (int64, error) {
	return q.UpdateAll(ctx, boil.GetContextDB(), cols)
}

// UpdateAll updates all rows with the specified column values.
func (q orderQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for orders")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for orders")
	}

	return rowsAff, nil
}

// UpdateAllG updates all rows with the specified column values.
func (o OrderSlice) UpdateAllG(ctx context.Context, cols M) (int64, error) {
	return o.UpdateAll(ctx, boil.GetContextDB(), cols)
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o OrderSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("models: update all requires at least one column argument")
	}

	colNames := make([]string, len(cols))
	args := make([]interface{}, len(cols))

	i := 0
	for name, value := range cols {
		colNames[i] = name
		args[i] = value
		i++
	}

	// Append all of the primary key values for each column
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), orderPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"orders\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, orderPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in order slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all order")
	}
	return rowsAff, nil
}

// DeleteG deletes a single Order record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *Order) DeleteG(ctx context.Context) (int64, error) {
	return o.Delete(ctx, boil.GetContextDB())
}

// Delete deletes a single Order record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Order) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Order provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), orderPrimaryKeyMapping)
	sql := "DELETE FROM \"orders\" WHERE \"id\"=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from orders")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for orders")
	}

	return rowsAff, nil
}

func (q orderQuery) DeleteAllG(ctx context.Context) (int64, error) {
	return q.DeleteAll(ctx, boil.GetContextDB())
}

// DeleteAll deletes all matching rows.
func (q orderQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no orderQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from orders")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for orders")
	}

	return rowsAff, nil
}

// DeleteAllG deletes all rows in the slice.
func (o OrderSlice) DeleteAllG(ctx context.Context) (int64, error) {
	return o.DeleteAll(ctx, boil.GetContextDB())
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o OrderSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), orderPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"orders\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, orderPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from order slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for orders")
	}

	return rowsAff, nil
}

// ReloadG refetches the object from the database using the primary keys.
func (o *Order) ReloadG(ctx context.Context) error {
	if o == nil {
		return errors.New("models: no Order provided for reload")
	}

	return o.Reload(ctx, boil.GetContextDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Order) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindOrder(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *OrderSlice) ReloadAllG(ctx context.Context) error {
	if o == nil {
		return errors.New("models: empty OrderSlice provided for reload all")
	}

	return o.ReloadAll(ctx, boil.GetContextDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *OrderSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := OrderSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), orderPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"orders\".* FROM \"orders\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, orderPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in OrderSlice")
	}

	*o = slice

	return nil
}

// OrderExistsG checks if the Order row exists.
func OrderExistsG(ctx context.Context, iD null.Int64) (bool, error) {
	return OrderExists(ctx, boil.GetContextDB(), iD)
}

// OrderExists checks if the Order row exists.
func OrderExists(ctx context.Context, exec boil.ContextExecutor, iD null.Int64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"orders\" where \"id\"=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if orders exists")
	}

	return exists, nil
}
