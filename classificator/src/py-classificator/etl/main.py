import csv

from proto.public_input_pb2 import Request, Response
from proto.internal_nodes_pb2 import Email, EtlOutput

from io import StringIO


def init(ctx):
    ctx.logger.info("[executing init]")


async def default_handler(ctx, data):
    ctx.logger.info("[executing default handler]")

    req = Request()
    data.Unpack(req)

    emails = []
    field_names = []
    with open(ctx.path(req.filename)) as csvfile:
        rows = csv.DictReader(csvfile)
        field_names = rows.fieldnames
        ctx.logger.info(field_names)
        for row in rows:
            emails.append(row)

    res = Response()
    res.message = f"Processing of {len(emails)} emails in progress"
    await ctx.send_early_reply(res)

    batch = StringIO()
    csv_writer = csv.DictWriter(batch, fieldnames=field_names)
    csv_writer.writeheader()
    count = 0
    batch_count = 0
    for email in emails:
        count += 1
        csv_writer.writerow(email)
        if count == req.batch_size:
            batch_key = f"emails_{ctx.get_request_id()}_{batch_count}"
            batch_count += 1

            await ctx.object_store.save(batch_key, bytes(batch.getvalue(), "utf-8"))

            etl_output = EtlOutput()
            etl_output.emails_key = batch_key
            await ctx.send_output(etl_output)

            emails = StringIO()
            csv.DictWriter(emails, field_names)
            csv_writer.writeheader()
            count = 0
    return

async def new_handler(ctx, req) -> None:
    ctx.logger.info("[executing new handler]")

    request = Request()
    req.Unpack(request)

    res = Response()
    res.message = f"Processing of {request.filename} emails in progress"
    await ctx.send_early_reply(res)

    etl_output = EtlOutput()
    etl_output.emails_key = request.filename

    await ctx.send_output(etl_output)

    return

custom_handlers = {
    "entrypoint": new_handler,
}