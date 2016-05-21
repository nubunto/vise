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

## The Vise API

Vise provides an API, whose routes and parameters are described:

 * `POST /api/save` note: all parameters are inside a FormData.
  * file: File to be uploaded.
  * days: Days which the file exists inside the server.
  * token: Optional. If provided, the server will store the file under this token. This token is provided for the first time this URL is called. It is not intended to be created by the user, although it could be used as such.
  * The object returned contains:
    * ok: boolean
    * message: string
    * user_token: string
  * Note that it is the client's responsibility to ensure that subsequent calls to `/api/save` provide the token. Otherwise, the server will think that this is a completely new user and assign a different token.
 * `GET /api/links` returns information about all the files. This is usually used for debugging purposes, but it is exposed anyway.
 * `GET /api/:token/links` returns links for all of the token's associated files.
 * There are also private information available via a JSON api. If you're interested, check out how to configure your vise in the file `main.go`.

## Vise is written in Go

Which is, for itself, quite a feature.
