// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package db

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
	"time"
)

type UsersRole string

const (
	UsersRoleAdmin  UsersRole = "admin"
	UsersRoleMember UsersRole = "member"
)

func (e *UsersRole) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = UsersRole(s)
	case string:
		*e = UsersRole(s)
	default:
		return fmt.Errorf("unsupported scan type for UsersRole: %T", src)
	}
	return nil
}

type NullUsersRole struct {
	UsersRole UsersRole `json:"users_role"`
	Valid     bool      `json:"valid"` // Valid is true if UsersRole is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullUsersRole) Scan(value interface{}) error {
	if value == nil {
		ns.UsersRole, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.UsersRole.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullUsersRole) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.UsersRole), nil
}

type UsersStatus string

const (
	UsersStatusActive   UsersStatus = "active"
	UsersStatusInactive UsersStatus = "inactive"
)

func (e *UsersStatus) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = UsersStatus(s)
	case string:
		*e = UsersStatus(s)
	default:
		return fmt.Errorf("unsupported scan type for UsersStatus: %T", src)
	}
	return nil
}

type NullUsersStatus struct {
	UsersStatus UsersStatus `json:"users_status"`
	Valid       bool        `json:"valid"` // Valid is true if UsersStatus is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullUsersStatus) Scan(value interface{}) error {
	if value == nil {
		ns.UsersStatus, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.UsersStatus.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullUsersStatus) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.UsersStatus), nil
}

type AppSetting struct {
	ID        uint64       `json:"id"`
	Name      string       `json:"name"`
	Value     string       `json:"value"`
	CreatedAt time.Time    `json:"created_at"`
	UpdatedAt time.Time    `json:"updated_at"`
	DeletedAt sql.NullTime `json:"deleted_at"`
}

type DepositCategory struct {
	ID          uint64       `json:"id"`
	Name        string       `json:"name"`
	Description string       `json:"description"`
	Status      string       `json:"status"`
	CreatedAt   time.Time    `json:"created_at"`
	UpdatedAt   time.Time    `json:"updated_at"`
	DeletedAt   sql.NullTime `json:"deleted_at"`
	Version     int32        `json:"version"`
}

type DepositMethod struct {
	ID                uint64       `json:"id"`
	DepositCategoryID int64        `json:"deposit_category_id"`
	Name              string       `json:"name"`
	Description       string       `json:"description"`
	MinDeposit        uint64       `json:"min_deposit"`
	MaxDeposit        uint64       `json:"max_deposit"`
	Rate              float64      `json:"rate"`
	Status            string       `json:"status"`
	CreatedAt         time.Time    `json:"created_at"`
	UpdatedAt         time.Time    `json:"updated_at"`
	DeletedAt         sql.NullTime `json:"deleted_at"`
	Version           int32        `json:"version"`
}

type Order struct {
	ID             uint64          `json:"id"`
	IDProvider     string          `json:"id_provider"`
	ProviderID     uint64          `json:"provider_id"`
	UserID         uint32          `json:"user_id"`
	Target         string          `json:"target"`
	Total          uint64          `json:"total"`
	BeginningValue sql.NullInt64   `json:"beginning_value"`
	Partial        sql.NullInt64   `json:"partial"`
	Status         string          `json:"status"`
	CreatedAt      time.Time       `json:"created_at"`
	UpdatedAt      time.Time       `json:"updated_at"`
	DeletedAt      sql.NullTime    `json:"deleted_at"`
	Version        int32           `json:"version"`
	Name           sql.NullString  `json:"name"`
	Price          sql.NullFloat64 `json:"price"`
}

