SET FOREIGN_KEY_CHECKS = 0;

-- Inserción de valores en Utils
INSERT INTO Utils (hours_until_match, exact_match_points, correct_result_match_points, champion_points, sub_champion_points) VALUES
(1, 4, 2, 10, 5);


-- Inserción de Roles de Usuario
INSERT INTO UserRoles (role_name, description) VALUES
('Admin', 'Administrator role'),
('Student', 'Student role');

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
('2023-06-20 19:00:00', 1, 11, NULL, NULL, 1, 1, 1), -- Argentina vs Canada
('2023-06-21 19:00:00', 8, 4, NULL, NULL, 1, 1, 1), -- Peru vs Chile
('2023-06-22 17:00:00', 6, 10, NULL, NULL, 1, 1, 2), -- Ecuador vs Venezuela
('2023-06-22 20:00:00', 15, 14, NULL, NULL, 1, 1, 2), -- Mexico vs Jamaica
('2023-06-23 17:00:00', 13, 2, NULL, NULL, 1, 1, 3), -- United States vs Bolivia
('2023-06-23 20:00:00', 9, 16, NULL, NULL, 1, 1, 3), -- Uruguay vs Panama
('2023-06-24 17:00:00', 5, 7, NULL, NULL, 1, 1, 4), -- Colombia vs Paraguay
('2023-06-24 18:00:00', 3, 12, NULL, NULL, 1, 1, 4), -- Brazil vs Costa Rica
('2023-06-25 17:00:00', 8, 11, NULL, NULL, 1, 1, 1), -- Peru vs Canada
('2023-06-25 20:00:00', 4, 1, NULL, NULL, 1, 1, 1), -- Chile vs Argentina
('2023-06-26 17:00:00', 6, 14, NULL, NULL, 1, 1, 2), -- Ecuador vs Jamaica
('2023-06-26 20:00:00', 10, 15, NULL, NULL, 1, 1, 2), -- Venezuela vs Mexico
('2023-06-27 17:00:00', 16, 13, NULL, NULL, 1, 1, 3), -- Panama vs United States
('2023-06-27 20:00:00', 2, 9, NULL, NULL, 1, 1, 3), -- Bolivia vs Uruguay
('2023-06-28 17:00:00', 7, 3, NULL, NULL, 1, 1, 4), -- Paraguay vs Brazil
('2023-06-28 20:00:00', 12, 5, NULL, NULL, 1, 1, 4), -- Costa Rica vs Colombia
('2023-06-29 19:00:00', 1, 8, NULL, NULL, 1, 1, 1), -- Argentina vs Peru
('2023-06-29 19:00:00', 11, 4, NULL, NULL, 1, 1, 1), -- Canada vs Chile
('2023-06-30 19:00:00', 14, 10, NULL, NULL, 1, 1, 2), -- Jamaica vs Venezuela
('2023-06-30 19:00:00', 15, 6, NULL, NULL, 1, 1, 2), -- Mexico vs Ecuador
('2023-07-01 20:00:00', 2, 16, NULL, NULL, 1, 1, 3), -- Bolivia vs Panama
('2023-07-01 20:00:00', 13, 9, NULL, NULL, 1, 1, 3), -- United States vs Uruguay
('2023-07-02 20:00:00', 12, 7, NULL, NULL, 1, 1, 4), -- Costa Rica vs Paraguay
('2023-07-02 20:00:00', 3, 5, NULL, NULL, 1, 1, 4); -- Brazil vs Colombia



