SET FOREIGN_KEY_CHECKS = 0;

-- Creación de la tabla User
CREATE TABLE User (
    document_id VARCHAR(10) PRIMARY KEY,
    email VARCHAR(255) NOT NULL,
    last_name VARCHAR(255) NOT NULL,
    first_name VARCHAR(255) NOT NULL,
    major VARCHAR(255),
    password VARCHAR(255) NOT NULL,
    role VARCHAR(50)
);

-- Creación de la tabla Teams
CREATE TABLE Teams (
    team_id INT PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(255) NOT NULL
);

-- Creación de la tabla Championships
CREATE TABLE Championships (
    championship_id INT PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(255) NOT NULL,
    year INT NOT NULL,
    country VARCHAR(255) NOT NULL,
    championship_type VARCHAR(255) NOT NULL,
    champion INT,
    subchampion INT,
    FOREIGN KEY (champion) REFERENCES Teams(team_id),
    FOREIGN KEY (subchampion) REFERENCES Teams(team_id)
);

-- Creación de la tabla ChampionshipGroups
CREATE TABLE ChampionshipGroups (
    group_id INT PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(255) NOT NULL,
    championship_id INT,
    referredCode VARCHAR(255),
    champion INT,
    subchampion INT,
    FOREIGN KEY (championship_id) REFERENCES Championships(championship_id),
    FOREIGN KEY (champion) REFERENCES Teams(team_id),
    FOREIGN KEY (subchampion) REFERENCES Teams(team_id)
);

-- Creación de la tabla User_ChampionshipGroups
CREATE TABLE User_ChampionshipGroups (
    document_id VARCHAR(10),
    group_id INT,
    PRIMARY KEY (document_id, group_id),
    FOREIGN KEY (document_id) REFERENCES User(document_id),
    FOREIGN KEY (group_id) REFERENCES ChampionshipGroups(group_id)
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
    FOREIGN KEY (team_local_id) REFERENCES Teams(team_id),
    FOREIGN KEY (team_visitor_id) REFERENCES Teams(team_id),
    FOREIGN KEY (championship_id) REFERENCES Championships(championship_id)
);

-- Creación de la tabla Predictions
CREATE TABLE Predictions (
    prediction_id INT PRIMARY KEY AUTO_INCREMENT,
    goals_local INT NOT NULL,
    goals_visitor INT NOT NULL,
    document_id VARCHAR(10),
    match_id INT,
    group_id INT,
    FOREIGN KEY (document_id) REFERENCES User(document_id),
    FOREIGN KEY (match_id) REFERENCES GameMatch(match_id),
    FOREIGN KEY (group_id) REFERENCES ChampionshipGroups(group_id)
);

-- Creación de la tabla Scores
CREATE TABLE Scores (
    score_id INT PRIMARY KEY AUTO_INCREMENT,
    document_id VARCHAR(10),
    match_id INT,
    points INT,
    FOREIGN KEY (document_id) REFERENCES User(document_id),
    FOREIGN KEY (match_id) REFERENCES GameMatch(match_id),
    UNIQUE (document_id, match_id)
);

