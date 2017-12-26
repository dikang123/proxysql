package proxysql

import (
	"database/sql"
	"fmt"
	"strconv"

	"github.com/juju/errors"
)

type Schedulers struct {
	Id         int64  `json:"id" db:"id"`
	Active     int64  `json:"active" db:"active"`
	IntervalMs int64  `json:"interval_ms" db:"interval_ms"`
	FileName   string `json:"filename" db:"filename"`
	Arg1       string `json:"arg1" db:"arg1"`
	Arg2       string `json:"arg2" db:"arg2"`
	Arg3       string `json:"arg3" db:"arg3"`
	Arg4       string `json:"arg4" db:"arg4"`
	Arg5       string `json:"arg5" db:"arg5"`
	Comment    string `json:"comment" db:"comment"`
}

const (
	/*add a new scheduler*/
	StmtAddOneScheduler = `
	INSERT 
	INTO 
		scheduler(filename,interval_ms,active,arg1,arg2,arg3,arg4,arg5) 
	VALUES(%q,%d,%d,%q,%q,%q,%q,%q)`

	/*delete a scheduler*/
	StmtDeleteOneScheduler = `
	DELETE 
	FROM 
		scheduler 
	WHERE id = %d`

	/*update a scheduler*/
	StmtUpdateOneScheduler = `
	UPDATE 
		scheduler 
	SET 
		active = %d,
		interval_ms=%d,
		filename = %q,
		arg1=%q,
		arg2=%q,
		arg3=%q,
		arg4=%q,
		arg5=%q,
		comment=%q 
	WHERE 
		id = %d`

	/*query all schedulers.*/
	StmtFindAllScheduler = `
	SELECT 
		id,
		active,
		interval_ms,
		filename,
		ifnull(arg1,""),
		ifnull(arg2,""),
		ifnull(arg3,""),
		ifnull(arg4,""),
		ifnull(arg5,""),
		comment 
	FROM 
		scheduler 
	LIMIT %d 
	OFFSET %d`
)

// query all schedulers
func FindAllSchedulerInfo(db *sql.DB, limit int64, skip int64) ([]Schedulers, error) {

	var allscheduler []Schedulers

	Query := fmt.Sprintf(StmtFindAllScheduler, limit, skip)

	rows, err := db.Query(Query)
	if err != nil {
		return []Schedulers{}, errors.Trace(err)
	}
	defer rows.Close()

	for rows.Next() {
		var tmpscheduler Schedulers
		err = rows.Scan(
			&tmpscheduler.Id,
			&tmpscheduler.Active,
			&tmpscheduler.IntervalMs,
			&tmpscheduler.FileName,
			&tmpscheduler.Arg1,
			&tmpscheduler.Arg2,
			&tmpscheduler.Arg3,
			&tmpscheduler.Arg4,
			&tmpscheduler.Arg5,
			&tmpscheduler.Comment,
		)

		if err != nil {
			continue
		}

		allscheduler = append(allscheduler, tmpscheduler)
	}

	return allscheduler, nil
}

//new NewSch
func NewSch(filename string, interval_ms int64) (*Schedulers, error) {

	sched := new(Schedulers)

	sched.FileName = filename
	sched.IntervalMs = interval_ms

	sched.Active = 0
	sched.Arg1 = "NULL"
	sched.Arg2 = "NULL"
	sched.Arg3 = "NULL"
	sched.Arg4 = "NULL"
	sched.Arg5 = "NULL"

	return sched, nil

}

// Set Scheduler Active status
func (sched *Schedulers) SetSchedulerActive(active int64) {
	if active >= 1 {
		sched.Active = 1
	} else {
		sched.Active = 0
	}
}

// Set Scheduler all Args
func (sched *Schedulers) SetSchedulerArg1(arg1 string) {
	sched.Arg1 = arg1
}

func (sched *Schedulers) SetSchedulerArg2(arg2 string) {
	sched.Arg2 = arg2
}

func (sched *Schedulers) SetSchedulerArg3(arg3 string) {
	sched.Arg3 = arg3
}

func (sched *Schedulers) SetSchedulerArg4(arg4 string) {
	sched.Arg4 = arg4
}

func (sched *Schedulers) SetSchedulerArg5(arg5 string) {
	sched.Arg5 = arg5
}

// Set scheduler interval_ms
func (sched *Schedulers) SetSchedulerIntervalMs(interval_ms int64) {
	switch {
	case interval_ms < 100:
		sched.IntervalMs = 100
	case interval_ms > 100000000:
		sched.IntervalMs = 100000000
	default:
		sched.IntervalMs = interval_ms
	}
}

//add a new scheduler
func (schld *Schedulers) AddOneScheduler(db *sql.DB) error {

	Query := fmt.Sprintf(StmtAddOneScheduler, schld.FileName, schld.IntervalMs, schld.Active, schld.Arg1, schld.Arg2, schld.Arg3, schld.Arg4, schld.Arg5)

	_, err := db.Exec(Query)
	if err != nil {
		return errors.Trace(err)
	}

	LoadSchedulerToRuntime(db)
	SaveSchedulerToDisk(db)

	return nil
}

//delete a scheduler
func (schld *Schedulers) DeleteOneScheduler(db *sql.DB) error {

	Query := fmt.Sprintf(StmtDeleteOneScheduler, schld.Id)

	result, err := db.Exec(Query)
	if err != nil {
		return errors.Trace(err)
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return errors.NotFoundf(strconv.Itoa(int(schld.Id)))
	}

	LoadSchedulerToRuntime(db)
	SaveSchedulerToDisk(db)
	return nil
}

//update a scheduler.
func (schld *Schedulers) UpdateOneSchedulerInfo(db *sql.DB) error {

	Query := fmt.Sprintf(StmtUpdateOneScheduler, schld.Active, schld.IntervalMs, schld.FileName, schld.Arg1, schld.Arg2, schld.Arg3, schld.Arg4, schld.Arg5, schld.Comment, schld.Id)

	result, err := db.Exec(Query)
	if err != nil {
		return errors.Trace(err)
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return errors.NotFoundf(strconv.Itoa(int(schld.Id)))
	}

	LoadSchedulerToRuntime(db)
	SaveSchedulerToDisk(db)

	return nil
}
