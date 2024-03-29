package models

import (
	"encoding/json"
	"time"

	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/validate"
	"github.com/gofrs/uuid"
)

// TimerConfig model struct
type TimerConfig struct {
	ID            uuid.UUID     `json:"id" db:"id"`
	CreatedAt     time.Time     `json:"created_at" db:"created_at"`
	UpdatedAt     time.Time     `json:"updated_at" db:"updated_at"`
	FocusTime     time.Duration `json:"focus_time" db:"focus_time"`
	BreakTime     time.Duration `json:"break_time" db:"break_time"`
	Interval      time.Duration `json:"interval" db:"interval"`
	TimerClientID uuid.UUID     `json:"timer_client_id" db:"timer_client_id"`
}

func NewTimerConfig(focusTime, breakTime, intervalTime time.Duration, timerClientID string) (*TimerConfig, error) {
	configUUID, err := uuid.NewV4()

	return &TimerConfig{
		configUUID,
		time.Now(),
		time.Now(),
		focusTime,
		breakTime,
		intervalTime,
		uuid.FromStringOrNil(timerClientID),
	}, err
}

// String is not required by pop and may be deleted
func (t TimerConfig) String() string {
	jt, _ := json.Marshal(t)
	return string(jt)
}

// TimerConfigs is not required by pop and may be deleted
type TimerConfigs []TimerConfig

// String is not required by pop and may be deleted
func (t TimerConfigs) String() string {
	jt, _ := json.Marshal(t)
	return string(jt)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (t *TimerConfig) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (t *TimerConfig) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (t *TimerConfig) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
