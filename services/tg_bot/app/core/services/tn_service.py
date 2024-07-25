import asyncio

from google.protobuf import empty_pb2

from app.core.notificator.notificator import NotificationSender


class TN:
    def __init__(self):
        self.notification_sender = NotificationSender()

    def send_notification(self, request, context):
        asyncio.run(self.notification_sender.send_notification("слышь уебище у тя новый колбэк бля"))
        return empty_pb2.Empty()
