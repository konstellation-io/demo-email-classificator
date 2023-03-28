from datetime import datetime

async def default_handler(ctx, _):
    ctx.logger.info("[executing default handler]")

    payload = b'x' * int(1024 * 1024 * 1024 * 1.0)
    ctx.logger.info(f"chosen payload has {len(payload)} bytes")

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
