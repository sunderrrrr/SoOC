-- Creating the Company table
CREATE TABLE Company (
    ID INT PRIMARY KEY,
    Name VARCHAR(255) NOT NULL,
    Adress VARCHAR(255),
    Owner VARCHAR(255),
    Contact VARCHAR(255)
);

-- Creating the Employers table
CREATE TABLE Employers (
    EmpID INT PRIMARY KEY,
    Name VARCHAR(255) NOT NULL,
    ComanyId INT,
    Level VARCHAR(255),
    Contact VARCHAR(255),
    FOREIGN KEY (ComanyId) REFERENCES Company(ID)
);

-- Creating the Dishes table
CREATE TABLE Dishes (
    DishID INT PRIMARY KEY,
    Name VARCHAR(255) NOT NULL,
    CompanyID INT,
    Price DECIMAL(10, 2),
    Include VARCHAR(255),
    FOREIGN KEY (CompanyID) REFERENCES Company(ID)
);

-- Creating the Order table
CREATE TABLE Orderz (
    OrderID INT PRIMARY KEY,
    Date DATE NOT NULL,
    CompanyID INT,
    Price DECIMAL(10, 2),
    Include VARCHAR(255),
    FOREIGN KEY (CompanyID) REFERENCES Company(ID)
);