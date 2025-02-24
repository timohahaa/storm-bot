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
    $1, $2
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
    $1, $2
)
`
	monthLinkStatsQuery = `
SELECT 
    CU.telegram_id
    , CL.link
FROM core.link CL
JOIN core.user CU ON CU.id = CL.user_id
WHERE EXTRACT(MONTH FROM CL.created_at) = $1
    AND CL.deleted_at IS NULL
`
)
