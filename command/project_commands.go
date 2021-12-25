package command

import (
	"compelo/event"
	"log"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type CreateNewProjectCommand struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

func (c *Compelo) CreateNewProject(cmd CreateNewProjectCommand) Response {
	c.Lock()
	defer c.Unlock()

	// TODO: Check if name is already taken.

	guid := uuid.New().String()
	c.raise(&event.ProjectCreated{
		GUID:         guid,
		Name:         cmd.Name,
		PasswordHash: hashAndSalt([]byte(cmd.Password)),
	})
	return Response{GUID: guid}
}

func hashAndSalt(pwd []byte) []byte {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err) // TODO: Handle me
	}
	return hash
}
