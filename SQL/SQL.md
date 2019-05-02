# SQL Notes

Best practice: keywords in uppercase, identifiers in lowercase

Primary key should be unrelated to data

## SELECT

```sql
SELECT first_name FROM person;
```

* Keywords: SELECT, FROM
* Identifiers: first_name, person

Tokens:
* SELECT clause: `SELECT first_name`
* SELECT list: `first_name`
* FROM clause: `FROM person`

### Using constants in SELECT list

Valid SQL:

```sql
SELECT 'Hello' as FirstWord,'World' as SecondWord;
```

Only once column names are introduced is a FROM clause required.

### Aliasing Table Name

Good practice to always qualify  column names with table name. Use table name aliases to shorten the query.

```sql
SELECT p.first_name, p.last_name FROM person p;
```

### DISTINCT

Return only unique rows.
```sql
SELECT DISTINCT p.first_name FROM person p;
```

When used with multiple columns, returns all distinct *combinations* of those columns.
```sql
SELECT DISTINCT p.first_name, p.last_name FROM person p;
```
Will return rows with matching `first_name` if `last_name` is distinct and vice versa.

## INSERT

Only works against a single table.

```sql
INSERT INTO contacts (first_name, last_name) VALUES ('Grayson', 'Gilmore');
```

Tokens:
* INSERT clause: `INSERT INTO contacts`
* VALUES clause: `VALUES ('Grayson', 'Gilmore')`

## UPDATE

```sql
UPDATE contacts SET last_name = 'Gilmore2' WHERE id = 1;
```

Tokens:
* UPDATE clause: `UPDATE contacts SET last_name = 'Gilmore2'`
* WHERE clause: `WHERE id = 1`

Without WHERE clause every row would be updated.

## DELETE

```sql
DELETE FROM contacts WHERE id = 1;
```

Tokens:
* DELETE clause: `DELETE FROM contacts`
* WHERE clause: `WHERE id = 1`

Without WHERE clause all rows deleted.