// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// FavoriteDao is the data access object for table favorite.
type FavoriteDao struct {
	table   string          // table is the underlying table name of the DAO.
	group   string          // group is the database configuration group name of current DAO.
	columns FavoriteColumns // columns contains all the column names of Table for convenient usage.
}

// FavoriteColumns defines and stores column names for table favorite.
type FavoriteColumns struct {
	Id           string //
	CreatedAt    string //
	UpdatedAt    string //
	DeletedAt    string //
	FavoriteId   string //
	FavoriteById string //
}

//  favoriteColumns holds the columns for table favorite.
var favoriteColumns = FavoriteColumns{
	Id:           "id",
	CreatedAt:    "created_at",
	UpdatedAt:    "updated_at",
	DeletedAt:    "deleted_at",
	FavoriteId:   "favorite_id",
	FavoriteById: "favorite_by_id",
}

// NewFavoriteDao creates and returns a new DAO object for table data access.
func NewFavoriteDao() *FavoriteDao {
	return &FavoriteDao{
		group:   "default",
		table:   "favorite",
		columns: favoriteColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *FavoriteDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *FavoriteDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *FavoriteDao) Columns() FavoriteColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *FavoriteDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *FavoriteDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *FavoriteDao) Transaction(ctx context.Context, f func(ctx context.Context, tx *gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
