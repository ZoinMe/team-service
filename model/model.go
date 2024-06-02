package model

import "time"

type UserRole string

const (
	RoleAdmin  UserRole = "Admin"
	RoleMember UserRole = "Member"
)

type Team struct {
	ID              int64     `json:"id"`
	Name            string    `json:"name"`
	Description     string    `json:"description"`
	Bio             string    `json:"bio"`
	ProfileImageURL string    `json:"profile_image_url"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

type TeamUser struct {
	ID       int64     `json:"id"`
	UserID   int64     `json:"user_id"`
	TeamID   int64     `json:"team_id"`
	JoinDate time.Time `json:"join_date"`
	Role     UserRole  `json:"role"`
}

type RequestStatus string

const (
	StatusPending  RequestStatus = "Pending"
	StatusApproved RequestStatus = "Approved"
	StatusRejected RequestStatus = "Rejected"
)

type Request struct {
	ID     int64         `json:"id"`
	UserID int64         `json:"user_id"`
	TeamID int64         `json:"team_id"`
	Status RequestStatus `json:"status"`
	SentAt time.Time     `json:"sent_at"`
}

type TechStack struct {
	ID         int64  `json:"id"`
	Technology string `json:"technology"`
	TeamID     int64  `json:"team_id"`
}

type Comment struct {
	ID        int64     `json:"id"`
	UserID    int64     `json:"user_id"`
	TeamID    int64     `json:"team_id"`
	Text      string    `json:"text"`
	ParentID  *int64    `json:"parent_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}