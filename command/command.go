package command

import (
	"compelo/event"
	"log"

	"github.com/google/uuid"
)

// Project is a single project in compelo.
type Project struct {
	event.EventMetaData
	guid string
	name string
}

// GUID returns the project's guid.
func (p *Project) GUID() string {
	return p.guid
}

// Name returns the project's name.
func (p *Project) Name() string {
	return p.name
}

// Compelo is the root aggregate.
type Compelo struct {
	projects map[string]Project

	changes []event.Event
	version int
	store   *event.Store
}

func (p *Compelo) Projects() map[string]Project {
	return p.projects
}

func New(store *event.Store, events []event.Event) *Compelo {
	p := &Compelo{
		projects: make(map[string]Project),
		store:    store,
	}

	for _, event := range events {
		p.on(event)
	}

	return p
}

// on handles projects events on the projects aggregate.
func (c *Compelo) on(e event.Event) {
	log.Println("Command handling event ", e.GetID(), e.EventType())

	switch e := e.(type) {
	case *event.ProjectCreated:
		c.projects[e.GUID] = Project{
			guid: e.GUID,
			name: e.Name,
		}
	case *event.ProjectRenamed:
		if project, ok := c.projects[e.GUID]; ok {
			project.name = e.NewName
			c.projects[e.GUID] = project
		}
	case *event.ProjectDeleted:
		delete(c.projects, e.GUID)
	}
	c.version++
}

func (c *Compelo) raise(event event.Event) error {
	// TODO: Handle error
	c.changes = append(c.changes, event)
	c.on(event)
	c.store.StoreEvent(event)
	return nil
}

// CreateNew creates a new project.
func (c *Compelo) CreateNew(name string) error {
	// TODO: Check if name is already taken.

	c.raise(&event.ProjectCreated{
		GUID: uuid.New().String(),
		Name: name,
	})

	return nil
}

// Rename renames a project.
func (c *Compelo) Rename(guid string, newName string) error {
	// TODO: Check if name is already taken.

	c.raise(&event.ProjectRenamed{
		GUID:    guid,
		NewName: newName,
	})

	return nil
}

// Delete deletes a project.
func (c *Compelo) Delete(guid string) error {
	// TODO: Check if project exists.

	c.raise(&event.ProjectDeleted{
		GUID: guid,
	})

	return nil
}
