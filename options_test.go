package dynamodbstore

import (
	"testing"

	"github.com/gorilla/sessions"
	"github.com/stretchr/testify/require"
)

func TestSessionOptions(t *testing.T) {
	assert := require.New(t)

	type args struct {
		opts []Option
	}
	tests := []struct {
		name string
		args args
		want *StoreOptions
	}{
		{
			name: "should update tablename",
			args: args{opts: []Option{TableName("testing")}},
			want: &StoreOptions{tableName: "testing"},
		},
		{
			name: "should update tablename and max age",
			args: args{opts: []Option{TableName("testing"), DefaultMaxAge(3600)}},
			want: &StoreOptions{tableName: "testing", defaultMaxAge: 3600},
		},
		{
			name: "should update tablename and session options",
			args: args{opts: []Option{TableName("testing"), SessionOptions(&sessions.Options{
				HttpOnly: true,
			})}},
			want: &StoreOptions{tableName: "testing", sessionOptions: &sessions.Options{
				HttpOnly: true,
			}},
		},
		{
			name: "should update tablename and key pairs",
			args: args{opts: []Option{TableName("testing"), KeyPairs([]byte("hello"))}},
			want: &StoreOptions{tableName: "testing", keyPairs: [][]byte{[]byte("hello")}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			so := &StoreOptions{}
			for _, opt := range tt.args.opts {
				opt(so)
			}
			assert.Equal(tt.want, so)
		})
	}
}
