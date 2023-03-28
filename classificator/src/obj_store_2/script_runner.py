from datetime import datetime

from proto.internal_nodes_pb2 import StatsStorerOutput

async def default_handler(ctx, req):
    ctx.logger.info("[executing default handler]")

    obj_store_output = StatsStorerOutput()
    req.Unpack(obj_store_output)

    filename = obj_store_output.message

    now = datetime.now()
    payload = await ctx.object_store.get(filename)
    ctx.logger.info(f"get: {datetime.now() - now}")
    ctx.logger.info(f"got payload with {len(payload)} bytes")

    now = datetime.now()
    await ctx.object_store.delete(filename)
    ctx.logger.info(f"delete: {datetime.now() - now}")
