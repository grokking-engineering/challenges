## Update (10am)

- `RPUSH`, `SADD` needs to support multiple values
- Command names are case-insensitive, parameter and values are case-sensitive and must be all lowercase.
- If a command doesn't specify return value, please return `OK` if successful, or follow the error code.
- Updated sample tests (see link in the sample tests section)
- `SET` will always overwrite the value for that key

# Grokking Challenge Finale

In this challenge, we will attempt to build an in-memory datastore. To make things
simple, let's follow the popular in-memory database Redis and try to build a
stripped down version of it.

Your job is to build a **Ledis** (Lite Redis) datastore that supports these
data structures: **string**, **list**, **set**.

The use of ready-made databases or libraries that handle the main gist of the challenge (Redis, Riak, RocksDB, LevelDB, PostgreSQL, MySQL etc) are not allowed.

However, the use of any other libraries/framework that help with the individual components of your implementation is allowed.


# Commands
Your Ledis datastore should support the following commands (We basically follow Redis interface).
All commands are case-insensitive. All parameter names are case-sensitive and must be all lowercased. If the server encounter an error processing a command, the server will return the appropriate error code instead of the result with the specified type. If a return value is not specified and the command succeed, please return code `OK` (see below).

- General:
  - `EXPIRE key seconds`: set a timeout on a key, `seconds` is a positive integer. Return 1 if the timeout is set, 0 if key doesn't exist (5pts)
  - `TTL key`: query the timeout of a key (5pts)
  - `DEL key`: delete a key (5pts)
  - `FLUSHDB`: clear all keys (5pts)
- String:
  - `SET key value`: set a string value, always overwriting what is already saved under key (5pts)
  - `GET key`: get a string value at key (5pts)
- List: List is an ordered collection (duplicates allowed) of string values
  - `LLEN key`: return length of a list (4pts)
  - `RPUSH key value1 [value2...]`: append 1 or more values to the list, create list if not exists, return length of list after operation (4pts)
  - `LPOP key`: remove and return the first item of the list (4pts)
  - `RPOP key`: remove and return the last item of the list (4pts)
  - `LRANGE key start stop`: return a range of element from the list (zero-based, inclusive of `start` and `stop`), `start` and `stop` are non-negative integers (4pts)
- Set: Set is a unordered collection of unique string values (duplicates not
  allowed)
  - `SADD key value1 [value2...]`: add values to set stored at `key` (4pts)
  - `SCARD key`: return the number of elements of the set stored at `key` (4pts)
  - `SMEMBERS key`: return array of all members of set (4pts)
  - `SREM key value1 [value2...]`: remove values from set (4pts)
  - `SINTER [key1] [key2] [key3] ...`: set intersection among all set stored in specified keys. Return array of members of the result set. (4pts)
- Snapshot: both commands have to be implemented correctly (30pts)
    - `SAVE`: save a snapshot
    - `RESTORE`: restore from the last snapshot

You should implement the data structures yourself. The use of ready-made databases, such as (Redis, Riak, RocksDB, LevelDB, PostgreSQL, MySQL etc) are not allowed.

You are allowed to use any other libraries/framework that help with the boilerplate of your implementation. For example, using framework to handle REST API routes are allowed.

# Specifications

## Values and Their Serializations
- Strings are 1 to 10000 ASCII characters and can only contain the
  following characters `abcdefghijklmnopqrstuvwxyz`.
- Keys are string
- Integers are 32 bit signed in decimal format without extraneous zeros at
  front, e.g. `0` or `-0` represents `0`; `01` is invalid.
- Response codes are strings with specific values.
    - `OK`
    - `EMEM`: Server is out of memory
    - `EINV`: Invalid parameter value or type
    - `ECOM`: Unknown command
    - `EKTYP`: Use of command on wrong key type, including un-set key

## Communication Protocol via HTTP
The HTTP request from the client to the server is based on the following format.

    POST /ledis HTTP/1.1
    Content-Type: application/json
    Content-Length: <length>

    {"command":"command string"}


In the above request, `command string` is the command as specified in the
**Commands** section.  The HTTP verb is required to be `POST`, and the `Content-Type` and
`Content-Length` headers are required. Any other HTTP headers must be valid,
but whether the server takes them into account is up to the server.

The response from your server should have the following format.

    HTTP/1.1 200 OK
    Content-Type: application/json
    Content-Length: <length>

    {"response": <result>}

The HTTP status code is required to be 200 and the headers `Content-Type` and
`Content-Length` are required. Any other headers must be valid but whether the
client takes them into account is up to the client. The result must be
serialized as a JSON object whose sole key must be `response` and whose value
must be the server returned value as specified in the **Commands** section,
i.e. This value (shown as `<result>` above) can either be a string, array,
integer or error code (depending on the command being called).

Below are some example request/response pairs. The lines prefixed with `> ` are
sent from the client while the lines prefixed with `< ` are sent from the
server. The lines prefixed with `# ` are comments meant to clarify what is
happening.


    # the client wants to set the value of the key key1 to be babelfish
    #
    > POST /ledis HTTP/1.1
    > Content-Type: plaintext
    > Content-Length:
    >
    > {"command":"command string"}
    #
    # the server response "OK", the request was successful
    #
    < HTTP/1.1 200 OK
    < Content-Type: application/json
    < Content-Length: 15
    <
    < {"response":"OK"}

# Sample Tests

https://docs.google.com/document/d/1Oec8SHviLVMJmULawCW5eh6q4NYQBokVPNoAM-aBR_Q/edit


# Grading
The detailed grading criteria is specified in each command in the **Commands** section in parentheses.

You'll be given a brand new server running Ubuntu on Digital Ocean. You will deploy and run your program on this server (exposing through an HTTP interface).

We will run our grading program and points to your server's IP address to perform the grading.

You will be awarded points based on the functionality and scalability of the system you implemented.
