package command

import (
	"compelo/event"
	"fmt"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type CreateNewProjectCommand struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

func (c *Compelo) CreateNewProject(cmd CreateNewProjectCommand) (Response, error) {
	c.Lock()
	defer c.Unlock()

	if err := c.checkUniqueConstraint(cmd.Name); err != nil {
		return Response{}, fmt.Errorf("project name is taken: %w", err)
	}

	guid := uuid.New().String()
	hash, err := hashAndSalt([]byte(cmd.Password))
	if err != nil {
		return Response{}, fmt.Errorf("cannot create project due to invalid password: %w", err)
	}

	c.raise(&event.ProjectCreated{
		GUID:         guid,
		Name:         cmd.Name,
		PasswordHash: hash,
	})

	return Response{GUID: guid}, nil
}

func hashAndSalt(pwd []byte) ([]byte, error) {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		return nil, fmt.Errorf("password hashing failed: %w", err)
	}
	return hash, nil
}
