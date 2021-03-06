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

// ItemSide is an object representing the database table.
type ItemSide struct {
	ItemID     int64      `boil:"item_id" json:"itemID" toml:"itemID" yaml:"itemID"`
	SideItemID int64      `boil:"side_item_id" json:"sideItemID" toml:"sideItemID" yaml:"sideItemID"`
	IsDefault  null.Int64 `boil:"is_default" json:"isDefault,omitempty" toml:"isDefault" yaml:"isDefault,omitempty"`
	Price      null.Int64 `boil:"price" json:"price,omitempty" toml:"price" yaml:"price,omitempty"`

	R *itemSideR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L itemSideL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var ItemSideColumns = struct {
	ItemID     string
	SideItemID string
	IsDefault  string
	Price      string
}{
	ItemID:     "item_id",
	SideItemID: "side_item_id",
	IsDefault:  "is_default",
	Price:      "price",
}

// Generated where

var ItemSideWhere = struct {
	ItemID     whereHelperint64
	SideItemID whereHelperint64
	IsDefault  whereHelpernull_Int64
	Price      whereHelpernull_Int64
}{
	ItemID:     whereHelperint64{field: "\"item_sides\".\"item_id\""},
	SideItemID: whereHelperint64{field: "\"item_sides\".\"side_item_id\""},
	IsDefault:  whereHelpernull_Int64{field: "\"item_sides\".\"is_default\""},
	Price:      whereHelpernull_Int64{field: "\"item_sides\".\"price\""},
}

// ItemSideRels is where relationship names are stored.
var ItemSideRels = struct {
	SideItem string
	Item     string
}{
	SideItem: "SideItem",
	Item:     "Item",
}

// itemSideR is where relationships are stored.
type itemSideR struct {
	SideItem *Item `boil:"SideItem" json:"SideItem" toml:"SideItem" yaml:"SideItem"`
	Item     *Item `boil:"Item" json:"Item" toml:"Item" yaml:"Item"`
}

// NewStruct creates a new relationship struct
func (*itemSideR) NewStruct() *itemSideR {
	return &itemSideR{}
}

// itemSideL is where Load methods for each relationship are stored.
type itemSideL struct{}

var (
	itemSideAllColumns            = []string{"item_id", "side_item_id", "is_default", "price"}
	itemSideColumnsWithoutDefault = []string{"side_item_id", "price"}
	itemSideColumnsWithDefault    = []string{"item_id", "is_default"}
	itemSidePrimaryKeyColumns     = []string{"item_id", "side_item_id"}
)

type (
	// ItemSideSlice is an alias for a slice of pointers to ItemSide.
	// This should generally be used opposed to []ItemSide.
	ItemSideSlice []*ItemSide

	itemSideQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	itemSideType                 = reflect.TypeOf(&ItemSide{})
	itemSideMapping              = queries.MakeStructMapping(itemSideType)
	itemSidePrimaryKeyMapping, _ = queries.BindMapping(itemSideType, itemSideMapping, itemSidePrimaryKeyColumns)
	itemSideInsertCacheMut       sync.RWMutex
	itemSideInsertCache          = make(map[string]insertCache)
	itemSideUpdateCacheMut       sync.RWMutex
	itemSideUpdateCache          = make(map[string]updateCache)
	itemSideUpsertCacheMut       sync.RWMutex
	itemSideUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

// OneG returns a single itemSide record from the query using the global executor.
func (q itemSideQuery) OneG(ctx context.Context) (*ItemSide, error) {
	return q.One(ctx, boil.GetContextDB())
}

// One returns a single itemSide record from the query.
func (q itemSideQuery) One(ctx context.Context, exec boil.ContextExecutor) (*ItemSide, error) {
	o := &ItemSide{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for item_sides")
	}

	return o, nil
}

// AllG returns all ItemSide records from the query using the global executor.
func (q itemSideQuery) AllG(ctx context.Context) (ItemSideSlice, error) {
	return q.All(ctx, boil.GetContextDB())
}

// All returns all ItemSide records from the query.
func (q itemSideQuery) All(ctx context.Context, exec boil.ContextExecutor) (ItemSideSlice, error) {
	var o []*ItemSide

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to ItemSide slice")
	}

	return o, nil
}

// CountG returns the count of all ItemSide records in the query, and panics on error.
func (q itemSideQuery) CountG(ctx context.Context) (int64, error) {
	return q.Count(ctx, boil.GetContextDB())
}

// Count returns the count of all ItemSide records in the query.
func (q itemSideQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count item_sides rows")
	}

	return count, nil
}

