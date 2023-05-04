import csv

from proto.public_input_pb2 import Request, Response
from proto.internal_nodes_pb2 import Email, EtlOutput

from io import StringIO


def init(ctx):
    ctx.logger.info("[executing init]")


async def default_handler(ctx, data):
    ctx.logger.info("[executing default handler]")

    # Unpack the request payload
    req = Request()
    data.Unpack(req)

    # Read the CSV file from the filename provided in the request and save the emails in a list
    emails = []
    with open(ctx.path(req.filename)) as csvfile:
        rows = csv.DictReader(csvfile)
        field_names = rows.fieldnames
        ctx.logger.info(field_names)
        for row in rows:
            emails.append(row)

    # Send an early reply to the client
    res = Response()
    res.message = f"Processing of {len(emails)} emails in progress"
    await ctx.send_early_reply(res)

    # Process the emails in batches and send the results to the next node using the object store
    batch = StringIO()
    csv_writer = csv.DictWriter(batch, fieldnames=field_names)
    csv_writer.writeheader()
    count = 0
    batch_count = 0
    for email in emails:
        count += 1
        csv_writer.writerow(email)

        # Send the batch to the object store when the batch size is reached
        if count == req.batch_size:
            batch_key = f"emails_{ctx.get_request_id()}_{batch_count}"
            batch_count += 1

            # Save the batch in the object store as a byte array using the batch key generated above
            await ctx.object_store.save(batch_key, bytes(batch.getvalue(), "utf-8"))

            # Send the batch key to the next node
            etl_output = EtlOutput()
            etl_output.emails_key = batch_key
            await ctx.send_output(etl_output)

            # Reset the batch
            emails = StringIO()
            csv.DictWriter(emails, field_names)
            csv_writer.writeheader()
            count = 0
    return
