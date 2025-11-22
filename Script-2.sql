CREATE DATABASE School;
-- DROP DATABASE School;
USE School;
CREATE TABLE Student(
	student_id INT AUTO_INCREMENT NOT NULL PRIMARY KEY,
	name VARCHAR(50) NOT NULL,
	group_name VARCHAR(50) NOT NULL,
	email VARCHAR(50) NOT NULL 
);

CREATE TABLE Subject(
	subject_id INT AUTO_INCREMENT NOT NULL PRIMARY KEY,
	name VARCHAR(50) NOT NULL
);

CREATE TABLE Grade(
	grade_id INT AUTO_INCREMENT NOT NULL PRIMARY KEY,
	student_id INT NOT NULL,
	subject_id INT NOT NULL,
	grade FLOAT NOT NULL,
	CONSTRAINT fk_grade_student FOREIGN KEY (student_id) REFERENCES Student(student_id),
	CONSTRAINT fk_grade_subject FOREIGN KEY (subject_id) REFERENCES Subject(subject_id)
);

INSERT INTO Student(name, group_name, email) VALUES
('Carlos', 'A', 'xd@gmail.com');

UPDATE Student
SET 
	name = "name 2",
	group_name = "A",
	email = "nuevo@gmail.com"
WHERE student_id = 2;

DELETE FROM Student WHERE student_id = 6

SELECT student_id, name, group_name, email FROM Student;

SELECT subject_id, name FROM Subject;

