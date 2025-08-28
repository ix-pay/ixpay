package utils

import (
	"errors"
	"strconv"
	"sync"
	"time"

	"github.com/ix-pay/ixpay/config"
)

func NewSnowflake(cfg *config.Config) (*Snowflake, error) {
	machineID, _ := strconv.ParseInt(cfg.MachineId, 10, 64)
	if machineID < 0 {
		return nil, errors.New("机器码异常")
	}

	sf, err := New(machineID)
	if err != nil {
		return nil, err
	}
	return sf, nil
}

const (
	epoch         int64 = 1609459200000 // 2021-01-01 00:00:00 UTC
	timestampBits uint8 = 41
	machineBits   uint8 = 10
	sequenceBits  uint8 = 12

	maxMachineID int64 = -1 ^ (-1 << machineBits)
	maxSequence  int64 = -1 ^ (-1 << sequenceBits)

	timestampShift = machineBits + sequenceBits
	machineShift   = sequenceBits
)

type Snowflake struct {
	mu        sync.Mutex
	lastStamp int64
	machineID int64
	sequence  int64
}

func New(machineID int64) (*Snowflake, error) {
	if machineID < 0 || machineID > maxMachineID {
		return nil, errors.New("machine ID out of range")
	}
	return &Snowflake{
		lastStamp: 0,
		machineID: machineID,
		sequence:  0,
	}, nil
}

func (s *Snowflake) Generate() int64 {
	s.mu.Lock()
	defer s.mu.Unlock()

	current := time.Now().UnixMilli()
	if current < s.lastStamp {
		panic("clock moved backwards")
	}

	if current == s.lastStamp {
		s.sequence = (s.sequence + 1) & maxSequence
		if s.sequence == 0 {
			for current <= s.lastStamp {
				current = time.Now().UnixMilli()
			}
		}
	} else {
		s.sequence = 0
	}

	s.lastStamp = current
	return ((current - epoch) << timestampShift) |
		(s.machineID << machineShift) |
		s.sequence
}
