package mongodb

import (
	"github.com/quangdangfit/log-service/pkg/paging"
)

type Option interface {
	apply(*option)
}

type option struct {
	filter interface{}
	sorter interface{}
	paging *paging.Pagination
	page   int
	limit  int
	hint   interface{}
}

type optionFn func(*option)

func (f optionFn) apply(opt *option) {
	f(opt)
}

func WithPaging(dPaging *paging.Pagination) Option {
	return optionFn(func(opt *option) {
		opt.page = int(dPaging.CurrentPage)
		opt.limit = int(dPaging.Limit)
		opt.paging = dPaging
	})
}

// WithFilter - default not filter
//
// # Use bson object to filter document
//
// see more: https://docs.mongodb.com/manual/reference/operator/query/
func WithFilter(filter interface{}) Option {
	return optionFn(func(opt *option) {
		opt.filter = filter
	})
}

// WithSorter - default sort descending by `_id` field
//
// { $sort: { <field1>: <sort order>, <field2>: <sort order> ... }}
//
// $sort takes a document that specifies the field(s) to sort by and the
// respective sort order. <sort order> can have one of the following values:
//
// 1  Sort ascending
//
// -1 Sort descending
//
// { $meta: "textScore" } Sort by the computed textScore metadata in descending
// order
func WithSorter(sorter interface{}) Option {
	return optionFn(func(opt *option) {
		opt.sorter = sorter
	})
}

func WithHint(hint interface{}) Option {
	return optionFn(func(opt *option) {
		opt.hint = hint
	})
}
