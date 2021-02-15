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
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// Period is an object representing the database table.
type Period struct {
	SampleRate int64   `boil:"sample_rate" json:"sample_rate" toml:"sample_rate" yaml:"sample_rate"`
	BufferSize int64   `boil:"buffer_size" json:"buffer_size" toml:"buffer_size" yaml:"buffer_size"`
	Frequency  int64   `boil:"frequency" json:"frequency" toml:"frequency" yaml:"frequency"`
	Period     float64 `boil:"period" json:"period" toml:"period" yaml:"period"`

	R *periodR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L periodL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var PeriodColumns = struct {
	SampleRate string
	BufferSize string
	Frequency  string
	Period     string
}{
	SampleRate: "sample_rate",
	BufferSize: "buffer_size",
	Frequency:  "frequency",
	Period:     "period",
}

// Generated where

type whereHelperint64 struct{ field string }

func (w whereHelperint64) EQ(x int64) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperint64) NEQ(x int64) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperint64) LT(x int64) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperint64) LTE(x int64) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperint64) GT(x int64) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperint64) GTE(x int64) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }
func (w whereHelperint64) IN(slice []int64) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelperint64) NIN(slice []int64) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereNotIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

type whereHelperfloat64 struct{ field string }

func (w whereHelperfloat64) EQ(x float64) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperfloat64) NEQ(x float64) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.NEQ, x)
}
func (w whereHelperfloat64) LT(x float64) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperfloat64) LTE(x float64) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelperfloat64) GT(x float64) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperfloat64) GTE(x float64) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}
func (w whereHelperfloat64) IN(slice []float64) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelperfloat64) NIN(slice []float64) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereNotIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

var PeriodWhere = struct {
	SampleRate whereHelperint64
	BufferSize whereHelperint64
	Frequency  whereHelperint64
	Period     whereHelperfloat64
}{
	SampleRate: whereHelperint64{field: "\"period\".\"sample_rate\""},
	BufferSize: whereHelperint64{field: "\"period\".\"buffer_size\""},
	Frequency:  whereHelperint64{field: "\"period\".\"frequency\""},
	Period:     whereHelperfloat64{field: "\"period\".\"period\""},
}

// PeriodRels is where relationship names are stored.
var PeriodRels = struct {
}{}

// periodR is where relationships are stored.
type periodR struct {
}

// NewStruct creates a new relationship struct
func (*periodR) NewStruct() *periodR {
	return &periodR{}
}

// periodL is where Load methods for each relationship are stored.
type periodL struct{}

var (
	periodAllColumns            = []string{"sample_rate", "buffer_size", "frequency", "period"}
	periodColumnsWithoutDefault = []string{"buffer_size", "frequency", "period"}
	periodColumnsWithDefault    = []string{"sample_rate"}
	periodPrimaryKeyColumns     = []string{"sample_rate", "buffer_size", "frequency"}
)

type (
	// PeriodSlice is an alias for a slice of pointers to Period.
	// This should generally be used opposed to []Period.
	PeriodSlice []*Period
	// PeriodHook is the signature for custom Period hook methods
	PeriodHook func(context.Context, boil.ContextExecutor, *Period) error

	periodQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	periodType                 = reflect.TypeOf(&Period{})
	periodMapping              = queries.MakeStructMapping(periodType)
	periodPrimaryKeyMapping, _ = queries.BindMapping(periodType, periodMapping, periodPrimaryKeyColumns)
	periodInsertCacheMut       sync.RWMutex
	periodInsertCache          = make(map[string]insertCache)
	periodUpdateCacheMut       sync.RWMutex
	periodUpdateCache          = make(map[string]updateCache)
	periodUpsertCacheMut       sync.RWMutex
	periodUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var periodBeforeInsertHooks []PeriodHook
var periodBeforeUpdateHooks []PeriodHook
var periodBeforeDeleteHooks []PeriodHook
var periodBeforeUpsertHooks []PeriodHook

var periodAfterInsertHooks []PeriodHook
var periodAfterSelectHooks []PeriodHook
var periodAfterUpdateHooks []PeriodHook
var periodAfterDeleteHooks []PeriodHook
var periodAfterUpsertHooks []PeriodHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Period) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range periodBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Period) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range periodBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Period) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range periodBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Period) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range periodBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Period) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range periodAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Period) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range periodAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Period) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range periodAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Period) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range periodAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Period) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range periodAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddPeriodHook registers your hook function for all future operations.
func AddPeriodHook(hookPoint boil.HookPoint, periodHook PeriodHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		periodBeforeInsertHooks = append(periodBeforeInsertHooks, periodHook)
	case boil.BeforeUpdateHook:
		periodBeforeUpdateHooks = append(periodBeforeUpdateHooks, periodHook)
	case boil.BeforeDeleteHook:
		periodBeforeDeleteHooks = append(periodBeforeDeleteHooks, periodHook)
	case boil.BeforeUpsertHook:
		periodBeforeUpsertHooks = append(periodBeforeUpsertHooks, periodHook)
	case boil.AfterInsertHook:
		periodAfterInsertHooks = append(periodAfterInsertHooks, periodHook)
	case boil.AfterSelectHook:
		periodAfterSelectHooks = append(periodAfterSelectHooks, periodHook)
	case boil.AfterUpdateHook:
		periodAfterUpdateHooks = append(periodAfterUpdateHooks, periodHook)
	case boil.AfterDeleteHook:
		periodAfterDeleteHooks = append(periodAfterDeleteHooks, periodHook)
	case boil.AfterUpsertHook:
		periodAfterUpsertHooks = append(periodAfterUpsertHooks, periodHook)
	}
}