type Provider struct {
	ID        uint64         `json:"id"`
	Name      string         `json:"name"`
	Link      string         `json:"link"`
	ApiKey    string         `json:"api_key"`
	ApiID     sql.NullString `json:"api_id"`
	Status    string         `json:"status"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt sql.NullTime   `json:"deleted_at"`
	Version   int32          `json:"version"`
}

type Service struct {
	ID                uint64       `json:"id"`
	Name              string       `json:"name"`
	Min               int32        `json:"min"`
	Max               int32        `json:"max"`
	Price             float64      `json:"price"`
	Margin            float64      `json:"margin"`
	ServiceCategoryID uint64       `json:"service_category_id"`
	ProviderID        uint64       `json:"provider_id"`
	Status            string       `json:"status"`
	CreatedAt         time.Time    `json:"created_at"`
	UpdatedAt         time.Time    `json:"updated_at"`
	DeletedAt         sql.NullTime `json:"deleted_at"`
	Version           int32        `json:"version"`
}

type ServiceCategory struct {
	ID        uint64       `json:"id"`
	Name      string       `json:"name"`
	Status    string       `json:"status"`
	CreatedAt time.Time    `json:"created_at"`
	UpdatedAt time.Time    `json:"updated_at"`
	DeletedAt sql.NullTime `json:"deleted_at"`
	Version   int32        `json:"version"`
}

type ServiceRecomendation struct {
	ID        uint64       `json:"id"`
	ServiceID uint64       `json:"service_id"`
	CreatedAt time.Time    `json:"created_at"`
	UpdatedAt time.Time    `json:"updated_at"`
	DeletedAt sql.NullTime `json:"deleted_at"`
	Version   int32        `json:"version"`
}

type Ticket struct {
	ID               uint64       `json:"id"`
	UserID           uint32       `json:"user_id"`
	TicketCategoryID uint64       `json:"ticket_category_id"`
	Title            string       `json:"title"`
	Description      string       `json:"description"`
	Status           string       `json:"status"`
	CreatedAt        time.Time    `json:"created_at"`
	UpdatedAt        time.Time    `json:"updated_at"`
	DeletedAt        sql.NullTime `json:"deleted_at"`
	Version          int32        `json:"version"`
}

type TicketCategory struct {
	ID          uint64       `json:"id"`
	Name        string       `json:"name"`
	Description string       `json:"description"`
	Status      string       `json:"status"`
	CreatedAt   time.Time    `json:"created_at"`
	UpdatedAt   time.Time    `json:"updated_at"`
	DeletedAt   sql.NullTime `json:"deleted_at"`
	Version     int32        `json:"version"`
}

type User struct {
	ID                    uint32       `json:"id"`
	UserStatusID          uint32       `json:"user_status_id"`
	UserTypeID            uint32       `json:"user_type_id"`
	FullName              string       `json:"full_name"`
	Email                 string       `json:"email"`
	Password              string       `json:"password"`
	PhoneNumber           string       `json:"phone_number"`
	Role                  UsersRole    `json:"role"`
	Status                UsersStatus  `json:"status"`
	EmailVerifiedAt       sql.NullTime `json:"email_verified_at"`
	PhoneNumberVerifiedAt sql.NullTime `json:"phone_number_verified_at"`
	ApiKey                string       `json:"api_key"`
	CreatedAt             time.Time    `json:"created_at"`
	UpdatedAt             time.Time    `json:"updated_at"`
	DeletedAt             sql.NullTime `json:"deleted_at"`
	Version               int32        `json:"version"`
}

type UserStatus struct {
	ID          uint32       `json:"id"`
	Name        string       `json:"name"`
	Description string       `json:"description"`
	CreatedAt   time.Time    `json:"created_at"`
	UpdatedAt   time.Time    `json:"updated_at"`
	DeletedAt   sql.NullTime `json:"deleted_at"`
	Version     int32        `json:"version"`
}

type UserType struct {
	ID          uint32       `json:"id"`
	Name        string       `json:"name"`
	Description string       `json:"description"`
	CreatedAt   time.Time    `json:"created_at"`
	UpdatedAt   time.Time    `json:"updated_at"`
	DeletedAt   sql.NullTime `json:"deleted_at"`
	Version     int32        `json:"version"`
}
