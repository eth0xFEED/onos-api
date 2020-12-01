# Generated by the protocol buffer compiler.  DO NOT EDIT!
# sources: onos/e2sub/subscription/subscription.proto
# plugin: python-betterproto
from dataclasses import dataclass
from typing import AsyncIterator, List, Optional

import betterproto
import grpclib


class Status(betterproto.Enum):
    """Status is a subscription status"""

    ACTIVE = 0
    PENDING_DELETE = 1


class EventType(betterproto.Enum):
    """EventType is a subscription event type"""

    NONE = 0
    ADDED = 1
    UPDATED = 2
    REMOVED = 3


class Encoding(betterproto.Enum):
    """Encoding indicates a payload encoding"""

    ENCODING_ASN1 = 0
    ENCODING_PROTO = 1


@dataclass(eq=False, repr=False)
class Lifecycle(betterproto.Message):
    """Lifecycle is the subscription lifecycle"""

    status: "Status" = betterproto.enum_field(1)

    def __post_init__(self) -> None:
        super().__post_init__()


@dataclass(eq=False, repr=False)
class Event(betterproto.Message):
    """Event is a subscription event"""

    type: "EventType" = betterproto.enum_field(1)
    subscription: "Subscription" = betterproto.message_field(2)

    def __post_init__(self) -> None:
        super().__post_init__()


@dataclass(eq=False, repr=False)
class ServiceModel(betterproto.Message):
    """ServiceModel is a service model definition"""

    id: str = betterproto.string_field(4)

    def __post_init__(self) -> None:
        super().__post_init__()


@dataclass(eq=False, repr=False)
class Payload(betterproto.Message):
    """Payload is a subscription payload"""

    encoding: "Encoding" = betterproto.enum_field(1)
    bytes: bytes = betterproto.bytes_field(2)

    def __post_init__(self) -> None:
        super().__post_init__()


@dataclass(eq=False, repr=False)
class Subscription(betterproto.Message):
    """Subscription is a subscription state"""

    id: str = betterproto.string_field(1)
    revision: int = betterproto.uint64_field(2)
    app_id: str = betterproto.string_field(3)
    e2_node_id: str = betterproto.string_field(4)
    service_model: "ServiceModel" = betterproto.message_field(5)
    payload: "Payload" = betterproto.message_field(6)
    lifecycle: "Lifecycle" = betterproto.message_field(7)

    def __post_init__(self) -> None:
        super().__post_init__()


@dataclass(eq=False, repr=False)
class AddSubscriptionRequest(betterproto.Message):
    """AddSubscriptionRequest a subscription request"""

    subscription: "Subscription" = betterproto.message_field(1)

    def __post_init__(self) -> None:
        super().__post_init__()


@dataclass(eq=False, repr=False)
class AddSubscriptionResponse(betterproto.Message):
    """AddSubscriptionResponse a subscription response"""

    subscription: "Subscription" = betterproto.message_field(1)

    def __post_init__(self) -> None:
        super().__post_init__()


@dataclass(eq=False, repr=False)
class RemoveSubscriptionRequest(betterproto.Message):
    """RemoveSubscriptionRequest a subscription delete request"""

    id: str = betterproto.string_field(1)

    def __post_init__(self) -> None:
        super().__post_init__()


@dataclass(eq=False, repr=False)
class RemoveSubscriptionResponse(betterproto.Message):
    """RemoveSubscriptionResponse a subscription delete response"""

    pass

    def __post_init__(self) -> None:
        super().__post_init__()


@dataclass(eq=False, repr=False)
class GetSubscriptionRequest(betterproto.Message):
    id: str = betterproto.string_field(1)

    def __post_init__(self) -> None:
        super().__post_init__()


@dataclass(eq=False, repr=False)
class GetSubscriptionResponse(betterproto.Message):
    subscription: "Subscription" = betterproto.message_field(1)

    def __post_init__(self) -> None:
        super().__post_init__()


@dataclass(eq=False, repr=False)
class ListSubscriptionsRequest(betterproto.Message):
    pass

    def __post_init__(self) -> None:
        super().__post_init__()


@dataclass(eq=False, repr=False)
class ListSubscriptionsResponse(betterproto.Message):
    subscriptions: List["Subscription"] = betterproto.message_field(1)

    def __post_init__(self) -> None:
        super().__post_init__()


@dataclass(eq=False, repr=False)
class WatchSubscriptionsRequest(betterproto.Message):
    noreplay: bool = betterproto.bool_field(1)

    def __post_init__(self) -> None:
        super().__post_init__()


@dataclass(eq=False, repr=False)
class WatchSubscriptionsResponse(betterproto.Message):
    event: "Event" = betterproto.message_field(1)

    def __post_init__(self) -> None:
        super().__post_init__()


class E2SubscriptionServiceStub(betterproto.ServiceStub):
    """
    SubscriptionService manages subscription and subscription delete requests
    """

    async def add_subscription(
        self, *, subscription: "Subscription" = None
    ) -> "AddSubscriptionResponse":
        """AddSubscription establishes E2 subscriptions on E2 Node."""

        request = AddSubscriptionRequest()
        if subscription is not None:
            request.subscription = subscription

        return await self._unary_unary(
            "/onos.e2sub.subscription.E2SubscriptionService/AddSubscription",
            request,
            AddSubscriptionResponse,
        )

    async def remove_subscription(
        self, *, id: str = ""
    ) -> "RemoveSubscriptionResponse":
        """RemoveSubscription removes E2 subscriptions on E2 Node."""

        request = RemoveSubscriptionRequest()
        request.id = id

        return await self._unary_unary(
            "/onos.e2sub.subscription.E2SubscriptionService/RemoveSubscription",
            request,
            RemoveSubscriptionResponse,
        )

    async def get_subscription(self, *, id: str = "") -> "GetSubscriptionResponse":
        """
        GetSubscription retrieves information about a specific subscription in
        the list of existing subscriptions
        """

        request = GetSubscriptionRequest()
        request.id = id

        return await self._unary_unary(
            "/onos.e2sub.subscription.E2SubscriptionService/GetSubscription",
            request,
            GetSubscriptionResponse,
        )

    async def list_subscriptions(self) -> "ListSubscriptionsResponse":
        """
        ListSubscriptions returns the list of current existing subscriptions
        """

        request = ListSubscriptionsRequest()

        return await self._unary_unary(
            "/onos.e2sub.subscription.E2SubscriptionService/ListSubscriptions",
            request,
            ListSubscriptionsResponse,
        )

    async def watch_subscriptions(
        self, *, noreplay: bool = False
    ) -> AsyncIterator["WatchSubscriptionsResponse"]:
        """WatchSubscriptions returns a stream of subscription changes"""

        request = WatchSubscriptionsRequest()
        request.noreplay = noreplay

        async for response in self._unary_stream(
            "/onos.e2sub.subscription.E2SubscriptionService/WatchSubscriptions",
            request,
            WatchSubscriptionsResponse,
        ):
            yield response
