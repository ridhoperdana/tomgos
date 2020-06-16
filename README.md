# tomgos

tomgos is a simple app to generate Golang struct based on
TOML file.

## Features

- Detect field type based on value of the TOML
- Detect multiple types of field
    - string
    - int64
    - map[string]interface
    - slice of type:
        - string
        - int
        - existing/inherited struct
        - map[string]interface
- Detect other struct (inherited struct) as field type 

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
    metadata = {}
    chef = ['ridho', 'budi']
    category_ids = [1, 3, 4]
    metas = [{}]
``` 

### Nested Struct
There are 2 ways for including nested struct as field type

1. Using `{video}` bracket followed by struct's name inside of it.
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

2. Write the children key toml inside the parent key.
    Example:
    ```toml
    [recipe]
        id = "this-is-id"
        title = "recipe title"
        description = "short description"
        cooking_time = 100
        portion = 1
        create_time = "1987-07-05T05:45:00Z"
   
        [recipe.video]
           url = "http://url.com"
        
        [[recipe.thumbnail]]
           url = "http://image.com" 
    
    [video]
        url = "http://url.com"
    
    [thumbnail]
        url = "http://image.com"
    ``` 


## How to Use

1. Build the binary. Run `make`
2. Run `./tomgos generate` to run the App.

## TO-DO

- [x] Support nested data structure
- [x] Support dynamic object
- [x] Support JSON descriptor
- [x] Support slice string/int field type
- [x] Support slice nested struct field type
- [x] Support slice of dynamic object field type
- [ ] Unit Test 