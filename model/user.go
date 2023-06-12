package model

import (
	"context"
	"github.com/chris-han-nih/ent-demo/ent"
	"github.com/chris-han-nih/ent-demo/ent/user"
	"net/url"
	"time"
)

type User struct {
	Id        int
	Name      string
	Age       int
	Rank      float64
	Active    bool
	CreatedAt time.Time
	Url       *url.URL
	Strings   []string
	State     string
}

func (u *User) Create(ctx context.Context, client *ent.Client) (int, error) {
	result, err := client.User.
		Create().
		SetAge(u.Age).
		SetRank(u.Rank).
		SetActive(u.Active).
		SetName(u.Name).
		SetURL(u.Url).
		SetStrings(u.Strings).
		SetState(user.State(u.State)).
		Save(ctx)
	if err != nil {
		return 0, err
	}
	return result.ID, nil
}
