package api

import (
	"context"
	"time"
)

// Identifier aliases the data type used for identifying entities within the system.
type Identifier []byte

// A User of the service. Could be a student or an instructor.
type User struct {
	ID           Identifier `json:"id"`
	Email        string     `json:"email" db:"email"`
	FirstName    string     `json:"firstName" db:"first_name"`
	LastName     string     `json:"lastName" db:"last_name"`
	IsInstructor bool       `json:"isInstructor"`
	CreatedAt    time.Time  `json:"createdAt" db:"created_at"`
	UpdatedAt    time.Time  `json:"updatedAt" db:"updated_at"`
}

// A Course that a user can subscribe to.
type Course struct {
	ID           Identifier `json:"id"`
	Name         string     `json:"name"`
	Description  string     `json:"description"`
	Category     string     `json:"category"`
	Tags         []string   `json:"tags"`
	InstructorID Identifier `json:"instructorId"`
	LessonOrder  []string   `json:"lesson_order"`
	ThreadID     Identifier `json:"threadId"`
	CreatedAt    time.Time  `json:"createdAt"`
	UpdatedAt    time.Time  `json:"updatedAt"`
}

// A Lesson is an individual part of a course.
type Lesson struct {
	ID          Identifier `json:"id"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	ContentAddr string     `json:"contentAddr"`
	CourseID    Identifier `json:"courseId"`
	ThreadID    Identifier `json:"threadId"`
	CreatedAt   time.Time  `json:"createdAt"`
	UpdatedAt   time.Time  `json:"updatedAt"`
}

// A Subscription for upgraded access.
type Subscription struct {
	ID     Identifier `json:"id"`
	UserID Identifier `json:"userId"`
	// TODO (erik): More to come once I figure out payment structure.
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

// A Purchase that links a user to owned courses.
type Purchase struct {
	ID        Identifier `json:"id"`
	UserID    Identifier `json:"userId"`
	CourseID  Identifier `json:"courseId"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
}

// A Comment is a post or response of some kind from a user of the platform.
type Comment struct {
	ID        Identifier `json:"id"`
	AuthorID  Identifier `json:"authorId"`
	ParentID  Identifier `json:"parentId"` // parent comment if replying
	Content   string     `json:"content"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
}

// A Thread is just a special comment that is the first of chain of replies. It can also have a title.
type Thread struct {
	Comment
	Title string `json:"title"`
}

// Creds are an email/password combo.
type Creds struct {
	Email    string `db:"email"`
	Password string `db:"password"`
}

// An Authenticator checks the validity of a set of credentials.
type Authenticator interface {
	Authenticate(ctx context.Context, creds Creds) (string, error)
}

// A UserService handles interactions with User objects.
type UserService interface {
	CreateUser(ctx context.Context, user User, password string) (Identifier, error)
	GetUser(ctx context.Context, id string) (User, error)
}

// A CourseService handles interactions with Course objects.
type CourseService interface {
	CreateCourse(ctx context.Context, course Course) (Identifier, error)
	GetCourse(ctx context.Context, id Identifier) (Course, error)
	ListCourses(ctx context.Context, category string) ([]Course, error)
}

type LessonService interface {
	CreateLesson(ctx context.Context, lesson Lesson) (Identifier, error)
	GetLesson(ctx context.Context, id Identifier)
}

// An API is the aggregation of all becomebitwise services.
type API interface {
	Authenticator() Authenticator
	CourseService() CourseService
	UserService() UserService
}
