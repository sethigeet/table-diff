-- Delete the original tables if they exist
drop table if exists mock_data_1;
drop table if exists mock_data_2;
-- Mock Data 1
create table MOCK_DATA_1 (
  id INT,
  first_name VARCHAR(50),
  last_name VARCHAR(50),
  email VARCHAR(50),
  gender VARCHAR(50),
  ip_address VARCHAR(20)
);
insert into MOCK_DATA_1 (
    id,
    first_name,
    last_name,
    email,
    gender,
    ip_address
  )
values (
    1,
    'Licha',
    'Stoddard',
    'lstoddard0@army.mil',
    'Non-binary',
    '175.142.80.103'
  );
insert into MOCK_DATA_1 (
    id,
    first_name,
    last_name,
    email,
    gender,
    ip_address
  )
values (
    2,
    'Alene',
    'Babbs',
    'ababbs1@squarespace.com',
    'Genderfluid',
    '110.144.137.245'
  );
insert into MOCK_DATA_1 (
    id,
    first_name,
    last_name,
    email,
    gender,
    ip_address
  )
values (
    3,
    'Sybilla',
    'Risley',
    'srisley2@unicef.org',
    'Non-binary',
    '208.54.64.211'
  );
insert into MOCK_DATA_1 (
    id,
    first_name,
    last_name,
    email,
    gender,
    ip_address
  )
values (
    4,
    'Sheeree',
    'Cowlin',
    'scowlin3@seattletimes.com',
    'Bigender',
    '168.228.147.167'
  );
insert into MOCK_DATA_1 (
    id,
    first_name,
    last_name,
    email,
    gender,
    ip_address
  )
values (
    5,
    'Cari',
    'Bryett',
    'cbryett4@bloomberg.com',
    'Genderfluid',
    '35.9.62.135'
  );
insert into MOCK_DATA_1 (
    id,
    first_name,
    last_name,
    email,
    gender,
    ip_address
  )
values (
    6,
    'Ken',
    'O''Duane',
    'koduane6@bbb.org',
    'Agender',
    '125.227.167.129'
  );
insert into MOCK_DATA_1 (
    id,
    first_name,
    last_name,
    email,
    gender,
    ip_address
  )
values (
    7,
    'Heriberto',
    'Basant',
    'hbasant7@globo.com',
    'Male',
    '168.250.127.190'
  );
insert into MOCK_DATA_1 (
    id,
    first_name,
    last_name,
    email,
    gender,
    ip_address
  )
values (
    8,
    'Cecilla',
    'Idiens',
    'cidiens8@ihg.com',
    'Agender',
    '89.118.80.149'
  );
insert into MOCK_DATA_1 (
    id,
    first_name,
    last_name,
    email,
    gender,
    ip_address
  )
values (
    9,
    'Filia',
    'Tenney',
    'ftenneya@yale.edu',
    'Genderfluid',
    '216.167.41.196'
  );
-- Mock Data 2
create table MOCK_DATA_2 (
  id INT,
  first_name VARCHAR(50),
  last_name VARCHAR(50),
  email VARCHAR(50),
  gender VARCHAR(50),
  ip_address VARCHAR(20)
);
insert into MOCK_DATA_2 (
    id,
    first_name,
    last_name,
    email,
    gender,
    ip_address
  )
values (
    1,
    'Licha',
    'Stoddard',
    'lstoddard0@army.mil',
    'Non-binary',
    '175.142.80.103'
  );
insert into MOCK_DATA_2 (
    id,
    first_name,
    last_name,
    email,
    gender,
    ip_address
  )
values (
    2,
    'Alene',
    'Babbses',
    'ababbs1@squarespace.com',
    'Genderfluid',
    '110.144.137.245'
  );
insert into MOCK_DATA_2 (
    id,
    first_name,
    last_name,
    email,
    gender,
    ip_address
  )
values (
    3,
    'Sybilla',
    'Risley',
    'srisley2@unicef.org',
    'Non-binary',
    '208.54.64.211'
  );
insert into MOCK_DATA_2 (
    id,
    first_name,
    last_name,
    email,
    gender,
    ip_address
  )
values (
    4,
    'Mathilda',
    'Phayre',
    'mphayre5@harvard.edu',
    'Polygender',
    '97.103.214.10'
  );
insert into MOCK_DATA_2 (
    id,
    first_name,
    last_name,
    email,
    gender,
    ip_address
  )
values (
    5,
    'Heriberto',
    'Basant',
    'hbasant7@globo.com',
    'Male',
    '168.250.127.190'
  );
insert into MOCK_DATA_2 (
    id,
    first_name,
    last_name,
    email,
    gender,
    ip_address
  )
values (
    6,
    'Cecilla',
    'Idiens',
    'cidiens8@ihg.com',
    'Agender',
    '89.118.80.149'
  );
insert into MOCK_DATA_2 (
    id,
    first_name,
    last_name,
    email,
    gender,
    ip_address
  )
values (
    7,
    'Merle',
    'Drake',
    'mdrake9@spotify.com',
    'Agender',
    '52.58.240.6'
  );
insert into MOCK_DATA_2 (
    id,
    first_name,
    last_name,
    email,
    gender,
    ip_address
  )
values (
    8,
    'Filia',
    'Tenney',
    'ftenneya@yale.edu',
    'Genderfluid',
    '216.167.41.196'
  );
insert into MOCK_DATA_2 (
    id,
    first_name,
    last_name,
    email,
    gender,
    ip_address
  )
values (
    9,
    'Loy',
    'Deny',
    'ldenyb@va.gov',
    'Agender',
    '191.130.237.99'
  );