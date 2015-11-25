Hi there,

Here's the answers for the 1st challenge of Grokking Engineering's techincal marathon.

## Feedback:
- If you spot any major problems in the answers, do tell us at <admin@grokkingengineering.org>. We will take that into considerations and update the answers if needed.

****

# Questions

## Question 1 (10pts)
Explain how HTTPS is more secure than HTTP?

#### Proposed Answer
There's two main reasons:

- *Encryption*
HTTP request are now made within a secure connection (encrypted via TLS/SSL).
Thus the data exchanged is not visible during man-in-the-middle attacks.

- *Trust establishment*
Web browsers know how to trust HTTPS websites if they can provide a valid
certificate (signed by trusted certificate authorities) and the cert correctly
identifies the website (e.g., when the browser visits "https://example.com", 
the received certificate is properly for "example.com" and not some other
entity)

#### Marking criteria
|                   Criteria                                 |     Points      |
|:-----------------------------------------------------------|----------------:|
| Mention one (1) reason correctly                           |               5 |
| Mention two (2) reasons correctly                          |              10 |
| From the 3rd correct reason onward, no point is given      |               0 |
| From the 3rd reason onward, penalty for wrong explanation  |              -1 |

****

## Question 2 (10pts)

You have a table with the following structure

```sql
-- postgresql-compatible SQL
-- each entry in user_purchases indicate that the user have bought an item
-- referrer_name have a few values: ‘amazon’, ‘facebook’, ‘friends-of-grokking-engineering’
CREATE TABLE user_purchases (
  user_id NUMERIC,
  item_id NUMERIC,
  referrer_name VARCHAR(255),
  created_at TIMESTAMP WITH TIME ZONE DEFAULT now(),
  updated_at TIMESTAMP WITH TIME ZONE DEFAULT now()
);
```

You make a lot of query to find items that a user purchased, similar to this

```sql
SELECT * FROM user_purchases WHERE referrer_name = 'amazon' AND user_id IN (123, 456, 789) AND created_at >= date('2015-07-01') AND created_at < date('2015-08-01');
```
The query is slow, and running `EXPLAIN` returns:

```sql
EXPLAIN SELECT * FROM user_purchases WHERE referrer_name = 'amazon' AND user_id IN (123, 456, 789) AND created_at >= date('2015-07-01') AND created_at < date('2015-08-01')

Seq Scan on user_purchases  (cost=0.00..3482332.76 rows=1 width=596)
  Filter: ((created_at >= '2015-07-01'::date) AND (created_at < '2015-08-01'::date) AND ((referrer_name)::text = 'amazon'::text) AND (user_id = ANY ('{123,456,789}'::numeric[])))
```

What would you do to improve the performance of the query (make it run faster), __independent__ of the server's infrastructure (CPU, memory, disk, etc)?

#### Answers

##### Major improvement
Reading what EXPLAIN returned, you can see that the queries triggered 
a sequencial scan on the entire table to get the results.
This is has O(n) complexity (w.r.t the number of entries).

Thus, imporving the performance means NOT doing O(N) scan 
is to create indexes on the columns:
user_id, item_id, referrer_name, created_at, updated_at columns.

The benefit of indexes will really show for user_id, item_id, created_at, updated_at,
because these have high-cardinality (more different values).
Index 's benefit won't be significant for referrer_name,
as it have many similar values.

Usually, the indexes are created using btree,
which would yield O(log N) search complexity.


##### Other improvements / valid points
We have noticed that there're other interesting answers.

As Grokking Team are by no means PostgreSQL guru (you guys seems to know much more),
we decided to look them up and explain what we think about them:

*Type casting solution, i.e: "not casting the date" or "convert the NUMERIC type to BIGINT"*

These are valid points, but as some datatypes / casting operations takes extra 
CPU cycles. However, their performance cost is likely smaller compared to the index.
These are not counted as valid answers
    
*Sharing solution, i.e "shard the table by date when it's too big"*
This is also a VERY good point, if you have already tried all other optimization.
In this case, however, we haven't started with the lowest hanging fruit: index.
These are not counted as valid answers

*Use different index type, i.e Hash index, GIST index)*

Wow, we didn't know much about this, so we looked it.
It turns out that while Hash index is indeed faster, it also have some problems,
like not being WAL-logged (during replication).
See http://www.depesz.com/2010/06/28/should-you-use-hash-index/
    
GIST index, on the other hand, provides a variety of indexing strategies based on your scenario.
See http://www.postgresql.org/docs/current/static/indexes-types.html
    
These are wonderful optimization on top of the standard b-tree indexes. Hence, we count themas valid answers
    
*Use compound index*

*Move referrer name to new table, add foreign key*

*Not using SELECT*

*Using UNION*

*Using OR instead of IN*


## Question 3 (10pts)

What can you do to get apples for free from the following (poorly written)
ecommerce service.

```html
<!DOCTYPE html>
<!-- the front-end -->
<html>
  <body>
    <form action="/purchase" method="POST">
      <p>Apples (1.5$ per kg)</p>
      <input id="quantity" type="text" name="quantity" />
      <input id="total" type="hidden" name="total" />
      <input type="submit" name="submit" id="submit" value="Submit" />
      </form>
    <script>
      document.getElementById('quantity').addEventListener('change', function() {
        document.getElementById('total').value = parseInt(document.getElementById('quantity').value, 10) * 1.5;
      });
    </script>
  </body>
</html>
```

