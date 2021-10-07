from faker import Faker

from rich import print
from rich.progress import track

fake = Faker("en_US")

TABLE_STRUCTURE = {
    "id": "SERIAL PRIMARY KEY",
    "username": "VARCHAR(50)",
    "name": "VARCHAR(100)",
    "email": "VARCHAR(50)",
    "gender": "VARCHAR(50)",
    "address": "VARCHAR(255)",
}


def get_create_table_query(table_name: str) -> str:
    query = f"CREATE TABLE {table_name} (\n"

    for k, v in TABLE_STRUCTURE.items():
        query += f"  {k} {v},\n"

    # Remove the extra comma from the end
    query = query[:-2]

    query += "\n);\n"

    return query


def get_drop_table_query(table_name: str) -> str:
    query = f"DROP TABLE IF EXISTS {table_name};\n"

    return query


def get_random_insert_query(table_name: str) -> str:
    query = f"INSERT INTO {table_name} ("

    for key in TABLE_STRUCTURE.keys():
        if key != "id":
            query += f"{key}, "

    query = query[:-2]

    query += ") VALUES ("

    simple_profile = fake.simple_profile()
    query += "'" + simple_profile["username"] + "', "
    query += "'" + simple_profile["name"] + "', "
    query += "'" + simple_profile["mail"] + "', "
    query += "'" + simple_profile["sex"] + "', "
    query += "'" + simple_profile["address"].replace("\n", "\\n") + "'"
    query += ");\n"

    return query


common_insert_queries = []
for i in range(1000):
    common_insert_queries.append(get_random_insert_query("mock_data_1"))
common_insert_queries = "\n".join(common_insert_queries)


# Make the file empty
print("Emptying the file...", end="")
with open("./create_mock_data_larger.sql", "w") as file:
    file.write("")
print("[bold green]Done[/bold green]")

file = open("./create_mock_data_larger.sql", "a")

print("Writing drop table queries...", end="")
file.write("-- Drop the old tables\n")
file.write(get_drop_table_query("mock_data_1"))
file.write(get_drop_table_query("mock_data_2"))
print("[bold green]Done[/bold green]")
file.write("\n")

print("Writing create table queries...", end="")
file.write("-- Create the tables\n")
file.write(get_create_table_query("mock_data_1"))
file.write(get_create_table_query("mock_data_2"))
print("[bold green]Done[/bold green]")
file.write("\n")

print("Writing insert queries...")
file.write("-- Insert the values\n")

for _ in track(range(99000), description="Writing insert queries for mock_data_1..."):
    file.write(get_random_insert_query("mock_data_1"))

file.write(common_insert_queries)

file.write("\n")

for _ in track(range(99000), description="Writing insert queries for mock_data_2..."):
    file.write(get_random_insert_query("mock_data_2"))

common_insert_queries = common_insert_queries.replace(
    "mock_data_1", "mock_data_2")
file.write(common_insert_queries)

print("[bold green]Done[/bold green]")

file.close()
