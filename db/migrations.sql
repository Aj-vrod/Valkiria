CREATE TABLE IF NOT EXISTS movies
(
    movie_id SERIAL PRIMARY KEY,
    name VARCHAR(255) UNIQUE NOT NULL,
    genre VARCHAR(255) NOT NULL
);

INSERT INTO movies VALUES
(
    1, 'Saw', 'horror'
),
(
    2, 'What heppens in Vegas', 'comedy'
),
(
    3, 'Die Hard', 'action'
),
(
    4, 'The Silence of the Lambs', 'thriller'
),
(
    5, 'The green mile', 'drama'
),
(
    6, 'The Lord of the Rings', 'fantasy'
),
(
    7, 'The Matrix', 'sci-fi'
),
(
    8, 'The Hobbit', 'adventure'
),
(
    9, 'The Good, the Bad and the Ugly', 'western'
),
(
    10, 'The Godfather', 'crime'
),
(
    11, 'Titanic', 'romance'
)
