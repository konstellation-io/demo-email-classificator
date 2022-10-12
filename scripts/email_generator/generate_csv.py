import random, csv
import sys

from faker import Faker

base_authors = [
    "human.resources@demo.com",
    "administration@demo.com",
    "john.doe@demo.com",
    "alice@demo.com",
    "bob@demo.com",
]
fake = Faker()

title_field = "title"
date_field = "date"
author_field = "author"
body_field = "body"


def get_random_author():
    if random.randint(0, 1):
        return base_authors[random.randint(0, len(base_authors) - 1)]
    return f'{fake.name().replace(" ", ".")}@demo.com'


def generate_email() -> [dict[str, any]]:
    random_text_formatted = fake.text(random.randint(100, 400)).replace("\n", "\\n")
    return {
        title_field: fake.text(random.randint(25, 40)),
        date_field: fake.date_time(),
        author_field: get_random_author(),
        body_field: random_text_formatted,
    }


def generate_emails(amount: int) -> [str]:
    return [generate_email() for _ in range(amount)]


def write_emails_csv(path, emails: [dict[str, any]]):
    with open(path, mode="w") as csv_file:
        field_names = [title_field, date_field, author_field, body_field]
        writer = csv.DictWriter(csv_file, fieldnames=field_names)
        writer.writeheader()
        [writer.writerow(email) for email in emails]


def main():
    try:
        file_path = sys.argv[1]
        amount = sys.argv[2]
        emails = generate_emails(int(amount))
        write_emails_csv(file_path, emails)
    except Exception as e:
        print(f'Error getting args {e}')
        print("Usage: generate_csv.py (amount_of_emails) (file_path)")


main()
