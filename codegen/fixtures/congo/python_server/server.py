# DO NOT EDIT THIS FILE. This file will be overwritten when re-running go-raml.

import signal

import gevent
from app import app
from gevent import monkey
from gevent.pool import Pool
from gevent.pywsgi import WSGIServer

monkey.patch_all()


server = WSGIServer(("", 5000), app, spawn=Pool(None))


def stop():
    server.stop()


gevent.signal(signal.SIGINT, stop)


if __name__ == "__main__":
    server.serve_forever()
