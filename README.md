ent-demo
---
https://entgo.io/docs/getting-started/

## Initial Setup
```bash
$ go mod init github.com/bradford-hamilton/ent-demo
```

## Generate the schema
```bash
$ go run entgo.io/ent/cmd/ent init User
```

## Add a field to the User schema
```go
//entdemo/ent/schema/user.go

package schema

import (
    "entgo.io/ent"
    "entgo.io/ent/schema/field"
)

// Fields of the User.
func (User) Fields() []ent.Field {
    return []ent.Field{
        field.Int("age").
            Positive(),
        field.String("name").
            Default("unknown"),
    }
}
```

## Run the code generation tool
```bash
$ go generate ./...
```