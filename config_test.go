package main

import (
	"reflect"
	"testing"
)

func TestInsertInMap(t *testing.T) {

	type args struct {
		s     []string
		value string
		dest  map[string]interface{}
	}

	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "allo",
			args: args{
				s:     []string{"a"},
				value: "allo",
				dest: map[string]interface{}{
					"asd": "qwe",
				},
			},
			want: map[string]interface{}{
				"a":   "allo",
				"asd": "qwe",
			},
		},
		{
			name: "b.a",
			args: args{
				s:     []string{"b", "a"},
				value: "allo",
				dest:  map[string]interface{}{},
			},
			want: map[string]interface{}{
				"b": map[string]interface{}{
					"a": "allo",
				},
			},
		},
		{
			name: "b.a already set",
			args: args{
				s:     []string{"b", "a"},
				value: "allo",
				dest: map[string]interface{}{
					"b": map[string]interface{}{
						"a": "bye",
					},
				},
			},
			want: map[string]interface{}{
				"b": map[string]interface{}{
					"a": "bye",
				},
			},
		},
		{
			name: "c.b.a",
			args: args{
				s:     []string{"c", "b", "a"},
				value: "allo",
				dest:  map[string]interface{}{},
			},
			want: map[string]interface{}{
				"c": map[string]interface{}{
					"b": map[string]interface{}{
						"a": "allo",
					},
				},
			},
		},
		{
			name: "c.b.a",
			args: args{
				s:     []string{"a", "b", "c"},
				value: "bye",
				dest: map[string]interface{}{
					"a": map[string]interface{}{
						"b": map[string]interface{}{
							"d": "allo",
						},
					},
				},
			},
			want: map[string]interface{}{
				"a": map[string]interface{}{
					"b": map[string]interface{}{
						"d": "allo",
						"c": "bye",
					},
				},
			},
		},
		{
			name: "c.b. asdasa",
			args: args{
				s:     []string{"files", "fileb", "size"},
				value: "20",
				dest: map[string]interface{}{
					"app": map[string]interface{}{
						"source": map[string]interface{}{
							"git":     "github.com",
							"version": "1.2.3",
						},
						"attributes": map[string]interface{}{
							"name": "myapp",
						},
					},
					"files": map[string]interface{}{
						"filea": map[string]interface{}{
							"name": "filea",
							"size": "12",
						},
						"fileb": map[string]interface{}{
							"name": "fileb",
						},
						"filec": map[string]interface{}{
							"name": "filec",
						},
					},
				},
			},
			want: map[string]interface{}{
				"app": map[string]interface{}{
					"source": map[string]interface{}{
						"git":     "github.com",
						"version": "1.2.3",
					},
					"attributes": map[string]interface{}{
						"name": "myapp",
					},
				},
				"files": map[string]interface{}{
					"filea": map[string]interface{}{
						"name": "filea",
						"size": "12",
					},
					"fileb": map[string]interface{}{
						"name": "fileb",
						"size": "20",
					},
					"filec": map[string]interface{}{
						"name": "filec",
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := insertInMap(tt.args.s, tt.args.value, tt.args.dest)
			if !reflect.DeepEqual(tt.want, got) {
				t.Errorf("insertInMap() = %v, want %v", got, tt.want)
			}
		})
	}
}
