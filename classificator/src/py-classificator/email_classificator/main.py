import csv
from io import StringIO
import random

from proto.internal_nodes_pb2 import (
    EtlOutput,
    EmailCategory,
    Email,
    CATEGORY_ADMINISTRATION,
    CATEGORY_REPARATIONS,
    CATEGORY_BILLING,
    CATEGORY_SPAM,
    ClassificatorOutput,
)


async def default_handler(ctx, req):
    # If the message is an early reply, do nothing
    if ctx.is_message_early_reply():
        return

    # Unpack the request payload
    etl_output = EtlOutput()
    req.Unpack(etl_output)

    ctx.logger.info(f"Processing batch {etl_output.emails_key}")

    # Read the emails batch from the object store using the key provided in the request
    emails_batch = await ctx.object_store.get(etl_output.emails_key)

    # Process the emails batch and categorize the emails
    rows = csv.DictReader(StringIO(emails_batch.decode("utf-8")))
    for row in rows:
        email = dict_to_email(row)

        res = ClassificatorOutput()
        res.email.CopyFrom(email)
        res.category = classify_email(email)

        # If the category of the email is "reparations", send the email to the "repairs" subtopic
        if res.category == CATEGORY_REPARATIONS:
            await ctx.send_output(res, "repairs")

        # Otherwise, send the email to the main topic
        await ctx.send_output(res)


def classify_email(_: Email) -> EmailCategory:
    # Randomly assign a category to the email
    categories = [CATEGORY_REPARATIONS, CATEGORY_ADMINISTRATION, CATEGORY_BILLING, CATEGORY_SPAM]
    return categories[random.randint(0, len(categories) - 1)]


def dict_to_email(raw_email: dict[str, str]) -> Email:
    # Parse the raw email into an Email object
    email = Email()
    email.title = raw_email.get("title")
    email.body = raw_email.get("body")
    email.author = raw_email.get("author")
    email.creation_date = raw_email.get("date")

    return email
