.open student

drop table if exists student;

create table if not exists student
(
id int primary key not null,
name varchar(100),
salary double(8,2)
);

.tables

.schema student

insert into student values(10,'abc',4500);
select * from student;

select "************************";

insert into student values(20, 'xyz', 3000), (30,'pqr',4000);
select * from student;
select "************************";

update student set salary = salary*1.10
where id = 10;
select * from student;
select "************************";

delete from student where id = 30;
select * from student;
select "************************";

replace into student values(30,'pqr',4000);
select * from student;
select "************************";

.quit
