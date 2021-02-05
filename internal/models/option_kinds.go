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

// OptionKind is an object representing the database table.
type OptionKind struct {
	ID         null.Int64 `boil:"id" json:"id,omitempty" toml:"id" yaml:"id,omitempty"`
	ItemKindID int64      `boil:"item_kind_id" json:"itemKindID" toml:"itemKindID" yaml:"itemKindID"`
	Name       string     `boil:"name" json:"name" toml:"name" yaml:"name"`

	R *optionKindR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L optionKindL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var OptionKindColumns = struct {
	ID         string
	ItemKindID string
	Name       string
}{
	ID:         "id",
	ItemKindID: "item_kind_id",
	Name:       "name",
}

// Generated where

var OptionKindWhere = struct {
	ID         whereHelpernull_Int64
	ItemKindID whereHelperint64
	Name       whereHelperstring
}{
	ID:         whereHelpernull_Int64{field: "\"option_kinds\".\"id\""},
	ItemKindID: whereHelperint64{field: "\"option_kinds\".\"item_kind_id\""},
	Name:       whereHelperstring{field: "\"option_kinds\".\"name\""},
}

// OptionKindRels is where relationship names are stored.
var OptionKindRels = struct {
	ItemKind    string
	KindOptions string
}{
	ItemKind:    "ItemKind",
	KindOptions: "KindOptions",
}

// optionKindR is where relationships are stored.
type optionKindR struct {
	ItemKind    *ItemKind   `boil:"ItemKind" json:"ItemKind" toml:"ItemKind" yaml:"ItemKind"`
	KindOptions OptionSlice `boil:"KindOptions" json:"KindOptions" toml:"KindOptions" yaml:"KindOptions"`
}

// NewStruct creates a new relationship struct
func (*optionKindR) NewStruct() *optionKindR {
	return &optionKindR{}
}

// optionKindL is where Load methods for each relationship are stored.
type optionKindL struct{}

var (
	optionKindAllColumns            = []string{"id", "item_kind_id", "name"}
	optionKindColumnsWithoutDefault = []string{}
	optionKindColumnsWithDefault    = []string{"id", "item_kind_id", "name"}
	optionKindPrimaryKeyColumns     = []string{"id"}
)

type (
	// OptionKindSlice is an alias for a slice of pointers to OptionKind.
	// This should generally be used opposed to []OptionKind.
	OptionKindSlice []*OptionKind

	optionKindQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	optionKindType                 = reflect.TypeOf(&OptionKind{})
	optionKindMapping              = queries.MakeStructMapping(optionKindType)
	optionKindPrimaryKeyMapping, _ = queries.BindMapping(optionKindType, optionKindMapping, optionKindPrimaryKeyColumns)
	optionKindInsertCacheMut       sync.RWMutex
	optionKindInsertCache          = make(map[string]insertCache)
	optionKindUpdateCacheMut       sync.RWMutex
	optionKindUpdateCache          = make(map[string]updateCache)
	optionKindUpsertCacheMut       sync.RWMutex
	optionKindUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

// OneG returns a single optionKind record from the query using the global executor.
func (q optionKindQuery) OneG(ctx context.Context) (*OptionKind, error) {
	return q.One(ctx, boil.GetContextDB())
}

// One returns a single optionKind record from the query.
func (q optionKindQuery) One(ctx context.Context, exec boil.ContextExecutor) (*OptionKind, error) {
	o := &OptionKind{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for option_kinds")
	}

	return o, nil
}

// AllG returns all OptionKind records from the query using the global executor.
func (q optionKindQuery) AllG(ctx context.Context) (OptionKindSlice, error) {
	return q.All(ctx, boil.GetContextDB())
}

// All returns all OptionKind records from the query.
func (q optionKindQuery) All(ctx context.Context, exec boil.ContextExecutor) (OptionKindSlice, error) {
	var o []*OptionKind

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to OptionKind slice")
	}

	return o, nil
}

// CountG returns the count of all OptionKind records in the query, and panics on error.
func (q optionKindQuery) CountG(ctx context.Context) (int64, error) {
	return q.Count(ctx, boil.GetContextDB())
}

// Count returns the count of all OptionKind records in the query.
func (q optionKindQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count option_kinds rows")
	}

	return count, nil
}

