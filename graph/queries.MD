```


query getRunwayByDesignator($clientID: String!, $designator: String!) {
  getRunwayByDesignator(clientID: $clientID, designator: $designator) {
    id
    designator
    contamination {
      key
      value
    }
  }
}

mutation updateRunway($clientID: String!, $input: GQRunwayInput!) {
  updateRunway(clientID: $clientID, input: $input) {
    id
    looseSand
    latestVersion
  }
}

subscription conflict($clientID: String!, $runwayID: String!){
  conflict(clientID: $clientID, runwayID: $runwayID) {
    id
    runwayID
  }
}

{
  "clientID": "1234",
  "runwayID": "61daef824d2590510ca14024"
  "input": {
    "id": "61daef824d2590510ca14024",
    "designator": "12-30",
    "looseSand": false,
    "latestVersion": 0
  }
}
```