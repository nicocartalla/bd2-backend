SET FOREIGN_KEY_CHECKS = 0;

CREATE TABLE User (
    user_id INT PRIMARY KEY AUTO_INCREMENT,
    email VARCHAR(255) NOT NULL,
    last_name VARCHAR(255) NOT NULL,
    first_name VARCHAR(255) NOT NULL,
    major VARCHAR(255),
    password VARCHAR(255) NOT NULL,
    role VARCHAR(50)
);

CREATE TABLE Championships (
    championship_id INT PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(255) NOT NULL,
    year INT NOT NULL,
    country VARCHAR(255) NOT NULL,
    championship_type VARCHAR(255) NOT NULL
);

CREATE TABLE UserGroups (
    group_id INT PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(255) NOT NULL,
    championship_id INT,
    FOREIGN KEY (championship_id) REFERENCES Championships(championship_id)
);

CREATE TABLE User_UserGroups (
    user_id INT ,
    group_id INT,
    PRIMARY KEY (user_id, group_id),
    FOREIGN KEY (user_id) REFERENCES User(user_id),
    FOREIGN KEY (group_id) REFERENCES UserGroups(group_id)
);

CREATE TABLE Teams (
    team_id INT PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(255) NOT NULL
);

CREATE TABLE GameMatch (
    match_id INT PRIMARY KEY AUTO_INCREMENT,
    match_date DATETIME NOT NULL,
    team_local_id INT,
    team_visitor_id INT,
    goals_local INT,
    goals_visitor INT,
    championship_id INT,
    FOREIGN KEY (team_local_id) REFERENCES Teams(team_id),
    FOREIGN KEY (team_visitor_id) REFERENCES Teams(team_id),
    FOREIGN KEY (championship_id) REFERENCES Championships(championship_id)
);

CREATE TABLE Predictions (
    prediction_id INT PRIMARY KEY AUTO_INCREMENT,
    goals_local INT NOT NULL,
    goals_visitor INT NOT NULL,
    user_id INT,
    match_id INT,
    group_id INT,
    FOREIGN KEY (user_id) REFERENCES User(user_id),
    FOREIGN KEY (match_id) REFERENCES GameMatch(match_id),
    FOREIGN KEY (group_id) REFERENCES UserGroups(group_id)
);

CREATE TABLE Scores (
    score_id INT PRIMARY KEY AUTO_INCREMENT,
    user_id INT,
    match_id INT,
    points INT NOT NULL,
    FOREIGN KEY (user_id) REFERENCES User(user_id),
    FOREIGN KEY (match_id) REFERENCES GameMatch(match_id)
);

CREATE TABLE Teams_Championships (
    team_id INT,
    championship_id INT,
    PRIMARY KEY (team_id, championship_id),
    FOREIGN KEY (team_id) REFERENCES Teams(team_id),
    FOREIGN KEY (championship_id) REFERENCES Championships(championship_id)
);

CREATE TABLE Utils (
    hours_until_match INT,
    exact_match_points INT,
    correct_result_match_points INT,
    champion_points INT,
    sub_champion_points INT
);

-- UTILS
INSERT INTO Utils (hours_until_match, exact_match_points, correct_result_match_points, champion_points, sub_champion_points) VALUES
(1, 4, 2, 10, 5);

-- Usuarios
INSERT INTO User (email, last_name, first_name, major, password, role) VALUES
('john.doe@example.com', 'Doe', 'John', 'Computer Science', 'password123', 'student'),
('jane.smith@example.com', 'Smith', 'Jane', 'Mathematics', 'password456', 'student'),
('mike.jones@example.com', 'Jones', 'Mike', 'Physics', 'password789', 'student'),
('laura.garcia@example.com', 'Garcia', 'Laura', 'Engineering', 'password321', 'student'),
('robert.brown@example.com', 'Brown', 'Robert', 'Economics', 'password654', 'student');

-- Campeonatos
INSERT INTO Championships (name, year, country, championship_type) VALUES
('Campeonato Uruguayo 2024', 2024, 'Uruguay', 'National'),
('Premier League', 2024, 'England', 'International');

-- Grupos de Usuarios
INSERT INTO UserGroups (name, championship_id) VALUES
('Group A', 1),
('Group B', 2),
('Group C', 1),
('Group D', 2);

-- Usuarios y Grupos
INSERT INTO User_UserGroups (user_id, group_id) VALUES
(1, 1),
(2, 1),
(3, 2),
(4, 3),
(5, 4);

-- Equipos
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

-- Partidos
INSERT INTO GameMatch (match_date, team_local_id, team_visitor_id, goals_local, goals_visitor, championship_id) VALUES
('2024-06-01 15:00:00', 1, 2, 2, 1, 1), -- Nacional 2 vs Peñarol 1 en el Campeonato Uruguayo
('2024-06-02 18:00:00', 3, 4, 0, 0, 1), -- Defensor Sporting 0 vs Danubio 0 en el Campeonato Uruguayo
('2024-06-03 16:00:00', 5, 6, 1, 1, 2), -- Manchester United 1 vs Liverpool 1 en la Premier League
('2024-06-04 20:00:00', 7, 8, 2, 3, 2); -- Chelsea 2 vs Arsenal 3 en la Premier League

-- Predicciones
INSERT INTO Predictions (goals_local, goals_visitor, user_id, match_id, group_id) VALUES
(2, 1, 1, 1, 1), -- John predice Nacional 2 vs Peñarol 1
(0, 0, 2, 2, 1), -- Jane predice Defensor Sporting 0 vs Danubio 0
(1, 1, 3, 3, 2), -- Mike predice Manchester United 1 vs Liverpool 1
(2, 3, 1, 4, 2), -- John predice Chelsea 2 vs Arsenal 3
(3, 1, 4, 1, 3), -- Laura predice Nacional 3 vs Peñarol 1
(0, 2, 5, 3, 4); -- Robert predice Manchester United 0 vs Liverpool 2

-- Puntajes
INSERT INTO Scores (user_id, match_id, points) VALUES
(1, 1, 5), -- John obtiene 5 puntos por el partido 1
(2, 2, 6), -- Jane obtiene 6 puntos por el partido 2
(3, 3, 5), -- Mike obtiene 5 puntos por el partido 3
(1, 4, 4), -- John obtiene 4 puntos por el partido 4
(4, 1, 3), -- Laura obtiene 3 puntos por el partido 1
(5, 3, 4); -- Robert obtiene 4 puntos por el partido 3

-- Equipos y Campeonatos
INSERT INTO Teams_Championships (team_id, championship_id) VALUES
(1, 1), -- Nacional en el Campeonato Uruguayo
(2, 1), -- Peñarol en el Campeonato Uruguayo
(3, 1), -- Defensor Sporting en el Campeonato Uruguayo
(4, 1), -- Danubio en el Campeonato Uruguayo
(5, 2), -- Manchester United en la Premier League
(6, 2), -- Liverpool en la Premier League
(7, 2), -- Chelsea en la Premier League
(8, 2); -- Arsenal en la Premier League

SET FOREIGN_KEY_CHECKS = 1;
