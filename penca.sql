SET FOREIGN_KEY_CHECKS = 0;

-- Creación de la tabla UserRoles - Admin, Student, etc
CREATE TABLE UserRoles (
    role_id INT PRIMARY KEY AUTO_INCREMENT,
    role_name VARCHAR(50) NOT NULL,
    description VARCHAR(255)
);

-- Creación de la tabla User
CREATE TABLE User (
    document_id VARCHAR(10) PRIMARY KEY,
    email VARCHAR(100) NOT NULL,
    last_name VARCHAR(50) NOT NULL,
    first_name VARCHAR(50) NOT NULL,
    major VARCHAR(30),
    password VARCHAR(255) NOT NULL,
    role_id INT,
    FOREIGN KEY (role_id) REFERENCES UserRoles(role_id)
);

-- Creación de la tabla GroupStages, que representan los grupos de un campeonato (por ejemplo, grupos de la fase de grupos de un mundial)
-- Grupo A, Grupo B, etc
CREATE TABLE GroupStages (
    group_s_id INT PRIMARY KEY AUTO_INCREMENT,
    group_s_name VARCHAR(50) NOT NULL,
    description VARCHAR(255),
    championship_id INT,
    FOREIGN KEY (championship_id) REFERENCES Championships(championship_id)
);

-- Creación de la tabla Teams
CREATE TABLE Teams (
    team_id INT PRIMARY KEY AUTO_INCREMENT,
    description VARCHAR(100),
    url_logo VARCHAR(255),
    name VARCHAR(255) NOT NULL
);

-- Creación de la tabla Championships
CREATE TABLE Championships (
    championship_id INT PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(255) NOT NULL,
    year INT NOT NULL,
    country VARCHAR(50) NOT NULL,
    championship_type VARCHAR(255) NOT NULL,
    champion INT,
    subchampion INT,
    FOREIGN KEY (champion) REFERENCES Teams(team_id),
    FOREIGN KEY (subchampion) REFERENCES Teams(team_id)
);


-- Creación de la tabla GameMatch
CREATE TABLE GameMatch (
    match_id INT PRIMARY KEY AUTO_INCREMENT,
    match_date DATETIME NOT NULL,
    team_local_id INT NOT NULL,
    team_visitor_id INT NOT NULL,
    goals_local INT,
    goals_visitor INT,
    championship_id INT,
    stage_id INT,
    group_s_id INT,
    FOREIGN KEY (team_local_id) REFERENCES Teams(team_id),
    FOREIGN KEY (team_visitor_id) REFERENCES Teams(team_id),
    FOREIGN KEY (championship_id) REFERENCES Championships(championship_id),
    FOREIGN KEY (stage_id) REFERENCES Stages(stage_id),
    FOREIGN KEY (group_s_id) REFERENCES GroupStages(group_s_id)
);

CREATE TABLE Predictions (
    prediction_id INT PRIMARY KEY AUTO_INCREMENT, 
    goals_local INT NOT NULL, 
    goals_visitor INT NOT NULL, 
    document_id VARCHAR(10), 
    match_id INT, 
    UNIQUE KEY unique_prediction (match_id, document_id), 
    FOREIGN KEY (document_id) REFERENCES User(document_id), 
    FOREIGN KEY (match_id) REFERENCES GameMatch(match_id)
);

CREATE TABLE PredictionsChampionships (
    prediction_c_id INT PRIMARY KEY AUTO_INCREMENT, 
    champion INT, 
    subchampion INT, 
    document_id VARCHAR(10), 
    championship_id INT, 
    UNIQUE KEY unique_prediction_championship (championship_id, document_id), 
    FOREIGN KEY (document_id) REFERENCES User(document_id), 
    FOREIGN KEY (championship_id) REFERENCES Championships(championship_id), 
    FOREIGN KEY (champion) REFERENCES Teams(team_id), 
    FOREIGN KEY (subchampion) REFERENCES Teams(team_id) 
);

-- Creación de la tabla Stages, que representan las etapas de un campeonato (por ejemplo, fase de grupos, octavos de final, etc)
CREATE TABLE Stages (
    stage_id INT PRIMARY KEY AUTO_INCREMENT,
    stage_name VARCHAR(50) NOT NULL,
    description VARCHAR(255)
);

