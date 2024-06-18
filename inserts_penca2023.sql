SET FOREIGN_KEY_CHECKS = 0;

-- Inserción de valores en Utils
INSERT INTO Utils (hours_until_match, exact_match_points, correct_result_match_points, champion_points, sub_champion_points) VALUES
(1, 4, 2, 10, 5);

-- Inserción de Roles de Usuario
INSERT INTO UserRoles (role_name, description) VALUES
('Admin', 'Administrator role'),
('Student', 'Student role');

INSERT INTO penca_ucu.`User` (document_id,email,last_name,first_name,major,password,role_id) VALUES
('0','admin@ucu.edu.uy','Cartalla','nico','Computer Engineering','$2a$09$/WRgKd9Dp8Wn26kT3c5wuOY7IVhJ.zOmmVz35xgXEa38X9dXSngte',1),
('11111111','nacho@ucu.edu.uy','Valle','Nacho','Computer Engineering','$2a$09$DYhB4Iwyi0srkIbNhUejPut7gufDfRcAbmOMkbMXQR9G9VtmRbIsm',2),
('22222222','nico@ucu.edu.uy','Cartalla','nico','Computer Engineering','$2a$09$66XFqeNscsW0y54eVQCxI.y1iCwxrovZ5cAzxsVOmqVuM3a/p/MQ6',2),
('33333333','nacho1@ucu.edu.uy','Valle','Nacho1','Computer Engineering','$2a$09$ZqlL1mN4HcXW.WwVRChJ6uh0FdPIurVN./XQUYKWbI.Y5KZL4ZPpO',2),
('44444444','nacho2@ucu.edu.uy','Valle','Nacho2','Computer Engineering','$2a$09$ez3vaeegjvDXKch1b7KDyeLaGATMjnxEypiq9Q6DmFWbqiKw4w86u',2),
('55555555','nacho3@ucu.edu.uy','Valle','Nacho3','Computer Engineering','$2a$09$uccrUSWhXoXgvfFqNPfMSOGIpPBnaeAy5OOlqdqjFBa7d2zkxCROe',2),
('66666666','anto@ucu.edu.uy','Mescia','Anto','Computer Engineering','$2a$09$f7OjKsyztCmu6itTAiioOeGHM5VpsKfRMLVLo/Nb1Sj3IpkmMuKXa',2);
INSERT INTO penca_ucu.`User` (document_id,email,last_name,first_name,major,password,role_id) VALUES
('49283227','nachovalle@ucu.edu.uy','Valle','Nacho','Computer Engineering','$2a$09$Pc2ZS0kFr4GEyxJzcXywUerapc2QiG67KiGz2TGIuiNXMeLRcR5tK',2);

-- Inserción de Campeonatos
INSERT INTO Championships (name, year, country, championship_type) VALUES
('Copa América 2023', 2023, 'Various', 'International');

-- Inserción de Grupos de la Copa América 2023
INSERT INTO GroupStages (group_s_name, description, championship_id) VALUES
('Grupo A', 'Grupo A de la Copa América 2023', 1),
('Grupo B', 'Grupo B de la Copa América 2023', 1),
('Grupo C', 'Grupo C de la Copa América 2023', 1),
('Grupo D', 'Grupo D de la Copa América 2023', 1);

-- Inserción de Etapas para la Copa América
INSERT INTO Stages (stage_name, description) VALUES
('Fase de Grupos', 'Primera etapa de la Copa América 2023'),
('Cuartos de Final', 'Segunda etapa de la Copa América 2023'),
('Semifinal', 'Tercera etapa de la Copa América 2023'),
('Final', 'Etapa final de la Copa América 2023');

-- Inserción de Equipos
INSERT INTO Teams (name, url_logo, description) VALUES
('Argentina', 'http://nicolascartalla.duckdns.org:65190/bd2-back/media/teams/Argentina.svg', 'Selección de Argentina'),
('Bolivia', 'http://nicolascartalla.duckdns.org:65190/bd2-back/media/teams/Bolivia.svg', 'Selección de Bolivia'),
('Brasil', 'http://nicolascartalla.duckdns.org:65190/bd2-back/media/teams/Brazil.svg', 'Selección de Brasil'),
('Chile', 'http://nicolascartalla.duckdns.org:65190/bd2-back/media/teams/Chile.svg', 'Selección de Chile'),
('Colombia', 'http://nicolascartalla.duckdns.org:65190/bd2-back/media/teams/Colombia.svg', 'Selección de Colombia'),
('Ecuador', 'http://nicolascartalla.duckdns.org:65190/bd2-back/media/teams/Ecuador.svg', 'Selección de Ecuador'),
('Paraguay', 'http://nicolascartalla.duckdns.org:65190/bd2-back/media/teams/Paraguay.svg', 'Selección de Paraguay'),
('Peru', 'http://nicolascartalla.duckdns.org:65190/bd2-back/media/teams/Peru.svg', 'Selección de Peru'),
('Uruguay', 'http://nicolascartalla.duckdns.org:65190/bd2-back/media/teams/Uruguay.svg', 'Selección de Uruguay'),
('Venezuela', 'http://nicolascartalla.duckdns.org:65190/bd2-back/media/teams/Venezuela.svg', 'Selección de Venezuela'),
('Canada', 'http://nicolascartalla.duckdns.org:65190/bd2-back/media/teams/Canada.svg', 'Selección de Canada'),
('Costa Rica', 'http://nicolascartalla.duckdns.org:65190/bd2-back/media/teams/Costa_Rica.svg', 'Selección de Costa Rica'),
('United States', 'http://nicolascartalla.duckdns.org:65190/bd2-back/media/teams/United_States.svg', 'Selección de United States'),
('Jamaica', 'http://nicolascartalla.duckdns.org:65190/bd2-back/media/teams/Jamaica.svg', 'Selección de Jamaica'),
('Mexico', 'http://nicolascartalla.duckdns.org:65190/bd2-back/media/teams/Mexico.svg', 'Selección de Mexico'),
('Panama', 'http://nicolascartalla.duckdns.org:65190/bd2-back/media/teams/Panama.svg', 'Selección de Panama');

