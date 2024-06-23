import asyncio
from app.api.server import serve

if __name__ == '__main__':
    asyncio.run(serve())