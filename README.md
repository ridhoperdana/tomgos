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

## How to Use

1. Build the binary. Run `make`
2. Run `./tomgos generate` to run the App.

## TO-DO

- Support nested data structure
- Support dynammic object