-- Inserción de Equipos en la Copa América 2023
INSERT INTO Teams_Championships (team_id, championship_id, group_s_id) VALUES
-- Grupo A
(1, 1, 1), -- Argentina en Grupo A
(8, 1, 1), -- Peru en Grupo A
(4, 1, 1), -- Chile en Grupo A
(11, 1, 1), -- Canada en Grupo A
-- Grupo B
(15, 1, 2), -- Mexico en Grupo B
(6, 1, 2), -- Ecuador en Grupo B
(10, 1, 2), -- Venezuela en Grupo B
(14, 1, 2), -- Jamaica en Grupo B
-- Grupo C
(13, 1, 3), -- United States en Grupo C
(9, 1, 3), -- Uruguay en Grupo C
(16, 1, 3), -- Panama en Grupo C
(2, 1, 3), -- Bolivia en Grupo C
-- Grupo D
(3, 1, 4), -- Brazil en Grupo D
(5, 1, 4), -- Colombia en Grupo D
(7, 1, 4), -- Paraguay en Grupo D
(12, 1, 4); -- Costa Rica en Grupo D

-- Inserción de Partidos de la fase de grupos de la Copa América 2023
INSERT INTO GameMatch (match_date, team_local_id, team_visitor_id, goals_local, goals_visitor, championship_id, stage_id, group_s_id) VALUES
('2023-06-20 19:00:00', 1, 11, 2, 0, 1, 1, 1), -- Argentina vs Canada
('2023-06-21 19:00:00', 8, 4, 1, 1, 1, 1, 1), -- Peru vs Chile
('2023-06-22 17:00:00', 6, 10, 3, 2, 1, 1, 2), -- Ecuador vs Venezuela
('2023-06-22 20:00:00', 15, 14, 0, 0, 1, 1, 2), -- Mexico vs Jamaica
('2023-06-23 17:00:00', 13, 2, 1, 0, 1, 1, 3), -- United States vs Bolivia
('2023-06-23 20:00:00', 9, 16, 2, 1, 1, 1, 3), -- Uruguay vs Panama
('2023-06-24 17:00:00', 5, 7, 0, 0, 1, 1, 4), -- Colombia vs Paraguay
('2023-06-24 18:00:00', 3, 12, 3, 1, 1, 1, 4), -- Brazil vs Costa Rica
('2023-06-25 17:00:00', 8, 11, 1, 2, 1, 1, 1), -- Peru vs Canada
('2023-06-25 20:00:00', 4, 1, 2, 2, 1, 1, 1), -- Chile vs Argentina
('2023-06-26 17:00:00', 6, 14, 1, 0, 1, 1, 2), -- Ecuador vs Jamaica
('2023-06-26 20:00:00', 10, 15, 2, 1, 1, 1, 2), -- Venezuela vs Mexico
('2023-06-27 17:00:00', 16, 13, 0, 1, 1, 1, 3), -- Panama vs United States
('2023-06-27 20:00:00', 2, 9, 3, 3, 1, 1, 3), -- Bolivia vs Uruguay
('2023-06-28 17:00:00', 7, 3, 1, 2, 1, 1, 4), -- Paraguay vs Brazil
('2023-06-28 20:00:00', 12, 5, 0, 0, 1, 1, 4), -- Costa Rica vs Colombia
('2023-06-29 19:00:00', 1, 8, 1, 3, 1, 1, 1), -- Argentina vs Peru
('2023-06-29 19:00:00', 11, 4, 2, 2, 1, 1, 1), -- Canada vs Chile
('2023-06-30 19:00:00', 14, 10, 0, 1, 1, 1, 2), -- Jamaica vs Venezuela
('2023-06-30 19:00:00', 15, 6, 1, 1, 1, 1, 2), -- Mexico vs Ecuador
('2023-07-01 20:00:00', 2, 16, 2, 0, 1, 1, 3), -- Bolivia vs Panama
('2023-07-01 20:00:00', 13, 9, 0, 1, 1, 1, 3), -- United States vs Uruguay
('2023-07-02 20:00:00', 12, 7, 1, 3, 1, 1, 4), -- Costa Rica vs Paraguay
('2023-07-02 20:00:00', 3, 5, 2, 1, 1, 1, 4); -- Brazil vs Colombia


-- Partidos de cuartos de final
INSERT INTO GameMatch (match_date, team_local_id, team_visitor_id, goals_local, goals_visitor, championship_id, stage_id, group_s_id) VALUES
('2023-07-05 17:00:00', 1, 14, 3, 1, 1, 2, NULL), -- Argentina vs Jamaica
('2023-07-05 20:00:00', 15, 11, 2, 1, 1, 2, NULL), -- Mexico vs Canada
('2023-07-06 17:00:00', 9, 5, 2, 1, 1, 2, NULL), -- Uruguay vs Colombia
('2023-07-06 20:00:00', 3, 13, 2, 0, 1, 2, NULL); -- Brazil vs United States

-- Partidos de semifinales
INSERT INTO GameMatch (match_date, team_local_id, team_visitor_id, goals_local, goals_visitor, championship_id, stage_id, group_s_id) VALUES
('2023-07-09 17:00:00', 1, 9, 1, 2, 1, 3, NULL), -- Argentina vs Uruguay
('2023-07-09 20:00:00', 15, 3, 1, 2, 1, 3, NULL); -- Mexico vs Brazil


-- Final
INSERT INTO GameMatch (match_date, team_local_id, team_visitor_id, goals_local, goals_visitor, championship_id, stage_id, group_s_id) VALUES
('2023-07-12 17:00:00', 9, 3, 2, 1, 1, 4, NULL); -- Uruguay vs Brazil



