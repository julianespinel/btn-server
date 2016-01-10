# using: https://github.com/rubenv/sql-migrate
cd ..
cd db/migrations
go get github.com/rubenv/sql-migrate/...
sql-migrate up
cd ../../
cd scripts
