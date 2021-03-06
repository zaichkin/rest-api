DROP TABLE brand CASCADE;
DROP TABLE gender CASCADE;
DROP TABLE category CASCADE;
DROP TABLE product CASCADE;
DROP TABLE workspace CASCADE;
DROP TABLE worker CASCADE;


CREATE TABLE brand(
    id SERIAL PRIMARY KEY, 
    title VARCHAR(150), 
    descript TEXT, 
    img VARCHAR(100)
);

CREATE TABLE gender 
(
    id SERIAL PRIMARY KEY, 
    title VARCHAR(30)
);

CREATE TABLE category 
(
    id SERIAL PRIMARY KEY, 
    parent INT,
    name VARCHAR(50)
);

CREATE TABLE product 
(
    id SERIAL, 
    title VARCHAR(150), 
    price INT, 
    size VARCHAR(50), 
    category INT, 
    gender INT, 
    brand INT, 
    descript TEXT, 
    img VARCHAR(100),
    FOREIGN KEY (brand) REFERENCES brand(id),
    FOREIGN KEY (gender) REFERENCES gender(id),
    FOREIGN KEY (category) REFERENCES category(id)
);

CREATE TABLE workspace 
(
    id SERIAL PRIMARY KEY, 
    parent INT,
    name VARCHAR(50)
);

CREATE TABLE worker 
(
    id SERIAL,
    name VARCHAR(100),
    descript TEXT,
    workspace INT,
    FOREIGN KEY (workspace) REFERENCES workspace(id)
);

INSERT INTO gender (title) VALUES ('мужской'), ('женский'), ('унисекс');

INSERT INTO brand(title, img, descript) 
VALUES 
('BUFF', ' ', 'Компания BUFF первая в мире создала уникальную бесшовную мультифункциональную бандану. Это универсальная многоцелевая повязка, которую можно носить множеством различных способов: как шапку, шарф, бандану, балаклаву, повязку и т.д.'), 
('OCUN', ' ',  'Бренд OCUN был основан в 1998 году. Ocun разрабатывает и производит снаряжение для скалолазания: оттяжки, беседки, крэшпады, стропы, а также большое количество одежды и аксессуаров. Команда Ocun состоит только из тех людей, которые по истине любят скалолазание и горы. Целью работы является поиск правильного баланса функциональности, качества, инноваций и элегантности. Каждый год снаряжение подвергается доработкам и модификации. Скальные туфли Ocun продаются во всем мире, тестируются лучшими спортсменами. Их активно используют Евгений Овчинников, Салават Рахметов и Ольга Бибик.'), 
('RED FOX', ' ',  'Российская компания RedFox - производитель высококачественных изделий для активного отдыха и спортивной экипировки. Мы рады предложить Вам одежду и специальное снаряжение, которые позволят Вам чувствовать себя комфортно в любых самых экстремальных условиях. Уверены, что современные технологии, материалы и дизайн, прекрасное сочетание элегантности, функциональности и комфорта и широчайший выбор моделей наверняка привлекут Ваше внимание.'), 
('ARCTERYX', ' ', 'Компания Arcteryx была основана в Канаде как производитель высокотехнологичной экипировки и снаряжения для экстремальных условий. Главная цель компании Arcteryx — создание одежды, обуви и экипировки (рюкзаков и беседок), самого лучшего качества насколько это возможно на данном этапе развития мировых технологий. Исключительная функциональность, продуманный дизайн, внимание к мелочам, качество материалов, надежность и износостойкость продукции — все это об изделиях от Arcteryx.'), 
('THE NORTH FACE', ' ',  'The North Face, Inc. — компания, специализирующаяся на производстве высококачественной технологичной спортивной, горной одежды, туристического инвентаря. Продукция предназначается для альпинистов, скалолазов, туристов и просто людей, ведущих активный образ жизни.'), 
('MOUNTAIN HARDWEAR', ' ',  'Компания Mountain Hardwear была основана в 1993 году небольшой группой энтузиастов, которые любили проводить всё свободное время на природе. Они заметили, что рынок outdoor экипировки меняется, и не в лучшую сторону - снижается качество продукции. Многие компании стали ориентироваться на менее привередливых пользователей. Mountain Hardwear была создана, чтобы противостоять этой тенденции, оставаясь верной принципам надежного снаряжения, которое годами будет служить путешественникам и спортсменам.');


INSERT INTO workspace(parent, name)
VALUES (0, 'Розничный магазин'), (0, 'Интернет-магазин'), (0, 'Оптовый отдел'),
       (1, 'Директор розничного магазина'), (1, 'Администратор'), (1, 'Эксперт по снаряжению');
SELECT * FROM workspace;
