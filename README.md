# Vise

Share unimportant and temporary files easily.

## Motivation

There are always those quirky files. Those files we rarely need, but when we do they are too far away.

Thanks to the power of the Internet™, we are now able to store that file with Vise.

## How does Vise works

When you first hit the homepage of Vise, you are presented with a very simple and straightforward interface.

You specify the file you whish to upload (one at a time, please) and how much that file is going to last.

Then you hit submit, and that's it. No login required. Completely anonymous.

Only caveat, though, you are bounded to see those uploaded files only in the computer (and browser) you've made that upload with.

## How does Vise works internally

Basically, Vise is composed of an HTTP endpoint that serves an index page, which is the front-end.

Vise also exposes an API, so it's easy to write programs that interface with it (documentation coming soon).

When you first submit an file, a lot of stuff is happening with Vise on the backend:

 * He noticed you're submitting a file for the first time, so he gives you a unique identifier.
 * He grabbed the file you're submitting and saved it on the server.
 * He sent you back the unique identifier so he can know where to look for your files next time.

On the front-end side, we keep this unique identifier locally using the browser storage. If you clear it, Vise will treat you like a new client.

## Vise is written in Go

Which is, for itself, quite a feature.
