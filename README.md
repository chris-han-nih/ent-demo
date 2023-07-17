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

### ID Field
`id` 필드는 스키마에 내장되어 있으며 선언이 필요하지 않습니다. SQL 기반 데이터베이스에서 해당 유형은 기본적으로 int(그러나 codegen 옵션으로 변경할 수 있음)이며 데이터베이스에서 자동 증가됩니다.

`id` 필드가 모든 테이블에서 고유하도록 구성하려면 스키마 마이그레이션을 실행할 때 WithGlobalUniqueID 옵션을 사용하십시오.
`client.Schema.Create(ctx, migrate.WithGlobalUniqueID(true))`

id 필드에 대해 다른 구성이 필요하거나 응용 프로그램(예: UUID)에서 엔터티 생성 시 id 값을 제공해야 하는 경우 기본 제공 id 구성을 재정의합니다. 예를 들어:
```go
// Fields of the Group.
func (Group) Fields() []ent.Field {
    return []ent.Field{
        field.Int("id").
            StructTag(`json:"oid,omitempty"`),
    }
}

// Fields of the Blob.
func (Blob) Fields() []ent.Field {
    return []ent.Field{
        field.UUID("id", uuid.UUID{}).
            Default(uuid.New).
            StorageKey("oid"),
    }
}

// Fields of the Pet.
func (Pet) Fields() []ent.Field {
    return []ent.Field{
        field.String("id").
            MaxLen(25).
            NotEmpty().
            Unique().
            Immutable(),
    }
}
```

### Database Type
기본적으로 모든 필드는 데이터베이스에서 동일한 유형으로 저장됩니다. 예를 들어, `int` 필드는 모든 데이터베이스에서 `int`로 저장됩니다. 그러나 필드 유형을 변경하려는 경우 `SchemaType` 메서드를 사용하여 데이터베이스 유형을 변경할 수 있습니다.
```go
package schema

import (
    "entgo.io/ent"
    "entgo.io/ent/dialect"
    "entgo.io/ent/schema/field"
)

// Card schema.
type Card struct {
    ent.Schema
}

// Fields of the Card.
func (Card) Fields() []ent.Field {
    return []ent.Field{
        field.Float("amount").
            SchemaType(map[string]string{
                dialect.MySQL:    "decimal(6,2)",   // Override MySQL.
                dialect.Postgres: "numeric",        // Override Postgres.
            }),
    }
}
```