Hi there,

Here's the first challenge for Grokking Engineering's techincal marathon.

## How to participate:

By now, we should have sent a code to you / your team.
You can go to the submission form at <http://goo.gl/forms/AybXZDgwnk> and write down your answers, together with the code to submit

## Submission rules:

- If you submit multiple times, only the the last submission count
- Bonus score is given for the first 10 submissions (based on you/your group's latest timestamp). First team to submit will receive +10 points, second team will receive +9 points, that goes all the way to the 10th team, who will receive +1 point

****

# Questions

## Question 1 (10pts)
Explain how HTTPS is more secure than HTTP?

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

## Question 3 (10pts)

What can you do to get apples for free from the following (poorly written) ecommerce service.

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
# backend
#!/usr/bin/python3

## the purchase request will get routed to this function
def purchase(request):
   quantity = int(request.POST.get('quantity', 0))
   total = float(request.POST.get('total', 0))
   request.user.balance -= total
   ship_apple_to(request.user, quantity)
   request.user.save()
   return HttpResponse('Enjoy your apples!')
```

## Question 4: (30pts)

Describe in as much detail as possible how an SQL query is executed, from the time it is entered into the console to the time the output is printed out.

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


