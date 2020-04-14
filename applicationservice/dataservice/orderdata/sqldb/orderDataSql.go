// Package sql represents SQL database implementation of the user data persistence layer
package sqldb

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jfeng45/order/app/logger"
	"github.com/jfeng45/order/domain/model"
	"github.com/jfeng45/order/tool/gdbc"
	"github.com/jfeng45/order/tool/timea"
	"github.com/pkg/errors"
	"time"
)

const (
	DELETE_ORDER        string = "delete from porder where order_number=?"
	QUERY_ORDER_BY_ID   string = "SELECT id, order_number, user_id, payment_id, status, created_time, updated_time FROM porder where id =?"
	QUERY_ORDER_BY_NAME        = "SELECT * FROM porder where order_number =?"
	QUERY_ORDER                = "SELECT * FROM porder "
	CREATE_PAYMENT               = "update porder set payment_id=?, status=?, updated_time=? where order_number=?"
	INSERT_ORDER               = "INSERT porder SET order_number =?, user_id=?, " +
		"status=?, created_time=?, updated_time=?"
)

// OrderDataSql is the SQL implementation of OrderDataInterface
type OrderDataSql struct {
	DB gdbc.SqlGdbc
}

func (ods *OrderDataSql) Remove(orderNumber string) (int64, error) {

	stmt, err := ods.DB.Prepare(DELETE_ORDER)
	if err != nil {
		return 0, errors.Wrap(err, "")
	}
	defer stmt.Close()

	res, err := stmt.Exec(orderNumber)
	if err != nil {
		return 0, errors.Wrap(err, "")
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "")
	}

	logger.Log.Debug("remove:row affected ", rowsAffected)
	return rowsAffected, nil
}

func (ods *OrderDataSql) Find(id int) (*model.Order, error) {
	rows, err := ods.DB.Query(QUERY_ORDER_BY_ID, id)
	if err != nil {
		return nil, errors.Wrap(err, "")
	}
	defer rows.Close()
	return retrieveOrder(rows)
}
func retrieveOrder(rows *sql.Rows) (*model.Order, error) {
	if rows.Next() {
		return rowsToOrder(rows)
	}
	return nil, nil
}
func rowsToOrder(rows *sql.Rows) (*model.Order, error) {
	var ct string
	var ut string
	order := &model.Order{}
	err := rows.Scan(&order.Id, &order.OrderNumber,&order.UserId, &order.P.Id, &order.Status, &ct, &ut)
	if err != nil {
		return nil, errors.Wrap(err, "")
	}
	createdTime, err := timea.Parse(timea.FORMAT_ISO8601_DATE_TIME, ct)
	if err != nil {
		return nil, errors.Wrap(err, "")
	}
	updatedTime, err := timea.Parse(timea.FORMAT_ISO8601_DATE_TIME, ut)
	if err != nil {
		return nil, errors.Wrap(err, "")
	}
	order.CreatedTime = createdTime
	order.UpdatedTime = updatedTime
	//logger.Log.Debug("rows to Order:", *order)
	return order, nil
}
func (ods *OrderDataSql) FindByNumber(orderNumber string) (*model.Order, error) {
	//logger.Log.Debug("call FindByName() and name is:", name)
	rows, err := ods.DB.Query(QUERY_ORDER_BY_NAME, orderNumber)
	if err != nil {
		return nil, errors.Wrap(err, "")
	}
	defer rows.Close()
	return retrieveOrder(rows)
}

func (ods *OrderDataSql) FindAll() ([]model.Order, error) {

	rows, err := ods.DB.Query(QUERY_ORDER)
	if err != nil {
		return nil, errors.Wrap(err, "")
	}
	defer rows.Close()
	users := []model.Order{}

	//var ds string
	for rows.Next() {
		user, err := rowsToOrder(rows)
		if err != nil {
			return users, errors.Wrap(err, "")
		}
		users = append(users, *user)

	}
	//need to check error for rows.Next()
	if err = rows.Err(); err != nil {
		return nil, errors.Wrap(err, "")
	}
	logger.Log.Debug("find user list:", users)
	return users, nil
}

func (ods *OrderDataSql) UpdatePayment( orderNumber string, paymentId int,status string ) (int64, error) {

	stmt, err := ods.DB.Prepare(CREATE_PAYMENT)

	if err != nil {
		return 0, errors.Wrap(err, "")
	}
	defer stmt.Close()
	res, err := stmt.Exec(paymentId, status, time.Now(),orderNumber)
	if err != nil {
		return 0, errors.Wrap(err, "")
	}
	rowsAffected, err := res.RowsAffected()

	if err != nil {
		return 0, errors.Wrap(err, "")
	}
	logger.Log.Debug("UpdatePayment: rows affected: ", rowsAffected)

	return rowsAffected, nil
}

func (ods *OrderDataSql) Insert(o *model.Order) (*model.Order, error) {

	stmt, err := ods.DB.Prepare(INSERT_ORDER)
	if err != nil {
		return nil, errors.Wrap(err, "")
	}
	defer stmt.Close()
	res, err := stmt.Exec(o.OrderNumber,o.UserId, o.Status,
		o.CreatedTime, o.UpdatedTime)
	if err != nil {
		return nil, errors.Wrap(err, "")
	}
	id, err := res.LastInsertId()
	if err != nil {
		return nil, errors.Wrap(err, "")
	}
	o.Id = int(id)
	//logger.Log.Debug("user inserted:", o)
	return o, nil
}

