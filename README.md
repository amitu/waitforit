# waitforit

TCP proxy that waits for a TCP server to start if its not running currently.

## Why?

My django server, that listens on 8000, auto restarts when I make any changes to
any file in my django project. Which is awesome.

Problem: Some projects I work on have grown particularly big, and restarting
takes 5-6 seconds. This happens to be longer than my "edit -> Cmd-S -> Tab to
browser -> Cmd-R". By this often server would not have restarted, and browser
shows me connection error.

So I keep reloading, Cmd-R Cmd-R like a maniac. Or tab back and forth betwen
Safari and PyCharm to see if django has started.

And this project is the solution.

Now I do not open localhost:8000, where my django server is running, but
localhost:8001, where waitforit is running. This guy takes a connection and
keeps retrying till django server comes up.

If you want to change these setting you can use flags.

Thats all folks. Since this project is bugfree, no need to raise any issues and
don't expect any more commits.
