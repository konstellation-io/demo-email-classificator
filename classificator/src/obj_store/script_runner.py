import os
from datetime import datetime

def init(ctx):
    ctx.logger.info("[executing init]")
    ctx.logger.info(f"OBJ STORE ENV VAR: {os.environ.get('KRT_NATS_OBJECT_STORE')}")
    ctx.logger.info(f"OBJECT STORE: {ctx.object_store}")

async def default_handler(ctx, _):
    ctx.logger.info("[executing default handler]")

    payload = b'x' * int(1024 * 1024 * 1024 * 1.0)
    now = datetime.now()
    await ctx.object_store.save("foo", payload)
    ctx.logger.info(f"put: {datetime.now() - now}")

    now = datetime.now()
    payload = await ctx.object_store.get("foo")
    ctx.logger.info(f"get: {datetime.now() - now}")
    ctx.logger.info(f"got payload with {len(payload)} bytes")

    now = datetime.now()
    await ctx.object_store.delete("foo")
    ctx.logger.info(f"delete: {datetime.now() - now}")
