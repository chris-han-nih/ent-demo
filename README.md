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
---
## Fields
### Types
현재 프레임워크에서 지원하는 타입은 다음과 같다.
- All Go numeric types. Like `int`, `uint8`, `float64`, etc.
- `bool`
- `string`
- `time.Time`
- `UUID`
- `[]byte` (SQL only).
- `JSON` (SQL only).
- `Enum` (SQL only).
- `Other` (SQL only).
```go
package schema

import (
    "time"
    "net/url"

    "github.com/google/uuid"
    "entgo.io/ent"
    "entgo.io/ent/schema/field"
)

// User schema.
type User struct {
    ent.Schema
}

// Fields of the user.
func (User) Fields() []ent.Field {
    return []ent.Field{
        field.Int("age").
            Positive(),
        field.Float("rank").
            Optional(),
        field.Bool("active").
            Default(false),
        field.String("name").
            Unique(),
        field.Time("created_at").
            Default(time.Now),
        field.JSON("url", &url.URL{}).
            Optional(),
        field.JSON("strings", []string{}).
            Optional(),
        field.Enum("state").
            Values("on", "off").
            Optional(),
        field.UUID("uuid", uuid.UUID{}).
            Default(uuid.New),
    }
}
```