-- Inserción de predicciones de partidos
INSERT INTO Predictions (goals_local, goals_visitor, document_id, match_id) VALUES
(2, 0, '1', 1),  -- Nacho predice Argentina vs Canada
(1, 1, '1', 2),  -- Nacho predice Peru vs Chile
(3, 2, '1', 3),  -- Nacho predice Ecuador vs Venezuela
(0, 0, '1', 4),  -- Nacho predice Mexico vs Jamaica
(1, 0, '1', 5),  -- Nacho predice United States vs Bolivia
(2, 1, '1', 6),  -- Nacho predice Uruguay vs Panama
(0, 0, '1', 7),  -- Nacho predice Colombia vs Paraguay
(3, 1, '1', 8),  -- Nacho predice Brazil vs Costa Rica
(1, 2, '1', 9),  -- Nacho predice Peru vs Canada
(2, 2, '1', 10), -- Nacho predice Chile vs Argentina
(1, 0, '1', 11), -- Nacho predice Ecuador vs Jamaica
(2, 1, '1', 12), -- Nacho predice Venezuela vs Mexico
(0, 1, '1', 13), -- Nacho predice Panama vs United States
(3, 3, '1', 14), -- Nacho predice Bolivia vs Uruguay
(1, 2, '1', 15), -- Nacho predice Paraguay vs Brazil
(0, 0, '1', 16), -- Nacho predice Costa Rica vs Colombia
(1, 3, '1', 17), -- Nacho predice Argentina vs Peru
(2, 2, '1', 18), -- Nacho predice Canada vs Chile
(0, 1, '1', 19), -- Nacho predice Jamaica vs Venezuela
(1, 1, '1', 20), -- Nacho predice Mexico vs Ecuador
(2, 0, '1', 21), -- Nacho predice Bolivia vs Panama
(0, 1, '1', 22), -- Nacho predice United States vs Uruguay
(1, 3, '1', 23), -- Nacho predice Costa Rica vs Paraguay
(0, 2, '1', 24), -- Nacho predice Brazil vs Colombia

(0, 2, '11212222', 1),  -- Vle N predice Argentina vs Canada
(2, 1, '11212222', 2),  -- Vle N predice Peru vs Chile
(1, 1, '11212222', 3),  -- Vle N predice Ecuador vs Venezuela
(0, 3, '11212222', 4),  -- Vle N predice Mexico vs Jamaica
(1, 0, '11212222', 5),  -- Vle N predice United States vs Bolivia
(1, 3, '11212222', 6),  -- Vle N predice Uruguay vs Panama
(2, 1, '11212222', 7),  -- Vle N predice Colombia vs Paraguay
(0, 2, '11212222', 8),  -- Vle N predice Brazil vs Costa Rica
(3, 1, '11212222', 9),  -- Vle N predice Peru vs Canada
(1, 2, '11212222', 10), -- Vle N predice Chile vs Argentina
(0, 0, '11212222', 11), -- Vle N predice Ecuador vs Jamaica
(2, 2, '11212222', 12), -- Vle N predice Venezuela vs Mexico
(1, 3, '11212222', 13), -- Vle N predice Panama vs United States
(3, 0, '11212222', 14), -- Vle N predice Bolivia vs Uruguay
(1, 1, '11212222', 15), -- Vle N predice Paraguay vs Brazil
(0, 3, '11212222', 16), -- Vle N predice Costa Rica vs Colombia
(2, 1, '11212222', 17), -- Vle N predice Argentina vs Peru
(0, 2, '11212222', 18), -- Vle N predice Canada vs Chile
(1, 0, '11212222', 19), -- Vle N predice Jamaica vs Venezuela
(2, 1, '11212222', 20), -- Vle N predice Mexico vs Ecuador
(3, 1, '11212222', 21), -- Vle N predice Bolivia vs Panama
(0, 2, '11212222', 22), -- Vle N predice United States vs Uruguay
(1, 0, '11212222', 23), -- Vle N predice Costa Rica vs Paraguay
(2, 1, '11212222', 24), -- Vle N predice Brazil vs Colombia