-- Inserción de predicciones de partidos
INSERT INTO Predictions (goals_local, goals_visitor, document_id, match_id) VALUES
(2, 0, '11111111', 1),  -- Nacho predice Argentina vs Canada
(1, 1, '11111111', 2),  -- Nacho predice Peru vs Chile
(3, 2, '11111111', 3),  -- Nacho predice Ecuador vs Venezuela
(0, 0, '11111111', 4),  -- Nacho predice Mexico vs Jamaica
(1, 0, '11111111', 5),  -- Nacho predice United States vs Bolivia
(2, 1, '11111111', 6),  -- Nacho predice Uruguay vs Panama
(0, 0, '11111111', 7),  -- Nacho predice Colombia vs Paraguay
(3, 1, '11111111', 8),  -- Nacho predice Brazil vs Costa Rica
(1, 2, '11111111', 9),  -- Nacho predice Peru vs Canada
(2, 2, '11111111', 10), -- Nacho predice Chile vs Argentina
(1, 0, '11111111', 11), -- Nacho predice Ecuador vs Jamaica
(2, 1, '11111111', 12), -- Nacho predice Venezuela vs Mexico
(0, 1, '11111111', 13), -- Nacho predice Panama vs United States
(3, 3, '11111111', 14), -- Nacho predice Bolivia vs Uruguay
(1, 2, '11111111', 15), -- Nacho predice Paraguay vs Brazil
(0, 0, '11111111', 16), -- Nacho predice Costa Rica vs Colombia
(1, 3, '11111111', 17), -- Nacho predice Argentina vs Peru
(2, 2, '11111111', 18), -- Nacho predice Canada vs Chile
(0, 1, '11111111', 19), -- Nacho predice Jamaica vs Venezuela
(1, 1, '11111111', 20), -- Nacho predice Mexico vs Ecuador
(2, 0, '11111111', 21), -- Nacho predice Bolivia vs Panama
(0, 1, '11111111', 22), -- Nacho predice United States vs Uruguay
(1, 3, '11111111', 23), -- Nacho predice Costa Rica vs Paraguay
(0, 2, '11111111', 24), -- Nacho predice Brazil vs Colombia

(0, 2, '22222222', 1),  -- nico@ucu.edu.uy predice Argentina vs Canada
(2, 1, '22222222', 2),  -- nico@ucu.edu.uy predice Peru vs Chile
(1, 1, '22222222', 3),  -- nico@ucu.edu.uy predice Ecuador vs Venezuela
(0, 3, '22222222', 4),  -- nico@ucu.edu.uy predice Mexico vs Jamaica
(1, 0, '22222222', 5),  -- nico@ucu.edu.uy predice United States vs Bolivia
(1, 3, '22222222', 6),  -- nico@ucu.edu.uy predice Uruguay vs Panama
(2, 1, '22222222', 7),  -- nico@ucu.edu.uy predice Colombia vs Paraguay
(0, 2, '22222222', 8),  -- nico@ucu.edu.uy predice Brazil vs Costa Rica
(3, 1, '22222222', 9),  -- nico@ucu.edu.uy predice Peru vs Canada
(1, 2, '22222222', 10), -- nico@ucu.edu.uy predice Chile vs Argentina
(0, 0, '22222222', 11), -- nico@ucu.edu.uy predice Ecuador vs Jamaica
(2, 2, '22222222', 12), -- nico@ucu.edu.uy predice Venezuela vs Mexico
(1, 3, '22222222', 13), -- nico@ucu.edu.uy predice Panama vs United States
(3, 0, '22222222', 14), -- nico@ucu.edu.uy predice Bolivia vs Uruguay
(1, 1, '22222222', 15), -- nico@ucu.edu.uy predice Paraguay vs Brazil
(0, 3, '22222222', 16), -- nico@ucu.edu.uy predice Costa Rica vs Colombia
(2, 1, '22222222', 17), -- nico@ucu.edu.uy predice Argentina vs Peru
(0, 2, '22222222', 18), -- nico@ucu.edu.uy predice Canada vs Chile
(1, 0, '22222222', 19), -- nico@ucu.edu.uy predice Jamaica vs Venezuela
(2, 1, '22222222', 20), -- nico@ucu.edu.uy predice Mexico vs Ecuador
(3, 1, '22222222', 21), -- nico@ucu.edu.uy predice Bolivia vs Panama
(0, 2, '22222222', 22), -- nico@ucu.edu.uy predice United States vs Uruguay
(1, 0, '22222222', 23), -- nico@ucu.edu.uy predice Costa Rica vs Paraguay
(2, 1, '22222222', 24), -- nico@ucu.edu.uy predice Brazil vs Colombia

(0, 2, '33333333', 1),  -- Nacho1 predice Argentina vs Canada
(2, 1, '33333333', 2),  -- Nacho1 predice Peru vs Chile
(1, 1, '33333333', 3),  -- Nacho1 predice Ecuador vs Venezuela
(0, 3, '33333333', 4),  -- Nacho1 predice Mexico vs Jamaica
(1, 0, '33333333', 5),  -- Nacho1 predice United States vs Bolivia
(1, 3, '33333333', 6),  -- Nacho1 predice Uruguay vs Panama
(2, 1, '33333333', 7),  -- Nacho1 predice Colombia vs Paraguay
(0, 2, '33333333', 8),  -- Nacho1 predice Brazil vs Costa Rica
(3, 1, '33333333', 9),  -- Nacho1 predice Peru vs Canada
(1, 2, '33333333', 10), -- Nacho1 predice Chile vs Argentina
(0, 0, '33333333', 11), -- Nacho1 predice Ecuador vs Jamaica
(2, 2, '33333333', 12), -- Nacho1 predice Venezuela vs Mexico
(1, 3, '33333333', 13), -- Nacho1 predice Panama vs United States
(3, 0, '33333333', 14), -- Nacho1 predice Bolivia vs Uruguay
(1, 1, '33333333', 15), -- Nacho1 predice Paraguay vs Brazil
(0, 3, '33333333', 16), -- Nacho1 predice Costa Rica vs Colombia
(2, 1, '33333333', 17), -- Nacho1 predice Argentina vs Peru
(0, 2, '33333333', 18), -- Nacho1 predice Canada vs Chile
(1, 0, '33333333', 19), -- Nacho1 predice Jamaica vs Venezuela
(2, 1, '33333333', 20), -- Nacho1 predice Mexico vs Ecuador
(3, 1, '33333333', 21), -- Nacho1 predice Bolivia vs Panama
(0, 2, '33333333', 22), -- Nacho1 predice United States vs Uruguay
(1, 0, '33333333', 23), -- Nacho1 predice Costa Rica vs Paraguay
(2, 1, '33333333', 24), -- Nacho1 predice Brazil vs Colombia