// ExistsG checks if the row exists in the table, and panics on error.
func (q itemSideQuery) ExistsG(ctx context.Context) (bool, error) {
	return q.Exists(ctx, boil.GetContextDB())
}

// Exists checks if the row exists in the table.
func (q itemSideQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if item_sides exists")
	}

	return count > 0, nil
}

// SideItem pointed to by the foreign key.
func (o *ItemSide) SideItem(mods ...qm.QueryMod) itemQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.SideItemID),
	}

	queryMods = append(queryMods, mods...)

	query := Items(queryMods...)
	queries.SetFrom(query.Query, "\"items\"")

	return query
}

// Item pointed to by the foreign key.
func (o *ItemSide) Item(mods ...qm.QueryMod) itemQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.ItemID),
	}

	queryMods = append(queryMods, mods...)

	query := Items(queryMods...)
	queries.SetFrom(query.Query, "\"items\"")

	return query
}

// LoadSideItem allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (itemSideL) LoadSideItem(ctx context.Context, e boil.ContextExecutor, singular bool, maybeItemSide interface{}, mods queries.Applicator) error {
	var slice []*ItemSide
	var object *ItemSide

	if singular {
		object = maybeItemSide.(*ItemSide)
	} else {
		slice = *maybeItemSide.(*[]*ItemSide)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &itemSideR{}
		}
		if !queries.IsNil(object.SideItemID) {
			args = append(args, object.SideItemID)
		}

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &itemSideR{}
			}

			for _, a := range args {
				if queries.Equal(a, obj.SideItemID) {
					continue Outer
				}
			}

			if !queries.IsNil(obj.SideItemID) {
				args = append(args, obj.SideItemID)
			}

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`items`),
		qm.WhereIn(`items.id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Item")
	}

	var resultSlice []*Item
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Item")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for items")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for items")
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.SideItem = foreign
		if foreign.R == nil {
			foreign.R = &itemR{}
		}
		foreign.R.SideItemItemSide = object
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if queries.Equal(local.SideItemID, foreign.ID) {
				local.R.SideItem = foreign
				if foreign.R == nil {
					foreign.R = &itemR{}
				}
				foreign.R.SideItemItemSide = local
				break
			}
		}
	}

	return nil
}

// LoadItem allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (itemSideL) LoadItem(ctx context.Context, e boil.ContextExecutor, singular bool, maybeItemSide interface{}, mods queries.Applicator) error {
	var slice []*ItemSide
	var object *ItemSide

	if singular {
		object = maybeItemSide.(*ItemSide)
	} else {
		slice = *maybeItemSide.(*[]*ItemSide)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &itemSideR{}
		}
		if !queries.IsNil(object.ItemID) {
			args = append(args, object.ItemID)
		}

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &itemSideR{}
			}

			for _, a := range args {
				if queries.Equal(a, obj.ItemID) {
					continue Outer
				}
			}

			if !queries.IsNil(obj.ItemID) {
				args = append(args, obj.ItemID)
			}

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`items`),
		qm.WhereIn(`items.id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Item")
	}

	var resultSlice []*Item
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Item")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for items")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for items")
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.Item = foreign
		if foreign.R == nil {
			foreign.R = &itemR{}
		}
		foreign.R.ItemSide = object
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if queries.Equal(local.ItemID, foreign.ID) {
				local.R.Item = foreign
				if foreign.R == nil {
					foreign.R = &itemR{}
				}
				foreign.R.ItemSide = local
				break
			}
		}
	}

	return nil
}

// SetSideItemG of the itemSide to the related item.
// Sets o.R.SideItem to related.
// Adds o to related.R.SideItemItemSide.
// Uses the global database handle.
func (o *ItemSide) SetSideItemG(ctx context.Context, insert bool, related *Item) error {
	return o.SetSideItem(ctx, boil.GetContextDB(), insert, related)
}

