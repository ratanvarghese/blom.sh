### Introduction

When this blog began, it was pretty much for my own amusement. I
picked technologies specifically because they seemed fun to play
with. Features were limited based on how easy I thought they would be
to implement myself. The layout was optimized not for reading experience, but for showing off SVG resizing and CSS transparency. I
started writing prose, but switched to listicles and even bullet
points simply because it was easy.

Was this all good? Bad? I had no idea. My own opinion reversed itself
every few days. As for other people, their feedback was limited to
direct messages and a log file that nobody read.

All the while, the wider internet was changing. So over the past few
weeks, I decided to enhance this site on several fronts.

### Server-Side Changes You Hopefully Will Not Notice

The first change was an overhaul of the content management system. The incumbent was a system I wrote myself in Golang:
[`blom`](https://github.com/ratanvarghese/blom). Golang is a fun
little language for certain tasks, but its unique strengths did
not shine through in `blom`.

I pretty much rewrote `blom` such that each task was conducted
using the most suitable language: file management was handled in bash
scripts, templating was handled using
[m4](https://www.gnu.org/software/m4/m4.html) and various other tools
were used as appropriate. This is obviously not something a reader
would notice directly, but it made adding new features and testing new layouts significantly easier.

Another enabling change was switching from a [simple NodeJS-based server](https://www.npmjs.com/package/http-server) to
[nginx](https://nginx.org/en/). In retrospect this was not necessary,
but it made some of the next changes easier to configure.

In the past, every time a file on this site was accessed, a line
would be added to a log file which I never, ever checked. However,
now I've installed a [web log analyzer](https://goaccess.io/) which
presents all the data with pretty colors. I used to feel like I was
shouting into the void, but now when I write I can be assured that
plenty of search engine bots are reading.

### HTTPS

Even before this blog existed, the world has been moving to using
[HTTPS by default for all websites](https://security.googleblog.com/2016/09/moving-towards-more-secure-web.html). `ratan.blog` has been very late to the party ... still, better late than never!

### Comments

I added a comment system! That [merits its own article](/adding-comments).

### Styling

[Here](/blog-enhancements/attachments/old.html) is what the blog used to look like.

I actually attempted [disable all the browser's default styling](https://meyerweb.com/eric/tools/css/reset/) and individually
write new styles for every possible HTML element. This was both
exhausting and probably a waste of time, but reading the
[MDN list of HTML elements](https://developer.mozilla.org/en-US/docs/Web/HTML/Element) was very educational. For one thing, I finally know
what a description list is (also called a definition list). As for
the CSS reset, I ended up just using the budget version:

	* {
	margin:0;
	padding:0;
	}

After eliminating the Lufthansa-esque color scheme, the next casualty
was the logo.

![old logo](/blog-enhancements/attachments/calculated_asymmetry.svg)

What was the point of that logo? Was it totally arbitrary? Well, I
don't actually remember. All I know is, the filename was
`calculated_asymmetry.svg`, but the logo was symmetrical using an axis
45&deg; from the horizontal. What a scandal!

![new logo](/blog-enhancements/attachments/calculated_asymmetry_new.svg)

The meaning of the new logo is that right-angled trapezoids are
underrated.

### Short URLs

So you know that feeling when some task seems so easy that there
would be no reason to test it, but then something goes wrong that
could have been easily detected while testing? Yeah,
that's what happened after I registered [r3n.me](http://r3n.me). I tried to
make a script that created alternate URLs for every existing article.
Unfortunately it created alternate URLs for those alternate URLs
every time it was run. Eventually there were 3 or 4 paths to every
article and each path had a corresponding entry in the RSS, Atom and
JSON Feeds.

If you were subscribed to any of those feeds and saw "re-runs" of old
articles, sorry about that.

On the plus side, there are short URLs for every article. They are
recorded in the [archive](/archive).

### Content

The third item will shock you!

=Description lists=
    I edited a previous article to use a description list, because description lists are cool and underrated (much like right-angled trapezoids).
=Listicles=
    I haven't removed the old listicles, but will probably not write more.
=More Frequent Updates=
    Just kidding.
