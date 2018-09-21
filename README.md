# craigslist-global
A Golang app that allows for Craigslist searches in more than one region

This little app boy is going to eventually search through all of Craigslist instead of just limiting searches to one baby-back little bitch area. I'm sick of having to do the same searches in Philadelphia and South Jersey. Shit's annoying for sure. Why wouldn't they just allow multiple area searches? Maybe they do, tbh. Who knows. Anyway, here we go fellas...

## Resources
- This baby has lots of info on the stuff we'll need to make the urls: https://www.craigslist.org/about/bulk_posting_interface

## TODO:
- Keep category when loading new page
- Load categories into select options in alphabetical order
- Properly escape spaces in search (right now it literally replaces every space with a +) - Investigate: https://golang.org/pkg/net/url/#QueryEscape
- Some sort of way to select multiple regions (You might want to search North Jersey, Central NJ, South Jersey, and Philadelphia but not Seattle or something). Also maybe some way to say "closest N regions to zip XXXXX"
