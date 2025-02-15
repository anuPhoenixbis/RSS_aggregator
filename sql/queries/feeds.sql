-- name: CreateFeed :one
INSERT INTO feeds(id , created_at , updated_at ,name , url , user_id)
VALUES($1 ,$2 ,$3 ,$4 ,$5 ,$6)

RETURNING *;
-- sql itself to generate the api key

-- name: GetFeeds :many
SELECT * FROM feeds ;

-- name: GetNextFeedsToFetch :many
SELECT * FROM feeds
ORDER BY last_fetched_at ASC NULLS FIRST
LIMIT $1;
-- FETCHING THEM IN ASCENDING ORDER BASED ON THE LAST FETCHED TIME

-- name: MarkFeedAsFetched :one 
UPDATE feeds
SET last_fetched_at = NOW(),
updated_at = NOW()
WHERE id = $1
RETURNING *;
-- UPDATING THE LAST FETCHED TIME OF THE FEED