// ExistsG checks if the row exists in the table, and panics on error.
func (q optionKindQuery) ExistsG(ctx context.Context) (bool, error) {
	return q.Exists(ctx, boil.GetContextDB())
}

// Exists checks if the row exists in the table.
func (q optionKindQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if option_kinds exists")
	}

	return count > 0, nil
}

// ItemKind pointed to by the foreign key.
func (o *OptionKind) ItemKind(mods ...qm.QueryMod) itemKindQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.ItemKindID),
	}

	queryMods = append(queryMods, mods...)

	query := ItemKinds(queryMods...)
	queries.SetFrom(query.Query, "\"item_kinds\"")

	return query
}

// KindOptions retrieves all the option's Options with an executor via kind_id column.
func (o *OptionKind) KindOptions(mods ...qm.QueryMod) optionQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"options\".\"kind_id\"=?", o.ID),
	)

	query := Options(queryMods...)
	queries.SetFrom(query.Query, "\"options\"")

	if len(queries.GetSelect(query.Query)) == 0 {
		queries.SetSelect(query.Query, []string{"\"options\".*"})
	}

	return query
}

// LoadItemKind allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (optionKindL) LoadItemKind(ctx context.Context, e boil.ContextExecutor, singular bool, maybeOptionKind interface{}, mods queries.Applicator) error {
	var slice []*OptionKind
	var object *OptionKind

	if singular {
		object = maybeOptionKind.(*OptionKind)
	} else {
		slice = *maybeOptionKind.(*[]*OptionKind)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &optionKindR{}
		}
		if !queries.IsNil(object.ItemKindID) {
			args = append(args, object.ItemKindID)
		}

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &optionKindR{}
			}

			for _, a := range args {
				if queries.Equal(a, obj.ItemKindID) {
					continue Outer
				}
			}

			if !queries.IsNil(obj.ItemKindID) {
				args = append(args, obj.ItemKindID)
			}

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`item_kinds`),
		qm.WhereIn(`item_kinds.id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load ItemKind")
	}

	var resultSlice []*ItemKind
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice ItemKind")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for item_kinds")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for item_kinds")
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.ItemKind = foreign
		if foreign.R == nil {
			foreign.R = &itemKindR{}
		}
		foreign.R.OptionKinds = append(foreign.R.OptionKinds, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if queries.Equal(local.ItemKindID, foreign.ID) {
				local.R.ItemKind = foreign
				if foreign.R == nil {
					foreign.R = &itemKindR{}
				}
				foreign.R.OptionKinds = append(foreign.R.OptionKinds, local)
				break
			}
		}
	}

	return nil
}

// LoadKindOptions allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (optionKindL) LoadKindOptions(ctx context.Context, e boil.ContextExecutor, singular bool, maybeOptionKind interface{}, mods queries.Applicator) error {
	var slice []*OptionKind
	var object *OptionKind

	if singular {
		object = maybeOptionKind.(*OptionKind)
	} else {
		slice = *maybeOptionKind.(*[]*OptionKind)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &optionKindR{}
		}
		args = append(args, object.ID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &optionKindR{}
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
		qm.From(`options`),
		qm.WhereIn(`options.kind_id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load options")
	}

	var resultSlice []*Option
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice options")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on options")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for options")
	}

	if singular {
		object.R.KindOptions = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &optionR{}
			}
			foreign.R.Kind = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if queries.Equal(local.ID, foreign.KindID) {
				local.R.KindOptions = append(local.R.KindOptions, foreign)
				if foreign.R == nil {
					foreign.R = &optionR{}
				}
				foreign.R.Kind = local
				break
			}
		}
	}

	return nil
}

// SetItemKindG of the optionKind to the related item.
// Sets o.R.ItemKind to related.
// Adds o to related.R.OptionKinds.
// Uses the global database handle.
func (o *OptionKind) SetItemKindG(ctx context.Context, insert bool, related *ItemKind) error {
	return o.SetItemKind(ctx, boil.GetContextDB(), insert, related)
}

