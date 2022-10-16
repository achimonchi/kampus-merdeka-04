CREATE TABLE students (
    id varchar(5) NOT NULL PRIMARY KEY,
    name varchar(100) NOT NULL,
    age int NOT NULL,
    grade int NOT NULL
);

INSERT INTO students (id,name,age,grade)
VALUES
    ("S-1", "James Bond", 20, 1),
    ("S-2", "James Rodriguez", 22, 1),
    ("S-3", "Luffy", 18, 2),
    ("S-4", "Spongebob", 30, 1);

CREATE TABLE domain(
    GlobalRank int,
    TldRank int,
    Domain varchar(255),
    TLD varchar(255),
    RefSubNets int,
    RefIps int,
    IDN_Domain varchar(255),
    IDN_TLD varchar(255),
    PrevGlobalRank int,
    PrevTldRank int,
    PrevRefSubNets int,
    PrevRefIps int,

)