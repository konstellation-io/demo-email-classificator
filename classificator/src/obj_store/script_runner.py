from datetime import datetime

from proto.internal_nodes_pb2 import StatsStorerOutput

async def default_handler(ctx, _):
    ctx.logger.info("[executing default handler]")

    filename = "foo"
    payload = b'x' * int(1024 * 1024 * 1024 * 1.0)
    ctx.logger.info(f"chosen payload has {len(payload)} bytes")

    now = datetime.now()
    await ctx.object_store.save(filename, payload)
    ctx.logger.info(f"put: {datetime.now() - now}")

    output = StatsStorerOutput()
    output.message = filename
    await ctx.send_output(output)
