package booking

type MemoryStore struct {
	bookings map[string]Booking // "A2" -> booking O(1)
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{bookings: map[string]Booking{}}
}

func (s *MemoryStore) Book(b Booking) error {
	// check if seat is available
	// if empty populate

	if _, exists := s.bookings[b.SeatID]; exists {
		return ErrSeatAlreadyBooked
	}

	s.bookings[b.SeatID] = b
	return nil
}

func (s *MemoryStore) ListBookings(movieID string) []Booking {

	var result []Booking
	for _, b := range s.bookings {
		if b.MovieID == movieID {
			result = append(result, b)
		}
	}
	return result
}
