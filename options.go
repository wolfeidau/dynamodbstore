package dynamodbstore

import (
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
)

// StoreOptions used to configure the dynamodb session store
type StoreOptions struct {
	tableName      string
	sessionOptions *sessions.Options
	keyPairs       [][]byte
	defaultMaxAge  int
}

// Option used to configure the store
type Option func(opts *StoreOptions)

// TableName update the table name
func TableName(name string) Option {
	return func(opts *StoreOptions) {
		opts.tableName = name
	}
}

// SessionOptions update the session options provided to github.com/gorilla/sessions
func SessionOptions(sessionOpts *sessions.Options) Option {
	return func(opts *StoreOptions) {
		opts.sessionOptions = sessionOpts
	}
}

// KeyPairs update the key pairs provided to github.com/gorilla/sessions
func KeyPairs(keyPairs ...[]byte) Option {
	return func(opts *StoreOptions) {
		opts.keyPairs = keyPairs
	}
}

// DefaultMaxAge update the default max age provided to github.com/gorilla/sessions
func DefaultMaxAge(age int) Option {
	return func(opts *StoreOptions) {
		opts.defaultMaxAge = age
	}
}

// NewDynamodbStoreWithOptions new dynamodb session store using options
func NewDynamodbStoreWithOptions(ddb *dynamodb.DynamoDB, options ...Option) *DynamodbStore {

	// setup a "default" store using the original values
	opts := &StoreOptions{
		tableName: DefaultTableName,
		sessionOptions: &sessions.Options{
			Path:   "/",
			MaxAge: sessionExpire,
		},
		defaultMaxAge: 60 * 20, // 20 minutes seems like a reasonable default
	}

	for _, opt := range options {
		opt(opts)
	}

	return &DynamodbStore{
		DB:            ddb,
		TableName:     opts.tableName,
		Codecs:        securecookie.CodecsFromPairs(opts.keyPairs...),
		Options:       opts.sessionOptions,
		DefaultMaxAge: opts.defaultMaxAge,
	}
}
