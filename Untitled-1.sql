-- SQLite (เปิดหน้านี้ ctrl + shift + p -> SQLite: New query)
-- วิธี run -> ctrl + shift + p -> SQLite: Use database -> เลือกไฟล์ที่จะใช้เป็น database (students.db) -> ลากคลุมบรรทัดที่ต้องการ run -> คลิกขวา -> Run Selected Query
INSERT INTO students (id, name, major, gpa)
VALUES ('001', 'test1', 'CS', 3.00);

INSERT INTO students (id, name, major, gpa)
VALUES ('002', 'test2', 'CS', 3.10);

SELECT * FROM students
