package event

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"log"

	"go.etcd.io/bbolt"
)

const eventsBucket = "events"

type RawEvent struct {
	ID        uint64          `json:"id"`
	EventType EventType       `json:"eventType"`
	EventData json.RawMessage `json:"eventData"`
}

type Store struct {
	db  *bbolt.DB
	bus *Bus

	unmarshalFns map[EventType]func([]byte) Event
}

func NewStore(bus *Bus, path string) *Store {
	db, err := bbolt.Open(path, 0666, nil)
	if err != nil {
		log.Fatal(err)
	}

	store := &Store{db: db, bus: bus}
	return store
}

func (s *Store) StoreEvent(event Event) error {
	err := s.db.Update(func(tx *bbolt.Tx) error {
		// Open the events bucket.
		tx.CreateBucketIfNotExists([]byte(eventsBucket))
		b := tx.Bucket([]byte(eventsBucket))

		// Generate ID for the event.
		id, _ := b.NextSequence()
		event.SetID(id)

		eventData, err := json.Marshal(event)
		if err != nil {
			return fmt.Errorf("marshal of event data failed: %w", err)
		}

		rawEvent := RawEvent{
			ID:        id,
			EventType: event.EventType(),
			EventData: eventData,
		}

		buf, err := json.Marshal(rawEvent)
		if err != nil {
			return fmt.Errorf("marshal of raw event failed: %w", err)
		}
		log.Println("Saving Event: " + string(buf))

		// Persist bytes to users bucket.
		return b.Put(itob(id), buf)
	})

	if err == nil {
		s.bus.Publish(event)
	}

	return err
}

func (s *Store) LoadEvents() ([]Event, error) {
	var events []Event
	err := s.db.View(func(tx *bbolt.Tx) error {
		tx.CreateBucketIfNotExists([]byte(eventsBucket))
		b := tx.Bucket([]byte(eventsBucket))
		c := b.Cursor()

		for k, v := c.First(); k != nil; k, v = c.Next() {
			var rawEvent *RawEvent
			err := json.Unmarshal(v, &rawEvent)
			if err != nil {
				return fmt.Errorf("unmarshal of raw event failed: %w", err)
			}

			if event, err := rawEvent.EventType.Unmarshal(rawEvent.EventData); err == nil {
				events = append(events, event)
			} else {
				return fmt.Errorf("unmarshal of event data failed: %w", err)
			}
		}
		return nil
	})

	if err == nil {
		for _, e := range events {
			s.bus.Publish(e)
		}
	}

	return events, err
}

// itob returns an 8-byte big endian representation of v.
func itob(v uint64) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, v)
	return b
}