// SetSideItem of the itemSide to the related item.
// Sets o.R.SideItem to related.
// Adds o to related.R.SideItemItemSide.
func (o *ItemSide) SetSideItem(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Item) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"item_sides\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 0, []string{"side_item_id"}),
		strmangle.WhereClause("\"", "\"", 0, itemSidePrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.ItemID, o.SideItemID}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	queries.Assign(&o.SideItemID, related.ID)
	if o.R == nil {
		o.R = &itemSideR{
			SideItem: related,
		}
	} else {
		o.R.SideItem = related
	}

	if related.R == nil {
		related.R = &itemR{
			SideItemItemSide: o,
		}
	} else {
		related.R.SideItemItemSide = o
	}

	return nil
}

// SetItemG of the itemSide to the related item.
// Sets o.R.Item to related.
// Adds o to related.R.ItemSide.
// Uses the global database handle.
func (o *ItemSide) SetItemG(ctx context.Context, insert bool, related *Item) error {
	return o.SetItem(ctx, boil.GetContextDB(), insert, related)
}

// SetItem of the itemSide to the related item.
// Sets o.R.Item to related.
// Adds o to related.R.ItemSide.
func (o *ItemSide) SetItem(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Item) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"item_sides\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 0, []string{"item_id"}),
		strmangle.WhereClause("\"", "\"", 0, itemSidePrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.ItemID, o.SideItemID}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	queries.Assign(&o.ItemID, related.ID)
	if o.R == nil {
		o.R = &itemSideR{
			Item: related,
		}
	} else {
		o.R.Item = related
	}

	if related.R == nil {
		related.R = &itemR{
			ItemSide: o,
		}
	} else {
		related.R.ItemSide = o
	}

	return nil
}

// ItemSides retrieves all the records using an executor.
func ItemSides(mods ...qm.QueryMod) itemSideQuery {
	mods = append(mods, qm.From("\"item_sides\""))
	return itemSideQuery{NewQuery(mods...)}
}

// FindItemSideG retrieves a single record by ID.
func FindItemSideG(ctx context.Context, itemID int64, sideItemID int64, selectCols ...string) (*ItemSide, error) {
	return FindItemSide(ctx, boil.GetContextDB(), itemID, sideItemID, selectCols...)
}

// FindItemSide retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindItemSide(ctx context.Context, exec boil.ContextExecutor, itemID int64, sideItemID int64, selectCols ...string) (*ItemSide, error) {
	itemSideObj := &ItemSide{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"item_sides\" where \"item_id\"=? AND \"side_item_id\"=?", sel,
	)

	q := queries.Raw(query, itemID, sideItemID)

	err := q.Bind(ctx, exec, itemSideObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from item_sides")
	}

	return itemSideObj, nil
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *ItemSide) InsertG(ctx context.Context, columns boil.Columns) error {
	return o.Insert(ctx, boil.GetContextDB(), columns)
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *ItemSide) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no item_sides provided for insertion")
	}

	var err error

	nzDefaults := queries.NonZeroDefaultSet(itemSideColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	itemSideInsertCacheMut.RLock()
	cache, cached := itemSideInsertCache[key]
	itemSideInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			itemSideAllColumns,
			itemSideColumnsWithDefault,
			itemSideColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(itemSideType, itemSideMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(itemSideType, itemSideMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"item_sides\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"item_sides\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT \"%s\" FROM \"item_sides\" WHERE %s", strings.Join(returnColumns, "\",\""), strmangle.WhereClause("\"", "\"", 0, itemSidePrimaryKeyColumns))
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
		return errors.Wrap(err, "models: unable to insert into item_sides")
	}

	var identifierCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	identifierCols = []interface{}{
		o.ItemID,
		o.SideItemID,
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, identifierCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, identifierCols...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for item_sides")
	}

CacheNoHooks:
	if !cached {
		itemSideInsertCacheMut.Lock()
		itemSideInsertCache[key] = cache
		itemSideInsertCacheMut.Unlock()
	}

	return nil
}

// UpdateG a single ItemSide record using the global executor.
// See Update for more documentation.
func (o *ItemSide) UpdateG(ctx context.Context, columns boil.Columns) (int64, error) {
	return o.Update(ctx, boil.GetContextDB(), columns)
}

