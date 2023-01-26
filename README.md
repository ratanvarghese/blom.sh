# Blom, the Content Manager

`blom` is the content manager for [my blog](http://ratan.blog). Other blogs should probably avoid using it.

This version has most of the features of the [old implementation](https://github.com/ratanvarghese/blom), except for 
+ JSON Feed pagination
+ Reading from `content.html`: all content must be provided as `content.md` (note that HTML can be trivially embedded in Markdown)
+ Golang HTML templates: these have been replaced with m4 based templates which are included in the repository.

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
+ [zig 0.9.0](https://ziglang.org)
+ [mxml](https://www.msweet.org/mxml/)
+ [mon13](https://github.com/ratanvarghese/mon13)
+ [gzip](https://gnu.org/software/gzip)
+ [jq](https://www.stedolan.github.io/jq/)
+ [markdown](https://daringfireball.net/projects/markdown/)

The generated website has further dependencies. Notably `blom` does not actually handle serving the files over HTTP: a seperate server such as nginx is needed. For local development purposes, you can use a simple HTTP server such as the one included in Python:

```
python -m http.server --directory .
```

## License

    blom.sh content management system
    Copyright (C) 2023 Ratan Abraham Varghese

    This program is free software: you can redistribute it and/or modify
    it under the terms of the GNU Affero General Public License as
    published by the Free Software Foundation, either version 3 of the
    License, or (at your option) any later version.

    This program is distributed in the hope that it will be useful,
    but WITHOUT ANY WARRANTY; without even the implied warranty of
    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
    GNU Affero General Public License for more details.

    You should have received a copy of the GNU Affero General Public License
    along with this program.  If not, see <https://www.gnu.org/licenses/>.
