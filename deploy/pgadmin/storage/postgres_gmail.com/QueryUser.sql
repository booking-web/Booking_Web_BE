SELECT u.user_id, u.email, u.full_name, r.role_name, uATC.file_url FROM users u 
	LEFT JOIN roles r ON r.role_id = u.role_id
	LEFT JOIN user_attachment uATC on uATC.user_id = u.user_id
	ORDER BY user_id ASC

INSERT INTO work_location(work_location_name) VALUES ('Bệnh viện Nhi Đồng Thành phố Hồ Chí Minh'),('Bệnh viện Nhi Đồng 2 Thành phố Hồ Chí Minh')

SELECT * FROM doctor_profile;

INSERT INTO doctor_profile(doctor_name, doctor_summary, exp_year, clinic_id, edu_location, degree, work_location, language, description) VALUES
('Bill','Test Doctor', 11, 1, 'Đại học Y dược thành phố Hồ Chí Minh', 'Thạc sĩ Y Khoa', 1, 1, 'Test Doctor'),
('Bill','Test Doctor', 11, 1, 'Đại học Y dược thành phố Hồ Chí Minh', 'Thạc sĩ Y Khoa', 2, 2, 'Test Doctor');

SELECT * FROM doctor_profile;

SELECT d.doctor_name,
	d.doctor_summary,
	d.exp_year,
	cl.clinic_name,
	d.edu_location,
	d.degree,
	array_agg(w.work_location_name ORDER BY w.work_location_id) AS work_location, 
	array_agg(language_name ORDER BY language_id),
	d.description
FROM doctor_profile d
LEFT JOIN work_location w ON w.work_location_id = d.work_location
LEFT JOIN language l ON l.language_id = d.language
LEFT JOIN clinic cl ON cl.clinic_id = d.clinic_id
GROUP BY d.doctor_name,
	d.doctor_summary,
	d.exp_year,
	cl.clinic_name,
	d.edu_location,
	d.degree,d.description;