(3, 1, '44444444', 1),  -- Nacho2 predice Argentina vs Canada
(1, 0, '44444444', 2),  -- Nacho2 predice Peru vs Chile
(1, 3, '44444444', 3),  -- Nacho2 predice Ecuador vs Venezuela
(0, 2, '44444444', 4),  -- Nacho2 predice Mexico vs Jamaica
(1, 1, '44444444', 5),  -- Nacho2 predice United States vs Bolivia
(2, 0, '44444444', 6),  -- Nacho2 predice Uruguay vs Panama
(0, 1, '44444444', 7),  -- Nacho2 predice Colombia vs Paraguay
(1, 3, '44444444', 8),  -- Nacho2 predice Brazil vs Costa Rica
(0, 2, '44444444', 9),  -- Nacho2 predice Peru vs Canada
(2, 1, '44444444', 10), -- Nacho2 predice Chile vs Argentina
(1, 1, '44444444', 11), -- Nacho2 predice Ecuador vs Jamaica
(0, 3, '44444444', 12), -- Nacho2 predice Venezuela vs Mexico
(2, 1, '44444444', 13), -- Nacho2 predice Panama vs United States
(0, 2, '44444444', 14), -- Nacho2 predice Bolivia vs Uruguay
(1, 3, '44444444', 15), -- Nacho2 predice Paraguay vs Brazil
(2, 2, '44444444', 16), -- Nacho2 predice Costa Rica vs Colombia
(1, 0, '44444444', 17), -- Nacho2 predice Argentina vs Peru
(2, 1, '44444444', 18), -- Nacho2 predice Canada vs Chile
(0, 0, '44444444', 19), -- Nacho2 predice Jamaica vs Venezuela
(3, 2, '44444444', 20), -- Nacho2 predice Mexico vs Ecuador
(1, 1, '44444444', 21), -- Nacho2 predice Bolivia vs Panama
(2, 3, '44444444', 22), -- Nacho2 predice United States vs Uruguay
(0, 1, '44444444', 23), -- Nacho2 predice Costa Rica vs Paraguay
(1, 0, '44444444', 24), -- Nacho2 predice Brazil vs Colombia

(2, 3, '55555555', 1),  -- Nacho3 predice Argentina vs Canada
(0, 0, '55555555', 2),  -- Nacho3 predice Peru vs Chile
(2, 1, '55555555', 3),  -- Nacho3 predice Ecuador vs Venezuela
(3, 2, '55555555', 4),  -- Nacho3 predice Mexico vs Jamaica
(1, 2, '55555555', 5),  -- Nacho3 predice United States vs Bolivia
(2, 1, '55555555', 6),  -- Nacho3 predice Uruguay vs Panama
(0, 3, '55555555', 7),  -- Nacho3 predice Colombia vs Paraguay
(1, 1, '55555555', 8),  -- Nacho3 predice Brazil vs Costa Rica
(3, 2, '55555555', 9),  -- Nacho3 predice Peru vs Canada
(1, 0, '55555555', 10), -- Nacho3 predice Chile vs Argentina
(0, 1, '55555555', 11), -- Nacho3 predice Ecuador vs Jamaica
(2, 0, '55555555', 12), -- Nacho3 predice Venezuela vs Mexico
(3, 1, '55555555', 13), -- Nacho3 predice Panama vs United States
(2, 1, '55555555', 14), -- Nacho3 predice Bolivia vs Uruguay
(0, 0, '55555555', 15), -- Nacho3 predice Paraguay vs Brazil
(3, 2, '55555555', 16), -- Nacho3 predice Costa Rica vs Colombia
(1, 1, '55555555', 17), -- Nacho3 predice Argentina vs Peru
(0, 2, '55555555', 18), -- Nacho3 predice Canada vs Chile
(1, 1, '55555555', 19), -- Nacho3 predice Jamaica vs Venezuela
(2, 1, '55555555', 20), -- Nacho3 predice Mexico vs Ecuador
(0, 3, '55555555', 21), -- Nacho3 predice Bolivia vs Panama
(1, 2, '55555555', 22), -- Nacho3 predice United States vs Uruguay
(0, 0, '55555555', 23), -- Nacho3 predice Costa Rica vs Paraguay
(3, 1, '55555555', 24), -- Nacho3 predice Brazil vs Colombia

