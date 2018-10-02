# craigslist-global
A Golang app that allows for Craigslist searches in more than one region.
Still very much a work in progress:
- Right now it's ironically limited to one geographical region, and that is Northeast USA (this will obviously change soon to fit the scope of the project)
- There are probably many bugs.
- Right now the only thing guaranteed to work decently is cto, ctd, cta, as that was the catalyst for writing it in the first place

## Author
- Michael Crinite

## Installation
- clone the repo
- Run:
```
$ GO111MODULE=on go build main.go
```

## Use
- Run:
```
$ GIN_MODE=release go run main.go
```
- Open `0.0.0.0:8095/` in your browser
- Select a category from the drop down
- Type your search text in the search bar
- Press the `Search` button

(This project might eventually be hosted by me if I feel like it but for now you'll have to run it yourself)

## Testing
Uhh.... not yet

## TODO:
- Add photos to each post (expand posts in general to look more like the actual Craigslist counterparts)
- Some sort of way to select multiple regions (You might want to search North Jersey, Central NJ, South Jersey, and Philadelphia but not Seattle or something). Also maybe some way to say "closest N regions to zip XXXXX"
- Get comprehensive lists of categories and regions for truly G L O B A L searching
- Prettier UI: Add CSS
- Properly escape spaces in search (right now it literally replaces every space with a +) - Investigate: https://golang.org/pkg/net/url/#QueryEscape
- Get the titles in a better way. Right now it's url parsing, when it should be grabbing it from the same place that CL gets theirs, which would mean parsing the list items as the root element rather than the post href


