import csv

from proto.public_input_pb2 import Response
from proto.internal_nodes_pb2 import EtlOutput, Email


def init(ctx):
    ctx.logger.info("[executing init]")
    emails = []
    with open("/krt-files/src/etl/emails.csv") as csvfile:
        rows = csv.DictReader(csvfile)
        # next(rows, None)
        for row in rows:
            emails.append(row)
    ctx.logger.info(emails[0])
    ctx.set("emails", emails)
    ctx.logger.info("[init] emails csv loaded")


async def default_handler(ctx, _):
    ctx.logger.info("[executing default handler]")
    emails = [dict_to_email(email) for email in ctx.get("emails")]
    res = Response()
    res.message = "Email processing in progress"
    await ctx.early_reply(res)
    etl_output = EtlOutput()
    etl_output.emails.extend(emails)
    await ctx.send_output(etl_output)
    return


def dict_to_email(raw_email: dict[str, str]) -> Email:
    email = Email()
    email.title = raw_email.get("title")
    email.body = raw_email.get("body")
    email.author = raw_email.get("author")
    email.creation_date = raw_email.get("date")

    return email