(3, 1, '12222', 1),  -- Vle N predice Argentina vs Canada
(1, 0, '12222', 2),  -- Vle N predice Peru vs Chile
(1, 3, '12222', 3),  -- Vle N predice Ecuador vs Venezuela
(0, 2, '12222', 4),  -- Vle N predice Mexico vs Jamaica
(1, 1, '12222', 5),  -- Vle N predice United States vs Bolivia
(2, 0, '12222', 6),  -- Vle N predice Uruguay vs Panama
(0, 1, '12222', 7),  -- Vle N predice Colombia vs Paraguay
(1, 3, '12222', 8),  -- Vle N predice Brazil vs Costa Rica
(0, 2, '12222', 9),  -- Vle N predice Peru vs Canada
(2, 1, '12222', 10), -- Vle N predice Chile vs Argentina
(1, 1, '12222', 11), -- Vle N predice Ecuador vs Jamaica
(0, 3, '12222', 12), -- Vle N predice Venezuela vs Mexico
(2, 1, '12222', 13), -- Vle N predice Panama vs United States
(0, 2, '12222', 14), -- Vle N predice Bolivia vs Uruguay
(1, 3, '12222', 15), -- Vle N predice Paraguay vs Brazil
(2, 2, '12222', 16), -- Vle N predice Costa Rica vs Colombia
(1, 0, '12222', 17), -- Vle N predice Argentina vs Peru
(2, 1, '12222', 18), -- Vle N predice Canada vs Chile
(0, 0, '12222', 19), -- Vle N predice Jamaica vs Venezuela
(3, 2, '12222', 20), -- Vle N predice Mexico vs Ecuador
(1, 1, '12222', 21), -- Vle N predice Bolivia vs Panama
(2, 3, '12222', 22), -- Vle N predice United States vs Uruguay
(0, 1, '12222', 23), -- Vle N predice Costa Rica vs Paraguay
(1, 0, '12222', 24), -- Vle N predice Brazil vs Colombia

(2, 3, '123', 1),  -- Vle N predice Argentina vs Canada
(0, 0, '123', 2),  -- Vle N predice Peru vs Chile
(2, 1, '123', 3),  -- Vle N predice Ecuador vs Venezuela
(3, 2, '123', 4),  -- Vle N predice Mexico vs Jamaica
(1, 2, '123', 5),  -- Vle N predice United States vs Bolivia
(2, 1, '123', 6),  -- Vle N predice Uruguay vs Panama
(0, 3, '123', 7),  -- Vle N predice Colombia vs Paraguay
(1, 1, '123', 8),  -- Vle N predice Brazil vs Costa Rica
(3, 2, '123', 9),  -- Vle N predice Peru vs Canada
(1, 0, '123', 10), -- Vle N predice Chile vs Argentina
(0, 1, '123', 11), -- Vle N predice Ecuador vs Jamaica
(2, 0, '123', 12), -- Vle N predice Venezuela vs Mexico
(3, 1, '123', 13), -- Vle N predice Panama vs United States
(2, 1, '123', 14), -- Vle N predice Bolivia vs Uruguay
(0, 0, '123', 15), -- Vle N predice Paraguay vs Brazil
(3, 2, '123', 16), -- Vle N predice Costa Rica vs Colombia
(1, 1, '123', 17), -- Vle N predice Argentina vs Peru
(0, 2, '123', 18), -- Vle N predice Canada vs Chile
(1, 1, '123', 19), -- Vle N predice Jamaica vs Venezuela
(2, 1, '123', 20), -- Vle N predice Mexico vs Ecuador
(0, 3, '123', 21), -- Vle N predice Bolivia vs Panama
(1, 2, '123', 22), -- Vle N predice United States vs Uruguay
(0, 0, '123', 23), -- Vle N predice Costa Rica vs Paraguay
(3, 1, '123', 24), -- Vle N predice Brazil vs Colombia

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
(3, 1, '1', 1),          -- Nacho predice Brazil campeón y Argentina subcampeón
(6, 4, '11212222', 1),   -- Vle N predice Ecuador campeón y Chile subcampeón
(9, 13, '12222', 1),     -- Vle N predice Uruguay campeón y United States subcampeón
(5, 12, '123', 1),       -- Vle N predice Colombia campeón y Costa Rica subcampeón
(3, 9, '49283227', 1);   -- Nacho Valle predice Brazil campeón y Uruguay subcampeón
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
('1', 1, 4),    -- Exacto: Argentina 2-0 Canada
('1', 2, 2),    -- Resultado: Peru 1-1 Chile (empate)
('1', 3, 4),    -- Exacto: Ecuador 3-2 Venezuela
('1', 4, 4),    -- Exacto: Mexico 0-0 Jamaica
('1', 5, 4),    -- Exacto: United States 1-0 Bolivia
('1', 6, 4),    -- Exacto: Uruguay 2-1 Panama
('1', 7, 4),    -- Exacto: Colombia 0-0 Paraguay
('1', 8, 4),    -- Exacto: Brazil 3-1 Costa Rica
('1', 9, 0),    -- Incorrecto: Peru 1-2 Canada
('1', 10, 2),   -- Resultado: Chile 2-2 Argentina (empate)
('1', 11, 4),   -- Exacto: Ecuador 1-0 Jamaica
('1', 12, 4),   -- Exacto: Venezuela 2-1 Mexico
('1', 13, 2),   -- Resultado: Panama 0-1 United States
('1', 14, 2),   -- Resultado: Bolivia 3-3 Uruguay (empate)
('1', 15, 2),   -- Resultado: Paraguay 1-2 Brazil
('1', 16, 4),   -- Exacto: Costa Rica 0-0 Colombia
('1', 17, 2),   -- Resultado: Argentina 1-3 Peru
('1', 18, 2),   -- Resultado: Canada 2-2 Chile (empate)
('1', 19, 2),   -- Resultado: Jamaica 0-1 Venezuela
('1', 20, 2),   -- Resultado: Mexico 1-1 Ecuador (empate)
('1', 21, 4),   -- Exacto: Bolivia 2-0 Panama
('1', 22, 2),   -- Resultado: United States 0-1 Uruguay
('1', 23, 2),   -- Resultado: Costa Rica 1-3 Paraguay
('1', 24, 4);   -- Exacto: Brazil 0-2 Colombia

