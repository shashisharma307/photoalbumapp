create database photo_album;
use photo_album;


CREATE TABLE `user` (
  `id` int(6)
  `user_id` int(6)  unsigned NOT NULL AUTO_INCREMENT,
  `fname` varchar(20) NOT NULL,
  `lastname` varchar(20) NOT NULL,
  `email` varchar(50) NOT NULL,
  `password` varchar(50) NOT NULL,
   `age` smallint NOT NULL,
   `contact` double NOT NULL,
	`gender` varchar(5) not null,
	PRIMARY KEY (`user_id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=latin1;

select * from user;

CREATE TABLE `album`(
	`album_id` int(6) unsigned NOT NULL AUTO_INCREMENT PRIMARY KEY,
     `album_name` varchar(25) not null,
     `created_at` datetime not null,
     `update_at` datetime not null,
	  `user_id` int(6) unsigned not null,
	  FOREIGN KEY (user_id) REFERENCES user(user_id)
)ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=latin1;


CREATE TABLE `images`(
	`image_id` int(6) unsigned NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `image_name` varchar(50) NOT NULL,
    `image_path` varchar(100),
	`album_id` int(6) unsigned not null,
    `user_id` int(6) unsigned not null,
FOREIGN KEY (album_id) REFERENCES album(album_id),	
FOREIGN KEY (user_id) REFERENCES user(user_id)
    
);



INSERT INTO 
	user(fname, lastname, email, password, age, contact, gender)
values
	('shashi','Sharma', 'shashi@test.com', 'ramgarh@1234', 30, 7838660434, 'm'),
	('nehil','Sharma', 'nehil@test.com', 'ramgarh@1234', 17, 9717120750, 'm'),
	('shailja','Sharma', 'shailja@test.com', 'ramgarh@1234', 16, 8851647173, 'f'),
	('namita','Sharma', 'namita@test.com', 'ramgarh@1234', 35, 9636493150, 'f'),
	('cp','Sharma', 'cp@test.com', 'ramgarh@1234', 42, 9414372448, 'm');
commit;


INSERT INTO 
	album(album_name, created_at, update_at, user_id)
values
	('vaishnodevi', now(), now(), 6),
    ('gujarda', now(), now(), 7),
    ('kuraj', now(), now(), 8),
    ('bhilwara', now(), now(), 9),
    ('badrinath', now(), now(), 10);	

select * from album;

INSERT INTO 
	images(image_name, image_path, album_id, user_id)
values
	('pic1', 'na', 11, 6),
    ('pic2', 'na', 11, 6),
    ('pic1', 'na', 12, 7),
    ('pic2', 'na', 12, 7),
    ('pic3', 'na', 12, 7),
	('pic1', 'na', 13, 8),
	('pic2', 'na', 13, 8),
	('pic3', 'na', 13, 8),
	('pic1', 'na', 14, 9),
	('pic2', 'na', 14, 9),
	('pic1', 'na', 15, 10),
	('pic2', 'na', 15, 6),
	('pic3', 'na', 15, 6);

select * from images;


drop table images;
drop table album;
drop table user;

