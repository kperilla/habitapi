# Design

## User Stories
- As a growing dev, I want the app to use CI/CD best practice to learn more about them
- As a developer, I want good documentation, diagrams, and an OpenAPI spec, to keep things well-ordered and well-designed
- As a user, I want the API to store my historical habit data, to view the accumulation
- As a user, I want to be able to log in to see my own stats/data
- As a user, I want to be able to store habit points/acts and store rewards, to be able to see my progress
- As a designer and user, I want a simple system of points that relate to habit 
actions, so that users are not discouraged or confused
- As a user/frontend, I want to be able to see my own stats/data with an endpoint, so that I can display them
- As a user, I want to be able to update or delete older entries, so that I can fix mistakes
- As a user, I want to be able to group related habits, so that they are organized
- As a designer, I want to have a recommendation for a cap on daily points to encourage a healthy progress
- As a user, I want to be able to ignore above recommendation if I want to
- As a user, I want to be able to have helpful tips, so that I can improve their habits
- As a user, I want a periodic digest of tips and help made by gen AI, to help with habits
- As a consumer, I may want a GraphQL API for historical data, to be able to get specific subsets of the data

## Specific Design

### Tech Stack
- Go REST API (TODO: look into best)
- DB
  - TODO: sql or nosql?
- CI/CD with GHA
- Hosted on GCP
- Domain: whatever is cheapest
- Kubernetes to define deployment

### API Design

REST API with the following endpoints:

- /habits - full CRUD
- /habit-group - full CRUD
- /deeds - full CRUD
- /rewards - full CRUD
- /users - full CRUD
- /data - GET
  - TODO: Need to see how best to deliver historical data; GraphQL?
