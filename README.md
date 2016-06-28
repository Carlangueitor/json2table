# json2table
A tiny cli tool to convert json to a nice table.

Using this json: 

```json
[
  {
    "name": "Jhon Doe",
    "age": 27,
    "job": "Web Developer"
  },
  {
    "name": "Charly Román",
    "age": 24,
    "job": "Software Engineer"
  }
```

We obtain the next table:

![json2table screenshot](https://raw.githubusercontent.com/carlangueitor/json2table/master/json2table.png)

## Usage

* From file: `json2table < jsonfile.json` 
* Pipe: `curl localhost:8000 | json2table`
