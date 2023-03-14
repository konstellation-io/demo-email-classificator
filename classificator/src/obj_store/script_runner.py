from datetime import datetime

from nats.js.api import StorageType, ObjectMeta, ObjectMetaOptions, ObjectStoreConfig


async def default_handler(ctx, _):
    ctx.logger.info("[executing default handler]")

    # max_chunk_size has to be lower than or equal to the max_payload set by the NATS server -- increase it if you
    # set a custom max_payload in the NATS server config
    max_chunk_size = 1024 * 1024  # 1MB
    obj_name = "foo.txt"
    bucket_name = "foo"
    storage_type = StorageType.MEMORY
    num_gb = 1  # change this to generate larger (or smaller) payloads
    payload = b'x' * 1024 * 1024 * 1024 * num_gb

    ometaopts = ObjectMetaOptions(max_chunk_size=max_chunk_size)
    ometa = ObjectMeta(name=obj_name, options=ometaopts)
    oscfg = ObjectStoreConfig(bucket=bucket_name, storage=storage_type)

    os = await ctx.js.create_object_store(config=oscfg)

    now = datetime.now()
    await os.put(meta=ometa, data=payload)
    ctx.logger.info("put: ", datetime.now() - now)

    now = datetime.now()
    obj_info = await os.get(name=obj_name)
    ctx.logger.info("get: ", datetime.now() - now)
    ctx.logger.info("info: ", obj_info)

    now = datetime.now()
    await os.delete(name=obj_name)
    ctx.logger.info("delete: ", datetime.now() - now)

    now = datetime.now()
    obj_info = await os.get(name=obj_name)
    ctx.logger.info("get (after delete): ", datetime.now() - now)
    ctx.logger.info("info: ", obj_info)

    if not (await ctx.js.delete_object_store(bucket=bucket_name)):
        ctx.logger.warning("Object Store %s could not be deleted", bucket_name)
