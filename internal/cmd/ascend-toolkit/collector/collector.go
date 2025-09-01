package collector

import (
	"context"
	"fmt"

	"ascend-common/devmanager"

	ascendv1 "github.com/innoai-tech/ascend-toolkit/pkg/apis/ascend/v1"
)

type Collector struct {
	ascendv1.Ascend

	Debug bool

	dm *devmanager.DeviceManager
}

func (c *Collector) Init(ctx context.Context) error {
	dm, err := devmanager.AutoInit("")
	if err != nil {
		return err
	}
	c.dm = dm
	return nil
}

func (c *Collector) Shutdown(ctx context.Context) error {
	if c.dm == nil {
		return nil
	}
	return c.dm.ShutDown()
}

func (c *Collector) Collect(ctx context.Context) error {
	if c.dm == nil {
		return nil
	}

	c.Type = c.dm.GetDevType()
	c.Version = c.dm.GetDcmiVersion()
	c.ProductTypes = c.dm.GetProductTypeArray()

	baseInfos, err := c.dm.GetChipBaseInfos()
	if err != nil {
		return err
	}

	for _, bi := range baseInfos {
		cc := &ChipCollector{}
		cc.DeviceID = bi.DeviceID
		cc.LogicID = bi.LogicID
		cc.PhysicID = bi.PhysicID
		cc.CardID = bi.CardID

		if err := cc.Collect(ctx, c.dm); err != nil {
			return err
		}

		if c.Debug {
			if err := cc.Err(); err != nil {
				fmt.Println(err)
			}
		}

		c.Chips = append(c.Chips, &cc.Chip)
	}

	return nil
}
