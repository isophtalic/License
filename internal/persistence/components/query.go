package components

import (
	"fmt"
	"math"
	"net/http"

	"github.com/isophtalic/License/internal/dto"
	customError "github.com/isophtalic/License/internal/error"
	"gorm.io/gorm"
)

/*
Query is custom query. <T> is type of model
*/
type Query[T interface{}] struct {
	db     *gorm.DB
	query  *gorm.DB
	value  *T
	values []T
}

/*
Initiate new instance of query.
*/
func NewQuery[T interface{}](model interface{}, database *gorm.DB) *Query[T] {
	return &Query[T]{
		db:     database.Model(model),
		query:  database.Model(model),
		value:  nil,
		values: nil,
	}
}

func (q *Query[T]) Origin() *gorm.DB {
	return q.query
}

/*
Like Flush() but return Query
*/
func (q *Query[T]) New() *Query[T] {
	q.Flush()
	return q
}

/*
Reset all values to initial values.
*/
func (q *Query[T]) Flush() {
	q.value = nil
	q.values = nil
	q.query = q.db
}

/*
Return one record. Use this func when finding one.
*/
func (q *Query[T]) Value() *T {
	return q.value
}

/*
Return many records. Use this func when finding many.
*/
func (q *Query[T]) Values() []T {
	return q.values
}

/*
Include: select specific attributes to return.
*/
func (q *Query[T]) Include(attributes ...string) *Query[T] {
	q.query = q.query.Select(attributes)
	return q
}

/*
Exclude: Omit specific attributes.
*/
func (q *Query[T]) Exclude(attributes ...string) *Query[T] {
	q.query = q.query.Omit(attributes...)
	return q
}

func (q *Query[T]) Where(conditions map[string]interface{}) *Query[T] {
	q.query = q.query.Where(conditions)
	return q
}

func (q *Query[T]) Or(conditions map[string]interface{}) *Query[T] {
	q.query = q.query.Or(conditions)
	return q
}

/*
Loading associate.
Param attributes should includes primary key and foreign key.
*/
func (q *Query[T]) Association(name string, attributes []string) *Query[T] {
	if len(attributes) > 0 {
		q.query = q.query.Preload(name, func(db *gorm.DB) *gorm.DB {
			return db.Select(attributes)
		})
	} else {
		q.query = q.query.Preload(name)
	}

	return q
}

/*
Pagination with limit & page & sort
*/
func (q *Query[T]) Pagination(pagination *dto.PaginationDTO) *Query[T] {
	var totalRows int64
	q.query.Count(&totalRows)
	pagination.TotalRows = totalRows
	pagination.TotalPages = int(math.Ceil(float64(totalRows) / float64(pagination.GetPerPage())))

	q.query = q.query.Scopes(func(db *gorm.DB) *gorm.DB {
		return db.Offset(pagination.GetOffset()).Limit(pagination.GetPerPage()).Order(pagination.GetSort())
	})
	return q
}

func (q *Query[T]) FindAll() *Query[T] {
	q.value = nil
	err := q.query.Find(&q.values).Error
	if err != nil {
		println(err.Error())
		q.values = nil
	}
	return q
}

/*
Find one with specific value of any attributes.
*/
func (q *Query[T]) FindOneByAttributes(attr map[string]interface{}) *Query[T] {
	q.values = nil
	var x T
	err := q.query.Where(attr).First(&x).Error
	if err == nil {
		q.value = &x
	}
	return q
}

/*
Find many with specific value of any attributes.
*/
func (q *Query[T]) FindManyByAttributes(attr map[string]interface{}) *Query[T] {
	q.value = nil
	var x []T
	err := q.query.Where(attr).Find(&x).Error
	if err == nil {
		q.values = x
	}
	return q
}

func (q *Query[T]) Create(data interface{}) *Query[T] {
	r := q.query.Create(data)
	if r.Error != nil {
		customError.Throw(http.StatusBadRequest, fmt.Sprintf("Query: %v", r.Error))
	}
	return q
}

func (q *Query[T]) Update(data interface{}) *Query[T] {
	r := q.query.Updates(data)
	if r.Error != nil {
		customError.Throw(http.StatusBadRequest, fmt.Sprintf("Query: %v", r.Error))
	}
	return q
}

func (q *Query[T]) DeleteByAttributes(attrs map[string]interface{}) *Query[T] {
	var x T
	r := q.query.Where(attrs).Delete(&x)
	if r.Error != nil {
		customError.Throw(http.StatusBadRequest, fmt.Sprintf("Query: %v", r.Error))
	}
	return q
}
