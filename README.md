RSS Aggregator using Golang



the sql file is converted to go database file by sqlc 

for the migration : 
goose up cmd:
goose postgres postgres://postgres:Anubhav0224@localhost:8000/rssagg up

to generate the go code :
sqlc generate 
