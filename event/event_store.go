package event

import (
	"encoding/binary"
	"encoding/json"
	"log"

	"go.etcd.io/bbolt"
)

type RawEvent struct {
	ID        uint64          `json:"id"`
	EventType EventType       `json:"eventType"`
	EventData json.RawMessage `json:"eventData"`
}

type EventStore struct {
	db  *bbolt.DB
	bus *EventBus

	unmarshalFns map[EventType]func([]byte) Event
}

func CreateEventStore(bus *EventBus) *EventStore {
	db, err := bbolt.Open("./compelo.db", 0666, nil)
	if err != nil {
		log.Fatal(err)
	}

	store := &EventStore{db: db, bus: bus}
	store.RegisterEvents()
	return store
}

func (s *EventStore) RegisterEvents() {
	s.unmarshalFns = make(map[EventType]func([]byte) Event)
	s.unmarshalFns[ProjectCreatedType] = (&ProjectCreated{}).UnmarshalFn()
	s.unmarshalFns[ProjectRenamedType] = (&ProjectRenamed{}).UnmarshalFn()
	s.unmarshalFns[ProjectDeletedType] = (&ProjectDeleted{}).UnmarshalFn()
}

func (s *EventStore) StoreEvent(event Event) {
	s.db.Update(func(tx *bbolt.Tx) error {
		// Open the events bucket.
		tx.CreateBucketIfNotExists([]byte("events"))
		b := tx.Bucket([]byte("events"))

		// Generate ID for the event.
		id, _ := b.NextSequence()
		event.SetID(id)

		eventData, err := json.Marshal(event)
		if err != nil {
			log.Fatal(err)
		}

		rawEvent := RawEvent{
			ID:        id,
			EventType: event.EventType(),
			EventData: eventData,
		}

		buf, err := json.Marshal(rawEvent)
		if err != nil {
			log.Fatal(err)
		}
		log.Println("Saving Event: " + string(buf))

		// Persist bytes to users bucket.
		return b.Put(itob(id), buf)
	})

	s.bus.Publish(event)
}

func (s *EventStore) LoadEvents() []Event {
	var events []Event

	s.db.View(func(tx *bbolt.Tx) error {
		// TODO: Check if bucket exists.
		b := tx.Bucket([]byte("events"))
		c := b.Cursor()

		for k, v := c.First(); k != nil; k, v = c.Next() {
			var rawEvent *RawEvent
			err := json.Unmarshal(v, &rawEvent)
			if err != nil {
				log.Fatal(err)
			}

			event := s.unmarshalFns[rawEvent.EventType](rawEvent.EventData)
			events = append(events, event)
			log.Println("Loaded event", event.EventType())
			s.bus.Publish(event)
		}
		return nil
	})
	return events
}

// itob returns an 8-byte big endian representation of v.
func itob(v uint64) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, v)
	return b
}
