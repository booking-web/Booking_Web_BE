SELECT u.user_id, u.email, u.full_name, r.role_name, uATC.file_url FROM users u 
	LEFT JOIN roles r ON r.role_id = u.role_id
	LEFT JOIN user_attachment uATC on uATC.user_id = u.user_id
	ORDER BY user_id ASC

INSERT INTO work_location(work_location_name) VALUES ('Bệnh viện Nhi Đồng Thành phố Hồ Chí Minh'),('Bệnh viện Nhi Đồng 2 Thành phố Hồ Chí Minh') 

SELECT * FROM doctor_profile;

INSERT INTO doctor(doctor_name, doctor_summary, exp_year, edu_location, degree, description, file_url, clinic_id) VALUES
('Bill','Test Doctor', 11, 'Đại học Y dược thành phố Hồ Chí Minh', 'Thạc sĩ Y Khoa', '', '',1) RETURNING doctor_id

INSERT INTO doctor_profile VALUES (2,1,1), (2,1,2)

INSERT INTO work_location(work_location_name) VALUES ('Bệnh viện chợ rẫy'), ('Bệnh viện Nhân Dân Gia Định'),('Bệnh viện Từ Vũ')
	
SELECT * FROM doctor_profile;
SELECT * FROM doctor;
SELECT * FROM clinic;
SELECT * FROM work_location

SELECT d.doctor_name,
	d.doctor_summary,
	d.exp_year,
	cl.clinic_name,
	d.edu_location,
	d.degree,
	w.work_location_name, 
	l.language_name,
	d.file_url,
	d.description
FROM doctor d
LEFT JOIN doctor_profile dp ON dp.doctor_id = d.doctor_id
LEFT JOIN work_location w ON w.work_location_id = dp.work_location
LEFT JOIN language l ON l.language_id = dp.language
LEFT JOIN clinic cl ON cl.clinic_id = d.clinic_id

DELETE FROM doctor where doctor_id = 6