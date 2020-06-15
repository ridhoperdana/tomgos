# tomgos

tomgos is a simple app to generate Golang struct based on
TOML file.

## Supported TOML

Just put a TOML file which represent your data. You can make/convert
it from JSON from this site: https://toolkit.site/format.html

Example working TOML:
```toml
[recipe]
    id = "this-is-id"
    title = "recipe title"
    description = "short description"
    cooking_time = 100
    portion = 1
    create_time = "1987-07-05T05:45:00Z"
``` 

### Nested Struct
To include nested object in the struct just use `{$struct_name}` on the
toml value.

Example:
```toml
[recipe]
    id = "this-is-id"
    title = "recipe title"
    description = "short description"
    cooking_time = 100
    portion = 1
    create_time = "1987-07-05T05:45:00Z"
    video = "{video}"

[video]
    url = "http://url.com"
``` 


## How to Use

1. Build the binary. Run `make`
2. Run `./tomgos generate` to run the App.

## TO-DO

- [x] Support nested data structure
- [ ] Support dynamic object
- [x] Support JSON descriptor