// Code generated by goctl. DO NOT EDIT!

package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	borrowFieldNames          = builder.RawFieldNames(&Borrow{})
	borrowRows                = strings.Join(borrowFieldNames, ",")
	borrowRowsExpectAutoSet   = strings.Join(stringx.Remove(borrowFieldNames, "`id`", "`create_time`", "`update_time`", "`create_at`", "`update_at`"), ",")
	borrowRowsWithPlaceHolder = strings.Join(stringx.Remove(borrowFieldNames, "`id`", "`create_time`", "`update_time`", "`create_at`", "`update_at`"), "=?,") + "=?"

	cacheBorrowIdPrefix = "cache:borrow:id:"
)

type (
	borrowModel interface {
		Insert(ctx context.Context, data *Borrow) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*Borrow, error)
		Update(ctx context.Context, data *Borrow) error
		Delete(ctx context.Context, id int64) error
	}

	defaultBorrowModel struct {
		sqlc.CachedConn
		table string
	}

	Borrow struct {
		Id         int64     `db:"id"`
		UserId     string    `db:"user_id"`  // 账号
		BookId     int64     `db:"book_id"`  // 书本号
		Isreturn   int64     `db:"isreturn"` // 是否归还，0未归还，1归还。
		BorrowTime time.Time `db:"borrow_time"`
		ReturnTime time.Time `db:"return_time"`
	}
)

func newBorrowModel(conn sqlx.SqlConn, c cache.CacheConf) *defaultBorrowModel {
	return &defaultBorrowModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`borrow`",
	}
}

func (m *defaultBorrowModel) Delete(ctx context.Context, id int64) error {
	borrowIdKey := fmt.Sprintf("%s%v", cacheBorrowIdPrefix, id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, borrowIdKey)
	return err
}

func (m *defaultBorrowModel) FindOne(ctx context.Context, id int64) (*Borrow, error) {
	borrowIdKey := fmt.Sprintf("%s%v", cacheBorrowIdPrefix, id)
	var resp Borrow
	err := m.QueryRowCtx(ctx, &resp, borrowIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", borrowRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, id)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultBorrowModel) Insert(ctx context.Context, data *Borrow) (sql.Result, error) {
	borrowIdKey := fmt.Sprintf("%s%v", cacheBorrowIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?)", m.table, borrowRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.UserId, data.BookId, data.Isreturn, data.BorrowTime, data.ReturnTime)
	}, borrowIdKey)
	return ret, err
}

func (m *defaultBorrowModel) Update(ctx context.Context, data *Borrow) error {
	borrowIdKey := fmt.Sprintf("%s%v", cacheBorrowIdPrefix, data.Id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, borrowRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.UserId, data.BookId, data.Isreturn, data.BorrowTime, data.ReturnTime, data.Id)
	}, borrowIdKey)
	return err
}

func (m *defaultBorrowModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheBorrowIdPrefix, primary)
}

func (m *defaultBorrowModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", borrowRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultBorrowModel) tableName() string {
	return m.table
}
