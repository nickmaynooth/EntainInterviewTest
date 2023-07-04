# Racing Service

The Racing Service provides functionality for retrieving races. It is implemented using Protocol Buffers (Proto) and Go.

## Proto Definitions

The Proto file `racing.proto` defines the messages and services used in the Racing Service.

### Messages

- `ListRacesRequest`: Represents the request message for listing races. It contains a `ListRacesRequestFilter` for specifying filter criteria.
- `ListRacesResponse`: Represents the response message for listing races. It contains a collection of `Race` resources.
- `ListRacesRequestFilter`: Represents the filter criteria for listing races. It includes `meeting_ids` for filtering races by meeting IDs, an optional `visible` field to filter races by visibility and a `order_by` field that allows you to sort in ascending or descending order for the advertised start time.
- `Race`: Represents a race resource with fields such as `id`, `meeting_id`, `name`, `number`, `visible`, `advertised_start_time` and `status`.

### RPC

- `ListRaces`: Retrieves a collection of all races based on the provided filter criteria.