Hackers Notes
==========================

pedro makes some notes, maybe inaccurate but pointers..

The idea of this app in present active state is to
make it work with sage accounts, and loads of spreadheets..

Idea is to remove need for windows workstations,
and replace with an active system and any interface..

It needs to be a full on system, but one that is integrable with
and this means importing stuff in a controlled way..
as well as complaining about things..

Upload a file and import it into the system, somehow
- import the sales export leger..
- and mark stuff and where it belongs..
- and then a pretend import to see stuff

This can be applied to any business..

eg
- spreadsheets cut and paste into a window
- spreasheets uplaods -
- odf and xls/xsex in go should be fun
- CSV and tsv



Overview
-------------------------
The system uses a relational db to stort key data
inc dates and planning and stuff

Uses and ORM for CRUD, but Database views for general queries..
ie all the joins ans stuff done as database views




Server Side
-------------------------

the golang stuff is a revel application
- http://revel.github.io/
- newer docs - http://revel-docs.daffodil.uk.com/manual

Golang newbies (like me)
- https://github.com/basti1302/go-lang-cheat-sheet


Structure
============

conf/
- the app config including routes

app/
- the revel application


app/controlers/
- the handlers for the requests

app/fp/
- the factory planner logic, models, structs etc

app/views/
- the main templates

app/views/staff/
- the backend templates when logged.in


===============
Database
===============
Database is initialised with the gorm..
I know, but what choices we have, maybe sqlx is better.. 
.. but later re orm and future once figure ot x/db issues

Currently..
-------------------------------
- using mysql, soon postgres, sqlite but wanted is postgis/spatical
- database is initialised in app/init.go and InitDB which it "gorm"
- https://github.com/jinzhu/gorm
- Get a "handle" to underlying go lang database with "gorm.DB.DB()" = golang native interface


Initialisation and Connection:
-------------------------------
- the db.* section in conf/app.conf
- https://github.com/daffodil/factory-planner/blob/master/conf/app.conf.skel#L15
- Its then started with revel.OnAppStart(InitDB) pointer
- https://github.com/daffodil/factory-planner/blob/master/app/init.go

Db Create
-------------
The database in created from the models..
- there are defined in app/fg/* eg accounts.Account
- these are created in app/fp/dev/db_create.go
- https://github.com/daffodil/factory-planner/blob/master/app/fp/dev/db_create.go#L23
- To create tables goto /ajax/dev/db/tables/create?drop=0
- drop=1 with drop all tables and recreate
- drop=0 will just update table
- the table names are defined in func(me MyStuff)TableName()string { return "my_stuffings" }

## Indexes ##

Creating the indexes is a pain, unless someone knows what to do with gorm..
So a hack all over the place in the DB_Index*{} methods eg
DB_IndexAccount - https://github.com/daffodil/factory-planner/blob/master/app/fp/accounts/accounts.go#L52
Also indexes return an error if they already exists
TODO find idndex and drop create

## Views  ##
- Database views are used extensively to query data..
- views are prefixed with "v_"
- Views are defined in app/fg/db_views.go
- https://github.com/daffodil/factory-planner/blob/master/app/fp/dev/db_views.go

There are generally two objects for each "table", eg Account and AccountView
 - a Account, for example is the "db table" and orm create, CRUD etc
 - but AccountView is an extender Order with more data
 - and a map on the view eg v_accounts
 
 Most queries use a view and in this example
 
 
 BEST Rules for now..
 
 I know we all want short ovjects..
 eg
 Contact.Active or Account.Active
 Therese clash on a view on contacts and accounts....
 therfore.. its ContactConActive and Account.AccActive
 as seperate Fields..
 
