package transfer

import (
	"context"
	"fmt"

	"github.com/babyname/fate/ent"
)

const MaxLimit = 1000

type Transfer interface {
	Start(ctx context.Context) error
}

type transferDatabase struct {
	Source *ent.Client
	Target *ent.Client
	Tables []string
}

func (t transferDatabase) Start(ctx context.Context) error {
	var err error
	err = t.Target.Schema.Create(ctx)
	if err != nil {
		return err
	}
	for _, table := range t.Tables {
		switch table {
		case "WuGeLucky":
			err = t.transferWuGeLucky(ctx)
		case "Character":
			err = t.transferCharacter(ctx)
		case "WuXing":
			err = t.transferWuXing(ctx)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

func (t transferDatabase) transferWuGeLucky(ctx context.Context) error {
	c, err := t.Source.WuGeLucky.Query().Count(ctx)
	if err != nil {
		return err
	}
	if c == 0 {
		return nil
	}

	for i := 0; i < c; i += MaxLimit {
		fmt.Println("insert database wugelucky:", i, "total", c)
		luckies, err := t.Source.WuGeLucky.Query().Limit(MaxLimit).Offset(i).All(ctx)
		if err != nil {
			return err
		}
		var luckyBluks []*ent.WuGeLuckyCreate
		for x := range luckies {
			lucky := t.Target.WuGeLucky.Create().SetWuGeLuckyWithOptional(luckies[x])
			luckyBluks = append(luckyBluks, lucky)
		}

		if len(luckyBluks) != 0 {
			_, err = t.Target.WuGeLucky.CreateBulk(luckyBluks...).Save(ctx)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (t transferDatabase) transferCharacter(ctx context.Context) error {
	c, err := t.Source.Character.Query().Count(ctx)
	if err != nil {
		return err
	}
	if c == 0 {
		return nil
	}

	for i := 0; i < c; i += MaxLimit {
		fmt.Println("insert database character:", i, "total", c)
		characters, err := t.Source.Character.Query().Limit(MaxLimit).Offset(i).All(ctx)
		if err != nil {
			return err
		}
		var bluks []*ent.CharacterCreate
		for x := range characters {
			character := t.Target.Character.Create().SetCharacterWithOptional(characters[x])
			bluks = append(bluks, character)
		}

		if len(bluks) != 0 {
			_, err = t.Target.Character.CreateBulk(bluks...).Save(ctx)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (t transferDatabase) transferWuXing(ctx context.Context) error {
	c, err := t.Source.WuXing.Query().Count(ctx)
	if err != nil {
		return err
	}
	if c == 0 {
		return nil
	}

	for i := 0; i < c; i += MaxLimit {
		fmt.Println("insert database wuxing:", i, "total", c)
		wuxings, err := t.Source.WuXing.Query().Limit(MaxLimit).Offset(i).All(ctx)
		if err != nil {
			return err
		}
		var bluks []*ent.WuXingCreate
		for x := range wuxings {
			wuxing := t.Target.WuXing.Create().SetWuXingWithOptional(wuxings[x])
			bluks = append(bluks, wuxing)
		}

		if len(bluks) != 0 {
			_, err = t.Target.WuXing.CreateBulk(bluks...).Save(ctx)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func newTransfer(c *DatabaseConfig) (*transferDatabase, error) {
	source, err := c.Source.Database().BuildClient()
	if err != nil {
		return nil, fmt.Errorf("could not open source database: %v", err)
	}
	target, err := c.Target.Database().BuildClient()
	if err != nil {
		return nil, fmt.Errorf("could not open target database: %v", err)
	}
	return &transferDatabase{
		Tables: c.Tables,
		Source: source,
		Target: target,
	}, nil
}

func NewTransfer(config *DatabaseConfig) (Transfer, error) {
	return newTransfer(config)
}

var _ Transfer = (*transferDatabase)(nil)