// SetItemKind of the optionKind to the related item.
// Sets o.R.ItemKind to related.
// Adds o to related.R.OptionKinds.
func (o *OptionKind) SetItemKind(ctx context.Context, exec boil.ContextExecutor, insert bool, related *ItemKind) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"option_kinds\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 0, []string{"item_kind_id"}),
		strmangle.WhereClause("\"", "\"", 0, optionKindPrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.ID}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	queries.Assign(&o.ItemKindID, related.ID)
	if o.R == nil {
		o.R = &optionKindR{
			ItemKind: related,
		}
	} else {
		o.R.ItemKind = related
	}

	if related.R == nil {
		related.R = &itemKindR{
			OptionKinds: OptionKindSlice{o},
		}
	} else {
		related.R.OptionKinds = append(related.R.OptionKinds, o)
	}

	return nil
}

// AddKindOptionsG adds the given related objects to the existing relationships
// of the option_kind, optionally inserting them as new records.
// Appends related to o.R.KindOptions.
// Sets related.R.Kind appropriately.
// Uses the global database handle.
func (o *OptionKind) AddKindOptionsG(ctx context.Context, insert bool, related ...*Option) error {
	return o.AddKindOptions(ctx, boil.GetContextDB(), insert, related...)
}

// AddKindOptions adds the given related objects to the existing relationships
// of the option_kind, optionally inserting them as new records.
// Appends related to o.R.KindOptions.
// Sets related.R.Kind appropriately.
func (o *OptionKind) AddKindOptions(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*Option) error {
	var err error
	for _, rel := range related {
		if insert {
			queries.Assign(&rel.KindID, o.ID)
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE \"options\" SET %s WHERE %s",
				strmangle.SetParamNames("\"", "\"", 0, []string{"kind_id"}),
				strmangle.WhereClause("\"", "\"", 0, optionPrimaryKeyColumns),
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

			queries.Assign(&rel.KindID, o.ID)
		}
	}

	if o.R == nil {
		o.R = &optionKindR{
			KindOptions: related,
		}
	} else {
		o.R.KindOptions = append(o.R.KindOptions, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &optionR{
				Kind: o,
			}
		} else {
			rel.R.Kind = o
		}
	}
	return nil
}

// OptionKinds retrieves all the records using an executor.
func OptionKinds(mods ...qm.QueryMod) optionKindQuery {
	mods = append(mods, qm.From("\"option_kinds\""))
	return optionKindQuery{NewQuery(mods...)}
}

// FindOptionKindG retrieves a single record by ID.
func FindOptionKindG(ctx context.Context, iD null.Int64, selectCols ...string) (*OptionKind, error) {
	return FindOptionKind(ctx, boil.GetContextDB(), iD, selectCols...)
}

// FindOptionKind retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindOptionKind(ctx context.Context, exec boil.ContextExecutor, iD null.Int64, selectCols ...string) (*OptionKind, error) {
	optionKindObj := &OptionKind{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"option_kinds\" where \"id\"=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, optionKindObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from option_kinds")
	}

	return optionKindObj, nil
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *OptionKind) InsertG(ctx context.Context, columns boil.Columns) error {
	return o.Insert(ctx, boil.GetContextDB(), columns)
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *OptionKind) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no option_kinds provided for insertion")
	}

	var err error

	nzDefaults := queries.NonZeroDefaultSet(optionKindColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	optionKindInsertCacheMut.RLock()
	cache, cached := optionKindInsertCache[key]
	optionKindInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			optionKindAllColumns,
			optionKindColumnsWithDefault,
			optionKindColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(optionKindType, optionKindMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(optionKindType, optionKindMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"option_kinds\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"option_kinds\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT \"%s\" FROM \"option_kinds\" WHERE %s", strings.Join(returnColumns, "\",\""), strmangle.WhereClause("\"", "\"", 0, optionKindPrimaryKeyColumns))
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
		return errors.Wrap(err, "models: unable to insert into option_kinds")
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
		return errors.Wrap(err, "models: unable to populate default values for option_kinds")
	}

CacheNoHooks:
	if !cached {
		optionKindInsertCacheMut.Lock()
		optionKindInsertCache[key] = cache
		optionKindInsertCacheMut.Unlock()
	}

	return nil
}

