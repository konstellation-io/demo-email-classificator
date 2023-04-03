import csv

from io import StringIO


def main():
    batch_size = 3
    emails = []
    with open("emails.csv") as csvfile:
        rows = csv.DictReader(csvfile)
        field_names = rows.fieldnames
        for row in rows:
            emails.append(row)

    batch = StringIO()
    csv_writer = csv.DictWriter(batch, fieldnames=field_names)
    csv_writer.writeheader()
    count = 0
    batch_count = 0
    for email in emails:
        count += 1
        csv_writer.writerow(email)
        if count == batch_size:
            batch_key = f"emails-{batch_count}"
            batch_count += 1

            print(batch_key)
            print(batch.getvalue())
            print(bytes(batch.getvalue(), 'utf-8'))

            emails = StringIO()
            csv.DictWriter(emails, field_names)
            csv_writer.writeheader()
            count = 0


main()