(1, 1, '66666666', 1),  -- Anto predice Argentina vs Canada
(0, 2, '66666666', 2),  -- Anto predice Peru vs Chile
(2, 0, '66666666', 3),  -- Anto predice Ecuador vs Venezuela
(1, 3, '66666666', 4),  -- Anto predice Mexico vs Jamaica
(2, 1, '66666666', 5),  -- Anto predice United States vs Bolivia
(1, 1, '66666666', 6),  -- Anto predice Uruguay vs Panama
(3, 2, '66666666', 7),  -- Anto predice Colombia vs Paraguay
(0, 1, '66666666', 8),  -- Anto predice Brazil vs Costa Rica
(2, 2, '66666666', 9),  -- Anto predice Peru vs Canada
(3, 1, '66666666', 10), -- Anto predice Chile vs Argentina
(1, 2, '66666666', 11), -- Anto predice Ecuador vs Jamaica
(0, 1, '66666666', 12), -- Anto predice Venezuela vs Mexico
(3, 0, '66666666', 13), -- Anto predice Panama vs United States
(2, 1, '66666666', 14), -- Anto predice Bolivia vs Uruguay
(0, 3, '66666666', 15), -- Anto predice Paraguay vs Brazil
(1, 0, '66666666', 16), -- Anto predice Costa Rica vs Colombia
(2, 1, '66666666', 17), -- Anto predice Argentina vs Peru
(3, 3, '66666666', 18), -- Anto predice Canada vs Chile
(1, 1, '66666666', 19), -- Anto predice Jamaica vs Venezuela
(2, 2, '66666666', 20), -- Anto predice Mexico vs Ecuador
(0, 3, '66666666', 21), -- Anto predice Bolivia vs Panama
(1, 2, '66666666', 22), -- Anto predice United States vs Uruguay
(3, 1, '66666666', 23), -- Anto predice Costa Rica vs Paraguay
(1, 0, '66666666', 24), -- Anto predice Brazil vs Colombia

(1, 1, '49283227', 1),  -- Nacho Valle predice Argentina vs Canada
(0, 2, '49283227', 2),  -- Nacho Valle predice Peru vs Chile
(2, 0, '49283227', 3),  -- Nacho Valle predice Ecuador vs Venezuela
(1, 3, '49283227', 4),  -- Nacho Valle predice Mexico vs Jamaica
(2, 1, '49283227', 5),  -- Nacho Valle predice United States vs Bolivia
(1, 1, '49283227', 6),  -- Nacho Valle predice Uruguay vs Panama
(3, 2, '49283227', 7),  -- Nacho Valle predice Colombia vs Paraguay
(0, 1, '49283227', 8),  -- Nacho Valle predice Brazil vs Costa Rica
(2, 2, '49283227', 9),  -- Nacho Valle predice Peru vs Canada
(3, 1, '49283227', 10), -- Nacho Valle predice Chile vs Argentina
(1, 2, '49283227', 11), -- Nacho Valle predice Ecuador vs Jamaica
(0, 1, '49283227', 12), -- Nacho Valle predice Venezuela vs Mexico
(3, 0, '49283227', 13), -- Nacho Valle predice Panama vs United States
(2, 1, '49283227', 14), -- Nacho Valle predice Bolivia vs Uruguay
(0, 3, '49283227', 15), -- Nacho Valle predice Paraguay vs Brazil
(1, 0, '49283227', 16), -- Nacho Valle predice Costa Rica vs Colombia
(2, 1, '49283227', 17), -- Nacho Valle predice Argentina vs Peru
(3, 3, '49283227', 18), -- Nacho Valle predice Canada vs Chile
(1, 1, '49283227', 19), -- Nacho Valle predice Jamaica vs Venezuela
(2, 2, '49283227', 20), -- Nacho Valle predice Mexico vs Ecuador
(0, 3, '49283227', 21), -- Nacho Valle predice Bolivia vs Panama
(1, 2, '49283227', 22), -- Nacho Valle predice United States vs Uruguay
(3, 1, '49283227', 23), -- Nacho Valle predice Costa Rica vs Paraguay
(1, 0, '49283227', 24); -- Nacho Valle predice Brazil vs Colombia

-- Inserción de predicciones de campeonatos
INSERT INTO PredictionsChampionships (champion, subchampion, document_id, championship_id) VALUES
(3, 1, '11111111', 1),   -- nico@ucu.edu.uy predice Brazil campeón y Argentina subcampeón
(3, 1, '22222222', 1),         -- Nacho predice Brazil campeón y Argentina subcampeón
(6, 4, '33333333', 1),     -- nacho2@ucu.edu.uy predice Ecuador campeón y Chile subcampeón
(9, 13, '44444444', 1),   -- nacho3@ucu.edu.uy predice Uruguay campeón y United States subcampeón
(5, 12, '55555555', 1),  -- anto@ucu.edu.uy predice Colombia campeón y Costa Rica subcampeón
(3, 9, '49283227', 1);  -- nachovalle@ucu.edu.uy predice Brazil campeón y Uruguay subcampeón

-- Resultados reales de los partidos de la fase de grupos de la Copa América 2023
UPDATE GameMatch SET goals_local = 2, goals_visitor = 0 WHERE match_id = 1;  -- Argentina vs Canada
UPDATE GameMatch SET goals_local = 1, goals_visitor = 1 WHERE match_id = 2;  -- Peru vs Chile
UPDATE GameMatch SET goals_local = 3, goals_visitor = 2 WHERE match_id = 3;  -- Ecuador vs Venezuela
UPDATE GameMatch SET goals_local = 0, goals_visitor = 0 WHERE match_id = 4;  -- Mexico vs Jamaica
UPDATE GameMatch SET goals_local = 1, goals_visitor = 0 WHERE match_id = 5;  -- United States vs Bolivia
UPDATE GameMatch SET goals_local = 2, goals_visitor = 1 WHERE match_id = 6;  -- Uruguay vs Panama
UPDATE GameMatch SET goals_local = 0, goals_visitor = 0 WHERE match_id = 7;  -- Colombia vs Paraguay
UPDATE GameMatch SET goals_local = 3, goals_visitor = 1 WHERE match_id = 8;  -- Brazil vs Costa Rica
UPDATE GameMatch SET goals_local = 1, goals_visitor = 2 WHERE match_id = 9;  -- Peru vs Canada
UPDATE GameMatch SET goals_local = 2, goals_visitor = 2 WHERE match_id = 10; -- Chile vs Argentina
UPDATE GameMatch SET goals_local = 1, goals_visitor = 0 WHERE match_id = 11; -- Ecuador vs Jamaica
UPDATE GameMatch SET goals_local = 2, goals_visitor = 1 WHERE match_id = 12; -- Venezuela vs Mexico
UPDATE GameMatch SET goals_local = 0, goals_visitor = 1 WHERE match_id = 13; -- Panama vs United States
UPDATE GameMatch SET goals_local = 3, goals_visitor = 3 WHERE match_id = 14; -- Bolivia vs Uruguay
UPDATE GameMatch SET goals_local = 1, goals_visitor = 2 WHERE match_id = 15; -- Paraguay vs Brazil
UPDATE GameMatch SET goals_local = 0, goals_visitor = 0 WHERE match_id = 16; -- Costa Rica vs Colombia
UPDATE GameMatch SET goals_local = 1, goals_visitor = 3 WHERE match_id = 17; -- Argentina vs Peru
UPDATE GameMatch SET goals_local = 2, goals_visitor = 2 WHERE match_id = 18; -- Canada vs Chile
UPDATE GameMatch SET goals_local = 0, goals_visitor = 1 WHERE match_id = 19; -- Jamaica vs Venezuela
UPDATE GameMatch SET goals_local = 1, goals_visitor = 1 WHERE match_id = 20; -- Mexico vs Ecuador
UPDATE GameMatch SET goals_local = 2, goals_visitor = 0 WHERE match_id = 21; -- Bolivia vs Panama
UPDATE GameMatch SET goals_local = 0, goals_visitor = 1 WHERE match_id = 22; -- United States vs Uruguay
UPDATE GameMatch SET goals_local = 1, goals_visitor = 3 WHERE match_id = 23; -- Costa Rica vs Paraguay
UPDATE GameMatch SET goals_local = 0, goals_visitor = 2 WHERE match_id = 24; -- Brazil vs Colombia

