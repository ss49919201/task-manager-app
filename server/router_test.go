package server

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_router_SetMiddleware(t *testing.T) {
	testFn := middleware(func(f http.HandlerFunc) http.HandlerFunc { return f })

	type fields struct {
		middlewareFunctions []middleware
	}
	type args struct {
		m middleware
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *router
	}{
		{
			"add one",
			fields{},
			args{testFn},
			&router{middlewareFunctions: []middleware{testFn}},
		},
		{
			"add nil",
			fields{},
			args{nil},
			&router{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &router{
				middlewareFunctions: tt.fields.middlewareFunctions,
			}
			got := r.SetMiddleware(tt.args.m)
			assert.Equal(t, len(tt.want.middlewareFunctions), len(got.middlewareFunctions))
		})
	}
}