-- Creación de la tabla Teams_Championships
CREATE TABLE Teams_Championships (
    team_id INT,
    championship_id INT,
    PRIMARY KEY (team_id, championship_id),
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

-- Inserción de valores en Utils
INSERT INTO Utils (hours_until_match, exact_match_points, correct_result_match_points, champion_points, sub_champion_points) VALUES
(1, 4, 2, 10, 5);

-- Inserción de Usuarios
INSERT INTO User (document_id, email, last_name, first_name, major, password, role) VALUES
('12345678', 'john.doe@example.com', 'Doe', 'John', 'Computer Science', 'password123', 'student'),
('87654321', 'jane.smith@example.com', 'Smith', 'Jane', 'Mathematics', 'password456', 'student'),
('56781234', 'mike.jones@example.com', 'Jones', 'Mike', 'Physics', 'password789', 'student'),
('43218765', 'laura.garcia@example.com', 'Garcia', 'Laura', 'Engineering', 'password321', 'student'),
('34567812', 'robert.brown@example.com', 'Brown', 'Robert', 'Economics', 'password654', 'student');

-- Inserción de Campeonatos
INSERT INTO Championships (name, year, country, championship_type) VALUES
('Campeonato Uruguayo 2024', 2024, 'Uruguay', 'National'),
('Premier League', 2024, 'England', 'International');

-- Inserción de Grupos de Usuarios en Campeonatos
INSERT INTO ChampionshipGroups (name, championship_id, referredCode) VALUES
('Group A', 1, 'CODE123'),
('Group B', 2, 'CODE456'),
('Group C', 1, 'CODE789'),
('Group D', 2, 'CODE012');

-- Inserción de Usuarios en Grupos de Campeonato
INSERT INTO User_ChampionshipGroups (document_id, group_id) VALUES
('12345678', 1),
('87654321', 1),
('56781234', 2),
('43218765', 3),
('34567812', 4);

-- Inserción de Equipos
INSERT INTO Teams (name) VALUES
-- Equipos del Campeonato Uruguayo
('Nacional'),
('Peñarol'),
('Defensor Sporting'),
('Danubio'),
-- Equipos de la Premier League
('Manchester United'),
('Liverpool'),
('Chelsea'),
('Arsenal');

-- Inserción de Equipos y Campeonatos
INSERT INTO Teams_Championships (team_id, championship_id) VALUES
(1, 1), -- Nacional en el Campeonato Uruguayo
(2, 1), -- Peñarol en el Campeonato Uruguayo
(3, 1), -- Defensor Sporting en el Campeonato Uruguayo
(4, 1), -- Danubio en el Campeonato Uruguayo
(5, 2), -- Manchester United en la Premier League
(6, 2), -- Liverpool en la Premier League
(7, 2), -- Chelsea en la Premier League
(8, 2); -- Arsenal en la Premier League

-- Inserción de Partidos del Campeonato Uruguayo (incluyendo todos los partidos)
INSERT INTO GameMatch (match_date, team_local_id, team_visitor_id, goals_local, goals_visitor, championship_id) VALUES
('2024-06-01 15:00:00', 1, 2, 2, 1, 1), -- Nacional 2 vs Peñarol 1
('2024-06-02 18:00:00', 3, 4, 0, 0, 1), -- Defensor Sporting 0 vs Danubio 0
('2024-06-05 14:00:00', 1, 3, 3, 2, 1), -- Nacional 3 vs Defensor Sporting
('2024-06-06 17:00:00', 2, 4, 1, 0, 1), -- Peñarol 1 vs Danubio
('2024-06-09 15:00:00', 3, 1, 2, 0, 1), -- Defensor Sporting 2 vs Nacional
('2024-06-10 18:00:00', 4, 2, 0, 2, 1), -- Danubio 0 vs Peñarol
('2024-06-13 14:00:00', 1, 4, 2, 2, 1), -- Nacional 2 vs Danubio
('2024-06-14 17:00:00', 2, 3, 1, 1, 1), -- Peñarol 1 vs Defensor Sporting
-- Otros partidos del Campeonato Uruguayo sin resultados definidos (asumimos fechas futuras para simplificación)
('2024-06-17 15:00:00', 4, 1, NULL, NULL, 1), -- Danubio vs Nacional
('2024-06-18 18:00:00', 3, 2, NULL, NULL, 1), -- Defensor Sporting vs Peñarol
('2024-06-20 15:00:00', 1, 4, NULL, NULL, 1), -- Nacional vs Danubio
('2024-06-22 18:00:00', 2, 3, NULL, NULL, 1), -- Peñarol vs Defensor Sporting
('2024-06-25 15:00:00', 3, 1, NULL, NULL, 1), -- Defensor Sporting vs Nacional
('2024-06-27 18:00:00', 4, 2, NULL, NULL, 1); -- Danubio vs Peñarol

-- Inserción de Partidos de la Premier League
INSERT INTO GameMatch (match_date, team_local_id, team_visitor_id, goals_local, goals_visitor, championship_id) VALUES
('2024-06-03 16:00:00', 5, 6, 1, 1, 2), -- Manchester United 1 vs Liverpool 1
('2024-06-04 20:00:00', 7, 8, 2, 3, 2), -- Chelsea 2 vs Arsenal 3
('2024-06-07 19:00:00', 5, 7, 2, 2, 2), -- Manchester United 2 vs Chelsea 2
('2024-06-08 21:00:00', 6, 8, 1, 3, 2), -- Liverpool 1 vs Arsenal 3
('2024-06-11 16:00:00', 7, 5, 0, 1, 2), -- Chelsea 0 vs Manchester United
('2024-06-12 20:00:00', 8, 6, 3, 1, 2), -- Arsenal 3 vs Liverpool
('2024-06-15 19:00:00', 5, 6, 3, 3, 2), -- Manchester United 3 vs Liverpool
('2024-06-16 21:00:00', 7, 8, 0, 1, 2); -- Chelsea 0 vs Arsenal

-- Inserción de Predicciones
INSERT INTO Predictions (goals_local, goals_visitor, document_id, match_id, group_id) VALUES
(2, 1, '12345678', 1, 1), -- John predice Nacional 2 vs Peñarol 1
(0, 0, '87654321', 2, 1), -- Jane predice Defensor Sporting 0 vs Danubio 0
(1, 1, '56781234', 3, 2), -- Mike predice Manchester United 1 vs Liverpool 1
(2, 3, '12345678', 4, 2), -- John predice Chelsea 2 vs Arsenal 3
(3, 1, '43218765', 1, 3), -- Laura predice Nacional 3 vs Peñarol 1
(0, 2, '34567812', 3, 4), -- Robert predice Manchester United 0 vs Liverpool 2
(3, 2, '12345678', 5, 1), -- John predice Nacional 3 vs Defensor Sporting 2
(1, 0, '87654321', 6, 1), -- Jane predice Peñarol 1 vs Danubio 0
(2, 2, '56781234', 7, 2), -- Mike predice Manchester United 2 vs Chelsea 2
(1, 3, '43218765', 8, 2), -- Laura predice Liverpool 1 vs Arsenal 3
(3, 1, '34567812', 5, 3), -- Robert predice Nacional 3 vs Defensor Sporting 1
(0, 0, '87654321', 7, 4), -- Emily predice Manchester City 0 vs Tottenham 0
(2, 0, '12345678', 9, 1), -- John predice Nacional 2 vs Defensor Sporting
(1, 2, '87654321', 10, 1), -- Jane predice Danubio 1 vs Peñarol
(0, 0, '56781234', 11, 2), -- Mike predice Chelsea 0 vs Manchester United
(3, 1, '43218765', 12, 2), -- Laura predice Arsenal 3 vs Liverpool
(2, 2, '34567812', 13, 3), -- Robert predice Nacional 2 vs Danubio
(1, 1, '87654321', 14, 4), -- Emily predice Peñarol 1 vs Defensor Sporting
(0, 3, '12345678', 15, 5), -- Alex predice Liverpool 0 vs Everton
(1, 0, '87654321', 16, 6), -- Maria predice Arsenal 1 vs Newcastle
(2, 1, '12345678', 17, 1), -- John predice Danubio 2 vs Nacional
(0, 2, '87654321', 18, 1), -- Jane predice Defensor Sporting 0 vs Peñarol
(1, 1, '56781234', 19, 2), -- Mike predice Nacional 1 vs Danubio
(2, 3, '43218765', 20, 2), -- Laura predice Peñarol 2 vs Defensor Sporting
(1, 0, '34567812', 21, 3), -- Robert predice Defensor Sporting 1 vs Nacional
(0, 0, '87654321', 22, 4); -- Emily predice Danubio 0 vs Peñarol

-- Inserción de Puntajes
INSERT INTO Scores (document_id, match_id, points) VALUES
('12345678', 1, 5), -- John obtiene 5 puntos por el partido 1
('87654321', 2, 6), -- Jane obtiene 6 puntos por el partido 2
('56781234', 3, 5), -- Mike obtiene 5 puntos por el partido 3
('12345678', 4, 4), -- John obtiene 4 puntos por el partido 4
('43218765', 1, 3), -- Laura obtiene 3 puntos por el partido 1
('34567812', 3, 4), -- Robert obtiene 4 puntos por el partido 3
('12345678', 5, 5), -- John obtiene 5 puntos por el partido 5
('87654321', 6, 6), -- Jane obtiene 6 puntos por el partido 6
('56781234', 7, 5), -- Mike obtiene 5 puntos por el partido 7
('43218765', 8, 4), -- Laura obtiene 4 puntos por el partido 8
('34567812', 5, 3), -- Robert obtiene 3 puntos por el partido 5
('87654321', 7, 4), -- Emily obtiene 4 puntos por el partido 7
('12345678', 9, 5), -- John obtiene 5 puntos por el partido 9
('87654321', 10, 6), -- Jane obtiene 6 puntos por el partido 10
('56781234', 11, 5), -- Mike obtiene 5 puntos por el partido 11
('43218765', 12, 4), -- Laura obtiene 4 puntos por el partido 12
('34567812', 13, 3), -- Robert obtiene 3 puntos por el partido 13
('87654321', 14, 4), -- Emily obtiene 4 puntos por el partido 14
('12345678', 15, 5), -- Alex obtiene 5 puntos por el partido 15
('87654321', 16, 6), -- Maria obtiene 6 puntos por el partido 16
('12345678', 17, 5), -- John obtiene 5 puntos por el partido 17
('87654321', 18, 6), -- Jane obtiene 6 puntos por el partido 18
('56781234', 19, 5), -- Mike obtiene 5 puntos por el partido 19
('43218765', 20, 4), -- Laura obtiene 4 puntos por el partido 20
('34567812', 21, 3), -- Robert obtiene 3 puntos por el partido 21
('87654321', 22, 4); -- Emily obtiene 4 puntos por el partido 22

SET FOREIGN_KEY_CHECKS = 1;
