package repository

import (
	"PowerX/internal/types"
	"PowerX/pkg/zerox"
	"context"
	"errors"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// BaseRepository 提供通用的 CRUD 操作
type BaseRepository[T any] struct {
	db *gorm.DB
}

// NewBaseRepository 创建新的 BaseRepository 实例
func NewBaseRepository[T any](db *gorm.DB) *BaseRepository[T] {
	return &BaseRepository[T]{db: db}
}

// Create 创建新记录，并返回创建后的对象
func (r *BaseRepository[T]) Create(ctx context.Context, obj *T) (*T, error) {
	query := r.db.WithContext(ctx)

	debug, ok := ctx.Value(zerox.DebugKey).(bool)
	// print(debug, ok)
	if ok && debug {
		query = query.Debug()
	}

	result := query.Create(obj)
	if result.Error != nil {
		return nil, result.Error
	}
	return obj, nil
}

// Upsert 插入或更新单个记录，并返回执行后的对象
func (r *BaseRepository[T]) Upsert(ctx context.Context, obj *T, uniqueFields []string) (*T, error) {
	// Upsert 依赖于 OnConflict 来进行冲突处理
	// uniqueFields 是用于识别冲突的唯一字段，可以根据业务进行定制（例如 uuid）

	query := r.db.WithContext(ctx)

	debug, ok := ctx.Value(zerox.DebugKey).(bool)
	// print(debug, ok)
	if ok && debug {
		query = query.Debug()
	}

	result := query.Clauses(clause.OnConflict{
		UpdateAll: true, // 如果冲突，则更新所有字段
	}).
		Create(obj)

	if result.Error != nil {
		return nil, result.Error
	}

	return obj, nil
}

// UpsertBatch 批量插入或更新记录，并返回执行后的对象列表
func (r *BaseRepository[T]) UpsertBatch(ctx context.Context, objs []*T, uniqueFields []string) ([]*T, error) {
	// 使用事务来保证批量 Upsert 的原子性
	tx := r.db.WithContext(ctx).Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}

	debug, ok := ctx.Value(zerox.DebugKey).(bool)
	// print(debug, ok)
	if ok && debug {
		tx = tx.Debug()
	}

	// 执行批量 Upsert 操作
	result := tx.Clauses(clause.OnConflict{
		UpdateAll: true, // 如果冲突，则更新所有字段
	}).Create(objs)

	if result.Error != nil {
		tx.Rollback()
		return nil, result.Error
	}

	// 提交事务并返回执行后的对象列表
	if err := tx.Commit().Error; err != nil {
		return nil, err
	}

	return objs, nil
}

// Update 更新记录，并返回更新后的对象
func (r *BaseRepository[T]) Update(ctx context.Context, obj *T) (*T, error) {
	query := r.db.WithContext(ctx)

	debug, ok := ctx.Value(zerox.DebugKey).(bool)
	// print(debug, ok)
	if ok && debug {
		query = query.Debug()
	}

	result := query.Save(obj)
	if result.RowsAffected == 0 {
		return nil, errors.New("record not found")
	}
	return obj, result.Error
}

// Patch 部分更新记录
func (r *BaseRepository[T]) Patch(ctx context.Context, uuid string, fields map[string]interface{}) (*T, error) {
	var obj T
	query := r.db.WithContext(ctx).Model(&obj)

	debug, ok := ctx.Value(zerox.DebugKey).(bool)
	// print(debug, ok)
	if ok && debug {
		query = query.Debug()
	}

	result := query.Where("uuid = ?", uuid).Updates(fields)
	if result.RowsAffected == 0 {
		return nil, errors.New("record not found")
	}
	return &obj, result.Error
}

// Delete 删除记录，并返回删除的对象
func (r *BaseRepository[T]) Delete(ctx context.Context, obj *T) (*T, error) {
	query := r.db.WithContext(ctx)

	debug, ok := ctx.Value(zerox.DebugKey).(bool)
	// print(debug, ok)
	if ok && debug {
		query = query.Debug()
	}

	result := query.Delete(obj)

	if result.RowsAffected == 0 {
		return nil, errors.New("record not found")
	}
	return obj, result.Error
}

