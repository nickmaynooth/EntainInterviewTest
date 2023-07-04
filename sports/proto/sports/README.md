Sports Service
The Sports Service provides functionality for retrieving sports events. It is implemented using Protocol Buffers (Proto) and Go.

Proto Definitions
The Proto file sports.proto defines the messages and services used in the Sports Service.

Messages
ListEventsRequest: Represents the request message for listing events. It contains a ListEventRequestFilter for specifying filter criteria.
ListEventsResponse: Represents the response message for listing events. It contains a collection of Event resources.
EventRequestById: Represents the request message for fetching an event by ID. It contains the id of the event to fetch.
ListEventRequestFilter: Represents the filter criteria for listing events. It includes ids for filtering events by IDs and an optional order_by field for sorting the events.
Resources
Event: Represents a sports event with fields such as id, name, number, and advertised_start_time.
RPC
ListEvents: Retrieves a collection of all events based on the provided filter criteria.
FetchEventById: Retrieves a single event based on the ID passed.

# Sports Service

The Sports Service provides functionality for retrieving sports events. It is implemented using Protocol Buffers (Proto) and Go.

## Proto Definitions

The Proto file sports.proto defines the messages and services used in the Sports Service.

### Messages

- `ListEventsRequest`: Represents the request message for listing races. It contains a `ListEventRequestFilter` for specifying filter criteria.
- `ListEventsResponse`: Represents the response message for listing races. It contains a collection of `Event` resources.
- `EventRequestById`: Represents the filter criteria for listing events. It includes `id` for filtering races by id an optional `ListEventRequestFilter` Represents the filter criteria for listing events. It includes ids for filtering events by IDs and an optional order_by field for sorting the events.
- `Event`: Represents a race resource with fields such as `id`, `name`, `number`, `advertised_start_time`.

### RPC

- `ListEvents`: Retrieves a collection of all events based on the provided filter criteria.
- `FetchEventById` Retrieves a single event based on the ID passed.