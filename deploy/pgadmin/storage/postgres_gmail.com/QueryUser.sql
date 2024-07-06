SELECT u.user_id, u.email, u.full_name, r.role_name, uATC.file_url FROM users u 
	LEFT JOIN roles r ON r.role_id = u.role_id
	LEFT JOIN user_attachment uATC on uATC.user_id = u.user_id
	ORDER BY user_id ASC