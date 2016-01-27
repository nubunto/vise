# Vise Database Architecture.

The old arch:

    "users": Bucket {
        "[token]": "[information on files stored as a JSON]"
    }

Was not good for many reasons: file updating/listing requires json parsing for each user.
This could be faster.

The proposed arch:

    "files": Bucket {
        "[file-token]": Bucket {
            "user": "[token]",
            "filename": "[filename]",
            "expires-in": "3",
        },
    }

By making use of Bolt's nested buckets, we can convert File information to the databases's data.

We can retrieve the file's content in a way that is secure to the poster: there's no info whatsoever required to render the file that the user doesn't need to know (we need to expose the file name).

This is also faster. To retrieve a user's files, we need to:

  * Make a slice of type Links
  * Find the bucket "users"
  * Scan all keys, parsing the value as JSON
  * Scan all those values, adding them to the slice (there's another loop in here for each user entry)
  * Return the Links slice

The second approach focuses on files rather than clients (which makes sense, since this is a anonymous file uploader).

One of the caveats is that we build a directory for each file, although we can change that logic fairly easily

To find all files of a user in the new architecture:

  * Make a slice of string
  * Cursor all keys of the "files" bucket,
  * Get the bucket for the file token, that is, the key of the "files" bucket,
  * Get the user token
  * If it matches the one we're looking for, add it to the slice
  * At the end, return the string slice.

Not a single JSON parsing in sight, and we also dropped from O(n^2) to O(n) (in theory). That ought to be faster.
