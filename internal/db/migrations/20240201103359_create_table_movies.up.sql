CREATE TABLE movies
(
    ID         INT NOT NULL,
    Title      VARCHAR(255),
    Year       VARCHAR(4),
    Rated      VARCHAR(10),
    Released   VARCHAR(20),
    Runtime    VARCHAR(10),
    Genre      VARCHAR(255),
    Director   VARCHAR(255),
    Writer     VARCHAR(255),
    Actors     VARCHAR(255),
    Plot       TEXT,
    Language   VARCHAR(255),
    Country    VARCHAR(255),
    Awards     VARCHAR(255),
    Poster     VARCHAR(255),
    Metascore  VARCHAR(5),
    ImdbRating VARCHAR(5),
    ImdbVotes  VARCHAR(20),
    ImdbID     VARCHAR(20),
    Type       VARCHAR(20),
    DVD        VARCHAR(20),
    BoxOffice  VARCHAR(20),
    Production VARCHAR(255),
    Website    VARCHAR(255),
    Response   VARCHAR(5),
    PRIMARY KEY (ID)
);


Commit;
