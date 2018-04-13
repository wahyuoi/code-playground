Get Date from Timestamp

`
select date('2018-03-20T00:00:00Z');
`

Get Hour from Timestamp using date_part

`
select date_part('hours', '2018-03-20T09:00Z'::timestamp);
`

Reference: https://www.postgresql.org/docs/current/static/functions-datetime.html
