-- employeesとdepartmentsテーブルを作成する
--  

create table departments (
    dept_id int primary key,
    dept_name text not null
);

create table employees (
    emp_id serial primary key,
    emp_name text not null,
    emp_dept_id int not null references departments(dept_id)
        on update cascade
        on delete restrict,
);