-- Usuario 2: Vle N (11212222)
('11212222', 1, 0),    -- Incorrecto: Argentina 0-2 Canada
('11212222', 2, 2),    -- Resultado: Peru 1-1 Chile (empate)
('11212222', 3, 2),    -- Resultado: Ecuador 1-1 Venezuela (empate)
('11212222', 4, 4),    -- Exacto: Mexico 0-0 Jamaica
('11212222', 5, 4),    -- Exacto: United States 1-0 Bolivia
('11212222', 6, 0),    -- Incorrecto: Uruguay 1-3 Panama
('11212222', 7, 0),    -- Incorrecto: Colombia 2-1 Paraguay
('11212222', 8, 2),    -- Resultado: Brazil 3-1 Costa Rica
('11212222', 9, 0),    -- Incorrecto: Peru 3-1 Canada
('11212222', 10, 2),   -- Resultado: Chile 1-2 Argentina
('11212222', 11, 4),   -- Exacto: Ecuador 1-0 Jamaica
('11212222', 12, 2),   -- Resultado: Venezuela 2-2 Mexico (empate)
('11212222', 13, 2),   -- Resultado: Panama 1-3 United States
('11212222', 14, 2),   -- Resultado: Bolivia 3-0 Uruguay
('11212222', 15, 2),   -- Resultado: Paraguay 1-1 Brazil (empate)
('11212222', 16, 0),   -- Incorrecto: Costa Rica 0-3 Colombia
('11212222', 17, 2),   -- Resultado: Argentina 2-1 Peru
('11212222', 18, 0),   -- Incorrecto: Canada 0-2 Chile
('11212222', 19, 2),   -- Resultado: Jamaica 1-0 Venezuela
('11212222', 20, 2),   -- Resultado: Mexico 2-1 Ecuador
('11212222', 21, 2),   -- Resultado: Bolivia 1-3 Panama
('11212222', 22, 2),   -- Resultado: United States 0-2 Uruguay
('11212222', 23, 0),   -- Incorrecto: Costa Rica 1-0 Paraguay
('11212222', 24, 2);   -- Resultado: Brazil 2-1 Colombia

