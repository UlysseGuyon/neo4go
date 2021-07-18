package neo4go

import (
	"testing"
)

func Test_validateManagerOptions(t *testing.T) {
	type args struct {
		opt ManagerOptions
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Should allow basic options",
			args: args{
				opt: ManagerOptions{
					URI:          "bolt://localhost:7687",
					DatabaseName: "neo4j",
				},
			},
			want: false,
		},
		{
			name: "Should not allow options without database name",
			args: args{
				opt: ManagerOptions{
					URI: "bolt://localhost:7687",
				},
			},
			want: true,
		},
		{
			name: "Should not allow options without URI",
			args: args{
				opt: ManagerOptions{
					DatabaseName: "neo4j",
				},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := validateManagerOptions(tt.args.opt); (got != nil) != tt.want {
				t.Errorf("validateManagerOptions() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isWriteQuery(t *testing.T) {
	type args struct {
		query string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Should not detect read query",
			args: args{query: "MATCH (u:User {id: 'abcd'}) RETURN collect(DISTINCT u) AS uList"},
			want: false,
		},
		{
			name: "Should detect write query CREATE",
			args: args{query: "CREATE (u:User {id: 'abcd'})"},
			want: true,
		},
		{
			name: "Should detect write query MERGE",
			args: args{query: "MERGE (u:User {id: 'abcd'})"},
			want: true,
		},
		{
			name: "Should detect write query DELETE",
			args: args{query: "MATCH (u:User) DETACH DELETE u"},
			want: true,
		},
		{
			name: "Should detect write query SET",
			args: args{query: "MATCH (u:User {id: 'abcd'}) SET u.id = 'abcde'"},
			want: true,
		},
		{
			name: "Should detect write query REMOVE",
			args: args{query: "MATCH (u:User {id: 'abcd'}) REMOVE u.id"},
			want: true,
		},
		{
			name: "Should detect write query FOREACH",
			args: args{query: "FOREACH (id IN ['abcd'] | C.R.E.A.T.E (:User {id: id}))"},
			want: true,
		},
		{
			name: "Should detect write query DROP",
			args: args{query: "DROP DATABASE neo4j IF EXISTS"},
			want: true,
		},
		{
			name: "Should detect write query ALTER",
			args: args{query: "ALTER USER alice SET PASSWORD $password CHANGE NOT REQUIRED"},
			want: true,
		},
		{
			name: "Should detect write query RENAME",
			args: args{query: "RENAME USER alice TO alice_delete"},
			want: true,
		},
		{
			name: "Should detect write query GRANT",
			args: args{query: "GRANT ROLE my_role TO alice"},
			want: true,
		},
		{
			name: "Should detect write query REVOKE",
			args: args{query: "REVOKE ROLE my_role FROM alice"},
			want: true,
		},
		{
			name: "Should detect write query DENY",
			args: args{query: "DENY READ {prop} ON GRAPH foo RELATIONSHIP Type TO my_role"},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsWriteQuery(tt.args.query); got != tt.want {
				t.Errorf("isWriteQuery() = %v, want %v", got, tt.want)
			}
		})
	}
}
