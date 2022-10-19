import csv

from proto.public_input_pb2 import Response
from proto.internal_nodes_pb2 import EtlOutput, Email


def init(ctx):
    ctx.logger.info("[executing init]")
    emails = []
    with open(ctx.path("src/etl/emails_high_load.csv")) as csvfile:
        rows = csv.DictReader(csvfile)
        for row in rows:
            emails.append(dict_to_email(row))
    ctx.set("emails", emails)
    ctx.logger.info("[init] emails csv loaded")


async def default_handler(ctx, _):
    ctx.logger.info("[executing default handler]")
    emails = ctx.get("emails")

    res = Response()
    res.message = f"Processing of {len(emails)} emails in progress"
    await ctx.send_early_reply(res)

    for email in emails:
        etl_output = EtlOutput()
        etl_output.email.CopyFrom(email)
        await ctx.send_output(etl_output)
    return


def dict_to_email(raw_email: dict[str, str]) -> Email:
    email = Email()
    email.title = raw_email.get("title")
    email.body = raw_email.get("body")
    email.author = raw_email.get("author")
    email.creation_date = raw_email.get("date")

    return email
