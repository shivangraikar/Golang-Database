.open kc.db
drop table if exists student;

create table if not exists student
(
rno int primary key not null,
name text not null,
marks int check (marks >= 0 and marks <= 100),
tel varchar(10),
email varchar(50),
location text default 'mumbai',
gender char(1),
dept text not null
);

insert into student values(10, 'amit', 89, '6677', null,  'mumbai', 'm' , 'comps');
insert into student values(11, 'sumit', 67, '6667', 's@gmail.com',  'thane', 'm' , 'extc');
insert into student values(12, 'ramit', 78, '6667', 'r@gmail.com',  'mumbai', 'm' , 'comps');
insert into student values(13, 'namit', 89, '9967', 'n@gmail.com',  'thane', 'm' , 'extc');
insert into student values(14, 'mit', 76, '967', 'm@gmail.com',  'mumbai', 'm' , 'comps');
insert into student values(15, 'seema', 79, '867', 's@gmail.com',  'mumbai', 'f' , 'extc');
insert into student values(16, 'reema', 78, '767', null,  'thane', 'm' , 'comps');
insert into student values(17, 'jyoti', 89, '567', 's@gmail.com',  'mumbai', 'f' , 'extc');
insert into student values(18, 'pooja', 89, '467', null,  'thane', 'f' , 'comps');
insert into student values(19, 'aarti', 91, '367', 's@gmail.com',  'mumbai', 'f' , 'extc');
insert into student values(20, 'divya', 93, '467', 's@gmail.com',  'mumbai', 'f' , 'comps');

select * from student;

.header on

select rno, name from student;

select name, marks from student;

select name || "has secured" ||  marks || 'marks' from student;
select name || "has secured" ||  marks || 'marks' as info from student;
select location from student;
select all location from student;
select distinct location from student;

--find all student from 'mumbai';
select * student from where location = 'mumbai';
select * student from where dept = 'extc';
select * student from where email is NULL;


select * from student where marks >= 40 and marks <= 80;
select * from student where marks between 40 and 80;
select * from student where rno == 13 or rno == 19 or rno == 16;
select * from student where rno in (13,19,16);

SELECT * FROM STUDENT WHERE NAME LIKE  '%IT%';
SELECT * FROM STUDENT WHERE NAME LIKE  'A%';
