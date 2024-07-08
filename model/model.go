package model

import "time"

type UserRole string

const (
	RoleAdmin  UserRole = "Admin"
	RoleMember UserRole = "Member"
)

type Team struct {
	ID              string     `json:"id"`
	Name            string    `json:"name"`
	Description     string    `json:"description"`
	Bio             string    `json:"bio"`
	ProfileImageURL string    `json:"profile_image_url"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

type TeamUser struct {
	ID       string     `json:"id"`
	UserID   string     `json:"user_id"`
	TeamID   string     `json:"team_id"`
	JoinDate time.Time `json:"join_date"`
	Role     UserRole  `json:"role"`
}

type RequestStatus string

const (
	StatusPending  RequestStatus = "Pending"
	StatusRejected RequestStatus = "Rejected"
)

type Request struct {
	ID     string         `json:"id"`
	UserID string         `json:"user_id"`
	TeamID string         `json:"team_id"`
	Status RequestStatus `json:"status"`
	SentAt time.Time     `json:"sent_at"`
}

type TechStack struct {
	ID         string  `json:"id"`
	Technology string `json:"technology"`
	TeamID     string  `json:"team_id"`
}

type Comment struct {
	ID        string     `json:"id"`
	UserID    string     `json:"user_id"`
	TeamID    string     `json:"team_id"`
	Text      string    `json:"text"`
	ParentID  *string    `json:"parent_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