-- Creación de la tabla Scores para almacenar los puntajes de los usuarios de cada partido 
CREATE TABLE Scores (
    score_id INT PRIMARY KEY AUTO_INCREMENT,
    document_id VARCHAR(10),
    match_id INT,
    points INT,
    FOREIGN KEY (document_id) REFERENCES User(document_id),
    FOREIGN KEY (match_id) REFERENCES GameMatch(match_id),
    UNIQUE (document_id, match_id)
);

CREATE TABLE ScoresChampionships (
    score_c_id INT PRIMARY KEY AUTO_INCREMENT,
    document_id VARCHAR(10),
    championship_id INT,
    points INT,
    FOREIGN KEY (document_id) REFERENCES User(document_id),
    FOREIGN KEY (championship_id) REFERENCES Championships(championship_id),
    UNIQUE (document_id, championship_id)
);

-- Creación de la tabla Teams_Championships
CREATE TABLE Teams_Championships (
    team_id INT,
    championship_id INT,
    group_s_id INT,
    PRIMARY KEY (team_id, championship_id),
    FOREIGN KEY (group_s_id) REFERENCES GroupStages(group_s_id),
    FOREIGN KEY (team_id) REFERENCES Teams(team_id),
    FOREIGN KEY (championship_id) REFERENCES Championships(championship_id)
);

-- Creación de la tabla Utils
CREATE TABLE Utils (
    hours_until_match INT,
    exact_match_points INT,
    correct_result_match_points INT,
    champion_points INT,
    sub_champion_points INT
);

CREATE TABLE Notifications (
    notification_id INT PRIMARY KEY AUTO_INCREMENT,
    document_id VARCHAR(10),
    match_id INT,
    notification_time DATETIME NOT NULL,
    notification_method VARCHAR(50) NOT NULL,
    FOREIGN KEY (document_id) REFERENCES User(document_id),
    FOREIGN KEY (match_id) REFERENCES GameMatch(match_id),
    UNIQUE (document_id, match_id, notification_method)
);

-- Inserción de valores en Utils
INSERT INTO Utils (hours_until_match, exact_match_points, correct_result_match_points, champion_points, sub_champion_points) VALUES
(1, 4, 2, 10, 5);

-- Inserción de Roles de Usuario
INSERT INTO UserRoles (role_name, description) VALUES
('Admin', 'Administrator role'),
('Student', 'Student role');


-- Inserción de Campeonatos
INSERT INTO Championships (name, year, country, championship_type) VALUES
('Copa América 2024', 2024, 'Various', 'International');


INSERT INTO GroupStages (group_s_name, description, championship_id) VALUES
('Grupo A', 'Grupo A de la Copa América 2024', 1),
('Grupo B', 'Grupo B de la Copa América 2024', 1),
('Grupo C', 'Grupo C de la Copa América 2024', 1),
('Grupo D', 'Grupo D de la Copa América 2024', 1);


-- Inserción de Etapas para la Copa América
INSERT INTO Stages (stage_name, description) VALUES
('Fase de Grupos', 'Primera etapa de la Copa América 2024'),
('Cuartos de Final', 'Segunda etapa de la Copa América 2024'),
('Semifinal', 'Tercera etapa de la Copa América 2024'),
('Final', 'Etapa final de la Copa América 2024');

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


-- Inserción de Equipos en la Copa América 2024
-- Grupo A (Argentina, Peru, Chile, Canada),  1
-- Grupo B (Mexico, Ecuador, Venezuela, Jamaica), 2
-- Grupo C (United States, Uruguay, Panama, Bolivia), 3
-- Grupo D (Brazil, Colombia, Paraguay, Costa Rica), 4
-- Inserción de Equipos en la Copa América 2024
-- Grupo A (Argentina, Peru, Chile, Canada), grupo_id = 1
-- Grupo B (Mexico, Ecuador, Venezuela, Jamaica), grupo_id = 2
-- Grupo C (United States, Uruguay, Panama, Bolivia), grupo_id = 3
-- Grupo D (Brazil, Colombia, Paraguay, Costa Rica), grupo_id = 4
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