-- Inserción de puntajes de usuarios por partidos (Scores)
INSERT INTO Scores (document_id, match_id, points) VALUES
-- Usuario 1: Nacho Valle
('11111111', 1, 4),    -- Exacto: Argentina 2-0 Canada
('11111111', 2, 2),    -- Resultado: Peru 1-1 Chile (empate)
('11111111', 3, 4),    -- Exacto: Ecuador 3-2 Venezuela
('11111111', 4, 4),    -- Exacto: Mexico 0-0 Jamaica
('11111111', 5, 4),    -- Exacto: United States 1-0 Bolivia
('11111111', 6, 4),    -- Exacto: Uruguay 2-1 Panama
('11111111', 7, 4),    -- Exacto: Colombia 0-0 Paraguay
('11111111', 8, 4),    -- Exacto: Brazil 3-1 Costa Rica
('11111111', 9, 0),    -- Incorrecto: Peru 1-2 Canada
('11111111', 10, 2),   -- Resultado: Chile 2-2 Argentina (empate)
('11111111', 11, 4),   -- Exacto: Ecuador 1-0 Jamaica
('11111111', 12, 4),   -- Exacto: Venezuela 2-1 Mexico
('11111111', 13, 2),   -- Resultado: Panama 0-1 United States
('11111111', 14, 2),   -- Resultado: Bolivia 3-3 Uruguay (empate)
('11111111', 15, 2),   -- Resultado: Paraguay 1-2 Brazil
('11111111', 16, 4),   -- Exacto: Costa Rica 0-0 Colombia
('11111111', 17, 2),   -- Resultado: Argentina 1-3 Peru
('11111111', 18, 2),   -- Resultado: Canada 2-2 Chile (empate)
('11111111', 19, 2),   -- Resultado: Jamaica 0-1 Venezuela
('11111111', 20, 2),   -- Resultado: Mexico 1-1 Ecuador (empate)
('11111111', 21, 4),   -- Exacto: Bolivia 2-0 Panama
('11111111', 23, 2),   -- Resultado: Costa Rica 1-3 Paraguay
('11111111', 24, 4),   -- Exacto: Brazil 0-2 Colombia

-- Usuario 2: Nico
('22222222', 1, 0),    -- Incorrecto: Argentina 0-2 Canada
('22222222', 2, 2),    -- Resultado: Peru 1-1 Chile (empate)
('22222222', 3, 2),    -- Resultado: Ecuador 3-2 Venezuela
('22222222', 4, 4),    -- Exacto: Mexico 0-0 Jamaica
('22222222', 5, 4),    -- Exacto: United States 1-0 Bolivia
('22222222', 6, 0),    -- Incorrecto: Uruguay 1-3 Panama
('22222222', 7, 2),    -- Resultado: Colombia 0-0 Paraguay (empate)
('22222222', 8, 0),    -- Incorrecto: Brazil 0-2 Costa Rica
('22222222', 9, 0),    -- Incorrecto: Peru 3-1 Canada
('22222222', 10, 2),   -- Resultado: Chile 2-2 Argentina (empate)
('22222222', 11, 4),   -- Exacto: Ecuador 1-0 Jamaica
('22222222', 12, 2),   -- Resultado: Venezuela 2-1 Mexico
('22222222', 13, 0),   -- Incorrecto: Panama 1-3 United States
('22222222', 14, 0),   -- Incorrecto: Bolivia 3-0 Uruguay
('22222222', 15, 2),   -- Resultado: Paraguay 1-1 Brazil (empate)
('22222222', 16, 4),   -- Exacto: Costa Rica 0-3 Colombia
('22222222', 17, 2),   -- Resultado: Argentina 1-3 Peru
('22222222', 18, 2),   -- Resultado: Canada 2-2 Chile (empate)
('22222222', 19, 2),   -- Resultado: Jamaica 0-1 Venezuela
('22222222', 20, 2),   -- Resultado: Mexico 1-1 Ecuador (empate)
('22222222', 21, 4),   -- Exacto: Bolivia 2-0 Panama
('22222222', 22, 2),   -- Resultado: United States 0-1 Uruguay
('22222222', 23, 2),   -- Resultado: Costa Rica 1-3 Paraguay
('22222222', 24, 2),   -- Resultado: Brazil 0-2 Colombia

