SET FOREIGN_KEY_CHECKS = 0;
CREATE TABLE User (
    user_id INT PRIMARY KEY,
    email VARCHAR(255) NOT NULL,
    last_name VARCHAR(255) NOT NULL,
    first_name VARCHAR(255) NOT NULL,
    major VARCHAR(255),
    password VARCHAR(255) NOT NULL,
    role VARCHAR(50)
);

CREATE TABLE Championships (
    championship_id INT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    year INT NOT NULL,
    country VARCHAR(255) NOT NULL,
    championship_type VARCHAR(255) NOT NULL
);

CREATE TABLE UserGroups (
    group_id INT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    championship_id INT,
    FOREIGN KEY (championship_id) REFERENCES Championships(championship_id)
);

CREATE TABLE User_UserGroups (
    user_id INT,
    group_id INT,
    PRIMARY KEY (user_id, group_id),
    FOREIGN KEY (user_id) REFERENCES User(user_id),
    FOREIGN KEY (group_id) REFERENCES UserGroups(group_id)
);

CREATE TABLE Teams (
    team_id INT PRIMARY KEY,
    name VARCHAR(255) NOT NULL
);

CREATE TABLE GameMatch (
    match_id INT PRIMARY KEY,
    date DATE NOT NULL,
    team_local_id INT,
    team_visitor_id INT,
    championship_id INT,
    FOREIGN KEY (team_local_id) REFERENCES Teams(team_id),
    FOREIGN KEY (team_visitor_id) REFERENCES Teams(team_id),
    FOREIGN KEY (championship_id) REFERENCES Championships(championship_id)
);

CREATE TABLE Results (
    result_id INT PRIMARY KEY,
    match_id INT,
    goals_local INT NOT NULL,
    goals_visitor INT NOT NULL,
    FOREIGN KEY (match_id) REFERENCES GameMatch(match_id)
);

CREATE TABLE Predictions (
    prediction_id INT PRIMARY KEY,
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
    score_id INT PRIMARY KEY,
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

-- Usuarios
INSERT INTO User (user_id, email, last_name, first_name, major, password, role) VALUES
(1, 'john.doe@example.com', 'Doe', 'John', 'Computer Science', 'password123', 'student'),
(2, 'jane.smith@example.com', 'Smith', 'Jane', 'Mathematics', 'password456', 'student'),
(3, 'mike.jones@example.com', 'Jones', 'Mike', 'Physics', 'password789', 'student');

-- Campeonatos
INSERT INTO Championships (championship_id, name, year, country, championship_type) VALUES
(1, 'Campeonato Uruguayo 2024', 2024, 'Uruguay', 'National'),
(2, 'Premier League', 2024, 'England', 'International');

-- Grupos de Usuarios
INSERT INTO UserGroups (group_id, name, championship_id) VALUES
(1, 'Group A', 1),
(2, 'Group B', 2);

-- Usuarios y Grupos
INSERT INTO User_UserGroups (user_id, group_id) VALUES
(1, 1),
(2, 1),
(3, 2);

-- Equipos
INSERT INTO Teams (team_id, name) VALUES
-- Equipos del Campeonato Uruguayo
(1, 'Nacional'),
(2, 'Peñarol'),
(3, 'Defensor Sporting'),
(4, 'Danubio'),
-- Equipos de la Premier League
(5, 'Manchester United'),
(6, 'Liverpool'),
(7, 'Chelsea'),
(8, 'Arsenal');

-- Partidos
INSERT INTO GameMatch (match_id, date, team_local_id, team_visitor_id, championship_id) VALUES
(1, '2024-06-01', 1, 2, 1), -- Nacional vs Peñarol en el Campeonato Uruguayo
(2, '2024-06-02', 3, 4, 1), -- Defensor Sporting vs Danubio en el Campeonato Uruguayo
(3, '2024-06-03', 5, 6, 2), -- Manchester United vs Liverpool en la Premier League
(4, '2024-06-04', 7, 8, 2); -- Chelsea vs Arsenal en la Premier League

-- Resultados
INSERT INTO Results (result_id, match_id, goals_local, goals_visitor) VALUES
(1, 1, 2, 1), -- Nacional 2 vs Peñarol 1
(2, 2, 0, 0), -- Defensor Sporting 0 vs Danubio 0
(3, 3, 1, 1), -- Manchester United 1 vs Liverpool 1
(4, 4, 2, 3); -- Chelsea 2 vs Arsenal 3

-- Predicciones
INSERT INTO Predictions (prediction_id, goals_local, goals_visitor, user_id, match_id, group_id) VALUES
(1, 2, 1, 1, 1, 1), -- John predice Nacional 2 vs Peñarol 1
(2, 0, 0, 2, 2, 1), -- Jane predice Defensor Sporting 0 vs Danubio 0
(3, 1, 1, 3, 3, 2), -- Mike predice Manchester United 1 vs Liverpool 1
(4, 2, 3, 1, 4, 2); -- John predice Chelsea 2 vs Arsenal 3

-- Puntajes
INSERT INTO Scores (score_id, user_id, match_id, points) VALUES
(1, 1, 1, 5), -- John obtiene 5 puntos por el partido 1
(2, 2, 2, 6), -- Jane obtiene 6 puntos por el partido 2
(3, 3, 3, 5), -- Mike obtiene 5 puntos por el partido 3
(4, 1, 4, 4); -- John obtiene 4 puntos por el partido 4

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
