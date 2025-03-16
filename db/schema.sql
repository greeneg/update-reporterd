--
-- File generated with SQLiteStudio v3.4.17 on Sun Mar 16 13:32:14 2025
--
-- Text encoding used: UTF-8
--
PRAGMA foreign_keys = off;
BEGIN TRANSACTION;

-- Table: Architectures
DROP TABLE IF EXISTS Architectures;

CREATE TABLE IF NOT EXISTS Architectures (
    Id           INTEGER  PRIMARY KEY AUTOINCREMENT
                          UNIQUE
                          NOT NULL,
    ArchName     STRING   UNIQUE
                          NOT NULL,
    CreationDate DATETIME NOT NULL
                          DEFAULT (CURRENT_TIMESTAMP) 
);

INSERT INTO Architectures (
                              Id,
                              ArchName,
                              CreationDate
                          )
                          VALUES (
                              1,
                              'noarch',
                              '2025-03-14 13:50:42'
                          );

INSERT INTO Architectures (
                              Id,
                              ArchName,
                              CreationDate
                          )
                          VALUES (
                              2,
                              'aarch64',
                              '2025-03-14 13:50:42'
                          );

INSERT INTO Architectures (
                              Id,
                              ArchName,
                              CreationDate
                          )
                          VALUES (
                              3,
                              'x86',
                              '2025-03-14 13:50:42'
                          );

INSERT INTO Architectures (
                              Id,
                              ArchName,
                              CreationDate
                          )
                          VALUES (
                              4,
                              'x86_64',
                              '2025-03-14 13:50:42'
                          );


-- Table: OperatingSystems
DROP TABLE IF EXISTS OperatingSystems;

CREATE TABLE IF NOT EXISTS OperatingSystems (
    Id           INTEGER  PRIMARY KEY AUTOINCREMENT
                          UNIQUE
                          NOT NULL,
    OsName       STRING   NOT NULL,
    OsVersion    STRING   NOT NULL,
    OsFamilyId   INTEGER  REFERENCES OsFamilies (Id) 
                          NOT NULL,
    OsArchId     INTEGER  REFERENCES Architectures (Id) 
                          NOT NULL,
    CreationDate DATETIME NOT NULL
                          DEFAULT (CURRENT_TIMESTAMP) 
);


-- Table: OrgUnits
DROP TABLE IF EXISTS OrgUnits;

CREATE TABLE IF NOT EXISTS OrgUnits (
    Id           INTEGER  PRIMARY KEY AUTOINCREMENT
                          UNIQUE
                          NOT NULL,
    Name         STRING   UNIQUE
                          NOT NULL,
    ShortName    STRING   UNIQUE
                          NOT NULL,
    CreationDate DATETIME NOT NULL
                          DEFAULT (CURRENT_TIMESTAMP) 
);


-- Table: OsFamilies
DROP TABLE IF EXISTS OsFamilies;

CREATE TABLE IF NOT EXISTS OsFamilies (
    Id           INTEGER  PRIMARY KEY AUTOINCREMENT
                          UNIQUE
                          NOT NULL,
    FamilyName   STRING   UNIQUE
                          NOT NULL,
    CreationDate DATETIME NOT NULL
                          DEFAULT (CURRENT_TIMESTAMP) 
);

INSERT INTO OsFamilies (
                           Id,
                           FamilyName,
                           CreationDate
                       )
                       VALUES (
                           1,
                           'linux',
                           '2025-03-14 13:50:42'
                       );

INSERT INTO OsFamilies (
                           Id,
                           FamilyName,
                           CreationDate
                       )
                       VALUES (
                           2,
                           'darwin',
                           '2025-03-14 13:50:42'
                       );

INSERT INTO OsFamilies (
                           Id,
                           FamilyName,
                           CreationDate
                       )
                       VALUES (
                           3,
                           'windows',
                           '2025-03-14 13:50:42'
                       );


-- Table: Packages
DROP TABLE IF EXISTS Packages;

CREATE TABLE IF NOT EXISTS Packages (
    Id           INTEGER  PRIMARY KEY AUTOINCREMENT
                          UNIQUE
                          NOT NULL,
    Name         STRING   NOT NULL
                          UNIQUE,
    CreationDate DATETIME NOT NULL
                          DEFAULT (CURRENT_TIMESTAMP) 
);


-- Table: PackagesByArchitecture
DROP TABLE IF EXISTS PackagesByArchitecture;

CREATE TABLE IF NOT EXISTS PackagesByArchitecture (
    Id             INTEGER  PRIMARY KEY AUTOINCREMENT
                            UNIQUE
                            NOT NULL,
    PackageId      INTEGER  REFERENCES Packages (Id) 
                            NOT NULL,
    ArchitectureId INTEGER  REFERENCES Architectures (Id) 
                            NOT NULL,
    CreationDate   DATETIME NOT NULL
                            DEFAULT (CURRENT_TIMESTAMP) 
);


-- Table: PackagesByPlatform
DROP TABLE IF EXISTS PackagesByPlatform;

