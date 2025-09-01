package collector

import (
	"context"
	"errors"
	"fmt"

	"ascend-common/devmanager"
	"ascend-common/devmanager/common"

	ascendv1 "github.com/innoai-tech/ascend-toolkit/pkg/apis/ascend/v1"
)

type ChipCollector struct {
	ascendv1.Chip

	err error
}

func (c *ChipCollector) Err() error {
	return c.err
}

func (c *ChipCollector) collect(fn func() error) {
	if err := fn(); err != nil {
		c.err = errors.Join(c.err, err)
	}
}

func (c *ChipCollector) Collect(ctx context.Context, device devmanager.DeviceInterface) error {
	c.collect(func() error {
		info, err := device.GetChipInfo(c.LogicID)
		if err != nil {
			return fmt.Errorf("get chip info failed: %w", err)
		}
		c.Info.Type = info.Type
		c.Info.Name = info.Name
		c.Info.Version = info.Version
		c.Info.NpuName = info.NpuName
		return nil
	})

	c.collect(func() error {
		m, err := device.GetDeviceMemoryInfo(c.LogicID)
		if err != nil {
			return fmt.Errorf("get device memory info failed: %w", err)
		}
		c.Memory.Size = m.MemorySize
		c.Memory.Used = m.MemorySize - m.MemoryAvailable
		c.Memory.Utilization = m.Utilization
		c.Memory.Frequency = m.Frequency
		return nil
	})

	c.collect(func() error {
		m, err := device.GetDeviceHbmInfo(c.LogicID)
		if err != nil {
			return fmt.Errorf("get high bandwidth memory info failed: %w", err)
		}
		c.Memory.HighBandwidthMemory.Size = m.MemorySize
		c.Memory.HighBandwidthMemory.Used = m.Usage
		c.Memory.HighBandwidthMemory.Frequency = m.Frequency
		c.Memory.HighBandwidthMemory.Temperature = m.Temp
		c.Memory.HighBandwidthMemory.BandwidthUtilization = m.BandWidthUtilRate
		return nil
	})

	c.collect(func() error {
		temp, err := device.GetDeviceTemperature(c.LogicID)
		if err != nil {
			return fmt.Errorf("get device voltage failed: %w", err)
		}
		c.Temperature = temp
		return nil
	})

	c.collect(func() error {
		vol, err := device.GetDeviceVoltage(c.LogicID)
		if err != nil {
			return fmt.Errorf("get device voltage failed: %w", err)
		}
		c.Voltage = vol
		return nil
	})

	c.collect(func() error {
		if device.GetDevType() == common.Ascend310P {
			cardPower, err := device.GetMcuPowerInfo(c.CardID)
			if err != nil {
				return fmt.Errorf("get card power info failed: %w", err)
			}
			c.Power = cardPower
			return nil
		}
		power, err := device.GetDevicePowerInfo(c.LogicID)
		if err != nil {
			return fmt.Errorf("get device power info failed: %w", err)
		}
		c.Power = power
		return nil
	})

	c.collect(func() error {
		freq, err := device.GetDeviceFrequency(c.DeviceID, common.AICoreCurrentFreq)
		if err != nil {
			return fmt.Errorf("get ai core current frequency failed: %w", err)
		}
		c.AICore.CurrentFrequency = freq
		return nil
	})

	c.collect(func() error {
		freq, err := device.GetDeviceFrequency(c.DeviceID, common.AICoreRatedFreq)
		if err != nil {
			return fmt.Errorf("get ai core rated frequency failed: %w", err)
		}
		c.AICore.RatedFrequency = freq
		return nil
	})

	c.collect(func() error {
		util, err := device.GetDeviceUtilizationRate(c.LogicID, common.AICore)
		if err != nil {
			return fmt.Errorf("get ai core utilization: %w", err)
		}
		c.AICore.Utilization = int(util)
		return nil
	})

	c.collect(func() error {
		util, err := device.GetDeviceUtilizationRate(c.LogicID, common.Overall)
		if err != nil {
			return fmt.Errorf("get overall utilization: %w", err)
		}
		c.Overall.Utilization = int(util)
		return nil
	})

	c.collect(func() error {
		util, err := device.GetDeviceUtilizationRate(c.LogicID, common.VectorCore)
		if err != nil {
			return fmt.Errorf("get vector core utilization: %w", err)
		}
		c.Vector.Utilization = int(util)
		return nil
	})

	c.collect(func() error {
		h, err := device.GetDeviceHealth(c.LogicID)
		if err != nil {
			return fmt.Errorf("get vector core utilization: %w", err)
		}
		c.Health = h
		return nil
	})

	c.collect(func() error {
		h, err := device.GetDeviceNetWorkHealth(c.LogicID)
		if err != nil {
			return fmt.Errorf("get vector core utilization: %w", err)
		}
		c.NetworkHealth = h
		return nil
	})

	return nil
}
