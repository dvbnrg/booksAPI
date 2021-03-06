DROP TABLE IF EXISTS books;

-- Create book table
CREATE TABLE IF NOT EXISTS books (
	id int,
	title varchar(30),
	author varchar(30),
	publisher varchar(30),
	publishdate varchar(10),
	rating int,
	checkedin boolean
);

INSERT INTO book.books VALUES ( 1, 'Hyperion', 'Dan Simmons', 'Doubleday', '1989', '91', TRUE);
INSERT INTO book.books VALUES ( 2, 'The Fall of Hyperion', 'Dan Simmons', 'Doubleday', '1990', '92', FALSE);
INSERT INTO book.books VALUES ( 3, 'Endymion', 'Dan Simmons', 'Doubleday', '1996', '90', TRUE);
INSERT INTO book.books VALUES ( 4, 'The Rise of Endymion', 'Dan Simmons', 'Doubleday', '1997', '90', FALSE);