-- Usuario 3: Nacho1
('33333333', 1, 0),    -- Incorrecto: Argentina 0-2 Canada
('33333333', 2, 2),    -- Resultado: Peru 1-1 Chile (empate)
('33333333', 3, 2),    -- Resultado: Ecuador 3-2 Venezuela
('33333333', 4, 4),    -- Exacto: Mexico 0-0 Jamaica
('33333333', 5, 4),    -- Exacto: United States 1-0 Bolivia
('33333333', 6, 0),    -- Incorrecto: Uruguay 1-3 Panama
('33333333', 7, 2),    -- Resultado: Colombia 0-0 Paraguay (empate)
('33333333', 8, 0),    -- Incorrecto: Brazil 0-2 Costa Rica
('33333333', 9, 0),    -- Incorrecto: Peru 3-1 Canada
('33333333', 10, 2),   -- Resultado: Chile 2-2 Argentina (empate)
('33333333', 11, 4),   -- Exacto: Ecuador 1-0 Jamaica
('33333333', 12, 2),   -- Resultado: Venezuela 2-1 Mexico
('33333333', 13, 0),   -- Incorrecto: Panama 1-3 United States
('33333333', 14, 0),   -- Incorrecto: Bolivia 3-0 Uruguay
('33333333', 15, 2),   -- Resultado: Paraguay 1-1 Brazil (empate)
('33333333', 16, 4),   -- Exacto: Costa Rica 0-3 Colombia
('33333333', 17, 2),   -- Resultado: Argentina 1-3 Peru
('33333333', 18, 2),   -- Resultado: Canada 2-2 Chile (empate)
('33333333', 19, 2),   -- Resultado: Jamaica 0-1 Venezuela
('33333333', 20, 2),   -- Resultado: Mexico 1-1 Ecuador (empate)
('33333333', 21, 4),   -- Exacto: Bolivia 2-0 Panama
('33333333', 22, 2),   -- Resultado: United States 0-1 Uruguay
('33333333', 23, 2),   -- Resultado: Costa Rica 1-3 Paraguay
('33333333', 24, 2),   -- Resultado: Brazil 0-2 Colombia

-- Usuario 4: Nacho2
('44444444', 1, 2),    -- Resultado: Argentina 3-1 Canada
('44444444', 2, 4),    -- Exacto: Peru 1-0 Chile
('44444444', 3, 0),    -- Incorrecto: Ecuador 1-3 Venezuela
('44444444', 4, 0),    -- Incorrecto: Mexico 0-2 Jamaica
('44444444', 5, 2),    -- Resultado: United States 1-1 Bolivia (empate)
('44444444', 6, 4),    -- Exacto: Uruguay 2-0 Panama
('44444444', 7, 4),    -- Exacto: Colombia 0-1 Paraguay
('44444444', 8, 2),    -- Resultado: Brazil 1-3 Costa Rica
('44444444', 9, 2),    -- Resultado: Peru 0-2 Canada
('44444444', 10, 2),   -- Resultado: Chile 1-2 Argentina
('44444444', 11, 2),   -- Resultado: Ecuador 1-1 Jamaica (empate)
('44444444', 12, 0),   -- Incorrecto: Venezuela 0-3 Mexico
('44444444', 13, 2),   -- Resultado: Panama 2-1 United States
('44444444', 14, 2),   -- Resultado: Bolivia 0-2 Uruguay
('44444444', 15, 2),   -- Resultado: Paraguay 1-3 Brazil
('44444444', 16, 2),   -- Resultado: Costa Rica 2-2 Colombia (empate)
('44444444', 17, 0),   -- Incorrecto: Argentina 1-0 Peru
('44444444', 18, 2),   -- Resultado: Canada 2-1 Chile
('44444444', 19, 2),   -- Resultado: Jamaica 1-0 Venezuela
('44444444', 20, 2),   -- Resultado: Mexico 2-1 Ecuador
('44444444', 21, 2),   -- Resultado: Bolivia 1-1 Panama (empate)
('44444444', 22, 2),   -- Resultado: United States 2-3 Uruguay
('44444444', 23, 4),   -- Exacto: Costa Rica 0-1 Paraguay
('44444444', 24, 2),   -- Resultado: Brazil 2-1 Colombia

-- Usuario 5: Nacho Valle (49283227)
('49283227', 1, 2),    -- Resultado: Argentina 1-1 Canada (empate)
('49283227', 2, 4),    -- Exacto: Peru 0-2 Chile
('49283227', 3, 4),    -- Exacto: Ecuador 2-0 Venezuela
('49283227', 4, 4),    -- Exacto: Mexico 1-3 Jamaica
('49283227', 5, 4),    -- Exacto: United States 2-1 Bolivia
('49283227', 6, 4),    -- Exacto: Uruguay 1-1 Panama (empate)
('49283227', 7, 4),    -- Exacto: Colombia 3-2 Paraguay
('49283227', 8, 2),    -- Resultado: Brazil 0-1 Costa Rica
('49283227', 9, 2),    -- Resultado: Peru 2-2 Canada (empate)
('49283227', 10, 4),   -- Exacto: Chile 3-1 Argentina
('49283227', 11, 4),   -- Exacto: Ecuador 1-2 Jamaica
('49283227', 12, 2),   -- Resultado: Venezuela 0-1 Mexico
('49283227', 13, 4),   -- Exacto: Panama 3-0 United States
('49283227', 14, 4),   -- Exacto: Bolivia 2-1 Uruguay
('49283227', 15, 4),   -- Exacto: Paraguay 3-1 Brazil
('49283227', 16, 2),   -- Resultado: Costa Rica 1-0 Colombia
('49283227', 17, 4),   -- Exacto: Argentina 2-1 Peru
('49283227', 18, 4),   -- Exacto: Canada 3-3 Chile
('49283227', 19, 4),   -- Exacto: Jamaica 1-1 Venezuela (empate)
('49283227', 20, 2),   -- Resultado: Mexico 2-2 Ecuador (empate)
('49283227', 21, 4),   -- Exacto: Bolivia 0-3 Panama
('49283227', 22, 4),   -- Exacto: United States 1-2 Uruguay
('49283227', 23, 4),   -- Exacto: Costa Rica 3-1 Paraguay
('49283227', 24, 4),   -- Exacto: Brazil 1-0 Colombia