-- Inserción de Partidos de la fase de grupos de la Copa América 2024
INSERT INTO GameMatch (match_date, team_local_id, team_visitor_id, goals_local, goals_visitor, championship_id, stage_id, group_s_id) VALUES
('2024-06-20 19:00:00', 1, 11, NULL, NULL, 1, 1, 1), -- Argentina vs Canada
('2024-06-21 19:00:00', 8, 4, NULL, NULL, 1, 1, 1), -- Peru vs Chile
('2024-06-22 17:00:00', 6, 10, NULL, NULL, 1, 1, 2), -- Ecuador vs Venezuela
('2024-06-22 20:00:00', 15, 14, NULL, NULL, 1, 1, 2), -- Mexico vs Jamaica
('2024-06-23 17:00:00', 13, 2, NULL, NULL, 1, 1, 3), -- United States vs Bolivia
('2024-06-23 20:00:00', 9, 16, NULL, NULL, 1, 1, 3), -- Uruguay vs Panama
('2024-06-24 17:00:00', 5, 7, NULL, NULL, 1, 1, 4), -- Colombia vs Paraguay
('2024-06-24 18:00:00', 3, 12, NULL, NULL, 1, 1, 4), -- Brazil vs Costa Rica
('2024-06-25 17:00:00', 8, 11, NULL, NULL, 1, 1, 1), -- Peru vs Canada
('2024-06-25 20:00:00', 4, 1, NULL, NULL, 1, 1, 1), -- Chile vs Argentina
('2024-06-26 17:00:00', 6, 14, NULL, NULL, 1, 1, 2), -- Ecuador vs Jamaica
('2024-06-26 20:00:00', 10, 15, NULL, NULL, 1, 1, 2), -- Venezuela vs Mexico
('2024-06-27 17:00:00', 16, 13, NULL, NULL, 1, 1, 3), -- Panama vs United States
('2024-06-27 20:00:00', 2, 9, NULL, NULL, 1, 1, 3), -- Bolivia vs Uruguay
('2024-06-28 17:00:00', 7, 3, NULL, NULL, 1, 1, 4), -- Paraguay vs Brazil
('2024-06-28 20:00:00', 12, 5, NULL, NULL, 1, 1, 4), -- Costa Rica vs Colombia
('2024-06-29 19:00:00', 1, 8, NULL, NULL, 1, 1, 1), -- Argentina vs Peru
('2024-06-29 19:00:00', 11, 4, NULL, NULL, 1, 1, 1), -- Canada vs Chile
('2024-06-30 19:00:00', 14, 10, NULL, NULL, 1, 1, 2), -- Jamaica vs Venezuela
('2024-06-30 19:00:00', 15, 6, NULL, NULL, 1, 1, 2), -- Mexico vs Ecuador
('2024-07-01 20:00:00', 2, 16, NULL, NULL, 1, 1, 3), -- Bolivia vs Panama
('2024-07-01 20:00:00', 13, 9, NULL, NULL, 1, 1, 3), -- United States vs Uruguay
('2024-07-02 20:00:00', 12, 7, NULL, NULL, 1, 1, 4), -- Costa Rica vs Paraguay
('2024-07-02 20:00:00', 3, 5, NULL, NULL, 1, 1, 4); -- Brazil vs Colombia


/* -- Inserción de Partidos de Cuartos de Final de la Copa América 2024
INSERT INTO GameMatch (match_date, team_local_id, team_visitor_id, goals_local, goals_visitor, championship_id, stage_id) VALUES
('2024-07-04 20:00:00', NULL, NULL, NULL, NULL, 1, 2), -- Cuartos de Final 1
('2024-07-05 20:00:00', NULL, NULL, NULL, NULL, 1, 2), -- Cuartos de Final 2
('2024-07-06 17:00:00', NULL, NULL, NULL, NULL, 1, 2), -- Cuartos de Final 3
('2024-07-06 20:00:00', NULL, NULL, NULL, NULL, 1, 2); -- Cuartos de Final 4

-- Inserción de Partidos de Semifinales de la Copa América 2024
INSERT INTO GameMatch (match_date, team_local_id, team_visitor_id, goals_local, goals_visitor, championship_id, stage_id) VALUES
('2024-07-09 19:00:00', NULL, NULL, NULL, NULL, 1, 3), -- Semifinal 1
('2024-07-10 19:00:00', NULL, NULL, NULL, NULL, 1, 3); -- Semifinal 2

-- Inserción de Partidos de Tercer Puesto y Final de la Copa América 2024
INSERT INTO GameMatch (match_date, team_local_id, team_visitor_id, goals_local, goals_visitor, championship_id, stage_id) VALUES
('2024-07-13 19:00:00', NULL, NULL, NULL, NULL, 1, 4), -- Tercer Puesto
('2024-07-14 19:00:00', NULL, NULL, NULL, NULL, 1, 4); -- Final
 */

CREATE TABLE UserNotifications


SET FOREIGN_KEY_CHECKS = 1;
