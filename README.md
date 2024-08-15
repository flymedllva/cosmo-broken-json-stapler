
# How to reproduce

1.
```shell
cd service1
make run
```

2.
```shell
npm install -g wgc@latest
wgc router compose -i graph.yaml -o config.json
```

3. use config.json for run router

4. run query

```graphql
query Test {
    test {
        data
    }
}
```

5. view result

```json
{
  "errors": [
    {
      "message": "Bad escaped character in JSON at position 26 (line 1 column 27)",
      "stack": "SyntaxError: Bad escaped character in JSON at position 26 (line 1 column 27)"
    }
  ]
}
```