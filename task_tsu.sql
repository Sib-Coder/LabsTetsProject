create table employees
(
    id         UUID primary key default gen_random_uuid(),
    name       varchar(100) not null,
    lastname   varchar(100) not null,
    surname varchar(100) null,
    gender  varchar(100) null,
    status   varchar(100) null,
    datebirth date,
    dateadded date

) ;

INSERT INTO  employees (name, lastname, surname, gender, status, datebirth, dateadded)
VALUES ( 'Daniil','Sinitsyn','Vsevolodovich','man', 'active','2001-01-11','2001-02-01');
INSERT INTO  employees (name, lastname, surname, gender, status,datebirth, dateadded)
VALUES ( 'Sib','Coder','Vsevolodovich','man', 'active','2002-02-01', '2002-02-25');
INSERT INTO  employees (name, lastname, surname, gender, status,datebirth, dateadded)
VALUES ( 'Roman','Lider','Andreevich','man', 'not active','2003-02-01', '2004-02-25');
INSERT INTO  employees (name, lastname, surname, gender, status,datebirth, dateadded)
VALUES ( 'Ksenia','Andrianova','Vadimovna','woman', 'active','2005-04-01', '2006-07-25');

--DELETE FROM employees WHERE name = 'Roman';