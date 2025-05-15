// Package st 时间处理
package st

import (
	"sync"
	"time"
)

// STime 时间模拟器
type STime struct {
	mockTime *time.Time
	lock     sync.RWMutex
}

var (
	sTime     *STime
	sTimeOnce sync.Once
)

// GetTime 获取时间单例
func GetTime() *STime {
	sTimeOnce.Do(func() {
		sTime = &STime{}
	})
	return sTime
}

// Now 获取当前时间，如果设置了模拟时间则返回模拟时间
func (t *STime) Now() time.Time {
	t.lock.RLock()
	defer t.lock.RUnlock()
	if t.mockTime != nil {
		return *t.mockTime
	}
	return time.Now()
}

// SetMockTime 设置模拟时间
func (t *STime) SetMockTime(mockTime time.Time) {
	t.lock.Lock()
	defer t.lock.Unlock()
	t.mockTime = &mockTime
}

// ResetMockTime 重置模拟时间，恢复使用系统时间
func (t *STime) ResetMockTime() {
	t.lock.Lock()
	defer t.lock.Unlock()
	t.mockTime = nil
}

func NowTime() time.Time {
	return GetTime().Now()
}
