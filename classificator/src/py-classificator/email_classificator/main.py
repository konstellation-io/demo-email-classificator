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
    if ctx.is_message_early_reply():
        return

    etl_output = EtlOutput()
    req.Unpack(etl_output)

    ctx.logger.info(f"Processing batch {etl_output.emails_key}")

    emails_batch = await ctx.object_store.get(etl_output.emails_key)

    rows = csv.DictReader(StringIO(emails_batch.decode("utf-8")))
    for row in rows:
        email = dict_to_email(row)

        res = ClassificatorOutput()
        res.email.CopyFrom(email)
        res.category = classify_email(email)

        if res.category == CATEGORY_REPARATIONS:
            await ctx.send_output(res, "repairs")

        await ctx.send_output(res)


def classify_email(_: Email) -> EmailCategory:
    categories = [CATEGORY_REPARATIONS, CATEGORY_ADMINISTRATION, CATEGORY_BILLING, CATEGORY_SPAM]
    return categories[random.randint(0, len(categories) - 1)]


def dict_to_email(raw_email: dict[str, str]) -> Email:
    email = Email()
    email.title = raw_email.get("title")
    email.body = raw_email.get("body")
    email.author = raw_email.get("author")
    email.creation_date = raw_email.get("date")

    return email

async def new_handler(ctx, req)-> None:
    if ctx.is_message_early_reply():
        return

    etl_output = EtlOutput()
    req.Unpack(etl_output)

    email = Email(
        title="test",
        body="test",
        author="test",
        creation_date="test",
    )
    res = ClassificatorOutput()
    res.email.CopyFrom(email)
    res.category = classify_email(email)

    if res.category == CATEGORY_REPARATIONS:
        await ctx.send_output(res, "repairs")

    await ctx.send_output(res)

custom_handlers = {
    "etl": new_handler,
}