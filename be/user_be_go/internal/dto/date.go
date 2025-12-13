package dto

import (
	"encoding/json"
	"strings"
	"time"
)

// Date 自定义日期类型，支持多种日期格式解析
type Date struct {
	time.Time
}

// UnmarshalJSON 实现 JSON 反序列化
// 支持以下格式：
// - "2006-01-02" (日期格式)
// - "2006-01-02T15:04:05Z07:00" (ISO 8601 格式)
// - "2006-01-02T15:04:05Z" (UTC 格式)
func (d *Date) UnmarshalJSON(data []byte) error {
	// 移除引号
	str := strings.Trim(string(data), `"`)
	if str == "" || str == "null" {
		return nil
	}

	// 尝试多种日期格式
	formats := []string{
		"2006-01-02",                    // 日期格式
		"2006-01-02T15:04:05Z07:00",     // ISO 8601 带时区
		"2006-01-02T15:04:05Z",          // UTC 格式
		"2006-01-02T15:04:05",           // 无时区
		time.RFC3339,                    // RFC3339
		time.RFC3339Nano,                // RFC3339 纳秒
	}

	var err error
	for _, format := range formats {
		d.Time, err = time.Parse(format, str)
		if err == nil {
			return nil
		}
	}

	return err
}

// MarshalJSON 实现 JSON 序列化
func (d Date) MarshalJSON() ([]byte, error) {
	if d.Time.IsZero() {
		return []byte("null"), nil
	}
	// 返回日期格式 "2006-01-02"
	return json.Marshal(d.Time.Format("2006-01-02"))
}

// ToTimePtr 转换为 *time.Time
func (d *Date) ToTimePtr() *time.Time {
	if d == nil || d.Time.IsZero() {
		return nil
	}
	return &d.Time
}

// NewDate 创建 Date 对象
func NewDate(t time.Time) *Date {
	return &Date{Time: t}
}

