insert into users (name, nick, email, password)
values
("miller", "miller00315", "miller@gmail.com", "$2a$10$c445dxy0b4Yaizy5i0oXk.k8DTjLB1x.LucLAtyDSG0Oj8PxP4iuC"),
("paulo", "paulo00315", "paulo@gmail.com", "$2a$10$c445dxy0b4Yaizy5i0oXk.k8DTjLB1x.LucLAtyDSG0Oj8PxP4iuC"),
("carlos", "carlos00315", "carlos@gmail.com", "$2a$10$c445dxy0b4Yaizy5i0oXk.k8DTjLB1x.LucLAtyDSG0Oj8PxP4iuC"),
("pedro", "pedro00315", "pedro@gmail.com", "$2a$10$c445dxy0b4Yaizy5i0oXk.k8DTjLB1x.LucLAtyDSG0Oj8PxP4iuC");

insert into followers(user_id, follower_id)
values
(1, 2),
(3, 1),
(1, 3);