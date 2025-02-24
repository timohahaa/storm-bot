package bot

const (
	createUserQuery = `
INSERT INTO core.user (
    telegram_id
    , is_admin
) VALUES (
    $1, $2
)
RETURNING 
    id
    , telegram_id
    , is_admin
`

	getUserQuery = `
SELECT 
    id
    , telegram_id
    , is_admin
FROM core.user
WHERE telegram_id = $1
    AND deleted_at IS NULL
`

	createLinkQuery = `
INSERT INTO core.link (
    user_id
    , chat_id 
    , link 
) VALUES (
    $1, $2, $3
)
RETURNING 
    id
    , user_id
    , chat_id
    , link
`
	createLinkQueryNoReturning = `
INSERT INTO core.link (
    user_id
    , chat_id 
    , link 
) VALUES (
    $1, $2, $3
)
`
	monthLinkStatsQuery = `
SELECT 
    CL.user_id
    , CL.link
FROM core.link CL
WHERE EXTRACT(MONTH FROM CL.created_at) = $1
    AND EXTRACT(YEAR FROM CL.created_at) = EXTRACT(YEAR FROM CURRENT_DATE)
    AND CL.deleted_at IS NULL
`
)
