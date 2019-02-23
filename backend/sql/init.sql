PRAGMA foreign_keys = off;
BEGIN TRANSACTION;

-- Table: Objectives
DROP TABLE IF EXISTS Objectives;

CREATE TABLE Objectives (
    ObjectivesID   INTEGER      PRIMARY KEY
                                UNIQUE
                                NOT NULL,
    SubmissionID   INTEGER      REFERENCES Submissions (SubmissionID) 
                                NOT NULL,
    PointValue     INTEGER      NOT NULL,
    ObjectiveType  VARCHAR (45) NOT NULL,
    ObjectiveMet   BOOLEAN      NOT NULL,
    ObjectiveValue VARCHAR (45) 
);


-- Table: Stories
DROP TABLE IF EXISTS Stories;

CREATE TABLE Stories (
    StoryID       INTEGER      PRIMARY KEY
                               NOT NULL
                               UNIQUE,
    Name          VARCHAR (45),
    MaxLength     INTEGER      NOT NULL,
    StoryComplete BOOLEAN      NOT NULL,
    UserID        INTEGER      REFERENCES Users (UserID),
    CurrentLength INTEGER      NOT NULL,
    CurrentStory  BOOLEAN      NOT NULL
);


-- Table: Submissions
DROP TABLE IF EXISTS Submissions;

CREATE TABLE Submissions (
    UserID       INTEGER REFERENCES Users (UserID),
    SubmissionID INTEGER PRIMARY KEY
                         NOT NULL
                         UNIQUE,
    Text         TEXT,
    Votes        INTEGER NOT NULL,
    Submitted    BOOLEAN NOT NULL,
    MaxLength    INT     NOT NULL,
    RoundID      INTEGER NOT NULL
                         REFERENCES TheRoundTable (RoundID) 
);


-- Table: TheRoundTable
DROP TABLE IF EXISTS TheRoundTable;

CREATE TABLE TheRoundTable (
    RoundID      INTEGER  PRIMARY KEY
                          NOT NULL
                          UNIQUE,
    StoryID      INTEGER  REFERENCES Stories (StoryID) 
                          NOT NULL,
    RoundNum     INTEGER  NOT NULL,
    EndTime      DATETIME NOT NULL,
    VoteTime     DATETIME NOT NULL,
    SubmissionID INTEGER  REFERENCES Submissions (SubmissionID) 
);


-- Table: Users
DROP TABLE IF EXISTS Users;

CREATE TABLE Users (
    UserID     INTEGER      PRIMARY KEY
                            UNIQUE
                            NOT NULL,
    Score      INT          NOT NULL,
    ScreenName VARCHAR (45),
    GoogleID   VARCHAR (45),
    FacebookID VARCHAR (45),
    TwitterID  VARCHAR (45),
    GithubID   VARCHAR (45),
    LinkedInID VARCHAR (45) 
);


COMMIT TRANSACTION;
PRAGMA foreign_keys = on;
