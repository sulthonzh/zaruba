from fastapi import FastAPI
from fastapi.staticfiles import StaticFiles
from sqlalchemy import create_engine
from helpers.transport import MessageBus, RMQMessageBus, RMQEventMap, LocalMessageBus


import os

def create_message_bus(mb_type: str) -> MessageBus:
    if mb_type == 'rmq':
        rmq_host = os.getenv('ZARUBA_SERVICE_NAME_RABBITMQ_HOST', 'localhost')
        rmq_user = os.getenv('ZARUBA_SERVICE_NAME_RABBITMQ_USER', 'root')
        rmq_pass = os.getenv('ZARUBA_SERVICE_NAME_RABBITMQ_PASS', 'toor')
        rmq_vhost = os.getenv('ZARUBA_SERVICE_NAME_RABBITMQ_VHOST', '/')
        rmq_event_map = RMQEventMap({})
        return RMQMessageBus(rmq_host, rmq_user, rmq_pass, rmq_vhost, rmq_event_map)
    return LocalMessageBus()

db_url = os.getenv('ZARUBA_SERVICE_NAME_SQLALCHEMY_DATABASE_URL', 'sqlite://')
mb_type = os.getenv('ZARUBA_SERVICE_NAME_MESSAGE_BUS_TYPE', 'local')
enable_route = os.getenv('ZARUBA_SERVICE_NAME_ENABLE_ROUTE_HANDLER', '1') != '0'
enable_event = os.getenv('ZARUBA_SERVICE_NAME_ENABLE_EVENT_HANDLER', '1') != '0'
static_url = os.getenv('ZARUBA_SERVICE_NAME_STATIC_URL', '/static')
static_directory = os.path.abspath(os.getenv('ZARUBA_SERVICE_NAME_STATIC_DIRECTORY')) if os.getenv('ZARUBA_SERVICE_NAME_STATIC_DIRECTORY') != '' else ''

engine = create_engine(db_url, echo=True)
app = FastAPI()
mb = create_message_bus(mb_type)

@app.on_event('shutdown')
def on_shutdown():
    mb.shutdown()
 
if static_directory != '':
    app.mount(static_url, StaticFiles(directory=static_directory), name='static')