create table if not exists questions(id serial, question_id int, text varchar(255), close boolean, primary key (id, question_id));
create table if not exists answers(
    id serial,
    question_id int,
    text varchar(255),
    correct boolean,
    close boolean,
    primary key (id)
);

insert into questions(question_id, text, close) values (1, '1 + 1', false);
insert into questions(question_id, text, close) values (2, '1 + 2', false);
insert into questions(question_id, text, close) values (3, '1 + 3', false);
insert into questions(question_id, text, close) values (4, '1 + 4', false);

insert into answers(question_id, text, correct, close) values (1, '0', false, false);
insert into answers(question_id, text, correct, close) values (1, '1', false, false);
insert into answers(question_id, text, correct, close) values (1, '2', true, false);
insert into answers(question_id, text, correct, close) values (1, '3', false, false);
insert into answers(question_id, text, correct, close) values (2, '0', false, false);
insert into answers(question_id, text, correct, close) values (2, '1', false, false);
insert into answers(question_id, text, correct, close) values (2, '2', false, false);
insert into answers(question_id, text, correct, close) values (2, '3', true, false);
insert into answers(question_id, text, correct, close) values (3, '1', false, false);
insert into answers(question_id, text, correct, close) values (3, '2', false, false);
insert into answers(question_id, text, correct, close) values (3, '3', false, false);
insert into answers(question_id, text, correct, close) values (3, '4', true, false);
insert into answers(question_id, text, correct, close) values (4, '2', false, false);
insert into answers(question_id, text, correct, close) values (4, '3', false, false);
insert into answers(question_id, text, correct, close) values (4, '4', false, false);
insert into answers(question_id, text, correct, close) values (4, '5', true, false);