-- Usuario 6: Nacho3
('55555555', 1, 2),    -- Resultado: Argentina 2-3 Canada
('55555555', 2, 2),    -- Resultado: Peru 0-0 Chile (empate)
('55555555', 3, 4),    -- Exacto: Ecuador 2-1 Venezuela
('55555555', 4, 2),    -- Resultado: Mexico 3-2 Jamaica
('55555555', 5, 2),    -- Resultado: United States 1-2 Bolivia
('55555555', 6, 4),    -- Exacto: Uruguay 2-1 Panama
('55555555', 7, 2),    -- Resultado: Colombia 0-3 Paraguay
('55555555', 8, 2),    -- Resultado: Brazil 1-1 Costa Rica (empate)
('55555555', 9, 2),    -- Resultado: Peru 3-2 Canada
('55555555', 10, 2),   -- Resultado: Chile 1-0 Argentina
('55555555', 11, 4),   -- Exacto: Ecuador 0-1 Jamaica
('55555555', 12, 2),   -- Resultado: Venezuela 2-0 Mexico
('55555555', 13, 2),   -- Resultado: Panama 3-1 United States
('55555555', 14, 4),   -- Exacto: Bolivia 2-1 Uruguay
('55555555', 15, 2),   -- Resultado: Paraguay 0-0 Brazil (empate)
('55555555', 16, 2),   -- Resultado: Costa Rica 3-2 Colombia
('55555555', 17, 2),   -- Resultado: Argentina 1-1 Peru (empate)
('55555555', 18, 2),   -- Resultado: Canada 0-2 Chile
('55555555', 19, 2),   -- Resultado: Jamaica 1-1 Venezuela (empate)
('55555555', 20, 4),   -- Exacto: Mexico 3-2 Ecuador
('55555555', 21, 4),   -- Exacto: Bolivia 0-3 Panama
('55555555', 22, 2),   -- Resultado: United States 1-2 Uruguay
('55555555', 23, 2),   -- Resultado: Costa Rica 0-0 Paraguay (empate)
('55555555', 24, 2),   -- Resultado: Brazil 1-0 Colombia

-- Usuario 7: Anto
('66666666', 1, 2),    -- Resultado: Argentina 1-1 Canada (empate)
('66666666', 2, 2),    -- Resultado: Peru 0-2 Chile
('66666666', 3, 4),    -- Exacto: Ecuador 2-0 Venezuela
('66666666', 4, 4),    -- Exacto: Mexico 1-3 Jamaica
('66666666', 5, 4),    -- Exacto: United States 2-1 Bolivia
('66666666', 6, 4),    -- Exacto: Uruguay 1-1 Panama (empate)
('66666666', 7, 4),    -- Exacto: Colombia 3-2 Paraguay
('66666666', 8, 2),    -- Resultado: Brazil 0-1 Costa Rica
('66666666', 9, 2),    -- Resultado: Peru 2-2 Canada (empate)
('66666666', 10, 4),   -- Exacto: Chile 3-1 Argentina
('66666666', 11, 4),   -- Exacto: Ecuador 1-2 Jamaica
('66666666', 12, 2),   -- Resultado: Venezuela 0-1 Mexico
('66666666', 13, 4),   -- Exacto: Panama 3-0 United States
('66666666', 14, 4),   -- Exacto: Bolivia 2-1 Uruguay
('66666666', 15, 4),   -- Exacto: Paraguay 3-1 Brazil
('66666666', 16, 2),   -- Resultado: Costa Rica 1-0 Colombia
('66666666', 17, 4),   -- Exacto: Argentina 2-1 Peru
('66666666', 18, 4),   -- Exacto: Canada 3-3 Chile
('66666666', 19, 4),   -- Exacto: Jamaica 1-1 Venezuela (empate)
('66666666', 20, 2),   -- Resultado: Mexico 2-2 Ecuador (empate)
('66666666', 21, 4),   -- Exacto: Bolivia 0-3 Panama
('66666666', 22, 4),   -- Exacto: United States 1-2 Uruguay
('66666666', 23, 4),   -- Exacto: Costa Rica 3-1 Paraguay
('66666666', 24, 4);   -- Exacto: Brazil 1-0 Colombia

-- Inserción de puntajes de usuarios por campeonatos (ScoresChampionships)
INSERT INTO ScoresChampionships (document_id, championship_id, points) VALUES
-- Usuario 1: Nacho Valle
('11111111', 1, 15),    -- 10 puntos por acertar al campeón (Brazil) y 5 puntos por acertar al subcampeón (Argentina)
-- Usuario 2: Nico
('22222222', 1, 15),    -- 10 puntos por acertar al campeón (Brazil) y 5 puntos por acertar al subcampeón (Argentina)
-- Usuario 4: Nacho2
('33333333', 1, 15),    -- 10 puntos por acertar al campeón (Brazil) y 5 puntos por acertar al subcampeón (Uruguay)
-- Usuario 5: Nacho Valle (49283227)
('49283227', 1, 15),    -- 10 puntos por acertar al campeón (Brazil) y 5 puntos por acertar al subcampeón (Uruguay)
-- Usuario 6: Nacho3
('55555555', 1, 15),    -- 10 puntos por acertar al campeón (Brazil) y 5 puntos por acertar al subcampeón (Argentina)
-- Usuario 7: Anto
('66666666', 1, 15);    -- 10 puntos por acertar al campeón (Brazil) y 5 puntos por acertar al subcampeón (Uruguay)

SET FOREIGN_KEY_CHECKS = 1;