// FindByCondition 根据条件查询并返回分页结果（指针数组）
func (r *BaseRepository[T]) FindByCondition(
	ctx context.Context,
	conditions map[string]interface{},
	page, pageSize int,
	callback func(db *gorm.DB, opt interface{}) *gorm.DB,
	opt interface{},
) (*types.Page[*T], error) {
	var results []*T
	var obj T

	// 构造查询条件
	query := r.db.WithContext(ctx)
	for key, value := range conditions {
		query = query.Where(key, value)
	}
	// 调用回调函数，让调用方自定义查询逻辑
	if callback != nil {
		query = callback(query, opt)
	}

	debug, ok := ctx.Value(zerox.DebugKey).(bool)
	// print(debug, ok)
	if ok && debug {
		query = query.Debug()
	}

	// 获取总记录数
	var totalCount int64
	countQuery := query.Model(&obj)
	resultCount := countQuery.Count(&totalCount)
	if resultCount.Error != nil {
		return nil, resultCount.Error
	}

	// 分页查询
	query = query.Limit(pageSize).Offset((page - 1) * pageSize)
	result := query.Find(&results)
	if result.Error != nil {
		return nil, result.Error
	}

	return &types.Page[*T]{
		List:      results,
		PageIndex: page,
		PageSize:  pageSize,
		Total:     totalCount,
	}, nil

}

// GetByID 通过 ID 获取记录（返回指针）
func (r *BaseRepository[T]) GetByID(ctx context.Context, id int64, callback func(*gorm.DB) *gorm.DB) (*T, error) {
	var obj T

	// 构造查询
	query := r.db.WithContext(ctx).Where("id = ?", id)

	if callback != nil {
		query = callback(query)
	}

	debug, ok := ctx.Value(zerox.DebugKey).(bool)
	// print(debug, ok)
	if ok && debug {
		query = query.Debug()
	}

	// 执行查询
	result := query.First(&obj)

	// 处理错误
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, nil // 查询不到记录返回 nil，而不是直接抛出错误
	}
	if result.Error != nil {
		return nil, result.Error // 处理其他数据库错误
	}
	return &obj, nil
}

// GetByUUID 通过 UUID 获取记录（返回指针）
func (r *BaseRepository[T]) GetByUUID(ctx context.Context, uuid string, callback func(*gorm.DB) *gorm.DB) (*T, error) {
	var obj T

	// 构造查询
	query := r.db.WithContext(ctx).Where("uuid = ?", uuid)
	// 调用回调函数，让调用方自定义查询逻辑
	if callback != nil {
		query = callback(query)
	}

	debug, ok := ctx.Value(zerox.DebugKey).(bool)
	// print(debug, ok)
	if ok && debug {
		query = query.Debug()
	}

	// 执行查询
	result := query.First(&obj)

	// 处理错误
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, nil // 查询不到记录返回 nil，而不是直接抛出错误
	}
	if result.Error != nil {
		return nil, result.Error // 处理其他数据库错误
	}
	return &obj, nil
}

// GetByCondition 根据条件查询单个记录（返回指针）
func (r *BaseRepository[T]) GetByCondition(ctx context.Context, conditions map[string]interface{}, callback func(*gorm.DB) *gorm.DB) (*T, error) {
	var obj T

	// 构造查询条件
	query := r.db.WithContext(ctx)
	for key, value := range conditions {
		query = query.Where(key, value)
	}

	// 调用回调函数，让调用方自定义查询逻辑
	if callback != nil {
		query = callback(query)
	}

	debug, ok := ctx.Value(zerox.DebugKey).(bool)
	// print(debug, ok)
	if ok && debug {
		query = query.Debug()
	}

	// 获取单条记录
	result := query.First(&obj)
	// 处理错误
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, nil // 查询不到记录返回 nil，而不是直接抛出错误
	}
	if result.Error != nil {
		return nil, result.Error // 处理其他数据库错误
	}

	return &obj, nil
}
