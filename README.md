## Database

Switch into the postgres user or simply fire up the postgresql client with the postgres role.
```
$ psql -U postgres
```

Create a role for the app
```
create role dude_expenses with createdb superuser login password '123456';
```

You need to create the new role as a superuser in order to be able to enable specific database extensions. If you have an existing user you have to alter the role with
```
alter role dude_expenses superuser;
```

Using the newly created role, create a database for development.
```
$ createdb -U dude_expenses -h localhost dude_expenses_development;
```

All app migrations live in `db/migrate`. You should execute them using the configuration you created above.
```
$ psql -h localhost -U dude_expenses -d dude_expenses_development
> create table...
```

## API

### Authorization

In order to access /expenses resources you must provide a valid JWT token in the Authorization header.

You can obtain a valid token by authenticating first. Use the authentication endpoint and provide a valid email and password.
```
$ curl -i -X POST https://dudeexpenses.com/users/authenticate -d '{ "email": "homer@simpsons.com", "password": "123456" }'
$ curl -i -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiMTMifQ.Ss0Cr47S6-V99R0E1dy_Tax8yPICiYS0SPDX4HLMD-2" https://dudeexpenses.com/expenses
```