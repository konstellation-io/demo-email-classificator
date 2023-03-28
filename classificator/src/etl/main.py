import csv

from proto.public_input_pb2 import ClassificatorRequest, Response
from proto.internal_nodes_pb2 import Email, EtlOutput


def init(ctx):
    ctx.logger.info("[executing init]")


async def default_handler(ctx, req):
    ctx.logger.info("[executing default handler]")

    data = ClassificatorRequest()
    req.Unpack(ClassificatorRequest)

    emails = []
    with open(ctx.path(data.filename)) as csvfile:
        rows = csv.DictReader(csvfile)
        for row in rows:
            emails.append(dict_to_email(row))

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
