Hi there,

Here's the answers for the 1st challenge of Grokking Engineering's techincal marathon.

## Feedback:
- If you spot any major problems in the answers, do tell us at <admin@grokkingengineering.org>. We will take that into considerations and update the answers if needed.

****

# Q1 (10 pts): Explain how HTTPS is more secure than HTTP?

### Proposed answer
There's two main reasons:

- __Encryption__:
HTTP request are now made within a secure connection (encrypted via TLS/SSL).
Thus the data exchanged is not visible during man-in-the-middle attacks.
- __Trust establishment__:
Web browsers can trust HTTPS websites if they can provide a valid
certificate (signed by trusted certificate authorities) and the cert correctly
identifies the website (i.e., when the browser visits https://example.com, it sees a valid certificate for example.com and not some other domain)

### Marking criteria
|                   Criteria                                 |     Points      |
|:-----------------------------------------------------------|----------------:|
| Mention one (1) reason correctly                           |               5 |
| Mention two (2) reasons correctly                          |              10 |
| From the 3rd correct reason onward, no point is given      |               0 |
| From the 3rd reason onward, penalty for wrong explanation  |              -1 |

****

# Q2 (10pts): Improve query performance

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

### Proposed answer

##### Major improvement
Reading what EXPLAIN returned, you can see that the queries triggered 
a sequencial scan on the entire table to get the results.
This is has O(n) complexity (w.r.t the number of entries).
Thus, the most effective way to improve the performance is NOT doing O(n) scan,
achieved by creating indexes on the columns:
user_id, item_id, referrer_name, created_at, updated_at columns.

The benefit of indexes will really show for user_id, item_id, created_at, updated_at,
because these have high-cardinality (more different values).
Index 's benefit won't be significant for referrer_name, as it have many similar values.

Usually, the indexes are created using btree, which would yield O(log n) search complexity.


##### Other improvements
We have noticed that there're other interesting answers.
As Grokking Team are by no means PostgreSQL guru (you guys seems to know much more),
we decided to look them up and explain what we think about them:

__Changing column to use different types, i.e: "don't cast the date" or "convert the NUMERIC type to BIGINT", etc__

It's true that some datatypes / casting operations takes extra CPU cycles.
However, the performance improvement gained from that is likely much smaller
compared that of adding index.

=> These are not counted as valid answers
    
__Shard the table, i.e "shard the table by date when it's too big"__

This is a VERY good point, if you have ALREADY tried all other optimization.
In this case, however, we haven't started with the lowest hanging fruit: index.

=> These are not counted as valid answers

__Use different index type, i.e Hash index, GIST index__

Wow, we didn't know much about this, so we looked it up.
It turns out that while Hash index is indeed faster, it also have some problems,
like not being WAL-logged (during replication).
See http://www.depesz.com/2010/06/28/should-you-use-hash-index/
    
GIST index, on the other hand, provides a variety of indexing strategies based on your scenario. If you have a specialize requirement here, it will probably work
See http://www.postgresql.org/docs/current/static/indexes-types.html
    
Depending on the situation, these can be great optimization on top of the standard b-tree indexes.

=> Because does involve indexing, we will take it as valid answers
    
__Use compound index, i.e: CREATE INDEX ON user_purchases (referrer_name, user_id, created_at)__

Postgres has this rule for multi-column indexes:

A single index scan can only use query clauses that use the index's columns with operators of its operator class and are joined with AND. For example, given an index on (a, b) a query condition like WHERE a = 5 AND b = 6 could use the index, but a query like WHERE a = 5 OR b = 6 could not directly use the index.

Assuming we run the same query in the example (only with AND & the 3 columns), your answers is only correct when you mention at least those 3 columns (user_id, referrer_name, created_at) in the compound index

=> We will take it as a valid answers if you don't conflict with the PostgreSQL rule above

__Move referrer name to new table, add foreign key__

We have some answers that try to move referrer_name to a new table and put foreign key constraint on it, similar to

```
CREATE TABLE referrers (
    id SERIAL NOT NULL PRIMARY KEY,
    name VARCHAR(255)
);
CREATE TABLE user_purchases (
    user_id INT,
    item_id INT,
    referrer_id INT NOT NULL REFERENCES referrers(id),
    created_at DATE DEFAULT CURRENT_DATE,
    updated_at DATE DEFAULT CURRENT_DATE
);
EXPLAIN SELECT * FROM user_purchases WHERE user_id = 234 AND referrer_id = 123 AND created_at > '2016-11-11'::date;
                                        QUERY PLAN                                         
-------------------------------------------------------------------------------------------
 Seq Scan on user_purchases  (cost=0.00..38.53 rows=1 width=20)
   Filter: ((created_at > '2016-11-11'::date) AND (user_id = 234) AND (referrer_id = 123))
```

See the result of EXPLAIN. You are essentially running sequencial anyway, so no improvement

=> Not accepted as valid answers
 
__Not using SELECT * and specify column__

It's also true that SELECT * need to read more bytes than only chossing the fields you need, but the cost is still much smaller compared to the improvement of index

=> This is not accepted as a valid answer

__Using UNION__

One team tried to run this query

```
EXPLAIN SELECT * FROM user_purchases 
WHERE user_id=123 AND referrer_name = 'amazon' AND created_at >= date('2015-07-01') AND created_at < date('2015-08-01')
UNION
SELECT * FROM user_purchases 
WHERE user_id= 456 AND referrer_name = 'amazon' AND created_at >= date('2015-07-01') AND created_at < date('2015-08-01')
UNION
SELECT * FROM user_purchases 
WHERE user_id=789 AND referrer_name = 'amazon' AND created_at >= date('2015-07-01') AND created_at < date('2015-08-01');

 HashAggregate  (cost=37.87..37.90 rows=3 width=596)
   ->  Append  (cost=0.00..37.83 rows=3 width=596)
         ->  Seq Scan on user_purchases  (cost=0.00..12.60 rows=1 width=596)
               Filter: ((created_at >= '2015-07-01'::date) AND (created_at < '2015-08-01'::date) AND (user_id = 123::numeric) AND ((referrer_name)::text = 'amazon'::text))
         ->  Seq Scan on user_purchases user_purchases_1  (cost=0.00..12.60 rows=1 width=596)
               Filter: ((created_at >= '2015-07-01'::date) AND (created_at < '2015-08-01'::date) AND (user_id = 456::numeric) AND ((referrer_name)::text = 'amazon'::text))
         ->  Seq Scan on user_purchases user_purchases_2  (cost=0.00..12.60 rows=1 width=596)
               Filter: ((created_at >= '2015-07-01'::date) AND (created_at < '2015-08-01'::date) AND (user_id = 789::numeric) AND ((referrer_name)::text = 'amazon'::text))
```

This is, uhm, much worse, as you have to do sequencial scan 3 times

__Using OR instead of IN()__

```
EXPLAIN SELECT item_id, user_id FROM user_purchases WHERE user_id = 234 OR user_id = 123 AND referrer_name = 'abc' AND created_at > '2016-11-11'::date;
                                                                     QUERY PLAN                                                                     
----------------------------------------------------------------------------------------------------------------------------------------------------
 Seq Scan on user_purchases  (cost=0.00..12.60 rows=1 width=64)
   Filter: ((user_id = 234::numeric) OR ((user_id = 123::numeric) AND ((referrer_name)::text = 'abc'::text) AND (created_at > '2016-11-11'::date)))
(2 rows)

EXPLAIN SELECT item_id, user_id FROM user_purchases WHERE user_id IN (123, 234)AND referrer_name = 'abc' AND created_at > '2016-11-11'::date;
                                                              QUERY PLAN                                                              
--------------------------------------------------------------------------------------------------------------------------------------
 Seq Scan on user_purchases  (cost=0.00..12.28 rows=1 width=64)
   Filter: ((user_id = ANY ('{123,234}'::numeric[])) AND (created_at > '2016-11-11'::date) AND ((referrer_name)::text = 'abc'::text))
```

There is no significant differneces here

=> This is not a valid answer

### Marking criteria
|                   Criteria                                        |     Points                              |
|:------------------------------------------------------------------|----------------------------------------:|
| Mention creating index correctly and create index on the 3 columns|               9                         |
| Mention any extra reasons and it's correct                        |  +1 (cap at 10)                         |
| Mention any incorrect fact after talking about the index          |  -1 (deduce once for all wrong answers)

****

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

### Proposed answer
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

### Marking criteria
|                        Criteria                                  |   Points  |
|:-----------------------------------------------------------------|----------:|
| Proposing 1 correct solution                                     |         5 |
| Proposing more than 1 correct solutions                          |        10 |

****

## Question 4: (30pts)

Describe in as much detail as possible how an SQL query is executed, from the time it is entered into the console to the time the output is printed out.

### Proposed answers

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

A lot of answers goes into details of these steps (too many, so we can't list them out here), as well as steps before and after it. Some details are specific to certain RDBMS, too.

We will try to see if you have covered thesre 4 steps, or is missing something, and give score based on that

### Marking criteria
|                        Criteria                                  |   Points  |
|:-----------------------------------------------------------------|----------:|
| Metion all 4 steps                                               |        30 |
| Missing each steps                                               |        -2 |


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