// UpdateG a single OptionKind record using the global executor.
// See Update for more documentation.
func (o *OptionKind) UpdateG(ctx context.Context, columns boil.Columns) (int64, error) {
	return o.Update(ctx, boil.GetContextDB(), columns)
}

// Update uses an executor to update the OptionKind.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *OptionKind) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	key := makeCacheKey(columns, nil)
	optionKindUpdateCacheMut.RLock()
	cache, cached := optionKindUpdateCache[key]
	optionKindUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			optionKindAllColumns,
			optionKindPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update option_kinds, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"option_kinds\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 0, wl),
			strmangle.WhereClause("\"", "\"", 0, optionKindPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(optionKindType, optionKindMapping, append(wl, optionKindPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update option_kinds row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for option_kinds")
	}

	if !cached {
		optionKindUpdateCacheMut.Lock()
		optionKindUpdateCache[key] = cache
		optionKindUpdateCacheMut.Unlock()
	}

	return rowsAff, nil
}

// UpdateAllG updates all rows with the specified column values.
func (q optionKindQuery) UpdateAllG(ctx context.Context, cols M) (int64, error) {
	return q.UpdateAll(ctx, boil.GetContextDB(), cols)
}

// UpdateAll updates all rows with the specified column values.
func (q optionKindQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for option_kinds")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for option_kinds")
	}

	return rowsAff, nil
}

// UpdateAllG updates all rows with the specified column values.
func (o OptionKindSlice) UpdateAllG(ctx context.Context, cols M) (int64, error) {
	return o.UpdateAll(ctx, boil.GetContextDB(), cols)
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o OptionKindSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), optionKindPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"option_kinds\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, optionKindPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in optionKind slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all optionKind")
	}
	return rowsAff, nil
}

// DeleteG deletes a single OptionKind record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *OptionKind) DeleteG(ctx context.Context) (int64, error) {
	return o.Delete(ctx, boil.GetContextDB())
}

// Delete deletes a single OptionKind record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *OptionKind) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no OptionKind provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), optionKindPrimaryKeyMapping)
	sql := "DELETE FROM \"option_kinds\" WHERE \"id\"=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from option_kinds")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for option_kinds")
	}

	return rowsAff, nil
}

func (q optionKindQuery) DeleteAllG(ctx context.Context) (int64, error) {
	return q.DeleteAll(ctx, boil.GetContextDB())
}

// DeleteAll deletes all matching rows.
func (q optionKindQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no optionKindQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from option_kinds")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for option_kinds")
	}

	return rowsAff, nil
}

// DeleteAllG deletes all rows in the slice.
func (o OptionKindSlice) DeleteAllG(ctx context.Context) (int64, error) {
	return o.DeleteAll(ctx, boil.GetContextDB())
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o OptionKindSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), optionKindPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"option_kinds\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, optionKindPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from optionKind slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for option_kinds")
	}

	return rowsAff, nil
}

// ReloadG refetches the object from the database using the primary keys.
func (o *OptionKind) ReloadG(ctx context.Context) error {
	if o == nil {
		return errors.New("models: no OptionKind provided for reload")
	}

	return o.Reload(ctx, boil.GetContextDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *OptionKind) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindOptionKind(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *OptionKindSlice) ReloadAllG(ctx context.Context) error {
	if o == nil {
		return errors.New("models: empty OptionKindSlice provided for reload all")
	}

	return o.ReloadAll(ctx, boil.GetContextDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *OptionKindSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := OptionKindSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), optionKindPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"option_kinds\".* FROM \"option_kinds\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, optionKindPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in OptionKindSlice")
	}

	*o = slice

	return nil
}

// OptionKindExistsG checks if the OptionKind row exists.
func OptionKindExistsG(ctx context.Context, iD null.Int64) (bool, error) {
	return OptionKindExists(ctx, boil.GetContextDB(), iD)
}

// OptionKindExists checks if the OptionKind row exists.
func OptionKindExists(ctx context.Context, exec boil.ContextExecutor, iD null.Int64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"option_kinds\" where \"id\"=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if option_kinds exists")
	}

	return exists, nil
}