```python
#!/usr/bin/python3
# backend

## the purchase request will get routed to this function
def purchase(request):
   quantity = int(request.POST.get('quantity', 0))
   total = float(request.POST.get('total', 0))
   request.user.balance -= total
   ship_apple_to(request.user, quantity)
   request.user.save()
   return HttpResponse('Enjoy your apples!')
```

#### Proposed Answer
If you read the code carefully, you'll notice that the amount deducted from the
user's balance is equal to the value of the `total` key in the POST request
body. The value of `total` is taken from the value of the hidden form element
with ID `total`, which is update every time the user change the amount of apples
in the `quantity` input element. Furthermore, there are no checks if the value
of `total` matches the `quantity` ordered on the server. Armed with this
knowledge, it is straightforward to construct a solution.

Posible solutions:

- Modify the value of `total` form element before clicking on `Submit` via the
JavaScript console available in most modern browsers. This solution is easier as
it is very probable that the server checks for user session via a secure cookie.
The request sent from the browser automatically includes this cookie.

- Alternatively, you can construct the request yourself with the help of the
`cURL` program or any HTTP library available to your favorite language. The
request's body has to include the `total` and `quantity` keys with desired
values (did you think of assigning a negative value to `total`? Free money!).
However, you'll also need to include the session secure cookie which you can
sniff for with networking tools like `ngrep` or the more convenient Network
tab in most modern browser's Developer Tools.

#### Marking criteria
|                        Criteria                                  |   Points  |
|:-----------------------------------------------------------------|----------:|
| Proposing 1 correct solution                                     |         5 |
| Proposing more than 1 correct solutions                          |        10 |

****

## Question 4: (30pts)

Describe in as much detail as possible how an SQL query is executed, from the time it is entered into the console to the time the output is printed out.

#### Answers

```
In general, there will be four steps to execute a query:

# Parse:
Here our SQL statement is converted into a structure.
For example, in Postgres, this will be a parse tree (in C)

# Analyze & rewrite:
Here, the SQL server analyzes and rewrites our query,
optimizing and simplifying it using a series of complex algorithms

# Plan:
Next, a plan for finding our data will be generated.
Like an obsessive compulsive person who won’t leave home
without every suitcase packed perfectly,
the RDBMS doesn’t run our query until it has a plan

# Execute:
Finally, our query will be executed.

The details will be different based on the RDBMS's implementation. 

For example, how Postgres is talked in great details here:
http://patshaughnessy.net/2014/10/13/following-a-select-statement-through-postgres-internals

And this is for MySQL:
http://adminlinux.blogspot.sg/2009/06/mysql-query-execution-basics.html
```

****

## Question 5: (30pts)

It's almost new year, so it's time for a challenge involving the number 2016.

Your task is to write a program that prints "2016" (without the quotes), without using any of the characters in `0123456789` (that is, no digits character is allowed).

Your program must also be independent of any external variables such as the system's date, or a seed of a random sequence.
Also, you can only use ASCII characters in your program (Unicode would be too easy, hehe). Any language in which digits are valid tokens is allowed. The shortest source code (in bytes) wins.

#### Examples of invalid programs

```ruby
#!/usr/bin/ruby
puts Time.now().year + 1
# not valid because you used system time
```

```ruby
#!/usr/bin/ruby
puts "ߠ".ord    # ߠ   is the Unicode character at #07e0 (2016 in decimal)
# not valid because you used Unicode
```

```bash
#!/bin/bash
curl -sX GET 'http://myurl.com/question5' # where you put the number 2016
# not valid because you are dependent on an external factor
```

#### Grading:
We will run your code on at <http://ideone.com>, so you can use any languages supported there.

Thus, when you submit, __please add a comment at the top saying what langauge you are using__.

Because as the source code gets smaller, it’s harder to weave off bytes, the score is given in logarithmic-ish scale as followed:

| Condition                                       | Score |
|-------------------------------------------------|-------|
| Not able to produce correct result              | 0     |
| Correct, source code >= 125 bytes               | 5     |
| Correct, 35 <= source code’s length < 125 bytes | 10    |
| Correct, 30 <= source code's length < 35 bytes  | 11    |
| Correct, 25 <= source code's length < 30 bytes  | 13    |
| Correct, 20 <= source code's length < 25 bytes  | 16    |
| Correct, 15 <= source code's length < 20 bytes  | 20    |
| Correct, 10 <= source code's length < 15 bytes  | 25    |
| Correct, <= 10 bytes                            | 30    |

### Answers

```
Since Unicode, system time & external sources are not allowed,
you will need to do some aritmethic.

One simple way is to take a character, convert it to ASCII code
and use that value.

This is easy, to do, but usually results in longer code

For example,

#!/usr/env/ruby
------------------------------------------------------------------
| ('bab'.to_i ?\r.ord) + ('q'.to_i ?\e.ord) - ('a'.to_i ?\r.ord) |
------------------------------------------------------------------

Explaination:
`?\r`     returns the character "\r" in ASCII (Carriage Return).
`?\r.ord`     returns 13 (ASCII code of "\r")
`'BC1'.to_i ?\r.ord`  convert "BC1" (2016 in base 13) to decimal
```

