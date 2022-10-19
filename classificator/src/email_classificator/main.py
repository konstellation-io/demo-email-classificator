import random

from proto.internal_nodes_pb2 import (
    EtlOutput,
    EmailCategory,
    Email,
    CATEGORY_ADMINISTRATION,
    CATEGORY_REPARATIONS,
    CATEGORY_BILLING,
    CATEGORY_SPAM,
    ClassificatorOutput,
)


async def default_handler(ctx, req):
    if ctx.is_message_early_reply():
        return
    etl_output = EtlOutput()
    req.Unpack(etl_output)
    email = etl_output.email
    category = classify_email(email)
    res = ClassificatorOutput()
    res.email.CopyFrom(email)
    res.category = category
    if category == CATEGORY_REPARATIONS:
        await ctx.send_output(res, "repairs")
    await ctx.send_output(res)


def classify_email(_: Email) -> EmailCategory:
    categories = [CATEGORY_REPARATIONS, CATEGORY_ADMINISTRATION, CATEGORY_BILLING, CATEGORY_SPAM]
    return categories[random.randint(0, len(categories) - 1)]
