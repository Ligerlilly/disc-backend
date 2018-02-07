### Curls
curl -X POST -H "Content-Type:application/graphql"  -d '{
  "query": "query GetAllDiscs {discs { title artist year id __typename}}",
  "variables": {}
}' http://localhost:8090