### Notifications Service

**Scope of this service**

- Keeping the user informed about the latest events in the system
- The client polls this service on a regular interval to get the data assiciated to the given client
- This is also responsible for handling asynchonous events when the server has any update for this client
  - Tell the client to request to refetch any data
  - Tell the client to perform certain actions on the client side as instructed by the server

**Architecture**

- This is a simple REST API service
- When there is any update corresponding to any user
  - Check for the priority of data
  - if the user has an active session (check the redis session-store) -> send a message to the client to request to refetch the data
  - don't care if the session is not active (because it is not going to be used anyway)
- Email notifications are sent if
  - the selected notification channel is email
  - the user has turned on email notifications
  - the user has not turned off email notifications for this particular event/organization
