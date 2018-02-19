# dynamodbstore [![GoDoc](http://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)](http://godoc.org/github.com/wolfeidau/dynamodbstore)

A session store backend for [gorilla/sessions](https://github.com/gorilla/sessions) which uses AWS [DynamoDB](https://aws.amazon.com/dynamodb). It uses the official [AWS Golang SDK](github.com/aws/aws-sdk-go).

# Installation

```
go get -u github.com/wolfeidau/dynamodbstore
```

# Usage

```go
    // you may want to configure this centrally in your application
    sess := session.Must(session.NewSession())

    ddb := dynamodb.New(sess)

    // secret-key should be generated
    secretKey := "secret-key-should-be-in-config"
    
    store, err := dynamodbstore.NewDynamodbStore(ddb, []byte(secretKey))
	if err != nil {
		log.Fatal(err)
	}

    // Get a session.
    session, err = store.Get(req, "session-key")
    if err != nil {
        log.Error(err)
    }

    // Add a value.
    session.Values["foo"] = "bar"

    // Save.
    if err = sessions.Save(req, rsp); err != nil {
        log.Fatalf("Error saving session: %v", err)
    }

    // Delete session.
    session.Options.MaxAge = -1
    if err = sessions.Save(req, rsp); err != nil {
        log.Fatalf("Error saving session: %v", err)
    }
```

# License

This code is released under MIT License, and is copyright Mark Wolfe.