-- Usuario 3: Vle N (12222)
('12222', 1, 2),    -- Resultado: Argentina 3-1 Canada
('12222', 2, 4),    -- Exacto: Peru 1-0 Chile
('12222', 3, 0),    -- Incorrecto: Ecuador 1-3 Venezuela
('12222', 4, 2),    -- Resultado: Mexico 0-2 Jamaica
('12222', 5, 2),    -- Resultado: United States 1-1 Bolivia (empate)
('12222', 6, 4),    -- Exacto: Uruguay 2-0 Panama
('12222', 7, 2),    -- Resultado: Colombia 0-1 Paraguay
('12222', 8, 2),    -- Resultado: Brazil 1-3 Costa Rica
('12222', 9, 0),    -- Incorrecto: Peru 0-2 Canada
('12222', 10, 2),   -- Resultado: Chile 2-1 Argentina
('12222', 11, 2),   -- Resultado: Ecuador 1-1 Jamaica (empate)
('12222', 12, 2),   -- Resultado: Venezuela 0-3 Mexico
('12222', 13, 4),   -- Exacto: Panama 2-1 United States
('12222', 14, 4),   -- Exacto: Bolivia 2-0 Uruguay
('12222', 15, 2),   -- Resultado: Paraguay 1-3 Brazil
('12222', 16, 2),   -- Resultado: Costa Rica 2-2 Colombia (empate)
('12222', 17, 0),   -- Incorrecto: Argentina 1-0 Peru
('12222', 18, 2),   -- Resultado: Canada 2-1 Chile
('12222', 19, 4),   -- Exacto: Jamaica 0-0 Venezuela
('12222', 20, 4),   -- Exacto: Mexico 3-2 Ecuador
('12222', 21, 4),   -- Exacto: Bolivia 1-1 Panama (empate)
('12222', 22, 4),   -- Exacto: United States 2-3 Uruguay
('12222', 23, 2),   -- Resultado: Costa Rica 0-1 Paraguay
('12222', 24, 4);   -- Exacto: Brazil 1-0 Colombia

-- Usuario 4: Vle N (123)
('123', 1, 4),    -- Exacto: Argentina 2-3 Canada
('123', 2, 2),    -- Resultado: Peru 0-0 Chile (empate)
('123', 3, 2),    -- Resultado: Ecuador 2-1 Venezuela
('123', 4, 4),    -- Exacto: Mexico 3-2 Jamaica
('123', 5, 2),    -- Resultado: United States 1-2 Bolivia
('123', 6, 4),    -- Exacto: Uruguay 2-1 Panama
('123', 7, 2),    -- Resultado: Colombia 0-3 Paraguay
('123', 8, 2),    -- Resultado: Brazil 1-1 Costa Rica (empate)
('123', 9, 4),    -- Exacto: Peru 3-2 Canada
('123', 10, 2),   -- Resultado: Chile 1-0 Argentina
('123', 11, 2),   -- Resultado: Ecuador 0-1 Jamaica
('123', 12, 2),   -- Resultado: Venezuela 2-0 Mexico
('123', 13, 4),   -- Exacto: Panama 3-1 United States
('123', 14, 2),   -- Resultado: Bolivia 2-1 Uruguay
('123', 15, 2),   -- Resultado: Paraguay 0-0 Brazil (empate)
('123', 16, 2),   -- Resultado: Costa Rica 3-2 Colombia
('123', 17, 2),   -- Resultado: Argentina 1-1 Peru (empate)
('123', 18, 2),   -- Resultado: Canada 0-2 Chile
('123', 19, 2),   -- Resultado: Jamaica 1-1 Venezuela (empate)
('123', 20, 2),   -- Resultado: Mexico 2-1 Ecuador
('123', 21, 2),   -- Resultado: Bolivia 3-0 Panama
('123', 22, 2),   -- Resultado: United States 1-2 Uruguay
('123', 23, 2),   -- Resultado: Costa Rica 0-0 Paraguay (empate)
('123', 24, 2);   -- Resultado: Brazil 3-1 Colombia

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
('49283227', 24, 4);   -- Exacto: Brazil 1-0 Colombia

-- Inserción de puntajes de usuarios por campeonatos (ScoresChampionships)
INSERT INTO ScoresChampionships (document_id, championship_id, points) VALUES
-- Usuario 1: Nacho Valle
('1', 1, 15),    -- 10 puntos por acertar al campeón (Brazil) y 5 puntos por acertar al subcampeón (Argentina)
-- Usuario 2: Vle N (11212222)
('11212222', 1, 0),    -- 0 puntos, no acertó ni el campeón ni el subcampeón
-- Usuario 3: Vle N (12222)
('12222', 1, 5),    -- 5 puntos por acertar al subcampeón (United States)
-- Usuario 4: Vle N (123)
('123', 1, 0),    -- 0 puntos, no acertó ni el campeón ni el subcampeón
-- Usuario 5: Nacho Valle (49283227)
('49283227', 1, 15);  -- 10 puntos por acertar al campeón (Brazil) y 5 puntos por acertar al subcampeón (Uruguay)

SET FOREIGN_KEY_CHECKS = 1;
