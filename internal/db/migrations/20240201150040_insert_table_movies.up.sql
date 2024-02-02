INSERT INTO movies (ID, Title, Year, Rated, Released, Runtime, Genre, Director, Writer, Actors, Plot, Language, Country,
                    Awards, Poster, Metascore, ImdbRating, ImdbVotes, ImdbID, Type, DVD, BoxOffice, Production, Website,
                    Response)
VALUES (1, 'The Wolf of Wall Street', '2013',
        'R', '25 Dec 2013', '180 min', 'Biography, Comedy, Crime', 'Martin Scorsese',
        'Terence Winter, Jordan Belfort', 'Leonardo DiCaprio, Jonah Hill, Margot Robbie',
        'In the early 1990s, Jordan Belfort teamed with his partner Donny Azoff and started brokerage firm Stratton Oakmont...',
        'English, French', 'United States', 'Nominated for 5 Oscars. 37 wins & 179 nominations total',
        'https://m.media-amazon.com/images/M/MV5BMjIxMjgxNTk0MF5BMl5BanBnXkFtZTgwNjIyOTg2MDE@._V1_SX300.jpg',
        '75', '8.2', '1,545,718', 'tt0993846', 'movie', '12 Dec 2015', '$116,900,694', 'N/A', 'N/A', 'True'),
       (2, 'Interstellar', '2014', 'PG-13', '07 Nov 2014', '169 min',
        'Adventure, Drama, Sci-Fi', 'Christopher Nolan', 'Jonathan Nolan, Christopher Nolan',
        'Matthew McConaughey, Anne Hathaway, Jessica Chastain',
        'Earth''s future has been riddled by disasters, famines, and droughts. There is only one way to ensure mankind''s survival: Interstellar travel. A newly discovered wormhole in the far reaches of our solar system allows a team of astronauts to go where no man has gone before, a planet that may have the right environment to sustain human life.',
        'English', 'United States, United Kingdom, Canada', 'Won 1 Oscar. 44 wins & 148 nominations total',
        'https://m.media-amazon.com/images/M/MV5BZjdkOTU3MDktN2IxOS00OGEyLWFmMjktY2FiMmZkNWIyODZiXkEyXkFqcGdeQXVyMTMxODk2OTU@._V1_SX300.jpg',
        '74', '8.7', '2,036,452', 'tt0816692', 'movie', '24 May 2016', '$188,020,017', 'N/A', 'N/A', 'True'),
       (3, 'Oppenheimer', '2023', 'R', '21 Jul 2023', '180 min',
        'Biography, Drama, History', 'Christopher Nolan', 'Christopher Nolan, Kai Bird, Martin Sherwin',
        'Cillian Murphy, Emily Blunt, Matt Damon',
        'The story of American scientist J. Robert Oppenheimer and his role in the development of the atomic bomb.',
        'English, German, Italian', 'United States, United Kingdom', '102 wins & 231 nominations',
        'https://m.media-amazon.com/images/M/MV5BMDBmYTZjNjUtN2M1MS00MTQ2LTk2ODgtNzc2M2QyZGE5NTVjXkEyXkFqcGdeQXVyNzAwMjU2MTY@._V1_SX300.jpg',
        '88', '8.4', '577,705', 'tt15398776', 'movie', '21 Nov 2023', '$326,102,235', 'N/A', 'N/A', 'True');

COMMIT;