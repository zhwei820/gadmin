package api

import (
	"github.com/gogf/gf/g"
	"github.com/gogf/gf/g/os/gtime"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetWhereFromQuery(t *testing.T) {

	t.Parallel()

	tests := []struct {
		querys map[string]interface{}
		wheres map[string]interface{}
	}{
		{
			querys: map[string]interface{}{
				"nickname__contains": "name_2",
			},
			wheres: map[string]interface{}{
				"nickname like BINARY ? ": "%name_2%",
			},
		},
		{
			querys: map[string]interface{}{
				"nickname__icontains": "NAME_2",
			},
			wheres: map[string]interface{}{
				"nickname like ? ": "%NAME_2%",
			},
		},
		{
			querys: map[string]interface{}{
				"nickname__range": "name_0,name_2",
			},
			wheres: map[string]interface{}{
				"nickname between ? AND ? ": []string{"name_0", "name_2"},
			},
		},
		{
			querys: map[string]interface{}{
				"nickname__in": "name_0,name_2",
			},
			wheres: map[string]interface{}{
				"nickname in (?) ": []string{"name_0", "name_2"},
			},
		},
		{
			querys: map[string]interface{}{
				"nickname__gt": "name_2",
			},
			wheres: map[string]interface{}{
				"nickname > ? ": "name_2",
			},
		},
		{
			querys: map[string]interface{}{
				"nickname__lt": "name_2",
			},
			wheres: map[string]interface{}{
				"nickname < ? ": "name_2",
			},
		},
		{
			querys: map[string]interface{}{
				"nickname__gte": "name_2",
			},
			wheres: map[string]interface{}{
				"nickname >= ? ": "name_2",
			},
		},
		{
			querys: map[string]interface{}{
				"nickname__lte": "name_2",
			},
			wheres: map[string]interface{}{
				"nickname <= ? ": "name_2",
			},
		},

		{
			querys: map[string]interface{}{
				"id__contains": 2,
			},
			wheres: map[string]interface{}{
				"id like BINARY ? ": 2,
			},
		},
		{
			querys: map[string]interface{}{
				"id__icontains": 2,
			},
			wheres: map[string]interface{}{
				"id like ? ": 2,
			},
		},

		{
			querys: map[string]interface{}{
				"id__range": []int{0, 2},
			},
			wheres: map[string]interface{}{
				"id between ? AND ? ": []int{0, 2},
			},
		},
		{
			querys: map[string]interface{}{
				"id__in": []int{0, 2},
			},
			wheres: map[string]interface{}{
				"id in (?) ": []int{0, 2},
			},
		},
		{
			querys: map[string]interface{}{
				"id__gt": 2,
			},
			wheres: map[string]interface{}{
				"id > ? ": 2,
			},
		},
		{
			querys: map[string]interface{}{
				"id__lt": 2,
			},
			wheres: map[string]interface{}{
				"id < ? ": 2,
			},
		},
		{
			querys: map[string]interface{}{
				"id__gte": 2,
			},
			wheres: map[string]interface{}{
				"id >= ? ": 2,
			},
		},
		{
			querys: map[string]interface{}{
				"id__lte": 2,
			},
			wheres: map[string]interface{}{
				"id <= ? ": 2,
			},
		},
	}
	db.SetDebug(true)
	table := createInitTable()
	defer dropTable(table)
	for _, test := range tests {
		assert.Equal(t, GetWhereFromQuery(test.querys), test.wheres)
		type User struct {
			Id         int
			Passport   string
			Password   string
			NickName   string
			CreateTime *gtime.Time
		}
		var users []*User
		_ = db.Table(table).Where(GetWhereFromQuery(test.querys)).Structs(&users)
		g.Dump("====================")
		g.Dump(test.wheres)
		g.Dump(users)
		continue
	}
}
