SELECT
    g.id as group_id,
    g.name,
    g.value,
    g.due_day,
    u.id as user_id,
    u.username
FROM groups g
LEFT JOIN user_groups ug ON g.id = ug.group_id
LEFT JOIN users u ON ug.user_id = u.id or u.id = g.owner_id
WHERE g.id = $1
