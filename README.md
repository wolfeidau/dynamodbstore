# dynamodbstore [![GoDoc](http://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)](http://godoc.org/github.com/wolfeidau/dynamodbstore)

A session store backend for [gorilla/sessions](https://github.com/gorilla/sessions) which uses AWS [DynamoDB](https://aws.amazon.com/dynamodb). It uses the official [AWS Golang SDK](github.com/aws/aws-sdk-go).

# Installation

```
go get -u github.com/wolfeidau/dynamodbstore
```

# Usage

```go
    sess := session.Must(session.NewSession())

    ddb := dynamodb.New(sess)

    store, err = NewDynamodbStore(ddb, []byte("secret-key"))
	if err != nil {
		t.Fatal(err.Error())
	}

    // Get a session.
    session, err = store.Get(req, "session-key")
    if err != nil {
        log.Error(err.Error())
    }

    // Add a value.
    session.Values["foo"] = "bar"

    // Save.
    if err = sessions.Save(req, rsp); err != nil {
        t.Fatalf("Error saving session: %v", err)
    }

    // Delete session.
    session.Options.MaxAge = -1
    if err = sessions.Save(req, rsp); err != nil {
        t.Fatalf("Error saving session: %v", err)
    }
```

# License

This code is released under MIT License, and is copyright Mark Wolfe.