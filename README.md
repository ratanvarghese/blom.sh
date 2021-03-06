# Blom, the Content Manager

`blom` is the content manager for [my blog](http://ratan.blog). Other blogs should probably avoid using it.

This version has most of the features of the [old implementation](https://github.com/ratanvarghese/blom), except for 
+ JSON Feed pagination
+ Reading from `content.html`: all content must be provided as `content.md` (note that HTML can be trivially embedded in Markdown)
+ Golang HTML templates: these have been replaced with m4 based templates which are included in the repository.
+ Parallelism: this was never a performance-sensitive task

Additionally, several features have been added:
+ Automatic GZIP compression
+ Including the article modification date in the article's page
+ Including links to the plain `content.md` and `content.html` in the article's page
+ Independence from the current working directory
+ Apply CSS classes to an entire article
+ A `_COMMENT_SECTION` below the article text

## Dependencies

+ a standard \*nix environment
  + [m4](https://www.gnu.org/software/m4/m4.html) is not quite an everyday tool, but is part of POSIX
+ [bash](https://tiswww.case.edu/php/chet/bash/bashtop.html)
+ [golang](https://golang.org), specifically the following packages:
  + the standard packages
  + [gorilla/feeds](https://github.com/gorilla/feeds)
  + [tqtime](https://github.com/ratanvarghese/tqtime)
  + **NOTE: The Golang portions have not been updated since 2019**
+ [gzip](https://gnu.org/software/gzip)
+ [jq](https://www.stedolan.github.io/jq/)
+ [markdown](https://daringfireball.net/projects/markdown/)

The generated website has further dependencies. Notably `blom` does not actually handle serving the files over HTTP.