CREATE TABLE IF NOT EXISTS PackagesByPlatform (
    Id           INTEGER  PRIMARY KEY AUTOINCREMENT
                          UNIQUE
                          NOT NULL,
    PackageId    INTEGER  REFERENCES Packages (Id) 
                          NOT NULL,
    PlatformId   INTEGER  REFERENCES OperatingSystems (Id) 
                          NOT NULL,
    CreationDate DATETIME NOT NULL
                          DEFAULT (CURRENT_TIMESTAMP) 
);


-- Table: PackagesByType
DROP TABLE IF EXISTS PackagesByType;

CREATE TABLE IF NOT EXISTS PackagesByType (
    Id            INTEGER  PRIMARY KEY AUTOINCREMENT
                           UNIQUE
                           NOT NULL,
    PackageId     INTEGER  REFERENCES Packages (Id) 
                           NOT NULL,
    PackageTypeId INTEGER  REFERENCES PackageTypes (Id) 
                           NOT NULL,
    CreationDate  DATETIME NOT NULL
                           DEFAULT (CURRENT_TIMESTAMP) 
);


-- Table: PackagesByVersion
DROP TABLE IF EXISTS PackagesByVersion;

CREATE TABLE IF NOT EXISTS PackagesByVersion (
    Id             INTEGER  PRIMARY KEY AUTOINCREMENT
                            UNIQUE
                            NOT NULL,
    PackageId      INTEGER  REFERENCES Packages (Id) 
                            NOT NULL,
    PackageVersion STRING   NOT NULL,
    CreationDate   DATETIME NOT NULL
                            DEFAULT (CURRENT_TIMESTAMP) 
);


-- Table: PackageTypes
DROP TABLE IF EXISTS PackageTypes;

CREATE TABLE IF NOT EXISTS PackageTypes (
    Id           INTEGER  PRIMARY KEY AUTOINCREMENT
                          UNIQUE
                          NOT NULL,
    TypeName     STRING   UNIQUE
                          NOT NULL,
    CreationDate DATETIME NOT NULL
                          DEFAULT (CURRENT_TIMESTAMP) 
);


-- Table: Roles
DROP TABLE IF EXISTS Roles;

CREATE TABLE IF NOT EXISTS Roles (
    Id           INTEGER  PRIMARY KEY AUTOINCREMENT
                          UNIQUE
                          NOT NULL,
    RoleName     STRING   UNIQUE
                          NOT NULL,
    Description  STRING   NOT NULL,
    CreationDate DATETIME NOT NULL
                          DEFAULT (CURRENT_TIMESTAMP) 
);

INSERT INTO Roles (
                      Id,
                      RoleName,
                      Description,
                      CreationDate
                  )
                  VALUES (
                      1,
                      'SYSTEM',
                      'Built-in system role',
                      '2025-03-14 13:50:42'
                  );

INSERT INTO Roles (
                      Id,
                      RoleName,
                      Description,
                      CreationDate
                  )
                  VALUES (
                      2,
                      'administrators',
                      'Accounts that have full administrative rights to the system',
                      '2025-03-14 13:50:42'
                  );


-- Table: Systems
DROP TABLE IF EXISTS Systems;

CREATE TABLE IF NOT EXISTS Systems (
    Id           INTEGER  PRIMARY KEY AUTOINCREMENT
                          UNIQUE
                          NOT NULL,
    FQDN         STRING   UNIQUE
                          NOT NULL,
    OsFamilyId   INTEGER  REFERENCES OsFamilies (Id) 
                          NOT NULL,
    OsId         INTEGER  REFERENCES OperatingSystems (Id) 
                          NOT NULL,
    ArchId       INTEGER  REFERENCES Architectures (Id) 
                          NOT NULL,
    CreationDate DATETIME NOT NULL
                          DEFAULT (CURRENT_TIMESTAMP) 
);


-- Table: UpdateCount
DROP TABLE IF EXISTS UpdateCount;

CREATE TABLE IF NOT EXISTS UpdateCount (
    Id            INTEGER  PRIMARY KEY AUTOINCREMENT
                           UNIQUE
                           NOT NULL,
    SystemId      INTEGER  REFERENCES Systems (Id) 
                           NOT NULL,
    Count         INTEGER  NOT NULL,
    CreationDate  DATETIME NOT NULL
                           DEFAULT (CURRENT_TIMESTAMP),
    LastCheckDate DATETIME NOT NULL
                           DEFAULT (CURRENT_TIMESTAMP) 
);


-- Table: Users
DROP TABLE IF EXISTS Users;

CREATE TABLE IF NOT EXISTS Users (
    Id                     INTEGER  PRIMARY KEY AUTOINCREMENT
                                    UNIQUE
                                    NOT NULL,
    UserName               STRING   UNIQUE
                                    NOT NULL,
    FullName               STRING   NOT NULL,
    Status                 BOOLEAN  NOT NULL,
    OrgUnitId              INTEGER  REFERENCES OrgUnits (Id) 
                                    NOT NULL,
    RoleId                 INTEGER  REFERENCES Roles (Id) 
                                    NOT NULL,
    PasswordHash           STRING   NOT NULL,
    CreationDate           DATETIME NOT NULL
                                    DEFAULT (CURRENT_TIMESTAMP),
    LastPasswordChangeDate DATETIME NOT NULL
                                    DEFAULT (CURRENT_TIMESTAMP) 
);


COMMIT TRANSACTION;
PRAGMA foreign_keys = on;