// Update uses an executor to update the ItemSide.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *ItemSide) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	key := makeCacheKey(columns, nil)
	itemSideUpdateCacheMut.RLock()
	cache, cached := itemSideUpdateCache[key]
	itemSideUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			itemSideAllColumns,
			itemSidePrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update item_sides, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"item_sides\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 0, wl),
			strmangle.WhereClause("\"", "\"", 0, itemSidePrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(itemSideType, itemSideMapping, append(wl, itemSidePrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update item_sides row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for item_sides")
	}

	if !cached {
		itemSideUpdateCacheMut.Lock()
		itemSideUpdateCache[key] = cache
		itemSideUpdateCacheMut.Unlock()
	}

	return rowsAff, nil
}

// UpdateAllG updates all rows with the specified column values.
func (q itemSideQuery) UpdateAllG(ctx context.Context, cols M) (int64, error) {
	return q.UpdateAll(ctx, boil.GetContextDB(), cols)
}

// UpdateAll updates all rows with the specified column values.
func (q itemSideQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for item_sides")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for item_sides")
	}

	return rowsAff, nil
}

// UpdateAllG updates all rows with the specified column values.
func (o ItemSideSlice) UpdateAllG(ctx context.Context, cols M) (int64, error) {
	return o.UpdateAll(ctx, boil.GetContextDB(), cols)
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o ItemSideSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), itemSidePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"item_sides\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, itemSidePrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in itemSide slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all itemSide")
	}
	return rowsAff, nil
}

// DeleteG deletes a single ItemSide record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *ItemSide) DeleteG(ctx context.Context) (int64, error) {
	return o.Delete(ctx, boil.GetContextDB())
}

// Delete deletes a single ItemSide record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *ItemSide) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no ItemSide provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), itemSidePrimaryKeyMapping)
	sql := "DELETE FROM \"item_sides\" WHERE \"item_id\"=? AND \"side_item_id\"=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from item_sides")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for item_sides")
	}

	return rowsAff, nil
}

func (q itemSideQuery) DeleteAllG(ctx context.Context) (int64, error) {
	return q.DeleteAll(ctx, boil.GetContextDB())
}

// DeleteAll deletes all matching rows.
func (q itemSideQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no itemSideQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from item_sides")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for item_sides")
	}

	return rowsAff, nil
}

// DeleteAllG deletes all rows in the slice.
func (o ItemSideSlice) DeleteAllG(ctx context.Context) (int64, error) {
	return o.DeleteAll(ctx, boil.GetContextDB())
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o ItemSideSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), itemSidePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"item_sides\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, itemSidePrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from itemSide slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for item_sides")
	}

	return rowsAff, nil
}

// ReloadG refetches the object from the database using the primary keys.
func (o *ItemSide) ReloadG(ctx context.Context) error {
	if o == nil {
		return errors.New("models: no ItemSide provided for reload")
	}

	return o.Reload(ctx, boil.GetContextDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *ItemSide) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindItemSide(ctx, exec, o.ItemID, o.SideItemID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *ItemSideSlice) ReloadAllG(ctx context.Context) error {
	if o == nil {
		return errors.New("models: empty ItemSideSlice provided for reload all")
	}

	return o.ReloadAll(ctx, boil.GetContextDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *ItemSideSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := ItemSideSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), itemSidePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"item_sides\".* FROM \"item_sides\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, itemSidePrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in ItemSideSlice")
	}

	*o = slice

	return nil
}

// ItemSideExistsG checks if the ItemSide row exists.
func ItemSideExistsG(ctx context.Context, itemID int64, sideItemID int64) (bool, error) {
	return ItemSideExists(ctx, boil.GetContextDB(), itemID, sideItemID)
}

// ItemSideExists checks if the ItemSide row exists.
func ItemSideExists(ctx context.Context, exec boil.ContextExecutor, itemID int64, sideItemID int64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"item_sides\" where \"item_id\"=? AND \"side_item_id\"=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, itemID, sideItemID)
	}
	row := exec.QueryRowContext(ctx, sql, itemID, sideItemID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if item_sides exists")
	}

	return exists, nil
}