// One returns a single period record from the query.
func (q periodQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Period, error) {
	o := &Period{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for period")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all Period records from the query.
func (q periodQuery) All(ctx context.Context, exec boil.ContextExecutor) (PeriodSlice, error) {
	var o []*Period

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Period slice")
	}

	if len(periodAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all Period records in the query.
func (q periodQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count period rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q periodQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if period exists")
	}

	return count > 0, nil
}

// Periods retrieves all the records using an executor.
func Periods(mods ...qm.QueryMod) periodQuery {
	mods = append(mods, qm.From("\"period\""))
	return periodQuery{NewQuery(mods...)}
}

// FindPeriod retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindPeriod(ctx context.Context, exec boil.ContextExecutor, sampleRate int64, bufferSize int64, frequency int64, selectCols ...string) (*Period, error) {
	periodObj := &Period{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"period\" where \"sample_rate\"=? AND \"buffer_size\"=? AND \"frequency\"=?", sel,
	)

	q := queries.Raw(query, sampleRate, bufferSize, frequency)

	err := q.Bind(ctx, exec, periodObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from period")
	}

	return periodObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Period) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no period provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(periodColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	periodInsertCacheMut.RLock()
	cache, cached := periodInsertCache[key]
	periodInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			periodAllColumns,
			periodColumnsWithDefault,
			periodColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(periodType, periodMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(periodType, periodMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"period\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"period\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT \"%s\" FROM \"period\" WHERE %s", strings.Join(returnColumns, "\",\""), strmangle.WhereClause("\"", "\"", 0, periodPrimaryKeyColumns))
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
		return errors.Wrap(err, "models: unable to insert into period")
	}

	var identifierCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	identifierCols = []interface{}{
		o.SampleRate,
		o.BufferSize,
		o.Frequency,
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, identifierCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, identifierCols...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for period")
	}

CacheNoHooks:
	if !cached {
		periodInsertCacheMut.Lock()
		periodInsertCache[key] = cache
		periodInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the Period.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Period) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	periodUpdateCacheMut.RLock()
	cache, cached := periodUpdateCache[key]
	periodUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			periodAllColumns,
			periodPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update period, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"period\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 0, wl),
			strmangle.WhereClause("\"", "\"", 0, periodPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(periodType, periodMapping, append(wl, periodPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update period row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for period")
	}

	if !cached {
		periodUpdateCacheMut.Lock()
		periodUpdateCache[key] = cache
		periodUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q periodQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for period")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for period")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o PeriodSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), periodPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"period\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, periodPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in period slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all period")
	}
	return rowsAff, nil
}

// Delete deletes a single Period record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Period) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Period provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), periodPrimaryKeyMapping)
	sql := "DELETE FROM \"period\" WHERE \"sample_rate\"=? AND \"buffer_size\"=? AND \"frequency\"=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from period")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for period")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q periodQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no periodQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from period")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for period")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o PeriodSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(periodBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), periodPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"period\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, periodPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from period slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for period")
	}

	if len(periodAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Period) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindPeriod(ctx, exec, o.SampleRate, o.BufferSize, o.Frequency)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *PeriodSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := PeriodSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), periodPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"period\".* FROM \"period\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, periodPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in PeriodSlice")
	}

	*o = slice

	return nil
}

// PeriodExists checks if the Period row exists.
func PeriodExists(ctx context.Context, exec boil.ContextExecutor, sampleRate int64, bufferSize int64, frequency int64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"period\" where \"sample_rate\"=? AND \"buffer_size\"=? AND \"frequency\"=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, sampleRate, bufferSize, frequency)
	}
	row := exec.QueryRowContext(ctx, sql, sampleRate, bufferSize, frequency)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if period exists")
	}

	return exists, nil
}